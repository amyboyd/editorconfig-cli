package editorconfig

import (
	"path/filepath"
	"sort"
	"strings"
)

func GetRulesToApplyToSourcePath(sourcePath string, cfs []ConfigFile) map[string]string {
	applicable := FilterConfigFilesToApplyToSourcePath(sourcePath, cfs)
	applicableSorted := SortConfigFilesByPrecendence(applicable)

	rules := make(map[string]string)

	for _, cf := range applicableSorted {
		for _, rule := range cf.DefaultRuleSet {
			rules[rule.Name] = rule.Value
		}

		for _, rs := range cf.FileConstrainedRuleSets {
			if rs.ConstraintRegexp.MatchString(sourcePath) {
				for _, rule := range rs.Rules {
					rules[rule.Name] = rule.Value
				}
			}
		}
	}

	if isIgnored, _ := rules["ignore"]; isIgnored == "true" {
		return make(map[string]string)
	}

	delete(rules, "root")

	if indentStyleValue, _ := rules["indent_style"]; indentStyleValue == "tab" {
		delete(rules, "indent_size")
	}

	return rules
}

func FilterConfigFilesToApplyToSourcePath(sourcePath string, cfs []ConfigFile) []ConfigFile {
	applicable := []ConfigFile{}

	for _, cf := range cfs {
		absSourcePath, err := filepath.Abs(sourcePath)
		if err != nil {
			ExitBecauseOfInternalError("Could not get absolute path for " + sourcePath)
		}
		absCfDir, err := filepath.Abs(cf.Dir())
		if err != nil {
			ExitBecauseOfInternalError("Could not get absolute path for " + cf.Dir())
		}
		if strings.HasPrefix(absSourcePath, absCfDir) {
			applicable = append(applicable, cf)
		}
	}

	return applicable
}

// ByPrecedence implements sort.Interface for []ConfigFile based on how many slashes are in the file's path.
type ByPrecedence []ConfigFile

func (a ByPrecedence) Len() int {
	return len(a)
}
func (a ByPrecedence) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByPrecedence) Less(i, j int) bool {
	return a[i].Precedence() < a[j].Precedence()
}

/**
 * @return Least important files first, most important files last.
 */
func SortConfigFilesByPrecendence(cfs []ConfigFile) []ConfigFile {
	sort.Sort(ByPrecedence(cfs))
	return cfs
}
