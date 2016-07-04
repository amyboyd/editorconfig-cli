package editorconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var filePathSeparatorRegex = regexp.QuoteMeta(string(filepath.Separator))

var endOfPathRegex, _ = regexp.Compile(`[` + filePathSeparatorRegex + `][^` + filePathSeparatorRegex + `]+$`)

var lfRegexp = regexp.MustCompile(`\n`)
var crRegexp = regexp.MustCompile(`\r`)
var crlfRegexp = regexp.MustCompile(`\r\n`)

var endsWithFinalNewLineRegexp = regexp.MustCompile(`(\n|\r|\r\n)$`)

var hasIndentationRegexp = regexp.MustCompile(`^[\t ]`)
var hasNoIndentationRegexp = regexp.MustCompile(`^([^\t ]|$)`)
var indentedWithMixedTabsAndSpacesRegexp = regexp.MustCompile(`^(\t+ +| +\t+)`)
var indentedWithTabsRegexp = regexp.MustCompile(`^\t+`)
var indentedWithTabsThenCommentLineRegexp = regexp.MustCompile(`^\t+ \*`)
var indentedWithSpacesRegexp = regexp.MustCompile(`^ +`)

func GetParentDir(path string) string {
	return endOfPathRegex.ReplaceAllString(path, "")
}

func ContainsString(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

var lineEndingsRegexp = regexp.MustCompile("(\r\n|\n|\r)")

func SplitIntoLines(s string) []string {
	return lineEndingsRegexp.Split(s, -1)
}

func ExitBecauseOfInternalError(err string) {
	fmt.Println(err)
	os.Exit(2)
}
