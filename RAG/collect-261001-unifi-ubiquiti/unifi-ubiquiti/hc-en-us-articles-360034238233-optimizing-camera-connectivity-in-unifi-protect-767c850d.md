---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-360034238233-optimizing-camera-connectivity-in-unifi-protect-767c850d
title: "hc-en-us-articles-360034238233-optimizing-camera-connectivity-in-unifi-protect-767c850d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-360034238233-optimizing-camera-connectivity-in-unifi-protect-767c850d.md
source_anchor: ""
source_lines: [1, 51]
sha256: 3c46b633b0e82fb497ea45b626e54c0e5828f905eea0ec0f6a5adaac56bea3c7
---

# hc-en-us-articles-360034238233-optimizing-camera-connectivity-in-unifi-protect-767c850d

Optimizing Camera Connectivity in UniFi Protect
This guide covers how to identify and resolve common connectivity issues that may affect your UniFi Protect camera performance and application access.
Classifying Your Connectivity Issue
Identifying the nature of your issue will help you address it more effectively:
- Camera Connectivity Issues: Result from a weak WiFi signal to the camera, affecting recording quality and creating gaps in footage.
- Client Connectivity Issues: Caused by a poor connection when accessing Protect remotely. Recordings will play in lower quality for smoother playback but are stored locally in full quality.
Steps to Diagnose Issues
- 
Access via Different Methods: Test access from various devices, operating systems, browsers, and networks:
  - Use different mobile devices (iOS, Android) and browsers (Chrome, Firefox, Safari).
  - Try both local networks and remote networks like workplace or public WiFi.
  - Have multiple users with varied roles attempt access.
- 
Check Network Stability:
  - Perform a speed test with the WiFiman app on the same network as your UniFi Console; networks below 2.5 Mbps may struggle.
  - Ensure no VPN or bandwidth-limiting application is active.
  - Confirm that your UniFi Console is not exceeding its camera capacity.
- Address High CPU Usage: If your computer's CPU utilization is near 100%, try reducing the number of simultaneous camera streams in the live view matrix.
Troubleshooting Slow Loading and Buffering
- Network Stability: Ensure your network connection is stable and strong.
- Bandwidth Limits: Confirm that bandwidth isn’t restricted by a VPN or a different subnet configuration.
- Camera Limits: Confirm your UniFi Console has not exceeded its supported camera limit.
- Device Performance: Check your computer’s CPU usage. High usage may impact playback.
- Direct Connection: Connect directly to your UniFi Console by entering its IP address into your browser’s navigation bar, while connected to the local network.
Troubleshooting Remote Access Issues
If remote access problems persist:
- Check for Ubiquiti Service Status: Visit status.ui.com to see if any UniFi services are down.
- User Permissions: Make sure you have the correct permissions for remote access. See UniFi Protect - Add and manage users.
- Firmware Updates: Confirm your CloudKey or UniFi Console firmware is up to date to support remote functionality.
Addressing Issues with Mobile App Access
- App Updates and Permissions: Ensure the UniFi Protect app is updated and has permission to access data over WiFi and cellular.
- Disable VPNs and Privacy Apps: VPNs and privacy applications may interfere with WebRTC, which Protect relies on. Try disabling VPNs and private DNS options, especially on Android devices.
- Troubleshoot Persistent App Issues: If the app continues to malfunction, reinstall it and try again.
Browser-Specific Issues
If Protect works in some browsers but not others, third-party browser extensions may be blocking connectivity:
- Disable Extensions: Try disabling suspected privacy or security-related browser extensions, such as VPNs, ad blockers, and WebRTC leak prevention tools.
- Test and Re-enable Extensions Individually: Re-enable each extension one by one to identify any conflicting software.
- 
Chrome Settings: Disable Chrome’s feature flag Anonymize local IPs exposed by WebRTC by entering chrome://flags/#enable-webrtc-hide-local-ips-with-mdns into your address bar and setting it to "Disabled."
Advanced Troubleshooting: WebRTC and NAT Issues
UniFi Protect uses WebRTC for connectivity between devices. Certain network configurations, like Symmetric NAT, can block these connections:
- Verify Network Requirements: Ensure both networks have reliable Internet access, adequate bandwidth, and proper firewall settings (outbound TCP on Port 443 and outbound UDP on Ports 0–65535).
- Symmetric NAT Detection: If a Symmetric NAT setup blocks connectivity, either set up a VPN connection or switch your router to a Cone NAT configuration.
- SSH Logging for Symmetric NAT: If you suspect Symmetric NAT on the console side, SSH into the UniFi Console and check logs for “symmetric” mentions to confirm.
Addressing Network-Specific Connectivity Issues
If Protect only fails on certain networks, focus on troubleshooting that network.
- 
Network-Specific Troubleshooting:
  - Ensure the UniFi Console has a stable connection and valid gateway IP and DNS settings.
  - Verify DNS resolution for key UniFi Protect service domains (e.g., Device.svc.ubnt.com, Global.stun.twilio.com).
  - Temporarily disable any custom firewall rules and verify WebRTC traffic is not blocked.
