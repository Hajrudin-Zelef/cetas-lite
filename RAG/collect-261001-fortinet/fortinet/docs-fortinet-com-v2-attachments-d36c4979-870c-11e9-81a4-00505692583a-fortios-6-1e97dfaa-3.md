---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6-1e97dfaa-3
title: "docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent", "distribution", "incident"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa.md
source_anchor: ""
source_lines: [257, 404]
sha256: 1677fb88806f265d5975f82b16efb73316a8f6e43796f6f9a0e92b69109c6440
---

# docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa

Building security into FortiOS 11
l belong to a locally attached subnet (local interface), or
l be in the routing domain of the FortiGate from another source (static route, RIP, OSPF, BGP).
If those conditions are not met, FortiOS silently drops the packet.
FIPS and Common Criteria
FortiOS has received NDPP, EAL2+, and EAL4+ based FIPS and Common Criteria certifications. Common Criteria
evaluations involve formal rigorous analysis and testing to examine security aspects of a product or system. Extensive
testing activities involve a comprehensive and formally repeatable process, confirming that the security product
functions as claimed by the manufacturer. Security weaknesses and potential vulnerabilities are specifically examined
during an evaluation.
To see Fortinet's complete history of FIPS/CC certifications go to the following URL and add Fortinet to the Vendor field:
https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search
PSIRT advisories
The FortiGuard Labs Product Security Incident Response Team (PSIRT) continually tests and gathers information about
Fortinet hardware and software products, looking for vulnerabilities and weaknesses. Any such findings are fed back to
Fortinet's development teams and serious issues are described along with protective solutions. The PSIRT regulatory
releases PSIRT advisories when issues are found and corrected. Advisories are listed at
https://www.fortiguard.com/psirt.
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

FortiOS ports and protocols
Communication to and from FortiOS is strictly controlled and only selected ports are opened for supported functionality
such as administrator logins and communication with other Fortinet products or services.
Accessing FortiOS using an open port is protected by authentication, identification, and encryption requirements. As
well, ports are only open if the feature using them is enabled.
FortiOS open ports
The following tables show the incoming and outgoing ports that are potentially opened by FortiOS.
Incoming ports
Purpose Protocol/Port
FortiAP-S Syslog, OFTP, Registration, Quarantine,
Log & Report
TCP/443
CAPWAP UDP/5246, UDP/5247
FortiAuthenticator Policy Authentication through Captive Portal TCP/1000
RADIUS disconnect TCP/1700
FortiClient Remote IPsec VPN access UDP/IKE 500, ESP (IP 50), NAT-T 4500
Remote SSL VPN access TCP/443
SSO Mobility Agent, FSSO TCP/8001
Compliance and Security Fabric TCP/8013 (by default; this port can be
customized)
FortiGate HA Heartbeat ETH Layer 0x8890, 0x8891, and 0x8893
HA Synchronization TCP/703, UDP/703
Unicast Heartbeat for Azure UDP/730
DNS for Azure UDP/53
FortiGuard Management TCP/541
AV/IPS UDP/9443
FortiManager AV/IPS Push UDP/9443
IPv4 FGFM management TCP/541
IPv6 FGFM management TCP/542
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

FortiOS ports and protocols 13
Incoming ports
Purpose Protocol/Port
FortiPortal API communications (FortiOS REST API,
used for Wireless Analytics)
TCP/443
3rd-Party Servers FSSO TCP/8001 (by default; this port can be
customized)
Others Web Admin TCP/80, TCP/443
Policy Override Authentication TCP/443, TCP/8008, TCP/8010
Policy Override Keepalive TCP/1000, TCP/1003
SSL VPN TCP/443
Outgoing ports
Purpose Protocol/Port
FortiAnalyzer Syslog, OFTP, Registration, Quarantine, Log
& Report
TCP/514
FortiAuthenticator LDAP, PKI Authentication TCP or UDP/389
RADIUS UDP/1812
FSSO TCP/8000
RADIUS Accounting UDP/1813
SCEP TCP/80, TCP/443
CRL Download TCP/80
External Captive Portal TCP/443
FortiGate HA Heartbeat ETH Layer 0x8890, 0x8891, and 0x8893
HA Synchronization TCP/703, UDP/703
Unicast Heartbeat for Azure UDP/730
DNS for Azure UDP/53
FortiGate Cloud Registration, Quarantine, Log & Report,
Syslog
TCP/443
OFTP TCP/514
Management TCP/541
Contract Validation TCP/443
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

FortiOS ports and protocols 14
Outgoing ports
Purpose Protocol/Port
FortiGuard AV/IPS Update TCP/443, TCP/8890
Cloud App DB TCP/9582
FortiGuard Queries UDP/53, UDP/8888, TCP/53, TCP/8888,
TCP/443 (as part of Anycast servers)
SDNS queries for DNS Filter UDP/53, TCP/853 (as part of Anycast
servers)
Registration TCP/80
Alert Email, Virus Sample TCP/25
Management, Firmware, SMS, FTM,
Licensing, Policy Override
TCP/443
Central Management, Analysis TCP/541
FortiManager IPv4 FGFM management TCP/541
IPv6 FGFM management TCP/542
Log & Report TCP or UDP/514
FortiGuard Queries UDP/53, UDP/8888, TCP/80, TCP/8888
FortiSandbox OFTP TCP/514
Others FSSO TCP/8001 (by default; this port can be
customized)
While a proxy is configured, FortiGate uses the following URLs to access the
FortiGuard Distribution Network (FDN):
l update.fortiguard.net
l service.fortiguard.net
l support.fortinet.com
Closing open ports
You can close open ports by disabling the feature that opens them. For example, if FortiOS is not managing a FortiAP
then the CAPWAP feature for managing FortiAPs can be disabled, closing the CAPWAP port.
The following sections of this document described a number of options for closing open ports:
l Use local-in policies to close open ports or restrict access on page 24
l Disable unused protocols on interfaces on page 23
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Security best practices
This section describes some techniques and best practices that you can use to improve FortiOS security.
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
l Review the Upgrade Path tool to determine the best path to take from your current version of FortiOS to the latest
version.
l Back up the current configuration.
Only FortiGate administrators who have read and write privileges can upgrade the FortiOS firmware.
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

