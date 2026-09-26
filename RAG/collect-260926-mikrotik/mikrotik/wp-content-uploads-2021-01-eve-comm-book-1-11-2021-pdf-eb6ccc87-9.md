---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-9
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Google", "United States"]
dates: []
keywords: ["compute"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [1044, 1383]
sha256: 7da969922f622acc16d821c2ed87348bc169083924708b22bbdcbaa8d86d6077
---

# Copyright (c) 2016, Andrea Dainese

 Mandatory Prerequisites: Internet must be reachable from your PC and VMware. EVE ISO 
installation requires internet access to get updates and install the latest EVE-COMM version 
from the EVE-NG repository. DNS must work as well, to check it, do a named ping, for example 
ping www.google.com 
3.3.1 Ubuntu Server Installation Phase 1 
Step 1: Create a bootable DVD disk or USB 
flash drive with an Ubuntu server image. 
Boot your server from ISO . Make sure that 
English is selected, Confirm with Enter 
 
 
Step 2:  Select the first Option “Install Ubuntu 
Server” Confirm with Enter 
 
 
 
Step 3:  Make sure that English is selected 
and confirm with Enter 
Step 4: You can select your own Location, 
or later, after management IP assignment,

EVE-NG Community Cookbook 
Version 1.11 
Page 35 of 165 © EVE-NG LTD 
 
location will be set automatically. You can 
leave United States. Confirm with Enter. 
 
 
Step 5: Configure the keyboard, leave “No” 
selected and confirm with enter 
 
Step 6:  Leave English (US)  as selection , 
confirm with Enter 
 
 
Step 7: Leave English (US) as selection and 
confirm with Enter 
 
Step 8: Select your management network 
adapter and confirm with Enter

EVE-NG Community Cookbook 
Version 1.11 
Page 36 of 165 © EVE-NG LTD 
Step 9: DHCP ENABLED 
Continue with Step 16 
 
Step 10: DHCP DISABLED/Static IP setup. 
If have not enabled DHCP in the network, 
you must assign an IP address manually. 
Continue with Enter. 
 
 
 
Step 11: Select “Configure network 
manually” and confirm with Enter
 
Step 12: Enter your desirable EVE 
management IP, using the Tab key select 
“Continue” and confirm with Enter
 
 
 
Step 13: Enter your subnet mask, using the 
Tab key select “Continue” and confirm with 
Enter 
Step 14:  Enter your Gateway IP, using the 
Tab key select “Continue” and confirm with 
Enter

EVE-NG Community Cookbook 
Version 1.11 
Page 37 of 165 © EVE-NG LTD 
  
 
 
Step 15: IMPORTANT: The name server 
must be able to resolve public DNS entries 
and will be used during the next install 
steps. Enter your name server IP,  using the 
Tab key select “Continue” and confirm with 
Enter 
 
 
Step 16:  Type your EVE server hostname, 
Example: eve-ng 
 
 
Step 17: Type your domain name. You are 
free to use any. Example: eve-ng.net 
 
Step 18: Type your Ubuntu username, 
Example: user

EVE-NG Community Cookbook 
Version 1.11 
Page 38 of 165 © EVE-NG LTD 
 
Step 19: Select a username (e.g. “user”) for  
your account and Continue 
 
Step 20: Enter a password for your new 
user
 
 
 
 
Step 21:  Re-enter your password and 
continue 
 
Step 22: If you want to use a weak password, 
click “Yes” on this screen. 
 
 
Step 23: Encrypt your Home directory, “No” 
 
Step 24: If your DNS and internet are 
working properly, Ubuntu will automatically 
detect your location and timezone. Confirm 
your timezone and continue with enter

EVE-NG Community Cookbook 
Version 1.11 
Page 39 of 165 © EVE-NG LTD 
 
 
Step 25: Select HDD partitioning method 
“Guided – use entire disk and set up LVM” 
 
Step 26: Select your disk partition, and 
confirm with enter
 
 
Step 27: Confirm write changes to disk with 
“Yes” and hit enter to continue 
 
Step 28: Select the volume size and continue 
 
 
Step 29: Confirm write the changes to disk 
with “Yes” and continue 
Step 30: If you have a proxy in use for your 
internet, enter your network proxy settings. If 
no proxy i s used, use the tab key to select 
Continue and confirm with enter.

EVE-NG Community Cookbook 
Version 1.11 
Page 40 of 165 © EVE-NG LTD 
  
 
Step 31: Select “No automatic updates” and 
Continue 
 
Step 32:  Using the Arrow keys select 
“OpenSSH server” for installation  and 
confirm with the Space key (*), continue with 
enter 
 
 
Step 33: Confirm “Install the GRUB 
bootloader to the master boot record” with 
“Yes” and continue with enter 
 
Step 34: REMOVE CD/DVD installation 
media and continue with enter

EVE-NG Community Cookbook 
Version 1.11 
Page 41 of 165 © EVE-NG LTD 
Step 35: Login in to your Ubuntu with the 
username created above (user/Test123 was 
the example)
 
 
Step 36:  Continue as root user. Enter the 
commands below, each followed by the enter 
key. 
 
sudo su 
 
Test123 
 
cd 
 
 
 
 
 
Step 37: Create root password 
 
sudo passwd root 
 
Repeat your desirable password twice; 
Example: eve 
 
 
 
Step 38:  Verify and set your hostname if you 
haven’t set it before 
 
nano /etc/hostname 
 
Edit it if necessary: eve-ng 
 
Confirm edit with ctrl+o followed by Enter 
And ctrl+x for Exit 
 
Step 39: Verify your host settings 
 
nano /etc/hosts 
 
Your assigned static IP will be bound to your 
server hostname and domain 
 
NOTE: in case if DHCP IP address is used, 
you will see 127.0.0.1 IP vs hostname 
 
Confirm edit with ctrl+o followed by enter 
And ctrl+x for Exit 
Step 40:  Edit permissions for root user to 
allow SSH access to EVE server 
 
nano /etc/ssh/sshd_config 
 
Find and edit PermitRootLogin to “yes” 
 
Confirm edit with ctrl+o followed by enter 
And ctrl+x for Exit 
 
Restart ssh service: 
 
sudo service ssh restart 
 
Step 41:  IMPORTANT 
SSH as root to your EVE server with Putty or any other telnet client program. 
Update the Ubuntu grub CMD Line with the following customized command. Make sure you 
enter this command below in a single line and confirm it with the enter key.

EVE-NG Community Cookbook 
Version 1.11 
Page 42 of 165 © EVE-NG LTD 
sed -i -e 's/GRUB_CMDLINE_LINUX_DEFAULT =.*/GRUB_CMDLINE_LINUX_DEFAULT="net.ifnames=0 
noquiet"/' /etc/default/grub 
 
Update GRUB, Followed by Enter 
 
update-grub 
 
 WARNING: DO NOT REBOOT your Ubuntu/EVE yet, proceed to step 42! 
 
Step 42: IMPORTANT 
Rename your Server interface name to eth0 
 
nano /etc/network/interfaces 
 
Before edit: 
 
After edit: 
 
 
Confirm your edit with ctrl+o followed by enter 
And ctrl+x to exit 
 
Reboot the EVE server 
 
reboot 
 
3.3.2 EVE Community Installation Phase 2 
Step 43: Start EVE Community installation with the following one-line command and hit enter 
 
wget -O - http://www.eve-ng.net/repo/install-eve.sh | bash -i 
 
Step 44: Reboot EVE 
reboot 
3.3.3 EVE Community Installation Phase 3 
 
Step 45: After the installation is completed, 
reboot EVE and follow the Management IP 
setup instructions in section 3.5.1. It is 
strongly recommended for bare-metal 
Step 46:  After your EVE is rebooted,

EVE-NG Community Cookbook 
Version 1.11 
Page 43 of 165 © EVE-NG LTD 
installations to use a static IP address. After 
the IP address setup, continue with Step 46 
 
Login to the EVE CLI and type:  
apt update 
apt upgrade 
reboot 
 
 
 
 IMPORTANT NOTE: You must prepare and upload at least a couple of images to start 
building your labs. Refer section 12 
3.4 Google Cloud Platform 
3.4.1 Google account 
Step 1: Connect to Google Cloud Platform (GCP 
https://console.cloud.google.com/getting-started 
 
 
 
Step 2: Sign into GCP. Create a new GCP account if you do not already have one.  
3.4.2 Goggle Cloud project 
Create new project. By default, GCP will offer you a project named “My First Project”. It can be 
used as well. 
 
Step 1. GCP top bar, click on “My First Project” 
 
 
 
Step 2. Next pop up window, click “NEW PROJECT”

EVE-NG Community Cookbook 
Version 1.11 
Page 44 of 165 © EVE-NG LTD 
Step 3. Enter your project name, and confirm “CREATE” 
 
This will take some time.  
 
Step 4. Navigate: Navigation Menu/Compute Engine/VM Instances 
 
 
 
Step 5. Navigate: top bar and select your newly created Project

EVE-NG Community Cookbook 
Version 1.11 
Page 45 of 165 © EVE-NG LTD 
 
Preparation of your Project can take some time. Wait until the VM Instance window finishes 
deployment and then press the “Create button.” 
 
3.4.3 Preparing Ubuntu boot disk template 
Step 1: Open the google cloud shell and press: “START CLOUD SHELL”

