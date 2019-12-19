package v1

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/suchaimi/finance/internal/config"
)

// ServerVersion represet the server version
type ServerVersion struct {
	Version string `json:"version"`
}

// Marshaled JSON
var versionJSON []byte

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{
		Version: config.Version,
	})

	if err != nil {
		panic(err)
	}
}

// VersionHandler comment
func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error writing version.")
	}
}
