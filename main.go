package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var version string
var commit string

const usage = `Usage: hclq [jq options] < file.hcl

hclq - Convert HCL into JSON

# Example: Basic HCL to JSON conversion (no jq)
$ echo 'x { y = "z" }' | hclq
{"x":{"y":"z"}}

# Example: Pass JSON to jq
$ echo 'x { y = "z" }' | hclq '.'
{
  "x": {
    "y": "z"
  }
}

# Example: Pass query to jq
$ echo 'x { y = "z" }' | ./hclq '.x.y'
"z"

# Example: Pass query and options to jq
$ echo 'x { y = "z" }' | ./hclq -r '.x.y'
z

# References
jq - https://stedolan.github.io/jq/manual/
`

func main() {
	// Quick & dirty version & help check
	// Not using `flag` to allow options and args passthrough to jq
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h":
			fallthrough
		case "--help":
			fmt.Println(usage)
			return
		case "-v":
			fallthrough
		case "--version":
			fmt.Printf("%s (%s)\n", version, commit)
			return
		}
	}

	logger := log.New(os.Stderr, "", 0)

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		logger.Fatalf("Failed to read file: %s\n", err)
	}

	file, diags := hclsyntax.ParseConfig(bytes, "file.hcl", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		logger.Fatalf("Failed to read file: %s\n", diags)
	}

	content, err := convertFile(file)

	if err != nil {
		logger.Fatalf("Failed to convert file: %v", err)
	}

	jb, err := json.Marshal(content)

	if err != nil {
		logger.Fatalf("Failed to generate JSON: %v", err)
	}

	if len(os.Args) < 2 {
		os.Stdout.Write(jb)
	} else {
		cmd := exec.Command("jq", os.Args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		stdin, err := cmd.StdinPipe()
		if err != nil {
			logger.Fatalf("Failed to get stdin pipe: %v", err)
		}

		cmd.Start()
		stdin.Write(jb)
		stdin.Close()
		cmd.Wait()
	}
}
