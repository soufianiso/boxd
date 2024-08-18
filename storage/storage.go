package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/soufianiso/letterboxd/types"
)



type PostgresStore struct {
	db *sql.DB
}



type Store interface {
	FetchAccount(int) (*types.User , error)	
}

func NewStore(connection string) (*PostgresStore, error){
	db , err := sql.Open("postgres",connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil{
		return nil,err 
	}

	return &PostgresStore{ db: db }, nil
}


func (p *PostgresStore) FetchAccount(id int) (*types.User,  error){
	user := &types.User{}
	query := `SELECT * FROM users WHERE id = $1`
	row := p.db.QueryRow(query, id)
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil  
}


