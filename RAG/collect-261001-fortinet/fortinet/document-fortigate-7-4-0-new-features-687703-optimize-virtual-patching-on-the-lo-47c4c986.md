---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-new-features-687703-optimize-virtual-patching-on-the-lo-47c4c986
title: "document-fortigate-7-4-0-new-features-687703-optimize-virtual-patching-on-the-lo-47c4c986"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2023-11-07"]
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-new-features-687703-optimize-virtual-patching-on-the-lo-47c4c986.md
source_anchor: ""
source_lines: [1, 39]
sha256: cbb531d0f461ed392c6f4e517f56d6bf1073984cb990695a303a0a869bdfc73c
---

# document-fortigate-7-4-0-new-features-687703-optimize-virtual-patching-on-the-lo-47c4c986

Optimize virtual patching on the local-in interface 7.4.2
|  | This information is also available in the FortiOS 7.4 Administration Guide: | 
Virtual patching is a method of mitigating vulnerability exploits by using the FortiGate's IPS engine to block known vulnerabilities. Virtual patching can be applied to traffic destined to the FortiGate by applying IPS signatures to the local-in interface using local-in policies.
When virtual patching is enabled in a local-in policy, the IPS engine queries the FortiGuard API server to:
- 
                                                    Obtain a list of vulnerabilities targeting the FortiGate on a particular version
- 
                                                    Optimize scanning by determining whether the session destined to the local-in interface on the FortiGate requires a scan. The session's port number and protocol are used to identify the services to be tagged.
If a tagged session lacks vulnerability signatures for the FortiOS version, then the IPS engine bypasses the session. This optimizes performance by only scanning and dropping sessions that are exploiting a vulnerability. Only SSL VPN and web GUI services are optimized for tagging. Other protocols are scanned, but do not have tag optimization.
Example
In this example, virtual patching is enabled for the local-in policy and the following scenarios are described:
- 
                                                    FortiGate with an SSL VPN vulnerability
- 
                                                    FortiGate with a web GUI vulnerability
- 
                                                    FortiGate with both an SSL VPN and web GUI vulnerability
To enable virtual patching:
- 
                                                    Enable virtual patching in the local-in policy: config firewall local-in-policy
    edit 1
        set intf "port2"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set service "ALL"
        set schedule "always"
        set virtual-patch enable
    next
