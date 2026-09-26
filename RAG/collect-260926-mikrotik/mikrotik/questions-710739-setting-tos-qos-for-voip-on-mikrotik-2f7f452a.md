---
id: collect-260926-mikrotik/mikrotik/questions-710739-setting-tos-qos-for-voip-on-mikrotik-2f7f452a
title: "questions-710739-setting-tos-qos-for-voip-on-mikrotik-2f7f452a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/qos/questions-710739-setting-tos-qos-for-voip-on-mikrotik-2f7f452a.md
source_anchor: ""
source_lines: [1, 7]
sha256: 581870c684e4dfe9ee10363a6f5a79cf6e4d2cfb0fecb2e54272a5cf9581f4f9
---

# questions-710739-setting-tos-qos-for-voip-on-mikrotik-2f7f452a

I have a Mikrotik router which looks like iptables. I need to setup QoS to ensure VoIP phones get top traffic priority. How do I know what ToS to use and how would I get it? I searched and can't seem to find the ToS for VoIP.
My rule set so far is as follows:
/ip firewall mangle
add chain=forward tos=XXX action=mark-packet new-packet-mark=VoIP passthrough=no comment="voip" disabled=no
/ queue tree
add name="ether1_voip" parent=ether1 packet-mark=VoIP limit-at=0 queue=default priority=2 \
    max-limit=0 burst-limit=0 burst-threshold=0 burst-time=0s disabled=no
