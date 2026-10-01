---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-98
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["compute", "copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [13777, 13936]
sha256: 78e0ec7aca0d63048bc2283cc6b4ae25c7c5c9dc4b8554c3880ef91b1b14691c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 CSPF provides a method of selecting a path in the MPLS domain. By default, the
                 optimization mode is used to compute a path from the egress to the ingress.
                 Compared with common path computation methods, the optimization mode
                 offers higher efficiency.
                 After this command is run, a path is calculated from the ingress to the egress.
         Step 8 (Optional) Configure the CSPF tie-breaking mode using either of the following
                methods:
                 A CSPF tie-breaking mode can be configured globally or for a specific tunnel. The
                 tunnel-specific tie-breaking configuration takes preference over the global
                 configuration, and the global configuration is used by a tunnel only if no tie-
                 breaking mode is configured for this tunnel.
                 ●      Configure a CSPF tie-breaking mode for a tunnel.
                        quit
                        interface tunnel tunnel-number
                        mpls te tie-breaking { least-fill | most-fill | random }

                 ●      Configure a CSPF tie-breaking mode for a node.
                        mpls te tie-breaking { least-fill | most-fill | random }

                        The default tie-breaking mode is random.

                 ----End

4.7.4 Configuring MPLS TE Tunnel Interfaces
Context
                 To set up an MPLS TE tunnel and customize tunnel attributes, you need to first
                 configure a tunnel interface. MPLS TE tunnel interface plays a crucial role in
                 setting up and maintaining an MPLS TE tunnel, as well as controlling how packets
                 are forwarded across the tunnel.

                         NOTE

                        MPLS TE tunnels forward MPLS packets, not IP packets. IP forwarding-related commands
                        are ineffective on tunnel interfaces.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       234
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                 Perform the following configuration on the ingress of an MPLS TE tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Create a tunnel interface and enter the tunnel interface view.
                 interface tunnel interface-number

         Step 3 Use one of the following methods to assign an IP address to the tunnel interface.
                 ●      Configure an IP address for the tunnel interface.
                        ip address ip-address { mask | mask-length } [ sub ]

                        Configure a primary IP address before you can add a secondary IP address for
                        a tunnel interface.
                 ●      Configure the tunnel interface to borrow the IP address of another interface.
                        ip address unnumbered interface interface-type interface-number

                              NOTE

                             A TE tunnel can be established on a tunnel interface without an IP address. However,
                             an IP address must be configured for the tunnel interface before it can forward traffic
                             over the TE tunnel. An MPLS TE tunnel is unidirectional and does not involve peer
                             address configuration. Therefore, you are advised to specify the ingress LSR ID as the
                             IP address of the tunnel interface, instead of configuring a unique IP address for the
                             interface.

         Step 4 Configure MPLS TE as the tunneling protocol.
                 tunnel-protocol mpls te

         Step 5 Configure the destination address for the tunnel.
                 destination ip-address

                 Generally, the egress LSR ID is used as the tunnel destination address. Different
                 tunnel types require specific destination IP addresses. Therefore, changing the
                 tunneling protocol to MPLS TE automatically removes the original tunnel
                 destination IP addresses. As a result, you need to reconfigure the destination
                 addresses for the tunnels.
         Step 6 Configure a tunnel ID.
                 mpls te tunnel-id tunnel-id

         Step 7 Configure RSVP-TE as the signaling protocol.
                 mpls te signal-protocol rsvp-te

         Step 8 (Optional) Configure bandwidth for the tunnel.
                 mpls te bandwidth ct0 ct0-value

                 Ensure that the bandwidth used by the tunnel does not exceed the maximum
                 reservable link bandwidth. Ignore this step if the TE tunnel is used only for
                 changing the data transmission path.
         Step 9 (Optional) Configure a tunnel name.
                 mpls te signalled tunnel-name signalled-tunnel-name

                 Generally, a tunnel interface name is used as a tunnel name. For example,
                 Tunnel10 can be used as both a tunnel interface name and a tunnel name. This
                 step is performed for the following purposes:

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      235
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration


                 ●      Facilitate tunnel management.
                 ●      Allow the device to interwork with a non-Huawei device that does not uses a
                        tunnel interface name as the tunnel name.
        Step 10 (Optional) Configure the route and label recording function.
                 mpls te record-route [ label ]

        Step 11 (Optional) Disable CSPF calculation during TE tunnel establishment.
                 mpls te cspf disable

                         NOTE

                        This command needs to be run only in inter-AS VPN Option C scenarios, and it is not
                        recommended in other scenarios.

                 ----End

4.7.5 (Optional) Disabling TE LSP Flapping Suppression
Context
                 TE LSP flapping suppression prevents high CPU usage stemming from TE LSP
                 flapping. This function can be disabled.
                 Perform the following configuration on the ingress of an MPLS TE tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enable MPLS on the local node and enter the MPLS view.
                 mpls

         Step 3 Enable MPLS TE globally.
                 mpls te

         Step 4 Disable TE LSP flapping suppression.
                 mpls te suppress-flapping disable

                 ----End

4.7.6 Verifying the Configuration
Procedure
                 ●      Run the display interface tunnel [ interface-number ] command to check the
                        operating status of tunnel interfaces.
                 ●      Run the display mpls lsp command to check CR-LSP information.
                 ●      Run the display mpls te tunnel [ verbose ] command to check MPLS TE
                        tunnel information.
                 ●      Run the display mpls te tunnel statistics command to check MPLS TE tunnel
                        statistics.
                 ●      Run the display mpls te tunnel-interface command to check tunnel
                        interface information on the local node.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                    236
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


