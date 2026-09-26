---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-21
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [3786, 3971]
sha256: 316bcac9e3400e12d5557c613b921792123c38b0872dd8539dcd4a3d67c858b4
---

# Copyright (c) 2016, Andrea Dainese

Step 1: MANDATORY: Configure your nodes and make sure you applied the vendor specific 
command to save the running configuration to NVRAM. If you do not save the configuration, it 
will not be exported and in the notification area, you will receive an error message stating the 
node cannot be exported. 
In this example the nodes have been configured with hostnames only and  the configurations 
have been saved to NVRAM. 
Step 2: In the example below a group of nodes were selected to export configurations. 
 
Step 3: Use “Export all CFGs” for selected nodes. Export configuration is 
completed. The notification area will display “Export All: done” when complete. 
10.3.4 Boot nodes from exported config set 
Step 1: Stop all nodes 
Step 2: Open sidebar and click Startup-configs. Make sure your config is set to ON and the 
nodes config switch is green  (switch on/off beside n ode). Press the green “Save” button (on 
the bottom) and all your nodes will boot with the exported config set after wiping them.

EVE-NG Community Cookbook 
Version 1.11 
Page 151 of 165 © EVE-NG LTD 
 
Step 3: Wipe nodes. For more information refer to section 8.1.3 
Step 4: Start nodes 
10.3.5 Edit exported configurations 
It is possible to edit your configurations for the nodes manually. 
Step 1: Select the node you want to edit the configuration o f and make your changes.  Click 
“Save” when you are finished. 
  
Step 2: Save the config for nodes with the green “Save” button on the bottom.  
 NOTE: you can manually copy/paste any configuration into the config set editor and 
apply it to  your node. Make sure your configuration interfaces match  the lab node’s 
interface names. 
10.3.6 Set lab to boot from none 
To reset your lab nodes’ configuration to factory default, follow the steps below: 
Step 1: Wipe nodes. Refer to section 10.3 for information about wiping nodes and the order of 
operations during boot.

EVE-NG Community Cookbook 
Version 1.11 
Page 152 of 165 © EVE-NG LTD 
Step 2: Open sidebar and click Startup -configs. Make sure your config is set to OFF and the 
nodes config switch is red (switch on/off beside node). Press the green “Save” button (on the 
bottom) and all your nodes will boot with no config/factory default after wiping them. 
 
Step 3: Start nodes 
10.3.7 Lab config script timeout 
Lab config script timeout is used when nodes are w aiting to boot from a config set. The node 
will literally wait during boot until the configuration is applied from the config set. 
Hit “More actions” and then “Edit lab” from the sidebar. Set the config script timeout in seconds. 
By default, this timer is set to 300 seconds for new labs.  
 NOTE: For heavy labs and nodes with 
long configurations, you can raise this 
timer to 600 seconds or higher.

EVE-NG Community Cookbook 
Version 1.11 
Page 153 of 165 © EVE-NG LTD 
11 EVE Troubleshooting 
11.1 CLI diagnostic information display commands 
11.1.1 Display full EVE Community diagnostic 
eve-info 
11.1.2 Display the currently installed EVE Community version: 
dpkg -l eve-ng 
 
11.1.3 Display if EVEs Intel VT-x/EPT option on/off: 
kvm-ok 
 
11.1.4 Display EVEs CPU INFO: 
lscpu 
 
11.1.5 Display EVEs HDD utilization.  
If the /boot only has a little space left you can refer to section 3.6.1.1. If the eve—ng—vg—root 
reaches 99% or 100% then you will need to expand the HDD in order to continue using EVE. 
The Solution to expand your HDD is described in section 11.1

EVE-NG Community Cookbook 
Version 1.11 
Page 154 of 165 © EVE-NG LTD 
df -h 
 
11.1.6 Display EVEs Bridge interface status 
brctl show 
 
11.1.7 Display EVEs system services status 
systemctl list-unit-files --state=enabled 
 
11.2 Expand EVEs System HDD 
 IMPORTANT NOTE: DO NOT expand your current/existing HDD on your EVE VM!

