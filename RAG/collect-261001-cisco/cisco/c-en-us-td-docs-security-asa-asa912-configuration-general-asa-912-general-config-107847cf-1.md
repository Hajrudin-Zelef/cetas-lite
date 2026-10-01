---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf-1
title: "c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf.md
source_anchor: ""
source_lines: [1, 77]
sha256: afc6fcfffc15bf4124491dace3a74bf821d0911c535212bd664818500571f224
---

# c-en-us-td-docs-security-asa-asa912-configuration-general-asa-912-general-config-107847cf

About Failover
Configuring failover requires two identical ASAs connected to each other through a dedicated failover link and, optionally, a state link. The health of the active units and interfaces is monitored to determine whether they meet the specific failover conditions. If those conditions are met, failover occurs.
Failover Modes
The ASA supports two failover modes, Active/Active failover and Active/Standby failover. Each failover mode has its own method for determining and performing failover.
- 
                                 In Active/Standby failover, one device functions as the Active unit and passes traffic. The second device, designated as the Standby unit, does not actively pass traffic. When a failover occurs, the Active unit fails over to the Standby unit, which then becomes Active. You can use Active/Standby failover for ASAs in single or multiple context mode.
- 
                                 In an Active/Active failover configuration, both ASAs can pass network traffic. Active/Active failover is only available to ASAs in multiple context mode. In Active/Active failover, you divide the security contexts on the ASA into 2 failover groups. A failover group is simply a logical group of one or more security contexts. One group is assigned to be Active on the primary ASA, and the other group is assigned to be active on the Secondary ASA. When a failover occurs, it occurs at the failover group level.
Both failover modes support stateful or stateless failover.
Failover System Requirements
This section describes the hardware, software, and license requirements for ASAs in a Failover configuration.
Hardware Requirements
The two units in a Failover configuration must:
-  
                                    		  
                                    Be the same model. In addition, for container instances, they must use the same resource profile attributes. For the Firepower 9300, High Availability is only supported between same-type modules; but the two chassis can include mixed modules. For example, each chassis has an SM-36 and SM-44. You can create High Availability pairs between the SM-36 modules and between the SM-44 modules.
-  
                                    		  
                                    Have the same number and types of interfaces. For the Firepower 2100 and Firepower 4100/9300 chassis, all interfaces must be preconfigured in FXOS identically before you enable Failover. If you change the interfaces after you enable Failover, make the interface changes in FXOS on the Standby unit, and then make the same changes on the Active unit. If you remove an interface in FXOS (for example, if you remove a network module, remove an EtherChannel, or reassign an interface to an EtherChannel), then the ASA configuration retains the original commands so that you can make any necessary adjustments; removing an interface from the configuration can have wide effects. You can manually remove the old interface configuration in the ASA OS.
-  
                                    		  
                                    Have the same modules installed (if any).
-  
                                    		  
                                    Have the same RAM installed.
If you are using units with different flash memory sizes in your Failover configuration, make sure the unit with the smaller flash memory has enough space to accommodate the software image files and the configuration files. If it does not, configuration synchronization from the unit with the larger flash memory to the unit with the smaller flash memory will fail.
Software Requirements
The two units in a Failover configuration must:
-  
                                    		  
                                    Be in the same context mode (single or multiple).
- 
                                    				
                                    For single mode: Be in the same firewall mode (routed or transparent). In multiple context mode, the firewall mode is set at the context-level, and you can use mixed modes.
-  
                                    		  
                                    Have the same major (first number) and minor (second number) software version. However, you can temporarily use different versions of the software during an upgrade process; for example, you can upgrade one unit from Version 8.3(1) to Version 8.3(2) and have failover remain active. We recommend upgrading both units to the same version to ensure long-term compatibility.
-  
                                    		  
                                    Have the same AnyConnect images. If the failover pair has mismatched images when a hitless upgrade is performed, then the clientless SSL VPN connection terminates in the final reboot step of the upgrade process, the database shows an orphaned session, and the IP pool shows that the IP address assigned to the client is “in use.”
- 
                                    				
                                    Be in the same FIPS mode.
- 
                                    				
                                    				
                                    (Firepower 4100/9300) Have the same flow offload mode, either both enabled or both disabled.
License Requirements
The two units in a failover configuration do not need to have identical licenses; the licenses combine to make a failover cluster license.
Failover and Stateful Failover Links
The failover link and the optional stateful failover link are dedicated connections between the two units. Cisco recommends to use the same interface between two devices in a failover link or a stateful failover link. For example, in a failover link, if you have used eth0 in device 1, use the same interface (eth0) in device 2 as well.
| Caution | All information sent over the failover and state links is sent in clear text unless you secure the communication with an IPsec tunnel or a failover key. If the ASA is used to terminate VPN tunnels, this information includes any usernames, passwords and preshared keys used for establishing the tunnels. Transmitting this sensitive data in clear text could pose a significant security risk. We recommend securing the failover communication with an IPsec tunnel or a failover key if you are using the ASA to terminate VPN tunnels. | 
Failover Link
The two units in a failover pair constantly communicate over a failover link to determine the operating status of each unit.
Failover Link Data
The following information is communicated over the failover link:
- 
                                       		  
                                       The unit state (active or standby)
- 
                                       		  
                                       Hello messages (keep-alives)
- 
                                       		  
                                       Network link status
- 
                                       		  
                                       MAC address exchange
- 
                                       		  
                                       Configuration replication and synchronization
Interface for the Failover Link
You can use an unused data interface (physical, subinterface, redundant, or EtherChannel) as the failover link; however, you cannot specify an interface that is currently configured with a name. The failover link interface is not configured as a normal networking interface; it exists for failover communication only. This interface can only be used for the failover link (and also for the state link). For most models, you cannot use a management interface for failover unless explicitly described below.
The ASA does not support sharing interfaces between user data and the failover link. You also cannot use separate subinterfaces on the same parent for the failover link and for data.
See the following guidelines for the failover link:
- 
                                       				
