---
id: collect-261001-fortinet/fortinet/questions-2173-how-to-get-a-list-of-ports-listening-in-a-fortigate-firewall-1a5c8c4c-1
title: "questions-2173-how-to-get-a-list-of-ports-listening-in-a-fortigate-firewall-1a5c8c4c"
domain: fortinet
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2014-05-09"]
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-fortinet/questions-2173-how-to-get-a-list-of-ports-listening-in-a-fortigate-firewall-1a5c8c4c.md
source_anchor: ""
source_lines: [1, 189]
sha256: 492121d99355f1ce9bf84e89a4ace9ae9cf38612ac5e75a34e4dc6b8c673d3d4
---

# questions-2173-how-to-get-a-list-of-ports-listening-in-a-fortigate-firewall-1a5c8c4c

Ports used by Fortinet was released May 9, 2014
Destination • Port Protocol(s) • Application(s) • Function(s)
21 TCP FTP • Log and Report uploads from FortiAnalyzer
• Anti-defacement backup and restoration (FTP). Listening on
FortiWeb
• FTP configuration backup from FortiWeb to other device
22 TCP SSH • SSH Command line based management:
• From Admin Workstation to Fortinet Device
22 TCP FTP over SSH • Log and Report uploads:
• To and from FortiCloud
• To and from FortiAnalyzer
• Anti-defacement backup and restoration (SSH/SCP) from FortiWeb
to other device
• SFTP configuration backup from FortiWeb to other device
23 TCP Telnet • Telnet Command line based management from Admin Workstation
to Fortinet devices
• HA (FGCP) between HA FortiGates
25 TCP SMTP • Alert Emails
• From FortiAnalyzer to SMTP Mail Server
• From FortiGate to SMTP Mail Server
• From FortiWeb to SMTP Mail Server
• Encrypted Virus Samples auto submitted to FortiGuard
49 TCP TACACS+ • TACACS+ from FortiAnalyzer
53 UDP DNS • DNS Lookups
• To DNS Servers
• To FortiGuard
4
4
53 UDP Fortinet Queries • FortiGuard Server List requests to FortiGuard
• AntiSpam or Web Filtering rating lookup queries to FortiGuard
• URL/AS rating lookup queries to FortiGuard
• Real-time Black List(RBL) lookup requests to RBL services
67 UDP DHCP • DHCP to and from FortiGate
68 UDP DHCP Relay • DHCP Relay to and from FortiGate
69 UDP TFTP • TFTP for backups, restoration, and firmware updates from FortiWeb
to other device
80 TCP • Default unsecure Web-based Management of Fortinet Device
• Admin Workstation to FortiAnalyzer
• Admin Workstation to FortiAuthenticator
• Admin Workstation to FortiGate
• Admin Workstation to FortiManager
• Admin Workstation to FortiWeb
80 TCP HTTP • Proxied HTTP traffic from FortiGate
80 TCP HTTP • Fortinet Device Registration to FortiGuard
• AV update requests from FortiClient to FortiManager
• Server health checks from FortiWeb to other device
• Predefined HTTP service. Only occurs if the service is used by a
policy, listening on FortiWeb
80 TCP Simple Certificate Enrollment
Protocol (SCEP)
• Issuing and revocation of digital certificates
• Listening on FortiAuthenticator
88 TCP Kerboros • Account Authentication traffic from FortiAuthenticator to Active
Directory Controllers
123 UDP NTP • Time Synchronization from Fortinet Device to NTP Server
135 TCP Client/Server (WMI, SEL) • FortiAuthenticator to Active Directory Controllers
137 UDP • Win Share to and from FortiAnalyzer (Not supported in FAZ v5.0/5.2)
• Anti-defacement backup and restoration (Windows-style share) from
FortiWeb to other device.
138 UDP • Win Share to and from FortiAnalyzer (Not supported in FAZ v5.0/5.2)
• Anti-defacement backup and restoration (Windows-style share) from
FortiWeb to other device.
5
5
139 TCP/UDP NetBIOS • Win Share to and from FortiAnalyzer (Not supported in FAZ v5.0/5.2)
• Anti-defacement backup and restoration (Windows-style share) from
FortiWeb to other device.
161 UDP Simple Network Management
Protocol (SNMP)
• SNMP Poll
• FortiManager to FortiGate
• Listening on FortiAuthenticator
• Listening on FortiWeb
162 UDP Simple Network Management
Protocol (SNMP) Traps
• To SysLog server
• To FortiAnalyzer
• To FortiManager
389 TCP/UDP LDAP • LDAP Lookups, Authentication Requests and Report queries
• PKI Authentication
• To Active Directory Domain Controllers
• To FortiAuthenticator
• To LDAP Server
443 TCP HTTPS • Default Secure Web-based Management of Fortinet Device
• Admin Workstation to Fortinet Device
• Firmware and Signature Downloads from FortiGuard
• FGD SMS to FortiGuard
• FC FTM to FortiGuard
• FC Licensing to FortiGuard
• Policy Override Auth to FortiGuard
• AntiVirus/IPS updates to FortiGuard
• URL/AS update requests to FortiGuard
• Remote Vulnerability Scan updates to FortiGuard
• Device Registration requests to FortiGuard
• Server health checks from FortiWeb to other devices
• Proxied HTTPS traffic from FortiGate to Proxy Server
• FSSO Portal and Widget traffic
6
6
443 TCP Representational state transfer
(REST) API / HTTP
• Listening on FortiAnalyzer
445 TCP Microsoft-DS Active Directory,
Windows shares
• Domain Controller Polling
• FortiAuthenticator to Active Directory Domain Controller
• Listening on FortiAnalyzer
• NTLM authentication queries.
• Anti-defacement backup and restoration (Windows-style share)
from FortiWeb to other device.
500 UDP IPsec • Secure SNMP over IPsec connection
• FortiGate to FortiAnalyzer
514 TCP/UDP Syslog messages OFTP • Device Registration
• From FortiManager to FortiAnalyzer
• From FortiGate to FortiAnalyzer
• Quarantined files to FortiAnalyzer
• Logs and Reports
• To SysLog server
• To FortiAnalyzer
• To FortiCloud
• To FortiManager
• OFTP for file submission and statistics exchange
• Between FortiGate and FortiSandbox (FortiCloud)
520 UDP Routing Information Protocol (RIP) • Listening on FortiGate
541 TCP • Device Registration • Central Management from FortiManager
• SSL Management Tunnel to FortiCloud
636 TCP Lightweight Directory Access
Protocol over TLS/SSL (LDAPS)
• Encrypted LDAP authentication traffic from
• Fortinet Devices to Active Directory Domain Controllers
• Fortinet Devices to LDAP servers (including FortiAuthenticator)
703 TCP FGCP L2 • HA Heartbeat between HA FortiGates
1000 TCP • Policy Override Keepalive listening on FortiGate
(Closed by default, but can be enabled)
7
7
1003 TCP • Policy Override Keepalive listening on FortiGate
(Closed by default, but can be enabled)
1812 TCP RADIUS • RADIUS Authentication Requests
• To FortiAuthenticator
• To RADIUS Server
1813 UDP RADIUS • RADIUS Accounting to FortiAuthenticator
2049 TCP NFS • Network File System listening on FortiAnalyzer (Not supported in
FAZ v5.0/5.2)
2302 TCP • HTTP or HTTPS administrative access to web-based manager's CLI
dashboard widget(v3.0 MR5 only)
• Listening on FortiAnalyzer
• Listening on FortiGate
2560 TCP Online Certificate Status Protocol
(OCSP)
• Obtaining the revocation status of an X.509 digital certificate,
listening on FortiAuthenticator
3000 TCP • Log aggregation listening on FortiAnalyzer
(Log aggregation server support requires model FortiAnalyzer
800 or greater)
3306 TCP • Remote MySQL database connection listening on FortiAnalyzer
3784 UDP BFD • Listening on FortiGate
4500 UDP IPsec • Secure SNMP over IPsec connection
• FortiGate to FortiAnalyzer
• FortiGate to FortiManager
5199 TCP • HA Heartbeat or synchronization listening on FortiManager
6055 UDP • HA heartbeat. Layer 2 multicast.
• From FortiWeb to other device
• Listening on FortiWeb
6056 UDP • HA configuration synchronization. Layer 2 multicast.
• From FortiWeb to other device
• Listening on FortiWeb
8
8
8000 TCP FSSO • Windows Active Directory Collector Agent for Fortinet Single Sign-On
• From Active Directory Collector to FortiGate
• From FortiAuthenticator to FortiGate
• From FortiGate to FortAuthenticator
8001 TCP SSO Mobiltity Agent • This port is used to pass userid and IP address information from
FortiClient to FortiAuthenticator.
(This functionality is not necessary for the completion of phase 1)
8002 TCP/UDP FSSO • UDP (for plain traffic), or TCP (for encrypted traffic)
• FortiAuthenticator listening for traffic - Hierarchical FSSO Info from
Tier Supplier
8003 TCP FSSO • FortiAuthenticator listening for traffic from DS/TS Agents with FSSO
Login information
8008 TCP • User authentication for policy override of HTTP traffic listening on
FortiGate
8009 TCP • FortiClient Portal listening on FortiGate 1000A, 3600A, and 5005FA2
only
8010 TCP • User authentication for policy override of HTTPS traffic from
FortiClient to FortiGate
(This port and IP address must be load balanced between all four
FortiGate 1500Ds)
