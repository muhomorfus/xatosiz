-- +goose Up
-- +goose StatementBegin
create extension if not exists "uuid-ossp";

create table "group" (
    uuid uuid default uuid_generate_v4() primary key
);

create table trace (
    uuid uuid default uuid_generate_v4() primary key,
    group_uuid uuid references "group"(uuid) on delete cascade,
    parent_uuid uuid references trace(uuid) on delete cascade,
    title text not null,
    error text,
    time_start timestamp not null,
    time_end timestamp,
    component text not null
);

create type priority_t as enum (
    'info',
    'warning',
    'error',
    'fatal'
);

create table event (
    uuid uuid default uuid_generate_v4() primary key,
    trace_uuid uuid references trace(uuid) on delete cascade,
    time timestamp not null,
    priority priority_t not null,
    message text not null,
    payload text,
    fixed boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table event;
drop type priority_t;
drop table trace;
drop table "group";
drop extension "uuid-ossp";
-- +goose StatementEnd
