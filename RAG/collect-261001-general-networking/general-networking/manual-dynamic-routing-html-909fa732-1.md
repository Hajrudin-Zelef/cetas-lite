---
id: collect-261001-general-networking/general-networking/manual-dynamic-routing-html-909fa732-1
title: "manual-dynamic-routing-html-909fa732"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "cost", "datacenter"]
source: docs/RAG/collect-261001-general-networking/manual-dynamic-routing-html-909fa732.md
source_anchor: ""
source_lines: [1, 75]
sha256: b988d3dbe81cc19d5ada1f6eb54822883af319d850660f9765fdbd421569b0bd
---

# manual-dynamic-routing-html-909fa732

Dynamic Routing (FRR)
Dynamic Routing (using routing protocols) is supported via an external plugin. Routing protocols support your network equipment in finding the best available path for your packets. We use Free Range Routing (FRR) to implement the various available protocols for dynamic routing.
These routing protocols are used to:
- Improve fault tolerance (if a connection breaks, a new route will be found if possible)
- Simplify administration (you have to add fewer routes manually)
It is not advisable to use dynamic routing in the following scenarios:
- When your network is small (it would be simpler to use static routes)
- If you are working in a highly isolated environment, where you have to be in control of every route in your network
Routing Protocols supported by the plugin include:
- RIPv1 and RIPv2
- OSPFv2 and v3
- BGPv4
Warning
Not all routing protocols will work in any setup because they may have to be direct neighbors. Consider the limitations of a routing protocol before using it.
Warning
It’s strongly advised to increase the kern.ipc.maxsockbuf value via Tunables. Go to and check if there is already a tunable for maxsockbuf and set it to 33554432 if it is lower. Otherwise add a new one with name above and the specified value.
Warning
Disabling a running routing daemon can be dangerous as it can lead to an inaccessible machine. If you want to disable a running routing daemon, make sure you do not lose routes which are required by your connection to this machine (for example when using SSH).
Installation
Go to  and select os-frr from the available plugins.
General setup
To use one or more of the protocols included, the plugin must be enabled in . Without any other service enabled this makes sure the zebra service is being configured, which is the coordinating master service which handles generic features such as logging and access to kernel routing.
Tip
By default logging should be enabled, which sends messages to the local logging and offers remote logging over syslog. Always make sure to choose a sensible log level (default is Notifications) and check the log in
Note
The plugins supports seamless reloads since version os-frr-1.4.3. When using older versions, there is no configuration reloading. This means there might be a temporary loss of service when saving settings. Normally this is only a small glitch, but in high traffic areas it might be something to take under consideration when performing maintenance.
Dynamic routing and high availability
In enterprise networks there is often a need to protect services against all sorts of failures. Dynamic routing helps to always provide a valid path for packets to travel. These nodes themselved might need to be configured more resilient to prevent single points of failures on the edges of your network.
In OPNsense high availability and failover is organised around carp, which makes it a logical choice to combine both technologies here as well.
There are different strategies ranging from disabling the daemon when in carp mode, to more fine grained control of route propagation when a machine is in backup mode.
Note
Unicast CARP is available to use the protocol across router boundaries. This can enable the use of CARP on WAN interfaces peering with eBGP neighbors if they are not connected to the same switch.
CARP failover mode
The most simple mode available. When a node becomes Backup it will stop the FRR services. When it returns to Master it will start the FRR services.
Note
Due to the nature of this option, it cannot be combined with other available CARP options.
OSPF[6]: CARP demote
This option registers a status monitor on top of the FRR logging feed to detect changes in link status. If OSPF cannot find its neighbors, it will make this machine less attractive by increasing the demotion factor.
The feature is inspired by OpenBSD’s handling of CARP demotion in ospfd (https://man.openbsd.org/ospfd.conf.5) and can be enabled
using the CARP demote checkbox in .
Note
Since the relevant neighbor negotiation messages are only being logged when the log level (in ) is configured to debug, the log will be more chatty when using this feature. When using a lower log level the status monitor is not expected to catch any relevant events.
OSPF[6]: Influence interface cost based on CARP status
FRR does not natively support interaction with CARP status as the variant in OpenBSD does (carp note in “depend on” keyword https://man.openbsd.org/ospfd.conf.5), this is where our next option comes into play.
Using the interface settings of an OSPF interface you can choose to adjust costs for that interface based on the CARP status of the selected virtual address. Go to and choose an interface, here you will find the following options that influence behaviour:
- Depend on (carp): 
  - Select a virtual address that this interface relies on. When this target is not in MASTER mode, the selected interface is considered demoted
- Cost (when demoted): 
  - Adjust the cost to this value when going to demoted state, usually one would use a high value here to prefer other routes first
- Cost: 
  - The standard cost, when provided will be used when in normal conditions. If it’s left blank FRR defaults will be used, which it will also rollback to when going back to master mode.
Dynamic Routing Protocols
For more detailed information, check out the FRR documentation.
| Options | Description | 
|---|---|
| Enable | This will activate the routing service. Without enabling it globally, none of the individual services will run. | 
| Profile | Control FRR’s default profile: traditional reflects defaults adhering mostly to IETF standards or common practices in wide-area internet routing. datacenter reflects a single administrative domain with intradomain links using aggressive timers. | 
| Enable CARP Failover | This will activate the routing service only on the master device. The backup device will stop the service completely. | 
| Enable SNMP AgentX Support | This will activate support for Net-SNMP AgentX. | 
| Enable logging | Sends logs to the OPNsense integrated syslog-ng service. | 
| Log Level | This is the detail level of the log. A higher level means more data is logged. | 
| Firewall Rules | Enable automatically created firewall rules, when additional policies are needed, disable this and define your own custom policies in the Firewall section. | 
| Description | Rule autogenerated by Firewall Rules option | 
|---|---|
| Pass OSPF | Allows OSPF traffic from configured network IP address ranges to multicast 224.0.0.0/24 (direction: in). | 
| Pass OSPF UNICAST | Allows OSPF unicast traffic from configured network IP address ranges to the local interface (direction: in). | 
| Pass OSPF | Allows OSPF traffic from the local interface to multicast 224.0.0.0/24 (direction: out). | 
| Pass OSPF UNICAST | Allows OSPF unicast traffic from the local interface to configured network IP address ranges (direction: out). | 
| Pass OSPF6 MULTICAST | Allows OSPFv3 traffic from link-local IPv6 (fe80::0/10) to multicast address ff02::5/128 (direction: in). | 
| Pass OSPF6 MULTICAST DR | Allows OSPFv3 traffic from link-local IPv6 (fe80::0/10) to multicast address ff02::6/128 (direction: in). | 
| Pass OSPF6 UNICAST | Allows OSPFv3 unicast traffic from link-local IPv6 (fe80::0/10) to the local interface (direction: in). | 
| Pass OSPF6 MULTICAST | Allows OSPFv3 traffic from the local interface to multicast address ff02::5/128 (direction: out). | 
| Pass OSPF6 MULTICAST DR | Allows OSPFv3 traffic from the local interface to multicast address ff02::6/128 (direction: out). | 
| Pass OSPF6 UNICAST | Allows OSPFv3 unicast traffic from the local interface to link-local IPv6 (fe80::0/10) (direction: out). | 
Attention
