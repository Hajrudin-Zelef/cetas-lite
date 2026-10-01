---
id: collect-261001-meraki/meraki/questions-780663-cisco-meraki-ad-authentication-on-ipads-not-working-9846d9c0
title: "questions-780663-cisco-meraki-ad-authentication-on-ipads-not-working-9846d9c0"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-meraki/questions-780663-cisco-meraki-ad-authentication-on-ipads-not-working-9846d9c0.md
source_anchor: ""
source_lines: [1, 6]
sha256: 210ff765561e273bf7333d7ce74e27a585178939b87b798a632c3e95b794038d
---

# questions-780663-cisco-meraki-ad-authentication-on-ipads-not-working-9846d9c0

I've had Cisco Meraki setup for a while for our iPad MDM and it's been great. I am now trying to take it a setup further by added in AD authentication at enrollment time.
I've followed this how-to from Cisco Meraki, I'm using the thrid option Active Directory via SM Agent.
Everything seems to be ok from the Meraki point of view. I get the green check mark next to the domain controller I just added under the status column.
On the iPad, when enrolling, it asks for a username and password, but whatever combo of username, email address etc that I try does not work.
Windows firewall is off on the servers and I've double checked the Firewall Information page from Meraki. All my required ports seem to be correctly opened. I also get a good connection status of the windows server in the Meraki client console.
I am missing a step?
