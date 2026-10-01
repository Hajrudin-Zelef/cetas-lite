---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-cannot-access-the-fortigate-web-admin-interface-eb410b1f-1
title: "fortigate-3-troubleshooting-tip-cannot-access-the-fortigate-web-admin-interface--eb410b1f"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-cannot-access-the-fortigate-web-admin-interface--eb410b1f.md
source_anchor: ""
source_lines: [1, 131]
sha256: fb04f087738e54ebcf61abd6c766c5a9fe29c60b1b6ed4ad293e73765f05d29a
---

# fortigate-3-troubleshooting-tip-cannot-access-the-fortigate-web-admin-interface--eb410b1f

Troubleshooting Tip: Cannot access the FortiGate web admin interface (GUI)
Description
This article describes some possible causes for non-working GUI access. In some cases, it is possible to reach the FortiGate unit through a Ping, Telnet, or SSH, yet not through the web admin GUI.
Scope
FortiGate.
Solution
Shortlist:
- The HTTP/HTTPS service is not enabled on the interface.
- Trusted hosts are enabled on all admin users, and the source IP that is in use is not listed as a trusted host.
- MTU along the path.
- GUI PORT (HTTP/HTTPS) is used on another service. Admin HTTPS port is 443, which conflicts with the DNS Server DNS-over-HTTPS.
- VIP Overlap.
- Local-in policy with a deny action.
- Issues with the HTTPS Daemon.
- Routing issue.
- TLS/SSL issue.
- HA Monitor interface.
- VDOMs issue.
- ACME listening port conflict (TCP 80 & TCP 443).
- IKE/IPsec port conflict if TCP port 443 is used as transport mode.
- Banned cipher suite.
- Traffic not reaching FortiGate.
- Blank GUI page due to high CPU or memory usage.
- Two different interfaces with IP addresses assigned in the same subnet on the same VDOM.
- GUI certificate validity.
- GUI certificate key length.
Details:
To initiate access, start by pinging the management IP address to verify that the FortiGate is actively listening on the specified management IP. This step helps isolate whether the FortiGate is operational and responsive, assuming that ping is enabled on the management port.
When the endpoint is unable to ping the FortiGate interface IP address, the next step is to determine if the endpoint is in the same broadcast domain using ipconfig /all. In some cases, the IP address must be assigned statically on the endpoint device if DHCP is not enabled on the FortiGate interface or when the endpoint is unable to connect to the internal DHCP server.
Connect to FortiGate via SSH through PuTTY as demonstrated below, ensuring that SSH is enabled on the management port:
If SSH access is unsuccessful, then use the following article to access the FortiGate via console cable and move on to the next steps:
Technical Tip: How to connect to the FortiGate and FortiAP console port
- Interface settings. GUI access, HTTP, and/or HTTPS have to be enabled on the interface.
CLI commands:
config system interface
    edit <interface name>
        set allowaccess ping http https
