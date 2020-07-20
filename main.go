package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	valuesPtr := flag.String("values", "", "Values for template in YAML format")
	templatePtr := flag.String("template", "", "Template file to render")
	outputPtr := flag.String("output", "", "(Optional) Output file. If not given, prints to stdout")
	flag.Parse()

	if *valuesPtr == "" {
		log.Fatal("Must provide a values file.")
	}

	if *templatePtr == "" {
		log.Fatal("Must provide a template file.")
	}

	templateFile, err := ioutil.ReadFile(*templatePtr)
	if err != nil {
		fmt.Println("Reading template file error", err)
		return
	}

	t := template.Must(template.New("").Parse(string(templateFile)))

	varFile, err := ioutil.ReadFile(*valuesPtr)
	if err != nil {
		fmt.Println("Reading values file error", err)
		return
	}

	m := map[string]interface{}{}
	//if err := json.Unmarshal(varFile, &m); err != nil {
	if err := yaml.Unmarshal(varFile, &m); err != nil {
		panic(err)
	}

	if *outputPtr == "" {
		err = t.Execute(os.Stdout, m)
	} else {
		f, err := os.Create(*outputPtr)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		err = t.Execute(w, m)
		w.Flush()
		f.Sync()
	}

	if err != nil {
		panic(err)
	}
}
