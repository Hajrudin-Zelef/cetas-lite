---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-318
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [47323, 47511]
sha256: 2ad1d65384dd60c62f374dcdb7c340065a3e9a82133bd06b175b64da88f65a39
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                             6 VPLS Configuration


                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 20 30
                        #
                        mpls lsr-id 2.2.2.9
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif20
                         ip address 8.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif30
                         ip address 9.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 8.1.1.0 0.0.0.255
                          network 9.1.1.0 0.0.0.255
                        #
                        return

                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 30 100 200 300
                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi a2 static
                         pwsignal ldp
                          vsi-id 2
                          peer 1.1.1.9
                         isolate spoken
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 1.1.1.9
                         remote-ip 1.1.1.9
                        #
                        interface Vlanif30
                         ip address 9.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif100
                         l2 binding vsi a2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   761
VPN Configuration
VPN Configuration                                                                             6 VPLS Configuration

                            hub-mode enable
                           #
                           interface Vlanif200
                            l2 binding vsi a2
                           #
                           interface Vlanif300
                            l2 binding vsi a2
                           #
                           interface 10GE1/0/1
                            port link-type trunk
                            port trunk allow-pass vlan 30
                           #
                           interface 10GE1/0/2
                            port link-type trunk
                            port trunk allow-pass vlan 100 200 300
                           #
                           interface LoopBack1
                            ip address 3.3.3.9 255.255.255.255
                           #
                           ospf 1
                            area 0.0.0.0
                             network 3.3.3.9 0.0.0.0
                             network 9.1.1.0 0.0.0.255
                           #
                           return



6.16 Configuring Static BFD for VPLS PW
Prerequisites
                    Before configuring static BFD for VPLS PW, you have completed the following
                    tasks:
                    ●      Configure network layer parameters for devices to communicate.
                    ●      Configure a VPLS single-segment PW.

Context
                    On an MPLS L2VPN, if PWs exist between PEs, you can configure BFD for PW to
                    rapidly detect link faults, so that upper-layer applications can quickly switch to
                    another link if a link encounters a fault.

                            NOTE

                           ● If the status of a PW is down, the BFD session can be established but cannot go up.
                           ● Static BFD for VPLS PW must be configured or cancelled on the PEs at both ends. If this
                             function is configured or cancelled only on the PE at one end, the PW status between
                             both ends will be inconsistent.


Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enable the BFD function globally, and enter the global BFD view.
                    bfd

         Step 3 Return to the system view.
                    quit


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                     762
VPN Configuration
VPN Configuration                                                                             6 VPLS Configuration


         Step 4 Create a BFD session for VPLS PW.
                    bfd session-name bind pw vsi vsi-name peer peer-address [ vc-id vc-id ]

                          NOTE

                        After a BFD session is created, you can run the min-tx-interval, min-rx-interval, and
                        detect-multiplier commands to modify session parameters. For details, see Configuring
                        BFD Parameters.

         Step 5 Configure a discriminator for the BFD session. Perform either of the following
                operations as required:
                    ●    Configure a local discriminator for the static BFD session.
                         discriminator local discr-value
                    ●    Configure a remote discriminator for the static BFD session.
                         discriminator remote discr-value

                          NOTE

                        ● Only static BFD sessions can be configured with local and remote discriminators.
                        ● The local discriminator on the local end must be the same as the remote discriminator
                          on the remote end. If they are different, the BFD session cannot go up.
                        ● The local and remote discriminators configured for a static BFD session can be changed.
                        ● Do not configure the same local and remote discriminators for different BFD sessions on
                          the same device.
                        ● When a BFD session is up, changing the local or remote discriminator will cause the
                          session to enter the administratively down state. The BFD session will recover
                          automatically without manual intervention.

                    ----End

Verifying the Configuration
                    Run the display bfd session command. The command output shows information
                    about the BFD session status, BFD session discriminators, BFD session type, and
                    type of the PW bound to the BFD session.


6.17 Configuring Common VPLS Parameters

6.17.1 Setting Common Parameters for a VSI
Context
                    Common parameters of a VSI include the VPLS encapsulation type, MTU for
                    negotiation, a tunnel policy, and description of the VSI.
                    Perform the following steps on the PEs at both ends of a PW.

