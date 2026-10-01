---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-10
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "distribution", "ethernet"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [234, 315]
sha256: 3df801ebceaf04de49a298d3e5ec6077e32b773c01688718cae2f847665d06ed
---

# ms-meraki-campus-lan-5d88fe48

When upgrading Meraki switches it is important that you allocate enough time in your upgrade window for each group or phase to ensure a smooth transition. Each upgrade cycle needs enough time to download the new version to the switches, perform the upgrade, allow the network to reconverge around protocols such as spanning tree and OSPF that may be configured in your network, and some extra time to potentially roll back if any issue is uncovered after the upgrade.
Meraki firmware release cycle consists of three stages during the firmware rollout process namely beta, release candidate (RC) and stable firmware. This cycle is covered in more detail in the Meraki Firmware Development Lifecycle section.
Please note that Meraki beta is fully supported by Meraki Technical Support and can be considered as an Early Field Deployment release. If you have any issues with the new beta firmware you can always roll back to the previous stable version, or the previously installed version if you roll back within 14 days
The high-level process for a switch upgrade involves the following:
- 
    The switch downloads the new firmware (time varies depending on your connection)
- 
    The switch starts a countdown of 20 minutes to allow any other switches downstream to finish their download
- 
    The switch reboots with its new firmware (about a minute)
- 
    Network protocols re-converge (varies depending on configuration)
Meraki Firmware Version Status will show with one of the following options:
Each firmware version now has an additional Status column as follows:
- Good (Green) status indicates that your network is set to the latest firmware release. Minor updates may be available, but no immediate action is required.
- 
    Warning (Yellow) status means that a newer stable major firmware or newer minor beta firmware is available that may contain security fixes, new features, and performance improvements. We recommend that you upgrade to the latest stable or beta firmware version.
- 
    Critical (Red) status indicates that the firmware for your network is out of date and may have security vulnerabilities and/or experience suboptimal performance. We highly recommend that you upgrade to the latest stable and latest beta firmware release.
For more information about Firmware Upgrades, please refer to the following FAQ document.
MS390 Specific Guidance
- Meraki continues to develop software capabilities for the MS390 platform, therefore it is important to refer to the firmware changelog before setting a firmware for your network which includes MS390 switches.
- Please ensure that the firmware selected includes support for MS390 build.
- Also pay attention to the new features section as well as the known issues related to this firmware.
Staged Upgrades Guidance
- To make managing complex switched networks simpler, Meraki supports automatic staged firmware updates
- This allows you to easily designate groups of switches into different upgrade stages
- When you are scheduling your upgrades you can easily mark multiple stages of upgrades (e.g. Stage1, Stage2 and Stage3)
- Each stage has to complete its upgrade process before proceeding to the next stage
- All members of a switch stack must be upgraded at the same time, within the same upgrade window.
- You cannot select an individual switch stack member to be upgraded; only the entire switch stack can be selected
- Switch stacks upgrade behavior; each stack member rebooting close to the same time and the stack then automatically re-forming as the members come online
This feature is currently not supported when using templates
Firmware Upgrade Barriers
- Firmware upgrade barriers is a built-in feature to prevent certain upgrade paths on devices running older firmware versions trying to upgrade to a build that would otherwise cause compatibility issues.
- Having devices use intermediary builds defined by Meraki will ensure a safe transition when upgrading your devices.
Here is an example of when firmware upgrade barriers come into effect. You might find yourself in a situation where you are unable to upgrade a device for an extended period of time due to uptime or business requirements. There is a switch in the network that is running MS 9.27 and would like to update to the latest stable version, which at the time of writing, is 11.30. Attempting to upgrade from 9.27 to 11.30 will not be a selectable option in the dashboard and administrators will have to upgrade to 10.35 first.
In order to complete the upgrade from the current version to the target version, two manual upgrades will be required. The first from your current to the intermediary version, and another from the intermediary to your target version.
Meraki Switches per Dashboard Network
General Guidance
- It is recommended to keep the total number of Meraki switches (e.g. Access AND Distribution) in a dashboard network within 400 for best performance of dashboard.
- If switch count exceeds 400 switches, it is likely to slow down the loading of the network topology/ switch ports page or result in display of inconsistent output.
- It is recommended to keep the total switch port count in a network to fewer than 8000 ports for reliable loading of the switch port page
There is no hard limit on the number of switches in a network, therefore please take this into consideration when you are planning for the whole Campus LAN network.
Cabling
General Guidance
- It is recommended to use Category-5e cables for switch ports up to 1Gbps
- While Category-5e cables can support multigigabit data rates upto 2.5/5 Gbps, external factors such as noise, alien crosstalk coupled with longer cable/cable bundle lengths can impede reliable link operation.
- Noise can originate from cable bundling, RFI, cable movement, lightning, power surges and other transient events.
- It is recommended to use Category-6a cabling for reliable multigigabit operations as it mitigates alien crosstalk by design
- Please ensure that you are using Approved Meraki SFPs and Accessories per hardware model
Meraki will only support the Approved Meraki SFPs and Accessories for use with MS and MX platforms. A number of Cisco converters have also been certified for use with Meraki MS switches:
- SFP-H10GB-CU1M
- SFP-H10GB-CU3M
- SFP-10G-SR-S
- SFP-10G-SR
Power Over Ethernet (PoE)
General Guidance
- MS platforms allocate power based on the actual drawn power from the client device
MS390 Specific Guidance
- MS390s allocate power based on the requested power from the client device (as opposed to the actual drawn power).
- It is recommended to calculate your power budget based on the maximum power mentioned on the client device data sheet (e.g. MR56 consumes 30W).
- This is based on the power class advertised using Layer 2 discovery protocols (e.g. LLDP, CDP). Refer to the following table for more information on the power class and the corresponding power values:
| Class | Maximum Power Level | 
|---|---|
| 0 (unknown class) | 15.4 W | 
| 1 | 4 W | 
| 2 | 7 W | 
| 3 | 15.4 W | 
| 4 | 30 W | 
| 5 | 45 W | 
| 6 | 60 W | 
| 7 | 75 W | 
| 8 | 90 W | 
IP Addressing and VLANs
General Guidance
- All Meraki MS platforms switchports are configured in Trunk mode with Native VLAN 1 by default with Management VLAN 1
- Even if it is undesirable to use Native VLAN 1, it is recommended to use it for provisioning the switches for ZTP purposes. Once the switches/stacks are online on dashboard in running steady, you can then change the Management VLAN as required. Remember to change port settings downstream first to avoid losing access to switches.
- Assign a dedicated management VLAN for your switches which has access to the Internet (More info here)
- Avoid overlapping subnets as this may lead to inconsistent routing and forwarding
- Dedicate /24 or /23 subnets for end-user access
- Do not configure a L3 interface for the management VLAN. Use L3 interfaces only for data VLANs. This helps in separating management traffic from end-user data
