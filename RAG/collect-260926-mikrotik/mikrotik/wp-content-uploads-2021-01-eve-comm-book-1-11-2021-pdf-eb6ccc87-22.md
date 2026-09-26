---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-22
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["copyright", "amd", "distribution", "ethernet", "intel", "liability"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [3972, 4211]
sha256: 1b560f9dc4472f1e3a466526f2463348cd36f47bd308fa8ea55d108492962691
---

# Copyright (c) 2016, Andrea Dainese

The image hdd inside the folder must be named correctly: 
Example: hda.qcow2 or virtioa.qcow2 
 
Full path Example: opt/unetlab/addons/qemu/acs-5.8.1.4/hda.qcow2 
The table of proper folder names is provided in our website: 
 
https://www.eve-ng.net/index.php/documentation/qemu-image-namings/ 
 
Supported HDD formats for the EVE images: 
lsi([a-z]+).qcow  lsia.qcow

EVE-NG Community Cookbook 
Version 1.11 
Page 159 of 165 © EVE-NG LTD 
hd([a-z]+).qcow  hda.qcow 
virtide([a-z]+).qcow  virtidea.qcow 
virtio([a-z]+).qcow  virtioa.qcow 
scsi([a-z]+).qcow  scsia.qcow 
sata([a-z]+).qcow  sataa.qcow 
12.2 How to prepare images for EVE 
How to add EVE-NG images please refer to: 
https://www.eve-ng.net/index.php/documentation/howtos/ 
12.3 How to add custom image template 
12.3.1 Templates folder choice 
 IMPORTANT NOTE: Starting from EVE-Community Version 2.0.3-107, EVE installation is 
autodetecting what kind of CPU manufacturer has your server: Intel or AMD, to choose proper 
templates set. You can check it manually on EVE cli: example below, showing that EVE has 
Intel CPU. 
root@eve-ng:~# lsmod | grep ^kvm_ 
kvm_intel             212992  74 
root@eve-ng:~# 
 
• If you have Intel CPU, then your template files are in 
"/opt/unetlab/html/templates/intel/" 
• If you have AMD CPU, then your template files are in 
"/opt/unetlab/html/templates/amd/" 
12.3.2 Prepare template file 
 NOTE: For templates development use templates folder which is matching your EVE server 
CPU manufacturer. 
Example below will be based for Intel CPU EVE custom image template. Use EVE cli or 
WinSCP/Filezilla to create template. 
Step 1:  Navigate to EVE location: /opt/unetlab/html/templates/intel/ 
Step 2: Choose your most suitable template from which you want to create your own image 
template. (example: newimage.yml)

EVE-NG Community Cookbook 
Version 1.11 
Page 160 of 165 © EVE-NG LTD 
 
Step 3: Make a copy from source template newimage.yml. Example: Using CLI create template 
and name it ngips.yml. 
cp /opt/unetlab/html/templates/intel/newimage.yml /opt/unetlab/html/templates/intel/ngips.yml 
You can create new template using WinSCP or Filezilla as well. 
 
IMOPRTANT: The new name of your template will be related to your image foldername. You r 
image foldername must start with prefix “ngips- “ 
Example: image foldername under /opt/unetlab/addons/qemu/ngips-6.5.0-115 
 
12.3.3 Prepare interface format and name lines 
EVE Community has included option to create various interface names, sequences and 
numbering. Please refer table below.  
Formula Template line format 
example 
Will 
produce

EVE-NG Community Cookbook 
Version 1.11 
Page 161 of 165 © EVE-NG LTD 
eth_format: <prefix>{<first value for slot: example 
1>}<separator>{<first value for port>-<number of 
port per slot: example 8>} 
eth_format: Gi{1}/{0-8} 
Gi1/0 
Gi1/1 
Gi1/2 
Gi1/3 
Gi1/4 
Gi1/5 
Gi1/6 
Gi1/7 
Gi2/0 
Gi2/1 
.... 
eth_format: <prefix>{<first value for slot:  example 
0>}<separator>{<first value for port> -<number of 
port per slot: example 4>} 
eth_format: Ge{0}/{0-4} 
Ge0/0 
Ge0/1 
Ge0/2 
Ge0/3 
Ge1/0 
Ge1/2 
Ge1/3 
Ge2/0 
Ge2/1 
Ge2/2 
.... 
eth_format: <prefix>{<first value>} eth_format: Gi{0} 
Gi0  
Gi1 
Gi2 
Gi3 
... 
eth_format: <prefix>{<first value>} eth_format: G0/{0} 
G0/0  
G0/1 
G0/2 
G0/3 
... 
eth_name: <prefix: Interface custom name> 
eth_name: 
- M1 
- T1 
- T2 
M1 
T1 
T2 
eth_name: <prefix: Interface custom name> 
eth_name: 
- MGMT 
- DATA 
- TRAFFIC 
MGMT 
DATA 
TRAFFIC 
 
Combined first named interface following by formatted interfaces Example: We have to 
set first node interface name “eth0/mgmt” and next following interfaces must start from eth1 
and change sequence accordingly. eth1, eth2,….,ethx 
As your node first interface will  be custom named (eth0/mgmt), therefore in the template 
“eth_name:” must be added before “eth_format:”

EVE-NG Community Cookbook 
Version 1.11 
Page 162 of 165 © EVE-NG LTD 
eth_name: 
- eth0/mgmt 
eth_format: eth{1} 
 
This adding will produce Node interfaces. 
 
12.3.4 Edit your new template file: 
For edit newly created template you ca n use WinSCP, FileZilla or cli. Example below shows 
template edit using cli and nano editor 
cd /opt/unetlab/html/templates/intel/ 
nano ngips.yml 
 
Change content , setting for various images can vary depends of vendor re quirements. The 
interface name lines please refer Section: 12.3.1 
 
 
# Copyright (c) 2016, Andrea Dainese 
# Copyright (c) 2018, Alain Degreffe 
# All rights reserved. 
# 
# Redistribution and use in source and binary forms, with or without 
# modification, are permitted provided that the following conditions are met: 
#     * Redistributions of source code must retain the above copyright 
#       notice, this list of conditions and the following disclaimer. 
#     * Redistributions in binary form must reproduce the above copyright 
#       notice, this list of conditions and the following disclaimer in the 
#       documentation and/or other materials provided with the distribution. 
#     * Neither the name of the UNetLab Ltd nor  the name of EVE-NG Ltd nor the 
#       names of its contributors may be used to endorse or promote products 
#       derived from this software without specific prior written permission. 
# 
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND 
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED 
# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE 
# DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY 
# DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES 
# (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; 
# LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND 
# ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT 
# (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS 
# SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

EVE-NG Community Cookbook 
Version 1.11 
Page 163 of 165 © EVE-NG LTD 
--- 
type: qemu 
name: NGIPS 
description: Cisco FirePower NGIPS 
cpulimit: 1 
icon: IPS.png 
cpu: 4 
ram: 8192 
ethernet: 3 
eth_name: 
- eth0/mgmt 
eth_format: eth{1} 
console: vnc 
shutdown: 1 
qemu_arch: x86_64 
qemu_version: 2.4.0 
qemu_nic: e1000 
qemu_options: -machine type=pc,accel=kvm -serial none -nographic -no-user-config  
  -nodefaults -display none -vga std -rtc base=utc -cpu host 
... 
 
Note: Qemu options in the line may vary per image requirements. Please check manufacturer 
advice how to run KVM image 
12.3.5 Prepare new icon for your template: 
Step 1 Use Filezilla or Winscp to copy your custom icon IPS.png (icon  filename IPS.png used 
in ngips.yml) 
 
This icon should be about 30-60 x 30-60 in the png format (switch.png is for example 65 x 33, 
8-bit/color RGBA) 
 
Step 2 Copy this new icon into /opt/unetlab/html/images/icons/ 
12.3.6 Template use 
 
Step 1 Create directory /opt/unetlab/addons/qemu/ngips-6.2.83  
 
mkdir /opt/unetlab/addons/qemu/ngips-6.2.83  
 
Step 4.2 Upload image NGIPS, Refer Section: 
  
 
 
12.4 How to hide unused images in the node list 
12.4.1 Creating new config.php file 
If your EVE Serve r does not have the config.php file in the /opt/unetlab/html/includes/ 
directory, then it must be created. 
Step 1. Use the EVE CLI. Make sure you are in the following EVE directory: 
/opt/unetlab/html/includes/ 
Step 2. Rename config.php.distributed (the template) to config.php.

EVE-NG Community Cookbook 
Version 1.11 
Page 164 of 165 © EVE-NG LTD 
mv config.php.distribution config.php 
 
12.4.2 Edit config.php file 
Step 1. Use vi or nano file editor to edit your config.php file. 
nano config.php 
 
