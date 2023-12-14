package app

import (
	"errors"
	"github.com/gocraft/dbr/v2"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/model"
	"log"
	"strings"
)

var (
	InvalidDefenseError      = errors.New("cannot insert invalid defense")
	DefenseTypeNotFoundError = errors.New("requested defense type not found")
	characterDefenseFields   = []string{"character_id", "damage_type", "defense_type"}
)

const (
	characterDefenseTable = "character_defense"
	defenseTypeTable      = "defense_type"
)

// characterDefenseRecord is used for binding to & from the character_defense table records.
type characterDefenseRecord struct {
	ID          int `db:"character_defense_id"`
	CharacterID int `db:"character_id"`
	DamageType  string
	DefenseType string
}

// toModel will convert a characterDefenseRecord to a gqlgen-generated *model.Defense type.
func (c *characterDefenseRecord) toModel() *model.Defense {
	return &model.Defense{
		DefenseType: model.DefenseType(strings.ToUpper(c.DefenseType)),
		DamageType:  model.DamageType(strings.ToUpper(c.DamageType)),
	}
}

// InsertCharacterDefense will insert a new character_defense record linked to the provided character id.
// Since the generated defense type and damage type are in uppercase, they will be converted to lower case for db insertion.
func (a *App) InsertCharacterDefense(characterID int, defense *model.Defense) error {
	if defense == nil {
		return InvalidDefenseError
	}

	defenseType := strings.ToLower(defense.DefenseType.String())
	damageType := strings.ToLower(defense.DamageType.String())

	sess := a.db.NewSession(nil)
	record := characterDefenseRecord{
		CharacterID: characterID,
		DamageType:  damageType,
		DefenseType: defenseType,
	}
	_, err := sess.InsertInto(characterDefenseTable).Columns(characterDefenseFields...).Record(&record).Exec()

	if err != nil {
		log.Printf("error attempting to insert character defense: %s", err.Error())
		return UnexpectedDBError
	}
	return nil
}

// GetCharacterDefenses returns a list of all the defenses that apply to the provided character id.
func (a *App) GetCharacterDefenses(characterID int) ([]*model.Defense, error) {
	sess := a.db.NewSession(nil)
	var records []characterDefenseRecord

	_, err := sess.Select("*").From(characterDefenseTable).Where("character_id = ?", characterID).Load(&records)
	if err != nil {
		log.Printf("error attempting to get character defenses for character id %d: %s", characterID, err.Error())
		return nil, UnexpectedDBError
	}

	var models []*model.Defense
	for _, record := range records {
		models = append(models, record.toModel())
	}

	return models, err
}

// GetDefenseTypeModifier will return a float64 that can be used to as a multiplicative damage modifier for the defense's damage type.
// Immunity = 0, Vulnerability = 2, Resistance = 0.5.
func (a *App) GetDefenseTypeModifier(defenseType model.DefenseType) (float64, error) {
	sess := a.db.NewSession(nil)

	defenseTypeStr := strings.ToLower(defenseType.String())
	var modifier float64
	err := sess.Select("modifier").From(defenseTypeTable).Where("type = ?", defenseTypeStr).LoadOne(&modifier)
	if err == nil {
		return modifier, nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return 1, DefenseTypeNotFoundError
	}

	log.Printf("error attempting to find defense type %s: %s", defenseTypeStr, err.Error())
	return 1, UnexpectedDBError
}
