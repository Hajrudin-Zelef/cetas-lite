---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-1
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Intel", "Nvidia"]
dates: ["2025-06-11", "2026-09-27"]
keywords: ["datacenter", "cpo", "dci", "ethernet", "gpu", "incident", "intel", "nvidia", "rubin"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [1, 114]
sha256: e9a3a5579ddd7130b0bc6b37ca94486b92b6a128fc7277697ab6a94ee2d006e4
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Guide technique dense — rédigé le 27/09/2026 — version 1.0**
**Public :** chef de service systèmes & énergies — chiffres, BOMs, dimensionnements, lien énergie.
**Méthode :** références produits vérifiées par recherche web le 27/09/2026.
Mentions utilisées dans tout le document :
- ✅ **Vérifié le 27/09/2026** = trouvé et corroboré par recherche web ce jour.
- ⚠️ **À vérifier** = ordre de grandeur usuel du marché / valeur normative, à confirmer sur datasheet ou devis distributeur.
- ❌ **Non trouvée au 27/09/2026** = recherche effectuée, aucune source fiable trouvée.

> Règle d'honnêteté : aucun prix « vérifié » n'est affirmé — les prix sont des
> fourchettes usuelles du marché marquées ⚠️, les hyperscalers ne publiant pas
> leurs tarifs. Les specs constructeur citées sont celles des datasheets.

---

## SOMMAIRE

- **Partie A — NICs serveur** (§1-15) : 25/100/200/400/800 GbE, PCIe, offloads, choix
- **Partie B — Fibre & transceivers** (§16-30) : formats, OM/OS, budgets optiques, DAC/AOC, nettoyage
- **Partie C — Architectures spine-leaf** (§31-42) : Clos, oversubscription, EVPN/VXLAN, rail-optimized
- **Partie D — Switches pour l'IA** (§43-56) : buffers, ECN, RoCEv2, IB vs Ethernet, Spectrum-X, UEC
- **Partie E — Châssis 800GbE clusters GPU** (§57-63) : radix, câblage, thermique
- **Partie F — DPU** (§64-72) : BlueField, offloads, quand en mettre
- **Partie G — CDN & edge** (§73-80) : principes, dimensionnement PoP
- **Partie H — Adressage, câblage, étiquetage** (§81-88) : plans, brassage, TIA-606
- **Partie I — Énergie** (§89-98) : conso NICs/switches/optiques, PUE, dimensionnement élec
- **Partie J — À venir (annonces vérifiées)** (§99-104) : 1.6T, UEC, CPO
- **Partie K — Approfondissements terrain** (§106-120) : supervision, tests, NCCL, stockage, sécurité, automation, TCO, migration, DCI, edge/5G, water-cooling, onduleurs, maintenance, retours terrain
- **Partie L — Fiches produits** (§126-133) : ConnectX-8, Thor 2/Ultra, BlueField-4, SN5600, TH6, Arista, E810
- **Partie M — Dimensionnements pas-à-pas** (§134-137) : cluster 2048 GPU, DC cloud, DCI, lab
- **Annexes** (§138-140) : diagnostic, table maître, index
- **Partie N — Annexes techniques denses** (§141-150) : catalogue transceivers, normes IEEE, arbre de dépannage, templates, acronymes, timeline, comparatif NOS, audit sécu, alertes Prometheus, FAQ
- **Partie O — Compléments ultra-denses** (§151-164) : comparatif 20 switches, lecture datasheet, budgets optiques, table câbles, protocole d'acceptation, RFP, pannes/MTTR, glossaire 2, fiches acteurs, refresh, multi-tenant, checklists, erreurs de câblage, conversions dB
- **Partie P — Derniers compléments** (§165-174) : étude 10 MW, test optiques, distances complètes, achat par profil, lexique FR-EN, rapport d'incident, formation, index tableaux, checklist congés, notes de version
- **Partie Q — Synthèses finales** (§175-181) : une page par débit, contrôles de congestion, connecteurs, template câblage, alimentations, FAQ 2, fin de vie
- **Partie R — Clôture opérationnelle** (§182-186) : matrice de décision, extraits de config, erreurs humaines chiffrées, ressources, revue de conception
- **Partie S — 110 décisions rapides** (§187) : le mémo ultime Si/Alors
- **§105** Glossaire (40 termes) — **§121** Quiz (10 Q/R) — **§122** Les 45 pièges terrain
- **§123** BOMs types — **§124** Checklist d'achat — **§125** Sources vérifiées

---

# PARTIE A — NICs SERVEUR : 25 / 100 / 200 / 400 / 800 GbE

## 1. Pourquoi la NIC est devenue un composant stratégique

Il y a 10 ans, la NIC était une commodité : on prenait du 10G Intel et on
n'y pensait plus. Trois ruptures ont changé la donne :

1. **L'IA générative** : l'entraînement distribué (all-reduce NCCL) sature le
   réseau inter-nœuds. Sur un cluster de 10 000 GPU, 1 % de perte de débit
   réseau = des millions de dollars de GPU qui attendent. La NIC fait partie
   du chemin critique du calcul.
2. **La vitesse a rattrapé le bus** : à 400 Gb/s, une NIC pousse autant que
   ce qu'un PCIe Gen5 x16 peut transporter (512 Gb/s théoriques). Le choix
   NIC ⇄ carte mère ⇄ CPU est un dimensionnement couplé, plus un achat isolé.
3. **Le transport a changé** : RDMA over Converged Ethernet (RoCEv2), puis
   Ultra Ethernet (UEC 1.0, ✅ 11/06/2025), déplacent la congestion, le
   multipath et la retransmission sélective **dans la NIC**. Deux NIC « 400G »
   de générations différentes n'ont plus les mêmes performances sur un fabric IA.

**Conséquence pour Zelef** : budgéter la NIC comme on budgète le GPU —
avec sa génération PCIe, sa conso (section 89) et son écosystème logiciel
(driver, firmware, OFED/DOCA).

## 2. Tableau maître : débit ⇄ format ⇄ PCIe minimal ⇄ usage

| Débit | Format dominant | Lanes élec. | PCIe minimal (x16 sauf note) | Usage typique 2026 |
|---|---|---|---|---|
| 25 GbE | SFP28 | 1×25G NRZ | Gen3 x8 | Legacy, management, OOB, vieux leafs |
| 100 GbE | QSFP28 | 4×25G NRZ | Gen4 x16 (Gen3 x16 limite) | Front-end, stockage, leafs généralistes |
| 200 GbE | QSFP56 | 4×50G PAM4 | Gen4 x16 | Transition, HPC, stockage NVMe-oF |
| 400 GbE | QSFP112 / QSFP-DD | 4×100G PAM4 | **Gen5 x16** | Backend IA standard, nœuds GPU 2024-2026 |
| 800 GbE | OSFP / QSFP-DD800 | 8×100G PAM4 | **Gen6 x16** | Backend IA 2025+, scale-out Rubin |

Notes :
- NRZ = 1 bit/symbole ; PAM4 = 2 bits/symbole. Le passage à PAM4 (50G/lane
  puis 100G/lane) est ce qui a permis 400/800G sans multiplier les fibres.
- Un slot PCIe Gen4 x16 plafonne à 256 Gb/s théoriques (~200 Gb/s utiles) :
  **y brancher une NIC 400G = la brider à ~50 %**. C'est le piège n°1 des
  achats NIC (voir §107, P1).
- PCIe Gen6 x16 = 128 Go/s théoriques (~1024 Gb/s) : juste assez pour 800G
  plein débit + overhead. ✅ ConnectX-8 documenté PCIe Gen6 (firmware docs
  NVIDIA, vérifié le 27/09/2026).

## 3. Le 25 GbE : mort ou toujours vivant ?

Le 25G (SFP28, 1 lane 25G NRZ) reste pertinent en 2026 pour :
- **OOB / management** : iDRAC/iLO, consoles — souvent encore 1G, mais les
  parcs neufs passent au 25G sur les leafs de management.
- **Nœuds légers** : hyperviseurs de services, monitoring, bastions.
- **Breakout** : un port 100G QSFP28 se décompose en 4×25G (câble breakout).

Références courantes (⚠️ gamme établie, vérifier dispo) : Intel XXV710,
NVIDIA ConnectX-4 Lx / ConnectX-6 Dx en mode 25G, Broadcom NetXtreme.
**Ne plus acheter de 10G neuf** sauf contrainte d'équipement existant : le
prix au Gb/s du 25G est inférieur depuis des années.

## 4. Le 100 GbE : le standard du front-end

Le 100G (QSFP28, 4×25G NRZ) est en 2026 le débit « banalisé » : serveurs
généralistes, stockage, interconnexion de baies, uplinks de leafs 25G.

**Références actuelles :**

| NIC | Débit/ports | PCIe | Points notables | Statut |
|---|---|---|---|---|
| NVIDIA ConnectX-6 Dx | 2×100G (ou 1×200G) | Gen4 x16 | RoCEv2, la valeur sûre du 100/200G | ✅ gamme établie |
| NVIDIA ConnectX-7 | 2×100/200G ou 1×400G | Gen5 x16 | RDMA, adaptatif, base des backend IA 400G | ✅ vérifié 27/09/2026 |
| Intel E810-CAM1/XXVDA2 | 2×100G QSFP28 | Gen4 x16 | 256 VF SR-IOV, ADQ, PTP, ~15-27 W | ✅ vérifié 27/09/2026 |
| Broadcom NetXtreme BCM575xx | 1/2×100G | Gen4 | Alternative OEM massive (Dell/HPE) | ⚠️ à vérifier |
| Marvell FastLinQ 45000 | 2×100G | Gen4 | Présent chez les OEM | ⚠️ à vérifier |

