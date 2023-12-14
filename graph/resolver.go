package graph

//go:generate go run github.com/99designs/gqlgen generate

import "github.com/k-nox/ddb-backend-developer-challenge/app"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app *app.App
}

// New returns a new resolver configured with an app dependency
func New(app *app.App) *Resolver {
	return &Resolver{
		app: app,
	}
}
