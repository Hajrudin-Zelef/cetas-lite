---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-14
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [984, 1126]
sha256: 2632980bb76e3967e7c2ca0f8c35998a3b566a33a375b251a9997556b32aa3ec
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                    2 GRE Configuration


Benefits
                    The keepalive function can detect the tunnel status and prevent data loss if the
                    remote end becomes unreachable, ensuring reliable data transmission.


2.3 Configuration Precautions for GRE

2.4 Configuring a GRE Tunnel

2.4.1 Configuring a Tunnel Interface

Prerequisites
                    Before configuring a tunnel interface, you have completed the following task:

                    ●   Configure a routing protocol to ensure IP connectivity between devices.


Context
                    A tunnel interface must be configured on each end of a GRE tunnel to be
                    established. You need to set the tunnel encapsulation type of the tunnel interfaces
                    to GRE and specify a source address (or source interface) and a destination
                    address for the interfaces. If the tunnel interfaces need to be advertised using
                    dynamic routing protocols, you also need to configure IP addresses for the tunnel
                    interfaces.

                    When configuring a source interface for a tunnel, do not specify the tunnel
                    interface of this tunnel as the source interface. Instead, specify the tunnel
                    interface of another tunnel.

                    The MTU value is valid only when a local device sends locally originated packets
                    over a GRE tunnel and is invalid when a local device forwards received packets
                    over a GRE tunnel.

                    A tunnel interface is a logical interface and goes down in the following situations:
                    ●   The destination address configured for the tunnel interface is unreachable or
                        is the IP address of the tunnel interface.
                    ●   The source interface configured for the tunnel interface is down.
                    ●   The IP address configured for the tunnel interface is invalid.
                    ●   The keepalive function is configured on the tunnel interface and detects that
                        the tunnel remote end is unreachable.

                    Perform the following steps on the devices at both ends of a GRE tunnel.


Procedure
         Step 1 Enter the system view.
                    system-view


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            11
VPN Configuration
VPN Configuration                                                                              2 GRE Configuration


         Step 2 Create a tunnel interface and enter the tunnel interface view.
                    interface tunnel interface-number

                          NOTE

                         The tunnel interface on a centralized GRE tunnel must be named using the slot ID, subcard
                         ID, and interface number. The tunnel interface's slot ID must be the same as the slot ID of
                         the tunnel where the source interface resides. If the slot IDs are different, the GRE tunnel
                         cannot be established.

         Step 3 (Optional) Configure a description for the tunnel interface.
                    description text

         Step 4 Set the tunnel encapsulation type to GRE.
                    tunnel-protocol gre

                          NOTE

                         Changing or deleting the tunnel mode of a tunnel interface will delete all tunnel-related
                         configurations on the interface, such as the MTU. Exercise caution when performing this
                         operation.

         Step 5 Configure a source address or source interface for the tunnel interface.
                    source { source-ip-address | interface-type interface-number }

         Step 6 Configure a destination address for the tunnel interface.
                    destination ip-address

         Step 7 (Optional) Set the MTU for the tunnel interface.
                    mtu mtu

         Step 8 (Optional) Enable path MTU auto-discovery on the tunnel interface.
                    tunnel pathmtu enable

                    By default, path MTU auto-discovery is disabled.
         Step 9 Configure an IP address for the tunnel interface. By default, no IP address is
                configured for a tunnel interface. If a dynamic routing protocol is used on the
                tunnel interface, configure an IP address for the tunnel interface. The IP addresses
                of tunnel interfaces on both ends of a tunnel must be on the same network
                segment. Perform one of the following operations:
                    ●    Configure an IPv4 address for the tunnel interface.
                         ip address ip-address { mask | mask-length } [ sub ]
                    ●    Configure the tunnel interface to borrow an IPv4 address.
                         ip address unnumbered interface interface-type interface-number
                    ●    Configure an IPv6 address for the tunnel interface.
                         ipv6 enable
                         ipv6 address ipv6-address prefix-length

                    ----End

2.4.2 (Optional) Enabling the Keepalive Function
Context
                    Before you configure a tunnel policy and set the VPN tunnel type to GRE, you
                    must enable the keepalive function. After this function is enabled, the local end
                    will not use a GRE tunnel with an unreachable remote end, preventing data loss.
                    This is because:

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                        12
VPN Configuration
VPN Configuration                                                                     2 GRE Configuration


                    ●    Before the keepalive function is enabled, the local tunnel interface may be up
                         even if the remote end is unreachable.
                    ●    After the keepalive function is enabled on the local end, the local tunnel
                         interface is set to down when the remote end is unreachable. In this case, the
                         local device does not select the unreachable GRE tunnel.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of a tunnel interface.
                    interface tunnel interface-number

         Step 3 Enable the keepalive function of GRE.
                    keepalive [ period period [ retry-times retry-times ] ]

                    As the keepalive function is unidirectional, enable it on both ends of the GRE
                    tunnel.

                    ----End

