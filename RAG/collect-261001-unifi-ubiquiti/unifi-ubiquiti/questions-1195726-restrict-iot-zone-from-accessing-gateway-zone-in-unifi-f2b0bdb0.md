---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1195726-restrict-iot-zone-from-accessing-gateway-zone-in-unifi-f2b0bdb0
title: "Restrict IoT zone from accessing gateway zone in Unifi"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1195726-restrict-iot-zone-from-accessing-gateway-zone-in-unifi-f2b0bdb0.md
source_anchor: ""
source_lines: [1, 7]
sha256: 6427d21252ee88188612e90300a7212c99437a655a6224861a93f96f6babaf91
---

# Restrict IoT zone from accessing gateway zone in Unifi

*Score : 1 | Source : https://serverfault.com/questions/1195726/restrict-iot-zone-from-accessing-gateway-zone-in-unifi*

I'm looking to clamp down on my IoT devices, so I've put most of them in an IoT zone and network. Right now the IoT to Gateway policies allow mDNS and all traffic. Looking at the Zone Matrix, the Hotspot to Gateway looks like the policies that I need (minus radius).
Am I on the right path? Are there other ports or concerns by restricting traffic to the Gateway zone from the IoT zone?
Here's my current (default?) Hotspot to Gateway policies.
