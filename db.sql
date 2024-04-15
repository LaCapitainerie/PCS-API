DROP TABLE IF EXISTS message;
DROP TABLE IF EXISTS ticket;
DROP TABLE IF EXISTS chat;
DROP TABLE IF EXISTS review_lessor_to_service;
DROP TABLE IF EXISTS review_traveler_to_service;
DROP TABLE IF EXISTS review_traveler_to_property;
DROP TABLE IF EXISTS reservation_service;
DROP TABLE IF EXISTS reservation_bill;
DROP TABLE IF EXISTS reservation;
DROP TABLE IF EXISTS property_service;
DROP TABLE IF EXISTS bill;
DROP TABLE IF EXISTS provider_licence;
DROP TABLE IF EXISTS service_type;
DROP TABLE IF EXISTS type_of_service;
DROP TABLE IF EXISTS service_unavailability;
DROP TABLE IF EXISTS service;
DROP TABLE IF EXISTS property_image;
DROP TABLE IF EXISTS property_unavailability;
DROP TABLE IF EXISTS property;
DROP TABLE IF EXISTS subscribe_traveler;
DROP TABLE IF EXISTS subscribe;
DROP TABLE IF EXISTS lessor;
DROP TABLE IF EXISTS provider;
DROP TABLE IF EXISTS traveler;
DROP TABLE IF EXISTS administrator;
DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    mail VARCHAR(320) NOT NULL,
    password VARCHAR(64) NOT NULL,
    register_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_connection_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE administrator (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    site VARCHAR(64),
    nickname VARCHAR(64) NOT NULL,
    user_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES "user"(uuid)
);

CREATE TABLE traveler (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    user_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES "user"(uuid)
);

CREATE TABLE provider (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    nickname VARCHAR(64) NOT NULL,
    user_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES "user"(uuid)
);

CREATE TABLE lessor (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    user_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES "user"(uuid)
);

