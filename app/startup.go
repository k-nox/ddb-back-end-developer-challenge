package app

import (
	"encoding/json"
	"errors"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/model"
	"io"
	"log"
	"os"
)

var StartupError = errors.New("error with app startup")

func parseJsonCharacter(jsonFilePath string) (*model.Character, error) {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		log.Printf("error opening json file: %s", err.Error())
		return nil, StartupError
	}
	defer func(json *os.File) {
		err := json.Close()
		if err != nil {
			log.Printf("error closing json file: %s", err.Error())
		}
	}(jsonFile)

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Printf("error reading json file: %s", err.Error())
		return nil, StartupError
	}

	var char model.Character

	err = json.Unmarshal(bytes, &char)
	if err != nil {
		log.Printf("error marshalling json: %s", err.Error())
		return nil, StartupError
	}

	char.CurrentHitPoints = char.MaxHitPoints

	return &char, nil
}

func (a *App) Startup(startingCharPath string) error {
	char, err := parseJsonCharacter(startingCharPath)
	if err != nil {
		return err
	}

	_, err = a.GetCharacterByName(char.Name)
	if err == nil {
		// character is already in db, no need to insert
		return nil
	}

	if err != nil && !errors.Is(err, CharNotFoundError) {
		return err
	}

	inserted, err := a.InsertCharacter(char)
	if err != nil {
		log.Fatalf("unable to insert starting character: %s", err.Error())
		return err
	}

	for _, defense := range char.Defenses {
		err := a.InsertCharacterDefense(inserted.ID, defense)
		if err != nil {
			log.Fatalf("unable to insert starting character's defenses: %s", err.Error())
			return err
		}
	}

	return nil
}
