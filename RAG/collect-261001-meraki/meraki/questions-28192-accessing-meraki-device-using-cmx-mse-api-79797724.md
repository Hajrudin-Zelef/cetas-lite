---
id: collect-261001-meraki/meraki/questions-28192-accessing-meraki-device-using-cmx-mse-api-79797724
title: "questions-28192-accessing-meraki-device-using-cmx-mse-api-79797724"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-28192-accessing-meraki-device-using-cmx-mse-api-79797724.md
source_anchor: ""
source_lines: [1, 6]
sha256: 36da11ca52b8d8017fed14a67da0db275cf9fb94163d0d97fd1ffe2733b8e1c7
---

# questions-28192-accessing-meraki-device-using-cmx-mse-api-79797724

How do I access a Meraki Device? Similar to the way in this tutorial. Cisco DevNet: CMX Mobility Services - Tutorials - MSE API Introduction
Here is snapshot of the Public IP of my Meraki device
I replaced the Public IP in the tutorial with the Public IP address of my meraki device- the APIs do not work. I am able to use those APIs when connecting to the Public IP specified int he tutorial.
There is a SSL certificate preloaded in the tutorial - the purpose of certificate is to ensure the identity of the remote computer (as in msesandbox.cisco.com). Do I need to install a certificate on my Meraki device too? My take is it should not be necessary - at least to start with I should be able to use http (instead of https).
Ping to the public IP address of the Meraki device does not work.
what am I missing here? Any suggestions / pointers? Please help.
