---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/guide-secure-unifi-cloud-controller-on-aws-lightsail-signed-with-lets-encrypt-ss-f5fdf716-1
title: "unifi_ssl_import.sh"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS", "Oracle"]
dates: ["2019-02"]
keywords: ["aws", "compute", "cost"]
source: docs/RAG/collect-261001-unifi-ubiquiti/guide-secure-unifi-cloud-controller-on-aws-lightsail-signed-with-lets-encrypt-ss-f5fdf716.md
source_anchor: ""
source_lines: [1, 158]
sha256: 95b40ede4f097d53c671d76cd34f44c05de63414d82c76d64a195ae013ec3555
---

# unifi_ssl_import.sh

I found a solution for how to navigate cloud key issues and wanted to set up a ZTP solution for Unifi hardware so I can direct ship equipment to the site, and provision it securely via internet without having to stand up a L2L tunnel.
Alright, lets get started…
This guide is applicable for any Ubuntu based install, but I’m going to utilize Amazon Lightsail for the demo, since at the time of writing, it’s the cheapest option I can find with enough compute resources and a static IP included.
2 GB RAM, 1 vCPU, 60 GB SSD
OPex (Recurring Cost) – 10$ per Month – As of February 2019
Guide
Dry Run
1. Set up LightSail instance
Login to – https://lightsail.aws.amazon.com
Spin up a Lightsail instance:
Set a name for the instance and provision it.
2. Create and attach static IP
Click on the instance name and click on the networking tab:
Click “Create Static IP”:
3. Open necessary ports
| TCP or UDP | Port Number | Usage | 
| TCP | 80 | Port used inform-URL for adoption. | 
| TCP | 443 | Port used for Cloud Access service. | 
| UDP | 3478 | Port used for STUN. | 
| TCP | 8080 | Port used for device and controller communication. | 
| TCP | 8443 | Port used for controller GUI/API as seen in a web browser. | 
| TCP | 8880 | Port used for HTTP portal redirection. | 
| TCP | 8843 | Port used for HTTPS portal redirection. | 
You can disable or lock down the ports as needed using IP-tables depending on your security posture
Post spotlight-
https://www.lammertbies.nl/comm/info/iptables.html#intr
4. Set up Unify packages
Add the Ubiquiti repository to /etc/apt/sources.list:
sudo echo "deb http://www.ubnt.com/downloads/unifi/debian stable ubiquiti" | sudo tee -a /etc/apt/sources.list
Add the Ubiquiti GPG Key:
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv 06E85760C0A52C50
Update the server’s repository information:
sudo apt-get update
Install JAVA 8 run time
You need Java Run-time 8 to run the UniFi Controller
Add Oracle’s PPA (Personal Package Archive) to your list of sources so that Ubuntu knows where to check for the updates. Use add–apt–repository command for that.
sudo add-apt-repository ppa:webupd8team/java -y sudo apt install java-common oracle-java8-installer
update your package repository by issuing the following command
sudo apt-get update
The oracle-java8-set-default package will automatically set Oracle JDK8 as default. Once the installation is complete we can check Java version.
java -version
java version "1.8.0_191"
MongoDB
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 0C49F3730359A14518585931BC711F9BA15703C6
echo "deb http://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.4.list
sudo apt update
Update. Retrieve the latest package information.
sudo apt update
sudo apt-get install apt-transport-https
Install UniFi Controller packages.
sudo apt install unifi
You should be able to Access the web interface and go through the initial setup wizard.
https://yourIPaddress:8443
5. Set up SSL using certbot and letsencrypt
Lets get that green-lock up in here shall we
So, a few things to note here… UniFi doesn’t really have a straightforward way to import certificates, you have to use the java keystore commands to import the cert, but there is a very handy script built by Steve Jenkins that makes this super easy.
First, we need to request a cert and sign it using lets encrypt certificate authority.
Let’s start with adding the repository and install the EFF certbot package – link
sudo apt-get update
sudo apt-get install software-properties-common 
sudo add-apt-repository universe 
sudo add-apt-repository ppa:certbot/certbot 
sudo apt-get update 
sudo apt-get install certbot
5.1 Update/add your DNS record and make sure its propagated (this is important)
Note - The DNS name should point to the static IP we attached to our light-sail instance
Im going to use the following A record for this example
unifyctrl01.multicastbits.com
Ping from the controller and make sure the server can resolve it.
ping unifyctrl01.multicastbits.com
You wont be able see any echo replies because ICMP is not allowed on the firewall rules in AWS - leave it as is we just need the server to see the IP resolving to DNS A record
5.2 Request the certificate
Issue the following command to start certbot in certonly mode
sudo certbot certonly
usage: 
  certbot [SUBCOMMAND] [options] [-d DOMAIN] [-d DOMAIN] ...
