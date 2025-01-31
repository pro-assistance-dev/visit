alter table event_messages
  alter column id set default uuid_generate_v4();

