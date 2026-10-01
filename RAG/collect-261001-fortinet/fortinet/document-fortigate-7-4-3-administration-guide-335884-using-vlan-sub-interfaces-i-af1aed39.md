---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-335884-using-vlan-sub-interfaces-i-af1aed39
title: "diagnose sys session filter policy  1"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-335884-using-vlan-sub-interfaces-i-af1aed39.md
source_anchor: ""
source_lines: [1, 93]
sha256: 2784fe3e6ceaa1bbe0b1bb6c72e8fb46b8a74cad58db061a056090998cc99fa0
---

# diagnose sys session filter policy  1

Using VLAN sub-interfaces in virtual wire pairs
Using VLAN sub-interfaces in virtual wire pairs
VLAN sub-interfaces, such as regular 802.1Q and 802.1ad (QinQ), are allowed to be members of a virtual wire pair.
Example
In this example, the FortiGate has two VLAN interfaces. The first interface is a QinQ (802.1ad) interface over the physical interface port3. The second interface is a basic 802.1Q VLAN interface over physical interface port5. These two interfaces are grouped in a virtual wire pair so that bi-directional traffic is allowed. This example demonstrates ICMP from the client (3.3.3.4) sent to the server (3.3.3.1).
To configure VLAN sub-interfaces in a virtual wire pair:
- 
                                                    Configure the QinQ interfaces: config system interface
    edit "8021ad-port3"
        set vdom "vdom1"
        set vlan-protocol 8021ad
        set device-identification enable
        set role lan
        set snmp-index 31
        set interface "port3"
        set vlanid 3
    next
    edit "8021Q"
        set vdom "vdom1"
        set device-identification enable
        set role lan
        set snmp-index 32
        set interface "8021ad-port3"
        set vlanid 33
    next
end
- 
                                                    Configure the 802.1Q interface: config system interface
    edit "8021q-port5"
        set vdom "vdom1"
        set device-identification enable
        set role lan
        set snmp-index 33
        set interface "port5"
        set vlanid 5
    next
end
- 
                                                    Configure the virtual wire pair: config system virtual-wire-pair
    edit "VWP1"
        set member "8021Q" "8021q-port5"
    next
end
- 
                                                    Configure the firewall policy: config firewall policy
    edit 1
        set name "1"
        set srcintf "8021Q" "8021q-port5"
        set dstintf "8021Q" "8021q-port5"
        set action accept
        set srcaddr "all"
        set dstaddr "all"
        set schedule "always"
        set service "ALL"
    next
end
To verify that bi-directional traffic passes through the FortiGate:
# diagnose sys session filter policy  1
# diagnose sys session list
session info: proto=1 proto_state=00 duration=18 expire=42 timeout=0 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3
origin-shaper=
reply-shaper=
per_ip_shaper=
class_id=0 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/0
state=may_dirty br npu 
statistic(bytes/packets/allow_err): org=168/2/1 reply=168/2/1 tuples=2
tx speed(Bps/kbps): 0/0 rx speed(Bps/kbps): 0/0
orgin->sink: org pre->post, reply pre->post dev=56->55/55->56 gwy=0.0.0.0/0.0.0.0
hook=pre dir=org act=noop 3.3.3.4:3072->3.3.3.1:8(0.0.0.0:0)
hook=post dir=reply act=noop 3.3.3.1:3072->3.3.3.4:0(0.0.0.0:0)
src_mac=08:5b:0e:71:bf:c6  dst_mac=d4:76:a0:5d:b2:de
misc=0 policy_id=1 pol_uuid_idx=534 auth_info=0 chk_client_info=0 vd=3
serial=00005f6c tos=ff/ff app_list=0 app=0 url_cat=0
rpdb_link_id=00000000 ngfwid=n/a
npu_state=0x4000c00 ofld-O ofld-R
npu info: flag=0x81/0x81, offload=8/8, ips_offload=0/0, epid=187/156, ipid=156/187, vlan=0x0005/0x0021
vlifid=156/187, vtag_in=0x0005/0x0021 in_npu=1/1, out_npu=1/1, fwd_en=0/0, qid=0/5
total session 1
DVLAN QinQ on NP7 platforms over virtual wire pairs
DVLAN 802.1ad and 802.1Q modes are supported on NP7 platforms over virtual wire pairs, which provides better performance and packet processing.
The default DVLAN mode is 802.1ad, but the DVLAN mode can be changed using diagnose npu np7 dvlan-mode <dvlan_mode> {<npid> | all}. The DVLAN mode can be applied to a specific NPID or all NPIDs. For example:
- diagnose npu np7 dvlan-mode 802.1AD 0 will set NP0 to work in 802.1ad mode.
- diagnose npu np7 dvlan-mode 802.1Q all will set all NPUs to work in 802.1Q mode.
|  | A reboot is required for custom DVLAN settings to take effect. To avoid any inconveniences or disruptions, changing the DVLAN settings should be done during a scheduled downtime or maintenance window. The DVLAN mode should only be changed if you are solely using the virtual wire pair (VWP) and are seeking to enhance performance. Enabling this feature may impact VLAN interfaces within your network. | 
In the virtual wire pair settings, the outer-vlan-id can be set. This is the same value as the outer provider-tag (S-Tag).
To configure the outer VLAN ID:
config system virtual-wire-pair
    edit "dvlan-test"
        set member "port33" "port34"
        set wildcard-vlan enable
        set outer-vlan-id 1234
    next
end
