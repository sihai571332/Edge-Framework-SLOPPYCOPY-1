#!/bin/bash



docker build -t broadcaster .

cd ../../shared_folder

docker run --network host -v "$(pwd):/go/mydata" broadcaster