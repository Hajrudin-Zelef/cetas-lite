---
id: collect-261001-cisco/cisco/c-en-us-support-docs-security-asa-5500-x-series-firewalls-212420-configure-firep-470bc69d-2
title: "c-en-us-support-docs-security-asa-5500-x-series-firewalls-212420-configure-firep-470bc69d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["accelerator"]
source: docs/RAG/collect-261001-cisco/c-en-us-support-docs-security-asa-5500-x-series-firewalls-212420-configure-firep-470bc69d.md
source_anchor: ""
source_lines: [106, 208]
sha256: 67d01c20ce42d9667358cd0a826191a30fe757d9286b464eb215f6852f551094
---

# c-en-us-support-docs-security-asa-5500-x-series-firewalls-212420-configure-firep-470bc69d

Encryption hardware device: Cisco ASA Crypto on-board accelerator (revision 0x1)
                             Boot microcode        : CNPx-MC-BOOT-2.00
                             SSL/IKE microcode     : CNPx-MC-SSL-SB-PLUS-0005
                             IPSec microcode       : CNPx-MC-IPSEC-MAIN-0026
                             Number of accelerators: 1
Baseboard Management Controller (revision 0x1) Firmware Version: 2.4
 0: Int: Internal-Data0/0    : address is a89d.21ce.fde6, irq 11
 1: Ext: GigabitEthernet0/0  : address is a89d.21ce.fdea, irq 10
 2: Ext: GigabitEthernet0/1  : address is a89d.21ce.fde7, irq 10
 3: Ext: GigabitEthernet0/2  : address is a89d.21ce.fdeb, irq 5
 4: Ext: GigabitEthernet0/3  : address is a89d.21ce.fde8, irq 5
 5: Ext: GigabitEthernet0/4  : address is a89d.21ce.fdec, irq 10
 6: Ext: GigabitEthernet0/5  : address is a89d.21ce.fde9, irq 10
 7: Int: Internal-Control0/0 : address is 0000.0001.0001, irq 0
 8: Int: Internal-Data0/1    : address is 0000.0001.0003, irq 0
   9: Ext: Management0/0       : address is a89d.21ce.fde6, irq 0
Management Interface Architecture (pre-7.4 releases)
The Management interface is divided into 2 logical interfaces: br1 (management0 on FPR2100/4100/9300 appliances) and diagnostic:
|  | Management - br1/management0 | Management - Diagnostic | 
| Purpose |  This interface is used in order to assign the FTD IP that is used for FTD/FMC communication. Terminates the sftunnel between FMC/FTD. Used as a source for rule-based syslogs. Provides SSH and HTTPS access to the FTD box. |  Provides remote access (for example, SNMP) to ASA engine. Used as a source for LINA-level syslogs, AAA, SNMP etc messages. | 
| Mandatory | Yes, since it is used for FTD/FMC communication (the sftunnel terminates on it) | No and it is not recommended to configure it. The recommendation is to use a data interface instead* (check the note below) | 
| Configure | This interface is configured during FTD installation (setup). Later you can modify the br1 settings as follows:  >configure network ipv4 manual 10.1.1.2 255.0.0.0 10.1.1.1Setting IPv4 network configuration. Network settings changed.  > Step 2. Update the FTD IP on FMC.  | The interface can be configured from FMC GUI: Navigate to Devices > Device Management, Select the Edit button and navigate to Interfaces  | 
| Restrict access |  By default, only the admin user can connect to the FTD br1 subinterface. To restrict SSH access is done with the use of the CLISH CLI   > configure ssh-access-list 10.0.0.0/8 | The access to the diagnostic interface can be controlled by FTD Devices > Platform Settings >  Secure Shell  and Devices > Platform Settings> HTTP  respectively  | 
| Verify | Method 1 - From FTD CLI:   > show network... =======[ br1 ]======= State : Enabled Channels : Management & Events Mode : MDI/MDIX : Auto/MDIX MTU : 1500 MAC Address : 18:8B:9D:1E:CA:7B ----------------------[ IPv4 ]----- Configuration : Manual Address : 10.1.1.2 Netmask : 255.0.0.0 Broadcast : 10.1.1.255 ----------------------[ IPv6 ]-----  Method 2 – From FMC GUI Devices > Device Management > Device > Management | Method 1 - From LINA CLI:   firepower# show interface ip brief.. Management1/1 192.168.1.1 YES unset up up  firepower# show run interface m1/1 ! interface Management1/1 management-only nameif diagnostic security-level 0 ip address 192.168.1.1 255.255.255.0 Method 2 – From FMC GUI Navigate to Devices > Device Management, select the Edit button and navigate to Interfaces | 
FTD Logging
- When a user configures FTD logging from Platform Settings, the FTD generates Syslog messages (same as on classic ASA) and can use any Data Interface as a source (includes the Diagnostic). An example of a syslog message that is generated in that case:
May 30 2016 19:25:23 firepower : %ASA-6-302020: Built inbound ICMP connection for faddr 192.168.75.14/1 gaddr 192.168.76.14/0 laddr 192.168.76.14/0
- On the other hand, when Access Control Policy (ACP) Rule-level logging is enabled the FTD originates these logs through the br1 logical interface as a source. The logs are originated from the FTD br1 subinterface:
Manage FTD with FDM (On-Box Management)
As from 6.1 version, an FTD that is installed on ASA5500-X appliances can be managed either by FMC (off-box management) or by Firepower Device Manager (FDM) (on-box management).
Output from FTD CLISH when the device is managed by FDM:
> show managers
Managed locally.
>
FDM it uses the br1 logical interface. This can be visualized as:
From FDM UI the management interface is accessible from the Device Dashboard > System Settings > Device Management IP:
Management Interface on FTD Firepower Hardware Appliances
FTD can be also installed on Firepower 2100, 4100 and 9300 hardware appliances. The Firepower chassis runs its own OS called FXOS while the FTD is installed on a module/blade.
FPR21xx appliance
FPR41xx appliance
FPR9300 appliance
On FPR4100/9300 this interface is only for the chassis management and cannot be used/shared with the FTD software that runs inside the FP module. For the FTD module allocate a separate data interface that for the FTD management.
On FPR2100 this interface is shared between the chassis (FXOS) and the FTD logical appliance:
> show network
===============[ System Information ]===============
Hostname                  : ftd623
Domains                   : cisco.com
DNS Servers               : 192.168.200.100
                            8.8.8.8
