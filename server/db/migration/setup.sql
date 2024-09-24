CREATE DATABASE Picker;
USE Picker;

CREATE TABLE Users(
	UserName VARCHAR(30) PRIMARY KEY CHECK( UserName <> "" ),
	UserPass VARCHAR(30) NOT NULL CHECK( UserPass <> "" ),
	UserToken VARCHAR(8) NOT NULL
);

CREATE TABLE FileUploads(
	UploadId INT PRIMARY KEY,
	UploadTime DATETIME NOT NULL,
	UploadMethod VARCHAR(15) CHECK (UploadMethod IN ("CreateFile", "CreateDir", "WriteFile", "RemoveFileDir", "RenameFileDir") ),
	SenderName VARCHAR(30),
	DirName VARCHAR(50),
	FOREIGN KEY (SenderName) REFERENCES Users(UserName)
);

CREATE TABLE WriteUploads(
	UploadId INT,
	FileContent VARCHAR(30) CHECK(FileContent <> ""),
	FOREIGN KEY (UploadId) REFERENCES FileUploads(UploadId)
);

CREATE TABLE FileRequests(
	SenderName VARCHAR(30),
	DirName VARCHAR(50),
	ReceiverName VARCHAR(30),
	LastUploadId INT,
	LastUploadTime DATETIME,
	FOREIGN KEY (SenderName) REFERENCES Users(UserName),
	FOREIGN KEY (ReceiverName) REFERENCES Users(UserName),
	FOREIGN KEY (LastUploadId) REFERENCES FileUploads(UploadId),
	FOREIGN KEY (DirName) REFERENCES FileUploads(DirName)
);

