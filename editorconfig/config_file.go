package editorconfig

import (
	"github.com/go-ini/ini"
	"regexp"
	"strings"
)

type ConfigFile struct {
	Path                    string
	DefaultRuleSet          RuleSet
	FileConstrainedRuleSets []FileConstrainedRuleSet
}

func (cf *ConfigFile) Dir() string {
	return GetParentDir(cf.Path)
}

func (cf *ConfigFile) IsRoot() bool {
	return cf.DefaultRuleSet.Get("root") == "true"
}

func (cf *ConfigFile) Precedence() int {
	return strings.Count(cf.Path, "/")
}

type Rule struct {
	Name  string
	Value string
}

type RuleSet []Rule

func (rs *RuleSet) Add(r Rule) {
	*rs = append(*rs, r)
}

func (rs RuleSet) Get(name string) string {
	for i := 0; i < len(rs); i++ {
		if rs[i].Name == name {
			return rs[i].Value
		}
	}
	return ""
}

type FileConstrainedRuleSet struct {
	Constraint       string
	ConstraintRegexp *regexp.Regexp
	Rules            RuleSet
}

func CreateConfigFileStruct(path string) ConfigFile {
	ini, err := ini.Load(path)
	if err != nil {
		ExitBecauseOfInternalError("Could not parse " + path)
	}

	cf := ConfigFile{}
	cf.Path = path
	cf.DefaultRuleSet = CreateRuleSetFromIniSectonName(ini, "")

	for _, sectionName := range ini.SectionStrings() {
		if sectionName == "DEFAULT" {
			continue
		}

		fcrs := FileConstrainedRuleSet{
			Constraint:       sectionName,
			ConstraintRegexp: ConvertWildcardPatternToGoRegexp(sectionName),
			Rules:            CreateRuleSetFromIniSectonName(ini, sectionName),
		}
		cf.FileConstrainedRuleSets = append(cf.FileConstrainedRuleSets, fcrs)
	}

	return cf
}

func CreateRuleSetFromIniSectonName(ini *ini.File, name string) RuleSet {
	section, err := ini.GetSection(name)
	if err != nil {
		ExitBecauseOfInternalError(err.Error())
		// return RuleSet{}
	}

	return CreateRuleSetFromIniSecton(section)
}

func CreateRuleSetFromIniSecton(section *ini.Section) RuleSet {
	rs := RuleSet{}

	for name, val := range section.KeysHash() {
		// The spec says rule names should be case-insensitive.
		name = strings.ToLower(name)

		rs.Add(Rule{
			Name:  name,
			Value: val,
		})
	}

	return rs
}
