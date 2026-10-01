---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-1
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [1, 82]
sha256: c6905045c1196a115abda025d83bff40338ceaa3b7a35d6cc07a8a8748778948
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

About Failover
Configuring failover requires two identical ASAs connected to each other through a dedicated failover link and, optionally, a state link. The health of the active units and interfaces is monitored to determine whether they meet the specific failover conditions. If those conditions are met, failover occurs.
Failover Modes
The ASA supports two failover modes, Active/Active failover and Active/Standby failover. Each failover mode has its own method for determining and performing failover.
- 
                                 In Active/Standby failover, one device functions as the Active unit and passes traffic. The second device, designated as the Standby unit, does not actively pass traffic. When a failover occurs, the Active unit fails over to the Standby unit, which then becomes Active. You can use Active/Standby failover for ASAs in single or multiple context mode.
- 
                                 In an Active/Active failover configuration, both ASAs can pass network traffic. Active/Active failover is only available to ASAs in multiple context mode. In Active/Active failover, you divide the security contexts on the ASA into 2 failover groups. A failover group is simply a logical group of one or more security contexts. One group is assigned to be Active on the primary ASA, and the other group is assigned to be active on the Secondary ASA. When a failover occurs, it occurs at the failover group level.
Both failover modes support stateful or stateless failover.
Failover system requirements
Failover system requirements are specifications that define the hardware, software, and license prerequisites for ASAs to function in a Failover configuration.
Hardware requirements
Hardware requirements are specifications that ensure two units in a Failover configuration can
- 
                                       
                                       maintain identical hardware configurations including model, interfaces, modules, and RAM
- 
                                       
                                       support proper synchronization and failover functionality, and
- 
                                       
                                       ensure reliable high availability operation.
Hardware requirement specifications
The two units in a Failover configuration must meet these requirements:
- 
                                       
                                       Be the same model. For the Firepower 9300, High Availability is only supported between same-type modules. However, the two chassis can include mixed modules. For example, each chassis has an SM-56, SM-48, and SM-40. You can create High Availability pairs between the SM-56 modules, between the SM-48 modules, and between the SM-40 modules.
- 
                                       
                                       Have the same number and types of interfaces. For the Firepower 2100 in Platform mode and Firepower 4100/9300 chassis, all interfaces must be preconfigured in FXOS identically before you enable Failover. If you change the interfaces after you enable Failover, make the interface changes in FXOS on the Standby unit, and then make the same changes on the Active unit. If you remove an interface in FXOS (for example, if you remove a network module, remove an EtherChannel, or reassign an interface to an EtherChannel), then the ASA configuration retains the original commands so that you can make any necessary adjustments; removing an interface from the configuration can have wide effects. You can manually remove the old interface configuration in the ASA OS.
- 
                                       
                                       Have the same modules installed (if any).
- 
                                       
                                       Have the same RAM installed.
If you are using units with different flash memory sizes in your Failover configuration, make sure the unit with the smaller flash memory has enough space to accommodate the software image files and the configuration files. If it does not, configuration synchronization from the unit with the larger flash memory to the unit with the smaller flash memory will fail.
Software requirements
Software requirements are configuration specifications that Failover units must meet to function properly in a failover configuration.
Required software specifications
The two units in a Failover configuration must meet these requirements:
- 
                                       
                                       Be in the same context mode (single or multiple).
- 
                                       
                                       For single mode: Be in the same firewall mode (routed or transparent). In multiple context mode, the firewall mode is set at the context-level, and you can use mixed modes.
- 
                                       
                                       Have the same major (first number) and minor (second number) software version. However, you can temporarily use different versions of the software during an upgrade process; for example, you can upgrade one unit from Version 8.3(1) to Version 8.3(2) and have failover remain active. We recommend upgrading both units to the same version to ensure long-term compatibility.
- 
                                       
                                       Have the same Secure Client images. If the failover pair has mismatched images when a hitless upgrade is performed, then the clientless SSL VPN connection terminates in the final reboot step of the upgrade process, the database shows an orphaned session, and the IP pool shows that the IP address assigned to the client is "in use."
- 
                                       
                                       Be in the same FIPS mode.
- 
                                       
                                       
                                       (Firepower 4100/9300) Have the same flow offload mode, either both enabled or both disabled.
License Requirements
The two units in a failover configuration do not need to have identical licenses; the licenses combine to make a failover cluster license.
Failover and stateful failover links
A failover link and stateful failover link are dedicated connections that
- 
                                    
                                    establish communication between the two units in a failover configuration
- 
                                    
                                    require the same interface to be used between corresponding devices for consistency, and
- 
                                    
                                    transmit critical failover and state information between the paired units.
Configuration recommendations
Cisco recommends to use the same interface between two devices in a failover link or a stateful failover link. For example, in a failover link, if you have used eth0 in device 1, use the same interface (eth0) in device 2 as well.
| Caution | All information sent over the failover and state links is sent in clear text unless you secure the communication with an IPsec tunnel or a failover key. If the ASA is used to terminate VPN tunnels, this information includes any usernames, passwords and preshared keys used for establishing the tunnels. Transmitting this sensitive data in clear text could pose a significant security risk. We recommend securing the failover communication with an IPsec tunnel or a failover key if you are using the ASA to terminate VPN tunnels. | 
Failover link
A failover link is a communication mechanism that enables two units in a failover pair to constantly communicate to determine the operating status of each unit.
Failover link data
This below information is communicated over the failover link:
- 
                                          
