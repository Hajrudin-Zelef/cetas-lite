---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/part-4
title: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 4)"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: []
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [161, 167]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: ecc76a60a859093ce1329ee23017f238b37bb3b4aefba952db836dfab9aa55ff
---

# Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 4)

- Tailscale SSH replaces the OpenSSH server with a Tailscale-aware daemon that verifies user identity and device posture against tailnet ACLs, removing the need to distribute SSH keys at all [secondary: https://tech-insider.org/tailscale-vs-wireguard-2026/ retrieved 2026-09-22].
- Access is governed by `ssh` rules in the tailnet policy file (`action: accept`, `src`, `dst`, `users`), and check mode (`"action": "check"`) can require explicit user approval per connection for privileged access [secondary: https://github.com/smana/cloud-native-ref/blob/HEAD/website/content/docs/decisions/0013-tailscale-over-bastion.md and https://github.com/tankpkg/packages/blob/HEAD/skills/tailscale-expert/references/security-features.md, both retrieved 2026-09-22].
- Session recording can be enforced with `enforceRecorder: true` for privileged SSH sessions [secondary: https://github.com/tankpkg/packages/blob/HEAD/skills/tailscale-expert/references/security-features.md retrieved 2026-09-22].
- A documented hardening posture pairs Tailscale SSH on servers with OS-level port 22 firewalled off, so the only SSH path is the identity-checked one [secondary: same source].

---

