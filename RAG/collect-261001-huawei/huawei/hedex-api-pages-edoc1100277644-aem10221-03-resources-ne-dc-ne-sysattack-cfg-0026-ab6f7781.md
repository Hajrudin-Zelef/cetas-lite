---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100277644-aem10221-03-resources-ne-dc-ne-sysattack-cfg-0026-ab6f7781
title: "hedex-api-pages-edoc1100277644-aem10221-03-resources-ne-dc-ne-sysattack-cfg-0026-ab6f7781"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100277644-aem10221-03-resources-ne-dc-ne-sysattack-cfg-0026-ab6f7781.md
source_anchor: ""
source_lines: [1, 19]
sha256: 571c09b91affb205cdecb5660c42fc35390473c7c19e440b43c7a0d059464fc5
---

# hedex-api-pages-edoc1100277644-aem10221-03-resources-ne-dc-ne-sysattack-cfg-0026-ab6f7781

After configuring CAR, run the corresponding display command to check statistics about packets discarded by CAR.
| Table 1 Displaying the weights of packets to be sent to the CPU in protocol queues of the protocol groups |  | 
|---|---|
| Operation | Command | 
|---|---|
| Query the weight of packets to be sent to the CPU in a specific queue of the whitelist protocol group. | display cpu-defend protocol-group whitelist queue { whitelist-bgp \| whitelist-ldp \| whitelist-management \| whitelist-multicast \| whitelist-reserve } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the user-defined flow protocol group. | display cpu-defend protocol-group user-defined-flow queue { user-define-flow-1 \| user-define-flow-2 \| user-define-flow-3 \| user-define-flow-4 \| user-define-flow-5 \| user-define-flow-6 \| user-define-flow-7 \| user-define-flow-8 } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the management protocol group. | display cpu-defend protocol-group management queue { dcn \| ftp \| ntp \| snmp \| ssh \| sshv6 \| syslog \| telnet } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the routing protocol group. | display cpu-defend protocol-group route-protocol queue { bgp \| bgpv6 \| isis \| ospf \| ospfv3 \| rip } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the multicast protocol group. | display cpu-defend protocol-group multicast queue { igmp \| multicast-reserve \| msdp \| pim } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the ARP protocol group. | display cpu-defend protocol-group arp queue { arp \| nd } weight configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the MPLS protocol group. | display cpu-defend protocol-group mpls queue { ldp \| oam-ping \| rsvp \| vxlan } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the user access protocol group. | display cpu-defend protocol-group access-user queue { bas-arp \| bas-igmp \| bas-nd \| bas-trigger \| dhcp \| dhcpv6 \| eapol \| l2tp \| lldp \| ppp \| vbas-reserve \| web } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the link-layer protocol group. | display cpu-defend protocol-group link-layer queue { 3ah \| bfd \| link-detect \| trunk \| y1731 \| interface-rdi \| lag-check \| lag-ping-trace \| mac-vlan } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the network-layer protocol group. | display cpu-defend protocol-group network-layer queue { clock \| default \| dns \| fragment \| gre \| hwtacas \| icmp \| icmpv6 \| ipv4-reserve \| ipv6-option \| nhrp \| vrrp \| radius-diameter } configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the system-message protocol group. | display cpu-defend protocol-group system-message queue system-message configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the blacklist protocol group. | display cpu-defend protocol-group blacklist queue blacklist configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the detection protocol group. | display cpu-defend protocol-group check-failed queue check-failed configuration slot slot-id | 
| Query the weight of packets to be sent to the CPU in a specific queue of the forwarding protocol group. | display cpu-defend protocol-group fwddata-to-cp queue forward-data configuration slot slot-id |
