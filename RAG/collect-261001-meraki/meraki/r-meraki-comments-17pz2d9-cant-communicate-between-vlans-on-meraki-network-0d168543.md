---
id: collect-261001-meraki/meraki/r-meraki-comments-17pz2d9-cant-communicate-between-vlans-on-meraki-network-0d168543
title: "r-meraki-comments-17pz2d9-cant-communicate-between-vlans-on-meraki-network-0d168543"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-17pz2d9-cant-communicate-between-vlans-on-meraki-network-0d168543.md
source_anchor: ""
source_lines: [1, 60]
sha256: 98748f166a494048d63abcb98304438fd3014cf00f7764ea2382f45fe00eded0
---

# r-meraki-comments-17pz2d9-cant-communicate-between-vlans-on-meraki-network-0d168543

Can't communicate between VLAN's on Meraki network 
        
    Hey all,
For starters, I'm a super newbie to Meraki. Have a decent amount of experience with Cisco CLI, but thus far that knowledge has not been useful with this configuration.
Here's a quick rundown of the equipment I was given:
-Client ordered brand new Meraki devices to overhaul their network that includes
- 
      (3) MS125-24P switches
- 
      (1) MS120-48LP Switch (being setup as the core switch)
- 
      (1) MX67 Firewall
Here's how I was instructed to set it up:
-Setup 2 VLAN's on the network: 1 for Management (Mgmt), 1 for Security (Sec).
- 
      Any device on the Mgmt Vlan should be able to access the devices on the Sec VLAN.
- 
      Devices on the Sec Vlan should NOT be able to access devices on the Mgmt VLAN
- 
      Both should have internet access
Here's the progress I've made:
- 
      Both VLANs have been setup and have internet access.
- 
      Sec -> Mgmt traffic is being denied via an outbound rule. Confirmed this was working, as a PC plugged into the Sec VLAN can't ping any of the switches, while a PC on the Mgmt VLAN can.
- 
      Where I am getting stuck is in allowing Mgmt -> Sec traffic. I have tried defining inbound and outbound rules that specify that the Mgmt VLAN is allowed to access the Sec VLAN. No dice. I tried defining a group policy with overriding firewall rules and manually assigning it to one of the connected PCs. No dice.
I am not sure what is not configured correctly at this point. The Outbound rules define an allow any any rule as its default rule. Doesn't that mean the devices should be able to ping each other without me having to define it explicitly? This has led me to believe that the switch ports are configured incorrectly. My boss was the one who went through and defined all the ports as access or trunk and I had assumed up to this point they were all done correctly.
Some port details that might be helpful:
- 
      Port 2 on the firewall is an access port for the Mgmt VLAN
- 
      Port 5 on the firewall is an access port for the Sec VLAN
- 
      PC1 is plugged into port 5 on the Core switch. Port 5 on the core switch is a trunk port that is allowing traffic from all VLANs.
- 
      PC2 is plugged into port 38 on the Core switch. Port 38 is an access port on the Sec VLAN.
If there is any other information you need to help diagnose this issue, please feel free to ask. I can't provide screenshots as that would be in violation of company policy, but I am more than happy to try and explain the settings I have setup in detail.
Thank you all in advance,
Special-Ad1396
UPDATE:
Hey all... So as it turns out it was the PC's themselves.
Both Windows Firewall and our AV were blocking the ping requests. Figured this out when I tried plugging in a blank PC with Windows firewall shut off and it works now... Sorry for the mistake.
Section des commentaires
So basically, L2 forwarding works as you can ping devices within the same vlan, whereas any inter-vlan traffic via the L3 (MX) is failing? You're correct in the outbound rules section being applicable, and that the L3 rules here are stateful. https://documentation.meraki.com/General_Administration/Cross-Platform_Content/Using_Layer_3_Firewall_Rules
Initial thoughts:
What are the port settings on the core facing the MX?
Are there any switch ACLs setup?
Have you taken any pcaps to see traffic even arrive at the L3 boundary (the MX) before being routed or if it's dropped?
I'll add to my bullet points from the post for less confusion:
Port 2 on the firewall is an access port for the Mgmt VLAN. This is connected to port 1 on the core switch, which is set as a trunk port allowing all VLANs.
Port 5 on the firewall is an access port for the Sec VLAN. This is connected to port 37 on the core switch, which is set as an access port for VLAN 10. (I tried switching this to a trunk port allowing all VLANs -No dice)
No ACL's. This is pretty much a brand new OOB setup, minus the VLAN's and switch ports configured.
Have not captured any packets yet. Next closest thing I have done is a tracert that hits the default gateway and then falls off. Weirdly enough, I can ping the default gateway of the other VLAN... Not sure if that is helpful information.
Why is the MX configured with access ports if we're trunking the other end of the link? Trunk both ends (and stp block the redundant uplink)?
Think how the MX would behave when receiving tagged packet on its Access port (which should be received as untagged).
Still thinking on this one.. but you say PC1 is plugged into a trunk port (port 5)? I think that may be the problem.. You need to plug PC1 into an access port that is on mgmt vlan.
Gave this a shot. Did not seem to change anything. Still unable to ping the PC on the Sec VLAN from the PC on the Mgmt VLAN.
Just for clarification, all those switches (MS120/125) are only layer 2. So I would be calling the MX the "Core" as it does all the Layer 3 routing. The links to all switches should be trunks (with native VLAN for management) Then access on the ports with the PCs
Was gonna say check windows FW but then got to the end of your post :)
