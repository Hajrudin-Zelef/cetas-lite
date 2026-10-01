---
id: collect-261001-cisco/cisco/theleran-cheat-sheets-cisco-ios-cli-pdf-6b241f43-1
title: "show cdp"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/theleran-cheat-sheets-cisco-ios-cli-pdf-6b241f43.md
source_anchor: ""
source_lines: [1, 449]
sha256: 06bee1d74e4387220c5f7b7d2a6d66aeb42bc2de1227c47c8c003cc5b440bb8b
---

# show cdp

Cisco IOS CLI Cheat Sheet
by 
Theleran
 via 
cheatography.com/77264/cs/18961/
Gene​ral
Gene​ral
> 
en
en
Enter Privleged Exec Mode
#
#
# 
config t
config t
Enter Global Config Mode
(config)#
(config)#
(config) # 
int 
int 
{type}
{type}
{number}
{number}
Enter Interface Config
Mode
(confi​g-if)#
(confi​g-if)#
(config) # 
vlan 
vlan 
{number}
{number}
Enter VLAN Config Mode
(confi​g-v​‐
(confi​g-v​‐
lan)#
lan)#
(config) # 
line con 0
line con 0
Enter Console Line Config
Mode
(confi​g-l​‐
(confi​g-l​‐
ine)#
ine)#
(config) # 
line vty 0 15
line vty 0 15
Enter VTY Line Config
Mode
(confi​g-l​‐
(confi​g-l​‐
ine)#
ine)#
(config) # 
no ip dom lo
no ip dom lo
Stops Router Domain Lookup
(config) # 
undebug all
undebug all
Stops all Debugs
# 
clock set {time} {date}
clock set {time} {date}
Sets manual Time/Date
# 
show file systems
show file systems
Lists available file systems
# 
exit
exit
Exits current mode/level
Hous​eke​eping
Hous​eke​eping
(config) # 
ho {name}
ho {name}
Set name of device
(config) # 
ena sec {password}
ena sec {password}
Set encypted password for Priv
Exec Mode
(config) # 
ser pass
ser pass
Encrypts All Passwords
(config) # 
banner motd #{Banner}#
banner motd #{Banner}#
Creates Message banner
(config) # 
security pass min
security pass min
{number}
{number}
Sets min password length
(config) # 
login block-for {time}
login block-for {time}
attempts {attempts} within {time}
attempts {attempts} within {time}
Login failure wait time set
(confi​g-line) # 
pass {password}
pass {password}
Sets password for Console
Line
(confi​g-line) # 
login
login
Makes passwords active, use
after every password config
 
Hous​eke​eping (cont)
Hous​eke​eping (cont)
(confi​g-line) # 
exec-t​imeout
exec-t​imeout
{time}
{time}
Sets login timeout
(confi​g-if) # 
shut
shut
 | 
no shut
no shut
Enables | Disables interface
(confi​g-if) #
des {descr​iption}
des {descr​iption}
Sets descri​ption of interface
# 
cop r s
cop r s
Copies Running Config to the
NVRAM
SSH Config
SSH Config
(config) # 
ip domain​-name
ip domain​-name
{Abxyz.com}
{Abxyz.com}
Sets Domain Name
(config) # 
cry key gen rsa genera​l-keys
cry key gen rsa genera​l-keys
mod 1024
mod 1024
Configs complexity of keys
(config) # 
username {name} secret
username {name} secret
{password}
{password}
Sets a UN & encrypted
Pass
(config) # 
line vty 0 15
line vty 0 15
Configs which VTY lines to
use
(confi​g-line) # 
login local
login local
Sets LOGIN
(confi​g-line) # 
transport input ssh
transport input ssh
*Defines transport potocol
to SSH
IP Routing Config
IP Routing Config
(config) # 
ip route
ip route
 
networ​k-a​‐
ddress
 
subnet​-mask
 {
ip-address
 |
exit-intf
 }
Static Route Command
(config) # 
ip route
ip route
 0.0.0.0 0.0.0.0
{
ip-address
 | 
exit-intf
 }
Default Static Route Command
(config) # 
ip route
ip route
 
networ​k-a​‐
ddress
 
subnet​-mask
 {
admin-​dis​‐
tance
 }
Floating Static Route Command
(Admin distance default value is
1)
VLAN Config
VLAN Config
(config) # 
vlan {vlan-id}
vlan {vlan-id}
Create a VLAN
(confi​g-vlan) # 
name {vlan-​‐
name {vlan-​‐
name}
name}
Specify a unique name to identify the
VLAN
(confi​g-vlan) # 
end
end
Return to the privileged EXEC mode
(config) # 
interface {inter​fac​‐
interface {inter​fac​‐
e_id}
e_id}
Enter interface config​uration mode
By 
Theleran
Theleran
cheatography.com/theleran/
 
