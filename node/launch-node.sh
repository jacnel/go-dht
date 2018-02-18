#! /bin/bash

go build startdhtnode.go
printf "\x1A" | ./startdhtnode 8403
