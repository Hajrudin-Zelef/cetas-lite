---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b-3
title: "c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["accelerator", "agent", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b.md
source_anchor: ""
source_lines: [177, 212]
sha256: adc6f75f0ce2dd10564afa0c95c5f2dd44ef86fa07ac37e1b4fc094ed37ff872
---

# c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b

| DISMAN-EVENT-MIB; OID:1.3.6.1.2.1.88 | mteTriggerTable, mteTriggerThresholdTable, mteObjectsTable, mteEventTable, mteEventNotificationTable | 
| DISMAN-EXPRESSION-MIB; OID:1.3.6.1.2.1.90 | expExpressionTable, expObjectTable, expValueTable | 
| ENTITY-SENSOR-MIB; OID: 1.3.6.1.2.1.99 | entPhySensorTable | 
| NAT-MIB; OID:1.3.6.1.2.1.123 | natAddrMapTable, natAddrMapIndex, natAddrMapName, natAddrMapGlobalAddrType, natAddrMapGlobalAddrFrom, natAddrMapGlobalAddrTo, natAddrMapGlobalPortFrom, natAddrMapGlobalPortTo, natAddrMapProtocol, natAddrMapAddrUsed, natAddrMapRowStatus | 
| CISCO-PTP-MIB; OID:1.3.6.1.4.1.9.9.760 | ciscoPtpMIBSystemInfo, cPtpClockDefaultDSTable, cPtpClockTransDefaultDSTable, cPtpClockPortTransDSTable | 
| CISCO-PROCESS-MIB; 1.3.6.1.4.1.9.9.109.1.1.1.1.7.1 1.3.6.1.4.1.9.9.109.1.1.1.1.7.2 to 1.3.6.1.4.1.9.9.109.1.1.1.1.7.(n+1) | cpmCPUTotal1minRev Associated parameters and values of cpmCPUTotal1minRev Examples:  | 
| Note | These three MIB OIDs can be used to track why remote access connections fail. | 
| Note | Not supported on the ASAv. | 
| Note | Provides information related to physical sensors, such as chassis temperature, fan RPM, power supply voltage, etc. Not supported on the ASAv platform. | 
| Note | Only MIBs corresponding to E2E Transparent Clock mode are supported. | 
Supported Traps (Notifications)
The following table lists the supported traps (notifications) and their associated MIBs.
| Table 5. Supported Traps (Notifications) |  |  | 
|---|---|---|
| Trap and MIB Name | Varbind List | Description | 
|---|---|---|
| authenticationFailure (SNMPv2-MIB) | — | For SNMP Version 1 or 2, the community string provided in the SNMP request is incorrect. For SNMP Version 3, a report PDU is generated instead of a trap if the auth or priv passwords or usernames are incorrect. The snmp-server enable traps snmp authentication command is used to enable and disable transmission of these traps. | 
| bgpBackwardTransition | bgpPeerLastError, bgpPeerState | The snmp-server enable traps peer-flap command is used to enable transmission of BGP peer-flap related trap. | 
| ccmCLIRunningConfigChanged (CISCO-CONFIG-MAN-MIB) | ccmHistoryRunningLastChanged, ccmHistoryEventTerminalType | The snmp-server enable traps config command is used to enable transmission of this trap. | 
| cefcFRUInserted (CISCO-ENTITY-FRU-CONTROL -MIB) | entPhysicalContainedIn | The snmp-server enable traps entity fru-insert command is used to enable this notification. | 
| cefcFRURemoved (CISCO-ENTITY-FRU-CONTROL -MIB) | entPhysicalContainedIn | The snmp-server enable traps entity fru-remove command is used to enable this notification. | 
| ceSensorExtThresholdNotification (CISCO-ENTITY-SENSOR-EXT -MIB) | entPhysicalName, entPhysicalDescr, entPhySensorValue, entPhySensorType, ceSensorExtThresholdValue | The snmp-server enable traps entity [power-supply-failure \| fan-failure \| cpu-temperature] command is used to enable transmission of the entity threshold notifications. This notification is sent for a power supply failure. The objects sent identify the fan and CPU temperature. The snmp-server enable traps entity fan-failure command is used to enable transmission of the fan failure trap.This trap does not apply to the Firepower 2100 series. The snmp-server enable traps entity power-supply-failure command is used to enable transmission of the power supply failure trap.This trap does not apply to the Firepower 2100 series. The snmp-server enable traps entity chassis-fan-failure command is used to enable transmission of the chassis fan failure trap. The snmp-server enable traps entity cpu-temperature command is used to enable transmission of the high CPU temperature trap. This trap does not apply to the Firepower 2100 series. The snmp-server enable traps entity power-supply-presence command is used to enable transmission of the power supply presence failure trap. The snmp-server enable traps entity power-supply-temperature command is used to enable transmission of the power supply temperature threshold trap. The snmp-server enable traps entity chassis-temperature command is used to enable transmission of the chassis ambient temperature trap. This trap does not apply to the Firepower 2100 series. The snmp-server enable traps entity accelerator-temperature command is used to enable transmission of the chassis accelerator temperature trap. | 
| cikeTunnelStart (CISCO-IPSEC-FLOW-MONITOR-MIB) | cikePeerLocalAddr, cikePeerRemoteAddr, cikeTunLifeTime | The snmp-server enable traps ikev2 start command is used to enable transmission of ikev2 start trap. | 
| cikeTunnelStop (CISCO-IPSEC-FLOW-MONITOR-MIB) | cikePeerLocalAddr, cikePeerRemoteAddr, cikeTunActiveTime | The snmp-server enable traps ikev2 stop command is used to enable transmission of ikev2 stop trap. | 
| cipSecTunnelStart (CISCO-IPSEC-FLOW-MONITOR -MIB) | cipSecTunLifeTime, cipSecTunLifeSize | The snmp-server enable traps ipsec start command is used to enable transmission of this trap. | 
| cipSecTunnelStop (CISCO-IPSEC-FLOW-MONITOR -MIB) | cipSecTunActiveTime | The snmp-server enable traps ipsec stop command is used to enable transmission of this trap. | 
| ciscoConfigManEvent (CISCO-CONFIG-MAN-MIB) | ccmHistoryEventCommandSource, ccmHistoryEventConfigSource, ccmHistoryEventConfigDestination | The snmp-server enable traps config command is used to enable transmission of this trap. | 
| ciscoRasTooManySessions (CISCO-REMOTE-ACCESS -MONITOR-MIB) | crasNumSessions, crasNumUsers, crasMaxSessionsSupportable, crasMaxUsersSupportable, crasThrMaxSessions | The snmp-server enable traps remote-access session-threshold-exceeded command is used to enable transmission of these traps. | 
| ciscoUFwFailoverStateChanged (CISCO-UNIFIED-FIREWALL-MIB) | gid, FOStatus | The snmp-server enable traps failover-state command is used to enable transmission of failover-state trap. | 
| clogMessageGenerated (CISCO-SYSLOG-MIB) | clogHistFacility, clogHistSeverity, clogHistMsgName, clogHistMsgText, clogHistTimestamp | Syslog messages are generated. The value of the clogMaxSeverity object is used to decide which syslog messages are sent as traps. The snmp-server enable traps syslog command is used to enable and disable transmission of these traps. | 
| clrResourceLimitReached (CISCO-L4L7MODULE-RESOURCE -LIMIT-MIB) | crlResourceLimitValueType, crlResourceLimitMax, clogOriginIDType, clogOriginID | The snmp-server enable traps connection-limit-reached command is used to enable transmission of the connection-limit-reached notification. The clogOriginID object includes the context name from which the trap originated. | 
| coldStart (SNMPv2-MIB) | — | The coldStart trap that occurs when the SNMP agent starts after the SNMP configuration. This trap also occurs when the agent starts after a system reboot. The snmp-server enable traps snmp coldstart command is used to enable and disable transmission of these traps. | 
| cpmCPURisingThreshold (CISCO-PROCESS-MIB) | cpmCPURisingThresholdValue, cpmCPUTotalMonIntervalValue, cpmCPUInterruptMonIntervalValue, cpmCPURisingThresholdPeriod, cpmProcessTimeCreated, cpmProcExtUtil5SecRev | The snmp-server enable traps cpu threshold rising command is used to enable transmission of the CPU threshold rising notification. The cpmCPURisingThresholdPeriod object is sent with the other objects. | 
| cufwClusterStateChanged (CISCO-UNIFIED-FIREWALL-MIB) | status | The snmp-server enable traps cluster-state command is used to enable transmission of cluster-state trap. | 
| entConfigChange (ENTITY-MIB) | — | The snmp-server enable traps entity config-change fru-insert fru-remove command is used to enable this notification. | 
| linkDown (IF-MIB) | ifIndex, ifAdminStatus, ifOperStatus | The linkdown trap for interfaces. The snmp-server enable traps snmp linkdown command is used to enable and disable transmission of these traps. | 
