package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"git.sr.ht/~jcmuller/strval"
	"gopkg.in/yaml.v3"
)

func main() {
	d, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	if e := yaml.Unmarshal(d, &data); e != nil {
		log.Fatal(e)
	}

	out, e := strval.Marshal(data)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("%s\n", out)
}
