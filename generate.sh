#!/bin/bash

svcs="./svc"
builds="../build"

MSGHOME="$HOME/src/go/src/github.com/RobotStudio/choreo-msg/msg"

includes="-I$MSGHOME -I."

[ "$#" -gt 0 ] || { echo "Please specify a language"; exit 1; }

PROTOC="$(which protoc)"
[ -f "$PROTOC" ] || { echo "Do you have protoc installed?"; exit 1; }

cd $svcs

lang_opts=
for lang in $*; do
  lang_opts="$lang_opts --${lang}_out=$builds/$lang"
  [ -d "${builds}/$lang" ] || mkdir -p "${builds}/$lang"
done

find . -iname '*.proto' -type f | while read -r f; do
  $PROTOC $includes $lang_opts "$f"
done

exit 0
