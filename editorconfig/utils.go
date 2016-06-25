package editorconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var filePathSeparatorRegex = regexp.QuoteMeta(string(filepath.Separator))

var endOfPathRegex, _ = regexp.Compile(`[` + filePathSeparatorRegex + `][^` + filePathSeparatorRegex + `]+$`)

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
