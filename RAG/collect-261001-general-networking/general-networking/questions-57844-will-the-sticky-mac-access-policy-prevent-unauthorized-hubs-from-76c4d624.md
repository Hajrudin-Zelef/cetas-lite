---
id: collect-261001-general-networking/general-networking/questions-57844-will-the-sticky-mac-access-policy-prevent-unauthorized-hubs-from-76c4d624
title: "questions-57844-will-the-sticky-mac-access-policy-prevent-unauthorized-hubs-from-76c4d624"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-57844-will-the-sticky-mac-access-policy-prevent-unauthorized-hubs-from-76c4d624.md
source_anchor: ""
source_lines: [1, 3]
sha256: 1d925f07ee64ad21ce558e1930ca23207f176a3c96b74ec507364789311f87a7
---

# questions-57844-will-the-sticky-mac-access-policy-prevent-unauthorized-hubs-from-76c4d624

As Jesse P explained, hubs do not have MAC addresses, but multiple devices connected to a hub would mean multiple MAC addresses on the switch interface, and what you suggest would detect that and prevent a situation where a hub is attaching multiple devices to a single switch interface. Unfortunately, it will be unable to detect the hub or a hub with a single device connected.
You must carefully consider your plan. For example, using a VoIP phone with a PC plugged into it could use two or three MAC addresses on the single switch interface. I have seen that mess up plans such as yours because you must allow more than one MAC address at a time for things to work correctly.
If what people are doing is to connect small switches, rather than hubs, then those switches would have MAC addresses, and may even be sending BPDUs. You could then configure something like bpduguard that will disable the switch interface when it receives BPDUs. That is a very common, and recommended, practice for access interfaces on a switch.