Certbot can obtain and install HTTPS/TLS/SSL certificates. By default,
it will attempt to use a webserver both for obtaining and installing the
certificate. The most common SUBCOMMANDS and flags are:
obtain, install, and renew certificates:
    (default) run   Obtain & install a certificate in your current webserver
    certonly        Obtain or renew a certificate, but do not install it
    renew           Renew all previously obtained certificates that are near expiry
    enhance         Add security enhancements to your existing configuration
   -d DOMAINS       Comma-separated list of domains to obtain a certificate for
5.3 Follow the wizard
Select the first option #1 (Spin up a temporary web server)
Enter all the information requested for the cert request.
This will save the certificate and the privet key generated to the following directory:
/etc/letsencrypt/live/DNSName/
All you need to worry about are these files:
- 
cert.pem
- 
fullchain.pem
- 
privkey.pem
6 Import the certificate to the UniFi controller
You can do this manually using the keytool-import
https://crosstalksolutions.com/secure-unifi-controller/
https://docs.oracle.com/javase/tutorial/security/toolsign/rstep2.html
But for this we are going to use the handy SSL import script made by Steven Jenkins
6.1 Download Steve Jenkins UniFi SSL Import Script
Copy the unifi_ssl_import.sh script to your server
wget https://raw.githubusercontent.com/stevejenkins/unifi-linux-utils/master/unifi_ssl_import.sh
6.2 Modify Script
Install Nano if you don’t have it (it’s better than VI in my opinion. Some disagree, but hey, I’m entitled to my opinion)
sudo apt-get install nano
nano unifi_ssl_import.sh
Change your hostname.example.com to the actual hostname you wish to use. In my case, I’m using
UNIFI_HOSTNAME=your_DNS_Record
Since we are using Ubuntu comment following three lines for Fedora/RedHat/CentOS
#UNIFI_DIR=/opt/UniFi
#JAVA_DIR=${UNIFI_DIR}
#KEYSTORE=${UNIFI_DIR}/data/keystore 
Uncomment following three lines for Debian/Ubuntu
UNIFI_DIR=/var/lib/unifi 
JAVA_DIR=/usr/lib/unifi
KEYSTORE=${UNIFI_DIR}/keystore
Since we are using Letsencrypt
LE_MODE=yes
here’s what i used for this demo
#!/usr/bin/env bash
# unifi_ssl_import.sh
# UniFi Controller SSL Certificate Import Script for Unix/Linux Systems
# by Steve Jenkins <http://www.stevejenkins.com/>
# Part of https://github.com/stevejenkins/ubnt-linux-utils/
# Incorporates ideas from https://source.sosdg.org/brielle/lets-encrypt-scripts
# Version 2.8
# Last Updated Jan 13, 2017
# CONFIGURATION OPTIONS
UNIFI_HOSTNAME=unifyctrl01.multicastbits.com
UNIFI_SERVICE=unifi
# Uncomment following three lines for Fedora/RedHat/CentOS
#UNIFI_DIR=/opt/UniFi
#JAVA_DIR=${UNIFI_DIR}
#KEYSTORE=${UNIFI_DIR}/data/keystore
# Uncomment following three lines for Debian/Ubuntu
UNIFI_DIR=/var/lib/unifi
JAVA_DIR=/usr/lib/unifi
KEYSTORE=${UNIFI_DIR}/keystore
# Uncomment following three lines for CloudKey
#UNIFI_DIR=/var/lib/unifi
#JAVA_DIR=/usr/lib/unifi
#KEYSTORE=${JAVA_DIR}/data/keystore
# FOR LET'S ENCRYPT SSL CERTIFICATES ONLY
# Generate your Let's Encrtypt key & cert with certbot before running this script
LE_MODE=yes
LE_LIVE_DIR=/etc/letsencrypt/live
# THE FOLLOWING OPTIONS NOT REQUIRED IF LE_MODE IS ENABLED
PRIV_KEY=/etc/ssl/private/hostname.example.com.key
SIGNED_CRT=/etc/ssl/certs/hostname.example.com.crt
CHAIN_FILE=/etc/ssl/certs/startssl-chain.crt
#rest of the script Omitted
6.3 Make script executable: 
chmod a+x unifi_ssl_import.sh
6.4 Run script: 
sudo ./unifi_ssl_import.sh
