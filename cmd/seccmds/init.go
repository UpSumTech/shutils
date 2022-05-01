package seccmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	parseShortDesc = `Prints examples of commands related to encryption or ssl etc`
	parseLongDesc  = `Prints examples of commands related to encryption or ssl etc`
	parseExample   = `
	### Example commands for certs and encryption etc
	shutils sec`
)

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "sec [no options!]",
		Short:            parseShortDesc,
		Long:             parseLongDesc,
		Example:          parseExample,
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`
runlevel # Get the current run level of the system
uptime # How long has this machine been running
who # who is currently logged into the system
w # List processes belonging to logged in users

last # List last logged in users
last reboot # List last time the system restarted
last root # List last time the system was accessed as root

wall 'Taking down system for maintenance' # Send a warning to other users to wrap up their work before maintenance window begins

# Scan for binaries that have suid bit set for user and group
find / -xdev -type f -perm /u+s,g+s -print # Useful for finding binaries you dont recognize

# Scenario : Running lsof shows deleted uid 999 hanging onto processes
# find user with UID 999
awk -v uid=999 -F ":" '$3==uid {print $1}' /etc/passwd
ps -U 999 # find processes owned by user 999

ls -i # List inode or physical addresses of files. Sometimes useful to find files with weird chars in names
zdump PST EST IST # list current time in different time zones

# Certs with CA
# openssl generate private key and CSR for example.com
openssl req -newkey rsa:2048 -nodes -keyout example.key -out example.csr
# only generate the CSR if the private key is already there
openssl req -key example.key -new -out example.csr
# use the private key and cert to generate the CSR to renew your certs
openssl x509 -in example.crt -signkey example.key -x509toreq -out example.csr

# Self signed certs
# openssl command to create private key and self signed cert
openssl req -newkey rsa:2048 -nodes -keyout example.key -x509 -days 365 -out example.crt
# openssl create self signed cert valid for 365 days from existing private key
openssl req -key example.key -new -x509 -days 365 -out exmaple.crt
# openssl create self signed cert from existing private key and CSR
openssl x509 -signkey example.key -in example.csr -req -days 365 -out example.crt

# check ssl connection with debug info
openssl s_client -debug -msg -connect example.com:80
# check cert expiry date for cert
openssl x509 -enddate -noout -in cert.pem
# view and verify contents of cert
openssl x509 -text -noout -verify -in example.crt
# get details of a cert
openssl x509 -enddate -startdate -email -subject -issuer -noout -in cert.pem
# get the details of a CSR (cert signing request)
openssl req -text -noout -verify -in example.csr
# verify that the certificate was signed by a particular CA
openssl verify -verbose -CAFile ca.crt example.crt
# verify that the private key mentioned below is a valid key
openssl rsa -check -in example.key

# convert PEM format key and cert to PKCS12 format
openssl pkcs12 -inkey private.key -in example.crt -export -out example.p12
# convert PKCS12 format to PEM format combined key and cert file
openssl pkcs12 -in example.p12 -nodes -out combined-example.crt

# keystore creations for use in java services
# generate server cert and save in a key store
keytool -genkeypair -alias server-keystore -keyalg RSA -keysize 4096 -keystore server-keystore.ks
# export server cert into a pem file
keytool -export -alias server-keystore -keystore server-keystore.ks -file server-cert.pem
# generate client keystore to e used in java services
keytool -genkey -alias client-keystore -keyalg RSA -keystore client-keystore.ks
# import server cert into client truststore
keytool -import -alias server-keystore -keystore client-truststore.ts -file server-cert.pem

# convert certs to pkcs file so that it can imported by ava keytools into jks format
openssl pkcs12 -export -inkey private-key.pem -in combined-chain-and-private-key.pem -name some-random-name -out some-random-name.p12
# import pkcs file into jks file using java keytools
keytool -importkeystore -srckeystore some-random-name.p12 -srcstoretype pkcs12 -destkeystore dest-keystore-to-use-with-java-services.jks
			`)
		},
	}

	return cmd
}
