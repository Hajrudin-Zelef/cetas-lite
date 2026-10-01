---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-unable-to-access-fortigate-gui-because-of-high-c-272c531c
title: "fortigate-3-troubleshooting-tip-unable-to-access-fortigate-gui-because-of-high-c-272c531c"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2026-06-20"]
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-unable-to-access-fortigate-gui-because-of-high-c-272c531c.md
source_anchor: ""
source_lines: [1, 4]
sha256: a9eadd3b5797b844ffa9185285f04347cdde647d352e71a616b301047d863854
---

# fortigate-3-troubleshooting-tip-unable-to-access-fortigate-gui-because-of-high-c-272c531c

Troubleshooting Tip: Unable to access FortiGate GUI because of high CPU due to httpsd process
| Description | This article describes how to get back into the FortiGate GUI after multiple HTTPS processes caused high CPU utilization. | 
| Scope | FortiGate. | 
| Solution | Scenario 1. High CPU is caused by GUI authentication requests.     Starting in FortiOS v7.6.4, a new daemon named http_authd has been introduced to manage administrative authentication processes.  The http_authd daemon centralizes all authentication activities related to administrative access on FortiGate devices. By consolidating these functions into a dedicated process, FortiOS improves the efficiency, consistency, and scalability of authentication handling.   In FortiOS v7.6.4 and above, the output:       config system interface     edit <interface_name>               unselect allowaccess https http     next  end   fnsysctl killall httpsd   Note: If the issue persists, perform hardening steps. Ensure that the external interface is disconnected before accessing the GUI again. This helps reduce high CPU usage caused by external attacks. Then perform the following actions:   Fix: Versions 7.6.7 and 8.0.0 introduce a new DoS mitigation mechanism to help protect against brute-force attempts and Slowloris-style attacks. Bug ID: 1256988: Brute-force attacks triggered a lot of leaving http_authd processes running and causing memory usage to steadily increase. Scenario 2. High CPU is caused by FortiFlow application lookup failures: Another possible cause of high CPU utilization by the httpsd process is repeated FortiFlow application lookup failures. Symptoms. The following symptoms may be observed:  Verify CPU utilization: Run: diagnose sys top 2 99 Example output: httpsd 13039 D 58.4 2.5 2 httpsd 13025 R 57.0 2.0 2 forticron 153 S 2.7 1.5 2 miglogd 240 S 0.5 1.6 2 node 164 S 0.3 3.1 3 In the affected condition, one of the httpsd processes may remain in the D state (uninterruptible sleep) while consuming a high percentage of CPU resources. Collect httpsd debug: Collect httpsd debug logs using the following commands: diagnose debug reset diagnose debug application httpsd -1 diagnose debug console timestamp enable diagnose debug enable The debug output may repeatedly display messages similar to the following: 2026-06-20 15:10:05 [httpsd] build_utm_app_lookup -- FortiFlow application query failed (-1)  2026-06-20 15:10:06 [httpsd] build_utm_app_lookup -- FortiFlow application query failed (-1)  2026-06-20 15:10:07 [httpsd] build_utm_app_lookup -- FortiFlow application query failed (-1) These messages indicate that the GUI is continuously attempting to perform FortiFlow application lookup requests that are failing, resulting in excessive processing by the httpsd daemon. Workaround: Disable FortiFlow application lookup: config log gui-display     set resolve-apps disable endAfter applying the workaround:  Related articles: Technical Tip: Regularly audit and restrict open ports on FortiGate public interfaces Technical Tip: System administrator best practices for FortiGate and FortiProxy Troubleshooting Tip: High CPU usage due to httpsd daemon on FortiGate Technical Tip: Brute-Force Attacks may cause the device to enter Conserve Mode with multiple http_authd daemons after upgrading to FortiOS v7.6.6 |
