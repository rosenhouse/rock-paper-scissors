package main

import (
	"log"
	"net/http"
	"os"

	"github.com/desmondrawls/rock-paper-scissors/web_ui"
)

func getEnvVarOrDefault(envVar string, defaultValue string) string {
	val := os.Getenv(envVar)
	if val != "" {
		return val
	}
	return defaultValue
}

func getListenAddress() string {
	ip := getEnvVarOrDefault("LISTEN_IP", "0.0.0.0")
	port := getEnvVarOrDefault("PORT", "8080")
	return ip + ":" + port
}

func main() {
	handler := &web_ui.Handler{}
	listenAddr := getListenAddress()
	log.Printf("server listening on http://%s\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, handler))
}
