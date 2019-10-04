#!/bin/bash

set -eufo pipefail

res=$(echo 'a "b" {   c = 42 }' | ./hclq '.a.b.c')

[[ "${res}" == "42" ]] || echo "ERROR: ${res} not equal to 42"
