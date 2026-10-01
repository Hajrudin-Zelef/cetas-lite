---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1199437-is-this-a-mac-address-collision-issues-on-a-unifi-router-2157c336
title: "questions-1199437-is-this-a-mac-address-collision-issues-on-a-unifi-router-2157c336"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Samsung"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1199437-is-this-a-mac-address-collision-issues-on-a-unifi-router-2157c336.md
source_anchor: ""
source_lines: [1, 19]
sha256: 42ff961c3e6f9bfae6cf75fb076832e9efbfc28a7f34da8ab0bb97273ba8d07b
---

# questions-1199437-is-this-a-mac-address-collision-issues-on-a-unifi-router-2157c336

I am using a Unifi Dream Machine for the router and DHCP. I have been adding some Meross Wi-Fi smart plugs around the office to reduce power wastage.
- I add them to Meross App and
- Then allocate static IP addresses to each.
On the 7th one,:
- I saw the name in the row (on the Unifi Clients screen) flashing between "Smart Plug" and "SM-S911B".
- The MAC address is c4:e7:ae:28:69:f2 , which is in the valid Meross range.
- When I clicked on it, on the right side bar, by the time I entered something (such as name or the static IP), the name would change and the fields would reset.
My initial thought was that there was an IP address collision:
- So I removed the plug and deleted the offline entry for the IPaddress in question.
- I plugged it in, and started seeing the same behaviour.
- I powered off both Samsung phones of this model, in the office. the problem disappeared
- I powered up the phones, the problem appeared.
- I removed their offline entries, and repowered the phones, the problem appeared.
- If I only have the phones on, one of them shows up with the same MAC (mentioned above).
To me, this rules out IP address collision and I think it may be a MAC address collision.
I know that the random MACs are supposed to have the random bit set, so the MAC addresses from Samsung should not be in the Meross range. But this is only true if the developer programmed it correctly.
So, my guess is that it is a MAC clash. I am blaming the Samsung's MAC randomization feature.
I turned off the random MAC on these phones and repowered them - but the problem remained.
Is there a way I can sort this out?
