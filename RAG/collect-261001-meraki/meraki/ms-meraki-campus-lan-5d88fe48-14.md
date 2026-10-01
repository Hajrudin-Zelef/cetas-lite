---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-14
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "throughput"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [485, 554]
sha256: db76c0e23eaf21a413cc9ba319318db002ad12a0850fa3540b4bb16cd41e312b
---

# ms-meraki-campus-lan-5d88fe48

The below diagram by no means should be considered as a recommended STP design but rather it is there to help you understand the interoperability considerations between the different platforms when running different STP protocols.
Physical Stacking - General
Migrating to a switch stack is an effective, flexible, and scalable solution to expand network capacity:
Benefits:
- Physical stacking will provide a high-performance and redundant access layer
- Physical stacking can also provide the network with ample bandwidth for an enterprise deployment
- Multiple uplinks can be used with cross-stack link aggregation to achieve more throughput to aggregation or core layers
- The switch stack behaves as a single device (characteristics and functionality of a single switch)
- The switch stack allows expansion of switch ports without having to manage multiple devices
- Switches can be added or removed from the switch stack without affecting the overall operation of the switch stack
General Guidance
- Create a full ring topology (i.e. stacking port 1 / switch 1 to stacking port 2 / switch 2, stacking port 1 / switch 2 to stacking port 2 / switch 3, etc)
- Finish your full ring topology by connecting stacking port 1 / switch x to stacking port 2 on switch 1)
- Use distributed uplinks across the stack such that they are equidistant (e.g. distance between uplinks is 2 hops). This will ensure that there are minimal hops across the stack for traffic to get to an uplink.
- Where applicable, use cross-stack link aggregation to increase your uplink capacity from access to distribution
Please refer to the below diagram for recommendations on stack uplinks:
For selected models, it is possible to stack different switch models together. Please refer to this document for more information on the supported platforms
In case of switch stacks, ensure that the management IP subnet does not overlap with the subnet of any configured L3 interface. Overlapping subnets on the management IP and L3 interfaces can result in packet loss when pinging or polling (via SNMP) the management IP of stack members. NOTE: This limitation does not apply to the MS390 series switches.
MS switches support one-to-one or many-to-one mirror sessions. Cross-stack port mirroring is available on Meraki stackable switches. Only one active destination port can be configured per switch/stack
Physical Stacking - All models except MS390/420/425
General Guidance
- Add the switch(s) to a dashboard network (Assuming they have already been claimed to your dashboard account)
- Power on each switch
- Connect a functional uplink to each switch such that it can access the Meraki Cloud (Please note that switches will use Management VLAN 1 by default so make sure the upstream device is configured accordingly)
- Set the firmware level for your switches from Organization > Firmware Upgrades (Consult the firmware changelog to choose the latest stable vs beta firmware)
- Wait until all your switches download firmware and come back online with the new firmware
- Power off all switches
- Disconnect uplink cables from all switches
- Connect the stacking cables to create a full ring topology
- Connect one uplink for the entire stack (Choose one of the ports used previously but only one port with one uplink for the entire stack)
- Power on each switch
- Wait for all switches to come online in dashboard and show the same firmware on the switch page
- Enable stacking on dashboard (Please note that dashboard might auto detect the stack and show it under Detected potential stacks)
- Provision the stack as required (either via Detected potential stacks or by manually selecting the switches and adding them into a stack)
- Where applicable, configure link aggregation to add more uplinks and connect the uplink cables to the designated ports on the selected switches
If required, IP addressing can be changed to different settings (e.g. Static IP address or a different management VLAN) after the stack has been properly configured and is showing online on dashboard
If the network is bound to a template, please follow the instructions here instead.
If you face problems with stacking switches, please check the common alerts here.
Adding a new switch(s) to an existing MS Switch Stack (all supported models except MS390/420/425)
- Add the switch(s) to the same dashboard network (Assuming they have already been claimed to your dashboard account)
- Power on each switch
- Connect a functional uplink to each switch such that it can access the Meraki Cloud (Please note that switches will use Management VLAN 1 by default so make sure the upstream device is configured accordingly)
- Wait until all your switches download firmware and come back online with the new firmware
- Power off all new switches
- Disconnect uplink cables from all switches
- Disconnect the stacking cable from stacking port 2 / switch 1 (keep it connected on the other end)
- Now connect the stacking cable to stacking port 2 / new switch (i.e. Have the last stack member connect to port 2 on the new switch)
- Connect the new members with stacking cables and ensure that you create a full ring topology (stacking port 1 / last switch to stacking port 2 / first switch)
- Power on the new switch(s)
- Wait for all new switches to come online in dashboard and show the same firmware on the switch page
- From Switch > Switch stacks choose your stack and click Manage members
- Provision the stack as required (by manually selecting the new switches and adding them into the existing stack)
- Where applicable, configure link aggregation to add more uplinks and connect the uplink cables to the designated ports on the selected switches
If the network is bound to a template, please follow the instructions here instead.
If you face problems with stacking switches, please check the common alerts here.
Physical Stacking - MS390
General Guidance
Do not stack more than 8 MS390 switches together. To install stacking cables; align the connector and connect the stack cable to the stack port on the switch back panel and finger-tighten the screws (clockwise direction).
- Add the switch(s) to a dashboard network (Assuming they have already been claimed to your dashboard account)
- Power on each new switch simultaneously
- Connect a functional uplink to each new switch such that it can access the Meraki Cloud (Please note that switches will use Management VLAN 1 by default so make sure the upstream device is configured accordingly)
- Set the firmware level for your switches from Organization > Firmware Upgrades (Consult the firmware changelog to choose the latest stable vs beta firmware, and make sure it supports the MS390 build)
- Wait until all your switches download firmware and come back online with the new firmware (this might take up to an hour)
- Navigate to Switch > Switch stacks
- Click Add one
- Select the switches to be added to the stack and click Create
- Power off all switches
- Disconnect uplink cables from all switches
- Connect the stacking cables to create a full ring topology (ensure that each connector is correctly aligned to the switch's stacking port it is connecting to and finger-tighten the screws in clockwise direction. Make sure the Cisco logo is on the top side of the connector) See below picture as an illustration; Green indicates correct insertion, Red indicates wrong insertion:
    
