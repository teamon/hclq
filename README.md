# `hclq` - HCL + JQ

Read HCL from stdin and pass to [jq](https://stedolan.github.io/jq/).

## Build

```bash
make
```

## Usage

```
Usage: hclq [jq options] < file.hcl

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
```

## References
- https://stedolan.github.io/jq/manual
- https://github.com/tmccombs/hcl2json