Management port           : 8305
IPv4 Default route
  Gateway                 : 10.62.148.129
==================[ management0 ]===================
State                     : Enabled
Channels                  : Management & Events
Mode                      : Non-Autonegotiation
MDI/MDIX                  : Auto/MDIX
MTU                       : 1500
MAC Address               : 70:DF:2F:18:D8:00
----------------------[ IPv4 ]----------------------
Configuration             : Manual
Address                   : 10.62.148.179
Netmask                   : 255.255.255.128
Broadcast                 : 10.62.148.255
----------------------[ IPv6 ]----------------------
Configuration             : Disabled
> connect fxos
Cisco Firepower Extensible Operating System (FX-OS) Software
...
firepower#
This screenshot is from Firepower Chassis Manager (FCM) UI on FPR4100 where a separate interface for FTD managment is allocated. In this example, Ethernet1/3 is chosen as the FTD management interface: p1
This can also be seen from the Logical Devices tab:p2
On FMC the interface is shown as diagnostic: p3
CLI Verification 
FP4100# connect module 1 console
Firepower-module1>connect ftd
Connecting to ftd console... enter exit to return to bootCLI
>
> show interface
… output omitted …
Interface Ethernet1/3 "diagnostic", is up, line protocol is up
  Hardware is EtherSVI, BW 10000 Mbps, DLY 1000 usec
        MAC address 5897.bdb9.3e0e, MTU 1500
        IP address unassigned
  Traffic Statistics for "diagnostic":
        1304525 packets input, 63875339 bytes
        0 packets output, 0 bytes
        777914 packets dropped
      1 minute input rate 2 pkts/sec,  101 bytes/sec
      1 minute output rate 0 pkts/sec,  0 bytes/sec
      1 minute drop rate, 1 pkts/sec
      5 minute input rate 2 pkts/sec,  112 bytes/sec
      5 minute output rate 0 pkts/sec,  0 bytes/sec
      5 minute drop rate, 1 pkts/sec
        Management-only interface. Blocked 0 through-the-device packets
… output omitted …
>
Integrate FTD with FMC - Management Scenarios
These are some of the deployment options that allows to manage FTD that runs on ASA5500-X devices from FMC.
Scenario 1. FTD and FMC on the same subnet.
This is the simplest deployment. As seen in the figure, the FMC is on the same subnet as the FTD br1 interface:
Scenario 2. FTD and FMC on different subnets. Control-plane does not go through the FTD.
In this deployment, the FTD must have a route towards the FMC and vice versa. On FTD the next hop is a L3 device (router):
