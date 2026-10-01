---
id: collect-261001-cisco/cisco/questions-2144-ios-changes-the-snmpv3-engineid-a8054339
title: "questions-2144-ios-changes-the-snmpv3-engineid-a8054339"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-2144-ios-changes-the-snmpv3-engineid-a8054339.md
source_anchor: ""
source_lines: [1, 15]
sha256: ab3e745a96a2fd5b1a8662865885a9f59b0f444b955d5a1ffa9dd9213b091cf8
---

# questions-2144-ios-changes-the-snmpv3-engineid-a8054339

I have encountered a very strange problem when configuring remote target engineId on a Catalyst 3750 switch. The switch changes the engineId, it puts a zero between the two last digits. As shown below it changes ...e166 to ...e1606. I have tried with removing the user, rebooting the switch and re-adding the user but the switch changes the engineId every time. I have configured tens of switches in the exact same way and has never come across this problem before. Does anybody know what the problem might be?
sw21(config)#snmp-server engineID remote 10.1.9.6 udp-port 162 b7a9d3ca99325e6b5fb2894a500e166
sw21#show snmp user
User name: trap
Engine ID: B7A9D3CA99325E6B5FB2894A500E1606
storage-type: nonvolatile        active
Authentication Protocol: SHA
Privacy Protocol: None
Group-name: sys
sw21#show version
....
Switch Ports Model              SW Version            SW Image
------ ----- -----              ----------            ----------
*    1 30    WS-C3750X-24       12.2(55)SE5           C3750E-UNIVERSALK9-M
     2 30    WS-C3750X-24       12.2(55)SE5           C3750E-UNIVERSALK9-M
