package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

func main() {
	var vals = flag.String("values", "", "Values JSON data")
	var run = flag.String("run", "", "Values JSON data")
	flag.Parse()

	log.Println("Vals:", *vals)
	log.Println("Run:", *run)

	var valsObjMap map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(*vals), &valsObjMap); err != nil {
		log.Fatalf("vals unmarshalling err: %v", err)
	}

	var runObjMap map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(*run), &runObjMap); err != nil {
		log.Fatalf("run unmarshalling err: %v", err)
	}

	fmt.Println("Deserializing Values...")
	valsOut, err := json.Marshal(valsObjMap)
	fmt.Println(string(valsOut), "Err:", err)

	fmt.Println("Deserializing Run...")
	runOut, err := json.Marshal(runObjMap)
	fmt.Println(string(runOut), "Err:", err)
}
