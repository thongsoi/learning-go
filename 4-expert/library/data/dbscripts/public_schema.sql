-- Initial public schema relates to Library 0.x

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

CREATE EXTENSION IF NOT EXISTS "plpgsql";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

SET search_path = public;
SET default_tablespace = '';

-- update updated at column
CREATE OR REPLACE FUNCTION update_updated_at_column() RETURNS TRIGGER
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$;

-- enum_user_role
CREATE TABLE enum_user_role (
    code integer NOT NULL,
    user_role text NOT NULL,
    CONSTRAINT enum_user_status_pk PRIMARY KEY (code)
);

-- enum_book_status
CREATE TABLE enum_book_status (
    code integer NOT NULL,
    book_status text NOT NULL,
    CONSTRAINT enum_book_status_pk PRIMARY KEY (code)
);

-- library_user
CREATE TABLE library_user (
    user_id uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    username text NOT NULL UNIQUE,
    user_password text NOT NULL,
    full_name text NOT NULL,
    user_role integer DEFAULT 1,
    token uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    CONSTRAINT library_user_pk PRIMARY KEY (user_id),
    CONSTRAINT fk_library_user_user_role FOREIGN KEY (user_role)
        REFERENCES enum_user_role (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

-- book
CREATE TABLE book (
    book_id uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    book_name text NOT NULL,
    author_name text NOT NULL,
    publisher text NOT NULL,
    book_description text,
    book_status integer DEFAULT 1,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    borrower_id uuid,
    CONSTRAINT book_pk PRIMARY KEY (book_id),
    CONSTRAINT fk_book_book_status FOREIGN KEY (book_status)
        REFERENCES enum_book_status (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_book_borrower_id FOREIGN KEY (borrower_id)
        REFERENCES library_user (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE INDEX book_book_status
ON book (book_status);

CREATE TRIGGER update_book_updated_at_column
    BEFORE UPDATE
    ON book
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();

