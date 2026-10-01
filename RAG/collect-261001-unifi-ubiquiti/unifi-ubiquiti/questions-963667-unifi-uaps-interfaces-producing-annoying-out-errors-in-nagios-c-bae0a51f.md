---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-963667-unifi-uaps-interfaces-producing-annoying-out-errors-in-nagios-c-bae0a51f
title: "questions-963667-unifi-uaps-interfaces-producing-annoying-out-errors-in-nagios-c-bae0a51f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-963667-unifi-uaps-interfaces-producing-annoying-out-errors-in-nagios-c-bae0a51f.md
source_anchor: ""
source_lines: [1, 10]
sha256: a9393e4c86ac356986bb0f66a644ba0f9bbea937993da1c926386504a4cbc497
---

# questions-963667-unifi-uaps-interfaces-producing-annoying-out-errors-in-nagios-c-bae0a51f

Since I have setup some UAPs and monitor them with my check_mk server over SNMP, I get flooded (a few hundreds a day) with warnings and critical messages about out-errors on the interfaces. What are these errors (percentage varies from 0.x% - 70% and more).
Example:
Host:     UAP-AC-PRO-5
Alias:    UAP-AC-PRO-5
Address:  192.16...
Service:  Interface 09
Event:    OK -> CRIT
Output:   CRIT - [ath3] (up) MAC: FC:EC:DA:7E:A5:FF, speed unknown, in: 583.15 B/s, out: 300.79 B/s, out-errors: 1.31%(!!) = 0.1%
Perfdata: in=583.147261;;;0; inucast=2.596055;;;; innucast=0;;;; indisc=0;;;; inerr=0;0.01;0.1;; out=300.792955;;;0; outucast=2.512848;;;; outnucast=0;;;; outdisc=0;;;; outerr=0.033283;0.01;0.1;; outqlen=0;;;0;
What are these messages and what do they tell me about my WIFI setup? I notice no issues on wifi no connection drops or packet loss so I'd like to get rid of the messages.
