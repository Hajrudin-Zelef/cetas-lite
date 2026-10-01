---
id: collect-261001-fortinet/fortinet/questions-58880-fortigate-800c-is-not-displaying-the-data-for-ids-sniffer-traffi-bf4fae7a
title: "questions-58880-fortigate-800c-is-not-displaying-the-data-for-ids-sniffer-traffi-bf4fae7a"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-58880-fortigate-800c-is-not-displaying-the-data-for-ids-sniffer-traffi-bf4fae7a.md
source_anchor: ""
source_lines: [1, 29]
sha256: b0f64660608e61b851b1207f845bfad5f23fc71f8633ebcc90e717f220663868
---

# questions-58880-fortigate-800c-is-not-displaying-the-data-for-ids-sniffer-traffi-bf4fae7a

Configure sniffer.
        set id {integer}   Sniffer ID. range[0-9999]
        set status {enable | disable}   Enable/disable the active status of the sniffer.
        set logtraffic {all | utm | disable}   Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy.
                all      Log all sessions accepted or denied by this policy.
                utm      Log traffic that has a security profile applied to it.
                disable  Disable all logging for this policy.
        set ipv6 {enable | disable}   Enable/disable sniffing IPv6 packets.
        set non-ip {enable | disable}   Enable/disable sniffing non-IP packets.
        set interface {string}   Interface name that traffic sniffing will take place on. size[35] - datasource(s): system.interface.name
        set host {string}   Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240). size[63]
        set port {string}   Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200). size[63]
        set protocol {string}   Integer value for the protocol type as defined by IANA (0 - 255). size[63]
        set vlan {string}   List of VLANs to sniff. size[63]
        set application-list-status {enable | disable}   Enable/disable application control profile.
        set application-list {string}   Name of an existing application list. size[35] - datasource(s): application.list.name
        set ips-sensor-status {enable | disable}   Enable/disable IPS sensor.
        set ips-sensor {string}   Name of an existing IPS sensor. size[35] - datasource(s): ips.sensor.name
        set dsri {enable | disable}   Enable/disable DSRI.
        set av-profile-status {enable | disable}   Enable/disable antivirus profile.
        set av-profile {string}   Name of an existing antivirus profile. size[35] - datasource(s): antivirus.profile.name
        set webfilter-profile-status {enable | disable}   Enable/disable web filter profile.
        set webfilter-profile {string}   Name of an existing web filter profile. size[35] - datasource(s): webfilter.profile.name
        set spamfilter-profile-status {enable | disable}   Enable/disable spam filter.
        set spamfilter-profile {string}   Name of an existing spam filter profile. size[35] - datasource(s): spamfilter.profile.name
        set dlp-sensor-status {enable | disable}   Enable/disable DLP sensor.
        set dlp-sensor {string}   Name of an existing DLP sensor. size[35] - datasource(s): dlp.sensor.name
        set ips-dos-status {enable | disable}   Enable/disable IPS DoS anomaly detection.
You can use this link for trouble shooting click here
