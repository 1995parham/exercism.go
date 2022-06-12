#!/bin/bash
set -e

echo "update each excercise to its latest"

find . -maxdepth 1 -mindepth 1 -type d -not -path '*/[.]*' -printf '%f\n' | xargs -I{} exercism download --exercise='{}' --track=go
