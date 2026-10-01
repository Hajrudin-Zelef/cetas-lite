---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-18
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [932, 972]
sha256: 677f70f1d1446384dbea68276e0c9c696758b89c9cdbad3c407d70433c4c1cb2
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | ip 				  routing Example:  Device(config)# ip routing  | Enables IP routing. | 
| Step 3 | router bgp  				autonomous-system Example:  Device(config)# router bgp 45000  | Enables a BGP routing process, assign it an AS number, and enter router configuration mode. The AS number can be from 1 to 65535, with 64512 to 65535 designated as private autonomous numbers. | 
| Step 4 | network  				network-number [mask  				network-mask] [route-map  				route-map-name] Example:  Device(config)# network 10.108.0.0  | Configures a network as local to this AS, and enter it in the BGP table. | 
| Step 5 | neighbor {ip-address \|  				peer-group-name}  				remote-as  				number Example:  Device(config)# neighbor 10.108.1.2 remote-as 65200  | Adds an entry to the BGP neighbor table specifying that the neighbor identified by the IP address belongs to the specified AS. For EBGP, neighbors are usually directly connected, and the IP address is the address of the interface at the other end of the connection. For IBGP, the IP address can be the address of any of the router interfaces. | 
| Step 6 | neighbor {ip-address \|  				peer-group-name}  				remove-private-as Example:  Device(config)# neighbor 172.16.2.33 remove-private-as  | (Optional) Removes private AS numbers from the AS-path in outbound routing updates. | 
| Step 7 | synchronization Example:  Device(config)# synchronization  | (Optional) Enables synchronization between BGP and an IGP. | 
| Step 8 | auto-summary Example:  Device(config)# auto-summary  | (Optional) Enables automatic network summarization. When a subnet is redistributed from an IGP into BGP, only the network route is inserted into the BGP table. | 
| Step 9 | bgp 				  graceful-restart Example:  Device(config)# bgp graceful-start  | (Optional) Enables NSF awareness on switch. By default, NSF awareness is disabled. | 
| Step 10 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 11 | show ip bgp 				  network  				network-number Example:  Device# show ip bgp network 10.108.0.0  | Verifies the configuration. | 
| Step 12 | show ip bgp 				  neighbor Example:  Device# show ip bgp neighbor  | Verifies that NSF awareness (Graceful Restart) is enabled on the neighbor. If NSF awareness is enabled on the switch and the neighbor, this message appears: Graceful Restart Capability: advertised and received If NSF awareness is enabled on the switch, but not on the neighbor, this message appears: Graceful Restart Capability: advertised | 
| Step 13 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Managing Routing Policy Changes
To learn if a BGP peer supports the route refresh capability and to reset the BGP session:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | show ip bgp 				  neighbors Example:  Device# show ip bgp neighbors  | Displays whether a neighbor supports the route refresh capability. When supported, this message appears for the router: Received route refresh capability from peer. | 
| Step 2 | clear ip bgp {* \|  				address \|  				peer-group-name} Example:  Device# clear ip bgp *  | Resets the routing table on the specified connection. | 
| Step 3 | clear ip bgp {* \|  				address \|  				peer-group-name}  				soft out Example:  Device# clear ip bgp * soft out  | (Optional) Performs an outbound soft reset to reset the inbound routing table on the specified connection. Use this command if route refresh is supported. | 
| Step 4 | show ip 				  bgp Example:  Device# show ip bgp  | Verifies the reset by checking information about the routing table and about BGP neighbors. | 
| Step 5 | show ip bgp 				  neighbors Example:  Device# show ip bgp neighbors  | Verifies the reset by checking information about the routing table and about BGP neighbors. | 
Configuring BGP Decision Attributes
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router bgp  				autonomous-system Example:  Device(config)# router bgp 4500  | Enables a BGP routing process, assign it an AS number, and enter router configuration mode. | 
| Step 3 | bgp best-path 				  as-path ignore Example:  Device(config-router)# bgp bestpath as-path ignore  | (Optional) Configures the router to ignore AS path length in selecting a route. | 
| Step 4 | neighbor {ip-address \|  				peer-group-name} next-hop-self Example:  Device(config-router)# neighbor 10.108.1.1 next-hop-self  | (Optional) Disables next-hop processing on BGP updates to a neighbor by entering a specific IP address to be used instead of the next-hop address. | 
| Step 5 | neighbor {ip-address \|  				peer-group-name}  				weight  				weight Example:  Device(config-router)# neighbor 172.16.12.1 weight 50  | (Optional) Assign a weight to a neighbor connection. Acceptable values are from 0 to 65535; the largest weight is the preferred route. Routes learned through another BGP peer have a default weight of 0; routes sourced by the local router have a default weight of 32768. | 
| Step 6 | default-metric  				number Example:  Device(config-router)# default-metric 300  | (Optional) Sets a MED metric to set preferred paths to external neighbors. All routes without a MED will also be set to this value. The range is 1 to 4294967295. The lowest value is the most desirable. | 
| Step 7 | bgp bestpath med 				  missing-as-worst Example:  Device(config-router)# bgp bestpath med missing-as-worst  | (Optional) Configures the switch to consider a missing MED as having a value of infinity, making the path without a MED value the least desirable path. | 
| Step 8 | bgp 				  always-compare med Example:  Device(config-router)# bgp always-compare-med  | (Optional) Configures the switch to compare MEDs for paths from neighbors in different autonomous systems. By default, MED comparison is only done among paths in the same AS. | 
| Step 9 | bgp bestpath med 				  confed Example:  Device(config-router)# bgp bestpath med confed  | (Optional) Configures the switch to consider the MED in choosing a path from among those advertised by different subautonomous systems within a confederation. | 
| Step 10 | bgp deterministic 				  med Example:  Device(config-router)# bgp deterministic med  | (Optional) Configures the switch to consider the MED variable when choosing among routes advertised by different peers in the same AS. | 
| Step 11 | bgp default 				  local-preference  				value Example:  Device(config-router)# bgp default local-preference 200  | (Optional) Change the default local preference value. The range is 0 to 4294967295; the default value is 100. The highest local preference value is preferred. | 
| Step 12 | maximum-paths  				number Example:  Device(config-router)# maximum-paths 8  | (Optional) Configures the number of paths to be added to the IP routing table. The default is to only enter the best path in the routing table. The range is from 1 to 16. Having multiple paths allows load-balancing among the paths. (Although the switch software allows a maximum of 32 equal-cost routes, the switch hardware will never use more than 16 paths per route.) | 
| Step 13 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 14 | show ip 				  bgp Example:  Device# show ip bgp  | Verifies the reset by checking information about the routing table and about BGP neighbors. | 
| Step 15 | show ip bgp 				  neighbors Example:  Device# show ip bgp neighbors  | Verifies the reset by checking information about the routing table and about BGP neighbors. | 
| Step 16 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
