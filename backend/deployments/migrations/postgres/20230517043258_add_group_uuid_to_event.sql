-- +goose Up
-- +goose StatementBegin
alter table event add column group_uuid uuid;
update event set group_uuid = (select t.group_uuid from trace t where t.uuid = event.trace_uuid);
create index event_group_idx on event using hash(group_uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index event_group_idx;
alter table event drop column group_uuid;
-- +goose StatementEnd
