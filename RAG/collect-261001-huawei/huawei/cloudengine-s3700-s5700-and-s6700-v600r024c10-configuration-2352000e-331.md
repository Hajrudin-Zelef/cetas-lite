---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-331
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [49495, 49638]
sha256: 6e007b0995a760e402a25b6ec4ea271149409ea1d23d179f048dad9ed089e58a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

6.19.5 Using Ping to Test PW Connectivity on a VPLS Network
Prerequisites
                    Before testing PW connectivity using the ping vpls command, ensure that the
                    virtual private LAN service (VPLS) network has been configured correctly.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                           794
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration


Context
                    VPLS PW ping can only use the label-alert test mode to enable the destination
                    device to forward packets to the CPU for processing, instead of forwarding them
                    to other devices.

                    Figure 6-47 Typical VPLS network




                    On the network shown in Figure 6-47, the VPLS PW ping process is as follows:
                    1.   An ingress node checks whether the PW exists based on the Virtual Switching
                         Instance (VSI) name, peer IP address, and PW ID. If not, an error message is
                         displayed, and the ping process ends.
                    2.   The ingress node constructs an MPLS Echo Request message, with destination
                         IP address 127.0.0.1/8 and TTL value 1 in the IP header. It then searches for
                         the corresponding LSP, adds an LSP label to the MPLS Echo Request message,
                         and forwards the message.
                    3.   The MPLS Echo Request message is forwarded along the LSP to the egress
                         node, which replies with an MPLS Echo Reply message.

Procedure
                    ●    LDP VPLS:
                         ping vpls [ -c echo-number | -m time-value | -s data-bytes | -t timeout-value | -r reply-mode | -exp
                         exp-value | -v | -g ] * vsi vsi-name peer peer-address [ negotiate-vc-id vc-id ]
                    ●    BGP VPLS:
                         ping vpls [ -c echo-number | -m time-value | -s data-bytes | -t timeout-value | -r reply-mode | -exp
                         exp-value | -v | -g ] * vsi vsi-name local-site-id remote-site-id

                    ----End

Example
                    Perform a ping operation to test LDP VPLS PW connectivity.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                               795
VPN Configuration
VPN Configuration                                                                                6 VPLS Configuration

                    <HUAWEI> ping vpls vsi a2 peer 10.1.1.1
                       PW PING : FEC 128 PSEUDOWIRE (NEW). Type = vlan, ID = 2 : 100 data bytes, press CTRL_C to break
                       Reply from 10.1.1.1: bytes=100 Sequence=1 time=60 ms
                       Reply from 10.1.1.1: bytes=100 Sequence=2 time=50 ms
                       Reply from 10.1.1.1: bytes=100 Sequence=3 time=60 ms
                       Reply from 10.1.1.1: bytes=100 Sequence=4 time=60 ms
                       Reply from 10.1.1.1: bytes=100 Sequence=5 time=60 ms
                    - -- FEC: FEC 128 PSEUDOWIRE (NEW). Type = vlan, ID = 2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 50/58/60 ms

                    The command output contains the following information:
                    ●    Response to each ping packet: If no response packet is received within the
                         timeout period, the message "Request time out" is displayed. If a response
                         packet is received, the number of data bytes, packet sequence number, TTL
                         value, and response time carried in the packet are displayed.
                    ●    Statistics about ping packets: These include the number of sent packets,
                         number of received packets, percentage of the packets that are not replied,
                         and the minimum, maximum and average response time.

6.19.6 Using Tracert to Test a PW Path on a VPLS Network
Prerequisites
                    Before testing PW connectivity using the tracert vpls command, ensure that the
                    VPLS network has been configured correctly.

Context
                    VPLS PW tracert can only use the label-alert test mode to enable the destination
                    device to forward packets to the CPU for processing, instead of forwarding them
                    to other devices.

                    Figure 6-48 Typical VPLS network




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                             796
VPN Configuration
VPN Configuration                                                                                          6 VPLS Configuration


                    On the network shown in Figure 6-48, the VPLS PW tracert process is as follows:
                    1.   An ingress node checks whether the PW exists based on the VSI name, peer IP
                         address, and PW ID. If not, an error message is displayed, and the tracert
                         process ends.
                    2.   The ingress node constructs an MPLS Echo Request message, with destination
                         IP address 127.0.0.1/8 and TTL value 1 in the IP header. It then searches for
                         the corresponding LSP, adds an LSP label to the MPLS Echo Request message,
                         and forwards the message.
                    3.   The MPLS Echo Request message is forwarded along the LSP to the egress
                         node, which replies with an MPLS Echo Reply message.
                    4.   After receiving the MPLS Echo Reply message, the ingress node determines
                         that the MPLS Echo Request message has reached the egress node and then
                         terminates the tracert operation.

Procedure
                    ●    LDP VPLS:
                         tracert vpls [ -exp exp-value | -f first-ttl | -m max-ttl | -r reply-mode | -t timeout-value | -g ] * vsi vsi-
                         name peer peer-address [ negotiate-vc-id vc-id ] [ full-lsp-path ] [ detail ]
                    ●    BGP VPLS:
                         tracert vpls [ -exp exp-value | -f first-ttl | -m max-ttl | -r reply-mode | -t timeout-value | -g ] * vsi vsi-
                         name local-site-id remote-site-id [ full-lsp-path ] [ detail ]

                    ----End

Example
                    Perform a tracert operation to test a PW path on a BGP VPLS network.
                    <HUAWEI> tracert vpls vsi test 10 10 full-lsp-path
                     PW Trace Route FEC: L2 VPN ENDPOINT. Sender VEID = 10, Remote VEID = 20, press CTRL_C to break
                    TTL Replier         Time Type       Downstream
                    0                      Ingress 10.1.1.2/[294929 32894 32888 ]
                    1    10.1.1.2      93 ms Transit 10.2.1.2/[32925 3 ]
                    2    10.2.1.2      1 ms Transit 10.3.1.2/[32881 ]
                    3    4.4.4.4      2 ms Egress

                    The preceding command output shows each node on the PW path and the
                    response time of each node.

6.19.7 Resetting BGP Connections Related to BGP L2VPN AD
Context
                    After the BGP configuration related to BGP L2VPN AD is changed, you can run
                    related commands in the user view to make the new configuration take effect
                    immediately.




Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                                      797

