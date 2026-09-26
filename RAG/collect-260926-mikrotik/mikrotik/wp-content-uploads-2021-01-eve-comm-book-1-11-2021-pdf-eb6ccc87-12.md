---
id: collect-260926-mikrotik/mikrotik/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87-12
title: "Copyright (c) 2016, Andrea Dainese"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/wp-content-uploads-2021-01-eve-comm-book-1-11-2021-pdf-eb6ccc87.md
source_anchor: ""
source_lines: [1882, 2138]
sha256: ba422201bae69b5d471576c9ec30e0ea72b73387b3bfa6bde9dab10c824a327b
---

# Copyright (c) 2016, Andrea Dainese

Type the below commands followed by Enter 
 apt update 
In case the prompt asks to confirm with Y/N, answer Yes. 
4.2 EVE-NG Community Upgrade 
Type commands followed by Enter 
 apt upgrade 
In case the prompt asks to confirm with Y/N, answer Yes. 
 IMPORTANT NOTE:  If you are upgrading EVE Community from older version, the 
installation may ask you to confirm additional! Information: 
 
Answer for prompt above is “N”

EVE-NG Community Cookbook 
Version 1.11 
Page 68 of 165 © EVE-NG LTD 
 
Answer for grub-pc version is: “Keep the local version currently installed” 
After the completion of the update and upgrade, reboot your EVE  Server. Type the following 
command and hit enter. 
reboot

EVE-NG Community Cookbook 
Version 1.11 
Page 69 of 165 © EVE-NG LTD 
5 Types of EVE management consoles 
 IMPORTANT NOTE: EVE Console TCP ports. EVE Community uses a static port 
range between 32678-40000.  
Formula is = 32768+128*POD+1 -> 32768+128*POD+128    POD: user id ( admin = 0 ) 
Exemple: you got admin (POD 0) + 2 users ( POD 1, POD 2 )  
32768+128*0+1(First port for POD0)   -> 32768+128*2+128(Last port of POD 2 ) = 32769 -> 
33152 
Port per user pod: 
POD First Port Last Port 
0 32769 32896 
1 32897 33024 
2 33025 33152 
3 33153 33280 
4 33281 33408 
5 33409 33536 
6 33537 33664 
7 33665 33792 
8 33793 33920 
9 33921 34048 
10 34049 34176 
EVE Community supports two different console types. 
5.1 Native console 
EVE Native console option requires locally installed software 
to access your lab nodes. To use the Native console option, 
you must have Administrator rights on your PC and ensure 
the TCP port range 32768-40000 is not blocked by a firewall 
or antivirus software. (See table above) 
 
 
 
5.1.1 Native Console: telnet 
Windows OS: You can use your preferred telnet program like Putty, SecureCRT or others. 
Example: Putty as native telnet client on Windows. 
To setup Windows native telnet client please follow section 3.7.1

EVE-NG Community Cookbook 
Version 1.11 
Page 70 of 165 © EVE-NG LTD 
 
 
Linux OS: You can use your preferred telnet program like the Native Terminal, SecureCRT, 
or others. 
Example: Telnet client from the native terminal on Linux Mint.  To setup Linux native telnet 
client please follow section 3.7.2 
 
MAC OSX: You can use your preferred telnet program like the native Terminal, SecureCRT, 
or others. 
Example: Telnet client from the native terminal on MAC OSX.  To setup MAC OSX native 
telnet client please follow section 3.7.3 
5.1.2 Native Console: Wireshark 
EVE Community has an integrated connection with natively installed Wireshark software on 
your PC. This allows live captures with Wireshark installed on the client machine. The EVE 
will capture natively installed Wireshark session. 
 
 IMPORTANT NOTE: Make sure you have installed Wireshark and EVE-NG client 
pack. It is strongly recommended if your Wireshark software is installed at your PC 
default location.

EVE-NG Community Cookbook 
Version 1.11 
Page 71 of 165 © EVE-NG LTD 
 
 IMPORTANT NOTE: The Wireshark wrapper located in your PC station must match 
your EVE root password.  Edit your EVE root password in the wireshark_wrapper.bat, 
if you had changed it during install. 
 
 
 
Example: Fortinet live interface port1 capture.

EVE-NG Community Cookbook 
Version 1.11 
Page 72 of 165 © EVE-NG LTD 
 
 
5.1.3 Native Console: VNC 
Windows OS: Recommended and tested is UltraVNC but any other compatible one can be 
used. 
Example: UltraVNC as Native VNC client on Windows.  To setup Windows native VNC client 
please follow section 3.7.1 
 
Linux OS: Remote Desktop Viewer for VNC Sessions. 
Example: Remote Desktop Viewer for VNC sessions on Linux Mint.  To setup Linux native 
Remote Desktop Viewer please follow section 3.7.2

