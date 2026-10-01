---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-9
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [1197, 1395]
sha256: 334fdfbf45a6d70d91f819833c7b87bd4ac69885df78a8519f2ba260a05728c4
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 45 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD45755 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.1 Establish and Maintain a Secure Configuration Process 
 Establish and maintain a secure configuration process for enterprise assets 
(end-user devices, including portable and mobile, non-computing/IoT devices, and 
servers) and software (operating systems and applications). Review and update 
documentation annually, or when significant enterprise changes occur that could 
impact this Safeguard. 
● ● ● 
v7 
5.1 Establish Secure Configurations 
 Maintain documented, standard security configuration standards for all 
authorized operating systems and software. 
● ● ●

Page 46 
2.3.2 Ensure only SNMPv3 is enabled (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Ensuring that only SNMPv3 service is enabled and SNMPv1, SNMPv2c are disabled. 
Rationale: 
SNMP Version 3 provides security enhancements that are not available in SNMP 
Version 1 or SNMP Version 2c. SNMP Versions 1 and 2c transmit data between the 
SNMP server and SNMP agent in clear text. SNMP Version 3 adds authentication and 
privacy options to secure protocol operations. Some firewalls need to be constantly 
monitored of its performance and status. Especially if the firewalls are critical to the 
operation. Enabling SNMPv3 will ensure that the firewall is monitored properly. 
Impact: 
Some older SNMP server that only run SNMPv1 or SNMPv2C will not be able to query 
to this firewall. 
Audit: 
From CLI, check to make sure that there is not any community for SNMPv1 or 
SNMPv2c and only SNMPv3 users are there. Also make sure that SNMP Agent is 
enabled.

Page 47 
FGT1 # config system snmp sysinfo 
FGT1 (sysinfo) # show 
config system snmp sysinfo 
 set status enable 
 ... 
end 
FGT1 (sysinfo) # end 
FGT1 # config system snmp community 
FGT1 (community) # show 
config system snmp community 
end 
FGT1 (community) # end 
FGT1 # config system snmp user 
FGT1 (user) # show 
config system snmp user 
    edit "snmp_test" 
        set security-level auth-priv 
        set auth-proto sha256 
        set auth-pwd ENC xxxxxx 
        set priv-proto aes256 
        set priv-pwd ENC xxxxxx 
    next 
end 
In the GUI, go to 
System -> SNMP. Make sure that SNMP agent is enabled.  Make sure that there 
is not any SNMPv1/2c community. Make sure that there is at least 1 SNMPv3 
user in the list. 
Remediation: 
To enable SNMP agent 
in CLI 
FGT1 # config system snmp sysinfo 
FGT1 (sysinfo) # set status enable 
FGT1 (sysinfo) # end 
In GUI, go to System -> SNMP and enable SNMP Agent. 
To delete SNMPv1/2c communities 
In this example, we'll delete community "public" 
in CLI 
FGT1 # config system snmp community 
FGT1 (community) # delete public 
FGT1 (community) # end 
FGT # 
In the GUI, go to 
System -> SNMP, select the community and click on the Delete button. 
To add SNMPv3 User 
in CLI

Page 48 
FGT1 # config system snmp user 
FGT1 (user) # edit "snmp_test" 
FGT1 (snmp_test) # set security-level auth-priv 
FGT1 (snmp_test) # set auth-proto sha256 
FGT1 (snmp_test) # set auth-pwd xxxx 
FGT1 (snmp_test) # set priv-proto aes256 
FGT1 (snmp_test) # set priv_pwd xxxx 
FGT1 (snmp_test) # end 
FGT1 # 
In the GUI, go to 
System -> SNMP, under SNMPv3, click on "Create New" button. Select 
"Authentication" and choose SHA256 as Authentication algorithm. Click 
"Change" to type in the password. Also select option "Private", choose AES256 
as Encryption Algorithm. Click on Change to change the password. Click "OK" 
to add the new user. Click apply to apply the new setting into the current 
config. 
Default Value: 
By default, SNMP agent is disabled. 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD45755 
2. https://docs.fortinet.com/document/fortigate/6.4.0/administration-
guide/457149/snmp-v3-users 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
12.3 Securely Manage Network Infrastructure 
 Securely manage network infrastructure. Example implementations include 
version-controlled-infrastructure-as-code, and the use of secure network 
protocols, such as SSH and HTTPS.  
 ● ● 
v7 
11.5 Manage Network Devices Using Multi-Factor 
Authentication and Encrypted Sessions 
 Manage all network devices using multi-factor authentication and encrypted 
sessions. 
 ● ●

Page 49 
2.4 Administrators and Admin Profiles

Page 50 
2.4.1 Ensure default 'admin' password is changed (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Before deploying any new FortiGate, it is important to change the password of the 
default admin account. 
It is also recommended that you change even the user name of the default admin 
account; however, since you cannot change the user name of an account that is 
currently in use, a second administrator account must be created in order to do this. 
Rationale: 
Default credentials are well documented by most vendors including Fortinet. Therefore, 
it will be one of the first things that will be tried to illegally gain access to the system. 
Impact: 
if not changed, then any scripts that use default credentials will be able to access the 
system. 
Audit: 
Using both CLI and GUI, in the username field put in "admin", leave the password field 
blank and proceed. If it's checked out, it means that the default password is still in place 
and needs to be changed. 
Remediation: 
In the CLI, to change the password of account "admin" 
FG1 # config system admin 
FG1 (admin) # edit "admin" 
FG1 (admin) # set password <your passwords> 
FG1 (admin) # end 
FG1 # 
To change the default password in the GUI:

Page 51 
1) Login to FortiGate with admin account 
2) Go to System > Administrators. 
3) Edit the admin account. 
4) Click Change Password. 
5) If applicable, enter the current password in the Old Password field. 
6) Enter a password in the New Password field, then enter it again in the 
Confirm Password field. 
7) Click OK. 
Default Value: 
By default, your FortiGate has an administrator account set up with the username admin 
and no password. In order to prevent unauthorized access to FortiGate, it is highly 
recommended that you add a password to this account. 
Username: admin The default admin account does not have any password. Just 
leave it blank 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD48763 
2. https://docs.fortinet.com/document/fortigate/6.2.0/cookbook/99980/default-
administrator-password 
Additional Information: 
In FortiOS 6.2.1 and later, adding a password to the admin administrator is mandatory. 
You will be prompted to configure it the first time you log in to the FortiGate using that 
account, after a factory reset, and after a new image installation. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.7 Manage Default Accounts on Enterprise Assets and 
Software 
 Manage default accounts on enterprise assets and software, such as root, 
administrator, and other pre-configured vendor accounts. Example 
implementations can include: disabling default accounts or making them 
unusable. 
● ● ● 
v7 
4.2 Change Default Passwords 
 Before deploying any new asset, change all default passwords to have values 
consistent with administrative level accounts. 
● ● ●

