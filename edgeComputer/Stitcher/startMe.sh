#!/bin/bash

counter=1

while [ 1 -gt 0 ] ; do

    file="/mydata/GO1-$counter.txt"
    if [ -e "$file" ] ; then

        rm /mydata/GO1-$counter.txt
        cd mydata
        xvfb-run python3 ../t.py -i ProcessImage -o $counter-A.jpg
        cd ProcessImage
	rm -fv *.jpg
	cd ..
        touch FOO-$counter.txt
        let counter=$counter+1

    fi

done
