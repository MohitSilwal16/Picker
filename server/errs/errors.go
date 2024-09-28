package errs

import "errors"

var (
	// Auth
	ErrInvalidUsernameFormat    = errors.New("USERNAME MUST BE B'TWIN 5-20 CHARS & ALPHANUMERIC")
	ErrInvalidPasswordFormat    = errors.New("PASSWORD MUST BE B'TWIN 8-20 CHARS, ATLEAST 1 UPPER, LOWER CASE & SPECIAL CHARACTER")
	ErrUsernameAlreadyUsedError = errors.New("USERNAME IS ALREADY USED")
	ErrInvalidCredentials       = errors.New("INVALID CREDENTIALS")
	ErrFailedToInitDir          = errors.New("USER REGISTERED BUT FAILED TO INIT DIR")
)

var (
	// Common
	ErrInternalServerError = errors.New("INTERNAL SERVER ERROR")
	ErrInvalidSessionToken = errors.New("INVALID SESSION TOKEN")
)

var (
	// File Watcher
	// Create & Write File
	ErrGivenPathIsDir = errors.New("GIVEN PATH IS DIRECTORY")

	// Create File
	ErrCannotCreateFile  = errors.New("CANNOT CREATE FILE")
	ErrFileAlreadyExists = errors.New("FILE ALREADY EXISTS")

	// Write File
	ErrCannotOpenFile      = errors.New("CANNOT OPEN FILE")
	ErrCannotWriteIntoFile = errors.New("CANNOT WRITE INTO FILE")
	ErrFileDoesntExists    = errors.New("FILE DOESN'T EXISTS")

	// Remove & Rename File
	ErrCannotRemoveFileDir = errors.New("CANNOT REMOVE FILE/DIRECTORY")
	ErrFileDirDoesntExists = errors.New("FILE/DIR DOESN'T EXISTS")

	// Rename File
	ErrCannotRenameFileDir  = errors.New("CANNOT RENAME FILE/DIRECTORY")
	ErrOldFileDoesntExists  = errors.New("OLD FILE DOESN'T EXISTS")
	ErrNewFileAlreadyExists = errors.New("NEW FILE ALREADY EXISTS")

	// Create Dir
	ErrCannotCreateDir  = errors.New("CANNOT CREATE DIRECTORY")
	ErrDirAlreadyExists = errors.New("DIRECTORY ALREADY EXISTS")
)
