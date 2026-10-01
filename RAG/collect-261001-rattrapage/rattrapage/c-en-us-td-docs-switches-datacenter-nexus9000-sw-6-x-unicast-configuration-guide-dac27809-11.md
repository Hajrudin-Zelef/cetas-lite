---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809-11
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809.md
source_anchor: ""
source_lines: [587, 662]
sha256: c5386d5e04b5607b394974287e1352d7a41e4699471b2ba8f6656c0ab8f282c2
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809

| Step 6 | timers throttle spf delay-time hold-time max-time  Example: switch(config-router)# timers throttle spf 3000 2000 5000 | Sets the SPF best-path schedule initial delay time, minimum hold time, and maximum wait time in milliseconds between SPF best-path calculations. The range is from 1 to 600000 milliseconds. The default values are a 200-ms delay time, 1000-ms hold time, and 5000-ms wait time. | 
| Step 7 | interface type slot/port  Example : switch(config)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 8 | ip ospf hello-interval seconds  Example: switch(config-if)# ip ospf hello-interval 30 | Sets the hello interval for this interface. The range is from 1 to 65535. The default is 10. | 
| Step 9 | ip ospf dead-interval seconds  Example: switch(config-if)# ip ospf dead-interval 30 | Sets the dead interval for this interface. The range is from 1 to 65535. | 
| Step 10 | ip ospf retransmit-interval seconds  Example: switch(config-if)# ip ospf retransmit-interval 30 | Sets the estimated time in seconds between LSAs transmitted from this interface. The range is from 1 to 65535. The default is 5. | 
| Step 11 | ip ospf transmit-delay seconds  Example: switch(config-if)# ip ospf transmit-delay 600 switch(config-if)# | Sets the estimated time in seconds to transmit an LSA to a neighbor. The range is from 1 to 450. The default is 1. | 
| Step 12 | show ip ospf  Example : switch(config-if)# show ip ospf | (Optional) Displays information about OSPF. | 
| Step 13 | copy running-config startup-config  Example: switch(config-if)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to control LSA flooding with the lsa-group-pacing option:
switch(config)# router ospf 201
Configuring Graceful Restart
Graceful restart is enabled by default. You can configure the following optional parameters for graceful restart in an OSPFv2 instance:
- Grace period—Configures how long neighbors should wait after a graceful restart has started before tearing down adjacencies.
- Helper mode disabled—Disables helper mode on the local OSPFv2 instance. OSPFv2 does not participate in the graceful restart of a neighbor.
- Planned graceful restart only—Configures OSPFv2 to support graceful restart only in the event of a planned restart.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
Ensure that all neighbors are configured for graceful restart with matching optional parameters set.
SUMMARY STEPS
4. (Optional) graceful-restart grace-period seconds
5. (Optional) graceful-restart helper-disable
6. (Optional) graceful-restart planned-only
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | graceful-restart  Example: switch(config-router)# graceful-restart | Enables a graceful restart. A graceful restart is enabled by default. | 
| Step 4 | graceful-restart grace-period seconds  Example: switch(config-router)# graceful-restart grace-period 120 | (Optional) Sets the grace period, in seconds. The range is from 5 to 1800. The default is 60 seconds. | 
| Step 5 | graceful-restart helper-disable  Example: switch(config-router)# graceful-restart helper-disable | (Optional) Disables helper mode. This feature is enabled by default. | 
| Step 6 | graceful-restart planned-only  Example: switch(config-router)# graceful-restart planned-only | (Optional) Configures a graceful restart for planned restarts only. | 
| Step 7 | show ip ospf instance-tag  Example : switch(config-if)# show ip ospf 201 | (Optional) Displays OSPF information. | 
| Step 8 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to enable a graceful restart if it has been disabled and set the grace period to 120 seconds:
switch(config)# router ospf 201
switch(config-router)# graceful-restart
Restarting an OSPFv2 Instance
You can restart an OSPv2 instance. This action clears all neighbors for the instance.
To restart an OSPFv2 instance and remove all associated neighbors, use the following command:
Configuring OSPFv2 with Virtualization
You can configure multiple OSPFv2 instances. You can also create multiple VRFs and use the same or multiple OSPFv2 instances in each VRF. You assign an OSPFv2 interface to a VRF.
Note Configure all other parameters for an interface after you configure the VRF for an interface. Configuring a VRF for an interface deletes all the configuration for that interface.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
SUMMARY STEPS
5. (Optional) maximum-paths paths
6. interface interface-type slot/port
8. ip-address ip-prefix/length
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | vrf context vrf-name  Example: switch(config)# vrf context RemoteOfficeVRF switch(config-vrf)# | Creates a new VRF and enters VRF configuration mode. | 
| Step 3 | router ospf instance-tag  Example: switch(config-vrf)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 4 | vrf vrf-name  Example: switch(config-router)# vrf RemoteOfficeVRF switch(config-router-vrf)# | Enters VRF configuration mode. | 
| Step 5 | maximum-paths paths  Example : switch(config-router-vrf)# maximum-paths 4 | (Optional) Configures the maximum number of equal OSPFv2 paths to a destination in the route table for this VRF. This feature is used for load balancing. | 
| Step 6 | interface interface-type slot/port  Example : switch(config-router-vrf)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 7 | vrf member vrf-name  Example: switch(config-if)# vrf member RemoteOfficeVRF | Adds this interface to a VRF. | 
| Step 8 | ip address ip-prefix/length  Example: switch(config-if)# ip address 192.0.2.1/16 | Configures an IP address for this interface. You must do this step after you assign this interface to a VRF. | 
| Step 9 | ip router ospf instance-tag area area-id  Example: switch(config-if)# ip router ospf 201 area 0 | Assigns this interface to the OSPFv2 instance and area configured. | 
| Step 10 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to create a VRF and add an interface to the VRF:
switch(config)# vrf context NewVRF
switch(config)# router ospf 201
switch(config)# interface ethernet 1/2
switch(config-if)# vrf member NewVRF
switch(config-if)# ip address 192.0.2.1/16
Verifying the OSPFv2 Configuration
To display the OSPFv2 configuration, perform one of the following tasks:
|  |  | 
|---|---|
| show ip ospf [instance-tag] [vrf vrf-name] | Displays information about one or more OSPF routing instances. The output includes the following area-level counts:  | 
| show ip ospf border-routers [ vrf { vrf-name \| all \| default \| management }] | Displays the OSPFv2 border router configuration. | 
| show ip ospf database [ vrf { vrf-name \| all \| default \| management }] | Displays the OSPFv2 link-state database summary. | 
| show ip ospf interface number [ vrf { vrf-name \| all \| default \| management }] | Displays OSPFv2-related interface information. | 
| show ip ospf lsa-content-changed-list neighbor-id interface - type number [ vrf { vrf-name \| all \| default \| management }] | Displays the OSPFv2 LSAs that have changed. | 
