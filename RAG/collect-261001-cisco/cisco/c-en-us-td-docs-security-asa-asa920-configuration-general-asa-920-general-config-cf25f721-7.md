---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721-7
title: "c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721.md
source_anchor: ""
source_lines: [459, 524]
sha256: 34fc28ef0b3b8a1292c36b339d08ab27f784978f5496afae2a30456219ae93fc
---

# c-en-us-td-docs-security-asa-asa920-configuration-general-asa-920-general-config-cf25f721

                                          No Link: The physical link for the interface is down.
- 
                                          
                                          Failed: No traffic is received on the interface, yet traffic is heard on the peer interface.
Failover times
Failover times are high availability mechanisms that:
- 
                                    
                                    initiate automatic switchover when specific failure conditions occur on the active unit,
- 
                                    
                                    monitor system health through configurable detection thresholds and timing parameters, and
- 
                                    
                                    ensure service continuity by transferring control to the standby unit when failures exceed defined limits.
Failover triggering events and timing parameters
These events trigger failover in a Firepower high availability pair:
- 
                                    
                                    More than 50% of the Snort instances on the active unit are down.
- 
                                    
                                    Disk space on the active unit is more than 90% full.
- 
                                    
                                    The no failover active command is run on the active unit or the failover active command is run on the standby unit.
- 
                                    
                                    The active unit has more failed interfaces than the standby unit.
- 
                                    
                                    Interface failure on the active device exceeds the threshold configured. By default, failure of a single interface causes failover. You can change the default value by configuring a threshold for the number of interfaces or a percentage of monitored interfaces that must fail for the failover to occur. If the threshold is exceeded on the active device, failover occurs. If the threshold is exceeded on the standby device, the unit moves to Fail state. To change the default failover criteria, enter this command in global configuration mode: Table 2. Interface policy command Command Purpose failover interface-policy num [%] hostname (config)# failover interface-policy 20%Changes the default failover criteria. When specifying a specific number of interfaces, the num argument can be from 1 to 250. When specifying a percentage of interfaces, the num argument can be from 1 to 100.
| Note | If you manually fail over using the CLI or ASDM, or you reload the ASA, the failover starts immediately and is not subject to the timers listed below. | 
| Table 3. ASA |  |  |  | 
|---|---|---|---|
| Failover condition | Minimum | Default | Maximum | 
|---|---|---|---|
| Active unit loses power, hardware goes down, or the software reloads or crashes. When any of these occur, the monitored interfaces or failover link do not receives any hello message. | 800 milliseconds | 15 seconds | 45 seconds | 
| Active unit main board interface link down. | 500 milliseconds | 5 seconds | 15 seconds | 
| Active unit 4GE module interface link down. | 2 seconds | 5 seconds | 15 seconds | 
| Active unit interface up, but connection problem causes interface testing. | 5 seconds | 25 seconds | 75 seconds | 
Configuration Synchronization
Failover includes various types of configuration synchronization.
Running Configuration Replication
Running configuration replication occurs when any one or both of the devices in the failover pair boot.
In Active/Standby failover, configurations are always synchronized from the active unit to the standby unit.
In Active/Active failover, whichever unit boots second obtains the running configuration from the unit that boots first, regardless of the primary or secondary designation of the booting unit. After both units are up, commands entered in the system execution space are replicated from the unit on which failover group 1 is in the active state.
When the standby/second unit completes its initial startup, it clears its running configuration (except for the failover commands needed to communicate with the active unit), and the active unit sends its entire configuration to the standby/second unit. When the replication starts, the ASA console on the active unit displays the message “Beginning configuration replication: Sending to mate,” and when it is complete, the ASA displays the message “End Configuration Replication to mate.” Depending on the size of the configuration, replication can take from a few seconds to several minutes.
On the unit receiving the configuration, the configuration exists only in running memory. You should save the configuration to flash memory according to Save Configuration Changes. For example, in Active/Active failover, enter the write memory all command in the system execution space on the unit that has failover group 1 in the active state. The command is replicated to the peer unit, which proceeds to write its configuration to flash memory.
| Note | During replication, commands entered on the unit sending the configuration may not replicate properly to the peer unit, and commands entered on the unit receiving the configuration may be overwritten by the configuration being received. Avoid entering commands on either unit in the failover pair during the configuration replication process. | 
File Replication
Configuration syncing does not replicate the following files and configuration components, so you must copy these files manually so they match:
- 
                                    Secure Client images
- 
                                    CSD images
- 
                                    Secure Client profiles The ASA uses a cached file for the Secure Client profile stored in cache:/stc/profiles, and not the file stored in the flash file system. To replicate the Secure Client profile to the standby unit, perform one of the following: 
  - 
                                          Enter the write standby command on the active unit.
  - 
                                          Reapply the profile on the active unit.
  - 
                                          Reload the standby unit.
- 
                                          
