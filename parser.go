package main

import "strings"

// YARARule represents the structure of a parsed YARA rule
type YARARule struct {
	RuleName  string
	Tags      []string
	Meta      map[string]string
	Strings   map[string]string
	Condition string
}

// ParseRule parses a YARA rule into a YARARule struct
func ParseRule(rule string) YARARule {
	lines := strings.Split(rule, "\n")
	var yaraRule YARARule
	yaraRule.Meta = make(map[string]string)
	yaraRule.Strings = make(map[string]string)

	metaSection := false
	stringsSection := false
	conditionSection := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		if strings.HasPrefix(line, "rule") && yaraRule.RuleName == "" {
			parts := strings.Fields(line)
			yaraRule.RuleName = strings.TrimRight(parts[1], ":")
			if len(parts) > 2 {
				yaraRule.Tags = parts[2:]
			}
		} else if line == "meta:" {
			metaSection = true
			stringsSection = false
			conditionSection = false
		} else if line == "strings:" {
			metaSection = false
			stringsSection = true
			conditionSection = false
		} else if line == "condition:" {
			metaSection = false
			stringsSection = false
			conditionSection = true
		} else {
			if metaSection {
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					yaraRule.Meta[key] = strings.Trim(value, "\"")
				}
			} else if stringsSection {
				parts := strings.SplitN(line, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					yaraRule.Strings[key] = strings.Trim(value, "\"")
				}
			} else if conditionSection {
				yaraRule.Condition += " " + line
			}
		}
	}
	yaraRule.Condition = strings.TrimSpace(yaraRule.Condition)
	return yaraRule
}
