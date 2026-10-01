---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-269
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [39501, 39649]
sha256: fabcaabfb60a309c8b4fb39d4e1ab7c16ba987667584acc7ca060efd9f7f68cb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 2 Enter the VSI view.
                    vsi vsi-name

         Step 3 Enter the VSI-BGP view.
                    pwsignal bgp

         Step 4 Enable flow label-based load balancing for the VSI.
                    flow-label { both | send | receive } [ static ]

                    By default, flow label-based load balancing is disabled for a VSI.

                          NOTE

                         ● Flow label-based load balancing can be enabled only when any of the following
                           conditions is true:
                             –     The receive parameter is configured on the local end, and the send parameter is
                                   configured on the remote end.
                             –     The send parameter is configured on the local end, and the receive parameter is
                                   configured on the remote end.
                             –     Both the send and receive parameters are configured on the local and remote
                                   ends.
                         ● If static is configured, flow label-based load balancing is statically configured.
                             If the static flow label-based load balancing configuration does not match on both ends,
                             the device discards packets with flow labels, causing packet loss.

                    ----End

6.7.7 (Optional) Configuring a Huawei Device to
Communicate with a Non-Huawei Device
Context
                    To enable a Huawei device to communicate with a non-Huawei device, disable
                    MTU check for a VSI.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the VSI view.
                    vsi vsi-name

         Step 3 Enter the VSI-BGP view.
                    pwsignal bgp

         Step 4 Configure the encapsulation type of BGP VPLS packets to comply with related
                standards.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                    633
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration

                    encapsulation rfc4761-compatible

                    By default, new VRPv8-based Huawei devices implement BGP VPLS in compliance
                    with industry standards, with 19 used as the encapsulation type of VPLS packets.
                    Existing VRPv5-based Huawei devices implement BGP VPLS based on the Huawei-
                    proprietary encapsulation types, with 4 used as the VLAN encapsulation type and
                    5 as the Ethernet encapsulation type. For existing VRPv5-based Huawei devices,
                    you can run the encapsulation rfc4761-compatible command on the devices to
                    configure the encapsulation type of BGP VPLS packets to comply with industry
                    standards. For new VRPv8-based Huawei devices, this command is meaningless
                    because it is reserved only for upgrade compatibility.
         Step 5 Disable MTU check for the VSI.
                    mtu-negotiate disable

                    By default, the MTU value for a VSI is 1500 bytes. On a BGP VPLS network, if a
                    Huawei device communicates with a non-Huawei device and the MTU values for
                    the same VSI on the two devices are different, the two devices cannot establish a
                    PW. After the mtu-negotiate disable command is run on a PE, the PE does not
                    perform any MTU check on VPLS packets received from the same VSI. If other
                    conditions for establishing a PW are met, the PW goes up.

                    ----End

6.7.8 Verifying the Configuration
Procedure
                    ●    Run the display vpls connection [ bgp | vsi vsi-name ] [ down | up ]
                         [ verbose ] command to check BGP VPLS connections.
                    ●    Run the display vsi remote bgp [ nexthop nexthop-address { export-vpn-
                         target vpn-target ] | route-distinguisher | route-distinguisher } command
                         to check remote VSI information.
                    ●    Run the display bgp l2vpn-ad routing-table vpls command to check VPLS
                         route information in the L2VPN AD address family.
                    ----End

6.7.9 Example for Configuring BGP VPLS
Networking Requirements
                    Figure 6-15 shows a backbone network built by an enterprise. The enterprise has
                    a large number of branch sites, only two of which are shown in this example. The
                    network environment often changes. Site1 connects to the backbone network by
                    connecting CE1 to PE1, while Site2 connects to the backbone network by
                    connecting CE2 to PE2. Users at Site1 and Site2 need to communicate at Layer 2,
                    and user information needs to be retained in Layer 2 packets when the packets
                    are transmitted over the backbone network.

                    Figure 6-15 Network diagram of configuring BGP VPLS
                         NOTE

                        In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    634
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration




Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure VPLS to transparently transmit Layer 2 packets on the backbone
                         network to implement Layer 2 communication between Site1 and Site2 and
                         to retain user information in Layer 2 packets when the packets are
                         transmitted over the backbone network.
                    2.   Use BGP VPLS to implement Layer 2 communication between CEs on the
                         enterprise network with many sites and a complex network environment.
                    3.   Configure an IGP on the backbone network for data transmission between
                         PEs on the public network.
                    4.   Configure basic MPLS functions and LDP on devices on the backbone network
                         to implement VPLS.
                    5.   Establish tunnels for transmitting data between PEs to prevent data from
                         being accessed by the public network.
                    6.   Enable MPLS L2VPN on PEs to implement VPLS.
                    7.   Enable PEs to function as BGP peers to exchange VPLS information, create a
                         VSI on PEs, specify BGP as the signaling protocol, specify the RD, VPN target,
                         and site ID, and bind AC interfaces to the VSI to implement BGP VPLS.

Procedure
         Step 1 Configure VLANs to which interfaces belong and assign IP addresses to the
                corresponding VLANIF interfaces.

                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan 10
                    [CE1-vlan10] quit
                    [CE1] interface vlanif 10
                    [CE1-Vlanif10] ip address 10.1.1.1 255.255.255.0
                    [CE1-Vlanif10] quit
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/1] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         635

