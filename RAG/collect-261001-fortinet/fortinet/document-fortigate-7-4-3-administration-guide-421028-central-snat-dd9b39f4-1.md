---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-421028-central-snat-dd9b39f4-1
title: "document-fortigate-7-4-3-administration-guide-421028-central-snat-dd9b39f4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-421028-central-snat-dd9b39f4.md
source_anchor: ""
source_lines: [1, 147]
sha256: 9688de31b50f56a8a5a0eb3bd912ce311921a6d9095ba56a3ca6d50e45b875ea
---

# document-fortigate-7-4-3-administration-guide-421028-central-snat-dd9b39f4

Central SNAT
Central SNAT
The central SNAT table enables you to define and control (with more granularity) the address translation performed by FortiGate. With the NAT table, you can define the rules for the source address or address group, and which IP pool the destination address uses.
FortiGate reads the NAT rules from the top down until it hits a matching rule for the incoming address. This enables you to create multiple NAT policies that dictate which IP pool is used based on source address, destination address, and source port. NAT policies can be rearranged within the policy list. NAT policies are applied to network traffic after a security policy.
The central SNAT table allows you to create, edit, delete, and clone central SNAT entries.
Central SNAT notes
- The central NAT feature is not enabled by default.
- If central NAT is enabled, the NAT option under IPv4 policies is skipped and SNAT must be done via central-snat-map . The firewall policy list and dialog boxes have messages and redirection links to show this information.
- If NGFW mode is policy-based, then it is assumed that central NAT (specifically SNAT) is enabled implicitly.
Sample configuration
To enable central SNAT from the GUI:
- 
                                                    In System > Settings, under System Operations Settings, enable Central SNAT.
- 
                                                    Click Apply.
To enable or disable central SNAT using the CLI:
config system settings
   set central-nat {enable | disable}
end
                                            When central NAT is enabled, Policy & Objects displays the Central SNAT section.
The Central SNAT policy has many options:
| Field | Description | 
|---|---|
| Type | Specify whether you are performing SNAT on IPv4 or IPv6. This option only appears when IPv6 is enabled under Feature Visibility. | 
| Incoming Interface | Specify one or more interfaces for the ingress traffic. | 
| Outgoing Interface | Specify one or more interfaces for the egress traffic. | 
| Source Address | Specify the address or address group of the source. | 
| Destination Address | Specify the address or address group of the destination. | 
| NAT | Enable or disable to perform NAT. When disabled, no source address translation will occur. | 
| IP Pool Configuration | Use outgoing interface address:  Use Dynamic IP Pool:  | 
| Protocol | Choose from any, TCP, UDP, SCTP, or specify the protocol number to match. For example, for ICMP, click specify with the protocol number 1. | 
| Explicit port mapping | Enable in order to match this NAT policy only when the following ports are a match:  Explicit port mapping cannot apply to some protocols which do not use ports, such as ICMP. When enabling a NAT policy which uses Explicit port mapping, always consider that ICMP traffic will not match this policy. When using IP Pools, only the Overload type IP Pool allows Explicit port mapping. When Explicit port mapping is applied, you must define an original source port range and a translated sort port range. The source port will map one to one with the translated port. Refer to Dynamic SNAT to understand how each IP Pool type works. | 
| Comments | Enter comments for this NAT policy. | 
| Enable this policy | Enable or disable this policy. | 
To configure central SNAT using the CLI:
config firewall central-snat-map
    edit <policyID number>
        set status {enable | disable}
        set orig-addr <valid address object preconfigured on the FortiGate>
        set srcintf <name of interface on the FortiGate>
        set dst-addr <valid address object preconfigured on the FortiGate>
        set dstintf <name of interface on the FortiGate>
        set protocol <integer for protocol number>
        set dst-port <integer for destination port or port range>
        set orig-port <integer for original port number>
        set nat-port <integer for translated port number>
        set comments <string>
    next
end
                                            |  | Setting the destination port for traffic matching is available when the protocols are TCP, UDP, or SCTP. | 
Example one
Apply SNAT to all traffic from port2 to port3.
To configure from the CLI:
config firewall central-snat-map
    edit 1
        set srcintf "port3"
        set dstintf "port2"
        set orig-addr "all"
        set dst-addr "all"
    next
