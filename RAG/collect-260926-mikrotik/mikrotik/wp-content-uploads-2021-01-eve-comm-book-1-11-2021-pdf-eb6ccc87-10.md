---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-10
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Google", "Intel"]
dates: []
keywords: ["compute", "intel", "licenses"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [1384, 1664]
sha256: 38602976a4155cb76f2b7b07993ab3f925a3f5584f1250bc5a4da7acac5f6a03
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 46 of 165 © EVE-NG LTD 
Step 2: create a nested Ubuntu 16.04 image model. Copy and paste the below command into 
the shell. Use copy/paste. crtl +c/ctrl +v. It is single line command. Confirm with “enter”: 
 
gcloud compute images create nested-ubuntu-xenial --source-image-
project=ubuntu-os-cloud --source-image-family=ubuntu-1604-lts --
licenses="https://www.google.com/compute/v1/projects/vm-
options/global/licenses/enable-vmx" 
 
 
You will get the following output when your image is ready: 
 
 
3.4.4 Creating VM 
Step 1: Navigate: Navigation Menu/Compute Engine/VM Instances and press “Create” 
 
 
 
 
Step 2: Assign the name for your VM 
 
Step 3: Set your own region and zone 
 
Step 4: Edit your Machine Configuration. General-Purpose. Choose the series of CPU platform, 
Preferred are Intel CPUs Skylake or Cascade. 
 
Step 5: Choose your desirable CPU and RAM settings. 
IMPORTANT: “Deploy a container image” must be UNCHECKED.

EVE-NG Community Cookbook 
Version 1.11 
Page 47 of 165 © EVE-NG LTD 
 
 
 
Step 6: Select Boot disk. Press Change 
 
  
 
Step 7. Select Custom images and the custom boot images you created previously. Choose 
HDD disk type and size. HDD size can vary depends of your needs.

EVE-NG Community Cookbook 
Version 1.11 
Page 48 of 165 © EVE-NG LTD 
 
 
 
Step 7: Allow http traffic and create VM 
 
3.4.5 EVE-NG-Community installation 
Step 1: Connect to the VM with the first option “Open in browser window”

EVE-NG Community Cookbook 
Version 1.11 
Page 49 of 165 © EVE-NG LTD 
 
 
Step 2: Launch installation with:  
 
Type the below command to become root: 
sudo -i 
 
Start EVE-COMM installation 
wget -O - http://www.eve-ng.net/repo/install-eve.sh | bash -i 
 
 
Step 3:  Update and upgrade your new EVE-COMM 
apt update 
 
apt upgrade 
Confirm with Y 
 
Step 4. Reboot EVE. Allow some time for reboot and then press “Reconnect”  
 
 
Step 5: IMPORTANT: Setup IP 
Once the IP wizard screen appears, press ctrl +c and type the below command to become root: 
sudo -i 
 
 
 
Now follow the IP setup wizard.

EVE-NG Community Cookbook 
Version 1.11 
Page 50 of 165 © EVE-NG LTD 
IMPORTANT: set IP as DHCP! 
 
Step 6: Dockers installation. After EVE is rebooted, reconnect the SSH session:  
 
Type command to become root: 
sudo -i 
 
Type command to update EVE 
apt update 
 
 
3.4.6 Access to Google Cloud EVE-COMM 
Use your public IP for accessing EVE via http. 
 
 
 
 
Default web login: admin/eve 
3.4.7 Optional: GCP Firewall rules for native console use 
Step 1: Navigate: Navigation menu/VPC Network/Firewall rules

EVE-NG Community Cookbook 
Version 1.11 
Page 51 of 165 © EVE-NG LTD 
 
 
Step 2: Create new firewall rule 
 
 
 
 
 
Step 3: Create an ingress FW rule; allow TCP ports 0-65535

EVE-NG Community Cookbook 
Version 1.11 
Page 52 of 165 © EVE-NG LTD 
 
Step 4: Create an egress FW rule; allow TCP ports 0-65535

EVE-NG Community Cookbook 
Version 1.11 
Page 53 of 165 © EVE-NG LTD 
 
Summary FW rules. 
 
 
3.5 EVE Management IP Address setup 
3.5.1 Management static IP address setup (preferred) 
The steps below will walk you through the network setup and a ssign a static management IP 
for EVE. 
Step 1: Log into the EVE CLI using the default 
login root/eve After login, type your preferred 
root password  for EVE , default is eve. 
Remember it for further use. Confirm with 
enter 
NOTE: Typed characters in the passwo rd 
field are not visible. 
Step 2:  Retype your root password again 
and confirm with enter.

EVE-NG Community Cookbook 
Version 1.11 
Page 54 of 165 © EVE-NG LTD 
 
 
 
 
Step 3: Choose your EVE VMs hostname. By 
default, it is eve-ng. You can leave it as it is. 
Confirm with enter 
 
Step 4: Type your domain name for your 
EVE VM. By default, it is example.com. The 
default value can be used as well. 
Confirm with enter
 
 
Step 5: Using the arrow keys, select the 
option “static”, confirm your selection with 
the space key, followed by enter 
 
Step 6: Type your desirable EVE 
management IP. Confirm with enter.

EVE-NG Community Cookbook 
Version 1.11 
Page 55 of 165 © EVE-NG LTD 
Step 7: Type the subnet mask of your EVE 
management network. Confirm with enter. 
 
Step 8: Type your networks gateway IP. 
Confirm with enter. 
 
 
Step 9: Type your networks primary DNS 
IP. Confirm with enter. 
IMPORTANT: DNS must be reachable and 
resolve public addresses. 
.  
Step 10: Type your network Secondary 
DNS IP. Confirm with Enter. 
IMPORTANT: DNS must be reachable and 
resolve public addresses. 
 
 
Step 11: Type your preferred NTP server IP. 
It can be left empty as well; in this case, your 
EVE VM will auto matically assign the time 
from its host. 
Step 12: If you have a proxy in use for your 
Internet, select the respective  proxy option 
and configure your proxy settings. By default, 
it is direct connection  (no proxy) . Confirm 
your selection with enter. EVE will reboot 
automatically.

EVE-NG Community Cookbook 
Version 1.11 
Page 56 of 165 © EVE-NG LTD 
  
 IMPORTANT NOTE: If you are setting up your management IP for the first time (fresh 
EVE installation), please return to the install section and complete installation phase 3. 
3.5.2 EVE Management IP address setup via DHCP 
The steps below will walk you through the network setup and assign a management IP for EVE 
via DHCP. 
Step 1: Log into the EVE CLI using the default 
login root/eve After login, type your preferred 
root password for EVE, default is eve. 
Remember it for further use. Confirm with 
enter 
NOTE: Typed characters in the password 
field are not visible. 
 
Step 2: Retype your root password again 
and confirm with enter. 
 
 
Step 3: Choose your EVE VMs hostname. By 
default, it is eve-ng. You can leave it as it is. 
Confirm with enter 
Step 4: Type your domain name for your 
EVE VM. By default, it is example.com. The 
default value can be used as well. 
Confirm with enter

EVE-NG Community Cookbook 
Version 1.11 
Page 57 of 165 © EVE-NG LTD 
  
 
Step 5: Using the arrow keys, select the 
option “dhcp”, confirm your selection with the 
space key, followed by enter 
 
Step 6: Type your preferred NTP server IP. 
It can be left empty as well; in this case, your 
EVE VM will automatically assign the time 
from its host. 
 
 
Step 7: If you have a proxy in use for your 
Internet, select the  respective proxy option 
and configure your proxy settings. By default, 
it is direct connection (no proxy). Confirm 
your selection with enter. EVE will reboot 
automatically.

EVE-NG Community Cookbook 
Version 1.11 
Page 58 of 165 © EVE-NG LTD 
 IMPORTANT NOTE: If you are setting up your management IP for the first time (fresh 
EVE installation), please return to the install section and complete installation phase 3. 
3.5.3 EVE Management IP address reset 
If for any reason you need to change these settings after the installation, you can rerun the IP 
setup wizard. Type the following command in the CLI and hit enter: 
rm -f /opt/ovf/.configured 
Then reboot. Once you log  into the CLI again, EVE will go t hrough the network setup again . 
Please follow the steps in section 3.5.1 for Static IP or 3.5.2 for DHCP IP. 
3.6 EVE-NG Community upgrade to EVE-NG Professional 
3.6.1 Mandatory Prerequisites  
 Mandatory Prerequisites: Internet must be reachable from your PC and VMware. EVE 
ISO installation requires internet access to get updates and install the latest EVE-PRO 
version from the EVE-NG repository. DNS must work as well, to check it, do a named 
ping, for example ping www.google.com  
3.6.1.1 EVE Community disk space 
You must have enough HDD space available. The main eve--ng--vg-root partition must have 
at least 10GByte free space while the boot partition must have at least 50Mbyte. To check how 
much space is available on your HDD, enter the following command into the CLI of EVE: 
df -h 
 
 
