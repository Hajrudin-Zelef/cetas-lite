---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-6
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "United States"]
dates: []
keywords: ["incident", "parameters"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [595, 789]
sha256: a4231d5ad2a90ec788d89c65707325f8092c367532b8d522cc4771dd919da5f4
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 24 
2.1.3 Ensure timezone is properly configured (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Sets the local time zone information so that the time displayed by the device is more 
relevant to those who are viewing it. 
Rationale: 
Having a correct time set on the device is important for two main reasons. The first 
reason is that digital certificates compare this time to the range defined by their Valid 
From and Valid To fields to define a specific validity period. The second reason is to 
have relevant time stamps when logging information. Whether you are sending 
messages to a Syslog server, sending messages to an SNMP monitoring station, or 
performing packet captures, timestamps have little usefulness if you cannot be certain 
of their accuracy. 
Impact: 
For many features to work, including scheduling, logging, and SSL-dependent features, 
the FortiOS system time must be accurate. 
Audit: 
In the CLI, do the following command and check the result of timezone filed in the 
output 
FGT1 # get system global 
... 
timezone            : (GMT-8:00) Pacific Time (US & Canada) 
... 
Or from GUI, do the following: 
1) login to FortiGate 
2) Go to 'System' -> 'Settings'. 
3) Time Zone and NTP settings are under 'System Time' 
Remediation: 
In this example, we will set Eastern Timezone (GMT-5:00) for the Fortigate. Each 
timezone will have its corresponding ID. To find the correct ID, when you type in the 
command "set timezone ", also type the question mark '?' to list all of the available 
timezones and their IDs. The ID of the Eastern Timezone is 12 
In the CLI:

Page 25 
FGT1 # config system global 
FGT1 (global) # set timezone 12 
FGT1 (global) # end 
FGT1 # 
In the GUI, do the following: 
1) after login to fortigate, go to 'System' -> 'Settings' 
2) select '(GMT-5:00) Eastern Time (US & Canada)' under 'System Time' 
Default Value: 
Default value is (GMT-8:00) Pacific Time (US & Canada) 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD49018 
2. https://docs.fortinet.com/document/fortigate/6.2.0/cookbook/512210/setting-the-
system-time 
3. https://docs.fortinet.com/document/fortigate/7.0.0/administration-
guide/512210/setting-the-system-time 
Additional Information: 
Daylight savings time is enabled by default, and can only be configured in the CLI. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.4 Standardize Time Synchronization 
 Standardize time synchronization. Configure at least two synchronized time 
sources across enterprise assets, where supported. 
 ● ● 
v7 
6.1 Utilize Three Synchronized Time Sources 
 Use at least three synchronized time sources from which all servers and 
network devices retrieve time information on a regular basis so that timestamps 
in logs are consistent. 
 ● ●

Page 26 
2.1.4 Ensure correct system time is configured through NTP 
(Automated) 
Profile Applicability: 
•  Level 1 
Description: 
You can either manually set the FortiOS system time, or configure the device to 
automatically keep its system time correct by synchronizing with a Network Time 
Protocol (NTP) server. 
These settings enable the use of primary and secondary NTP servers to provide 
redundancy in case of a failure involving the primary NTP server. 
Rationale: 
NTP enables the device to maintain accurate time and date when receiving updates 
from a reliable NTP server. Accurate timestamps are critical when correlating events 
with other systems, troubleshooting, or performing investigative work. Logs and certain 
cryptographic functions, such as those utilizing certificates, rely on accurate time and 
date parameters. In addition, rules referencing a Schedule object will not function as 
intended if the device’s time and date are incorrect. For additional security, 
authenticated NTP can be utilized. If Symmetric Key authentication is selected, only 
SHA1 should be used, as MD5 is considered severely compromised. 
Impact: 
For many features to work, including scheduling, logging, and SSL-dependent features, 
the FortiOS system time must be accurate. 
Audit: 
In the CLI:

Page 27 
FGT1 # diag sys ntp status 
synchronized: yes, ntpsync: enabled, server-mode: enabled 
  
