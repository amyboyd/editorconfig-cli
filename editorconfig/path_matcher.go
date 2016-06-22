package editorconfig

/*
 * Converts a string from a .editorconfig file to a Go-compatible regex, according to the rules
 * documented under "Wildcard Patterns" here:
 * http://docs.editorconfig.org/en/master/editorconfig-format.html#patterns
 *
 * *	Matches any string of characters, except path separators (/)
 * **	Matches any string of characters
 * ?	Matches any single character
 * [seq]	Matches any single character in seq
 * [!seq]	Matches any single character not in seq
 * {s1,s2,s3}	Matches any of the strings given (separated by commas, can be nested)
 * {num1..num2}	Matches any integer numbers between num1 and num2, where num1 and num2 can be either positive or negative
 */

import (
	"regexp"
	"strconv"
	"strings"
)

var metaCharsRegexp = []*regexp.Regexp{
	// Characters .+$^!()
	regexp.MustCompile(`([\.\+\$\^\!\(\)])`),
	// [ not followed by a ]
	regexp.MustCompile(`(\[[^\]]*)$`),
	// ] not preceded by a [
	regexp.MustCompile(`^([^\[]*\])`),
}

func ConvertWildcardPatternToGoRegexp(pattern string) *regexp.Regexp {
	if pattern == "*" {
		// A single * seems to be universally used to mean every file, despite the official docs
		// showing that ** is correct and * should only match top-level files. However, we adapt
		// to what is used in the real world to be practical.
		return regexp.MustCompile(".")
	}

	originalPattern := pattern

	for _, r := range metaCharsRegexp {
		pattern = r.ReplaceAllString(pattern, "\\$1")
	}

	// Handle **
	pattern = strings.Replace(pattern, `**`, `.+`, -1)

	// Handle *
	pattern = strings.Replace(pattern, `*`, `[^/\\]+`, -1)

	// Handle ?
	pattern = strings.Replace(pattern, `?`, `.`, -1)

	// [seq] should work already.

	// Handle [!seq]
	pattern = strings.Replace(pattern, `[\!`, `[^`, -1)

	// Handle {s1,s2,s3}
	for i := 1; i < 7; i++ {
		find := `\{([^,}]+)` + strings.Repeat(`,([^,}]+)`, i) + `\}`
		replace := "($1"
		for ii := 1; ii <= i; ii++ {
			replace += "|$" + strconv.Itoa(ii+1)
		}
		replace += ")"
		pattern = regexp.MustCompile(find).ReplaceAllString(pattern, replace)
	}

	// Handle {num1..num2}
	// @todo - This is currently not fully supported. If there is a numeric range, we only check
	// that numbers are present; we don't check if the numbers present are within the correct range.
	pattern = regexp.MustCompile(`\{-?\d+\\.\\.-?\d+\}`).ReplaceAllString(pattern, `[-0-9]+`)

	r, err := regexp.Compile(pattern)
	if err != nil {
		ExitBecauseOfInternalError("A file pattern could not be parsed: " + originalPattern)
	}

	return r
}
