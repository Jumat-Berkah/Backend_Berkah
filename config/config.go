package config

import "os"

var privateKey = os.Getenv("PRIVATE_KEY")
var publicKey = os.Getenv("PUBLIC_KEY")