FROM tensorflow/tensorflow:latest

RUN apt-get update
RUN DEBIAN_FRONTEND='noninteractive' apt-get install -y --no-install-recommends software-properties-common

RUN add-apt-repository ppa:deadsnakes/ppa
RUN apt-get update
RUN DEBIAN_FRONTEND='noninteractive' apt-get install -y --no-install-recommends python3 python3-pip python3-setuptools 
RUN pip3 install --upgrade pip

RUN apt-get install bash

RUN apt-get update && apt-get install -y libsm6 libxext6
RUN apt-get update && apt-get install -y libxrender-dev
RUN apt-get update && apt-get install -y xvfb

RUN pip3 install tensorflow keras opencv-python image matplotlib
ADD . /src

WORKDIR /src

ENTRYPOINT ["bash","detect.sh"]