CREATE TABLE universities (
    id CHAR(36) CHARACTER SET utf8 NOT NULL PRIMARY KEY COMMENT 'UUID',
    name VARCHAR(128) NOT NULL,
    meeting_station VARCHAR(1024) NOT NULL,
    UNIQUE (name)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE genders (
    id CHAR(1) CHARACTER SET utf8 NOT NULL PRIMARY KEY COMMENT 'UUID',
    name VARCHAR(128) NOT NULL,
    UNIQUE (name)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE users (
    id CHAR(36) CHARACTER SET utf8 NOT NULL PRIMARY KEY COMMENT 'UUID',
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(128) NOT NULL,
    user_name VARCHAR(128) NOT NULL,
    email VARCHAR(128) NOT NULL,
    password VARCHAR(128) NOT NULL,
    university_id CHAR(36) NOT NULL REFERENCES universities(id) ON DELETE CASCADE,
    birth_date DATE NOT NULL,
    year INTEGER NOT NULL,
    gender_id CHAR(36) NOT NULL REFERENCES genders(id) ON DELETE CASCADE,
    icon_url VARCHAR(1028) NOT NULL,
    instagram_id VARCHAR(256) NOT NULL,
    good INTEGER NOT NULL DEFAULT 0,
    bad INTEGER NOT NULL DEFAULT 0,
    UNIQUE (email),
    UNIQUE (user_name),
    INDEX (university_id, birth_date, year, gender_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
