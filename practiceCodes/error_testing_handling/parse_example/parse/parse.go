package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
	Index int    // The index into the space-separated list of words.
	Word  string // The word that generated the parse error.
	Err   error  // The raw error that precipitated this error, if any.
}

// String returns a human-readable error message.
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	// fmt.Println("Hi")
	// fmt.Println(fields)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("Here")
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}

/*
This fields2numbers function is defined at line 37.
It iterates over the slice fields, at line 41, and
converts each field to a number num at line 42. If strconv.Atoi
results in an error because the field is not a string,
this is handled with an if condition ( see the implementation
	from line 43 and line 45), causing panic and displaying a
	ParseError with the detailed info of the problem. If everything
	is ok, num is appended to the slice nums at line 46 and returned.

The Parse function starts with a defer of an anonymous function
call (implemented from line 22 to line 30). This tries to recover
 from any panic that has happened and returns the error that occurred
 as err.

In main.go, starting at line 9, we use the parse package we
discussed above. We build a slice of strings to be parsed.
A for-loop ( from line 17 to line 25), iterates over this slice,
applying Parse to each slice ex at line 19.
The slice of integers nums that is returned is
printed at line 24, unless there is an error which
is handled (from line 20 to line 23).
This produces the output you can see in the terminal.

*/
