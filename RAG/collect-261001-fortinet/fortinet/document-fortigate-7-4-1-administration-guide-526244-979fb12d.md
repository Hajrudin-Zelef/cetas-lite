---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-526244-979fb12d
title: "document-fortigate-7-4-1-administration-guide-526244-979fb12d"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-526244-979fb12d.md
source_anchor: ""
source_lines: [1, 31]
sha256: a48dcc7eb4c5784d70ef0af36d7a9d0618054ac9f4dc8a7f912a7937ed42a409
---

# document-fortigate-7-4-1-administration-guide-526244-979fb12d

Interface policies
Interface policies
Interface policies are implemented before the security policies and are only flow-based. They are configured in the CLI.
This feature allows you to attach a set of IPS policies with the interface instead of the forwarding path, so packets can be delivered to IPS before entering the firewall. This feature is used for following IPS deployments:
- 
                                                    One-Arm: By defining interface policies with IPS and DoS anomaly checks and enabling sniff-mode on the interface, the interface can be used for one-arm IDS.
- 
                                                    IPv6 IPS: IPS inspection can be enabled through interface IPv6 policy.
- 
                                                    Scan traffic that is destined to the FortiGate.
- 
                                                    Scan and log traffic that are silently dropped or flooded by Firewall or Multicast traffic.
IPS sensors can be assigned to an interface policy. Both incoming and outgoing packets are inspected by IPS sensor (signature).
To configure an interface policy:
config firewall interface-policy
    edit 1
        set status enable
        set comments 'test interface policy #1'
        set logtraffic utm
        set interface "port2"
        set srcaddr all
        set dstaddr all
        set service "ALL"
        set application-list-status disable
        set ips-sensor-status disable
        set dsri disable
        set av-profile-status enable
        set av-profile default
        set webfilter-profile-status disable
    next
end
