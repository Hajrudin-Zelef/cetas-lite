---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-30-1
title: "Routing Tables"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-30.md
source_anchor: ""
source_lines: [1, 95]
sha256: eed736a60d19aecfd57de6d42b51136a150b34b2852c19ec886149b74da14f8f
---

# Routing Tables

By default, all routes are added to the "main" routing table as it was before. From a configuration point of view, the biggest differences are routing table limit increase, routing table monitoring differences, and how routes are added to specific routing tables (see next example)

v7 introduces a new menu /routing route, which shows all address family routes as well as all filtered routes with all possible route attributes. `/ip route` and `/ipv6 route` menus are used to add static routes and for simplicity show only basic route attributes.

For more in-depth information on routing see this article (IP Routing).

Another new change is that most common route print requests are processed by the routing process which significantly improves the speed compared to v6.

# Use of Routing Tables and Policy Routing

The main difference from v6 is that the routing table must be added to the `/routing table` menu before actually referencing it anywhere in the configuration.  And **fib** parameter should be specified if the routing table is intended to push routes to the  FIB.

The routing rule configuration is the same except for the menu location (instead of `/ip route rule`, now it is `/routing rule`).

Let's consider a basic example where we want to resolve 8.8.8.8 only in the routing table named myTable to the gateway 172.16.1.1:


Instead of routing rules, you could use mangle to mark packets with routing-mark, the same way as it was in ROSv6.

When you are using mangle to set routing-marks be careful not to mark "local" traffic. If your rule is marking also local traffic, then router's ip addresses will not be reachable because, by default, mangle is the first in the list and overrides local lookup:

For example basic multiwan setup:

Packet coming from LAN interface with destination address 192.168.1.1 will be marked by mangle rule and processed by the "myTable" routing table. In that table only route to reach 192.168.1.1 is via default route, which means that packet will be forwarded to the 1.1.1.1 gateway.

There are two options to fix the problem:

- exclude local from being marked by adding to mangle rule: **`!dst-address-type=local`** 
- in /routing/rule menu move first rule (action=mangle) below (action=lookup table=local)

Which approach to use depends on the complexity of the setup.

# Nexhop lookup

Consider an example from v6:

Gateway 10.0.0.1 is recursively resolved through C using the smallest referring scope (scope 20 from route B), both routes are active. Now we change both A and B at the same time:

Suddenly, applying an update to route A makes the gateway of route B inactive. This is because in v6 there is only one gateway object per address.

v7 keeps multiple gateway objects per address, one for each combination of scope and gateway check.

When `target-scope` or gateway check of a route is changed, ROS v7 ***will not affect other routes***, as it does in v6. In v7 target-scope and gateway check are properties that are internally attached to the gateway, not to the route.

# OSPF Configuration

OSPFv3 and OSPFv2 are now merged into one single menu `/routing ospf`. At the time of writing this article, there are no default instances and areas.

To start both OSPFv2 and OSPF v3 instances, first, you need to create an instance for each and then add an area to the instance.

At this point, you are ready to start OSPF on the network interface. In the case of IPv6, you add either interface on which you want to run OSPF (the same as ROSv6) or the IPv6 network. In the second case, OSPF will automatically detect the interface. Here are some interface configuration examples:

ROSv7 uses templates to match the interface against the template and apply configuration from the matched template.  OSPF menus `interface` and `neighbor` contains read-only entries purely for status monitoring.

~~All route distribution control is now done purely with routing filter select, no more redistribution knobs in the instance~~ (Since the v7.1beta7 redistribution knob is back, you still need to use routing filters to set route costs and type if necessary). This gives greater flexibility on what routes from which protocols you want to redistribute.

For example, let's say you want to redistribute only static IPv4 routes from the 192.168.0.0/16 network range.

The default action of the routing filter chain is "reject"

# BGP Configuration

There is a complete redesign of the BGP configuration compared to ROSv6. The first biggest difference is that there is no more **instance** **`peer`** configuration menus. Instead, we have **`connection`**, **`template`** and **`session`** menus.

The reason for such a structure is to strictly split parameters that are responsible for connection and parameters that are BGP protocol specific.

Let's start with the Template. It contains all BGP protocol-related configuration options. It can be used as a template for dynamic peers and apply a similar config to a group of peers. Note that this is not the same as peer groups on Cisco devices, where the group is more than just a common configuration.

By default, there is a default template that requires you to set your own AS.

Starting from v7.1beta4 template parameters are exposed in the "connection" configuration. This means that the template is not mandatory anymore, allowing for an easier basic BGP connection setup, similar to what it was in ROSv6.

Most of the parameters are similar to ROSv6 except that some are grouped in the output and input section making the config more readable and easier to understand whether the option is applied on input or output. If you are familiar with CapsMan then the syntax is the same, for example, to specify the output selection chain you set `output.filter-chain=myBgpChain`.

You can even inherit template parameters from another template, for example:

Another important aspect of the new routing configuration is the global Router ID, which sets router-id and group peers in one instance. RouterOS adds a default ID which picks instance-id from any interface's highest IP. The default BGP template by default is set to use the "default" ID.

If for any reason you need to tweak or add new instances it can be done in `/routing id` menu.

Very interesting parameters are `input.``affinity` `and` **`output.affinity`**, they allow control in which process input and output of active session will be processed:

- **alone** - input and output of each session are processed in its own process, most likely the best option when there are a lot of cores and a lot of peers
- **afi, instance, vrf, remote-as** - try to run input/output of new session in process with similar parameters
- **main** - run input/output in the main process (could potentially increase performance on single-core even possibly on multicore devices with small amount of cores)
- **input** - run output in the same process as input (can be set only for output affinity)

Now that we have parameters set for the template we can add BGP connections. A minimal set of parameters are `remote.address`, `template, connect`, `listen` and `local.role`

Connect and listen to parameters specify whether peers will try to connect and listen to a remote address or just connect or just listen. It is possible that in setups where peer uses the multi-hop connection `local.address` must be configured too (similar as it was with `update-source` in ROSv6).

It is not mandatory to specify a remote AS number. ROS v7 can determine remote ASN from an open message. You should specify the remote AS only when you want to accept a connection from that specific AS.

