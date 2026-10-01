---
id: collect-261001-general-networking/general-networking/manual-aliases-html-a3e5f567-2
title: "manual-aliases-html-a3e5f567"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/manual-aliases-html-a3e5f567.md
source_anchor: ""
source_lines: [100, 180]
sha256: 3fa7162db463dfe2dc884a95725a13b260579a46a4a37f8ecefb0cf151feeb22
---

# manual-aliases-html-a3e5f567

of the csv. The file itself should be compressed using gzip
This format requires a [zip] file containing the following csv files:
|  |  |  |  | 
|---|---|---|---|
| Filename | Purpose | Format | Example | 
|---|---|---|---|
| %prefix%-locations-en.csv | maps geo locations to iso countries | geoname_id,,,,country_iso_code | 1,,,,NL | 
| %prefix%-IPv4.csv | IPv4 networks | network,geoname_id | 2.21.241.0/28,1 | 
| %prefix%-IPv6.csv | IPv6 networks | network,geoname_id | 2001:470:1f15:210::/64,1 | 
The %prefix% can be used to identify the product and/or vendor, in MaxMind’s case these files are named
GeoLite2-Country-Locations-en.csv, GeoLite2-Country-Blocks-IPv4.csv, GeoLite2-Country-Blocks-IPv6.csv for example.
Tip
Geo IP lists can be rather large, especially when using IPv6. When creating rules, always try to minimize the number of addresses needed in your selection. A selection of all countries in the world not being the Netherlands can usually be rewritten as only addresses from the Netherlands for example.
Tip
If the number of items is larger than the allocated alias size, you can assign more memory to aliases.
Network group
Combine different network type aliases into one, this type of alias accepts other host type aliases (networks, hosts, …).
Although nesting is possible with other alias types as well, this type only displays valid aliases easing administration, functionally
a Networks type alias can do the same but uses a different presentation.
Dynamic IPv6 Host
An IPv6 Dynamic Host is used where the system is using a dynamic prefix on the LAN, a tracking interface. When the prefix changes, either due to the ISP changing the prefix at will or the prefix changes when the WAN connection is reset, any alias containing an address of a client such as a server on the LAN would no longer be valid.
For example, you obtain a prefix 2001:db8:2222:2800::/56. You have a /56 prefix and if the tracking id was set to 0 for your LAN, you would have an address range on your LAN of 2001:db8:2222:2800:: to 2001:db8:2222:2800:FFFF:FFFF:FFFF:FFFF.
You want to run a server on your LAN that is accessible from the WAN so you give it a static address of 2001:db8:2222:2800:1000:1000::1 and create a rule allowing traffic to access the server.
When your prefix changes, that static address is no longer valid, so you must use the Dynamic IPv6 Host to create an alias address for the firewall entry that automatically tracks the prefix and changes the rule.
The Dynamic Host Alias will always split on the /64 boundary, it will take the upper 64 bits from the interface you select and the lower 64 bits from the address you enter. It does not matter what size your prefix delegation is.
Create a new IPv6 Dynamic Host alias and enter only the suffix of the address, in this example, we will enter the lower 64 bits of the address, you would enter ::1000:1000:0000:1, note the ‘::’ at the start of the address, you MUST always start the address with a ‘::’. You do not need to enter a size after the address i.e. /128 as that is automatically assumed.
Select the interface you wish to use for the source of the upper 64 bits, in this case we will select the LAN interface.
When the prefix changes, the alias address will then be updated in the firewall rules, let’s say your prefix changes to 2001:db8:2222:3200::/56 the rule updates and the entry for your server in the firewall would update automatically to be 2001:db8:2222:3200:1000:1000::1
Let’s take another example, you have a /48 prefix delegation, you have two LAN interfaces and a server on each. You would need to create two separate Dynamic IPv6 Host entries, one for each LAN. For simplicities sake we will use the same address for each server on each interface, you would enter ::aaaa:bbbb:cccc:0001 as the address.
| Upper 64 bits, taken from LAN 1 Interface | Lower 64 bits - Your server address | 
|---|---|
| Server 1: 2a02:1234:5678:0000 | aaaa:bbbb:cccc:0001 | 
| Server 1 GUA address is: 2a02:1234:5678:0000:aaaa:bbbb:cccc:0001 |  | 
| Upper 64 bits, taken from LAN 2 Interface | Lower 64 bits - Your server address | 
|---|---|
| Server 2: 2a02:1234:5678:0001 | aaaa:bbbb:cccc:0001 | 
| Server 2 GUA address is: 2a02:1234:5678:0001:aaaa:bbbb:cccc:0001 |  | 
The prefix changes, in this case we have a /48 prefix, so the new prefix is 2a02:1234:5679/48 our aliases would update to give us the following addresses:
| LAN 1: Server 1 GUA address is: | 2a02:1234:5679:0000:aaaa:bbbb:cccc:0001 | 
| LAN 2: Server 2 GUA address is: | 2a02:1234:5679:0001:aaaa:bbbb:cccc:0001 | 
You may enter multiple addresses, for example if you have several servers on the same LAN segment, just add the suffix for each one. In the example below we have three servers.
BGP ASN
With this alias type you are able to select networks by their responsible parties.
Using BGP parties announce the addresses they are responsible for to each other.
For example Cloudflare uses AS number 13335, Microsoft is known to use 8075.
More background and how addresses are assigned is explained on wikipedia
External
The contents for external alias types is not administered via our normal alias service and can be practical in scenarios where you want to push new entries from external programs. Such as specific lockout features or external tools feeding access control to your firewall.
In you can always inspect the current contents of the external alias and add or remove entries immediately.
Tip
When changing alias contents which are used on firewall rules with state tracking enabled, you might need to remove the specific state before the new rule turns active. (see )
Tip
Since external alias types won’t be touched by OPNsense, you can use pfctl directly in scripts to manage
its contents. (e.g. pfctl -t MyAlias -T add 10.0.0.3 to add 10.0.0.3 to MyAlias)
OpenVPN group
This alias type offers the possibility to build firewall policies for logged in OpenVPN users by the group they belong to as configured in .
The current users that are logged into OpenVPN can be inspected via , the alias just follows this information and flushes the attached addresses to the item in question.
For example, when a user named fred which is a member of group remote_users logs into OpenVPN and received a tunnel address
of 10.10.10.2, the alias containing “remote_users” would include this address as well.
Note
For this mechanism to work, the common-name of the user certificate must match the username exactly, which is the case by default if the certificate has been created from the user manager.
Tip
When using LDAP (Active directory), you can synchronise group membership to avoid double administration in OPNsense.
Internal (automatic)
Internal aliases are prefixed with __ so they are easy to identify and can’t overlap with any user defined ones.
These aliases help you to determine what the content is for some internal concepts such as “LAN network”. Using
the  menu item you can inspect their contents at any time.
Using Aliases in Firewall Rules
Aliases can be used in firewall rules to ease administration of large lists. For instance we might need a list of remote IP addresses that should have access to certain services, when anything changes we only need to update the list.
Let’s create a simple alias to allow 3 remote IP addresses access to an ipsec server for a site to site tunnel connection:
- 192.168.100.1
- 192.168.200.2
- 192.168.202.2
We call our list remote_ipsec and update our firewall rules accordingly.
Note
The list icon identifies a rule with an alias.
Export / Import
The alias admin page () contains a download and an upload button in the footer of the table, with this feature you can
merge aliases into the configuration and download a json formatted list of all aliases in the system.
Since data is validated before insertion, it shouldn’t be possible to import defective data (if the import fails, a list of errors is presented).
Tip
