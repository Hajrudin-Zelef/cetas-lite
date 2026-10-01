---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-5
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [352, 461]
sha256: 45a6c3ac19c4d8df15b9cdc082511b645d5c9c09f97a7221114ce8bc628283c8
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

on either of the Cisco StackWise Virtual members.
The Cisco StackWise Virtual active switch synchronizes the Multicast Forwarding Information Base (MFIB) state to the Cisco
StackWise Virtual standby switch. On both the member switches, all the multicast routes are loaded in the hardware, with replica
expansion table (RET) entries programmed for only local, outgoing interfaces. Both the member switches are capable of performing
hardware forwarding.
Note
To avoid multicast route changes as a result of a switchover, we recommend that all the links carrying multicast traffic
be configured as MEC rather than Equal Cost Multipath (ECMP).
For packets traversing an SVL, all Layer 3 multicast replications occur on the egress switch. If there are multiple receivers
on the egress switch, only one packet is replicated and forwarded over the SVL, and then replicated to all the local egress
ports.
Software Features
Software features run only on the Cisco StackWise Virtual active switch. Incoming packets to the Cisco StackWise Virtual
standby switch that require software processing are sent across an SVL to the Cisco StackWise Virtual active switch.
Dual-Active Detection
If the standby switch detects a complete loss of the SVL, it assumes the active switch has failed and will take over as the
active switch. However, if the original Cisco StackWise Virtual active switch is still operational, both the switches will
now be Cisco StackWise Virtual active switches. This situation is called a dual-active scenario. This scenario can have adverse
effects on network stability because both the switches use the same IP addresses, SSH keys, and STP bridge IDs. Cisco StackWise
Virtual detects a dual-active scenario and takes recovery action. DAD link is the dedicated link used to mitigate this.
If the last available SVL fails, the Cisco StackWise Virtual standby switch cannot determine the state of the Cisco StackWise
Virtual active switch. To ensure network uptime without delay, the Cisco StackWise Virtual standby switch then assumes the
Cisco StackWise Virtual active role. The original Cisco StackWise Virtual active switch enters recovery mode and brings down
all its interfaces, except the SVL and the management interfaces.
Note
On the Cisco Catalyst 9500X Series Switches:
Dynamic addition and removal of SVL and DAD links are supported. If the switch is already operating in SVL mode, a device
restart is not required for the SVL and DAD link addition or removal configuration to take effect.
If a user tries to remove the last active SVL link, the user is notified of a stack split through a syslog message.
Dual-Active-Detection Link with Fast Hello
To use the dual-active fast hello packet detection method, you must provision a direct ethernet connection between the two
Cisco StackWise Virtual switches. You can dedicate up to four links for this purpose.
The two switches start with exchanging dual-active hello messages containing information about the initial switch states.
If all SVLs fail and a dual-active scenario occurs, each switch will trigger an exchange of the dual-active hello messages
which allows it to recognize that there is a dual-active scenario from the peer's messages.
This initiates recovery actions as described in the Recovery Actions section. If a switch does not receive an expected dual-active fast hello message from the peer before the timer expires,
the switch assumes that the link is no longer capable of dual-active detection.
Note
Do not use the same port for StackWise Virtual Link and dual-active detection link.
Dual-Active Detection with enhanced PAgP
Port aggregation protocol (PAgP) is a Cisco proprietary protocol used for managing EtherChannels. If a StackWise Virtual MEC
terminates on a Cisco switch, you can run PAgP protocol on the MEC. If PAgP is running on the MECs between the StackWise Virtual
switch and an upstream or downstream switch, the StackWise Virtual can use PAgP to detect a dual-active scenario. The MEC
must have at least one port on each switch of the StackWise Virtual setup.
Enhanced PAgP is an extension of the PAgP protocol. In virtual switch mode, ePAgP messages include a new type length value
(TLV) which contains the ID of the StackWise Virtual active switch. Only switches in virtual switch mode send the new TLV.
When the StackWise Virtual standby switch detects SVL failure, it initiates SSO and becomes StackWise Virtual active. Subsequent
ePAgP messages sent to the connected switch from the newly StackWise Virtual active switch contain the new StackWise Virtual
active ID. The connected switch sends ePAgP messages with the new StackWise Virtual active ID to both StackWise Virtual switches.
If the formerly StackWise Virtual active switch is still operational, it detects the dual-active scenario because the StackWise
Virtual active ID in the ePAgP messages changes.
Note
To avoid PAgP flaps and to ensure that dual-active detection functions as expected, the stack MAC persistent wait timer must
be configured as indefinite using the command stack-mac persistent timer 0 .
Recovery Actions
A Cisco StackWise Virtual active switch that detects a dual-active condition shuts down all of its non-SVL or non-DAD interfaces
to remove itself from the network. The switch then waits in recovery mode until the SVLs recover. You should physically repair
the SVL failure and the switch automatically reloads and restores itself as the standby switch. To enable the switch to remain
in recovery mode after restoring the SVL links, see Disabling Recovery Reload section.
Implementing Cisco StackWise Virtual
The two-node solution of Cisco StackWise Virtual is normally deployed at the aggregation layer. Two switches are connected
over an SVL.
Cisco StackWise Virtual combines the two switches into a single logical switch with a large number of ports, offering a single
point of management. One of the member switches is the active and works as the control and management plane, while the other
one is the standby. The virtualization of multiple physical switches into a single logical switch is only from a control and
management perspective. Because of the control plane being common, it may look like a single logical entity to peer switches.
The data plane of the switches are converged, that is, the forwarding context of a switch might be passed to the other member
switch for further processing when traffic is forwarded across the switches. However, the common control plane ensures that
all the switches have equivalent data plane entry for each forwarding entity.
An election mechanism that determines which switch is Cisco StackWise Virtual active and which one is a control plane standby,
is available. The active switch is responsible for management, bridging and routing protocols, and software data path. These
are centralized on the active switch supervisor of the Cisco StackWise Virtual active switch.
How to Configure Cisco StackWise Virtual
Configuring Cisco StackWise Virtual Settings
To enable StackWise Virtual, perform the following procedure on both the switches:
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password, if prompted.
Step 2
switch switch-number renumber new switch -number
Example:
Device# switch 1 renumber 2
(Optional) Reassigns the switch number.
The default switch number will be 1. The valid values for the new switch number are 1 and 2.
Step 3
switch switch-number priority priority-number
Example:
Device# switch 1 priority 5
(Optional) Assigns the priority number.
The default priority number is 1. The highest priority number is 15.
Step 4
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 5
stackwise-virtual
Example:
Device(config)# stackwise-virtual
Enables Cisco StackWise Virtual and enters stackwise-virtual submode.
Step 6
domain id
Example:
Device(config-stackwise-virtual)# domain 2
