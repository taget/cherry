#!/bin/bash
# Author: Eli Qiao <qiaoliyong@gmail.com>

set -x
html=$(find /home/taget/shifei.me/ -type f -name *.html)

for i in $html; do
    ./html2pdf  $i
done
