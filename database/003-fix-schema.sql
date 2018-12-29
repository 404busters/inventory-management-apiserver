BEGIN;
CREATE TABLE item_type_preset (
  id uuid primary key,
  display_name text not null
);

CREATE TABLE item_type_preset_item_type (
  preset uuid references item_type_preset(id),
  item_type uuid references item_type(id),
  require_number int not null,
  primary key (preset, item_type)
);

ALTER TABLE inventory_alias ADD COLUMN alias_type text not null default '';
ALTER TABLE inventory_alias ALTER COLUMN alias_type DROP DEFAULT;
ALTER TABLE inventory_alias ADD CONSTRAINT unique_alias_type UNIQUE (inventory, alias_type);

ALTER TABLE transport_inventory DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;

COMMIT;