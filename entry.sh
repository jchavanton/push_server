#!/bin/bash

if [ "$1" = "" ]; then
	CMD="/main"
else
        CMD="$*"
fi

echo "Running [$CMD]"
exec $CMD
echo "exiting ..."
