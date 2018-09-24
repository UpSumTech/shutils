#! /usr/bin/env bash

# go to home directory
cd ~
# same as above
cd /Users/sumanmukherjee

# go to root dir
cd /

# Create a file foo.txt
touch ~/code/foo.txt

# list files like foowhatever
ls ~/code/foo*
# list files like whatever.txt
ls ~/code/*.txt
# same as above
ls /Users/sumanmukherjee/code/*.txt

# Add some text to foo.txt
echo 'first line' >> ~/code/foo.txt
echo 'second line' >> ~/code/foo.txt

# Download a file and specify a new filename
curl http://example.com/file.html -o new_file.html

# Get a file from internet on your standard out
curl http://example.com/file.html

# Find file in current dir
find . -type f -name '*.txt'

# Find file in current /Users/sumanmukherjee/code
find /Users/sumanmukherjee/code -type f -name '*.txt'

# Find file in current /Users/sumanmukherjee/code named like foo-whatever.txt
find /Users/sumanmukherjee/code -type f -name 'foo*.txt'

# Find dir in current /Users/sumanmukherjee/code
find /Users/sumanmukherjee/code -type d -name 'example*'

# Filter for BCIT in file.txt
cat /Users/sumanmukherjee/code/file.txt | grep 'BCIT'
# This is the same as above only with ~ in place of /Users/sumanmukherjee
cat ~/code/file.txt | grep 'BCIT'

# Get the first 5 lines of file.txt
cat ~/code/file.txt | head -n 5
# Get the first 5 lines of file.txt and put it in head.txt
cat ~/code/file.txt | head -n 5 > head.txt

# Get the last 5 lines of file.txt
cat ~/code/file.txt | tail -n 5
# Get the last 5 lines of file.txt and put it in tail.txt
cat ~/code/file.txt | tail -n 5 > tail.txt

# get the 4th and 5th line of file.txt
cat ~/code/file.txt | head -n 5 | tail -n 2

# Get the 4th and 5th line that matches foo from file.txt
cat ~/code/file.txt | grep 'foo' | head -n 5 | tail -n 2

# Get a file from internet, find text whatever and get the last 10 lines from that search into newtail.txt
curl http://example.com/file.html | grep 'whatever' | tail -n 10 > newtail.txt

# zip up the dir to compressed.zip
zip -r compressed.zip /path/to/dir
