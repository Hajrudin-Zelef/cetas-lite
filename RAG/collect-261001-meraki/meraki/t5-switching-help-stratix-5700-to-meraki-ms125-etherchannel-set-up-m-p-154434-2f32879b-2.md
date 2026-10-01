---
id: collect-261001-meraki/meraki/t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b-2
title: "t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b.md
source_anchor: ""
source_lines: [19, 87]
sha256: 324e6285de4113472cf84cc6424be1d63fa94fb18360e2669d19457c5d476ccd
---

# t5-switching-help-stratix-5700-to-meraki-ms125-etherchannel-set-up-m-p-154434-2f32879b

			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-12-2022 04:27 AM
OK I'm on my computer now so I'll be able to type a longer answer than on my phone.
You gave alot of explanation about your topology in text form which I will not analyse.  It's too much   However I do see you have multiple of these stratix switches that connect over different links towards the meraki switch.
So you just need to focus on what is actually going wrong.
1) your port channel initially comes up
 -> this means your lacp negotiation and timers are matching and this works
2) STP is working on individual links and you have a stable topology
3) you first gave the info that the interfaces are shut down due to STP channel misconfig
Point 3 is where you need to focus.
Use case for the channel-misconfig feature:
Port-channel misconfig is a feature that is primarily meant for static port-channels.  If you for example create a port channel on a switch and you connect both links to another switch.  That will work fine.  However if you were to connect one link to one switch and another to a totally different switch (not a stackmember or linecard on the same switch) then you would get really weird STP behavior on your network.
Operation:
So the misconfig checks for incoming BPDU's.  If you have a working port-channel on the neighbor switch it will only send BPDU's over one of the links, however if you mis uplink to another switch your switch will receive BPDU's from both switches and the bridge-id will not match.  Following this the spanning-tree etherchannel guard misconfig will fire and will put your ports into err-disable until your err-disable recovery kicks in and re-enables the ports and the cycle continues after receiving new BPDU's
In your case, you are presumably correctly uplinking your switches however you are still receiving BPDU's with different bridge id's how is this possible...?
Normal BPDU's from STP, RSTP and MSTP are sent with the normal destination MAC address 01:80:c2:00:00:00. These are understood by the Meraki MS switch and will be processed. However Cisco switches by default run PVST+ or RPVST and these use 01:00:0c:cc:cc:cd. Meraki switches do not run PVST+ or PVST so these frames are seen as regular multicast frames and are just forwarded on.
But you are running MSTP... so what gives..?
Well in a Cisco MSTP network whenever you have a boundary port. That is a port that links to a switches that runs some other flavor of STP it will do a PVST simulation and will send Cisco PVST+ or RPVST BPDU's on each VLAN on that port for compatibility. So if you connect two cisco switches to the same Meraki MS switch those switches will see both BPDU's from the Meraki switch but will also see each other just as if they were directly connected and this WILL cause problems.
So how to fix this:
2 options
- You turn OFF PVST simulation if the firmware on the Stratix switches allow sit. This way it will prevent those switches to send PVST frames on boundary ports
- If that is not an option or you have another device running PVST+ like a router with a switch module or ports in a bridge group then the other choice is to just disable the guard feature with following global config command: no spanning-tree etherchannel guard misconfig. This will turn off the check but please take care in using lacp where possible and be very carefull around static port-channels.
Good luck!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-11-2022 02:27 PM
The channel misconfig happens if the Cisco device receives BPDU's from other Cisco devices behind the MS switch.  Is it possible there are other Cisco products at the other end of the MS125 switch?
If there are you will have to make sure they also run MSTP and disable PVST simulation.  Or else those switches will send PVST BPDU's on each VLAN and the MS125 will just pass them through.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-11-2022 05:05 PM
The topology is 2 sections of OT network VLAN3 / 192.168.3.x and VLAN 4 / 192.168.4.x are tied together via a fiber connection between 2 Meraki MS125 switches. The OT networks have several Stratix 5700 switches each (12) on VLAN 3 and (7) on VLAN 4. The OT Servers hang off the Meraki switches VLAN 3 has 2 severs on VLAN 3 Meraki ports and VLAN 4 has 1 server on VLAN 4 Meraki port. There is also a router / firewall on the VLAN 3 Meraki that routes between VLANs and other things. Anyway, all the 5700 switches use MSTP and the Meraki's use RSTP but everything is fine. EtherChannel between 5700s is straight forward and very easy to do, although you have to use CLI to enable / disable individual ports in an EtherChannel, Once the EtherChannel is configured the device manager web based config tool for 5700s wont allow you to effect the individual ports of an EtherChannel, only the full channel can have setting changes. Anyway, Everything works well as is but there is a single point of failure on VLAN3 where a single cable controls the connectivity of all OT devices on VLAN 3 to the servers. I was hoping to add redundancy to the system by creating an EtherChannel between the 5700 that ties the entire VLAN 3 OT network to the servers hanging off the VLAN 3 Meraki. Each time I try to configure the channel just bounces up and down. It will stay up for about 3 - 4 minutes then all links will drop out for about 30 seconds then come back up. When connectivity is lost the Channel disappears from the MSTP table. It doesn't say BLK or ALT it goes away completely. I have tried lowering the STP priority on the 5700 to force it to Master. It works it becomes master but the same thing happens it falls out for 30 seconds every 3-4 minutes. Then I tried disabling RSTP on the trunk ports of the Meraki ( lost all connectivity) Error messages are always the same in syslog: %PM-4-ERR_DISABLE: Channel-misconfig (STP) error detected on Po1, putting Po1 in err-disable state. I wouldn't normally be so free with config changes on my network but I am in shutdown so production doesn't need access to servers till 7/13. Any advice would be great. I'm an just about out of ideas.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-12-2022 04:27 AM
OK I'm on my computer now so I'll be able to type a longer answer than on my phone.
You gave alot of explanation about your topology in text form which I will not analyse.  It's too much   However I do see you have multiple of these stratix switches that connect over different links towards the meraki switch.
So you just need to focus on what is actually going wrong.
1) your port channel initially comes up
 -> this means your lacp negotiation and timers are matching and this works
2) STP is working on individual links and you have a stable topology
3) you first gave the info that the interfaces are shut down due to STP channel misconfig
Point 3 is where you need to focus.
Use case for the channel-misconfig feature:
Port-channel misconfig is a feature that is primarily meant for static port-channels.  If you for example create a port channel on a switch and you connect both links to another switch.  That will work fine.  However if you were to connect one link to one switch and another to a totally different switch (not a stackmember or linecard on the same switch) then you would get really weird STP behavior on your network.
Operation:
