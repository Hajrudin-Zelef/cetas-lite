---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fo-60be8d6e-3
title: "FortiGuard configuration"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license", "training"]
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-fo-60be8d6e.md
source_anchor: ""
source_lines: [567, 795]
sha256: bc2b534ca8582a9825cf87a9fe209b91471ae1c74347f109d41f6487b111e0a7
---

# FortiGuard configuration

        set schedule always
        set logtraffic all
        set users guest1
    next
end
config user local
    edit guest1
        set type password
        set passwd <PASSWORD>
    next
end
config authentication scheme
    edit local-basic
        set method basic
        set user-database local-user-db
    next
end
config authentication rule
    edit local-basic-rule
        set srcaddr all
        set ip-based disable
        set active-auth-method local-basic
    next
end
config firewall policy
    edit 1
        set srcintf port3
        set dstintf port1
        set srcaddr all
        set dstaddr all
        set action accept
        set schedule always
        set service dns
        set fsso disable
        set nat enable
    end
end
On the FortiGate using the proxy:
config system fortiguard
    set proxy-server-ip 192.168.254.251
    set proxy-server-port 8080
    set proxy-username guest1
    set proxy-password <PASSWORD>
end
Test FortiCloud/FortiGuard-related communication as appropriate for the target FortiOS release.
🔐 Never place real production credentials inside shared documentation or GitHub repositories.
FortiOS firmware images include Fortinet Internet Service Database (ISDB) information.
Depending on the image/update state, the built-in ISDB can be:
Partial / lightweight ISDB
        |
        v
Core Fortinet services
Full ISDB
        |
        v
Large database with current service objects
The lightweight database helps policies and policy routes referencing Fortinet services continue functioning after a firmware upgrade before the full database is retrieved.
diagnose firewall internet-service list
This displays ISDB entries.
config firewall policy
    edit 1
        set srcintf port1
        set dstintf port3
        set srcaddr all
        set dstaddr all
        set internet-service enable
        set internet-service-id 1245187 1245326 1245324 1245325
        set action accept
        set schedule always
        set logtraffic all
        set fsso disable
    end
end
ISDB IDs are release/database dependent. Do not assume an ID has the same meaning across FortiOS versions.
After a firmware upgrade/reboot, FortiGate can perform an automatic ISDB update after startup.
Conceptually:
Firmware Upgrade
       |
       v
FortiGate Reboot
       |
       v
Initial built-in ISDB
       |
       v
Automatic update
       |
       v
Current ISDB
The notes specify an approximate startup update delay of:
~5 minutes
Check:
diagnose autoupdate versions | grep internet -a 6
Air-gapped environments may not permit direct Internet access.
Common examples:
- OT networks
- Industrial environments
- Critical infrastructure
- Isolated data centers
Architecture:
Connected Environment
        |
        | Download license/package
        v
Offline Transfer
        |
        v
Air-Gapped FortiGate
FortiGuard packages such as AV and IPS may need to be manually transferred.
Licensing can also be manually transferred on supported hardware appliances.
According to the supplied training notes, manual licensing for air-gapped environments is supported for supported FortiGate hardware appliances running FortiOS 7.2.0 or later, but not FortiGate VM appliances. Verify current platform support before deployment.
Example:
execute restore manual-license ftp a.lic 192.168.20.200
The exact transfer method and syntax should be checked against the target FortiOS release.
Example:
execute restore ips tftp nids-720-19.261.pkg 192.168.20.200
Then debug update processing:
diagnose debug application updated -1
diagnose debug enable
After troubleshooting:
diagnose debug disable
A FortiGate in a restricted environment might use:
config system fortiguard
    set protocol https
    set port 443
    set load-balance-servers 1
    set auto-join-forticloud disable
    set update-server-location usa
end
The exact configuration depends on the architecture and whether any local update infrastructure exists.
The following options may appear in config system fortiguard depending on FortiOS version:
set update-ffdb enable
set update-uwdb enable
set update-extdb enable
set update-build-proxy enable
Used for FortiGuard/FortiOS database update functionality.
Unified Weight Database.
Can contribute to threat weighting and security intelligence.
Provides extended security database information.
Can be relevant when using an explicit web proxy for supported update workflows.
Version Note: Do not blindly copy all options into production. CLI availability and behavior are FortiOS-release dependent.
Example:
config system fortiguard
    set webfilter-force-off disable
    set webfilter-cache enable
    set webfilter-cache-ttl 3600
    set webfilter-timeout 15
end
Concept:
URL Request
    |
    v
FortiGate
    |
    +---- Cached result ----> Apply rating
    |
    +---- No cache ---------> FortiGuard
                                  |
                                  v
                              Rating result
A consolidated example:
config system fortiguard
    set protocol https
    set port 443
    set load-balance-servers 1
    set update-server-location automatic
    set update-ffdb enable
    set update-uwdb enable
    set update-extdb enable
    set update-build-proxy enable
    set antispam-force-off disable
    set antispam-cache enable
    set antispam-cache-ttl 1800
    set antispam-cache-mpercent 2
    set antispam-timeout 7
    set outbreak-prevention-force-off disable
    set outbreak-prevention-cache enable
    set outbreak-prevention-cache-ttl 300
    set outbreak-prevention-cache-mpercent 2
    set outbreak-prevention-timeout 7
    set webfilter-force-off disable
    set webfilter-cache enable
    set webfilter-cache-ttl 3600
    set webfilter-timeout 15
    set sdns-server-ip 208.91.112.220
    set sdns-server-port 53
    set source-ip 0.0.0.0
end
Important: Treat this as a reference template, not a universal copy/paste configuration.
                    FortiGuard Problem
                           |
                           v
                    Is license valid?
                      /          \
                    NO            YES
                    |              |
              Fix licensing        v
                              DNS resolution?
                              /          \
                            NO            YES
                            |              |
                       Fix DNS             v
                                    Routing working?
                                      /       \
                                    NO         YES
                                    |           |
                               Fix routing      v
                                         Required ports?
                                          /          \
                                        NO            YES
                                        |              |
                                   Allow traffic       v
                                                TLS validation?
                                                 /        \
                                               NO          YES
                                               |             |
                                      Check certificates     v
                                                        Service issue?
                                                         /       \
                                                       YES       NO
                                                       |          |
                                              Debug service      Check
                                              communication     FortiGuard
| Purpose | Command | 
|---|---|
| FortiGuard/service communication | diagnose sys service-communication | 
| Autoupdate debugging | diagnose debug application updated -1 | 
| Enable debug | diagnose debug enable | 
| Disable debug | diagnose debug disable | 
