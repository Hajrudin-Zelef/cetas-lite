---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809-10
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809.md
source_anchor: ""
source_lines: [511, 586]
sha256: fd241b9decdd3108231a84473f7d5e4d9d77173093966e4a8c8dc32a4591d4cf
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809

- Wait for BGP—Sends stub router advertisements until BGP converges.
Note You should not save the running configuration of a router when it is configured for a graceful shutdown because the router continues to advertise a maximum metric after it is reloaded.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
SUMMARY STEPS
3. max-metric router-lsa [external-lsa [ max-metric-value ]] [include-stub] [on-startup { seconds | wait-for bgp tag }] [summary-lsa [ max-metric-value ]]
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | max-metric router-lsa [ external-lsa [ max-metric-value] ] [ include-stub ] [ on-startup { seconds \| wait-for bgp tag }] [ summary-lsa [ max-metric-value ]] Example: switch(config-router)# max-metric router-lsa | Configures OSPFv2 stub route advertisements. | 
| Step 4 | copy running-config startup-config  Example: switch(config-router)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to enable the stub router advertisements on startup for the default 600 seconds:
switch(config)# router ospf 201
Configuring the Administrative Distance of Routes
You can set the administrative distance of routes added by OSPFv2 into the RIB.
The administrative distance is a rating of the trustworthiness of a routing information source. A higher value indicates a lower trust rating. Typically, a route can be learned through more than one routing protocol. The administrative distance is used to discriminate between routes learned from more than one routing protocol. The route with the lowest administrative distance is installed in the IP routing table.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
See the guidelines and limitations for this feature in the “Guidelines and Limitations for OSPFv2” section.
SUMMARY STEPS
5. route-map map-name [ permit | deny ] [ seq ]
6. match route-type route-type
7. match ip route-source prefix-list name
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | [ no ] table-map map-name  Example: switch(config-router)# table-map foo | Configures the policy for filtering or modifying OSPFv2 routes before sending them to the RIB. You can enter up to 63 alphanumeric characters for the map name. | 
| Step 4 | exit  Example: switch(config-router)# exit switch(config)# | Exits router configuration mode. | 
| Step 5 | route-map map-name [ permit \| deny ] [ seq ]  Example: switch(config)# route-map foo permit 10 switch(config-route-map)# | Creates a route map or enters route-map configuration mode for an existing route map. Use seq to order the entries in a route map. Note The permit option enables you to set the distance. If you use the deny option, the default distance is applied. | 
| Step 6 | match route-type route-type  Example: switch(config-route-map)# match route-type external | Matches against one of the following route types:  | 
| Step 7 | match ip route-source prefix-list name  Example: switch(config-route-map)# match ip route-source prefix-list p1 | Matches the IPv4 route source address or router ID of a route to one or more IP prefix lists. Use the ip prefix-list command to create the prefix list. | 
| Step 8 | match ip address prefix-list name  Example: switch(config-route-map)# match ip address prefix-list p1 | Matches against one or more IPv4 prefix lists. Use the ip prefix-list command to create the prefix list. | 
| Step 9 | set distance value  Example: switch(config-route-map)# set distance 150 | Sets the administrative distance of routes for OSPFv2. The range is from 1 to 255. | 
| Step 10 | copy running-config startup-config  Example: switch(config-route-map)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to configure the OSPFv2 administrative distance for inter-area routes to 150, for external routes to 200, and for all prefixes in prefix list p1 to 190:
switch(config)# router ospf 201
switch(config-router)# table-map foo
switch(config)# route-map foo permit 10
switch(config-route-map)# match route-type inter-area
switch(config-route-map)# set distance 150
switch(config)# route-map foo permit 20
switch(config-route-map)# match route-type external
switch(config-route-map)# set distance 200
switch(config)# route-map foo permit 30
switch(config-route-map)# match ip route-source prefix-list p1
Modifying the Default Timers
OSPFv2 includes a number of timers that control the behavior of protocol messages and shortest path first (SPF) calculations. OSPFv2 includes the following optional timer parameters:
- LSA arrival time—Sets the minimum interval allowed between LSAs that arrive from a neighbor. LSAs that arrive faster than this time are dropped.
- Pacing LSAs—Sets the interval at which LSAs are collected into a group and refreshed, checksummed, or aged. This timer controls how frequently LSA updates occur and optimizes how many are sent in an LSA update message (see the “Flooding and LSA Group Pacing” section).
- Throttle LSAs—Sets the rate limits for generating LSAs. This timer controls how frequently LSAs are generated after a topology change occurs.
- Throttle SPF calculation—Controls how frequently the SPF calculation is run.
At the interface level, you can also control the following timers:
- Retransmit interval—Sets the estimated time between successive LSAs.
- Transmit delay—Sets the estimated time to transmit an LSA to a neighbor.
See the “Configuring Networks in OSPFv2” section for information about the hello interval and dead timer.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
SUMMARY STEPS
4. timers lsa-group-pacing seconds
5. timers throttle lsa start-time hold-interval max-time
6. timers throttle spf delay-time hold-time max-time
8. ip ospf hello-interval seconds
9. ip ospf dead-interval seconds
10. ip ospf retransmit-interval seconds
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | timers lsa-arrival msec  Example: switch(config-router)# timers lsa-arrival 2000 | Sets the LSA arrival time in milliseconds. The range is from 10 to 600000. The default is 1000 milliseconds. | 
| Step 4 | timers lsa-group-pacing seconds  Example: switch(config-router)# timers lsa-group-pacing 200 | Sets the interval in seconds for grouping LSAs. The range is from 1 to 1800. The default is 10 seconds. | 
| Step 5 | timers throttle lsa start-time hold-interval max-time  Example: switch(config-router)# timers throttle lsa 3000 | Sets the rate limit in milliseconds for generating LSAs with the following timers: start-time —The range is from 0 to 5000 milliseconds. The default value is 0 milliseconds. hold-interva l—The range is from 50 to 30,000 milliseconds. The default value is 5000 milliseconds. max-time —The range is from 50 to 30,000 milliseconds. The default value is 5000 milliseconds. | 
