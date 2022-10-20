CREATE TABLE user_management.users (
	id BIGINT auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	username varchar(100) NOT NULL,
	password varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	PRIMARY KEY(id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_unicode_ci;