---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-internet-is-not-working-with-one-of-the-sd-wan-m-8b38e10a
title: "fortigate-3-troubleshooting-tip-internet-is-not-working-with-one-of-the-sd-wan-m-8b38e10a"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2024-10-10"]
keywords: []
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-internet-is-not-working-with-one-of-the-sd-wan-m-8b38e10a.md
source_anchor: ""
source_lines: [1, 153]
sha256: c28ec1315dcbe4dd57a289582506514082ea4eee1b3ee4d8cc2dbb53aa9c57e9
---

# fortigate-3-troubleshooting-tip-internet-is-not-working-with-one-of-the-sd-wan-m-8b38e10a

**Description**


This article describes how to handle an issue where the Internet is not working with one of the SD-WAN member when IP pool is called in the policy.


**Scope**


FortiGate.

**Solution**


Users may face an issue while accessing the internet when there is an outgoing interface as an SD-WAN with more than one WAN interface, such as WAN1 and WAN2, in the same firewall policy, and where the IP pool is used to perform SNAT.


**config system interface** 

    edit "wan1"

        set alias to_ISP1

        set mode dhcp

    next

        edit "wan2"

            set alias to_ISP2

            set ip 1.1.10.1 255.255.255.248

        next

    end**config system sdwan**

    set status enable

**config members**

    edit 1

        set interface "wan1"

    next

        edit 2

            set interface "wan2"

            set gateway 1.1.10.6

        next

    end

end

**config firewall ippool**

    edit "1.1.10.1"

        set startip 1.1.10.1

        set endip 1.1.10.5

    next

end



**config firewall policy**

edit 1

set srcintf "port2"

set dstintf "sdwan"

set srcaddr "all"

set dstaddr "all"

set action accept

set schedule "always"

set service "ALL"

set logtraffic all

set logtraffic-start enable

set ippool enable

set poolname "1.1.10.1"

set nat enable

next




In this case, if traffic goes via WAN1 with SNAT as 1.1.10.1-1.1.10.5, internet will not work.

To overcome the issue, it is necessary to set 'associated-interface' in the IP pool. This option will only be available from CLI.


**config firewall ippool**

    edit "1.1.10.1"

        set type overload

        set startip 1.1.10.1

        set endip 1.1.10.5

        set associated-interface 'wan2'       <- Assigning the wan2 interface.

        set comments ''

    next

end


With this, the IP pool will be applied only when traffic is going via the WAN2 interface. Traffic going via WAN1 will use the WAN1 interface IP for SNAT.

The debugs for such session would resemble the excerpt below:


2024-10-10 09:12:36 id=65308 trace_id=5 func=print_pkt_detail line=5920 msg="vd-root:0 received a packet(proto=1, 192.168.30.3:1->8.8.8.8:2048) tun_id=0.0.0.0 from port2. type=8, code=0, id=1, seq=16."

2024-10-10 09:12:36 id=65308 trace_id=5 func=init_ip_session_common line=6110 msg="allocate a new session-00060766"

2024-10-10 09:12:36 id=65308 trace_id=5 func=__vf_ip_route_input_rcu line=1988 msg="find a route: flag=00000000 gw-192.168.10.20 via wan1"

2024-10-10 09:12:36 id=65308 trace_id=5 func=__iprope_tree_check line=535 msg="gnum-100004, use addr/intf hash, len=2"

2024-10-10 09:12:36 id=65308 trace_id=5 func=get_new_addr line=1274 msg="find SNAT: IP-1.1.10.4(from IPPOOL), port-60418"

2024-10-10 09:12:36 id=65308 trace_id=5 func=fw_forward_handler line=998 msg="Allowed by Policy-1: SNAT"

2024-10-10 09:12:36 id=65308 trace_id=5 func=ip_session_confirm_final line=3128 msg="npu_state=0x100, hook=4"

2024-10-10 09:12:36 id=65308 trace_id=5 func=__ip_session_run_tuple line=3461 msg="SNAT 192.168.30.3->1.1.10.4:60418"


**Related article:**
