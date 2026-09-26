# INDEX — Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services

Corpus `etape6-phased4-segmentation-qos-multicast` · **17 fichiers** · 770 lignes source · ~10858 mots · partition exacte de `docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-segmentation-qos-multicast/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services](00-segmentation-qos-multicast/overview.md) | 1–45 | deep-dive | reference |
| 02 | [Wave 2 — VRF: VRF-lite, leaking, multi-tenancy](00-segmentation-qos-multicast/wave-2-vrf-vrf-lite-leaking-multi-tenancy.md) | 46–91 | deep-dive | reference |
| 03 | [3.5 Zero-trust in the DC](00-segmentation-qos-multicast/3-5-zero-trust-in-the-dc.md) | 92–101 | deep-dive | reference |
| 04 | [Wave 4 — QoS: classification, marking, queuing](00-segmentation-qos-multicast/wave-4-qos-classification-marking-queuing.md) | 102–150 | deep-dive | reference |
| 05 | [6.1 The three-layer lossless model](00-segmentation-qos-multicast/6-1-the-three-layer-lossless-model.md) | 151–191 | deep-dive | reference |
| 06 | [7.3 Anycast-RP (RFC 3446)](00-segmentation-qos-multicast/7-3-anycast-rp-rfc-3446.md) | 192–206 | deep-dive | reference |
| 07 | [Wave 8 — Multicast in VXLAN overlays + IGMP](00-segmentation-qos-multicast/wave-8-multicast-in-vxlan-overlays-igmp.md) | 207–245 | deep-dive | reference |
| 08 | [9.3 NTP/PTP timing](00-segmentation-qos-multicast/9-3-ntp-ptp-timing.md) | 246–256 | deep-dive | reference |
| 09 | [Coverage audit & final verification log](00-segmentation-qos-multicast/coverage-audit-final-verification-log.md) | 257–330 | deep-dive | reference |
| 10 | [Wave 11 — VRF deep-dive: RD/RT mechanics, VRF-aware services](00-segmentation-qos-multicast/wave-11-vrf-deep-dive-rd-rt-mechanics-vrf-aware-services.md) | 331–379 | deep-dive | reference |
| 11 | [Wave 13 — RoCE congestion control algorithms & Ultra Ethernet](00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md) | 380–425 | deep-dive | reference |
| 12 | [Wave 15 — Services deep-dive: DHCPv6/SLAAC, DNS, NTP/NTS, PTP profiles](00-segmentation-qos-multicast/wave-15-services-deep-dive-dhcpv6-slaac-dns-ntp-nts-ptp-prof.md) | 426–485 | deep-dive | reference |
| 13 | [Wave 17 — Multicast in practice: market data, TRM checklist, troubleshooting](00-segmentation-qos-multicast/wave-17-multicast-in-practice-market-data-trm-checklist-trou.md) | 486–534 | deep-dive | reference |
| 14 | [Final verification (all 18 waves)](00-segmentation-qos-multicast/final-verification-all-18-waves.md) | 535–611 | deep-dive | reference |
| 15 | [Wave 20 — Config walkthroughs: RoCEv2 lossless QoS (Dell OS10)](00-segmentation-qos-multicast/wave-20-config-walkthroughs-rocev2-lossless-qos-dell-os10.md) | 612–658 | deep-dive | reference |
| 16 | [Wave 21 — Config walkthroughs: anycast-RP + TRM outline](00-segmentation-qos-multicast/wave-21-config-walkthroughs-anycast-rp-trm-outline.md) | 659–708 | deep-dive | reference |
| 17 | [Wave 22 — Reference tables](00-segmentation-qos-multicast/wave-22-reference-tables.md) | 709–770 | deep-dive | reference |

## Par tâche

- **reference** — [Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services](00-segmentation-qos-multicast/overview.md), [Wave 2 — VRF: VRF-lite, leaking, multi-tenancy](00-segmentation-qos-multicast/wave-2-vrf-vrf-lite-leaking-multi-tenancy.md), [3.5 Zero-trust in the DC](00-segmentation-qos-multicast/3-5-zero-trust-in-the-dc.md), [Wave 4 — QoS: classification, marking, queuing](00-segmentation-qos-multicast/wave-4-qos-classification-marking-queuing.md), [6.1 The three-layer lossless model](00-segmentation-qos-multicast/6-1-the-three-layer-lossless-model.md), [7.3 Anycast-RP (RFC 3446)](00-segmentation-qos-multicast/7-3-anycast-rp-rfc-3446.md), [Wave 8 — Multicast in VXLAN overlays + IGMP](00-segmentation-qos-multicast/wave-8-multicast-in-vxlan-overlays-igmp.md), [9.3 NTP/PTP timing](00-segmentation-qos-multicast/9-3-ntp-ptp-timing.md), [Coverage audit & final verification log](00-segmentation-qos-multicast/coverage-audit-final-verification-log.md), [Wave 11 — VRF deep-dive: RD/RT mechanics, VRF-aware services](00-segmentation-qos-multicast/wave-11-vrf-deep-dive-rd-rt-mechanics-vrf-aware-services.md), [Wave 13 — RoCE congestion control algorithms & Ultra Ethernet](00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md), [Wave 15 — Services deep-dive: DHCPv6/SLAAC, DNS, NTP/NTS, PTP profiles](00-segmentation-qos-multicast/wave-15-services-deep-dive-dhcpv6-slaac-dns-ntp-nts-ptp-prof.md), [Wave 17 — Multicast in practice: market data, TRM checklist, troubleshooting](00-segmentation-qos-multicast/wave-17-multicast-in-practice-market-data-trm-checklist-trou.md), [Final verification (all 18 waves)](00-segmentation-qos-multicast/final-verification-all-18-waves.md), [Wave 20 — Config walkthroughs: RoCEv2 lossless QoS (Dell OS10)](00-segmentation-qos-multicast/wave-20-config-walkthroughs-rocev2-lossless-qos-dell-os10.md), [Wave 21 — Config walkthroughs: anycast-RP + TRM outline](00-segmentation-qos-multicast/wave-21-config-walkthroughs-anycast-rp-trm-outline.md), [Wave 22 — Reference tables](00-segmentation-qos-multicast/wave-22-reference-tables.md)

## Par acteur

- **Google** (1) — [00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md](00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md)
- **Huawei** (4) — [00-segmentation-qos-multicast/wave-2-vrf-vrf-lite-leaking-multi-tenancy.md](00-segmentation-qos-multicast/wave-2-vrf-vrf-lite-leaking-multi-tenancy.md), [00-segmentation-qos-multicast/7-3-anycast-rp-rfc-3446.md](00-segmentation-qos-multicast/7-3-anycast-rp-rfc-3446.md), [00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md](00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md), [00-segmentation-qos-multicast/wave-21-config-walkthroughs-anycast-rp-trm-outline.md](00-segmentation-qos-multicast/wave-21-config-walkthroughs-anycast-rp-trm-outline.md)
- **Meta** (1) — [00-segmentation-qos-multicast/6-1-the-three-layer-lossless-model.md](00-segmentation-qos-multicast/6-1-the-three-layer-lossless-model.md)
- **Nvidia** (5) — [00-segmentation-qos-multicast/wave-4-qos-classification-marking-queuing.md](00-segmentation-qos-multicast/wave-4-qos-classification-marking-queuing.md), [00-segmentation-qos-multicast/6-1-the-three-layer-lossless-model.md](00-segmentation-qos-multicast/6-1-the-three-layer-lossless-model.md), [00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md](00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md), [00-segmentation-qos-multicast/wave-17-multicast-in-practice-market-data-trm-checklist-trou.md](00-segmentation-qos-multicast/wave-17-multicast-in-practice-market-data-trm-checklist-trou.md), [00-segmentation-qos-multicast/wave-20-config-walkthroughs-rocev2-lossless-qos-dell-os10.md](00-segmentation-qos-multicast/wave-20-config-walkthroughs-rocev2-lossless-qos-dell-os10.md)

## Par date

- **2026-09-22** — [00-segmentation-qos-multicast/overview.md](00-segmentation-qos-multicast/overview.md), [00-segmentation-qos-multicast/wave-2-vrf-vrf-lite-leaking-multi-tenancy.md](00-segmentation-qos-multicast/wave-2-vrf-vrf-lite-leaking-multi-tenancy.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–45 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/overview.md |
| 46–91 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-2-vrf-vrf-lite-leaking-multi-tenancy.md |
| 92–101 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/3-5-zero-trust-in-the-dc.md |
| 102–150 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-4-qos-classification-marking-queuing.md |
| 151–191 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/6-1-the-three-layer-lossless-model.md |
| 192–206 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/7-3-anycast-rp-rfc-3446.md |
| 207–245 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-8-multicast-in-vxlan-overlays-igmp.md |
| 246–256 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/9-3-ntp-ptp-timing.md |
| 257–330 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/coverage-audit-final-verification-log.md |
| 331–379 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-11-vrf-deep-dive-rd-rt-mechanics-vrf-aware-services.md |
| 380–425 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet.md |
| 426–485 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-15-services-deep-dive-dhcpv6-slaac-dns-ntp-nts-ptp-prof.md |
| 486–534 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-17-multicast-in-practice-market-data-trm-checklist-trou.md |
| 535–611 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/final-verification-all-18-waves.md |
| 612–658 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-20-config-walkthroughs-rocev2-lossless-qos-dell-os10.md |
| 659–708 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-21-config-walkthroughs-anycast-rp-trm-outline.md |
| 709–770 | etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-22-reference-tables.md |

