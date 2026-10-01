---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-2-administration-guide-632796-ospf-with-ipsec-vpn-for-net-e677642e-2
title: "document-fortigate-7-4-2-administration-guide-632796-ospf-with-ipsec-vpn-for-net-e677642e"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-2-administration-guide-632796-ospf-with-ipsec-vpn-for-net-e677642e.md
source_anchor: ""
source_lines: [48, 53]
sha256: 7953637d1d2fed901c241f366e1a6685d402131b02394d9eb2dba51d43857bbc
---

# document-fortigate-7-4-2-administration-guide-632796-ospf-with-ipsec-vpn-for-net-e677642e

To check VPN and OSPF states using diagnose and get commands:
- Run the HQ1 # diagnose vpn ike gateway list command. The system should return the following:vd: root/0 name: pri_HQ2 version: 1 interface: port1 11 addr: 172.16.200.1:500 -> 172.16.202.1:500 virtual-interface-addr: 10.10.10.1 -> 10.10.10.2 created: 1024s ago IKE SA: created 1/1 established 1/1 time 0/0/0 ms IPsec SA: created 1/3 established 1/2 time 0/5/10 ms id/spi: 45 d184777257b4e692/e2432f834aaf5658 direction: responder status: established 1024-1024s ago = 0ms proposal: aes128-sha256 key: 9ed41fb06c983344-189538046f5ad204 lifetime/rekey: 86400/85105 DPD sent/recv: 00000003/00000000 vd: root/0 name: sec_HQ2 version: 1 interface: port2 12 addr: 172.17.200.1:500 -> 172.17.202.1:500 virtual-interface-addr: 10.10.11.1 -> 10.10.11.2 created: 346s ago IKE SA: created 1/1 established 1/1 time 0/0/0 ms IPsec SA: created 1/1 established 1/1 time 0/10/15 ms id/spi: 48 d909ed68636b1ea5/163015e73ea050b8 direction: initiator status: established 0-0s ago = 0ms proposal: aes128-sha256 key: b9e93c156bdf4562-29db9fbafa256152 lifetime/rekey: 86400/86099 DPD sent/recv: 00000000/00000000
- Run the HQ1 # diagnose vpn tunnel list command. The system should return the following:list all ipsec tunnel in vd 0 name=pri_HQ2 ver=1 serial=1 172.16.200.1:0->172.16.202.1:0 tun_id=172.16.202.1 bound_if=11 lgwy=static/1 tun=intf/0 mode=auto/1 encap=none/528 options[0210]=create_dev frag-rfc accept_traffic=1 proxyid_num=1 child_num=0 refcnt=14 ilast=2 olast=2 ad=/0 stat: rxp=102 txp=105 rxb=14064 txb=7816 dpd: mode=on-demand on=1 idle=20000ms retry=3 count=0 seqno=3 natt: mode=none draft=0 interval=0 remote_port=0 proxyid=pri_HQ2 proto=0 sa=1 ref=2 serial=1 auto-negotiate src: 0:0.0.0.0/0.0.0.0:0 dst: 0:0.0.0.0/0.0.0.0:0 SA: ref=3 options=18227 type=00 soft=0 mtu=1438 expire=42254/0B replaywin=2048 seqno=6a esn=0 replaywin_lastseq=00000067 itn=0 life: type=01 bytes=0/0 timeout=42932/43200 dec: spi=1071b4ee esp=aes key=16 032036b24a4ec88da63896b86f3a01db ah=sha1 key=20 3962933e24c8da21c65c13bc2c6345d643199cdf enc: spi=ec89b7e3 esp=aes key=16 92b1d85ef91faf695fca05843dd91626 ah=sha1 key=20 2de99d1376506313d9f32df6873902cf6c08e454 dec:pkts/bytes=102/7164, enc:pkts/bytes=105/14936 name=sec_HQ2 ver=1 serial=2 172.17.200.1:0->172.17.202.1:0 tun_id=172.17.202.1 bound_if=12 lgwy=static/1 tun=intf/0 mode=auto/1 encap=none/528 options[0210]=create_dev frag-rfc accept_traffic=1 proxyid_num=1 child_num=0 refcnt=14 ilast=3 olast=0 ad=/0 stat: rxp=110 txp=114 rxb=15152 txb=8428 dpd: mode=on-demand on=1 idle=20000ms retry=3 count=0 seqno=3 natt: mode=none draft=0 interval=0 remote_port=0 proxyid=sec_HQ2 proto=0 sa=1 ref=2 serial=1 auto-negotiate src: 0:0.0.0.0/0.0.0.0:0 dst: 0:0.0.0.0/0.0.0.0:0 SA: ref=3 options=18227 type=00 soft=0 mtu=1438 expire=42927/0B replaywin=2048 seqno=2 esn=0 replaywin_lastseq=00000002 itn=0 life: type=01 bytes=0/0 timeout=42931/43200 dec: spi=1071b4ef esp=aes key=16 bcdcabdb7d1c7c695d1f2e0f5441700a ah=sha1 key=20 e7a0034589f82eb1af41efd59d0b2565fef8d5da enc: spi=ec89b7e4 esp=aes key=16 234240b69e61f6bdee2b4cdec0f33bea ah=sha1 key=20 f9d4744a84d91e5ce05f5984737c2a691a3627e8 dec:pkts/bytes=1/68, enc:pkts/bytes=1/136
- Run the HQ1 # get router info ospf neighbor command. The system should return the following:OSPF process 0, VRF 0: Neighbor ID Pri State Dead Time Address Interface 2.2.2.2 1. Full/ - 00:00:37 10.10.10.2 pri_HQ2 2.2.2.2 1. Full/ - 00:00:32 10.10.11.2 sec_HQ2
- Run the HQ1 # get router info routing-table ospf command. The system should return the following:Routing table for VRF=0 O 172.16.101.0/24 [110/20] via 10.10.10.2, pri_HQ2 , 00:03:21 In case the primary tunnel is down after route convergence.
- Run the HQ1 # get router info routing-table ospf command. The system should return the following:Routing table for VRF=0 O 172.16.101.0/24 [110/110] via 10.10.11.2, sec_HQ2 , 00:00:01
