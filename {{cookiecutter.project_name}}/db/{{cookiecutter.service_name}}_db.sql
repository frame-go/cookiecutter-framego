CREATE TABLE IF NOT EXISTS object_tab (
	id bigint NOT NULL,
	name varchar(100) NOT NULL,
	description varchar(1000) NOT NULL,
	status smallint NOT NULL,
	create_time timestamptz NOT NULL,
	update_time timestamptz NOT NULL,
	PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_object_tab_create_time ON object_tab (create_time);
