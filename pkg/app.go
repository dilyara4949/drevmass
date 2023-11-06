package pkg

import "github.com/dilyara4949/drevmass/internal/postgres"

type Application struct {
	Env *Env
	Pql postgres.Database
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Pql = NewDatabase(app.Env)
	return *app
}

func (app *Application)CloseDBConnection() {
	app.CloseDBConnection()
}