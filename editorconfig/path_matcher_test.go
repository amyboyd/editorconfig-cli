package editorconfig

import (
	"testing"
)

func TestConvertWildcardPatternToGoRegexp(t *testing.T) {
	RunConvertWildcardPatternToGoRegexp := func(input string, expected string) {
		result := ConvertWildcardPatternToGoRegexp(input).String()
		if result != expected {
			t.Error("For " + input + " got " + result + " but expected " + expected)
		}
	}

	// Test that special regex characters are escaped.
	RunConvertWildcardPatternToGoRegexp(`[hello :).sql`, `\[hello :\)\.sql`)
	RunConvertWildcardPatternToGoRegexp(`]hello(`, `\]hello\(`)

	// Test *
	RunConvertWildcardPatternToGoRegexp(`*.go`, `[^/\\]+\.go`)
	RunConvertWildcardPatternToGoRegexp(`folder/*.go`, `folder/[^/\\]+\.go`)

	// Test **
	RunConvertWildcardPatternToGoRegexp(`**.go`, `.+\.go`)
	RunConvertWildcardPatternToGoRegexp(`folder (copy)/**.go`, `folder \(copy\)/.+\.go`)

	// Test ?
	RunConvertWildcardPatternToGoRegexp(`words-beginning-with-?.txt`, `words-beginning-with-.\.txt`)
	RunConvertWildcardPatternToGoRegexp(`words beginning with ??.txt`, `words beginning with ..\.txt`)

	// Test [seq]
	RunConvertWildcardPatternToGoRegexp(`hexadecimal-ids/[a-f0-9]/document`, `hexadecimal-ids/[a-f0-9]/document`)

	// Test [!seq]
	RunConvertWildcardPatternToGoRegexp(`names-not-beginning-with-[!A-G]`, `names-not-beginning-with-[^A-G]`)

	// Test {s1,s2,s3}
	RunConvertWildcardPatternToGoRegexp(`animals/{aardvark,bunny}/pictures`, `animals/(aardvark|bunny)/pictures`)
	RunConvertWildcardPatternToGoRegexp(`animals/{aardvark,bunny,cheetah,donkey}/pictures`, `animals/(aardvark|bunny|cheetah|donkey)/pictures`)

	// Test {num1..num2} (this is not yet fully supported)
	RunConvertWildcardPatternToGoRegexp(`photos/{-500..999}.jpg`, `photos/[-0-9]+\.jpg`)

	// Test everything together.
	RunConvertWildcardPatternToGoRegexp(`*/**/[a-z]/{photos,videos}/{0..5}.*`, `[^/\\]+/.+/[a-z]/(photos|videos)/[-0-9]+\.[^/\\]+`)
}
