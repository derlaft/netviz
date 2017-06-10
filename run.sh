#!/bin/bash

ls /mnt/*.pcap | xargs -P4 -n1 bash run_single.sh
