---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-11
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [1559, 1776]
sha256: 082118f6b9b3843328b8e32138acf2dfd2322f2ec7f955ac4f602f46d496d5b1
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 56 
System -> Admin Profiles, select the profile and click on "Edit". 
Stage 2: verify the admin accounts. In the CLI: 
FGT1 #config system admin 
FGT1 (admin) # edit "support1" 
FGT1 (support1) # show full 
config system admin 
  edit "support1" 
    ... 
    set accprofile "tier_1" 
    ... 
  next 
end 
In the GUI, go to 
System -> Administrators, select the account and click "Edit" 
Remediation: 
In this example, I would like to provide the profile "tier_1" the ability to view and modify 
address objects. This sub-privilege is under fwgrp privilege. 
In CLI 
FGT1 # config system accprofile 
FGT1 (accprofile) # edit "tier_1" 
FGT1 (tier_1) # set fwgrp custom 
FGT1 (tier_1) # config fwgrp-permission 
FGT1 (fwgrp-permission) # set address read-write 
FGT1 (fwgrp-permission) # end 
FGT1 (tier_1) # end 
FGT1 # 
For the GUI, go to 
System -> Admin Profiles, select "tier_1" and click "Edit". On "Firewall", 
click on "Custom" and then click on "Read/Write" option for "Address". 
In the next example, I would like to assign the profile "tier_1" to the account "support1". 
In the CLI 
FGT1 # config system admin 
FGT1 (admin) # edit "support1" 
FGT1 (support1) # set accprofile "tier_1" 
FGT1 (support1) # end 
FGT1 # 
For the GUI, go to 
System -> Administrators, select "support1" and click "Edit". Under 
"Administrator Profile", select "tier_1". 
Default Value: 
By default, there are only 2 profiles: prof_admin and super_admin. You have to select a 
profile to create an admin account, the system will not automatically choose for you.

Page 57 
References: 
1. https://docs.fortinet.com/document/fortigate/latest/administration-
guide/294491/administrator-profiles 
Additional Information: 
You cannot change the profile of the account which you are currently logging in as. 
The profile "super_admin" cannot be deleted or modified. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
5.4 Restrict Administrator Privileges to Dedicated 
Administrator Accounts 
 Restrict administrator privileges to dedicated administrator accounts on 
enterprise assets. Conduct general computing activities, such as internet 
browsing, email, and productivity suite use, from the user’s primary, non-privileged 
account. 
● ● ● 
v7 
4.3 Ensure the Use of Dedicated Administrative Accounts 
 Ensure that all users with administrative account access use a dedicated or 
secondary account for elevated activities. This account should only be used for 
administrative activities and not internet browsing, email, or similar activities. 
● ● ●

Page 58 
2.4.4 Ensure idle timeout time is configured (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
The idle timeout period is the amount of time that an administrator will stay logged in to 
the GUI without any activity. 
Rationale: 
Best practice dictates settings admin idle timeout to prevent the risk of unauthorized 
access to the device by preventing someone from using a logged-in GUI on a PC that 
has been left unattended. 
Impact: 
This is to prevent someone from accessing the FortiGate if the management PC is left 
unattended. 
Audit: 
To check the idle timeout in the GUI: 
1) Login to FortiGate 
2) Go to 'System' > 'Settings'. 
3) In the 'Administration Settings' section, check the 'Idle timeout' value 
in minutes. 
To check the idle timeout in the CLI: 
get system global 
check the value of admintimeout in minutes 
Remediation: 
To change the idle timeout in the GUI: 
1) Login to FortiGate with Super Admin privileges 
2) Go to 'System' > 'Settings'. 
3) In the 'Administration Settings' section, set the 'Idle timeout' value to 
five minutes by typing 5. 
4) Click Apply. 
To change the idle timeout in the CLI:

Page 59 
config system global 
 set admintimeout 5 
end 
Default Value: 
By default, it is set to five minutes. 
References: 
1. https://docs.fortinet.com/document/fortigate/6.2.0/cookbook/215451/setting-the-
idle-timeout-time 
Additional Information: 
A setting of higher than 15 minutes will have a negative effect on a security rating score. 
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

Page 60 
2.4.5 Ensure only encrypted access channels are enabled 
(Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Allow only HTTPS access to the GUI and SSH access to the CLI 
Rationale: 
By only allowing encrypted access, we are making it harder to use "Man in the Middle" 
attack to sniff login credentials. 
Audit: 
In the CLI, when verifying the network interface, make sure that http and telnet are not 
in the allowaccess list 
FG1 # config system interface 
FG1 (interface) # edit port1 
FG1 (port1) # show 
config system interface 
 edit "port1" 
  ... 
  set allowaccess ssh https ping snmp 
  ... 
 next 
end 
In the web GUI, click on 
Network -> Interfaces, select the interface and click "Edit". In the 
interface setting page, make sure that HTTP and Telnet are not selected in 
the section "Administrative Access" 
Remediation: 
If HTTP or Telnet is in the allowaccess list, you will have to set that list again with the 
same elements except for http or telnet 
FG1 # config system interface 
FG1 (interface) # edit port1 
FG1 (port1) # set allowaccess ssh https ping snmp 
FG1 (port1) # end 
FG1 # 
In the web GUI, click on

Page 61 
Network -> Interfaces, select the interface and click "Edit". In the 
interface setting page, uncheck HTTP and Telnet in the section 
"Administrative Access". 
Default Value: 
By default, HTTP and Telnet are not enabled on any interface. 
References: 
1. https://docs.fortinet.com/document/fortigate/6.0.0/handbook/909236/configuring-
administrative-access-to-interfaces 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
3.10 Encrypt Sensitive Data in Transit 
 Encrypt sensitive data in transit. Example implementations can include: 
Transport Layer Security (TLS) and Open Secure Shell (OpenSSH). 
 ● ● 
v7 
4.5 Use Multifactor Authentication For All Administrative 
Access 
 Use multi-factor authentication and encrypted channels for all administrative 
account access. 
 ● ●

Page 62 
2.4.6 Apply Local-in Policies (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Configure Local-in Policies to control inbound traffic that is destined to a FortiGate 
interface. 
Rationale: 
Local-in Policies allow for more granular and specific control of all types of traffic that 
are destined for a FortiGate interface. They are not limited to management only 
protocols so they can extend past "trusted host" configurations and can be configured 
with source and destination addresses as well as services specifically. 
Impact: 
Local-in Policies are processed before "trusted host" configurations so it is important to 
validate that management access will be maintained once the Local-in policies are put 
in place. 
Audit: 
To review Local-in Policies you can enable the feature to see them in the GUI by going 
to 
System > Feature Visibility and turning on "Local-in policies" under the 
Additional Features Section. This will then add the section under "Policies 
and Objects" there will now be a section for "Local-in Policies" 
It can also be viewed through the CLI: 
config firewall local-in-policy  
show 
Remediation: 
Local-in Policies can only be configured through the CLI:

