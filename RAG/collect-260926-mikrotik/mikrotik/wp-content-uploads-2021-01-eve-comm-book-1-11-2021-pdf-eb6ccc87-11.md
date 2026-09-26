---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-11
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Apple", "Broadcom", "Microsoft"]
dates: []
keywords: ["ethernet", "license"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [1665, 1881]
sha256: 65ae3610742eebd15c56b6fa7f78ce2d2653a7d54ff9f0a88e7054dce58d1e89
---

# Copyright (c) 2016, Andrea Dainese

To free up space on the /boot, enter the following command, hit enter and confirm with “y” 
apt autoremove 
3.6.1.2 Verify current EVE Community version 
You have to make sure that your EVE Community Edition is of version (v2.0.3-86) or later. You 
must be able to reach the internet from your PC, VMware or Server.  
 
To check your current EVE-NG version, enter the following command

EVE-NG Community Cookbook 
Version 1.11 
Page 59 of 165 © EVE-NG LTD 
dpkg -l eve-ng 
 
You can also verify your current EVE version from the WEB GUI.  Top menu bar,  System, 
System status. 
 
You can check the version number of the newest currently available Community version on the 
EVE-NG Community site: http://www.eve-ng.net/community. 
3.6.1.3 Steps to upgrade to the latest EVE Community version 
Type the following commands below and hit enter after each. 
 apt update 
In case of any Y/N prompt, answer Yes. 
apt upgrade 
In case of any Y/N prompt, answer Yes. 
reboot 
3.6.2 Upgrading EVE Community to EVE-NG Professional 
 WARNING: Please be ready to purchase a license when upgrading, as you will not be 
able to start any nodes until a valid license has been activated on your EVE. 
To upgrade to EVE -NG Pro, issue the following commands into the CLI of EVE follow ed by 
enter. 
apt update

EVE-NG Community Cookbook 
Version 1.11 
Page 60 of 165 © EVE-NG LTD 
apt install eve-ng-pro 
 
reboot 
 
After the reboot continue with the below commands, followed by enter 
apt update 
 
apt install eve-ng-dockers 
 
reboot 
 
Continue to the EVE-NG Pro license purchase section of the website and follow the remaining 
instructions. 
3.7 Native telnet console management setup 
If you prefer to use a natively installed telnet client to manage nodes inside EVE, follow the 
steps below: 
3.7.1 Windows Native Console 
Step 1: Download the EVE Windows Client 
integration pack: 
http://www.eve-ng.net/downloads/windows-
client-side-pack 
 
Step 2: Install it as administrator 
 
Step 3: Leave the option for UltraVNC 
checked. UltraVNC is very tiny and the 
preferred VNC client for Windows by EVE. 
 
Step 4: Continue with Next. When it ask s to 
choose Ultra VNC Options, only leave the 
UltraVNC Viewer  checked, the rest is not 
needed.

EVE-NG Community Cookbook 
Version 1.11 
Page 61 of 165 © EVE-NG LTD 
Step 5: Continue with Next and finish the 
installation. 
 
By default, EVE Windows Client Integration will install Putty as your Telnet Client. The default 
location for the  EVE Win dows Client Integration  software and .re g files is: “C:\Program 
Files\EVE-NG” 
Set the default telnet program manually in Windows 10. Example: SecureCRT 
Step 1: Go to: Windows Settings/Apps/Default Apps/Choose Default Apps by Protocol  
Step 2: Set your default Telnet program: 
 
 NOTE: The first time click on the type of link that is used to access a running node 
inside EVE via telnet, the  browser will ask to choose the telnet program. If you ha ve 
prepared your default telnet program with the instructions above, you have to choose 
your default Telnet program. 
Example: Firefox browser: 
 
Set your default application, check the box “Remember my choice telnet links” and click Open 
link 
3.7.2 Linux Native Console 
The steps below will show how to setup the native consoles pack for Linux Mint 18 (Ubuntu): 
Step 1: Go to the EVE Linux Side 
integration pack download page: 
http://www.eve-ng.net/downloads/linux-
client-side 
Step 2: Open the link to GitHub 
https://github.com/SmartFinn/eve-ng-
integration

EVE-NG Community Cookbook 
Version 1.11 
Page 62 of 165 © EVE-NG LTD 
Step 3: Scroll down to the installation part 
 
 
 
Step 4: Login as root to your Linux system and enter the commands below: 
NOTE: An internet connection is required. Enter each command line below one after the 
other 
sudo add-apt-repository ppa:smartfinn/eve-ng-integration 
 
sudo apt-get update 
 
sudo apt-get install eve-ng-integration 
 For other Linux native console setup options please refer to: 
https://github.com/SmartFinn/eve-ng-integration 
 
3.7.3 MAC OSX Native Console 
Telnet Protocol: 
OSX Sierra (and older releases) is ready to use for the telnet protocol.  
 
 
 
For High Sierra, a telnet binary must be added (Apple decided to remove it and it is not present 
anymore on the latest OSX releases).

EVE-NG Community Cookbook 
Version 1.11 
Page 63 of 165 © EVE-NG LTD 
 
Procedure to install a previous telnet binary: 
Download telnet and ftp binaries from eve:  
http://your_eve_ip/files/osx.zip  (to be updated) Please contact to EVE Live chat for this 
package. 
Step 1: Reboot the Mac and hold down the “Command” and “R” key simultaneously after you 
hear the start-up chime, this will boot OSX into Recovery Mode 
Step 2: When the “OSX Utilities” screen appears, pull down the ‘Utilities’ menu at the top of the 
screen instead, and choose “Terminal” 
Step 3: Type the following command into the terminal then hit enter: 
 
crutil disable; reboot 
 
Step 4: When the OSX reboot is done, extract the osx.zip to your home directory 
Step 5:  Copy the files to /usr/bin and set the permissions using the terminal utility:  
 
sudo –i 
 
cp telnet ftp /usr/bin ; chmod 555 /usr/bin/telnet; chmod 555 /usr/bin/ftp 
 
chown root:wheel /usr/bin/telnet /usr/bin/ftp

EVE-NG Community Cookbook 
Version 1.11 
Page 64 of 165 © EVE-NG LTD 
1. Reboot the Mac and hold down Command + R keys simultaneously after you hear 
the startup chime, this will boot OSX into Recovery Mode 
2. When the “OSX Utilities” screen appears, pull down the ‘Utilities’ menu at the top of 
the screen instead, and choose “Terminal” 
Type the following command into the terminal then hit enter: 
crutil enable; reboot 
 
VNC Protocol: 
Download Chicken of VNC at: https://sourceforge.net/projects/chicken/files/Chicken-
2.2b2.dmg/download  
Install and use it as default VNC Client 
RDP Protocol: 
Download and install the Microsoft Remote Desktop on the App Store: 
 
 
3.8 Login to the EVE WEB GUI 
Login to the EVE management UI: 
http://<your_eve_ip>/ 
Default user access:

EVE-NG Community Cookbook 
Version 1.11 
Page 65 of 165 © EVE-NG LTD 
User: admin 
Password: eve 
 NOTE: You can change your EVE WEB Admin passw ord, please refer to section  
6.3.1.2 
 IMPORTANT NOTE: You must prepare and upload at least a couple of images to start 
building your labs. Refer to section 12

EVE-NG Community Cookbook 
Version 1.11 
Page 66 of 165 © EVE-NG LTD 
4 EVE-NG Community Update & Upgrade 
 Prerequisites: Internet access and working DNS on your EVE-NG is required. 
Verify your internet reachability with named ping. Example: ping www.google.com  
ping www.google.com 
 
If your ping is success, follow next step for update. If named ping has no success, please verify 
your DNS IP assigned for EVE or firewall. Some cases ping can be blocked by FW, but Internet 
and DNS are capable to make update/upgrade. 
OPTION for bare EVE installations which has bnx2x Broadcom Ethernet  drivers, please 
rewrite your driver to the newest linux-firmware: 
sudo apt-get -o Dpkg::Options::="--force-overwrite" install linux-firmware 
 IMPORTANT NOTE: before you start your EVE Community update & upgrade, please free 
up your EVE Community from older kernel packages: 
apt autoremove 
 
4.1 EVE-NG Community Update 
It is strongly recommended to keep your EVE-NG up to date. To update and upgrade, SSH to 
your EVE CLI.  
To verify your current EVE -NG version, please follow “ CLI diagnostic information display 
commands” in section 11.1.1. You can verify your current EVE version from the System/System 
Status tab on the top menu of the WEB GUI as well.

EVE-NG Community Cookbook 
Version 1.11 
Page 67 of 165 © EVE-NG LTD 
The newest version of EVE-NG can be verified by checking the official website: http://www.eve-
ng.net/community/community-2. The main page will display the latest EVE -NG version and 
correct steps to update. 
 
 
