package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/MohitSilwal16/Picker/server/myerrors"
	"github.com/MohitSilwal16/Picker/server/pb"
	"github.com/MohitSilwal16/Picker/server/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// Prepared Statements
var stmtRegister *sql.Stmt
var stmtLogin *sql.Stmt
var stmtVerifySessionToken *sql.Stmt
var stmtUpdateSessionToken *sql.Stmt
var stmtIsUsernameValid *sql.Stmt
var stmtGetUsernameBySessionToken *sql.Stmt
var stmtLogout *sql.Stmt

func generateUniqueSessionToken() (string, error) {
	newSessionToken := utils.TokenGenerator()
	for {
		duplicateToken, err := IsSessionTokenValid(newSessionToken)
		if err != nil {
			return "", err
		}
		if !duplicateToken {
			break
		}
		newSessionToken = utils.TokenGenerator()
	}
	return newSessionToken, nil
}

func InitMaria() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	dbUser := os.Getenv("SQL_USER")
	dbPass := os.Getenv("SQL_PASS")
	dbName := os.Getenv("SQL_DB_NAME")
	dbPort := os.Getenv("SQL_PORT")

	if dbUser == "" || dbName == "" || dbPass == "" || dbPort == "" {
		return errors.New("DATABASE NAME, USER & PASS NOT SPECIFIED IN .ENV FILE")
	}

	dbUrl := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", dbUser, dbPass, dbPort, dbName)

	db, err = sql.Open("mysql", dbUrl)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Connection with Maria DB isn't Established")
		return err
	}
	log.Println("Connection with Maria DB is Established")

	stmtVerifySessionToken, err = db.Prepare("SELECT 1 FROM users WHERE UserToken = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for SessionToken Validation")
		return err
	}

	stmtRegister, err = db.Prepare("INSERT INTO users (UserName, UserPass, UserToken) VALUE (?, ?, ?);")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for User Registration")
		return err
	}

	stmtLogin, err = db.Prepare("SELECT 1 FROM users WHERE UserName = ? AND UserPass = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for User Login")
		return err
	}

	stmtUpdateSessionToken, err = db.Prepare("UPDATE users SET UserToken = ? WHERE UserName = ? AND UserPass = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for SessionToken Updation")
		return err
	}

	stmtIsUsernameValid, err = db.Prepare("SELECT 1 FROM users WHERE UserName = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement for Username Validation")
		return err
	}

	stmtGetUsernameBySessionToken, err = db.Prepare("SELECT UserName FROM users WHERE UserToken = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement to Get Username By Session Token")
		return err
	}

	stmtLogout, err = db.Prepare("UPDATE users SET UserToken = '' WHERE UserToken = ?;")
	if err != nil {
		log.Println("Error: Failed to Prepare Statement to Log Out")
		return err
	}

	return nil
}

// DB Methods
func IsSessionTokenValid(sessionToken string) (bool, error) {
	if sessionToken == "" {
		return false, nil
	}

	rows, err := stmtVerifySessionToken.Query(sessionToken)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// Returns True if Session Token is Valid
	return rows.Next(), nil
}

func Register(authRequest *pb.AuthRequest) (string, error) {
	rows, err := stmtIsUsernameValid.Query(authRequest.Name)
	if err != nil {
		return "", nil
	}
	defer rows.Close()

	if rows.Next() {
		return "", myerrors.ErrUsernameAlreadyUsedError
	}

	newSessionToken, err := generateUniqueSessionToken()
	if err != nil {
		return "", err
	}

	_, err = stmtRegister.Exec(authRequest.Name, authRequest.Pass, newSessionToken)
	if err != nil {
		return "", err
	}

	return newSessionToken, nil
}

func Login(authRequest *pb.AuthRequest) (string, error) {
	rows, err := stmtLogin.Query(authRequest.Name, authRequest.Pass)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", myerrors.ErrInvalidCredentials
	}

	newSessionToken, err := generateUniqueSessionToken()
	if err != nil {
		return "", err
	}

	_, err = stmtUpdateSessionToken.Exec(newSessionToken, authRequest.Name, authRequest.Pass)
	if err != nil {
		return "", err
	}

	return newSessionToken, nil
}

func GetUsernameBySessionToken(sessionToken string) (string, error) {
	if sessionToken == "" {
		return "", myerrors.ErrInvalidSessionToken
	}

	rows, err := stmtGetUsernameBySessionToken.Query(sessionToken)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var username string
	if rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			return "", err
		}
		return username, nil
	}
	return "", myerrors.ErrInvalidSessionToken
}

func Logout(sessionToken string) (bool, error) {
	res, err := stmtLogout.Exec(sessionToken)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected >= 1 {
		return true, nil
	}
	return false, nil
}
