---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-30-2
title: "Routing Tables"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-30.md
source_anchor: ""
source_lines: [96, 176]
sha256: dba9b2d2e83e36ecedef06f98debed8f7f70fe14cd36c96a301151035e58052f
---

# Routing Tables

Peer role is now a mandatory parameter, for basic setups, you can just use ibgp, ebgp (more information on available roles can be found in the corresponding RFC draft https://datatracker.ietf.org/doc/draft-ietf-idr-bgp-open-policy/?include_text=1), keep in mind that at the moment capabilities, communities, and filtering described in the draft is not implemented.

Very basic iBGP set up to listen on the whole local network for connections:

Now you can monitor the status of all connected and disconnected peers from `/routing bgp session` menu.

Other great debugging information on all routing processes can be monitored from `/routing stats` menu

Route filtering differs a bit from ROSv6. In the BGP template, you can now specify output.filter-chain, output.filter-select, input.filter as well as several input.accept-* options.

Now input.accept-* allows filtering incoming messages directly before they are even parsed and stored in memory, that way significantly reducing memory usage. Regular input filter chain can only reject prefixes which means that it will still eat memory and will be visible in /routing route table as "not active, filtered",

A very basic example of a BGP input filter to accept prefixes from 192.168.0.0/16 subnet without modifying any attributes. For other prefixes subtract 1 from the received local pref value and set IGP metric to value from OSPF ext. Additionally, we will accept only specific prefixes from the address list to reduce memory usage

If the routing filter chain is not specified BGP will try to advertise every active route it can find in the routing table

The default action of the routing filter chain is "drop"

## Monitoring Advertisements

RouterOS v7 by default disables monitoring of the BGP output. This allows to significantly reduce resource usage on setups with large routing tables.

To be able to see output advertisements several steps should be taken:

- enable "output.keep-sent-attributes" in BGP connection configuration
- run "dump-saved-advertisements" from BGP session menu
- view saved output from "/routing/stats/pcap" menu

## Networks

Lastly, you might notice that the **`network`** menu is missing and probably wondering how to advertise your own networks. Now networks are added to the firewall address-list and referenced in the BGP configuration.

Following ROSv6 network configuration:

would translate to v7 as:

It is alsopossible to automatically create blackhole route for each BGP network:

There is more configuration to be done when adding just one network but offers simplicity when you have to deal with a large number of networks. v7 even allows specifying for each BGP connection its own set of networks.

In v7 it is not possible to turn off synchronization with IGP routes (the network will be advertised only if the corresponding IGP route is present in the routing table).

# Routing Filters

Starting from ROSv7.1beta4, the routing filter configuration is changed to a script-like configuration. The rule now can have "if .. then" syntax to set parameters or apply actions based on conditions from the "if" statement.

Multiple rules without action are stacked in a single rule and executed in order like a firewall, the reason is that the "set" parameter order is important and writing one "set"s per line, allows for an easier understanding from top to bottom on what actions were applied.

For example, match static default route and apply action accept can be written in one config rule:


For example, ROSv6 rule "/routing filter add chain=ospf_in prefix=172.16.0.0/16 prefix-length=24 protocol=static action=accept" converted to ROSv7 would be:

Another example, to match prefixes from the 172.16.0.0/16 range with prefix length equal to 24 and set BGP med and prepend values

It is also possible to match prefix length range like this


Filter rules now can be used to match or set communities,  large communities, and extended communities from the community list:

If there are a lot of community sets, that need to be applied in multiple rules, then it is possible to define community sets and use them to match or set:

Since route-target is encoded in extended community attribute to change or match RT you need to operate on extended community attribute, for example:

# RPKI

RouterOS implements an RTR client. You connect to the server which will send route validity information. This information then can be used to validate routes in route filters against a group with "rpki-validate" and further in filters "match-rpki" can be used to match the exact state.

For more info refer to the RPKI documentation.

# RIP Configuration

To start RIP, the instance should be configured. There you should select which routes will be redistributed by RIP and if it will redistribute the default route.

Then interface-template should be configured. There is no need to define networks in ROS version 7 as it was in version 6.

Now the basic configuration is completed on one router. RIP neighbor router should be configured in a similar way.

In ROS v7 the neighbors will appear only when there are routes to be sent or/and to be received.

Prefix lists from ROSv6 are deprecated, now all the filtering must be done by the routing filters.
