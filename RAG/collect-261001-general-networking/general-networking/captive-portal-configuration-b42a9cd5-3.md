---
id: collect-261001-general-networking/general-networking/captive-portal-configuration-b42a9cd5-3
title: "captive-portal-configuration-b42a9cd5"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Microsoft"]
dates: []
keywords: ["benchmark", "incident"]
source: docs/RAG/collect-261001-general-networking/captive-portal-configuration-b42a9cd5.md
source_anchor: ""
source_lines: [81, 99]
sha256: 5e5e86bf4e63cfe968cea57aebf8296c3bcd09fc63ac800caeff8718536ad332
---

# captive-portal-configuration-b42a9cd5

A splash page that appears once isn't proof of a finished deployment. Meraki captive portal rollouts commonly fail when the client reaches the walled garden but can't reach an authentication dependency, when RADIUS or SAML responses time out, when a voucher is redeemed twice, or when an Apple or Android device caches an old IPSK and keeps presenting stale credentials.
Start troubleshooting from the client path. Confirm the guest device receives the expected network configuration, resolves the portal service, reaches the allowed authentication endpoints, and receives the intended policy after login. Then test a second device type. Captive portal detection can produce both false positives and false negatives when a portal interferes with operating-system connectivity checks, a behavior highlighted in Microsoft's captive portal guidance.
Vertical checks that match the environment
- Education: Isolate the guest VLAN, apply content controls where required, retain the logs your policy requires, and combine voucher rotation with reauthentication. Don't expose campus management interfaces just because the user is on a trusted-looking device.
- Retail: Use sensible per-client bandwidth caps, keep the initial social-login or email path short, and define MV Sense dwell-time thresholds that operations can interpret. A dashboard full of unowned metrics won't improve the store experience.
- Corporate BYOD: Prefer IPSK with SAML or Azure AD and posture checks over a shared PSK. Conditional access should determine whether the identity is eligible, while Meraki policies determine what the device can reach.
Pre-launch validation table
| Issue or Vertical | Likely Cause | Recommended Fix or Best Practice | 
|---|---|---|
| Client stuck in walled garden | Missing allowed endpoint or incorrect redirect | Review portal, DNS, payment, and OS connectivity-check exceptions | 
| RADIUS or SAML timeout | Identity service isn't reachable from the guest path | Test reachability and response behavior before production | 
| Voucher redeemed twice | Weak session or MAC-binding policy | Define reuse rules, expiration, and staff recovery procedure | 
| Apple or Android keeps old IPSK | Credential caching on the device | Revoke the key, clear the saved network, and retest reauthentication | 
| Education guest network | Internal resources exposed | Enforce guest VLAN isolation and default-deny LAN rules | 
| Retail guest Wi-Fi | Excessive bandwidth consumption | Apply a per-client cap and monitor peak sessions | 
| Corporate BYOD | Shared PSK creates uncontrolled access | Use individual IPSK credentials with SAML or Azure AD | 
A practical hardening pattern is to isolate the guest SSID from trusted VLANs, allow only portal, DNS, payment, and operating-system connectivity-check destinations before login, and serve the portal over a valid HTTPS certificate. A captive portal setup guide also gives a lab benchmark using 100 concurrent connections, a 30-minute idle timeout, a 120-minute hard timeout, and per-user caps of 8000 Kbit/s down and 2500 Kbit/s up. Treat those as starting points for testing, not universal production values. Tight timeouts can protect the WAN, but they can also force hotel and retail users to authenticate again too often.
Before approval, verify guest-VLAN DNS resolution, redirects on current Android and iOS devices plus a MacBook, RADIUS reachability where used, certificate validity, voucher behavior, IPSK revocation, roaming, and analytics event firing. Review firmware and rogue AP checks regularly, and log portal events in a way that supports both incident response and privacy review.
Splash Access provides external captive portals for Cisco Meraki guest Wi-Fi, including customizable splash pages, IPSK and voucher workflows, SAML or Azure AD integration, social Wi-Fi options, geo-fenced coupons, billing connections, and analytics that can work alongside Meraki MV Sense. Visit Splash Access to map those capabilities to your hotel, education, retail, or corporate BYOD rollout and plan the access policies before deployment.
