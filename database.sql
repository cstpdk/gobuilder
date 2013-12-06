CREATE TABLE project(
    name CHAR(100) PRIMARY KEY NOT NULL,
    descript TEXT,
    git CHAR(100),
    gitbranch CHAR(100),
    buildkey CHAR(64)
);

CREATE TABLE user(
    username CHAR(100) PRIMARY KEY NOT NULL,
    password NOT NULL,
    email CHAR(100),
    role CHAR(10)
);

CREATE TABLE build(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user CHAR(100),
    complete INT DEFAULT 0
);
