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
)

type characterRecord struct {
	Id               int `db:"character_id"`
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
		ID:               c.Id,
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

func (a *App) InsertCharacter(char *model.Character) error {
	if char == nil || char.Stats == nil {
		return InvalidCharError
	}
	sess := a.db.NewSession(nil)
	res, err := sess.InsertInto(characterTable).
		Pair("name", char.Name).
		Pair("max_hit_points", char.MaxHitPoints).
		Pair("current_hit_points", char.CurrentHitPoints).
		Pair("level", char.Level).
		Pair("strength", char.Stats.Strength).
		Pair("dexterity", char.Stats.Dexterity).
		Pair("constitution", char.Stats.Constitution).
		Pair("intelligence", char.Stats.Intelligence).
		Pair("wisdom", char.Stats.Wisdom).
		Pair("charisma", char.Stats.Charisma).
		Exec()
	if err != nil {
		log.Printf("error attempting to insert character with name %s: %s", char.Name, err.Error())
		return UnexpectedDBError
	}
	log.Println(res)
	return nil
}
