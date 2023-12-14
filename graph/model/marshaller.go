package model

import "strings"

// UnmarshalJSON is defined on the DefenseType so that it implements the json.Unmarshaler interface.
// This is used when parsing the starting character json file into a *model.Character
func (d *DefenseType) UnmarshalJSON(data []byte) error {
	return d.UnmarshalGQL(strings.ToUpper(removeOuterQuotes(string(data))))
}

// UnmarshalJSON is defined on the DamageType so that it implements the json.Unmarshaler interface.
// This is used when parsing the starting character json file into a *model.Character
func (d *DamageType) UnmarshalJSON(data []byte) error {
	return d.UnmarshalGQL(strings.ToUpper(removeOuterQuotes(string(data))))
}

// removeOuterQuotes transforms a string such as `"foo"` to `foo`.
// If there are no outer quotes, the string is returned unchanged.
func removeOuterQuotes(s string) string {
	out := s
	if len(out) > 0 && out[0] == '"' {
		out = out[1:]
	}
	if len(out) > 0 && out[len(out)-1] == '"' {
		out = out[:len(out)-1]
	}
	return out
}
