---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fo-60be8d6e-4
title: "FortiGuard configuration"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fo-60be8d6e.md
source_anchor: ""
source_lines: [796, 919]
sha256: f079b8882eb5b7f5e3dbcf10be4cf84d58af9bd374b8edfefbaf5141aae6e971
---

# FortiGuard configuration

| ISDB entries | diagnose firewall internet-service list | 
| Update versions | diagnose autoupdate versions | 
| IoT daemon debug | diagnose debug application iotd -1 | 
| Device list | diagnose user device list | 
| Disable local CIDB signatures | diagnose cid sigs disable | 
- Default route exists.
- DNS works.
- FortiGate can reach the Internet.
- Required FortiGuard ports are allowed.
- Upstream firewall is not blocking FortiGuard.
- Proxy settings are correct if used.
- FortiGate is registered.
- FortiGuard contract is valid.
- Required subscription exists.
- System time is correct.
- Certificate chain is trusted.
- DNS resolves to the expected service.
- No TLS interception is breaking validation.
- Required TLS versions are supported.
- Automatic update is enabled.
- FortiGuard location is appropriate.
- Update server is reachable.
- AV/IPS package signature validates.
- Update debug has no errors.
- FortiGate can reach FortiManager.
- Update server is configured.
- Rating server is configured.
- Correct port is reachable.
- Fallback behavior is understood.
FortiGuard
   |
   +-- AV Updates
   |
   +-- IPS Updates
   |
   +-- Web Rating
   |
   +-- AntiSpam
   |
   +-- Application Intelligence
   |
   +-- IoT Intelligence
   |
   +-- Threat Intelligence
   |
   +-- Security Databases
   |
   +-- Cloud Services
| Topic | Remember | 
|---|---|
| AV/IPS packages | Digitally signed by Fortinet CA | 
| Anycast | Helps reach an appropriate FortiGuard service location | 
| Web/Spam | Can use communication mechanisms that are not proxy-friendly | 
| FortiManager | Can act as a local FortiGuard update/rating source | 
| Air-Gap | Requires offline/manual update/licensing workflows | 
| IoT Detection | Uses FortiGuard intelligence when local identification is insufficient | 
| ISDB | Built-in database allows FortiGate policies to continue functioning during update transition | 
| Malware statistics | Used to improve Fortinet threat intelligence | 
| TLS | Certificate/trust validation is part of secure communication | 
| Cache | Reduces repeated FortiGuard lookups | 
| Diagnostics | diagnose sys service-communication is a key starting point | 
# FortiGuard configuration
config system fortiguard
    set protocol https
    set port 443
    set update-server-location automatic
end
# Automatic updates
config system autoupdate schedule
    set status enable
    set frequency automatic
end
# Update tunneling
config system autoupdate tunneling
    set address <PROXY-IP>
    set port <PORT>
    set username <USERNAME>
    set password <PASSWORD>
    set status enable
end
# FortiGuard diagnostics
diagnose sys service-communication
# Update debug
diagnose debug application updated -1
diagnose debug enable
# Stop debugging
diagnose debug disable
# IoT diagnostics
diagnose debug application iotd -1
diagnose debug enable
diagnose user device list
# ISDB
diagnose firewall internet-service list
diagnose autoupdate versions | grep internet -a 6
                           ┌─────────────────────┐
                           │    FortiGuard Labs   │
                           │ Threat Intelligence  │
                           └──────────┬──────────┘
                                      |
                              FortiGuard Network
                                      |
                 ┌────────────────────┼────────────────────┐
                 |                    |                    |
              Anycast              Regional           FortiManager
                 |                 Servers              Local FDN
                 |                    |                    |
                 └────────────────────┼────────────────────┘
                                      |
                                  FortiGate
                                      |
       ┌──────────────┬───────────────┼───────────────┬──────────────┐
       |              |               |               |              |
      AV             IPS          Web Filter       AntiSpam         IoT
       |              |               |               |              |
       └──────────────┴───────────────┼───────────────┴──────────────┘
                                      |
                               Security Decision
FortiGuard is not simply an “update server”; it is a distributed security-intelligence ecosystem that provides signatures, ratings, threat intelligence, device intelligence, and security databases through multiple communication models such as Anycast, direct connectivity, proxy-based access, and FortiManager-local services.
FortiGuard · FortiGate FortiGuard · FortiOS FortiGuard · FortiGuard Anycast · FortiGuard troubleshooting · FortiGuard update · FortiManager FortiGuard · FortiGate IoT Detection · FortiGate ISDB · FortiGate air gap · FortiGate offline licensing · FortiGuard proxy · FortiGuard AV IPS update · FortiGuard Labs · NSE4 · NSE7 · FortiGate CLI · Fortinet Security
- 
YouTube — SheynShield 
  - Fortinet NSE content
  - FortiGate troubleshooting
  - Network Security Engineering
