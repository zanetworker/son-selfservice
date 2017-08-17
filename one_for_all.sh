#!/bin/bash

BUILD_PATH=./selfservice-backend
cd $BUILD_PATH && chmod +x build.sh && ./build.sh

if [ $? -eq 0 ]
  then
    echo "Done Building Backend Image!"
    echo "Building the whole project"
    docker-compose up --build
  else
     echo "Something went wrong building back-end base image!"
    #statements
fi


# source build.sh
