DROP TABLE IF EXISTS property;
DROP TABLE IF EXISTS subscribe_traveler;
DROP TABLE IF EXISTS subscribe;
DROP TABLE IF EXISTS lessor;
DROP TABLE IF EXISTS provider;
DROP TABLE IF EXISTS traveler;
DROP TABLE IF EXISTS administrator;
DROP TABLE IF EXISTS user;

CREATE TABLE user (
    uuid UUID PRIMARY KEY,
    mail VARCHAR(320) NOT NULL,
    password VARCHAR(64) NOT NULL,
    register_date TIMESTAMP WITH TIME ZONE NOT NULL,
    last_connection_date TIMESTAMP WITH TIME ZONE
);

CREATE TABLE administrator (
    uuid UUID PRIMARY KEY,
    site VARCHAR(64),
    nickname VARCHAR(64) NOT NULL,
    user_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES user(uuid)
);

CREATE TABLE traveler (
    uuid UUID PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    user_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES user(uuid)
);

CREATE TABLE provider (
    uuid UUID PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    nickname VARCHAR(64) NOT NULL,
    user_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES user(uuid)
);

CREATE TABLE lessor (
    uuid UUID PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    user_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES user(uuid)
);

CREATE TABLE subscribe (
    uuid UUID PRIMARY KEY,
    type VARCHAR(64) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

CREATE TABLE subscribe_traveler (
    uuid UUID PRIMARY KEY,
    begin_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    traveler_uuid UUID NOT NULL,
    subscribe_uuid UUID NOT NULL,
    FOREIGN KEY (traveler_uuid) REFERENCES traveler(uuid),
    FOREIGN KEY (subscribe_uuid) REFERENCES subscribe(uuid)
);

CREATE TABLE property (
    uuid UUID PRIMARY KEY,
    address VARCHAR(64) NOT NULL,
    city VARCHAR(64) NOT NULL,
    zip_code VARCHAR(6) NOT NULL,
    country VARCHAR(64) NOT NULL,
    room INTEGER NOT NULL,
    bathroom INTEGER NOT NULL,
    description TEXT,
    administrator_validation BOOLEAN DEFAULT FALSE,
    lessor_uuid UUID NOT NULL,
    FOREIGN KEY (lessor_uuid) REFERENCES lessor(uuid)
);