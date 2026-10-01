---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-10
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [1193, 1356]
sha256: f3550242623c9892bbd8e4423e41324f212ce09719f5d5f2dc91ded43e9168d2
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 46 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
12.8 Establish and Maintain Dedicated Computing 
Resources for All Administrative Work 
 Establish and maintain dedicated computing resources, either physically or 
logically separated, for all administrative tasks or tasks requiring administrative 
access. The computing resources should be segmented from the enterprise's 
primary network and not be allowed internet access. 
  ● 
v7 
11.7 Manage Network Infrastructure Through a Dedicated 
Network 
 Manage the network infrastructure across network connections that are 
separated from the business use of that network, relying on separate VLANs or, 
preferably, on entirely different physical connectivity for management sessions for 
network devices. 
 ● ●

Page 47 
Internal Only - General 
1.2.5 Set 'access-class' for 'line vty' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
The 'access-class' setting restricts incoming and outgoing connections between a 
particular vty (into a Cisco device) and the networking devices associated with 
addresses in an access list. 
Rationale: 
Restricting the type of network devices, associated with the addresses on the access-
list, further restricts remote access to those devices authorized to manage the device 
and reduces the risk of unauthorized access. 
Impact: 
Applying 'access'class' to line VTY further restricts remote access to only those devices 
authorized to manage the device and reduces the risk of unauthorized access. 
Conversely, using VTY lines with 'access class' restrictions increases the risks of 
unauthorized access. 
Audit: 
Perform the following to determine if the ACL is set: 
Verify you see the access-class defined 
hostname#sh run | sec vty <line-number> <ending-line-number> 
Remediation: 
Configure remote management access control restrictions for all VTY lines. 
hostname(config)#line vty <line-number> <ending-line-number> 
hostname(config-line)# access-class <vty_acl_number> in 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a2.html#GUID-
FB9BC58A-F00A-442A-8028-1E9E260E54D3

Page 48 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
12.8 Establish and Maintain Dedicated Computing 
Resources for All Administrative Work 
 Establish and maintain dedicated computing resources, either physically or 
logically separated, for all administrative tasks or tasks requiring administrative 
access. The computing resources should be segmented from the enterprise's 
primary network and not be allowed internet access. 
  ● 
v7 
11.7 Manage Network Infrastructure Through a Dedicated 
Network 
 Manage the network infrastructure across network connections that are 
separated from the business use of that network, relying on separate VLANs or, 
preferably, on entirely different physical connectivity for management sessions for 
network devices. 
 ● ●

Page 49 
Internal Only - General 
1.2.6 Set 'exec-timeout' to less than or equal to 10 minutes for 
'line aux 0' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
If no input is detected during the interval, the EXEC facility resumes the current 
connection. If no connections exist, the EXEC facility returns the terminal to the idle 
state and disconnects the incoming session. 
Rationale: 
This prevents unauthorized users from misusing abandoned sessions. For example, if 
the network administrator leaves for the day and leaves a computer open with an 
enabled login session accessible. There is a trade-off here between security (shorter 
timeouts) and usability (longer timeouts). Review your local policies and operational 
needs to determine the best timeout value. In most cases, this should be no more than 
10 minutes. 
Impact: 
Organizations should prevent unauthorized use of unattended or abandoned sessions 
by an automated control. Enabling 'exec-timeout' with an appropriate length of minutes 
or seconds prevents unauthorized access of abandoned sessions. 
Audit: 
Perform the following to determine if the timeout is configured: 
Verify you return a result 
hostname#show running-config all | sec line aux 0 
Remediation: 
Configure device timeout (10 minutes or less) to disconnect sessions after a fixed idle 
time. 
hostname(config)#line aux 0 
hostname(config-line)#exec-timeout <timeout_in_minutes> <timeout_in_seconds> 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/D_through_E.html#GUID-76805E6F-9E89-4457-
A9DC-5944C8FE5419

Page 50 
Internal Only - General 
Additional Information: 
Some Cisco devices don't even have an auxiliary port, therefore this control will fail. 
Adjusting this control logic: check for existence of auxiliary port (ie check for 'none_exist' 
of a section with 'aux') combined with a logical OR with the current control. in addition 
adjusting the regex to match the control's title. Adjusting the 'audit' section. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.3 Configure Automatic Session Locking on Enterprise 
Assets 
 Configure automatic session locking on enterprise assets after a defined period 
of inactivity. For general purpose operating systems, the period must not exceed 
15 minutes. For mobile end-user devices, the period must not exceed 2 minutes. 
● ● ● 
v7 16.11 Lock Workstation Sessions After Inactivity 
 Automatically lock workstation sessions after a standard period of inactivity. ● ● ●

Page 51 
Internal Only - General 
1.2.7 Set 'exec-timeout' to less than or equal to 10 minutes 'line 
console 0' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
If no input is detected during the interval, the EXEC facility resumes the current 
connection. If no connections exist, the EXEC facility returns the terminal to the idle 
state and disconnects the incoming session. 
Rationale: 
This prevents unauthorized users from misusing abandoned sessions. For example, if 
the network administrator leaves for the day and leaves a computer open with an 
enabled login session accessible. There is a trade-off here between security (shorter 
timeouts) and usability (longer timeouts). Review your local policies and operational 
needs to determine the best timeout value. In most cases, this should be no more than 
10 minutes. 
Impact: 
Organizations should prevent unauthorized use of unattended or abandoned sessions 
by an automated control. Enabling 'exec-timeout' with an appropriate length reduces the 
risk of unauthorized access of abandoned sessions. 
Audit: 
Perform the following to determine if the timeout is configured: 
Verify you return a result 
hostname#show running-config all | section line con 0 
Remediation: 
Configure device timeout (10 minutes or less) to disconnect sessions after a fixed idle 
time. 
hostname(config)#line con 0 
hostname(config-line)#exec-timeout <timeout_in_minutes> <timeout_in_seconds> 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/D_through_E.html#GUID-76805E6F-9E89-4457-
A9DC-5944C8FE5419

