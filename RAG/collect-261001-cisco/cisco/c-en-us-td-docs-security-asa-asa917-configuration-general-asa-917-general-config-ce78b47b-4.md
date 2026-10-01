---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b-4
title: "c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b.md
source_anchor: ""
source_lines: [213, 232]
sha256: cab1c01777d41a7753f7b9f731c2ba06ee34cc8f9bab3e2244a169658c59d2b8
---

# c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b

| linkUp (IF-MIB) | ifIndex, ifAdminStatus, ifOperStatus | The linkup trap for interfaces. The snmp-server enable traps snmp linkup command is used to enable and disable transmission of these traps. | 
| mteTriggerFired (DISMAN-EVENT-MIB) | mteHotTrigger, mteHotTargetName, mteHotContextName, mteHotOID, mteHotValue, cempMemPoolName, cempMemPoolHCUsed | The snmp-server enable traps memory-threshold command is used to enable the memory threshold notification. The mteHotOID is set to cempMemPoolHCUsed. The cempMemPoolName and cempMemPoolHCUsed objects are sent with the other objects. | 
| mteTriggerFired (DISMAN-EVENT-MIB) | mteHotTrigger, mteHotTargetName, mteHotContextName, mteHotOID, mteHotValue, ifHCInOctets, ifHCOutOctets, ifHighSpeed, entPhysicalName | The snmp-server enable traps interface-threshold command is used to enable the interface threshold notification. The entPhysicalName objects are sent with the other objects. | 
| natPacketDiscard (NAT-MIB) | ifIndex | The snmp-server enable traps nat packet-discard command is used to enable the NAT packet discard notification. This notification is rate limited for 5 minutes and is generated when IP packets are discarded by NAT because mapping space is not available. The ifIndex gives the ID of the mapped interface. | 
| ospfNbrStateChange | ospfRouterId, ospfNbrIpAddr, ospfNbrAddressLessIndex, ospfNbrRtrId, ospfNbrState | The snmp-server enable traps peer-flap command is used to enable transmission of OSPF peer-flap related trap. | 
| warmStart (SNMPv2-MIB) | — | The warmStart trap that occurs when the SNMP agent restarts for the first time. This trap also occurs when the agent restarts after a SNMP configuration change where, all the SNMP host configuration are removed and a fresh SNMP configuration is done. The snmp-server enable traps snmp warmstart command is used to enable and disable transmission of these traps. | 
| Note | For cluster and HA nodes, post a reload, if the interfaces reboot time exceeds 5 minutes (preset threshold), the trap is dropped. When the cluster and HA nodes have rebooted sucessfully, all other traps are sent as expected. | 
| Note | This notification is only sent in multimode when a security context is created or removed. | 
| Note | For ASA5585 models, the SNMP engine is changed to use the netsnmp version 5.8 library and the following OIDs are not available in the library:  | 
Interface Types and Examples
The interface types that produce SNMP traffic statistics include the following:
-  
                                 		  
                                 Logical—Statistics collected by the software driver, which are a subset of physical statistics.
-  
                                 		  
                                 Physical—Statistics collected by the hardware driver. Each physical named interface has a set of logical and physical statistics associated with it. Each physical interface may have more than one VLAN interface associated with it. VLAN interfaces only have logical statistics. Note 
 For a physical interface that has multiple VLAN interfaces associated with it, be aware that SNMP counters for ifInOctets and ifOutoctets OIDs match the aggregate traffic counters for that physical interface. 
-  
                                 		  
