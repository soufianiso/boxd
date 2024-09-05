package movies

import (
	"database/sql"
)


type Storage struct{
	db *sql.DB
}


type Store interface{
	getAllmovies(string) error
	getMovieById(string) error
	
}


func NewStorage(db *sql.DB) *Storage{
	return &Storage{db:db}
}


func(s *Storage) getAllmovies(a string) error {
	// database logic	
	return nil
}

func(s *Storage) getMovieById(a string) error {
	// database logic	
	return nil
}






