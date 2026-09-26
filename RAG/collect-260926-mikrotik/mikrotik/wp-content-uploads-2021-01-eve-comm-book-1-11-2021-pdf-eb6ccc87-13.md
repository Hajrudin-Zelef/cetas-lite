---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-13
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [2139, 2352]
sha256: ca4df5632d0785a661c4a3efcf8ab828d3fa5e2588972f51276be5927a812c61
---

# Copyright (c) 2016, Andrea Dainese

 
6.2.1.4 Export Folder 
Select the folder(s) you wish to export from your EVE and press Export.  
 
 
Save the exported file as .zip to your local PC. The exported zip file is ready to import to another 
EVE instance. 
 
If your browser is set to save downloaded files to a default directory, your exported file will be 
saved in the browsers default downloads directory. 
6.2.1.5 Import Folder 
 IMPORTANT: Importable file MUST be in .zip format, do NOT unzip the file. 
Step 1: Press the Import button.

EVE-NG Community Cookbook 
Version 1.11 
Page 81 of 165 © EVE-NG LTD 
 
Step 2: Choose the zipped file that contains EVE folders with labs. 
 
Step 3: Press the Upload Button 
 
Step 4: After you made sure your folder is imported and has all its content (labs), you can close 
the upload session. 
 
6.2.2 Lab files Management 
You can manage created labs from the main EVE file manager window

EVE-NG Community Cookbook 
Version 1.11 
Page 82 of 165 © EVE-NG LTD 
 
6.2.2.1 Create Lab 
Click on the New Lab button and refer to section 8.1 
 
6.2.2.2 Delete Lab 
Select the lab or labs you wish to delete and then press the Delete button 
 
6.2.2.3 Clone Lab 
The cloning feature provides a very convenient way to duplicate original labs to share with 
others or base another lab on it. 
Cloned labs will copy exported configs (on supported nodes ) but  will not copy saved 
states/configurations in Qemu nodes like Windows hosts, Cisco ISE, or other Qemu nodes that 
are not supported by the export config feature. Please refer to section 10.1 for more information 
on configuration export for labs. 
Step 1: Select the lab you wish to clone and move the mouse pointer (blue) to that lab, an extra 
option will appear. Click on Clone.

EVE-NG Community Cookbook 
Version 1.11 
Page 83 of 165 © EVE-NG LTD 
 
Step 2: Your lab will be cloned with all you r exported configurations or configuration sets with 
a new name. 
 
Step 3: The lab has been cloned lab and can be renamed to your liking. Move the mouse pointer 
to the cloned lab and choose Rename. 
 
Step 4: Rename it, and click OK to confirm 
 
6.2.2.4 Move Lab 
Step 1: Select the lab you wish to Move and move the mouse pointer (blue) to that lab, an extra 
option will appear. Choose Move to. 
 
Step 2:  Type the path to the new destination and 
confirm by clicking Move 
   
 
 
6.2.2.5 Export Lab 
Select the Lab(s) you wish to export from your EVE Server and press Export.

EVE-NG Community Cookbook 
Version 1.11 
Page 84 of 165 © EVE-NG LTD 
 
Save exported file as .zip to your local PC. The exported zip file is ready to import into another 
EVE. 
 
If your browser is set to save downloaded files to default directory, your exported file will be 
saved in the browsers default downloads directory. 
6.2.2.6 Import Labs 
 IMPORTANT: Importable file MUST be in .zip format, do NOT unzip the file.  
Step 1: Press the Import button. 
 
Step 2: Choose the zipped file which contains the EVE labs.

EVE-NG Community Cookbook 
Version 1.11 
Page 85 of 165 © EVE-NG LTD 
 
Step 3: Press the Upload Button 
 
Step 4: After you made sure your lab is imported, you can close the upload session. 
 
6.3 EVE Management Dropdown Menu 
6.3.1 EVE User management 
The User Management page, under the 
Management dropdown, will allow Admin 
accounts to manage other user accounts. 
 
6.3.1.1 Creating a new EVE User 
Step 1: Open the User management submenu. Management>User management and click Add 
user

EVE-NG Community Cookbook 
Version 1.11 
Page 86 of 165 © EVE-NG LTD 
 
