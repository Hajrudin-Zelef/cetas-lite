---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fo-60be8d6e-1
title: "FortiGuard configuration"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["distribution", "latency", "license", "training"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fo-60be8d6e.md
source_anchor: ""
source_lines: [1, 261]
sha256: babda7bd003c79c5f26949740d52eb03a4fa4a53280bce0951286050eb091bda
---

# FortiGuard configuration

FortiGuard Services, Updates, Anycast, FortiManager, IoT Detection, Proxy & Air-Gap
Practical reference for FortiOS Administration / NSE4 / NSE7
FortiGuard provides cloud-based security intelligence and update services for FortiGate.
Common FortiGuard services include:
- Antivirus (AV)
- Intrusion Prevention System (IPS)
- Application Control
- AntiSpam
- Web Filtering
- Web Application Firewall (WAF)
- Botnet / threat intelligence
- IoT device identification
- Security rating / reputation services
- Firmware and security database updates
For normal cloud-based FortiGuard operation:
FortiGate
   |
   | HTTPS / DNS / FortiGuard protocols
   v
FortiGuard Distribution Network (FDN)
The FortiGate generally needs Internet connectivity to:
- Validate its FortiGuard license.
- Reach FortiGuard services.
- Download security databases.
- Receive updated security intelligence.
AV and IPS packages are cryptographically signed by Fortinet.
During an update, FortiGate validates the package signature before accepting it.
| Level | Behavior | 
|---|---|
| level-0 | Accept unsigned package | 
| level-1 | Warn and request confirmation | 
| level-2 | Reject unsigned package | 
If no level is configured, the effective behavior is generally Level 1.
⚠️ Security levels are preconfigured on the BIOS on supported platforms.
FortiGuard Package
       |
       v
Signature Validation
       |
       +---- Valid ----> Accept
       |
       +---- Invalid --> Security-level dependent behavior
FortiGuard automatic update behavior can operate at intervals determined by the FortiGuard update mechanism and contract state.
Typical automatic update intervals can range from approximately:
10 minutes
      |
      v
...
      |
      v
60 minutes
config system autoupdate schedule
    set status enable
    set frequency automatic
end
The update interval can be influenced by the percentage of remaining valid contract time.
Conceptually:
More remaining contract
        |
        v
Shorter update interval
        |
        v
More frequent updates
Example from the training scenario:
~70% contract remaining
        -> ~10 minute interval
~50% contract remaining
        -> ~20 minute interval
Exam Tip: Do not assume that every FortiGate always uses exactly the same update interval. Hardware/platform, service state, FortiOS version and FortiGuard policy can affect behavior.
On supported hardware/platforms, FortiGate can maintain a persistent secure connection to FortiGuard when the feature is available.
Conceptually:
FortiGate
   |
   | Persistent HTTPS connection
   v
FortiGuard
   |
   | New update notification
   v
fds_notify
   |
   | Download request
   v
FortiGuard Update Server
The fds_notify daemon waits for update notifications and then initiates the update download.
This feature availability is platform/version dependent.
The Improve IPS Quality option allows FortiGate to send attack-related information to FortiGuard.
Purpose:
Attack detected
      |
      v
FortiGate
      |
      | Security telemetry
      v
FortiGuard
      |
      v
Threat intelligence / signature improvement
The objective is to help FortiGuard improve IPS signatures as attacks evolve.
PUP/PUA detection allows FortiGate Antivirus functionality to identify:
- Potentially Unwanted Applications
- Grayware
- Other potentially unwanted software
This can provide an additional detection layer beyond conventional malware signatures.
FortiGate can use a proxy/tunnel for certain FortiGuard update operations.
Example:
config system autoupdate tunneling
    set address 1.2.3.4
    set port 1344
    set username test
    set password 123
    set status enable
end
Proxy tunneling is supported only for specific services, including:
- Registration
- Antivirus updates
- IPS updates
For FortiGate VMs, proxy tunneling can additionally be used for license validation.
A critical distinction:
AV / IPS Update
       |
       +---- Proxy tunneling supported
Web Filtering / Spam Rating
       |
       +---- UDP-based communication
Web filtering and spam-filtering services may use UDP communication such as:
UDP/53
UDP/8888
UDP traffic cannot simply be redirected through a traditional proxy.
Therefore, even if FortiOS supports HTTPS-based FortiGuard communication for some services, do not assume all FortiGuard traffic can traverse the same proxy.
For environments without direct Internet access, FortiManager can provide local update and rating services.
Architecture:
                 Internet
                    |
                    v
              FortiManager
             Local Services
              /          \
             /            \
       Update              Rating
       AV/IPS              Web/Spam
          \                  /
           \                /
              FortiGate
This is particularly useful in:
- Closed networks
- Restricted networks
- OT environments
- Data centers
- Air-gapped or partially isolated environments
Example:
config system central-management
    set type fortimanager
    set fmg 192.168.254.200
    config server-list
        edit 1
            set server-type update
            set server-address 192.168.254.200
        next
        edit 2
            set server-type rating
            set server-address 192.168.254.200
        next
    end
    set fmg-update-port 443
    set include-default-servers enable
end
| Server Type | Typical Purpose | 
|---|---|
| update | AV / IPS and update-related services | 
| rating | Web filtering / AntiSpam rating | 
set include-default-servers enable
This allows FortiGate to retain default FortiGuard servers as fallback when configured FortiManager connectivity is unavailable.
Older FortiOS/FortiManager implementations may use different ports or mechanisms. Always verify the target release.
FortiGuard can use Anycast to improve global connectivity and reduce latency.
                    FortiGuard
                 Global Anycast IP
                         |
          +--------------+--------------+
          |              |              |
        Region A       Region B       Region C
          |              |              |
       FortiGate       FortiGate       FortiGate
Instead of manually selecting a distant FortiGuard server, DNS/Anycast routing can direct the FortiGate toward an appropriate service location.
globalupdate.fortinet.net
globalguardservice.fortinet.net
update.fortiguard.net
service.fortiguard.net
securewf.fortiguard.net
usupdate.fortinet.net
usguardservice.fortinet.net
Non-Anycast examples:
usupdate.fortiguard.net
usservice.fortiguard.net
ussecurewf.fortiguard.net
euupdate.fortinet.net
euguardservice.fortinet.net
config system fortiguard
    set update-server-location automatic
end
Possible location behavior can include:
automatic
usa
eu
automatic is generally associated with selecting an appropriate FortiGuard location, commonly using Anycast where supported.
Example:
config system fortiguard
    set protocol https
    set port 443
end
Depending on the FortiGuard service and FortiOS release, other ports/protocols may be involved.
Commonly encountered:
HTTPS  443
DNS    53
UDP    8888
Do not test FortiGuard connectivity by checking only TCP/443.
A complete troubleshooting process should consider:
DNS
 |
 +--> FortiGuard hostname resolution
 |
Network
 |
 +--> Routing
 |
Firewall
 |
 +--> Required ports
 |
TLS
 |
 +--> Certificate / trust
 |
FortiGuard
 |
 +--> Service/license availability
FortiGate can cache certain FortiGuard results locally.
The cache can store web filtering results locally for a configurable period.
Concept:
Client
  |
  v
FortiGate
  |
  +---- Cache hit ----> Local decision
  |
  +---- Cache miss ---> FortiGuard
Example:
config system fortiguard
    set webfilter-cache enable
    set webfilter-cache-ttl 3600
    set webfilter-timeout 15
end
config system fortiguard
    set antispam-cache enable
    set antispam-cache-ttl 1800
    set antispam-cache-mpercent 2
    set antispam-timeout 7
end
