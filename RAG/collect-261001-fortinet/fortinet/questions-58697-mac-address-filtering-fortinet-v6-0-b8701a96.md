---
id: collect-261001-fortinet/fortinet/questions-58697-mac-address-filtering-fortinet-v6-0-b8701a96
title: "questions-58697-mac-address-filtering-fortinet-v6-0-b8701a96"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2019-04-24"]
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-58697-mac-address-filtering-fortinet-v6-0-b8701a96.md
source_anchor: ""
source_lines: [1, 23]
sha256: 77352e2f80c0c98a9eef1b1acdcbd1f48019a9687268641d45600047eb35bf71
---

# questions-58697-mac-address-filtering-fortinet-v6-0-b8701a96

I am having trouble configuring mac filtering. I do not use DHCP but am trying to figure out a way to set a mac address on a port and create a rule to block that specific mac address. So when I try and send traffic to that port, I am unable to. Any thoughts?
- 
        You need to be more specific. What switch model(s) do you have? Also, understand that security using addresses (IP or MAC) is often worse than no security because you get a false sense of security. It is extremely easy for someone to change the MAC address on a device to be that of an allowed device. Please edit your question to include more information.Ron Maupin– Ron Maupin ♦2019-04-24 13:38:49 +00:00Commented Apr 24, 2019 at 13:38
1 Answer 1
I assume that you are using Fortigate firewall. You can use below codes for your work
config firewall ipmacbinding setting
set bindthroughfw {enable | disable}  - this is enabling IPMAC binding to get through a Firewall.
set bindtofw {enable | disable}  - this will check an IP MAC binding combination to allow access TO the firewall
set undefinedhost {allow | block} - this defines how the Firewall will treat traffic that has not been bound
end
config firewall ipmacbinding table
edit <index_int> - the number in the IP/MAC binding table
set ip <address_ipv4> - IP address value
set mac <address_hex>  - MAC address value
set name <name_str> - the name which may be used for this binding
set status {enable | disable} - is the binding now enabled
end
config system interface
edit <interface name>
set ipmac {enable | disable }   - enable to enable mac binding on interface
next
end
In addition to that if you want to config MAC address control you can use below article which published by fortinate
