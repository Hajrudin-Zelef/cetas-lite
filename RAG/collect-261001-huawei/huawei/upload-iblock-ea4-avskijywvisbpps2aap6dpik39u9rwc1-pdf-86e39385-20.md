---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-20
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [2192, 2324]
sha256: a6f93c47d9ac4ce4b3bf8140420b4e1a3e569eb93dfc00a90611fa44193db653
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

The USG6000 series supports in-depth protocol decoding, multi-layer carrier 
protocol decoding, compression and decompression, and normalization to 
prevent application-layer detection evasion. 
 File blocking by type 
The USG6000 series filters application-layer files by type to block high-risk 
files and confidential files In addition, the USG6000 series filters transferred 
files by file name extension and real file type. 
During the filtering by real file type, the USG6000 series identifies the real 
type of transferred files based on content to prevent detection evasion. 
The USG6000 series decompresses compressed files and filters the files by real 
file type. 
 File blocking by data 
The USG6000 series implements in-depth analysis on file content and filters 
files by data to prevent information leaks and unauthorized information input. 
During the filtering by real file type, the USG6000 series identifies the real 
type of transferred files based on content to prevent evasion. 
The USG6000 series decompresses file packages and filters the files by real 
file type. 
The USG6000 series also supports data normalization to prevent detection 
evasion using coding technologies. 
HTTPS Traffic Defense 
HTTPS traffic defense of the USG6000 series analyzes HTTPS traffic and 
provides application-layer protection after the decryption to prevent malicious 
traffic from evading detection through the HTTPS channel. 
HTTPS traffic defense helps the USG6000 series decrypt HTTPS traffic. After 
the decryption is complete, the USG6000 series processes the traffic as it does 
to HTTP traffic. HTTPS traffic defense has the following operations: 
 SSL proxy 
SSL traffic cannot be decrypted in listening mode. To decrypt HTTPS 
traffic, the USG6000 series implements SSL proxy as follows: 
1. After receiving an SSL negotiation request from a client, the USG6000 
series serves as a server and negotiates with the client using its own 
certificate to set up an SSL tunnel. 
2. The USG6000 series initiates SSL negotiation to the real server to set up 
an SSL tunnel. 
3. The USG6000 series works as a transparent proxy server and forwards the 
traffic of the client and server. After receiving traffic from one tunnel end, 
the USG6000 series decrypts the traffic, implements application-layer 
detection, encrypts the traffic, and sends it to the other tunnel end. 
 Application-layer detection

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  50 
   
 
After SSL traffic is decrypted, the USG6000 series processes the HTTPS URL 
filtering as it does to HTTP traffic. 
3.14 Sound Maintenance and Management System 
Diversified Management Methods 
The USG6000 series performs local or remote maintenance using the following 
methods: 
 Local configuration and maintenance using the console port 
 Local or remote configuration and maintenance using Telnet 
 Secure Shell (SSH) maintenance and management. It provides 
information security guarantee and powerful authentication on an insecure 
network to defend against attacks such as IP spoofing and plain-text 
password interception. 
 Web- and Webs-based GUI configuration and maintenance 
 Unified management by Huawei NMS 
SNMP-based Terminal System Management 
The USG6000 series supports SNMP (v1/v2/v3) and the Client/Server model 
and can be managed by the NMS workstation such as Huawei eSight. 
3.15 Comprehensive Log Report System 
The USG6000 series collects statistics on the interface traffic and sessions 
during its operating to provide reference for the NMS, generate log information 
for other modules to make decisions, or deliver these information to users for 
debugging use. The users can customize logs by configuring the USG6000 
series to collect statistics only on the interested information. 
Logs are used to check the operating status of the device, analyze the network 
status, and locate the problem, providing references for system diagnosis and 
maintenance. 
The system log of the USG6000 series provides an after-the-event audit mode. 
A router provides detailed logs on all operation records and attacks, as well as 
log query and filtering methods to facilitate log query and analysis. 
The generated log information can be displayed using the console port or 
Telnet. It can be saved on a device or exported to the log server through the 
syslog protocol. 
Local Log Storage 
The USG6000 series supports hard disk cards to store generated logs.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  51 
   
 
When no log server is configured, you can use the local hard disk to store logs. 
If the local hard disk is full, you can enable the USG6000 series to discard the 
latest logs or use the latest logs to overwrite the oldest logs. 
You can export log files from the hard disk to prevent log loss. 
Log Server 
To receive and store router logs, Huawei has launched dedicated log server 
software. Based on this software, users can conveniently browse, query, and 
analyze logs. The log server software consists of the front-end management and 
back-end process parts. Front-end management provides operations, such as 
database configuration, log configuration, and log category query. Back-end 
processes include the log collection and monitoring processes. You can use the 
log server software to customize receiving log types and provides log storage, 
query, export, and backup functions. 
Two Log Export Modes 
The USG6000 series supports the output of syslogs in text. In addition, the 
USG6000 series can create information tables based on flow status and 
generate fast binary logs for the heavy traffic passing through. Compared with 
syslogs, binary logs better suits the scenario in which log contents is massive 
and therefore require a higher network speed. 
Abundant Logs 
The USG6000 series provides complete and unified log information. The types 
of logs include: 
 Traffic log 
The USG6000 series generates traffic logs by flow for the passing traffic. A log 
of this type contains the source address, source port, destination address, 
destination port, Internet-access user, application, flow start time, flow end 
time, and flow status. For a flow that uses NA T, the related log also contains 
information about the address and port after NA T. 
You can view global traffic conditions by user and application to learn about 
the bandwidth usage and security policy implementation. 
 Attack defense log 
When massive attacks occur, the USG6000 series applies the queue mechanism 
to provide log alarm information for the attack defense feature that routers 
support, and generates alarms in SYSLOG mode. Alarm information includes 
the attack source (source address) and attack type. 
 Threat log 
When detecting threats, the USG6000 series generates threat logs. The threat 
logs record the detected network threats such as viruses, intrusions, DDoS 
attacks, botnets, and worms and the defense against them. They help you learn 
about the current and historical threat events, modify policies, and take defense 
measures.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  52 
   
 
