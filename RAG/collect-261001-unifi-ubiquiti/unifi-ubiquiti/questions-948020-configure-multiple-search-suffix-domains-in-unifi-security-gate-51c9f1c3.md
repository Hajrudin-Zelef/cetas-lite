---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-948020-configure-multiple-search-suffix-domains-in-unifi-security-gate-51c9f1c3
title: "questions-948020-configure-multiple-search-suffix-domains-in-unifi-security-gate-51c9f1c3"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-948020-configure-multiple-search-suffix-domains-in-unifi-security-gate-51c9f1c3.md
source_anchor: ""
source_lines: [1, 22]
sha256: 98f2b6c76c9ff04c64b7cb4c694f7d74ea0d61a860bab9e7747e1793c43e82a4
---

# questions-948020-configure-multiple-search-suffix-domains-in-unifi-security-gate-51c9f1c3

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
6
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I would like to have the DHCP server in my Unifi Security Gatway include multiple domain search entries as part of DHCP option 119 so that I can use shortnames for multiple suffixes like:
host -> host.example.com
anotherhost -> anotherhost.home.arpa
container -> container.somehost.lxd
Currently the Domain Name: example.com entry in the Network settings is used for the option domain-search "example.com" entry in /opt/vyatta/etc/dhcpd.conf, but this field does not allow you to specify multiple entries (probably for good reason).
I attempted to enable a custom DHCP option for Code 119, but this seems to use a raw hex value in the config file like option domain-search 65:78:61:6d:70:6c:65:2e:63:6f:6d:20:68:6f:6d:65:2e:61:72:70:61; for a value of "example.com home.arpa". It also doesn't clear the original entry, which likely causes issues with how the values get encoded on the wire. I tried various values but none seem to show up correctly in /var/lib/dhcp/dhclient.leases
How can I configure multiple search suffixes for my Unifi network?
I'm sure you've figured this out by now, but you can set a "text" type for a custom DHCP option, then just enter a comma separated domain list. Note: it overrides the default provided by the DHCP server, so you need to include that too.
I know that this is quite old, however using the text type as a custom DHCP option didn't work for me.
I needed to add a custom code 119 of type hex array and convert the search list to hex.
I wrote a simple python script to convert the domain list and just adding the output from that script worked like a charm for me.
You can now do this in v2.0.9-hotfix.6, but I haven't been able to find the answer anywhere. Here's what worked for me:
configure
set service dhcp-server shared-network-name [your dhcp network name] subnet [your IP subnet] domain-name "subnet1.x.com subnet2.x.com x.com"
commit ; save
You should end up with correctly formatted option domain-name and option domain-search lines in /opt/vyatta/etc/dhcpd.conf.