EVE-NG Community Cookbook 
Version 1.11 
Page 73 of 165 © EVE-NG LTD 
 
MAC OSX: Preferred VNC program: Chicken VNC 
Example: Chicken VNC as Native VNC client on MAC OSX.  To setup MAC OSX native RDP 
Viewer client please follow section 3.7.3 
5.1.4 Native Console: RDP 
Windows OS: Windows Native RDP. 
Example: Windows RDP session to Win10 host in the lab.   
 
Linux OS: Remote Desktop Viewer as RDP session to lab Win10 host. 
Example: RDP session to Win10 host in the lab. To setup Linux native Remote Desktop 
Viewer please follow section 3.7.2

EVE-NG Community Cookbook 
Version 1.11 
Page 74 of 165 © EVE-NG LTD 
MAC OSX: Remote Desktop Viewer as RDP session to lab Win10 host. 
Example: RDP session to Win10 host in the lab.   
To setup MAC OSX native RDP Viewer client please follow section 3.7.3 
 
5.2 HTML5 console 
The EVE Community HTML5 console 
provides a clientless solution for 
managing labs and node sessions. 
Management is achieved directly through 
the browser by opening new browser 
window. It is very convenient for 
Corporate users with restricted 
Workstation permissions (Locked Telnet, 
vnc, rdp). 
 
 
 
 
 
5.2.1 HTML5 Console: Telnet 
HTML5 Telnet console opens telnet sessions in the new browser window.

EVE-NG Community Cookbook 
Version 1.11 
Page 75 of 165 © EVE-NG LTD 
5.2.2 HTML5 Console: VNC 
HTML5 VNC opens VNC sessions in the new browser window. 
 
5.2.3 HTML5 Console: RDP for Windows 
HTML5 RDP console opens RDP session s in the new browser window. For Windows  7, 
Windows Server 2008. 
During Windows machine image installation, you can allow RDP sessions to be used for access 
to Windows host. If your Windows host has enabled RDP session, edit windows node settings 
and set RDP console. Give time to boot this node and RDP session will opens in new browser 
tab.

EVE-NG Community Cookbook 
Version 1.11 
Page 76 of 165 © EVE-NG LTD

EVE-NG Community Cookbook 
Version 1.11 
Page 77 of 165 © EVE-NG LTD 
6 EVE WEB GUI Management 
6.1 EVE Management Page 
The Main EVE management window 
 
6.1.1 Management buttons 
 
Button Description 
 
Select All or Deselect All folders or labs in the EVE tree 
 
Create/Add new Lab 
 
Change selected item name. To use this option, please select the folder or lab 
that you want to rename. You must not rename the Shared folder, the Users 
folder or any folder inside the Users folder. 
 
Move selected item(s) to a different location. To use this option, please select 
the folder(s) or lab(s) that you want to move.

EVE-NG Community Cookbook 
Version 1.11 
Page 78 of 165 © EVE-NG LTD 
 
Delete selected folders or labs. You must not d elete the Shared folder, the 
Users folder or any folder inside the Users folder. 
 
Import an EVE lab or lab folder  from a previous export . Import file must be in 
.zip format 
 
Export EVE lab or folder. Select folder(s) and /or labs you wish to export and 
select this option. The export is saved to your local PC in .zip format and is 
ready to import to another EVE. 
 
Toggle the s orting folders and labs b etween alphabetical and last edit date  
(ascending/descending cannot be changed currently). 
 
Refresh current folder content 
6.1.2 Management tabs 
 
 
Tab Description 
 
Returns back to the EVE Home Management screen. 
 
 
 
Management dropdown, opening the management submenu. 
 
Management submenu, refer to sections: 6.3 
 
System dropdown. 
 
System submenu, refer to section 6.4

EVE-NG Community Cookbook 
Version 1.11 
Page 79 of 165 © EVE-NG LTD 
 
 
 
Information dropdown 
 
Information submenu, for details see section 6.5 
6.2 Folders and Lab files management 
This section will explain how to manage folders and labs on the EVE management page. 
6.2.1 Folders Management 
6.2.1.1 Create folder 
Type the new folder name and click “Add Folder” 
6.2.1.2 Delete folder 
Select the f older you wish to delete and press 
Delete.  
 NOTE: All folder content will be deleted as 
well. 
 
 
6.2.1.3 Move Folder 
Select the f older you wish to move and press the 
Move to button.

EVE-NG Community Cookbook 
Version 1.11 
Page 80 of 165 © EVE-NG LTD 
 
 
Type and select the target destination for your folder and 
confirm by clicking on Move. 
 
