# INDEX — Step 6 — Phase D wave 2: EVPN-VXLAN overlay

Corpus `etape6-phased2-evpn-vxlan` · **16 fichiers** · 879 lignes source · ~6301 mots · partition exacte de `docs/RAG/etape6_phaseD2_evpn_vxlan.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-evpn-vxlan/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Step 6 — Phase D wave 2: EVPN-VXLAN overlay](00-evpn-vxlan/overview.md) | 1–53 | deep-dive | reference |
| 02 | [1.2 Packet format (RFC 7348 §4)](00-evpn-vxlan/1-2-packet-format-rfc-7348-4.md) | 54–100 | deep-dive | quantization |
| 03 | [1.4 VTEP and NVE](00-evpn-vxlan/1-4-vtep-and-nve.md) | 101–147 | deep-dive | reference |
| 04 | [1.6 Underlay requirements](00-evpn-vxlan/1-6-underlay-requirements.md) | 148–192 | deep-dive | reference |
| 05 | [Wave 2 — EVPN control plane](00-evpn-vxlan/wave-2-evpn-control-plane.md) | 193–245 | deep-dive | reference |
| 06 | [2.3 EVPN building blocks: EVI, RD, RT, ESI](00-evpn-vxlan/2-3-evpn-building-blocks-evi-rd-rt-esi.md) | 246–306 | deep-dive | reference |
| 07 | [2.6 Distributed anycast gateway](00-evpn-vxlan/2-6-distributed-anycast-gateway.md) | 307–360 | deep-dive | reference |
| 08 | [3.2 ESI (Ethernet Segment Identifier)](00-evpn-vxlan/3-2-esi-ethernet-segment-identifier.md) | 361–412 | deep-dive | reference |
| 09 | [3.5 Aliasing and backup paths](00-evpn-vxlan/3-5-aliasing-and-backup-paths.md) | 413–478 | deep-dive | reference |
| 10 | [4.3 EVPN-VXLAN ↔ MPLS-VPN stitching (DCI L3 gateway)](00-evpn-vxlan/4-3-evpn-vxlan-mpls-vpn-stitching-dci-l3-gateway.md) | 479–526 | deep-dive | reference |
| 11 | [Wave 5 — Vendor implementations and interop](00-evpn-vxlan/wave-5-vendor-implementations-and-interop.md) | 527–582 | deep-dive | reference |
| 12 | [5.2 Arista EOS — Vxlan1 interface + MLAG/ESI](00-evpn-vxlan/5-2-arista-eos-vxlan1-interface-mlag-esi.md) | 583–629 | deep-dive | reference |
| 13 | [5.3 NVIDIA — Cumulus Linux (FRR) and SONiC](00-evpn-vxlan/5-3-nvidia-cumulus-linux-frr-and-sonic.md) | 630–696 | deep-dive | actor-profile |
| 14 | [5.5 Aruba CX (HPE Aruba Networking)](00-evpn-vxlan/5-5-aruba-cx-hpe-aruba-networking.md) | 697–745 | deep-dive | reference |
| 15 | [5.8 Gaps — Wave 5](00-evpn-vxlan/5-8-gaps-wave-5.md) | 746–821 | deep-dive | reference |
| 16 | [6.3 Common failure modes](00-evpn-vxlan/6-3-common-failure-modes.md) | 822–879 | deep-dive | reference |

## Par tâche

- **actor-profile** — [5.3 NVIDIA — Cumulus Linux (FRR) and SONiC](00-evpn-vxlan/5-3-nvidia-cumulus-linux-frr-and-sonic.md)
- **quantization** — [1.2 Packet format (RFC 7348 §4)](00-evpn-vxlan/1-2-packet-format-rfc-7348-4.md)
- **reference** — [Step 6 — Phase D wave 2: EVPN-VXLAN overlay](00-evpn-vxlan/overview.md), [1.4 VTEP and NVE](00-evpn-vxlan/1-4-vtep-and-nve.md), [1.6 Underlay requirements](00-evpn-vxlan/1-6-underlay-requirements.md), [Wave 2 — EVPN control plane](00-evpn-vxlan/wave-2-evpn-control-plane.md), [2.3 EVPN building blocks: EVI, RD, RT, ESI](00-evpn-vxlan/2-3-evpn-building-blocks-evi-rd-rt-esi.md), [2.6 Distributed anycast gateway](00-evpn-vxlan/2-6-distributed-anycast-gateway.md), [3.2 ESI (Ethernet Segment Identifier)](00-evpn-vxlan/3-2-esi-ethernet-segment-identifier.md), [3.5 Aliasing and backup paths](00-evpn-vxlan/3-5-aliasing-and-backup-paths.md), [4.3 EVPN-VXLAN ↔ MPLS-VPN stitching (DCI L3 gateway)](00-evpn-vxlan/4-3-evpn-vxlan-mpls-vpn-stitching-dci-l3-gateway.md), [Wave 5 — Vendor implementations and interop](00-evpn-vxlan/wave-5-vendor-implementations-and-interop.md), [5.2 Arista EOS — Vxlan1 interface + MLAG/ESI](00-evpn-vxlan/5-2-arista-eos-vxlan1-interface-mlag-esi.md), [5.5 Aruba CX (HPE Aruba Networking)](00-evpn-vxlan/5-5-aruba-cx-hpe-aruba-networking.md), [5.8 Gaps — Wave 5](00-evpn-vxlan/5-8-gaps-wave-5.md), [6.3 Common failure modes](00-evpn-vxlan/6-3-common-failure-modes.md)

## Par acteur

- **Nvidia** (4) — [00-evpn-vxlan/overview.md](00-evpn-vxlan/overview.md), [00-evpn-vxlan/5-3-nvidia-cumulus-linux-frr-and-sonic.md](00-evpn-vxlan/5-3-nvidia-cumulus-linux-frr-and-sonic.md), [00-evpn-vxlan/5-8-gaps-wave-5.md](00-evpn-vxlan/5-8-gaps-wave-5.md), [00-evpn-vxlan/6-3-common-failure-modes.md](00-evpn-vxlan/6-3-common-failure-modes.md)

## Par date

- **2026-04-23** — [00-evpn-vxlan/wave-2-evpn-control-plane.md](00-evpn-vxlan/wave-2-evpn-control-plane.md)
- **2026-09-22** — [00-evpn-vxlan/overview.md](00-evpn-vxlan/overview.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–53 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/overview.md |
| 54–100 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/1-2-packet-format-rfc-7348-4.md |
| 101–147 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/1-4-vtep-and-nve.md |
| 148–192 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/1-6-underlay-requirements.md |
| 193–245 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/wave-2-evpn-control-plane.md |
| 246–306 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/2-3-evpn-building-blocks-evi-rd-rt-esi.md |
| 307–360 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/2-6-distributed-anycast-gateway.md |
| 361–412 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/3-2-esi-ethernet-segment-identifier.md |
| 413–478 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/3-5-aliasing-and-backup-paths.md |
| 479–526 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/4-3-evpn-vxlan-mpls-vpn-stitching-dci-l3-gateway.md |
| 527–582 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/wave-5-vendor-implementations-and-interop.md |
| 583–629 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/5-2-arista-eos-vxlan1-interface-mlag-esi.md |
| 630–696 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/5-3-nvidia-cumulus-linux-frr-and-sonic.md |
| 697–745 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/5-5-aruba-cx-hpe-aruba-networking.md |
| 746–821 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/5-8-gaps-wave-5.md |
| 822–879 | etape6-phased2-evpn-vxlan/00-evpn-vxlan/6-3-common-failure-modes.md |

