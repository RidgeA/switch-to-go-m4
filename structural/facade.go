package structural

type database struct {
	connected bool
}

var db = database{}

func (d database) connect() {

}

func (d database) startTransaction() {

}

func (d database) query(q string) {

}

func (d database) closeTransaction() {

}

func Query(query string) {
	if !db.connected {
		db.connect()
	}
	db.startTransaction()
	db.query(query)
	db.closeTransaction()
}
