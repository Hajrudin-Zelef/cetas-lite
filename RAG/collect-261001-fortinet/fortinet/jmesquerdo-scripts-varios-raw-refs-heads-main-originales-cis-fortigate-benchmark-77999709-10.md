---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-10
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [1396, 1558]
sha256: 9ba68d869eab3056ee8beb53bb618be1c4e50a9b3ff339b70ef0589bdf7c8fc2
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 52 
2.4.2 Ensure all the login accounts having specific trusted hosts 
enabled (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Configure an administrative account to be accessible only to someone who is using a 
trusted host. You can set a specific IP address for the trusted host or use a subnet. 
Rationale: 
Access to a firewall to perform administrative tasks should only come from specific 
network segments reserved for administrators only. This additional layer of security 
ensure that no one from anywhere else on the network able to login even with correct 
credentials. 
Impact: 
All access, from legitimate or illegitimate users, outside of allowed segment will be 
stopped. Thus, administrators working remotely will have to make sure that they have 
access to jump hosts that sit in the allowed segment. 
Audit: 
This example is to check if trusted hosts option is enabled for account "test_admin" and 
which trusted hosts are in the list 
FG1 # config system admin 
FG1 (admin) # edit "test_admin" 
FG1 (test_admin) # show 
config system admin 
 edit "test_admin" 
  ... 
  set trusthost1 10.0.0.0 255.255.255.0 
  set trusthost2 192.168.10.0 255.255.255.0 
         ... 
 next 
end  
In the web GUI, go to

Page 53 
System -> Administrators, select the account and click on edit. In the 
account setting page, make sure that "Restrict login to trusted hosts" are 
enabled and all the allowed hosts / subnets are in the list of trusted Host. 
Please take note that certain versions of FortiOS will only show the first 3 
trusted hosts in the list. If you want to see more, you have to click on the 
"+" sign as if you're adding a new item into the list. Keep clicking until 
you see an empty field of trusted host. That's when you know that you have 
reached the bottom of the list. 
Remediation: 
To remove a trusted host item from the list in CLI 
FG1 # config system admin 
FG1 (admin) # edit "test_admin" 
FG1 (test_admin) # unset trusthost1 
FG1 (test_admin) # end 
FG1 # 
To add a trusted host into the list in CLI 
FG1 # config system admin 
FG1 (admin) # edit "test_admin" 
FG1 (test_admin) # set trusthost6 1.1.1.1 255.255.255.255 
FG1 (test_admin) # end 
FG1 # 
Before adding an item, please make sure that it does not already exist. For example, if 
trusthost3 is already in the list, using it again will over-ride the existing host/network. 
In the web GUI, go to 
System -> Administrators, select the account and click on edit. In the 
account setting page, make sure that "Restrict login to trusted hosts" are 
enabled and all the allowed hosts / subnets are in the list of trusted Host. 
Please take note that certain versions of FortiOS will only show the first 3 
trusted hosts in the list. If you want to see more, you have to click on the 
"+" sign as if you're adding a new item into the list. Keep clicking until 
you see an empty field of trusted host. That's when you know that you have 
reached the bottom of the list. To add another trusted host, fill in the 
empty field of the new "Trusted Host". To remove a trusted host, simply erase 
everything in the field of that corresponding host. 
Default Value: 
By default, each account is accessible from everywhere , the host value is 0.0.0.0/0 
References: 
1. https://docs.fortinet.com/document/fortigate/6.0.0/cookbook/222079/using-a-
trusted-host-optional

Page 54 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
5.4 Restrict Administrator Privileges to Dedicated 
Administrator Accounts 
 Restrict administrator privileges to dedicated administrator accounts on 
enterprise assets. Conduct general computing activities, such as internet browsing, 
email, and productivity suite use, from the user’s primary, non-privileged account. 
● ● ● 
v8 
12.8 Establish and Maintain Dedicated Computing 
Resources for All Administrative Work 
 Establish and maintain dedicated computing resources, either physically or 
logically separated, for all administrative tasks or tasks requiring administrative 
access. The computing resources should be segmented from the enterprise's 
primary network and not be allowed internet access. 
  ● 
v7 
4.6 Use of Dedicated Machines For All Administrative 
Tasks 
 Ensure administrators use a dedicated machine for all administrative tasks or 
tasks requiring administrative access. This machine will be segmented from the 
organization's primary network and not be allowed Internet access. This machine 
will not be used for reading e-mail, composing documents, or browsing the Internet. 
  ● 
v7 
11.6 Use Dedicated Machines For All Network 
Administrative Tasks 
 Ensure network engineers use a dedicated machine for all administrative tasks 
or tasks requiring elevated access. This machine shall be segmented from the 
organization's primary network and not be allowed Internet access. This machine 
shall not be used for reading e-mail, composing documents, or surfing the Internet. 
 ● ● 
v7 
11.7 Manage Network Infrastructure Through a Dedicated 
Network 
 Manage the network infrastructure across network connections that are 
separated from the business use of that network, relying on separate VLANs or, 
preferably, on entirely different physical connectivity for management sessions for 
network devices. 
 ● ●

Page 55 
2.4.3 Ensure admin accounts with different privileges having their 
correct profiles assigned (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Verify that users with access to the Fortinet should only have the minimum privileges 
required for that particular user. 
Rationale: 
In some organizations, there are needs to create different levels of administrative 
accounts. For example, technicians from tier 1 support should not have total access to 
the system as compared with a tier 3 support. 
Audit: 
There are 2 stages to audit. Stage 1: verify the profile. Here is how to verify in the CLI: 
FGT1 # config system accprofile 
FGT1 (accprofile) # edit "tier_1" 
FGT1 (tier_1) # show full 
config system accprofile 
    edit "tier_1" 
        set comments '' 
        set secfabgrp read 
        set ftviewgrp read 
        set authgrp none 
        set sysgrp none 
        set netgrp read 
        set loggrp none 
        set fwgrp custom 
        set vpngrp none 
        set utmgrp none 
        set wifi none 
        set admintimeout-override disable 
        config fwgrp-permission 
            set policy none 
            set address none 
            set service none 
            set schedule none 
        end 
    next 
end 
FGT1 (tier_1) # 
If the following privileges are set to "custom", please also check the sub-privileges of the 
customized ones to make sure that only the right privileges are allowed: fwgrp, sysgrp, 
netgrp, loggrp, utmgrpset. 
In the GUI, go to

