---
id: collect-261001-cisco/cisco/t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04-2
title: "t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04.md
source_anchor: ""
source_lines: [49, 315]
sha256: 024efaf53adea638e0ac29c985fd55d22437c300033f314644c19422efbdd648
---

# t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04

			Catalyst 3000
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-30-2024 06:06 AM
Hi,
It all looks good, meaning you have ports Gi0/1, Gi0/2 in VLAN10 and Gi0/3, Gi0/4 in VLAN 20 (except that on STP output for VLAN 20, port Gi0/4 does not show up; did you not paste complete output or have disabled STP on that port via BPDFilter?); assuming hosts connected to these ports have IP addresses from the correct subnet, you should be able to have connectivity between PC's and default gateway which is the switch, while based on routing also between hosts in different VLAN's assuming you'v set the correct gateway IP address (the switch) on the hosts. Not sure what OS are the hosts running, however, try to ping the switch from hosts (the other way around, from switch to hosts may not work as maybe hosts have firewall turned on which filters ICMP packets). Based on the ARP and MAC table, you should at leat be able to ping from host 192.168.10.2 to switch which is 192.168.10.1.
ARP entries on the switch will show up only if there is IP communication between switch and host (if you ping the switch from all hosts, the switch should have all ARP entries); MAC entries will show up on switch only if the hosts sends any kind of traffic for switch to learn the MAC address.
Best,
Cristian.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 01:12 PM
In order to route between vlan, you need the command "ip routing", that's it. AS your switch have ipbase license, it is able to accept the command and do the needfull
Factory reset is not reason to stopping working. Can you share the show running-config?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 02:14 PM
Thank you for the response, greatly appreciated. I did run the "ip routing" command. Please see the show running-config output and please see if there are any mistakes in the cisco switch setup.
Building Configuration
Current Configuration
version 15.2
no service pad
no service password-encryption
!
hostname xxxxxx
!
boot-start-marker
boot-end-marker
!
!
no aaa new-model
system mtu routing 1500
!
!
!
!
!
ip routing
!
!
!
!
!
!
spanning-tree mode rapid-pvst
spanning-tree extend system-id
!
!
!
vlan internal allocation policy ascending
!
!
!
!
!
interface GigabitEthernet0/1
switchport access vlan 10
switchport mode access
!
interface GigabitEthernet0/2
switchport access vlan 10
switchport mode access
!
interface GigabitEthernet0/3
switchport access vlan 20
switchport mode access
!
interface GigabitEthernet0/4
switchport access vlan 20
switchport mode access
!
interface GigabitEthernet0/5
!
interface GigabitEthernet0/6
!
interface GigabitEthernet0/7
!
interface GigabitEthernet0/8
!
interface Vlan1
no ip address
!
interface Vlan10
ip address 192.168.10.1 255.255.255.0
!
interface Vlan20
ip address 192.168.20.1 255.255.255.0
!
ip forward-protocol nd
!
ip http server
ip http secure server
!
line con 0
line vty 5 15
!
!
end
-----------
Note about the host Linux PCs am testing with :
The hosts I am testing them with have some bridges setup for QEMU and maybe expecting tags - so if the switch configuration looks good, I can switch start looking for issues on the connected hosts. so far, all I have done on the hosts are these commands
sudo ip addr add 192.168.10.2/255.255.255.0 dev enp0s25 (a static ip in the 10 subnet)
sudo ip route add 0.0.0.0 via 192.168.10.1 --> which is Vlan 10's IP on the switch
Please let me know, thank you so much for the help!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 02:19 PM - edited 10-28-2024 02:40 PM
There is nothing need except
No shut down
Needed under vlan SVI
That all
MHM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 06:46 AM
thanks @MHM Cisco World can you please confirm if this is what you mean ?
#Assign Ports
interface GigabitEthernet0/<x>
no shutdown ?????
switchport mode access
switchport access vlan <vlan id>
or is this what you mean ?
#enable VLAN to VLAN communications
ip routing
no shutdown ?????
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 07:09 AM
I see this before
show ip interface brief
Vlan1                        unassigned         Yes   unset      up                     down
Vlan10                      192.168.10.1      Yes   unset      up                     down
Vlan20                      192.168.20.1     Yes   unset      up                     down
then you now share below
show ip interface brief
Vlan1 unassigned Yes unset up down
Vlan10 192.168.10.1 YES manual up up
Vlan20 192.168.20.1 YES manual up up
So the VLAN SVI are both UP not problem
only check if you connect correct PC to correct VLAN, i.e. PC with IP in subnet 192.168.10.x must connect to port assign to vlan 10 and PC with IP in subnet 192.168.20.x must connect to port assign to vlan20
MHM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 02:31 PM - edited 10-28-2024 02:37 PM
Send the commnad of "show ip int br" please
did you create the vlan on the switch with the command
conf t
vlan 10
exit
vlan 20
?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 06:43 AM
@Flavio Miranda here is the output - my apologies for the delay, I do not have remote access to this particular environment makes it hard to respond after evening.
show ip interface brief
Vlan1 unassigned Yes unset up down
Vlan10 192.168.10.1 YES manual up up
Vlan20 192.168.20.1 YES manual up up
GigabitEthernet0/1 unassigned YES unset up up
GigabitEthernet0/2 unassigned YES unset up up
GigabitEthernet0/3 unassigned YES unset up up
GigabitEthernet0/4 unassigned YES unset up up
GigabitEthernet0/5 unassigned YES unset down down
GigabitEthernet0/6 unassigned YES unset down down
GigabitEthernet0/7 unassigned YES unset down down
GigabitEthernet0/8 unassigned YES unset down down
GigabitEthernet0/9 unassigned YES unset down down
GigabitEthernet0/10 unassigned YES unset down down
For your question#2 this is the notes and sequence I used to create them
Config 1/4
----------
enable
config termninal
hostname <name>
show vlan
vlan <vlan id>
#Assign Ports
interface GigabitEthernet0/<x>
switchport mode access
switchport access vlan <vlan id>
#Assign IPs
interface vlan <vlan id>
ip address <match the subnet ids, start with 1, not zero> <subnet mask 255.255.255.0>
#to check everything
ip interface brief
#enable VLAN to VLAN communications
ip routing
Please let me know if you see something not right
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 07:40 AM
@bob0198labops dont worry
Something has change since your first post. The vlan were down
show ip interface brief
Vlan1 unassigned Yes unset up down
Vlan10 192.168.10.1 Yes unset up down
Vlan20 192.168.20.1 Yes unset up down
They are up now
show ip interface brief
Vlan1 unassigned Yes unset up down
Vlan10 192.168.10.1 YES manual up up
Vlan20 192.168.20.1 YES manual up up
Still can not communicate?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 08:41 AM
Yes, I just tried with 3 devices connected none of them are able to ping each other. I tried pinging the Vlan Ips, the other device ips, all come back with "Destination Host Unreachable"
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-29-2024 09:25 AM
