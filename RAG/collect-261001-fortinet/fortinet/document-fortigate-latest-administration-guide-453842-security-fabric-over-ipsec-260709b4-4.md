---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4-4
title: "document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4.md
source_anchor: ""
source_lines: [427, 447]
sha256: 9c5a54c90b78c59a13b7736cb08a153578f12f8ef636108bb37860267516c686
---

# document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4

To authorize the downstream FortiGate (HQ2) on the root FortiGate (HQ1):
- 
                                                    In the root FortiGate (HQ1), go to System > Firmware & Registration. The table highlights the connected FortiGate with its serial numbers that is unauthorized.
- 
                                                    Select the unauthorized device and click Authorization > Authorize. After authorization, the downstream FortiGate (HQ2) appears in the Security Fabric widget. This means the downstream FortiGate (HQ2) has successfully joined the Security Fabric.
To check the Security Fabric over IPsec VPN:
- 
                                                    On the root FortiGate (HQ1), go to Security Fabric > Physical Topology. The root FortiGate (HQ1) is connected by the downstream FortiGate (HQ2) with VPN icon in the middle.
- 
                                                    On the root FortiGate (HQ1), go to Security Fabric > Logical Topology. The root FortiGate (HQ1) VPN interface To-HQ2 is connected by downstream FortiGate (HQ2) VPN interface To-HQ1 with VPN icon in the middle.
To run diagnostics:
- 
                                                    To view the downstream FortiGate pending authorization on root FortiGate (HQ1): HQ1 # diagnose sys csf authorization pending-list Serial IP Address HA-Members Path ------------------------------------------------------------------------------------ FG101ETK18002187 0.0.0.0 FG3H1E5818900718:FG101ETK18002187
- 
                                                    To view the downstream FortiGate (HQ2) after it joins the Security Fabric: HQ1 # diagnose sys csf downstream
 1:     FG101ETK18002187 (10.10.10.3) Management-IP: 0.0.0.0 Management-port:0 parent: FG3H1E5818900718
        path:FG3H1E5818900718:FG101ETK18002187
        data received: Y downstream intf:To-HQ1 upstream intf:To-HQ2 admin-port:443
        authorizer:FG3H1E5818900718
- 
                                                    To view the root FortiGate (HQ1) on the downstream FortiGate (HQ2) after joining the Security Fabric: HQ2 # diagnose sys csf upstream Upstream Information: Serial Number:FG3H1E5818900718 IP:10.10.10.1 Connecting interface:To-HQ1 Connection status:Authorized
