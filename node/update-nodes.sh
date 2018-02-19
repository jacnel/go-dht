#! /bin/bash

printf "cd go/src/dht/node/\ngit pull\ngo build testnode.go\ngo build startdhtnode.go\nexit\n" | ssh -i dht-testing.pem ec2-user@ec2-54-208-29-162.compute-1.amazonaws.com

printf "cd go/src/dht/node/\ngit pull\ngo build testnode.go\ngo build startdhtnode.go\nexit\n" | ssh -i dht-testing.pem ec2-user@ec2-54-211-127-45.compute-1.amazonaws.com

printf "cd go/src/dht/node/\ngit pull\ngo build testnode.go\ngo build startdhtnode.go\nexit\n" | ssh -i dht-testing.pem ec2-user@ec2-75-101-226-165.compute-1.amazonaws.com 

printf "cd go/src/dht/node/\ngit pull\ngo build testnode.go\ngo build startdhtnode.go\nexit\n" | ssh -i dht-testing.pem ec2-user@ec2-34-233-120-248.compute-1.amazonaws.com 
