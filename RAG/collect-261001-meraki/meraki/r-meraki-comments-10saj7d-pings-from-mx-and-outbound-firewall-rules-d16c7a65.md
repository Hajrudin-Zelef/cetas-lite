---
id: collect-261001-meraki/meraki/r-meraki-comments-10saj7d-pings-from-mx-and-outbound-firewall-rules-d16c7a65
title: "r-meraki-comments-10saj7d-pings-from-mx-and-outbound-firewall-rules-d16c7a65"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-10saj7d-pings-from-mx-and-outbound-firewall-rules-d16c7a65.md
source_anchor: ""
source_lines: [1, 11]
sha256: 2476b82f79d291765b42bb6438dcd2cc8030d5f547884f5a9e08269be17e093d
---

# r-meraki-comments-10saj7d-pings-from-mx-and-outbound-firewall-rules-d16c7a65

Pings from MX and outbound firewall rules 
        
    Are sourced from the MX affected by the L3 outbound firewall rules? If I have a rule that is supposed to block inter-VLAN traffic, will the MX vlan 10 interface be able to ping a device in vlan 20?
Right now I have a L3 outbound firewall rule that includes a rule that denies RFC1918 addresses to RFC1918 addresses (using policy objects that include the CIDR ranges). However, the MX unit can still ping from the VLAN 10 interface to a local device in VLAN 20.
Unfortunately, I do not have a way to get into a device on the LAN to test the rules, so I need to use the MX unit. But I cannot tell if the rules are taking effect, or if the MX-sourced pings are an exception.
Section des commentaires
Is the MX where all your layer 3 interfaces for internal routing are? We use a ms425 for internal layer 3 routing and you can create ACL rules there to block any inter vlan traffic. Unfortunately ACLs are only available at the switch level. If I'm not mistaken the layer 3 firewall rules only apply between WAN and LAN traffic. You can test this by creating a rule to block vlan 10 to vlan 20 and see if you get any hits.
All the layer 3 interfaces are on the MX units. The switches only do layer 2 - the switches trunk to the MX units.
The MX will not block its own traffic - the 'ping' tool from the MX is therefore not a good way to test functionality.
You'll need to fully confirm with a test client
Thank you very much for this response! This is exactly what I needed to know.
