package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var HOST string

func main() {
	HOST = os.Args[2]
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

	outputPath := os.Args[3]
	os.MkdirAll(outputPath, os.ModePerm)

	docs := buildPaginatedIndexDocuments(resourceType, header, rows, 5)
	for i, doc := range docs {
		bytes, err := json.Marshal(doc)
		if err != nil {
			log.Fatal(err)
		}
		var out string
		// if i == 0, we build 2 files: type.json and type-0.json for pagination
		if i == 0 {
			out = filepath.Join(outputPath, fmt.Sprintf("%s.json", resourceType))
			err = ioutil.WriteFile(out, bytes, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		out = filepath.Join(outputPath, fmt.Sprintf("%s-%d.json", resourceType, i))
		err = ioutil.WriteFile(out, bytes, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	objDocs := buildObjectDocuments(resourceType, header, rows)
	for _, doc := range objDocs {
		bytes, err := json.Marshal(doc)
		if err != nil {
			log.Fatal(err)
		}
		outDir := filepath.Join(outputPath, resourceType)
		os.MkdirAll(outDir, os.ModePerm)
		out := filepath.Join(outDir, fmt.Sprintf("%s.json", doc.Data[0].ID))
		err = ioutil.WriteFile(out, bytes, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// build rewrites
	rewrites := fmt.Sprintf("/%s.json page=:p /%s-:p.json 200!\n", resourceType, resourceType)
	ioutil.WriteFile(filepath.Join(outputPath, "_redirects"), []byte(rewrites), os.ModePerm)
}

type Document struct {
	Meta  map[string]interface{} `json:"meta,omitempty"`
	Links map[string]string      `json:"links,omitempty"`
	Data  []Object               `json:"data"`
}

type Object struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes map[string]string `json:"attributes"`
}

func buildPaginatedIndexDocuments(objType string, header []string, rows [][]string, pageSize int) []Document {
	var docs []Document

	var objs []Object
	for _, row := range rows {
		obj := buildObject(objType, header, row)

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
			"total-pages": len(docs),
		}
	}
	// set links
	for i := range docs {
		links := make(map[string]string)
		if i > 0 {
			links["prev"] = fmt.Sprintf("%s/%s.json?page=%d", HOST, objType, i-1)
		}
		if i < len(docs)-1 {
			links["next"] = fmt.Sprintf("%s/%s.json?page=%d", HOST, objType, i+1)
		}

		links["first"] = fmt.Sprintf("%s/%s.json?page=0", HOST, objType)
		links["last"] = fmt.Sprintf("%s/%s.json?page=%d", HOST, objType, len(docs)-1)
		docs[i].Links = links
	}

	return docs
}

func buildObjectDocuments(objType string, header []string, rows [][]string) []Document {
	var docs []Document

	for _, row := range rows {
		obj := buildObject(objType, header, row)
		doc := Document{Data: []Object{obj}}
		docs = append(docs, doc)
	}

	return docs
}

func buildObject(objType string, header []string, row []string) Object {
	obj := Object{Type: objType}
	kv := row2map(header, row)
	obj.ID = kv["id"]
	delete(kv, "id")
	obj.Attributes = kv

	return obj
}

// row2map normalize each columns name and build a map for the row
func row2map(header []string, row []string) map[string]string {
	r := make(map[string]string)
	for i, h := range header {
		r[strings.ToLower(h)] = row[i]
	}
	return r
}
