package pkg

import "strings"

/*
	Complete the method/function so that it converts dash/underscore delimited words into camel casing.
	The first word within the output should be capitalized only if the original word was capitalized.

	Examples
	"the-stealth-warrior" gets converted to "theStealthWarrior"
	"The_Stealth_Warrior" gets converted to "TheStealthWarrior"
*/
func ToCamelCase(s string) string {
	if s != "" {
		titleCase := strings.Title(strings.Replace(strings.ToLower(s), "_", "-", -1))
		camelCase := strings.Replace(titleCase, "-", "", -1)
		return s[:1] + camelCase[1:]
	}

	return ""
}
