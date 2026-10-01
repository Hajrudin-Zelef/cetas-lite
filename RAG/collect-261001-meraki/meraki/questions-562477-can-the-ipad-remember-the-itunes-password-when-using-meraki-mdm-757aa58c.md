---
id: collect-261001-meraki/meraki/questions-562477-can-the-ipad-remember-the-itunes-password-when-using-meraki-mdm-757aa58c
title: "questions-562477-can-the-ipad-remember-the-itunes-password-when-using-meraki-mdm-757aa58c"
domain: meraki
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-562477-can-the-ipad-remember-the-itunes-password-when-using-meraki-mdm-757aa58c.md
source_anchor: ""
source_lines: [1, 26]
sha256: c5d9cfdd3ffb8f590f25d981b11e4c05e00c6847752654f78d418e6777ea0284
---

# questions-562477-can-the-ipad-remember-the-itunes-password-when-using-meraki-mdm-757aa58c

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
7
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm currently using Cisco Meraki MDM to manage apps that get pushed out to company-owned iPads. In order for the company to retain ownership of the apps we purchase, apps are installed through the company iTunes account.
The problem is that when the apps installed with that account have updates to install, the user is being prompted for the password of the company iTunes account. I don't want to give the users this password for several reasons - to keep the password from being changed, to keep apps from being purchased on the company account, etc.
Is there a way to set the iPad to remember the password for the iTunes account so the user is not prompted for it when updates are installed? And if not, is there a way to bypass this issue when pushing out apps through MDM?
Does the
user have to enter the Apple ID password for every app installed?
When
Apps are deployed via MDM, Apple requires an Apple ID and password for
the app to be installed. Apps downloaded and installed via Apple
Configurator do not require entering an Apple ID and password,
however, the iOS device has to be physically connected to the OS X
device running Apple Configurator.
With iOS 6, the device caches a
users password for 15 minutes. If you install FREE apps in batches via
Systems Manager with iOS 6, you will have to enter the password once
instead of doing it for every single App.
Paid apps redeemed via VPP
redemption codes still require the user to enter a password for EVERY
app. Devices running pre-iOS 6 will be required to enter a password
for every app regards of whether it is free or paid
Some further Google searching says that some MDM applications have an unsupported way of doing what you want, but Meraki doesn't. Blame Apple for this requirement.
