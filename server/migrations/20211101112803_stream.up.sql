create table streams
(
  id uuid default uuid_generate_v4() not null primary key,

  poster_id  uuid constraint poster_id_fk references file_infos on update cascade on delete cascade
);
