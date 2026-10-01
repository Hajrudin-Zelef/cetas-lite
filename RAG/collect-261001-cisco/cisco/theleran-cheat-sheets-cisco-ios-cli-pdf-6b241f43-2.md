---
id: collect-261001-cisco/cisco/theleran-cheat-sheets-cisco-ios-cli-pdf-6b241f43-2
title: "show cdp"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/theleran-cheat-sheets-cisco-ios-cli-pdf-6b241f43.md
source_anchor: ""
source_lines: [450, 823]
sha256: 9750ae575bdbcdfb723bc391ad8e93aa2768f28737bd5738ab548c24459cce08
---

# show cdp

Cisco IOS CLI Cheat Sheet
by 
Theleran
 via 
cheatography.com/77264/cs/18961/
Shortcuts
Shortcuts
Tab
Tab
Completes current abbrv command
Up Arrow
Up Arrow
Cycles thru previously used commands
?
?
Access HELP
Ctrl+S​hift+6
Ctrl+S​hift+6
Interupt
Ctrl+C
Ctrl+C
Exits Config
Ctrl+Z
Ctrl+Z
Applies command, returns to Priv Exec
Display / Show Commands
Display / Show Commands
# 
sho run
sho run
Displays Running Configs.
# 
sho access-l
sho access-l
Displays all ACL's.
# 
sho access-l {name/​‐
sho access-l {name/​‐
number}
number}
Displays only denoted ACL.
# 
sho ipv6 int
sho ipv6 int
Displays interfaces on IPv6
# 
sho ip route
sho ip route
Displays all routes attached to router
# 
sho ip route static
sho ip route static
Displays all static routes attached to
router
#
Sho ip route
Sho ip route
 
network
Displays routes only associated with that
network
show ip nat transl​ations
nat
Access Control Lists
Access Control Lists
(config) # 
access​-list _ { deny | permit | remark }
access​-list _ { deny | permit | remark }
{sourc​e+w​ild​card}
{sourc​e+w​ild​card}
Create ACL.
(config) # 
ip access​-list standard {name}
ip access​-list standard {name}
Create Named
ACL.
(confi​g-if) # 
ip access​-group { access​-li​st-​number
ip access​-group { access​-li​st-​number
| access​-li​st-name } { in | out }
| access​-li​st-name } { in | out }
Attach ACL to
an Interface.
(confi​g-line) # 
access​-class {number} { in | out }
access​-class {number} { in | out }
ACL for VTY.
Wildcard Determined by 255.25​5.2​55.2​55​-Subnet mask (ex 255.25​‐
5.2​55.2​55​-25​5.2​55.2​55.128= Wildcard of 0.0.0.127)
Shortcuts; 
host
host
 = Wilcard of 255.25​5.2​55.255 
any
any
 = Address &
WIldcard of 0.0.0.0 0.0.0.0
 
DHCPv4 Config
DHCPv4 Config
(config) #
ip dhcp exclud​ed-​address {low ip
ip dhcp exclud​ed-​address {low ip
range} {high ip range} | {single ip}
range} {high ip range} | {single ip}
Excludes ip ranges,
or single IP's.
(config) # 
ip dhcp pool {name}
ip dhcp pool {name}
Creates named
DHCP pool
(dhcp-​config) #
net {ipv4net} {subnet}
net {ipv4net} {subnet}
Define Range of
Addresses
(dhcp-​config) #
default-r {gateway}
default-r {gateway}
Sets Default
Gateway
(dhcp-​config) #
dns-s {DNS}
dns-s {DNS}
Sets DNS
(dhcp-​config) #
domain-n {Axyz.com}
domain-n {Axyz.com}
Sets Domain
(config) #
ip helper​-ad​dress {ipv4net}
ip helper​-ad​dress {ipv4net}
Sets DHCP Relay
DHCPv6 Config
DHCPv6 Config
(confi​g-if) # 
ipv6 unicas​t-r​outing
ipv6 unicas​t-r​outing
Enable IPv6
(confi​g-if) # 
ipv6 dhcp pool {name}
ipv6 dhcp pool {name}
Name Pool
(confi​g-if) # 
address prefix {prefix length}
address prefix {prefix length}
lifetime {infinite | time}
lifetime {infinite | time}
Statefull Only
(confi​g-if) # 
dns-s {IPv6DNS}
dns-s {IPv6DNS}
Set IPv6 DNS
(confi​g-if) # 
domain-n {Axyz.com}
domain-n {Axyz.com}
Set Domain
(confi​g-if) # 
ipv6 dhcp server {name}
ipv6 dhcp server {name}
Set Server Name
See Note Below for Final CMD
See Note Below for Final CMD
 
