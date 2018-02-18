#! /bin/bash

printf "cd go/src/dht/node/\nsh launch-node.sh &\n" | ssh -i dht-testing.pem ec2-user@ec2-54-208-29-162.compute-1.amazonaws.com


printf "cd go/src/dht/node/\nsh launch-node.sh &\n" | ssh -i dht-testing.pem ec2-user@ec2-54-211-127-45.compute-1.amazonaws.com

printf "cd go/src/dht/node/\nsh launch-node.sh &\n" | ssh -i dht-testing.pem ec2-user@ec2-75-101-226-165.compute-1.amazonaws.com

printf "cd go/src/dht/node/\nsh launch-node.sh &\n" | ssh -i dht-testing.pem ec2-user@ec2-34-233-120-248.compute-1.amazonaws.com

go build testnode.go

./testnode 4 5000 10
sh clear-tables.sh

./testnode 4 5000 100
sh clear-tables.sh

./testnode 4 5000 1000
sh clear-tables.sh

./testnode 4 5000 10000
sh clear-tables.sh

printf "sudo reboot" | ssh -i dht-testing.pem ec2-user@ec2-54-208-29-162.compute-1.amazonaws.com

printf "sudo reboot" | ssh -i dht-testing.pem ec2-user@ec2-54-211-127-45.compute-1.amazonaws.com

printf "sudo reboot" | ssh -i dht-testing.pem ec2-user@ec2-75-101-226-165.compute-1.amazonaws.com

printf "sudo reboot" | ssh -i dht-testing.pem ec2-user@ec2-34-233-120-248.compute-1.amazonaws.com
