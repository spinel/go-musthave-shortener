BEGIN;

CREATE TABLE IF NOT EXISTS entities (
	id serial not null,
    code character varying unique not null,
	url character varying not null,
    user_uuid uuid not null,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp,
	PRIMARY KEY(id)	
);

COMMIT;
