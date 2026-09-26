---
id: collect-260926-mikrotik/mikrotik/firewall-rules-and-interface-lists-best-practice
title: "etc..."
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/firewall-rules-and-interface-lists-best-practice.md
source_anchor: ""
source_lines: [1, 63]
sha256: 8741b9b27b0b2f7fdb7a24bcf190a9c6fd2d81d7ded54af77e001ad105f94532
---

# etc...

Hi everyone,

When building a firewall for a router with multiple networks with different access requirements (for example: VLANs), there are multiple ways to use firewall rules, including:

1. Add a separate firewall rule per VLAN-interface. Multiple firewall rules for the same resource (example: allow DNS queries to router, rule repeated many times for different VLANs).
2. Add a single firewall rule per service/resource, and use interface lists to determine membership of VLANs and whether that rule applies to them.

For example we can have option 1:

```
/interface/list
add name=Trusted
add name=Guest
add name=IoT
add name=Security
/interface/list/member
add interface=vlan10-trusted list=Trusted
add interface=vlan20-guest list=Guest
add interface=vlan30-iot list=IoT
add interface=vlan40-security list=Security
/ip/firewall/filter
add action=accept chain=input comment="accept DNS (UDP) from Trusted" dst-port=53 in-interface-list=Trusted protocol=udp
add action=accept chain=input comment="accept DNS (TCP) from Trusted" dst-port=53 in-interface-list=Trusted protocol=tcp
add action=accept chain=input comment="accept DNS (UDP) from Guest" dst-port=53 in-interface-list=Guest protocol=udp
add action=accept chain=input comment="accept DNS (TCP) from Guest" dst-port=53 in-interface-list=Guest protocol=tcp
add action=accept chain=forward comment"accept from Trusted to IoT" in-interface-list=Trusted out-interface-list=IoT
add action=accept chain=forward comment"accept from Trusted to Security" in-interface-list=Trusted out-interface-list=Security
add action=accept chain=forward comment"accept from Security to IoT" in-interface-list=Security out-interface-list=IoT
# etc...
```

And option 2:

```
/interface/list
add name=Trusted
add name=Guest
add name=IoT
add name=Security
# nested lists here to control access through membership:
add name=DNSAccess include=Trusted,Guest
add name=IoTAccess include=Trusted,Security
add name=SecurityAccess include=Trusted
/interface/list/member
add interface=vlan10-trusted list=Trusted
add interface=vlan20-guest list=Guest
add interface=vlan30-iot list=IoT
add interface=vlan40-security list=Security
/ip/firewall/filter
add action=accept chain=input comment="accept DNS (UDP) from DNSAccess" dst-port=53 in-interface-list=DNSAccess protocol=udp
add action=accept chain=input comment="accept DNS (TCP) from DNSAccess" dst-port=53 in-interface-list=DNSAccess protocol=tcp
add action=accept chain=forward comment"accept from IoTAccess to IoT" in-interface-list=IoTAccess out-interface-list=IoT
add action=accept chain=forward comment="accept from SecurityAccess to Security" in-interface-list=SecurityAccess out-interface-list=Security
# etc...
```

As you can see, option 1 has simpler interface list configuration, keeping firewall/access-control complexity in the firewall filter rules, but has more repetition of rules for the same service. If you add a VLAN or want to allow some access, you need to create potentially dozens of new firewall filter rules for the same services, you could end up with 50+ rules just to allow DNS access, repeating the same thing just for each VLAN that is allowed access. Becomes more complex with more interfaces, but access control is contained in firewall, and interface lists are simpler.

Option 2 has more complex interface list configuration, and a lot of access control has now moved there, but firewall filter rules are simplified and now policy is defined by membership of interface list, more similar to zone style (if member of DNSAccess, they have DNSAccess, etc...). Add a new interface and want it to have DNS, DHCP, etc...? Just add it to the appropriate interface list(s). You only need a single firewall rule per service (DNS, DHCP, etc...). You can control how coarse or fine-grained you want it by nesting/including lists and having a smaller number of "main" lists that control access to multiple services at once (InternetAccess, BasicServicesAccess, etc...). Firewall is simplified and easier to understand, but interface lists get more complex and the nested membership (using include=) might make it more difficult to quickly understand what is going on.

I want to find out what the MikroTik community's view on these methods are, if there's another method that's considered better.

Thank you!
