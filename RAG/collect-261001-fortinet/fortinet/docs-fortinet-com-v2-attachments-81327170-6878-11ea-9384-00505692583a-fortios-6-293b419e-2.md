---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6-293b419e-2
title: "docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e.md
source_anchor: ""
source_lines: [137, 257]
sha256: 74034e1858b28fcea17e74c11b36ed198a564499617118a3c830daa89d92e528
---

# docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e

Building security into FortiOS 8
Secure password storage
The passwords, and private keys used in certificates, that are stored on the FortiGate are encrypted using a predefined
private key, and encoded when displayed in the CLI and configuration file. System admin passwords are hashed with
SHA256 and encoded before being displayed.
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
Configuration backup
The FortiGate configuration file has important information that should always be kept secured, including details about
your network, users, credentials, passwords, and keys. There are many reasons to back up your configuration, such as
disaster recovery, preparing for migrating to another device, and troubleshooting. Evaluate the risk involved if your
configurations were exposed, and manage your risk accordingly.
When backing up your configuration, consider the following steps to safeguard the file:
l Enable Encryption when backing up the configuration.
l Store the configuration file in a secure location.
l Delete old configuration files that are no longer needed.
If a configuration file must be shared with a third party for auditing, troubleshooting, or any other reasons, consider only
providing a section of the file and not the entire file. Otherwise, consider the following steps:
l Enable Encryption when backing up the configuration and only share the password with the intended party.
l Manually replace the passwords in the backed up configuration file.
l Request that the configuration file be deleted after the intended purpose has been satisfied.
Maintainer account
Administrators with physical access to a FortiGate appliance can use a console cable and a special administrator
account called maintainer to log into the CLI. When enabled, the maintainer account can be used to log in from the
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS 9
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
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Building security into FortiOS 10
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
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

