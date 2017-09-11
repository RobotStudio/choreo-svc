#!/bin/bash

src="./src"
builds="../svc"
langs=(
  "go"
  "cpp"
  "python"
  "java"
)

proto_deps=(
  "github.com/RobotStudio/choreo-msg"
)


[ "$#" -gt 0 ] && langs=("$@")
#[ "$#" -gt 0 ] || { echo "Please specify a language"; exit 1; }

PROTOC="$(which protoc)"
[ -f "$PROTOC" ] || { echo "Do you have protoc installed?"; exit 1; }

cd "$src" || { echo "Could not cd into $src"; exit 1; }

[ -d "$builds" ] && rm -rf "$builds"
mkdir -p "$builds"

# Assemble lang params
lang_opts=
for lang in "${langs[@]}"; do
  lang_opts="$lang_opts --${lang}_out=$builds"
done

# Assemble include params
deps=.
for dep in "${proto_deps[@]}"; do
  deps="$deps:${GOPATH}/src/${dep}/src"
done

# Find and build
find . -iname '*.proto' -type f | while read -r f; do
  echo "$PROTOC" "-I$deps" "$lang_opts" "$f"
  # shellcheck disable=SC2086
  "$PROTOC" -I"$deps" $lang_opts "$f"
done

exit 0
