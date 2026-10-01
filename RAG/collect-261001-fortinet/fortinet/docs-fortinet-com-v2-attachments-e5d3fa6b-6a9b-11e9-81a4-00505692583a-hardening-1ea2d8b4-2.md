---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening-1ea2d8b4-2
title: "docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4.md
source_anchor: ""
source_lines: [135, 275]
sha256: b21d88091504a827dc7d5e51e450dce1adabd6185c39b16e724dd96061c6423b
---

# docs-fortinet-com-v2-attachments-e5d3fa6b-6a9b-11e9-81a4-00505692583a-hardening--1ea2d8b4

Building security into FortiOS 8
hash value. For example:
config system admin
edit "admin"
set accprofile "super_admin"
set vdom "root"
set password ENC SH2nlSm9QL9tapcHPXIqAXvX7vBJuuqu22hpa0JX0sBuKIo7z2g0Kz/+0KyH4E=
next
end
Pre-shared keys in IPSec phase-1 configurations are stored in plain text. In the configuration file these pre-shared keys
are encoded. The encoding consists of encrypting the password with a fixed key using DES (AES in FIPS mode) and
then Base64 encoding the result.
Maintainer account
Administrators with physical access to a FortiGate appliance can use a console cable and a special administrator
account called maintainer to log into the CLI. When enabled, the maintainer account can be used to log in from the
console after a hard reboot. The password for the maintainer account is bcpb followed by the FortiGate serial number.
An administrator has 60-seconds to complete this login. See the Fortinet knowledgebase or Resetting a lost Admin
password for details.
The only action the maintainer account has permissions to perform is to reset the passwords of super_admin accounts.
Logging in with the maintainer account requires rebooting the FortiGate. FortiOS generates event log messages when
you login with the maintainer account and for each password reset.
The maintainer account is enabled by default; however, there is an option to disable this feature. The maintainer
account can be disabled using the following command:
config system global
set admin-maintainer disable
end
If you disable this feature and lose your administrator passwords you will no longer be able to
log into your FortiGate.
Administrative access security
Secure administrative access features:
l SSH, Telnet, and SNMP are disabled by default. If required, these admin services must be explicitly enabled on
each interface from the GUI or CLI.
l SSHv1 is disabled by default. SSHv2 is the default version.
l SSLv3 and TLS1.0 are disabled by default. TLSv1.1 and TLSv1.2 are the SSL versions enabled by default for
HTTPS admin access.
l HTTP is disabled by default, except on dedicated MGMT, DMZ, and predefined LAN interfaces. HTTP redirect to
HTTPS is enabled by default.
l Thestrong-crypto global setting is enabled by default and configures FortiOS to use strong ciphers (AES,
3DES) and digest (SHA1) for HTTPS/SSH/TLS/SSL functions.
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS 9
l SCP is disabled by default. Enabling SCP allows downloading the configuration file from the FortiGate as an
alternative method of backing up the configuration file. To enable SCP:
config system global
set admin-scp enable
end
l DHCP is enabled by default on the dedicated MGMT interface and on the predefined LAN port (defined on some
FortiGate models).
l The default management access configuration for FortiGate models with dedicated MGMT, DMZ, WAN, and LAN
interfaces is shown below. Outside of the interfaces listed below, management access must be explicitly enabled
on interfaces – management services are enabled on specific interfaces and not globally.
l Dedicated management interface
l Ping
l FMG-Access(fgfm)
l CAPWAP
l HTTPS
l HTTP
l Dedicated WAN1/WAN2 interface
l Ping
l FMG-Access(fgfm)
l Dedicated DMZ interface
l Ping
l FMG-Access(fgfm)
l CAPWAP
l HTTPS
l HTTP
l Dedicated LAN interface
l Ping
l FMG-Access(fgfm)
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
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS 10
l Netbios forwarding
l Ident accept
For more information, see Disable unused protocols on interfaces on page 24.
TCP sequence checking
FortiOS uses TCP sequence checking to ensure a packet is part of a TCP session. By default, anti-replay protection is
strict, which means that if a packet is received with sequence numbers that fall out of the expected range, FortiOS drops
the packet. Strict anti-replay checking performs packet sequence checking and ICMP anti-replay checking with the
following criteria:
l The SYN, FIN, and RST bit cannot appear in the same packet.
l FortiOS does not allow more than 1 ICMP error packet to go through before it receives a normal TCP or UDP
packet.
l If FortiOS receives an RST packet, FortiOS checks to determine if its sequence number in the RST is within the un-
ACKed data and drops the packet if the sequence number is incorrect.
l For each new session, FortiOS checks to determine if the TCP sequence number in a SYN packet has been
calculated correctly and started from the correct value.
Reverse path forwarding
FortiOS implements a mechanism called Reverse Path Forwarding (RPF), or Anti Spoofing, to block an IP packet from
being forwarded if its source IP does not:
l belong to a locally attached subnet (local interface), or
l be in the routing domain of the FortiGate from another source (static route, RIP, OSPF, BGP).
If those conditions are not met, FortiOS silently drops the packet.
FIPS and Common Criteria
FortiOS has received NDPP, EAL2+, and EAL4+ based FIPS and Common Criteria certifications. Common Criteria
evaluations involve formal rigorous analysis and testing to examine security aspects of a product or system. Extensive
testing activities involve a comprehensive and formally repeatable process, confirming that the security product
functions as claimed by the manufacturer. Security weaknesses and potential vulnerabilities are specifically examined
during an evaluation.
To see Fortinet's complete history of FIPS/CC certifications go to the following URL and add Fortinet to the Vendor
field:
https://csrc.nist.gov/projects/cryptographic-module-validation-program/validated-modules/search
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS 11
PSIRT advisories
The FortiGuard Labs Product Security Incident Response Team (PSIRT) continually tests and gathers information
about Fortinet hardware and software products, looking for vulnerabilities and weaknesses. Any such findings are fed
back to Fortinet's development teams and serious issues are described along with protective solutions. The PSIRT
regulatory releases PSIRT advisorieswhen issues are found and corrected. Advisories are listed at
https://www.fortiguard.com/psirt.
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

FortiOS ports and protocols 12
FortiOS ports and protocols
Communication to and from FortiOS is strictly controlled and only selected ports are opened for supported functionality
such as administrator logins and communication with other Fortinet products or services.
Accessing FortiOS using an open port is protected by authentication, identification, and encryption requirements. As
well, ports are only open if the feature using them is enabled.
FortiOS open ports
The following diagram and tables shows the incoming and outgoing ports that are potentially opened by FortiOS.
Incoming ports
Purpose Protocol/Port
FortiAP-S Syslog, OFTP, Registration,
Quarantine, Log & Report
TCP/443
CAPWAP UDP/5246, UDP/5247
FortiOS Handbook - Hardening your FortiGate Fortinet Technologies Inc.

