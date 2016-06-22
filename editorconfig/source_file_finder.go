package editorconfig

import (
	"bytes"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var containsDotDirectoryRegex = regexp.MustCompile(`(^|/)\.[^/]+/`)

var fileExtensions = GetSourceFileExtensions()

var fileExtensionsRegex = CreateSourceFileExtensionRegex(fileExtensions)

func GetSourceFileExtensions() []string {
	var file, err = Asset("data/file-extensions.txt")
	if err != nil {
		ExitBecauseOfInternalError("Could not open list of file extensions")
	}

	reader := bytes.NewBuffer(file)
	exts := []string{}

	for {
		ext, _ := reader.ReadBytes(' ')
		if len(ext) == 0 {
			break
		}
		ext = ext[0 : len(ext)-1]
		exts = append(exts, string(ext))

		// Skip the rest of the line (a description of what the file extension means).
		reader.ReadBytes('\n')
	}

	if len(exts) == 0 {
		ExitBecauseOfInternalError("No file extensions found")
	}

	return exts
}

func CreateSourceFileExtensionRegex(fileExtensions []string) *regexp.Regexp {
	for i, e := range fileExtensions {
		fileExtensions[i] = regexp.QuoteMeta(e)
	}

	r, err := regexp.Compile("\\.(" + strings.Join(fileExtensions, "|") + ")$")
	if err != nil {
		ExitBecauseOfInternalError("Could not compile regex: " + err.Error())
	}

	return r
}

func FindSourceFiles(searchPaths []string) ([]string, error) {
	files := []string{}

	for _, searchPath := range searchPaths {
		if searchPath == "/" {
			errMessage := "The path / was given.\n" +
				"Exiting because this could be a mistake," +
				" e.g. an environment variable infront of / might be unintentionally empty.\n" +
				"Here is an example of why we won't run on /:\n" +
				"http://www.theregister.co.uk/2015/01/17/scary_code_of_the_week_steam_cleans_linux_pcs/"
			return nil, cli.NewExitError(errMessage, 2)
		}

		_ = filepath.Walk(searchPath, func(path string, fileInfo os.FileInfo, err error) error {
			// Don't add paths that don't exist.
			if err != nil {
				return nil
			}

			// Only add files, not directories.
			if fileInfo.IsDir() {
				return nil
			}

			// Don't run on files inside dot directories. These directories are likely to
			// be source control directories like ".git" or ".svn", or temporary directories
			// like ".vagrant" or ".idea".
			if containsDotDirectoryRegex.MatchString(path) {
				return nil
			}

			// Only keep files with extensions that we know are textual.
			if !fileExtensionsRegex.MatchString(path) {
				return nil
			}

			files = append(files, path)

			return nil
		})
	}

	return files, nil
}
