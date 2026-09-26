---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/part-12
title: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 12)"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["benchmark", "benchmarks", "consumer"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [461, 466]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 04ed588f744175c263d29768d6606d50f85947b534cffaab15c8974087c684ad
---

# Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 12)

- Do not compare headline numbers across the WireGuard-vs-OpenVPN benchmarks cited in section 11.3: the datazone.de 2026 benchmark [independent], the voxihost write-up [secondary], and consumer-VPN comparisons [secondary] use different hardware, ciphers, stream counts, and DCO states.
- The only safe directional summary: kernel-path VPNs (WireGuard, OpenVPN with DCO, IPsec/XFRM) substantially outperform userspace tun-based OpenVPN on the same hardware; among kernel-path options, results depend on cipher, MTU, and CPU crypto acceleration [secondary/independent].
- For mesh overlays, relay-vs-direct path selection dominates real-world performance more than protocol crypto speed: a DERP-relayed Tailscale path can underperform a direct WireGuard link regardless of cipher efficiency [secondary].

---

