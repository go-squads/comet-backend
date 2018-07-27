package repository

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"io"
	"math/rand"

	"github.com/go-squads/comet-backend/appcontext"
	"github.com/go-squads/comet-backend/domain"
)

type UserRepository struct {
	db *sql.DB
}

const (
	getUserIdQuery   = "SELECT id FROM users WHERE username = $1 AND password = $2"
	getUserSaltQuery = "SELECT salt FROM users WHERE username = $1"
	insertTokenQuery = "UPDATE users SET token = $1 WHERE id = $2"
)

func getRandomString() string {
	const chars = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 256)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func concatPasswordAndSalt(password string, salt string) string {
	var buffer bytes.Buffer

	buffer.WriteString(password)
	buffer.WriteString(salt)

	return buffer.String()
}

func hashString(stringPassword string) string {
	hasher := sha256.New()
	io.WriteString(hasher, stringPassword)
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	return hashedPassword
}

func (self UserRepository) LogIn(credentials domain.User) string {
	var userId int
	var userSalt string

	err = self.db.QueryRow(getUserSaltQuery, credentials.Username).Scan(&userSalt)
	if err != nil {
		// user does not exist
		return ""
	}

	passwordWithSalt := concatPasswordAndSalt(credentials.Password, userSalt)
	hashedPassword := hashString(passwordWithSalt)

	err = self.db.QueryRow(getUserIdQuery, credentials.Username, hashedPassword).Scan(&userId)
	if err != nil {
		// password incorrect
		return ""
	}

	token := getRandomString()

	self.db.Exec(insertTokenQuery, token, userId)

	return token
}

func GetUserRepository() UserRepository {
	return UserRepository{
		db: appcontext.GetDB(),
	}
}
