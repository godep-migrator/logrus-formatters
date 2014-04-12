package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/meatballhat/logrus-formatters/l2met"
)

func main() {
	addr := ":8540"

	log := logrus.New()
	log.Formatter = l2met.NewFormatter("l2met-web")

	log.WithFields(logrus.Fields{"addr": addr}).Info("Listening")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.WithFields(logrus.Fields{
			"event":  "http-request",
			"url":    req.URL,
			"method": req.Method,
		}).Info("Request handled")
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(addr, nil)
}
