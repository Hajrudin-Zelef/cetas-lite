---
id: collect-261001-general-networking/general-networking/manual-aliases-html-a3e5f567-3
title: "manual-aliases-html-a3e5f567"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["advisory", "cyber"]
source: docs/RAG/collect-261001-general-networking/manual-aliases-html-a3e5f567.md
source_anchor: ""
source_lines: [181, 223]
sha256: 6cf7b50d8cbada1f9d386938b7b8a7e524b58bf70bf851c2ea6c32c32232b840
---

# manual-aliases-html-a3e5f567

When performing migrations, sometimes its easier to change multiple items at once in a text editor. This feature can easily be used to facilitate that, with limiting risk of a broken configuration (since items are validated equally as single item input would do).
Add new entries using our API
The endpoints from the alias_util can easily be used to push new entries into an alias (or remove existing ones). In case of an external alias these items won’t be persistent over reboots, which can be practical in some use-cases (large frequent changing lists for example).
The document “Use the API” contains the steps needed to create an api key and secret, next you can just call the same endpoint the user interface would.
Below you see how to add 10.0.0.2 to an alias named MyAlias using an insecure connection (self-signed cert) on
the host opnsense.firewall with curl. The verbose option provides more details about the data exchanged between the
two machines.
curl \
  --header "Content-Type: application/json" \
  --basic \
  --user "key:secret" \
  --request POST \
  --insecure \
  --verbose \
  --data  '{"address":"10.0.0.2"}' \
  https://opnsense.firewall/api/firewall/alias_util/add/MyAlias
Note
Adding aliases using /api/firewall/alias_util/add/ is only supported for Host, Network and External type aliases
Exclusions
Pf firewall tables support exceptions (or exclusion) of addresses. This feature can be used in one Alias or in combined (Network group type) Aliases. See (https://www.freebsd.org/doc/handbook/firewalls-pf.html 30.3.2.4).
Nesting
For host and network alias types nesting is possibility, this can simplify management a lot since single items can be named properly and grouped into sections for administration.
For example, we define 4 servers among 2 critical using different rulesets:
- server_a {10.0.1.1}
- server_b {10.0.1.2}
- server_c {10.0.1.100}
- server_d {10.0.1.200}
- critical_servers {server_a , server_b}
- other_servers {server_c , server_d}
- servers { critical_servers , other_servers}.
The alias servers will contain all 4 addresses after configuration.
There is also a possibility to combine different Aliases with Aliases, consisting of exclusions. For example, there is Alias “FireHOL” that use extensive externl drop-list and two Aliases that contains subnet and hosts exclusions. It is possible to create Network group (combined) Alias (“FireHOL_with_exclusions”):
- FireHOL {https://raw.githubusercontent.com/firehol/blocklist-ipsets/master/firehol_level1.netset}
- subnets_exclusions {!127.0.0.0/8, !0.0.0.0/8}
- hosts_exclusions {!8.8.8.8}
- FireHOL_with_exclusions {FireHOL, subnets_exclusions, hosts_exclusions}
FireHOL_with_exclusions Alias will contain all records from FireHOL Alias excluding addresses from exclusions Aliases.
It’s always good to check if an address is included in the Alias via
Spamhaus
The Spamhaus Don’t Route Or Peer Lists DROP (Don’t Route Or Peer) and DROPv6 are advisory “drop all traffic” lists, consisting of netblocks that are “hijacked” or leased by professional spam or cyber-crime operations (used for dissemination of malware, trojan downloaders, botnet controllers). The DROP and DROPv6 lists are a tiny subset of the SBL, designed for use by firewalls and routing equipment to filter out the malicious traffic from these netblocks.
Source : https://www.spamhaus.org/drop/
- Downloads
To setup the DROP and DROPv6 lists in combination with the firewall rules, read: Configure Spamhaus DROP
