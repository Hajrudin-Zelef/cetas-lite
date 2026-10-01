---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-21
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [2325, 2504]
sha256: 36f67862fee6a30cac4334e686af2f695e23800c378910585ab5f67a789e6149
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

 URL filtering log 
The USG6000 series implements URL filtering on intranet users who initiate 
web access based on the specified policy and records URL logs of the users. 
URL filtering logs help you learn about the URL access behaviors, alarms and 
blocking events generated when intranet users access URLs, and causes of the 
alarms and blocking events. 
 Data filtering log 
The USG6000 series implements data filtering and generates logs for the traffic 
that matches data filtering conditions. Data filtering logs help you learn about 
the risky user behaviors, alarms and blocking events generated when intranet 
users transfer files, send and receive emails, and access websites, and causes of 
the alarms and blocking events. 
 Mail filtering log 
The USG6000 series implements mail filtering and generates logs for the 
traffic that matches mail filtering conditions. Mail filtering logs help you learn 
about the protocol types, attachment quantities, and attachment sizes of user 
emails and the causes that legitimate emails are blocked and take appropriate 
measures. 
 Operation log 
Operation logs record all operations performed by administrators on the 
USG6000 series. The USG6000 series helps you learn about the logins, logouts, 
and configuration operations of all administrators and the device management 
history and enhance device security. 
 System log 
System logs record all key events during the system operating. Based on the 
logs, you can learn about the operating status of the device and locate the fault. 
 User activity log 
The USG6000 series logs user activities when the users access the Internet. 
User activity logs help you learn about user behaviors and user online records, 
such as the login time, Internet-access duration, and IP and MAC addresses 
used for login, discover abnormal user login and access behaviors, and take 
immediate measures. 
 Policy matching log 
The USG6000 series logs policy matching events. Policy matching logs help 
you learn about the events that policies are matched, determine whether 
policies are correctly configured and effective, and locate faults. 
 Audit log 
The USG6000 series supports the behavior audit and content audit functions 
and generates audit logs on user Internet-access behaviors and key contents. 
Based on audit logs, you can view the network behaviors of users. 
 Traffic monitoring log

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  53 
   
 
The USG6000 series monitors traffic by security zone and IP address, checks 
whether the rate or connection quantity reaches the upper limit or lower limit. 
The USG6000 series generates alarms and records logs when the upper limit is 
hit, and generates alarms to instruct the system to recover when the lower limit 
is hit. 
 Blacklist log 
The USG6000 series automatically adds the source IP address of any 
illegitimate user that it has detected to the blacklist and generates a blacklist log 
that records the host IP address and blacklisting reason. 
 Statistics information 
Flow statistics are recorded to help you learn about the operating status of a 
router. The flow statistics include total connection quantity, current connection 
quantity and half-open connection quantity, peak connection quantity, and 
discarded packet quantity. 
Statistics on attack packet quantities help you learn about the status of attack 
events. 
Diversified Reports 
The USG6000 series provides diversified reports that combine log information 
and intuitively display the information. You can customize reports to obtain 
only the data of your concern. 
Reports can be sent in an email to the administrator at the scheduled time. 
 Traffic report 
Traffic reports of the USG6000 series analyze traffic statistics, rankings, and 
trends by source address, destination address, user, application, application 
category, and application subcategory. 
The USG6000 series summarizes data of traffic logs and generates intuitive 
reports in different dimensions, which provide you visibility into network 
traffic status and help you determine traffic management methods. 
 Threat report 
Threat reports of the USG6000 series analyze threat times trends and rankings 
by threat type, user, attacker, target, threat name, virus, and attack defense. 
The USG6000 series summarizes data of threat logs and generates intuitive 
reports in different dimensions, which provide you visibility into latest threat 
behaviors, attackers, and victims and help you determine security defense 
methods. 
 URL report 
URL reports of the USG6000 series analyze URL access statistics, rankings, 
and trends by URL type and website. 
The USG6000 series summarizes data of URL logs and generates intuitive 
reports in different dimensions, which provide you visibility into the URLs or 
websites that are access the most times and users who frequently access 
illegitimate URLs and help you determine URL filtering policies.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  54 
   
 
 Policy matching report 
Policy matching reports of the USG6000 series analyze statistics on matching 
times and rankings by policy. 
The USG6000 series summarizes data of policy matching logs and generates 
intuitive reports in different dimensions, which provide you visibility into 
policy configuration and effectiveness and help you optimize policies.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  55 
   
 
4 Typical Networking 
4.1 Attack Defense 
Figure 4-1 Attack defense networking diagram 
DMZ
SYSLOG
WWW
DNS
Mail
LAN
Internal office area
PC PC
LAN Switch
USG
Hacker
Hacker
Internet
Router
Government or 
enterprise network
 
 
 The USG6000 series is deployed at the network ingress to prevent various 
attacks launched from the Internet and intranet. 
 The available deployment mode is as follows: using the mirroring port of 
the device, LAN Switch, and unified security gateway to defend against 
various attacks. 
 With the powerful anti-DoS function, the USG6000 series protects the 
resource hosts in the intranet to the greatest extent. 
 The USG6000 series can work in transparent or routing mode to meet 
different networking requirements. 
4.2 NAT 
Combined with the policy-based NAT function, the USG6000 series establishes 
a more secure network environment using the secure filtering function over

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  56 
   
 
NA T applications to enhance capabilities in defending against attacks and 
preventing unauthorized access. The following figure shows NAT networking 
diagram of the USG6000 series. 
Internet
USG
Providing NAT
   USG
Providing NAT
 
Log serverRADIUS server
 
Intranet
   
FTP server Mail server Web server
DMZ
Branch office
Mobile office
 
 Specific users in the enterprise can access the Internet, e-commerce, and 
online banking systems, and a shield is set up between the intranet and the 
Internet. 
 Remote branch offices or trustworthy partners can access internal servers 
(such as the web server and FTP server) in the DMZ through the firewall, 
but cannot access other intranet resources. 
 Internet users cannot access resources on the intranet and the DMZ or 
launch any attacks. 
 The USG6000 series supports bi-directional NA T and NAT ALG between 
the security zones of different security levels. 
 The efficient logging function provides NAT logs. 
4.3 Hot Standby 
The USG6000 series supports hot standby, which ensures uninterrupted 
services even when the active and standby firewalls are switching over. The 
following figure shows the networking diagram for hot standby.

