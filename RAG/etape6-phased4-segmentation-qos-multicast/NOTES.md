# NOTES — corpus `etape6-phased4-segmentation-qos-multicast`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-segmentation-qos-multicast` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services` — 770 lignes, 14 chunks.

- Verification log (gaps / conflicts / unverified)
- Wave 1 — VLAN design at scale
- Wave 2 — VRF: VRF-lite, leaking, multi-tenancy
- Wave 3 — Microsegmentation & policy
- Wave 4 — QoS: classification, marking, queuing
- Wave 5 — Data Center Bridging: PFC, ETS, DCBX, QCN
- Wave 6 — ECN, RoCEv2 lossless tuning, vendor QoS models
- Wave 7 — Multicast: PIM-SM/SSM, RP design, MSDP, Anycast-RP
- Wave 8 — Multicast in VXLAN overlays + IGMP
- Wave 9 — Network services: DHCP relay, DNS, NTP/PTP
- Coverage audit & final verification log
- Wave 10 — EVPN control plane and segmentation (route types, RD/RT, IRB)
- Wave 11 — VRF deep-dive: RD/RT mechanics, VRF-aware services
- Wave 12 — QoS deep-dive: MQC walkthrough, Nexus system classes, Arista model
- Wave 13 — RoCE congestion control algorithms & Ultra Ethernet
- Wave 14 — Multicast deep-dive: PIM mechanics, BSR/Auto-RP, mVPN
- Wave 15 — Services deep-dive: DHCPv6/SLAAC, DNS, NTP/NTS, PTP profiles
- Final verification (post Waves 10–15)
- Wave 16 — Segmentation in practice: compliance, DMZ, service insertion
- Wave 17 — Multicast in practice: market data, TRM checklist, troubleshooting
- Wave 18 — QoS in practice: traffic-class plan, buffer math, validation
- Final verification (all 18 waves)
- Wave 19 — Config walkthroughs: EVPN multi-tenant VRF (NX-OS)
- Wave 20 — Config walkthroughs: RoCEv2 lossless QoS (Dell OS10)
- Wave 21 — Config walkthroughs: anycast-RP + TRM outline
- Final verification (all 21 waves)
- Wave 22 — Reference tables
- Final verification (all 22 waves)

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
