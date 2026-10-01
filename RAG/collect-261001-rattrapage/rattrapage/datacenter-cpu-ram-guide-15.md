---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-15
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Samsung"]
dates: ["2026-06-01", "2026-07-23", "2026-09-27"]
keywords: ["amd", "arr", "capex", "clearwater forest", "datacenter", "dram", "intel"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [2298, 2456]
sha256: 283c38982a49c47f671f5f7d33569fa4284ac0721759e43423af7521aec66406
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

**Q7.** Pourquoi faut-il obturer les U vides d'une baie avec des caches ?

**Q8.** Votre onduleur est dimensionné sur la consommation moyenne des
serveurs. Quel est le risque au redémarrage après une coupure ?

**Q9.** Citez trois différences entre un RDIMM et un MRDIMM, et un cas où le
MRDIMM ne sert à rien.

**Q10.** Un commercial vous propose du DDR5-6400 premium pour un cluster de
virtualisation généraliste. Que répondez-vous, et que vérifiez-vous d'abord ?

## 126. Réponses du quiz

**R1.** 12 × 6400 MT/s × 8 octets = 614 Go/s ; / 192 cœurs = **3,2 Go/s/cœur**.
C'est faible pour du HPC (les codes gourmands veulent 5-10 Go/s/cœur) :
le 9965 est fait pour le scale-out (web, conteneurs), pas pour la
simulation. Pour du HPC, un 9755 (4,8 Go/s/cœur) ou du MRDIMM.

**R2.** AMX = accélérateur matriciel par cœur, conçu pour les opérations
d'inférence (INT8/BF16). À cœurs équivalents, le débit d'inférence d'un
Xeon P-core dépasse largement celui d'un EPYC qui n'a que l'AVX-512/VNNI.
Pour l'inférence CPU, le choix se fait sur l'accélérateur, pas sur le
nombre de cœurs.

**R3.** **12× 64 Go** : les 12 canaux sont remplis → pleine bande passante
(537 Go/s théoriques à 5600). Avec 6× 128 Go, 6 canaux sont vides → moitié
moins de débit pour la même capacité, souvent plus cher.

**R4.** L'on-die ECC ne protège que l'intérieur des puces DRAM ; le side-band
ECC (barrettes ECC Registered) protège aussi le bus et remonte les erreurs
au système. **Seul le side-band ECC suffit en production** : l'on-die seul
laisse passer les erreurs de bus silencieusement.

**R5.** Non : le 2S ajoute l'interconnexion inter-socket (UPI/Infinity
Fabric, quelques dizaines de watts), un 2e jeu de VRM et de DIMM, et les
deux sockets ne descendent jamais aussi bas en idle qu'un seul. Comptez
**10-20 % de plus** à charge égale (ordre de grandeur).

**R6.** I = 15 000 / (1,732 × 400 × 0,95) ≈ **22,8 A par phase**. Avec la règle
des 50 % (une PDU doit tenir seule toute la baie), il faut une PDU triphasée
**32 A** par baie (2 PDU : A et B), chargée à ~36 % en nominal.

**R7.** Un U vide sans cache crée un court-circuit aéraulique : l'air frais
passe à côté des serveurs au lieu de les traverser → +10 °C possibles à
l'entrée des serveurs, ventilateurs qui s'emballent (loi cubique), sur-
consommation et bruit.

**R8.** L'appel de courant au démarrage simultané (ou le pic de charge)
dépasse la moyenne : l'onduleur surchargé disjoncte ou bascule en bypass
au moment où on a le plus besoin de lui. **On dimensionne sur le pic + 25 %.**

**R9.** Différences : (1) le MRDIMM multiplexe 2 rangs simultanément →
8800 MT/s vs 6400 ; (2) il coûte ~100 $ de plus par barrette 64 Go ;
(3) il consomme nettement plus (+437 W mesurés sur 12 barrettes en charge
mémoire). Inutile pour : virtualisation généraliste, workloads peu sensibles
à la bande passante — le surcoût (prix + énergie) n'apporte rien.

