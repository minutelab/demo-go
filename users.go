package demo

// User is a user stored in the database
type User struct {
	ID    string
	Name  string
	Email string
}

// Store a user in the DB
func Store(u User) error {
	return execOnce("INSERT INTO users (id, name, email) VALUES ($1,$2,$3)", u.ID, u.Name, u.Email)
}

// Get a user by ID
func Get(id string) (*User, error) {
	var u User
	if err := db.Get(&u, "Select * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &u, nil
}
