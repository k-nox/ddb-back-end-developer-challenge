package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/k-nox/ddb-backend-developer-challenge/graph/model"
	"io"
	"log"
	"os"
)

func parseStartingCharacter() (*model.Character, error) {
	jsonFile, err := os.Open("briv.json")
	if err != nil {
		return nil, err
	}
	defer func(json *os.File) {
		err := json.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	bytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var char model.Character

	err = json.Unmarshal(bytes, &char)
	if err != nil {
		return nil, err
	}

	return &char, nil
}

func (a *App) Startup() error {
	char, err := parseStartingCharacter()
	if err != nil {
		return err
	}

	_, err = a.GetCharacterByName(char.Name)
	if err == nil {
		// character is already in db, no need to insert
		return nil
	}

	if err != nil && !errors.Is(err, CharNotFound) {
		return err
	}
	_, err = a.InsertCharacter(*char)
	if err != nil {
		log.Fatalf("unable to insert starting character: %s", err.Error())
		return err
	}
	return nil
}
