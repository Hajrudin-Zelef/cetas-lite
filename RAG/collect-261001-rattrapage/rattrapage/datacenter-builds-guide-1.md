---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-1
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "Samsung"]
dates: ["2026-09-27"]
keywords: ["datacenter", "amd", "blackwell", "dram", "gpu", "hbm", "hbm3", "hyperscaler", "intel", "liquid cooling"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [1, 148]
sha256: 356d82fd52ced94e86dc95dc6c88be084204d387a44623030eb93c526790f8f5
---

# Datacenter Builds — Le guide des BOMs

**Builds serveurs complets par workload + montage d'un datacenter.**
Chaque build = une BOM complète et chiffrée : châssis, CPU, RAM, stockage, GPU,
NICs, alimentation — avec ordre de prix total et consommation électrique estimée.
Puis : rack, PDU, câblage, refroidissement, onduleurs, mise en service, PME vs hyperscaler.

**Public :** Zelef, chef de service systèmes & énergies. Ce guide parle BOM, watts et euros.
**Langue :** français. **Date de référence des prix : 27/09/2026.**

> **Avertissement honnête.** Les prix serveur sont des prix publics indicatifs relevés
> sur le web le 27/09/2026 (boutiques européennes, configurateurs, snapshots de marché).
> Un vrai devis intégrateur varie de ±15 à 30 % selon volumes, remises et délais.
> Les TDP sont des ordres de grandeur : vérifiez la fiche constructeur exacte avant de
> dimensionner l'électrique. Quand une info n'a pas pu être vérifiée, c'est marqué
> « à vérifier ». Quand un composant est introuvable, c'est marqué
> « non trouvé au 27/09/2026 ». Rien n'est inventé.

---

## 1. Comment utiliser ce guide

1. Lisez la **partie A (sections 2 à 16)** : l'état du marché et les formules de
   dimensionnement. C'est le socle : sans ça, les BOMs ne sont que des listes de courses.
2. Pour chaque workload, prenez la BOM **S** (petit) ou **M** (moyen), ou les variantes
   **double GPU / quad GPU** quand c'est pertinent.
3. Additionnez les consommations estimées : c'est l'entrée de la **partie D**
   (rack, PDU, refroidissement, onduleurs).
4. La **partie E** donne la méthode pour chiffrer seul, la **partie F** le montage
   datacenter complet, la **partie G** l'écart PME/hyperscaler.
5. Glossaire (30+ termes), quiz (10 questions + réponses) et pièges terrain en fin.

Toutes les BOMs sont en **euros HT indicatifs**, sauf mention contraire.
Taux de change retenu : 1 $ ≈ 0,92 € (ordre de grandeur, septembre 2026 — à vérifier
au jour du devis).

---

## 2. Conventions de prix et mentions de vérification

| Mention | Signification |
|---|---|
| **vérifié le 27/09/2026** | Prix ou référence relevé sur le web à cette date (source citée) |
| **à vérifier** | Ordre de grandeur plausible, pas de source directe ce jour-là |
| **non trouvé au 27/09/2026** | Recherché, introuvable publiquement |
| **TDP indicatif** | Valeur constructeur arrondie ; vérifier la fiche exacte |

Règle de lecture des BOMs : colonne « Prix » = prix unitaire public constaté ;
« Total » = quantité × prix. Les totaux sont arrondis. La main-d'œuvre d'intégration
n'est pas incluse (compter 5 à 10 % du matériel pour un montage + burn-in pro).

---

## 3. État du marché au 27/09/2026 — ce qui a changé

Trois faits structurent tous les chiffrages de ce guide.

**1. La DRAM serveur a explosé.** Prix contrat d'un module DDR5 RDIMM 64 Go :
**≈ 1 500 $ en septembre 2026**, soit 5,5× le niveau d'un an plus tôt (272 $) ;
spot jusqu'à 3 100 $ (TrendForce via presse coréenne, vérifié le 27/09/2026).
Cause : Samsung, SK hynix et Micron allouent leurs capacités au HBM, à la DDR5
serveur et aux SSD entreprise — l'IA agentique tire la demande. Conséquence
directe : **la RAM est redevenue le premier poste de coût d'un serveur**
(devant le CPU sur beaucoup de builds). Dimensionnez la RAM au plus juste,
privilégiez 1 DIMM par canal, et négociez les gros volumes tôt.

**2. Les GPU restent sous tension d'allocation.** Snapshot de marché mars 2026
(vérifié le 27/09/2026) : L40S 48 Go à 8 610–8 900 $, RTX PRO 6000 96 Go à
9 450–9 800 $, H100 80 Go SXM neuf à 25 000–35 000 $ (occasion 18 000–22 000 $),
H200 141 Go à 30 000–40 000 $, B200 à 30 000–40 000 $. Délais Blackwell PRO :
3 à 7 mois début 2026. Le MI300X AMD (192 Go HBM3) se traite autour de
10 000–12 000 $ pièce en standalone — l'alternative mémoire/prix pour l'inférence
gros modèles.

**3. Le refroidissement liquide est devenu mainstream.** Enquête AFCOM 2026 :
36 % des opérateurs ont déployé du liquid cooling, 28 % prévoient de le faire sous
12–24 mois ; 37 % ont adopté ou envisagent le rear-door heat exchanger (RDHx),
forme dominante car la moins invasive. Le direct-to-chip monophasé représente
≈ 55 % du marché DTC 2026. Repères de densité : air OK jusqu'à ~20–30 kW/rack,
hybride 50–100 kW, DTC obligatoire au-delà de 100 kW/rack. Côté puissance, le
**800VDC** arrive : produits commerciaux Vertiv, Schneider Electric, Eaton, Delta
attendus au **2ᵉ semestre 2026** (vérifié le 27/09/2026).

---

## 4. CPU : AMD EPYC domine le rack en 2026

Pour les workloads de ce guide, **AMD EPYC 9005 « Turin »** (socket SP5, DDR5
12 canaux, PCIe 5.0) est le choix par défaut : densité de cœurs, lanes PCIe
(128 en mono-socket) et prix/cœur. Intel Xeon 6 (Granite Rapids) reste pertinent
quand l'écosystème l'impose (certains appliances, AMX pour l'inférence CPU),
mais en 2026 le rapport prix/perf du rack généraliste est chez AMD.
Prix publics constatés sur configurateur européen (vérifiés le 27/09/2026) :

| CPU | Cœurs/Threads | TDP indicatif | Prix constaté |
|---|---|---|---|
| EPYC 9124 (Genoa) | 16/32 | 200 W | ≈ 750 € |
| EPYC 9115 (Turin) | 16/32 | 200 W | ≈ 185 € (prix promo constaté — à vérifier, anomalie possible) |
| EPYC 9275F | 24/48 | 320 W | ≈ 2 784 € |
| EPYC 9355P | 32/64 | 280 W | ≈ 2 891 € |
| EPYC 9375F | 32/64 | 320 W | ≈ 3 893 € |
| EPYC 9455P | 48/96 | 300 W | ≈ 3 406 € |
| EPYC 9475F | 48/96 | 360 W | ≈ 4 790 € |
| EPYC 9555P | 64/128 | 360 W | ≈ 4 830 € |
| EPYC 9575F | 64/128 | 400 W | ≈ 6 731 € |
| EPYC 9655P | 96/192 | 400 W | ≈ 5 208 € |
| EPYC 9825 | 144/288 | 390 W | ≈ 7 774 € |
| EPYC 9845 | 160/320 | 390 W | ≈ 8 764 € |

Règles terrain :
- **Suffixe P** = mono-socket optimisé (moins cher, le bon choix en 1P).
  **Suffixe F** = fréquence haute (bases de données, VDI, latence).
- Au-delà de 64 cœurs, vérifiez que votre workload scale : beaucoup d'applis
  paient des cœurs qu'elles n'utilisent pas. Le sweet spot 2026 : **32–64 cœurs**.
- En bi-socket, la bande passante mémoire double mais la latence NUMA apparaît :
  réservez le 2P au HPC, aux grosses DB et à la virtualisation dense.

---

## 5. RAM DDR5 : le poste qui fait mal en 2026

Repères (vérifiés le 27/09/2026) :
- RDIMM DDR5-5600 ECC 64 Go : **≈ 1 500 $ le module** en contrat (≈ 1 380 €),
  spot jusqu'à ≈ 3 100 $ en pointe. Il y a un an : 272 $.
- RDIMM DDR5 128 Go : ≈ 2 800–3 000 $ constatés sur configurateurs (≈ 2 600–2 750 €).
- RDIMM DDR5 96 Go : position intermédiaire, ≈ 2 100 $ (à vérifier).

Règles de dimensionnement :
1. **1 DIMM par canal, 12 canaux remplis** sur EPYC = pleine bande passante.
   12 × 64 Go = 768 Go est la configuration standard 2026.
2. Ne mélangez ni capacités ni rangs sur un même canal si vous voulez la
   fréquence nominale (le BIOS dégrade sinon).
3. ECC Registered obligatoire en serveur. Le prix au Go chute avec la densité
   du module… sauf en 2026 où tout est cher : comparez 12×64 Go vs 12×96 Go
   au moment du devis, l'écart varie vite.
4. Pour les DB : RAM = working set + 30 % de marge (section 33).
   Pour Ceph OSD : 4–5 Go de RAM par To de stockage + base OS (section 43).

**Astuce budget :** sur un nœud Ceph ou backup, 384 Go (12×32 Go) suffisent
souvent — n'achetez pas 768 Go « au cas où » à 1 500 $ le module.

---

## 6. Stockage : NVMe U.2, le standard du datacenter

Le 2,5" U.2 NVMe PCIe 4.0/5.0 est le format de référence. Prix constatés
(vérifiés le 27/09/2026) :

