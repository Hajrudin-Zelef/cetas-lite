---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration
title: "8.4 Exit nodes and Mullvad integration"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: ["2026-04", "2026-05", "2026-05-22", "2026-07", "2026-09-22"]
keywords: ["distribution", "pricing"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [313, 333]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 67ccfece6ecd1fea3636947c7983ffb77cd52a9626d1c2cc49bf8ec35c0b8bd5
---

# 8.4 Exit nodes and Mullvad integration

- A tailnet is flat by default: every device can reach every other device on every port until ACLs restrict it [secondary: https://github.com/peculiarengineer-mk/peculiarengineer/blob/HEAD/src/content/blog/tailscale-private-networking-workers-to-prod.md retrieved 2026-09-22].
- ACLs are written in a JSON/HuJSON policy file with `acls` (accept rules over `src`/`dst` with optional ports), `groups`, `tagOwners`, and a separate `ssh` section for Tailscale SSH rules [secondary: same source].
- Tags group machines by role (`--advertise-tags=tag:prod` at join time); once a machine is tagged, the tag — not the joining user — owns the node, and only `tagOwners` can re-tag it, so tags and owners must be defined before join or the join is rejected [secondary: same source].
- A documented least-privilege pattern restricts worker nodes to specific ports on production (e.g. only Redis 6379 and MongoDB 27017), blocks worker-to-worker traffic entirely, and reserves full access for `autogroup:admin` [secondary: same source].
- A 2026 security-hardening checklist for Tailscale recommends: no `*:*` rules, tags on all server devices, `tests` covering critical paths, Tailscale SSH with OS port 22 firewalled, key expiry disabled for infrastructure but enabled for user devices, ephemeral auth keys for CI/CD, auth keys in a secrets manager, device approval for enrollments, subnet-router OS firewall restrictions, and Tailnet Lock for high-security tailnets [secondary: https://github.com/tankpkg/packages/blob/HEAD/skills/tailscale-expert/references/security-features.md retrieved 2026-09-22].

### 8.4 Exit nodes and Mullvad integration

- Exit nodes route a device's internet traffic through another tailnet node, functioning as a self-hosted VPN egress; operational checklists recommend egress filtering on exit nodes [secondary: https://github.com/tankpkg/packages/blob/HEAD/skills/tailscale-expert/references/security-features.md retrieved 2026-09-22].
- Tailscale offers Mullvad exit nodes, integrating Mullvad's VPN egress into the tailnet; a source dated 7 April 2026 described the Mullvad integration as beta [secondary: https://github.com/jordithijsman/mullvad-tailscale-macos and https://github.com/ralphalberti/dotfiles/blob/HEAD/docs/references/tailscale-and-mullvad.md, both retrieved 2026-09-22].
- The current (2026-09-22) status of the Mullvad exit-node integration — whether it remains beta or is generally available, and its pricing — was not verified against an official source in this pass [unverified].

### 8.5 Headscale: self-hosted control plane

- Headscale is the open-source, self-hosted implementation of the Tailscale control protocol, letting operators run their own coordination server instead of Tailscale's hosted SaaS [secondary: https://www.infralovers.com/blog/2026-05-22-headscale-self-hosted-tailscale-alternative/ retrieved 2026-09-22].
- Headscale 0.28.0 was the listed target in May 2026 matrices [secondary: https://github.com/pinggy-io/pinggy_website/blob/HEAD/content/blog/top_open_source_tailscale_alternatives.md retrieved 2026-09-22]; Headscale 0.29.2 was listed in a July 2026 matrix [secondary: https://github.com/yaleizhou/learn-skills.dev/blob/HEAD/data/skills-md/iuliandita/skills/networking/SKILL.md retrieved 2026-09-22].
- The self-hosting trade-off: operators keep metadata and key distribution under their own control but assume responsibility for the control plane's availability, upgrades, backups, and security — the same responsibility shift as any self-hosted coordination service [secondary: https://www.infralovers.com/blog/2026-05-22-headscale-self-hosted-tailscale-alternative/ retrieved 2026-09-22].
- Headscale appears in Kubernetes-oriented networking notes as the self-hosted control-server option alongside Tailscale [secondary: https://github.com/geomachine/kubernetes-notes/blob/HEAD/network/README.md retrieved 2026-09-22].

---

