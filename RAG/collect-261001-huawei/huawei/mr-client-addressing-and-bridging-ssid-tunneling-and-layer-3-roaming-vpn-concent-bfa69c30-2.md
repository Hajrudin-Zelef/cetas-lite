---
id: collect-261001-huawei/huawei/mr-client-addressing-and-bridging-ssid-tunneling-and-layer-3-roaming-vpn-concent-bfa69c30-2
title: "mr-client-addressing-and-bridging-ssid-tunneling-and-layer-3-roaming-vpn-concent-bfa69c30"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/mr-client-addressing-and-bridging-ssid-tunneling-and-layer-3-roaming-vpn-concent-bfa69c30.md
source_anchor: ""
source_lines: [61, 62]
sha256: 3cc493eec954ef2347cf28af584d7fab7f45160fd5ddb81df5fd8fd526b52b6e
---

# mr-client-addressing-and-bridging-ssid-tunneling-and-layer-3-roaming-vpn-concent-bfa69c30

Note: IPsec tunnels between peers never traverse the Cloud. The VPN registry simply acts as a broker allowing peers to exchange connection specific information. The actual IPsec tunnel is always peer-to-peer.
The above example does not consider upstream NAT firewalls doing PAT (NAT overload). If the upstream firewalls are doing PAT, source ports will change when Register-Request packets are sent out. Rewritten ports will be provided to the VPN registry and will be shared to other registered peers as such.
