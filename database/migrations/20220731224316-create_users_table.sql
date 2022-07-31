-- +migrate Up
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT,
    username VARCHAR(120) NOT NULL,
    email VARCHAR(120) NOT NULL,
    phone VARCHAR(12),
    status TINYINT NOT NULL DEFAULT(1),
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY(id),
    CONSTRAINT username_unique UNIQUE(username),
    CONSTRAINT email_unique UNIQUE(email),
    CONSTRAINT phone_unique UNIQUE(phone)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS users;