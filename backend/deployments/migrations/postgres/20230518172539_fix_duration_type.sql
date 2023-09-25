-- +goose Up
-- +goose StatementBegin
alter table alert_config drop column duration;
alter table alert_config add column duration text not null default '10m';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table alert_config drop column duration;
alter table alert_config add column duration interval not null default '00:10:00'::interval;
-- +goose StatementEnd
