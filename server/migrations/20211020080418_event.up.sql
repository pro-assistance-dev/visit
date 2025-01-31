create table events
(
  id uuid default uuid_generate_v4() not null primary key,
  name  varchar ,
  slug  varchar ,
  type  varchar ,
  description varchar ,
  place varchar ,
  organizer varchar,

  thumbnail_id  uuid constraint thumbnail_id_fk references file_infos on update cascade on delete cascade
);
