package auth

import (
	"database/sql"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserCreate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite", "./users.db")
	if err != nil {
		return err
	}

	// Create users table if it doesn't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'user',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	// Create admin user if it doesn't exist
	return createDefaultAdmin()
}

func createDefaultAdmin() error {
	// Check if admin user exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'admin'").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// Create default admin user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			INSERT INTO users (username, email, password_hash, role) 
			VALUES (?, ?, ?, ?)`,
			"admin", "admin@secauto.local", string(hashedPassword), "admin")

		if err != nil {
			return err
		}
		log.Println("Default admin user created: admin/admin123")
	}

	return nil
}

func AuthenticateUser(username, password string) (*User, error) {
	var user User
	var passwordHash string

	err := db.QueryRow(`
		SELECT id, username, email, password_hash, role, created_at, updated_at 
		FROM users WHERE username = ?`, username).Scan(
		&user.ID, &user.Username, &user.Email, &passwordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(userData UserCreate) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO users (username, email, password_hash, role) 
		VALUES (?, ?, ?, ?)`,
		userData.Username, userData.Email, string(hashedPassword), userData.Role)

	if err != nil {
		log.Printf("Failed to create user %s: %v", userData.Username, err)
		return err
	}

	log.Printf("Successfully created user: %s with role: %s", userData.Username, userData.Role)
	return nil
}

func GetUsers() ([]User, error) {
	rows, err := db.Query(`
		SELECT id, username, email, role, created_at, updated_at 
		FROM users ORDER BY created_at DESC`)
	if err != nil {
		log.Printf("Failed to query users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("Failed to scan user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	log.Printf("Retrieved %d users from database", len(users))
	return users, nil
}

func DeleteUser(userID int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	return err
}

func GetUserByID(userID int) (*User, error) {
	var user User
	err := db.QueryRow(`
		SELECT id, username, email, role, created_at, updated_at 
		FROM users WHERE id = ?`, userID).Scan(
		&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
