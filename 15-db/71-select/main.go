package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// SQLite-safe configuration
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := createTable(ctx, db); err != nil {
		log.Fatal(err)
	}

	_, err = createUser(ctx, db, "Mario", "mario@example.com", "secret123")
	if err != nil {
		log.Fatal(err)
	}

	users, err := getUsers(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}

func createTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, schema)
	return err
}

func createUser(ctx context.Context, db *sql.DB, name, email, password string) (int64, error) {
	stmt := `
	INSERT INTO users (name, email, hashed_password)
	VALUES (?, ?, ?)
	`

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	res, err := db.ExecContext(ctx, stmt, name, email, string(hash))
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func getUserByEmail(ctx context.Context, db *sql.DB, email string) (*User, error) {
	stmt := `
	SELECT id, name, email, hashed_password, created_at
	FROM users
	WHERE email = ?
	`

	var u User
	err := db.QueryRowContext(ctx, stmt, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.HashedPassword, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func getUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	stmt := `
	SELECT id, name, email, hashed_password, created_at
	FROM users
	`

	rows, err := db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.HashedPassword,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, rows.Err()
}
