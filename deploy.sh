#!/bin/bash

echo "=============================="
echo -n "Running tests... "
go test -v ./tests/...

if [ $? -ne 0 ]; then
    echo "failed"
    exit
fi

echo "=============================="
echo -n "Building... "
go build -o build/telegramud
echo "done"
echo "=============================="
echo "Stopping remote application... "
ssh vds screen -X -S mud quit
echo "done"
echo "=============================="
echo "Deploying... "
scp build/telegramud vds:~
echo "done"
echo "=============================="
echo -n "Restarting application... "
echo "failed"
echo "=============================="
