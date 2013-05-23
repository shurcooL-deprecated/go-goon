package goon

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"go/parser"
	"go/printer"
	"go/token"
	"go/ast"
	"go/format"

	"os/exec"

	//. "gist.github.com/5258650.git"
	//"runtime/debug"

	. "gist.github.com/5286084.git"
)

var _ ast.Ident

//var _ = debug.Stack()
//var _ = GetLines("")

// dumpState contains information about the state of a dump operation.
type dumpState struct {
	w                io.Writer
	depth            int
	pointers         map[uintptr]int
	ignoreNextType   bool
	ignoreNextIndent bool
	cs               *ConfigState
}

// indent performs indentation according to the depth level and cs.Indent
// option.
func (d *dumpState) indent() {
	if d.ignoreNextIndent {
		d.ignoreNextIndent = false
		return
	}
	d.w.Write(bytes.Repeat([]byte(d.cs.Indent), d.depth))
}

// unpackValue returns values inside of non-nil interfaces when possible.
// This is useful for data types like structs, arrays, slices, and maps which
// can contain varying types packed inside an interface.
func (d *dumpState) unpackValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface && !v.IsNil() {
		v = v.Elem()
	}
	return v
}

// dumpPtr handles formatting of pointers by indirecting them as necessary.
func (d *dumpState) dumpPtr(v reflect.Value) {
	// Remove pointers at or below the current depth from map used to detect
	// circular refs.
	for k, depth := range d.pointers {
		if depth >= d.depth {
			delete(d.pointers, k)
		}
	}

	// Keep list of all dereferenced pointers to show later.
	pointerChain := make([]uintptr, 0)

	// Figure out how many levels of indirection there are by dereferencing
	// pointers and unpacking interfaces down the chain while detecting circular
	// references.
	nilFound := false
	cycleFound := false
	indirects := 0
	ve := v
	for ve.Kind() == reflect.Ptr {
		if ve.IsNil() {
			nilFound = true
			break
		}
		indirects++
		addr := ve.Pointer()
		pointerChain = append(pointerChain, addr)
		if pd, ok := d.pointers[addr]; ok && pd < d.depth {
			cycleFound = true
			indirects--
			break
		}
		d.pointers[addr] = d.depth

		ve = ve.Elem()
		if ve.Kind() == reflect.Interface {
			if ve.IsNil() {
				nilFound = true
				break
			}
			ve = ve.Elem()
		}
	}

	// Display type information.
	d.w.Write(bytes.Repeat(ampersandBytes, indirects))

	// Display dereferenced value.
	switch {
	case nilFound == true:
		d.w.Write(nilBytes)

	case cycleFound == true:
		d.w.Write(circularBytes)

	default:
		d.ignoreNextType = true
		d.dump(ve)
	}
}

// dumpSlice handles formatting of arrays and slices.  Byte (uint8 under
// reflection) arrays and slices are dumped in hexdump -C fashion.
func (d *dumpState) dumpSlice(v reflect.Value) {
	// Handle byte (uint8 under reflection) arrays and slices uniquely.
	numEntries := v.Len()
	if (numEntries > 0) && (v.Index(0).Kind() == reflect.Uint8) {
		// We need an addressable interface to convert the type back into a byte
		// slice.  However, the reflect package won't give us an interface on
		// certain things like unexported struct fields in order to enforce
		// visibility rules.  We use unsafe to bypass these restrictions since
		// this package does not mutate the values.
		vs := v
		if !vs.CanInterface() || !vs.CanAddr() {
			vs = unsafeReflectValue(vs)
		}
		vs = vs.Slice(0, numEntries)

		// Type assert a uint8 slice and hexdump it.  Also fix indentation
		// based on the depth.
		iface := vs.Interface()
		if buf, ok := iface.([]uint8); ok {
			indent := strings.Repeat(d.cs.Indent, d.depth)
			str := indent + hex.Dump(buf)
			str = strings.Replace(str, "\n", "\n"+indent, -1)
			str = strings.TrimRight(str, d.cs.Indent)
			d.w.Write([]byte(str))
			return
		}
		// We shouldn't ever get here, but the return is intentionally in the
		// above if statement to ensure we fall through to normal behavior if
		// the type assertion fails for some reason.
	}

	// Recursively call dump for each item.
	for i := 0; i < numEntries; i++ {
		d.dump(d.unpackValue(v.Index(i)))
		d.w.Write(commaNewlineBytes)
	}
}

func IsZeroValue(v reflect.Value) bool {
	if !v.CanInterface() || !reflect.Zero(v.Type()).CanInterface()/* || reflect.Slice == v.Kind()*/ {
		return false
	}
	return reflect.Zero(v.Type()).Interface() == v.Interface()
}

