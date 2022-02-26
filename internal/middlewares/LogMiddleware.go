package middlewares

import (
	"encoding/json"
	"github.com/lcnssantos/iothub/pkg"
	"log"
	"net/http"
	"os"
	"time"
)

type LogMiddleware struct {
}

type Request struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

type Log struct {
	Host    string    `json:"host"`
	Version string    `json:"version"`
	Time    time.Time `json:"time"`
	Data    Request   `json:"data"`
}

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

func (this LogMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		logMessage := Log{
			Host:    hostname,
			Version: pkg.Version,
			Time:    time.Now(),
			Data: Request{
				Method:   r.Method,
				Endpoint: r.RequestURI,
			},
		}
		jsonMessage, _ := json.Marshal(logMessage)

		log.Println(string(jsonMessage))

		next.ServeHTTP(w, r)
	})
}
