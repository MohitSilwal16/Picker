
CREATE DATABASE Picker;
USE Picker;

CREATE TABLE user(
	username VARCHAR(30) PRIMARY KEY,
	password VARCHAR(30),
	session_token VARCHAR(8)
);

CREATE TABLE file_request(
	sender VARCHAR(30),
	dir_name VARCHAR(50),
	receiver VARCHAR(30),
	last_upload_id INT,
	FOREIGN KEY (sender) REFERENCES USER(username),
	FOREIGN KEY (receiver) REFERENCES USER(username)
);

CREATE TABLE upload_user1_dir1(
	upload_id INT PRIMARY KEY AUTO_INCREMENT,
	upload_method VARCHAR(15) CHECK( upload_method IN ("InitDir", "CreateDir", "CreateFile", "WriteFile", "RenameFileDir", "RemoveFileDir") ),
	upload_time DATETIME
);

SHOW TABLES;
