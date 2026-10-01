---
id: collect-261001-meraki/meraki/questions-24503-setting-up-a-new-meraki-ms220-48fp-switch-to-replace-a-pre-exist-072310c3
title: "questions-24503-setting-up-a-new-meraki-ms220-48fp-switch-to-replace-a-pre-exist-072310c3"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-24503-setting-up-a-new-meraki-ms220-48fp-switch-to-replace-a-pre-exist-072310c3.md
source_anchor: ""
source_lines: [1, 17]
sha256: b125f5d2293c8ede2017e80e5eb4e05b35cb12b6459ab29cd1af88e5ba7ab5f1
---

# questions-24503-setting-up-a-new-meraki-ms220-48fp-switch-to-replace-a-pre-exist-072310c3

I have a new Cisco Meraki Switch that is replacing a DellForce10 switch that is already in place. I need to configure the Meraki switch to how the Dell switch is setup. The only thing that is significant with the switch and how they are setup is the use of VLAN's. On the Dell switch I do not believe there is a Native VLAN setup on it as seen below.
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
To my understanding the Native VLAN will be left at 1 on the Meraki switch? The Meraki switch is a 48 port switch with four fiber ports, and the network connection for the switch will be setup with the connection coming through the fiber port 49. Would I have Port 49 setup to configure as type would be trunk, Native VLAN 1, and Allowed VLAN all.
Then would I have all of the ports 1-48 setup as access ports? So if we were to go by the Dell VLAN setup from above. On port 1, would I have it setup as Type: access, VLAN: 105? Below is what I am seeing for the window for the configuration of port 1/1. If someone could please help me in trying to figure out if this is the right way to set up the new switch as the old one? Thanks