// dump is the main workhorse for dumping a value.  It uses the passed reflect
// value to figure out what kind of object we are dealing with and formats it
// appropriately.  It is a recursive function, however circular data structures
// are detected and handled properly.
func (d *dumpState) dump(v reflect.Value) {
	// Handle invalid reflect values immediately.
	kind := v.Kind()
	if kind == reflect.Invalid {
		d.w.Write(invalidAngleBytes)
		return
	}

	// Handle pointers specially.
	if kind == reflect.Ptr {
		d.indent()
		d.w.Write(openParenBytes)
		d.w.Write([]byte(TypeStringWithoutPackagePrefix(v)))
		d.w.Write(closeParenBytes)
		d.w.Write(openParenBytes)
		d.dumpPtr(v)
		d.w.Write(closeParenBytes)
		return
	}

	// Print type information unless already handled elsewhere.
	var shouldPrintClosingBr bool = false
	if !d.ignoreNextType {
		d.indent()
		d.w.Write(openParenBytes)
		d.w.Write([]byte(TypeStringWithoutPackagePrefix(v)))
		d.w.Write(closeParenBytes)
		d.w.Write(openParenBytes)
		shouldPrintClosingBr = true
	}
	d.ignoreNextType = false

	// Call Stringer/error interfaces if they exist and the handle methods flag
	// is enabled
	if !d.cs.DisableMethods {
		if (kind != reflect.Invalid) && (kind != reflect.Interface) {
			if handled := handleMethods(d.cs, d.w, v); handled {
				return
			}
		}
	}

	switch kind {
	case reflect.Invalid:
		// Do nothing.  We should never get here since invalid has already
		// been handled above.

	case reflect.Bool:
		printBool(d.w, v.Bool())

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		printInt(d.w, v.Int(), 10)

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		printUint(d.w, v.Uint(), 10)

	case reflect.Float32:
		printFloat(d.w, v.Float(), 32)

	case reflect.Float64:
		printFloat(d.w, v.Float(), 64)

	case reflect.Complex64:
		printComplex(d.w, v.Complex(), 32)

	case reflect.Complex128:
		printComplex(d.w, v.Complex(), 64)

	case reflect.Array, reflect.Slice:
		d.w.Write([]byte(TypeStringWithoutPackagePrefix(v)))
		d.w.Write(openBraceNewlineBytes)
		d.depth++
		if (d.cs.MaxDepth != 0) && (d.depth > d.cs.MaxDepth) {
			d.indent()
			d.w.Write(maxNewlineBytes)
		} else {
			d.dumpSlice(v)
		}
		d.depth--
		d.indent()
		d.w.Write(closeBraceBytes)

	case reflect.String:
		d.w.Write([]byte(strconv.Quote(v.String())))

	case reflect.Interface:
		// If we got here, it's because interface is nil
		// See https://github.com/davecgh/go-spew/issues/12
		d.w.Write(nilBytes)

	case reflect.Ptr:
		// Do nothing.  We should never get here since pointers have already
		// been handled above.

	case reflect.Map:
		d.w.Write([]byte(TypeStringWithoutPackagePrefix(v)))
		d.w.Write(openBraceNewlineBytes)
		d.depth++
		if (d.cs.MaxDepth != 0) && (d.depth > d.cs.MaxDepth) {
			d.indent()
			d.w.Write(maxNewlineBytes)
		} else {
			keys := v.MapKeys()
			for _, key := range keys {
				d.dump(d.unpackValue(key))
				d.w.Write(colonSpaceBytes)
				d.ignoreNextIndent = true
				d.dump(d.unpackValue(v.MapIndex(key)))
				d.w.Write(commaNewlineBytes)
			}
		}
		d.depth--
		d.indent()
		d.w.Write(closeBraceBytes)

	case reflect.Struct:
		d.w.Write([]byte(TypeStringWithoutPackagePrefix(v)))
		d.w.Write(openBraceBytes)
		d.depth++
		if (d.cs.MaxDepth != 0) && (d.depth > d.cs.MaxDepth) {
			d.indent()
			d.w.Write(maxNewlineBytes)
		} else {
			vt := v.Type()
			numFields := v.NumField()
			if numFields > 0 {
				d.w.Write(newlineBytes)
			}
			for i := 0; i < numFields; i++ {
				//if !IsZeroValue(d.unpackValue(v.Field(i))) {
				if true {
					d.indent()
					vtf := vt.Field(i)
					d.w.Write([]byte(vtf.Name))
					d.w.Write(colonSpaceBytes)
					d.ignoreNextIndent = true
					d.dump(d.unpackValue(v.Field(i)))
					d.w.Write(commaBytes)
					d.w.Write(newlineBytes)
					/*d.w.Write([]byte("\t// "))
					d.w.Write(openParenBytes)
					d.w.Write([]byte(d.unpackValue(v.Field(i)).Type().String()))
					d.w.Write(closeParenBytes)*/
				}
			}
		}
		d.depth--
		d.indent()
		d.w.Write(closeBraceBytes)

	case reflect.Uintptr:
		printHexPtr(d.w, uintptr(v.Uint()))

	case reflect.UnsafePointer, reflect.Chan, reflect.Func:
		printHexPtr(d.w, v.Pointer())

	// There were not any other types at the time this code was written, but
	// fall back to letting the default fmt package handle it in case any new
	// types are added.
	default:
		if v.CanInterface() {
			fmt.Fprintf(d.w, "%v", v.Interface())
		} else {
			fmt.Fprintf(d.w, "%v", v.String())
		}
	}

	if shouldPrintClosingBr {
		d.w.Write(closeParenBytes)
	}
}

