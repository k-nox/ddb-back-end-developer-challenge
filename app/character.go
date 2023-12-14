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

// characterRecord is used for binding to & from the `character` table in the sqlite database.
type characterRecord struct {
	ID                 int `db:"character_id"`
	Name               string
	MaxHitPoints       int
	CurrentHitPoints   int
	TemporaryHitPoints dbr.NullInt64
	Level              int
	Strength           int
	Dexterity          int
	Constitution       int
	Wisdom             int
	Intelligence       int
	Charisma           int
}

// toModel converts a characterRecord to a gqlgen-generated *model.Character.
func (c characterRecord) toModel() *model.Character {
	return &model.Character{
		ID:                 c.ID,
		Name:               c.Name,
		MaxHitPoints:       c.MaxHitPoints,
		CurrentHitPoints:   c.CurrentHitPoints,
		TemporaryHitPoints: nullInt64ToPtr(c.TemporaryHitPoints),
		Level:              c.Level,
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

// GetCharacterByName returns a *model.Character with information by performing a select based on the provided name.
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

// InsertCharacter creates a new character record from a *model.Character.
// If the character has nil stats, it is rejected and not inserted.
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

// GetCharacterByID selects a character from the database by character_id and converts it to a *model.Character.
func (a *App) GetCharacterByID(id int) (*model.Character, error) {
	var record characterRecord
	sess := a.db.NewSession(nil)

	err := sess.Select("*").From(characterTable).Where("character_id = ?", id).LoadOne(&record)
	if err == nil {
		return record.toModel(), nil
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return nil, CharNotFoundError
	}

	log.Printf("error attempting to select char by id %d: %s", id, err.Error())
	return nil, UnexpectedDBError
}

// UpdateHitPoints will update the character's current_hit_points to a new value.
func (a *App) UpdateHitPoints(id int, newHitPoints int) error {
	sess := a.db.NewSession(nil)
	_, err := sess.Update(characterTable).Set("current_hit_points", newHitPoints).Where("character_id = ?", id).Exec()
	if err != nil {
		log.Printf("error attempting to update char hitpoints for char id %d: %s", id, err.Error())
		return UnexpectedDBError
	}
	return nil
}

// UpdateTemporaryHitPoints will update a character's temporary_hit_points to a new value or nil.
func (a *App) UpdateTemporaryHitPoints(id int, newTempHitPoints *int) error {
	sess := a.db.NewSession(nil)
	_, err := sess.Update(characterTable).Set("temporary_hit_points", newTempHitPoints).Where("character_id = ?", id).Exec()
	if err != nil {
		log.Printf("error attempting to update char temp hit points for char id %d: %s", id, err.Error())
		return UnexpectedDBError
	}

	return nil
}

// nullInt64ToPtr converts a dbr.NullInt64 type to a *int.
// If the provided value is 0, it returns nil.
func nullInt64ToPtr(i dbr.NullInt64) *int {
	if i.Valid && i.Int64 != 0 {
		out := int(i.Int64)
		return &out
	}
	return nil
}
