---
id: collect-261001-general-networking/general-networking/r-networking-comments-17k9wph-sporadic-vlan-connectivity-issues-471a616c
title: "r-networking-comments-17k9wph-sporadic-vlan-connectivity-issues-471a616c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/r-networking-comments-17k9wph-sporadic-vlan-connectivity-issues-471a616c.md
source_anchor: ""
source_lines: [1, 28]
sha256: 2cc78a4b0abd203987062fa32722f6fe6a7f7c2dada1e69cb46ef09533a69113
---

# r-networking-comments-17k9wph-sporadic-vlan-connectivity-issues-471a616c

sporadic vlan connectivity issues 
        
        
        
    
    
    Hello,
I am having a few issues with a mish mash network. Running HA pair merakis with latest firmware with2 site-to-sites, then are going into two Cisco MS120-8's. Those are going through a Cisco SG200, and then into a Unifi 48 port switch, then out into eight unifi 48 port switches. There is also a Cisco ASA that has 2 different site-to-site vpns connecting to the Cisco SG200. The wireless is all unifi as well if it matters. The unifi network is being controlled with a unifi cloud key gen2. The unifi switches and controller are all up to date with firmware too.
The meraki site-to-site's are going to our other buildings that also have merakis with different vlans on them. In our main building we have a few seperate vlans setup. The ASA's site-to-sites are going off to other companies that needed to be able to monitor specific equipment (this network belongs to a medical building). The other companies are not using merakis and our initial setup trying to get meraki to work with them wasn't working and we had an ASA laying around that wasn't very old. I will try and make this easy and just use common addresses. Lets say:
      vlan 1   192.168.1.x
vlan 2   192.168.2.x
vlan 3   192.168.3.x  - over site-to-site
vlan 4   192.168.4.x - over site-to-site
vlan 5   192.168.5.x
    
We have terminal servers user connect to everyday. The terminal servers are on vlan 1. About 4-5 times a day, the vlan communication stops for about 5 seconds at a time, enough to end the session. So if someone from vlan 2 is connected to the terminal server, the connection drops, they reconnect, then its fine for maybe an hour, or maybe 15 min, or maybe 5 hours. Its a coin toss. If you are on vlan 1 and are connected to the terminal server, you stay connected all day. I have setup ping plotter with multiple computers on different vlans and have noticed if you are on vlan 2 and are pinging vlan 2, it never drops, same with the other vlans. It only drops when crossing vlans. All servers are updates, all switches are updates. I let the primary meraki fail over to the secondary meraki and the issue continues. I have checked the logs of the unif controller and do not see any errors when the issue occurs. The SG200 doesn't seem to tell me anything.
Anyone have any ideas or suggetions on what I should be checking next? I would love to be able to find an error somewhere instead of just starting to replace equipment.
Section des commentaires
Thank you to those who responded. It turns out this has been happening longer than I thought, just no one said anything and without any errors anywhere, I hadn't seen it going on.
I made a mistake and it turns out even though the firmware on all of the unifi devices were up-to-date, I had missed an update for the Unifi Network (not the control-key firmware, the part above it that says Network). I do not really check the unifi app on my phone, but installed it for testing sake and discovered it there.
I will still be going through the suggestions everyone gave me just to help learn and have a better idea for next time. I very much appreciate the assistance.
Check the VLAN configuration on all devices, especially the Cisco ASA and Cisco MS120-8 switches, to ensure there are no conflicting settings. It's possible that the intermittent drops in communication are related to VLAN misconfigurations or issues with VLAN routing. Also, ensure that your VLANs are properly segregated and that there are no IP conflicts or address overlaps.
How do you segregate VLANs?
you're going to need a drawing to chase this down.
I would think things like looking at switch logs for events during time of issue is where to start. interface counters, resets, spanning tree, ipsec rekey, drops, routing flaps, environment/facilities issues like power events. Start with getting everything in syslog, and ntp. With all devices agreeing on a standard exact time, and logging you at least have a place to start looking. Next try to find a specific test case, ideally one you can control. A specific host to a specific destination. map out every single possible device between the two, look at each device and each interface.
without topology, lets try this.
for troubleshooting sake, pick a downtime and during the downtime, unplug redundant links. e.g. you have 2 cables LACP from 1 switch stack to another, disconnect one of the physical link and see if it resolve, then vice versa.
for consistency, automate your testings using 1 script
