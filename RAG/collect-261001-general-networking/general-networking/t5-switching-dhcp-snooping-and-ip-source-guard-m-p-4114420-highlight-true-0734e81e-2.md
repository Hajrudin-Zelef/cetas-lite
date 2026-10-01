---
id: collect-261001-general-networking/general-networking/t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e-2
title: "t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e.md
source_anchor: ""
source_lines: [150, 247]
sha256: d179b95ee92613882a38527da4bf22a0af091704a2e5e495e18f66f322bd2813
---

# t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e

			LAN Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-07-2020 12:46 AM
DHCP snooping needs to trust ports that will receive ingress DHCP server traffic, otherwise snooping will drop those packets, and the binding will never be learned.
Any port that will carry traffic from a valid DHCP server (the trunk port used for your uplink) should be trusted for DHCP snooping using:
interface <interface>
 ip dhcp snooping trust
!
If you are running DHCP snooping on the switch that the DHCP server is connected to, you will need to trust that access port as well. In this case, I'm wondering if the DHCP servers in question are on downstream switches. If that's the case, the core switch would need trust statements since you're running DHCP relay. If they are on access ports those ports would need to be trusted as well.
In terms of troubleshooting ...
You can verify all bindings on a switch using the operation command:
show ip source binding
Alternatively you could use show ip dhcp snooping binding, but this would exclude any bindings from other sources, such as static bindings configured manually.
When IP source guard is enabled on a port, ingress IP packets with a source address not in the binding table will be dropped. This is likely the behavior you're experiencing, and the reason would be because the association of MAC, IP, VLAN and switchport is not yet in the binding table. That is the primary function of this feature.
The combination with DHCP snooping with IP source guard or dynamic ARP inspection can create problems for switch reboots, with bindings being lost. This is because not all devices retry DHCP on link up-down events. To mitigate this you should configure a backup URL for the DHCP snooping binding database which can be loaded on boot so the switch knows existing leases.
Typically, this is done using a TFTP server that allows write access:
ip dhcp snooping database tftp://1.2.3.4/dhcp-snooping/hostname
If you have a host that does not perform DHCP you can manually configure a static binding using the following:
ip source binding <MAC> vlan <VLAN> interface <interface>
If you are troubleshooting and would like to manually add an entry to the DHCP snooping binding table, you can do so from operational mode with the following:
ip dhcp snooping binding <MAC> vlan <VLAN> interface <interface> expiry <seconds>
This will add an entry which will expire after N seconds (typically whatever your default lease time is).
You can also manually delete a snooping binding with:
clear ip dhcp snooping binding <IP>
For IP source guard, you can verify the operational status using:
show ip verify source
This will either show an IP address if it has been learned and is permitted, deny-all if no matching entries are known, or inactive if DHCP snooping is no enabled on the VLAN.
Once you get DHCP snooping and IP source guard enabled, I strongly recommend enabling DAI or dynamic ARP inspection as well. IP source guard will prevent IP packets but not filter ARP, so DAI is a similar feature specific to ARP.
To enable DAI you would first add trust statements to all your trunk links between switches which would handle source traffic from a router:
ip arp inspection trust
Then enable globally per-VLAN similar to DHCP snooping:
ip arp inspection vlan 100-200
And finally, there is an access port configuration you should apply to increase the limit of ARP packets per second permitted, since the default of 15 is too low and will result in frequent err-disabled ports. The following generally works well for most environments:
interface <interface>
 ip arp inspection limit rate 150 burst interval 5
!
You can view the operational state of DAI with:
show ip arp inspection 
Finally, a few notes on using DHCP snooping, IP source guard, and DAI at scale:
For larger L2 deployments you should also trust downstream switch links for both DHCP snooping and DAI to avoid the additional CPU of those entries being learned. In our template we apply the trust statement to all links between switches regardless of upstream or downstream for this reason.
If you run a large L2 environment (e.g. over 1,000 hosts flowing through a switch) on IOS XE 16.x you might also find that CPU is running high because of a service called IPDT or IP device tracking.
In IOS-XE the implementation of DHCP snooping was reworked as part of IPDT, so enabling this feature will also enable IPDT on the switch. IPDT is generally tuned for a L3 access model and does not scale well without additional configuration to ignore links between switches. This is more of a problem on the 3650 and 3850 series switches than the 9300 series due to significantly better CPU on the 9300, but is still good to apply as a safe guard.
Globally, create a custom IPDT policy:
device-tracking policy DT-NOTRACK
 trusted-port
 security level glean
 device-role switch
 no protocol ndp
 no protocol dhcp6
 no protocol arp
 no protocol dhcp4
 no protocol udp
 tracking disable
!
device-tracking binding reachable-lifetime 300 stale-lifetime 30 down-lifetime 30
!
Note that some of this configuration is no longer needed, or is redundant to work around bugs, depending on the specific release, but this template should work across most IOS XE releases regardless of train.
Then on each trunk interface between switches, where you have already trusted DHCP snooping and ARP inspection, apply the DT-NOTRACK policy you just defined:
interface <interface>
 device-tracking attach-policy DT-NOTRACK
!
If you don't have any access ports connected to your "CS" core switch, I would recommend against running DHCP snooping there and keeping this as an access switch only feature set.
I'd also recommend using descriptions on ports, and using a spanning-tree priority other than 0. The default is 32,772 so if you want something safely below that, 28,672 would be fine for your access layer, and 16,384 for your core would leave you enough room for a future higher-priority bridge if needed. If you use the default "root primary" it will attempt to use 24,576, and "root secondary" will be 28,672, so leaving access switches at 32,772 and instead protecting access ports with BPDUguard would be a more common approach that is less likely to cause confusion.
Example access port:
interface <interface>
 description #
 switchport mode access
 switchport access vlan 100
 switchport nonegotiate
 switchport block unicast
 switchport port-security maximum 8
 switchport port-security aging time 5
 switchport port-security aging type inactivity
 switchport port-security violation shutdown
 switchport port-security
 ip arp inspection limit rate 150 burst interval 5
 no mdix auto
 ip verify source
 spanning-tree portfast
 spanning-tree bpduguard enable
!
Example trunk port:
interface <interface>
 switchport mode trunk
 switchport nonegoitate
 device-tracking attach-policy DT-NOTRACK
 switchport trunk native vlan 999
 switchport trunk allowed vlan 100-200
 ip arp inspection trust
 ip dhcp snooping trust
!
 
