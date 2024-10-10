create extension if not exists "uuid-ossp";

create table if not exists customers (
  id uuid primary key default uuid_generate_v4(),
  username varchar not null,
  created_at timestamp default now(),
  updated_at timestamp
);

create table if not exists orders (
  id uuid primary key default uuid_generate_v4(),
  customer_id uuid not null references customers(id),
  total decimal(10, 2) not null,
  created_at timestamp default now(),
  updated_at timestamp
);