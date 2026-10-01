---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-421028-central-snat-dd9b39f4-2
title: "document-fortigate-7-4-3-administration-guide-421028-central-snat-dd9b39f4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-421028-central-snat-dd9b39f4.md
source_anchor: ""
source_lines: [148, 252]
sha256: 92925f0b872409af8c7381e1f28934e256e31b8f34db48b13c22d3c8c9201cb8
---

# document-fortigate-7-4-3-administration-guide-421028-central-snat-dd9b39f4

Traffic with original port in the range between 50000-65535 will be NAT'd with the Overload type IP Pool. The mapped port is in the same port range:
session info: proto=17 proto_state=01 duration=3 expire=176 timeout=0 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3
origin-shaper=
reply-shaper=
per_ip_shaper=
class_id=0 ha_id=0 policy_dir=0 tunnel=/ helper=dns-udp vlan_cos=0/255
state=may_dirty 
statistic(bytes/packets/allow_err): org=71/1/1 reply=123/1/1 tuples=2
tx speed(Bps/kbps): 23/0 rx speed(Bps/kbps): 40/0
orgin->sink: org pre->post, reply pre->post dev=9->6/6->9 gwy=192.168.2.1/192.168.0.10
hook=post dir=org act=snat 192.168.0.10:52540->4.2.2.1:53(192.168.2.201:52540)
hook=pre dir=reply act=dnat 4.2.2.1:53->192.168.2.201:52540(192.168.0.10:52540)
dst_mac=04:d5:90:5f:a2:2a
misc=0 policy_id=2 auth_info=0 chk_client_info=0 vd=0
serial=00011399 tos=ff/ff app_list=0 app=0 url_cat=0
sdwan_mbr_seq=0 sdwan_service_id=0
rpdb_link_id=00000000 rpdb_svc_id=0 ngfwid=n/a
npu_state=0x040000
Traffic with original port outside the range of 50000-65535 will be NAT'd to the outgoing interface IP:
session info: proto=6 proto_state=01 duration=3 expire=3597 timeout=3600 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3
origin-shaper=
reply-shaper=
per_ip_shaper=
class_id=0 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/255
state=may_dirty 
statistic(bytes/packets/allow_err): org=2262/10/1 reply=2526/11/1 tuples=2
tx speed(Bps/kbps): 741/5 rx speed(Bps/kbps): 828/6
orgin->sink: org pre->post, reply pre->post dev=9->6/6->9 gwy=192.168.2.1/192.168.0.10
hook=post dir=org act=snat 192.168.0.10:49805->142.250.68.66:443(192.168.2.86:62214)
hook=pre dir=reply act=dnat 142.250.68.66:443->192.168.2.86:62214(192.168.0.10:49805)
pos/(before,after) 0/(0,0), 0/(0,0)
dst_mac=04:d5:90:5f:a2:2a
misc=0 policy_id=2 auth_info=0 chk_client_info=0 vd=0
serial=0001139a tos=ff/ff app_list=0 app=0 url_cat=0
sdwan_mbr_seq=0 sdwan_service_id=0
rpdb_link_id=00000000 rpdb_svc_id=0 ngfwid=n/a
npu_state=0x040000
Protocols which do not use ports, such as ICMP, will be NAT'd to the outgoing interface IP:
session info: proto=1 proto_state=00 duration=7 expire=59 timeout=0 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3
origin-shaper=
reply-shaper=
per_ip_shaper=
class_id=0 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/255
state=may_dirty 
statistic(bytes/packets/allow_err): org=480/8/1 reply=480/8/1 tuples=2
tx speed(Bps/kbps): 66/0 rx speed(Bps/kbps): 66/0
orgin->sink: org pre->post, reply pre->post dev=9->6/6->9 gwy=192.168.2.1/192.168.0.10
hook=post dir=org act=snat 192.168.0.10:1->4.2.2.1:8(192.168.2.86:62209)
hook=pre dir=reply act=dnat 4.2.2.1:62209->192.168.2.86:0(192.168.0.10:1)
dst_mac=04:d5:90:5f:a2:2a
misc=0 policy_id=2 auth_info=0 chk_client_info=0 vd=0
serial=0001138b tos=ff/ff app_list=0 app=0 url_cat=0
sdwan_mbr_seq=0 sdwan_service_id=0
rpdb_link_id=00000000 rpdb_svc_id=0 ngfwid=n/a
npu_state=0x040000
Example four
In the following example, two central SNAT rules will be created:
- 
                                                    Rule 3 will have a destination port set and IP pool test-ippool4-3 applied.
