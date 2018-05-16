#! /usr/bin/env bash

# To debug init logs
cat /var/log/cloud-init-output.log # To see the logs of the bootstrap process on the NAT box
cat /var/log/cloud-init.log # To see the init logs of the ec2 box
cat /var/lib/cloud/data/status.json # To view the current status of the ec2 box
cat /var/lib/cloud/data/result.json # To view the result of the init process

# To get instance metadata from inside the instance
curl http://169.254.169.254/latest/
curl http://169.254.169.254/latest/dynamic/instance-identity/document # This gives the instance identity data
curl http://169.254.169.254/latest/meta-data/ # This gives the meta-data of the instance
