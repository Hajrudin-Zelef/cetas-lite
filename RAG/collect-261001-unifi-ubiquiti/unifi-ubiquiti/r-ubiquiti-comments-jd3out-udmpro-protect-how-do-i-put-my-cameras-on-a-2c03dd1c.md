---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/r-ubiquiti-comments-jd3out-udmpro-protect-how-do-i-put-my-cameras-on-a-2c03dd1c
title: "r-ubiquiti-comments-jd3out-udmpro-protect-how-do-i-put-my-cameras-on-a-2c03dd1c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/r-ubiquiti-comments-jd3out-udmpro-protect-how-do-i-put-my-cameras-on-a-2c03dd1c.md
source_anchor: ""
source_lines: [1, 8]
sha256: 9100a26f8e8887d2b3fd38c1fa1b31bb79b974c5d316352de46643c8a513b316
---

# r-ubiquiti-comments-jd3out-udmpro-protect-how-do-i-put-my-cameras-on-a-2c03dd1c

I'm new to Unifi.  I tried simply creating a new separate vlan for video and then assigning the ports the video cameras on connected to to the VLAN but this did not work - Unifi showed the cameras as "offline" until I put them back on the default network (presumably this is untagged).

I'm not a network admin expert but I have configured Juniper, Cisco and other equipment in the past and have a passing familiarity with L2 protocols.  Can someone point me to some documentation (written or video) that shows me how to isolate my video cameras on a separate vlan?
  

Obviously an attacker should not be able to unplug a cable from a camera, plug that cable into a laptop and use that connection as a way into the network ... Feel free to mention any other best practices (e.g. MAC addr white listing or other things)

As mentioned in the subject line, I have a UDM-Pro (with HDD) and some cameras. I also have a Ubiquiti PoE lite (8 port PoE, 8 regular port) switch and a few unmanaged netgear switches. My plan is for all of my Protect cameras to be connected through Ubiquiti switches only (i.e. no camera on a netgear unmanaged).
