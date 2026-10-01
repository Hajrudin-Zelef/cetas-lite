---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-16
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [2614, 2854]
sha256: 5a9dccec233d9b4c9cff0a11e0e68d12221fc6e66cba51f73f0c9d4351b5ae1a
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 102 
config system automation-action 
    edit "Quarantine on FortiSwitch + FortiAP" 
        set description "Default automation action configuration for 
quarantining a MAC address on FortiSwitches and FortiAPs." 
        set action-type quarantine 
    next 
    edit "Quarantine FortiClient EMS Endpoint" 
        set description "Default automation action configuration for 
quarantining a FortiClient EMS endpoint device." 
        set action-type quarantine-forticlient 
    next 
end 
config system automation-trigger 
    edit "Compromised Host - High" 
        set description "Default automation trigger configuration for when a 
high severity compromised host is detected." 
    next 
end 
config system automation-stitch 
    edit "Compromised Host Quarantine" 
        set description "Default automation stitch to quarantine a high 
severity compromised host on FortiAPs, FortiSwitches, and FortiClient EMS." 
        set status disable 
        set trigger "Compromised Host - High" 
        config actions 
            edit 1 
                set action "Quarantine on FortiSwitch + FortiAP" 
            next 
            edit 2 
                set action "Quarantine FortiClient EMS Endpoint" 
            next 
        end 
    next 
end 
Default Value: 
Not enabled 
5.2 Fabric Connectors 
Security Fabric Connector Configuration

Page 103 
5.2.1 Configure Root FortiGate for Security Fabric 
Configuring and identifying the root FortiGate within the Security Fabric

