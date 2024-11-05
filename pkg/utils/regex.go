package utils

import "regexp"

// ValidateByRegex 判断正则
func ValidateByRegex(str, rule string) bool {
	re := regexp.MustCompile(rule)
	return re.MatchString(str)
}
