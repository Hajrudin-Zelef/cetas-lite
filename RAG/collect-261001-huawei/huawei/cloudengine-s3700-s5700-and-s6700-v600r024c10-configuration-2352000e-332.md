---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-332
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "ethernet"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [49639, 49787]
sha256: a3549d866d2f3a1af64ae2b39770332dac4d6923943f52db4cbc933e8dc614f0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                   6 VPLS Configuration


                    Table 6-15 Resetting BGP connections related to BGP L2VPN AD

                     Operation                  Command

                     Reset all BGP              reset bgp l2vpn-ad all
                     connections related
                     to BGP L2VPN AD.

                     Reset all EBGP             reset bgp l2vpn-ad external
                     connections.

                     Reset all IBGP             reset bgp l2vpn-ad internal
                     connections.

                     Reset the connection       reset bgp l2vpn-ad { as-number-dot | ipv4-address }
                     of a specified BGP
                     peer.

                     Reset BGP                  reset bgp l2vpn-ad group group-name
                     connections of a
                     specified peer group.




6.20 Troubleshooting VPLS

6.20.1 VPLS Interconnection Fails
Fault Symptom
                    When a Huawei device is connected to a non-Huawei device using VPLS, the VPLS
                    status is not up.

Procedure
                    ●   Connecting to a Juniper device
                        When a Huawei device is connected to a Juniper device using BGP VPLS, BGP
                        peers are down. As a result, no VPLS PW can be established. To rectify the
                        fault, perform the following operations:
                        a.   Configure the L2VPN address family as a unicast address family on the
                             Juniper device.
                        b.   Set the default-offset to 1 on the Huawei device. The CE ID cannot be 0.
                    ●   Connecting to a Cisco device
                        The latest RFC defines that the PW encapsulation type for BGP VPLS is 19.
                        Huawei devices support only Ethernet encapsulation and VLAN encapsulation.
                        When a Huawei device connects to a Cisco device using VPLS and the VPLS
                        status is not up, perform the following operations:
                        a.   Configure the encapsulation type of BGP VPLS packets to comply with
                             related standards.
                             encapsulation rfc4761-compatible


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            798
VPN Configuration
VPN Configuration                                                                           6 VPLS Configuration


                               The device is enabled to convert the VPLS encapsulation type to 19 when
                               sending VPLS packets, provided that the current encapsulation type is not
                               19.

                               When receiving VPLS packets with the encapsulation type 19, the device
                               automatically converts the encapsulation type according to the VPLS
                               encapsulation type on the link.
                         b.    Disable MTU check for VSIs on PEs.
                               mtu-negotiate disable

                               By default, the MTU value for a VSI is 1500 bytes. On a BGP VPLS
                               network, if a Huawei device communicates with a non-Huawei device
                               and the MTU values for the same VSI on the two devices are different,
                               the two devices cannot establish a PW.

                               Some non-Huawei devices do not support MTU check for VSIs. If a
                               Huawei device needs to communicate with a non-Huawei device through
                               BGP VPLS, run the mtu-negotiate disable command to disable MTU
                               check on the Huawei device.

                    ----End

6.20.2 The VSI State Is Up on Only One End

Fault Symptom
                    After VPLS is configured, the VSI is up on only one end.


Procedure
         Step 1 Check whether multiple AC interfaces on the local end are bound to the VSI.
                    display vsi [ name vsi-name ] [ verbose ]

                    If fewer than two interfaces are bound to the VSI, you need to configure the local
                    and remote ends to bind more than two AC interfaces to the VSI.

                          NOTE

                        If two or more AC interfaces are bound to the VSI, the VSI can be up.

         Step 2 Check whether the remote end specifies the local end as a UPE.

                    If the remote end specifies the local end as a UPE, the remote end does not
                    instruct the local end to withdraw labels after the remote AC interface is faulty.
                    This may cause the VSI to be up on only one end.

                    ----End

6.20.3 An LDP VPLS VSI Cannot Go Up

Fault Symptom
                    After LDP VPLS is configured, a VSI cannot go up.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                799
VPN Configuration
VPN Configuration                                                                                6 VPLS Configuration


Procedure
         Step 1 Check whether the encapsulation types on both ends are the same.
                    <HUAWEI> display vsi name tt
                    Vsi                    Mem PW Mac               Encap      Mtu Vsi
                    Name                      Disc Type Learn Type             Value State
                    --------------------------------------------------------------------------
                    tt                    static ldp unqualify vlan        1500 up

                    ●     If the encapsulation types on both ends are different, run the encapsulation
                          { ethernet | vlan } command in the VSI view to change the encapsulation
                          type on one end to be the same as that on the other end.
                    ●     If the encapsulation types on both ends are the same, go to step 2.
                           NOTE

                         A VSI can be up only when the encapsulation types configured on both ends are the same.

         Step 2 Check whether the MTU values on both ends are the same.
                    <HUAWEI> display vsi name tt
                    Vsi                    Mem PW Mac               Encap      Mtu Vsi
                    Name                      Disc Type Learn Type             Value State
                    --------------------------------------------------------------------------
                    tt                    static ldp unqualify vlan        1500 up

                    ●     If the MTU values on both ends are different, run the mtu mtu-value
                          command in the VSI view to change the MTU value on one end to be the
                          same as that on the other end.
                    ●     If the MTU values on both ends are the same, go to step 3.
                           NOTE

                         A VSI can be up only when the MTU values configured for the two ends are the same.

