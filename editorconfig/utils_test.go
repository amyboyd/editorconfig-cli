package editorconfig

import (
	"reflect"
	"strconv"
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

func TestMustGetFileAsString(t *testing.T) {
	license := MustGetFileAsString("../LICENSE")
	if !strings.Contains(license, "MIT License") || !strings.Contains(license, "THE SOFTWARE IS PROVIDED \"AS IS\"") {
		t.Error("Could not read file")
	}
}

func TestGetNumberOfLeftSpaces(t *testing.T) {
	for i := 0; i < 20; i++ {
		if GetNumberOfLeftSpaces(strings.Repeat(" ", i)) != i {
			t.Error("Wrong number of spaces returned when string starts with " + strconv.Itoa(i) + " spaces")
		}
	}
}
