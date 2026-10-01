---
id: collect-261001-general-networking/general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa-3
title: "switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa.md
source_anchor: ""
source_lines: [119, 195]
sha256: 02a1234bff4e1b50f84b16fc6900df834114bde450b17f47838576079f3866f7
---

# switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa

Note: Due to the management IP behavior described above, accessing the Local Status Page from any member will direct to the local status page of the active switch, regardless of which stack member you are physically connected to.
- 
    Based on the number of members in a stack and switch ports being used boot time can vary.
- 
    Rebooting a stack member from the dashboard using the "Reboot device" tool will reboot only that individual stack member. The same applies when performing a factory reset on a stack member. (This does not apply to MS390's or Catalyst switches, which will reboot the entire stack).
- 
    Physically power cycling switch stack members may cause other stack members to be unreachable for a short time, while the management plane re-initializes.
Note: If an active member fails, Catalyst switches on CS firmware may briefly pause traffic to reinitialize the Meraki management container on the new active member.
Adding a new member to the stack
- Add the new switch into a dashboard network of the existing switch stack. For that, you may follow Adding and Removing Devices from Dashboard Networks article.
- Connect the new switch to an uplink to bring it online and ensure it checks in with the Meraki Dashboard.
- 
    Upgrade the switch to the same firmware build as that running on the switch stack, using the Firmware Upgrade Manager under Organization > Monitor > Firmware upgrades,
- 
    Before adding the new member to an existing stack make sure the total number of VLANs is limited to 1000. E.g., If an existing stack uses 1000 VLANs (e.g., VLANs 1-1000) and the new member switch ports are configured with an additional 500 unique VLANs (e.g., VLANs 2001-2500), the total number of VLANs in the combined stack would be 1500. This exceeds the limit. The dashboard will not allow the new member to be added to the stack and will show the following error:
- 
    Navigate to Switching > Monitor > Switch stacks and select the existing stack you want to add the switch to.
- 
    Check that the switches in the existing stack have all fetched the new configuration. To verify this, navigate to Switching > Monitor > Switches and select a switch in the stack. Look for 'Configuration status' in the column on the left of the switch details page and check if the status reads 'Up to date'.
- 
    In Switching > Monitor > Switch stacks > Manage members add the new switch to the existing stack.
NOTE: If adding a switch with Layer 3 configurations to a Layer 3 switch stack, there are two options:
- Clone the configuration from the added member switch onto the switch stack.
- Ignore the existing layer 3 configuration on the added member.
8. Power off the new switch, physically stack the new switch to the existing stack in a ring fashion and power it on.
It is highly recommended to keep the entire switch stack powered on during the addition of a new member to prevent the new member from assuming the active role.
Configuring StackWise Virtual
Prerequisites for StackWise Virtual
- Both switches in the Cisco StackWise Virtual pair must be directly connected to each other.
- Both switches in the Cisco StackWise Virtual pair must be of the same switch model.
- Both switches in the Cisco StackWise Virtual pair must be running the same software version.
- All the ports used for configuring a StackWise Virtual Link (SVL) must share the same speed. For example, you cannot configure a 10G or a 40G port to form an SVL, simultaneously.
Requirements
- The dashboard network and all switches must be running IOS XE 17.18.1 or higher.
- Between 2-8 StackWise Virtual link interfaces must be configured.
- 1 dual-active-detection interface must be configured.
- Each switch must have its own uplink interface.
- Uplink ports cannot be configured as an SVL or DAD interface.
- The switch with the lowest MAC address will be elected as the active. This switch MUST have a root port if it is not configured as the root bridge.
- The switches will reboot multiple times during the StackWise Virtual setup to initiate the process and provision members. This is expected behavior.
- At this time, only Catalyst 9500 High performance series and C9550 models are supported.
Limitations
- The switch stack must be deleted prior to removing members from the network.
- A StackWise Virtual member cannot be replaced without first deleting the stack and reconfiguring
- Once StackWise Virtual has been configured, the SVL and DAD interfaces cannot be changed. To do so, the stack must be deleted and reconfigured.
Configuration
- Go to Switching > Switch stacks and click on the "Create stack" button
- Provide a name for the stack, select two C9500s from the dropdown, and then click the "Create" button
- 
    Select SVL interfaces 
Selecting an SVL interface on one switch will automatically select the same interface on the other switch.
- 
    Select DAD interfaces 
 
Note - Selecting a DAD interface on one switch will automatically select the same interface on the other switch
- 
    Click the acknowledgement checkbox and click configure
The configuration workflow will take approximately 30 minutes to complete.
Verifying StackWise Virtual Configuration Using the Cloud CLI
StackWise Virtual can be monitored and verified by using the Cloud CLI tool.
Verifying StackWise Virtual link interfaces
Verifying StackWise Virtual dual active detection interfaces
Verifying switch stack
Configuring a Flexible Switch Stack
Up to eight Meraki MS420/425 switches can be configured in a flexible stack to allow for high-speed communication between devices.
Flexible stacking is available on MS420 and MS425 switches, which do not have dedicated stacking ports. Any SFP+ interface on these switches can be configured as a stack port. On the MS425, the QSFP+ ports can also be configured as stack ports. Please note that 10 Gb/s is the minimum speed required to support flexible stacking. This section describes flexible stacking.
Only like-models can be stacked. For example, MS350-48 and MS350-24X can be stacked, but MS250-48 cannot be stacked with a MS350-48.
Physical stacking is available on MS150, MS210, MS225, MS250, MS350, MS355, MS390, MS410, and MS450 switches, which include dedicated stacking ports. For physical stacking, check the Configuring a Physical Switch Stack section of this article.
On the MS420 and MS425 series switches, you have the flexibility to use any of the front switch ports as either ethernet (default) or stacking. This option is available under the port configuration and can be easily modified by just selecting enable from the dropdown.
Converting a link aggregate to a stacking port is not a supported configuration and may result in unexpected behavior.
Once this configuration is made and the switches have downloaded the new configuration, it is recommended to follow a similar ring topology as mentioned above for the overall switch port cabling. Switch ports configured as stack ports will show up with a new symbol on the nodes status page to indicate that it is configured for stacking.
Flexible Switch Stack Configuration Steps
The following steps explain how to prepare a group of switches for flexible stacking, how to stack them together, and how to configure the stack in dashboard:
- Add the switches into a dashboard network. This can be a new dashboard network for these switches, or an existing network with other switches. For that, you may follow Adding and Removing Devices from Dashboard Networks article. Do not configure the stack in dashboard yet.
- Connect an uplink to each switch. Ensure that the uplink switch ports are different than the intended stacking ports.
- Power on all the switches, then wait several minutes for them to download the latest firmware and updates from dashboard. The switches may reboot during this process.
    