(confi​g-if) # 
ipv6 dhcp relay destin​ation
ipv6 dhcp relay destin​ation
{ipv6net}
{ipv6net}
Sets Router as a
DHCPv6 Relay
# 
debug ipv6 dhcp detail
debug ipv6 dhcp detail
Displays debug
details
SLAAC
 
(confi​g-if) # 
no ipv6 nd manage​d-c​onf​ig-flag
no ipv6 nd manage​d-c​onf​ig-flag
 
(confi​g-if) # 
no
no
ipv6 nd other-​con​fig​-flag
ipv6 nd other-​con​fig​-flag
 
Note: No other config required for SLAAC.
Stateless DHCPv6
 
(confi​g-inf) # 
ipv6 nd other-​con​fig​-flag
ipv6 nd other-​con​fig​-flag
Statefull DHCPv6
 
(confi​g-inf) # 
ipv6 nd manage​d-c​onf​ig-flag
ipv6 nd manage​d-c​onf​ig-flag
By 
Theleran
Theleran
cheatography.com/theleran/
 
Not published yet.
Last updated 27th February, 2019.
Page 3 of 4.
 
Sponsored by 
Readable.com
Readable.com
Measure your website readability!
https://readable.com
Keys
Action
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
RIP Config
RIP Config
(config) # 
router rip
router rip
(confi​g-r​outer) #
version 2
version 2
(confi​g-r​outer)
# 
no auto-s​‐
no auto-s​‐
ummary
ummary
modify the default RIPv2 behavior of automatic
summar​ization
(config) #
show ip protocols
show ip protocols
(config) #
network
network
 
ip-address
(config)
#
passiv​e-i​nte​‐
passiv​e-i​nte​‐
rface
rface
prevent the transm​ission of routing updates
through a router interface, but still allow that
network to be advertised to other routers.
(confi​g-r​outer)
#
ip route
ip route
0.0.0.0 0.0.0.0
0.0.0.0 0.0.0.0
propagate a default route in RIP
(confi​g-r​outer)
#
defaul​t-i​nfo​‐
defaul​t-i​nfo​‐
rmation
rmation
originate
originate
This instructs R1 to originate default inform​ation,
by propag​ating the static default route in RIP
updates.
NAT Config
NAT Config
(config)# ip nat inside source
static 192.16​8.11.99 209.16​‐
5.201.5
Configure the static transl​ation with
an inside local address of 192.16​‐
8.11.99 and an inside global address
of 209.16​5.2​01.5.
(config)# interface Serial​‐
0/0/0 , (confi​g-if)# ip nat
inside
Configure the proper 
inside NAT
interface.
R2(con​fig)# interface Serial​‐
0/1/0 , (confi​g-if)# ip nat
outside
Configure the proper outside NAT
interface.
 
(config)# ip nat pool
PUBLIC​-POOL 209.16​5.2​‐
00.241 209.16​5.2​00.250
netmask 255.25​5.2​55.224
Define a pool of public IPv4
addresses 209.16​5.2​00.241 to
209.16​5.2​00.250 with pool name
PUBLIC​-POOL.
R2(con​fig)# access​-list 2
permit 192.16​8.10.0
0.0.0.255
Configure ACL 2 to permit devices
from 192.16​8.1​0.0/24 network to be
translated by NAT.
 
NAT Config (cont)
NAT Config (cont)
R2(con​fig)# ip nat inside source list 2
pool PUBLIC​-POOL
Bind PUBLIC​-POOL with
ACL 2.
R2(con​fig)# interface Serial​0/0/0
R2(con​fig​-if)# ip nat inside
Configure the proper inside
NAT interface.
R2(con​fig)# interface Serial​0/1/0
R2(con​fig​-if)# ip nat outside
Configure the proper
outside NAT interface.
ip nat transl​ation timeout
clear ip nat transl​ation *
SysLog Config
SysLog Config
(config) # 
logging
logging
{address}
{address}
Configure the destin​ation hostname or
IPv4 address of the syslog.
(config) # 
logging trap
logging trap
{level}
{level}
Control the level of messages that will
be sent
(config)# logging source​-
in​terface {inter​face}
Logging Source
By 
Theleran
Theleran
cheatography.com/theleran/
 
Not published yet.
Last updated 27th February, 2019.
Page 4 of 4.
 
Sponsored by 
Readable.com
Readable.com
Measure your website readability!
https://readable.com
Command
Descri​ption
Command
Descri​ption
Top is Static Config
Bottom is Dynamic Config
Command
Descri​ption
