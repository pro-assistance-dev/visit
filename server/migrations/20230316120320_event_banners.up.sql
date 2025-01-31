CREATE TABLE event_banners (
  id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
  event_id uuid ,
  file_info_id uuid
);
