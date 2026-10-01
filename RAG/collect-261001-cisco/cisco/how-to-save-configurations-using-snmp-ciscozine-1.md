---
id: collect-261001-cisco/cisco/how-to-save-configurations-using-snmp-ciscozine-1
title: "how-to-save-configurations-using-snmp-ciscozine"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention"]
source: docs/RAG/collect-261001-cisco/how-to-save-configurations-using-snmp-ciscozine.md
source_anchor: ""
source_lines: [1, 174]
sha256: a671b41e8d85ee3d9cba42a990eb0ff171597d116d4fb760f113c733691eb107
---

# how-to-save-configurations-using-snmp-ciscozine

Everyone knows there are software to **get the configuration using SNMP**; but how can you copy the configuration if you don’t have any tool?

Let me explain what is SNMP before show you how to implement it.

Simple Network Management Protocol (**SNMP**) is an “Internet-standard protocol for managing devices on IP networks”. Devices that typically support SNMP include routers, switches, servers, workstations, printers, modem racks, and more. It is used mostly in **network management** **systems** to monitor network-attached devices for conditions that warrant administrative attention. SNMP uses an extensible design, where the **available information** is defined **by management information bases** (MIBs). MIBs describe the structure of the management data of a device subsystem; they use a hierarchical namespace containing **object identifiers** (OID). Each OID identifies a variable that can be read or set via SNMP.

In this tutorial, I will use the


OK, let’s start :)

