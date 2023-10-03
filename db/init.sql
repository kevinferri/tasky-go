-- maybe better to just handle the schema in supabase

create table if not exists snippets (
  id uuid default uuid_generate_v4(),
  title varchar(64),

  primary key(id)
);