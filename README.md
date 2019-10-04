# `hclq` - HCL + JQ

Read HCL from stdin and pass to [jq](https://stedolan.github.io/jq/).

## Build

```bash
go build
```

## Usage

```bash
hclq '.some.field' < file.hcl
```
