---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-8
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [1004, 1196]
sha256: 5bc27e5809c6e2950be3ff56f611a7e6da07f93fa503bd4b76f4673b2ff11ea9
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 39 
config system password-policy 
 set status enable 
 set apply-to admin-password ipsec-preshared-key 
 set minimum-length 8 
 set min-lower-case-letter 1 
 set min-upper-case-letter 1 
 set min-non-alphanumeric 1 
 set min-number 1 
 set expire-status enable 
 set expire-day 90 
 set reuse-password disable 
end 
or From GUI, do the following 
1) log in to FortiGate as Super Admin 
2) Go to 'System' -> 'Settings' 
3) find the 'password Policy' Section 
4) Default 'Password scope' is 'Off', change it to 'Both' 
5) set 'Minimum length' to '8' 
6) Enable 'Character requirements' 
7) set minimum '1' in the filed of 'Upper Case', 'Lower Case', 'Numbers (0-
9)' and 'Special' 
8) Disable 'Allow password reuse' 
9) Enable 'Password expiration' and set it to 90 
Default Value: 
By Default, Password Policy is disabled, can be checked from CLI as follows: 
config system password-policy 
    set status disable 
end 
Or from GUI as follows: 
1) log in to FortiGate as Super Admin 
2) Go to 'System '-> 'Settings' 
3) find the 'password Policy' Section 
4) Default 'Password scope' is 'Off' 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD31021 
2. https://docs.fortinet.com/document/fortigate/7.0.0/cli-reference/11620/config-
system-password-policy 
3. https://docs.fortinet.com/document/fortigate/7.0.0/administration-
guide/364729/password-policy 
Additional Information: 
Consider the following to ensure better security: 
• Do not use passwords that are obvious, such as the company name, 
administrator names, or other obvious words or phrases.

Page 40 
• Use numbers in place of letters, for example: passw0rd. 
• Administrator passwords can be up to 64 characters. 
• Include a mixture of numbers, symbols, and upper and lower case letters. 
• Use multiple words together, or possibly even a sentence, for example: 
correcthorsebatterystaple. 
• Use a password generator. 
• Change the password regularly and always make the new password unique and 
not a variation of the existing password. for example, do not change from 
password to password1. 
• Make note of the password and store it in a safe place away from the 
management computer, in case you forget it; or ensure at least two people know 
the password in the event one person becomes unavailable. Alternatively, have 
two different admin logins. 
FortiGate allows you to create a password policy for administrators and IPsec pre-
shared keys. With this policy, you can enforce regular changes and specific criteria for a 
password policy, including: 
• The minimum length, between 8 and 64 characters. 
• If the password must contain uppercase (A, B, C) and/or lowercase (a, b, c) 
characters. 
• If the password must contain numbers (1, 2, 3). 
• If the password must contain special or non-alphanumeric characters: !, @, #, $, 
%, ^, &, *, (, and ) 
• Where the password applies (admin or IPsec or both). 
• The duration of the password before a new one must be specified. 
• The minimum number of unique characters that a new password must include. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
5.2 Use Unique Passwords 
 Use unique passwords for all enterprise assets. Best practice implementation 
includes, at a minimum, an 8-character password for accounts using MFA and a 
14-character password for accounts not using MFA.  
● ● ● 
v7 
4.4 Use Unique Passwords 
 Where multi-factor authentication is not supported (such as local administrator, 
root, or service accounts), accounts will use passwords that are unique to that 
system. 
 ● ●

Page 41 
2.2.2 Ensure administrator password retries and lockout time are 
configured (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Failed login attempts can indicate malicious attempts to gain access to your network. To 
prevent this security risk, FortiGate is preconfigured to limit the number of failed 
administrator login attempts. After the maximum number of failed login attempts is 
reached, access to the account is blocked for the configured lockout period. 
Rationale: 
When you login and fail to enter the correct password you could be a valid user, or a 
hacker attempting to gain access. For this reason, best practices dictate to limit the 
number of failed attempts to login before a lockout period where you cannot login for a 
certain period of time. lockout period will minimize the hacker attempts to gain access to 
firewall. 
Impact: 
Attackers will keep attempting to access the device through brute force attacks without 
any interruption which may lead to a successful login. 
Audit: 
To check the lockout options, from CLI: 
get system global 
from the output, check the value of the below fields 
admin-lockout-threshold and admin-lockout-duration 
Remediation: 
To configure the lockout options, from CLI: 
config system global 
 set admin-lockout-threshold 3 
 set admin-lockout-duration 60 
end 
Default Value: 
By default, the number of password retry attempts is set to three, allowing the 
administrator a maximum of three attempts at logging in to their account before they are 
locked out for a set amount of time (by default, 60 seconds). 
To configure the lockout options, from CLI:

Page 42 
config system global 
 set admin-lockout-threshold 3 
 set admin-lockout-duration 60 
end 
References: 
1. https://docs.fortinet.com/document/fortigate/6.2.0/cookbook/631730/setting-the-
administrator-password-retries-and-lockout-time 
Additional Information: 
The number of attempts and the default wait time before the administrator can try to 
enter a password again can be configured using the CLI. 
A maximum of ten retry attempts can be configured, and the lockout period can be 1 to 
2147483647 seconds (over 68 years). 
The higher the retry attempts, the higher the risk that someone might be able to guess 
the password. 
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

Page 43 
2.3 SNMP

Page 44 
2.3.1 Ensure SNMP agent is disabled (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
The Simple Network Management Protocol (SNMP) server is used to listen for SNMP 
commands from an SNMP management system, execute the commands or collect the 
information and then send results back to the requesting system. 
Rationale: 
The SNMP server can communicate using SNMP v1, which transmits data in the clear 
and does not require authentication to execute commands. Unless absolutely 
necessary, it is recommended that the SNMP service not be used. If SNMP is required 
the server should be configured to use only SNMPv3. 
Impact: 
SNMP servers will not be able to query the Fortigate devices that have SNMP agents 
disabled. 
Audit: 
on CLI, run the following commands to check whether SNMP agent is disabled. 
FGT1 # config system snmp sysinfo 
FGT1 (sysinfo) # show full 
config system snmp sysinfo 
    set status disable 
    ...  
end 
On the GUI, select System -> SNMP, make sure that SNMP Agent is disabled. 
Remediation: 
On the CLI, run the following command to disable the agent 
FGT1 # config system snmp sysinfo 
FGT1 (sysinfo) # set status disable 
FGT1 (sysinfo) # end 
On the GUI, select System -> SNMP, disable SNMP agent 
Default Value: 
SNMP agent is disabled by default.

