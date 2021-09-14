#!/bin/bash

if [-n "$(gofmt -l .)"] then
  echo "GO code is not formatted:"
  gofmt -d .
  exit 1
fi