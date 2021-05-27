package main

import (
	"os"
	"bytes"
	"io/ioutil"
	"path/filepath"
	"html/template"
	"fmt"
	"log"
)

var outDir = "./build"
var docPath = "./src/index.html"

// @todo: break into small helpers

func main()  {

	ctx, readError := ioutil.ReadFile(docPath)

	if readError != nil {
		log.Fatal(readError)
	}

	ext := filepath.Ext(docPath)

	if makeError := os.MkdirAll(outDir, os.ModePerm); makeError != nil {
		log.Fatal(makeError)
	}

	outSrc := filepath.Join(outDir, "index.html") 

	outDoc, pathError := os.Create(outSrc)

	defer outDoc.Close()

	if pathError != nil {
		log.Fatal(pathError)
	} 

	sArray := []string{}

	for _, c := range ctx {
		sArray = append(sArray, string(c))
	}

	fmt.Println("ctx: ", sArray, outSrc, ext) // rm

	var buf bytes.Buffer

	if _, bigError := buf.Write(ctx); bigError != nil {
		log.Fatal(bigError)
	}

	tmp := template.New(filepath.Base(docPath))

	tmp, parseError := tmp.Parse(buf.String())

	if parseError != nil {
		log.Fatal(parseError)
	}

	// @todo: create template context https://github.com/julvo/tinystatic/blob/b6e07748b0c5d9b7a50c2edad4ca97fd6f3705a8/route.go#L180

	// tmpCtx := map[string]interface{}{}
	// tmpCtx["test"] = struct{fname, lname string}{"Bruce", "Lee"}

	buf.Reset()

	// @params: (wr io.Writer, d interface{}) https://golang.org/pkg/text/template/#Template.Execute
	if err :=  tmp.Execute(outDoc, []interface{}{}); err != nil {
		log.Fatal(err)
	}
	
}


// Refs
// tinystatic https://github.com/julvo/tinystatic/blob/master/route.go