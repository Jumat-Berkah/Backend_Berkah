package config

import (
	"Backend_berkah/helper"
	"Backend_berkah/model"
	"os"
)

var mongoURI = os.Getenv("JUMAT_BERKAH")

var mongoInfo = model.DBInfo{ // menggunakan model.DBinfo
	DBString: mongoURI,
	DBName:   "jumat_berkah",
}

var DB, ErrorMongoconn = helper.MongoConnect(mongoInfo) // menggunakan apdb.MongoConnect