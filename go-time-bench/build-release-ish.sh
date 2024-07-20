#!/bin/bash

go build -ldflags "-s -w" -o release
time ./release