---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-8
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [525, 638]
sha256: e41874893f6ff6ba774a86f722dea0b516fdf2cc22b67d05a525972c91a84724
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

- 
                                    Local Certificate Authorities (CAs)
- 
                                    ASA images
- 
                                    ASDM images
Command Replication
After startup, commands that you enter on the active unit are immediately replicated on the standby unit. You do not have to save the active configuration to flash memory to replicate the commands.
In Active/Active failover, commands entered in the system execution space are replicated from the unit on which failover group 1 is in the active state.
Failure to enter the commands on the appropriate unit for command replication to occur causes the configurations to be out of synchronization. Those changes may be lost the next time the initial configuration synchronization occurs.
The following commands are replicated to the standby ASA:
- 
                                    All configuration commands except for mode, firewall , and failover lan unit
- 
                                    copy running-config startup-config
- 
                                    delete
- 
                                    mkdir
- 
                                    rename
- 
                                    rmdir
- 
                                    write memory
The following commands are not replicated to the standby ASA:
- 
                                    All forms of the copy command except for copy running-config startup-config
- 
                                    All forms of the write command except for write memory
- 
                                    debug
- 
                                    failover lan unit
- 
                                    firewall
- 
                                    show
- 
                                    terminal pager and pager
Config-sync optimization
Config-sync optimization is a configuration synchronization enhancement that
- 
                                       
                                       compares configuration hash values between active and joining devices to determine if full synchronization is necessary
- 
                                       
                                       enables faster device rejoining by skipping full configuration synchronization when hash values match, and
- 
                                       
                                       reduces maintenance window and upgrade time during failover operations.
Guidelines and limitations of configuration sync optimization
When a device reboots or rejoins after a suspend or resume failover, the joining device clears its running configuration. The active device then sends its entire configuration to the joining device for a full configuration synchronization. If the active device has a large configuration, this process can take several minutes.
The configuration sync optimization functionality enables comparing the configuration of the joining device and the active device by exchanging configuration hash values. If the hash computed on both active and joining devices match, the joining device skips full configuration synchronization and rejoin the failover configuration. This functionality ensures faster peering and reduces maintenance window and upgrade time.
Guidelines and limitations:
- 
                                       
                                       The configuration sync optimization functionality is enabled by default.
- 
                                       
                                       ASA multiple context mode supports configuration sync optimization by sharing the context order during full configuration synchronization, allowing comparison of context order during subsequent node-rejoin.
- 
                                       
                                       If you configure passphrase and failover IPsec key, then configuration sync optimization is not effective as the hash value computed in the active and standby devices differs.
- 
                                       
                                       If you configure the device with dynamic ACL or SNMPv3, configuration sync optimization is not effective.
- 
                                       
                                       Active device synchronizes full configuration with flapping LAN links as default behavior. During failover flaps between active and standby devices, configuration sync optimization is not triggered and devices perform a full configuration synchronization.
- 
                                       
                                       Configuration sync optimization gets triggered when the failover configuration recovers from an interruption or loss of network communication between the active and standby devices.
Monitoring config-sync optimization:
- 
                                       
                                       Syslog messages are generated displaying whether the hash values computed on the active and joining unit match, does not match, or if the operation timeout expires.
- 
                                       
                                       The syslog message also displays the time elapsed, from the time of sending the hash request to the time of getting and comparing the hash response.
Monitoring commands for configuration sync optimization.
- 
                                       
                                       show failover config-sync checksum: Displays information about the device status and checksum.
- 
                                       
                                       show failover config-sync configuration: Displays information about the device configuration and checksum.
- 
                                       
                                       show failover config-sync status: Displays status of configuration sync optimization functionality.
| Note | The Configuration State in show failover state command displays the config sync state status during HA joining. This state does not reflect the later config deployments or changes, and replication on device until a resync has been initiated. | 
Active/standby failover
Active/standby failover is a high availability configuration that
- 
                                    
                                    lets you use a standby ASA device to take over the functionality of a failed unit,
- 
                                    
                                    enables the standby unit to become the active unit when the active unit fails, and
- 
                                    
                                    allows the failed unit to come back online as the standby unit if the problem is temporary.
The failed unit must be replaced if the problem is not temporary. Set the standby unit to primary before replacing the failed unit to retain the configuration of the secondary unit.
For multiple context mode, the ASA can fail over the entire unit (including all contexts) but cannot fail over individual contexts separately.
Primary/secondary roles and active/standby status
Primary/secondary roles and active/standby status are failover configuration concepts that
- 
                                       
                                       define unit hierarchy, with primary units taking precedence during simultaneous startup
- 
                                       
                                       determine IP address and MAC address usage patterns during failover operations, and
- 
                                       
