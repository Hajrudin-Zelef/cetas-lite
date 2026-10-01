---
id: collect-261001-general-networking/general-networking/manual-dnsmasq-html-7bc6b938-5
title: "manual-dnsmasq-html-7bc6b938"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/manual-dnsmasq-html-7bc6b938.md
source_anchor: ""
source_lines: [358, 505]
sha256: 8ecc5ecbba5fa49bcacbcf786b73aaa821205b2b4c39d32ce9ad8df881ca3bbe
---

# manual-dnsmasq-html-7bc6b938

In a network, we have different clients that should receive different boot images depending on if they require a BIOS or EFI boot.
By using DHCP tags, we can configure this behavior by matching DHCP options and combining them with a DHCP boot directive.
Go to and create two tags:
| Option | Value | 
|---|---|
| Name | IsBIOS | 
| Option | Value | 
|---|---|
| Name | IsEFI | 
Go to
We will match the DHCP option client-arch[93] which has multiple possibilities when it comes to the client architecture.
Value 0 matches x86 BIOS and value 7 matches EFI BC (EFI x64). Choose the correct values to match your specific clients.
| Option | Value | 
|---|---|
| Type | Match | 
| Option | client-arch[93] | 
| Tag [set] | IsBIOS | 
| Value | 0 | 
| Option | Value | 
|---|---|
| Type | Match | 
| Option | client-arch[93] | 
| Tag [set] | IsEFI | 
| Value | 7 | 
Go to
Create two boot entries that serve the correct image to matching clients. We assume the requests are on LAN, though it can be left empty if these boot images should be served on any interfaces. Adjust IP addresses and filenames to fit your environment.
| Option | Description | 
|---|---|
| Interface | LAN | 
| Tag | IsBIOS | 
| Filename | undionly.kpxe | 
| Servername | 192.168.99.10 | 
| Server address | 192.168.99.10 | 
| Option | Description | 
|---|---|
| Interface | LAN | 
| Tag | IsEFI | 
| Filename | snponly.efi | 
| Servername | 192.168.99.10 | 
| Server address | 192.168.99.10 | 
Apply the new configuration, and check the PXE boot server if clients request the correct boot image files.
DHCPv4 for small HA setups
In addition to the setup described above, Dnsmasq can be a viable option in a HA setup in small and medium sized network environments.
In contrast to KEA DHCP, it does not offer lease synchronization. Each Dnsmasq instance is a separate entity.
The main tricks to make this work are the following options:
- Go to :
Set this on the current master:
| Option | Value | 
|---|---|
| DHCP reply delay | Do not set a value here, we want the master to respond first. | 
| Disable HA sync | X | 
Set this on the current backup:
| Option | Value | 
|---|---|
| DHCP reply delay | 10 (10 seconds is a good starting point) | 
| Disable HA sync | X | 
Note
This means, each DHCP Discover will be answered by the master. If the master does not respond for 10 seconds, the backup server will respond. It’s important to choose a high enough delay time, otherwise the behavior can be unpredictable in busy networks. The disabled HA sync ensures that the DHCP general settings are not synced between master and backup.
- Go to :
With LAN as example, set this on the current master:
| Option | Value | 
|---|---|
| Interface | LAN | 
| Start address | 192.168.1.100 | 
| End address | 192.168.1.199 | 
| Disable HA sync | X | 
Set this on the current backup:
| Option | Value | 
|---|---|
| Interface | LAN | 
| Start address | 192.168.1.200 | 
| End address | 192.168.1.220 | 
| Disable HA sync | X | 
Note
Now both master and backup have their own pool in the LAN network. The pool on master is larger, since it will respond to most DHCP discovers. If the master does not respond, the backup server will serve an IP address from its available pool. Since the pools do not overlap, there cannot be an IP address conflict between clients. The disabled HA sync ensures that these pools are not synchronized.
Tip
Reservations for single hosts created in can still be synchronized. They count as their own single IP address pools outside of the defined DHCP ranges. This means both servers will serve the same IP address to a host when queried. There cannot be an IP address conflict in this case. Set the MAC address of the host in the Hardware address field.
With this setup, a simple and efficient HA setup with automatic DNS registration is possible. Yet for larger scalable setups with big IP address ranges in many VLANs, KEA DHCP might be the better choice due to its robust HA synchronization options.
DHCPv6 and Router Advertisements for small HA setups
Just as with DHCPv4, the same type of configuration can be done for DHCPv6 with a few minor adjustments.
Since IPv6 uses DAD (Duplicate Address Detection), you do not need to create separate pools. SLAAC and DAD will take care of avoiding duplicates.
Special care must be taken for the Router Advertisements. Since both master and backup will send them at the same time, the current default gateway must be determined by priority and router lifetime.
- Go to :
Set this on the current master:
| Option | Value | 
|---|---|
| Interface | LAN | 
| Start address | :: | 
| Constructor | LAN | 
| RA Mode | ra-stateless | 
| RA Priority | High | 
| RA Interval | 10 | 
| RA Router Lifetime | 30 | 
| Disable HA sync | X | 
Set this on the current backup:
| Option | Value | 
|---|---|
| Interface | LAN | 
| Start address | :: | 
| Constructor | LAN | 
| RA Mode | ra-stateless | 
| RA Priority | Normal | 
| RA Interval | 10 | 
| RA Router Lifetime | 30 | 
| Disable HA sync | X | 
As final step, go to
Enable the checkbox Router Advertisements on both master and backup and apply the configuration.
Both master and backup will now advertise their link local addresses as default gateway. As long as clients receive the RA priority high packets,
they prefer the master as the current IPv6 default gateway. When the master goes offline, the RA interval is sent every 10 seconds, yet after 30 seconds
the RA router lifetime will be reached and the master will be deprecated from the clients routing table. The backup will now be installed as new
IPv6 default route.
As soon as the master comes back online, the higher RA priority will make clients shift back eventually.
Note
This whole process is not seamless, it takes some time. At least as long as the dysfunct IPv6 route is not deprecated by the clients, IPv6 will still be routed to the non-existing link local address of the offline master.
Attention
Do not set the RA Interval and RA Router Lifetime too low, as clients could potentially loose their default routes in busy networks. The bare minimum for RA Router Lifetime should be (RA Interval*3).
Dnsmasq as primary DNS resolver
This is a small complementory section how to configure Dnsmasq as the primary DNS resolver for your network combined with Unbound as recurser.
It is useful if you rely on features like dynamic IPv6 networks with PTR records registered via DHCP, or the Firewall Alias (IPset) feature.
The drawbacks are Unbound Statistics or Blocklist features based on client IP, as the client will always be 127.0.0.1.
The benefits are a less complicated configuration and less adjustments in Unbound if new networks get introduced.
- Go to and set:
| Option | Value | 
|---|---|
| Enable | X | 
| Listen Port | 53053 | 
- Go to and set:
| Option | Value | 
|---|---|
| Enable | X | 
| Listen Port | 53 | 
| Do not forward to system defined DNS servers | X (This will force Dnsmasq to only use forwarding specified in the domains tab) | 
| Do not forward private reverse lookups | X | 
- Go to and set:
| Option | Value | 
|---|---|
| Sequence | 1 | 
| Domain | * (This will match all domains) | 
| IP address | 127.0.0.1 (Unbound listens on this IP address and port) | 
| Port | 53053 | 
Apply the configuration and test DNS resolution with a client.
Firewall Alias (IPset)
Dnsmasq has a powerful feature, it can add resolved IP addresses to firewall aliases.
This is quite useful in restricted networks or to gather statistics.
As example, you provide a guest network, but users should only access example.com. With a normal firewall alias, this might be challenging,
as the domain might use multiple subdomains that serve additional content. It could also use a CDN to load balance content across different servers with dynamically
changing IP addresses per client.
With a Dnsmasq managed alias, this becomes rather simple as it will automatically add new IPv4 and IPv6 addresses as soon as they are requested by clients.
