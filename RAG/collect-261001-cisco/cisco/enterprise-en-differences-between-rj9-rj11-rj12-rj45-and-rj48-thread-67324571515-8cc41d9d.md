---
id: collect-261001-cisco/cisco/enterprise-en-differences-between-rj9-rj11-rj12-rj45-and-rj48-thread-67324571515-8cc41d9d
title: "enterprise-en-differences-between-rj9-rj11-rj12-rj45-and-rj48-thread-67324571515-8cc41d9d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/enterprise-en-differences-between-rj9-rj11-rj12-rj45-and-rj48-thread-67324571515-8cc41d9d.md
source_anchor: ""
source_lines: [1, 28]
sha256: 027f1da805e1e1a116614bb9825e7b66157d91c85e0c2dbc05f75b1b12c2c581
---

# enterprise-en-differences-between-rj9-rj11-rj12-rj45-and-rj48-thread-67324571515-8cc41d9d

Hello everyone,
There are several common variants of Registered Jack, which are usually distinguished by a specific number after "RJ". This distinction corresponds to the location and the number of wires/connectors, and they vary in size. This article will give a brief introduction.
All of these represent the same 4P4C connector, in which case the change in quantity does not indicate any change in functionality. RJ9 is also known as RJ10 and RJ22.
RJ9 is typically used to connect a telephone headset (sometimes called a listening device) to its body. Therefore, RJ9 does not interface with the public network. It has the least number of pins (4). So, it's the smallest of the registered jacks.
The pinout and color code for RJ9/RJ10/RJ22 are shown in the following table:
| Locations | Pin | Function | Color Code | 
| 1 | T0+ | Transfer0+ | Black | 
| 2 | R0- | Receive0- | Red | 
| 3 | R1+ | Receive1+ | Green | 
| 4 | T1- | Transfer1- | Yellow | 
RJ11 belongs to the 6P4C modular connector category and is slightly larger than RJ9/RJ10/RJ22, which carry audio and control signals. They connect phones to the Public Switched Telephone Network (PSTN) network and are also used for asymmetric digital subscriber line (ADSL) and modem connections.
The 4 wires are used as two pairs so that the RJ11 can make two connections to connect two phone lines. Normally, one pair is used and the other pair is used as a backup.
The RJ12 is a 6P6C connector that has 6 positions and 6 wires. So the only difference between RJ12 and RJ11 is that RJ12 has two extra wires. Because the number of bits is the same, the size of RJ12 and RJ11 is the same,
The RJ12 can be connected 3 times using the available 3 pairs of wires. The RJ12 is primarily used for system phones, but it can also be used in place of the RJ11 connector. But the reverse is not possible because RJ11 can only make two connections, while RJ12 can make three.
The pinouts and color codes for the 6PGC connector are as follows:
| Locations | PIN | RJ11 | RJ12 | Color Code | 
| 1 | Transfer+ | NA | T3 | White/Orange | 
| 2 | Transfer+ | T2 | T2 | Black | 
| 3 | Receive- | R1 | R1 | Red | 
| 4 | Transfer+ | T1 | T1 | Green | 
| 5 | Receive- | R2 | R2 | Yellow | 
| 6 | Receive- | NA | R3 | Blue/brown | 
RJ45 is a well-known connector that is widely used for Ethernet and data connections. Therefore, they are also called Ethernet connectors. This is an 8P8C connector. Because of the large number of bits, RJ45 is slightly larger than RJ12. RJ45 terminates Cat6 and Cat5e standard coaxial cables.
The RJ45 pinouts are as follows:
The RJ48 connector is also an 8P8C connector, which is similar to the RJ45 connector. But neither connector is the same. Although RJ48 is used in the networking domain, it is specifically used to terminate T1 and ISDN networks. Another advantage of RJ48 over RJ45 is that it can be used for longer distance communication.
RJ48 is used to terminate shielded twisted pair (STP) cables and comes in different variations, including RJ48C, RJ48X, and RJ48S. The RJ48 pin combination can be 1, 2, 7, 8, or 1, 2, 4, or 5.
The RJ48 pins are as follows:
The above are the more common registered jack variants, which are simply compared and sorted out for reference only. Knowing the differences between them, it's helpful to choose the right RJ type for us.
