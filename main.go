// This tool translates annotated Go files into Pulumi component schema metadata, which is in JSON Schema. This
// allows a component author to begin by writing the same Go structures, which is often more natural. It also has
// the benefit of using these same structures in the implementation of the component itself.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// This tool simply takes an array of files to parse. These files must include only Go types of the
	// expected kinds: resource definitions and annotated struct types. It will issue an error for anything else.
	if len(os.Args) < 3 {
		log.Fatalf("error: usage: [PULUMI-PKG-NAME] [GO-SOURCE-PKG]")
	}
	sch, err := Generate(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	// Now serialize the schema into JSON and print it out.
	b, err := json.Marshal(sch)
	if err != nil {
		log.Fatalf("error: serializing schema to JSON: %s", err.Error())
	}

	fmt.Printf("%s\n", string(b))
}
