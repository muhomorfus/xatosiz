-- +goose Up
-- +goose StatementBegin
alter table alert_hit
    drop constraint alert_hit_config_uuid_fkey,
    add constraint alert_hit_config_uuid_fkey
        foreign key (config_uuid)
        references alert_config(uuid)
        on delete cascade;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table alert_hit
    drop constraint alert_hit_config_uuid_fkey,
    add constraint alert_hit_config_uuid_fkey
        foreign key (config_uuid)
            references alert_config(uuid);
-- +goose StatementEnd
