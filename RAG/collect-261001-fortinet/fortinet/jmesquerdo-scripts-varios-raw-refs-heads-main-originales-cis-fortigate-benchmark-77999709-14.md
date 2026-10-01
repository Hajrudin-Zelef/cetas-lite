---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-14
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["incident", "zero-day"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [2212, 2415]
sha256: 91cda74b0ec47effb3d22da39205c55b1836038596aa9bd20242917e4d2b70d2
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 81 
3.6 Ensure logging is enabled on all firewall policies (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Logging should be enabled for all firewall policies including the default implicit deny 
policy. 
Rationale: 
Firewall policies should log for all traffic (both allow and deny policies). This enables 
SOC or security analyst to do further investigations on security incidents especially on 
threat hunting or incident response activities. Although there are many data sources that 
can provide DNS query logs (AD or EDR), but this option should be enabled out of best 
practice and with assumption that no other data sources is available. 
Impact: 
By default, when creating firewall policies, logging option is not enabled. Also, the 
default implicit deny policy is not logged. This creates data gap in threat hunting or 
incident response activities. 
Audit: 
Go to "Policy & Objects" > "Firewall Policy". 
Validate that logging is enabled on all firewall policies. 
Remediation: 
Review firewall policies and ensure that: 
For allowed policies, "Log Allowed Traffic" is set on "All Sessions" option 
For denied policies, "Log Violation Traffic" is enabled. 
Default Value: 
Disabled 
4 Security Profiles 
This section contains best practices related to FortiGate security features, including: 
• Inspection modes 
• Antivirus 
• Web filter 
• Filtering based on YouTube channel 
• DNS filter

Page 82 
• Application control 
• Intrusion prevention 
• File filter 
• Email filter 
• Data leak prevention 
• VoIP solutions 
• ICAP 
• Web application firewall 
• SSL & SSH Inspection 
• Custom signatures 
• Overrides

Page 83 
4.1 Intrusion Prevention System (IPS) 
Intrusion Prevention System (IPS) Security profiles

Page 84 
4.1.1 Detect Botnet Connections (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Interfaces which are classified as "WAN" and are used by a policy should use an IPS 
sensor which block or monitor outgoing connections to botnet sites. 
Rationale: 
Blocking outgoing connections to known Botnets should be utilized in a Defense In 
Depth network design 
Audit: 
Review all firewall policies that have a "WAN" interface as the destination and ensure 
that an IPS sensor with "Scan Outgoing Connections to Botnet Sites" is set to "Block" 
Remediation: 
Apply an IPS Sensor with "Scan Outgoing Connections to Botnet Sites" set to "Block" 
on all firewall policies with traffic exiting the network to a "WAN" interface.

Page 85 
4.2 Antivirus

Page 86 
4.2.1 Ensure Antivirus Definition Push Updates are Configured 
(Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Ensure FortiGate is configured to accept antivirus definition push updates 
Rationale: 
Ensure that the FortiGate will accept push updates from FortiGuard to ensure the most 
up to date signature databases are present on the device. 
Audit: 
GUI (FortiOS 6): 
Access the FortiGate administrative web access page and go to System > 
FortiGuard under "FortiGuard Updates" validate "Accept push updates"  is 
enabled. 
GUI (FortiOS 7): 
Access the FortiGate administrative web access page and go to System > 
FortiGuard under "FortiGuard Updates" ensure that the "Scheduled updates" is 
set to "Automatic". 
CLI (FortiOS 6): 
config system autoupdate push-update 
get (Validate status is enable) 
CLI (FortiOS 7): 
config system autoupdate schedule 
show (Validate that there are no output, meaning it is already set as 
"automatic" 
Remediation: 
GUI (FortiOS 6): 
Access the FortiGate administrative web access page and go to System > 
FortiGuard under 'FortiGuard Updates" enable "Accept push updates".   
GUI (FortiOS 7): 
Access the FortiGate administrative web access page and go to System > 
FortiGuard under "FortiGuard Updates" ensure that the "Scheduled updates" is 
set to "Automatic".  
CLI (FortiOS 6):

Page 87 
config system autoupdate 
set status enable 
end 
CLI (FortiOS 7): 
config system autoupdate schedule 
set status enable 
set frequency automatic 
end 
Default Value: 
Disable (on FortiOS 6) 
Enabled and set to automatic (on FortiOS 7) 
References: 
1. https://docs.fortinet.com/document/fortigate/7.0.10/administration-guide/547335

Page 88 
4.2.2 Apply Antivirus Security Profile to Policies (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Ensuring that traffic traversing between networks on the FortiGate have an Antivirus 
Security profile inspecting it. 
Rationale: 
Traffic moving between "interfaces" on the FortiGate should have firewall policies 
applied with an antivirus security profile applied. 
Audit: 
Review all firewall policies and ensure that traffic has an antivirus security profile 
assigned for inspection 
Remediation: 
Review firewall policies and apply an appropriate antivirus security profile to policies

Page 89 
4.2.3 Enable Outbreak Prevention Database (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Ensure FortiGate AV inspection uses outbreak prevention database as an added layer 
of protection on top of antivirus' signature-based detection. 
Rationale: 
Antivirus mainly uses signature for malware blocking. By enabling "FortiGuard outbreak 
prevention database", FortiGate can leverage on 3rd party malware hash signatures 
curated by the FortiGuard as an additional protection layer. 
The hash signatures are obtained from FortiGuard's Global Threat Intelligence 
database. The antivirus database queries FortiGuard with the hash of a scanned file. If 
FortiGuard returns a match, the scanned file is deemed to be malicious. 
Audit: 
GUI: 
Go to "Security Profiles" > "AntiVirus" > select AV profile 
Validate that "Use FortiGuard outbreak prevention database" is enabled. 
CLI: 
FGT1 # config antivirus profile 
 
FGT1 (profile) # show 
Validate that for each traffic protocol, "set outbreak-prevention block" is configured. 
Remediation: 
Review Antivirus Security Profiles and validate that "Use FortiGuard outbreak 
prevention database" is enabled. 
Default Value: 
Disabled 
References: 
1. https://docs.fortinet.com/document/fortigate/7.0.9/administration-
guide/889364/fortiguard-outbreak-prevention

Page 90 
4.2.4 Enable AI /heuristic based malware detection (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
AI /heuristic based detection should be enabled. 
Rationale: 
The AV Engine AI malware detection model integrates into regular AV scanning to help 
detect potentially malicious Windows Portable Executables (PEs) in order to mitigate 
zero-day attacks. It is an additional layer of protection on top of traditional antivirus 
protection. 
In version 6.x, it is named "Heuristic detection". On version 7.x, Fortinet has renamed 
this to AI based detection. 
Audit: 
Configuration and verification can be only done on CLI. 
On FortiOS 6.4.x 
FGT1 # show antivirus heuristic 
Validate that it is in "block" mode. 
On FortiOS 7.x: 
FGT1 # show antivirus settings | grep machine-learning-detection 
Validate that it is enabled. 
Remediation: 
FGT1 # config antivirus settings 
 
FGT1 (settings) # set machine-learning-detection enable 
Default Value: 
Disabled (for version 6.4.x) 
Enabled (for version 7.x) 
References: 
1. https://docs.fortinet.com/document/fortigate/6.4.11/cli-reference/517620/config-
antivirus-heuristic 
2. https://docs.fortinet.com/document/fortigate/7.0.0/new-features/773410/ai-based-
malware-detection

