---
id: collect-261001-general-networking/general-networking/manual-aliases-html-a3e5f567-1
title: "manual-aliases-html-a3e5f567"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/manual-aliases-html-a3e5f567.md
source_anchor: ""
source_lines: [1, 99]
sha256: 2c08d16ab71f57d03731b7f40d3226465ff4a53b0a7d0c16d96c85eb08553dbe
---

# manual-aliases-html-a3e5f567

Aliases
Aliases are named lists of networks, hosts or ports that can be used as one entity by selecting the alias name in the various supported sections of the firewall. These aliases are particularly useful to condense firewall rules and minimize changes.
Aliases can be added, modified and removed via .
Overview
The page opens with an overview of the aliases available, including some internal ones available from installed software. At the top of the page you can search for aliases or preselect various types or categories to which they belong.
In order to gain some insights into the current status of the plugins, two additional (meta) fields are being displayed, being:
- Loaded# 
  - Shows the amount of entries planned to install into the table, in case there’s not enough memory available to load the item in question, one can easily find the alias causing the overflow in table entries (bar at the right top of the page)
- Last updated 
  - Shows the last timestamp from the entries saved to disk.
Note
The fields above are only used for aliases that contain either networks or hosts, port type aliases are part of the rule and thus not visible in any table.
Alias Types
OPNsense offers the following alias types:
| Type | Description | 
|---|---|
| Hosts | Single hosts by IP or Fully Qualified Domain Name or host exclusions (starts with “!” sign) | 
| Networks | Entire network p.e. 192.168.1.1/24 or network exclusion eg !192.168.1.0/24 | 
| Ports | Port numbers or a port range like 20:30 | 
| MAC addresses | MAC address or partial mac addresses like f4:90:ea | 
| URL (IPs) | A table of IP addresses that are fetched once | 
| URL Tables (IPs) | A table of IP addresses that are fetched on regular intervals. | 
| URL Table in JSON format (IPs) | A table of IP addresses that are fetched on regular intervals. (using a json structure) | 
| GeoIP | Select countries or whole regions | 
| Network group | Combine different network type aliases into one | 
| Dynamic IPv6 Host | A Host entry that will auto update on a prefixchange | 
| BGP ASN | Maps autonomous system (AS) numbers to networks where they are responsible for. | 
| OpenVPN group | Map user groups to logged in OpenVPN users | 
| Internal (automatic) | Internal aliases which are managed by the product | 
| External (advanced) | Externally managed alias, this only handles the placeholder. Content is set from another source (plugin, api call, etc) | 
Hosts
Hosts can be entered as a single IP address, a range (separated with a minus sign, e.g. 10.0.0.1-10.0.0.10)
or a fully qualified domain name.
When using a fully qualified domain name, the name will be resolved periodically (default is each 300 seconds).
Apply changes and look at the content of our newly created pf table.
Go to and select our newly created youtube table.
As you can see there are multiple IP addresses for this domain.
Tip
To change the alias domain resolve interval, go to and set Aliases Resolve Interval to the number of seconds to refresh.
Hosts type Aliases can contain exclusion hosts. Exclusion addresses starts with “!” sign (eg !192.168.0.1) and can be used to exclude hosts from Network Group Aliases.
Warning
Please note that the Flush action is not persistent!
“flush” means flush the current contents of the alias, which will be repopulated when it’s not an external type, so flush in most cases isn’t very useful.
Same behaviour applies to the API call alias_util flush
Networks
Networks are specified in Classless Inter-Domain Routing format (CIDR). Use the the correct CIDR mask for each entry. For instance a /32 specifies a single IPv4 host, or /128 specifies a single IPv6 host, whereas /24 specifies 255.255.255.0 and /64 specifies a normal IPv6 network. Network type Aliases can contain exclusion hosts or networks. Exclusion addresses starts with “!” sign (eg !192.168.0.0/24) and can be used to exclude hosts or networks from current Alias or Network Group Alias
Apart from the CIDR notation, one could also use a wildcard mask to match ranges of hosts or networks.
Tip
To match all servers ending at .1 in the 192.168.X.1 networks, use a wildcard definition like 192.168.0.1/0.0.255.0
Ports
Ports can be specified as a single number or a range using a colon :. For instance to add a range of 20 to 25 one would enter 20:25 in the Port(s) section.
MAC addresses
Hardware mac addresses can be specified as a (partial) hex value, such as F4:90:EA to match all addresses from
Deciso or f4:90:ea:00:00:01 to match a single item (the input is case insensitive).
The way these aliases function is approximately the same as hostnames in host type aliases, they are resolved on periodic
intervals from the arp and ndp tables.
Warning
Please be aware that hardware addresses can be spoofed (https://en.wikipedia.org/wiki/MAC_spoofing), which doesn’t make filters on them more secure than ip addresses in any way.
Note
Since mappings between addresses and mac addresses are resolved periodically the actual situation can differ, you can always check to inspect the current contents of the alias.
URL Tables
URL tables can be used to fetch a list of IP addresses from a remote server.
You can specify a Refresh frequency` to determine how often this information should be updated.
Note
The content of the file being fetched should contain one IPv[4|6] address per line, lines that start with a whitespace
, colon (,), semicolon (;), pipe (|) or hash (#) will be ignored.
URL Table in JSON format (IPs)
URL tables can be used to fetch a list of IP addresses from a remote server and parse their contents when in JSON format, similar to our standard (text based) url table.
You can use a Path expression to select data from the container, in some cases, when content is “flat” you just need a
single path reference. For example the spamhause drop list contains a json
file per row with a field cidr.
More advanced scenarios are also possible as our parser supports jq, some (simple) examples can be found below in the table below.
| Content | Path Expression | Topic | 
|---|---|---|
| https://ip-ranges.amazonaws.com/ip-ranges.json | .prefixes[] \| select(.region==”us-east-1”) \| select(.service==”EC2”) \| .ip_prefix | All ip addresses belonging to service EC2 in region us-east-1 | 
| https://api.github.com/meta | .web + .api + .git \| .[] | All of GitHubs web, api and git addresses | 
| https://endpoints.office.com/endpoints/worldwide?clientrequestid=b10c5ed1-bad1-445f-b386-b919946339a7 | .[] \| select(.serviceArea==”Exchange”) \| select(“.ips”)\| .ips \| .[]? | Exchange networks from Microsoft | 
Tip
Use https://play.jqlang.org/ to fiddle with the jq language before pasting content and path expression in an alias.
GeoIP
With GeoIP aliases you can select one or more countries or whole continents to block or allow. Use the toggle all checkbox to select all countries within the given region.
To use GeoIP, when not using our business edition, you need to configure a source in the tab, our software supports formats offered by IPinfo and MaxMind.
Note
In our experience IPinfo offers a much bigger and more detailed dataset, do make sure you increase to a higher number when using different countries from the list. As of this writing, the total size of the list is ~7 million entries.
Although you’re not obligated to use one of the defined services, the following documents explain how to use the option of your choice:
Tip
When using the Business Edition, you can leave the Url field empty so the firewall will download the IPinfo database provided
from our mirrors automatically.
Below you will find a detailed specification our software can detect and process automatically.
This format is a simple comma separated file containing the following elements in this order:
- network
- country
- country_code
- continent
- continent_code
- asn
- as_name
- as_domain
Our software only uses network and country_code, these should be mentioned in the header
