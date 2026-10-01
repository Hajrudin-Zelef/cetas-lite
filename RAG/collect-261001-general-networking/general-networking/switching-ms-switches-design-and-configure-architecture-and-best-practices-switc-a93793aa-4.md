---
id: collect-261001-general-networking/general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa-4
title: "switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa.md
source_anchor: ""
source_lines: [196, 246]
sha256: 9d935c5926e3367f9eab984ae8445b0b4a54454b51d5d8e819e65a263b8ac67d
---

# switching-ms-switches-design-and-configure-architecture-and-best-practices-switc-a93793aa

  - The power LEDs on the front of each switch will blink during this process.
  - Once the switches are done downloading and installing firmware, their power LEDs will stay solid white or green.
- Choose (but do not yet connect) two switch ports per switch to be the dedicated stacking ports. Switch stacks should be connected in a ring topology (as shown in the following image). Ensure that the stacking ports are different than the switch uplink switch port. Do not actually connect the switch ports yet. This will be done in Step 6.
It is recommended to configure and use identical switch port types as flexible stacking ports. For example, 2 x 10Gb/s (SFP+) or 2 x 40Gb/s (QSFP) interfaces can be connected together as flexible stacking ports.
Please note that 10 Gb/s is the minimum speed required to support flexible stacking. 
- Configure the intended switch ports for stacking in dashboard under Switching > Monitor > Switch ports:
- Connect the switch stack via the intended stacking ports like the image shown in step 4.
- Navigate to Switching > Monitor > Switch stacks.
- Configure the switch stack in dashboard. If dashboard has already detected the correct stack under Detected potential stacks, click Provision this stack to automatically configure the stack.
Otherwise, to configure the stack manually:
- Click add one / Add a stack
Note: The MS AutoStacking process applies to flexible switch stacks as well, please see the note above regarding auto-provisioning.
- Select the checkboxes of the switches you would like to stack, name the stack, and then click Create.
- Disconnect all but one switch's uplink, which will be the uplink for the switch stack.
Viewing and Creating Your Stacks
The Switch stacks page gives you quick access to all of the configured stacks in the network as well as provides easy configuration options for new stacks that are being deployed. Clicking on "Add a stack" or when there is a detected stack you will be able to easily configure a new physical stack.
Check the Health of the Stack
Viewing a Stack
In order to check the stack status visually simply click any row on the stacks List. This will take you to the overview of the stack selected. From here you can easily get a feel for connected switch ports and which switches are contained in the stack. We've included the capability to blink the LEDs on switches in the stack to easily indicate which switch it is for anyone who is on site looking at the stack.
Managing Stack Members
To add or remove stack members simply click on the tab labeled Manage members and select the switch(es) that you want to add or remove from the stack and click either add or remove switches.
Deleting a Stack
To delete a stack in its entirety, browse to Switching > Switch stacks, from there select the checkbox of the stack in question and then click on the "Delete stacks" button.
You will then be prompted with a warning regarding Layer 3 configuration whether or not any such configuration exists on the stack:
Clicking confirm will then successfully delete the stack and return your switches back to stand-alone operation and configuration. It is recommended that the switches are allowed time to fetch configuration and are then powered down and stack cables removed.
Removing a Stack Member
To remove a stack member from a switch stack, browse to Switching > Switch stacks, from there select the stack in question and head over to manage members. Select the switch that you would like to remove and then click on "Remove switches"
If any of the following configuration existed on the switch prior to adding it to the stack, then they will be recovered upon removing it from or deleting the stack:
- Link aggregates
- Mirrored switch ports
- Switched virtual interface (SVIs)
- Internet Group Management Protocol (IGMP) snooping
- Spanning-Tree Protocol (STP) priority
When a Catalyst switch (including MS390) is removed from a stack, it retains the switch ID it had while in the stack. A factory reset does not change the switch ID, and there is currently no way to modify this ID on switches using Cloud configuration.
Attempting to remove a stack member from a pair of 2 switches will fail and result in a "Stacks must have at least 2 switches" error. Instead, the "Delete stacks" button should be used at the Switch stacks overview.
Viewing Switch Stack Role Information
When using physical or flexible stacking, one switch in the stack acts as the active switch, managing all of the stack functions such as spanning-tree. It may occasionally be useful to understand which switch is the active switch and which are stack members.
Beginning with MS 15.18, the stack role is displayed on the switch details page for an individual switch accessed from the Switching > Monitor > Switches page if the switch in question belongs to a switch stack. The "switch stack" section of the left-hand column displays one of three stack roles:
- Active
- Standby
- Member
Switches other than the MS390 have one active switch in the stack and all other switches are member switches. 
MS390 stacks also have one standby switch. The standby becomes the active unit in the event of a failure of that switch.
Replacing and Cloning Stack Members
The steps below should be used for the following use cases:
- Replacing a Stack Member
    
  - A stacked switch has failed and needs to be replaced.
  - A stacked switch needs to be replaced in a stack with 8 switches.
- Cloning a Stack Member
    
