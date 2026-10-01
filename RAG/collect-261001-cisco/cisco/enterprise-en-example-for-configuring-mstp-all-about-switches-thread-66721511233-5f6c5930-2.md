---
id: collect-261001-cisco/cisco/enterprise-en-example-for-configuring-mstp-all-about-switches-thread-66721511233-5f6c5930-2
title: "enterprise-en-example-for-configuring-mstp-all-about-switches-thread-66721511233-5f6c5930"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-example-for-configuring-mstp-all-about-switches-thread-66721511233-5f6c5930.md
source_anchor: ""
source_lines: [51, 87]
sha256: d517b62cc9c952c57f237d047518b920b60cfca3489d61c8ffe95f134acfbbde
---

# enterprise-en-example-for-configuring-mstp-all-about-switches-thread-66721511233-5f6c5930

| <HUAWEI> system-view [HUAWEI] sysname SwitchD [SwitchD] vlan batch 2 to 4094 [SwitchD] interface gigabitethernet0/0/1 [SwitchD-GigabitEthernet0/0/1] port link-type trunk [SwitchD-GigabitEthernet0/0/1] port trunk allow-pass vlan 2 to 4094 [SwitchD-GigabitEthernet0/0/1] quit [SwitchD] interface gigabitethernet0/0/2 [SwitchD-GigabitEthernet0/0/2] port link-type trunk [SwitchD-GigabitEthernet0/0/2] port trunk allow-pass vlan 2 to 4094 [SwitchD-GigabitEthernet0/0/2] quit | 
Step 2: Configure switches to work in MSTP mode.
On Huawei switches, MSTP is enabled by default, so you can skip this step.
| [SwitchA] stp mode mstp | 
| [SwitchB] stp mode mstp | 
| [SwitchC] stp mode mstp | 
| [SwitchD] stp mode mstp | 
Step 3: Configure the region named RG1 and specify the mapping between VLANs and MSTIs.
The region configuration on the four devices must be consistent so that the loop can be eliminated.
| [SwitchA] stp region-configuration [SwitchA-mst-region] region-name RG1 //Configure the region name RG1. [SwitchA-mst-region] instance 1 vlan 1 to 200 //By default, all VLANs are mapped to MSTI 0. Here, VLANs 1 to 200 are mapped to MSTI 1, and VLANs 201 to 4094 are mapped to MSTI 0. [SwitchA-mst-region] active region-configuration //Activate the region configuration. [SwitchA-mst-region] quit | 
| [SwitchB] stp region-configuration [SwitchB-mst-region] region-name RG1 [SwitchB-mst-region] instance 1 vlan 1 to 200 [SwitchB-mst-region] active region-configuration [SwitchB-mst-region] quit | 
| [SwitchC] stp region-configuration [SwitchC-mst-region] region-name RG1 [SwitchC-mst-region] instance 1 vlan 1 to 200 [SwitchC-mst-region] active region-configuration [SwitchC-mst-region] quit | 
| [SwitchD] stp region-configuration [SwitchD-mst-region] region-name RG1 [SwitchD-mst-region] instance 1 vlan 1 to 200 [SwitchD-mst-region] active region-configuration [SwitchD-mst-region] quit | 
Step 4: Configure the root bridge and secondary root bridge.
Configure SwitchA as the root bridge and SwitchB as the secondary root bridge in MSTI 0, and configure SwitchA as the secondary root bridge and SwitchB as the root bridge in MSTI 1.
| [SwitchA] stp instance 0 root primary //You can also use the stp priority 0 command to set the STP priority to 0. Running the stp priority 0 command is equivalent to running the stp root primary command. [SwitchA] stp instance 1 root secondary //You can also use the stp priority 4096 command to set the STP priority to 4096. The stp priority 4096 command is equivalent to the stp root secondary command. | 
| [SwitchB] stp instance 0 root secondary [SwitchB] stp instance 1 root primary | 
Step 5: Disable STP on GE0/0/3 interfaces of SwitchC and SwitchD.
| [SwitchC] interface gigabitethernet0/0/3 [SwitchC-GigabitEthernet0/0/3] stp disable [SwitchC-GigabitEthernet0/0/3] quit | 
| [SwitchD] interface gigabitethernet0/0/3 [SwitchD-GigabitEthernet0/0/3] stp disable [SwitchD-GigabitEthernet0/0/3] quit | 
Step 6: Enable STP globally.
On Huawei X7 series switches, STP is enabled by default, so you can skip this step.
| [SwitchA] stp enable | 
| [SwitchB] stp enable | 
| [SwitchC] stp enable | 
| [SwitchD] stp enable | 
Step 7: Verify the configuration.
Check brief information about MSTP. You can view the port roles and states.
Configuration file of SwitchA
| # sysname SwitchA # vlan batch 2 to 4094 # stp instance 0 root primary stp instance 1 root secondary stp enable # stp region-configuration region-name RG1 instance 1 vlan 1 to 200 active region-configuration # interface GigabitEthernet0/0/1 port link-type trunk port trunk allow-pass vlan 2 to 4094 # interface GigabitEthernet0/0/2 port link-type trunk port trunk allow-pass vlan 2 to 4094 # return | 
Configuration file of SwitchB
| # sysname SwitchB # vlan batch 2 to 4094 # stp instance 0 root secondary stp instance 1 root primary stp enable # stp region-configuration region-name RG1 instance 1 vlan 1 to 200 active region-configuration # interface GigabitEthernet0/0/1 port link-type trunk port trunk allow-pass vlan 2 to 4094 # interface GigabitEthernet0/0/2 port link-type trunk port trunk allow-pass vlan 2 to 4094 # return | 
Configuration file of SwitchC
| # sysname SwitchC # vlan batch 2 to 4094 # stp enable # stp region-configuration region-name RG1 instance 1 vlan 1 to 200 active region-configuration # interface GigabitEthernet0/0/1 port link-type trunk port trunk allow-pass vlan 2 to 4094 # interface GigabitEthernet0/0/2 port link-type trunk port trunk allow-pass vlan 2 to 4094 # interface GigabitEthernet0/0/3 stp disable # return | 
Configuration file of SwitchD
| # sysname SwitchD # vlan batch 2 to 4094 # stp mode stp # interface GigabitEthernet0/0/1 port link-type trunk port trunk allow-pass vlan 2 to 4094 # interface GigabitEthernet0/0/2 port link-type trunk port trunk allow-pass vlan 2 to 4094 # interface GigabitEthernet0/0/3 stp disable # return | 
★★★Summary★★★ All About Huawei Switch Features and Configurations
