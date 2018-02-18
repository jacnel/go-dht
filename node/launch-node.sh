#! /bin/bash

cd ~/go/src/dht/node/
go build ./startdhtnode.go
./startdhtnode 8403
