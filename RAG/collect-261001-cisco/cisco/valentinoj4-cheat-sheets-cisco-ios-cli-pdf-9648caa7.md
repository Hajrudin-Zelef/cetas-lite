---
id: collect-261001-cisco/cisco/valentinoj4-cheat-sheets-cisco-ios-cli-pdf-9648caa7
title: "valentinoj4-cheat-sheets-cisco-ios-cli-pdf-9648caa7"
domain: cisco
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["dram"]
source: docs/RAG/collect-261001-cisco/valentinoj4-cheat-sheets-cisco-ios-cli-pdf-9648caa7.md
source_anchor: ""
source_lines: [1, 336]
sha256: af93f058c086f21a92f719ee9fcf76527ee0597912faebcf3eace7e19d52a986
---

# valentinoj4-cheat-sheets-cisco-ios-cli-pdf-9648caa7

Cisco IOS CLI Cheat Sheet
by 
ValentinoJ4
 via 
cheatography.com/198096/cs/41870/
Basic Device Config​uration
Basic Device Config​uration
Enter Privie​leged Exec mode
Switch> 
enable
enable
Leave Privie​leged Exec mode
Switch# 
disable
disable
Enter Global Config​uration mode
Switch# 
configure terminal
configure terminal
Config​uring a Device Name
(Hostname)
Switch​(co​nfig)# 
hostname
hostname
NAME
Config​uring a Console Port
(Console)
Switch​(co​nfig)# 
line console
line console
 0
Config​uring a Console Password
(Console)
Switch​(co​nfi​g-l​ine)# 
password
password
PASSWORD
Activating password checking on
the Console Port
Switch​(co​nfi​g-l​ine)# 
login
login
Activating Userna​me/​Pas​sword
checking on the Console Port
Switch​(co​nfi​g-l​ine)# 
login local
login local
Creating a timeout on the Console
Port
Switch​(co​nfi​g-l​ine)# 
exec-t​‐
exec-t​‐
imeout
imeout
 MINUTES SECONDS
Backup one level
Switch​(co​nfi​g-l​ine)# 
exit
exit
Backup one level
Switch​(co​nfig)# 
exit
exit
Backup all levels
Switch# 
CTRL + Z
CTRL + Z
 or 
CTRL + C
CTRL + C
The Show Commands to Know!
The Show Commands to Know!
Shows inform​​ation about the switch and its
interf​​aces, RAM, NVRAM, flash, IOS, etc
Switch# 
show
show
version
version
Shows the current config​​ur​ation file stored in
DRAM.
Switch# 
show
show
runnin​​g-​c​onfig
runnin​​g-​c​onfig
Shows the config​​ur​ation file stored in NVRAM
which is used when the device boots.
Switch# 
show
show
startu​p-c​​onfig
startu​p-c​​onfig
Shows an overview of all interf​​aces, their
physical status, protocol status and ip address if
assigned.
Switch# 
show ip
show ip
interface brief
interface brief
Shows any Descri​ptions you've configured on
your individual interf​aces.
Switch# 
show
show
interfaces
interfaces
descri​​ption
descri​​ption
Shows the status of all interfaces like connected
or not, speed, duplex, trunk or access VLAN.
Switch# 
show
show
interfaces status
interfaces status
 
Awesome Shortcuts
Awesome Shortcuts
Create shortcuts for long
commands
Router​(co​​nfig)# 
alias exec
alias exec
SHIP 
show ip interface brief
show ip interface brief
Configure a Banner Message that
shows "​eve​ryw​her​e"
Router​(co​​nfig)#
 banner motd 
 banner motd 
$
THIS IS MY MESSAGE $
Encrypt all plain-text passwords
stored on the devic
Router​(co​​nfig)# 
service
service
passwo​rd-​enc​ryption
passwo​rd-​enc​ryption
Stop those "​pop​-up​" messages
from cutting through your CLI
Router​(co​​nfig)#
 logging synchr​‐
 logging synchr​‐
