---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-297
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [43899, 44062]
sha256: 65d75d7c23aae064860c9c275a83d0954b4e5889739a85b3d6d0297a0ab68e85
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

         Step 5 Configure VSI peer relationships between SPEs and UPEs.
                    peer peer-address [ negotiation-vc-id vc-id ] [ tnl-policy policy-name ] upe

         Step 6 Create a PW and enter its view.
                    peer peer-address [ negotiation-vc-id vc-id ] pw pw-name


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                      704
VPN Configuration
VPN Configuration                                                                       6 VPLS Configuration


         Step 7 Associate spoke PW status with hub PW status.
                    track hub-pw

                    By default, spoke PW status is not associated with hub PW status.

                    This command applies only to spoke PWs. You must perform step 5 to create a
                    spoke PW first.

                    After this command is run, SPE1 instructs the UPE to switch traffic to the
                    secondary spoke PW after detecting that all connected hub PWs are down. The
                    protect-group group-name and protect-mode pw-redundancy master
                    commands must have been run to create a PW protection group in master/slave
                    PW redundancy mode, and the primary and secondary PWs have joined the group.

                    ----End

6.11.5 (Optional) Manually Switching Traffic Between PWs in
a PW Protection Group

Context
                    If you want to maintain the device with the primary PW configured in a VPLS PW
                    redundancy scenario, you can switch traffic from the primary PW to the secondary
                    PW, and switch traffic back to the primary PW after the device stabilizes.

                    Perform the following steps on a UPE:

                          NOTE

                        Only the PWs that work in master/slave PW redundancy mode support manual switching.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the created VSI.
                    vsi vsi-name

         Step 3 Enter the VSI-LDP view.
                    pwsignal ldp

         Step 4 Create a PW protection group and enter its view.
                    protect-group group-name

         Step 5 Perform primary/secondary PW switchover in the PW protection group.
                    protect-switch manual

         Step 6 Perform primary/secondary PW switchback in the PW protection group.
                    protect-switch clear

                          NOTE

                        The interval between PW switchover and switchback must be longer than 15s.

                    ----End

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                             705
VPN Configuration
VPN Configuration                                                                6 VPLS Configuration


6.11.6 Verifying the Configuration
Procedure
                    ●   Run the display vsi [ name vsi-name ] [ verbose ] command to check VSI
                        information.
                    ●   Run the display vsi name vsi-name protect-group [ group-name [ verbose |
                        history ] ] command to check information about the PW protection group of
                        a specified VSI.
                    ----End


6.12 Configuring VPLS Interworking
Prerequisites
                    Before configuring VPLS interworking, you have completed the following tasks:
                    ●   Configure interface IP addresses and routes on PEs and SPEs to ensure IP
                        connectivity.
                    ●   Configure LSR IDs and enable MPLS on PEs and SPEs.
                    ●   Establish tunnels between PEs and between PEs and SPEs to transmit service
                        traffic.

6.12.1 Understanding VPLS Interworking
Context
                    LDP VPLS uses LDP signaling messages that carry the FEC 128 TLV, requires PWs
                    to be manually configured, and has low device performance requirements. BGP AD
                    VPLS uses BGP to automatically discover VPLS members and uses LDP signaling
                    messages that carry the FEC 129 TLV to automatically establish PWs. LDP VPLS
                    applies to VPLS networks with a small number of sites, whereas BGP AD VPLS
                    applies to VPLS networks requiring a large number of PWs. As networks expand,
                    LDP VPLS networks need to be connected to BGP AD VPLS networks. Therefore,
                    interworking between LDP VPLS and BGP AD VPLS networks becomes an urgent
                    issue. Interworking between LDP VPLS and BGP AD VPLS has been developed to
                    meet this requirement, which allows for seamless interconnection between LDP
                    and BGP AD VPLS networks.

Implementation
                    To enable an LDP VPLS network to communicate with a BGP AD VPLS network,
                    the edge nodes between the two networks must support both LDP VPLS and BGP
                    AD VPLS. In Figure 6-28, PE1 supports LDP VPLS, PE3 supports BGP AD VPLS, and
                    the edge nodes PE2 and PE4 support both LDP VPLS and BGP AD VPLS. PE1, PE2,
                    and PE4 form the LDP VPLS network, while PE3, PE2, and PE4 form the BGP AD
                    VPLS network. LDP VPLS uses LDP signaling messages carrying the FEC 128 TLV to
                    establish and maintain PWs. BGP AD VPLS uses BGP to automatically discover
                    neighbors and uses LDP signaling messages carrying the FEC 129 TLV to establish
                    and maintain PWs. Signaling negotiation for LDP VPLS is independent of that for

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           706
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                    BGP AD VPLS. For details about PW establishment and maintenance in LDP VPLS,
                    see 6.6 Configuring LDP VPLS. For details about member discovery and PW
                    establishment and maintenance in BGP AD VPLS, see 6.8 Configuring BGP AD
                    VPLS. After PWs are established, PEs can exchange data packets over these PWs.
                    Data packet encapsulation on an LDP VPLS network is similar to that on a BGP AD
                    VPLS network. For the packet encapsulation process on PWs, see VPLS
                    Encapsulation Types.

                    Figure 6-28 Interworking between LDP VPLS and BGP AD VPLS




6.12.2 Configuring Interworking Between LDP VPLS and BGP
AD VPLS
Context
                    On the network shown in Figure 6-29, an LDP VPLS network is deployed among
                    PE1, PE2, and PE3, and a BGP AD VPLS network is deployed among PE2, PE3, and
                    PE4. To configure interworking between LDP VPLS and BGP AD VPLS, configure
                    hybrid VSIs on PE2 and PE3 (edge nodes). In this scenario, the hybrid VSIs on PE2
                    and PE3 establish LDP PWs with the VSI on PE1 and establish BGP AD PWs with
                    the VSI on PE4.

                    Figure 6-29 Network diagram of configuring interworking between LDP VPLS and
                    BGP AD VPLS




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         707
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


Procedure
         Step 1 Establish LDP PWs. The configuration is similar to that in "Configuring LDP VPLS."
                For details, see 6.6.2 Creating a VSI and Configuring LDP Signaling.

