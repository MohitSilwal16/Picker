package handler

import "errors"

var ErrInternalServer = errors.New("INTERNAL SERVER ERROR")

// Create & Write File
var ErrGivenPathIsDir = errors.New("GIVEN PATH IS DIRECTORY")

// Create File
var ErrCannotCreateFile = errors.New("CANNOT CREATE FILE")
var ErrFileAlreadyExists = errors.New("FILE ALREADY EXISTS")

// Write File
var ErrCannotOpenFile = errors.New("CANNOT OPEN FILE")
var ErrCannotWriteIntoFile = errors.New("CANNOT WRITE INTO FILE")
var ErrFileDoesntExists = errors.New("FILE DOESN'T EXISTS")

// Remove & Rename File
var ErrCannotRemoveFileDir = errors.New("CANNOT REMOVE FILE/DIRECTORY")
var ErrFileDirDoesntExists = errors.New("FILE/DIR DOESN'T EXISTS")

// Rename File
var ErrCannotRenameFileDir = errors.New("CANNOT RENAME FILE/DIRECTORY")
var ErrOldFileDoesntExists = errors.New("OLD FILE DOESN'T EXISTS")
var ErrNewFileAlreadyExists = errors.New("NEW FILE ALREADY EXISTS")

// Create Dir
var ErrCannotCreateDir = errors.New("CANNOT CREATE DIRECTORY")
var ErrDirAlreadyExists = errors.New("DIRECTORY ALREADY EXISTS")
