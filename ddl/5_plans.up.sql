CREATE TABLE plans (
    id CHAR(36) CHARACTER SET utf8 NOT NULL PRIMARY KEY COMMENT 'UUID',
    spot_id CHAR(36) NOT NULL REFERENCES spots(spot_id) ON DELETE CASCADE,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE user_plans (
    id CHAR(36) CHARACTER SET utf8 NOT NULL PRIMARY KEY COMMENT 'UUID',
    plan_id CHAR(36) NOT NULL REFERENCES plans(id) ON DELETE CASCADE,
    canceled TIMESTAMP,
    finished TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
