package editorconfig

/*
 * Converts a string from a .editorconfig file to a Go-compatible regex. Refer to the documentation
 * in `docs/file-pattern-matchers.md`
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
		// As documented in `docs/file-pattern-matchers.md`, we deviate from the official
		// documentation here and make * match every file.
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
