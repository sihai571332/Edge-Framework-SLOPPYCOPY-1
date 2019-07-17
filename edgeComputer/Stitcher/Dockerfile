FROM jjanzic/docker-python3-opencv
#FROM ubuntu

RUN apt-get update && apt-get install -y python3-pip
RUN apt-get update && apt-get install -y python3
RUN apt-get update && apt-get install bash
RUN pip3 install opencv-python

RUN apt-get update && apt-get install -y libsm6 libxext6
RUN apt-get update && apt-get install -y libxrender-dev
RUN apt-get update && apt-get install -y xvfb


RUN pip3 install pillow
RUN pip3 install imutils
COPY startMe.sh /
COPY t.py /

CMD ["bash"]

CMD ["bash","startMe.sh"]

#ENTRYPOINT ["bash"]