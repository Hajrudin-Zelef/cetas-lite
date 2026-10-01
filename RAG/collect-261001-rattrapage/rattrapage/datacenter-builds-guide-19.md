---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-19
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "Samsung"]
dates: ["2026-09-16", "2026-09-27"]
keywords: ["datacenter", "amd", "compute", "dram", "gpu", "intel", "luna", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [3329, 3537]
sha256: b793684b3ab6c9af874bfab0ef762b483adcef50ed22113a3a9c10caf54abfff
---

# Datacenter Builds — Le guide des BOMs

```
PROJET : _______________  DATE : ___________  VERSION : ___
| # | Composant | Référence exacte | Qté | Prix unit. | Total | Délai | Fournisseur |
|---|-----------|------------------|-----|------------|-------|-------|-------------|
| 1 |           |                  |     |            |       |       |             |
...
|   | SOUS-TOTAL MATÉRIEL |            |       |             |
|   | Intégration (10 %)  |            |       |             |
|   | TOTAL PROJET        |            |       |             |
P_max calculée : _____ kW | P_réaliste : _____ kW | PDU : _________
Refroidissement : ________________ | Onduleur : ________________
```

---

## 202. Gabarit — plan de phases PDU

```
RACK : _____  PDU : _____ (A/B)  Date : _____
| Phase | Équipements | P (kW) | I (A) | % charge |
|-------|-------------|--------|-------|----------|
| L1    |             |        |       |          |
| L2    |             |        |       |          |
| L3    |             |        |       |          |
| TOTAL |             |        |       |          |
Écart max entre phases : ____ % (cible < 20 %)
```

---

## 203. Gabarit — plan de câblage (extrait NetBox)

```
Câble ID | De | Vers | Type | Longueur | Couleur | Date | Par
R01-C001 | R01-TOR1-E25 | R01-SRV03-E0 | DAC 25G | 3 m | bleu | ... | ...
```

---

## 204. Tableau de conversion — unités

| De | Vers | Facteur |
|---|---|---|
| kW | BTU/h | × 3 412 |
| kW | kcal/h | × 860 |
| To | Tio | × 0,9095 |
| Gb/s | Go/s | / 8 |
| 1 To à 10 Gb/s | temps de transfert | ≈ 15 min |
| 1 To à 25 Gb/s | temps de transfert | ≈ 6 min |
| 1 kW pendant 1 an | kWh | × 8 760 |
| 1 kW à 0,20 €/kWh/an | €/an | 1 752 € |

---

## 205. Dimensionnement rapide — l'antisèche du chef

```
RAM DB        = working set × 1,3
RAM Ceph      = 5 Go × To bruts + 4 Go
IOPS          = pic mesuré × 1,5
Réseau        = débit / 0,7 (jamais > 70 %)
PSU/PDU       = P_max × 1,25, chargée à 80 %
Onduleur (kW) = Σ P_réaliste × 1,25
Froid (m³/h)  = kW × 300
Batteries     = P × autonomie / 0,95
Rack          = 42U, 1 200 mm, 2 PDU A/B
Marge         = +30 % capacité à 3 ans
```

---

## 206. Normes et docs à connaître

| Référence | Sujet |
|---|---|
| ASHRAE TC 9.9 | Environnement thermique datacenter (18–27 °C) |
| Uptime Institute Tier | Niveaux de disponibilité |
| EN 50600 | Conception datacenter (Europe) |
| ISO 27001 A.11 | Sécurité physique |
| NIST 800-88 | Effacement des médias |
| FIPS 140-3 | Certification HSM |
| RGPD art. 32 | Sécurité des traitements (logs, rétention) |

---

## 207. Index des BOMs du guide

| Build | Section BOM | Section conso | Prix | P_max |
|---|---|---|---|---|
| Compute S | §18 | §19 | ≈ 29 750 € | 0,62 kW |
| Compute M | §20 | §21 | ≈ 56 975 € | 1,26 kW |
| PG OLTP S | §28 | §28 | ≈ 46 665 € | 0,79 kW |
| PG OLTP M | §29 | §29 | ≈ 100 760 € | 1,43 kW |
| PG Warehouse | §30 | §30 | ≈ 144 125 € | 1,54 kW |
| Ceph OSD NVMe S | §42 | §42 | ≈ 49 320 € | 0,78 kW |
| Ceph OSD NVMe M | §43 | §43 | ≈ 96 180 € | 1,49 kW |
| MON/MGR (×3) | §44 | §44 | ≈ 40 455 € | 1,35 kW |
| Ceph hybride S | §54 | §54 | ≈ 34 650 € | 0,74 kW |
| Ceph hybride M | §55 | §55 | ≈ 67 680 € | 1,53 kW |
| VDI S (100 users) | §63 | §63 | ≈ 64 205 € | 1,35 kW |
| VDI M vGPU (100 users) | §64 | §64 | ≈ 129 130 € | 2,68 kW |
| SIEM S | §73 | §73 | ≈ 55 895 € | 0,95 kW |
| SIEM M | §74 | §74 | ≈ 111 605 € | 1,60 kW |
| 2× MI300X | §82 | §82 | ≈ 77 560 € | 2,30 kW |
| 4× MI300X | §83 | §83 | ≈ 144 440 € | 4,40 kW |
| 2× L40S | §92 | §92 | ≈ 47 485 € | 1,38 kW |
| 4× L40S | §93 | §93 | ≈ 96 540 € | 2,69 kW |
| 2× RTX PRO 6000 | §94 | §94 | ≈ 50 585 € | 1,88 kW |
| 4× RTX PRO 6000 | §95 | §95 | ≈ 103 740 € | 3,74 kW |
| Backup S | §102 | §102 | ≈ 25 370 € | 0,70 kW |
| Backup M | §103 | §103 | ≈ 48 535 € | 0,85 kW |
| Firewall 1G HA | §110 | §110 | ≈ 940 € | 0,03 kW |
| Firewall 10G HA | §111 | §111 | ≈ 6 660 € | 0,24 kW |
| Firewall 25G HA | §112 | §112 | ≈ 14 890 € | 0,50 kW |
| HSM (paire) | §115–116 | §116 | devis | 0,60 kW |
| Rack 42U équipé | §138 | — | ≈ 9 300 € | — |
| Salle 8 racks 60 kW | §181 | §181 | ≈ 2,78 M€ | 60 kW |

---

## 208. Sources des prix (27/09/2026)

- CPU EPYC, châssis 1U/2U, RAM configurateurs : rect.coreto.de / rect.coreto-europe.com,
  ahead-it.eu, scan.co.uk (relevés 24–27/09/2026).
- GPU : snapshot Fusionww mars 2026 (via sourcebyspec.com), MSRP NVIDIA
  (TechRadar), prix cloud RunPod/Spheron/DigitalOcean (relevés 09/2026).
- SSD Samsung PM9A3 : tech-america.com, hssl.us (relevés 09/2026).
- DRAM : TrendForce via sedaily.com (16/09/2026).
- PDU APC, racks : dell.com, superwarehouse.com (relevés 09/2026).
- Firewall : shop.opnsense.com (fiches DEC), info.netgate.com (comparatif
  janvier 2026).
- Refroidissement/800VDC : datacentremagazine.com, upsite.com (AFCOM 2026),
  techtimes.com, finance.biggo.com (relevés 09/2026).
- HSM : itbrief.co.uk (annonce Thales Luna 8, 09/2026), peerspot.com.

---

## 209. Ce que ce guide ne couvre pas (volontairement)

- Le câblage électrique amont (TGBT, transfo) : domaine de l'électricien.
- Le génie civil (dalle, désenfumage) : bureau d'études.
- Les licences logicielles détaillées (VMware, Veeam) : devis éditeurs.
- La cybersécurité applicative : voir guides wazuh/fail2ban/ssh.
- L'administratif (ICPE, déclarations) : selon votre pays et puissance.

---

## 210. Historique des versions du guide

| Version | Date | Changements |
|---|---|---|
| 1.0 | 27/09/2026 | Version initiale : 12 workloads, 26 BOMs, datacenter complet |

---

## 211. Dernier mot — l'ordre des priorités

Si vous ne retenez que cinq choses :

1. **La RAM se mesure** (working set × 1,3) — à 1 500 $/64 Go, l'à-peu-près
   coûte des dizaines de milliers d'euros.
2. **Les watts s'additionnent** (§125) — tout le dimensionnement électrique
   et froid en découle.
3. **Le réseau se dimensionne au pic** (rebuild, boot storm, backup) —
   pas à la moyenne.
4. **Le backup se teste** (restore mensuel, immuabilité) — sinon il n'existe pas.
5. **L'as-built se livre** — une salle sans documentation n'est pas une salle,
   c'est un risque.

Bon build. Et quand le commercial vous dit « c'est standard », sortez la
check-list §122.

---

## 212. Comparatif CPU détaillé — EPYC Turin vs Xeon 6 (2026)

| Critère | AMD EPYC 9005 (Turin) | Intel Xeon 6 (Granite Rapids) |
|---|---|---|
| Cœurs max | 160 (9845) | 128 (6980P) |
| Canaux mémoire | 12× DDR5 | 12× DDR5 / MRDIMM |
| Lanes PCIe | 128 (5.0) en 1P | 96 (5.0) |
| TDP max | ~500 W | ~500 W |
| Prix/cœur | le plus bas du marché | +20–40 % (à vérifier) |
| Atouts | densité, prix, écosystème large | AMX (IA), QAT (crypto), certains ISV |
| Verdict guide | **choix par défaut** | si l'appli l'exige |

Le Xeon 6 se justifie pour : appliances certifiées Intel, inférence CPU avec
AMX (jusqu'à 2× vs sans), QAT pour le chiffrement/IPsec à haut débit.
Sinon, EPYC.

---

## 213. Quel EPYC pour quel workload — tableau de choix

| Workload | CPU recommandé | Pourquoi |
|---|---|---|
| Virtualisation dense | 9655P (96c) | cœurs pas chers |
| DB OLTP | 9375F / 9475F (F) | fréquence, latence |
| HPC | 9655P / 9845 | cœurs + bande passante |
| Ceph OSD | 9355P / 9555P | cœurs/€, lanes PCIe |
| VDI | 9655P (96c) | vCPU par € |
| Inférence | 9455P / 9555P | équilibre, lanes pour GPU |
| Firewall 25G+ | 9124 (16c) | suffit, économe |
| Backup | 9355P (32c) | dédup multithread |

---

## 214. Comparatif SSD NVMe entreprise (2026)

