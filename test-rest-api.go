// Copyright 2018 by Gregory Mirsky. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const codeLength = 10
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/about", About)
	myRouter.HandleFunc("/test", Test)
	myRouter.HandleFunc("/hello/{name}", Hello)
	myRouter.HandleFunc("/hostname", Myhostname)
	log.Fatal(http.ListenAndServe(":8080", myRouter)) //Port ==> 23450
}

//RandStringBytesMask returns random alphabetical string to be used as a
//transaction code.
func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

//myAPIStructure structure to pass to JSON NewEncoder
type myAPIStructure struct {
	Code          string `json:"code"`
	Message       string `json:"message"`
	MessageFormat string `json:"message_format"`
}

// GetOutboundIP returns the primary outbound IP address of the host or container
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

// Myhostname returns the ReST API response from http://localhost:8080/hostanme
func Myhostname(w http.ResponseWriter, r *http.Request) {
	ip := GetOutboundIP().String()
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	response := myAPIStructure{
		Code:          RandStringBytesMask(codeLength),
		Message:       name + " [" + ip + "]",
		MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

//Test returns the ReST API response from http://localhost:8080/test
func Test(w http.ResponseWriter, r *http.Request) {
	response := myAPIStructure{
		Code:          RandStringBytesMask(codeLength),
		Message:       "Testing 123. This is a test. Just a test.",
		MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

//About returns the ReST API response from http://localhost:8080/about
func About(w http.ResponseWriter, r *http.Request) {
	response := myAPIStructure{
		Code:          RandStringBytesMask(codeLength),
		Message:       "Docker Test ReST API container.",
		MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}

//Hello returns the ReST API response from http://localhost:8080/hello/younamehere
func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := myAPIStructure{
		Code:          RandStringBytesMask(codeLength),
		Message:       "Hello! " + name,
		MessageFormat: "text"}
	json.NewEncoder(w).Encode(response)
}
