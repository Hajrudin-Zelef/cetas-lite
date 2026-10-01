---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-262
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [38405, 38546]
sha256: 5c291852bb11312d6e24cbe415d5df1385dae4d051b83ace26475eb2f8f9244e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                                   ● Flow label-based load balancing can be enabled only when any of the
                                     following conditions is true:
                                       ●     The receive parameter is configured on the local end, and the send
                                             parameter is configured on the remote end.
                                       ●     The send parameter is configured on the local end, and the receive
                                             parameter is configured on the remote end.
                                       ●     Both the send and receive parameters are configured on the local and
                                             remote ends.
                                   ● If static is configured, flow label-based load balancing is statically configured.
                                       If the static flow label-based load balancing configuration does not match on
                                       both ends, the device discards packets with flow labels, causing packet loss.
                                   ● Configuring flow labels will cause PWs to be renegotiated. As a result,
                                     services go down and then up.

                    ----End

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        614
VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


6.6.5 (Optional) Configuring a Huawei Device to
Communicate with a Non-Huawei Device

Context
                    When a Huawei device interworks with a non-Huawei device, configure LDP VPLS
                    parameters based on the non-Huawei device's configuration.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the MPLS L2VPN view.
                    mpls l2vpn

         Step 3 Disable the device from sending Notification messages over LDP PWs.
                    mpls l2vpn default martini

                    By default, a device sends Notification messages over LDP PWs.

                    Before running this command, delete the configured VPLS service.

                    ----End

6.6.6 Verifying the Configuration

Procedure
                    ●    Run the display vsi [ name vsi-name ] [ verbose ] command to check VSI
                         information.
                    ●    Run the display vsi services { all | vsi-name | interface interface-type
                         interface-number | vlan vlan-id | interface interface-name } command to
                         check information about the AC interface associated with the VSI.
                    ●    Run the display mpls label-stack vpls vsi vsi-name peer peer-id vc-id vc-id
                         command to check label stack information in VPLS scenarios.
                    ●    Run the display vsi remote ldp [ [ router-id ip-address ] [ pw-id pw-id ] |
                         [ verbose ] | [ unmatch ] ] command to check remote VSI information.
                    ●    Run the display vpls connection [ ldp | vsi vsi-name ] [ down | up ]
                         [ verbose ] command to check LDP VPLS connections.
                    ●    Run the display vsi pw out-interface [ vsi vsi-name ] command to check the
                         outbound interface information of VSI PWs.
                    ●    Run the display l2vpn vsi-list tunnel-policy policy-name command to check
                         information about the tunnel policies applied to VSIs.
                    ●    Run the display vpls forwarding-info [ vsi vsi-name [ peer peer-address
                         [ negotiation-vc-id vc-id | remote-site site-id ] ] | state { up | down } ]
                         [ verbose ] command to check VSI forwarding information.
                    ●    Run the display admin-vsi binding [ admin-vsi vsi-name ] command to
                         check the binding between the management VSI and service VSIs.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            615
VPN Configuration
VPN Configuration                                                                            6 VPLS Configuration


                    ●    Run the display vsi { name vsi-name peer-info [ peer-ip-address ] | peer-
                         info } command to check the peer PW status.
                    ----End

6.6.7 Example for Configuring LDP VPLS
Networking Requirements
                    Figure 6-10 shows a backbone network built by an enterprise. The enterprise has
                    a small number of branch sites, only two of which are shown in this example.
                    Site1 connects to the backbone network by connecting CE1 to PE1, while Site2
                    connects to the backbone network by connecting CE2 to PE2. Users at Site1 and
                    Site2 need to communicate at Layer 2, and user information needs to be retained
                    in Layer 2 packets when the packets are transmitted over the backbone network.

                    Figure 6-10 Network diagram of configuring LDP VPLS
                          NOTE

                         In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




Precautions
                    During the configuration, note the following:
                    ●    PEs on the same L2VPN must be configured with the same VSI ID.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VPLS to transparently transmit Layer 2 packets on the backbone
                         network to implement Layer 2 communication between Site1 and Site2 and
                         to retain user information in Layer 2 packets when the packets are
                         transmitted over the backbone network.
                    2.   Configure LDP VPLS to implement Layer 2 communication between CEs
                         because the enterprise network has a small number of sites.
                    3.   Configure an IGP on the backbone network for data transmission between
                         PEs on the public network.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     616
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


                    4.   Configure basic MPLS functions and LDP on devices on the backbone network
                         to implement VPLS.
                    5.   Establish tunnels for transmitting data between PEs to prevent data from
                         being accessed by the public network.
                    6.   Enable MPLS L2VPN on PEs to implement VPLS.
                    7.   Create VSIs on PEs, configure LDP as the signaling protocol, and bind VSIs to
                         AC interfaces to implement LDP VPLS.

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

