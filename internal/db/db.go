package db

import (
	"database/sql"
	"errors"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

// Database is an alias of sql.DB
type Database struct {
	internal *sql.DB
}

// InitDatabase is to initialize database
func InitDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

// Rows queries the number of table rows.
func (db *Database) Rows() (int, error) {
	var count int

	row := db.internal.QueryRow("SELECT COUNT(*) FROM `jokes`")

	err := row.Scan(&count)
	if err != nil {
		return 0, errors.New("出错啦！")
	}

	return count, nil
}

// Random returns a random choosed joke.
func (db *Database) Random() (string, error) {
	r, err := db.Rows()
	if err != nil {
		return "", errors.New("出错啦！")
	}

	uid := rand.Intn(r) + 1

	return db.GetJokeOf(uid)
}

// GetJokeOf returns a joke by given uid.
func (db *Database) GetJokeOf(uid int) (string, error) {
	r, err := db.Rows()
	if err != nil {
		return "", errors.New("出错啦！")
	}

	if uid <= 0 || uid > r {
		return "", errors.New("超出范围啦！")
	}

	var content string
	var sqlStatement = `SELECT content FROM jokes WHERE uid=$1`

	row := db.internal.QueryRow(sqlStatement, uid)

	err = row.Scan(&content)
	if err != nil {
		return "", errors.New("出错啦！")
	}

	return content, nil
}

// RandomN returns random choosed N jokes.
func (db *Database) RandomN(n int) ([]string, error) {
	// r, err := db.Rows()
	// if err != nil {
	// 	return nil, err
	// }

	return []string{}, nil
}
