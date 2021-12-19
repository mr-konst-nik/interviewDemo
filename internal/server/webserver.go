package server

import (
	"fmt"
	"interviewDemo/internal/logger"
	"net/http"
)

func StartSRV(httpListen string) {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(httpListen, nil); err != nil {
		logger.Log.Sugar().Errorf("Error in starting WEB server %s", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello wolrd!")
	logger.Log.Sugar().Debugf("New visit: %s", r.URL)
}
