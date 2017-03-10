#!/bin/bash
export DEST="./.exvim.TR"
export TOOLS="/home/user/exvim//vimfiles/tools/"
export TMP="${DEST}/_ID"
export TARGET="${DEST}/ID"
sh ${TOOLS}/shell/bash/update-idutils.sh
