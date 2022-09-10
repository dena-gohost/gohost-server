CREATE TABLE spots (
    id CHAR(36) CHARACTER SET utf8 NOT NULL PRIMARY KEY COMMENT 'UUID',
    name VARCHAR(256) NOT NULL,
    description VARCHAR(1024) NOT NULL,
    image_url VARCHAR(1024) NOT NULL,
    address VARCHAR(1024) NOT NULL,
    meeting_station VARCHAR(256) NOT NULL,
    meeting_hour CHAR(5) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    INDEX (NAME)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE spot_universities (
    id CHAR(36) CHARACTER SET utf8 NOT NULL PRIMARY KEY COMMENT 'UUID',
    spot_id CHAR(36) NOT NULL REFERENCES spots(id) ON DELETE CASCADE,
    university_id CHAR(36) NOT NULL REFERENCES universities(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    UNIQUE (spot_id, university_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE entries (
    id CHAR(36) CHARACTER SET utf8 NOT NULL PRIMARY KEY COMMENT 'UUID',
    user_id CHAR(36) NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    date DATE NOT NULL,
    spot_id CHAR(36) NOT NULL REFERENCES spots(spot_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    UNIQUE (user_id, date)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;