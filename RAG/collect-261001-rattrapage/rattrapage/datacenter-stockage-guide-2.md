---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-2
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom"]
dates: ["2026-09-27"]
keywords: ["arr", "compute"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [155, 319]
sha256: 6d82755300c84208cb3c9bc1ccccfe3ee096dd05bc055feb91daccf33d8f417f
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Le « 4U 60 baies » top-load (chargement par le dessus) : 60 disques 3,5" SAS/SATA,
2 backplanes expander, 2+2 alimentations. Exemple vérifié : **Supermicro
SSG-6049P-E1CR60H** — 4U, 60 baies 3,5" SAS3/SATA3 hot-swap + 2 baies 2,5" arrière,
alimentations 2000 W redondantes Titanium (vérifié le 27/09/2026, fiches revendeurs).
Poids à plein : **~50-60 kg** — prévoir 2 personnes + rails lourds.

## 15. Topologie 4U 90+ baies

Des châssis 4U à 90-102 baies 3,5" existent chez les hyperscalers et certains OEM
(architectures « top-load double rangée »). Référence précise **non trouvée au
27/09/2026** dans les catalogues publics généralistes consultés → **à vérifier**
auprès des intégrateurs. Retiens : au-delà de 60 baies, le goulot devient
l'expander et la puissance (60 HDD × 10 W = 600 W rien que pour les disques).

## 16. Critère de choix n°1 : le backplane

| Type de backplane | Débit par baie | Usage |
|---|---|---|
| Direct (pas d'expander) | plein débit HBA | NVMe, perf maximale |
| Expander SAS 12G simple | partagé | HDD/SSD SATA, capacitaire |
| Tri-mode (SAS/SATA/NVMe) | négocié par baie | mixte, évolutif |
| EDSFF E3.S direct PCIe | Gen5 x4 par baie | dense NVMe 2026 |

Un backplane expander 12G x4 vers le HBA = 4,8 Go/s pour 24 disques : OK pour HDD,
**catastrophique pour 24 NVMe Gen4** (qui demandent ~7 Go/s chacun). Le backplane
se choisit **après** le workload, jamais avant.

## 17. Critère : profondeur, rails, poids

Un 4U 60 baies fait ~70-80 cm de profondeur (vérifié : 767 mm pour le SSG-6049P).
Vérifie : profondeur utile du rack (1000 mm mini conseillé), rails supportant
60+ kg, dégagement arrière pour câbles + PDU. Un châssis plein qui ne sort pas
de ses rails = maintenance impossible sans arrêt.

## 18. Alimentations : redondance et rendement Titanium

| Niveau 80 PLUS | Rendement à 50 % de charge |
|---|---|
| Gold | 92 % |
| Platinum | 94 % |
| Titanium | 96 % |

Vérifié : le SSG-6049P-E1CR60H est livré en 2000 W redondants **Titanium**
(27/09/2026). Sur un nœud 60 baies qui tire ~900 W en continu, passer de Gold
à Titanium économise ~35-40 W en permanence, soit ~350 kWh/an — à multiplier
par le nombre de nœuds et le PUE. Pour ton métier : **impose Titanium sur tout
nœud de stockage neuf**.

## 19. Refroidissement : le flux d'air décide

Les disques 3,5" empilés en top-load créent des zones mortes. Règles :
ventilateurs hot-swap redondants (N+1), sondes par zone, et **ne jamais faire
fonctionner un 60 baies avec des baies vides sans obturateurs** : l'air prend
le chemin le plus facile et les disques restants chauffent. Température cible
disques : < 40 °C en régime établi (au-delà, l'AFR grimpe).

## 20. Référence vérifiée : Supermicro SSG-6049P-E1CR60H

Fiche consolidée (vérifié le 27/09/2026, plusieurs revendeurs) :
4U, 60 baies 3,5" hot-swap SAS3/SATA3 + 2 baies 2,5" arrière, double Xeon Scalable,
24 slots DDR4 (jusqu'à 6 To), 3 slots PCIe, alimentations 2000 W redondantes
Titanium, expander SAS3 12 Gb/s. Statut catalogue : **fin de vie (EOL)** chez
certains revendeurs — pour du neuf 2026, vise la génération suivante (W790/X13)
ou du reconditionné garanti. Prix : **à vérifier** (sur devis, configuration dépendante).

## 21. Références vérifiées : Supermicro EDSFF all-flash

Communiqué Supermicro (vérifié le 27/09/2026) :
- **SSG-121E-NE316R** : 1U, 16 baies E3.S
- **SSG-221E-NE324R** : 2U, 32 baies E3.S
- **SSG-121E-NES24R** : 1U, 24 baies E1.S
- Capacité : 0,5 Po en 1U16, **1 Po en 2U32** (avec SSD ~30 To)
- PCIe Gen5, architecture NUMA symétrique, 2 slots PCIe Gen5 x16 + 2 AIOM.
C'est la trajectoire « pétaoctet par U » : le format à viser pour tout projet
all-flash neuf à partir de 2026.

## 22. Références marché : Dell, HPE, Lenovo

Dell PowerEdge (série R, ex. R760xd), HPE ProLiant (DL380), Lenovo ThinkSystem
(SR650) : les trois généralistes proposent des 2U 12-24 baies et des 4U/5U
capacitaires. Modèles exacts et options backplane **à vérifier** au catalogue
du trimestre (les gammes tournent vite). En pratique : compare toujours à
iso-backplane et iso-alimentation, pas à iso-prix catalogue.

## 23. BOM exemple : nœud 60 baies HDD (capacitaire froid)

| Ligne | Qté | Ordre de grandeur |
|---|---|---|
| Châssis 4U 60 baies + 2×2000 W Titanium | 1 | à vérifier (devis) |
| Carte mère double socket + 2 CPU 16c | 1 | à vérifier |
| RAM 128 Go ECC DDR4/DDR5 | 1 lot | à vérifier |
| HBA 12G 16 ports externes (ex. 9500-16e) | 1 | ~300-500 € (ordre de grandeur public) |
| HDD 20-30 To CMR 7200 tr/min | 60 | ~15-20 €/To → 18 000-36 000 € |
| SSD boot 2,5" (miroir) | 2 | à vérifier |
| Rails lourds + obturateurs | 1 lot | à vérifier |

Capacité brute : 60 × 20 To = 1,2 Po (1,09 Pio). Avec EC 8+3 : ~0,87 Po utiles.

## 24. BOM exemple : nœud NVMe E3.S (performant)

| Ligne | Qté | Ordre de grandeur |
|---|---|---|
| Châssis 2U 32 baies E3.S Gen5 | 1 | à vérifier |
| 2 CPU + 512 Go DDR5 | 1 lot | à vérifier |
| 2× NIC 100 GbE | 2 | à vérifier |
| SSD E3.S 7,68 To Gen5 (ex. CD9P-R, D7-PS1010) | 32 | ~350-600 €/To → 86 000-147 000 € |
| 2×1600 W Titanium | inclus | — |

Capacité brute : 32 × 7,68 To = 245,8 To. Réplica 3 → ~82 To utiles par nœud.

## 25. Calcul d'air simplifié (règle terrain)

1 W à évacuer ≈ 1,7 m³/h d'air pour ΔT = 2 °C (ordre de grandeur, à vérifier
selon hygrométrie). Un nœud 60 baies à 900 W demande donc ~1 500 m³/h.
Vérifie que la rangée et les CRAC suivent **avant** d'empiler 10 nœuds
(9 kW, 15 000 m³/h) dans une rangée prévue pour du compute léger.

## 26. Montage rack : sécurité

60 baies pleines ≈ 55-65 kg. Règles : rails **dédiés lourds** (pas les rails
fins du 1U), 2 personnes minimum, chariot élévateur de rack si disponible,
disques installés **après** mise en rack quand c'est possible. Un 4U qui bascule
d'un rack, c'est un accident du travail + 60 disques morts.

---

# PARTIE C — CARTES HBA

## 27. Principe d'une HBA

Une HBA (Host Bus Adapter) expose les disques **tels quels** au système
d'exploitation : pas de RAID, pas de cache, pas d'abstraction. C'est le mode
« dumb pipe » intelligent : le logiciel (ZFS, Ceph, mdadm) voit chaque disque,
son SMART, ses erreurs. Pour tout SDS moderne : **HBA obligatoire, RAID matériel
interdit sur les disques de données**.

## 28. IT mode vs IR mode

| Mode | Nom Broadcom | Comportement |
|---|---|---|
| IT (Initiator-Target) | HBA / « eHBA » | JBOD pur, chaque disque visible |
| IR (Integrated RAID) | MegaRAID | RAID 0/1/5/6/10/50/60 matériel |

Le même silicium peut souvent basculer d'un firmware à l'autre (ex. 9400/9500 :
firmware HBA vs MegaRAID). **Vérifie le firmware livré** : un « 9500-16i » peut
arriver en IR si mal commandé. Pour Ceph/ZFS : IT ou rien.

## 29. Pourquoi IT mode pour ZFS et Ceph

ZFS et Ceph gèrent eux-mêmes redondance, checksums, scrub et rebuild. Un RAID
matériel entre les deux :
- masque les erreurs disque (le contrôleur « corrige » en silence),
- casse l'alignement des rebuilds (rebuild RAID 6 de 60×20 To = semaines),
- ajoute un point de panne (cache, batterie/supercondensateur).
Exception : les **disques système** (boot) peuvent rester en RAID 1 matériel.

## 30. Broadcom HBA 9400-16i (génération 12G)

La 9400-16i (SAS3408/SAS3416, PCIe 3.0, 12G SAS) est la HBA 12G la plus répandue
des générations précédentes. En 2026 elle reste pertinente pour du HDD/SATA
capacitaire d'occasion/reconditionné. Prix public : **à vérifier** (marché gris
actif). Pour du neuf NVMe/Gen4 : vise la 9500.

## 31. Broadcom HBA 9500-8i / 9500-16i / 9500-8e / 9500-16e (vérifié)

