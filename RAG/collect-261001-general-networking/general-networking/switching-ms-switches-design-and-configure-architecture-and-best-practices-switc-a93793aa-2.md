---
id: collect-261001-general-networking/general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa-2
title: "switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa.md
source_anchor: ""
source_lines: [76, 118]
sha256: 0ba4c15d94270783f03cdc123ea739887e643a9203333034aad0eb76c59bd506
---

# switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa

- Add the switches into a dashboard network. For that, follow the Adding and Removing Devices from Dashboard Networks article. This can be a new dashboard network for these switches, or an existing network with other switches. Do not configure the stack in dashboard yet.
- With all switches powered off and links disconnected, connect the switches together via stacking cables in a ring topology (as shown in the following image). To create a full ring, start by connecting switch 1/stack port 1 to switch 2/stack port 2, then switch 2/stack port 1 to switch 3/stack port 2 and so forth, with the bottom switch connecting to the top switch to complete the ring.
 
  - Make sure that the stacking cables are connected correctly. If there are LED lights for the stack cables, it should turn orange when a cable is detected and green when the stack status is good:
        
    - Correct orientation - The black tab on the cable heads should face up
    - Incorrect orientation - The black tab on the cable heads are facing down
  - Correct orientation - The black tab on the cable heads should face up
- Make sure that the stacking cables are connected correctly. If there are LED lights for the stack cables, it should turn orange when a cable is detected and green when the stack status is good:
        
- Connect one cable to act as a single uplink from 1 switch of the stack. Power on all the switches, then wait several minutes for them to download the latest firmware and updates from dashboard. The switches may reboot during this process.
    
  - The power LEDs on the front of each switch will blink during this process.
  - Once the switches are done downloading and installing the firmware, their power LEDs will stay solid white or green.
- The dashboard will now automatically attempt to provision detected potential stacks for you. This has the same effect as clicking the "Provision this stack" button, and the dashboard will only attempt this once. If the dashboard is unable to auto-provision the stack, the switches will appear on the same page in the Detected stacks that failed auto provisioning section, and the stack will need to be manually provisioned using the Provision this stack button once the cause of the provisioning failure has been fixed.
When rebooting a stack member from the dashboard using the "Reboot device" tool will reboot only that individual stack member. The same applies when performing a factory reset on a stack member. This does not apply to MS390 or C9300/L/X stacked switches, for the behavior of these models please refer to the section Stacking Catalyst Switches
Creating a stack manually
If the dashboard does not detect the potential stack, you can configure the stack manually. Navigate to Switching > Monitor > Switch stacks. Click the link to add one or the Add a stack button, depending on the option available:
Next, select the checkboxes of the switches you would like to stack, enter a name for the stack, and click Create.
Validate the switches are now successfully stacked via Switching > Monitor > Switch stacks
6. If unsuccessful, review and correct the error and reattempt provisioning.
NOTE: After the switch stack is up and running, multiple uplinks can be added for redundancy.
Stacking Catalyst Switches
Follow the steps below to setup a Catalyst switch stack or watch the Cisco Meraki MS390 Stack Setup Walkthrough video.
- Add the switches into a dashboard network. For that, you may follow Adding and Removing Devices from Dashboard Networks article. This can be a new dashboard network for these switches, or an existing network with other switches. Do not configure the stack in dashboard yet.
- With all switches powered off and links disconnected, connect the switches together via stacking cables in a ring topology (as shown in the following image). To create a full ring, start by connecting switch 1/stack port 1 to switch 2/stack port 2, then switch 2/stack port 1 to switch 3/stack port 2 and so forth, with the bottom switch connecting to the top switch to complete the ring. While connecting the stacking cables, ensure that each connector is correctly aligned to the switch's stacking port it is connecting to and finger-tighten the screws (clockwise direction). Make sure the Cisco logo is on the top side of the connector as shown in image.
- 
    Connect one uplink to a single switch in the stack, then power on only that switch. Wait until the switch connects to Dashboard and upgrade to desired network version is completed. Power on the remaining switches, then wait several minutes for them to download the latest firmware and updates from dashboard. The switches may reboot during this process.
- 
    Download the same firmware build using the Firmware Upgrade Manager under Organization > Monitor > Firmware upgrades, if they are not already set for this. This helps ensure each switch is running the same firmware build. Please note that it might take close to an hour for the switches to upgrade.
- 
    Navigate to Switching > Monitor > Switch stacks
- 
    Configure the switch stack in dashboard. If dashboard has already detected the correct stack under Detected potential stacks, click Provision this stack to automatically configure the stack. Otherwise, click the link to add one or the Add a stack button, depending on the option available:
It is expected that not all member switches will appear online until the stack is fully configured in the dashboard, as instructed in step 6.
Based on the number of members in a stack, the switch stack may require approximately one hour to become operational and connect to the Meraki Dashboard. 
By default the switch members are sorted alphanumerically from top to bottom on the switch stack page; this can be changed by updating the names of the relevant stack members.
The MS AutoStacking process applies to Catalyst switches as well, please see the note above regarding auto-provisioning.
- Select the checkboxes of the switches you would like to stack, name the stack, and then click Create.
7. Ensure that all switches have downloaded the latest configuration. To verify this, navigate to Switching > Monitor > Switches and select the MS390 switch. Look for CONFIG in the column on the left of the switch details page and check if the status reads Up to date.
- 
    When using static IP addressing for switch management interfaces rather than DHCP, the management IP address cannot be individually altered for each stack member through the dashboard. Once the stack is configured, all members will display the management IP address of the active (primary/master) switch, where the control plane resides. It is important to note that if a change to the management IP is made after the stack has been configured, that change will apply to the active switch and will be reflected across all stack members. The management IP is shared and not unique to each switch within the stack. 
 
