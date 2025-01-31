create table contacts
(
    id          uuid default uuid_generate_v4() not null primary key,
    description varchar,
    latitude    varchar,
    longitude   varchar,
    time        varchar
);

create table emails (
    id              uuid default uuid_generate_v4() not null primary key,
    address         varchar,
    description     varchar,
    contact_id uuid references contacts on update cascade on delete cascade,
    main            boolean
);

create table phones (
    id              uuid default uuid_generate_v4() not null primary key,
    number         varchar,
    description     varchar,
    contact_id uuid references contacts on update cascade on delete cascade,
    main            boolean
);

create table websites (
    id              uuid default uuid_generate_v4() not null primary key,
    address         varchar,
    description     varchar,
    contact_id uuid references contacts on update cascade on delete cascade,
    main            boolean
);

create table addresses
(
    id        uuid default uuid_generate_v4() not null primary key,
    region    varchar,
    city      varchar,
    street    varchar,
    building  varchar,
    flat      varchar,
    zip       integer,
    contact_id uuid references contacts on update cascade on delete cascade,
    region_id varchar,
    city_id   varchar,
    street_id varchar,
    b_id      varchar
);


