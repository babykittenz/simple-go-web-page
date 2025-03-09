package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"simple-go-web-page/models"
	"simple-go-web-page/pets"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"), nil)
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"), nil)
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}
func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, cat, nil)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)

}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("A mixed breed of unknown origin. Probably has German Shepard").
		SetColor("Black and white").
		SetAge(3).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("calico").
		SetWeight(5).
		SetDescription("This cat is pretty cool").
		SetColor("Black").
		SetAge(3).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// Since we are using the adapter pattern, this handler does not care where it gets the data from;
	// it will simply use whatever is stored in app.catService.
	catBreeds, err := app.App.CatService.GetAllBreeds()
	if err != nil {
		log.Println(err)
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	// Setup toolbox
	var t toolbox.Tools
	// Get species from URL itself.
	species := chi.URLParam(r, "species")
	// Get breed from the URL.
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)

	// Create a pet from abstract factory
	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// Write the result as JSON.
	_ = t.WriteJSON(w, http.StatusOK, pet)
}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {
	// Get the breed
	breed, _ := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")
	fmt.Println(breed)

	// Get the dog of the month from database
	dom, _ := app.App.Models.Dog.GetDogOfMonthById(1)

	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2023-11-01")

	// Create dog and decorate it
	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:               1,
			DogName:          "Sam",
			BreedId:          breed.ID,
			Color:            "Black & Tan",
			DateOfBirth:      dob,
			SpayedOrNeutered: true,
			Description:      "Sam is a very good boy.",
			Weight:           20,
			Breed:            *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	// Serve the web page
	data := make(map[string]any)
	data["dog"] = dog

	app.render(w, "dog-of-month.page.gohtml", &templateData{Data: data})
}
