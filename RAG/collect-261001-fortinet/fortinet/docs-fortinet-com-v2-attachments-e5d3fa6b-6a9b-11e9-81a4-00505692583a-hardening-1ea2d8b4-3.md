---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening-1ea2d8b4-3
title: "docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4.md
source_anchor: ""
source_lines: [276, 417]
sha256: 63abe2ec792183855a6579b0f3792c87d0782fc84433a7f536b1424f48e76f37
---

# docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4

FortiOS ports and protocols 13
Incoming ports
Purpose Protocol/Port
FortiAuthenticator RADIUS UDP/1812
FSSO TCP/8000
FortiGate HA Heartbeat ETH Layer 0x8890, 0x8891, and
0x8893
HA Synchronization TCP/703, UDP/703
FortiGuard Management TCP/541
AV/IPS UDP/9443
FortiManager AV/IPS Push UDP/9443
SSH CLI Management TCP/22
Management TCP/541
SNMP Poll UDP/161, UDP/162
FortiGuard Queries TCP/443
FortiPortal API communications
(FortiOS REST API, used for Wireless
Analytics)
TCP/443
Others Web Admin TCP/80, TCP/443
FSSO TCP/8000
Policy Override Authentication TCP/443, TCP/8008
FortiClient Portal TCP/8009
Policy Override Keepalive TCP/1000, TCP/1003
SSL VPN TCP/10443
3rd-Party Servers FSSO TCP/8000
Outgoing ports
Purpose Protocol/Port
FortiAnalyzer Syslog, OFTP, Registration,
Quarantine, Log & Report
TCP/514
FortiAuthenticator LDAP, PKI Authentication TCP or UDP/389
FortiCloud Registration, Quarantine, Log
& Report, Syslog
TCP/443
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

FortiOS ports and protocols 14
Outgoing ports
Purpose Protocol/Port
OFTP TCP/514
Management TCP/541
Contract Validation TCP/443
FortiGate HA Heartbeat ETH Layer 0x8890, 0x8891, and
0x8893
HA Synchronization TCP/703, UDP/703
FortiGuard AV/IPS Update TCP/443, TCP/8890
Cloud App DB TCP/9582
FortiGuard Queries UDP/53, UDP/8888
DNS UDP/53, UDP/8888
Registration TCP/80
Alert Email, Virus Sample TCP/25
Management, Firmware, SMS, FTM,
Licensing, Policy Override
TCP/443
Central Management, Analysis TCP/541
FortiManager Management TCP/541
IPv6 FGFM connection TCP/542
Log & Report TCP or UDP/514
Secure SNMP UDP/161, UDP/162
FortiGuard Queries TCP/8890, UDP/53
FortiSandbox OFTP TCP/514
Note that, while a proxy is configured, FortiGate uses the following URLs to access
the FortiGuard Distribution Network (FDN):
l update.fortiguard.net
l service.fortiguard.net
l support.fortinet.com
Closing open ports
You can close open ports by disabling the feature that opens them. For example, if FortiOS is not managing a FortiAP
then the CAPWAPfeature for managing FortiAPs can be disabled, closing the CAPWAPport.
The following sections of this document described a number of options for closing open ports:
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

FortiOS ports and protocols 15
l Use local-in policies to close open ports or restrict access on page 25
l Disable unused protocols on interfaces on page 24
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 16
Security best practices
This chapter describes some techniques and best practices that you can use to improve FortiOS security.
Install the FortiGate unit in a physically secure location
A good place to start with is physical security. Install your FortiGate in a secure location, such as a locked room or one
with restricted access. A restricted location prevents unauthorized users from getting physical access to the device.
If unauthorized users have physical access, they can disrupt your entire network by disconnecting your FortiGate (either
by accident or on purpose). They could also connect a console cable and attempt to log into the CLI. Also, when a
FortiGate unit reboots, a person with physical access can interrupt the boot process and install different firmware.
Register your product with Fortinet Support
You need to register your Fortinet product with Fortinet Support to receive customer services, such as firmware updates
and customer support. You must also register your product for FortiGuard services, such as up-to-date antivirus and IPS
signatures. To register your product the Fortinet Support website.
Keep your FortiOS firmware up to date
Always keep FortiOS up to date. The most recent version is the most stable and has the most bugs fixed and
vulnerabilities removed. Fortinet periodically updates the FortiGate firmware to include new features and resolve
important issues.
After you register your FortiGate, you can receive notifications on FortiGate GUI about firmware updates. You can
update the firmware directly from the GUI or by downloading firmware updates from the Fortinet Support website.
Before you install any new firmware, be sure to follow these steps:
l Review the release notes for the latest firmware release.
l Review the UpgradePath tool to determine the best path to take from your current version of FortiOS to the latest
version.
l Back up the current configuration.
Only FortiGate administrators who have read and write privileges can upgrade the FortiOS firmware.
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Security best practices 17
System administrator best practices
This section describes a collection of changes you can implement to make administrative access to the GUI and CLI
more secure.
Disable administrative access to the external (Internet- facing) interface
When possible, don’t allow administration access on the external (Internet-facing) interface.
To disable administrative access, go to Network > Interfaces,edit the external interface and disable HTTPS, PING,
HTTP, SSH, and TELNET under Administrative Access.
From the CLI:
config system interface
edit <external-interface-name>
unset allowaccess
end
Allow only HTTPS access to the GUI and SSH access to the CLI
For greater security never allow HTTP or Telnet administrative access to a FortiGate interface, only allow HTTPS and
SSH access. You can change these settings for individual interfaces by going to Network > Interfacesand adjusting
the administrative access to each interface.
From the CLI:
config system interface
edit <interface-name>
set allowaccess https ssh
end
Require TLS 1.2 for HTTPS administrator access
Use the following command to require TLS 1.2 for HTTPS administrator access to the GUI:
config system global
set admin-https-ssl-versions tlsv1-2
end
TLS 1.2 is currently the most secure SSL/TLS supported version for SSL-encrypted administrator access.
Re-direct HTTP GUI logins to HTTPS
Go to System > Settings > Administrator Settings and enable Redirect to HTTPS to make sure that all
attempted HTTP login connections are redirected to HTTPS.
From the CLI:
config system global
set admin-https-redirect enable
end
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

