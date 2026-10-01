---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-21
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1096, 1150]
sha256: b24d4fb418c289d970c3963f25e763e781cacc077e4a21aedce68fa89b2c0c8c
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 4 | aggregate-address  				address mask  				as-set Example:  Device(config-router)# aggregate-address 10.0.0.0 255.0.0.0 as-set  | (Optional) Generates AS set path information. This command creates an aggregate entry following the same rules as the previous command, but the advertised path will be an AS_SET consisting of all elements contained in all paths. Do not use this keyword when aggregating many paths because this route must be continually withdrawn and updated. | 
| Step 5 | aggregate-address  				address-mask  				summary-only Example:  Device(config-router)# aggregate-address 10.0.0.0 255.0.0.0 summary-only  | (Optional) Advertises summary addresses only. | 
| Step 6 | aggregate-address  				address mask  				suppress-map  				map-name Example:  Device(config-router)# aggregate-address 10.0.0.0 255.0.0.0 suppress-map map1  | (Optional) Suppresses selected, more specific routes. | 
| Step 7 | aggregate-address  				address mask  				advertise-map  				map-name Example:  Device(config-router)# aggregate-address 10.0.0.0 255.0.0.0 advertise-map map2  | (Optional) Generates an aggregate based on conditions specified by the route map. | 
| Step 8 | aggregate-address  				address mask  				attribute-map  				map-name Example:  Device(config-router)# aggregate-address 10.0.0.0 255.0.0.0 attribute-map map3  | (Optional) Generates an aggregate with attributes specified in the route map. | 
| Step 9 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 10 | show ip bgp neighbors [advertised-routes] Example:  Device# show ip bgp neighbors  | Verifies the configuration. | 
| Step 11 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring Routing Domain Confederations
You must specify a confederation identifier that acts as the autonomous system number for the group of autonomous systems.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router bgp  				autonomous-system Example:  Device(config)# router bgp 100  | Enters BGP router configuration mode. | 
| Step 3 | bgp confederation 				  identifier  				autonomous-system Example:  Device(config)# bgp confederation identifier 50007  | Configures a BGP confederation identifier. | 
| Step 4 | bgp confederation peers  				autonomous-system [autonomous-system  				...] Example:  Device(config)# bgp confederation peers 51000 51001 51002  | Specifies the autonomous systems that belong to the confederation and that will be treated as special EBGP peers. | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show ip bgp 				  neighbor Example:  Device# show ip bgp neighbor  | Verifies the configuration. | 
| Step 7 | show ip bgp 				  network Example:  Device# show ip bgp network  | Verifies the configuration. | 
| Step 8 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring BGP Route Reflectors
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router bgp  				autonomous-system Example:  Device(config)# router bgp 101  | Enters BGP router configuration mode. | 
| Step 3 | neighbor {ip-address \|  				peer-group-name}  				route-reflector-client Example:  Device(config-router)# neighbor 172.16.70.24 route-reflector-client  | Configures the local router as a BGP route reflector and the specified neighbor as a client. | 
| Step 4 | bgp 				  cluster-id  				cluster-id Example:  Device(config-router)# bgp cluster-id 10.0.1.2  | (Optional) Configures the cluster ID if the cluster has more than one route reflector. | 
| Step 5 | no bgp 				  client-to-client reflection Example:  Device(config-router)# no bgp client-to-client reflection  | (Optional) Disables client-to-client route reflection. By default, the routes from a route reflector client are reflected to other clients. However, if the clients are fully meshed, the route reflector does not need to reflect routes to clients. | 
| Step 6 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 7 | show ip 				  bgp Example:  Device# show ip bgp  | Verifies the configuration. Displays the originator ID and the cluster-list attributes. | 
| Step 8 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring Route Dampening
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router bgp  				autonomous-system Example:  Device(config)# router bgp 100  | Enters BGP router configuration mode. | 
| Step 3 | bgp 				  dampening Example:  Device(config-router)# bgp dampening  | Enables BGP route dampening. | 
| Step 4 | bgp dampening  				half-life reuse suppress 				  max-suppress [route-map  				map] Example:  Device(config-router)# bgp dampening 30 1500 10000 120  | (Optional) Changes the default values of route dampening factors. | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show ip bgp flap-statistics [{regexp  				regexp} \| {filter-list  				list} \| {address mask [longer-prefix]}] Example:  Device# show ip bgp flap-statistics  | (Optional) Monitors the flaps of all paths that are flapping. The statistics are deleted when the route is not suppressed and is stable. | 
| Step 7 | show ip bgp 				  dampened-paths Example:  Device# show pi bgp dampened-paths  | (Optional) Displays the dampened routes, including the time remaining before they are suppressed. | 
| Step 8 | clear ip bgp flap-statistics [{regexp  				regexp} \| {filter-list  				list} \| {address mask [longer-prefix]} Example:  Device# clear ip bgp flap-statistics  | (Optional) Clears BGP flap statistics to make it less likely that a route will be dampened. | 
| Step 9 | clear ip bgp 				  dampening Example:  Device# clear ip bgp dampening  | (Optional) Clears route dampening information, and unsuppress the suppressed routes. | 
| Step 10 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Monitoring and Maintaining BGP
You can remove all contents of a particular cache, table, or database. This might be necessary when the contents of the particular structure have become or are suspected to be invalid.
You can display specific statistics, such as the contents of BGP routing tables, caches, and databases. You can use the information to get resource utilization and solve network problems. You can also display information about node reachability and discover the routing path your device’s packets are taking through the network.
The table given below lists the privileged EXEC commands for clearing and displaying BGP. For explanations of the display fields, see the Cisco IOS IP Command Reference, Volume 2 of 3: Routing Protocols, Release 12.4.
| Table 11  IP BGP Clear and Show 		  Commands |  | 
|---|---|
| clear ip bgp address | Resets a particular BGP connection. | 
| clear ip bgp * | Resets all BGP connections. | 
| clear ip bgp peer-group tag | Removes all members of a BGP peer group. | 
| show ip bgp prefix | Displays peer groups and peers not in peer groups to which the prefix has been advertised. Also displays prefix attributes such as the next hop and the local prefix. | 
| show ip bgp cidr-only | Displays all BGP routes that contain subnet and supernet network masks. | 
