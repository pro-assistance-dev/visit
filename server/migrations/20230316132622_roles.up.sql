CREATE TABLE roles (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR,
    label VARCHAR,
    start_page VARCHAR
);