---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-configuring-fortigate-ha-and-ospf-graceful-restart-ta-f401555d
title: "t5-fortigate-technical-tip-configuring-fortigate-ha-and-ospf-graceful-restart-ta-f401555d"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["fine-tuning", "parameters"]
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-configuring-fortigate-ha-and-ospf-graceful-restart-ta-f401555d.md
source_anchor: ""
source_lines: [1, 220]
sha256: f774e374d496ce9c876ca30046af84a3a82099d69be7dee16d481256883c49c6
---

# t5-fortigate-technical-tip-configuring-fortigate-ha-and-ospf-graceful-restart-ta-f401555d

**Description**


This article describes how to configure the FortiGate HA cluster and the OSPF settings (using the graceful-restart feature) so that each router (the FortiGate and its peer(s)) will keep the OSPF routes in their routing table(s) to avoid traffic interruption during an HA failover.

**Scope**


- All supported versions of FortiOS.
- FortiGate is running in NAT and HA mode.


**Solution**

**Diagram.**


**Expectations, requirements.**


This process will result in an HA cluster with one or more OSPF peers that will failover without traffic interruption. 

**Configuration.**


On a FortiGate HA cluster, the OSPF router daemon process is only running on the Primary unit. When there is an HA failover, a new OSPF process will be launched on the newly elected primary.

 Even though the FortiGate has all the routes, if the peer sees the FortiGate as unresponsive, it will remove all the routes from its routing table, and traffic will be interrupted:

 To avoid traffic interruption, the following three steps must be undertaken:

1. Check that the remote peer will not delete the routes.
2. Check that the FortiGate cluster will keep the OSPF routes in the kernel ('**get router info kernel** ' command).
3. Fine-tune timers.


1. Check that the remote peer will not delete the routes.

This can be achieved with an OSPF graceful restart. 'Graceful restart' is an OSPF capability. It is an Internet standard defined in RFC 3623. This capability needs to be configured on both peers. By default, the graceful-restart is disabled on FortiGate.

 Configuration snapshot:


**config router ospf set restart-mode graceful-restart        config area**            set router-id 30.1.1.2

**config area**

edit 0.0.0.0

next

**end**
                

**config network**

edit 1

set prefix 30.1.1.0 255.255.255.0

next

edit 2

set prefix 60.1.1.0 255.255.255.0

next

end

**config redistribute "connected"**

end

**config redistribute "static"**

end

**config redistribute "rip"**

end

**config redistribute "bgp"**

end

**config redistribute "isis"**

end


1.  **Check that the FortiGate cluster will keep the OSPF routes in the kernel (get router info kernel command).**

When the FortiGate is configured in an HA cluster, FIB is synchronized between the HA cluster members.

After an HA failover , the packets are forwarded immediately because FIB is synchronized.


The synchronized routes on the secondary will have a limited lifetime and a lower priority.

The lifetime of these routes can be configured through the 'route-ttl', 'route-wait', and 'route-hold' parameters in the system HA configuration.


**config system ha**

     set route-ttl 60

    set route-wait 60

    set route-hold 60

end


**route-ttl:**

Controls how long HA routes are kept in the FIB of a cluster unit after it has been promoted as Primary.

Increasing 'route-ttl' can help provide a seamless failover.

The route-ttl range is 5 to 3600 seconds. The default route-ttl time is 10 seconds.**route-wait:** 

The time the primary unit waits after receiving the routing table update before sending the update to the subordinate units in the cluster.

The route-wait range is 0 to 3600 seconds. The default route-wait is 0 seconds.

route-hold:

The time that the primary unit waits between sending routing table updates to subordinate units in a cluster.

The route hold range is 0 to 3600 seconds. The default route hold time is 10 seconds.


1.  **Fine-tuning timers.**

Other main timers can be tuned:**restart-period(default 120):** 

Time needed for neighbours to restart(sec).

This is the number of seconds to wait for the HELLO Message before removing the routes.

The restart-period should be less or equal to the route-ttl

If the restart-period is higher than the route-ttl, traffic interruption can be expected as kernel routes will be removed without learning new routes. 

Consider tuning these counters with the following two criteria:

- The time to detect a real OSPF peer failure.
- The maximum time allowed for a restart.


**Note:**

When graceful-restart is enabled, it will delay the time at which a real network/peer failure will be detected. This will result in downtime that can be as long as the route-ttl.

It is, therefore, important that those timers be configured to a value that suits the network requirements.**Verification**.

Verify configuration with the output of the following CLI commands :


**get router info ospf status**

Routing Process "ospf 0" with ID 30.1.1.1

 Process uptime is 1 minute

 Process bound to VRF default

 Conforms to RFC2328, and RFC1583Compatibility flag is disabled

 Supports only single TOS(TOS0) routes

 Supports opaque LSA

Supports Graceful Restart

 SPF schedule delay 5 secs, Hold time between two SPFs 10 secs

 Refresh timer 10 secs


**get router info ospf status**


Check timers in the CLI command:


**get router info ospf neighbor a.b.c.d**

OSPF process 0:

  Neighbor 20.1.1.2, interface address 20.1.1.2

     In the area 0.0.0.0 via interface wan2

     Neighbor priority is 1, State is Full, 6 state changes

     DR is 20.1.1.2, BDR is 20.1.1.1

     Options is 0x42 (*|O|-|-|-|-|E|-)

     Dead timer due in 00:00:32

     Neighbor is up for 00:03:18

**Troubleshooting**.


 Check other OSPF-related information with the following commands:


**get router info ospf neighbor****get router info ospf status****get router info ospf route****get router info routing-table database**


For more in-depth information:


**diagnose ip router ospf level info****diagnose ip router ospf all enable**

**Related articles:**
