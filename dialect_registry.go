package dsc

var dialect = NewDefaultDialect()
var datastoreDialectableRegistry = make(map[string]DatastoreDialect)

//RegisterDatastoreDialect register DatastoreDialect for a driver.
func RegisterDatastoreDialect(driver string, dialectable DatastoreDialect) {
	datastoreDialectableRegistry[driver] = dialectable
}

//GetDatastoreDialect returns DatastoreDialect for passed in driver.
func GetDatastoreDialect(driver string) DatastoreDialect {
	if result, ok := datastoreDialectableRegistry[driver]; ok {
		return result
	}
	panic("failed to lookup datastore dialectabe: " + driver)
}

func init() {
	RegisterDatastoreDialect("mysql", newMySQLDialect())
	RegisterDatastoreDialect("pg", newPgDialect())
	RegisterDatastoreDialect("postgres", newPgDialect())
	RegisterDatastoreDialect("ora", newOraDialect())
	RegisterDatastoreDialect("mssql", newMsSQLDialect())
	RegisterDatastoreDialect("sqlite3", newSQLLiteDialect())
	RegisterDatastoreDialect("cql", newCasandraDialect())
	RegisterDatastoreDialect("odbc", newOdbcDialect())
	RegisterDatastoreDialect("ndjson", &fileDialect{})
	RegisterDatastoreDialect("tsv", &fileDialect{})
	RegisterDatastoreDialect("csv", &fileDialect{})

}
