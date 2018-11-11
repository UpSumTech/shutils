#! /usr/bin/env bash

strace ls -i foo # Trace the system calls for the command
ltrace ls -i foo # Trace the library calls for the command
