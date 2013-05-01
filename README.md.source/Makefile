SHELL := /bin/bash

../README.md: * ../tests/*.go
	cat 1.txt \
		<(go run ../tests/small_example.go) \
		2.txt \
		<(go run ../tests/complete_example.go) \
		3.txt \
		<(go run ../tests/large_example.go) \
		4.txt \
		> ../README.md
