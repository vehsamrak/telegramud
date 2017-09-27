#!/usr/bin/env bash

case $1 in
'')
    $0 start
;;

'reload')
    echo "Reloading application ..."
    docker exec -it telegramud_application pkill game
    docker restart telegramud_application
;;

'restart')
    $0 stop
    $0 start
;;

'start')
    echo "Starting containers ..."
    docker-compose -f docker-compose.yml -f docker-compose-local.yml up -d
;;

'stop')
    echo "Stopping containers ..."
    docker-compose -f docker-compose.yml -f docker-compose-local.yml down
;;

'migrate')
    echo "Running migrations ..."
    docker-compose -f docker-compose.yml -f docker-compose-local.yml run application go run -v game.go -migrate
;;

'build')
    echo "Building containers ..."
    docker-compose -f docker-compose.yml -f docker-compose-local.yml up -d --build
;;

esac
