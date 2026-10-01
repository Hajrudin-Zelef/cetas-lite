---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-72652-pfsense-unifi-ap-hp-switch-and-vlans-691e3a37-1
title: "questions-72652-pfsense-unifi-ap-hp-switch-and-vlans-691e3a37"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-72652-pfsense-unifi-ap-hp-switch-and-vlans-691e3a37.md
source_anchor: ""
source_lines: [1, 98]
sha256: 5504fece46135009d1b9c1dfb15b3c01ad799e740fda3d85e5c84487913544d9
---

# questions-72652-pfsense-unifi-ap-hp-switch-and-vlans-691e3a37

I have a Unifi AP with a single MYNET SSID connected to my HP 2520 switch on port 23. I have a pfSense LAN interface connected on port 1. AP clients get DHCP (192.168.1.0/24) from pfSense and can access Internet, etc. I have static leases for all my "known" devices but also a DHCP pool so new devices can get on to get configured.
I have OpenVPN (client) configured on pfSense and policy route ports 443, etc over it. Some websites drop VPN traffic though, so thought I'd setup a "vpn bypass" SSID On the AP and give it a completely different subnet. I could tell pfSense to just route the entire network out the WAN interface instead.
I added vlan 200 on the AP and gave it NOVPN for an SSID. On pfSense, I followed some tutorials and setup vlan 200 on the LAN interface. I have the vlan interface enabled, added the firewall "pass" rule for the it. Then, setup DHCP for a completely different subnet (192.168.2.0/24). I don't have static leases configured for it yet, if it matters.
When a client leaves MYNET and connects to NOVPN, the client doesn't get an IP. tcpdump shows the DHCP request come in from the client, and pfSense replies, but that's all I see going on. On the client, the connection times out/fails. On pfSense, I see a DHCP reply going to the client's static lease for MYNET. The reply is originating from x.x.1.1 (pfSense LAN interface). I was expecting it to come from x.x.2.1 which is the vlan 200 address for pfSense.
11:44:16.625014 IP 192.168.1.1.67 > 192.168.1.66.68: BOOTP/DHCP, Reply, length 300
I think this might be happening because I don't have any VLANs on my switch maybe? I tried setting up vlan 200 on my switch but got lost with tagged/untagged/GVRP, etc. Do I need to trunk/tag the AP and pfSense ports somehow since they will be carrying multiple VLAN traffic? I tried a few things from this HP link but couldn't quite figure out what was needed to make this work. Any thoughts?
Current HP ports assignments:
===========================- TELNET - MANAGER MODE -============================
               Switch Configuration - VLAN - VLAN Port Assignment
  Port   DEFAULT_VLAN  no_vpn_vlan    |  Port   DEFAULT_VLAN  no_vpn_vlan
  ---- + ------------  ------------   |  ---- + ------------  ------------
  1    | No            Untagged       |  13   | No            Untagged
  2    | No            Untagged       |  14   | No            Untagged
  3    | No            Untagged       |  15   | No            Untagged
  4    | No            Untagged       |  16   | No            Untagged
  5    | No            Untagged       |  17   | No            Untagged
  6    | No            Untagged       |  18   | No            Untagged
  7    | No            Untagged       |  19   | No            Untagged
  8    | No            Untagged       |  20   | No            Untagged
  9    | No            Untagged       |  21   | No            Untagged
  10   | No            Untagged       |  22   | No            Untagged
  11   | No            Untagged       |  23   | No            Untagged
  12   | No            Untagged       |  24   | No            Untagged
 Actions->   Cancel     Edit     Save     Help
Port/Trunk settings (all ports are set the same as #1):
  Port    Type      Enabled      Mode      Flow Ctrl  Group  Type
  ----  --------- + -------  ------------  ---------  -----  -----
  1     1000T     | Yes      Auto          Disable
  ...
UPDATE: adding switch configs below:
running config
switch-2520G# show run
Running configuration:
; J9299A Configuration Editor; Created on release #J.14.54
hostname "switch-2520G" 
vlan 1 
   name "DEFAULT_VLAN" 
   no untagged 1-24 
   no ip address 
   exit 
vlan 200 
   name "no_vpn_vlan" 
   untagged 1-24 
   ip address 192.168.1.10 255.255.255.0 
   exit 
auto-tftp 192.168.100.120 "/tftp"
banner motd "HP SWITCH
"
include-credentials
password manager user-name "x..." sha1 "x..."
no telnet-server
ip authorized-managers 192.168.1.220 255.255.255.255 access manager
ip authorized-managers 172.22.200.220 255.255.255.0 access manager
ip authorized-managers 192.168.1.220 255.255.255.255 access manager
ip ssh public-key operator "ssh-rsa ..."
ip ssh public-key operator "ssh-rsa ..."
snmp-server community "x..." operator
snmpv3 engineid "xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx"
aaa authentication ssh login public-key
no tftp server
IP info
switch-2520G# show ip
 Internet (IP) Service
  Default Gateway :                
  Default TTL     : 64   
  Arp Age         : 20  
  Domain Suffix   :                               
  DNS server      :                                         
  VLAN                 | IP Config  IP Address      Subnet Mask     Proxy ARP
  -------------------- + ---------- --------------- --------------- ---------
  DEFAULT_VLAN         | Disabled 
  no_vpn_vlan          | Manual     192.168.1.1      255.255.255.0   No
 
VLAN info
switch-2520G# show vlan
 Status and Counters - VLAN Information
  Maximum VLANs to support : 256                  
  Primary VLAN : DEFAULT_VLAN
  Management VLAN :             
  VLAN ID Name                 Status       Voice Jumbo
  ------- -------------------- ------------ ----- -----
  1       DEFAULT_VLAN         Port-based   No    No   
  200     no_vpn_vlan          Port-based   No    No   
 
VLAN 1 (DEFAULT_VLAN)
switch-2520G# show vlan 1
 Status and Counters - VLAN Information - VLAN 1
  VLAN ID : 1      
  Name : DEFAULT_VLAN        
  Status : Port-based  Voice : No 
  Jumbo : No 
  Port Information Mode     Unknown VLAN Status    
  ---------------- -------- ------------ ----------
 
  Overridden Port VLAN configuration
  Port Mode        
  ---- ------------
 
