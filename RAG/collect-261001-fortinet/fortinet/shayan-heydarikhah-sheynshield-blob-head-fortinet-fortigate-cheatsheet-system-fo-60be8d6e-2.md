---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fo-60be8d6e-2
title: "FortiGuard configuration"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["distribution", "license"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fo-60be8d6e.md
source_anchor: ""
source_lines: [262, 566]
sha256: d42773d664ef242c231555478e0645a76455994548c74f8826a65fa813cf0325
---

# FortiGuard configuration

Cache values and supported options are release dependent.
FortiGuard services can have force-off controls.
Example:
config system fortiguard
    set antispam-force-off disable
    set webfilter-force-off disable
end
Conceptually:
force-off enable
      |
      v
Service disabled
force-off disable
      |
      v
Service available
Example configuration:
config system fortiguard
    set outbreak-prevention-force-off disable
    set outbreak-prevention-cache enable
    set outbreak-prevention-cache-ttl 300
    set outbreak-prevention-cache-mpercent 2
    set outbreak-prevention-timeout 7
end
The cache helps reduce repeated external lookups.
Example:
config system fortiguard
    set sdns-server-ip 208.91.112.220
    set sdns-server-port 53
end
SDNS-related communication can be important for FortiGuard DNS/security services.
FortiGate can periodically send encrypted security statistics to FortiGuard.
Potential telemetry includes statistics related to:
- Antivirus
- IPS
- Botnet IP lists
- Application Control
The information can include device-related metadata such as:
- FortiGate IP address
- Serial number
- Country
The purpose is to improve Fortinet threat intelligence and security services.
The notes specify:
60 minutes
config system global
    set fds-statistics enable
    set fds-statistics-period 60
end
FortiGuard can use submitted malware statistics to identify active threats.
Conceptually:
New malware activity
        |
        v
FortiGate telemetry
        |
        v
FortiGuard analysis
        |
        v
Signature activity evaluation
        |
        +---- Active ----> Active AV database
        |
        +---- Inactive --> Extended AV database
If activity for an inactive threat reappears, its signature can be promoted back into the active database.
FortiGate and FortiGuard establish a secure TLS connection before exchanging protected data.
Conceptually:
FortiGate                         FortiGuard
   |                                  |
   |---- TLS ClientHello ------------>|
   |                                  |
   |<--- ServerHello + Certificate ---|
   |                                  |
   |---- Certificate / Key Info ----->|
   |                                  |
   |<--- Certificate Verify ----------|
   |<--- Finished --------------------|
   |                                  |
   |---- Finished ------------------->|
   |                                  |
   |<==== Encrypted Communication ===>|
The certificates must belong to a trusted Fortinet certificate chain.
During TLS validation, FortiGate evaluates the server identity and certificate status.
Important failure conditions include:
Certificate CN/SAN mismatch
        |
        v
DNS-resolved domain != certificate identity
        |
        v
TLS handshake aborted
Other abort conditions can include:
- Invalid OCSP status
- Revoked issuer CA
- Untrusted certificate chain
FortiGuard connectivity can use certificate status validation such as OCSP.
Conceptually:
FortiGate
   |
   | ClientHello
   | + OCSP status request
   v
FortiGuard
   |
   | Certificate + OCSP status
   v
FortiGate
   |
   +---- Valid ----> Continue TLS
   |
   +---- Invalid --> Abort
Exact certificate-validation behavior can vary by FortiOS release and service.
Example:
config system fortiguard
    set fortiguard-anycast enable
    set fortiguard-anycast-source fortinet
end
Depending on the FortiOS release, alternative source options may be available.
A useful conceptual distinction:
| Feature | Direct FortiGuard | Local FortiManager | 
|---|---|---|
| Anycast | Available | Local architecture | 
| Internet dependency | Usually yes | Can reduce direct dependency | 
| AV/IPS updates | Yes | Yes | 
| Rating services | Yes | Yes | 
| Local control | Lower | Higher | 
| Closed environment | Limited | Useful | 
| Centralized update distribution | Limited | Strong | 
For large environments:
FortiGuard
     |
     v
FortiManager
     |
     +---- FGT-01
     +---- FGT-02
     +---- FGT-03
     +---- FGT-04
diagnose sys service-communication
Useful for inspecting FortiGuard/service communication status.
diagnose debug application updated -1
diagnose debug enable
Use this when troubleshooting update behavior.
Remember to stop debugging after troubleshooting:
diagnose debug disable
Check:
System
  >
FortiGuard
  >
License
Also inspect:
Log & Report
  >
System Events
  >
General System Events
Cloud/service communication logging can also help identify:
- License problems
- FortiGuard connectivity issues
- Update failures
- Authentication problems
FortiGuard Labs provides online lookup resources.
Use it to determine:
- URL category
- URL rating
- Classification
It can also be used to request reevaluation of incorrectly categorized URLs.
Useful for researching:
- Viruses
- Botnet C&C
- IPS signatures
- Vulnerabilities
- Mobile malware
FortiGuard Threat Encyclopedia
Search application signatures and classifications.
FortiGuard Application Control
FortiGuard IoT Detection is a subscription-based service that can help identify devices not completely identified by the local device database.
Architecture:
Unknown Device
      |
      v
FortiGate Device Detection
      |
      v
Local CIDB
      |
      +---- Known ----> Device identified
      |
      +---- Unknown --> FortiGuard
                            |
                            v
                       Device analysis
                            |
                            v
                       Result returned
The FortiGate should:
- Be registered with FortiCare.
- Have an appropriate IoT Detection license.
- Be connected to a supported FortiGuard Anycast service.
- Have device detection enabled on the relevant interface.
Device
  |
  | DHCP / MAC / HTTP / traffic characteristics
  v
FortiGate
  |
  | Device information
  v
FortiGuard Collection Server
  |
  v
FortiGuard Query Server
  |
  v
Device identification
  |
  v
FortiGate Device Inventory
View results from:
Dashboard
  >
Users & Devices
  >
Device Inventory
For troubleshooting/testing:
diagnose cid sigs disable
This disables the local device signature database so FortiGate can query FortiGuard for device identification.
diagnose debug application iotd -1
diagnose debug enable
diagnose user device list
FortiAP can collect information about connected devices and provide it to FortiGate for device identification.
Potential data sources include:
DHCP
MAC address
HTTP
Device behavior
Traffic characteristics
Architecture:
Client Device
      |
      v
    FortiAP
      |
      v
  FortiGate
      |
      v
 FortiGuard IoT
      |
      v
Device intelligence
      |
      v
  FortiGate
Example:
config wireless-controller setting
    edit ap-1
        set device-weight 1
        set device-holdoff 5
        set device-idle 1440
    end
end
The notes describe weighting of different device-identification sources.
Conceptually:
FortiGuard intelligence
        +
DHCP information
        +
Behavioral analysis
        |
        v
Confidence score
        |
        v
Device classification
Exact scoring and supported ranges are version/platform dependent. Verify the CLI reference for the target FortiOS release.
Controls how long FortiGate waits before sending device information toward FortiGuard.
Example:
device-holdoff = 5 minutes
Controls how long inactive device information remains in the table.
Example:
device-idle = 1440 minutes
FortiGate can use an explicit proxy for certain FortiCloud/FortiGuard communication.
Concept:
FortiGate
    |
    | Proxy authentication
    v
Explicit Proxy
    |
    v
Internet
    |
    v
FortiCloud / FortiGuard
⚠️ Not every FortiGuard service supports these proxy settings. Web Filter service traffic, in particular, may follow a different communication mechanism.
config firewall proxy-policy
    edit 1
        set proxy explicit-web
        set dstintf port1
        set srcaddr all
        set dstaddr all
        set service webproxy
        set action accept
