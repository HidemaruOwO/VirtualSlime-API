#!/bin/bash

url="https://v-sli.me/data.ts"

cache_dir="cache"

if [ ! -d "$cache_dir" ]; then
    mkdir "$cache_dir"
fi

curl -o "$cache_dir/data.ts" "$url"

echo "ファイルがダウンロードされました。"

