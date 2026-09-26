---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-18
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-18.md
source_anchor: ""
source_lines: [1, 24]
sha256: 8f78ff1b17d7a9743b4a3c2b75455885717e21c9b8dbb042b5821b0eeb008b95
---

# Summary

Internet Group Management Protocol (IGMP) proxy can implement multicast routing. It is forwarding IGMP frames and is commonly used when there is no need for a more advanced protocol like PIM.

**IGMP proxy features:**

- The simplest way how to do multicast routing;
- Can be used in topologies where PIM-SM is not suitable for some reason;
- It takes slightly less resources than PIM-SM;
- Ease of configuration.

On the other hand, IGMP proxy is not well suited for complicated multicast routing setups. Compared to PIM-based solutions, IGMP proxy does not support more than one upstream interface and routing loops are not detected or avoided.

By default, IGMP proxy upstream interface will send IGMPv3 membership reports and it will detect what IGMP version the upstream device (e.g. multicast router) is using based on received queries. In case IGMPv1/v2 queries are received, the upstream port will fall back to the lower IGMP version. It will convert back to IGMPv3 when IGMPv1/v2 querier present timer (400s) expires. Downstream interfaces of IGMP proxy will only send IGMPv2 queries.

RouterOS v7 has IGMP proxy configuration available in the main **system** package. Older RouterOS versions need an additional **multicast** package installed in order to use IGMP proxy. See more details about Packages.

# Examples

To forward all multicast data coming from the ether2 interface to the downstream bridge interface, where subscribers are connected, use the configuration below. Both interfaces should have an IP address.

You may also need to configure `alternative-subnets` on the upstream interface in case the multicast sender address is in an IP subnet that is not directly reachable from the local router:

To enable `quick-leave`, use the setting below:
