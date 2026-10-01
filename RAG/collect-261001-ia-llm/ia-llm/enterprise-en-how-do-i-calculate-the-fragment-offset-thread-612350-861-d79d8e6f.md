---
id: collect-261001-ia-llm/ia-llm/enterprise-en-how-do-i-calculate-the-fragment-offset-thread-612350-861-d79d8e6f
title: "enterprise-en-how-do-i-calculate-the-fragment-offset-thread-612350-861-d79d8e6f"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei"]
dates: ["2020-04-02"]
keywords: []
source: docs/RAG/collect-261001-ia-llm/enterprise-en-how-do-i-calculate-the-fragment-offset-thread-612350-861-d79d8e6f.md
source_anchor: ""
source_lines: [1, 22]
sha256: a76e88c87a3f190cdbbdfae6d3eb01b26187bd9952e7d997d5a6e5cb61991225
---

# enterprise-en-how-do-i-calculate-the-fragment-offset-thread-612350-861-d79d8e6f

Huawei Enterprise Support Community 
- User Guide
- English-Intl
- Login
Sprout2020-04-02 14:56:54What's New:2020-04-02 22:56:54
6112130
Hi all,
How do I calculate the fragment offset?
For example:
The data part of a data packet has 3800 bytes (using the fixed header), and the MTU of the interface is 1420 bytes. How do I calculate the fragment offset?
The right answer is 0,175,350.
Thanks.
More Articles
1CPU attack defence
2Port security
3How to solve the eNSP 40 error
4OSPF Packet Transmission Mode
5IPv6 Tips
Hi Sprout,
The fragment offset is the relative position of a fragment in the original packet. The offset unit is 8 bytes. That is, the length of each fragment must be an integral multiple of 8 bytes (64 bits).
A total of 3800 bytes need to be transmitted, and a maximum of 1400 bytes can be transmitted at a time (because 20 bytes of the fixed IP header need to be added, the MTU=1420 bytes). Therefore, only three times of transmission are required, that is, 1400 bytes + 1400 bytes + 1000 bytes (3800 bytes in total).
In this way, the fragment offset is 0 during the first transmission. For the second transmission, the offset is 175 (1400/8). For the third transmission, the fragment offset is 2800/8 = 350.
