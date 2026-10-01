---
id: collect-261001-cisco/cisco/r-cisco-comments-8tzzw8-having-a-hard-time-with-port-channeling-0c35de0f
title: "r-cisco-comments-8tzzw8-having-a-hard-time-with-port-channeling-0c35de0f"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2018-26-06"]
keywords: []
source: docs/RAG/collect-261001-cisco/r-cisco-comments-8tzzw8-having-a-hard-time-with-port-channeling-0c35de0f.md
source_anchor: ""
source_lines: [1, 115]
sha256: 7fe39169f0df0498b2c7e98521118af756d5ea04f24679ddba370a7ae2a47fb7
---

# r-cisco-comments-8tzzw8-having-a-hard-time-with-port-channeling-0c35de0f

Having a hard time with port channeling
UPDATE: solution down below
      Original thread
I have this complex setup that I'm not use to.  I'm trying to etherchannel, but I'm working with a chassis that I'm unfamiliar with.  I thought I had it setup correctly from just my knowledge and experience alone, but LACP won't negotiate in Windows Server.  Here is my basic network diagram I'm working with: https://imgur.com/a/Q1Orv3N
    
This is a broken down one that is prob easier to follow on what I'm trying to achieve: https://imgur.com/a/WPlbcAO
I think the problem isnt so much the ether channel on the cisco side or even the VRTX External NIC side. I think I may have to port channel the internal NICS too maybe?
Here is the Cisco 3750-x config
interface Port-channel1
switchport trunk encapsulation dot1q
switchport trunk native vlan 101
switchport trunk allowed vlan 10
switchport mode trunk
!
interface GigabitEthernet2/0/1
description VRTX CMC Management
switchport access vlan 101
switchport trunk native vlan 101
switchport mode access
spanning-tree portfast
!
interface GigabitEthernet2/0/2
description VRTX Gi0/1
switchport trunk encapsulation dot1q
switchport trunk native vlan 101
switchport trunk allowed vlan 10
switchport mode trunk
channel-protocol lacp
channel-group 1 mode active
!
interface GigabitEthernet2/0/3
description VRTX Gi0/2
switchport trunk encapsulation dot1q
switchport trunk native vlan 101
switchport trunk allowed vlan 10
switchport mode trunk
channel-protocol lacp
channel-group 1 mode active
Config for Internal Dell Switch
interface gigabitethernet0/1
channel-group 1 mode auto
switchport mode trunk
switchport trunk native vlan none
switchport trunk allowed vlan remove 1-9,11-4094
!
 interface gigabitethernet0/2
 channel-group 1 mode auto
 switchport mode trunk
 switchport trunk native vlan none
 switchport trunk allowed vlan remove 1-9,11-4094
!
 interface tengigabitethernet1/1
 switchport access vlan 10
 !
 interface tengigabitethernet1/2
 switchport access vlan 10
 !
 interface tengigabitethernet1/3
 switchport access vlan 10
 !
interface tengigabitethernet1/4
switchport access vlan 10
 !
 interface tengigabitethernet2/1
switchport access vlan 10
 !
 interface tengigabitethernet2/2
switchport access vlan 10
 !
 interface tengigabitethernet2/3
switchport access vlan 10
 !
 interface tengigabitethernet2/4
switchport access vlan 10
 !
 interface Port-channel1
switchport mode trunk
switchport trunk native vlan 10
switchport trunk allowed vlan remove 1-9,11-4094
 !
 interface oob
 ip address 10.10.1.101 255.255.255.0
 no ip address dhcp
      UPDATE (6/26/2018):
UPDATE (6/26/2018):
So I decided to do the lag group on the internal interfaces as well. I first ended up getting my one blade working by creating a seperate LAG group and assigned it to the 0/1-4 interfaces.  Then my LACP team in windows came right up. I assigned this same LAG group with ID2 to my 2nd blade, but it didnt work.  I then created a 3rd LAG group and assigned it to 1/1-4  and then that blades LACP team came up.  BOOM!
    
      Proof
Here is the dell GUI side of it: https://imgur.com/a/94lHKHb
Here is the Dell R1-2210 config of it all:  https://imgur.com/a/Kl7hynW
Here it is in Windows Server:  https://imgur.com/a/K07m5Fu
    
Section des commentaires
Your Cisco side looks good. What's the output on the vrtx side for "show lacp port-channel 1"?
can I also see a "show etherchannel summary" on the cisco side?
Dell Switch (VRTX Side)
console#show lacp port-channel 1 Port-Channel Po1 Port Type Unknown Attached Lag id: Actor System Priority:1 MAC Address: e4:f0:04:a7:c7:c9 Admin Key: 1000 Oper Key: 1000 Partner System Priority:0 MAC Address: 00:00:00:00:00:00 Oper Key: 0 console#
Cisco
actually, u know what. I think I removed the port channel on the VRTX side for testing. Was gonna redo it, so its not there I dont think. I'll have to add it back
I wonder if its because my internal interfaces are all in access mode?
Not familiar with Dell switches but any reason there’s no native vlan on the Dell gigabit interfaces?
Also the native vlans on the etherchannels on both ends don’t match.
maybe missed it? I've added those in.. Thanks
That's fair enough, misses happen. Did it work? Presume you made the native vlans to match on both ends as well?
Dell Switches don’t show the complete config I think. There a command to show everything but I can’t remember. I’m not too familiar with them though. If you put in the command but it doesn’t show in the configuration there’s a good chance it’s active anyway.
Do a channel group mode active on the actual interfaces
on the dell side? You can't, it says bad parameter. I dont think it takes that command
Cisco side.
I think I may have gotten it!
U mind sharing the solution?
edited the original thread with solution!
What was the solution?
edited the original thread with solution!
Commentaire supprimé par un membre de l’équipe de modération
so far, tried both , no luck. Thank you. We'll get there!
