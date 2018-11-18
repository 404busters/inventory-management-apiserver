ROLLBACK;
BEGIN;

CREATE TABLE location (
  id uuid primary key,
  name text not null,
  created_at timestamp without time zone not null default current_timestamp,
  updated_at timestamp without time zone not null default current_timestamp
);

CREATE TABLE item_type (
  id uuid primary key,
  name text not null,
  description text,
  created_at timestamp without time zone not null default current_timestamp,
  updated_at timestamp without time zone not null default current_timestamp,
  deleted_at timestamp without time zone default null
);

CREATE TABLE item_type_tag (
  id uuid primary key,
  name text not null,
  created_at timestamp without time zone not null default current_timestamp,
  updated_at timestamp without time zone not null default current_timestamp
);

CREATE TABLE item_type_tag_item_type (
  item_type uuid references item_type(id),
  tag uuid references item_type_tag(id),
  primary key (item_type, tag),
  created_at timestamp without time zone not null default current_timestamp,
  updated_at timestamp without time zone not null default current_timestamp
);

CREATE TABLE inventory (
  id uuid primary key,
  item_type uuid not null references item_type(id),
  last_seen_location uuid not null references location(id),
  status text not null,
  created_at timestamp without time zone not null default current_timestamp,
  updated_at timestamp without time zone not null default current_timestamp,
  deleted_at timestamp without time zone default null
);

CREATE TABLE "user" (
  id uuid primary key,
  name text not null,
  created_at timestamp without time zone not null default current_timestamp,
  updated_at timestamp without time zone not null default current_timestamp
);

CREATE TABLE transport (
  id uuid primary key,
  person_in_charge uuid not null references "user"(id),
  location uuid not null references location(id),
  event_type text not null,
  notes text,
  created_at timestamp without time zone default current_timestamp,
  updated_at timestamp without time zone default current_timestamp,
  deleted_at timestamp without time zone default null
);

CREATE TABLE transport_inventory (
  transport uuid not null references transport(id),
  inventory uuid not null references inventory(id),
  created_at timestamp without time zone default current_timestamp,
  updated_at timestamp without time zone default current_timestamp,
  deleted_at timestamp without time zone default null,
  primary key (transport, inventory)
);

COMMIT;