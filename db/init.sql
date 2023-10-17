create table if not exists users (
  id uuid default uuid_generate_v4() not null unique,
  email varchar(64) not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,

  primary key(id)
);

create table if not exists snippets (
  id uuid default uuid_generate_v4() not null unique,
  title varchar(64) not null,
  created_by uuid references users not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,

  primary key(id)
);