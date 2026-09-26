---
id: collect-260926-mikrotik/mikrotik/configuring-bgp-filters-on-mikrotik-routeros-v7-part1-justin-wilson-j2sw
title: "configuring-bgp-filters-on-mikrotik-routeros-v7-part1-justin-wilson-j2sw"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/configuring-bgp-filters-on-mikrotik-routeros-v7-part1-justin-wilson-j2sw.md
source_anchor: ""
source_lines: [1, 85]
sha256: 6e22a86683e4ab830c214fc3be0fc67e1bfc28bcd4426f106a891cbe62e28ff2
---

# configuring-bgp-filters-on-mikrotik-routeros-v7-part1-justin-wilson-j2sw

BGP (Border Gateway Protocol) is essential for dynamic routing in modern networks. MikroTik’s RouterOS version 7 introduces significant updates to BGP functionality, making it more efficient and flexible. This guide walks you through configuring BGP filters on MikroTik RouterOS v7 to manage route propagation effectively.

#### Prerequisites

Before diving into the configuration, ensure:

- Your MikroTik router is running RouterOS v7 or later.
- BGP is properly set up and peering with your desired neighbors.
- You have a basic understanding of routing and BGP principles.

#### Step 1: Access Your Router

Log in to your MikroTik router using WinBox, WebFig, or via the terminal (SSH). For this guide, we’ll use the terminal interface for configurations.

#### Step 2: Define BGP Filters

BGP filters allow you to control which prefixes are accepted, rejected, or modified before being advertised or learned. Filters are configured under the `/routing filter` menu.

##### Create a Filter

Here’s how to create a filter that matches specific prefixes:

```
/routing filter add chain=bgp-in prefix=192.168.0.0/16 action=accept
/routing filter add chain=bgp-in prefix=10.0.0.0/8 action=reject
```
- `chain` : Name of the filter chain (e.g.,`bgp-in` or`bgp-out` ).
- `prefix` : The subnet to match.
- `action` : The action to perform (e.g.,`accept` ,`reject` , or`discard` ).

You can also match routes based on:

- **Prefix length** : Add`prefix-length=16-24` to specify a range.
- **BGP community** : Use`bgp-communities=community-name` .

#### Step 3: Apply Filters to Peers

After defining your filters, apply them to the desired BGP peer.

`/routing bgp peer set [find name="peer_name"] in-filter=bgp-in out-filter=bgp-out`
Replace `peer_name` with the name of your BGP peer.

- `in-filter` : Filter applied to routes received from the peer.
- `out-filter` : Filter applied to routes sent to the peer.

#### Step 4: Verify Configuration

Once filters are applied, verify that they’re functioning as expected.

##### Check BGP Routes

`/routing bgp route print`
This command displays all BGP routes. Use it to confirm that only allowed routes are present. If you are ullign in full route table I would suggest you use Winbox and filter for specific routes.

##### Debugging and Logs

If the configuration isn’t behaving as expected, enable debugging for BGP:

`/log print where topics~"bgp"`
Logs can help identify misconfigurations or unexpected behavior.

#### Example Scenario

Here’s a practical example of using BGP filters:

- **Requirement** : Accept only prefixes from AS 65001 and block all others.

1. Create a filter to match AS 65001:`/routing filter add chain=bgp-in bgp-as-path=^65001\$ action=accept /routing filter add chain=bgp-in action=reject`  - `bgp-as-path=^65001$` : Matches routes originating from AS 65001.
  - The second rule rejects all other routes.
2. Apply the filter to the BGP peer:`/routing bgp peer set [find name="peer_name"] in-filter=bgp-in`
3. Verify routes from the peer:`/routing bgp route print where bgp-as-path~"65001"`

#### Conclusion

By leveraging BGP filters on MikroTik RouterOS v7, you can finely control route advertisements and acceptance, ensuring your network operates securely and efficiently. Always test configurations in a lab environment before deploying them to production to avoid disruptions. Happy routing!

**j2networks family of sites**

https://j2sw.com

https://startawisp.info

https://indycolo.net

#packetsdownrange #routethelight
