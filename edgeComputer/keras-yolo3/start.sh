#!/bin/bash

docker build -t flow .

cd .. && cd Memory1

docker run --net=host -ti -v "$(pwd):/mydata" flow
