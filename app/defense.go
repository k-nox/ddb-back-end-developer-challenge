package app

import (
	"errors"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/model"
	"log"
	"strings"
)

var (
	InvalidDefenseError    = errors.New("cannot insert invalid defense")
	characterDefenseFields = []string{"character_id", "damage_type", "defense_type"}
)

const characterDefenseTable = "character_defense"

type characterDefenseRecord struct {
	ID          int `db:"character_defense_id"`
	CharacterID int `db:"character_id"`
	DamageType  string
	DefenseType string
}

func (c *characterDefenseRecord) toModel() *model.Defense {
	return &model.Defense{
		DefenseType: model.DefenseType(strings.ToUpper(c.DefenseType)),
		DamageType:  model.DamageType(strings.ToUpper(c.DamageType)),
	}
}

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
