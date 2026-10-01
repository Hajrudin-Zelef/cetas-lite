---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-best-practices-555436-hardening-784ac110-1
title: "document-fortigate-7-6-0-best-practices-555436-hardening-784ac110"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-best-practices-555436-hardening-784ac110.md
source_anchor: ""
source_lines: [1, 80]
sha256: adf15bb2d557db7722f703941358787b78432162d983bf7e45a2af89bb013c06
---

# document-fortigate-7-6-0-best-practices-555436-hardening-784ac110

Hardening
Hardening
System hardening reduces security risk by eliminating potential attack vectors and shrinking the system's attack surface. Some of the best practices described previously in this document contribute to the hardening of the FortiGate with additional hardening steps listed here.
Physical security
Install the FortiGate in a physically secure location. Physical access to the FortiGate can allow it to be bypassed, or other firmware could be loaded after a manual reboot.
If the FortiGate cannot be physical secured:
- 
                                                    Ensure USB firmware and configuration installation are disabled. They are disabled by default: config system auto-install
    set auto-install-config disable
    set auto-install-image disable
end
- 
                                                    Enable port security (802.1x) to prevent unauthorized devices from forwarding traffic.
Vulnerability - monitoring PSIRT
Product Security Incident Response Team (PSIRT) continually tests and gathers information about Fortinet hardware and software products, looking for vulnerabilities and weaknesses. The findings are sent to the Fortinet development teams, and serious issues are described, along with protective solutions, in advisories listed at https://www.fortiguard.com/psirt.
Firmware
Keep the FortiOS firmware up to date. The latest patch release has the most fixed bugs and vulnerabilities, and should be the most stable. Firmware is periodically updated to add new features and resolve important issues.
- 
                                                    Read the release notes. The known issues may include issues that affect your business.
- 
                                                    Do not use out of support firmware. Review the Product Life Cycle > Software page and plan to upgrade before the FortiOS End of Support (EOS) date, which is when Fortinet Support services for the firmware version expire.
- 
                                                    Use a federated update to upgrade the firmware of all devices. This process follows the upgrade path to ensure a smooth transition. See Upgrading all device firmware by following the upgrade path (federated update) for more information.
- 
                                                    For standalone FortiGates, enable automatic firmware updates to automatically update firmware based on the FortiGuard upgrade path. Only upgrades to the latest patch of the current minor version are performed, for example from 7.6.1 to 7.6.2. See Enabling automatic firmware updates for more information.
- 
                                                    In the event a the user is unable to immediately apply a patch to their device, they have the option to temporarily activate virtual patching within their local-in policies. See Virtual patching on the local-in management interface for more information.
Encrypted protocols
Use encrypted protocols whenever possible, for example:
- 
                                                    LDAPS instead of LDAP
- 
                                                    RADSEC over TLS instead of RADIUS
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
|  | When configuring an LDAP connection to an Active Directory server, an administrator must provide Active Directory user credentials.  To secure RADIUS connections, consider using RADSEC over TLS instead. See Configuring a RADSEC client. | 
Strong ciphers
Force higher levels of encryption and strong ciphers. Strong crypto is enabled by default:
config system global
    set strong-crypto enable
    set ssl-static-key-ciphers disable
    set dh-params 8192
end
                                            See FortiGate encryption algorithm cipher suites for more information.
FortiGuard databases
Ensure that FortiGuard databases, such as AS, IPS, and AV, are updated punctually. Optionally, send an alert if they are out of date.
Penetration testing
Test your FortiGate to try to gain unauthorized access, or hire a penetration testing company to verify your work.
Denial of service
Denial of service (DoS) is a type of attack meant to disable a machine or network causing inaccessibility to the resource or users. Most often this is accomplished by overwhelming the target with more information than it can handle, resulting in a crash. DoS policies, which look for anomalous traffic patterns, are checked before the more resource intensive security policies to help prevent this.
The following guidelines can be used to get started with DoS policies. These policies can be applied to incoming traffic from your local network or internet, depending on your particular network.
- 
                                                    Enable anomaly logging and keep the action as monitor for some time. This is to observe and understand what expected traffic looks like so that you may tune thresholds to have small margins, and therefore more protection. Keep note of false alarms. If they are too frequent, you should adjust your policy accordingly.
- 
                                                    Enable the following DoS policy anomalies to help prevent targeted attacks: 
  - 
                                                            tcp_syn_flood
  - 
                                                            tcp_port_scan
  - 
                                                            tcp_src_session
  - 
                                                            tcp_dst_session
  - 
                                                            ip_src_session
  - 
                                                            ip_dst_session
 If you have an idea of your traffic rates for the preceding traffic patterns, you may adjust the threshold. Otherwise, begin with the default and adjust after a period of observing normal traffic. For more information, see DoS policy in the FortiOS Administration Guide.
- 
                                                            
