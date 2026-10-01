---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-7-0-3-administration-guide-667795-400867a1
title: "index-php-document-fortigate-7-0-3-administration-guide-667795-400867a1"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-7-0-3-administration-guide-667795-400867a1.md
source_anchor: ""
source_lines: [1, 87]
sha256: 5f74c60327d26395535474f22b40fd3719674b0cab4f7828f1bda95236f47464
---

# index-php-document-fortigate-7-0-3-administration-guide-667795-400867a1

General use cases
General use cases
There are three scenarios in which the FortiOS session initiation protocol (SIP) solution is usually deployed:
- The SIP server is in a private network that is protected from the internet by a FortiGate.
- The SIP clients are in a private network that is protected from the internet by a FortiGate.
- The SIP server is in a private network, such as a corporation's internal network or an ISP's network, that is protected from the internet by a FortiGate. The SIP clients are in a remote private network, such as a SOHO network, and behind a NAT device that is not aware of SIP applications.
The following VIP, NAT, and HNT examples show configurations for these common scenarios.
VIP
A FortiGate with SIP Application Layer Gateway (ALG) or SIP session helper protects the SIP server from the internet, while SIP phones from the internet need to register to the SIP server and establish calls through it.
A VIP needs to be configured for the SIP server, and the VIP must be applied in a firewall policy for the phones to send REGISTER messages through the FortiGate from port1 to port2.
Only one firewall policy needs to be configured for all SIP phones on both the internet and private network to register to the SIP server through port1 and set up SIP calls. This example assumes either SIP ALG or SIP session helper is enabled.
To configure the VIP for the SIP server:
config firewall vip
    edit "VIP_for_SIP_Server"
        set extip 172.20.120.50
        set extintf "port1"
        set mappedip "10.11.101.50"
    next
end
                                            To configure the firewall policy:
config firewall policy
    edit 1
        set srcintf "port1"
        set dstintf "port2"
        set srcaddr "all"
        set dstaddr "VIP_for_SIP_Server"
        set action accept
        set schedule "always"
        set service "SIP"
    next
end
                                            |  | Setting service toSIP and notall in the firewall policy can improve protection by restricting the data traffic passing through the FortiGate to the SIP call traffic only. | 
NAT
A FortiGate with SIP ALG or SIP session helper protects the SIP phones and the internal network from the internet, while SIP phones in the internal network need to register to the SIP server installed on the internet and establish calls through it.
One firewall policy needs to be configured with NAT enabled for SIP phones to send REGISTER messages through the FortiGate from port2 to port1. This example assumes either SIP ALG or SIP session helper is enabled.
To configure the firewall policy:
config firewall policy
    edit 1
        set srcintf "port2"
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "SIP"
        set nat enable
    next
end
                                            HNT
A FortiGate with SIP ALG protects the SIP server from the internet, while SIP phones are in remote private networks behind NAT devices that are not aware of the SIP application. This is only supported in proxy mode.
In this example, the SIP server is located in an ISP's service cloud that is protected by the FortiGate SIP ALG, and the SIP phones are installed in the home networks of the ISP's customers.
The SIP messages traversing the remote NAT devices might have their IP addresses translated by the NAT device at the network layer, but untranslated at the SIP application layer because those NAT devices are not aware of the SIP applications. This causes problems in a SIP session initiated process. Special configurations for the hosted NAT traversal (HNT) are required to resolve this issue.
To configure the FortiGate with HNT support for SIP phones A and B to set up calls with each other:
- Identify port1 as the external interface:config system interface
    edit "port1"
        set external enable
    next
end
- Configure the VIP for the SIP server:config firewall vip
    edit "VIP_for_SIP_Server"
        set extip 10.21.101.10
        set extintf "port1"
        set mappedip "10.30.120.20"
    next
end
- Configure a VoIP profile with HNT enabled:config voip profile
    edit "hnt"
        config sip
            set hosted-nat-traversal enable
            set hnt-restrict-source-ip enable
        end
    next
endhosted-nat-traversal must be enabled.hnt-restrict-source-ip does not have to be enabled,  but can be enabled  to restrict the RTP packets' source IP to be the same as the SIP packets' source IP.
- Apply the VoIP profile and VIP in a firewall policy for phone A and B to register and set up SIP calls through the FortiGate and SIP server:config firewall policy
    edit 1
        set srcintf "port1"
        set dstintf "port2"
        set srcaddr "all"
        set dstaddr "VIP_for_SIP_Server"
        set action accept
        set schedule "always"
        set service "SIP"
        set utm-status enable
        set voip-profile "hnt"
        set nat enable
    next
endnat must be enabled in the firewall policy.
