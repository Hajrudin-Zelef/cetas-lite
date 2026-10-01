---
id: collect-261001-meraki/meraki/r-meraki-comments-wro6v2-troubleshooting-mx-switch-configuration-d5286aa3
title: "r-meraki-comments-wro6v2-troubleshooting-mx-switch-configuration-d5286aa3"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-wro6v2-troubleshooting-mx-switch-configuration-d5286aa3.md
source_anchor: ""
source_lines: [1, 55]
sha256: f6e0528d24f6b30b6d484464c7f5d128ea4f70e2734cf4a4307d9590212fc44d
---

# r-meraki-comments-wro6v2-troubleshooting-mx-switch-configuration-d5286aa3

Troubleshooting MX > Switch configuration 
        
    Wondering if someone could weigh on on this scenario, as I'm confused what happened.
Has anyone else run into this situation?
Network Info:
What we have:
[MX]----(Trunk - Allow All/Native VLAN 2)----[SW]---(Trunk - Allow All/Native VLAN 8)----[AP]
I can provide more info of course, I just have to be careful since this is a customer network.
- 
      All Meraki gear - MX84, MS225-48LP, and MR33's.
- 
      MX is static of course, the Switch and AP's were static, but since we rolled back, we left them as DHCP for the next attempt.
- 
      DHCP is at the MX, DHCP scopes are defined for the old and new VLANs.
- 
      Using DNS servers located at a data center on existing and new VLANs.
- 
      New VLANs are enabled on the VPN tunnel at the MX.
- 
      Switch Mgmt. VLAN on Dashboard is configured as VLAN 2.
- 
      All IP's are in completely different subsets, and all /24.
- 
      The switches are not routing, nor have any L3 interfaces. Strictly layer 2.
The goal:
Change VLANs from 2 & 8 to 21 and 81, and update the subnets as well. Basically just replicating what we already have, but with different VLAN's and subnets:
[MX]----(Trunk - Allow All/Native VLAN 21)----[SW]---(Trunk - Allow All/Native VLAN 81)----[AP]
Troubleshooting:
I started at the AP's, and since they were set to static IP, I converted them to DHCP and then cycled the corresponding switchport to reboot them.
They rebooted, came up with an IP in the DHCP range for VLAN 8 as expected.
I then changed the native VLAN for each AP, waited a moment, then cycled the port once again to reboot each AP. As expected, they rebooted and came back with an IP in the DHCP range VLAN 81.
The problems started when I started to move upstream to the switch; when attempting to change the native VALN from 2 to 21 on the uplink port, we lost connection with the switch temporarily, but then received the error on the dashboard that DNS was misconfigured. On the MX side, we configured the port connecting to the Switch's uplink port to native VLAN 21, trunk, allowing all VLANs. Again, basically the same, just a VLAN change and IP change. I should also add that the switch was set to DHCP at this time.
The DNS servers the switch would have used are located on the remote end of the auto VPN tunnel (this is one of many sites that connect back over VPN to a central MX in a data center), but they are the same DNS servers being used on the AP's we changed downstream from the switch, and the AP's were working fine during this time.
There is no overlap with this sites subnets and any other site's subnets.
VLAN 21 & 81 are enabled on the VPN.
I had Meraki support on the line and they confirmed that the switches were sending the DNS queries for the Meraki dashboard, and confirmed through the Pcap that the MX was receiving responses from the DNS servers as expected, but that the LAN interface back to the switch was not showing that they were passing the DNS responses back to the switch. Again the same configuration here.
We changed the native VLAN back to 2, and the switch worked fine. At this point, I reverted all changes and backed out of the maintenance until I could figure out what was going on.
Currently the switch is still on DHCP until we try to cutover again.
The Rub:
No cellular fail-over, and because of physical distance, we cannot be on site to physically troubleshoot the switch. We have a WattBox out there to help with power cycling, but that won't resolve a switch pulling down a incorrect configuration from the dashboard. And it's offline from the network (apparently for the last 5 months)
I'm thinking this is a simple misconfiguration, but my experience in Meraki is not all encompassing, so I'm reaching out for help. Any insight is greatly appreciated.
Section des commentaires
Have you tried setting the native vlan the same across your network?
Hi, thanks for the reply; do you mean like this?
[MX]----(Trunk - Allow All/Native VLAN 21)----[SW]---(Trunk - Allow All/Native VLAN 21)----[AP]
Yes, the native VLAN for your Meraki devices should be consistent across all uplinks and also made the mgmt VLAN for your switch(es).
Correct. What is the result?
Typically I would configure a routed link if moving traffic between VLANs, but if there is no need to do so, I would expect the frames VLAN ID to remain the same, even across Trunks, which means a consistent native VLAN.
So when updating trunk lines you want to move from the outside in. In your example ap first, then switch, then mx. If doing this remotely, you have to remember dashboard changes are not instant so expect to loss connection to the end device when the change happens due to the vlan mismatch. Once the further out device gets updated then you update the link on the closer device. So AP port first, when it downloads its config and losses connection, the do the switch port the AP uses. This advice is only for doing this remotely as you are on the internet side of the mx and want to work back to you.
Also ensure the vlans can ve reached by the destination before changing it
It sounds like a route back to those VLANs or did you not change ips? Does the switch or the Mx have the default gateway ip?
Change at ap first then switch then mx
It seems your initial config and your operations were spot on. You also don’t need to match your AP native VLAN to your Switch native Vlan. I my own designs I use an unused native VLAN on links between switches. I used the switch mgmt vlan as native between upstream firewall and switch and I use ap mgmt VLAN as native on ap ports.
The only insights I can give are. When changing uplinks towards MX’es make sure you have all VLANs allowed both on the mx And on the switch before changing the native on the switch before on the mx.
You could test by temporarily using another VLAN as native to the mx that is not your switch mgmt and test if your traffic goes through on your new switch VLAN so you can be sure there is no blocking on the firewall. And then move that new VLAN to native.
