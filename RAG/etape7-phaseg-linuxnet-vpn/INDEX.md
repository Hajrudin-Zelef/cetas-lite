# INDEX — Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring

Corpus `etape7-phaseg-linuxnet-vpn` · **18 fichiers** · 834 lignes source · ~9616 mots · partition exacte de `docs/RAG/etape7_phaseG_linuxnet_vpn.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-linuxnet-vpn/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring](00-linuxnet-vpn/overview.md) | 1–62 | deep-dive | reference |
| 02 | [2. iptables migration and nftables](00-linuxnet-vpn/2-iptables-migration-and-nftables.md) | 63–122 | deep-dive | reference |
| 03 | [4.2 Hardening checklist (sshd)](00-linuxnet-vpn/4-2-hardening-checklist-sshd.md) | 123–160 | deep-dive | reference |
| 04 | [Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 4)](00-linuxnet-vpn/part-4.md) | 161–167 | deep-dive | reference |
| 05 | [5. fail2ban and CrowdSec: reactive brute-force defenses](00-linuxnet-vpn/5-fail2ban-and-crowdsec-reactive-brute-force-defenses.md) | 168–197 | deep-dive | reference |
| 06 | [6. Wazuh: architecture, 2026 releases, FIM, vulnerability detection, Cloud pricing](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md) | 198–245 | deep-dive | pricing |
| 07 | [6.7 Alternatives: Security Onion and Graylog](00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md) | 246–262 | deep-dive | reference |
| 08 | [7. WireGuard: protocol, kernel integration, wg-quick](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md) | 263–312 | deep-dive | reference |
| 09 | [8.4 Exit nodes and Mullvad integration](00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md) | 313–333 | deep-dive | reference |
| 10 | [9. ZeroTier: overlay networking, pricing, planet/moon architecture](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md) | 334–385 | deep-dive | pricing |
| 11 | [11. OpenVPN 2.6/2.7 and DCO](00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md) | 386–460 | deep-dive | reference |
| 12 | [Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 12)](00-linuxnet-vpn/part-12.md) | 461–466 | deep-dive | reference |
| 13 | [14. Adoption signals (2026)](00-linuxnet-vpn/14-adoption-signals-2026.md) | 467–519 | deep-dive | reference |
| 14 | [Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 14)](00-linuxnet-vpn/part-14.md) | 520–534 | deep-dive | reference |
| 15 | [18. Source URLs (retrieved 2026-09-22; verbatim)](00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md) | 535–618 | deep-dive | reference |
| 16 | [19. Appendix A — annotated configuration references](00-linuxnet-vpn/19-appendix-a-annotated-configuration-references.md) | 619–676 | deep-dive | reference |
| 17 | [19.3 OpenSSH hardening snippet with Match blocks and CA trust](00-linuxnet-vpn/19-3-openssh-hardening-snippet-with-match-blocks-and-ca-trus.md) | 677–750 | deep-dive | reference |
| 18 | [19.6 Tailscale ACL policy (least-privilege example)](00-linuxnet-vpn/19-6-tailscale-acl-policy-least-privilege-example.md) | 751–834 | deep-dive | regulation |

## Par tâche

- **pricing** — [6. Wazuh: architecture, 2026 releases, FIM, vulnerability detection, Cloud pricing](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md), [9. ZeroTier: overlay networking, pricing, planet/moon architecture](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md)
- **reference** — [Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring](00-linuxnet-vpn/overview.md), [2. iptables migration and nftables](00-linuxnet-vpn/2-iptables-migration-and-nftables.md), [4.2 Hardening checklist (sshd)](00-linuxnet-vpn/4-2-hardening-checklist-sshd.md), [Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 4)](00-linuxnet-vpn/part-4.md), [5. fail2ban and CrowdSec: reactive brute-force defenses](00-linuxnet-vpn/5-fail2ban-and-crowdsec-reactive-brute-force-defenses.md), [6.7 Alternatives: Security Onion and Graylog](00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md), [7. WireGuard: protocol, kernel integration, wg-quick](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md), [8.4 Exit nodes and Mullvad integration](00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md), [11. OpenVPN 2.6/2.7 and DCO](00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md), [Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 12)](00-linuxnet-vpn/part-12.md), [14. Adoption signals (2026)](00-linuxnet-vpn/14-adoption-signals-2026.md), [Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring (part 14)](00-linuxnet-vpn/part-14.md), [18. Source URLs (retrieved 2026-09-22; verbatim)](00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md), [19. Appendix A — annotated configuration references](00-linuxnet-vpn/19-appendix-a-annotated-configuration-references.md), [19.3 OpenSSH hardening snippet with Match blocks and CA trust](00-linuxnet-vpn/19-3-openssh-hardening-snippet-with-match-blocks-and-ca-trus.md)
- **regulation** — [19.6 Tailscale ACL policy (least-privilege example)](00-linuxnet-vpn/19-6-tailscale-acl-policy-least-privilege-example.md)

## Par acteur

- **Apple** (2) — [00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md), [00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md)
- **Google** (2) — [00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md), [00-linuxnet-vpn/14-adoption-signals-2026.md](00-linuxnet-vpn/14-adoption-signals-2026.md)
- **Microsoft** (1) — [00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md)
- **OpenAI** (1) — [00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md](00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md)
- **United States** (1) — [00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md)

## Par date

- **2025-04** — [00-linuxnet-vpn/4-2-hardening-checklist-sshd.md](00-linuxnet-vpn/4-2-hardening-checklist-sshd.md)
- **2025-10** — [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md)
- **2026-01** — [00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md](00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md)
- **2026-01-15** — [00-linuxnet-vpn/4-2-hardening-checklist-sshd.md](00-linuxnet-vpn/4-2-hardening-checklist-sshd.md), [00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md](00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md)
- **2026-02** — [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md)
- **2026-03** — [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md), [00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md](00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md), [00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md)
- **2026-04** — [00-linuxnet-vpn/4-2-hardening-checklist-sshd.md](00-linuxnet-vpn/4-2-hardening-checklist-sshd.md), [00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md](00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md), [00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md), [00-linuxnet-vpn/part-14.md](00-linuxnet-vpn/part-14.md)
- **2026-05** — [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md), [00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md](00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md), [00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md), [00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md](00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md)
- **2026-05-22** — [00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md](00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md), [00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md](00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md)
- **2026-06** — [00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md), [00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md)
- **2026-07** — [00-linuxnet-vpn/2-iptables-migration-and-nftables.md](00-linuxnet-vpn/2-iptables-migration-and-nftables.md), [00-linuxnet-vpn/4-2-hardening-checklist-sshd.md](00-linuxnet-vpn/4-2-hardening-checklist-sshd.md), [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md), [00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md), [00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md](00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md), [00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md), [00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md](00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md)
- **2026-08** — [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md), [00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md)
- **2026-08-29** — [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md)
- **2026-09** — [00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md](00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md), [00-linuxnet-vpn/part-14.md](00-linuxnet-vpn/part-14.md)
- **2026-09-06** — [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md)
- **2026-09-22** — [00-linuxnet-vpn/overview.md](00-linuxnet-vpn/overview.md), [00-linuxnet-vpn/2-iptables-migration-and-nftables.md](00-linuxnet-vpn/2-iptables-migration-and-nftables.md), [00-linuxnet-vpn/4-2-hardening-checklist-sshd.md](00-linuxnet-vpn/4-2-hardening-checklist-sshd.md), [00-linuxnet-vpn/part-4.md](00-linuxnet-vpn/part-4.md), [00-linuxnet-vpn/5-fail2ban-and-crowdsec-reactive-brute-force-defenses.md](00-linuxnet-vpn/5-fail2ban-and-crowdsec-reactive-brute-force-defenses.md), [00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md](00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md), [00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md](00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md), [00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md](00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md), [00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md](00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md), [00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md](00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md), [00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md](00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md), [00-linuxnet-vpn/14-adoption-signals-2026.md](00-linuxnet-vpn/14-adoption-signals-2026.md), [00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md](00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md), [00-linuxnet-vpn/19-6-tailscale-acl-policy-least-privilege-example.md](00-linuxnet-vpn/19-6-tailscale-acl-policy-least-privilege-example.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–62 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/overview.md |
| 63–122 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/2-iptables-migration-and-nftables.md |
| 123–160 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/4-2-hardening-checklist-sshd.md |
| 161–167 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/part-4.md |
| 168–197 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/5-fail2ban-and-crowdsec-reactive-brute-force-defenses.md |
| 198–245 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/6-wazuh-architecture-2026-releases-fim-vulnerability-detecti.md |
| 246–262 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/6-7-alternatives-security-onion-and-graylog.md |
| 263–312 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/7-wireguard-protocol-kernel-integration-wg-quick.md |
| 313–333 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/8-4-exit-nodes-and-mullvad-integration.md |
| 334–385 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/9-zerotier-overlay-networking-pricing-planet-moon-architectu.md |
| 386–460 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/11-openvpn-2-6-2-7-and-dco.md |
| 461–466 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/part-12.md |
| 467–519 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/14-adoption-signals-2026.md |
| 520–534 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/part-14.md |
| 535–618 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/18-source-urls-retrieved-2026-09-22-verbatim.md |
| 619–676 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/19-appendix-a-annotated-configuration-references.md |
| 677–750 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/19-3-openssh-hardening-snippet-with-match-blocks-and-ca-trus.md |
| 751–834 | etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/19-6-tailscale-acl-policy-least-privilege-example.md |

