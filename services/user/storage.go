package user


import(
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}


type Store interface {
	CreateUser(string) string
}

func NewStorage(db *sql.DB) *Storage{
	return &Storage{db: db}
}


func(s *Storage) CreateUser(b string) string{
	return b	
}
