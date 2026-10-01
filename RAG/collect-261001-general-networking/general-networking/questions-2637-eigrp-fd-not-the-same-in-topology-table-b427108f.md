---
id: collect-261001-general-networking/general-networking/questions-2637-eigrp-fd-not-the-same-in-topology-table-b427108f
title: "EIGRP FD not the same in topology table"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-2637-eigrp-fd-not-the-same-in-topology-table-b427108f.md
source_anchor: ""
source_lines: [1, 15]
sha256: 6e7aa732c1ded71bd0f68f4894965eeddd145f8df62a366d6cbf63c14dc2105f
---

# EIGRP FD not the same in topology table

*Score : 9 | Source : https://networkengineering.stackexchange.com/questions/2637/eigrp-fd-not-the-same-in-topology-table*

In EIGRP when showing topology table, i notice in some entries are like below:
P 10.1.6.0/24, 1 successors, FD is 793600
            via 10.1.2.2 (2195456/281600), Serial0/0/0
           via 10.1.3.2 (77081600/281600), Serial0/0/1
It seems the successor should be 10.1.2.2, but why its FD shows 2195456, not equal to 793600?

---

### Reponse — score 3

When you manipulate the various k-values in order to affect route preference, this sort of output will sometimes be displayed until you reset your EIGRP neighbors. You can accomplish this by entering the following command(assuming Cisco gear): clear ip eigrp (AS#) neighbors where (AS#) is the EIGRP autonomous system number. The few times I have seen output similar to that posted(all in lab scenarios, btw), this has fixed it.
