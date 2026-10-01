---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-20-unifi-l3-adoption-with-dhcp-option-43-on-pfsense-m-c7f55b22
title: "posts-20-unifi-l3-adoption-with-dhcp-option-43-on-pfsense-m-c7f55b22"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-20-unifi-l3-adoption-with-dhcp-option-43-on-pfsense-m-c7f55b22.md
source_anchor: ""
source_lines: [1, 35]
sha256: d9edecb25f8891c484c4d8947d422f5f5a88c417fd940fb6facdf016ecbc0e25
---

# posts-20-unifi-l3-adoption-with-dhcp-option-43-on-pfsense-m-c7f55b22

Unifi L3 Adoption
Recently I installed a lot of new Unifi Sites in local Schools, due to Corona the demand of a working Wifi in School Environments boosted. Most sites we Setup run their Network split to different VLANs, so a simple Layer 2 Adoption is not always possible. But Unifi Controllers luckily have a way to solve this possible adoption issue. So I got used to set both, DHCP Option 43 and unifi DNS name on each site, no matter if VLAN (L3 Structure) or a flat L2 Network. I just set the Options and done.
DHCP Option 43
Depending on what you use as DHCP Server Configuration of the “DHCP Vendor Specific Option 43" is quite different. Most of them have in common that you first need to convert the Static IPv4 of your unifi Controller into a HEX Value. In Front of the HEX-Address we have 01:04: this is for Suboption 01 and length 04:
Tiny DHCP Option 43 Generator
I added a small Form and table to this Post to quickly give you a rough overview how to configure which DHCP. Just fill in your static IPv4 Controller Address. If any other DHCP Server is missing and you like me to add it let me know the Options. But this should give enough of an idea how to do it for your setup.
| DEVICE | COMMENT | SETTINGS | 
|---|---|---|
| pfSense | Add Custom DHCP Option Number 43 with Type String | Number: 43Type: StringValue: 01:04:... | 
| OPNSense | Go to Services > DHCPv4 > [the relevant interface] > Additional Options | Number: 43Type: StringValue: "01:04:..." | 
| Sophos UTM | Net Services > DHCP > New DHCP Option | Name: unifiTyp: hexHex: 01:04:...Scope: Vendor IDVendor ID: ubnt | 
| Mikrotik RouterOS via GUI | IP > DHCP Server > Click the "+" sign to create a new option | Name: unifiCode: 43Value: ... | 
| Mikrotik RouterOS via CLI | SSH or Terminal into your Mikrotik and type: | /ip dhcp-server option add code=43 name=unifi value=.../ip dhcp-server network set 0 dhcp-option=unifi | 
| Cisco CLI | Create or modify a DHCP Pool. (help) Then add option: | option 43 hex ... | 
| Peplink Router | Go to DHCP Option 43 and enter those Values | Type: ByteHex-Button: checkedValue: ... | 
| Windows Server | Go to Scope Options and add Vendor Specific Info | Binary: ... | 
| DD-WRT | services tab, enabled Dnsmasq, Additional Options: | dhcp-option=vendor:ubnt,1,... | 
| DNSMasq | Add this option to your pool config | dhcp-option=vendor:ubnt,1,... | 
| ISC DHCP server (dhcpd.conf) | Linux ISC DHCP server via dhcpd.conf. | Add to your dhcpd.conf: option space ubnt; option ubnt.unifi-address code 1 = ip-address; class "ubnt" { match if substring (option vendor-class-identifier, 0, 4) = "ubnt"; option vendor-class-identifier "ubnt"; vendor-option-space ubnt; }Add to the subnet part, short example: subnet 10.10.10.0 netmask 255.255.255.0 { range 10.10.10.100 10.10.10.160; option ubnt.unifi-address ...; } | 
| Palo Alto | Add DHCP Option 43 (Pan OS 7.1) | Name: unifiOption Code: 43Vendor Class ID: ubntOption Type: IP AddressOption Value: ... | 
Unifi adopting via DNS
Secondary to the Option 43 I also add the hostname unifi into my local DNS Server pointing as A Record to my static IPv4 of the Unifi Controller. Don’t forget to also specify a domain if your DHCP Server sets a local domain. So if your local domain is “.home.arpa” add for Example:
unifi.home.arpa A 192.168.1.2
I indeed prefer to use actually both methods, because depending on your local domain I already had some problems where the DNS only way didn’t work as expected.
Screenshots Example pfSense
Most times we use pfSense as Router, DHCP and local DNS. As a quick Example here are some Screenshots how to configure DHCP 43 and DNS in a pfSense:
- Services > DNS Resolver > General Settings > Host Override
- Services > DHCP Server > Additional BOOTP/DHCP Options
Comments
Nothing is loaded from external services until you click.
With a Fediverse or Mastodon account you can reply to this post. Replies from any compatible server are shown below once loaded.
Mastodon replies
Web mentions
Leave a comment
Reply to this post on Mastodon and your reply will show up here.
