---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-15
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [555, 626]
sha256: de0fb02704bd6e1a89a316e6e34c52b2f6859b2151999ed50c274ff3b2a23ab8
---

# ms-meraki-campus-lan-5d88fe48

- Connect one uplink for the entire stack (Choose one of the ports used previously but only one port with one uplink for the entire stack)
- Power on each switch
- Wait for all switches to come online in dashboard and show the same firmware on the switch page
- Enable stacking on dashboard (Please note that dashboard might auto detect the stack and show it under Detected potential stacks)
- Provision the stack as required (either via Detected potential stacks or by manually selecting the switches and adding them into a stack)
- Where applicable, configure link aggregation to add more uplinks and connect the uplink cables to the designated ports on the selected switches
Please note that all MS390 switches in a stack will show the same management IP address as there is only one control plane running on the primary switch. It is recommended to configure the same IP address on all switches to ensure that traffic uses the same IP during failover scenarios
If required, IP addressing can be changed to different settings (e.g. Static IP address or a different management VLAN) on all stack members after the stack has been properly configured and is showing online on dashboard. Again, it is recommended to configure the same IP address on all switches to ensure that traffic uses the same IP during failover scenarios
If a member needs to be removed from a stack, please amend to its unique IP address details before removing it from the stack.
Rebooting a member from dashboard (or by power recycle) will reboot all members in a stack.
Factory resetting a member will reboot all members in a stack
If you have already configured settings in your dashboard network with port settings etc, please ensure that the switch/stack has a maximum of 1000 VLANs. For example, If you have an existing stack with each port set to Native VLAN 1, 1-1000 and the new member ports are set to native VLAN 1; allowed VLANs: 1,2001-2500 then your total number of VLAN in the stack will be 1000(1-1000)+500(2001-2500) = 1500. Dashboard will not allow the new member to be added to the stack and will show an error
If the network is bound to a template, please follow the instructions here instead.
If you face problems with stacking switches, please check the common alerts here.
MS390 Stack IP Address Provisioning Sequence for Best Results:
- Claim your MS390s into a dashboard network (do not create a stack)
- Set the firmware to 11.31+
- Connect an uplink to each switch (members un-stacked)
- Ensure that the stacking cables are not connected to any member
- Power on switches (members un-stacked)
- Have DHCP available on native VLAN 1
- Wait for firmware to be loaded and configuration to be synced
- Power off switches
- Disconnect all uplinks from all switches
- Connect stacking cables to all members to form a ring topology
- Connect one uplink to one member (only one link for the stack)
- Power on switches and wait for them to come online on dashboard
- Create a stack on dashboard by adding all members
- Wait for the stack ports to show online on all members in dashboard
- Observe the IP address used on the stack members (should be the same for all members)
- Click on the IP address of each switch and change settings from DHCP to Static. Configure the IP address that is used for the stack for each switch member
- Configure Link aggregation as needed and add more uplinks accordingly
- Make sure to abide to the maximum VLAN count as described in the above section when you provision your MS390 stack/switches
Please note that the Primary switch owns the Management IP and will resolve ARP requests to its own MAC address
Adding a new MS390 switch(s) to an existing MS390 Switch Stack
- Add the switch(s) to the same dashboard network (Assuming they have already been claimed to your dashboard account)
- Power on each new switch
- Connect a functional uplink to each new switch such that it can access the Meraki Cloud (Please note that switches will use Management VLAN 1 by default so make sure the upstream device is configured accordingly)
- Wait until all your switches download firmware and come back online with the new firmware
- From Switch > Switch stacks choose your stack and click Manage members
- Provision the stack as required (by manually selecting the new switches and adding them into the existing stack)
- Power off all new switches
- Disconnect uplink cables from all switches
- Disconnect the stacking cable from stacking port 2 / switch 1 (keep it connected on the other end)
- Now connect the stacking cable to stacking port 2 / new switch (i.e. Have the last stack member connect to port 2 on the new switch)
- Connect the new members with stacking cables and ensure that you create a full ring topology (stacking port 1 / last switch to stacking port 2 / first switch)
- Power on the new switch(s)
- Wait for all new switches to come online in dashboard and show the same firmware on the switch page
- Where applicable, configure link aggregation to add more uplinks and connect the uplink cables to the designated ports on the selected switches
If you have already configured settings in your dashboard network with port settings etc, please ensure that the switch/stack has a maximum of 1000 VLANs. For example, If you have an existing stack with each port set to Native VLAN 1, 1-1000 and the new member ports are set to native VLAN 1; allowed VLANs: 1,2001-2500 then your total number of VLAN in the stack will be 1000(1-1000)+500(2001-2500) = 1500. Dashboard will not allow the new member to be added to the stack and will show an error
If the network is bound to a template, please follow the instructions here instead.
If you face problems with stacking switches, please check the common alerts here.
StackPower for MS390
StackPower is an innovative feature that aggregates all the available power in a stack of switches and manages it as one common power pool for the entire stack. StackPower feature is introduced for the first time ever in Meraki Switching portfolio with the MS390s.
By pooling & distributing power across MS390s using a series of StackPower cables, StackPower provides simple and resilient power distribution across the stack. Below is the back panel of MS390 depicting the location of StackPower ports.
Guidance and steps to deploy StackPower:
- StackPower is only supported on MS390s with MS15+
- Do not add more than 4 x MS390 switches in power-stack
- If need be, split your MS390 switches into two power-stack units within a single Data stack (For instance, if you have a total of 5 MS390 switches in a data-stack, you can configure 3 switches in one power-stack setup and the rest 2 switches in another power-stack setup as shown below)
- Connect the end of the cable with a green band to either StackPower port on the first switch
- Align the connector correctly, and insert it into a StackPower port on the switch rear panel.
- Connect the end of the cable with the yellow band to another switch
- Hand-tighten the captive screws to secure the StackPower cable connectors in place.
- StackPower feature doesn't need any dashboard configuration but is automatically enabled when the cables are installed
If after connecting the cables you do not see the power-stack in dashboard, please contact Meraki support for further troubleshooting
Physical Stacking - MS420/425
General Guidance
Please note that 10 Gb/s is the minimum speed required to support flexible stacking.
Please use identical ports on both ends for stacking ports (e.g. both 10Gbps SFP+ or 40Gbps QSFP)
- Add the switch(s) to a dashboard network (Assuming they have already been claimed to your dashboard account)
- Connect an uplink to all your switch(s) such that it can access the Meraki Cloud (Please note that switches will use Management VLAN 1 by default so make sure the upstream device is configured accordingly)
- Please ensure that your uplink port is different from the intended stacking ports
