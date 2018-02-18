#! /bin/bash

go build startdhtnode.go
./startdhtnode 8403
echo "\x1A"
