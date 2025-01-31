CREATE TABLE chairpersons (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  name varchar,
  regalis varchar,
  event_id uuid ,
  photo_id uuid
);
