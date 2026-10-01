---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/manual-kea-html-34416966-2
title: "manual-kea-html-34416966"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/manual-kea-html-34416966.md
source_anchor: ""
source_lines: [80, 173]
sha256: e1f7bcff822ac6a974241dee84bb92a2f801f65886edffce0a3fbf24f57c753d
---

# manual-kea-html-34416966

| Override no update | Ignores the client’s wishes for no DDNS updates to be performed. | 
| Override client update | Ignores the client’s delegation requests. Causes Kea to perform Dynamic DNS updates even though the client indicated its intention to perform the updates itself. | 
| Update on renew | Instructs the server to always update the DNS information when a lease is renewed, even if its DNS information has not changed. This allows Kea to self-heal if it was previously unable to add DNS entries or they were somehow lost by the DNS server. May impact performance, especially for servers with numerous clients that renew often. | 
| Conflict resolution mode | Controls how DDNS conflicts with DHCID records are handled. The default enforces client ownership via DHCID. | 
DHCPv6
| Option | Description | 
|---|---|
| Subnet | Subnet to use, should be large enough to hold the specified pools and reservations | 
| Interface | Select which interface this subnet belongs to | 
| Dynamic Prefix | Use the identity association prefix allocated to this interface and generate subnet and pools automatically. DHCP options that are not auto collected are unaffected by prefix changes and remain static. | 
| Allocator | Select allocator method to use when offering leases to clients. | 
| PD Allocator | Select allocator method to use when offering prefix delegations to clients | 
| Description | You may enter a description here for your reference (not parsed). | 
| Pools | List of pools, one per line in range or subnet format (e.g. 2001:db8:1::-2001:db8:1::100, 2001:db8:1::/80). Leave this blank if you do not want to offer dynamic leases (i.e: “Deny unknown clients”) | 
| Valid lifetime | Valid lifetime for this subnet scope. | 
| DHCP option data |  | 
| Auto collect option data | Automatically update option data for relevant attributes such as dns servers when applying settings from the gui. When using a dynamic prefix in a subnet, this will set the correct primary IP address automatically. | 
| DNS servers | DNS servers to offer to the clients | 
| Domain search | The domain search list to offer to the client | 
| Options | Select custom DHCPv6 options that were created in the options tab. | 
| Dynamic DNS |  | 
| DNS forward zone | DNS zone where DHCP clients should be registered (e.g. “home.arpa.”). | 
| DNS reverse zone | Full reverse DNS zone receiving PTR updates (e.g. “8.b.d.0.1.0.0.2.ip6.arpa.”). This will not be dynamically adjusted if the subnet is configured with a dynamic prefix. | 
| DNS qualifying suffix | If a DHCP client only sends a hostname in option 81, append this suffix to create an FQDN (e.g. “home.arpa.”). | 
| DNS server address | Authoritative DNS server receiving dynamic updates. | 
| DNS server port | Port of the authoritative DNS server receiving dynamic updates. Leave empty to use default (53). | 
| TSIG key name | TSIG key name used for secure DNS updates. | 
| TSIG key secret | Base64 encoded TSIG key secret. | 
| TSIG key algorithm | Algorithm used for TSIG authentication with the DNS server (e.g. hmac-sha256) | 
| Override no update | Ignores the client’s wishes for no DDNS updates to be performed. | 
| Override client update | Ignores the client’s delegation requests. Causes Kea to perform Dynamic DNS updates even though the client indicated its intention to perform the updates itself. | 
| Update on renew | Instructs the server to always update the DNS information when a lease is renewed, even if its DNS information has not changed. This allows Kea to self-heal if it was previously unable to add DNS entries or they were somehow lost by the DNS server. May impact performance, especially for servers with numerous clients that renew often. | 
| Conflict resolution mode | Controls how DDNS conflicts with DHCID records are handled. The default enforces client ownership via DHCID. | 
| Option | Description | 
|---|---|
| Subnet | Subnet to use, should be large enough to hold the specified prefix. | 
| Prefix | The prefix that will be used as prefix delegation pool. | 
| Prefix length | The length of the prefix for the prefix delegation pool. | 
| Delegated length | The length of each delegated prefix offered via the prefix delegation pool. | 
| Description | You may enter a description here for your reference (not parsed). | 
DHCPv4
| Option | Description | 
|---|---|
| Subnet | Subnet this reservation belongs to | 
| IP address | IP address to offer to the client | 
| MAC address | MAC address of the client in question | 
| Client ID | ID of the client in question. Per default this is preferred over MAC addresses. Disable “Match client-id” in the subnet to skip the Client ID. | 
| Hostname | Offer a hostname to the client | 
| Description | You may enter a description here for your reference (not parsed). | 
| DHCP option data |  | 
| Auto collect option data | Automatically update option data for relevant attributes as routers, dns servers and ntp servers when applying settings from the gui. | 
| Routers (gateway) | Default gateways to offer to the clients | 
| Static routes | Static routes that the client should install in its routing cache, defined as dest-ip1,router-ip1,dest-ip2,router-ip2 | 
| DNS servers | DNS servers to offer to the clients | 
| Domain name | The domain name to offer to the client, set to this firewall’s domain name when left empty | 
| Domain search | The domain search list to offer to the client | 
| NTP servers | Specifies a list of IP addresses indicating NTP (RFC 5905) servers available to the client. | 
| Time servers | Specifies a list of RFC 868 time servers available to the client. | 
| Next server | Next server IP address | 
| TFTP server | TFTP server address or FQDN | 
| TFTP bootfile name | Boot filename to request | 
| Options | Select custom DHCPv4 options that were created in the options tab. | 
DHCPv6
| Option | Description | 
|---|---|
| Subnet | Subnet this reservation belongs to | 
| IP address | IP address to offer to the client | 
| MAC address | MAC address of the client in question | 
| DUID | DUID of the client in question | 
| Hostname | Offer a hostname to the client | 
| Domain search | The domain search list to offer to the client | 
| Options | Select custom DHCPv6 options that were created in the options tab. | 
| Description | You may enter a description here for your reference (not parsed). | 
| Option | Description | 
|---|---|
| Description | You must enter a description here. It is used to reference this option inside reservations and subnets. | 
| Match DHCP option |  | 
| Match Code | The server will only send the option defined in “Set DHCP option” if a client first sends the option defined in “Match DHCP option”. Leave empty to always send the option. | 
| Match Encoding | Encoding used to evaluate the match condition. “Hex” supports all encapsulated and structured options generically via payload in hexadecimal byte pairs. | 
| Match Data | Data to match against the selected DHCP option. | 
| Set DHCP option |  | 
| Set Code | DHCP option to offer to the client. | 
| Set Encoding | Choose the encoding type. “Hex” supports all encapsulated and structured options generically via payload in hexadecimal byte pairs. | 
| Set Data | Payload to send to a client. | 
| Force | Always send the option, also when the client does not ask for it in the parameter request list. | 
| Option | Description | 
|---|---|
| Name | Peer name, there should be one entry matching this machines “This server name” | 
| Role | This peers role | 
| Url | This specifies the URL of our server instance, which should use a different port than the control agent. For example http://my-host:8001/ | 
Note
Define HA peers for this cluster. All nodes should contain the exact same definitions (usually two hosts, a primary and a standby host)
Configuration examples
DHCPv4 for medium/large HA setups