Page 104 
5.2.1.1 Ensure Security Fabric is Configured (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Ensure Root FortiGate is configured as security fabric root 
Rationale: 
Without a root FortiGate configured the security fabric is not functional and can not be 
leveraged 
Impact: 
Without Security Fabric enabled visibility and management of traffic throughout an 
organization is decreased and individual FortiGate management becomes more 
intensive 
Audit: 
Review through the GUI: 
To Validate root FortiGate status go to "Security Fabric" -> Fabric 
Connectors and then select "Security Fabric Setup"  
Validate that the root FortiGate has status set to enabled and the Security 
Fabric Role set to "Serve as Fabric Root" 
Ensure that FortiAnalyzer settings are correct and that there is a defined 
Fabric name as well as interfaces selected that will "Allow other Security 
Fabric Devices to Join". 
Remediation: 
Remediation through the GUI: 
To configure root FortiGate status go to "Security Fabric" -> Fabric 
Connectors and then select "Security Fabric Setup"  
On the root FortiGate set the status to enabled and the Security Fabric Role 
to "Serve as Fabric Root" 
 
Configure FortiAnalyzer settings when prompted and define a Fabric name as 
well as interfaces that will "Allow other Security Fabric Devices to Join". 
Default Value: 
Disabled 
6 VPN

Page 105 
6.1 SSL VPN 
SSL VPN Best Practices

Page 106 
6.1.1 Apply a Trusted Signed Certificate for VPN Portal (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Apply a signed certificate from a trusted Certificate Authority (CA) to the SSL VPN portal 
to allow users to connect securely with confidence 
Rationale: 
Having an unsigned or self signed certificate leaves connections open to man-in-the-
middle attacks and could allow users to connect to untrusted servers 
Audit: 
GUI: 
Access the FortiGate administrative web access page and go to VPN > SSL-VPN 
Settings and assign a signed certificate in the dropdown for "Server 
Certificate" 
Remediation: 
Import a signed certificate from a trusted CA through the GUI 
System > Certificates > Import and then assign the certificate to the SSL VPN 
portal by going to VPN > SSL-VPN Settings and selecting the proper 
certificate in the dropdown for "Server Certifcate" 
Default Value: 
Self Signed Factory installed certificate

Page 107 
6.1.2 Enable Limited TLS Versions for SSL VPN (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Enable and disable TLS versions and Cipher suites for more granular control of SSL 
VPN connections and enforcing more secure connections. 
Rationale: 
Limiting TLS versions to more secure versions as well as enforcing stronger ciphers 
increases the security of the SSL VPN connections 
Audit: 
CLI: 
Config vpn ssl settings 
get (Validate ssl-max-prot-ver and ssl-min-proto-ver as well as algorithm and 
banned-cipher) 
Remediation: 
CLI: 
config vpn ssl settings 
set ssl-max-prot-ver *** {Configure max TLS Version supported} 
set ssl-min-proto ver *** {set minimum support TLS version} 
set banned-cipher *** {add cipher suite to banned list and prevent it from 
being used} 
set algorithm high {use high algorithms} 
Default Value: 
ssl-max-proto-ver : tls1-3 ssl-min-proto-ver : tls1-2 banned-cipher : algorithm : high

Page 108 
7 Users and Authentication 
This section provides best practice related to Users and devices including: 
• Endpoint control and compliance 
• Users and user Groups Definition 
• Guest Management 
• LDAP, RADIUS, and TACACS+ Servers 
• Authentication Settings 
• FortiTokens 
• PKI 
• Configuring the maximum login attempts and lockout period

Page 109 
7.1 Configuring the maximum login attempts and lockout period 
(Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Configure maximum user log in attempts and lockout period 
Rationale: 
Failed user log in attempts can indicate an attempt to gain access to the network. 
Limiting the number of attempts before the account is locked for a determined amount 
of time helps slow down brute force attempts and impedes malicious attempts to access 
user accounts. 
Audit: 
CLI: 
config user setting 
get (Validate auth-lockout-threshold *  the number of attempts before locking 
out the account and auth-lockout duration * the duration of the lockout once 
the failed attempts is met) 
Remediation: 
CLI: 
config user setting 
 
set auth-lockout-threshold 5 
 
end 
 
config user setting 
 
set auth-lockout-duration 300 
 
end 
Default Value: 
auth-lockout-threshold: 3 auth-lockout-duration: 0 
8 Logs and Reports 
This section provides best practices related to logging and reporting in FortiGate.

Page 110 
8.1 Enable Logging 
How to enable logging on the FortiGate device.

Page 111 
8.1.1 Enable Event Logging (Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Enabling event logging to allow for log generation and review. 
Rationale: 
Enabling event logging generates logs that can be stored for later review or auditing or 
can be ingested by another system (SIEM, Analyzer) for monitoring and response 
Audit: 
CLI: 
config log eventfilter 
get 
(validate event is enabled) 
end 
Remediation: 
Access the FortiGate administrative web access page and go to 
Log & Report > Log Settings enable Event Logging. 
CLI: 
config log eventfilter 
set event enable 
end

Page 112 
8.2 Encrypt Logs Sent to FortiAnalyzer / FortiManager 
Ensure that logs sent to FortiAnalyzer or FortiManager are encrypted during 
transmission.

Page 113 
8.2.1 Encrypt Log Transmission to FortiAnalyzer / FortiManager 
(Automated) 
Profile Applicability: 
•  Level 2 
Description: 
Enable encryption for logs that are sent to FortiAnalyzer or FortiManager 
Rationale: 
Provides encryption for logs that are sent to FortiAnalyzer or FortiManager to prevent 
logs being collected and viewed as they traverse the network. 
Audit: 
CLI: 
config log fortianalyzer setting 
get (validate enc-algorithm is set to high) 
end 
Remediation: 
GUI: 
Access the FortiGate administrative web access page and go to Log & Report > 
Log Settings and when configuring Remote logging to 
FortiAnalyzer/FortiManager select "Encrypt log transmission" 
CLI: 
config log fortianalyzer setting 
set enc-algorithm high 
end

Page 114 
8.3 Centralized Logging and Reporting 
Logging and Reporting should be done to a Centralized device

