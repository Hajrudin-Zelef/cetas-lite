---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-6
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: ["2026-09-27"]
keywords: ["attention", "datacenter", "dram", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [774, 931]
sha256: 3a5cacbf3849851d4ada6801052691cea0a588b03c58be8091840f745ba3721c
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

Exemple : 12 canaux × 256 Go LRDIMM = **3 To par socket** en 1 DPC.
C'est le format des bases in-memory (SAP HANA) et des grosses instances
Java. En dessous de 1 To par socket, le LRDIMM n'apporte rien d'autre
qu'une facture plus élevée.

## 37. 3DS : la densité extrême (Three-Dimensional Stacking)

Le 3DS empile les puces DRAM verticalement (TSV — vias traversants) dans
une barrette RDIMM ou LRDIMM : jusqu'à **256 Go en RDIMM 3DS**, davantage
en LRDIMM.

| Avantage | Inconvénient |
|---|---|
| Densité record par slot | prix très élevé (indicatif : 2-4× le RDIMM standard) |
| Pas de changement de carte mère | consommation et chaleur concentrées |
| Idéal pour maximiser To/socket | disponibilité parfois limitée |

Usage réel : serveurs de virtualisation très denses (beaucoup de VM, peu
de débit par VM) et appliances. Pour 95 % des projets, des RDIMM 64/96 Go
standard font le travail pour 3× moins cher au Go.

## 38. MRDIMM : rappel et positionnement (voir section 17)

Pour mémoire dans la partie RAM : le MRDIMM est électriquement un RDIMM
DDR5 avec multiplexeur de rangs. Points d'attention spécifiques :

- **Compatibilité** : un slot MRDIMM accepte des RDIMM classiques (le
  contrôleur s'adapte), mais l'inverse n'est pas vrai en pratique — il faut
  un contrôleur qui parle MRDIMM (Xeon 6900P aujourd'hui, Venice demain).
- **Population** : 1 DPC recommandé pour atteindre 8800 MT/s ; 2 DPC fait
  chuter la fréquence (section 47).
- **Surcoût énergétique** : +437 W mesurés sur un serveur 1 CPU en charge
  mémoire intensive (Phoronix — vérifié le 27/09/2026). À intégrer au
  dimensionnement électrique (partie H), pas seulement au devis.

## 39. Tableau comparatif des formats de barrettes

| Format | Principe | Capacité typique | Fréquence max | Prix relatif (indicatif) | Usage |
|---|---|---|---|---|---|
| UDIMM ECC | non bufferisé | 16-64 Go | 5600 | 1 | stations de travail, petits serveurs |
| RDIMM | registre | 16-128 Go | 6400 | 1,2-1,5 | standard serveurs |
| RDIMM 3DS | DRAM empilées | 128-256 Go | 6400 | 2-4 | densité sans LRDIMM |
| LRDIMM | buffer complet | 128-256 Go | 5600-6400 | 1,5-2 | > 1,5 To/socket |
| MRDIMM G1 | rangs multiplexés | 32-256 Go | 8800 | 1,5-2 | HPC/IA/DB (Intel) |
| MRDIMM G2 | rangs multiplexés | à venir | 12800 | à venir | Venice / Xeon futurs |

Note UDIMM ECC : existe en DDR5 pour les serveurs d'entrée de gamme 1P
peu denses. Dès que vous dépassez 4-6 barrettes, passez en RDIMM.

## 40. DDR5 : fréquences réellement disponibles (vérifié le 27/09/2026)

Le standard JEDEC DDR5 monte à **6400 MT/s** ; au-delà, c'est du MRDIMM
ou de l'overclocking (hors serveur sérieux). Fréquences serveur réelles :

| Fréquence | Statut 2026 | Usage |
|---|---|---|
| DDR5-4800 | en fin de vie (générations Genoa) | renouvellement ancien parc |
| DDR5-5200 | disponible | entrée de gamme |
| DDR5-5600 | **standard courant** | la plupart des serveurs neufs |
| DDR5-6000 | disponible | bon compromis prix/perf |
| DDR5-6400 | max JEDEC, max Turin/6900P | HPC, DB exigeantes |
| DDR5-8000 | natif Xeon 6+ / Venice (SP7) | nouvelle génération |
| MRDIMM-8800 | production (Intel) | HPC/IA |
| MRDIMM-12800 | attendu T1 2027 | Venice, futurs Xeon |

**Règle d'achat** : ne payez la fréquence max que si votre workload est
sensible à la bande passante (section 30). En virtualisation généraliste,
du DDR5-5600 bien populé bat du DDR5-6400 mal populé (canaux vides =
bande passante perdue, section 46).

## 41. ECC : pourquoi c'est obligatoire en serveur

ECC (Error Correcting Code) : 8 bits supplémentaires par 64 bits de données
(72 bits au total par transfert), capables de **corriger toute erreur 1 bit
et de détecter les erreurs 2 bits** (SECDED).

Pourquoi c'est non négociable :
- À l'échelle d'un To de RAM, le taux d'erreurs « soft » (rayons cosmiques,
  bruit thermique) produit des **erreurs par mois**, pas par décennie.
- Sans ECC : 1 bit qui bascule = donnée corrompue silencieusement, crash
  kernel, ou pire : corruption de base de données **qui se réplique dans
  les sauvegardes**.
- Avec ECC : l'erreur 1 bit est corrigée à la volée et **comptée** ; quand
  les erreurs corrigées se multiplient sur une barrette, le BMC alerte
  **avant** la panne (maintenance prédictive, section 55).

Coût : ~12,5 % de puces en plus, quelques % de prix. Le rapport
coût/bénéfice est le plus élevé de tout le serveur. **Aucun serveur de
production ne tourne sans ECC. Point.**

## 42. On-die ECC vs side-band ECC : ne pas confondre

La DDR5 a introduit une confusion marketing qu'il faut dissiper :

| Type | Où | Corrige | Suffit en serveur ? |
|---|---|---|---|
| On-die ECC | dans chaque puce DRAM (DDR5) | erreurs internes à la puce | **non** |
| Side-band ECC | 2 puces supplémentaires par rang (barrette ECC) | erreurs sur le bus + puces | **oui** |

- L'on-die ECC existe sur **toutes** les DDR5, y compris non-ECC : il ne
  protège que l'intérieur des puces, pas le bus mémoire.
- Le vrai ECC serveur = **side-band** (barrettes « ECC » Registered).
  C'est lui qui remonte les erreurs au système d'exploitation et au BMC.
- Une barrette « DDR5 » vendue sans mention ECC n'a que l'on-die ECC :
  **insuffisant pour un serveur de production.**

## 43. Rank, bank, bank group : le vocabulaire

- **Rank (rang)** : ensemble de puces DRAM lues/écrites simultanément pour
  fournir 64 bits (+8 ECC). 1R = un rang, 2R = deux rangs par barrette.
  Le contrôleur n'active qu'un rang à la fois.
- **Bank / bank group** : subdivisions internes des puces. La DDR5 a
  32 banques en 8 groupes (contre 16 banques en DDR4) : plus de parallélisme
  interne, d'où une partie des gains DDR5 à fréquence égale.
- **Pourquoi ça compte** : un 2R offre un peu plus de bande passante
  effective qu'un 1R (entrelacement des rangs) et permet des capacités
  supérieures ; mais 2 rangs × 2 DPC = 4 rangs par canal, ce qui peut
  forcer une baisse de fréquence (section 47).

En pratique : à capacité égale, préférez le 2R pour la performance, sauf
si la QVL de l'OEM recommande le contraire pour votre fréquence cible.

## 44. Capacités réelles des barrettes en 2026 (indicatif)

| Capacité | Format | Disponibilité (indicatif) |
|---|---|---|
| 16 Go | RDIMM 1R | courante, entrée de gamme |
| 32 Go | RDIMM 1R/2R | la plus vendue avec 64 Go |
| 48 Go | RDIMM 2R | courante (puces 24 Gb) |
| 64 Go | RDIMM 2R | **standard datacenter** |
| 96 Go | RDIMM 2R | courante (puces 24 Gb) |
| 128 Go | RDIMM 2R / 3DS | disponible, prix élevé |
| 256 Go | LRDIMM / 3DS | in-memory, prix très élevé |

Le « sweet spot » prix/Go en 2026 se situe sur **64 et 96 Go**
(indicatif — le marché mémoire est tendu, voir section 2). Les 128 Go+
ne se justifient que si vous manquez de slots (12 canaux × 128 Go =
1,5 To/socket en RDIMM standard, ce qui couvre déjà 90 % des besoins).

## 45. Règles de population des slots : principes

Les 12 canaux d'un socket SP5 = 12 slots en 1 DPC (24 slots en 2 DPC sur
les cartes 2 DPC). Règles universelles :

1. **Remplissez tous les canaux** : 12 barrettes identiques > 8 barrettes
   plus grosses. Chaque canal vide = 1/12e de bande passante en moins.
2. **Symétrie stricte** : mêmes capacité, rang et fréquence sur tous les
   canaux d'un socket ; idéalement la même référence partout.
3. **1 DPC avant 2 DPC** : une barrette par canal donne la fréquence max ;
   la 2e barrette fait souvent baisser la fréquence (section 47).
4. **Symétrie inter-sockets** : en 2S, les deux sockets doivent avoir la
   même configuration mémoire (capacité et population), sinon NUMA
   déséquilibré.
5. **Suivez le manuel de la carte mère** : l'ordre de remplissage des
   slots (A1, B1…) est normé par l'OEM. Un slot « sauté » = canal perdu.

Schéma : 12 canaux, 1 DPC, population correcte :

