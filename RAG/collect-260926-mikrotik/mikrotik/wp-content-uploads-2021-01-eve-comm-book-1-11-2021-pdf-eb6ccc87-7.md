---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-7
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Google", "Intel", "Microsoft", "United States"]
dates: []
keywords: ["distribution", "intel", "memory"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [508, 777]
sha256: 406d5712babf3d955e77b72aee72f5fcba594f4f645c58ec47f71dc435d13daf
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 12 of 165 © EVE-NG LTD 
• Ubuntu Server 16.04 LTS as platform for bare metal 
• Google Cloud Platform 
2.3 Unsupported hardware and systems 
The following are currently not supported: 
• VirtualBox virtualization 
• Citrix XenServer 
• Microsoft HyperV 
• Ubuntu 17.X or 18.x as platform

EVE-NG Community Cookbook 
Version 1.11 
Page 13 of 165 © EVE-NG LTD 
3 Installation 
3.1 VMware Workstation or VM Player 
3.1.1 VMware workstation EVE  VM installation using ISO image 
(preferred) 
Download EVE-NG Community ISO distribution image: 
https://www.eve-ng.net/downloads/eve-ng-2 
3.1.1.1 EVE VM Setup and Settings 
Step 1: Create a New Virtual machine 
 
Step 2: Select “I will install the operating 
system later”

EVE-NG Community Cookbook 
Version 1.11 
Page 14 of 165 © EVE-NG LTD 
Step 3: Select a Guest Operating system: 
Linux and select the version: Ubuntu 64-bit 
 
 
Step 4: Enter the name for your EVE-
COMM VM and select Location where your 
EVE VM will be stored on the host PC. 
 
 
Step 5: Type your desirable HDD size and 
select “Store virtual disk as single file”. 
 
Step 6: Press Customize Hardware

EVE-NG Community Cookbook 
Version 1.11 
Page 15 of 165 © EVE-NG LTD 
Step 7: Assign desirable memory 
 
Step 8: Set Processors “Number of 
processors” and “Number of cores per 
processor”. Set Intel VT-x/EPT Virtualization 
engine to ON (checked).  
NOTE: VMware Player will display only one 
CPU option: Number of processors. 
 
 
Step 9a: Select your desirable Network 
Adapter. Laptop PC 
NOTE: It is recommended to choose the 
NAT adapter option for Laptops to avoid 
EVE management interface IP changes. 
This can happen anytime the laptop is 
connected to a different SSID 
  
Step 9b: Select your desirable Network 
Adapter. Desktop PC 
NOTE: Desktop PC EVE management 
interface can be either NAT or Bridged to 
home LAN subnet.

EVE-NG Community Cookbook 
Version 1.11 
Page 16 of 165 © EVE-NG LTD 
 
Step 10: Select CD/DVD Option: “use ISO 
image file.” Browse to your downloaded 
EVE-Community.iso (actual name can be 
different) file 
 
Step 11: Confirm VM Settings. 
 
3.1.1.2 EVE-NG VM Installation steps 
 Mandatory Prerequisites: Internet must be reachable from your PC and VMware. EVE ISO 
installation requires internet access to get updates and install the latest EVE-PRO version from 
the EVE-NG repository. DNS must work as well, to check it, do a named ping, for example ping 
www.google.com 
EVE VM Installation from ISO has 3 Phases 
Phase 1 (Ubuntu installation) 
Step 1: Power ON EVE VM. Chose English 
and confirm with Enter. 
Step 2: Be sure that “Install EVE VM” is 
highlighted. Confirm with Enter.

EVE-NG Community Cookbook 
Version 1.11 
Page 17 of 165 © EVE-NG LTD 
  
 
Step 3: Make sure that English is selected 
and confirm with Enter. 
 
Step 4: You can select your own Location, 
or later, after management IP assignment, 
location will be set automatically. You can 
leave United States. Confirm with Enter. 
 
 
Step 5: DHCP ENABLED, EVEs hostname 
by default is eve-ng. You can change it if 
you wish. Using the Tab key select continue 
and confirm with Enter. Continue to Step 14 
Step 6: DHCP DISABLED/Static IP setup. If 
you have not enabled DHCP in the network, 
you must assign an IP address manually. 
Confirm Continue with Enter.

EVE-NG Community Cookbook 
Version 1.11 
Page 18 of 165 © EVE-NG LTD 
  
 
Step 7: Confirm selection “Configure network 
manually” with Enter 
 
Step 8: Enter your desirable EVE 
management IP, using the Tab key select 
“Continue” and confirm with Enter 
 
 
Step 9: Enter your subnet mask, using the 
Tab key select “Continue” and confirm with 
Enter 
Step 10: Enter your Gateway IP, using the 
Tab key select “Continue” and confirm with 
Enter

EVE-NG Community Cookbook 
Version 1.11 
Page 19 of 165 © EVE-NG LTD 
  
 
Step 11: IMPORTANT. The name server 
must be able to resolve public DNS entries 
and will be used during the next install 
steps. Enter your name server IP, using the 
Tab key select “Continue” and confirm with 
Enter 
 
Step 12: EVEs hostname by default is eve-
ng. It can be changed if you wish, using the 
Tab key select continue and confirm with 
Enter

EVE-NG Community Cookbook 
Version 1.11 
Page 20 of 165 © EVE-NG LTD 
Step 13: Enter your networks domain name. 
You are free to use anything you like, for 
example: eve-ng.net 
Using the Tab key select continue and 
confirm with Enter 
 
Step 14: If your DNS IP settings are correct, 
Ubuntu will detect your location automatically 
by conn ecting to Ubuntu servers . Confirm 
with Enter.  
 
 
Step 15: If you have a proxy in use for your 
internet access, enter your netwo rk proxy 
settings. If no proxy is used, select Continue 
with the Tab key and confirm with Enter. 
 
Step 16:  Select no automatic  updates and 
confirm with Enter. Security updates can 
later be run manually from EVE cli. 
 
EVE VM Installation Phase 2 (EVE installation) 
Step 17:  After the “Finish the installation”  
screen appeared, DO NOT remove CD ISO 
from the VM or hit Enter continue. First, we 
have to verify that EVE is ready for the 
installation phase 2. 
Step 18: Without powering off your EVE VM, 
open the  EVE VM settings and make sure  
that CD/DVD ISO “Device status connected” 
and “Connect at power on ” is checked. 
Confirm with OK.

EVE-NG Community Cookbook 
Version 1.11 
Page 21 of 165 © EVE-NG LTD 
  
 
Step 19: Return to the EVE console screen 
and continue with Enter , the EVE VM will 
reboot and finish the installation phase 2 
 
Step 20: Once the EVE login screen appears, 
login to the  CLI with root/eve and continue 
with installation phase 3 
 
EVE VM Installation Phase 3 (Management IP setup and updates) 
Step 21: Setup EVEs Management IP 
address. A Static IP address setup is 
preferred.  
Follow steps in section: 
3.5.1 for static IP, 3.5.2 for DHCP IP 
Step 22:  After your EVE is rebooted,  
Login to EVE CLI and type:  
apt update 
apt upgrade 
 
 
 IMPORTANT NOTE: You must prepare and upload at least a couple of images to start 
building your labs. Refer to section 12

EVE-NG Community Cookbook 
Version 1.11 
Page 22 of 165 © EVE-NG LTD 
3.1.2 VMware workstation OVF deployment 
Download EVE-NG Community OVF image zip file, place it in the dedicated HDD storage for 
EVE VM and unzip it: 
https://www.eve-ng.net/index.php/download/#DL-COMM 
3.1.2.1 Deployment and VM machine settings 
Step 1: VMware workstation or VM Player, 
Menu File/Open 
 
Step 2:  Browse your downloaded and 
unzipped EVE -COMM, EVE -COMM-VM.ovf, 
followed by Open 
 
 
Step 3: Browse your desired EVE VM store 
destination followed by Import 
 
Step 4: Open your EVE VM Settings and set 
the desired RAM.

EVE-NG Community Cookbook 
Version 1.11 
Page 23 of 165 © EVE-NG LTD 
Step 5: IMPORTANT Set CPU Number of 
Cores and number of cores per processor. 
Set Intel VT-x/EPT Virtualization engine to 
ON (checked).  
NOTE: VMware Player will display only one 
CPU option: Number of processors.  
 
Step 6: Laptop PC Select your desirable 
Network Adapter. 
NOTE: It is recommended to choose the 
NAT adapter option for Laptops to avoid 
EVE management interface IP changes. 
This can happen anytime the laptop is 
connected to a different SSID. 
 
 
Step 7: Desktop PC Select your desirable 
Network Adapter.  
NOTE: Desktop PC EVE management 
interface can be either NAT or Bridged to 
home LAN subnet. 
 
Step 8: Power ON your EVE VM and follow 
Management IP setup instructions described 
in section 3.5.1 for Static IP or 3.5.2 for 
DHCP IP.

