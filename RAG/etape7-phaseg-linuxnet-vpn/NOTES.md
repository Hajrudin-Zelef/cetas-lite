# NOTES — corpus `etape7-phaseg-linuxnet-vpn`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-linuxnet-vpn` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring` — 834 lignes, 12 chunks.

- 1. Linux NAT: MASQUERADE, SNAT, DNAT, port forwarding, and conntrack
- 2. iptables migration and nftables
- 3. UFW and firewalld
- 4. OpenSSH: 2026 versions, hardening, certificates, bastions
- 5. fail2ban and CrowdSec: reactive brute-force defenses
- 6. Wazuh: architecture, 2026 releases, FIM, vulnerability detection, Cloud pricing
- 7. WireGuard: protocol, kernel integration, wg-quick
- 8. Tailscale: mesh VPN, pricing, ACLs, exit nodes, Mullvad, Headscale
- 9. ZeroTier: overlay networking, pricing, planet/moon architecture
- 10. strongSwan and IKEv2 IPsec: site-to-site and road-warrior
- 11. OpenVPN 2.6/2.7 and DCO
- 12. Nebula: lighthouse-based mesh
- 13. Comparison matrices
- 14. Adoption signals (2026)
- 15. Homelab and SMB use cases
- 16. Operational failure modes and pitfalls
- 17. Gaps, conflicts, and unverifiable claims
- 18. Source URLs (retrieved 2026-09-22; verbatim)
- 19. Appendix A — annotated configuration references
- 20. Appendix B — selection flowcharts (text form)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
