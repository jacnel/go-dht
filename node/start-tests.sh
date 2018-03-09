!# /bin/bash

printf "cd go/src/dht/node/\n./testnode $1 $2 $3 0\n" | ssh -i dht-testing.pem ec2-user@ec2-107-23-129-34.compute-1.amazonaws.com &

printf "cd go/src/dht/node/\n./testnode $1 $2 $3 1\n" | ssh -i dht-testing.pem ec2-user@ec2-54-173-150-90.compute-1.amazonaws.com &

printf "cd go/src/dht/node/\n./testnode $1 $2 $3 2\n" | ssh -i dht-testing.pem ec2-user@ec2-54-87-229-45.compute-1.amazonaws.com &

printf "cd go/src/dht/node/\n./testnode $1 $2 $3 3\n" | ssh -i dht-testing.pem ec2-user@ec2-34-228-83-193.compute-1.amazonaws.com &
