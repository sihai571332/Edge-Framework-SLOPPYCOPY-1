#!/bin/bash

docker build -t send .

cd ../shared_folder

docker run -it --network host -v "$(pwd):/shared" send