---
id: collect-261001-huawei/huawei/questions-35830-d8a39b29
title: "questions-35830-d8a39b29"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-huawei/questions-35830-d8a39b29.md
source_anchor: ""
source_lines: [1, 22]
sha256: f97e6320d72800ed7a56d568c931d16cd930375d3f9af5400bc5bbd856e01860
---

# questions-35830-d8a39b29

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We've received an OLT ma5608T and an ONT hg8245H for a short period of time - we want to see if Huawei is better than Zhone.
We don't have a full documentation for this OLT because we're not a partner of Huawei.
My question is: is it possible to configure ONT services without accessing the ONT directly? On Zhone MXK OLT you'd do something like this in order to configure a PPPoE client on an ONT:
I work for Huawei and I have configured PPPoE, both client and server, recently for an AR200 router, to be honest with you, I don't know how much of a difference exist between my equipment and yours.
To configure the PPPoE client on AR200 you need the following:
Configure a Dialer-rule allowing IP
Configure a Dialer interface
Configure dialer queue-length on dialer interface
Configure dialer idle time on dialer interface
Configure dialer user name on dialer interface
Configure dialer group on dialer interface, this must match dialer rule name.
Bind dialer interface to physical interface
For PPP authentication, you use the dialer interface to configure either PAP or CHAP.
Hope this helps you a bit, currently I have not the configuration code on me; but if you need them just ask
After few days of asking on many different question boards and forums, I've found out that it's practically impossible to achieve this on Huawei. You need to use tr069 or directly access an ONT either by HTTP or SSH/Telnet. Also there's an application called iManager U2000. But still... there's absolutely no way of configuring a PPPoE client on hg8245H by using OMCI.
Huawei cannot do it through OMCI. They have implemented a machanism to achieve PPPoE credential delivery to their ONT via and XML file that resides on some virtual FTTP server that resides either on their OLT or NMS. You need to first download this file then extract the PPPoE info from it. PPPoE client on ONT will configure its client to process further.