end
                                            Example two
Apply an IP Pool to all traffic from port3 to port2 that are TCP. NAT all other traffic using the outgoing interface IP.
To configure from the CLI:
config firewall ippool
    edit "Overload-IPPOOL"
        set startip 192.168.2.201
        set endip 192.168.2.202
    next
end
config firewall central-snat-map
    edit 1
        set srcintf "port3"
        set dstintf "port2"
        set orig-addr "all"
        set dst-addr "all"
        set protocol 6
        set nat-ippool "Overload-IPPOOL"
    next
    edit 2
        set srcintf "port3"
        set dstintf "port2"
        set orig-addr "all"
        set dst-addr "all"
    next
end
                                            To collect session table output from the CLI:
diagnose sys session list
The TCP session (protocol 6) is NAT'd with Overload-IPPOOL to 192.168.2.201:
session info: proto=6 proto_state=05 duration=14 expire=0 timeout=3600 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3
origin-shaper=
reply-shaper=
per_ip_shaper=
class_id=0 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/255
state=may_dirty 
statistic(bytes/packets/allow_err): org=860/7/1 reply=555/8/1 tuples=2
tx speed(Bps/kbps): 60/0 rx speed(Bps/kbps): 38/0
orgin->sink: org pre->post, reply pre->post dev=9->6/6->9 gwy=192.168.2.1/192.168.0.10
hook=post dir=org act=snat 192.168.0.10:49531->23.57.57.114:443(192.168.2.201:61776)
hook=pre dir=reply act=dnat 23.57.57.114:443->192.168.2.201:61776(192.168.0.10:49531)
pos/(before,after) 0/(0,0), 0/(0,0)
dst_mac=04:d5:90:5f:a2:2a
misc=0 policy_id=2 auth_info=0 chk_client_info=0 vd=0
serial=00011065 tos=ff/ff app_list=0 app=0 url_cat=0
sdwan_mbr_seq=0 sdwan_service_id=0
rpdb_link_id=00000000 rpdb_svc_id=0 ngfwid=n/a
npu_state=0x040000
A UDP session (protocol 17) is NAT'd to the outgoing interface IP address 192.168.2.86:
session info: proto=17 proto_state=01 duration=16 expire=163 timeout=0 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3
origin-shaper=
reply-shaper=
per_ip_shaper=
class_id=0 ha_id=0 policy_dir=0 tunnel=/ helper=dns-udp vlan_cos=0/255
state=may_dirty 
statistic(bytes/packets/allow_err): org=59/1/1 reply=187/1/1 tuples=2
tx speed(Bps/kbps): 3/0 rx speed(Bps/kbps): 11/0
orgin->sink: org pre->post, reply pre->post dev=9->6/6->9 gwy=192.168.2.1/192.168.0.10
hook=post dir=org act=snat 192.168.0.10:52177->4.2.2.1:53(192.168.2.86:61770)
hook=pre dir=reply act=dnat 4.2.2.1:53->192.168.2.86:61770(192.168.0.10:52177)
dst_mac=04:d5:90:5f:a2:2a
misc=0 policy_id=2 auth_info=0 chk_client_info=0 vd=0
serial=00011061 tos=ff/ff app_list=0 app=0 url_cat=0
sdwan_mbr_seq=0 sdwan_service_id=0
rpdb_link_id=00000000 rpdb_svc_id=0 ngfwid=n/a
npu_state=0x040000
Example three
Apply an IP Pool to all traffic from port3 to port2 that have a specific original port range, mapping the ports to the same NAT'd port range. Nat all other traffic using the outgoing interface IP.
To configure from the CLI:
config firewall central-snat-map
    edit 1
        set srcintf "port3"
        set dstintf "port2"
        set orig-addr "all"
        set dst-addr "all"
        set orig-port 50000-65535
        set nat-ippool "Overload-IPPOOL"
        set nat-port 50000-65535
    next
    edit 2
        set srcintf "port3"
        set dstintf "port2"
        set orig-addr "all"
        set dst-addr "all"
    next
end
                                            To collect session table output from the CLI:
diagnose sys session list
