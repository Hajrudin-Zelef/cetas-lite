---
id: collect-261001-meraki/meraki/questions-24495-setting-up-vlan-on-new-meraki-switch-afd36e03
title: "questions-24495-setting-up-vlan-on-new-meraki-switch-afd36e03"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-24495-setting-up-vlan-on-new-meraki-switch-afd36e03.md
source_anchor: ""
source_lines: [1, 21]
sha256: 5000def355932f9b47b0b64119e43acd2851933d8f8c71308d9eff90b12d5ac5
---

# questions-24495-setting-up-vlan-on-new-meraki-switch-afd36e03

Hi we have DellForce10 switches and we are changing them to Cisco Meraki switches. On the Dell switch we have the following Vlan setup:
Covington-CustomerCheckout-S25>show vlan
Codes: * - Default VLAN, G - GVRP VLANs, P - Primary, C - Community, I - Isolated
Q: U - Untagged, T - Tagged
   x - Dot1x untagged, X - Dot1x tagged
   G - GVRP tagged, M - Vlan-stack
    NUM    Status    Description                     Q Ports
    1      Active                                    T Po1(Gi 0/21-22)
                                                    U Gi 0/9
    102    Active    Wireless                        T Po1(Gi 0/21-22)
                                                     U Gi 0/18,20
    105    Active    DHCP_Clients                    T Po1(Gi 0/21-22)
                                                     U Gi 0/1-8,10,12-17,19
    106    Active    Static_Clients                  T Po1(Gi 0/21-22)
    109    Active    Production                      T Po1(Gi 0/21-22)
    110    Active    VOIP                            T Po1(Gi 0/21-22)
    112    Active    Management                      T Po1(Gi 0/21-22)
    220    Active    Printers                        T Po1(Gi 0/21-22)
                                                     U Gi 0/11
    999    Inactive
Now, the Meraki switch has Native and allowed VLANS and I am not sure how to set it up. Below are the options that I have to set it up with. I am also a little confused on what the untagged and tagged traffic means on the the Dell switch and how that would translate onto the configuration of the Meraki Switch.
