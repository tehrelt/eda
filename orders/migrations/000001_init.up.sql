create extension if not exists "uuid-ossp";

create table if not exists orders (
  id uuid primary key default uuid_generate_v4(),
  who varchar not null,
  total decimal(10, 2) not null,
  created_at timestamp default now(),
  updated_at timestamp
);