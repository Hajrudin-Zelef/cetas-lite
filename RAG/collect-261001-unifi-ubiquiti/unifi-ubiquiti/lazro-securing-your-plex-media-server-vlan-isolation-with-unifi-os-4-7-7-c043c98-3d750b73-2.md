---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/lazro-securing-your-plex-media-server-vlan-isolation-with-unifi-os-4-7-7-c043c98-3d750b73-2
title: "lazro-securing-your-plex-media-server-vlan-isolation-with-unifi-os-4-7-7-c043c98-3d750b73"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["containment"]
source: docs/RAG/collect-261001-unifi-ubiquiti/lazro-securing-your-plex-media-server-vlan-isolation-with-unifi-os-4-7-7-c043c98-3d750b73.md
source_anchor: ""
source_lines: [192, 218]
sha256: a6f39f10d9322b59c942d365e024c9dc687c9bbefa7dab85ad8e2196b604accb
---

# lazro-securing-your-plex-media-server-vlan-isolation-with-unifi-os-4-7-7-c043c98-3d750b73

- Wrong destination port: Verify the service is actually listening on that port
- Policy order: Deny rules above Allow rules will block traffic
- Zone assignment: Device not actually on the expected VLAN/zone
External Access Shows Relay Connection
Problem: Plex accessible but mobile app uses relay
Solutions:
- Reset app cache (Profile → Settings → Advanced → Reset Cache)
- Force quit and restart the app
- Sign out and back into Plex account
- Test on cellular data (not home WiFi)
- Verify port forwarding shows success on canyouseeme.org
- Check Plex Remote Access settings show correct public port
Conclusion
Network segmentation through VLAN isolation represents a fundamental shift in how we approach home network security. Rather than treating your entire home network as a single trusted zone, this guide has demonstrated how to implement defense-in-depth principles that are commonly used in enterprise environments.
By isolating your Plex media server and related services on dedicated VLANs, you’ve accomplished several critical security objectives:
Containment: If your media server were ever compromised — whether through a software vulnerability, weak credentials, or malicious media files — the attacker’s lateral movement across your network is severely restricted. They cannot freely pivot to other devices on different VLANs without first breaching your carefully constructed firewall policies.
Principle of Least Privilege: Each firewall policy you’ve created embodies this security principle. Your Apple TV can access Plex on port 32400, but nothing more. Your client networks can reach Deluge on port 8112, but are blocked from everything else on that VLAN. This granular control means that even legitimate devices only have the minimum access necessary to function.
Visibility and Control: With traffic now flowing through explicit policy definitions, you have unprecedented visibility into what’s happening on your network. UniFi’s traffic analytics can show you exactly which devices are communicating with your media server, how much bandwidth they’re consuming, and whether any unexpected connection attempts are occurring.
External Access Without Compromise: The port forwarding configuration demonstrates how to safely expose services to the internet. By using port translation (external 32401 → internal 32400) and region-based restrictions, you’ve added multiple layers of protection that go beyond simple port forwarding. The non-standard external port reduces automated scanning exposure, while the internal isolation ensures that even if someone gains access to Plex, they haven’t gained access to your entire network.
The techniques covered in this guide — VLAN creation, zone-based firewall policies, port forwarding with translation, and local DNS configuration — are all transferable skills that extend beyond just securing a media server. These same principles can be applied to IoT devices, home automation systems, guest networks, security cameras, and any other services that benefit from isolation.
As you continue to expand your home network, remember that security is not a destination but an ongoing process. Regularly review your firewall policies, keep your services updated, monitor your network traffic for anomalies, and adjust your security posture as new threats emerge and your network requirements evolve.
The investment you’ve made in properly segmenting your network today will pay dividends in both security and peace of mind for years to come.
References and Resources
- UniFi Documentation: https://help.ui.com/
- Plex Support: https://support.plex.tv/
- Network Segmentation Best Practices: NIST Special Publication 800–125B
- CanYouSeeMe Port Checker: https://www.canyouseeme.org/