end
- Trusted host configuration. If 'trusted hosts' are configured, the IP address of the computer used for the GUI access must be allowed as a trusted host. A whole subnet can be allowed as a trusted host. By default, trusted host settings are not configured, and administrative access is not restricted to any specific user IP addresses. Sample trusted host configuration:
GUI: Define Trusted hosts by going to System -> Administrators.
- See the inbound request: there is no reply.
- The debug flow will show the 'policy 0, drop'.
- It will not indicate anything about TRUSTED HOSTS.
- Debug HTTPS will not show any log.
diagnose sniffer packet any "host 10.1.1.10 and port 443" 4 0
port3 in 10.1.1.1.55826 -> 10.1.1.10.443: syn 3127611448
port3 in 10.1.1.1.55825 -> 10.1.1.10.443: syn 2440393440
Debug flow:
msg="iprope_in_check() check failed on policy 0, drop"
Check if trusted hosts are configured in all admin users, which is the case:
show sys admin
config system admin
edit "admin"
set trusthost1 10.1.7.0 255.255.255.0
set accprofile "super_admin"
...omit
next
end
Add the source IP as a trusted host:
config system admin
(admin) # edit admin
(admin) # set trusthost2 10.1.1.0/24
(admin) # show
config system admin
edit "admin"
set trusthost1 10.1.7.0 255.255.255.0
set trusthost2 10.1.1.0 255.255.255.0
next
- MTU along the path. After the first few synchronisation and handshake packets, the web admin GUI HTTP and HTTPS packets can become larger than 1500 bytes.
For example:
When a FortiGate network interface is connected to a network segment that supports such extended-size packets. 
For Telnet or SSH, packets typically remain of a smaller size.
To then be able to use the web admin GUI, the fragmentation must be allowed at certain points in the network infrastructure (points where a jumbo frame reaches a network segment that does not support it), or jumbo frames must be allowed along the whole communication path.
Note about Jumbo frames: Jumbo frames are packets that are larger than the standard 1500 maximum transmission unit (MTU) size. Jumbo frames increase data transfer speeds by carrying more data per frame, reducing the overhead from headers.
All networks that carry jumbo frames must have network units that all support jumbo frames. Otherwise, jumbo frames will be dropped when they reach network devices that do not support them.
- Admin access ports. By default, for admin login via GUI, the HTTPS port is configured to 443 and the HTTP port to 80.
If those default settings are changed, access to the GUI will not be possible without specifying the custom port used at the end of the URL. To verify which HTTPS/HTTP ports are configured for admin access:
show full | grep admin-port
    set admin-port 8080  <-- The ports were changed from the default.
show full | grep admin-sport
set admin-sport 8443 <-- The ports were changed from the default.
CLI Reference:
config system dns-server
    edit "fortilink" <-- It can be any FortiGate Interface where the user is trying to log in. 
        set doh disable <--     
    next
end
If the default ports have been changed, consider directly accessing the GUI using the specific port that is currently defined: http(s)://<address_of_appliance>:<custom port>.
For example: http://192.168.0.101:222, where 222 is the non-default port used to access the GUI via HTTP. If the ports need to be changed to a new value or the default value, use the following syntax for HTTP access:
config system global
set admin-port <integer>
end
- The existing virtual IP is overriding the admin HTTP or HTTPS ports.
When a Virtual IP (VIP) has the same IP address as the FortiGate interface and forwards the same ports used for HTTP/HTTPS access (example 80 or 443), the VIP will override the administrative access.
This should either be removed or changed such that it does not overlap with FortiGate HTTP/HTTPS ports. This can be verified by checking the VIP list on FortiGate (Policy & Objects -> Virtual IPs) or running the debug flow.
Troubleshooting:
The sniffer will show that the INBOUND request has been forwarded to another IP.
HUB01 # diagnose sniffer packet any "host 192.168.247.1 and port 443" 4 0 a
 port2 in 192.168.247.1.57530 -> 192.168.247.20.443: syn 174545504 
 port4 out 192.168.247.1.57530 -> 10.255.255.11.443: syn 174545504 <-- Same Source IP and Same Source Port.
A debug flow will show the traffic matches a VIP:
diagnose debug reset
diagnose debug flow filter clear
diagnose debug flow show function-name enable
diagnose debug flow show iprope enable
diagnose debug flow filter saddr 192.168.247.1 <--Adjust to the source IP of the testing PC.
diagnose debug flow filter daddr 192.168.247.20 <-- Adjust to the GUI FortiGate IP.
diagnose debug flow filter dport 443 <-- Adjust the port if it is not the default port.
diagnose debug console timestamp enable
diagnose debug flow trace start 1
diagnose debug enable
Among all the lines it will receive on the DEBUG, the following will appear in the first lines:
msg="find DNAT: IP-10.255.255.11, port-0(fixed port)"
This indicates there is a VIP matching the request. Check the VIPs on the GUI under Policy & Objects -> Virtual IPs.
- Check if any local-in policy is configured to deny access to the related interface.
config firewall local-in-policy
show full
edit 1
set intf "wan1"
set srcaddr "all"
set srcaddr-negate disable
set dstaddr-negate disable
set action deny
set service "HTTPS"
set service-negate disable
set schedule ''
set status enable
set comments ''
next
end
end
