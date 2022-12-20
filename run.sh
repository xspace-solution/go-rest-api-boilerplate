#!/bin/bash

echo "Installing and Compiling..."
go install .
echo "running Go app"
go run .
echo "App Stopped"
