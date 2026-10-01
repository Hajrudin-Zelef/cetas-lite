---
id: collect-261001-rattrapage/rattrapage/githubfoam-acfbb15b7918cb492fafd3e18e7edc38-4136301d-3
title: "githubfoam-acfbb15b7918cb492fafd3e18e7edc38-4136301d"
domain: rattrapage
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/githubfoam-acfbb15b7918cb492fafd3e18e7edc38-4136301d.md
source_anchor: ""
source_lines: [126, 200]
sha256: 2e0d2b93f595bc14c9985465a6da3c5bf05c05f311e012206eb2fd1ce01ec1d1
---

# githubfoam-acfbb15b7918cb492fafd3e18e7edc38-4136301d

|  | Logging to FortiAnalyzer stores the logs and provides log analysis | 
|  | If a security fabric is established, you can create rules to trigger actions based on the logs. | 
|  | For example, sending an email if the FortiGate configuration is changed, or running a CLI script if a host is compromised. | 
|  | If you are using a standalone logging server, integrating an analyzer application or server allows you to parse the raw logs into meaningful data. | 
|  | #========================================================================================================================================== | 
|  | #Firewall Hardening | 
|  | # https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/862226/policies#LocalInPolicies | 
|  | The principle of least privilege (PoLP) is an information security concept which maintains that | 
|  | a user or entity should only have access to the specific data, resources and applications needed to complete a required task | 
|  | Use local-in policies | 
|  | Note that extra care should be taken when configuring a local-in policy, as an incorrect configuration could inadvertently deny traffic for SSL VPN, dynamic routing protocols, HA, and other FortiGate features. | 
|  | Policies that allow traffic should apply to a specific interface, and not the any interface. | 
|  | Security policies are evaluated in order. When traffic matches a policy, further policies are not processed. | 
|  | Put the most specific policies at the top of the list, and follow the least privilege access principle | 
|  | Policies | 
|  | Put the most specific, or narrow, policies at the top of the policy list. | 
|  | Do not use the all or any objects in a policy, except when routing to the internet. | 
|  | Do not override the implicit deny policy. | 
|  | Use users in policies. This makes the policy more specific and reduces the chances of unintended traffic matching. | 
|  | Virtual IPs | 
|  | Policies that include VIPs, or that have match-vip enabled, have priority over other policies. | 
|  | Interface aliases | 
|  | It might not be possible to use the same interface on each FortiGate for the same function. | 
|  | Add aliases to the interfaces so that policies are easier to understand. For example, a policy that controls traffic | 
|  | between you network and your phones switch is clearer if it shows LAN to Phones, instead of port4 to port2. | 
|  | Network > Interfaces > mgmt > Alias | 
|  | #========================================================================================================================================== | 
|  | #Firewall Hardening | 
|  | #https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/555436/hardening#PhysicalSec | 
|  | Physical access to the FortiGate can allow it to be bypassed, or other firmware could be loaded after a manual reboot. | 
|  | If the FortiGate cannot be physical secured: | 
|  | Disable USB firmware and configuration installation: | 
|  | config system auto-install | 
|  | set auto-install-config disable | 
|  | set auto-install-image disable | 
|  | end | 
|  | Enable port security (802.1x) to prevent unauthorized devices from forwarding traffic. | 
|  | Optionally, disable the maintainer account. Note that doing this will make you unable to recover administrator access using a console connection is all of the administrator credentials are lost. | 
|  | #========================================================================================================================================== | 
|  | #Firewall Hardening | 
|  | # Vulnerability - monitoring PSIRT https://www.fortiguard.com/psirt?product=FortiOS | 
|  | #Firmware | 
|  | Keep the FortiOS firmware up to date. The latest patch release has the most fixed bugs and vulnerabilities, and should be the most stable. | 
|  | Read the release notes. The known issues may include issues that affect your business. | 
|  | #Encrypted protocols | 
|  | Use encrypted protocols whenever possible, for example, | 
|  | SNMPv3 instead of SNMP, | 
|  | SSH instead of telnet, | 
|  | OSPF MD5 authentication, | 
|  | SCP instead of FTP or TFTP, | 
|  | NTP authentication, | 
|  | and encrypted logging instead of TCP. | 
|  | #===================================================================== | 
|  | #upgrade the firmware on an HA cluster in the same way as on a standalone FortiGate | 
|  | #Interrupted upgrade is disabled by default | 
|  | #An interrupted upgrade upgrades all cluster members at the same time. T | 
|  | config system ha | 
|  | set uninterruptible-upgrade disable | 
|  | end | 
|  | #===================================================================== | 
|  | Firewall Hardening | 
|  | If a security fabric is established, you can create rules to trigger actions based on the logs. | 
|  | For example, sending an email if the FortiGate configuration is changed, or running a CLI script if a host is compromised. | 
|  | https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/691328/logging-and-reporting | 
|  | #===================================================================== | 
|  | Alert by email notification? | 
|  | FortiGuard databases | 
|  | Ensure that FortiGuard databases, such as AS, IPS, and AV, are updated punctually. Optionally, send an alert if they are out of date. | 
|  | https://docs.fortinet.com/document/fortigate/7.2.0/best-practices/555436/hardening#FortiGuardDatabase | 
|  | #===================================================================== | 
  
    Sign up for free
    to join this conversation on GitHub.
    Already have an account?
    Sign in to comment
