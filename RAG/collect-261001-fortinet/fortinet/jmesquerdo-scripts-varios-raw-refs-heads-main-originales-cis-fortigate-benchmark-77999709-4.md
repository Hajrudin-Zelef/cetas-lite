---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-4
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark", "exploit", "safeguards"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [206, 425]
sha256: a99b0b7943a2958d332d792a7d59cb04f08cff308db28e822ba0144b98b4e64f
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 9 
Impact Statement  
Any security, functionality, or operational consequences that can result from following 
the recommendation. 
Audit Procedure  
Systematic instructions for determining if the target system complies with the 
recommendation  
Remediation Procedure 
Systematic instructions for applying recommendations to the target system to bring it 
into compliance according to the recommendation. 
Default Value 
Default value for the given setting in this recommendation, if known. If not known, either 
not configured or not defined will be applied.  
References 
Additional documentation relative to the recommendation.  
CIS Critical Security Controls® (CIS Controls®) 
The mapping between a recommendation and the CIS Controls is organized by CIS 
Controls version, Safeguard, and Implementation Group (IG). The Benchmark in its 
entirety addresses the CIS Controls safeguards of (v7) “5.1 - Establish Secure 
Configurations” and (v8) '4.1 - Establish and Maintain a Secure Configuration Process” 
so individual recommendations will not be mapped to these safeguards. 
Additional Information  
Supplementary information that does not correspond to any other field but may be 
useful to the user.

Page 10 
Profile Definitions  
The following configuration profiles are defined by this Benchmark: 
• Level 1 
Items in this profile intend to: 
o be practical and prudent; 
o provide a clear security benefit; and 
o not negatively inhibit the utility of the technology beyond acceptable 
means. 
• Level 2 
This profile extends the "Level 1" profile. Items in this profile exhibit one or more 
of the following characteristics: 
o are intended for environments or use cases where security is paramount 
o acts as a defense in depth measure 
o may negatively inhibit the utility or performance of the technology.

Page 11 
 
 
Acknowledgements 
This Benchmark exemplifies the great things a community of users, vendors, and 
subject matter experts can accomplish through consensus collaboration. The CIS 
community thanks the entire consensus team with special recognition to the following 
individuals who contributed greatly to the creation of this guide: 
 
Contributor 
Jayesh Rajan  
Darren Freidel  
Huy Vu Tran 
Mohammed Khalid Babiker Yousif 
Robert Loehmann  
Kent Wade 
Eric Leong  
 
Editor 
Madhukar Saxena

Page 12 
Recommendations 
1 Network Settings 
This section provides best practices related to Network/IP, DNS settings, DHCP server, 
static routing, Policy routing, and dynamic routing.

Page 13 
1.1 Ensure DNS server is configured (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Fortinet uses the Domain Name Service (DNS) to translate host names into IP 
addresses. To enable DNS lookups, you must specify the primary DNS server for your 
system. You can also specify secondary and tertiary DNS servers. When resolving host 
names, the system consults the primary name server. If a failure or time-out occurs, the 
system consults the secondary name server 
Rationale: 
The purpose is to perform the resolution of system hostnames to Internet Protocol (IP) 
addresses. 
Audit: 
In CLI: 
FGT1 # config system dns 
FGT1 (dns) # show 
config system dns 
    set primary <ip_address> 
    set secondary <ip_address> 
    ... 
end 
In the GUI, go to Networks -> DNS. The Fortigate uses either the default FortiGuard 
DNS or customized DNS 
Remediation: 
In this example, we will assign 8.8.8.8 as primary DNS and 8.8.4.4 as secondary DNS. 
In CLI: 
FGT1 # config system dns 
FGT1 (dns) # set primary 8.8.8.8 
FGT1 (dns) # set secondary 8.8.4.4 
FGT1 (dns) # end 
FGT1 # 
In the GUI, go to Networks -> DNS. Click on "Specify" and put in 8.8.8.8 as "Primary 
DNS Server" and 8.8.4.4 as "Secondary DNS Server" 
Default Value: 
Default primary DNS server is 208.91.112.53. Default secondary DNS server is 
208.91.112.52

Page 14 
References: 
1. https://docs.fortinet.com/document/fortigate/6.4.1/administration-
guide/903162/important-dns-cli-commands 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.9 Configure Trusted DNS Servers on Enterprise Assets 
 Configure trusted DNS servers on enterprise assets. Example implementations 
include: configuring assets to use enterprise-controlled DNS servers and/or 
reputable externally accessible DNS servers.  
 ● ● 
v7 
11.1 Maintain Standard Security Configurations for 
Network Devices 
 Maintain standard, documented security configuration standards for all 
authorized network devices. 
 ● ●

Page 15 
1.2 Ensure intra-zone traffic is not always allowed (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
This is to make sure that only specific, authorized traffic are allowed between networks 
in the same zone. 
Rationale: 
This adds an extra layer of protection between different networks 
Audit: 
In this example, we'll verify the zone DMZ. 
In CLI: 
FGT1 # config system zone 
FGT1 (zone) # edit DMZ 
FGT1 (DMZ) # show full 
config system zone 
    edit "DMZ" 
        ... 
        set intrazone deny 
        ... 
    next 
end 
In the GUI, click on Network -> Interfaces, select the zone and click on "Edit". Make 
sure that the option "Block intra-zone traffic" is enabled. 
Remediation: 
In this example, we'll turn of intra-zone traffic in the zone DMZ. 
In CLI: 
FGT1 # config system zone 
FGT1 (zone) # edit DMZ 
FGT1 (DMZ) # set intrazone deny 
FGT1 (DMZ) # end 
FGT1 # 
In the GUI, click on Network -> Interfaces, select the zone and click on "Edit" and turn 
on "Block intra-zone traffic" 
Default Value: 
By default, intra-zone traffic is blocked

Page 16 
References: 
1. https://docs.fortinet.com/document/fortigate/6.2.0/cookbook/116821/zone 
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
2.10 Physically or Logically Segregate High Risk 
Applications 
 Physically or logically segregated systems should be used to isolate and run 
software that is required for business operations but incur higher risk for the 
organization. 
  ●

Page 17 
1.3 Disable all management related services on WAN port 
(Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Enabling any management related services on WAN interface is high risk. Management 
related services such as HTTPS, HTTP, ping, SSH, SNMP, and Radius should be 
disabled on WAN. 
Rationale: 
Management related services should only be enabled on management interface. This is 
part of defending the firewall from attacks and reducing attack surface. For WAN related 
services such as IPSec and SSLVPN, make use of local-in-policy (refer to CIS Section 
2.4) to tighten firewall defenses. 
Impact: 
Enabling management related services on WAN port is convenient but it exposes the 
firewall to unnecessary risks. Vulnerabilities found on vendor devices are commonly 
related to management services and opening access to these allows attackers to exploit 
its vulnerabilities. 
Audit: 
On GUI: 
Go to "Network" > "Interfaces". 
 
Identify WAN interface and validate that HTTPS, HTTP, PING, SSH, SNMP, and 
Radius Accounting is not enabled in "Administrative Access" section. 
On CLI: 
`FGT1 # show system interface` 
Identify WAN interface and validate that "set allowaccess" does not have ping, https, 
http, ssh, snmp or radius-acct configured. 
Remediation: 
On GUI: 
Go to "Network" > "Interfaces". 
Review WAN interface and disable HTTPS, HTTP, ping, SSH, SNMP, and Radius 
services. 
On CLI:

