package main

import (
	"fmt"
	"github.com/gernest/utron"
	c "langAnalysis/controllers"
	"langAnalysis/models"
	"langAnalysis/search"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello dear friend :)")
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	err = app.Model.Register(&models.Settings{}, &models.HhBase{}, &models.GitHubBase{})
	if err != nil {
		log.Fatal(err)
	}

	// CReate Models tables if they dont exist yet
	app.Model.AutoMigrateAll()

	// Get search data
	search.Start(app)

	// Register Controller
	app.AddController(c.Controller)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
