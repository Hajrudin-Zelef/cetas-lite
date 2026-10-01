---
id: collect-261001-fortinet/fortinet/questions-638283-can-i-have-an-alert-when-a-ipsec-is-down-in-fortigate-80c-b46cfdea
title: "questions-638283-can-i-have-an-alert-when-a-ipsec-is-down-in-fortigate-80c-b46cfdea"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-638283-can-i-have-an-alert-when-a-ipsec-is-down-in-fortigate-80c-b46cfdea.md
source_anchor: ""
source_lines: [1, 8]
sha256: d0709f0eeb45a15142d61d47a402d101a1c92279a50f7d5cfa94f66aa5ba04ca
---

# questions-638283-can-i-have-an-alert-when-a-ipsec-is-down-in-fortigate-80c-b46cfdea

We have a fortigate 80c and 60D with an IPSEC VPN .
Is there a way to use the log (or other tool) to send an email alert when the tunnel is down?
I couldn't see anything on the Log & Report tab
Do you have an existing monitoring platform? For VPN tunnels, I tend not to trigger alerts from the endpoint firewalls (e.g. SNMP), but monitor different sides of the link... A ping check from your monitoring solution is pretty darn good for this, as the tunnel will either be up and passing traffic or down.
Despite this, you can configure email alerts for certain functions/events in the Fortigate firewall.
  Configure an alert message that will notify you of activities that are important to be aware about. For example, a branch office does not have a FortiGate administrator so you need to know, at all times, that the IPSec VPN tunnel is up and running. An alert email notification message can be configured for sending only IPSec tunnel errors.
Steps:
Go to Log&Report > Log Config > Alert E-mail. Make sure you have a valid SMTP server configuration. Create a new alert and filter on IPSec tunnel errors.
