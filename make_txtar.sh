#!/usr/bin/env bash

set -e

tmp=$(mktemp -d dentwatcher.XXXXX)

if [ -z "${tmp+x}" ] || [ -z "$tmp" ]; then
    echo "Error: \$tmp is not set or is an empty string."
    exit 1
fi

{
    rg --files . \
        | grep -v $tmp/filelist.txt \
        | grep -vE 'dentwatcher$' \
        | grep -v README.org \
        | grep -v make_txtar.sh \
        | grep -v go.sum \
        | grep -v go.mod \
        | grep -v Makefile \
        | grep -v cmd/main.go \
        | grep -v logger.go \
        # | grep -v dentwatcher.go \

} | tee $tmp/filelist.txt
tar -cf $tmp/dentwatcher.tar -T $tmp/filelist.txt
mkdir -p $tmp/dentwatcher
tar xf $tmp/dentwatcher.tar -C $tmp/dentwatcher
rg --files $tmp/dentwatcher
txtar-c $tmp/dentwatcher | pbcopy

rm -rf $tmp
