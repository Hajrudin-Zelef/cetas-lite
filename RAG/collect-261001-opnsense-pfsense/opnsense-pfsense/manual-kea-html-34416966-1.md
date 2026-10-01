---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/manual-kea-html-34416966-1
title: "manual-kea-html-34416966"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/manual-kea-html-34416966.md
source_anchor: ""
source_lines: [1, 79]
sha256: 2316bb18f37c0eed79e976337689d66f2b4cef3251a348e53405ed3ba5eba061
---

# manual-kea-html-34416966

KEA DHCP
Kea is the next generation of DHCP software, developed by Internet Systems Consortium (ISC).
It is considered the replacement for ISC-DHCP in larger HA enabled setups and synergizes well with radvd for HA enabled router advertisements.
Currently it is not possible to register hostnames dynamically between KEA and Unbound, only static reservations will be synchronized on an Unbound service restart.
Control Agent
The Kea Control Agent (CA) is a daemon which exposes a RESTful control interface for managing Kea servers. When building a high available dhcp setup, the control agent is a requirement for these kind of setups.
| Option | Description | 
|---|---|
| Enabled | Enable control agent | 
| Bind address | Address on which the RESTful interface should be available, usually this is localhost (127.0.0.1) | 
| Bind port | Choose an unused port for communication here. | 
Note
Although the control agent is required to use high availability peers, it does not have to listen on a non loopback address. The peer configuration by default uses the so called “Multi-Threaded Configuration (HA+MT)”, in which case it starts a separate listener for the HA communication.
DDNS Agent
The Kea DHCP DDNS (D2) server is a middleware between the DHCP servers, and authoritative DNS servers. Enabling it is a requirement if dynamic DNS updates (RFC2136) should be sent when clients are assigned an IP address in configured subnets.
| Option | Description | 
|---|---|
| Enabled | Enable DDNS server. To send updates to an authoritative nameserver, configure Dynamic DNS inside the DHCPv4 and DHCPv6 subnets. | 
| Manual config | Disable configuration file generation and manage the file (/usr/local/etc/kea/kea-dhcp-ddns.conf) manually. | 
| Bind address | Address on which the DHCP DDNS server interface should be available; usually this is localhost (127.0.0.1). | 
| Bind port | Portnumber to use for the DHCP DDNS server interface; default is 53001. | 
Kea DHCPv4/v6
This is the DHCPv4/v6 service available in KEA, which offers the following tab sheets with their corresponding settings:
| Option | Description | 
|---|---|
| Service |  | 
| Enabled | Enable DHCPv4/v6 server. | 
| Manual config | Disable configuration file generation and manage the file (/usr/local/etc/kea/kea-dhcp4.conf) or (/usr/local/etc/kea/kea-dhcp6.conf) manually. | 
| General settings |  | 
| Interfaces | Select interfaces to listen on. | 
| Valid lifetime | Defines how long the addresses (leases) given out by the server are valid (in seconds) | 
| Firewall rules | Automatically add a basic set of firewall rules to allow dhcp traffic, more fine grained controls can be offered manually when disabling this option. | 
| Socket type** (DHCPv4 only) | Socket type used for DHCP communication. | 
| Socket retries | Sometimes interfaces can be slow to come up or be unavailable temporarily. This option defines how many times KEA should retry the socket binding. | 
| Socket retry wait time | Defines the wait time in milliseconds between socket retry attempts. | 
| Decline Probation Period | Defines how long an address that has been detected as duplicate via DHCPDECLINE will be prevented to be given out to other clients. | 
| MAC sources (DHCPv6 only) | The DHCPv6 protocol does not provide any completely reliable way to retrieve hardware addresses of clients. To mitigate that issue, a number of mechanisms are available. Each of these mechanisms works in certain cases, but may not in others. Whether the mechanism works in a particular deployment is somewhat dependent on the network topology and the technologies used. Please note that this influences PD route installation, since the source MAC address of the client is required to target the link-local route. It also influences MAC based reservations. | 
| Lease Expiration |  | 
| Affinity lifetime | Defines in seconds for how long a returning client will be able to retrieve the same lease. | 
| Reclamation delay | The interval in seconds between the completion of the previous reclamation cycle and the start of the next one. | 
| Reclamation initiation | This parameter controls the server wait time in seconds between each lease reclamation procedure. | 
| Maximum reclamation time | Defines an upper limit in milliseconds to the length of time a lease reclamation procedure may take. Use “0” to disable the time limit. | 
| Maximum reclamation leases | Defines the maximum number of reclaimed leases that can be processed at one time. Use “0” to set it to unlimited. | 
| Cleanup circles | This parameter specifies how many consecutive clean-up cycles must end with remaining leases to be processed before a warning is printed. | 
| High Availability |  | 
| Enabled | Enable High availability hook, requires the Control Agent to be enabled as well. | 
| This server name | The name of this server, should match with one of the entries in the HA peers. Leave empty to use this machines hostname | 
| Max Unacked clients | This specifies the number of clients which send messages to the partner but appear to not receive any response. A higher value needs a busier environment in order to consider a member down, when set to 0, any network disruption will cause a failover to happen. | 
DHCPv4
| Option | Description | 
|---|---|
| Subnet | Subnet to use, should be large enough to hold the specified pools and reservations | 
| Description | You may enter a description here for your reference (not parsed). | 
| Pools | List of pools, one per line in range or subnet format (e.g. 192.168.0.100 - 192.168.0.200 , 192.0.2.64/26). Leave this blank if you do not want to offer dynamic leases (i.e: “Deny unknown clients”) | 
| Valid lifetime | Valid lifetime for this subnet scope. | 
| Match client-id | By default, KEA uses client-identifiers instead of MAC addresses to locate clients, disabling this option changes back to matching on MAC address which is used by most dhcp implementations. | 
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
| IPv6-only Preferred (Option 108) | The number of seconds for which the client should disable DHCPv4. The minimum value is 300 seconds. | 
| Options | Select custom DHCPv4 options that were created in the options tab. | 
| Dynamic DNS |  | 
| DNS forward zone | DNS zone where DHCP clients should be registered (e.g. “home.arpa.”). | 
| DNS reverse zone | Full reverse DNS zone receiving PTR updates (e.g. “200.10.10.in-addr.arpa.”). | 
| DNS qualifying suffix | If a DHCP client only sends a hostname in option 81, append this suffix to create an FQDN (e.g. “home.arpa.”). | 
| DNS server address | Authoritative DNS server receiving dynamic updates. | 
| DNS server port | Port of the authoritative DNS server receiving dynamic updates. Leave empty to use default (53). | 
| TSIG key name | TSIG key name used for secure DNS updates. | 
| TSIG key secret | Base64 encoded TSIG key secret. | 
| TSIG key algorithm | Algorithm used for TSIG authentication with the DNS server (e.g. hmac-sha256) | 
