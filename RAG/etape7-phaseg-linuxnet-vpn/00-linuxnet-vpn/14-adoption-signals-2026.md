---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/14-adoption-signals-2026
title: "14. Adoption signals (2026)"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: ["Google"]
dates: ["2026-04", "2026-09", "2026-09-22"]
keywords: ["advisory", "agent", "agents", "benchmarks", "consumer", "embedding", "mcp", "pricing"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [467, 534]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 30075c611747da5bb0dd8fb94dd65561e6e4d89b9dc481beb534de1d50f58da5
---

# 14. Adoption signals (2026)

## 14. Adoption signals (2026)

- Tailscale is claimed at 5 million users in 2026 secondary coverage [secondary: https://tech-insider.org/tailscale-vs-wireguard-2026/ retrieved 2026-09-22] — vendor-adjacent figure, treat as directional [unverified as official].
- Tailscale's 2026 integration surface (IdPs: Google Workspace, Entra ID, Okta, Auth0, GitHub, GitLab, OIDC, SAML, JumpCloud; posture: CrowdStrike, SentinelOne, Jamf, Kandji, Intune, Kolide; CI: GitHub Actions, GitLab CI, CircleCI, Buildkite, Jenkins) indicates enterprise-oriented adoption breadth [secondary: same source].
- WireGuard's adoption is measured by embedding: major consumer VPNs (Mullvad, Proton VPN, IVPN, NordLynx), Cloudflare WARP, firewall distributions (OPNsense, pfSense, OpenWrt, MikroTik RouterOS), and the Linux kernel itself [secondary: same source].
- Security Onion reports over 2 million downloads historically [secondary: https://gbhackers.com/security-onion-24-10-released/ retrieved 2026-09-22].
- Wazuh's ecosystem activity in 2026 includes third-party MCP/API integrations tested against 4.14.7 and active homelab/SMB deployment guides, indicating healthy community adoption [secondary: https://github.com/gensecaihq/wazuh-mcp-server/commit/239b94e062c0870ef9caa2db3f2455e2c25092d0 and https://github.com/duresa7/homelab/blob/HEAD/Guides/Wazuh.md, retrieved 2026-09-22].
- ZeroTier is rated 4.8/5 on G2 [secondary: https://www.g2.com/products/zerotier-one/pricing retrieved 2026-09-22].

---

## 15. Homelab and SMB use cases

### 15.1 Remote access without a public bastion

- A 2026 cloud-native platform decision record replaced its SSH bastion with Tailscale: a subnet router in a private subnet (no public IP) advertises VPC CIDRs, identity ACLs replace security-group thinking, and Tailscale SSH provides keyless per-connection SSH; the acknowledged costs are a third-party control plane on the access path and one remaining EC2 instance to patch [secondary: https://github.com/smana/cloud-native-ref/blob/HEAD/website/content/docs/decisions/0013-tailscale-over-bastion.md retrieved 2026-09-22].
- For a single VPS, a documented 2026 pattern combines Tailscale for network-layer zero trust, ACLs restricting SSH to named users, and an SSH CA for certificate auth — with the OS SSH optionally bound to the tailnet IP on a non-standard port [secondary: https://github.com/pengzz9527/selfvps/blob/HEAD/content/en/post/ssh-zero-trust-tailscale-certificate-auth.md retrieved 2026-09-22].
- Where Tailscale is not an option, certificate-backed ProxyJump through one hardened jump host remains the recommended pattern [secondary].

### 15.2 Homelab SIEM

- Documented 2026 homelab Wazuh deployments run the full stack (manager, indexer, dashboard) on a single Ubuntu host, enroll a mix of Linux/Windows/macOS agents over TCP 1514/1515, and use agent groups to push FIM policy (e.g. a `proxmox` group covering all hypervisor nodes) [secondary: https://github.com/duresa7/homelab/blob/HEAD/Guides/Wazuh.md retrieved 2026-09-22].
- A Proxmox-homelab Wazuh plan verified every upstream fact against the actual `v4.14.7` tag (after catching a wrong `6.0.0` tag assumption), and scoped Security Onion out of the same iteration — a good illustration of verify-then-deploy discipline [secondary: https://github.com/stevedwray/proxmox-homelab/blob/HEAD/docs/wazuh-stack/plan.md retrieved 2026-09-22].
- Common homelab mistakes observed in guides: agents pointed at a stale manager IP after migration, dashboard/API ports not opened on the host firewall, and under-provisioned indexer storage [secondary: https://github.com/duresa7/homelab/blob/HEAD/Guides/Wazuh.md retrieved 2026-09-22].

### 15.3 SMB selection heuristics

- Fewer than ~10 nodes, no compliance mandate: Tailscale (or Headscale if SaaS is unacceptable) plus Tailscale SSH and ACLs; host firewall (UFW/firewalld) default-deny; key-only OpenSSH on anything reachable [secondary synthesis of cited guides].
- 10–100 nodes with compliance needs: Wazuh for host monitoring/FIM/vulnerability detection, possibly paired with Security Onion for network visibility; CrowdSec for collaborative banning on public-facing services [secondary synthesis].
- Multi-site connectivity: WireGuard for simple site-to-site links; strongSwan IKEv2 where EAP/RADIUS integration or vendor interop is required; ZeroTier where Layer-2 adjacency is needed [secondary synthesis].
- These heuristics are editorial synthesis from the cited sources, not vendor recommendations [secondary].

---

## 16. Operational failure modes and pitfalls

- **conntrack exhaustion** on NAT gateways: intermittent new-connection failures under load; monitor fill percentage and tune timeouts, not just the max [secondary: section 1.3 sources].
- **UFW lockout:** enabling UFW over SSH before allowing the SSH port [secondary: section 3 sources].
- **firewalld permanent/runtime split:** rules that vanish on reboot or never apply [secondary: section 3 sources].
- **Double banning:** fail2ban jails plus CrowdSec scenarios on the same logs producing conflicting ban states [secondary: section 5 sources].
- **Flat tailnet:** deploying Tailscale without ACLs leaves every device reachable from every device [secondary: section 8 sources].
- **Tag ownership lockout:** joining with `--advertise-tags` before defining `tagOwners` gets the join rejected [secondary: section 8 sources].
- **WireGuard behind NAT without keepalive:** idle peers become unreachable; set `PersistentKeepalive` [secondary].
- **Wrong DCO module:** `ovpn-dco` with OpenVPN 2.7+ or the in-kernel `ovpn` module with 2.6.x [official: section 11 sources].
- **strongSwan behind NAT without NAT-T:** IKE succeeds but ESP fails; ensure UDP 4500 is open [secondary].
- **Wazuh agent pointing at a stale manager IP** after infrastructure migration; **indexer disk exhaustion** from unrotated indices [secondary: section 15 sources].
- **SSH agent forwarding through bastions** exposing identities to a compromised intermediate host; prefer ProxyJump [secondary: section 4 sources].
- **Empty-principal SSH certificates** issued by a misconfigured CA granting wildcard access — fixed server-side in OpenSSH 10.3, but CA hygiene still matters [secondary: section 4 sources].

---

## 17. Gaps, conflicts, and unverifiable claims

- **Wazuh Cloud pricing:** $1,467 vs $1,449 top-tier monthly figures conflict between two secondary sources (both September 2026 snapshots); neither confirmed against the official pricing page [secondary].
- **Tailscale pricing:** "Starter $6/user/month" vs "Standard $8/user/month" entry-tier conflict between two secondary sources; Premium agrees at $18/user/month [secondary].
- **ZeroTier pricing:** 10 vs 25 free devices, and $18/10-devices/month vs $2/device/month, across G2, Slashdot, and a third-party comparison [secondary].
- **Mullvad exit nodes:** described as beta in an April 2026 source; current status unverified [unverified].
- **Tailscale "5 million users":** vendor-adjacent claim in secondary coverage, not confirmed against an official source [unverified as official].
- **ZeroTier planet/moon operational detail:** no dedicated source retrieved; included from general knowledge and flagged [unverified].
- **sshuttle:** no dedicated source retrieved in this pass; 2026 maintenance status and performance unverified [unverified].
- **Graylog:** listed as an alternative but not researched in depth [unverified].
- **strongSwan 6.0.7 CVE-2026-47895:** reported via FreeBSD port notes; upstream advisory text not retrieved [secondary].
- **VPN benchmarks:** methodologies differ across sources and must not be blended; no single "N× faster" claim is supported [independent/secondary].
- **conntrack sizing numbers** (262,144 / 1,048,576) are sourced examples, not recommendations [secondary].
- **OpenSSH "latest version":** 10.5 is the newest version found in secondary coverage (linuxiac), but official release notes were not retrieved; confirm on official mirrors before citing as latest [secondary].

---