First of all, check if your PC has the SNMP suite; if not, install the net-snmp software (http://net-snmp.sourceforge.net/).

So, to retrieve the configuration, run these commands:

snmpset -c [snmp-community-string] -v 2c [ip-device] 1.3.6.1.4.1.9.9.96.1.1.1.1.2.336 i 1
snmpset -c [snmp-community-string] -v 2c [ip-device] 1.3.6.1.4.1.9.9.96.1.1.1.1.3.336 i 4
snmpset -c [snmp-community-string] -v 2c [ip-device] 1.3.6.1.4.1.9.9.96.1.1.1.1.4.336 i 1
snmpset -c [snmp-community-string] -v 2c [ip-device] 1.3.6.1.4.1.9.9.96.1.1.1.1.5.336 a [ip-tftp-server]
snmpset -c [snmp-community-string] -v 2c [ip-device] 1.3.6.1.4.1.9.9.96.1.1.1.1.6.336 s [file-name]
snmpset -c [snmp-community-string] -v 2c [ip-device] 1.3.6.1.4.1.9.9.96.1.1.1.1.14.336 i 1

where:

- snmp-community-string is the community key of your cisco router.
- ip-device is the ip address of your Cisco device.
- ip-tftp-server is the ip address of your tftp server.
- file-name is the name where the running configuration will be saved.

**Remember:** The command syntax are the same on linux or windows devices.

**Example #1:** Suppose you have a Cisco Router (192.168.1.1) with the “c1sc0zine” community string, a tftp server with 192.168.1.100 ip address and you want to save the running configuration to the tftp-server (in the ciscozine.txt file).


The client sends the “snmp set” commands to the router, then the router sends the running configuration to the tftp server using the tftp protocol.

Below the command list:

snmpset -c c1sc0zine -v 2c 192.168.1.1 1.3.6.1.4.1.9.9.96.1.1.1.1.2.336 i 1
snmpset -c c1sc0zine -v 2c 192.168.1.1 1.3.6.1.4.1.9.9.96.1.1.1.1.3.336 i 4
snmpset -c c1sc0zine -v 2c 192.168.1.1 1.3.6.1.4.1.9.9.96.1.1.1.1.4.336 i 1
snmpset -c c1sc0zine -v 2c 192.168.1.1 1.3.6.1.4.1.9.9.96.1.1.1.1.5.336 a 192.168.1.100
snmpset -c c1sc0zine -v 2c 192.168.1.1 1.3.6.1.4.1.9.9.96.1.1.1.1.6.336 s ciscozine.txt
snmpset -c c1sc0zine -v 2c 192.168.1.1 1.3.6.1.4.1.9.9.96.1.1.1.1.14.336 i 1

But what is the meaning of the last part of each command? See you below an explanation:

1.3.6.1.4.1.9.9.96.1.1.1.1.2.336 i 1

ccCopyProtocol: The protocol file transfer protocol that should be used to copy the configuration file over the network. If the config file transfer is to occur locally on the SNMP agent, the method of transfer is left up to the implementation, and is not restricted to the protocols below. The object can be:

1. **tftp**
2. ftp
3. rcp
4. scp
5. sftp


1.3.6.1.4.1.9.9.96.1.1.1.1.3.336 i 4

ccCopySourceFileType: Specifies the type of file to copy from. The object can be:

1. networkFile
2. iosFile
3. startupConfig
4. **runningConfig**
5. terminal
6. fabricStartupConfig


1.3.6.1.4.1.9.9.96.1.1.1.1.4.336 i 1

ccCopyDestFileType: specifies the type of file to copy to. The object can be:

1. **networkFile**
2. iosFile
3. startupConfig
4. runningConfig
5. terminal
6. fabricStartupConfig


1.3.6.1.4.1.9.9.96.1.1.1.1.5.336 a 192.168.1.100

ccCopyServerAddress: The IP address of the TFTP server to copy the configuration file. In this case the tftp server is 192.168.1.100.


snmpset -c c1sc0zine -v 2c 192.168.1.1 1.3.6.1.4.1.9.9.96.1.1.1.1.6.336 s ciscozine.txt

ccCopyFileName: The file name (including the path, if applicable) of the file.


1.3.6.1.4.1.9.9.96.1.1.1.1.14.336 i 1

ccCopyEntryRowStatus: The status of this table entry. Once the entry status is set to active, the associated entry cannot be modified until the request completes (ccCopyState transitions to ‘successful’ or ‘failed’ state). The object can be:

1. **active**
2. notInService
3. notReady
4. createAndGo
5. createAndWait
6. destroy

**Note:** When you run this command, the router will send the running configuration to the ftp-server!


**Remember:** The object types can be:

- i: INTEGER
- u: unsigned INTEGER
- t: TIMETICKS
- a: IPADDRESS
- o: OBJID
- s: STRING
- x: HEX STRING
- d: DECIMAL STRING
- b: BITS
- U: unsigned int64
- I: signed int64
- F: float
- D: double


**Example #2:** What can you do if you are not able to install/use the net-snmp suite on your pc? It’s simple! Use your router!


Not everyone knows that the IOS has an embedded hidden commands for snmp operation. In fact, if you want to save the startup configuration to the tftp server, run these commands on your router:

snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.2.333 integer 1
snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.3.333 integer 3
snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.4.333 integer 1
snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.5.333 ip-address 192.168.1.100
snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.6.333 string ciscozine.txt
snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.14.333 integer 1

as you notice, the last value of the second line is “3” and not “4”; this because I want save the startup configuration, not the running configuration!

The output will be:

Ciscozine-test#snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.2.333 integer 1
SNMP Response: reqid 8, errstat 0, erridx 0 
 ccCopyTable.1.2.333 = 1
Ciscozine-test#
Ciscozine-test#snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.3.333 integer 3
SNMP Response: reqid 9, errstat 0, erridx 0 
 ccCopyTable.1.3.333 = 3
Ciscozine-test#
Ciscozine-test#snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.4.333 integer 1
SNMP Response: reqid 10, errstat 0, erridx 0 
 ccCopyTable.1.4.333 = 1
Ciscozine-test#
Ciscozine-test#snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.5.333 ip-address 192.168.1.100
SNMP Response: reqid 11, errstat 0, erridx 0 
 ccCopyTable.1.5.333 = 192.168.1.100
Ciscozine-test#
Ciscozine-test#snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.6.333 string ciscozine.txt
SNMP Response: reqid 12, errstat 0, erridx 0 
 ccCopyTable.1.6.333 = ciscozine.txt
Ciscozine-test#
Ciscozine-test#snmp set v2c 192.168.1.1 c1sc0zine oid 1.3.6.1.4.1.9.9.96.1.1.1.1.14.333 integer 1
SNMP Response: reqid 13, errstat 0, erridx 0 
 ccCopyTable.1.14.333 = 1
Ciscozine-test#

… and your tftp server will receive the startup configuration!

**Note:** What is the **last OID element** (in these two examples the number “333”)? It’s a **random number**! Each time you copy to or from a device using SNMP, choose a random number. **This number creates a row instance**. It must be the same everywhere in your command. Once you use a specific number, it can not be used again before it times out. **The timeout is five minutes**. If you use the same number within the five minutes, you get an error (SNMP: Inconsistent value.)

References:

what if i don’t have dedicated tftp server in my LAN and i have installed cisco tftp server software on my personal laptop

so it’ll work????

