-- +goose Up
-- +goose StatementBegin
create index trace_parent_idx on trace using hash(parent_uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index trace_parent_idx;
-- +goose StatementEnd