CREATE TABLE subscribe (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type VARCHAR(64) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

CREATE TABLE subscribe_traveler (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    begin_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    traveler_uuid UUID NOT NULL,
    subscribe_uuid UUID NOT NULL,
    FOREIGN KEY (traveler_uuid) REFERENCES traveler(uuid),
    FOREIGN KEY (subscribe_uuid) REFERENCES subscribe(uuid)
);

CREATE TABLE property (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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

CREATE TABLE property_unavailability (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    property_uuid UUID NOT NULL,
    FOREIGN KEY (property_uuid) REFERENCES property(uuid)
);

CREATE TABLE property_image (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    path VARCHAR(255) NOT NULL,
    property_uuid UUID NOT NULL,
    FOREIGN KEY (property_uuid) REFERENCES property(uuid)
);

CREATE TABLE service(
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    price NUMERIC(10,2) NOT NULL,
    target_customer VARCHAR(8) NOT NULL, -- Peut prendre que les valeurs "all" "lessor" ou "traveler"
    address VARCHAR(64) NOT NULL,
    city VARCHAR(64) NOT NULL,
    zip_code VARCHAR(6) NOT NULL,
    country VARCHAR(64) NOT NULL,
    range_action INTEGER, -- null = infinie
    description TEXT,
    provider_uuid UUID NOT NULL,
    FOREIGN KEY (provider_uuid) REFERENCES provider(uuid)
);

CREATE TABLE service_unavailability(
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    service_uuid UUID NOT NULL,
    FOREIGN KEY (service_uuid) REFERENCES service(uuid)
);

CREATE TABLE type_of_service(
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type VARCHAR(64) NOT NULL,
    licence BOOLEAN DEFAULT FALSE
);

CREATE TABLE service_type
(
    service_uuid UUID NOT NULL,
    type_of_service_uuid UUID NOT NULL,
    FOREIGN KEY (service_uuid) REFERENCES service (uuid),
    FOREIGN KEY (type_of_service_uuid) REFERENCES type_of_service (uuid),
    PRIMARY KEY (service_uuid, type_of_service_uuid)
);

CREATE TABLE provider_licence(
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP,
    validation BOOLEAN DEFAULT FALSE,
    path_proof VARCHAR(255) NOT NULL,
    provider_uuid UUID NOT NULL,
    FOREIGN KEY (provider_uuid) REFERENCES provider(uuid)
);

CREATE TABLE bill(
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    price NUMERIC(10, 2) NOT NULL,
    date TIMESTAMP NOT NULL,
    type VARCHAR(64),
    content TEXT
);

CREATE TABLE property_service(
    property_uuid UUID NOT NULL,
    service_uuid UUID NOT NULL,
    bill_uuid UUID NOT NULL,
    date TIMESTAMP NOT NULL,
    FOREIGN KEY (bill_uuid) REFERENCES bill(uuid),
    FOREIGN KEY (property_uuid) REFERENCES property(uuid),
    FOREIGN KEY (service_uuid) REFERENCES service(uuid),
    PRIMARY KEY (property_uuid, service_uuid)
);

CREATE TABLE reservation(
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    traveler_uuid UUID NOT NULL,
    property_uuid UUID NOT NULL,
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    FOREIGN KEY (traveler_uuid) REFERENCES traveler(uuid),
    FOREIGN KEY (property_uuid) REFERENCES property(uuid)
);

CREATE TABLE reservation_bill(
    reservation_uuid UUID NOT NULL,
    bill_uuid UUID NOT NULL,
    FOREIGN KEY (reservation_uuid) REFERENCES reservation(uuid),
    FOREIGN KEY (bill_uuid) REFERENCES bill(uuid),
    PRIMARY KEY (reservation_uuid, bill_uuid)
);

CREATE TABLE reservation_service(
    reservation_uuid UUID NOT NULL,
    service_uuid UUID NOT NULL,
    date TIMESTAMP NOT NULL,
    FOREIGN KEY (reservation_uuid) REFERENCES reservation(uuid),
    FOREIGN KEY (service_uuid) REFERENCES service(uuid),
    PRIMARY KEY (reservation_uuid, service_uuid)
);

CREATE TABLE review_traveler_to_property(
    traveler_uuid UUID NOT NULL,
    property_uuid UUID NOT NULL,
    note numeric(10, 1) NOT NULL,
    comment TEXT,
    FOREIGN KEY (traveler_uuid) REFERENCES traveler(uuid),
    FOREIGN KEY (property_uuid) REFERENCES property(uuid),
    PRIMARY KEY (traveler_uuid, property_uuid)
);

CREATE TABLE review_traveler_to_service(
    traveler_uuid UUID NOT NULL,
    service_uuid UUID NOT NULL,
    note numeric(10, 1) NOT NULL,
    comment TEXT,
    FOREIGN KEY (traveler_uuid) REFERENCES traveler(uuid),
    FOREIGN KEY (service_uuid) REFERENCES service(uuid),
    PRIMARY KEY (traveler_uuid, service_uuid)
);

CREATE TABLE review_lessor_to_service (
    lessor_uuid UUID NOT NULL,
    service_uuid UUID NOT NULL,
    note numeric(10, 1) NOT NULL,
    comment TEXT,
    FOREIGN KEY (lessor_uuid) REFERENCES lessor(uuid),
    FOREIGN KEY (service_uuid) REFERENCES service(uuid),
    PRIMARY KEY (lessor_uuid, service_uuid)
);

CREATE TABLE chat (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    view BOOLEAN DEFAULT FALSE
);

CREATE TABLE ticket
(
    uuid        UUID PRIMARY KEY,
    type        VARCHAR(64) NOT NULL,
    state       VARCHAR(16) NOT NULL,
    description TEXT        NOT NULL,
    chat_uuid  UUID        NOT NULL,
    FOREIGN KEY (chat_uuid) REFERENCES chat (uuid)
);

CREATE TABLE message (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    type VARCHAR(5), -- "text" ou "image"
    user_uuid UUID NOT NULL,
    chat_uuid UUID NOT NULL,
    FOREIGN KEY (user_uuid) REFERENCES "user"(uuid),
    FOREIGN KEY (chat_uuid) REFERENCES chat(uuid)
);