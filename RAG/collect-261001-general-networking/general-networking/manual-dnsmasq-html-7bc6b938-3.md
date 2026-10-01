---
id: collect-261001-general-networking/general-networking/manual-dnsmasq-html-7bc6b938-3
title: "manual-dnsmasq-html-7bc6b938"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-dnsmasq-html-7bc6b938.md
source_anchor: ""
source_lines: [112, 244]
sha256: f967f70e3f00cf45d8e2fa14c8fb8b6439c417c050edc2c7555cee69edcd8e5a
---

# manual-dnsmasq-html-7bc6b938

| RA MTU | Optional MTU to send to clients via Router Advertisements. If unsure leave empty. | 
| RA Interval | Time (seconds) between Router Advertisements. | 
| RA Router Lifetime | The lifetime of the route may be changed or set to zero, which allows a router to advertise prefixes but not a route via itself. When using HA, setting a short timespan here is advised for faster IPv6 failover. A good combination could be 10 seconds RA interval and 30 seconds RA router lifetime. Going lower than that can pose issues in busy networks. | 
| Mode | Mode flags to set for this range, ‘static’ means no addresses will be automatically assigned. | 
| Lease time | Defines how long the addresses (leases) given out by the server are valid (in seconds). Set 0 for infinite; be careful as this might deplete the pool. | 
| Domain Type | Choose if the domain will only match clients in this range, or all clients in any subnets on the selected interface. If you create both IPv4 and IPv6 ranges, setting this to “Interface” on both ranges is recommended. | 
| Domain | Offer the specified domain to machines in this range. | 
| Disable HA sync | Ignore this range from being transferred or updated by HA sync. | 
| Description | You may enter a description here for your reference (not parsed). | 
| Modes | M-Bit | O-Bit | A-Bit | Default Route | DHCPv6 | SLAAC | 
|---|---|---|---|---|---|---|
| default | 1 | 1 | 0 | advertised | stateful | no | 
| ra-only | 0 | 0 | 0 | advertised | no | no | 
| slaac | 1 | 0 | 1 | advertised | both | yes | 
| ra-stateless | 0 | 1 | 1 | advertised | stateless | yes | 
This is what the RA Flags (Bits) mean:
  - M - Managed address configuration:
  - The client should use stateful DHCPv6 to obtain an IPv6 address (and implicitly O information).
  - O - Other configuration:
  - The client should use DHCPv6 to obtain other information (e.g., DNS server, Domain).
  - A - Autonomous address-configuration:
  - The client can use SLAAC to self-assign an IPv6 address based on the advertised prefix.
