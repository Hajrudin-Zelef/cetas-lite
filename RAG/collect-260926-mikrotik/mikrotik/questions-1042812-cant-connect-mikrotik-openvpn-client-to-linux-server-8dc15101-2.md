---
id: collect-260926-mikrotik/mikrotik/questions-1042812-cant-connect-mikrotik-openvpn-client-to-linux-server-8dc15101-2
title: "ip addr"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2020-11-17"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-1042812-cant-connect-mikrotik-openvpn-client-to-linux-server-8dc15101.md
source_anchor: ""
source_lines: [135, 164]
sha256: a1f3f69a9d875f56c6d10b94763ab9afbbf433c4f75cd35b17a6a3c0e57b027e
---

# ip addr

2020-11-17 01:53:40 us=841342 195.111.111.2:34720 TCPv4_SERVER READ [1282] from [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ ] pid=2 DATA len=1268
2020-11-17 01:53:40 us=841519 195.111.111.2:34720 VERIFY OK: depth=1, CN=Easy-RSA CA
2020-11-17 01:53:40 us=841600 195.111.111.2:34720 VERIFY OK: depth=0, CN=mikrotik
2020-11-17 01:53:40 us=841809 195.111.111.2:34720 TCPv4_SERVER WRITE [77] to [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ 2 ] pid=3 DATA len=51
2020-11-17 01:53:40 us=846294 195.111.111.2:34720 TCPv4_SERVER READ [22] from [AF_INET]195.111.111.2:34720: P_ACK_V1 kid=0 [ 3 ]
2020-11-17 01:53:40 us=891974 195.111.111.2:34720 TCPv4_SERVER READ [303] from [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ ] pid=3 DATA len=289
2020-11-17 01:53:40 us=892064 195.111.111.2:34720 TCPv4_SERVER WRITE [259] to [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ 3 ] pid=4 DATA len=233
2020-11-17 01:53:40 us=896730 195.111.111.2:34720 TCPv4_SERVER READ [22] from [AF_INET]195.111.111.2:34720: P_ACK_V1 kid=0 [ 4 ]
2020-11-17 01:53:40 us=896755 195.111.111.2:34720 Control Channel: TLSv1.2, cipher TLSv1.2 ECDHE-RSA-AES256-GCM-SHA384, 2048 bit RSA
2020-11-17 01:53:40 us=896771 195.111.111.2:34720 [mikrotik] Peer Connection Initiated with [AF_INET]195.111.111.2:34720
2020-11-17 01:53:40 us=896788 mikrotik/195.111.111.2:34720 MULTI_sva: pool returned IPv4=10.101.0.2, IPv6=(Not enabled)
2020-11-17 01:53:40 us=896826 mikrotik/195.111.111.2:34720 OPTIONS IMPORT: reading client specific options from: ccd-tcp/mikrotik
2020-11-17 01:53:40 us=896884 mikrotik/195.111.111.2:34720 MULTI: Learn: 10.101.0.2 -> mikrotik/195.111.111.2:34720
2020-11-17 01:53:40 us=896894 mikrotik/195.111.111.2:34720 MULTI: primary virtual IP for mikrotik/195.111.111.2:34720: 10.101.0.2
2020-11-17 01:53:40 us=896902 mikrotik/195.111.111.2:34720 MULTI: internal route 192.168.47.0/24 -> mikrotik/195.111.111.2:34720
2020-11-17 01:53:40 us=896911 mikrotik/195.111.111.2:34720 MULTI: Learn: 192.168.47.0/24 -> mikrotik/195.111.111.2:34720
2020-11-17 01:53:40 us=896967 mikrotik/195.111.111.2:34720 Outgoing Data Channel: Cipher 'AES-256-CBC' initialized with 256 bit key
2020-11-17 01:53:40 us=896978 mikrotik/195.111.111.2:34720 Outgoing Data Channel: Using 160 bit message hash 'SHA1' for HMAC authentication
2020-11-17 01:53:40 us=896987 mikrotik/195.111.111.2:34720 Incoming Data Channel: Cipher 'AES-256-CBC' initialized with 256 bit key
2020-11-17 01:53:40 us=896995 mikrotik/195.111.111.2:34720 Incoming Data Channel: Using 160 bit message hash 'SHA1' for HMAC authentication
2020-11-17 01:53:40 us=943781 mikrotik/195.111.111.2:34720 TCPv4_SERVER READ [56] from [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ ] pid=4 DATA len$
2020-11-17 01:53:40 us=943836 mikrotik/195.111.111.2:34720 PUSH: Received control message: 'PUSH_REQUEST'
2020-11-17 01:53:40 us=943858 mikrotik/195.111.111.2:34720 SENT CONTROL [mikrotik]: 'PUSH_REPLY,route-gateway 10.101.0.1,topology subnet,ping 10,ping-resta$
2020-11-17 01:53:40 us=943876 mikrotik/195.111.111.2:34720 TCPv4_SERVER WRITE [22] to [AF_INET]195.111.111.2:34720: P_ACK_V1 kid=0 [ 4 ]
2020-11-17 01:53:40 us=943918 mikrotik/195.111.111.2:34720 TCPv4_SERVER WRITE [156] to [AF_INET]195.111.111.2:34720: P_CONTROL_V1 kid=0 [ ] pid=5 DATA len$
2020-11-17 01:53:41 us=251 mikrotik/195.111.111.2:34720 TCPv4_SERVER READ [22] from [AF_INET]195.111.111.2:34720: P_ACK_V1 kid=0 [ 5 ]
2020-11-17 01:53:41 us=76161 mikrotik/195.111.111.2:34720 Connection reset, restarting [0]
2020-11-17 01:53:41 us=76196 mikrotik/195.111.111.2:34720 SIGUSR1[soft,connection-reset] received, client-instance restarting
2020-11-17 01:53:41 us=76272 TCP/UDP: Closing socket
Thanks in advance!
