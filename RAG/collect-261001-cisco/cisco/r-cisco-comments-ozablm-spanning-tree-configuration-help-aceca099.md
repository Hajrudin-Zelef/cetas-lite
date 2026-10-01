---
id: collect-261001-cisco/cisco/r-cisco-comments-ozablm-spanning-tree-configuration-help-aceca099
title: "r-cisco-comments-ozablm-spanning-tree-configuration-help-aceca099"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["pruning"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-ozablm-spanning-tree-configuration-help-aceca099.md
source_anchor: ""
source_lines: [1, 52]
sha256: 35c8c4c457728f468e1a7b55b835636ae80e174acf12a0b01f06d1025da3f16c
---

# r-cisco-comments-ozablm-spanning-tree-configuration-help-aceca099

Spanning Tree Configuration Help
So, i have a Router-Switch 1-Switch 2 network. Between switch I configured a trunk for let's say vlan 10, 20, 40. No problem so far. One day I need to add vlan 30 in switch 2 because I need an access to that vlan in one of the port. So I did the usual; add vlan, add allowed vlan in the trunking port of both switch 1 and 2, and configure the port needed as access vlan 30. The problem is, it doesn't work. other vlan can work normally, just that particular vlan 30 isn't working.
After troubleshooting stuff, I noticed one thing: vlan 30 in switch 2 is acting as the root bridge. Other vlans are not. So I checked switch 1, and yes, the vlan 30 in switch 1 isn't the root. So I changed vlan 30 in switch 1 into root. Problem is, it still doesn't work. Now BOTH switch 1 and 2 acting as the root for vlan 30. Tried many things, increasing the priority of vlan 30 in switch 2, deleting and recreating vlan 30 in switch 2, configure vlan 30 as secondary root, etc2 and it's still not working. Honestly I'm out of idea right now.
Anyone here has encountered this problem and maybe help solve this? The 2 switches are both 2960s.
Thanks in advance.
Edit: Here is my run conf and spanning tree for switch 1: https://pastebin.com/kSNHdzc6 and switch 2: https://pastebin.com/6D5Ankdd Another info: I have another switch connected to SWITCH-1 which also trunkin vlan 30 and it's fine, I can create access on the port in that switch. Only SWITCH-2 is having a problem. Another thing, I can ping the gateway of vlan 30 in Switch-2, but if I create interface vlan 30 with vlan 30 ip address, suddenly I can't ping the gateway unless I delete that interface vlan.
Section des commentaires
Are you using VTP? Which you type "show vlan" on SW1, do you see VLAN 30 listed? If not you might not have added the Layer 2 VLAN for it.
If so, SW1 one should be the VTP "server" and SW2 should be "client" and both switches should have the same VTP domain.
so....
SW1:
spanning-tree vlan 1-600 priority 4096 ( Only set this on the Master VTP Server Switch at the site, do not use this command on secondary switches )
vtp mode server
vtp domain "xxx"
SW2:
vtp mode client
vtp domain "xxx"
You can also used VTP mode transparent if you want to manage the VLAN on each switch separately.
Also, make sure on the trunk ports on each switch that you are allowing VLAN 30 by either allowing all:
switchport trunk allowed vlan all (once issued this command won't show on the port config)
OR:
switchport trunk allowed vlan 10,20,30,40 (if you are specifically allowing only specific VLAN's - This command WILL show on the port config)
I have to check again, but iirc I use transparent vtp. Not really a fan of vtp server-client setup. I already allowed vlan 10,20,30,40 on both end of trunk. For switch 1 I already configured vlan 30 root primary AND priority 4096 before. I also configured vlan 30 priority 60654 (forgot the specific number but it was the biggest priority number) just to make sure. And yeah, it's still not working for vlan 30.
If switch 1 and 2 are both root for vlan 30, then it usually means they don't see vlan 30 bpdus from each other.
Make sure that the L2 connectivity is good between the switches for vlan 30.
Verify that vlan 30 is created in vlan database of both switches with "show vlan".
Verify that vlan 30 is allowed across both trunk interfaces with "show interface trunk"
Verify that stp for vlan 30 is in a forwarding state for both trunks with "show spanning-tree vlan 30".
Check log buffer for any relevant error messages with "show logging"
Is the L2 connectivity condition different from each vlan? FYR I edited OP with SW-1 and 2 config. show int trunk also showed that vlan 30 allowed, active, forwarding and not pruned on both switches.
Without VTP, switches won't propagate VLANs across a L2 domain. If you don't have it defined on switch 1, switch 2 has no vlan 30 path to the router. You can either statically configure vlan 30 on SW1 or configure VTP to propagate VLANs.
Static setup
Switch 1:
VTP setup (https://www.cisco.com/c/en/us/support/docs/lan-switching/vtp/10558-21.html)
On both switches:
I use transparent on SWITCH-2, but Server on SWITCH-1. Not really a fan of VTP Server-client config. But I already make the same necessary vlan on both switches.
Transparent will propagate VTP to downstream switches, but won't interact with it locally. You can run both as a server and it'll work fine - each will update the other. I'd avoid transparent unless specifically necessary, and simply remove VTP if you don't like the functionality.
Given your setup, I'd check to make sure VTP pruning isn't trolling you. If you run "show interface trunk" and scroll to the bottom of the output, you should see a section for "Vlans in spanning tree forwarding state and not pruned" - make sure VL30's in there on both ends.
Before you begin, as a precaution let's grab an export of all your VLAN IDs and names:
Copy and paste that to notepad, and just to make things quick and efficient if we need them lets make into a script:
Ok, with a little bit of luck we won't need that, but if we do, we've got it.
Now let's fix up your STP and VTP config:
Switch-1 Configuration:
Switch-2 Configuration:
I'm currently using pvst instead rapid-pvst. does it affect something in my config? Anyway I've updated OP with the config file if you want to take a look. Thanks.
If your equipment supports rapid STP, use rapid STP.
Technically, no. pvst works, and isn't causing any problems.
But traditional STP (or PVST) takes forever to reconverge compared to rapid pvst.
If you don't want to use VTP, that's cool. Just use transparent mode and make sure you've created all of the VLANs in both switches.
The general rule of thumb is for your STP root to be as close to the default-gateway as possible.
I think you said the gateway is connected to SW-1.
Why do you want SW-2 to have a lower priority for VLAN 30, and serve as the root for that VLAN?
