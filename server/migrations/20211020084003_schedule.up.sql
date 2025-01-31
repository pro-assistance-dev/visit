create table schedules
(
  id uuid default uuid_generate_v4() not null primary key,
  name  varchar not null unique,
  description varchar,
  schedule_date date,

  event_id uuid constraint event_id_fk references events on update cascade on delete cascade
);
