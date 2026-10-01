---
id: collect-261001-cisco/cisco/c-en-us-support-docs-security-asa-5500-x-series-firewalls-212420-configure-firep-470bc69d-1
title: "c-en-us-support-docs-security-asa-5500-x-series-firewalls-212420-configure-firep-470bc69d"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2016-03-28"]
keywords: ["accelerator"]
source: docs/RAG/collect-261001-cisco/c-en-us-support-docs-security-asa-5500-x-series-firewalls-212420-configure-firep-470bc69d.md
source_anchor: ""
source_lines: [1, 105]
sha256: 5e2b224a46d204d3d4000c16e8a506440e29630198894b08a1044f2451594228
---

# c-en-us-support-docs-security-asa-5500-x-series-firewalls-212420-configure-firep-470bc69d

Introduction
This document describes the operation and configuration of the Management Interface on Cisco Secure Firewall Threat Defense (formelly knowns as Firepower Threat Defense (FTD)).
Note: FTD software versions 7.4 and later support merged Management and Diagnostic interfaces, also known as the Converged Management Interface (CMI).
CMI is enabled by default, and can be disabled by users. For more details see section "Merged management and diagnostic interfaces" at https://www.cisco.com/c/en/us/td/docs/security/secure-firewall/management-center/device-config/770/management-center-device-config-77/get-started-device-management.html.
The information covered in this document is mainly applicable to firewalls with disabled CMI (unmerged diagnostic and management interfaces).
Prerequisites
Requirements
There are no specific requirements for this document.
Components Used
- FTD that runs on ASA5508-X hardware appliance
- FTD that runs on ASA5512-X hardware appliance
- FTD that runs on FPR9300 hardware appliance
- FMC that runs on 6.1.0 (build 330)
The information in this document was created from the devices in a specific lab environment. All of the devices used in this document started with a cleared (default) configuration. If your network is live, ensure that you understand the potential impact of any command.
Background Information
The purpose of this document is to demonstrate:
- FTD Management interface architecture on ASA5500-X devices
- FTD Management interface when FDM is used
- FTD Management interface on FP41xx/FP9300 series
- The Management interface on Cisco Secure Firewall devices (1xxx, 12xx, 31xx, 42xx)
- FTD/Firepower Management Center (FMC) integration scenarios
Configure
Post-7.4 Software Releases
As of 7.4, by default the firewall uses the Converged Management Interface (CMI) mode:
FTD3100# show management-interface convergence
management-interface convergence
NAT is used as an internal implementation of CMI:
- An internal private IPv4 address (203.0.113.130) is automatically configured in data plane for IPv4.
- An internal private IPv6 address (fd00:0:1:1::2) is automatically configured in data plane for IPv6.
FTD3100# show interface ip brief       
Interface                  IP-Address      OK?           Method Status      Protocol
Internal-Data0/1           unassigned      YES           unset  up          up 
Port-channel5              192.0.2.1       YES           unset  up          up 
Ethernet1/1                unassigned      YES           unset  admin down  down
Ethernet1/2                unassigned      YES           unset  admin down  down
Ethernet1/5                unassigned      YES           unset  admin down  down
Ethernet1/6                unassigned      YES           unset  admin down  down
Ethernet1/7                unassigned      YES           unset  admin down  down
Ethernet1/8                unassigned      YES           unset  admin down  down
Ethernet1/9                unassigned      YES           unset  admin down  down
Ethernet1/10               unassigned      YES           unset  admin down  down
Ethernet1/11               unassigned      YES           unset  admin down  down
Ethernet1/12               unassigned      YES           unset  admin down  down
Ethernet1/13               unassigned      YES           unset  admin down  down
Ethernet1/14               unassigned      YES           unset  admin down  down
Ethernet1/15               unassigned      YES           unset  admin down  down
Ethernet1/16               unassigned      YES           unset  admin down  down
Internal-Control1/1        unassigned      YES           unset  up          up 
Internal-Data1/1           169.254.1.1     YES           unset  up          up 
Internal-Data1/2           unassigned      YES           unset  up          up 
Management1/1              203.0.113.130   YES           unset  up          up  <----
For more details refer to this document https://www.cisco.com/c/en/us/support/docs/security/secure-firewall-threat-defense/222680-claifry-the-purpose-of-ip-address-203-0.html
Pre-7.4 Releases
Management Interface on ASA 5500-X Devices
Note: Most of the ASA 5500-X devices have been announced as End-of-Life and End-of-Sale: https://www.cisco.com/c/en/us/products/security/asa-5500-series-next-generation-firewalls/eos-eol-notice-listing.html
When an FTD image is installed on a ASA 55xx device, the management interface is shown as Management1/1. On 5512/15/25/45/55-X devices this becomes Management0/0. From the FTD Command Line Interface (CLI) this can be verified in the show tech-support output.
Connect to the FTD console and run the command:
> show tech-support
-----------------[ BSNS-ASA5508-1 ]-----------------
Model                     : Cisco ASA5508-X Threat Defense (75) Version 6.1.0 (Build 330)
UUID                      : 04f55302-a4d3-11e6-9626-880037a713f3
Rules update version      : 2016-03-28-001-vrt
VDB version               : 270
----------------------------------------------------
Cisco Adaptive Security Appliance Software Version 9.6(2)
Compiled on Tue 23-Aug-16 19:42 PDT by builders
System image file is "disk0:/os.img"
Config file at boot was "startup-config"
firepower up 13 hours 43 mins
Hardware:   ASA5508, 8192 MB RAM, CPU Atom C2000 series 2000 MHz, 1 CPU (8 cores)
Internal ATA Compact Flash, 8192MB
BIOS Flash M25P64 @ 0xfed01000, 16384KB
Encryption hardware device : Cisco ASA Crypto on-board accelerator (revision 0x1)
                             Number of accelerators: 1
 1: Ext: GigabitEthernet1/1  : address is d8b1.90ab.c852, irq 255
 2: Ext: GigabitEthernet1/2  : address is d8b1.90ab.c853, irq 255
 3: Ext: GigabitEthernet1/3  : address is d8b1.90ab.c854, irq 255
 4: Ext: GigabitEthernet1/4  : address is d8b1.90ab.c855, irq 255
 5: Ext: GigabitEthernet1/5  : address is d8b1.90ab.c856, irq 255
 6: Ext: GigabitEthernet1/6  : address is d8b1.90ab.c857, irq 255
 7: Ext: GigabitEthernet1/7  : address is d8b1.90ab.c858, irq 255
 8: Ext: GigabitEthernet1/8  : address is d8b1.90ab.c859, irq 255
 9: Int: Internal-Data1/1    : address is d8b1.90ab.c851, irq 255
