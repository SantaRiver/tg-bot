-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.users (
  id serial primary key,
  name varchar(255) not null,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table users;
-- +goose StatementEnd
