---
id: collect-261001-fortinet/fortinet/document-fortigate-7-0-0-best-practices-555436-930478d8-1
title: "document-fortigate-7-0-0-best-practices-555436-930478d8"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["asic", "incident"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-0-0-best-practices-555436-930478d8.md
source_anchor: ""
source_lines: [1, 93]
sha256: ba57be30bc2abf45301e62c6cd214d26f14de1dd9e034a98a44447c5eb8ee92d
---

# document-fortigate-7-0-0-best-practices-555436-930478d8

Hardening
Hardening
System hardening reduces security risk by eliminating potential attack vectors and shrinking the system's attack surface. Some of the best practices described previously in this document contribute to the hardening of the FortiGate with additional hardening steps listed here.
Physical security
Install the FortiGate in a physically secure location. Physical access to the FortiGate can allow it to be bypassed, or other firmware could be loaded after a manual reboot.
If the FortiGate cannot be physical secured:
- 
                                                    Disable USB firmware and configuration installation: config system auto-install
    set auto-install-config disable
    set auto-install-image disable
end
- 
                                                    Enable port security (802.1x) to prevent unauthorized devices from forwarding traffic.
- 
                                                    Optionally, disable the maintainer account. Note that doing this will make you unable to recover administrator access using a console connection is all of the administrator credentials are lost.
Vulnerability - monitoring PSIRT
Product Security Incident Response Team (PSIRT) continually tests and gathers information about Fortinet hardware and software products, looking for vulnerabilities and weaknesses. The findings are sent to the Fortinet development teams, and serious issues are described, along with protective solutions, in advisories listed at https://www.fortiguard.com/psirt.
Firmware
Keep the FortiOS firmware up to date. The latest patch release has the most fixed bugs and vulnerabilities, and should be the most stable. Firmware is periodically updated to add new features and resolve important issues.
- 
                                                    Read the release notes. The known issues may include issues that affect your business.
- 
                                                    Do not use out of support firmware. Review the product lifecycle and plan to upgrade before the firmware expires.
Encrypted protocols
Use encrypted protocols whenever possible, for example:
- 
                                                    LDAPS instead of LDAP
- 
                                                    SNMPv3 instead of SNMP
- 
                                                    SSH instead of telnet
- 
                                                    OSPF MD5 authentication
- 
                                                    SCP instead of FTP or TFTP
- 
                                                    NTP authentication
- 
                                                    Encrypted logging instead of TCP
|  | When configuring an LDAP connection to an Active Directory server, an administrator must provide Active Directory user credentials.  | 
Strong ciphers
Force higher levels of encryption and strong ciphers:
config system global
    set strong-crypto enable
    set ssl-static-key-ciphers disable
    set dh-params 8192
end
                                            See FortiGate encryption algorithm cipher suites for more information.
|  | The ssh-hmac-md5 andssh-cbc-cipher commands are removed in FortiOS 7.0.2 and later. See Enabling individual ciphers in the SSH administrative access protocol for details. | 
FortiGuard databases
Ensure that FortiGuard databases, such as AS, IPS, and AV, are updated punctually. Optionally, send an alert if they are out of date.
Penetration testing
Test your FortiGate to try to gain unauthorized access, or hire a penetration testing company to verify your work.
Denial of service
Denial of service (DoS) is a type of attack meant to disable a machine or network causing inaccessibility to the resource or users. Most often this is accomplished by overwhelming the target with more information than it can handle, resulting in a crash. DoS policies, which look for anomalous traffic patterns, are checked before the more resource intensive security policies to help prevent this.
The following guidelines can be used to get started with DoS policies. These policies can be applied to incoming traffic from your local network or internet, depending on your particular network.
- Enable anomaly logging and keep the action as monitor for some time. This is to observe and understand what expected traffic looks like so that you may tune thresholds to have small margins, and therefore more protection. Keep note of false alarms. If they are too frequent, you should adjust your policy accordingly.
- Enable the following DoS policy anomalies to help prevent targeted attacks: 
  - tcp_syn_flood
  - tcp_port_scan
  - tcp_src_session
  - tcp_dst_session
  - ip_src_session
  - ip_dst_session
 If you have an idea of your traffic rates for the preceding traffic patterns, you may adjust the threshold. Otherwise, begin with the default and adjust after a period of observing normal traffic. For more information, see DoS protection in the FortiOS Administration Guide.
- Where possible, enable ASIC DoS for offloading using network processor ASICs. The FortiOS Hardware Acceleration Guide contains more information about DoS-related NP6 ASIC features, such as configuring NP6 anomaly protection and using the host protection engine (HPE) to protect the FortiGate from DoS attacks.
Secure password storage
The passwords, and private keys used in certificates, that are stored on the FortiGate are encrypted using a predefined private key, and encoded when displayed in the CLI and configuration file. System admin passwords are hashed with SHA256 and encoded before being displayed.
Passwords cannot be decrypted without the private key and are not shown anywhere in clear text. The private key is required on other FortiGates to restore the system from a configuration file. In an HA cluster, the same key should be used on all of the units.
To enhance password security, specify a custom private key for the encryption process. This ensures that the key is only known by you.
FortiGate models with a Trusted Platform Module (TPM) can store the master encryption password, which is used to generate the master encryption key, on the TPM. For more information, see Trusted platform module support.
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
The FortiGate configuration file has important information that should always be kept secured, including details about your network, users, credentials, passwords, and keys. There are many reasons to back up your configuration, such as disaster recovery, preparing for migrating to another device, and troubleshooting. Evaluate the risk involved if your configurations were exposed, and manage your risk accordingly.
When backing up your configuration, consider the following steps to safeguard the file:
- 
                                                    Enable Encryption when backing up the configuration.
- 
                                                    Store the configuration file in a secure location.
- 
                                                    Delete old configuration files that are no longer needed.
If a configuration file must be shared with a third party for auditing, troubleshooting, or any other reasons, consider only providing a section of the file and not the entire file. Otherwise, consider the following steps:
- 
                                                    Enable Encryption when backing up the configuration and only share the password with the intended party.
- 