10: Int: Internal-Data1/2    : address is 0000.0001.0002, irq 0
11: Int: Internal-Control1/1 : address is 0000.0001.0001, irq 0
12: Int: Internal-Data1/3    : address is 0000.0001.0003, irq 0
13:   Ext: Management1/1       : address is d8b1.90ab.c851, irq 0
14: Int: Internal-Data1/4    : address is 0000.0100.0001, irq 0
ASA5512-X:
> show tech-support
-------------------[ FTD5512-1 ]--------------------
Model                     : Cisco ASA5512-X Threat Defense (75) Version 6.1.0 (Build 330)
UUID                      : 8608e98e-f0e9-11e5-b2fd-b649ba0c2874
Rules update version      : 2016-03-28-001-vrt
VDB version               : 270
----------------------------------------------------
Cisco Adaptive Security Appliance Software Version 9.6(2)
Compiled on Fri 18-Aug-16 15:08 PDT by builders
System image file is "disk0:/os.img"
Config file at boot was "startup-config"
firepower up 4 hours 37 mins
Hardware:   ASA5512, 4096 MB RAM, CPU Clarkdale 2793 MHz, 1 CPU (2 cores)
            ASA: 1764 MB RAM, 1 CPU (1 core)
Internal ATA Compact Flash, 4096MB
BIOS Flash MX25L6445E @ 0xffbb0000, 8192KB
