---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-7
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [790, 1003]
sha256: 52d2bfb54e40755110bf9c767610a35b26a2fb959d7675c738911ce4b242b3f7
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 30 
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

Page 31 
2.1.6 Ensure the latest firmware is installed (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Check against Fortinet website to make sure that the latest stable firmware is installed. 
Rationale: 
Fortinet periodically updates the FortiGate firmware to include new features and resolve 
important issues. After you have registered your FortiGate unit, firmware updates can 
be downloaded from the Fortinet Customer Service & Support website. 
It is important to constantly keep the firmware up-to-date to prevent any new well-known 
exploitation. 
Audit: 
First, check for the latest firmware version available by going to 
https://docs.fortinet.com/upgrade-tool, select your product from the Current Product 
drop-down menu then select the upgrade to FortiOS Version which will gives you the 
latest version available. 
Second, verify the current firmware on your system. 
In the CLI: 
FGT1 # get system status 
... 
Version: Fortigate-100D v6.2.7,build1190,201216 (GA) 
... 
FGT1 # 
In the GUI: 
go to Dashboard -> Status -> System information and check for Firmware. 
At the same time, go to https://www.fortiguard.com/psirt?product=FortiOS and check for 
vulnerabilities that your existing version might have.

Page 32 
Remediation: 
First, determine the upgrade path recommended by Fortinet. If you have not upgraded 
the system for a long time, it is not recommended to upgrade straight to the latest 
version as the configuration could be lost. Fortinet provides a tool to recommend an 
upgrade path for all of its products. 
Go to https://docs.fortinet.com/upgrade-tool. Choose your product from the "Current 
Product" drop-down menu, the "current FortiOS version", and the latest firmware 
version available for that model from "Upgrade to FortiOS Version". Click "Go". Write 
down the path and then click on "Download" to download all the necessary versions. 
The second step is to download the required FortiOS firmware/s. Go to 
https://support.fortinet.com and login. Go to Support -> Firmware Download. Once 
there, select the product and click on "Upgrade Path". Choose the specific model of the 
hardware, the current firmware version and the latest firmware version available for that 
model. Click "Go". Write down the path and then click on "Download" to download all 
the necessary versions. 
The last step is to install the new firmwares in the order provided by the "Upgrade tool". 
It is recommended to use GUI to perform this task as it would be much easier. 
In the GUI, click on 
System -> Firmware, then click on "Browse" to select the next firmware file. 
Then click on "Upgrade". You might have to perform this step multiple times 
if you follow the upgrade path. 
Default Value: 
There is no default firmware. The hardware comes with the latest firmware at the time it 
was manufactured. 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=10948 
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
2.2 Ensure Software is Supported by Vendor 
 Ensure that only software applications or operating systems currently supported 
by the software's vendor are added to the organization's authorized software 
inventory. Unsupported software should be tagged as unsupported in the inventory 
system. 
● ● ●

Page 33 
Controls 
Version Control IG 1 IG 2 IG 3 
v7 
8.2 Ensure Anti-Malware Software and Signatures are 
Updated 
 Ensure that the organization's anti-malware software updates its scanning 
engine and signature database on a regular basis. 
● ● ● 
v7 
11.4 Install the Latest Stable Version of Any Security-
related Updates on All Network Devices 
 Install the latest stable version of any security-related updates on all network 
devices. 
● ● ●

Page 34 
2.1.7 Disable USB Firmware and configuration installation 
(Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Disable USB port auto install feature for config and firmware 
Rationale: 
Disabling USB port for auto install prevents a USB from being connected with a 
manipulated configuration or incorrect firmware from being connected and loaded 
automatically. 
Audit: 
CLI: 
config system auto-install 
get (verify that set auto-install-config and set auto-install-image are 
disabled) 
Remediation: 
CLI: 
config system auto-install 
    set auto-install-config disable 
    set auto-install-image disable 
end 
Default Value: 
config system auto-install set auto-install-config enable set auto-install-image enable 
end

Page 35 
2.1.8 Disable static keys for TLS (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Disable support for static keys on TLS sessions terminating on the FortiGate 
Rationale: 
Prevent TLS sessions terminating on the FortiGate from using static SSL keys 
Audit: 
CLI: 
config system global 
get (Validate that ssl-static-key-ciphers disable is set) 
Remediation: 
CLI: 
config system global 
 
set ssl-static-key-ciphers disable 
 
end 
Default Value: 
set ssl-static-key-ciphers enable

Page 36 
2.1.9 Enable Global Strong Encryption (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Enable FortiOS to only use strong encryption and allow only strong ciphers for 
communication 
Rationale: 
Audit: 
CLI: 
config system global 
get (validate strong-crypto is enabled) 
Remediation: 
CLI: 
config system global 
 
set strong-crypto enable 
 
end 
Default Value: 
strong-crypto : enable

Page 37 
2.2 Password Policy 
This Section contains criteria for local passwords such as complexity and restrictions. 
The best practice is to use named accounts, and if possible a back-end authentication 
solution such as Active Directory or (best case) a two-factor authentication solution. 
However, local credentials will always exist, if only to account for the failure of a back-
end authentication solution.

Page 38 
2.2.1 Ensure 'Password Policy' is enabled (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
It is important to use secure and complex passwords for preventing unauthorized 
access to the FortiGate device. 
Rationale: 
Attackers can use Brute force password software to launch more than just dictionary 
attacks. such Attacks can discover common passwords where a letter is replaced by a 
number or symbol. 
Impact: 
Weak passwords can be easily discovered by hackers which leads to unauthorized 
access to FortiGate and depends on the access privilege of the compromised account 
the attacker may modify the settings. 
Audit: 
currently implemented password policy can be shown from GUI or CLI 
From CLI, type 
get system password-policy  
From GUI, 
Or from GUI as follows: 
1) log in to FortiGate with a user with at least read-only privileges  
2) Go to 'System' -> 'Settings' 
3) find and check the status of the 'password Policy' Section 
Remediation: 
can be modified from CLI or GUI 
From CLI, do the following:

