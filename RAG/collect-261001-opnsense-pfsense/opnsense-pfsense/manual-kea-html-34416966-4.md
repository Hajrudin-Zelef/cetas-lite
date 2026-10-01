---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/manual-kea-html-34416966-4
title: "manual-kea-html-34416966"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-opnsense-pfsense/manual-kea-html-34416966.md
source_anchor: ""
source_lines: [317, 415]
sha256: eac45f759e3c298b6e15e167d7ac9a2dfaef7563592021ab863b033a99558e43
---

# manual-kea-html-34416966

For reverse zone updates enable the advanced mode inside a subnet. Add your DNS reverse zone to the existing forward configuration. Please note that reverse zone updates will be sent to the same DNS server as the forward zone updates.
Some clients might send client specific flags to avoid reverse zone updates. You can override that behavior with Override no update and Override client update.
Prefix Delegation (IA_PD)
Kea supports prefix delegation with static or dynamic prefixes. A prefix delegation is most commonly used for router behind router setups, yet also in client implementations that run their own VMs.
Route Installation
Whenever an IA_PD lease is acknowledged, a route targeting the link-local address of the requesting DHCPv6 client will be automatically installed.
Since lease files are synchronized in high availability mode, the routes will also be installed and cleaned up on both peers.
Note
If the MAC address for a client route installation is not found, take a look at the MAC sources option in the general DHCPv6 settings. It influences how client
MAC addresses are constructed per default. The current default ipv6-link-local will construct the MAC out of an EUI-64 link-local address.
This should work for most clients, yet if they use random link-local addresses, duid would be the next best option.
Static Prefix
As an example setup, we will use unique local addresses (ULA) to lease an IA_NA address (/128 IPv6 address) and a IA_PD prefix (/56 IPv6 prefix) to a requesting client.
Prefix: fd80::/48
- Go to and follow through these tabs:
| Option | Value | 
|---|---|
| Service |  | 
| Enabled | X | 
| General settings |  | 
| Interfaces | LAN | 
| Firewall rules | X | 
For the IA_NA address pool, we take the first /52 prefix (fd80::/52) of the available /48 prefix (fd80::/48)
| Option | Value | 
|---|---|
| Subnet | fd80::/48 | 
| Pools | fd80::100 - fd80::199 (/52 will be auto calculated via the pool) | 
For the IA_PD pool, we take the second /52 prefix (fd80:0:0:1000::/52), and lease up to 16 prefixes (fd80:0:0:1000::/56 - fd80:0:0:10F0::/56) to clients.
| Option | Value | 
|---|---|
| Subnet | fd80::/48 | 
| Prefix | fd80:0:0:1000:: | 
| Prefix length | 52 | 
| Delegated length | 56 | 
After applying the configuration, clients will receive an IA_NA address (e.g., fd80::100/128) and an IA_PD prefix (e.g., fd80:0:0:1000::/56).
Dynamic Prefix
As an example setup, our provider has provided us a prefix via DHCPv6 on our WAN interface.
Prefix: 2001:db8:1234::/56
We will use Identity association mode to carve out a prefix on LAN that is big enough to host a PD pool.
- Go to and set the following configuration:
To reserve a prefix range, the combination of the hexadecimal value Assign prefix ID and the decimal length value Reserved prefix range is used. On our LAN interface, we start with an assigned prefix ID of 0, which marks the first /64 network available. We reserve a /60 prefix for KEA’s subnet on this interface, so we count up 16x /64 networks via the reserved prefix range.
LAN will now reserve the hexadecimal prefix IDs 0-F.
| Option | Value | 
|---|---|
| IPv6 Configuration Type | Identity association | 
| Parent interface | WAN | 
| Assign prefix ID | 0 | 
| Reserved prefix range | 16 | 
In this example we want to reserve a /61 prefix, so our decimal reserved prefix range is 8. Since our LAN interface already reserves the hexadecimal prefix IDs 0-F, for OPT1 we start at the hexadecimal prefix ID 10.
OPT1 will now reserve the hexadecimal prefix IDs 10-17.
| Option | Value | 
|---|---|
| IPv6 Configuration Type | Identity association | 
| Parent interface | WAN | 
| Assign prefix ID | 10 | 
| Reserved prefix range | 8 | 
In this example we want to reserve a /62 prefix, so our decimal reserved prefix range is 4. Since our LAN interface reserves the hexadecimal prefix IDs 0-F, and our OPT1 interface the hexadecimal prefix IDs 10-17, for OPT1 we start at the hexadecimal prefix ID 18.
OPT2 will now reserve the hexadecimal prefix IDs 18-1B.
| Option | Value | 
|---|---|
| IPv6 Configuration Type | Identity association | 
| Parent interface | WAN | 
| Assign prefix ID | 18 | 
| Reserved prefix range | 4 | 
Attention
If you change these ranges later or remove interfaces, ensure you also update the KEA configuration. If an interface is removed, also remove the dynamic subnet from KEA. If prefix ID ranges are changed, ensure the delegated length in a PD pool is updated with a new value that fits into that network. If not followed, KEA will emit log messages with details and may fail to start.
- Next, go to and configure the dynamic PD pool for LAN:
| Option | Value | 
|---|---|
| Service |  | 
| Enabled | X | 
| General settings |  | 
| Interfaces | LAN | 
| Firewall rules | X | 
The subnet pool is automatically calculated. Since our example prefix ID range is from 0-F, the calculated subnet size will be 2001:db8:1234::/60.
This subnet will be automatically split into two subnets:
the first subnet 2001:db8:1234::/61 will host the IA_NA pool 2001:db8:1234::/64
the second subnet 2001:db8:1234:8::/61 will host the IA_PD pool.
| Option | Value | 
|---|---|
| Interface | LAN | 
| Dynamic Prefix | X | 
| Auto collect option data | X (optional, if you also want to send a dynamic DNS server) | 
For the IA_PD pool, the automatically calculated IA_PD prefix of the subnet is used. In our example that is 2001:db8:1234:8::/61.
This is the range which can be delegated to other routers. We can set the delegated length to control how many prefixes can be leased from
this pool. In our case we need 2 delegated prefixes, so we set a delegated length of /62.
| Option | Value | 
|---|---|
| Subnet | LAN | 
| Delegated length | 62 | 
Note
By splitting your ISP provided prefix smartly, each of your internal networks can have dynamic prefix delegation ranges.
After applying the configuration, clients will receive an IA_NA address (e.g., 2001:db8:1234::100/128) and an IA_PD prefix (e.g., 2001:db8:1234:8::/62).
Attention
Using HA in combination with dynamic prefix delegation is not recommended. When using a DHCPv6 provided ISP prefix, both HA peers would likely get different prefixes from the ISP, which would cause problems with the HA setup since the KEA configurations would differ between peers. For an HA setup, using a static IPv6 prefix is a requirement to ensure a single routing identity.
Leases DHCPv4/v6
This page offers an overview of the (non static) leases being offered by KEA DHCPv4/v6.
Tip
There are action buttons to quickly register and find reservations.
