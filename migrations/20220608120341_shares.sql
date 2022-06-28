-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.shares
(
    id              serial primary key,
    figi            varchar(16)  not null,
    class_code      varchar(16)  not null,
    currency        varchar(8)   not null,
    name            varchar(128) not null,
    country_of_risk varchar(8)   not null,
    created_at      timestamp    not null default current_timestamp,
    updated_at      timestamp    not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table shares;
-- +goose StatementEnd
