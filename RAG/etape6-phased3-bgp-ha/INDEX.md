# INDEX — Phase D3 — BGP underlay and high availability in the data center

Corpus `etape6-phased3-bgp-ha` · **16 fichiers** · 755 lignes source · ~10440 mots · partition exacte de `docs/RAG/etape6_phaseD3_bgp_ha.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-bgp-ha/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Phase D3 — BGP underlay and high availability in the data center](00-bgp-ha/overview.md) | 1–18 | deep-dive | reference |
| 02 | [Wave 1 — BGP as the data-center underlay: design patterns, addressing, and tuning](00-bgp-ha/wave-1-bgp-as-the-data-center-underlay-design-patterns-addre.md) | 19–74 | deep-dive | reference |
| 03 | [1.7 BFD integration with BGP: session bring-up order and pitfalls](00-bgp-ha/1-7-bfd-integration-with-bgp-session-bring-up-order-and-pitf.md) | 75–96 | deep-dive | reference |
| 04 | [Wave 2 — Route policy in the DC fabric: prefix lists, route maps, communities, RPKI](00-bgp-ha/wave-2-route-policy-in-the-dc-fabric-prefix-lists-route-maps.md) | 97–144 | deep-dive | regulation |
| 05 | [3.1 Cisco vPC (virtual Port Channel) — NX-OS](00-bgp-ha/3-1-cisco-vpc-virtual-port-channel-nx-os.md) | 145–190 | deep-dive | reference |
| 06 | [3.5 Aruba CX VSX (Virtual Switching Extension) — AOS-CX](00-bgp-ha/3-5-aruba-cx-vsx-virtual-switching-extension-aos-cx.md) | 191–236 | deep-dive | reference |
| 07 | [Wave 4 — First-hop redundancy, anycast gateway, ISSU/NSF, and convergence](00-bgp-ha/wave-4-first-hop-redundancy-anycast-gateway-issu-nsf-and-con.md) | 237–294 | deep-dive | reference |
| 08 | [Wave 5 — File verification and coverage audit](00-bgp-ha/wave-5-file-verification-and-coverage-audit.md) | 295–341 | deep-dive | reference |
| 09 | [Wave 6 — BGP session security, dynamic peering, and EVPN overlay control-plane design](00-bgp-ha/wave-6-bgp-session-security-dynamic-peering-and-evpn-overlay.md) | 342–402 | deep-dive | reference |
| 10 | [7.4 Layer 3 over MC-LAG pairs](00-bgp-ha/7-4-layer-3-over-mc-lag-pairs.md) | 403–426 | deep-dive | reference |
| 11 | [Wave 8 — Detection extras, FHRP tracking, control-plane protection, and GR procedures](00-bgp-ha/wave-8-detection-extras-fhrp-tracking-control-plane-protecti.md) | 427–481 | deep-dive | reference |
| 12 | [9.3 Route reflector scaling and hierarchy](00-bgp-ha/9-3-route-reflector-scaling-and-hierarchy.md) | 482–511 | deep-dive | reference |
| 13 | [Wave 10 — Monitoring, telemetry, and the DC BGP troubleshooting checklist](00-bgp-ha/wave-10-monitoring-telemetry-and-the-dc-bgp-troubleshooting-.md) | 512–562 | deep-dive | reference |
| 14 | [Appendix — Consolidated open items (all waves)](00-bgp-ha/appendix-consolidated-open-items-all-waves.md) | 563–636 | deep-dive | reference |
| 15 | [11.2 BGP timer defaults reference (control-plane)](00-bgp-ha/11-2-bgp-timer-defaults-reference-control-plane.md) | 637–685 | deep-dive | reference |
| 16 | [11.5 Acronym glossary (this file)](00-bgp-ha/11-5-acronym-glossary-this-file.md) | 686–755 | deep-dive | reference |

## Par tâche

- **reference** — [Phase D3 — BGP underlay and high availability in the data center](00-bgp-ha/overview.md), [Wave 1 — BGP as the data-center underlay: design patterns, addressing, and tuning](00-bgp-ha/wave-1-bgp-as-the-data-center-underlay-design-patterns-addre.md), [1.7 BFD integration with BGP: session bring-up order and pitfalls](00-bgp-ha/1-7-bfd-integration-with-bgp-session-bring-up-order-and-pitf.md), [3.1 Cisco vPC (virtual Port Channel) — NX-OS](00-bgp-ha/3-1-cisco-vpc-virtual-port-channel-nx-os.md), [3.5 Aruba CX VSX (Virtual Switching Extension) — AOS-CX](00-bgp-ha/3-5-aruba-cx-vsx-virtual-switching-extension-aos-cx.md), [Wave 4 — First-hop redundancy, anycast gateway, ISSU/NSF, and convergence](00-bgp-ha/wave-4-first-hop-redundancy-anycast-gateway-issu-nsf-and-con.md), [Wave 5 — File verification and coverage audit](00-bgp-ha/wave-5-file-verification-and-coverage-audit.md), [Wave 6 — BGP session security, dynamic peering, and EVPN overlay control-plane design](00-bgp-ha/wave-6-bgp-session-security-dynamic-peering-and-evpn-overlay.md), [7.4 Layer 3 over MC-LAG pairs](00-bgp-ha/7-4-layer-3-over-mc-lag-pairs.md), [Wave 8 — Detection extras, FHRP tracking, control-plane protection, and GR procedures](00-bgp-ha/wave-8-detection-extras-fhrp-tracking-control-plane-protecti.md), [9.3 Route reflector scaling and hierarchy](00-bgp-ha/9-3-route-reflector-scaling-and-hierarchy.md), [Wave 10 — Monitoring, telemetry, and the DC BGP troubleshooting checklist](00-bgp-ha/wave-10-monitoring-telemetry-and-the-dc-bgp-troubleshooting-.md), [Appendix — Consolidated open items (all waves)](00-bgp-ha/appendix-consolidated-open-items-all-waves.md), [11.2 BGP timer defaults reference (control-plane)](00-bgp-ha/11-2-bgp-timer-defaults-reference-control-plane.md), [11.5 Acronym glossary (this file)](00-bgp-ha/11-5-acronym-glossary-this-file.md)
- **regulation** — [Wave 2 — Route policy in the DC fabric: prefix lists, route maps, communities, RPKI](00-bgp-ha/wave-2-route-policy-in-the-dc-fabric-prefix-lists-route-maps.md)

## Par acteur

- **Nvidia** (5) — [00-bgp-ha/overview.md](00-bgp-ha/overview.md), [00-bgp-ha/wave-1-bgp-as-the-data-center-underlay-design-patterns-addre.md](00-bgp-ha/wave-1-bgp-as-the-data-center-underlay-design-patterns-addre.md), [00-bgp-ha/3-1-cisco-vpc-virtual-port-channel-nx-os.md](00-bgp-ha/3-1-cisco-vpc-virtual-port-channel-nx-os.md), [00-bgp-ha/wave-4-first-hop-redundancy-anycast-gateway-issu-nsf-and-con.md](00-bgp-ha/wave-4-first-hop-redundancy-anycast-gateway-issu-nsf-and-con.md), [00-bgp-ha/11-5-acronym-glossary-this-file.md](00-bgp-ha/11-5-acronym-glossary-this-file.md)

## Par date

- **2026-05** — [00-bgp-ha/3-5-aruba-cx-vsx-virtual-switching-extension-aos-cx.md](00-bgp-ha/3-5-aruba-cx-vsx-virtual-switching-extension-aos-cx.md)
- **2026-09-22** — [00-bgp-ha/overview.md](00-bgp-ha/overview.md), [00-bgp-ha/wave-2-route-policy-in-the-dc-fabric-prefix-lists-route-maps.md](00-bgp-ha/wave-2-route-policy-in-the-dc-fabric-prefix-lists-route-maps.md), [00-bgp-ha/wave-5-file-verification-and-coverage-audit.md](00-bgp-ha/wave-5-file-verification-and-coverage-audit.md), [00-bgp-ha/11-5-acronym-glossary-this-file.md](00-bgp-ha/11-5-acronym-glossary-this-file.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–18 | etape6-phased3-bgp-ha/00-bgp-ha/overview.md |
| 19–74 | etape6-phased3-bgp-ha/00-bgp-ha/wave-1-bgp-as-the-data-center-underlay-design-patterns-addre.md |
| 75–96 | etape6-phased3-bgp-ha/00-bgp-ha/1-7-bfd-integration-with-bgp-session-bring-up-order-and-pitf.md |
| 97–144 | etape6-phased3-bgp-ha/00-bgp-ha/wave-2-route-policy-in-the-dc-fabric-prefix-lists-route-maps.md |
| 145–190 | etape6-phased3-bgp-ha/00-bgp-ha/3-1-cisco-vpc-virtual-port-channel-nx-os.md |
| 191–236 | etape6-phased3-bgp-ha/00-bgp-ha/3-5-aruba-cx-vsx-virtual-switching-extension-aos-cx.md |
| 237–294 | etape6-phased3-bgp-ha/00-bgp-ha/wave-4-first-hop-redundancy-anycast-gateway-issu-nsf-and-con.md |
| 295–341 | etape6-phased3-bgp-ha/00-bgp-ha/wave-5-file-verification-and-coverage-audit.md |
| 342–402 | etape6-phased3-bgp-ha/00-bgp-ha/wave-6-bgp-session-security-dynamic-peering-and-evpn-overlay.md |
| 403–426 | etape6-phased3-bgp-ha/00-bgp-ha/7-4-layer-3-over-mc-lag-pairs.md |
| 427–481 | etape6-phased3-bgp-ha/00-bgp-ha/wave-8-detection-extras-fhrp-tracking-control-plane-protecti.md |
| 482–511 | etape6-phased3-bgp-ha/00-bgp-ha/9-3-route-reflector-scaling-and-hierarchy.md |
| 512–562 | etape6-phased3-bgp-ha/00-bgp-ha/wave-10-monitoring-telemetry-and-the-dc-bgp-troubleshooting-.md |
| 563–636 | etape6-phased3-bgp-ha/00-bgp-ha/appendix-consolidated-open-items-all-waves.md |
| 637–685 | etape6-phased3-bgp-ha/00-bgp-ha/11-2-bgp-timer-defaults-reference-control-plane.md |
| 686–755 | etape6-phased3-bgp-ha/00-bgp-ha/11-5-acronym-glossary-this-file.md |

