-- +goose Up
-- +goose StatementBegin
alter table trace drop column error;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table trace add column error text;
-- +goose StatementEnd
