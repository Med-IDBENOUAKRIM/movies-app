package main

import (
	"net/http"
)

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(w, http.StatusOK, envelope{"health": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
