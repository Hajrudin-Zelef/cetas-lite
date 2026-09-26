---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-2.md
source_anchor: ""
source_lines: [1, 46]
sha256: db4e628fe1afa12b821326b92d72d74d83807354611d0294fdfee2734493cfe0
---

# Overview

Policy routing is the method to steer traffic matching certain criteria to a certain gateway. This can be used to force some customers or specific protocols from the servers (for example HTTP traffic) to always be routed to a certain gateway. It can even be used to steer local and overseas traffic to different gateways.

RouterOS implements several components that can be used to achieve said task:

- routing tables
- routing rules
- firewall mangle marking

# Routing Tables

A router can have multiple routing tables with its own set of routes routing the same destination to different gateways.

Tables can be seen and configured from the `/routing/table` menu.

By default, RouterOS has only the '**main**' routing table:

If a custom routing table is required, it should be defined in this menu prior to using it anywhere in the configuration.

Let's consider a basic example where we have two gateways 172.16.1.1 and 172.16.2.1 and we want to resolve 8.8.8.8 only in the routing table named '**myTable**' to the gateway 172.16.2.1:

For a user-created table to be able to resolve the destination, the main routing table should be able to resolve the destination too.

In our example, the **main** routing table should also have a route to destination 8.8.8.8 or at least a default route, since the default route is dynamically added by the DHCP for safety reasons it is better to add 8.8.8.8 also in the main table.

But configuration above is not enough, we need a method to force the traffic to actually use our newly created table. RouterOS gives you two options to choose from:

- firewall mangle - it gives more control over the criteria to be used to steer traffic, for example, per connection or per packet balancing, etc. For more info on how to use mangle marking see Firewall Marking examples.
- routing rules - a basic set of parameters that can be used to quickly steer traffic. This is the method we are going to use for our example.

It is not recommended to use both methods at the same time or you should know exactly what you are doing. If you really do need to use both mangle and routing rules in the same setup then keep in mind that mangle has higher priority, meaning if the mangle marked traffic can be resolved in the table then route rules will never see this traffic.

Routing table count is limited to 4096 unique tables.

# Routing Rules

Routing rules allow steering traffic based on basic parameters like a source address, a destination address, or in-interface as well as other parameters.

For our example, we want to select traffic with destination 8.8.8.8 and do not fall back to the **main** table:

Lets's say that we know that customer is connected to ether4 and we want only that customer to route 8.8.8.8 to a specific gateway. We can use the following rule:

If for some reason the gateway used in our table goes down, the whole lookup will fail and the destination will not be reachable. In active-backup setups we want the traffic to be able to fall back to the **main** table. To do that change the action from `lookup-only-in-table` to `lookup`.

Also, routing rules can be used as a very "basic firewall". Let's say we do not want to allow a customer connected to ether4 to be able to access the 192.168.1.0/24 network:
