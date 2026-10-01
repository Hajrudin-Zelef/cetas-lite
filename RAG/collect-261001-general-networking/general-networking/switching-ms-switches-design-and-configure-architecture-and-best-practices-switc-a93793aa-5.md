---
id: collect-261001-general-networking/general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa-5
title: "switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa.md
source_anchor: ""
source_lines: [247, 330]
sha256: 4b05b6b051cc18cff79802b4ae2d56741308b3e5249a9522ca44d13d0dc31104
---

# switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa

  - A stacked switch needs to be replaced, but the new switch should be up and running before the replacement occurs.
  - A new switch needs to be added and requires the same switch port configurations of another stack member.
Note: All the following instructions are the same for both physical and flexible switch stacks.
Replacing and Cloning a Stack Member Video
Replacing a Stack Member
The following steps will clone the original stack member and remove it from the stack:
If the active stack member is replaced there will be a short loss in connectivity when it is powered down while a new active member is elected.
When the Active Member of a Classic MS Stack reboots, LACP links can take up to ~90 seconds to come back up when connected to certain neighbors (i.e. Catalyst/MS390).
For Catalyst switches (including MS390) on CS firmware, it is required you power down the entire switch stack before adding / replacing a new member to avoid packet forwarding or configuration sync issues. This is only a requirement for stacks on CS firmware and not needed for stacks running IOS-XE firmware.
- Claim the new/replacement switch in the inventory (Using the Organization Inventory):
- Navigate to Organization > Inventory
- Click the Claim button
- Enter the serial number of the new switch. If replacing multiple members, list all serials
- Click Claim
- Add the switch to the network containing the stack
- Select the switch to be added to the network
- Click Add to...
- Select the network and Add to existing
Note: After the switch has been added to the network and before it is added to the stack or replaced, it should be brought online individually and updated to the same firmware build as the rest of the stack. Failing to do so can prevent the switch from stacking successfully. The configured firmware build for the network can be verified under Organization > Firmware Upgrades. A flashing white or green LED on the status light on the switch indicates that a firmware upgrade is in progress.
- (Optional) Edit the name of the new switch
- Navigate to Switching > Monitor > Switches
- Select the new switch
- Click the next to the title to rename the switch
- Clone and replace the stack member
- Navigate to Switching > Monitor > Switch stacks
- Select the existing stack
- Navigate to the Clone and replace member tab
- Select the source switch to be replaced
- Select the destination switch which will replace the source switch
- Click Clone switch
- Power off the stack member to be replaced.
- Physically swap the switches.
The old switch can then be repurposed as a standalone switch or removed from the network (Adding and Removing Devices from Dashboard Networks).
Cloning a Stack Member
The following steps will clone the original stack member without removing it from the stack.
- Claim the new/replacement switch in the inventory:
- Navigate to Organization > Inventory
- Click the Claim button
- Enter the serial number of the new switch. If replacing multiple members, list all serials
- Click Claim
- Add the switch to the network containing the stack
- Select the switch to be added to the network
- Click Add to...
- Select the network and Add to existing
- (Optional) Edit the name of the new switch
- Navigate to Switching > Monitor > Switches
- Select the new switch
- Click the next to the title to rename the switch
- Clone the stack member
- Navigate to Switching > Monitor > Switches
- Select the replacement switch
- Click on Clone as highlighted below
- Search for the original stack member name or mac address and select it.
- Click Clone.
- Add the new switch to a stack.
- Navigate to Switching > Monitor > Switch stacks
- Select the existing stack
- Navigate to the Manage members tab. In the Add members section, select the switch to add
- Click Add Switches
Replacing a Stack Member for Device Configuration Source
If your new switch does not already have a cloud ID please follow the instructions here first to generate one.
- Do not power down the entire switch stack.
- Power off the faulty switch and remove it from the switch stack.
- With the replacement switch powered off, physically add it to the switch stack.
- Power on the new switch.
- Add the new switch's cloud ID to the same network as the switch stack.
- Select Device Configuration and enter the same credentials used for the stack.
- Wait several minutes for the dashboard to detect the new member and for it to appear in the switch stack.
Switch Replacement Walkthrough for Stacks Bound to a Template
Layer 3 Interface Configuration
For configuration details and caveats regarding Layer 3 switch stack routing, please reference the article MS Layer 3 Switching and Routing.
Common Alerts
Ensure all stack members are configured on dashboard, online and connected via their stacking ports.
Note: If connected and configured correctly, the alert will disappear within up to 1 hour. If the error persists, please contact Cisco Meraki Technical Support for further troubleshooting.
This switch's current stack members differ from the dashboard configuration/Misconfigured Switch.
This can occur in the following scenarios:
- Stack members are configured on dashboard, but not all members are connected via their stacking ports.
- A stack member has failed or is powered off.
This switch is not connected to a stack/Switch not connected to stack.
This can occur in the following scenarios:
- The switch is configured on dashboard as a stack member, but is not connected to a stack.
This switch does not have a stack configuration/Unconfigured Switch.
This can occur in the following scenarios:
- The switch is physically connected as a stack, but not configured on dashboard as a stack member.