endBecause the IPS engine can currently only tag services related to SSL VPN and web GUI signatures, all other protocols are scanned when service is set toALL . However, you can bypass scanning of other protocols, such as SSH and FTP, by settingservice to onlyHTTPS .
- 
                                                    Observe the outcome of the following scenarios: 
  - 
                                                            In this example, FortiOS has an SSL VPN vulnerability. The IPS engine drops SSL VPN traffic to the local-in interface on the FortiGate and bypasses web GUI traffic. Traffic for other services is scanned and passed to the interface. Following is a log of the SSL VPN traffic that was dropped because of the vulnerability. Bypassed web GUI traffic did not generate any logs. # diagnose ips vpatch fmwp-status Enabled FMWP signatures: 3 10002887 FortiOS.SSL-VPN.Heap.Buffer.Overflow. 1: date=2023-11-07 time=14:53:44 eventtime=1699325624346021995 tz="+1200" logid="0419016384" type="utm" subtype="ips" eventtype="signature" level="alert" vd="root" severity="critical" srcip=10.1.100.22 srccountry="Reserved" dstip=10.1.100.1 dstcountry="Reserved" srcintf="port2" srcintfrole="undefined" dstintf="root" dstintfrole="undefined" sessionid=284 action="dropped" proto=6 service="HTTPS" policyid=1 attack="FortiOS.SSL-VPN.Heap.Buffer.Overflow." srcport=53250 dstport=11443 hostname="myfortigate.example" url="/error" httpmethod="POST" direction="outgoing" attackid=10002887 ref="http://www.fortinet.com/ids/VID10002887" incidentserialno=99614721 msg="vPatch: FortiOS.SSL-VPN.Heap.Buffer.Overflow." crscore=50 craction=4096 crlevel="critical"
  - 
                                                            In this example, FortiOS has a web GUI vulnerability. The IPS engine drops web GUI traffic to the local-in interface on the FortiGate and bypasses SSL VPN traffic. Traffic for other services is scanned and passed to the interface. Following is a log of the web GUI traffic that was dropped because of the vulnerability. Bypassed SSL VPN traffic did not generate any logs. # diagnose ips vpatch fmwp-status Enabled FMWP signatures: 2 10002156 FortiOS.NodeJS.Proxy.Authentication.Bypass. 10002890 FortiOS.HTTPD.Content-Length.Memory.Corruption. 1: date=2023-11-07 time=14:55:15 eventtime=1699325715311370215 tz="+1200" logid="0419016384" type="utm" subtype="ips" eventtype="signature" level="alert" vd="root" severity="critical" srcip=10.1.100.22 srccountry="Reserved" dstip=10.1.100.1 dstcountry="Reserved" srcintf="port2" srcintfrole="undefined" dstintf="root" dstintfrole="undefined" sessionid=304 action="dropped" proto=6 service="HTTPS" policyid=1 attack="FortiOS.NodeJS.Proxy.Authentication.Bypass." srcport=53622 dstport=443 hostname="127.0.0.1:9980" url="/api/v2/cmdb/system/admin" agent="Node.js" httpmethod="GET" direction="outgoing" attackid=10002156 ref="http://www.fortinet.com/ids/VID10002156" incidentserialno=99614722 msg="vPatch: FortiOS.NodeJS.Proxy.Authentication.Bypass." crscore=50 craction=4096 crlevel="critical"
  - 
                                                            In this example, FortiOS has an SSL VPN and a web GUI vulnerability. The IPS engine drops both SSL VPN and web GUI traffic to the local-in interface on the FortiGate. Traffic for other services is scanned and passed to the interface. Following is a log of the SSL VPN and web GUI traffic that was dropped because of the vulnerability. # diagnose ips vpatch fmwp-status Enabled FMWP signatures: 3 10002156 FortiOS.NodeJS.Proxy.Authentication.Bypass. 10002887 FortiOS.SSL-VPN.Heap.Buffer.Overflow. 10002890 FortiOS.HTTPD.Content-Length.Memory.Corruption. 1: date=2023-11-07 time=06:42:44 eventtime=1699296164649894963 tz="+1200" logid="0419016384" type="utm" subtype="ips" eventtype="signature" level="alert" vd="root" severity="critical" srcip=10.1.100.22 srccountry="Reserved" dstip=10.1.100.1 dstcountry="Reserved" srcintf="port2" srcintfrole="undefined" dstintf="root" dstintfrole="undefined" sessionid=1094 action="dropped" proto=6 service="HTTPS" policyid=1 attack="FortiOS.SSL-VPN.Heap.Buffer.Overflow." srcport=44164 dstport=10443 hostname="myfortigate.example" url="/error" httpmethod="POST" direction="outgoing" attackid=10002887 ref="http://www.fortinet.com/ids/VID10002887" incidentserialno=116392250 msg="vPatch: FortiOS.SSL-VPN.Heap.Buffer.Overflow." crscore=50 craction=4096 crlevel="critical" 2: date=2023-11-07 time=06:42:09 eventtime=1699296129458704870 tz="+1200" logid="0419016384" type="utm" subtype="ips" eventtype="signature" level="alert" vd="root" severity="critical" srcip=10.1.100.22 srccountry="Reserved" dstip=10.1.100.1 dstcountry="Reserved" srcintf="port2" srcintfrole="undefined" dstintf="root" dstintfrole="undefined" sessionid=1066 action="dropped" proto=6 service="HTTPS" policyid=1 attack="FortiOS.NodeJS.Proxy.Authentication.Bypass." srcport=42352 dstport=443 hostname="127.0.0.1:9980" url="/api/v2/cmdb/system/admin" agent="Node.js" httpmethod="GET" direction="outgoing" attackid=10002156 ref="http://www.fortinet.com/ids/VID10002156" incidentserialno=116392236 msg="vPatch: FortiOS.NodeJS.Proxy.Authentication.Bypass." crscore=50 craction=4096 crlevel="critical"
-
