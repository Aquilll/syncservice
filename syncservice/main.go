package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"syncservice/api"
	"syncservice/provider"
	"syncservice/queue"
	"syncservice/worker"
)

func main() {
	providersList := os.Getenv("PROVIDERS")
	bufferSizeStr := os.Getenv("QUEUE_BUFFER_SIZE")
	salesforceRPSStr := os.Getenv("SALESFORCE_RPS")
	hubspotRPSStr := os.Getenv("HUBSPOT_RPS")
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		log.Fatal("Missing required environment variable: PORT")
	}

	if providersList == "" || bufferSizeStr == "" {
		log.Fatal("Missing required environment variables")
	}

	providers := strings.Split(providersList, ",")
	bufferSize, _ := strconv.Atoi(bufferSizeStr)
	salesforceRPS, _ := strconv.Atoi(salesforceRPSStr)
	hubspotRPS, _ := strconv.Atoi(hubspotRPSStr)

	pq := queue.NewPartitionedQueue(providers, bufferSize)
	handler := &api.APIHandler{PartitionedQueue: pq}
	http.HandleFunc("/internal/crud", handler.CrudHandler)
	http.HandleFunc("/webhook", handler.WebhookHandler)

	//starts worker for different providers in separate thread.
	go worker.StartWorkerForProvider("salesforce", &provider.SalesforceProvider{}, pq, salesforceRPS)
	go worker.StartWorkerForProvider("hubspot", &provider.HubSpotProvider{}, pq, hubspotRPS)

	log.Println("Starting HTTP server on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
