package main

import (
	"github.com/tombenke/go-12f-common/v2/apprun"
	"github.com/tombenke/go-12f-common/v2/must"
	"github.com/tombenke/go-application-template/internal/application"
)

func main() {
	// Make and run an application via ApplicationRunner
	must.Must(apprun.MakeAndRun(&application.Config{}, application.NewApplication))
}
