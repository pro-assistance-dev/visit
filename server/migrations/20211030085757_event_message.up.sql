create table event_messages
(
  id uuid default uuid_generate_v4() not null primary key,
  message  varchar ,
  created_on timestamp default current_timestamp not null,
  event_id uuid constraint event_id_fk references events on update cascade on delete cascade,
    user_id uuid constraint user_id_fk references users on update cascade on delete cascade
);
