-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table brand (
  id int unsigned not null,
  name varchar(255) not null,
  designer varchar(255),
  created_at timestamp(6) not null default current_timestamp(6),
  updated_at timestamp(6) not null default current_timestamp(6),
  primary key(id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table brand;
