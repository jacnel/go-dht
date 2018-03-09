#! /bin/bash

printf "cd go/src/dht/node/\ngit pull\ngo build testnode.go\ngo build startdhtnode.go\nexit\n" | ssh -i dht-testing.pem ec2-user@ec2-107-23-129-34.compute-1.amazonaws.com

printf "cd go/src/dht/node/\ngit pull\ngo build testnode.go\ngo build startdhtnode.go\nexit\n" | ssh -i dht-testing.pem ec2-user@ec2-54.173.150.90.compute-1.amazonaws.com

printf "cd go/src/dht/node/\ngit pull\ngo build testnode.go\ngo build startdhtnode.go\nexit\n" | ssh -i dht-testing.pem ec2-user@ec2-54-87-229-45.compute-1.amazonaws.com

printf "cd go/src/dht/node/\ngit pull\ngo build testnode.go\ngo build startdhtnode.go\nexit\n" | ssh -i dht-testing.pem ec2-user@ec2-34-228-83-193.compute-1.amazonaws.com