func TypeStringWithoutPackagePrefix(v reflect.Value) string {
	//return v.Type().String()[len(v.Type().PkgPath())+1:]		// TODO: Error checking?
	//return v.Type().PkgPath()
	//return v.Type().String()
	//return v.Type().Name()

	/*x := v.Type().String()
	if strings.HasPrefix(x, "main.") {
		x = x[len("main."):]
	}
	return x*/

	px := v.Type().String()
	prefix := px[0 : len(px)-len(strings.TrimLeft(px, "*"))]		// Split "**main.Lang" -> "**" and "main.Lang"
	x := px[len(prefix):]
	if strings.HasPrefix(x, "main.") {
		x = x[len("main."):]
	}
	return prefix + x

	/*x = string(debug.Stack())//GetLine(string(debug.Stack()), 0)
	//x = x[1:strings.Index(x, ":")]
	//spew.Printf(">%s<\n", x)
	//panic(nil)
	//st := string(debug.Stack())
	//debug.PrintStack()

	return x*/
}

// fdump is a helper function to consolidate the logic from the various public
// methods which take varying writers and config states.
func fdump(cs *ConfigState, w io.Writer, a ...interface{}) {
	for _, arg := range a {
		if arg == nil {
			w.Write(interfaceBytes)
			w.Write(nilParenBytes)
			w.Write(newlineBytes)
			continue
		}

		d := dumpState{w: w, cs: cs}
		d.pointers = make(map[uintptr]int)
		d.dump(reflect.ValueOf(arg))
		d.w.Write(newlineBytes)
	}
}

// Dumps to []byte
func bdump(a ...interface{}) []byte {
	var buf bytes.Buffer
	fdump(&config, &buf, a...)
	return Gofmt3(buf.String())
}

// Dumps to string
func Sdump(a ...interface{}) string {
	return string(bdump(a...))
}

// Dumps to stdout
func Dump(a ...interface{}) {
	os.Stdout.Write(bdump(a...))
}

// Noop
func Gofmt0(str string) []byte {
	return []byte(str)
}

// TODO: Replace with go1.1's go/format
func Gofmt1(str string) []byte {
	if expr, err := parser.ParseExpr(str); nil == err {
		var buf bytes.Buffer
		// This loses the formatting spacing information due to NewFileSet
		printer.Fprint(&buf, token.NewFileSet(), expr)
		return buf.Bytes()
	}
	return nil
}

// TODO: Replace with go1.1's go/format
// Mimics gofmt's default internal behaviour
func Gofmt2(x string) []byte {
	fset := token.NewFileSet()
	// Ok I give up basically reimplementing private code of gofmt here, useless work cuz go1.1 will have go/format
	// So I'll just use gofmt binary for now
	if file, err := parser.ParseFile(fset, "", "package p; func _() {" + x + "}", parser.ParseComments); nil == err {
		var buf bytes.Buffer
		// The following printer.Config tries to mimic the (current) default gofmt behaviour
		(&printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}).Fprint(&buf, fset, file)
		buf.Write(newlineBytes)
		return buf.Bytes()
	} else {
		panic(err)
	}
	return []byte("gofmt error!\n" + x)
}

// TODO: Replace with go1.1's go/format
// Actually executes gofmt binary as a new process
func Gofmt3(str string) []byte {
	// TODO: The gofmt path is hardcoded, this is bad
	cmd := exec.Command("/usr/local/go/bin/gofmt")

	// TODO: Error checking and other niceness
	// http://stackoverflow.com/questions/13432947/exec-external-program-script-and-detect-if-it-requests-user-input
	in, err := cmd.StdinPipe()
	CheckError(err)
	go func() {
		_, err = in.Write([]byte(str))
		CheckError(err)
		err = in.Close()
		CheckError(err)
	}()

	data, err := cmd.Output()
	if nil != err {
		return []byte("gofmt error!\n" + str)
	}
	return data
}

// TODO: Can't use it until go/format is fixed to be consistent with gofmt, currently it strips comments out of partial Go programs
func Gofmt4(str string) []byte {
	formattedSrc, err := format.Source([]byte(str))
	if nil != err {
		return []byte("gofmt error (" + err.Error() + ")!\n" + str)
	}
	return formattedSrc
}