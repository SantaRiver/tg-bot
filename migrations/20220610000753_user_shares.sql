-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.user_shares
(
    id         serial primary key,
    user_id    integer   not null,
    share_id   integer   not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table user_shares;
-- +goose StatementEnd
