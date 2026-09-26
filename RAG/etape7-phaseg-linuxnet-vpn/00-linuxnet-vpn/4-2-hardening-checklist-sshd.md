---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/4-2-hardening-checklist-sshd
title: "4.2 Hardening checklist (sshd)"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: ["2025-04", "2026-01-15", "2026-04", "2026-07", "2026-09-22"]
keywords: ["agent", "agents", "distribution", "research"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [123, 160]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 6ace233d345d7cbc1cda6a98fbee72c14a6bfea74d06853c82c9ac3d0c5bb424
---

# 4.2 Hardening checklist (sshd)

- OpenSSH 10.3 (and 10.3p1) was released on 2 April 2026 [secondary: https://cybersecuritynews.com/openssh-10-3-release/ retrieved 2026-09-22].
- The headline 10.3 fix was a shell-injection vulnerability in the `-J` (ProxyJump) command-line option: user and host names passed via `-J` or `-oProxyJump="..."` on the command line were not validated, allowing shell injection when those values came from adversarial input; the fix validates command-line values, while configuration-file entries remain unvalidated [secondary: https://cybersecuritynews.com/openssh-10-3-release/ retrieved 2026-09-22].
- OpenSSH 10.3 also corrected a subtle `sshd` certificate behavior: certificates issued with an empty principals section had been treated as a wildcard, effectively authenticating as any user trusting the issuing CA via `authorized_keys`; the new behavior rejects this dangerous edge case [secondary: https://cybersecuritynews.com/openssh-10-3-release/ retrieved 2026-09-22].
- OpenSSH 10.4 was released on 6 July 2026 with security patches, protocol hardening, and early post-quantum cryptography support, distributed through the mirrors listed on the official OpenSSH site [secondary: https://cybersecuritynews.com/openssh-10-4-security-fixes/ retrieved 2026-09-22].
- OpenSSH 10.5 fixed further security flaws; secondary coverage framed the release as the project's response to AI-assisted bug discovery [secondary: https://linuxiac.com/openssh-10-5-fixes-security-flaws-as-project-responds-to-ai-assisted-bug-discovery/ retrieved 2026-09-22].
- Version-history floors relevant to 2026 deployments: `ed25519-sk` hardware keys and default touch-to-sign arrived in 8.2; `ssh-rsa`/SHA-1 host keys were disabled by default in 8.8; DSA support was removed in 10.0 (April 2025); the default key exchange became the post-quantum hybrid `mlkem768x25519-sha256` in 10.0; `PerSourcePenalties` (in-daemon rate limiting) arrived in 9.8 [secondary: https://github.com/fuyutarow/dotfiles/blob/HEAD/agents/skills/securing-remote-access/SKILL.md retrieved 2026-09-22].
- As of mid-2026 the OpenSSH release series stood at 10.3 per one practitioner reference [secondary: same source], but 10.4 and 10.5 coverage above shows the series advanced through the summer; treat "latest" claims as date-sensitive and check the official mirrors before upgrading [secondary].

### 4.2 Hardening checklist (sshd)

- Disable password authentication and root password login once key-based access is verified: `PasswordAuthentication no`, `PermitRootLogin no` (or `prohibit-password`), `PermitEmptyPasswords no` [secondary: https://github.com/mr-mainbytelabs/technical-docs-portfolio/blob/HEAD/ssh-hardening-sop.md retrieved 2026-09-22].
- Restrict who can log in with `AllowUsers`/`AllowGroups`, set `MaxAuthTries 3` and `LoginGraceTime 60`, and always validate with `sshd -t` before restarting [secondary: https://oneuptime.com/blog/post/2026-01-15-enable-ssh-ubuntu-desktop/view retrieved 2026-09-22].
- Pin modern algorithms explicitly on sensitive hosts: KEX `sntrup761x25519-sha512@openssh.com,curve25519-sha256,...`, ciphers `chacha20-poly1305@openssh.com,aes256-gcm@openssh.com,...`, MACs `hmac-sha2-512-etm@openssh.com,...`; verify what the local build supports with `ssh -Q kex`, `ssh -Q cipher`, `ssh -Q mac` [secondary: https://github.com/matteobisi/msbiro.net/blob/HEAD/content/posts/back-to-basics-sshd-hardening.md retrieved 2026-09-22].
- Avoid weak MACs still present in some defaults: 64-bit UMAC variants and SHA-1-based MACs are considered deprecated by hardening guides [secondary: https://github.com/am0rphous/cheatsheets/blob/HEAD/Linux%20%F0%9F%90%A7/Network/SSH-Hardening.md retrieved 2026-09-22].
- Prefer Ed25519 keys (`ssh-keygen -t ed25519`) for new deployments; use 4096-bit RSA only where legacy compatibility demands it [secondary: https://oneuptime.com/blog/post/2026-01-15-enable-ssh-ubuntu-desktop/view retrieved 2026-09-22].
- Changing the SSH port is obscurity, not security: it reduces log noise but is not a substitute for key-only authentication [secondary: same source].

### 4.3 Certificates instead of authorized_keys sprawl

- OpenSSH certificates replace per-host `authorized_keys` distribution: a CA signs short-lived user certificates, and hosts trust the CA via `TrustedUserCAKeys`; the daemon then accepts any certificate the CA signed [secondary: https://github.com/fuyutarow/dotfiles/blob/HEAD/agents/skills/securing-remote-access/SKILL.md retrieved 2026-09-22].
- Setting `TrustedUserCAKeys` does not automatically disable `authorized_keys`; both are accepted unless `AuthorizedKeysFile none` is set explicitly [secondary: same source].
- The operational philosophy is "don't revoke, expire": certificate revocation lists (KRLs) reintroduce the per-host distribution problem certificates were meant to solve, so short TTLs are the primary control and active revocation is kept for emergencies [secondary: same source].
- A 2026-09 practitioner guide documents combining Tailscale ACLs with an SSH CA for certificate-based auth with revocation handling on VPS fleets [secondary: https://github.com/pengzz9527/selfvps/blob/HEAD/content/en/post/ssh-zero-trust-tailscale-certificate-auth.md retrieved 2026-09-22].

### 4.4 Bastions and ProxyJump

- The modern bastion pattern is `ProxyJump` (`ssh -J bastion internal`) rather than SSH agent forwarding: forwarding an agent lets a compromised intermediate host use the agent's identities for onward connections, while ProxyJump keeps keys on the client [secondary: https://github.com/fuyutarow/dotfiles/blob/HEAD/agents/skills/securing-remote-access/SKILL.md retrieved 2026-09-22].
- A 2026 platform decision record documents replacing a traditional SSH bastion with Tailscale (mesh VPN plus identity ACLs), noting the trade-off: the bastion host disappears but a third-party control plane now sits on the access path, and the subnet router itself still needs sizing and patching [secondary: https://github.com/smana/cloud-native-ref/blob/HEAD/website/content/docs/decisions/0013-tailscale-over-bastion.md retrieved 2026-09-22].
- For fleets where a mesh VPN is not an option, certificate-backed ProxyJump through a hardened, audited jump host remains the standard pattern; enforce per-source rate limiting (`PerSourcePenalties` on 9.8+) and MFA at the bastion [secondary].

### 4.5 sshuttle

- sshuttle is a transparent VPN-over-SSH tool: it forwards traffic through an SSH connection without requiring administrator access or a TUN device on the client side in its classic mode [unverified: no dedicated sshuttle source was retrieved in this pass; included for scope completeness].
- Because sshuttle tunnels over SSH, it inherits SSH authentication and encryption, but it is not a substitute for a real VPN where kernel-level tunneling, UDP transport, or always-on connectivity is required [unverified].
- **Research gap:** sshuttle versioning, 2026 maintenance status, and performance characteristics were not covered by retrieved sources; verify against the upstream project before citing specifics [unverified].

### 4.6 Tailscale SSH

