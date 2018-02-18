#! /bin/bash

go build startdhtnode.go
echo "\x1A" | ./startdhtnode 8403
