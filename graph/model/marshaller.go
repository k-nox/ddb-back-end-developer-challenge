package model

import "strings"

func (d *DefenseType) UnmarshalJSON(data []byte) error {
	return d.UnmarshalGQL(strings.ToUpper(removeOuterQuotes(string(data))))
}

func (d *DamageType) UnmarshalJSON(data []byte) error {
	return d.UnmarshalGQL(strings.ToUpper(removeOuterQuotes(string(data))))
}

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
