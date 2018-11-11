#!/usr/bin/env bash

# Try to see if port is accepting connections on remote machine
telnet 54-190-204-211.us-west-2.compute.amazonaws.com 22

# To debug init logs
cat /var/log/cloud-init-output.log # To see the logs of the bootstrap process on the NAT box
cat /var/log/cloud-init.log # To see the init logs of the ec2 box
cat /var/lib/cloud/data/status.json # To view the current status of the ec2 box
cat /var/lib/cloud/data/result.json # To view the result of the init process

# To get instance metadata from inside the instance
curl http://169.254.169.254/latest/
curl http://169.254.169.254/latest/dynamic/instance-identity/document # This gives the instance identity data
curl http://169.254.169.254/latest/meta-data/ # This gives the meta-data of the instance

# installing certbot to generate letencrypt certs on an ec2 machine that you can use for your domain
add-apt-repository ppa:certbot/certbot # add the ppa to apt
apt-get install -y -qq certbot python3-certbot-dns-route53 # install certbot and the certbot plugin to validate route53 domains

# Generate the certs with certbot
certbot certonly --dns-route53 --expand --noninteractive --agree-tos --email developer@example.com -d example.com