---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad-2
title: "docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad.md
source_anchor: ""
source_lines: [64, 137]
sha256: c650d26195c74bef477d3b49e8402177b827c7bc5f7384cfb7c33c5749b1961e
---

# docs-network-security-tutorials-how-to-configure-opnsense-nat-4928d4ad

| Option | Description | 
|---|---|
| Disabled | Check this option to disable the rule without removing it. | 
| No RDR (NOT) | Enabling this option prevents traffic matching this rule from being redirected and a redirect rule is not created. Hint: this option is rarely used; don't use it unless you're sure you know what you're doing. | 
| Interface | Which interface the rule should apply to. The majority of the time, this will be WAN. | 
| TCP/IP version | IPv4, IPv6 or both. | 
| Protocol | In typical scenarios, this will be TCP | 
| Source | Where the traffic comes from. Click `Advanced` to see the other source settings | 
| Source / Invert | Invert match in `Source` field. | 
| Source port range | When applicable, the source port on which we should match. This is almost always random and almost never equals the destination port range (and should almost always be 'any'). | 
| Destination / Invert | Invert match in `Destination` field. | 
| Destination | Where the traffic is going | 
| Destination port range | Service port(s) the traffic is using. For this mapping, specify the port or port range for the packet's destination when using the TCP or UDP protocols. | 
| Redirect target IP | Where to redirect the traffic to. Enter the internal IP address of the server to which the ports will be mapped. | 
| Redirect target port | Which port to use (when using TCP and/or UDP). Enter the port number for the machine with the IP address you entered above. In the case of a port range, specify the range's starting port (the end port will be calculated automatically). | 
| Pool Options | This option is explained in the previous section. The default is to use Round robin. Only Round Robin types are compatible with Host Aliases. Subnets of any type can be used. **Round Robin:** Iterates over the translation addresses.**Random:** Chooses an address at random from the translation address pool**Source Hash:** Determines the translation address by hashing the source address, ensuring that the redirection address is always the same for a given source.**Bitmask:** Uses the subnet mask while keeping the last portion the same; 172.16.10.50 → x.x.x.50.**Sticky Address:** When using the Random or Round Robin pool types, the Sticky Address option ensures that a specific source address is always mapped to the same translation address. | 
| Description | A description to easily find the rule in the overview. | 
| Set local tag | You can mark a packet matching this rule and use this mark to match on other NAT/filter rules. | 
| Match local tag | Check for a tag set by another rule. | 
| No XMLRPC sync | Prevent this rule from being synced to a backup host. (Checking this on the backup host has no effect.) | 
| NAT reflection | This option is explained in the previous section. Leave this on the default unless you have a good reason not to. | 
| Filter rule association | Associate this with a regular firewall rule. | 

## Configure One-to-one NAT

One-to-one NAT, as the name suggests, will translate two IP addresses one-to-one rather than one-to-many, as is more common.

To configure the One-to-One NAT in OPNsense you may navigate to `Firewall` → `NAT` → `One-to-One`. An overview of 1:1 NAT rules can be found here.

**Figure 2.** *One-to-One NAT configuration in OPNsense*

To add new One-to-One NAT rules, you may click the `+` button in the upper right corner.

The following fields are available when adding a 1:1 mapping rule:

| Option | Description | 
|---|---|
| Disabled | Check this option to disable the rule without removing it. | 
| Interface | Which interface the rule should apply to. The majority of the time, this will be WAN. | 
| Type | BINAT (default) or NAT. | 
| External network | Enter the starting address of the external subnet for the 1:1 mapping or network. If no subnet mask is provided, the subnet mask from the internal address below will be applied to this IP address. This is the address or network to/from which traffic will be translated. | 
| Protocol | In typical scenarios, this will be TCP | 
| Source | Enter the internal subnet for the 1:1 mapping. The subnet size specified for the source will be applied to the external subnet, when none is provided. | 
| Source / Invert | Invert match in `Source` field. | 
| Destination / Invert | Invert match in `Destination` field. | 
| Destination | The destination network packages should match, when used to map external networks, this is usually any | 
| Description | A description to easily find the rule in the overview. | 
| NAT reflection | This option is explained in the previous section. Leave this on the default unless you have a good reason not to. | 

## Configure Outbound NAT (SNAT)

Outbound NAT is also known as `Source NAT` or `SNAT`. When a client on an internal network sends an outbound request, the gateway must change the source IP to the gateway's external IP, because the outside server will be unable to respond otherwise.

If you only have one external IP address, you should leave the Outbound NAT options set to automatic. If you have multiple IP addresses, however, you may want to change the settings and add some custom rules.

To configure the Outbound NAT in OPNsense you may navigate to `Firewall` → `NAT` → `Outbound` . An overview of outbound rules can be found here.

**Figure 3.** *Outbound NAT configuration in OPNsense*

The following modes are available for outbound NAT configuration in OPNsense:

| Outbound NAT Mode | Description | 
|---|---|
| Automatic outbound NAT rule generation | The default and is good for most cases. | 
| Manual outbound NAT rule generation | No automatic rules are generated. Outbound NAT rules are created manually. | 
| Hybrid outbound NAT rule generation | Automatic rules are added, but manual rules can also be added. | 
| Disable outbound NAT rule generation | Disables outbound NAT. This is used for transparent bridges, for example. | 

To add new Outbound NAT rules, you may select either the `Manual outbound NAT rule generation` or `Hybrid outbound NAT rule generation` option and then click `Save` button.

New rules can be added, by clicking the `+` button in the upper right corner.

The following fields are available when adding an outbound rule:

