package main

import (
	"fmt"
	"io"	
	"encoding/json"
	"net/http"
	"os"
	"log"
)

type JsonResponse struct {
	Status  int
	Message string
}

func welcome(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Welcome"))
}

func fileUpload(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	var response JsonResponse

	file, header, err := request.FormFile("file")
	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		log.Println(response.Message)
		data, err := json.MarshalIndent(response, "", "  ")
		if err != nil {}
		fmt.Fprintln(rw, string(data))
		return
	}
	defer file.Close()

	out, err := os.Create("tmp/"+header.Filename)
	if err != nil {
		response.Status = 400
		response.Message = "Unable to create the file for writing. Check your write access privilege"
		log.Println(response.Message)
		data, err := json.MarshalIndent(response, "", "  ")
		if err != nil {}
		fmt.Fprintln(rw, string(data))
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		log.Println(response.Message)
		data, err := json.MarshalIndent(response, "", "  ")
		if err != nil {}
		fmt.Fprintln(rw, string(data))
	}

	response.Status = 200
	response.Message = "File uploaded successfully : " + header.Filename
	data, err := json.MarshalIndent(response, "", "  ")
	if err != nil {}
	fmt.Fprintln(rw, string(data))
	log.Println(response.Message)
}

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/fileUpload", fileUpload)
	log.Fatal(http.ListenAndServe(":3000", nil));
}
