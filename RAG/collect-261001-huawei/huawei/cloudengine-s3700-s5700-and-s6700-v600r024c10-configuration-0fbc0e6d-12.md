---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-12
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [1005, 1151]
sha256: 329b4908609399296b7b9ce78c68f831cc5b97a653d44df6273eda95ac4c5a36
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Basic Concepts
                 ●      Tunnel ID
                        The system automatically allocates to each tunnel a tunnel ID, which uniquely
                        identifies a tunnel interface for a specific upper layer application, such as VPN
                        or route management. A tunnel ID is locally significant.
                        A tunnel ID is 32 bits long. The length of each field contained in a tunnel ID
                        varies according to the tunnel type. Figure 2-10 shows the structure of a
                        tunnel ID.

                        Figure 2-10 Structure of a tunnel ID




                        The fields are described as follows:

                        Table 2-2 Fields in a tunnel ID

                         Field                      Description

                         Token                      Index used to search an MPLS LFIB for a specific
                                                    entry

                         Sequence Number            Sequence number of a tunnel ID

                         Slot Number                Slot number of an outbound interface that sends
                                                    packets


                 ●      NHLFE
                        A next hop label forwarding entry (NHLFE) is used to guide MPLS packet
                        forwarding.
                        An NHLFE contains information about the tunnel ID, outbound interface, next
                        hop, outgoing label, label operation, etc.
                 ●      ILM
                        An incoming label map (ILM) entry maps an incoming label to a set of
                        NHLFEs.
                        An ILM entry contains information about the tunnel ID, incoming label,
                        inbound interface, label operation, etc.
                        A transit node creates ILM entries to map labels with NHLFEs. The node can
                        then search the ILM table for all label forwarding information. This process is
                        similar to searching the FIB based on destination IP addresses.
                 ●      FTN
                        An FEC-to-NHLFE (FTN) entry maps FEC to a set of NHLFEs.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               14
MPLS Configuration
MPLS Configuration                                                         2 Basic MPLS Configuration


                        You can obtain FTN information by searching for non-0x0 token values in a
                        FIB. FTN entries are only available on ingress nodes.

2.2.4 MPLS MTU
                 The maximum transmission unit (MTU) size defines the maximum number of
                 bytes that a device can transmit in a packet. It plays an important role when two
                 devices communicate on a network. If the MTU value of a packet exceeds the
                 maximum size supported by a receiver or transit device, the packet is fragmented
                 during transmission, increasing the workload of the network. The packet may even
                 be discarded during transmission, affecting services. To prevent such issues,
                 configure a device (LSR) to calculate the MTU size before sending packets for
                 communication using the following formula:

Fundamentals
                 LDP LSP forwarding and common IP forwarding differ greatly in terms of
                 implementation mechanism but share a large number of similar aspects about the
                 MTU. If MTU values are correctly negotiated before packet transmission, packets
                 can successfully reach the receiver without packet fragmentation and reassembly.

                 For a FEC, an LSR calculates the smallest value among all MTU values advertised
                 by preferred next-hop LSRs as well as the MTU value of the local outbound
                 interface. The LSR then adds the calculated MTU value to the MTU TLV of Label
                 Mapping messages to be sent to the upstream device. When the MTU value
                 changes (for example, when the local outbound interface or the MTU
                 configuration changes), the LSR recalculates the MTU value and advertises the
                 new value to its upstream device through Label Mapping messages.


2.3 Configuration Precautions for MPLS Basics

2.4 Configuring the MPLS MTU
Context
                 The MTU value determines the maximum size (in bytes) of a data packet that can
                 be transmitted across a data link. If the size of a packet exceeds the maximum
                 size supported by a receiver or transit device, the packet is fragmented during
                 transmission, increasing the workload of the network. In some cases, the packet
                 may even be discarded during transmission, affecting services. To prevent such
                 issues, configure a device (LSR) to calculate the MTU size before sending packets
                 for communication using the following formula:

                 LDP MTU value = Min { All MTU values advertised by downstream devices, MTU
                 value of the local outbound interface}. You then need to configure the device to
                 add the calculated LDP MTU value to the MTU TLV in Label Mapping messages,
                 so that the value is advertised to the upstream device when those messages are
                 sent upstream. When the MTU value changes (for example, when the local
                 outbound interface or the MTU configuration changes), the LSR recalculates the
                 MTU value and advertises the new value to its upstream device through Label

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            15
MPLS Configuration
MPLS Configuration                                                                2 Basic MPLS Configuration


                 Mapping messages. The MTU size of a local outbound interface is determined as
                 follows:
                 ●      If the MPLS MTU value is not configured on the interface, the interface MTU
                        value is used.
                 ●      If both the MPLS MTU value and interface MTU value are configured on the
                        interface, the smaller value of the two values is used.
                 In this manner, MPLS determines the size of MPLS packets to be forwarded on the
                 ingress based on the LDP MTU value. This prevents large packets from being sent
                 on the ingress and prevents forwarding failures on transit nodes.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Configure the device to add the MTU TLV of a specified type in Label Mapping
                messages to be sent.
                 mtu-signalling [ apply-tlv ]

         Step 4 Return to the system view.
                 quit

         Step 5 Enter the view of an MPLS-enabled interface.
                 interface interface-type interface-number

         Step 6 Switch the interface working mode to Layer 3.
                 undo portswitch

                 Determine whether to perform this step based on the current interface working
                 mode.

                         NOTE

                        If multiple interfaces need to be switched to the Layer 3 mode, run the undo portswitch
                        batch interface-type { interface-number1 [ to interface-number2 ] } &<1-10> command in
                        the system view to switch these interfaces to the Layer 3 mode in batches.

         Step 7 Configure the MPLS MTU for the interface.
                 mpls mtu mtu

