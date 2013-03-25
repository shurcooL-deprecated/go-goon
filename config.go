package goon

import (
)

// ConfigState houses the configuration options used by spew to format and
// display values.  There is a global instance, Config, that is used to control
// all top-level Formatter and Dump functionality.  Each ConfigState instance
// provides methods equivalent to the top-level functions.
//
// The zero value for ConfigState provides no indentation.  You would typically
// want to set it to a space or a tab.
type ConfigState struct {
	// Indent specifies the string to use for each indentation level.  The
	// global config instance that all top-level functions use set this to a
	// single space by default.  If you would like more indentation, you might
	// set this to a tab with "\t" or perhaps two spaces with "  ".
	Indent string

	// MaxDepth controls the maximum number of levels to descend into nested
	// data structures.  The default, 0, means there is no limit.
	//
	// NOTE: Circular data structures are properly detected, so it is not
	// necessary to set this value unless you specifically want to limit deeply
	// nested data structures.
	MaxDepth int

	// DisableMethods specifies whether or not error and Stringer interfaces are
	// invoked for types that implement them.
	DisableMethods bool

	// DisablePointerMethods specifies whether or not to check for and invoke
	// error and Stringer interfaces on types which only accept a pointer
	// receiver when the current type is not a pointer.
	//
	// NOTE: This might be an unsafe action since calling one of these methods
	// with a pointer receiver could technically mutate the value, however,
	// in practice, types which choose to satisify an error or Stringer
	// interface with a pointer receiver should not be mutating their state
	// inside these interface methods.
	DisablePointerMethods bool

	// ContinueOnMethod specifies whether or not recursion should continue once
	// a custom error or Stringer interface is invoked.  The default, false,
	// means it will print the results of invoking the custom error or Stringer
	// interface and return immediately instead of continuing to recurse into
	// the internals of the data type.
	//
	// NOTE: This flag does not have any effect if method invocation is disabled
	// via the DisableMethods or DisablePointerMethods options.
	ContinueOnMethod bool
}

// Config is the active configuration of the top-level functions.
// The configuration can be changed by modifying the contents of spew.Config.
var Config ConfigState = ConfigState{Indent: "\t"}
