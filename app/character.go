package app

import (
	"errors"
	"github.com/gocraft/dbr/v2"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/model"
	"log"
)

const (
	characterTable = "character"
)

var (
	CharNotFoundError = errors.New("no character found")
	InvalidCharError  = errors.New("cannot insert invalid character")
	characterFields   = []string{"name", "max_hit_points", "current_hit_points", "level", "strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"}
)

type characterRecord struct {
	ID               int `db:"character_id"`
	Name             string
	MaxHitPoints     int
	CurrentHitPoints int
	Level            int
	Strength         int
	Dexterity        int
	Constitution     int
	Wisdom           int
	Intelligence     int
	Charisma         int
}

func (c characterRecord) toModel() *model.Character {
	return &model.Character{
		ID:               c.ID,
		Name:             c.Name,
		MaxHitPoints:     c.MaxHitPoints,
		CurrentHitPoints: c.CurrentHitPoints,
		Level:            c.Level,
		Stats: &model.Stats{
			Strength:     c.Strength,
			Dexterity:    c.Dexterity,
			Constitution: c.Constitution,
			Intelligence: c.Intelligence,
			Wisdom:       c.Wisdom,
			Charisma:     c.Charisma,
		},
	}
}

func (a *App) GetCharacterByName(name string) (*model.Character, error) {
	var record characterRecord
	sess := a.db.NewSession(nil)
	err := sess.Select("*").From(characterTable).Where("name = ?", name).LoadOne(&record)
	if err == nil {
		return record.toModel(), nil
	}
	if errors.Is(err, dbr.ErrNotFound) {
		return nil, CharNotFoundError
	}
	log.Printf("error attempting to select char by name %s: %s", name, err.Error())
	return nil, UnexpectedDBError
}

func (a *App) InsertCharacter(char *model.Character) (*model.Character, error) {
	if char == nil || char.Stats == nil {
		return nil, InvalidCharError
	}
	sess := a.db.NewSession(nil)
	record := characterRecord{
		Name:             char.Name,
		MaxHitPoints:     char.MaxHitPoints,
		CurrentHitPoints: char.CurrentHitPoints,
		Level:            char.Level,
		Strength:         char.Stats.Strength,
		Dexterity:        char.Stats.Dexterity,
		Constitution:     char.Stats.Constitution,
		Wisdom:           char.Stats.Wisdom,
		Intelligence:     char.Stats.Intelligence,
		Charisma:         char.Stats.Charisma,
	}
	err := sess.InsertInto(characterTable).Columns(characterFields...).Record(&record).Returning("character_id").Load(&record.ID)
	if err != nil {
		log.Printf("error attempting to insert character with name %s: %s", char.Name, err.Error())
		return nil, UnexpectedDBError
	}
	return record.toModel(), nil
}
