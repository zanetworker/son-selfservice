#!/bin/bash

sleep 5
if curl http://selfservice-frontend:5000; then
  echo "Tests passed!"
  exit 0
else
  echo "Tests failed!"
  exit 1
fi