- 
                                                    Rule 5 will have IP pool test-ippool4-1 applied but will not set the destination port.
Example traffic will then be passed to see how the rule is matched.
To test central SNAT rule destination port support:
- 
                                                    Configure central SNAT rule 3 with the destination port range specified: config firewall ippool
    edit "test-ippool4-3"
        set startip 172.16.200.150
        set endip 172.16.200.150
    next
end
config firewall central-snat-map
    edit 3
        set srcintf "port24"
        set dstintf "port17"
        set orig-addr "all"
        set dst-addr "all"
        set protocol 6
        set nat-ippool "test-ippool4-3"
        set dst-port 80-443
    next
end
- 
                                                    Configure central SNAT rule 5: config firewall ippool
    edit "test-ippool4-1"
        set startip 172.16.200.151
        set endip 172.16.200.151
    next
end
config firewall central-snat-map
    edit 5
        set srcintf "port24"
        set dstintf "port17"
        set orig-addr "all"
        set dst-addr "all"
        set nat-ippool "test-ippool4-1"
    next
end
- 
                                                    Send HTTP traffic to pass through the FortiGate that is expected to match central SNAT rule 3. IP pool test-ippool4-3 will perform source NAT.
- 
                                                    Check the session to review for expected behavior: # diagnose sys session list session info: proto=6 proto_state=01 duration=2 expire=3599 timeout=3600 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3 origin-shaper= reply-shaper= per_ip_shaper= class_id=0 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/255 state=log may_dirty f00 statistic(bytes/packets/allow_err): org=1800/31/1 reply=77304/60/1 tuples=2 tx speed(Bps/kbps): 602/4 rx speed(Bps/kbps): 25854/206 orgin->sink: org pre->post, reply pre->post dev=24->17/17->24 gwy=172.16.200.55/10.1.100.42 hook=post dir=org act=snat 10.1.100.42:46731->172.16.200.55:80(172.16.200.150:46731) hook=pre dir=reply act=dnat 172.16.200.55:80->172.16.200.150:46731(10.1.100.42:46731) pos/(before,after) 0/(0,0), 0/(0,0) misc=0 policy_id=99 pol_uuid_idx=15864 auth_info=0 chk_client_info=0 vd=0 serial=00003c37 tos=ff/ff app_list=0 app=0 url_cat=0 rpdb_link_id=00000000 ngfwid=n/a npu_state=0x4000001 no_offload no_ofld_reason: disabled-by-policy total session 1
- 
                                                    Send PING traffic to pass through the FortiGate that is expected to match central SNAT rule 5. IP pool test-ippool4-1 will perform source NAT.
- 
                                                    Check the session to review for expected behavior: # diagnose sys session list session info: proto=1 proto_state=00 duration=2 expire=59 timeout=0 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3 origin-shaper= reply-shaper= per_ip_shaper= class_id=0 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/255 state=log may_dirty f00 statistic(bytes/packets/allow_err): org=252/3/1 reply=252/3/1 tuples=2 tx speed(Bps/kbps): 99/0 rx speed(Bps/kbps): 99/0 orgin->sink: org pre->post, reply pre->post dev=24->17/17->24 gwy=172.16.200.55/10.1.100.42 hook=post dir=org act=snat 10.1.100.42:36732->172.16.200.55:8(172.16.200.151:36732) hook=pre dir=reply act=dnat 172.16.200.55:36732->172.16.200.151:0(10.1.100.42:36732) misc=0 policy_id=99 pol_uuid_idx=15864 auth_info=0 chk_client_info=0 vd=0 serial=00003f62 tos=ff/ff app_list=0 app=0 url_cat=0 rpdb_link_id=00000000 ngfwid=n/a npu_state=0x4000001 no_offload no_ofld_reason: disabled-by-policy total session 1
