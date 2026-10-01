---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-16
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [627, 655]
sha256: 83ba1961b1461284ba026738be93fccdaf98243016a785475ee69de8005d1316
---

# ms-meraki-campus-lan-5d88fe48

- Set the firmware level for your switches from Organization > Firmware Upgrades (Consult the firmware changelog to choose the latest stable vs beta firmware)
- Power on each switch
- Wait until all your switches download firmware and come back online with the new firmware
- Configure the designated stacking port with the stacking enabled
- Connect the stacking cables to create a full ring topology
- Connect one uplink (or link aggregate) for the entire stack and remove all other uplinks
- Enable stacking on dashboard (Please note that dashboard might auto detect the stack and show it under Detected potential stacks)
- Provision the stack as required (either via Detected potential stacks or by manually selecting the switches and adding them into a stack)
- Where applicable, configure link aggregation to add more uplinks and connect the uplink cables to the designated ports on the selected switches
Converting a link aggregate to a stacking port is not a supported configuration and may result in unexpected behavior.
If the network is bound to a template, please follow the instructions here instead.
If you face problems with stacking switches, please check the common alerts here.
Physical Stacking - Replacing a Stack Member
Replacing a stack member can be useful in one of these occasions:
- A failed switch that is RMA'd and needs to be replaced with a new one (like for like)
- A switch that is being migrated to another switch (e.g. larger switch, PoE enabled, etc)
General Guidance
- Power off the stack member to be replaced
- Claim the new/replacement switch in the inventory
- Add the switch to the network containing the stack
- Edit the name of the switch if required (For instance, to resemble the old switch e.g. SW-SFO-#5-02)
- Power on the switch that is replacing the old one
- Connect a functional uplink to one of the ports on the switch
- Wait for the switch to come online and update its firmware to the one configured on your network (Refer to Organization > Firmware Upgrades and check the Switch details page)
- Navigate to Switch > Switch stacks
- Select the existing stack
- Navigate to the failed switch that is RMA'd and needs to be replaced with a new one (like for like)
- Follow one of the following options:
    
