#!/bin/bash
export DEST="./.exvim.TR"
export TOOLS="/home/user/exvim//vimfiles/tools/"
export TMP="${DEST}/_symbols"
export TARGET="${DEST}/symbols"
sh ${TOOLS}/shell/bash/update-symbols.sh
