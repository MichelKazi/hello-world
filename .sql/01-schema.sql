CREATE USER admin
    WITH PASSWORD 'admin123'
    CREATEDB;

CREATE DATABASE dev
    WITH OWNER admin;

\connect dev;

CREATE EXTENSION pgcrypto;

CREATE SCHEMA helloworld AUTHORIZATION admin;

CREATE TABLE helloworld.person(
  Id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  Age integer,
  Person varchar(128),
  FavoriteFood varchar(128)
);

GRANT ALL PRIVILEGES 
    ON ALL TABLES 
    IN SCHEMA helloworld 
    TO admin;
