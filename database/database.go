package database

type Db interface {
	Migrate()
	Connect()
	Disconnect()
}
