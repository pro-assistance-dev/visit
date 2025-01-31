create table schedule_items
(
  id uuid default uuid_generate_v4() not null primary key,
  theme  varchar,
  description varchar ,
  start_time time,
  end_time time,

  schedule_id uuid constraint schedule_id_fk references schedules on update cascade on delete cascade
);
