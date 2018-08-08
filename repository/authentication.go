package repository

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"io"
	"math/rand"

	"fmt"
	"github.com/go-squads/comet-backend/appcontext"
	"github.com/go-squads/comet-backend/domain"
	"log"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

const (
	tokenLength              = 64
	getUserIdQuery           = "SELECT id FROM users WHERE username = $1 AND password = $2"
	getUserSaltQuery         = "SELECT salt FROM users WHERE username = $1"
	insertTokenQuery         = "UPDATE users SET token = $1 WHERE id = $2"
	checkTokenAvailableQuery = "SELECT token FROM users"
	userRoleQuery            = "SELECT role FROM users WHERE token = $1"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStringGenerator() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	strGen := make([]rune, tokenLength)
	for i := range strGen {
		strGen[i] = letter[rand.Intn(len(letter))]
	}
	return string(strGen)
}

func getRandomString() string {
	const chars = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, tokenLength)
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

func (self UserRepository) LogIn(credentials domain.User) (string, string, string) {
	var userId int
	var fullname string
	var userRole string
	var userSalt string

	err = self.db.QueryRow(getUserSaltQuery, credentials.Username).Scan(&userSalt)
	if err != nil {
		// user does not exist
		return "", "", ""
	}

	passwordWithSalt := concatPasswordAndSalt(credentials.Password, userSalt)
	hashedPassword := hashString(passwordWithSalt)

	err = self.db.QueryRow(getUserIdQuery, credentials.Username, hashedPassword).Scan(&userId)
	if err != nil {
		// password incorrect
		return "", "", ""
	}

	token := randomStringGenerator()

	self.db.Exec(insertTokenQuery, token, userId)
	err = self.db.QueryRow("SELECT name,role FROM users WHERE id = $1 ", userId).Scan(&fullname, &userRole)
	if err != nil {
		log.Println(err.Error())
	}

	return token, fullname, userRole
}

func (self UserRepository) ValidateUserToken(token string) bool {
	var dbToken string
	var rows *sql.Rows
	isValidate := true

	rows, err = self.db.Query(checkTokenAvailableQuery)
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&dbToken)
		if token == dbToken {
			isValidate = true
		} else if token == "" {
			isValidate = false
		} else if token != dbToken {
			isValidate = false
		}
	}

	fmt.Println(isValidate)
	return isValidate
}

func (self UserRepository) SetUserRoleBased(token string) {
	_, err = self.db.Exec("SET ROLE " + self.getUserRoleBase(token))
	if err != nil {
		fmt.Println(err)
	}
}

func (self UserRepository) getUserRoleBase(token string) string {
	var role string

	err = self.db.QueryRow(userRoleQuery, token).Scan(&role)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(role)
	return role
}

func GetUserRepository() UserRepository {
	return UserRepository{
		db: appcontext.GetDB(),
	}
}
