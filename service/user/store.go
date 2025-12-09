package user

import ("github.com/Bois1/ecomm/types"
	 	"database/sql"
		
)


type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail( email string) (*types.User, error){
	rows, err := s.db.Query("SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		err = rows.Scan(
			&user.ID, 
			&user.FirstName, 
			&user.LastName, 
			&user.Email, 
			&user.Password, 
			&user.CreatedAt)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil
}

func (s *Store) CreateUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}