---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-78-1
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-78.md
source_anchor: ""
source_lines: [1, 133]
sha256: 2251ba2de3a07042213d58f0788f4d790e121bd00b5597810f6bd0bfd3dc4c66
---

# Overview

Simple Network Management Protocol (SNMP) is an Internet-standard protocol for managing devices on IP networks. SNMP can be used to graph various data with tools such as CACTI, MRTG, or The Dude.

SNMP write support is only available for some OIDs. For supported OIDs SNMP v1, v2 or v3 write is supported.

SNMP will respond to the query on the interface SNMP request was received from forcing responses to have same source address as request destination sent to the router

SNMP tool collects data from different services running on the system. If, for some reason, communication between SNMP and some service is taking longer time than expected (30 seconds per service, 5 minutes for routing service), you will see a warning in the log stating "timeout while waiting for program" or "SNMP did not get OID data within expected time, ignoring OID". After that, this service will deny SNMP requests for a while before even trying to get requested data again.

This error has nothing to do with SNMP service itself. In most cases, such an error is printed when some slow or busy service is monitored through SNMP, and quite often, it is a service that should not be monitored through SNMP, and proper solution in such cases is to skip such OIDs on your monitoring tool.

# Quick Configuration

To enable SNMP in RouterOS:

You can also specify administrative contact information in the above settings. All SNMP data will be available to communities configured in the *community* menu.

# General Properties

| **Sub-menu:**`/snmp` | 
|---|

This sub menu allows to enable SNMP and to configure general settings.

| Property | Description | 
|---|---|
| **contact** (*string* ; Default:**""** ) | Contact information | 
| **enabled** (*yes \| no* ; Default:**no** ) | Used to disable/enable SNMP service | 
| **engine-id** (*string* ; Default:**""** ) | For SNMP v3, used as part of the identifier. You can configure the suffix part of the engine id using this argument. If the SNMP client is not capable to detect set engine-id value then this prefix hex has to be used 0x80003a8c04 | 
| **location** (*string* ; Default:**""** ) | Location information | 
| **trap-community** (*string* ; Default:**public** ) | Which communities configured in the *community* menu to use when sending out the trap. | 
| **trap-generators** (*interfaces \| start-trap* ; Default: ) | What action will generate traps:  | 
| **trap-interfaces** (*string \| all* ; Default: ) | List of interfaces that traps are going to be sent out. | 
| **trap-target** (*list of IP/IPv6* ; Default:**0.0.0.0** ) | IP (IPv4 or IPv6) addresses of SNMP data collectors that have to receive the trap | 
| **trap-version** (*1\|2\|3* ; Default:**1** ) | A version of SNMP protocol to use for trap | 
| **src-address** (*IPv4 or IPv6 address* ; Default:**::** ) | Force the router to always use the same IP source address for all of the SNMP messages | 
| **vrf** (*VRF name* ; default value:**main** ) | Set VRF on which service is listening for incoming connections | 

the engine-id field holds the suffix value of engine-id, usually, SNMP clients should be able to detect the value, as SNMP values, as read from the router. However, there is a possibility that this is not the case. In which case, the engine-ID value has to be set according to this rule: <engine-id prefix> + <hex-dump suffix>, so as an example, if you have set 1234 as suffix value you have to provide 80003a8c04 + 31323334, combined hex (the result) is 80003a8c0431323334

# Community Properties

| **Sub-menu:**`/snmp community` | 
|---|

This sub-menu allows to set up access rights for the SNMP data.

There is little security in v1 and v2c, just Clear text community string („username“) and the ability for Limiting access by IP address.

In the production environment, SNMP v3 should be used as that provides security - Authorization (User + Pass) with MD5/SHA1, Encryption with DES and AES).

Default settings only have one community named *public* without any additional security settings. These settings should be considered insecure and should be adjusted according to the required security profile.

## **Properties**

| Property | Description | 
|---|---|
| **address** (*IP/IPv6 address* ; Default:**0.0.0.0/0** ) | Addresses from which connections to SNMP server is allowed | 
| **authentication-password** (*string* ; Default:**""** )*sensitive* | Password used to authenticate the connection to the server (SNMPv3). Password must be at least 8 characters in length. | 
| **authentication-protocol** (*MD5 \| SHA1* ; Default:**MD5** ) | The protocol used for authentication (SNMPv3) | 
| **encryption-password** (*string* ; Default:**""** )*sensitive* | the password used for encryption (SNMPv3). Password must be at least 8 characters in length. | 
| **encryption-protocol** (*DES \| AES* ; Default:**DES** ) | encryption protocol to be used to encrypt the communication (SNMPv3). AES (see rfc3826) available since v6.16. | 
| **name** (*string* ; Default: ) | Name of the SNMP community. | 
| **read-access** (*yes \| no* ; Default:**yes** ) | Whether read access is enabled for this community | 
| **security** (*authorized \| none \| private* ; Default:**none** ) | Security levels:  | 
| **write-access** (*yes \| no* ; Default:**no** ) | Whether write access is enabled for this community | 

# Management information base (MIB)

The Management Information Base (MIB) is the database of information maintained by the agent that the manager can query. You can download the latest MikroTik RouterOS MIB file from here: https://mikrotik.com/download/tools

Used MIBs in RouterOS:

- MIKROTIK-MIB
- MIB-2
- HOST-RESOURCES-MIB
- IF-MIB
- IP-MIB
- IP-FORWARD-MIB
- IPV6-MIB
- BRIDGE-MIB
- DHCP-SERVER-MIB
- CISCO-AAA-SESSION-MIB
- ENTITY-MIB
- UPS-MIB
- SQUID-MIB

# Object identifiers (OID)

Each OID identifies a variable that can be read via SNMP. Although the MIB file contains all the needed OID values, you can also print individual OID information in the console with the **print oid** command at any menu level:

# Traps

SNMP traps enable the router to notify the data collector of interface changes and SNMP service status changes by sending traps. It is possible to send out traps with security features to support SNMPv1 (no security). SNMPv2 and variants and SNMPv3 with encryption and authorization.

For SNMPv2 and v3 you have to set up an appropriately configured community as a *trap-community* to enable required features (password or encryption/authorization).

# SNMP write

SNMP write allows changing router configuration with SNMP requests. Consider securing access to the router or to router's SNMP, when SNMP and write-access are enabled.

To change settings by SNMP requests, use the command below to allow SNMP to write for the selected community.

## System Identity

It's possible to change router system identity by SNMP set command.

- *snmpset* - SNMP application used for SNMP SET requests to set information on a network entity;
- *public* - router's community name;
- *192.168.0.0* - IP address of the router;
- *1.3.6.1.2.1.1.5.0* - SNMP value for router's identity;

SNMPset command above is equal to the RouterOS command:

## Reboot

It's possible to reboot the router with SNMP set command, you need to set the value for reboot SNMP settings, which is not equal to 0.

- **1.3.6.1.4.1.14988.1.1.7.1.0** , SNMP value for the router reboot;
- **s 1** , snmpset command to set value, value should not be equal to 0;

Reboot SNMPset command is equal to the RouterOS command:

## Run Script

SNMP write allows running scripts on the router from the **system script** menu when you need to set value for the SNMP setting of the script.

- **X** , script number, numeration starts from 1;
- **s 1** , snmpset command to set value, the value should not be equal to 0;

The same command on RouterOS:

