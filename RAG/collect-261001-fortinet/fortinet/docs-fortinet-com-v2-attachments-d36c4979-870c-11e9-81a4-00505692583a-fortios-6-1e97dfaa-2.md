---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6-1e97dfaa-2
title: "docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa.md
source_anchor: ""
source_lines: [135, 256]
sha256: 9409b9123007f9ebe8e920621e68bd7fcb963b3e1789be730fdf72b0354fbf15
---

# docs-fortinet-com-v2-attachments-d36c4979-870c-11e9-81a4-00505692583a-fortios-6--1e97dfaa

Building security into FortiOS 8
Secure password storage
The passwords, and private keys used in certificates, that are stored on the FortiGate are encrypted using a predefined
private key, and encoded when displayed in the CLI and configuration file.
Passwords cannot be decrypted without the private key and are not shown anywhere in clear text. The private key is
required on other FortiGates to restore the system from a configuration file. In an HA cluster, the same key should be
used on all of the units.
To enhance password security, specify a custom private key for the encryption process. This ensures that the key is only
known by you.
FortiGate models with a Trusted Platform Module (TPM) can store the master encryption password, which is used to
generate the master encryption key, on the TPM. For more information, see Trusted platform module support.
To configure your own private encryption key:
config system global
set private-data-encryption enable
end
Please type your private data encryption key (32 hexadecimal numbers):
********************************
Please re-enter your private data encryption key (32 hexadecimal numbers) again:
********************************
Your private data encryption key is accepted.
Maintainer account
Administrators with physical access to a FortiGate appliance can use a console cable and a special administrator
account called maintainer to log into the CLI. When enabled, the maintainer account can be used to log in from the
console after a hard reboot. The password for the maintainer account is bcpb followed by the FortiGate serial number.
An administrator has 60-seconds to complete this login. See the Fortinet knowledge base or Resetting a lost Admin
password for details.
The only action the maintainer account has permissions to perform is to reset the passwords of super_admin accounts.
Logging in with the maintainer account requires a hard boot of the FortiGate. FortiOS generates event log messages
when you log in with the maintainer account and for each password reset.
The maintainer account is enabled by default; however, there is an option to disable this feature. The maintainer account
can be disabled using the following command:
config system global
set admin-maintainer disable
end
If you disable this feature and lose your administrator passwords you will no longer be able to
log into your FortiGate.
Administrative access security
Secure administrative access features:
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS 9
l SSH, Telnet, and SNMP are disabled by default. If required, these admin services must be explicitly enabled on
each interface from the GUI or CLI.
l SSHv1 is disabled by default. SSHv2 is the default version.
l SSLv3 and TLS1.0 are disabled by default. TLSv1.1 and TLSv1.2 are the SSL versions enabled by default for
HTTPS admin access.
l HTTP is disabled by default, except on dedicated MGMT, DMZ, and predefined LAN interfaces. HTTP redirect to
HTTPS is enabled by default.
l The strong-crypto global setting is enabled by default and configures FortiOS to use strong ciphers (AES,
3DES) and digest (SHA1) for HTTPS/SSH/TLS/SSL functions.
l SCP is disabled by default. Enabling SCP allows downloading the configuration file from the FortiGate as an
alternative method of backing up the configuration file. To enable SCP:
config system global
set admin-scp enable
end
l DHCP is enabled by default on the dedicated MGMT interface and on the predefined LAN port (defined on some
FortiGate models).
l The default management access configuration for FortiGate models with dedicated MGMT, DMZ, WAN, and LAN
interfaces is shown below. Outside of the interfaces listed below, management access must be explicitly enabled on
interfaces – management services are enabled on specific interfaces and not globally.
l Dedicated management interface
l Ping
l FMG-Access (fgfm)
l CAPWAP
l HTTPS
l HTTP
l Dedicated WAN1/WAN2 interface
l Ping
l FMG-Access (fgfm)
l Dedicated DMZ interface
l Ping
l FMG-Access (fgfm)
l CAPWAP
l HTTPS
l HTTP
l Dedicated LAN interface
l Ping
l FMG-Access (fgfm)
l CAPWAP
l HTTPS
l HTTP
Non-factory SSL certificates
Non-factory SSL certificates should be used for the administrator and SSL VPN portals. Your certificate should identify
your domain so that remote users can recognize the identity of the server or portal that they are accessing through a
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS 10
trusted CA.
The default Fortinet factory self-signed certificates are provided to simplify initial installation and testing. Using these
certificates leaves you vulnerable to man-in-the-middle attacks, where an attacker spoofs your certificate, compromises
your connection, and steals your personal information.
It is highly recommended that you purchase a server certificate from a trusted CA to allow remote users to connect to
SSL VPN with confidence. Your administrator web portal should also be configured with a server certificate from a
trusted CA. See Purchase and import a signed SSL certificate for information.
Network security
This section describes FortiOS and FortiGate network security features.
Network interfaces
The following are disabled by default on each FortiGate interface:
l Broadcast forwarding
l STP forwarding
l VLAN forwarding
l L2 forwarding
l Netbios forwarding
l Ident accept
For more information, see Disable unused protocols on interfaces on page 23.
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
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

