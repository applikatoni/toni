#!/bin/bash
# Builds and packages the toni binary

set -e

version=$(cat toni.go | grep "VERSION\s*=" | awk '{print $NF}' | sed 's/\"//g')
target="toni-${version}-$(go env GOOS)-$(go env GOARCH)"
executable="toni"

rm -rf ./builds/$target
mkdir ./builds/$target

go build -o ./builds/$target/$executable ./ || exit 1

cp ./README.md ./builds/$target/
cp ./LICENSE ./builds/$target/
cp ./sample.toni.yml ./builds/$target/
echo ${version} >> ./builds/$target/VERSION

tar czvfC ./builds/$target.tar.gz ./builds $target

rm -rf ./builds/$target
