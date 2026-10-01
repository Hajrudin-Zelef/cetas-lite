---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-21
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["regulation"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [3315, 3525]
sha256: 877f00aa09c69203848a3dd97d13a4405bf6dc83a0a4afd0136d9ee307e535c8
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 125 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.5 Collect Detailed Audit Logs 
 Configure detailed audit logging for enterprise assets containing sensitive data. 
Include event source, date, username, timestamp, source addresses, destination 
addresses, and other useful elements that could assist in a forensic investigation. 
 ● ● 
v7 
6.3 Enable Detailed Logging 
 Enable system logging to include detailed information such as an event source, 
date, user, timestamp, source addresses, destination addresses, and other useful 
elements. 
 ● ●

Page 126 
Internal Only - General 
2.2.3 Set 'logging console critical' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Verify logging to device console is enabled and limited to a rational severity level to 
avoid impacting system performance and management. 
Rationale: 
This configuration determines the severity of messages that will generate console 
messages. Logging to console should be limited only to those messages required for 
immediate troubleshooting while logged into the device. This form of logging is not 
persistent; messages printed to the console are not stored by the router. Console 
logging is handy for operators when they use the console. 
Impact: 
Logging critical messages at the console is important for an organization managing 
technology risk. The 'logging console' command should capture appropriate severity 
messages to be effective. 
Audit: 
Perform the following to determine if the feature is enabled: 
Verify a command string result returns 
 
hostname#show run | incl logging console 
Remediation: 
Configure console logging level. 
 
hostname(config)#logging console critical 
Default Value: 
Tthe default is to log all messages 
Additional Information: 
The console is a slow display device. In message storms some logging messages may 
be silently dropped when the console queue becomes full. Set severity levels 
accordingly.

Page 127 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.5 Collect Detailed Audit Logs 
 Configure detailed audit logging for enterprise assets containing sensitive data. 
Include event source, date, username, timestamp, source addresses, destination 
addresses, and other useful elements that could assist in a forensic investigation. 
 ● ● 
v7 
6.3 Enable Detailed Logging 
 Enable system logging to include detailed information such as an event source, 
date, user, timestamp, source addresses, destination addresses, and other useful 
elements. 
 ● ●

Page 128 
Internal Only - General 
2.2.4 Set IP address for 'logging host' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Log system messages and debug output to a remote host. 
Rationale: 
Cisco routers can send their log messages to a Unix-style Syslog service. A syslog 
service simply accepts messages and stores them in files or prints them according to a 
simple configuration file. This form of logging is best because it can provide protected 
long-term storage for logs (the devices internal logging buffer has limited capacity to 
store events.) In addition, logging to an external system is highly recommended or 
required by most security standards. If desired or required by policy, law and/or 
regulation, enable a second syslog server for redundancy. 
Impact: 
Logging is an important process for an organization managing technology risk. The 
'logging host' command sets the IP address of the logging host and enforces the logging 
process. 
Audit: 
Perform the following to determine if a syslog server is enabled: 
Verify one or more IP address(es) returns 
 
hostname#sh log | incl logging host 
Remediation: 
Designate one or more syslog servers by IP address. 
 
hostname(config)#logging host {syslog_server} 
Default Value: 
System logging messages are not sent to any remote host. 
References: 
1. http://www.cisco.com/en/US/docs/ios/netmgmt/command/reference/nm_09.html#
wp1082864

Page 129 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
13.1 Centralize Security Event Alerting 
 Centralize security event alerting across enterprise assets for log correlation 
and analysis. Best practice implementation requires the use of a SIEM, which 
includes vendor-defined event correlation alerts. A log analytics platform configured 
with security-relevant correlation alerts also satisfies this Safeguard. 
 ● ● 
v8 13.11 Tune Security Event Alerting Thresholds 
 Tune security event alerting thresholds monthly, or more frequently.   ● 
v7 
6.6 Deploy SIEM or Log Analytic tool 
 Deploy Security Information and Event Management (SIEM) or log analytic tool 
for log correlation and analysis. 
 ● ● 
v7 
6.8 Regularly Tune SIEM 
 On a regular basis, tune your SIEM system to better identify actionable events 
and decrease event noise. 
  ●

Page 130 
Internal Only - General 
2.2.5 Set 'logging trap informational' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Limit messages logged to the syslog servers based on severity level informational. 
Rationale: 
This determines the severity of messages that will generate simple network 
management protocol (SNMP) trap and or syslog messages. This setting should be set 
to either "debugging" (7) or "informational" (6), but no lower. 
Impact: 
Logging is an important process for an organization managing technology risk. The 
'logging trap' command sets the severity of messages and enforces the logging 
process. 
Audit: 
Perform the following to determine if a syslog server for SNMP traps is enabled: 
Verify "level informational" returns 
 
hostname#sh log | incl trap logging 
Remediation: 
Configure SNMP trap and syslog logging level. 
 
hostname(config)#logging trap informational 
Default Value: 
Disabled 
References: 
1. http://www.cisco.com/en/US/docs/ios/netmgmt/command/reference/nm_09.html#
wp1015177

Page 131 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.5 Collect Detailed Audit Logs 
 Configure detailed audit logging for enterprise assets containing sensitive data. 
Include event source, date, username, timestamp, source addresses, destination 
addresses, and other useful elements that could assist in a forensic investigation. 
 ● ● 
v7 
6.3 Enable Detailed Logging 
 Enable system logging to include detailed information such as an event source, 
date, user, timestamp, source addresses, destination addresses, and other useful 
elements. 
 ● ●

Page 132 
Internal Only - General 
2.2.6 Set 'service timestamps debug datetime' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Configure the system to apply a time stamp to debugging messages or system logging 
messages 
Rationale: 
Including timestamps in log messages allows correlating events and tracing network 
attacks across multiple devices. Enabling service timestamp to mark the time log 
messages were generated simplifies obtaining a holistic view of events enabling faster 
troubleshooting of issues or attacks. 
Impact: 
Logging is an important process for an organization managing technology risk and 
establishing a timeline of events is critical. The 'service timestamps' command sets the 
date and time on entries sent to the logging host and enforces the logging process. 
Audit: 
Perform the following to determine if the additional detail is enabled: 
Verify a command string result returns 
 
hostname#sh run | incl service timestamps 
Remediation: 
Configure debug messages to include timestamps. 
 
hostname(config)#service timestamps debug datetime {<em>msec</em>} show-
timezone 
Default Value: 
Time stamps are applied to debug and logging messages. 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/R_through_setup.html#GUID-DC110E59-D294-
4E3D-B67F-CCB06E607FC6

