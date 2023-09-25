-- +goose Up
-- +goose StatementBegin
drop index trace_parent_idx;
create index trace_group_idx on trace using hash(group_uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index trace_group_idx;
create index trace_parent_idx on trace using hash(parent_uuid);
-- +goose StatementEnd
