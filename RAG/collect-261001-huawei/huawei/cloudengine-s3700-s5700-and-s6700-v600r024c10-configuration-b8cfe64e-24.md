---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-b8cfe64e-24
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e.md
source_anchor: ""
source_lines: [3059, 3193]
sha256: 6ded9ae9736cf3549490ed9b4b097885451c29aa0c148d07a2aa1e0c386122a3
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--b8cfe64e

Procedure
         Step 1 Configure a traffic classifier.
                    For details about how to configure a traffic classifier, see 3.4 Configuring a Traffic
                    Classifier in "MQC Configuration".
         Step 2 Configure a traffic behavior.
                    1.   Enter the system view.
                         system-view
                    2.   Create a traffic behavior and enter the traffic behavior view, or enter the view
                         of an existing traffic behavior.
                         traffic behavior behavior-name
                    3.   Configure a redirection action.
                         –    Redirection to an interface
                              redirect interface interface-type interface-number [ fail-action forward ]


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                         54
QoS Configuration
QoS Configuration                                                                             7 Redirection Configuration


                                      NOTE

                                     Generally, if a packet needs to be redirected to the outbound interface, the
                                     outbound interface needs to be added to the VLAN corresponding to the packet.
                         –      Redirection to a single next-hop IP address
                                redirect [ vpn-instance vpn-instance-name ] nexthop { ip-address [ track nqa admin-name
                                test-name [ reaction probe-failtimes fail-times ] ] } &<1-16> [ fail-action discard ] [ low-
                                precedence ]
                                redirect ipv6 [ vpn-instance vpn-instance-name ] nexthop { ipv6-address [ track nqa admin-
                                name test-name [ reaction probe-failtimes fail-times ] ] } &<1-16> [ fail-action discard ]
                         –      Redirection to multiple next-hop IP addresses
                                redirect load-balance [ vpn-instance vpn-instance-name ] nexthop { ip-address [ track nqa
                                admin-name test-name [ reaction probe-failtimes fail-times ] ] } &<1-16> [ fail-action
                                discard ] [ low-precedence ]
                                redirect ipv6 load-balance [ vpn-instance vpn-instance-name ] nexthop { ipv6-address [ track
                                nqa admin-name test-name [ reaction probe-failtimes fail-times ] ] } &<1-16> [ fail-action
                                discard ]

                                Only the S5732-H-V2, S6730E-H-V2, S6730-H-V2, S6780-H, S6750-H,
                                S6750E-S, S6750-S, S5755-S, S5755E-H, S5755-H, S5735R-S-V2, S5735E-
                                S-V2, S5735-S-V2, S5735I-S-V2, and S5735I-H-V2 support redirection to
                                multiple next-hop IP addresses.
                                The S3710-H does not support redirection to multiple next-hop IP
                                addresses.
                         –      Redirection to a remote next-hop address
                                redirect remote [ vpn-instance vpn-instance-name ] { ip-address [ track nqa admin-name test-
                                name [ reaction probe-failtimes fail-times ] ] } &<1-16> [ exact ] [ low-precedence ]
                                redirect remote [ vpn-instance vpn-instance-name ] ipv6-address &<1-16> [ exact ]

                                      NOTE

                                     If low-precedence is specified, the device does not support redirection of IPv6
                                     packets or IP packets encapsulated in tunnel mode.
                    4.   Exit the traffic behavior view.
                         quit

         Step 3 Configure a traffic policy.
                    1.   Create a traffic policy and enter the traffic policy view, or enter the view of an
                         existing traffic policy.
                         traffic policy policy-name

                    2.   Bind a traffic behavior and a traffic classifier to the traffic policy.
                         classifier classifier-name behavior behavior-name [ precedence precedence-value ]

                    3.   Exit the traffic policy view.
                         quit

         Step 4 Apply the traffic policy.
                    ●    Apply a traffic policy to the system.
                         a.     Apply a traffic policy to the system.
                                traffic-policy policy-name global [ slot slot-id ] { inbound | outbound }

                    ●    Apply a traffic policy to an interface.
                         a.     Enter the interface view.
                                interface interface-type interface-number

                         b.     Apply a traffic policy to the interface.
                                traffic-policy policy-name { inbound | outbound }


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                              55
QoS Configuration
QoS Configuration                                                                          7 Redirection Configuration


                        c.    Exit the interface view.
                              quit

                    ●   Apply a traffic policy to a VLAN.
                        a.    Create a VLAN and enter the VLAN view.
                              vlan vlan-id

                        b.    Apply a traffic policy to the VLAN.
                              traffic-policy policy-name { inbound | outbound }

                        c.    Exit the VLAN view.
                              quit

                    ●   Apply a traffic policy to a QoS group.
                        If the same traffic policy needs to be applied to multiple VLANs or interfaces,
                        you are advised to add the VLANs or interfaces to the same QoS group and
                        then apply the traffic policy to the QoS group.
                        a.    Create a QoS group and enter the QoS group view.
                              qos group group-name

                        b.    Add a specified interface or VLAN to the QoS group.
                              group-member { interface { interface-type interface-num | interface-name [ to interface-type
                              interface-num | interface-name ] } &<1-8> | vlan { vlanid [ to vlanid ] } &<1-8> }
                              Only one type of member can be specified.
                        c.    Apply a traffic policy to the QoS group.
                              traffic-policy policy-name { inbound | outbound }

                        d.    Exit the QoS group view.
                              quit

                    ----End

Verifying the Configuration

                    Operation                                           Command

                    Check the configured traffic classifiers.           display traffic classifier [ classifier-
                                                                        name ]

                    Check the configured traffic behaviors.             display traffic behavior [ behavior-
                                                                        name ]

                    Check the configured traffic policy.                display traffic policy [ policy-name
                                                                        [ classifier classifier-name ] ]

                    Check the traffic policy application                display traffic-policy applied-record
                    records.