Tip
For other RA modes not listed here, visit the dnsmasq man page.
| Option | Description | 
|---|---|
| Type | “Set” option to send it to a client in a DHCP offer or “Match” option to dynamically tag clients that send it in the initial DHCP request. | 
| Option | DHCPv4 option to offer to the client. | 
| Option6 | DHCPv6 option to offer to the client. | 
| Interface | This adds a single interface as a tag so this DHCP option can match the interface of a DHCP range. | 
| Tag | If the optional tags are given, then this option is only sent when all the tags are matched. Can be optionally combined with an interface tag. The special address 0.0.0.0 or [::] is taken to mean “the address of the machine running dnsmasq”. When using “Match”, leave empty to match on the option only. | 
| Tag [set] | Tag to set for requests matching this range which can be used to selectively match dhcp options. | 
| Value | Value (or values) to send to the client. The special address 0.0.0.0 or [::] is taken to mean “the address of the machine running dnsmasq”. When using “Match”, leave empty to match on the option only. Send multiple values as a comma-separated list. E.g., 192.168.1.1,192.168.1.2 . | 
| Force | Always send the option, even when the client does not ask for it in the parameter request list. | 
| Description | You may enter a description here for your reference (not parsed). | 
| Option | Description | 
|---|---|
| Interface | This adds a single interface as tag so this DHCP boot option can match the interface of a DHCP range. | 
| Tag | Only offer this boot image to the clients matched by the given tag. Can be optionally combined with an interface tag. | 
| Filename | The boot image file name. | 
| Servername | The name of the server which serves the boot image. | 
| Server address | The address of the server which serves the boot image. | 
| Description | You may enter a description here for your reference (not parsed). | 
| Option | Description | 
|---|---|
| Tag | An alphanumeric label which marks a network so that DHCP options may be specified on a per-network basis. | 
Note
Interfaces set tags automatically, you do not need to set tags for them. Just select the interface in a DHCP range or DHCP option for the match to happen.
Advanced settings
To configure options that are not available in the gui one can add custom configuration files on the firewall itself.
Files can be added in /usr/local/etc/dnsmasq.conf.d/, these should use as extension .conf (e.g. custom-options.conf).
When more files are placed inside the directory, all will be included in alphabetical order.
Warning
It is the sole responsibility of the administrator which places a file in the extension directory to ensure that the configuration is valid.
Configuration examples
DHCPv4 with DNS registration
Dnsmasq can be used as a DNS forwarder. Though in our recommended setup, we will not use it as our default DNS server.
We will use Unbound as primary DNS server for our clients, and only forward some internal zones to Dnsmasq which manages the hostnames of DHCP registered leases.
This requires Dnsmasq to run with a non-standard port other than 53.
- Go to and set:
| Option | Value | 
|---|---|
| Enable | X | 
| Listen Port | 53053 | 
- Press Apply
Afterwards we can configure Unbound to forward the zones to Dnsmasq.
- Go to and set:
| Option | Value | 
|---|---|
| Enable | X | 
| Listen Port | 53 | 
- Press Apply
- Go to and create an entry for each DHCP range you plan to configure.
In our example, we configure query forwarding for 2 networks:
lan.internal - 192.168.1.0/24
guest.internal - 192.168.10.0/24
| Option | Value | 
|---|---|
| Domain | lan.internal | 
| Server IP | 127.0.0.1 | 
| Server Port | 53053 | 
- Press Save and add next
| Option | Value | 
|---|---|
| Domain | 1.168.192.in-addr.arpa | 
| Server IP | 127.0.0.1 | 
| Server Port | 53053 | 
- Press Save and Apply
Note
The first entry is for the forward lookup (A-Record), the second for the reverse lookup (PTR-Record).
Tip
If all PTR records for 192.168.0.0/16 should be handled by Dnsmasq, creating a single entry with 168.192.in-addr.arpa is enough.
| Option | Value | 
|---|---|
| Domain | guest.internal | 
| Server IP | 127.0.0.1 | 
| Server Port | 53053 | 
- Press Save and add next
| Option | Value | 
|---|---|
| Domain | 10.168.192.in-addr.arpa | 
| Server IP | 127.0.0.1 | 
| Server Port | 53053 | 
- Press Save and Apply
Note
.internal is the IANA and ICANN approved TLD (Top Level Domain) for internal use. If you instead own a TLD, e.g., example.com, you could create a zone
that is not used on the internet, e.g., lan.internal.example.com.
Now that we have the DNS infrastructure set up, we can configure DHCP.
- Go to and set:
| Option | Value | 
|---|---|
| Interface | LAN, GUEST (The network interfaces which will serve DHCP, this registers firewall rules) | 
| Do not forward to system defined DNS servers | X (Unless Domains are specified in Dnsmasq: Domains, this will disable forwarding behavior) | 
| DHCP fqdn | X | 
| DHCP default domain | internal (or leave empty to use this system’s domain) | 
| DHCP register firewall rules | X | 
Note
DHCP fqdn will do two things:
- Make sure all devices are registered in DNS with the configured domain name appended, e.g. smartphone.lan.internal .
This ensures thatsmartphone can exist in bothlan.internal andguest.internal .
- Register the DHCP domain name as local, which will make Dnsmasq authoritative for this domain, ensuring NXDOMAIN is returned
for devices querying unknown hostnames within this local domain.
- Press Apply
As next step we define the DHCP ranges for our interfaces.
- Go to and set:
| Option | Value | 
|---|---|
| Interface | LAN | 
| Start address | 192.168.1.100 | 
| End address | 192.168.1.199 | 
| Domain | lan.internal | 
- Press Save and Apply
Note