**R10.** Je refuse le premium : en virtualisation généraliste, le goulot
n'est pas la bande passante mémoire. Je vérifie d'abord que **tous les
canaux sont remplis** (c'est ça qui fait le débit), puis je prends la
fréquence standard (5600) et j'investis la différence en capacité ou en
contrat de maintenance.

## 127. BOM type n°1 : nœud virtualisation 2S (indicatif)

Base du cas chiffré de la section 31 — devis à faire valider par l'OEM.

| # | Composant | Spécification | Qté |
|---|---|---|---|
| 1 | Châssis | 2U, 2S, redondance fans N+1 confirmée | 1 |
| 2 | CPU | AMD EPYC 9655 (96 c., 400 W) ou équiv. | 2 |
| 3 | Dissipateurs | 2U, compatibles 400 W | 2 |
| 4 | RAM | 64 Go DDR5-5600 RDIMM ECC (QVL) | 24 |
| 5 | Stockage | NVMe 3,84 To U.2, DWPD ≥ 1 | 8 |
| 6 | Réseau | 2× 25 GbE SFP28 | 1 carte |
| 7 | Alimentations | 1600 W 80 PLUS Titanium, redondantes | 2 |
| 8 | Rails | télescopiques, compatibles baie | 1 kit |
| 9 | BMC | licence management incluse | 1 |
| 10 | Garantie | 5 ans, sur site J+1 | 1 |

Ordre de grandeur CAPEX : **20 000-30 000 €** (indicatif, très dépendant
des remises). Postes à ne jamais rogner : RAM (QVL), alimentations
(Titanium), garantie 5 ans.

## 128. BOM type n°2 : nœud HPC 2S + MRDIMM (indicatif)

Base du cas chiffré de la section 33.

| # | Composant | Spécification | Qté |
|---|---|---|---|
| 1 | Châssis | 2U, 2S, 500 W/CPU validés, N+1 fans | 1 |
| 2 | CPU | Intel Xeon 6980P (128 c., 500 W) | 2 |
| 3 | Refroidissement | dissipateurs 500 W (ou kit DLC) | 2 |
| 4 | RAM | 64 Go MRDIMM-8800 ECC (QVL) | 24 |
| 5 | Stockage | NVMe 7,68 To (scratch local) | 2 |
| 6 | Réseau | 200 GbE (InfiniBand ou RoCE) | 1 carte |
| 7 | Alimentations | 2000 W+ 80 PLUS Titanium, redondantes | 2 |
| 8 | Rails + CMA | kit complet | 1 |
| 9 | Garantie | 3-5 ans, sur site J+1 | 1 |

Ordre de grandeur CAPEX : **45 000-60 000 €** (indicatif). Et surtout :
**2 kW électriques + 2 kW de froid par nœud** à prévoir en salle (sections
78, 92) — le CAPEX n'est que la moitié du sujet.

## 129. Checklist d'achat finale : les règles d'or

1. **Workload d'abord** : nature (latence vs débit), puis CPU, puis RAM.
2. **Bande passante par cœur** calculée (section 30) avant tout choix dense.
3. **Licences chiffrées avant le CPU** (section 25) — le poste n°1 du TCO.
4. **Tous les canaux mémoire remplis**, barrettes identiques, QVL (45, 115).
5. **1 DPC** sauf besoin de capacité avéré (section 47).
6. **TDP et PPT** : refroidissement et électrique dimensionnés sur le max
   (sections 9, 83).
7. **Redondance N+1** ventilateurs **et** alimentations, confirmée par écrit
   sur votre config exacte (sections 59, 89).
8. **Format 2U par défaut** ; 1U seulement si la densité l'impose vraiment
   (sections 69-71).
9. **Baie mesurée** : profondeur, kW/rack, PDU 2× (A/B) monitorées
   (sections 73, 78, 89).
10. **Onduleur sur le pic + 25 %**, triphasé au-delà de 10 kW, arrêt ordonné
    testé (sections 87, 88).
11. **PUE annualisé** connu de votre salle ; objectif ≤ 1,5 (sections 81, 82).
12. **TCO 5 ans** présenté avec chaque devis, jamais le CAPEX seul
    (sections 101, 102).
13. **BIOS/BMC/microcode à jour** avant production, par vagues (104).
14. **Fin de vie négociée à l'achat** : reprise, destruction certifiée (122).
15. **Tout ce qui est « vérifié » dans ce guide a une date ; tout le reste
    est à mesurer chez vous** (wattmètre, PDU monitorées, POC).

## 130. Sources, limites et maintenance du guide

**Sources principales (recherche web du 27/09/2026)** : communiqués et
fiches techniques AMD (EPYC 9005 Turin, EPYC 9006 Venice — Advancing AI
22-23/07/2026) et Intel (Xeon 6, Xeon 6+ Clearwater Forest — Computex
01/06/2026, Diamond Rapids) ; couvertures Tom's Hardware, Phoronix (tests
MRDIMM-8800 : +24 % HPCG, +437 W mesurés), VideoCardz, HotHardware,
TechSpot, Wccftech ; analyses CXL/DDR6 (JEDEC, Samsung/SK Hynix/Micron).

**Limites** : prix = tarifs publics au lancement (remises non incluses) ;
consommations par config = estimations d'ingénierie ; chiffres constructeur
= chiffres fournisseur ; tout point marqué « indicatif » ou « à vérifier »
doit être validé par devis, fiche technique ou mesure.

**Maintenance du guide** : à relire à chaque génération (Venice T4 2026,
Diamond Rapids 2027, MRDIMM Gen2 T1 2027, DDR6 2028-2029). Les sections
marquées d'une date de vérification sont celles à revérifier en priorité.

---

*Fin du guide — 130 sections. Rédigé le 27/09/2026.*

---

## 131. EPYC 8004 « Siena » : la gamme edge 1P (complément)

Pour mémoire (hors datacenter dense, mais utile en edge) : la série 8004
sur socket **SP6**, jusqu'à 64 cœurs Zen 4c, 6 canaux DDR5, TDP contenus
(~70-225 W, indicatif). Cible : telco, edge, stockage froid, appliances.

