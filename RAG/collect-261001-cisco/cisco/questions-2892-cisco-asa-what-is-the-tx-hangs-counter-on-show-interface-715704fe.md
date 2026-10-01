---
id: collect-261001-cisco/cisco/questions-2892-cisco-asa-what-is-the-tx-hangs-counter-on-show-interface-715704fe
title: "questions-2892-cisco-asa-what-is-the-tx-hangs-counter-on-show-interface-715704fe"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "full-duplex"]
source: docs/RAG/collect-261001-cisco/questions-2892-cisco-asa-what-is-the-tx-hangs-counter-on-show-interface-715704fe.md
source_anchor: ""
source_lines: [1, 18]
sha256: e2638fd24fbd917e50ae3ef1b51cfad75f1351691fcfda7fcb1a899c98ae9486
---

# questions-2892-cisco-asa-what-is-the-tx-hangs-counter-on-show-interface-715704fe

Consider the following output from show interface:
Interface GigabitEthernet0/1 "inside", is up, line protocol is up
  Hardware is i82546GB rev03, BW 1000 Mbps, DLY 10 usec
    Full-Duplex(Full-duplex), 1000 Mbps(1000 Mbps)
    Input flow control is unsupported, output flow control is off
    MAC address 5057.aaaa.25d7, MTU 1500
    IP address 10.0.0.7, subnet mask 255.255.255.0
    2708954646 packets input, 1614638330819 bytes, 0 no buffer
    Received 786860 broadcasts, 0 runts, 0 giants
    7262 input errors, 7262 CRC, 0 frame, 0 overrun, 0 ignored, 0 abort
    0 pause input, 0 resume input
    0 L2 decode drops
    2622997029 packets output, 1295347863682 bytes, 6396 underruns
    0 pause output, 0 resume output
    0 output errors, 0 collisions, 90 interface resets
    0 late collisions, 0 deferred
    0 input reset drops, 7908 output reset drops, 31 tx hangs
The last line says 31 tx hangs. What does this mean?
