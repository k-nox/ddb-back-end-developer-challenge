// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Character struct {
	ID                 int        `json:"id"`
	Name               string     `json:"name"`
	MaxHitPoints       int        `json:"hitPoints"`
	CurrentHitPoints   int        `json:"currentHitPoints"`
	TemporaryHitPoints *int       `json:"temporaryHitPoints,omitempty"`
	Level              int        `json:"level"`
	Stats              *Stats     `json:"stats"`
	Defenses           []*Defense `json:"defenses"`
}

type DamageInput struct {
	CharacterID int        `json:"characterId"`
	DamageType  DamageType `json:"damageType"`
	Roll        int        `json:"roll"`
}

type Defense struct {
	DefenseType DefenseType `json:"defense"`
	DamageType  DamageType  `json:"type"`
}

type HealInput struct {
	CharacterID int `json:"characterId"`
	Roll        int `json:"roll"`
}

type Stats struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}

type DamageType string

const (
	DamageTypeBludgeoning DamageType = "BLUDGEONING"
	DamageTypePiercing    DamageType = "PIERCING"
	DamageTypeSLAShing    DamageType = "SLASHING"
	DamageTypeFire        DamageType = "FIRE"
	DamageTypeCold        DamageType = "COLD"
	DamageTypeAcid        DamageType = "ACID"
	DamageTypeThunder     DamageType = "THUNDER"
	DamageTypeLightning   DamageType = "LIGHTNING"
	DamageTypePoison      DamageType = "POISON"
	DamageTypeRadiant     DamageType = "RADIANT"
	DamageTypeNecrotic    DamageType = "NECROTIC"
	DamageTypePsychic     DamageType = "PSYCHIC"
	DamageTypeForce       DamageType = "FORCE"
)

var AllDamageType = []DamageType{
	DamageTypeBludgeoning,
	DamageTypePiercing,
	DamageTypeSLAShing,
	DamageTypeFire,
	DamageTypeCold,
	DamageTypeAcid,
	DamageTypeThunder,
	DamageTypeLightning,
	DamageTypePoison,
	DamageTypeRadiant,
	DamageTypeNecrotic,
	DamageTypePsychic,
	DamageTypeForce,
}

func (e DamageType) IsValid() bool {
	switch e {
	case DamageTypeBludgeoning, DamageTypePiercing, DamageTypeSLAShing, DamageTypeFire, DamageTypeCold, DamageTypeAcid, DamageTypeThunder, DamageTypeLightning, DamageTypePoison, DamageTypeRadiant, DamageTypeNecrotic, DamageTypePsychic, DamageTypeForce:
		return true
	}
	return false
}

func (e DamageType) String() string {
	return string(e)
}

func (e *DamageType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DamageType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DamageType", str)
	}
	return nil
}

func (e DamageType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DefenseType string

const (
	DefenseTypeImmunity      DefenseType = "IMMUNITY"
	DefenseTypeResistance    DefenseType = "RESISTANCE"
	DefenseTypeVulnerability DefenseType = "VULNERABILITY"
)

var AllDefenseType = []DefenseType{
	DefenseTypeImmunity,
	DefenseTypeResistance,
	DefenseTypeVulnerability,
}

func (e DefenseType) IsValid() bool {
	switch e {
	case DefenseTypeImmunity, DefenseTypeResistance, DefenseTypeVulnerability:
		return true
	}
	return false
}

func (e DefenseType) String() string {
	return string(e)
}

func (e *DefenseType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DefenseType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DefenseType", str)
	}
	return nil
}

func (e DefenseType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
