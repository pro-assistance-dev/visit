create table users
(
  id uuid default uuid_generate_v4() not null primary key,
  email  varchar not null unique,
  password varchar not null,
  last_name varchar,
  first_name varchar,
  patron_name varchar,
  company varchar,
  post varchar,
  phone varchar
);
