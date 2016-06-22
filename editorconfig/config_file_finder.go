package editorconfig

import (
	"os"
	"path/filepath"
	"strings"
)

var configFiles = []ConfigFile{}

var checkedDirectories = []string{}

func FindConfigFiles(sourceFilePaths []string) []ConfigFile {
	configFiles = []ConfigFile{}

	for _, path := range sourceFilePaths {
		// Convert to absolute path so we can go all the way up the file system path looking for configs.
		path, err := filepath.Abs(path)
		if err != nil {
			ExitBecauseOfInternalError("Could not get absolute path for " + path)
		}

		path = GetParentDir(path)

		var searchParents bool
		for strings.Contains(path, "/") {
			searchParents = CheckDirectoryForConfigFile(path)
			if !searchParents {
				break
			}
			path = GetParentDir(path)
		}

		if searchParents {
			// `path` is now the volume root without a trailing slash, e.g. '' for '/' or 'C:' for 'C:/'.
			CheckDirectoryForConfigFile(path + "/")
		}
	}

	return configFiles
}

/**
 * @return Whether or not to search parent directories.
 */
func CheckDirectoryForConfigFile(dir string) bool {
	for _, cd := range checkedDirectories {
		if dir == cd {
			return false
		}
	}

	stat, err := os.Stat(dir)
	if err != nil {
		ExitBecauseOfInternalError("Could not stat directory: " + dir)
	}
	if !stat.IsDir() {
		ExitBecauseOfInternalError("Expected this path to be a directory, but it is not: " + dir)
	}

	checkedDirectories = append(checkedDirectories, dir)

	possibleFile := GetConfigFilePathInDirectory(dir)
	stat, err = os.Stat(possibleFile)
	if err != nil {
		return true
	}
	if stat.IsDir() {
		return true
	}
	if stat.Size() == 0 {
		return true
	}

	cf := CreateConfigFileStruct(possibleFile)
	configFiles = append(configFiles, cf)

	return !cf.IsRoot()
}

func GetConfigFilePathInDirectory(dir string) string {
	if strings.HasSuffix(dir, "/") {
		// This will only be possible in Unix for the '/' directory.
		return dir + ".editorconfig"
	} else {
		return dir + "/.editorconfig"
	}
}
