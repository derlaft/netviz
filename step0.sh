#!/bin/bash

tshark -Tfields -r $1 -e frame.len -e frame.protocols -e frame.time_delta -e tcp.stream -e udp.stream | \
  while read len proto delta stream; do 
    echo $stream \
      $len \
      $delta \
      $(echo "ibase=16; $(echo "$proto" | md5sum | head -c2 | tr '[a-z]' '[A-Z]')" | bc -ql);
  done
