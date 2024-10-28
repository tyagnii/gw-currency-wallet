package db

type DataBase interface {
	Connect()
	Query()
	Post()
}
