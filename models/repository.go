package models

import "database/sql"

// Repository is the database repository. Anything that implements
// this interface must implement all the methods included here.
type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(name string) (*DogBreed, error)
	GetDogOfMonthById(id int) (*DogOfMonth, error)
}

// mysqlRepository is a simple wrapper for the *sql.DB type. This is
// used to return a MySQL/MariaDB repository.
type mysqlRepository struct {
	DB *sql.DB
}

// newMysqlRepository is a convenience factory method to return a new mysqlRepository.
func newMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
