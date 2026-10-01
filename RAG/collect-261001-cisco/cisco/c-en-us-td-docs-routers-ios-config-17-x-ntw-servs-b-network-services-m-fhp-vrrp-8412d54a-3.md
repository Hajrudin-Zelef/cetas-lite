---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp-8412d54a-3
title: "c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "preemption"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a.md
source_anchor: ""
source_lines: [203, 428]
sha256: 79774ce952a4fb162a7a2dfae85494b50ad8c36d6e8606848adf5c91ebd8746b
---

# c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a

VRRP, it is a good idea to do so before enabling VRRP.
Procedure
Step 1
enable
Example:
Router> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Router# configure terminal
Enters global configuration mode.
Step 3
interfacetypenumber
Example:
Router(config)#GigabitEthernet 0/0/0
Enters interface configuration mode.
Step 4
ipaddressip-addressmask
Example:
Router(config-if)# ip address 172.16.6.5 255.255.255.0
Configures the router to take over as virtual primary router for a VRRP group if it has a higher priority than the current
virtual primary router.
The default delay period is 0 seconds.
The router that is IP address owner will preempt, regardless of the setting of this command.
Step 8
vrrpgrouptimerslearn
Example:
Router(config-if)# vrrp 10 timers learn
Configures the router, when it is acting as virtual router backup for a VRRP group, to learn the advertisement interval used
by the virtual primary router.
Step 9
exit
Example:
Router(config-if)# exit
Exits interface configuration mode.
Step 10
novrrpsso
Example:
Router(config)# no vrrp sso
(Optional) Disables VRRP support of SSO.
VRRP support of SSO is enabled by default.
EnablingVerifying VRRP
Procedure
Step 1
enable
Example:
Router> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Router# configure terminal
Enters global configuration mode.
Step 3
interfacetypenumber
Example:
Router(config)# interfaceGigabitEthernet 0/0/0
Enters interface configuration mode.
Step 4
ipaddressip-addressmask
Example:
Router(config-if)# ip address 172.16.6.5 255.255.255.0
Configures an IP address for an interface.
Step 5
vrrpgroupipip-address [secondary]
Example:
Router(config-if)# vrrp 10 ip 172.16.6.1
Enables VRRP on an interface.
After you identify a primary IP address, you can use the vrrpip command again with the secondary keyword to indicate additional IP addresses supported by this group.
Note
All routers in the VRRP group must be configured with the same primary address and a matching list of secondary addresses
for the virtual router. If different primary or secondary addresses are configured, the routers in the VRRP group will not
communicate with each other and any misconfigured router will change its state to primary.
Step 6
showvrrp [briefall] | interface]
Example:
Router(config-if)#show vrrp brief
Interface Grp Pri Time Own Pre State Master addr Group addr
BD10 1 100 9609 Y Backup 10.1.0.2 10.1.0.10
BD10 5 200 90218 Y Master 10.1.0.1 10.1.0.50
BD10 100 100 3609 Backup 10.1.0.2 10.1.0.100
(Optional) Displays a brief or detailed status of one or all VRRP groups on the router.
Step 7
showvrrpinterfacetypenumber [brief]
Example:
Router(config)# interfaceGigabitEthernet 0/0/0
Router)config-if)#show vrrp interface bdi10
BDI10 - Group 10
G1
State is Master
Virtual IP address is 10.0.0.5
Virtual MAC address is 0000.5e00.010a
Advertisement interval is 10.000 sec
Preemption enabled, delay min 380 secs
Priority is 110
Master Router is 10.0.0.2 (local), priority is 110
Master Advertisement interval is 10.000 sec
Master Down interval is 30.570 sec
FLAGS: 1/1
(Optional) Displays the VRRP groups and their status on a specified interface.
Step 8
end
Example:
Router(config-if)# end
Returns to privileged EXEC mode.
Configuring VRRP Object Tracking
Note
If a VRRP group is the IP address owner, its priority is fixed at 255 and cannot be reduced through object tracking.
Router(config)# track 2 interface serial 6 line-protocol
Configures an interface to be tracked where changes in the state of the interface affect the priority of a VRRP group.
This command configures the interface and corresponding object number to be used with the
vrrptrack command.
The
line-protocol keyword tracks whether the interface is up. The
iprouting keyword also checks that IP routing is enabled and active on the interface.
You can also use the
trackiproute command to track the reachability of an IP route or a metric type object.
Step 4
interfacetypenumber
Example:
Router(config)# interface Ethernet 2
Enters interface configuration mode.
Step 5
vrrpgroupipip-address
Example:
Router(config-if)# vrrp 1 ip 10.0.1.20
Enables VRRP on an interface and identifies the IP address of the virtual router.
Step 6
vrrpgroupprioritylevel
Example:
Router(config-if)# vrrp 1 priority 120
Sets the priority level of the router within a VRRP group.
Step 7
vrrpgrouptrackobject-number [decrementpriority]
Example:
Router(config-if)# vrrp 1 track 2 decrement 15
Configures VRRP to track an object.
Step 8
end
Example:
Router(config-if)# end
Returns to privileged EXEC mode.
Step 9
showtrack [object-number]
Example:
Router# show track 1
Displays tracking information.
Configuring VRRP Text Authentication
Before you begin
Interoperability with vendors that may have implemented the RFC 2338 method is not enabled.
Text authentication cannot be combined with MD5 authentication for a VRRP group at any one time. When MD5 authentication
is configured, the text authentication field in VRRP hello messages is set to all zeros on transmit and ignored on receipt,
provided the receiving router also has MD5 authentication enabled.
Configures an interface type and enters interface configuration mode.
Step 4
ipaddressip-addressmask [secondary]
Example:
Router(config-if)# ip address 10.0.0.1 255.255.255.0
Specifies a primary or secondary IP address for an interface.
Step 5
vrrpgroupauthenticationtexttext-string
Example:
Router(config-if)# vrrp 1 authentication text textstring1
Authenticates VRRP packets received from other routers in the group.
If you configure authentication, all routers within the VRRP group must use the same authentication string.
The default string is cisco.
Note
All routers within the VRRP group must be configured with the same authentication string. If the same authentication string
is not configured, the routers in the VRRP group will not communicate with each other and any misconfigured router will change
its state to primary.
Step 6
vrrpgroupipip-address
Example:
Router(config-if)# vrrp 1 ip 10.0.1.20
Enables VRRP on an interface and identifies the IP address of the virtual router.
Step 7
Repeat Steps 1 through 6 on each router that will communicate.
—
Step 8
end
Example:
Router(config-if)# end
Returns to privileged EXEC mode.
Configuration Examples for VRRPv2
Example: Configuring VRRP
In the following example, Router A and Router B each belong to three VRRP groups.
In the configuration, each group has the following properties:
Group 1:
Virtual IP address is 10.1.0.10.
Router A will become the primary for this group with priority 120.
Advertising interval is 3 seconds.
Preemption is enabled.
Group 5:
Router B will become the primary for this group with priority 200.
Advertising interval is 30 seconds.
Preemption is enabled.
Group 100:
Router A will become the primary for this group first because it has a higher IP address (10.1.0.2).
In the following example, the tracking process is configured to track the state of the line protocol on serial interface
0/1. VRRP on Ethernet interface 1/0 then registers with the tracking process to be informed of any changes to the line protocol
state of serial interface 0/1. If the line protocol state on serial interface 0/1 goes down, then the priority of the VRRP
group is reduced by 15.
Router# show vrrp
Ethernet1/0 - Group 1
State is Master
Virtual IP address is 10.0.0.3
Virtual MAC address is 0000.5e00.0101
Advertisement interval is 1.000 sec
Preemption is enabled
min delay is 0.000 sec
Priority is 105
Track object 1 state Down decrement 15
Master Router is 10.0.0.2 (local), priority is 105
Master Advertisement interval is 1.000 sec
Master Down interval is 3.531 sec
Router# show track
Track 1
Interface Serial0/1 line-protocol
Line protocol is Down (hw down)
1 change, last change 00:06:53
Tracked by:
VRRP Ethernet1/0 1
Example: VRRP Text Authentication
