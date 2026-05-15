#!/bin/bash
# Needed for recursive glob to work
shopt -s globstar
basedir=$(dirname "$0")
source $basedir/.kickoff.conf
if [[ "$APP_NAME" == "go-application-template" || "$APP_NAME" == "" ]]; then
    echo "Please change APP_NAME in $basedir/.kickoff.conf to the name of your application."
    exit 1
fi

export files="$(grep -rl go-application-template | grep -Ev '^\.kickoff\.')"

# Replace names and references in files
for f in $files; do
    echo "Processing $basedir/$f"
    sed "s/go-application-template/$APP_NAME/g" -i $basedir/$f
done
