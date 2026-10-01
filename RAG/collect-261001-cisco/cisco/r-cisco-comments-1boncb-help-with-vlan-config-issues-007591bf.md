---
id: collect-261001-cisco/cisco/r-cisco-comments-1boncb-help-with-vlan-config-issues-007591bf
title: "r-cisco-comments-1boncb-help-with-vlan-config-issues-007591bf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-1boncb-help-with-vlan-config-issues-007591bf.md
source_anchor: ""
source_lines: [1, 64]
sha256: 90b0c08a2a84757badb0839fe04c43190312a43c8bca487a3d72135ee8295ff2
---

# r-cisco-comments-1boncb-help-with-vlan-config-issues-007591bf

Help with VLan config issues 
        
    [** SOLVED **]
Hello r/cisco! I am trying to configure a network in a lab environment and need a bit of help. I've been through my book, boson netsim, and google and am still stuck on this problem.
I have a Catalyst 2950 switch that I am trying to create 2 VLans on. The switch connects via port f0/9 to an 1841 series router. To keep it simple lets just use a class B address /24.
There is more but this is the problem area. On Vlan 1 i have clients(win8) and new connections, on Vlan 2 I have a database server and a backup server (all running Windows Server 2k12). For simplicity I have disabled all firewalls.
Here is how i configured the 2950 At this point i just want the computer to ping each-other
(config)#hostname InternalSwitch
(config)#enable secret grouponeadmin
(config)#line vty 0 4
(config-line)#password grouponeadmin
(config-line)#login
(config)#no logging console
(config)#vlan database
(vlan)#vtp transparent
(config)#int vlan 1
(config-if)#ip address 100.10.2.2 255.255.255.0
(config)ip default-gateway 100.10.2.1
(config)#vlan 2
(config-vlan)#name Database
(config)#int vlan 2
(config-if)#ip address 100.10.1.2 255.255.255.0
(config)#int f0/1
(config-if)#switchport access vlan 2
(config-if)#spanning-tree portfast
(config)#int f0/2
(config-if)#switchport access vlan 2
(config-if)#spanning-tree portfast
(config)#int f0/9
(config-if)#switchport mode trunk
(config-if)#switchport trunk allowed vlan all
(config-if)#channel-group 1 mode on
(config)#spanning-tree vlan 1
(config)#spanning-tree vlan 2
(config)# write memory
So at this point I have my 2 vlans, I have assigned them on 2 subnets and given them a port to jump at. The computers on my database vlan can ping each-other at this point.
Next i configured the router port f0/1.1 100.10.2.1 and f0/1.2 100.10.1.1
All of the tutorials and help have used non 802.1q, but with my hardware i can only use 802.1q for trunking. There HAS to be a command I'm missing here. Sorry if this is no way near enough info to identify a problem. As stated, I am a network noobie :(
edit: formatting
edit: SOLVED
Each port on the switch needed to have have "switchport mode access" if they were in a vlan
The router needed an entry for Vlan 2 in its local database.
TY for all your help guys!
Section des commentaires
Check your config on the router. The 2950 will not do any inter-vlan routing so you need to depend on the router for this. The router config should look something like this:
Also verify that your VLANs actually exist (show vlan brief) and not just the interfaces. A 2950 can only have one active SVI, so you may as well nuke the IP address on one of the VLANs as well (no int vlan2). Run "show int trunk" and "show spanning-tree vlan [1|2]" to verify that the trunks are up and the VLANs are actually forwarding.
assuming you meant f0/1.1 in the first line
Yes, sorry.
2950's are L2 only. They can only have one Vlan interface active at any given time. What is the port channel for? You only have one port connecting to the router, correct?
On the switch port connecting to the router: swit trun encap dot1q
One the router:
f0/1.1
encap dot1 1
f0/1.2
encap dot1 2
Don't forget:
I have forgotten that many times and been baffled.
You have channel-group 1 mode on your trunk port, but i only see one trunk port. is there an ether-channel? If not you don't need this.
Also you can only have one active management VLAN on the 2950 at a time. VLAN 2 does not need a int vlan defined. This is not causing the issue of the computers not seeing each other, just something you need to know.
Are you sure of the cabling between the switch and the router? Can you ping the switch from the router?
Do the workstations and the server have the correct default gateway set for the vlan they are on? The clients should have a default gateway of 100.10.2.2 and the servers should be set to 100.10.1.2.
the gateways are secure. I will test the communication between the switch and router when i get back to the lab. the config is a mis match of different things i tried to get this to work. I guess what is bothering me the most is what do i have to configure for the trunk line. all the tutorials say that 1841 series routers auto-configure trunk's but since it didn't work I have been desperate to hit the switch that opens the gates.
Yeap im getting a layer 3 issue. that seems the be the root of all this. still trying to figure out why i can't ping back and forth. i removed Vlan 2's IP address making Vlan 1's address the one the switch identifies itself by (100.10.2.2) and on the Router placed vlan 1 in the f0/1.1 sub interface (100.10.2.1). still can't get them to talk. trying different physical ports now to make sure it isn't physical.
You've got to remember the 2950 is a L2 only switch. The IP addresses on that switch are irrelevant to anything you're trying to do. They are only there for managing the switch and only one of them can be active at a time.
