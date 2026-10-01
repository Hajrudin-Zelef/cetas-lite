---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-1
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Samsung"]
dates: ["2024-10-10", "2026-09-27"]
keywords: ["amd", "benchmarks", "inference", "intel"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [1, 147]
sha256: c2d276eb915bc2ddab64c150499964b6ef5c4c04b31e811f0dc468ccf09bf57b
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

## Guide technique — présent et futur vérifié

**Rédigé le 27/09/2026 — pour Zelef, chef de service systèmes & énergies.**
**Angle : chiffres, BOMs, dimensionnements, et lien permanent avec l'énergie**
**(consommation, refroidissement, alimentation/onduleur).**

> Convention de ce guide :
> - « vérifié le 27/09/2026 » = information confirmée par recherche web ce jour-là.
> - « indicatif » = ordre de grandeur d'ingénierie, à valider par devis/fiche technique.
> - « non trouvé au 27/09/2026 » = cherché, introuvable ce jour-là.
> - Les prix CPU cités sont les tarifs publics constructeur au lancement (1KU, USD) ;
>   les prix réellement payés dépendent du volume, de l'OEM et de la négociation.

### Sommaire des parties

- Partie A — AMD EPYC (sections 1 à 10)
- Partie B — Intel Xeon 6 (sections 11 à 22)
- Partie C — 1, 2, 4 sockets & NUMA (sections 23 à 33)
- Partie D — RAM serveur I : formats et technologies (sections 34 à 45)
- Partie E — RAM serveur II : dimensionnement chiffré (sections 46 à 56)
- Partie F — Ventilation serveurs (sections 57 à 68)
- Partie G — Formats 1U/2U/4U, châssis, rails (sections 69 à 80)
- Partie H — Lien énergie : PUE, conso, onduleur, PDU (sections 81 à 92)
- Partie I — À venir : roadmaps vérifiées (sections 93 à 104)
- Partie J — 18 pièges terrain (sections 105 à 122)
- Partie K — Glossaire, quiz, BOMs, checklist (sections 123 à 130)

---

## 1. Objectif et méthode de lecture

Ce guide répond à une question simple : **comment choisir, dimensionner et
alimenter un serveur moderne sans se tromper d'un facteur 2 sur la facture
électrique ni d'un facteur 10 sur les licences.** Il couvre le CPU (AMD EPYC,
Intel Xeon 6), la RAM serveur (DDR5, RDIMM/LRDIMM/MRDIMM, CXL), la ventilation
et les châssis, avec un fil rouge permanent : l'énergie.

Méthode de lecture conseillée :
- Vous achetez un serveur cette année → lisez les parties A, B, C, D, E puis H.
- Vous dimensionnez une baie ou une salle → parties F, G, H en priorité.
- Vous préparez un budget 2027-2028 → partie I, puis H (TCO).
- Vous êtes pressé → sections 10, 19, 30, 48, 85, 87, 129 (les règles d'or).

Tout ce qui est présenté comme un fait produit vérifiable porte sa source ou
sa date de vérification. Tout le reste est explicitement marqué comme
estimation. Un guide technique qui invente des chiffres est pire que pas de
guide du tout.

## 2. Sources et honnêteté des chiffres

Les faits « vérifié le 27/09/2026 » de ce guide proviennent de la recherche web
effectuée ce jour : communiqués AMD/Intel, couvertures de lancement (Tom's
Hardware, Phoronix, VideoCardz, HotHardware, TechSpot, Wccftech), tableaux de
références constructeur et analyses indépendantes.

Limites assumées :
- Les benchmarks constructeur (AMD vs Intel et inversement) sont des chiffres
  **fournisseur**, obtenus sur des configurations choisies. Ils donnent une
  direction, pas une vérité. Les mesures indépendantes (Phoronix, ServeTheHome)
  restent la référence.
- Les prix publics CPU sont ceux du lancement ; les remises OEM/volume les
  font varier de 20 à 50 %. Pour les DIMM, alimentations et châssis, seuls des
  ordres de grandeur « indicatifs » sont donnés : le marché mémoire fluctue
  fortement (tensions d'approvisionnement signalées par Samsung et Micron pour
  2026-2027 — vérifié le 27/09/2026).
- Les consommations « par config » de la partie H sont des **estimations
  d'ingénierie** (TDP × facteurs d'utilisation + overhead), pas des mesures au
  wattmètre. Elles servent à dimensionner, pas à facturer.

## 3. Panorama des générations AMD EPYC

AMD a pris le leadership du serveur x86 en deux vagues : d'abord le rapport
cœurs/euro avec Rome/Milan, puis la densité et l'efficacité avec Genoa/Bergamo
et Turin. Le tableau ci-dessous fixe les repères (vérifié le 27/09/2026 pour
les générations 9004/9005/9006 ; générations antérieures = historique établi).

