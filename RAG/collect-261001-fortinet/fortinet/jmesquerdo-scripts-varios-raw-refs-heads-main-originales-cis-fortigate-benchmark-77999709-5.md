---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-5
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [426, 594]
sha256: 833b8c2f89d2a569cdba855cc9fb11dee10c8d8d8b46a697675e0c2fd768fddd
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 18 
FGT1 # config system interface 
FGT1 (interface) # edit "port1" 
FGT1 (port1) # unselect allowaccess ping https ssh snmp http radius-acct 
Note: 
1. Interface name may differ based on deployment. For this example, port1 is 
deployed as WAN interface. 
2. "unselect allowaccess" will only show services that you have enabled. If you 
have not enabled snmp on that interface, then snmp option will not be available. 
2 System Settings 
This topic contains information and best practices about FortiGate administration and 
system configuration.

Page 19 
2.1 General Settings

Page 20 
2.1.1 Ensure 'Pre-Login Banner' is set (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Configure a pre-login banner, ideally approved by the organization’s legal team. This 
banner should, at minimum, prohibit unauthorized access, provide notice of logging or 
monitoring, and avoid using the word “welcome” or similar words of invitation. 
Rationale: 
Through a properly stated login banner, the risk of unintentional access to the device by 
unauthorized users is reduced. Should legal action take place against a person 
accessing the device without authorization, the login banner greatly diminishes a 
defendant’s claim of ignorance. 
Impact: 
Login banners provide a definitive warning to any possible intruders that may want to 
access the FortiGate that certain types of activity are illegal, but at the same time, it also 
advises the authorized and legitimate users of their obligations relating to acceptable 
use. 
Audit: 
Run the following command in the CLI to verify the pre-login-banner is enabled: 
FG1 # get system global 
 ... 
 pre-login-banner    : enable 
 ... 
end 
In the GUI, to verify the content of the pre-login disclaimer message: 
1) go to 'System' -> 'Replacement Messages' 
2) from the top right side, select 'Extended View' 
3) find 'Pre-login Disclaimer Message' 
Remediation: 
Run the following command in the CLI to enable the pre-login-banner: 
FG1 # config system global 
FG1 (global) # set pre-login-banner enable 
FG1 (global) # end 
FG1 # 
In the GUI, to edit the content of the pre-login disclaimer message:

Page 21 
1. go to 'System' -> 'Replacement Messages' -> 'Extended View' -> 'Pre-login 
Disclaimer Message'. The edit screen is on the bottom right corner of the page. 
Click on "Save" after the editing is done. 
Default Value: 
the 'Pre-Login Banner' is disabled by default 
FG1 # config system global 
FG1 (global) # show 
config system global 
 ... 
 set pre-login-banner disable 
 ... 
end 
the warning message default value is as follows: 
PRE WARNING: 
This is a private computer system. Unauthorized access or use  
is prohibited and subject to prosecution and/or disciplinary  
action. All use of this system constitutes consent to  
monitoring at all times and users are not entitled to any  
expectation of privacy. If monitoring reveals possible evidence 
of violation of criminal statutes, this evidence and any other  
related information, including identification information about  
the user, may be provided to law enforcement officials. 
If monitoring reveals violations of security regulations or 
unauthorized use, employees who violate security regulations or 
make unauthorized use of this system are subject to appropriate  
disciplinary action. 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD33887 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.2 Establish and Maintain a Secure Configuration 
Process for Network Infrastructure 
 Establish and maintain a secure configuration process for network devices. 
Review and update documentation annually, or when significant enterprise 
changes occur that could impact this Safeguard. 
● ● ● 
v7 
5.1 Establish Secure Configurations 
 Maintain documented, standard security configuration standards for all 
authorized operating systems and software. 
● ● ●

Page 22 
2.1.2 Ensure 'Post-Login-Banner' is set (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Sets the banner after users successfully login. This is equivalent to Message of the Day 
(MOTD) in some other systems. 
Rationale: 
Network banners are electronic messages that provide notice of legal rights to users of 
computer networks. From a legal standpoint, banners have four primary functions. 
First, banners may be used to generate consent to real-time monitoring under Title III. 
Second, banners may be used to generate consent to the retrieval of stored files and 
records pursuant to ECPA. Third, in the case of government networks, banners may 
eliminate any Fourth Amendment "reasonable expectation of privacy" that government 
employees or other users might otherwise retain in their use of the government's 
network under O'Connor v. 
Impact: 
When post-login banner is enabled, some automated-script might be affected because 
both CLI and GUI need an acceptance action (press "A" or "Accept") to continue. 
Audit: 
Run the following command in the CLI to verify the post-login-banner is enabled: 
FG1 # get system global 
 ... 
 post-login-banner    : enable 
 ... 
In the GUI, to verify the content of the post-login disclaimer message: 
1) go to 'System' -> 'Replacement Messages' 
2) from the top right side, select 'Extended View' 
3) find 'Post-login Disclaimer Message' 
Remediation: 
Run the following command in the CLI to enable the post-login-banner: 
FG1 # config system global 
FG1 (global) # set post-login-banner enable 
FG1 (global) # end 
FG1 # 
In the GUI, to edit the content of the post-login disclaimer message, go to

Page 23 
System -> Replace Messages -> Extended View -> "Post-login Disclaimer 
Message". The edit screen is on the bottom right  corner of the page. Click 
on "Save" after the editing is done. 
Default Value: 
POST WARNING: This is a private computer system. Unauthorized access or use is 
prohibited and subject to prosecution and/or disciplinary action. All use of this system 
constitutes consent to monitoring at all times and users are not entitled to any 
expectation of privacy. If monitoring reveals possible evidence of violation of criminal 
statutes, this evidence and any other related information, including identification 
information about the user, may be provided to law enforcement officials. If monitoring 
reveals violations of security regulations or unauthorized use, employees who violate 
security regulations or make unauthorized use of this system are subject to appropriate 
disciplinary action. 
%%LAST_SUCCESSFUL_LOGIN%% %%LAST_FAILED_LOGIN%% 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD33887 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.2 Establish and Maintain a Secure Configuration 
Process for Network Infrastructure 
 Establish and maintain a secure configuration process for network devices. 
Review and update documentation annually, or when significant enterprise 
changes occur that could impact this Safeguard. 
● ● ● 
v7 
11.1 Maintain Standard Security Configurations for 
Network Devices 
 Maintain standard, documented security configuration standards for all 
authorized network devices. 
 ● ●

