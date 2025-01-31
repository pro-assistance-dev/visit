alter table users add column
check_uuid uuid DEFAULT uuid_generate_v4() not null;
