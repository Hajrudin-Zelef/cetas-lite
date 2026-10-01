---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-709167-configuring-an-ssl-ssh-ins-fdd63f52-2
title: "document-fortigate-latest-administration-guide-709167-configuring-an-ssl-ssh-ins-fdd63f52"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-709167-configuring-an-ssl-ssh-ins-fdd63f52.md
source_anchor: ""
source_lines: [57, 81]
sha256: a347ff6f617457efb43fd4ed702ea38ce8e746999bd41e6e53e916a953dae53f
---

# document-fortigate-latest-administration-guide-709167-configuring-an-ssl-ssh-ins-fdd63f52

                                                                            Keep Untrusted & Allow: Allow the server certificate and keep it untrusted.
  - 
                                                                            Block: Block the certificate.
  - 
                                                                            Trust & Allow: Allow the server certificate and re-sign it as trusted.
 Log SSL anomalies Enable this feature to record and log traffic sessions containing invalid certificates. By default, SSL anomalies logging is enabled. Logs are generated in the UTM log type under the SSL subtype when invalid certificates are detected.
- 
                                                                            
- 
                                                    Click OK.
Inspecting all ports
The behavior of inspecting all ports can be different between flow and proxy mode inspection when the inspection mode is configured in a firewall policy.
In proxy mode inspection, when deep inspection is enabled:
- 
                                                    If Inspect all ports is disabled, only the ports specified in the Protocol Port Mapping section will be scanned.
- 
                                                    If Inspect all ports is enabled, all ports will be scanned.
In flow mode inspection, when deep-inspection is enabled:
- 
                                                    All ports will be scanned, regardless of whether or not Inspect all ports is enabled or disabled.
When performing certificate inspection instead of deep inspection, flow and proxy mode inspection behave the same:
- 
                                                    If Inspect all ports is disabled, only the ports specified in the Protocol Port Mapping section will be scanned.
- 
                                                    If Inspect all ports is enabled, all ports will be scanned.
