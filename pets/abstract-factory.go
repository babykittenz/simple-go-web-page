package pets

import (
	"errors"
	"fmt"
	"simple-go-web-page/models"
)

type AnimalInterface interface {
	Show() string
}
type DogFromFactory struct {
	Pets *models.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pets.Breed.Breed)
}

type CatFromFactory struct {
	Pets *models.Cat
}

func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pets.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}
type DogAbstractFactory struct {
}

func (dff *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pets: &models.Dog{},
	}
}

type CatAbstractFactory struct {
}

func (cff *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pets: &models.Cat{},
	}
}
func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("invalid species supplied")
	}
}