onous
onous
Save your config​uration
Router# 
wr 
wr 
or 
write 
write 
or 
copy run
copy run
star
star
Some Fun "​Stu​die​s"
Some Fun "​Stu​die​s"
LearnCisco (Confi​guring a Cisco Router)
A Very Thorough Deep
Dive
Lock down your Cisco Router (from the
NSA?)
Hardening a Cisco
Router
Config​uring a Switch Management Interface
Config​uring a Switch Management Interface
Access a specific switch VLAN
interface SVI (common VLAN 1)
Switch​(co​nfig)#
 interface
 interface
VLAN
VLAN
 #
Configure a reachable IP address
and Subnet Mask
Switch​(co​nfig-if )# 
ip address
ip address
ADDRESS MASK
Activate the Switch management
interface
Switch​(co​nfig-if )#
 no shut
 no shut
Exit the switch VLAN interface
(SVI)
Switch​(co​nfig-if )#
 exit
 exit
Configure a default gateway for the
switch to send upstream traffic out
of the local LAN
Switch​(co​nfig)# 
ip defaul​t-g​‐
ip defaul​t-g​‐
ateway 
ateway 
IP-ADD​RES​S-O​F-U​‐
PST​REA​M-R​OUTER
Config​uring a Router Network Interface
Config​uring a Router Network Interface
Locate your Router interfaces
(learn their design​ations)
Router# 
show ip interface
show ip interface
brief
brief
Access Global Config​uration mode
Router# 
configure terminal
configure terminal
Access a specific Router interface
Router​(co​nfig)#
 interface
 interface
gi0/0/0
Configure a reachable IP address
and Subnet Mask
Router​(co​nfi​g-if)# 
ip address
ip address
ADDRESS MASK
Activate the Router interface
Router​(co​​nf​ig-if )#
 no shut
 no shut
By 
ValentinoJ4
ValentinoJ4
cheatography.com/valentinoj4/
 
Published 28th December, 2023.
Last updated 28th December, 2023.
Page 1 of 2.
 
Sponsored by 
Readable.com
Readable.com
Measure your website readability!
https://readable.com

Cisco IOS CLI Cheat Sheet
by 
ValentinoJ4
 via 
cheatography.com/198096/cs/41870/
Config​​uring SSH for Remote Management
Config​​uring SSH for Remote Management
Configure the device hostname
Switch​(co​nfig)# 
hostname
hostname
NAME
Configure the Doman Name the device
will operate on
Switch​(co​​nfig)# 
ip
ip
domain​​-name
domain​​-name
 EXAMPL​‐
E.COM
Configure a username and password
for remote management
Switch​(co​​nfig)# 
username
username
admin 
password
password
 cisco
Generate encryption keys to "​obf​usc​‐
ate​" the management traffic
Switch​(co​​nfig)# 
crypto key
crypto key
generate rsa
generate rsa
Configure a minimum of 1024 bits for
encryption security
How many bits in the
modulus [512]: 1024
Define the SSH version to use (older
versions have exploi​ts/​vul​ner​abi​lities)
Switch​(co​​nfig)#
 ip ssh
 ip ssh
version
version
 2
Access the VTY lines (used for Telnet
and SSH)
Switch​(co​​nfig)#
 line vty 
 line vty 
0
4
Activate SSH on the VTY lines
Switch​(co​​nf​ig-​line)#
transport input
transport input
 ssh
Require a username and password
combo
Switch​(co​​nf​i​g​-l​​ine)# 
login
login
local
local
Save your config​uration
Switch# 
wr
wr
 
By 
ValentinoJ4
ValentinoJ4
cheatography.com/valentinoj4/
 
Published 28th December, 2023.
Last updated 28th December, 2023.
Page 2 of 2.
 
Sponsored by 
Readable.com
Readable.com
Measure your website readability!
https://readable.com
