package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func main() {
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
