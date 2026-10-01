---
id: collect-261001-cisco/cisco/questions-907-cisco-catalyst-4500e-modules-installation-0689bbf6
title: "questions-907-cisco-catalyst-4500e-modules-installation-0689bbf6"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-907-cisco-catalyst-4500e-modules-installation-0689bbf6.md
source_anchor: ""
source_lines: [1, 10]
sha256: 02dbce9f0ec372c82ad487258c7b1dd8c76bf631c0428439f1625c75a8aa1ce0
---

# questions-907-cisco-catalyst-4500e-modules-installation-0689bbf6

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
4500 line cards are hot swapable, so the answer to the first three questions is yes (assuming a 45xxR chassis).
A couple of caveats, for question 2, the supervisor has to meet the requirements for redundany. For instance, you would not be able to insert a Sup7 in a chassis with a SupIV. Additional configuration may be required as well to enable the redundancy properly.
For question 3, you shouldn't remove an active supervisor (while you can, this isn't best practice). If the supervisor you wish to remove is currently active, failover to the standby supervisor first.
The answer to the last question is yes, you can cause damage to the chassis and/or line card if you insert them incorretly (i.e. this has nothing to do with the chassis being running, but a purely mechanical problem). The levers should be open when inserted, the card should glide in easily until the levers start to close by themselves, and the final insertion should be done with the levers. If you have to use A LOT of force, it may not be aligned properly and you could result in damage.
On the 4500s, you also need to make sure you are using the proper slots. Do not try to install a supervisor module in a linecard slot or vice versa.
hot swapping of supervisors is supported without disrupting system operation
It is obviously advised to pull out the backup one, as you runs these in a "master/backup" scenario, with the "master" syncing CEF adjacencies, RIB updates, TCAM entries etc, to the "backup". So removing the "backup" SUP typically has no effect. Removing the "master" SUP should have no effect if you have things like Nonstop Forwarding and Stateful Switchover configured, if not, your results will vary depending on your set up.
4 No - These products are designed to be hot swappable (having said that, very very rarely they may malfunction. For example, pulling out a supervisor engine when you have two installed, could cause an outage, even though supervisor redundancy has been enabled with not stop forwarding etc. Sh*t happens, software does crash, so my recommendation would be to perform these supposedly "non service effecting" procedures during an announced maintenance period, just in case.).
In addition to what ytti said, the 4500 cards need quite "firm" pressure to seat properly. The extraction levers must rest all the way against the card when inserted. Make sure the cards are sliding into the "rails" properly, a twisted card can damage the card or the backplane connector
