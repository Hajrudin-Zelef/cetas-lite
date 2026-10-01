---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-220066768-updating-and-installing-self-hosted-unifi-network-se-abf3ecf8
title: "hc-en-us-articles-220066768-updating-and-installing-self-hosted-unifi-network-se-abf3ecf8"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-220066768-updating-and-installing-self-hosted-unifi-network-se-abf3ecf8.md
source_anchor: ""
source_lines: [1, 59]
sha256: 766e1f01d98d1a5b4aa0757590483ec0d5c3f7046af2f0327bd8f1dff6404305
---

# hc-en-us-articles-220066768-updating-and-installing-self-hosted-unifi-network-se-abf3ecf8

(Legacy) Updating and Installing Self-Hosted UniFi Network Servers (Linux)
| Looking for the next generation of UniFi self-hosting? The UniFi OS Server is the new standard for self-hosting UniFi, offering support for advanced features like Organizations, IdP Integration, and Site Magic SD-WAN. Learn more. | 
This article provides the steps to install and update the self-hosted UniFi Network Server to the current stable release on a Debian or Ubuntu system via APT (Advanced Package Tool). If you run into issues following the process described in this article, please take a look at the scripts provided here to install and update the UniFi Network Application.
Requirements
In order to update the UniFi Network Server via APT, it is necessary to create source files or edit lines in an existing sources.list file with Linux text editors: vi or nano. The repo structure should be permanent, but if there are any changes, they will be pointed out in the UniFi Network software version release posts, found in the Release section of the Community.
Before upgrading the UniFi Network Server, make sure that you have backed up the UniFi Network Database. You will need to make sure that the user has sudo permissions. For more information about adding a user to sudo list, see this Debian article.
Ubuntu or Debian-Based Scripts
Refer to this Community post.
UniFi Network APT Steps
- 
Install required packages before you begin with the following command: sudo apt-get update && sudo apt-get install ca-certificates apt-transport-https
- 
Use the following command to add a new source list: echo 'deb [ arch=amd64,arm64 ] https://www.ui.com/downloads/unifi/debian stable ubiquiti' | sudo tee /etc/apt/sources.list.d/100-ubnt-unifi.list
- Add the GPG Keys.
  - 
Install the following trusted key into /etc/apt/trusted.gpg.d :sudo wget -O /etc/apt/trusted.gpg.d/unifi-repo.gpg https://dl.ui.com/unifi/unifi-repo.gpg
- 
- Install the correct version of MongoDB.
  - 
New UniFi Install: Run the following command: echo "deb [trusted=yes] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/3.6 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.6.list
sudo apt-get updateNote: Depending on your Linux distribution, this step may not be necessary. If MongoDB is not included with your distro and isn't available in its repository, refer to the MongoDB installation guide (Ubuntu | Debian).
  - Existing Install (UniFi Network v7.5 or Newer): Skip this step.
  - Existing Install (UniFi Network v7.4 or Prior): Refer to MongoDB's Help Article for instructions on updating MongoDB to v.3.6 or later.
- 
- 
Install and upgrade the UniFi Network application with the following command: sudo apt-get update && sudo apt-get install unifi -y
- The UniFi Network application should now be accessible at the computer's configured local or public IP address, by typing that IP address in a browser's navigation bar (Chrome is recommended). If it is not launching, use the following command:
sudo service unifi start
When using the commands above, it is assumed you have sudo and wget installed. More information about sudo can be found here, and wget here.
Other Helpful Commands
- To stop the UniFi service: sudo service unifi stop
- To restart the UniFi service: sudo service unifi restart
- To see the status of UniFi service: sudo service unifi status
"Testing" refers to the next generation release, which is not released to the general public yet. "Stable" refers to the current stable release, which is supported by Ubiquiti and described in this article. "Old Stable" is the previous stable release, which has been replaced by the new stable release.
| Suite Name | Code Name | Archived Code Names | 
| oldstable | unifi-5.10 | These code names have been archived and are no longer supported | 
| stable | unifi-5.11 |  | 
| testing | * | unifi3, unifi4, unifi-5.3, unifi-5.4, unifi-5.5, unifi-5.6, unifi-5.7, unifi-5.8, unifi-5.9 | 
*testing is empty at the moment
Log Files Location
Log files are essential for any troubleshooting. Find them here:
- /usr/lib/unifi/logs/server.log
- /usr/lib/unifi/logs/mongod.log
If your application is running on a Unix/Linux based system, then you will require superuser (sudo) privileges to access these log files.
Notes and Tips
- This Community post shares some scripts to install and update the UniFi Network Application for any recent Ubuntu or Debian version.
- If you are installing in a VM or a headless server, you may encounter entropy issues. This could be anything from slow service start/restart to complete service failure. The fix is to install haveged . This is an external link for a tutorial on the subject.
- Since UniFi Network version 5.6.x the UniFi service does not run as root. This means that you cannot bind to privileged ports (<1024). The application will fail to start if you try to use these ports.
- The following affects APT versions 1.5 onward (Ubuntu 17.10 and Debian Sid or newer). A recent version of the apt-secure man page stated: "Since version 1.5 changes in the information contained in the Release file about the repository need to be confirmed before APT continues to apply updates from this repository", meaning that when performing an update from a major version to the next (for example 5.5.x to 5.6.x) the apt-get update will result in an error.
  - To fix this run the command the following way: apt-get update --allow-releaseinfo-change
  - If you receive an error stating the command is not understood in combination with the other options, users have reported that issuing the following two commands has fixed it: issue apt clean and hit enter, followed byapt-get update and enter.
- To fix this run the command the following way: 
- The UniFi Network application, when installed on Debian and Ubuntu, will not have a GUI since it's being run as a Service. Please use the service command for starting, stopping and restarting the UniFi Network application.
- If you see the following error on your DNS server, a user reports solving this by forcing the system to use 8.8.8.8 as DNS server, not his ISP.
- At the moment Ubiquiti does not support arm64, hence it is not available via the repo. A download and manual installation will be necessary for this.
- 
In the Add the GPG Keys, method B: Source: External Link. For users behind restrictive firewalls the following command will allow them to import the GPG key: apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 06E85760C0A52C50
- Because of the Java 8 dependency, when using UniFi Network 5.7.X some additional steps might be needed. See this Community post for details.
- The command apt-get isapt in Ubuntu older than 16.04.
