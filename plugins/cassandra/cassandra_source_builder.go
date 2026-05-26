package cassandra

/*
This file is only for the builder API.
*/

func Hosts(hosts string) *CassandraSource { _ = "STUB: not implemented"; return nil }

func (s *CassandraSource) Keyspace(keyspace string) *CassandraSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *CassandraSource) From(table string) *CassandraSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *CassandraSource) Select(selectClause string) *CassandraSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *CassandraSource) Where(whereClause string) *CassandraSource {
	_ = "STUB: not implemented"
	return nil
}
