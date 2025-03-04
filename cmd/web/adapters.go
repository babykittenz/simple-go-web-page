package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"simple-go-web-page/models"
)

// CatBreedsInterface is simply our target interface, which defines all the methods that
// any type which implements this interface must have.
type CatBreedsInterface interface {
	GetAllCatBreeds() ([]*models.CatBreed, error)
}

// RemoteService is the Adaptor type. It embeds a DataInterface interface
// (which is critical to the pattern).
type RemoteService struct {
	Remote CatBreedsInterface
}

// GetAllBreeds is the function on RemoteService which lets us
// call any adaptor which implements the DataInterface type.
func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
	return rs.Remote.GetAllCatBreeds()
}

// JSONBackend is the JSON adapter, which needs to satisfy the CatBreedsInterface by
// have a GetAllCatsBreeds method.
type JSONBackend struct{}

// GetAllCatBreeds is necessary so that JSONBackend satisfies the CatBreedsInterface requirements.
func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")
	if err != nil {
		return nil, err
	}
	// no resource leak
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var breeds []*models.CatBreed

	// we unmarshal body into breeds
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	type catBreeds struct {
		XMLName struct{}           `xml:"cat-breeds"`
		Breeds  []*models.CatBreed `xml:"cat-breed"`
	}

	var breeds catBreeds

	err = xml.Unmarshal(body, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds.Breeds, nil

}
