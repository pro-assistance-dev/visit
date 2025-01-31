create table users_events_actives
(
  id uuid default uuid_generate_v4() not null primary key,
  minutes int,
  event_time varchar,
  created_at timestamp default current_timestamp not null,
  event_id uuid constraint event_id_fk references events on update cascade on delete cascade,
  user_id uuid constraint user_id_fk references users on update cascade on delete cascade,
  UNIQUE (user_id, event_id, event_time)
);
