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
	CharNotFound = errors.New("no character found")
)

func (a *App) GetCharacterByName(name string) (*model.Character, error) {
	var char model.Character
	sess := a.db.NewSession(nil)
	err := sess.Select("*").From(characterTable).Where("name = ?", name).LoadOne(&char)
	if err == nil {
		return &char, nil
	}
	if errors.Is(err, dbr.ErrNotFound) {
		return nil, CharNotFound
	}
	log.Printf("error attempting to select char by name %s: %s", name, err.Error())
	return nil, UnexpectedDBError
}

func (a *App) InsertCharacter(char model.Character) (*model.Character, error) {
	var out model.Character
	sess := a.db.NewSession(nil)
	err := sess.InsertInto(characterTable).Columns("*").Record(&char).Load(&out)
	if err != nil {
		log.Printf("error attempting to insert character with name %s: %s", char.Name, err.Error())
		return nil, UnexpectedDBError
	}

	return &out, nil
}
