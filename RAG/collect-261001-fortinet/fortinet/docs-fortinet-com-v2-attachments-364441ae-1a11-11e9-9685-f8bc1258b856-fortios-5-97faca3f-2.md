---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5-97faca3f-2
title: "docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["distribution", "incident"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f.md
source_anchor: ""
source_lines: [177, 361]
sha256: 35133c6fe2aea432a0e60c91c9be58b6ec250eccc5adc2a5f9d6693ac2956b54
---

# docs-fortinet-com-v2-attachments-364441ae-1a11-11e9-9685-f8bc1258b856-fortios-5--97faca3f

Building security into FortiOS Network security
l SCP is disabled by default. Enabling SCP allows downloading the configuration file from the FortiGate as an
alternative method of backing up the configuration file. To enable SCP:
config system global
set admin-scp enable
end
l DHCP is enabled by default on the dedicated MGMT interface and on the predefined LAN port (defined on some
FortiGate models).
l On FortiGate models with dedicated MGMT interfaces, dedicated DMZ interfaces, dedicated WAN interfaces, and
pre-defined LAN interfaces, the default management access on interfaces is shown below. Outside of the interfaces
listed below, management access must be explicitly enabled on interfaces – management services are enabled on
specific interfaces and not globally.
l Dedicated Management Interface
l Ping
l FMG-Access (fgfm)
l CAPWAP
l HTTPS
l HTTP
l Dedicated WAN1/WAN2 Ports
l Ping
l FMG-Access (fgfm)
l Dedicated DMZ Port
l Ping
l FMG-Access (fgfm)
l CAPWAP
l HTTPS
l HTTP
l Pre-Defined LAN Port
l Ping
l FMG-Access (fgfm)
l CAPWAP
l HTTPS
l HTTP
Network security
This section describes FortiOS and FortiGate network security features.
Network interfaces
The following are disabled by default on each FortiGate interface:
l Broadcast forwarding
l STP forwarding
l VLAN forwarding
l L2 forwarding
Hardening
Fortinet Technologies Inc.
9

FIPS and Common Criteria Building security into FortiOS
l Netbios forwarding
l Ident accept
For more information, see Disable unused protocols on interfaces on page 23.
TCP sequence checking
FortiOS uses TCP sequence checking to ensure a packet is part of a TCP session. By default, anti-replay
protection is strict, which means that if a packet is received with sequence numbers that fall out of the expected
range, FortiOS drops the packet. Strict anti-replay checking performs packet sequence checking and ICMP anti-
replay checking with the following criteria:
l The SYN, FIN, and RST bit cannot appear in the same packet.
l FortiOS does not allow more than 1 ICMP error packet to go through before it receives a normal TCP or UDP packet.
l If FortiOS receives an RST packet, FortiOS checks to determine if its sequence number in the RST is within the un-
ACKed data and drops the packet if the sequence number is incorrect.
l For each new session, FortiOS checks to determine if the TCP sequence number in a SYN packet has been
calculated correctly and started from the correct value.
Reverse path forwarding
FortiOS implements a mechanism called Reverse Path Forwarding (RPF), or Anti Spoofing, to block an IP packet
from being forwarded if its source IP does not:
l belong to a locally attached subnet (local interface), or
l be in the routing domain of the FortiGate from another source (static route, RIP, OSPF, BGP).
If those conditions are not met, FortiOS silently drops the packet.
FIPS and Common Criteria
FortiOS has received NDPP, EAL2+, and EAL4+ based FIPS and Common Criteria certifications. Common
Criteria evaluations involve formal rigorous analysis and testing to examine security aspects of a product or
system. Extensive testing activities involve a comprehensive and formally repeatable process, confirming that the
security product functions as claimed by the manufacturer. Security weaknesses and potential vulnerabilities are
specifically examined during an evaluation.
To see Fortinet's complete history of FIPS/CC certifications go to the following URL and add Fortinet to the Vendor
field:
https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search
PSIRT advisories
The FortiGuard Labs Product Security Incident Response Team (PSIRT) continually tests and gathers information
about Fortinet hardware and software products, looking for vulnerabilities and weaknesses. Any such findings are
fed back to Fortinet's development teams and serious issues are described along with protective solutions. The
10 Hardening
Fortinet Technologies Inc.

Building security into FortiOS PSIRT advisories
PSIRT regulatory releases PSIRT advisories when issues are found and corrected. Advisories are listed at
https://www.fortiguard.com/psirt.
Hardening
Fortinet Technologies Inc.
11

FortiOS open ports FortiOS ports and protocols
FortiOS ports and protocols
Communication to and from FortiOS is strictly controlled and only selected ports are opened for supported
functionality such as administrator logins and communication with other Fortinet products or services.
Accessing FortiOS using an open port is protected by authentication, identification, and encryption requirements.
As well, ports are only open if the feature using them is enabled.
FortiOS open ports
The following diagram and tables shows the incoming and outgoing ports that are potentially opened by FortiOS.
For more details about open ports and the communication protocols that FortiOS uses, see the document Fortinet
Communication Ports and Protocols.
12 Hardening
Fortinet Technologies Inc.

FortiOS ports and protocols FortiOS open ports
Incoming ports
Purpose Protocol/Port
FortiAP-S Syslog, OFTP, Registration,
Quarantine, Log & Report
TCP/443
CAPWAP UDP/5246, UDP/5247
FortiAuthenticator RADIUS UDP/1812
FSSO TCP/8000
FortiGate HA Heartbeat ETH Layer 0x8890, 0x8891, and 0x8893
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
Policy Override Authentication TCP/443, TCP/8008, TCP/8010
FortiClient Portal TCP/8009
Policy Override Keepalive TCP/1000, TCP/1003
SSL VPN TCP/10443
3rd-Party Servers FSSO TCP/8000
Hardening
Fortinet Technologies Inc.
13

FortiOS open ports FortiOS ports and protocols
Outgoing ports
Purpose Protocol/Port
FortiAnalyzer Syslog, OFTP, Registration,
Quarantine, Log & Report
TCP/514
FortiAuthenticator LDAP, PKI Authentication TCP or UDP/389
FortiCloud Registration, Quarantine, Log & Report,
Syslog
TCP/443
OFTP TCP/514
Management TCP/541
Contract Validation TCP/443
FortiGate HA Heartbeat ETH Layer 0x8890, 0x8891, and 0x8893
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
14 Hardening
Fortinet Technologies Inc.

FortiOS ports and protocols Closing open ports
Note that, while a proxy is configured, FortiGate uses the following URLs to access
the FortiGuard Distribution Network (FDN):
l update.fortiguard.net
l service.fortiguard.net
l support.fortinet.com
Closing open ports
You can close open ports by disabling the feature that opens them. For example, if FortiOS is not managing a
FortiAP then the CAPWAP feature for managing FortiAPs can be disabled, closing the CAPWAP port.
The following sections of this documnent described a number of options for closing open ports:
l Use local-in policies to close open ports or restrict access on page 24
l Disable unused protocols on interfaces on page 23
Hardening
Fortinet Technologies Inc.
15

