---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-cannot-access-the-fortigate-web-admin-interface-eb410b1f-2
title: "fortigate-3-troubleshooting-tip-cannot-access-the-fortigate-web-admin-interface--eb410b1f"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2024-08-28"]
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-cannot-access-the-fortigate-web-admin-interface--eb410b1f.md
source_anchor: ""
source_lines: [132, 209]
sha256: 43b78a61b9e381f2751b4bce40056025ff52c81d129f41573c91c9417baa4545
---

# fortigate-3-troubleshooting-tip-cannot-access-the-fortigate-web-admin-interface--eb410b1f

- If the issue happens after creating IPSEC DIAULUP VPN, check if ike-tcp port is changed to port 443 from the default port 4500. This should be checked if the firewall is on FortiOS v7.6.1 or later.
In v7.6.3, a warning will be presented to remind users that the HTTPS port conflicts with the ike-tcp-port.
- Restart the HTTPS Daemon: If none of the processes above fixed the issue, try restarting the HTTPS Daemon.
- Routing Issue: If there are two default routes to 0.0.0.0 with the same distance and priority (for example, 5), it will not be possible to access the GUI. To find default routes in the CLI, run 'get router info routing-table database | grep 0.0.0.0'.
- Change the distance for one route to 10 as an example.
- Now it should be possible to log in to the GUI, and it should not freeze or hang.
- Disable the setting: Retrieve the default Gateway from the server on the internal network interface.
- If trying to access FortiGate using the WAN interface, make sure that the route is active or valid in the routing table.
By default, TLS 1.1 and TLS 1.2 are enabled when accessing the FortiGate GUI via a web browser.
To verify what version is enabled, run the following commands:
config system global
get | grep 'min-proto'
-  HA Monitor Interface GUI Issue: The FortiGate can reach the GUI via a username/password from the local interface IP, like 192.168.1.99, but it cannot connect through other specified LAN interfaces using the same credentials. It is essential to verify if these interfaces are part of the HA Monitor interface, and if they are, eliminate them; additionally, create a specific interface for HA Monitor. It will enable user verification for GUI access.
Notes:
- When an interface is configured as an HA monitor, its primary purpose is to monitor the health and status of the HA cluster. This configuration may impose certain restrictions, such as limited management access.
- Access Restrictions: The HA monitor interface may not be set up to allow management access (for example, GUI access) because of its monitoring role.
-  VDOMs issue:
config global
config system global
get | grep 'min-proto'
To change this setting from the CLI:
config system global
set admin-https-ssl-versions (shift + ?) <- To list the available TLS version.
tlsv1-0 TLS 1.0.
tlsv1-1 TLS 1.1.
tlsv1-2 TLS 1.2.
set admin-https-ssl-versions tlsv1-2 <- With this setting, only TLS 1.2 is allowed.
end
From v6.4, tlsv1-0 is no longer supported and instead, tlsv1-3 was introduced:
config system global
set admin-https-ssl-versions
tlsv1-1 TLS 1.1.
tlsv1-2 TLS 1.2.
    tlsv1-3 TLS 1.3.
end 
-  ACME listening port conflict (TCP 80 & TCP 443): If the FortiGate GUI admin is listening to the WAN interface, which is also the ACME listening interface, then it is necessary not to use the default (TCP 80/443) port for the GUI management.
config system acme
    set interface "wan1" <-----
end
config system interface
    edit "wan1"
        set allowaccess ping https ssh http telnet fgfm <-----
    next
end
-  The default IKE-TCP port value of 443 applies only to new FortiGate configurations running FortiOS 7.6.1 or later. When upgrading to FortiOS 7.6.1 or later, the pre‑existing ike‑tcp‑port setting is preserved. If both the ike‑tcp‑port and the administrative port are set to 443, the administration page cannot be accessed through the interface where the IPsec tunnel terminates. To resolve this, enable administrative GUI access on a different interface or change the administrative port.
-  Banned cipher suite: Administrators can configure the system to block or disable weak ciphers for SSH and HTTPS connections. However, if the cipher suite used by the web browser to establish the HTTPS session is disabled, access to the web-based management interface (web GUI) will be blocked:
config system global
    set admin-https-ssl-banned-ciphers <>
If the admin-https-ssl-banned-ciphers set, try to unset or change the banned only weak cipher encryption. Here are the specific cipher suites supported by each TLS version: FortiGate encryption algorithm cipher suites
-  Traffic not Reaching FortiGate: If packets are not reaching the web server or interface, or responses from the web interface are not transmitted back to the user, it will not be possible to access the GUI. To verify, run a packet capture via CLI and see how the TCP handshake is made and if follow-up packets are visible.
The command is:
diagnose sniffer packet any 'port 8443' 4 0 a <------ Replace the port if the web interface is reachable via a different port.
Filters, adding a host IP, if the connecting IP is known, can be added:
diagnose sniffer packet any 'host 192.168.48.2 and port 8443' 4 0 a
Example output:
FGT# diagnose sniffer packet any 'port 8443' 4 0 a
Using Original Sniffing Mode
interfaces=[any]
filters=[port 8443]
2024-08-28 16:36:30.527027 port1 in 192.168.48.2.56662 -> 10.191.19.1728443: syn 1919112407 
2024-08-28 16:36:30.527238 port1 out 10.191.19.1728443 -> 192.168.48.2.56662: syn 3263049518 ack 1919112408 
2024-08-28 16:36:30.527648 port1 in 192.168.48.2.56662 -> 10.191.19.1728443: ack 3263049519
Legacy  (for v5.6 and above only). admin-server-cert (for v5.6 and above only).
Enable the following debug and try to access the GUI again:
diagnose debug application httpsd -1
diagnose debug enable
-  Blank GUI page due to high CPU or memory usage: If the GUI opens with a blank (white) screen and no error message, such as connection refused or timeout, check system CPU and memory utilization. When CPU or memory is high due to the httpsd process, the httpsd process may accept connections but fail to render the GUI. If httpsd shows high usage, restart it as described in point 8. Verify this using the command below:
diagnose system top
In some cases, CPU and memory usage remain stable, but the GUI page is blank. This behavior is often caused by continuous brute-force login attempts. To verify if there are multiple attempts, run the following command:
diagnose alertconsole list | grep login
To address this issue, disable HTTPS access on the Public Interface if it is not needed and/or apply security hardening for admin access. Refer to the following article: Technical Tip: Hardening best practices: Secure network and devices.
- Two different interfaces with IP address assigned in the same subnet on same VDOM:  If two different interfaces are assigned IP in the same network in the same VDOM, it can create issues with intermittent GUI access. It is recommended to avoid assigning IP in same network to two different interfaces. This will only occur when subnet overlap is allowed (allow-subnet-overlap is enabled) as described in Technical Tip: Enable subnet overlap to set IP addresses of multiple interfaces in the same subnet.
- The GUI certificate might have expired. Check the certificate below to make sure it is valid:
config system global
    set admin-server-cert <new_cert>
end
- The FortiGate GUI cannot be accessed after upgrading to FortiOS 7.6.1, 7.4.8, or 7.2.11 if the admin server certificate is using an RSA key of less than 2048 bits: GUI cannot be accessed when using a server certificate with an RSA 1024 bit key.
Related documents:
