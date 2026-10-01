---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1795865-tp-link-tl-sg108e-opnsense-vlan-set-up-unsure-of-what-f7ea1197
title: "TP Link TL-SG108E OPNSENSE VLAN set up unsure of what"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1795865-tp-link-tl-sg108e-opnsense-vlan-set-up-unsure-of-what-f7ea1197.md
source_anchor: ""
source_lines: [1, 53]
sha256: c7fab0924f34fc5312565e5d18abbaef6468b038c3e007d275b6437ae7ab8fc4
---

# TP Link TL-SG108E OPNSENSE VLAN set up unsure of what

*Score : 0 | Source : https://superuser.com/questions/1795865/tp-link-tl-sg108e-opnsense-vlan-set-up-unsure-of-what*

I have been following this guide:
https://homenetworkguy.com/how-to/set-up-a-fully-functioning-home-network-using-opnsense/
Which has been great up until the switch configuration for the VLANS where I have become lost with tagged and untagged ports.
My potential VLAN configuration is as follows:
| VLAN | IP | NIC Port | Switch Port | Tagged | Untagged | Notes | 
|---|---|---|---|---|---|---|
| LAN | 192.168.1.1 | 0 | 1 | N | Y | Dedicated LAN management port | 
| USER | 192.168.10.1 | 1 | 2 | Y | N | Can access all other VLANs | 
| GUEST | 192.168.20.1 | 3 | 3 | Y | N | Any unknown devices should be routed here. Does not have access to other VLANs | 
| IOT | 192.168.30.1 | 3 | 3 | Y | N | Does not have access to other VLANs | 
| IPCAM | 192.168.40.1 | 3 | 3 | Y | N | No Internet access. Can access USER for NAS | 
| WORK | 192.168.50.1 | 2 | 4 | Y | N | Does not have access to other VLANs | 
Should I add the WAP (my old router) to port 5 on the switch which would be tagged as well as it would expect VLAN traffic?
When any unknown device is plugged into the switch would it default to LAN? If the MAC address is known it would be routed to its companion VLAN?

---

### Reponse — score 1

When any unknown device is plugged into the switch would it default to LAN?
It will default to whichever VLAN is assigned as "untagged" on that port.
(If your switch has separate "PVID" settings, then packets sent by the device and received by the switch will default to the VLAN that was set as the PVID for that port – but for proper operation, the PVID of each port should always match the untagged VLAN of that port.)
If the MAC address is known it would be routed to its companion VLAN?
No, it doesn't work that way at all. Your VLAN assignment is not based the MAC address, nor on the devices being "known" in general1. Normally 802.1Q VLANs are statically assigned to switch ports and any device connected to a given switch port always goes into the same VLANs.
OpnSense cannot influence VLAN assignment – all it can do is completely refuse to issue an IP address to an "unknown in this VLAN" device, e.g. if your Wi-Fi access point puts your toaster in the "USER" VLAN instead of the "IOT" VLAN, you can make OpnSense refuse to issue a DHCP lease, but you can't make it move the device to the "IOT" VLAN.
1 Although MAC-based VLAN assignment is technically quite possible, that is a feature you'll only find in switches and APs more towards the "enterprise" side, e.g. in TP-Link JetStream/Omada series – usually as part of 802.1X (aka "WPA-Enterprise" for Wi-Fi). Certainly not in the TL-SG108E, which only does static 802.1Q tag-based VLANs at most.
I have become lost with tagged and untagged ports.
"Tagged" and "Untagged" does not apply to ports in the first place. Instead it applies to each VLAN on each port – you wouldn't just have a single column like in your example; you would have a matrix that describes which ports have which VLANs tagged. The same VLAN can (and usually will) be "tagged" on the OpnSense port, yet "untagged" on another port.
(Keep in mind, the whole point of "tagging" is to allow a single port – such as the one going to your OpnSense system – to carry multiple VLANs at once. However, devices need to understand the tagging, so it's usually the switch's job to receive tagged VLANs from OpnSense and expose them as untagged on the other ports.)
So half of your table makes sense, but half of it does not. You need two separate tables: one that describes only the assignments of VLAN IDs to IP subnets (and maybe to OpnSense interfaces), and a second table that describes VLAN assignments to switch ports. For example:
| Interface | VLAN ID | IP | Notes | 
|---|---|---|---|
| LAN | 1 | 192.168.1.1/24 | Dedicated LAN management port | 
| USER | 2 | 192.168.10.1/24 | Can access all other VLANs | 
| GUEST | 3 | 192.168.20.1/24 | Any unknown devices should be routed here. /Not going to work/Does not have access to other VLANs | 
| IOT | 4 | 192.168.30.1/24 | Does not have access to other VLANs | 
| IPCAM | 5 | 192.168.40.1/24 | No Internet access. Can access USER for NAS | 
| WORK | 6 | 192.168.50.1/24 | Does not have access to other VLANs | 
and:
| Switch port | VLAN 1 (LAN) | VLAN 2 (USER) | VLAN 3 (GUEST) | VLAN 4 (IOT) | VLAN 5 (IPCAM) | VLAN 6 (WORK) | 
|---|---|---|---|---|---|---|
| 1 (OpnSense) | Untagged | Tagged | Tagged | Tagged | Tagged | Tagged | 
| 2 (Windows PC) | - | Untagged | – | – | – | - | 
| 4 (IP camera) | - | - | - | - | Untagged | - | 
| 5 (Cheap generic AP) | - | - | Untagged | - | - | - | 
| 6 (Multi-SSID AP) | Untagged | Tagged | Tagged | Tagged | - | - | 
In this example, the only ports that have any "tagged" VLANs are those going to devices which understand tagged VLANs (and can be trusted with them). Everything else only gets a single untagged VLAN (aka an "access port").
For example, OpnSense understands 802.1Q VLAN tags – each "VLAN" interface corresponds to a tag; same with any Linux device. On the other hand, Windows PCs generally can't handle tagged VLANs well (unless they run Hyper-V or have a "server-grade" NIC with its fancy drivers).
If your switch has a "PVID" setting, the PVID of each port must match the ID of untagged VLAN on that port, for symmetry between sent and received packets. (When the switch receives packets without a VLAN tag, it assumes they belong to the 'PVID' VLAN – exactly the opposite of sending packets from the 'untagged' VLAN.)
