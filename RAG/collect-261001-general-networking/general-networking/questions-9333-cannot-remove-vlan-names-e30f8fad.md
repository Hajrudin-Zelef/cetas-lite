---
id: collect-261001-general-networking/general-networking/questions-9333-cannot-remove-vlan-names-e30f8fad
title: "questions-9333-cannot-remove-vlan-names-e30f8fad"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-9333-cannot-remove-vlan-names-e30f8fad.md
source_anchor: ""
source_lines: [1, 28]
sha256: b2996795edd96670f6d7d24a4b03ee1979fae6101aafbb05bc35ebcdbd839acc
---

# questions-9333-cannot-remove-vlan-names-e30f8fad

Hi I was practicing creating vlans and assigning them to ports however I cannot remove the names.
CCNA-SWITCH1#show vlan
VLAN Name                             Status    Ports
---- -------------------------------- --------- -------------------------------
1    default                          active    Fa0/1, Fa0/2, Fa0/3, Fa0/4
                                                Fa0/5, Fa0/6, Fa0/7, Fa0/8
                                                Fa0/9, Fa0/10, Fa0/11, Fa0/12
                                                Fa0/13, Fa0/14, Fa0/15, Fa0/16
                                                Fa0/17, Fa0/18, Fa0/19, Fa0/20
                                                Fa0/21, Fa0/22, Fa0/23, Fa0/24
                                                Gi0/1, Gi0/2
1002 fddi-default                     act/unsup
1003 trcrf-default                    act/unsup
1004 fddinet-default                  act/unsup
1005 trbrf-default                    act/unsup
CCNA-SWITCH1#show ip interface brief
Interface              IP-Address      OK? Method Status                Protocol
Vlan1                  unassigned      YES manual administratively down down
Vlan2                  unassigned      YES manual up                    down
Vlan3                  unassigned      YES manual up                    down
FastEthernet0/1        unassigned      YES unset  down                  down
FastEthernet0/2        unassigned      YES unset  down                  down
FastEthernet0/3        unassigned      YES unset  down                  down
FastEthernet0/4        unassigned      YES unset  down                  down
FastEthernet0/5        unassigned      YES unset  down                  down
FastEthernet0/6        unassigned      YES unset  down                  down
 --More--
I tried to remove vlan 2 and 3 i used the command no vlan 2 etc.. but they still showing
