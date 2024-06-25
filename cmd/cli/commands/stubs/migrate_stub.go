package stubs

type PostgresqlStubs struct {
}

// CreateUp Create up migration content.
func (receiver PostgresqlStubs) CreateUp() string {
	return `CREATE TABLE DummyTable (
  id SERIAL PRIMARY KEY NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp NULL
);
`
}

// CreateDown Create down migration content.
func (receiver PostgresqlStubs) CreateDown() string {
	return `DROP TABLE IF EXISTS DummyTable;
`
}

// UpdateUp Update up migration content.
func (receiver PostgresqlStubs) UpdateUp() string {
	return `ALTER TABLE DummyTable ADD column varchar(255) NOT NULL;
`
}

// UpdateDown Update down migration content.
func (receiver PostgresqlStubs) UpdateDown() string {
	return `ALTER TABLE DummyTable DROP COLUMN column;
`
}
