package model_test

import (
	"encoding/json"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func jsonStr(s string) []byte {
	out, _ := json.Marshal(s)
	return out
}

func TestDamageType_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		data     []byte
		expected model.DamageType
	}{
		{
			name:     "should correctly parse valid damage type if type is in lower case",
			data:     jsonStr("fire"),
			expected: model.DamageTypeFire,
		},
		{
			name:     "should correctly parse valid damage type if type is in mixed case",
			data:     jsonStr("aCiD"),
			expected: model.DamageTypeAcid,
		},
		{
			name:     "should correctly parse valid damage type if type is in upper case",
			data:     jsonStr("THUNDER"),
			expected: model.DamageTypeThunder,
		},
		{
			name: "should error if damage type is not valid",
			data: jsonStr("abc"),
		},
		{
			name: "should error if damage type is an empty string",
			data: jsonStr(""),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var damageType model.DamageType
			if c.expected.IsValid() {
				require.NoError(t, json.Unmarshal(c.data, &damageType))
				require.Equal(t, c.expected, damageType)
			} else {
				require.Error(t, json.Unmarshal(c.data, &damageType))
			}
		})
	}
}

func TestDefenseType_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		data     []byte
		expected model.DefenseType
	}{
		{
			name:     "should correctly parse valid damage type if type is in lower case",
			data:     jsonStr("resistance"),
			expected: model.DefenseTypeResistance,
		},
		{
			name:     "should correctly parse valid damage type if type is in mixed case",
			data:     jsonStr("iMMunIty"),
			expected: model.DefenseTypeImmunity,
		},
		{
			name:     "should correctly parse valid damage type if type is in upper case",
			data:     jsonStr("VULNERABILITY"),
			expected: model.DefenseTypeVulnerability,
		},
		{
			name: "should error if damage type is not valid",
			data: jsonStr("abc"),
		},
		{
			name: "should error if damage type is an empty string",
			data: jsonStr(""),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var defenseType model.DefenseType
			if c.expected.IsValid() {
				require.NoError(t, json.Unmarshal(c.data, &defenseType))
				require.Equal(t, c.expected, defenseType)
			} else {
				require.Error(t, json.Unmarshal(c.data, &defenseType))
			}
		})
	}
}
