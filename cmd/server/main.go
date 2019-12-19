package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/suchaimi/finance/internal/api"
	"github.com/suchaimi/finance/internal/config"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("Version", config.Version).Debug("starting server.")

	router, err := api.NewRouter()
	if err != nil {
		logrus.WithError(err).Fatal("Error building website router")
	}

	const addr = "0.0.0.0:8088"
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}
}
