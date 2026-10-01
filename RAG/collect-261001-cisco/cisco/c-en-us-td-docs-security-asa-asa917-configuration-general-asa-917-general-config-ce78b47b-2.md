---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b-2
title: "c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b"
domain: cisco
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "aws", "ethernet", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b.md
source_anchor: ""
source_lines: [101, 176]
sha256: c4cf789e8f0dde954e6131b00a43df4f09d9078997e72459f2bfc3c8017bf909
---

# c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b

[70]    1.3.6.1.2.1.31.1.1.1.6. ifHCInOctets
--More--
SNMP Object Identifiers
Each Cisco system-level product has an SNMP object identifier (OID) for use as a MIB-II sysObjectID. The CISCO-PRODUCTS-MIB and the CISCO-ENTITY-VENDORTYPE-OID-MIB includes the OIDs that can be reported in the sysObjectID object in the SNMPv2-MIB, Entity Sensor MIB and Entity Sensor Threshold Ext MIB. You can use this value to identify the model type. The following table lists the sysObjectID OIDs for ASA and ISA models.
| Table 2. SNMP Object Identifiers |  |  | 
|---|---|---|
| Product Identifier | sysObjectID | Model Number | 
|---|---|---|
| ASAv | ciscoASAv (ciscoProducts 1902) | Cisco Adaptive Security Virtual Appliance (ASAv) | 
| ASAv System Context | ciscoASAvsy (ciscoProducts 1903) | Cisco Adaptive Security Virtual Appliance (ASAv) System Context | 
| ASAv Security Context | ciscoASAvsc (ciscoProducts 1904) | Cisco Adaptive Security Virtual Appliance (ASAv) Security Context. | 
| ISA 30004C Industrial Security Appliance | ciscoProducts 2268 | ciscoISA30004C | 
| CISCO ISA30004C with 4 GE Copper Security Context | ciscoProducts 2139 | ciscoISA30004Csc | 
| CISCO ISA30004C with 4 GE Copper System Context | ciscoProducts 2140 | ciscoISA30004Csy | 
| ISA 30002C2F Industrial Security Appliance | ciscoProducts 2267 | ciscoISA30002C2F | 
| CISCO ISA30002C2F with 2 GE Copper ports + 2 GE Fiber Security Context | ciscoProducts 2142 | ciscoISA30002C2Fsc | 
| CISCO ISA30002C2F with 2 GE Copper ports + 2 GE Fiber System Context | ciscoProducts 2143 | ciscoISA30002C2Fsy | 
| Cisco Industrial Security Appliance (ISA) 30004C Chassis | cevChassis 1677 | cevChassisISA30004C | 
| Cisco Industrial Security Appliance (ISA) 30002C2F Chassis | cevChassis 1678 | cevChassisISA30002C2F | 
| Central Processing Unit Temperature Sensor for ISA30004C Copper SKU | cevSensor 187 | cevSensorISA30004CCpuTempSensor | 
| Central Processing Unit Temperature Sensor for ISA30002C2F Fiber | cevSensor 189 | cevSensorISA30002C2FCpuTempSensor | 
| Processor Card Temperature Sensor for ISA30004C Copper SKU | cevSensor 192 | cevSensorISA30004CPTS | 
| Processor Card Temperature Sensor for ISA30002C2F Fiber SKU | cevSensor 193 | cevSensorISA30002C2FPTS | 
| Power Card Temperature Sensor for ISA30004C Copper SKU | cevSensor 197 | cevSensorISA30004CPowercardTS | 
| Power Card Temperature Sensor for ISA30002C2F Fiber SKU | cevSensor 198 | cevSensorISA30002C2FPowercardTS | 
| Port Card Temperature Sensor for ISA30004C | cevSensor 199 | cevSensorISA30004CPortcardTS | 
| Port Card Temperature Sensor for ISA30002C2F | cevSensor 200 | cevSensorISA30002C2FPortcardTS | 
| Central Processing Unit for ISA30004C Copper SKU | cevModuleCpuType 329 | cevCpuISA30004C | 
| Central Processing Unit for ISA30002C2F Fiber SKU | cevModuleCpuType 330 | cevCpuISA30002C2F | 
| Modules ISA30004C, ISA30002C2F | cevModule 111 | cevModuleISA3000Type | 
| 30004C Industrial Security Appliance Solid State Drive | cevModuleISA3000Type 1 | cevModuleISA30004CSSD64 | 
| 30002C2F Industrial Security Appliance Solid State Drive | cevModuleISA3000Type 2 | cevModuleISA30002C2FSSD64 | 
| Cisco ISA30004C/ISA30002C2F Hardware Bypass | cevModuleISA3000Type 5 | cevModuleISA3000HardwareBypass | 
| Cisco Secure Firewall 240P | ciscoCsf240P (ciscoProducts 3424) | ciscoCsf240P | 
| FirePOWER 4140 Security Appliance, 1U with embedded security module 36 | ciscoFpr4140K9 (ciscoProducts 2293) | FirePOWER 4140 | 
| FirePOWER 4120 Security Appliance, 1U with embedded security module 24 | ciscoFpr4120K9 (ciscoProducts 2294) | FirePOWER 4120 | 
| FirePOWER 4110 Security Appliance, 1U with embedded security module 12 | ciscoFpr4110K9 (ciscoProducts 2295) | FirePOWER 4110 | 
| FirePOWER 4110 Security Module 12 | ciscoFpr4110SM12 (ciscoProducts 2313) | FirePOWER 4110 Security Module 12 | 
| FirePOWER 4120 Security Module 24 | ciscoFpr4120SM24 (ciscoProducts 2314) | FirePOWER 4110 Security Module 24 | 
| FirePOWER 4140 Security Module 36 | ciscoFpr4140SM36 (ciscoProducts 2315) | FirePOWER 4110 Security Module 36 | 
| FirePOWER 4110 Chassis | cevChassis 1714 | cevChassisFPR4110 | 
| FirePOWER 4120 Chassis | cevChassis 1715 | cevChassisFPR4120 | 
| FirePOWER 4140 Chassis | cevChassis 1716 | cevChassisFPR4140 | 
| FirePOWER 4K Fan Bay | cevContainer 363 | cevContainerFPR4KFanBay | 
| FirePOWER 4K Power Supply Bay | cevContainer 364 | cevContainerFPR4KPowerSupplyBay | 
| FirePOWER 4120 Supervisor Module | cevModuleFPRType 4 | cevFPR4120SUPFixedModule | 
| FirePOWER 4140 Supervisor Module | cevModuleFPRType 5 | cevFPR4140SUPFixedModule | 
| FirePOWER 4110 Supervisor Module | cevModuleFPRType 7 | cevFPR4110SUPFixedModule | 
| Cisco FirePOWER 4110 Security Appliance, Threat Defense | cevChassis 1787 | cevChassisCiscoFpr4110td | 
| Cisco FirePOWER 4120 Security Appliance, Threat Defense | cevChassis 1788 | cevChassisCiscoFpr4120td | 
| Cisco FirePOWER 4140 Security Appliance, Threat Defense | cevChassis 1789 | cevChassisCiscoFpr4140td | 
| Cisco Firepower 9000 Security Module 24, Threat Defense | cevChassis 1791 | cevChassisCiscoFpr9000SM24td | 
| Cisco Firepower 9000 Security Module 24 NEBS, Threat Defense | cevChassis 1792 | cevChassisCiscoFpr9000SM24Ntd | 
| Cisco Firepower 9000 Security Module 36, Threat Defense | cevChassis 1793 | cevChassisCiscoFpr9000SM36td | 
| Cisco Secure Firewall Threat Defense Virtual, VMware | cevChassis 1795 | cevChassisCiscoFTDVVMW | 
| Cisco Firewall Threat Defense Virtual, AWS | cevChassis 1796 | cevChassisCiscoFTDVAWS | 
Physical Vendor Type Values
Each Cisco chassis or standalone system has a unique type number for SNMP use. The entPhysicalVendorType OIDs are defined in the CISCO-ENTITY-VENDORTYPE-OID-MIB. This value is returned in the entPhysicalVendorType object from the ASA, ASAv, or ASASM SNMP agent. You can use this value to identify the type of component (module, power supply, fan, sensors, CPU, and so on). The following table lists the physical vendor type values for the ASA models.
| Table 3. Physical Vendor Type Values |  | 
|---|---|
| Item | entPhysicalVendorType OID Description | 
|---|---|
| Gigabit Ethernet port | cevPortGe (cevPort 109) | 
| Cisco Adaptive Security Virtual Appliance | cevChassisASAv (cevChassis 1451) | 
Supported Tables and Objects in MIBs
The following table lists the supported tables and objects for the specified MIBs.
In multi-context mode, these tables and objects provide information for a single context. If you want data across contexts, you need to sum them. For example, to get overall memory usage, sum the cempMemPoolHCUsed values for each context.
| Table 4. Supported Tables and Objects in MIBs |  | 
|---|---|
| MIB Name and OID | Supported Tables and Objects | 
|---|---|
| CISCO-ENHANCED-MEMPOOL-MIB; OID:1.3.6.1.4.1.9.9.221 | cempMemPoolTable, cempMemPoolIndex, cempMemPoolType, cempMemPoolName, cempMemPoolAlternate, cempMemPoolValid. For a 32-bit memory system, poll using the 32-bit memory counters—cempMemPoolUsed, cempMemPoolFree,cempMemPoolUsedOvrflw, cempMemPoolFreeOvrflw, cempMemPoolLargestFree, cempMemPoolLowestFree, cempMemPoolUsedLowWaterMark, cempMemPoolAllocHit, cempMemPoolAllocMiss, cempMemPoolFreeHit, cempMemPoolFreeMiss, cempMemPoolLargestFreeOvrflw, cempMemPoolLowestFreeOvrflw, cempMemPoolUsedLowWaterMarkOvrflw, cempMemPoolSharedOvrflw. For a 64-bit memory system, poll using the 64-bit memory counters—cempMemPoolHCUsed, cempMemPoolHCFree, cempMemPoolHCLargestFree, cempMemPoolHCLowestFree, cempMemPoolHCUsedLowWaterMark, cempMemPoolHCShared | 
| CISCO-REMOTE-ACCESS-MONITOR-MIB; OID:1.3.6.1.4.1.9.9.392 | crasNumTotalFailures, crasNumSetupFailInsufResources, crasNumAbortedSessions | 
| CISCO-ENTITY-SENSOR-EXT-MIB; OID:1.3.6.1.4.1.9.9.745 | ceSensorExtThresholdTable | 
| CISCO-L4L7MODULE-RESOURCE-LIMIT-MIB; OID:1.3.6.1.4.1.9.9.480 | ciscoL4L7ResourceLimitTable | 
| CISCO-TRUSTSEC-SXP-MIB; OID:1.3.6.1.4.1.9.9.720 | ctsxSxpGlobalObjects, ctsxSxpConnectionObjects, ctsxSxpSgtObjects | 
