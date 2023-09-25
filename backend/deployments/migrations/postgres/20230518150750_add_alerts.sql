-- +goose Up
-- +goose StatementBegin
create table alert_config (
    uuid uuid default uuid_generate_v4() primary key,
    message_expression text not null,
    min_priority priority_t not null,
    duration interval not null,
    min_rate int not null,
    comment text not null
);

create table alert_hit (
    uuid uuid default uuid_generate_v4() primary key,
    config_uuid uuid not null references alert_config(uuid),
    time timestamp not null
);

create table alert (
    uuid uuid default uuid_generate_v4() primary key,
    message text not null,
    event_uuid uuid not null references event(uuid),
    time timestamp not null,
    solved boolean not null default false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table alert;
drop table alert_hit;
drop table alert_config;
-- +goose StatementEnd
