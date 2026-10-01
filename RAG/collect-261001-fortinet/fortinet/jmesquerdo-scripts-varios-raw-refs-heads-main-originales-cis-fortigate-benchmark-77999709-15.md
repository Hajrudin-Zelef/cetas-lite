---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-15
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [2416, 2613]
sha256: a59e0713cc64dcda801f29a888bfdea84e9a1410bb9f5eca60b71196d38d9537
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 91 
4.2.5 Enable grayware detection on antivirus (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Grayware detection should be enabled. 
Rationale: 
Usage of grayware is generally not allowed in strict company policies and some 
graywares can be used for malicious intent. If the file passes the virus scan, it can be 
checked for grayware. Grayware signatures are kept up to date in the same manner as 
the antivirus definitions. 
Audit: 
CLI: 
FGT1 # show antivirus settings | grep grayware 
Validate that grayware detection is enabled. 
Remediation: 
FGT1 # config antivirus settings 
 
FGT1 (settings) # set grayware enable 
Default Value: 
Enabled 
References: 
1. https://community.fortinet.com/t5/FortiGate/Technical-Tip-Configuration-options-
about-antivirus/ta-p/191939

Page 92 
4.3 DNS Filter

Page 93 
4.3.1 Enable Botnet C&C Domain Blocking DNS Filter 
(Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Enable Botnet C&C domain blocking to block botnet access at the DNS name resolving 
stage 
Rationale: 
Blocking botnet website access at the DNS resolution stage provides an additional layer 
of defense. 
Audit: 
GUI: 
Review DNS filters under Security Profiles > DNS Filter and ensure that 
"redirect botnet C&C requests to Block portal" is enabled and that policies 
allowing DNS traffic have a DNS Filter Security profile applied 
Remediation: 
Review DNS Filter Security Profiles and validate that "Redirect botnet C&C requests to 
Block Portal" is enabled and that firewall policies that have DNS traffic have a DNS 
Filter security profile applied with that option enabled

Page 94 
4.3.2 Ensure DNS Filter logs all DNS queries and responses 
(Manual) 
Profile Applicability: 
•  Level 1 
Description: 
DNS filter should log all DNS queries and responses. 
Rationale: 
DNS filter should log all DNS queries and responses (whether if the DNS category is 
blocked, monitored, or allowed). This enables SOC or security analyst to do further 
investigations on security incidents especially on threat hunting or incident response 
activities. Although there are many data sources that can provide DNS query logs (AD 
or EDR), but this option should be enabled out of best practice and with assumption that 
no other data sources is available. 
Impact: 
By default, allowed DNS is not logged. This creates data gap in threat hunting or 
incident response activities. 
Audit: 
GUI: 
Go to "Security Profiles" > "DNS Filter" > select DNS Filter profile 
Validate that "Log all DNS queries and responses" is enabled. 
CLI: 
FGT1 # config dnsfilter profile 
 
FGT1 (profile) # show 
Validate that "set log-all-domain enable" is configured on DNS Filter profile. 
Remediation: 
Review DNS Filter Security Profiles and validate that "Log all DNS queries and 
responses" is enabled. 
Default Value: 
Disabled

Page 95 
References: 
1. https://community.fortinet.com/t5/FortiGate/Technical-Tip-FortiGate-Static-DNS-
filter-behavior-in-logging/ta-p/223110

Page 96 
4.4 Application Control 
Application Control Security profiles

Page 97 
4.4.1 Block high risk categories on Application Control (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Ensure FortiGate Application Control blocks high risk application to reduce attack 
surface. 
Rationale: 
High risk applications such as those in "P2P" and "Proxy" are known for spreading 
malwares. Other than that, some of these traffic is encrypted and therefore is able to 
bypass network security inspection (for those without decryption implemented). Blocking 
these applications from running eliminates this risk. 
If any application that falls under "P2P" and "Proxy" requires to be allowed based on 
organization's policy, that specific application needs to be under "Monitor" mode in the 
"Application and Filter Override" configuration. 
Audit: 
GUI: 
Go to "Security Profiles" > "Application Control" > select App Control 
profile 
Validate that "P2P" and "Proxy" category is blocked. 
Remediation: 
Review Application Control Security Profiles and validate that "P2P" and "Proxy" 
category is blocked. 
Default Value: 
Disabled on default profile

Page 98 
4.4.2 Block applications running on non-default ports (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Ensure FortiGate Application Control blocks applications running on non-default ports. 
Rationale: 
Running application on non-default ports is not directly a threat, but can be an indication 
of something unexpected. For example, HTTPS runs on port 443. Potentially, if attacker 
starts a rogue HTTPS server on port 10443, it could be used for data exfiltration. 
Audit: 
GUI: 
Go to "Security Profiles" > "Application Control" > select App Control 
profile 
Validate that "Block applications detected on non-default ports" option is enabled. 
Remediation: 
GUI: 
Go to "Security Profiles" > "Application Control" > select App Control 
profile 
 
Enable "Block applications detected on non-default ports" option 
CLI: 
FGT1 # config application list 
 
FGT1 (list) # edit <profile name> 
 
FGT1 (<profile name>) # set enforce-default-app-port enable 
Default Value: 
Disabled 
References: 
1. https://attack.mitre.org/techniques/T1571/

Page 99 
4.4.3 Ensure all Application Control related traffic are logged 
(Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Ensure no category is set to "Allow" on FortiGate Application Control. 
Rationale: 
Any category that is set as "Allow" on Application Control will not be logged. This 
creates visibility gap on security investigation. This includes "Unknown Applications" 
category. 
Impact: 
Visibility gap, affects incident forensics and response. 
Audit: 
GUI: 
Go to "Security Profiles" > "Application Control" > select App Control 
profile 
Validate that no "Allow" action is set on any categories. 
Remediation: 
Review Application Control Security Profiles and validate that no "Allow" action is set on 
any categories. 
Default Value: 
"Unknown Applications category is set as "Allow" 
5 Security Fabric 
This Section provides best practice related to configuring Fortinet Security Fabric.

Page 100 
5.1 Automation

Page 101 
5.1.1 Enable Compromised Host Quarantine (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Default automation trigger configuration for when a high severity compromised host is 
detected. 
Rationale: 
By enabling this feature you protect your environment against compromised hosts. 
Default automation stitch to quarantine a high severity compromised host on FortiAPs, 
FortiSwitches, and FortiClient EMS. 
Audit: 
GUI 
Security Fabric>Automation> 
Verify Compromised Host Quarantine is enabled. 
Remediation: 
GUI 
Security Fabric>Automation 
Edit and change Disabled to Enabled 
CLI

