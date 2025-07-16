package stringman

import "strings"

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func AddLastnameFirstname(firstname string, lastname string) string {

	capf := capitalize(firstname)
	capl := capitalize(lastname)
	return capf + " " + capl
}
