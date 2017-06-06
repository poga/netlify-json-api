package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	filename := path.Base(os.Args[1])
	resourceType := filename[0 : len(filename)-len(filepath.Ext(filename))]

	r := csv.NewReader(f)
	// read header
	header, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	docs := buildPaginatedIndexDocuments(resourceType, header, rows, 5)
	fmt.Printf("%v\n", docs)
	bytes, err := json.Marshal(docs)
	fmt.Println(string(bytes[:]))
}

type Document struct {
	Meta  map[string]interface{} `json:"meta"`
	Links map[string]string      `json:"links"`
	Data  []Object               `json:"data"`
}

type Object struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes map[string]string `json:"attributes"`
}

func buildPaginatedIndexDocuments(objType string, header []string, rows [][]string, pageSize int) *[]Document {
	var docs []Document

	var objs []Object
	for _, row := range rows {
		obj := Object{Type: objType}
		kv := row2map(header, row)
		obj.ID = kv["id"]
		delete(kv, "id")
		obj.Attributes = kv

		objs = append(objs, obj)

		if (len(objs)) == pageSize {
			doc := Document{}
			doc.Data = objs
			docs = append(docs, doc)
			objs = make([]Object, 0)
		}
	}

	if len(objs) > 0 {
		doc := Document{}
		doc.Data = objs
		docs = append(docs, doc)
		objs = make([]Object, 0)
	}
	// set meta
	for i := range docs {
		docs[i].Meta = map[string]interface{}{
			"total": len(docs),
		}
	}
	// set links
	for i := range docs {
		links := make(map[string]string)
		if i > 0 {
			links["prev"] = fmt.Sprintf("/%s.json?page=%d", objType, i-1)
		}
		if i <= len(docs)-1 {
			links["next"] = fmt.Sprintf("/%s.json?page=%d", objType, i+1)
		}

		links["first"] = fmt.Sprintf("/%s.json?page=0", objType)
		links["last"] = fmt.Sprintf("/%s.json?page=%d", objType, len(docs)-1)
		docs[i].Links = links
	}

	return &docs
}

// row2map normalize each columns name and build a map for the row
func row2map(header []string, row []string) map[string]string {
	r := make(map[string]string)
	for i, h := range header {
		r[strings.ToLower(h)] = row[i]
	}
	return r
}
