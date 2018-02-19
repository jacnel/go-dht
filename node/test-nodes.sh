#! /bin/bash

sh clear-tables.sh

go build testnode.go
sleep 5

./testnode 40 5000 10
sh clear-tables.sh

./testnode 40 5000 100
sh clear-tables.sh

./testnode 40 5000 1000
sh clear-tables.sh

./testnode 40 5000 10000
sh clear-tables.sh
