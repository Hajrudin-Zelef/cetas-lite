---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac-4
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac.md
source_anchor: ""
source_lines: [173, 237]
sha256: 99354b05e0c1265061b4ddde9e6a3e662b8c2d23efe89159dc6b87438e3541c8
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac

| Step 12 | set dampening halflife reuse suppress max-suppress-time Example:  Device(config-route-map)# set dampening 30 1500 10000 120  | Sets BGP route dampening factors. | 
| Step 13 | set local-preference value Example:  Device(config-route-map)# set local-preference 100  | Assigns a value to a local BGP path. | 
| Step 14 | set origin {igp \| egp as \| incomplete} Example:  Device(config-route-map)#set origin igp  | Sets the BGP origin code. | 
| Step 15 | set as-path {tag \| prepend as-path-string} Example:  Device(config-route-map)# set as-path tag  | Modifies the BGP autonomous system path. | 
| Step 16 | set level {level-1 \| level-2 \| level-1-2 \| stub-area \| backbone} Example:  Device(config-route-map)# set level level-1-2  | Sets the level for routes that are advertised into the specified area of the routing domain. The stub-area and backbone are OSPF NSSA and backbone areas. | 
| Step 17 | set metric metric value Example:  Device(config-route-map)# set metric 100  | Sets the metric value to give the redistributed routes (for EIGRP only). The metric value is an integer from -294967295 to 294967295. | 
| Step 18 | set metricbandwidth delay reliability loading mtu Example:  Device(config-route-map)# set metric 10000 10 255 1 1500  | Sets the metric value to give the redistributed routes (for EIGRP only):  | 
| Step 19 | set metric-type {type-1 \| type-2} Example:  Device(config-route-map)# set metric-type type-2  | Sets the OSPF external metric type for redistributed routes. | 
| Step 20 | set metric-type internal Example:  Device(config-route-map)# set metric-type internal  | Sets the multi-exit discriminator (MED) value on prefixes advertised to external BGP neighbor to match the IGP metric of the next hop. | 
| Step 21 | set weight number Example:  Device(config-route-map)# set weight 100  | Sets the BGP weight for the routing table. The value can be from 1 to 65535. | 
| Step 22 | end Example:  Device(config-route-map)# end  | Returns to privileged EXEC mode. | 
| Step 23 | show route-map Example:  Device# show route-map  | Displays all route maps configured or only the one specified to verify configuration. | 
| Step 24 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
How to Control Route Distribution
Although each of Steps 3 through 14 in the following section is optional, you must enter at least one match route-map configuration command and one set route-map configuration command.
| Note | The keywords are the same as defined in the procedure to configure the route map for redistritbution. | 
The metrics of one routing protocol do not necessarily translate into the metrics of another. For example, the RIP metric is a hop count, and the IGRP metric is a combination of five qualities. In these situations, an artificial metric is assigned to the redistributed route. Uncontrolled exchanging of routing information between different routing protocols can create routing loops and seriously degrade network operation.
If you have not defined a default redistribution metric that replaces metric conversion, some automatic metric translations occur between routing protocols:
-  
                                    			 
                                    RIP can automatically redistribute static routes. It assigns static routes a metric of 1 (directly connected).
-  
                                    			 
                                    Any protocol can redistribute other routing protocols if a default mode is in effect.
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | router { rip \| ospf \| eigrp} Example:  Device(config)# router eigrp 10  | Enters router configuration mode. | 
| Step 3 | redistribute protocol [process-id] {level-1 \| level-1-2 \| level-2} [metric metric-value] [metric-type type-value] [match internal \| external type-value] [tag tag-value] [route-map map-tag] [weight weight] [subnets] Example:  Device(config-router)# redistribute eigrp 1   | Redistributes routes from one routing protocol to another routing protocol. If no route-maps are specified, all routes are redistributed. If the keyword route-map is specified with no map-tag , no routes are distributed. | 
| Step 4 | default-metric number Example:  Device(config-router)# default-metric 1024  | Cause the current routing protocol to use the same metric value for all redistributed routes ( RIP and OSPF). | 
| Step 5 | default-metric bandwidth delay reliability loading mtu Example:  Device(config-router)# default-metric 1000 100 250 100 1500  | Cause the EIGRP routing protocol to use the same metric value for all non-EIGRP redistributed routes. | 
| Step 6 | end Example:  Device(config-router)# end  | Returns to privileged EXEC mode. | 
| Step 7 | show route-map Example:  Device# show route-map  | Displays all route maps configured or only the one specified to verify configuration. | 
| Step 8 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Policy-Based Routing
Restrictions for Configuring PBR
- 
                                       
                                       Policy-based routing (PBR) is not supported to forward traffic into GRE tunnel. This applies to PBR applied on any interface and forwarding traffic into GRE tunnel (by means of PBR next-hop or default next-hop or set interface).
- 
                                       
                                       PBR is not supported on GRE tunnel itself (applied under the GRE tunnel itself).
Information About Policy-Based Routing
You can use policy-based routing (PBR) to configure a defined policy for traffic flows. By using PBR, you can have more control over routing by reducing the reliance on routes derived from routing protocols. PBR can specify and implement routing policies that allow or deny paths based on:
-  
                                    		  
                                    Identity of a particular end system
-  
                                    		  
                                    Application
-  
                                    		  
                                    Protocol
You can use PBR to provide equal-access and source-sensitive routing, routing based on interactive versus batch traffic, or routing based on dedicated links. For example, you could transfer stock records to a corporate office on a high-bandwidth, high-cost link for a short time while transmitting routine application data such as e-mail over a low-bandwidth, low-cost link.
With PBR, you classify traffic using access control lists (ACLs) and then make traffic go through a different path. PBR is applied to incoming packets. All packets received on an interface with PBR enabled are passed through route maps. Based on the criteria defined in the route maps, packets are forwarded (routed) to the appropriate next hop.
- Route map statement marked as
                                    		  permit is processed as follows: 
                                    		  
                                    
  - A match command can match on
                                          				length or multiple ACLs. A route map statement can contain multiple match
                                          				commands. Logical or algorithm function is performed across all the match
                                          				commands to reach a permit or deny decision. 
                                          				
