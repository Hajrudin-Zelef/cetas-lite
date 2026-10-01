---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-19
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [2055, 2191]
sha256: e26dbdaef923b014872453b785382e3263d0309e3234d0383b422efaae049b08
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

SSL VPN 
Secure Sockets Layer (SSL) VPN is transparent to users and easy to manage, 
making it an attractive remote access security solution. An enterprise can 
expand its intranet anywhere on the Internet, including the PCs and Internet 
information platform, therefore promoting employee productivity, protecting 
enterprise data, as well as enabling partners and consultants to access the 
intranet. 
Based on the SSL/TLS supported by all standard browsers, the SSL VPN has 
enhanced SSL/TLS functions. 
SSL ensures data communication security in the following aspects: 
 Authentication: Before the SSL connection is established, mutual 
authentication between the client and server is required, and the 
authentication uses the digital certificate. The authentication can be 
performed either by the client on the server or mutually on each other. 
 Confidentiality: An encryption algorithm is used to encrypt the data to be 
transmitted. 
 Integrity: A data authentication algorithm is used to verify the received 
data. 
The proxy function of the USG6000 series helps you access web resources on 
the intranet using a web browser. Web proxy fully outstands the ease-of-use of 
SSL VPN. When a remote user sends a request to access intranet pages using a 
web browser, the USG6000 series receives and forwards the request to the 
intranet server, and sends the server response in web pages to the user. During 
transmission on the Internet, information of the web pages is encrypted in the 
SSL tunnel to ensure that web resources on the intranet are provided to remote 
users securely and truly. 
3.13 Application-Layer Security 
Service Awareness (SA) 
Traditional firewalls identifies applications and applies policies by port. If an 
application uses an ephemeral port for communication, the application may 
evade the detection of firewalls. 
SA of the USG6000 series implements in-depth analysis on packet payload to 
identify the real application type of traffic. It has the following features: 
 Multiple identification methods 
The USG6000 series uses several methods to accurately identify common 
protocols such as HTTP and applications such as facebook and WebMail. 
 Predefined identification rule database

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  47 
   
 
The USG6000 series incorporates a predefined rule database to identify 
applications. The rule database can be updated online to identify 
ever-increasing new applications. 
Huawei predefined rule database supports more than 5000 protocols and 
applications to meet identification requirements. 
 User-defined identification rules 
The USG6000 series also supports user-defined rules for application 
identification to meet differentiated requirements. 
You can define conditions such as the IP address, port, and content matching in 
application identification rules to identify protocols or applications that are not 
covered by the predefined rules. 
Intrusion Prevention System (IPS) 
IPS of the USG6000 series, based on in-depth application identification, 
implements application-layer analysis and detection on the traffic to accurately 
identify various network attack behaviors and defend against the attacks. The 
USG6000 series detects threats such as botnets, Trojan horses, and worms and 
attacks such as the SQL injection and XSS attacks. 
 Deployment mode 
Off-line deployment: The USG6000 series implements security detection, but 
not defense action or traffic cleaning. Traffic is free of any impact. 
In-line detection deployment: The USG6000 series implements security 
detection, but not defense actions. It modifies only some QoS and TTL 
information but does not discard packets. 
In-line defense deployment: The USG6000 series implements security 
detection and traffic cleaning. When a security threat is detected, the USG6000 
series applies defense actions, such as discarding packets, modifying packets, 
and limiting traffic. 
 Major features 
Detection based on predefined rules: You can configure predefined rules for 
users, including the policies that defend against vulnerability-based attacks, 
botnets, Trojan horses, worms, SQL injection attacks, and XSS attacks. You 
can choose and generate a set of signatures by object, severity, operating 
system, protocol type, and threat type and formulate predefined rules based on 
the signatures. You can also define exception signatures to exempt some 
objects. 
Detection based on user-defined rules: You can configure user-defined rules 
when necessary. A user-defined rule consists of the user-defined object and rule 
body. The rule body contains identification conditions for the decoded fields. 
Such a user-defined rule helps you flexibly meet the detection requirements for 
IPS. 
Correlation detection: The USG6000 series provides predefined correlation 
detection for some threats to identify the relationship between security threat 
events. Such correlation detection helps you discover in-depth threats.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  48 
   
 
Anti-evasion: Hackers may evade the detection of IPS to attack the target 
device or server. Anti-evasion ensures accurate detection, without missing any 
attacks or threats. 
Updates of the engine and signature database: The USG6000 series supports 
the online and offline updates of the engine and signature database to defend 
against new threats on the live network. 
Antivirus (AV) 
The A V feature of the USG6000 series implements application-layer inspection 
on traffic to analyze transmitted files, detect viruses, and blocks the transfer of 
virus-infected files, protecting the customer's server and PC. 
The USG6000 series provides the following A V functions: 
 Powerful application-layer protocol parsing 
The USG6000 series implements powerful application-layer protocol parsing 
to analyze file transfer actions and scan files for viruses. 
 Diversified file types 
The USG6000 series supports diversified file types, decompresses file 
packages for virus scanning, and identifies the real file types based on content 
to prevent detection evasion that may be conspired by changing file name 
extensions. 
 Flow-based A V detection 
The USG6000 series supports flow-based A V detection for high defense 
performance. 
 Update of the virus signature database 
The virus signature database can be updated for the device to detect new 
viruses on the live network. A V detection will not be interrupted while the virus 
signature database is updated. 
Data Filtering 
Data filtering of the USG6000 series implements application-layer analysis on 
the transmitted data, detects and blocks data at the application layer based on 
predefined filtering policies, and reduces the risks of unauthorized file transfers 
and sensitive information transmission. 
Data filtering consists of protocol data filtering, file blocking by type, and file 
blocking by data. They have different scanning and filtering objectives: 
 Protocol data filtering 
Some application-layer protocols carry information in protocol contents, such 
as the web page, forum, micro-blogging, and email contents. You can configure 
policies to filter protocol contents. 
Based on in-depth protocol identification, the USG6000 series identifies traffic 
that uses an ephemeral port to prevent detection evasion and misjudgment.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  49 
   
 
