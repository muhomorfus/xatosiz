-- +goose Up
-- +goose StatementBegin
alter table alert drop constraint alert_event_uuid_fkey;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table alert add constraint alert_event_uuid_fkey foreign key event_uuid references event(uuid);
-- +goose StatementEnd
