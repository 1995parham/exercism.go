default:
    @just --list

lint:
    #!/usr/bin/env bash

    for project in $(find . -maxdepth 1 -mindepth 1 -type d -not -path '*/[.]*'); do
      cd "$project"
      echo "$project"
      echo
      golangci-lint run --fix
      echo
      cd ..
    done

test:
    #!/usr/bin/env bash

    for project in $(find . -maxdepth 1 -mindepth 1 -type d -not -path '*/[.]*'); do
      cd "$project"
      echo "$project"
      echo
      go test -v
      echo
      cd ..
    done
