---
id: collect-261001-huawei/huawei/network-ptmngsys-web-tsrev-ar-en-content-ar-46-edesk-wifi-internet-slow-edesk-wi-f09d6975
title: "network-ptmngsys-web-tsrev-ar-en-content-ar-46-edesk-wifi-internet-slow-edesk-wi-f09d6975"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2011-04-10"]
keywords: ["voice"]
source: docs/RAG/collect-261001-huawei/network-ptmngsys-web-tsrev-ar-en-content-ar-46-edesk-wifi-internet-slow-edesk-wi-f09d6975.md
source_anchor: ""
source_lines: [1, 60]
sha256: 8bfc9c8dbc38821c6d47d57f5c5c40513e8dc9437eae1b6ca87323650f7caf9a
---

# network-ptmngsys-web-tsrev-ar-en-content-ar-46-edesk-wifi-internet-slow-edesk-wi-f09d6975

Ping the gateway from a STA. If the ping operation fails or a large number of packets are lost, locate the fault as follows:
In Wi-Fi scenarios, some STAs may scan the local network segment and sends ARP Request packets to all IP addresses on the entire LAN. If the ARP Request packets received by the device may exceed the CPU load threshold for triggering CPCAR, they are discarded by the device. However, the device discards the ARP Request packets randomly. Normal ARP Request packets from other STAs may be incorrectly discarded in this process, leading to a failure to learn ARP entries. The STAs will fail to ping the gateway.
Run the display cpu-defend statistics command for multiple times to observe the count of discarded ARP Request packets. That is, whether the Drop Packets value of arp-request increases in the command output. If so, the ARP Request packets are discarded.
<Huawei> display cpu-defend statistics
-----------------------------------------------------------------
Packet Type               Pass Packets        Drop Packets
-----------------------------------------------------------------
8021X                                0                   0
arp-miss                             5                   0
arp-reply                         8090                   0
arp-request                    1446576              127773
……
unknown-packet                   66146                   0
voice                                0                   0
vrrp                                 0                   0
-----------------------------------------------------------------
 If the ARP Request packets are continuously discarded, run the packet-type packet-type rate-limit rate-value command in the CPU attack defense policy view to set the CPCAR of arp-request to 512.
<Huawei> system-view
[Huawei] cpu-defend policy arp
[Huawei-cpu-defend-policy-arp] packet-type arp-request rate-limit 512
[Huawei-cpu-defend-policy-arp] quit
[Huawei] cpu-defend-policy arp 
After strict ARP learning is configured, the device learns only ARP entries for ARP Reply packets in response to ARP Request packets sent by itself, but does not learn the ARP entries for the ARP packets received from other devices. Upon an exception in ARP learning, if strict ARP learning is enabled, the device cannot learn ARP Request packets sent from STAs, leading to a failure to ping the gateway from the STAs.
In this case, run the display current-configuration | include arp command to view all ARP configurations. If arp learning strict or arp learning strict { force-enable | trust } is displayed in the command output, strict ARP learning is enabled. Disable this function.
To disable strict ARP learning, run the following command:
<Huawei> system-view
[Huawei] undo arp learning strict
If strict ARP learning is configured on an interface, disable this function in the interface view:
<Huawei> system-view
<Huawei> interface vlanif 100
[Huawei-Vlanif100] undo arp learning strict
An IP address conflict on the network will lead to frequent route flapping, greatly affecting user services. Run the display arp ip-conflict track command to check records of detected IP address conflicts.
<Huawei> display arp ip-conflict track
    Conflict type       : Remote IP Confilct
    IP address          : 100.1.1.1
    System time         : 2011-04-10 06:59:43
    Conflict count      : 1
    Suppress count      : 0
    Old interface       : GE1/0/1
    Receive interface   : GE1/0/2
    Old VLAN/CEVLAN     : 100/0
    Receive VLAN/CEVLAN : 100/0
    Old MAC             : 00e0-ca63-8141
    Receive MAC         : 00e0-ca63-8142
 If DHCP is used to allocate IP addresses on the network, run the display ip pool command to check the configured IP address pool to determine IP address conflicts in the address pool.
<Huawei> display ip pool name huawei conflict
  Pool-name        : huawei
  Pool-No          : 1
  Lease            : 1 Days 0 Hours 0 Minutes
 ……
Client-ID format as follows:
   DHCP  : mac-address                 PPPoE   : mac-address
   IPSec : user-id/portnumber/vrf      PPP     : interface index
   L2TP  : cpu-slot/session-id         SSL-VPN : user-id/session-id
  ----------------------------------------------------------------
  Index              IP             Client-ID    Type       Left   Status
  ----------------------------------------------------------------
   109        192.168.0.110           -           -          -    Conflict     
----------------------------------------------------------------
 You can locate the corresponding STAs by MAC address based on conflicted IP addresses. If IP addresses are statically configured, modify the IP address. If IP addresses are allocated using DHCP, shorten the DHCP lease, expand the address pool network segment, or configure automatic reclaim of conflicting IP addresses in the address pool.
