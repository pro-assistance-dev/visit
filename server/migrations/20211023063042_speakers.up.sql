create table speakers
(
  id uuid default uuid_generate_v4() not null primary key,
  name  varchar,
  regalias  varchar,

  schedule_item_id uuid constraint schedule_items_id_fk references schedule_items on update cascade on delete cascade
);
