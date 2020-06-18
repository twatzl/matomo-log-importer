# Matomo Log Importer

The official Matomo log import script is written in Python 2.7 (deprecated), hard to read and does not work 
with access logs from Traefik. Instead of trying to understand and fix the code I decided to write my own importer in Go. 

## Current Status

Project is freshly started and in development. Not ready for use.

## Goals

Provide an application that can be used to read access.log files from Traefik and add the tracking data to Matomo.

Although I will try to keep open the possibility to add other formats filters and so on I will only implement what I
need for my own use case.

Contributions are welcome though.

