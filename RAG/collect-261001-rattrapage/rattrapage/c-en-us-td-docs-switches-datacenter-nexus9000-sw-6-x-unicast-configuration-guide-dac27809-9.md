---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809-9
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809.md
source_anchor: ""
source_lines: [441, 510]
sha256: 6fb9cf73a4aa6a3b4e26cf38cd7b47a3abd952d4cfee401333bac9df275937b5
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809

- Default information originate—Generates an AS External (type 5) LSA for a default route to the external autonomous system.
Note Default information originate ignores match statements in the optional route map.
Note If you redistribute static routes, Cisco NX-OS also redistributes the default static route.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
SUMMARY STEPS
3. redistribute { bgp id | direct | eigrp id | isis id | ospf id | rip id | static } route-map map-name
4. default-information originate [ always ] [ route-map map-name ]
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | redistribute { bgp id \| direct \| eigrp id \| isis id \| ospf id \| rip id \| static } route-map map-name  Example: switch(config-router)# redistribute bgp route-map FilterExternalBGP | Redistributes the selected protocol into OSPF through the configured route map. Note If you redistribute static routes, Cisco NX-OS also redistributes the default static route. | 
| Step 4 | default-information originate [ always ] [ route-map map-name]  Example: switch(config-router)# default-information-originate route-map DefaultRouteFilter | Creates a default route into this OSPF domain if the default route exists in the RIB. Use the following optional keywords:  Note This command ignores match statements in the route map. | 
| Step 5 | default-metric cost  Example: switch(config-router)# default-metric 25 | Sets the cost metric for the redistributed routes. This command does not apply to directly connected routes. Use a route map to set the default metric for directly connected routes. Note If you do not specify a metric, OSPF uses a default value of 20 when redistributing routes from all protocols except Border Gateway Protocol (BGP) routes, which use a metric of 1. | 
| Step 6 | copy running-config startup-config  Example: switch(config-router)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to redistribute the Border Gateway Protocol (BGP) into OSPF:
switch(config)# router ospf 201
switch(config-router)# redistribute bgp route-map FilterExternalBGP
Limiting the Number of Redistributed Routes
Route redistribution can add many routes to the OSPFv2 route table. You can configure a maximum limit to the number of routes accepted from external protocols. OSPFv2 provides the following options to configure redistributed route limits:
- Fixed limit—Logs a message when OSPFv2 reaches the configured maximum. OSPFv2 does not accept any more redistributed routes. You can optionally configure a threshold percentage of the maximum where OSPFv2 logs a warning when that threshold is passed.
- Warning only—Logs a warning only when OSPFv2 reaches the maximum. OSPFv2 continues to accept redistributed routes.
- Withdraw—Starts the timeout period when OSPFv2 reaches the maximum. After the timeout period, OSPFv2 requests all redistributed routes if the current number of redistributed routes is less than the maximum limit. If the current number of redistributed routes is at the maximum limit, OSPFv2 withdraws all redistributed routes. You must clear this condition before OSPFv2 accepts more redistributed routes.
- You can optionally configure the timeout period.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
SUMMARY STEPS
3. redistribute { bgp id | direct | eigrp id | isis id | ospf id | rip id | static } route-map map-name
4. redistribute maximum-prefix max [ threshold ] [ warning-only | withdraw [ num-retries timeout ]]
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | redistribute { bgp id \| direct \| eigrp id \| isis id \| ospf id \| rip id \| static } route-map map-name  Example: switch(config-router)# redistribute bgp route-map FilterExternalBGP | Redistributes the selected protocol into OSPF through the configured route map. | 
| Step 4 | redistribute maximum-prefix max [ threshold ] [ warning-only \| withdraw [ num-retries timeout ]]  Example: switch(config-router)# redistribute maximum-prefix 1000 75 warning-only | Specifies a maximum number of prefixes that OSPFv2 distributes. The range is from 1 to 65535. Optionally specifies the following:  | 
| Step 5 | show running-config ospf  Example: switch(config-router)# show running-config ospf | (Optional) Displays the OSPFv2 configuration. | 
| Step 6 | copy running-config startup-config  Example: switch(config-router)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to limit the number of redistributed routes into OSPF:
switch(config)# router ospf 201
switch(config-router)# redistribute bgp route-map FilterExternalBGP
Configuring Route Summarization
You can configure route summarization for inter-area routes by configuring an address range that is summarized. You can also configure route summarization for external, redistributed routes by configuring a summary address for those routes on an ASBR. For more information, see the “Route Summarization” section.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
SUMMARY STEPS
3. area area-id range ip-prefix/length [ no-advertise ] [ cost cost ]
4. summary-address ip-prefix/length [ no-advertise | tag tag-id ]
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | area area-id range ip-prefix/length [ no-advertise ] [ cost cost ]  Example: switch(config-router)# area 0.0.0.10 range 10.3.0.0/16 | Creates a summary address on an ABR for a range of addresses and optionally does not advertise this summary address in a Network Summary (type 3) LSA. The cost range is from 0 to 16777215. | 
| Step 4 | summary-address ip-prefix/length [ no-advertise \| tag tag]  Example: switch(config-router)# summary-address 10.5.0.0/16 tag 2 | Creates a summary address on an ASBR for a range of addresses and optionally assigns a tag for this summary address that can be used for redistribution with route maps. | 
| Step 5 | show ip ospf summary-address  Example : switch(config-router)# show ip ospf summary-address | (Optional) Displays information about OSPF summary addresses. | 
| Step 6 | copy running-config startup-config  Example: switch(config-router)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to create summary addresses between areas on an ABR:
switch(config)# router ospf 201
switch(config-router)# area 0.0.0.10 range 10.3.0.0/16
switch(config-router)# copy running-config startup-config
This example shows how to create summary addresses on an ASBR:
switch(config)# router ospf 201
switch(config-router)# summary-address 10.5.0.0/16
Configuring Stub Route Advertisements
Use stub route advertisements when you want to limit the OSPFv2 traffic through this router for a short time. For more information, see the “OSPFv2 Stub Router Advertisements” section.
Stub route advertisements can be configured with the following optional parameters:
- On startup—Sends stub route advertisements for the specified announce time.
