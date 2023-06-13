BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "Post" (
	"UserId"	INTEGER,
	"Text"	TEXT,
	"Id"	INTEGER UNIQUE,
	FOREIGN KEY("UserId") REFERENCES "User"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "Comments" (
	"PostId"	INTEGER,
	"UserId"	INTEGER,
	"Text"	TEXT,
	"Id"	INTEGER UNIQUE,
	FOREIGN KEY("UserId") REFERENCES "User"("Id"),
	FOREIGN KEY("PostId") REFERENCES "Post"("Id"),
	PRIMARY KEY("Id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "User" (
	"Id"	INTEGER UNIQUE,
	"Pseudo"	TEXT NOT NULL UNIQUE,
	"Password"	TEXT NOT NULL,
	"email"	TEXT NOT NULL UNIQUE,
	PRIMARY KEY("Id")
);
INSERT INTO "User" VALUES (1,'hi','$2a$14$IRPNcoGyNdeLsKoyC..kyOYVwtVjSmpOzo1zKbEKgKpmfN7KCNBi.','mathis@ynov.com');
INSERT INTO "User" VALUES (2,'Mathis','$2a$14$iHkgubHiXUDP2ig3ENLlgOLbWJZlnZZn.ltHSUm48NXu7ZDTrN1SW','mathis.silotia@ynov.com');
INSERT INTO "User" VALUES (3,'Mathis2','$2a$14$xCL/EszaXeXVPzqHxCYLcu6r8ooLN6sCcOXnb/mJJ2J9dhXNg4s.e','mathis.silotia2@ynov.com');
COMMIT;
