package user

import (
	"database/sql"
	"fmt"
	"log"

	// "log"

	_ "github.com/lib/pq"
	"github.com/soufianiso/boxd/types"
)

type Storage struct {
	db *sql.DB
}


type Store interface {
	CreateUser(*types.User, string) error
	GetUserByEmail(string) (*types.User, error)
}

func NewStorage(db *sql.DB) *Storage{
	return &Storage{db: db}
}


func(s *Storage) CreateUser(user *types.User,  password string,) error{
	_, err := s.db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES ($1,$2,$3,$4)", user.FirstName, user.LastName, user.Email, password,)
	if err != nil{
		fmt.Println(err)	
		return err
	}

	return nil
}

func(s *Storage) GetUserByEmail(email string) (*types.User, error){
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil{
		log.Print(err)
		return  nil, err 
	}
	
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
