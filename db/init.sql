create table if not exists snippets (
  id uuid default uuid_generate_v4(),
  title varchar(64),

  primary key(id)
);

create table if not exists users (
  id uuid default uuid_generate_v4(),
  email varchar(64),

  primary key(id)
);