Not published yet.
Last updated 27th February, 2019.
Page 1 of 4.
 
Sponsored by 
Readable.com
Readable.com
Measure your website readability!
https://readable.com
Command
Descri​ption
Displays As
Command
Descri​ption
Command
Descri​ption
Command
Descri​ption
Command
Descri​ption

Cisco IOS CLI Cheat Sheet
by 
Theleran
 via 
cheatography.com/77264/cs/18961/
VLAN Config (cont)
VLAN Config (cont)
(confi​g-if) # 
switchport mode access
switchport mode access
Set the port to access
mode.
(confi​g-if) # 
switchport access vlan
switchport access vlan
vlan_id
Assign the port to a
VLAN.
(confi​g-if) # 
end
end
Return to the privileged
EXEC mode.
(config) #
show vlan brief
show vlan brief
Display the contents of
the vlan.dat file
(confi​g-if) # 
mls qos trust
mls qos trust
 
[
cos
cos
 | device
cisco-​phone | dscp | ip-pre​ced​ence]
Set the trusted state of an
interface
(confi​g-if) # 
switchport voice vlan
switchport voice vlan
 
vlan-#
Assign a voice VLAN to a
port
(confi​g-if) # 
switchport mode trunk
switchport mode trunk
Configure a switch port
on one end of a trunk link
(confi​g-if) # 
switchport trunk native
switchport trunk native
 
vlan
#
Configure native VLAN
For a Catalyst switch, the 
erase startu​p-c​onfig
erase startu​p-c​onfig
 command must
accompany the 
{(config) #
delete vlan.dat
delete vlan.dat
} command prior to reload to
restore the switch to its factory default condition.
PAT Config
PAT Config
(config)# ip nat pool NAT-PO​‐
OL-​OVE​RLOAD 209.16​5.2​‐
00.241 209.16​5.2​00.250
netmask 255.25​5.2​55.224
Define a pool of public IPv4
addresses 209.16​5.2​00.241 to
209.16​5.2​00.250 with pool name
NAT-PO​OL-​OVE​RLOAD.
(config)# access​-list 3 permit
10.0.0.0 0.255.2​55.255
Configure ACL 3 to permit devices
from 10.0.0.0/8 network to be
translated by NAT.
(config)# ip nat inside source
list 3 pool NAT-PO​OL-​OVE​‐
RLOAD overload
Bind NAT-PO​OL-​OVE​RLOAD with
ACL 3.
(config)# interface Serial​0/0/0
R2(con​fig​-if)# ip nat inside
Configure the proper inside NAT
interface.
R2(con​fig)# interface Serial​‐
0/1/0 
R2(con​fig​-if)# ip nat
outside
Configure the proper outside NAT
interface.
 
CDP Config
CDP Config
# show cdp
Display the status of CDP on R1.
R1# configure terminal
R1(con​fig)# cdp run R1(con​‐
fig)# interface s0/0/0 R1(con​‐
fig​-if)# no cdp enable R1(con​‐
fig​-if)# end
Enter Global Config​urE​nable CDP
globally on R1. Disable CDP on
interface S0/0/0. Use end
command to exit Global Config​‐
uration mode.
# show cdp neighbors
Display the list of CDP neighbors on
R1.
# show cdp neighbors detail
Display more details from the list of
CDP neighbors on R1.
Clock & NTP Config
Clock & NTP Config
# show clock detail
Display the clock
(config)# clock
timezone PST -8
R1(con​fig)# Clock
summer​-time PDT
recurring
Set the clock time zone to PST (Pacific
Standard Time), which is 8 hours later than
GMT (-8). Set PDT (Pacific Daylight Time)
to summer time recurring.
(config)# ntp server
209.16​5.2​00.225
Configure R1 to use an external public NTP
server with an IP address of 209.16​5.2​‐
00.225.
# show ntp associ​‐
ations
Verify that R1 is associated with the NTP
server at IP address 209.16​5.2​00.225.
LLDP Config
LLDP Config
# show lldp
Display the status of LLDP
(config)# lldp run R1(con​fig)#
interface s0/0/0 R1(con​fig​-if)# no
lldp transmit
Enable LLDP globally on R1.
Disable LLDP on interface
S0/0/0.
# show lldp neighbors
Display the list of LLDP
neighbors
# show lldp neighbors detail
Display more details from the
list of LLDP neighbors
By 
Theleran
Theleran
cheatography.com/theleran/
 
Not published yet.
Last updated 27th February, 2019.
Page 2 of 4.
 
Sponsored by 
Readable.com
Readable.com
Measure your website readability!
https://readable.com
Command
Descri​ption
Command
Descri​ption
Command
Descri​ption
Command
Descri​ption

