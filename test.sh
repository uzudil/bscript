#!/bin/bash
go build -o runner
rm -f test.log
for f in src/tests/*.b; do
    echo "   Running test: $f"
    ./runner -source $f >> test.log 2>&1
    if [ $? -ne "0" ]; then
        echo "   *** Failed $f, see test.log for details"
        exit 1
    fi
done
echo "Success."