Step 2: The Add New User management window will pop up. Fill in the main information about 
your EVE user 
 
Step 3: The POD number is a value assigned to user accounts automatically. POD number s 
are like user profiles inside of EVE and are a unique value for every user Think of PODs like a 
virtual rack of equipment for each user. Admins can assign a preferred number between 1-128. 
Please keep POD numbers unique between users! 
Step 4: Press ADD 
  
6.3.1.2 Edit EVE User 
Step 1: Open the User management submenu. Management -> User management and choose 
which user you want to edit.

EVE-NG Community Cookbook 
Version 1.11 
Page 87 of 165 © EVE-NG LTD 
Step 2:  The Edit user management window will pop up. Now you can edit necessa ry user 
information, roles , or access time . Confirm settings by pressing Edit at the bottom of the 
window. 
 
6.3.1.3 User monitoring 
There is a dropdown menu next to “Add User” called “More Info” that can provide additional 
information about your users. Click the checkbox next to the relevant information that you would 
like displayed. Additional columns will be added for each checkbox that is chosen. 
 
6.4 EVE System Dropdown menu 
 The EVE System dropdown contains the system utilization status, log 
files, and an option to stop all running nodes on the server.

EVE-NG Community Cookbook 
Version 1.11 
Page 88 of 165 © EVE-NG LTD 
6.4.1 System status 
The System Status page, under the System Dropdown, will show EVE 
server resource utilization, the number of  running nodes per template,  
current running versions of EVE and Qemu , and the current sta tus of the 
UKSM and CPU Limit options. 
 
UKSM – “Ultra KSM (kernel same-page merging) is a Linux kernel feature that allows the KVM 
hypervisor to share identical memory pages among different process or virtual machines on the 
same server.” It can be disabl ed globally for EVE on this page. It is recommended to keep 
UKSM enabled.  
CPU Limit – CPU limit is used to limit CPU 
overloads during the nodes run time. It acts 
like a smart CPU usage option. If a running 
node reaches 80% CPU u tilization, the CPU 
Limit feature throttles CPU use for this node 
to 50% until process usage drop s under 30% 
for a period of 1 minute. 
It is recommended to keep the Global CPU 
Limit option enabled. 
CPU Limit can be turned for individual nodes 
in a lab. EVE node templates are set , by 
default, with the recommended CPU limit 
settings. An Unchecked CPU Limit option 
means that this node will boot without CPU 
limit. 
Reference: 
https://searchservervirtualization.techtarget.com/definition/KSM-kernel-samepage-merging  
6.4.2 System logs 
The System logs page, under the System Dropdown, will display EVE 
server log information

EVE-NG Community Cookbook 
Version 1.11 
Page 89 of 165 © EVE-NG LTD 
 
In the menu you can select a specific log file for inspection. 
 
6.4.3 Stop All Nodes 
The Stop All Nodes option, under the System Dropdown, is an 
option that stops all running nodes on the EVE server.  This 
option is accessible only by Admin users. 
  
6.5 EVE Information Dropdown menu 
The Eve Information Dropdown contains links to  the EVE 
Website, EVE forum, EVE YouTube channel , and the web-
based EVE Live Help chat. 
To join the EVE Forum, in order to make posts or download 
materials, a forum user account must be created.  
To join the EVE Live Chat for support, please use your Google account for access, or create a 
new user account for this chat. Please note the forum and live chat use separate user accounts. 
6.6 Other Tab line info 
 
Other items on the top menu are: Real-time clock, a shortcut to edit the currently logged in user, 
and a sign-out button. 
6.7 Lab preview and global settings 
Once you click on a lab in the folder tree, a main window on the right side will display schematic 
content of the lab as well as lab management options like open, edit, and delete.

EVE-NG Community Cookbook 
Version 1.11 
Page 90 of 165 © EVE-NG LTD 
 
6.7.1 Lab preview window 
The lab preview window display s the schematic position of nodes and their connectivity . The 
Scale option allows you change the lab preview size. 
 
6.7.2 Lab preview buttons 
In the lab preview, these buttons allow you to manage the selected lab. 
Button Description 
 Opens the Lab to the Topology Canvas  
 Opens the Labs Global Settings. Refer to section 6.7.4 for more info. 
 Deletes the lab

