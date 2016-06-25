package editorconfig

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetParentDir(t *testing.T) {
	if GetParentDir("a/b/c.txt") != "a/b" {
		t.Error()
	}
	if GetParentDir("/") != "/" {
		t.Error()
	}
	if GetParentDir("/a/b") != "/a" {
		t.Error()
	}
}

func TestSplitIntoLines(t *testing.T) {
	result := SplitIntoLines("Aardvark\nBunny\rCat\r\nDolphin\n")
	expected := []string{"Aardvark", "Bunny", "Cat", "Dolphin", ""}
	if !reflect.DeepEqual(result, expected) {
		t.Error("Did not split string into lines correctly, got lines: " + strings.Join(result, ", "))
	}
}
