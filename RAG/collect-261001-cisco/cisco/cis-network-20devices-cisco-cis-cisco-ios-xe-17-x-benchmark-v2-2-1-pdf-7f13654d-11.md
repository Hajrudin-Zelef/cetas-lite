---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-11
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [1357, 1510]
sha256: 4858b973c8d8f2cc957c1211be0df16bf3a4826aea5e8bbc14b3194e0324ae61
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 52 
Internal Only - General 
Additional Information: 
adjusting the 'audit procedure': removing the 'Note' and changing the command to use 
from 'sh run | sec line con 0' to 'show running-config all | section line con 0'. 
Adjusting the regex from '^\s*(exec-timeout)\s*((10)|([0-9]))\s*$' to '^\s*(exec-
timeout)\s+((10\s+0\s*$)|([0-9]\s+([0-9]\s*$|[1-5][0-9]\s*$)))' 
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

Page 53 
Internal Only - General 
1.2.8 Set 'exec-timeout' to less than or equal to 10 minutes 'line 
vty' (Automated) 
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
Verify you return a result NOTE: If you set an exec-timeout of 10 minutes, this will not 
show up in the configuration 
hostname#sh line vty <tty_line_number> | begin Timeout 
Remediation: 
Configure device timeout (10 minutes or less) to disconnect sessions after a fixed idle 
time. 
hostname(config)#line vty {line_number} [ending_line_number] 
hostname(config-line)#exec-timeout <<span>timeout_in_minutes> 
<timeout_in_seconds</span>> 
References: 
1. https://www.cisco.com/c/en/us/td/docs/switches/datacenter/mds9000/sw/comma
nd/b_cisco_mds_9000_cr_book/l_commands.html#wp3716128869

Page 54 
Internal Only - General 
Additional Information: 
Adjusting the 'Audit Procedure': command currently listed will show the settings and 
status of line vty xx, but then you need to check all line vty from 0 to 98... And it doesn't 
really match what the Artifact is checking... I would suggest to use the 'show running-
config all | section line vty' command in the 'audit procedure' and look for the line 
starting with exec-timeout. 
adjusting the artifact for the regex to match with the control's title (...less than or equal to 
10 minutes). All line vty should have exec-timeout less than or equal to 10 minutes, then 
adjusting the Artifact's "# of config lines to match" from 'at least one' to 'all' 
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

Page 55 
Internal Only - General 
1.2.9 Set 'http Secure-server' limit (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Device management includes the ability to control the number of administrators and 
management sessions that manage a device. Limiting the number of allowed 
administrators and sessions per administrator based on account type, role, or access 
type is helpful in limiting risks related to denial-of-service (DoS) attacks. 
Rationale: 
This requirement addresses concurrent sessions for administrative accounts and does 
not address concurrent sessions by a single administrator via multiple administrative 
accounts. The maximum number of concurrent sessions should be defined based upon 
mission needs and the operational environment for each system. At a minimum, limits 
must be set for SSH, HTTPS, account of last resort, and root account sessions. Center 
for Internet Security recommends a limit of 2 
Audit: 
The result should show ip http secure-server with max connections on following line 
hostname#show run | inc ip http secure-server 
Remediation: 
hostname(config)#ip http max-connections 2 
References: 
1. NIST SP 800-53 :: AC-10 
Additional Information: 
limiting http max connections is useful only if http server or http secure-server, at least 
one, is enable. if HTTP server is disabled (Artifact #2 "HTTP server is disabled") and 
HTTPS server is disabled (Artifact #3 "HTTPS server is disabled"), then this control is 
passed, otherwise it needs to limit connections to a maximum of 2 (Artifact #1 "Max http 
secure-server limit"). Adjusting the regex from '^\sip\s+http\s+max-connections\s+0-2?$' 
to '^\sip\s+http\s+max-connections\s+[0-2]\s*$'

Page 56 
Internal Only - General 
CIS Controls: 
Controls Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v7 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 57 
Internal Only - General 
1.2.10 Set 'exec-timeout' to less than or equal to 10 min on 'ip 
http' (Automated) 
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
This prevents unauthorized users from misusing abandoned sessions. For example, if 
the network administrator leaves for the day and leaves a computer open with an 
enabled login session accessible. There is a trade-off here between security (shorter 
timeouts) and usability (longer timeouts). Review your local policies and operational 
needs to determine the best timeout value. In most cases, this should be no more than 
10 minutes. 
Audit: 
Perform the following to determine if the timeout is configured: 
sh run | beg ip http timeout-policy 
Remediation: 
Configure device timeout (10 minutes or less) to disconnect sessions after a fixed idle 
time. 
ip http timeout-policy idle 600 life {nnnn} requests {nn} 
Default Value: 
disabled