ipv4 server(ntp2.fortiguard.com) 208.91.114.23 -- reachable(0xff) S:3 T:54 
    server-version=4, stratum=1 
    reference time is e12361d5.f27e0322 -- UTC Wed Sep 11 12:06:45 2019 
    clock offset is -0.001569 sec, root delay is 0.000000 sec 
    root dispersion is 0.010269 sec, peer dispersion is 19 msec 
  
ipv4 server(ntp1.fortiguard.com) 208.91.115.123 -- reachable(0xff) S:3 T:54 
selected 
    server-version=4, stratum=1 
    reference time is e12361d4.4f8b22a5 -- UTC Wed Sep 11 12:06:44 2019 
    clock offset is -0.000652 sec, root delay is 0.000000 sec 
    root dispersion is 0.010284 sec, peer dispersion is 8 msec 
  
ipv4 server(ntp2.fortiguard.com) 208.91.113.71 -- reachable(0xff) S:3 T:54 
    server-version=4, stratum=2 
    reference time is e12361d6.4caf57ab -- UTC Wed Sep 11 12:06:46 2019 
    clock offset is -0.004814 sec, root delay is 0.000137 sec 
    root dispersion is 0.011154 sec, peer dispersion is 3 msec 
  
ipv4 server(ntp1.fortiguard.com) 208.91.113.70 -- reachable(0xff) S:3 T:54 
    server-version=4, stratum=2 
    reference time is e123617b.c98e2059 -- UTC Wed Sep 11 12:05:15 2019 
    clock offset is -0.005106 sec, root delay is 0.000122 sec 
    root dispersion is 0.013382 sec, peer dispersion is 6 msec 
Remediation: 
You can only customize NTP setting using CLI. In this example, we'll assign 
pool.ntp.org as primary NTP server and 1.1.1.1 as secondary NTP server. 
FGT1 # config system ntp 
FGT1 (ntp) # set type custom 
FGT1 (ntp) # config ntpserver 
FGT1 (ntpserver) # edit 1 
FGT1 (1) # set server pool.ntp.org 
FGT1 (1) # next 
FGT1 (ntpserver) # edit 2 
FGT1 (2) # set server 1.1.1.1 
FGT1 (2) # end 
FGT1 (ntp) # end 
FGT1 # 
Default Value: 
By default, Fortinet uses the NTPs server of the FortiGuard 
References: 
1. https://docs.fortinet.com/document/fortigate/6.2.0/cookbook/512210/setting-the-
system-time 
2. https://kb.fortinet.com/kb/documentLink.do?externalID=FD49018

Page 28 
3. https://docs.fortinet.com/document/fortigate/7.0.0/administration-
guide/512210/setting-the-system-time 
Additional Information: 
Daylight savings time is enabled by default, and can only be configured in the CLI. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.4 Standardize Time Synchronization 
 Standardize time synchronization. Configure at least two synchronized time 
sources across enterprise assets, where supported. 
 ● ● 
v7 
6.1 Utilize Three Synchronized Time Sources 
 Use at least three synchronized time sources from which all servers and 
network devices retrieve time information on a regular basis so that timestamps 
in logs are consistent. 
 ● ●

Page 29 
2.1.5 Ensure hostname is set (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Changes the device default hostname. 
Rationale: 
The device hostname plays an important role in asset inventory and identification as a 
security requirement, but also in the public keys and certificate deployments as well as 
when correlating logs from different systems during an incident handling. 
Audit: 
In CLI 
get system global 
 ... 
 hostname            : FG1 
 ... 
In GUI, go to 'System' -> 'Settings', check the field 'Hostname' 
Remediation: 
In CLI, set the hostname to 'New_FGT1' as follows: 
FGT1 # config system global 
FGT1 (global) # set hostname "New_FGT1" 
FGT1 (global) # end 
New_FGT1 # 
or In GUI, go to 'System' -> 'Settings', update the field 'Hostname' with the new 
hostname, and click "Apply" 
Default Value: 
The default value of the hostname is the model number of the unit. Example: 'FortiGate 
2000E' 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD48765

