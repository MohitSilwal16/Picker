
USE Picker;

SHOW TABLES
WHERE Tables_in_picker = "user";

SELECT *
FROM USER;

SELECT *
FROM uploader_nimesh_demo;

SELECT *
FROM uploader_nimesh_test;

DROP TABLE uploader_nimesh_demo;
DROP TABLE uploader_nimesh_test;


DESCRIBE user;
DESCRIBE file_request;
DESCRIBE upload_admin_dir;
