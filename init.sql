CREATE DATABASE IF NOT EXISTS "ACTIONS_COMPANIES_DB";

USE "ACTIONS_COMPANIES_DB";

CREATE TABLE IF NOT EXISTS "actions" (
    "ticker" VARCHAR(5) NOT NULL,
    "target_from" VARCHAR NOT NULL,
    "target_to" VARCHAR NOT NULL,
    "company" VARCHAR NOT NULL,
    "action" VARCHAR NOT NULL,
    "brokerage" VARCHAR NOT NULL,
    "rating_from" VARCHAR NOT NULL,
    "rating_to" VARCHAR NOT NULL,
    "time" DATE NOT NULL,
    PRIMARY KEY ("ticker")
);

CREATE TABLE IF NOT EXISTS "users" (
    "id" VARCHAR(255) NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    PRIMARY KEY ("id")
);