EVE-NG Community Cookbook 
Version 1.11 
Page 155 of 165 © EVE-NG LTD 
11.2.1 Expand HDD on VMware Workstation 
Expanding your EVEs system HDD is achieved by adding an additional HDD to your EVE VM. 
Step 1: Stop all your labs and shutdown EVE. 
Use EVE CLI command: shutdown -h now 
Step 2: Go to edit VM settings and add a new Hard driv e. 
Then click Next. 
Step 3: Leave the recommended SCSI HDD option and then 
click Next 
Step 4: Make sure you have selected the option “Create a 
new Virtual disk.”  
Step 5: Set your desirable HDD Size; example 200GB. 
Step 6: Make sure you have set  the option “Store Virtual disk as a single file” and then click 
Next  
Step 7: Optional: Specify the location of where your new HDD will be stored, then click Finish. 
Step 8:  Boot your EVE VM, HDD siz e will be expanded automatically. To verify, use the 
command to verify HDD utilization referenced in section 11.1.5 
11.2.2 Expand your HDD on ESXi 
Expanding your EVEs system HDD is achieved by adding an additional HDD to your EVE VM. 
Step 1: Stop all your labs and shutdown EVE. 
Use EVE CLI command: shutdown -h now 
Step 2:  Go to edit VM settings and add a new Hard 
drive. Then click Next  
Step 3: Make sure you have selected the option “Create 
a new Virtual disk.” Then click Next 
Step 4: Set your desirable HDD Size; example 200GB. 
Step 5: It is recommended to set the Thick Provision Lazy Zeroed HDD option. 
Step 6: Specify the location of where your new HDD will be stored and then click Next 
Step 7: Leave the recommended SCSI HDD option as is and click Finish. 
Step 8: Boot your EVE VM, the HDD size will be expanded automatically. To verify, use the 
command to verify HDD utilization referenced in section 11.1.5

EVE-NG Community Cookbook 
Version 1.11 
Page 156 of 165 © EVE-NG LTD 
11.2.3 Expand your HDD on a Bare Metal EVE Server 
It is a complicated process to expand a HDD for a bare metal EVE server. Please open a ticket 
in our Live chat support for advice. 
http://www.eve-ng.net/live-helpdesk 
Use a google account to join in the Live Chat or create new chat account. 
11.3 Reset Management IP 
Type the following commands into the CLI followed by enter: 
rm -f /opt/ovf/.configured 
 
su – 
 
http://www.eve-ng.net/documentation/installation/bare-installIP address setup wiza rd. Please 
follow the steps in section 3.5.1 for Static IP or 3.5.2  for DHCP IP setup. 
11.4 EVE Community SQL Database recovery 
Starting from EVE Community version 2.0.3-95, you can recover SQL user database in case 
of disaster: 
unl_wrapper -a restoredb 
 
11.5 EVE Log files 
EVE log Files can be obtained from the System Logs page under the System dropdown menu
 
Use the menu to collect log file data you are interested in.

EVE-NG Community Cookbook 
Version 1.11 
Page 157 of 165 © EVE-NG LTD 
 
11.6 EVE cli diagnostic info 
Use EVE cli to obtain your EVE information: 
eve-info

EVE-NG Community Cookbook 
Version 1.11 
Page 158 of 165 © EVE-NG LTD 
12 Images for EVE 
Images must be uploaded and prepared before they can be used in labs . The best way to 
upload images is to use the  WinSCP tool for Windows environment or FileZilla for MAC  OSX 
and Linux. 
Link to download WinSCP: 
https://winscp.net/eng/download.php 
Link to download FileZilla: 
https://filezilla-project.org/ 
To access EVE, use SSH protocol (port 22).  
Supported images for EVE are stored in the three locations: 
• IOL (IOS on Linux), /opt/unetlab/addons/iol/bin/ 
• Dynamips images, /opt/unetlab/addons/dynamips 
• Qemu images, /opt/unetlab/addons/qemu 
12.1 Qemu image naming table 
 IMPORTANT NOTE: Intel VT-X/EPT must be enabled to run Qemu nodes in EVE. For 
information on how to enable this option, Refer to section 3: EVE Installation. 
The directory names used for QEMU images are very sensitive and must match the table below 
exactly in order to work. 
 
Ensure your image folder name starts as per the table. After the " -" you can add whatever you 
like to label the image. We recommend using the version of your image.  
 
Folder name examples: 
 
firepower6-FTD-6.2.1 
acs-5.8.1.4 
 
