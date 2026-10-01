---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-17
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "agent"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [2855, 3063]
sha256: ee643f368bcea69fba4a67c5dc2368ffb5b4c71d17d6d9669f62858ffb5a8027
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 115 
8.3.1 Centralized Logging and Reporting (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Device logs should be sent to a centralized device for log collection, retention, and 
reporting. This could be a SIEM. syslog device, FortiAnalyzer, FortiManager, etc. 
Rationale: 
Centralized logging allows for more reliable log retention and more enriched log data for 
review and reporting. 
Audit: 
Review log settings through the administrative web page go to Log & Report > 
Log Settings and validate under "Remote Logging and Archiving" that logs are 
being offloaded to another device. 
Remediation: 
Configure a remote server for logs to be sent to. 
Access the FortiGate administrative web access page and to to Log & Report > 
Log Settings and under "Remote Logging and Archiving" configure a remote 
server to send logs to.

Page 116 
Appendix: Summary Table 
CIS Benchmark Recommendation Set 
Correctly 
Yes No 
1 Network Settings 
1.1 Ensure DNS server is configured (Automated)   
1.2 Ensure intra-zone traffic is not always allowed (Manual)   
1.3 Disable all management related services on WAN port 
(Manual) 
  
2 System Settings 
2.1 General Settings 
2.1.1 Ensure 'Pre-Login Banner' is set (Automated)   
2.1.2 Ensure 'Post-Login-Banner' is set (Automated)   
2.1.3 Ensure timezone is properly configured (Manual)   
2.1.4 Ensure correct system time is configured through NTP 
(Automated) 
  
2.1.5 Ensure hostname is set (Automated)   
2.1.6 Ensure the latest firmware is installed (Manual)   
2.1.7 Disable USB Firmware and configuration installation 
(Automated) 
  
2.1.8 Disable static keys for TLS (Automated)   
2.1.9 Enable Global Strong Encryption (Automated)   
2.2 Password Policy 
2.2.1 Ensure 'Password Policy' is enabled (Automated)   
2.2.2 Ensure administrator password retries and lockout time 
are configured (Automated) 
 

Page 117 
CIS Benchmark Recommendation Set 
Correctly 
Yes No 
2.3 SNMP 
2.3.1 Ensure SNMP agent is disabled (Automated)   
2.3.2 Ensure only SNMPv3 is enabled (Automated)   
2.4 Administrators and Admin Profiles 
2.4.1 Ensure default 'admin' password is changed (Manual)   
2.4.2 Ensure all the login accounts having specific trusted 
hosts enabled (Manual) 
  
2.4.3 Ensure admin accounts with different privileges having 
their correct profiles assigned (Manual) 
  
2.4.4 Ensure idle timeout time is configured (Automated)   
2.4.5 Ensure only encrypted access channels are enabled 
(Automated) 
  
2.4.6 Apply Local-in Policies (Manual)   
2.5 High Availability 
2.5.1 Ensure High Availability Configuration (Automated)   
2.5.2 Ensure "Monitor Interfaces" for High Availability Devices 
is Enabled (Automated) 
  
2.5.3 Ensure HA Reserved Management Interface is 
Configured (Manual) 
  
3 Policy and Objects 
3.1 Ensure that unused policies are reviewed regularly 
(Manual) 
  
3.2 Ensure that policies do not use  "ALL" as Service 
(Automated) 
  
3.3 Ensure Policies are Uniquely Named (Manual)  

Page 118 
CIS Benchmark Recommendation Set 
Correctly 
Yes No 
3.4 Ensure there are no Unused Policies (Manual)   
3.5 Ensure firewall policy denying all traffic to/from Tor or 
malicious server IP addresses using ISDB (Manual) 
  
3.6 Ensure logging is enabled on all firewall policies 
(Manual) 
  
4 Security Profiles 
4.1 Intrusion Prevention System (IPS) 
4.1.1 Detect Botnet Connections (Manual)   
4.2 Antivirus 
4.2.1 Ensure Antivirus Definition Push Updates are Configured 
(Automated) 
  
4.2.2 Apply Antivirus Security Profile to Policies (Manual)   
4.2.3 Enable Outbreak Prevention Database (Automated)   
4.2.4 Enable AI /heuristic based  malware detection 
(Automated) 
  
4.2.5 Enable grayware detection on antivirus (Automated)   
4.3 DNS Filter 
4.3.1 Enable Botnet C&C Domain Blocking DNS Filter 
(Automated) 
  
4.3.2 Ensure DNS Filter logs all DNS queries and responses 
(Manual) 
  
4.4 Application Control 
4.4.1 Block high risk categories on Application Control 
(Manual) 
 

Page 119 
CIS Benchmark Recommendation Set 
Correctly 
Yes No 
4.4.2 Block applications running on non-default ports 
(Automated) 
  
4.4.3 Ensure all Application Control related traffic are logged 
(Manual) 
  
5 Security Fabric 
5.1 Automation 
5.1.1 Enable Compromised Host Quarantine (Automated)   
5.2 Fabric Connectors 
5.2.1 Configure Root FortiGate for Security Fabric 
5.2.1.1 Ensure Security Fabric is Configured (Automated)   
6 VPN 
6.1 SSL VPN 
6.1.1 Apply a Trusted Signed Certificate for VPN Portal 
(Manual) 
  
6.1.2 Enable Limited TLS Versions for SSL VPN (Manual)   
7 Users and Authentication 
7.1 Configuring the maximum login attempts and lockout 
period (Automated) 
  
8 Logs and Reports 
8.1 Enable Logging 
8.1.1 Enable Event Logging (Automated)   
8.2 Encrypt Logs Sent to FortiAnalyzer / FortiManager 
8.2.1 Encrypt Log Transmission to FortiAnalyzer / 
FortiManager (Automated) 
 

Page 120 
CIS Benchmark Recommendation Set 
Correctly 
Yes No 
8.3 Centralized Logging and Reporting 
8.3.1 Centralized Logging and Reporting (Automated)  

Page 121 
Appendix: CIS Controls v7 IG 1 Mapped 
Recommendations 
Recommendation Set 
Correctly 
Yes No 
2.1.1 Ensure 'Pre-Login Banner' is set   
2.1.5 Ensure hostname is set   
2.1.6 Ensure the latest firmware is installed   
2.2.2 Ensure administrator password retries and lockout time 
are configured   
2.3.1 Ensure SNMP agent is disabled   
2.4.1 Ensure default 'admin' password is changed   
2.4.3 Ensure admin accounts with different privileges having 
their correct profiles assigned  

Page 122 
Appendix: CIS Controls v7 IG 2 Mapped 
Recommendations 
Recommendation Set 
Correctly 
Yes No 
1.1 Ensure DNS server is configured   
2.1.1 Ensure 'Pre-Login Banner' is set   
2.1.2 Ensure 'Post-Login-Banner' is set   
2.1.3 Ensure timezone is properly configured   
2.1.4 Ensure correct system time is configured through NTP   
2.1.5 Ensure hostname is set   
2.1.6 Ensure the latest firmware is installed   
2.2.1 Ensure 'Password Policy' is enabled   
2.2.2 Ensure administrator password retries and lockout time 
are configured   
2.3.1 Ensure SNMP agent is disabled   
2.3.2 Ensure only SNMPv3 is enabled   
2.4.1 Ensure default 'admin' password is changed   
2.4.2 Ensure all the login accounts having specific trusted 
hosts enabled   
2.4.3 Ensure admin accounts with different privileges having 
their correct profiles assigned   
2.4.4 Ensure idle timeout time is configured   
2.4.5 Ensure only encrypted access channels are enabled   
3.2 Ensure that policies do not use  "ALL" as Service  

