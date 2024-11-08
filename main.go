package main

import (
	route "Backend_berkah/routes"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("jumat_berkah", route.URL)
}