---
id: collect-260926-rattrapage/rattrapage/questions-717038-mikrotik-switched-vlan-configuration-97bd1bc0
title: "questions-717038-mikrotik-switched-vlan-configuration-97bd1bc0"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/ai-llm/questions-717038-mikrotik-switched-vlan-configuration-97bd1bc0.md
source_anchor: ""
source_lines: [1, 50]
sha256: 3cdf600054ab7e6c73e0704e644ccdae4ffe7e5c2a535ed69b1bc03044162fdf
---

# questions-717038-mikrotik-switched-vlan-configuration-97bd1bc0

With Routerboard 450G I want to configure the 5 Ethernet ports as follows:
- ether1: vlan3, untagged
- ether2: vlan1, untagged
- ether3: vlan2, untagged
- ether4: vlan2, untagged
- ether5: vlan1-vlan3, tagged
Within vlan1, the device should have IP 10.7.1.3/24.
However, when I want to ping the Routerboard it is not reachable. I configured the device the following way:
- Include ether1 in switching:
/interface ethernet switch print
Flags: I - invalid
 #   NAME                                    TYPE          MIRROR-SOURCE                                  MIRROR-TARGET                                  SWITCH-ALL-PORTS
 0   switch1                                 Atheros-8316  none                                           none                                           yes
- Switch all ports together:
/interface ethernet print
Flags: X - disabled, R - running, S - slave
 #    NAME                                      MTU MAC-ADDRESS       ARP        MASTER-PORT                                    SWITCH
 0 R  ether1                                   1500 E4:8D:8C:18:D5:A1 enabled    none                                           switch1
 1  S ether2                                   1500 E4:8D:8C:18:D5:A2 enabled    ether1                                         switch1
 2  S ether3                                   1500 E4:8D:8C:18:D5:A3 enabled    ether1                                         switch1
 3  S ether4                                   1500 E4:8D:8C:18:D5:A4 enabled    ether1                                         switch1
 4 RS ether5                                   1500 E4:8D:8C:18:D5:A5 enabled    ether1                                         switch1
- Configure ports 1-4 to remove VLAN tags and port 5 to add them; set default VLAN IDs:
/interface ethernet switch port print
Flags: I - invalid
 #   NAME                                                         SWITCH                                                         VLAN-MODE VLAN-HEADER    DEFAULT-VLAN-ID
 0   ether1                                                       switch1                                                        secure    always-strip                 3
 1   ether2                                                       switch1                                                        secure    always-strip                 1
 2   ether3                                                       switch1                                                        secure    always-strip                 2
 3   ether4                                                       switch1                                                        secure    always-strip                 2
 4   ether5                                                       switch1                                                        secure    add-if-missing               0
 5   switch1-cpu                                                  switch1                                                        fallback  leave-as-is                  0
- Wire the VLANs together. As can be seen, ether2 and ether5 are connected through vlan1:
/interface ethernet switch vlan print
Flags: X - disabled, I - invalid
 #   SWITCH                                                                         VLAN-ID PORTS
 0   switch1                                                                              3 ether1
                                                                                            ether5
 1   switch1                                                                              1 ether2
                                                                                            ether5
 2   switch1                                                                              2 ether3
                                                                                            ether4
                                                                                            ether5
- Finally, add IP address to ether2:
/ip addr print
Flags: X - disabled, I - invalid, D - dynamic
 #   ADDRESS            NETWORK         INTERFACE
 0   10.7.1.3/24        10.7.1.0        ether2
Port 5 is directly connected to a managed switch which outputs vlan1-vlan3 tagged. The Routerboard cannot be pinged.
Do I miss something in the configuration or do I understand the concept wrong? (To my understanding, it would be sufficient to add the IP to ether2, because it is switched to vlan1 on ether5)
