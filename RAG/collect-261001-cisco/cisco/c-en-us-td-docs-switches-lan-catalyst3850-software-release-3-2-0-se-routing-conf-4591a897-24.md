---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-24
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1267, 1306]
sha256: 267e23989d0511293b48465c0cbb1ccd985c6726af4d9d3c2f4d2afb947e598b
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 5 | is-type {level-1 \|  				level-1-2 \|  				level-2-only} Example:  Device(config-router)# is-type level-2-only  | (Optional) Configures the router to act as a Level 1 (station) router, a Level 2 (area) router for multi-area routing, or both (the default): | 
| Step 6 | exit Example:  Device(config-router)# end  | Returns to global configuration mode. | 
| Step 7 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Specifies an interface to route IS-IS, and enter interface configuration mode. If the interface is not already configured as a Layer 3 interface, enter the no switchport command to put it into Layer 3 mode. | 
| Step 8 | ip router isis [area 				  tag] Example:  Device(config-if)# ip router isis tag1  | Configures an IS-IS routing process for ISO CLNS on the interface and attach an area designator to the routing process. | 
| Step 9 | clns router isis [area 				  tag] Example:  Device(config-if)# clns router isis tag1  | Enables ISO CLNS on the interface. | 
| Step 10 | ip address  				ip-address-mask Example:  Device(config-if)# ip address 10.0.0.5 255.255.255.0  | Define the IP address for the interface. An IP address is required on all interfaces in an area enabled for IS-IS if any one interface is configured for IS-IS routing. | 
| Step 11 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 12 | show isis [area tag]  				database detail Example:  Device# show isis database detail  | Verifies your entries. | 
| Step 13 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring IS-IS Global Parameters
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | clns 				  routing Example:  Device(config)# clns routing  | Enables ISO connectionless routing on the switch. | 
| Step 3 | router 				  isis Example:  Device(config)# router isis  | Specifies the IS-IS routing protocol and enters router configuration mode. | 
| Step 4 | default-information 				  originate [route-map  				map-name] Example:  Device(config-router)# default-information originate route-map map1  | (Optional) Forces a default route into the IS-IS routing domain. If you enter route-map map-name, the routing process generates the default route if the route map is satisfied. | 
| Step 5 | ignore-lsp-errors Example:  Device(config-router)# ignore-lsp-errors  | (Optional) Configures the router to ignore LSPs with internal checksum errors, instead of purging the LSPs. This command is enabled by default (corrupted LSPs are dropped). To purge the corrupted LSPs, enter the no ignore-lsp-errors router configuration command. | 
| Step 6 | area-password  				password Example:  Device(config-router)# area-password 1password  | (Optional Configures the area authentication password, which is inserted in Level 1 (station router level) LSPs. | 
| Step 7 | domain-password  				password Example:  Device(config-router)# domain-password 2password  | (Optional) Configures the routing domain authentication password, which is inserted in Level 2 (area router level) LSPs. | 
| Step 8 | summary-address  				address mask [level-1 \|  				level-1-2 \|  				level-2] Example:  Device(config-router)# summary-address 10.1.0.0 255.255.0.0 level-2  | (Optional) Creates a summary of addresses for a given level. | 
| Step 9 | set-overload-bit [on-startup {seconds \|  				wait-for-bgp}] Example:  Device(config-router)# set-overload-bit on-startup wait-for-bgp  | (Optional) Sets an overload bit (a hippity bit) to allow other routers to ignore the router in their shortest path first (SPF) calculations if the router is having problems.  | 
| Step 10 | lsp-refresh-interval  				seconds Example:  Device(config-router)# lsp-refresh-interval 1080  | (Optional) Sets an LSP refresh interval in seconds. The range is from 1 to 65535 seconds. The default is to send LSP refreshes every 900 seconds (15 minutes). | 
| Step 11 | max-lsp-lifetime  				seconds Example:  Device(config-router)# max-lsp-lifetime 1000  | (Optional) Sets the maximum time that LSP packets remain in the router database without being refreshed. The range is from 1 to 65535 seconds. The default is 1200 seconds (20 minutes). After the specified time interval, the LSP packet is deleted. | 
| Step 12 | lsp-gen-interval  				[level-1 \|  				level-2]  				lsp-max-wait  				[lsp-initial-wait 				  lsp-second-wait] Example:  Device(config-router)# lsp-gen-interval level-2 2 50 100  | (Optional) Sets the IS-IS LSP generation throttling timers:  | 
| Step 13 | spf-interval  				[level-1 \|  				level-2]  				spf-max-wait  				[spf-initial-wait 				  spf-second-wait] Example:  Device(config-router)# spf-interval level-2 5 10 20  | (Optional) Sets IS-IS shortest path first (SPF) throttling timers.  | 
| Step 14 | prc-interval  				prc-max-wait  				[prc-initial-wait 				  prc-second-wait] Example:  Device(config-router)# prc-interval 5 10 20  | (Optional) Sets IS-IS partial route computation (PRC) throttling timers.  | 
| Step 15 | log-adjacency-changes 				[all] Example:  Device(config-router)# log-adjacency-changes all  | (Optional) Sets the router to log IS-IS adjacency state changes. Enter all to include all changes generated by events that are not related to the Intermediate System-to-Intermediate System Hellos, including End System-to-Intermediate System PDUs and link state packets (LSPs). | 
| Step 16 | lsp-mtu  				size Example:  Device(config-router)# lsp mtu 1560  | (Optional) Specifies the maximum LSP packet size in bytes. The range is 128 to 4352; the default is 1497 bytes. | 
| Step 17 | partition 				  avoidance Example:  Device(config-router)# partition avoidance  | (Optional) Causes an IS-IS Level 1-2 border router to stop advertising the Level 1 area prefix into the Level 2 backbone when full connectivity is lost among the border router, all adjacent level 1 routers, and end hosts. | 
| Step 18 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 19 | show 				  clns Example:  Device# show clns  | Verifies your entries. | 
| Step 20 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Note | If any link in the network has a reduced MTU size, you must change the LSP MTU size on all routers in the network. | 
Configuring IS-IS Interface Parameters
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | interface  				interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Specifies the interface to be configured and enter interface configuration mode. If the interface is not already configured as a Layer 3 interface, enter the no switchport command to put it into Layer 3 mode. | 
| Step 3 | isis metric  				default-metric [level-1 \|  				level-2] Example:  Device(config-if)# isis metric 15  | (Optional) Configures the metric (or cost) for the specified interface. The range is from 0 to 63. The default is 10. If no level is entered, the default is to apply to both Level 1 and Level 2 routers. | 
| Step 4 | isis hello-interval 				{seconds \|  				minimal} [level-1 \|  				level-2] Example:  Device(config-if)# isis hello-interval minimal  | (Optional) Specifies the length of time between hello packets sent by the switch. By default, a value three times the hello interval seconds is advertised as the holdtime in the hello packets sent. With smaller hello intervals, topological changes are detected faster, but there is more routing traffic. | 
