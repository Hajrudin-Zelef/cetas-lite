---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-68-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "preemption"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-68.md
source_anchor: ""
source_lines: [110, 186]
sha256: cd55d1f58be3e353a6031d60c212565b4410267e64aa9dbcd2f89757cc2236c9
---

# Summary

- If priority in advertisement packet is 0;
- When Preemption_Mode is set to yes and Priority in the ADVERTISEMENT is lower than the local Priority

After the transition to Master state node is:

- in IPv4 broadcasts gratuitous ARP request;
- in IPv6 sends an unsolicited ND Neighbor Advertisement for every associated IPv6 address.

In other cases, advertisement packets will be discarded. When the shutdown event is received, transit to Init state.

Preemption mode is ignored if the Owner router becomes available.

### Master state

When the MASTER state is set, the node functions as a forwarding router for IPv4/IPv6 addresses associated with the VR.

In IPv4 networks, the Master node responds to ARP requests for the IPv4 address associated with the VR. In IPv6 networks Master node:

- responds to ND Neighbor Solicitation message for the associated IPv6 address;
- sends ND Router Advertisements for the associated IPv6 addresses.


If the advertisement packet is received by master node:

- If priority is 0, send advertisement immediately;
- If priority in advertisement packet is greater than nodes priority then transit to the backup state;
- If priority in advertisement packet is equal to nodes priority and primary IP Address of the sender is greater than the local primary IP Address, then transit to the backup state;
- Ignore advertisement in other cases.

When the shutdown event is received, send the advertisement packet with priority=0 and transit to Init state.

## Connection tracking synchronization

Similar to different High availability features, RouterOS v7 supports VRRP connection tracking synchronization.

The VRRP connection tracking synchronization requires that RouterOS connection tracking is running. By default, connection tracking is working in `auto` mode. If VRRP devices do not contain any firewall rules, you need to manually enable connection tracking:

To sync connection tracking entries configure the device as follows:

Verify configuration in the logging section:

Connection tracking entries are synchronized only from the Master to the Backup device.

When both **sync-connection-tracking****preemption-mode**

If multiple VRRP interfaces are configured between two units and `sync-connection-tracking=yes` is required, it must be enabled only on one of the VRRP interfaces, preferably the one designated as the `group-authority`.

# Configuring VRRP

## IPv4

Setting up Virtual Router is quite easy, only two actions are required - create VRRP interface and set Virtual Routers IP address.

For example, add VRRP to ether1 and set VRs address to 192.168.1.1

Notice that only the 'interface' parameter was specified when adding VRRP. It is the only parameter required to be set manually, other parameters if not specified will be set to their defaults: `vrid=1, priority=100` and `authentication=none`.

Address on the VRRP interface must have /32 netmask if the address configured on VRRP is from the same subnet as on the router's any other interface.

Before VRRP can operate correctly correct IP address is required on ether1. In this example, it is 192.168.1.2/24.

## IPV6

To make VRRP work in IPv6 networks, several additional options must be enabled - v3 support is required and the protocol type should be set to IPv6:

Now when the VRRP interface is set, we can add a global address and enable ND advertisement:

No additional address configuration is required as it is in the IPv4 case. IPv6 uses link-local addresses to communicate between nodes.

# Parameters

VRRP interface parameters.

**Sub-menu:** `/interface vrrp`

#### Writable settings

