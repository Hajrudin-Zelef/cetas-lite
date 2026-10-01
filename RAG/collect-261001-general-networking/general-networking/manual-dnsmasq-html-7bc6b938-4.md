---
id: collect-261001-general-networking/general-networking/manual-dnsmasq-html-7bc6b938-4
title: "manual-dnsmasq-html-7bc6b938"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-dnsmasq-html-7bc6b938.md
source_anchor: ""
source_lines: [245, 357]
sha256: b2f9807cad50573288da8f077082974697c2f4438b77e79e522826088bed9739
---

# manual-dnsmasq-html-7bc6b938

If a host receives a DHCP lease from this range, and it advertises a hostname, it will be registered under the chosen domain name.
E.g., a host named nas01 will become nas01.lan.internal. A client can query this FQDN to receive the current IP address.
Attention
If you plan to use partial IPv6 addresses in ranges with a constructor, enable the advanced mode and set Domain Type to Interface.
This will register any subnets on the chosen interface to the selected domain. This is the only way dynamic DNS registration succeeds
when the IPv6 prefix is dynamic.
| Option | Value | 
|---|---|
| Interface | GUEST | 
| Start address | 192.168.10.100 | 
| End address | 192.168.10.199 | 
| Domain | guest.internal | 
- Press Save and Apply
Tip
Creating a DHCP range will automatically send out common DHCP options to requesting clients, without explicitly configuring them.
This is an incomplete overview which highlights some default DHCP options:
| DHCP Option | Default | Description | 
|---|---|---|
| router[3] | IPv4 address of the interface that received the DHCP Request. | The default gateway the client should use. In this case the OPNsense. | 
| dns-server[6] | IPv4 address of the interface that received the DHCP Request. | The DNS server the client should use. In this case Unbound on the OPNsense. | 
| domain-name[15] | Domain set in a DHCP Range, or the default system domain if none could be matched. | The domain name the client should use, to construct short names to FQDNs in DNS lookups | 
| client fqdn[81] | A combination of client hostname and domain, the result of the DDNS registration. | The full qualified domain name the client should use. | 
Note
Only some usecases require setting these options manually, e.g., the IPv4 address of the router and dns-server in high availability setups with CARP.
Attention
If Dnsmasq does not start, check that ISC-DHCP and KEA DHCP are not active since they will block the bindable ports this DHCP server requires. It is also a good idea to check for the error message.
Now that the setup is complete, the following will happen in regards of DHCP and DNS.
- A new device (e.g. a smartphone) joins the LAN network and sends a DHCP Discover broadcast.
- Dnsmasq receives this broadcast on port 67 and responds with a DHCP offer, containing an available IP address and DHCP options for router[3] and dns-server[6].
- The device sends a DHCP request to request the available IP address, and possibly send its own hostname.
- Dnsmasq acknowledges the request.
Our smartphone now has the following IP configuration:
- IP address: 192.168.1.100
- Default Gateway: 192.168.1.1
- DNS Server: 192.168.1.1
At the same time, Dnsmasq registers the DNS hostname of the smartphone (if it exists). Since we configured the FQDN option and domain in the DHCP range, the name of the
smartphone will be: smartphone.lan.internal..
When a client queries Unbound for exactly smartphone.lan.internal., the configured query forwarding sends the request to the DNS server responsible for lan.internal.
which is our configured Dnsmasq listening on 127.0.0.1:53053. Dnsmasq responds to this query and will resolve the current A record of smartphone.lan.internal. to
192.168.1.100, sending this information to Unbound which in return sends the response back to the client that initially queried.
Tip
You can usually resolve a hostname in your network by querying for e.g. smartphone. This works because client systems
recognize that a FQDN is not used, and will therefore suffix the request with their domain name received from Dnsmasq, transforming
the query to smartphone.lan.internal..
As you can see, this is a highly integrated and simple setup which leverages just the available DHCP and DNS standards with no trickery involved.
DHCPv6 and Router Advertisements
DHCPv6 and Router Advertisements can run at the same time as DHCPv4, just specify another range.
Attention
DHCPv6 does not have a router option like DHCPv4. To push the default gateway to clients you must use Router Advertisements. This can be done with Dnsmasq, but also by a different service like .
In this example, we add a DHCPv6 range and Router Advertisements to our LAN interface. The following configuration sets stateless DHCPv6 and SLAAC. This means clients will use a SLAAC address but query additional DHCPv6 options, e.g. DNS Server.
- Go to and set:
| Option | Value | 
|---|---|
| Interface | LAN | 
| Start address | ::1000 | 
| End address | ::2000 | 
| Constructor | LAN | 
| RA Mode | slaac | 
With the mode set to slaac, clients will generate a SLAAC address and an additional DHCPv6 address (stateful DHCPv6).
If clients should only generate a SLAAC address, set the mode to ra-stateless (stateless DHCPv6).
Attention
If you use a constructor and a custom domain for the range, enable the advanced mode and set Domain Type to Interface.
This will register any subnets on the chosen interface to the selected domain. Otherwise all names fall back to the default system domain.
As final step, go to  and enable Router Advertisements.
Press Apply to activate the new configuration.
Tip
The DNS server will be sent automatically via RDNSS and DHCPv6 option. The IP address will be this firewall. If you want to change this behavior, create your own DHCPv6 options in
DHCP reservations
A DHCP reservation will always assign the same IPv4 and IPv6 addresses to a client.
For an IPv4 reservation, a DHCPv4 range should exist. If this DHCPv4 range should only serve reservations, set it to static.
For an IPv6 reservation, a DHCPv6 range must be configured which sets slaac as Router Advertisement option.
This sets the A bit so that clients can generate a SLAAC address and receive an additional DHCPv6 lease.
If a different Router Advertisement daemon is used, ensure it runs in Assisted mode.
Tip
Reservations will reserve the IP address inside a range, meaning the reserved IP will not be offered to dynamic clients.
A dynamic range like 192.168.1.100-192.168.1.199 and a reservation like 192.168.1.101 are valid and there will be no collisions.
The reservation can also be outside the dynamic range, but it is not recommended for simple setups as the dynamic dns registration with dhcp-fqdn will not work correctly.
Attention
Setting the range mode to static is not required for reservations. It is for specific usecases where a range should not serve any unknown dynamic clients.
Note
As all clients configure a tag with the receiving interface name automatically, DHCP options that are tagged with an interface will automatically match the reservations.
Here are a few examples for DHCP reservations. This assumes we already created ranges for LAN and GUEST as outlined in the previous sections.
Go to
| Option | Value | 
|---|---|
| Host | smartphone | 
| IP addresses | 192.168.1.150 | 
| Hardware addresses | aa:bb:cc:dd:ee:ff | 
- Press Save and Apply
Attention
Setting a domain in the reservation has no effect on the dynamic dns registration; it will only create a static host override.
Dnsmasq will always combine the host with a domain configured in a matching dhcp range.
This is especially important for partial IPv6 reservations, as they cannot be resolved before the dynamic dns registration has finished.
| Option | Value | 
|---|---|
| Host | smartphone | 
| IP addresses | ::1234 | 
| Client identifier | 00:03:00:01:aa:bb:cc:dd:ee:ff | 
- Press Save and Apply
Attention
A Hardware address will not work for IPv6 reservations. It must be the device unique identifier (DUID). This example uses the common DUID-LL type.
Tip
Setting a partial IPv6 address will ensure it uses the same constructor as the configured DHCPv6 ranges.
| Option | Value | 
|---|---|
| Host | smartphone | 
| IP addresses | 192.168.1.150::1234 | 
| Client identifier | 00:03:00:01:aa:bb:cc:dd:ee:ff | 
| Hardware addresses | aa:bb:cc:dd:ee:ff | 
- Press Save and Apply
Tip
This combines both IPv4 and IPv6 reservations in the same configuration item.
DHCP boot
