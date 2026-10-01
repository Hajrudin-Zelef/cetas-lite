---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-administration-guide-783883-configuring-fortiguard-upda-d8838c30
title: "Configuring FortiGuard updates"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-administration-guide-783883-configuring-fortiguard-upda-d8838c30.md
source_anchor: ""
source_lines: [1, 14]
sha256: 955a27e0af4635b6c38cd7647e8191704df2d449b25bb8ec4e190ab4a6bfc857
---

# Configuring FortiGuard updates

# Configuring FortiGuard updates

###### To configure FortiGuard updates:

1. 
                                                    Go to *System > FortiGuard*
2. 
                                                    Scroll down to the *FortiGuard Updates* section.
3. 
                                                    Configure the options for connecting and downloading definition files: Immediately download updates The option can be enabled on 2U and larger hardware models when the FortiGuard are servers are connected in anycast mode. The FortiGate forms a secure, persistent connection with FortiGuard to get notifications of new updates through an HTTPS connection. The FortiGate uses the fds_notify daemon to wait for the notification, then makes another connection to the FortiGuard server to download the updates. Scheduled Updates Enable to schedule updates to be sent to the FortiGate at the specified time or automatically. See Scheduled updates and Automatic updates. Improve IPS quality Enable to send information to the FortiGuard servers when an attack occurs. This can help keep the FortiGuard database current as attacks evolve, and improve IPS signatures. Use extended IPS signature package Enable to use the extended IPS database, that includes protection from legacy attacks, along with the regular IPS database that protects against the latest common and in-the-wild attacks. AntiVirus PUP/PUA Enable antivirus grayware checks for potentially unwanted applications. Update server location The FortiGuard update server location. See Update server location for details.
4. 
                                                    Click *Apply* .
