---
id: collect-260926-mikrotik/mikrotik/questions-1037334-dual-gateway-setup-in-mikrotik-0084002e
title: "questions-1037334-dual-gateway-setup-in-mikrotik-0084002e"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1037334-dual-gateway-setup-in-mikrotik-0084002e.md
source_anchor: ""
source_lines: [1, 46]
sha256: 250acbec0d0d4ed07a580a0a82a1982cbc80767267cc3012f5a3e550d6e07842
---

# questions-1037334-dual-gateway-setup-in-mikrotik-0084002e

Preliminary setup: You will need to set up the two WAN connections.  Assume IFC1 to be the first WAN port (e.g. ether1) interface and IFC2 to be the second WAN port (e.g. ether5).
/interface list member add interface=IFC1 list=WAN
/interface list member add interface=IFC2 list=WAN
/ip dhcp-client add interface=IFC1 default-route-distance=1
/ip dhcp-client add interface=IFC2 default-route-distance=1
Some of this may already be configured.  If using a different type of WAN connection such as PPPoE then adjust accordingly.
When the DHCP client connects, routes should be added, below is an example.
[admin@mikrotik] > /ip route print
Flags: X - disabled, A - active, D - dynamic,
C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme,
B - blackhole, U - unreachable, P - prohibit
 #      DST-ADDRESS        PREF-SRC        GATEWAY            DISTANCE
 0 ADS  0.0.0.0/0                          yyy.yy.yyy.1              1
 1 ADS  0.0.0.0/0                          xxx.xxx.xxx.1             1
 2 ADC  xxx.xxx.xxx.0/22   xxx.xxx.xxx.xxx IFC2                      0
...
If you don't see two 0.0.0.0/0 routes with the correct gateways, you can create them as static routes (/ip route add ...).  This is basic Mikrotik stuff beyond the scope of this answer.
Differentiating between the two WAN connections will be done using policy routing.  On each default WAN route set a routing mark.
/ip route set 0 routing-mark=unid2rm
/ip route set 1 routing-mark=unid3rm
The values unid2rm and unid3rm are arbitrary text strings.  This means the Mikrotik will send packets marked with a given routing mark through the given gateway IP.
Second, configure the firewall.  Look at the FORWARD chain.  The goal here is to exclude marked connections for the secondary (non-default) WAN from going through fasttrack.
/ip firewall filter add 8 chain=forward action=fasttrack-connection connection-state=established,related connection-mark=!unid2cm
/ip firewall filter add 9 chain=forward action=accept connection-state=established,related
Change the numbers 8 and 9 so that the rules are positioned at the start of your FORWARD chain.
Now in the PREROUTING chain, mark the incoming connections.  If you're only steering outbound connections, you can skip this.
/ip firewall mangle add 3 chain=prerouting action=mark-connection new-connection-mark=unid2cm passthrough=no connection-mark=no-mark in-interface=IFC1
/ip firewall mangle add 4 chain=prerouting action=mark-connection new-connection-mark=unid3cm passthrough=no connection-mark=no-mark in-interface=IFC2
The next step involves the MANGLE table and is where the magic happens for you.  Here you mark a connection based on the criteria you set.
/ip firewall mangle add 5 chain=prerouting action=mark-connection new-connection-mark=unid2cm passthrough=yes dst-address=221.35.12.5 connection-mark=no-mark in-interface-list=LAN
This is saying "when there is a connection to 221.35.12.5 which is not already marked, mark it with connection mark unid2cm".  You can set it based on source IP, MAC, etc. and add as many of these as needed.  Then the companion is:
/ip firewall mangle add 6 chain=prerouting action=mark-routing new-routing-mark=unid2rm passthrough=no connection-mark=unid2cm in-interface-list=LAN
/ip firewall mangle add 7 chain=output action=mark-routing new-routing-mark=unid2rm passthrough=no connection-mark=unid2cm
/ip firewall mangle add 8 chain=prerouting action=mark-routing new-routing-mark=unid3rm passthrough=no connection-mark=unid3cm in-interface-list=LAN
/ip firewall mangle add 9 chain=output action=mark-routing new-routing-mark=unid3rm passthrough=no connection-mark=unid3cm
Note the difference between unid2rm and unid2cm.  These rules take packets from a marked connection and give them a routing mark.  The routing mark is then used in Mikrotik's route table as mentioned earlier.
Every setup is a bit different, so you might need to play around to make it work for you.
Added: Mikrotik usually sets up NAT automatically on the WAN interface list.  The very first step at the top of the answer ensures that both interfaces are in the list.
[admin@mikrotik] > /ip firewall nat print
Flags: X - disabled, I - invalid, D - dynamic
 0    ;;; defconf: masquerade
      chain=srcnat action=masquerade out-interface-list=WAN log=no
      log-prefix=""
     
    
/ip route printand/ip firewall xxx printwherexxxisfilter,mangle,nat,raw.
