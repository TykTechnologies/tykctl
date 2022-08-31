package main

import (
	"ara-client-sdk/swagger-gen"
	"context"
	"fmt"
	"log"
)

func main() {
	config := swagger.Configuration{
		BasePath:      "http://localhost:4010",
		DefaultHeader: make(map[string]string),
	}
	client := swagger.NewAPIClient(&config)
	list, _, err := client.EntitlementsApi.AdminGetEnt(context.Background(), "ssss")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("hello worked again")
	fmt.Printf("%+v\n", list.Payload)
}
