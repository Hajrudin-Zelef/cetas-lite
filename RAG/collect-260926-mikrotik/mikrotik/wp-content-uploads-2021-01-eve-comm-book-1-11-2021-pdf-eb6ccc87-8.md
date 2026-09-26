---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-8
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Intel", "United States"]
dates: []
keywords: ["intel"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [778, 1043]
sha256: eb35189ae62ca46454f047ac597bf30d6a13a7fef2630785b7e5a4aa681a8978
---

# Copyright (c) 2016, Andrea Dainese

EVE-NG Community Cookbook 
Version 1.11 
Page 24 of 165 © EVE-NG LTD 
 IMPORTANT NOTE: You must prepare and upload at least a couple of images to start 
building your labs. Refer to section 12 
3.1.2.2 OVF VM update to the latest EVE version 
Step 9: Make sure if your EVE OVF VM is up to date to the newest EVE version. 
Follow the steps described in section 4. 
 
3.1.2.3 OVF VM HDD Size expansion 
 IMPORTANT NOTE:  DO NOT expand the current EVE OV F HDD. To expand your EVE 
system size, please follow Troubleshooting section 11.2 
3.2 VMware ESXi 
3.2.1 VMware ESXi EVE installation using ISO image (preferred) 
Download EVE-NG Community ISO installation image: 
http://www.eve-ng.net/downloads/eve-ng-2 
3.2.1.1 EVE-NG ESXi VM Setup and Settings 
Step 1: Upload EVE ISO image to the ESXi 
store. 
 
Step 2: Create NEW VM

EVE-NG Community Cookbook 
Version 1.11 
Page 25 of 165 © EVE-NG LTD 
Step 3: Enter the name for your EVE-PRO 
VM and select Guest Operating system 
Linux and version: Ubuntu 64-bit 
 
Step 4: Select Location where your EVE VM 
will be stored in HDD. 
 
 
Step 5: IMPORTANT Customize your EVE 
VM CPU Settings. Set CPU Number of 
Cores and number of cores per processor. 
Set Intel VT-x/EPT Virtualization to ON 
(checked). 
 
Step 6: Assig desirable RAM for your EVE 
 
 
Step 7: Set the size of HDD for your new 
EVE VM. It is recommended to set “Thick 
Provisioned eagerly provisioned”. Server 
EVE HDD is recommended to set at least 
500Gb 
Step 8: Set your Management network. 
Adapter type VMXNET3

EVE-NG Community Cookbook 
Version 1.11 
Page 26 of 165 © EVE-NG LTD 
 
NOTE: Additional Network A dapters c an be 
added for further use. 
 
Step 9:  Add new device to your EVE VM, 
CD/DVD 
 
Step 10:  Set DVD drive to “Datastore ISO 
File” and browse your uploaded EVE -
PRO.iso. Make sure that Status is checked 
ON, “Connect at power on” 
 
3.2.1.2 EVE-NG ESXi VM Installation steps 
 Mandatory Prerequisites: Internet must be reachable from your PC and VMware. EVE 
ISO installation requires internet access to get updates and install the latest EVE-PRO 
version from the EVE-NG repository. DNS must work as well, to check it, do  a named 
ping, for example ping www.google.com  
EVE ESXi VM Installation from ISO has 3 Phases 
Phase 1 (Ubuntu installation) 
Step 1: Power ON EVE VM. Chose English 
and confirm with Enter. 
Step 2: Be sure if “Install EVE VM” is 
highlighted. Confirm with Enter.

EVE-NG Community Cookbook 
Version 1.11 
Page 27 of 165 © EVE-NG LTD 
  
 
Step 3: Make sure if English is selected and 
confirm with Enter. 
 
Step 4: You can select your own Location, 
or later, after management IP assignment, 
location will be set automatically. You can 
leave United States. Confirm with Enter 
 
 
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
Page 28 of 165 © EVE-NG LTD 
  
 
Step 7: Confirm selection “Configure network 
manually” with Enter 
 
Step 8: Enter your desirable EVE 
management IP, using the Tab key select 
“Continue” and confirm with Enter 
 
 
Step 9: Correct your subnet mask, using the 
Tab key select “Continue” and confirm with 
Step 10: Correct your Gateway IP, using 
the Tab key select “Continue” and confirm

EVE-NG Community Cookbook 
Version 1.11 
Page 29 of 165 © EVE-NG LTD 
Enter 
 
with Enter
 
 
Step 11: IMPORTANT. Name server must 
respond to the Internet and will be used 
during the next install steps. Enter your 
name server IP.  Using the Tab key select 
“Continue” and confirm with Enter 
 
Step 12: EVE hostname by default is eve-
ng. It can be changed if you wish. Using the 
Tab key select continue and confirm with 
Enter 
 
 
Step 13: Enter your network domain name. 
You are free to use any, for example:  
eve-ng.net 
Step 14: If your DNS IP settings are correct, 
Ubuntu will detect your location from 
Internet. Confirm with Enter.

EVE-NG Community Cookbook 
Version 1.11 
Page 30 of 165 © EVE-NG LTD 
Using the Tab key select continue and 
confirm with Enter 
 
 
 
Step 15:  If you have proxy in use for your 
internet, assign your network proxy settings. 
If no proxy in use, with Tab key select 
Continue and confirm with Enter. 
 
Step 16:  Select no automatic updates and 
confirm with Enter. Security updates ca n be 
run later manually from EVE cli. 
 
EVE VM Installation Phase 2 (EVE installation) 
Step 17:  After the “Finish the installation” 
screen appears, DO NOT remove CD ISO 
from VM or hit Enter continue. We have to 
verify settings for EVE installation Phase 2 . 
Follow step 9. 
Step 18: Without powering off the EVE VM, 
open the  EVE VM settings and make sure 
that CD/DVD ISO “Device status connected” 
and “Connect at power on” is checked. 
Confirm with OK.

EVE-NG Community Cookbook 
Version 1.11 
Page 31 of 165 © EVE-NG LTD 
 
 
 
Step 19: Return back to EVE console screen 
and confirm Continue with Enter , EVE VM 
will reboot and continue Phase 2 installation 
 
Step 20:  Once EVE login screen appeared, 
login in CLI with root/eve and follow 
installation Phase 3 
 
EVE VM Installation Phase 3 (Management IP setup and updates) 
Step 21:  Setup EVE Management IP 
address. A Static IP address setup is 
preferred  
Follow steps in section : 
3.5.1 for static IP, 3.5.2 for DHCP IP 
Step 22:  After your EVE is rebooted,  
Login to EVE CLI and type:  
apt update 
apt upgrade

EVE-NG Community Cookbook 
Version 1.11 
Page 32 of 165 © EVE-NG LTD 
Step 2 3:  On the EVE CLI prompt, reboot 
EVE by typing  
reboot 
 
 IMPORTANT NOTE: You must prepare and upload at least a couple of images to start 
building your labs. Refer to section 12 
3.2.2 VMware ESXi OVF deployment 
Download EVE-NG Community OVF image zip file, place it in the dedicated HDD storage for 
EVE VM and unzip it: 
https://www.eve-ng.net/index.php/download/#DL-COMM 
3.2.2.1 ESXi OVF VM Setup and Settings 
Step 1: ESXi Host, Create/Register VM 
 
Step 2: Set option Deploy a virtual machine 
from an OVF or OVA file 
 
 
Step 3: Type the name for your new EVE 
VM and browse to select your all 
downloaded EVE OVF files 
Step 4: Select the storage where your EVE 
VM will be deployed.

EVE-NG Community Cookbook 
Version 1.11 
Page 33 of 165 © EVE-NG LTD 
  
 
Step 5: Select your Management network 
and Thick Disk provisioning. EVE OVF 
HDD is only 40Gb large. It is recommended 
after installation to add extra HDD. Section 
11.2 
 
Step 6:  IMPORTANT Open VM Settings. 
Set the quantity of CPUs and number of 
cores per socket. Set Intel VT-x/EPT 
Hardware Virtualization engine to ON 
(checked).  
 
 
Step 7: Set desirable RAM for your EVE. 
 
Step 8: Power ON your EVE VM and follow 
Management IP setup instructions described 
in section 3.5.1 for Static IP or 3.5.2 for 
DHCP IP.

EVE-NG Community Cookbook 
Version 1.11 
Page 34 of 165 © EVE-NG LTD 
 IMPORTANT NOTE: You must prepare and upload at least a couple of images to start 
building your labs. Refer to section 12 
3.2.2.2 ESXi OVF VM update to the latest EVE version 
Make sure that your EVE OVF VM is up to date with the newest EVE version. 
Follow the steps described in section 4 for upgrade instructions 
 
3.2.2.3 ESXi OVF VM HDD Size expansion 
 NOTE: IMPORTANT! DO NOT expand the current EVE OVF HDD. To expand your EVEs 
system disk size, please follow the troubleshooting section 11.2 
3.3 Bare hardware server EVE installation 
Download Ubuntu Server 16.04.6 LTS ISO image: 
http://tw.archive.ubuntu.com/ubuntu-cd/16.04/ubuntu-16.04.6-server-amd64.iso 
 
