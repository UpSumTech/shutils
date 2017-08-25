#! /usr/bin/env bash

ls -lah # Display all files in current dir in human readable format
ls -ltr # Display files newest last
du -sh ~/**/* | sort -rn # Display dirs in sorted sizes
df -h # Display free disk space in human readable format