| Génération | Série | Nom de code | Cœurs max/socket | Socket | Mémoire | PCIe |
|---|---|---|---|---|---|---|
| 1re (2017) | 7001 | Naples | 32 | SP3 | DDR4-2666, 8 canaux | Gen3 |
| 2e (2019) | 7002 | Rome | 64 | SP3 | DDR4-3200, 8 canaux | Gen4 |
| 3e (2021) | 7003 | Milan | 64 | SP3 | DDR4-3200, 8 canaux | Gen4 |
| 4e (2022) | 9004 | Genoa | 96 | SP5 | DDR5-4800, 12 canaux | Gen5 |
| 4e dense (2023) | 9004 | Bergamo (Zen 4c) | 128 | SP5 | DDR5-4800, 12 canaux | Gen5 |
| 5e (10/2024) | 9005 | Turin (Zen 5/5c) | 192 | SP5 | DDR5-6400, 12 canaux | Gen5 |
| 6e (annoncée 07/2026) | 9006 | Venice (Zen 6/6c) | 256 | SP7 | DDR5-8000, 16 canaux | Gen6 |

Points à retenir :
- **SP5 est le socket de la génération 9004/9005** : un serveur SP5 Genoa
  accepte en principe un CPU Turin après mise à jour BIOS/microcode (drop-in
  upgrade — vérifier la matrice de compatibilité de l'OEM, section 104).
- Turin double la bande passante mémoire par rapport à Genoa (DDR5-6400 contre
  DDR5-4800) à nombre de canaux égal (12).
- La vraie rupture plateforme arrive avec Venice : **nouveau socket SP7**,
  16 canaux, PCIe Gen6 (voir section 93).

## 4. EPYC 9005 « Turin » : positionnement

Turin (lancé le 10/10/2024, vérifié le 27/09/2026) est la 5e génération EPYC.
Deux variantes d'architecture coexistent dans la même gamme commerciale :

- **Turin « scale-up »** : cœurs Zen 5 classiques (jusqu'à 16 CCD), gravure
  4 nm, jusqu'à 128 cœurs / 256 threads, 512 Mo de L3 max. Cible : performance
  par cœur, bases de données, virtualisation exigeante.
- **Turin « scale-out »** : cœurs Zen 5c denses (jusqu'à 12 CCD), gravure
  3 nm, jusqu'à **192 cœurs / 384 threads**, 384 Mo de L3 max. Cible : cloud,
  conteneurs, scale-out, inference IA.

Chiffres clés constructeur (à prendre comme ordre de grandeur) :
- IPC : +17 % en entreprise/cloud, jusqu'à +37 % en HPC/IA vs génération
  précédente (chiffres AMD — vérifié le 27/09/2026).
- AVX-512 avec chemin de données 512 bits complet (contre 2×256 bits sur
  Zen 4) : gros gain sur HPC et IA, voir section 22.
- Contrôleur DDR5 avec ECC dynamique « post-package repair », chiffrement des
  liens PCIe, **CXL 2.0** (voir section 56).
- TDP max : 500 W (contre 400 W sur Genoa) — le refroidissement suit, voir
  partie F.

## 5. Zen 5 vs Zen 5c : bien choisir sa variante

| Critère | Zen 5 (classique) | Zen 5c (dense) |
|---|---|---|
| Objectif | Perf/cœur maximale | Débit/cœur, densité |
| Cœurs max/socket | 128 | 192 |
| Fréquence boost max | 5,0 GHz (réf. F) | 3,7 GHz |
| L3 max | 512 Mo | 384 Mo |
| Finesse de gravure CCD | 4 nm | 3 nm |
| Workloads types | DB, ERP, virtualisation dense, HPC | Cloud, conteneurs, web, inference |
| Exemples de réf. | 9755, 9575F, 9175F | 9965, 9845, 9825, 9745 |

Règle pratique : **si votre workload est sensible à la latence ou à la
fréquence mono-thread (bases de données transactionnelles, certains ERP),
prenez du Zen 5 classique voire une référence « F ». Si votre workload se
parallélise bien (VM, conteneurs, serveurs web, encodage), le Zen 5c donne
plus de débit par euro et par watt.**

Piège classique : comparer deux CPU uniquement au nombre de cœurs. Un 64
cœurs Zen 5 à 5,0 GHz bat un 96 cœurs Zen 5c à 2,3 GHz sur tout workload
mal parallélisé. Le dimensionnement commence par la nature du workload,
pas par le catalogue.

## 6. Tableau des références EPYC 9005 « Turin »

Références principales, prix publics 1KU USD au lancement, TDP par défaut
(vérifié le 27/09/2026 via couvertures de lancement ; base/boost en GHz).

