#!/bin/bash

go build
go install
sudo chown root ../../../../bin/network-scanner
sudo chmod ugo+s ../../../../bin/network-scanner
