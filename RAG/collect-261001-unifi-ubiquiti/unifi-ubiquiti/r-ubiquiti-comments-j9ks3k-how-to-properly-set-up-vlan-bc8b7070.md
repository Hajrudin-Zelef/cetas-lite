---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-j9ks3k-how-to-properly-set-up-vlan-bc8b7070
title: "r-ubiquiti-comments-j9ks3k-how-to-properly-set-up-vlan-bc8b7070"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-j9ks3k-how-to-properly-set-up-vlan-bc8b7070.md
source_anchor: ""
source_lines: [1, 64]
sha256: 3154bdaaace11d6074bf2378a1fbd04a8aefc4037db0bf823848e0e06871c2da
---

# r-ubiquiti-comments-j9ks3k-how-to-properly-set-up-vlan-bc8b7070

How to properly set up VLAN? 
        
        
        
    
    
    I have a VM that needs to be completely isolated from LAN with the exception of one computer (actually the machine hosting it) and able to access WAN. I've tried everything I can think of. Assigned VLAN in UniFi controller by MAC, client refuses to work on the correct subnet (i.e. main LAN is 10.10.1.0 and I'm trying to put this one client on 10.10.2.0).
So client would be for example 10.10.2.20 and needs to be able to communicate with 10.10.1.100 but absolutely nothing else on the 10.10.1.0 subnet.
I've tried using firewall rules instead but I can't get that to work either. Using a UDM Pro.
What am I doing wrong and why am I so dumb?
Section des commentaires
Hello! Thanks for posting on r/Ubiquiti!
This subreddit is here to provide unofficial technical support to people who use or want to dive into the world of Ubiquiti products. If you haven’t already been descriptive in your post, please take the time to edit it and add as many useful details as you can.
Please read and understand the rules in the sidebar, as posts and comments that violate them will be removed. Please put all off topic and picture posts in the weekly off topic thread that is stickied to the top of the subreddit.
If you see people spreading misinformation, trying to mislead others, or other inappropriate behavior, please report it!
I am a bot, and this action was performed automatically. Please contact the moderators of this subreddit if you have any questions or concerns.
As long you don't use the new 6.x.x controller, you will be able to properly setup VLANs.
Still on 5.14.22
Is there a thread somewhere documenting it’s an issue? I’m currently struggling with VLANs
Just an example:
https://www.reddit.com/r/UNIFI/comments/itml87/beware_unifi_6x_vlan_issues/
https://www.reddit.com/r/Ubiquiti/comments/j2d3qr/version_6x_bug_list/
Commentaire supprimé par le membre
Sorry I have to upvote this simple comment because I understand the struggle. Take a +1
What’s your virtualization host running? Usually this is done by setting up a tagged vlan network in the host and assigning that network to the VM. You also need to make sure the tagged vlan is set up on the switch and gateway and assigned to the virtualization host.
VMWare Player 15. Windows 10 host and guest.
Second this ;
but also I have to mention that if your VM literally only ever needs to communicate with the host (no Internet) that it doesn't need full access to the network anyway. You probably didn't mean that ; but if you did, then the config also doesn't need to involve anything outside the host.
My bad on the rest. lol :)
u/salamithot does mention the VM needs WAN access AND needs to be isolated from anything else on the LAN. That means they need the VLAN and a host-only network.
Can't you just set the network to host-only on your hypervisor?
Does this allow the VM to access the internet?
No, only to the host machine
Happy cake day!!
My post is old, but it still applies.
https://www.reddit.com/r/Ubiquiti/comments/85kbro/usg_firewall_rules_for_iot_and_security_cameras/dvy2h29/
Broken down for your scenario
For your situation I'll draw up this network setup.
LAN: 10.10.1.0/24 corporate (keep this for management), you can tag it or keep it untagged to a VLAN and use this for management.
Main: 10.10.2.0/24 corporate (this will be your home/private network) tag this VLAN 2 (or whatever number you want)
VM: 10.10.3.0/24 corporate (VM network), tag this VLAN 3 (or whatever number you want)
LAN_IN: (Note, you can use any name for your rules, these are just examples of the names in this scenario. All other settings listed need to be enabled accordingly).
2000 : name=Allow Main to VLANs : Before Predefined Rules : Allow : All protocols : Established and Related States checked : Source and Destination left empty
2001 : name=Block VM : Before Predefined Rules : Drop : All protocols : Source Network - VM: Destination left empty
WAN_IN:
2000 : name=WANIN Block VM: Before Predefined Rules : Drop : All protocols : Source Network - VM : Destination left empty
(NOTE: this is not necessary as long as you don't have port forwarding enabled for anything to the VMnetwork, as the LAN_IN rule is sufficient enough)
WAN_OUT:
2000 : name=WANOUT Block VM: Before Predefined Rules : Drop : All protocols : Source Network - VM: Destination left empty
(NOTE: this is not necessary as long as you don't have port forwarding enabled for anything to the VM network, as the LAN_IN rule is sufficient enough)
For the subnets choose whatever you want. 192.168...., 10.0..., 172.1...
Then, you can create an exception for the whatever machine you need to access the VM network IP address.
I've already done some of this but I'll check if there's anything missing.
How is your server connected to the UDM pro. Are there any switches in between?
One Mikrotik CRS305-1G-4S+IN between. Using SFP.
Have you set the VLAN on the Port, the Host is connected to?
And have you set the VLAN on the virtual NIC on the Host?
Wait so if I have the 10.0.2.0/24 subnet as VLAN 437 I can just set the virtual NIC in VMWare to use VLAN 437?
Assigning a client a VLAN tag via the controller doesn't actually work. The closest thing would be LLDP-MED, which detects a device type (voip) and assigns it to a predefined VLAN.
For your use case, I recommend setting up a trunk port, with only the VLANs you need to pass to the host, the default behavior of a Unifi switch is to create a trunk with all VLANs. Based on your description, this trunk should contain at least two VLANs, a management VLAN and guest VLAN. The management VLAN will give the host an IP address on a subnet accessible to the device you will be using to configure and manage it. While the guest VLAN will be used by the host and the client to communicate with each other and allow the client access to the internet. You will need to make a virtual switch on the host that has the guest VLAN ID. Creating a virtual switch should automatically create a virtual interface on your host that is assigned to that virtual switch, or at least this is the expected behavior on VMware and Hyper-V. On the guest, you will need to assign it's virtual interface to the same virtual switch. Lastly, you will need to create a firewall rule on the LAN IN side, that drops all packets from the guest VLAN (identified by it's subnet) going to any network. As it's on the LAN IN side, this firewall rule will apply when any other subnet on your network receives packets from the guest VLAN. but it will not stop the guest and host from communicating with the UDM and it will allow those packets to be routed through the WAN.
Screenshots of your unifi settings would be a great start
Also what is the layout of your network in regards to how the devices are plugged (where is the VM sitting and your other systems). Put some labels in there so we know what VLANs your devices are supposed to be sitting in
Port isolation would work, or not?
First, use a subnet calculator to make sure you are configuring the proper network and IP assignments. Then make sure you configure your VLAN properly. Lastly, you have to route to get on/off the VLAN, so make sure you have a gateway configured on the VLAN to do that.
