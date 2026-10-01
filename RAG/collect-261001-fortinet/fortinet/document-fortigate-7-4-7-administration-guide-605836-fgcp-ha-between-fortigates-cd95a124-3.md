---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates-cd95a124-3
title: "document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2023-05", "2023-05-29"]
keywords: ["apache", "ethernet"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124.md
source_anchor: ""
source_lines: [206, 304]
sha256: d42f5bc2695cffc1c42b294a12ba310cfb9d7941b00e6064b8f55dade2ec3450
---

# document-fortigate-7-4-7-administration-guide-605836-fgcp-ha-between-fortigates--cd95a124

To test session synchronization in the FGCP cluster:
- 
                                                    On PC1, verify the IP address and gateway: root@pc1:~# ifconfig eth1
eth1      Link encap:Ethernet  HWaddr 00:0c:29:a0:60:d6 
          inet addr:10.1.100.11  Bcast:10.1.100.255  Mask:255.255.255.0					 
          ...
root@pc1:~# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         10.1.100.1      0.0.0.0         UG    0      0        0 eth1
10.1.100.0      0.0.0.0         255.255.255.0   U     0      0        0 eth1
10.6.30.0       0.0.0.0         255.255.255.0   U     0      0        0 eth0
169.254.0.0     0.0.0.0         255.255.0.0     U     1000   0        0 eth0
- 
                                                    Using Wget, initiate a large file download with HTTP that will maintain a long session: root@pc1:~# wget http://172.16.200.55/big100MB.html   --keep-session-cookies --limit-rate=128k --progress=dot -S -r --delete-after
--2023-05-29 14:55:33--  http://172.16.200.55/big100MB.html
Connecting to 172.16.200.55:80... connected.
HTTP request sent, awaiting response... 
  HTTP/1.1 200 OK
  Date: Mon, 29 May 2023 21:55:41 GMT
  Server: Apache/2.4.18 (Ubuntu)
  Last-Modified: Thu, 01 Dec 2016 00:17:35 GMT
  ETag: "6126784-5428dbf967ad3"
  Accept-Ranges: bytes
  Content-Length: 101869444
  Vary: Accept-Encoding
  Keep-Alive: timeout=5, max=100
  Connection: Keep-Alive
  Content-Type: text/html
Length: 101869444 (97M) [text/html]
Saving to: '172.16.200.55/big100MB.html'
     0K .......... .......... .......... .......... ..........  0%  199K 8m18s
    50K .......... .......... .......... .......... ..........  0%  100K 12m26s
   100K .......... .......... .......... .......... ..........  0%  200K 11m3s
   150K .......... .......... .......... .......... ..........  0%  100K 12m25s
   200K .......... .......... .......... .......... ..........  0%  100K 13m14s
   250K .......... .......... .......... .......... ..........  0%  200K 12m24s
- 
                                                    On the primary FortiGate (FG-1800F-DC), check the session information: # diagnose sys session filter dport 80 # diagnose sys session list session info: proto=6 proto_state=01 duration=5 expire=3594 timeout=3600 flags=00000000 socktype=0 sockport=0 av_idx=0 use=4 origin-shaper= reply-shaper= per_ip_shaper= class_id=0 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/255 state=may_dirty npu synced log-start statistic(bytes/packets/allow_err): org=112/2/1 reply=60/1/1 tuples=2 tx speed(Bps/kbps): 0/0 rx speed(Bps/kbps): 0/0 orgin->sink: org pre->post, reply pre->post dev=13->14/14->13 gwy=0.0.0.0/0.0.0.0 hook=pre dir=org act=noop 10.1.100.11:54752->172.16.200.55:80(0.0.0.0:0) hook=post dir=reply act=noop 172.16.200.55:80->10.1.100.11:54752(0.0.0.0:0) pos/(before,after) 0/(0,0), 0/(0,0) misc=0 policy_id=1 pol_uuid_idx=15767 auth_info=0 chk_client_info=0 vd=0 serial=00000d80 tos=ff/ff app_list=0 app=0 url_cat=0 rpdb_link_id=00000000 ngfwid=n/a npu_state=0x4000c00 ofld-O ofld-R npu info: flag=0x81/0x81, offload=9/9, ips_offload=0/0, epid=133/132, ipid=132/133, vlan=0x0000/0x0000 vlifid=132/133, vtag_in=0x0000/0x0000 in_npu=1/1, out_npu=1/1, fwd_en=0/0, qid=12/12 total session: 1
- 
                                                    On the secondary FortiGate (FG-1800F), check that the session is synchronized: # diagnose sys session filter dport 80 # diagnose sys session list session info: proto=6 proto_state=01 duration=47 expire=3552 timeout=3600 flags=00000000 socktype=0 sockport=0 av_idx=0 use=3 origin-shaper= reply-shaper= per_ip_shaper= class_id=0 ha_id=0 policy_dir=0 tunnel=/ vlan_cos=0/255 state=dirty may_dirty npu syn_ses statistic(bytes/packets/allow_err): org=0/0/0 reply=0/0/0 tuples=2 tx speed(Bps/kbps): 0/0 rx speed(Bps/kbps): 0/0 orgin->sink: org pre->post, reply pre->post dev=13->14/14->13 gwy=0.0.0.0/0.0.0.0 hook=pre dir=org act=noop 10.1.100.11:54752->172.16.200.55:80(0.0.0.0:0) hook=post dir=reply act=noop 172.16.200.55:80->10.1.100.11:54752(0.0.0.0:0) pos/(before,after) 0/(0,0), 0/(0,0) misc=0 policy_id=1 pol_uuid_idx=0 auth_info=0 chk_client_info=0 vd=0 serial=00000d80 tos=ff/ff app_list=0 app=0 url_cat=0 rpdb_link_id=00000000 ngfwid=n/a npu_state=0x4000000 npu info: flag=0x00/0x00, offload=0/0, ips_offload=0/0, epid=0/0, ipid=0/0, vlan=0x0000/0x0000 vlifid=0/0, vtag_in=0x0000/0x0000 in_npu=0/0, out_npu=0/0, fwd_en=0/0, qid=0/0 no_ofld_reason: total session: 1
To test failover in the FGCP cluster:
- 
                                                    On the switch connected to port5 of the primary FortiGate, change port2's status to be down: config switch physical-port 
    edit port2
        set status down 
    next
end
- 
                                                    Check the HA status on the primary FortiGate (FG-1800F-DC), which now becomes the secondary device: # get system ha status  
HA Health Status: 
    WARNING: FG180FTK*******1 has mondev down; 
Model: FortiGate-1800F
Mode: HA A-P
Group Name: Example_cluster
Group ID: 0
Debug: 0
Cluster Uptime: 0 days 1:16:13
Cluster state change time: 2023-05-29 20:08:56
Primary selected using:
    <2023/05/29 20:08:56> vcluster-1: FG180FTK*******2 is selected as the primary because the value 0 of link-failure + pingsvr-failure is less than peer member FG180FTK*******1.
    <2023/05/29 19:11:14> vcluster-1: FG180FTK*******1 is selected as the primary because its uptime is larger than peer member FG180FTK*******2.
    <2023/05/29 18:59:45> vcluster-1: FG180FTK*******2 is selected as the primary because its uptime is larger than peer member FG180FTK*******1.
    <2023/05/29 18:59:45> vcluster-1: FG180FTK*******1 is selected as the primary because its override priority is larger than peer member FG180FTK*******2.
ses_pickup: enable, ses_pickup_delay=disable
override: disable
...
Secondary   : FortiGate-1800F , FG180FTK*******1, HA cluster index = 1
Primary     : FortiGate-1800F , FG180FTK*******2, HA cluster index = 0
number of vcluster: 1
vcluster 1: standby 169.254.0.1
Secondary: FG180FTK*******1, HA operating index = 1
Primary: FG180FTK*******2, HA operating index = 0
- 
                                                    Check the HA status on the new primary FortiGate (FG-1800F): # get system ha status   
HA Health Status: 
    WARNING: FG180FTK*******1 has mondev down; 
Model: FortiGate-1800F
Mode: HA A-P
Group Name: Example_cluster
Group ID: 0
Debug: 0
Cluster Uptime: 0 days 1:19:9
Cluster state change time: 2023-05-29 20:08:56
Primary selected using:
    <2023/05/29 20:08:56> vcluster-1: FG180FTK*******2 is selected as the primary because the value 0 of link-failure + pingsvr-failure is less than peer member FG180FTK*******1.
    <2023/05/29 19:11:14> vcluster-1: FG180FTK*******1 is selected as the primary because its uptime is larger than peer member FG180FTK*******2.
    <2023/05/29 18:59:45> vcluster-1: FG180FTK*******2 is selected as the primary because its uptime is larger than peer member FG180FTK*******1.
    <2023/05/29 18:55:03> vcluster-1: FG180FTK*******2 is selected as the primary because it's the only member in the cluster.
ses_pickup: enable, ses_pickup_delay=disable
override: disable
...
Primary     : FortiGate-1800F , FG180FTK*******2, HA cluster index = 0
Secondary   : FortiGate-1800F , FG180FTK*******1, HA cluster index = 1
number of vcluster: 1
vcluster 1: work 169.254.0.1
Primary: FG180FTK*******2, HA operating index = 0
Secondary: FG180FTK*******1, HA operating index = 1
- 
