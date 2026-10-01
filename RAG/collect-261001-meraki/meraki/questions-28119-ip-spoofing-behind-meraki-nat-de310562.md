---
id: collect-261001-meraki/meraki/questions-28119-ip-spoofing-behind-meraki-nat-de310562
title: "questions-28119-ip-spoofing-behind-meraki-nat-de310562"
domain: meraki
role: reference
task: reference
actors: []
dates: ["2016-02-24"]
keywords: []
source: docs/RAG/collect-261001-meraki/questions-28119-ip-spoofing-behind-meraki-nat-de310562.md
source_anchor: ""
source_lines: [1, 6]
sha256: 45f3ed07dafc90d53dad0c1bbad1d21292e3a3a334a8883a7db58329fda95882
---

# questions-28119-ip-spoofing-behind-meraki-nat-de310562

I am using Meraki MX80 for NAT. There are a bunch of internal subnets, and the NAT works as expected. A colleague raised the question, what happens if a packet spoofs its source IP to something that's not in those internal subnets. Will such packets be dropped?
- 
        Depending on how you have the device configured, the traffic not matching an address range for NAT could be routed normally. You will need to provide the configuration (sanitize any public addresses).Ron Maupin– Ron Maupin ♦2016-02-24 06:07:58 +00:00Commented Feb 24, 2016 at 6:07
1 Answer 1
Agree with Ron Maupin. More details would need to be provided.
With that being said, the default behavior would be to drop the packet if it doesn't hold the route for the IP address. Your packets would send out but the firewall wouldn't be able to return the ingress packets over the same session, thus your route would break. If your device that was spoofing the IP was also configured to respond to the same spoofed IP, then it may work. Of course, none of the spoofed traffic would work if your Meraki wasn't set to NAT the said spoofed subnet/IP in the first place.
