---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-3
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom"]
dates: ["2026-09-27"]
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [320, 491]
sha256: f28bea87523b3fb7da1a87a1606c3b4fc5e9241f69e84979c34d64d1114da2dc
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Fiche consolidée (vérifié le 27/09/2026, fiches revendeurs) :
- **Tri-mode** : 12 Gb/s SAS, 6 Gb/s SATA, PCIe 4.0 (NVMe)
- Hôte : PCIe 4.0 x8 (16 GT/s par voie)
- Connecteurs : SFF-8654 SlimSAS (internes : 1× pour 8i, 2× pour 16i ;
  externes : SFF-8644 pour les versions « e »)
- Capacité : **1024 périphériques SAS/SATA, 32 NVMe**
- Débit : jusqu'à 13 700 Mo/s (séquentiel 256K), 3 M IOPS (4K lecture aléatoire)
- MTBF : 5 000 000 h à 40 °C ; garantie 3 ans
- Conso : **6,12 W (8e), 8,74 W (16e)** — vérifié, utile au bilan énergie
- Outils : LSI Storage Authority, storcli, HII (UEFI)
- UBM ready (SFF-TA-1005) pour baies U.3
- Prix public constaté : **à vérifier** (varie ~250-450 € selon version)

## 32. Broadcom 9600-16e / MegaRAID 9670 (24G, vérifié)

Fiche consolidée (vérifié le 27/09/2026) :
- **24 Gb/s SAS** (SAS-4), SATA 6G, PCIe 4.0 NVMe (tri-mode)
- Hôte : PCIe 4.0 x16 (9670W-16i) ou x8 (9670-24i)
- Connecteurs : **SFF-8654 SlimSAS 24G** (interne), SFF-8674 (externe)
- Débits SAS négociés : 22,5 / 12 / 6 Gb/s par voie (le 24G effectif ≈ 22,5G
  après encodage — détail vérifié sur fiche 9600)
- Devices : 240 SAS/SATA, 32 NVMe ; cache 8 Go + CacheVault CVPM05 (versions RAID)
- RAID 0/1/5/6/10/50/60 + JBOD
- **DataBolt2** : permet du 24G vers le contrôleur avec des disques/backplanes
  12G/6G existants (vérifié)
- Disponibilité : en stock chez les revendeurs consultés le 27/09/2026.

## 33. Microchip Adaptec SmartHBA / SmartIOC 2200 (24G, vérifié)

Microchip a démontré la **première interopérabilité 24G SAS de bout en bout**
avec KIOXIA (communiqué vérifié, 2021, toujours la référence publique) :
SmartROC 3200 (RAID-on-Chip PCIe Gen4 tri-mode), **SmartIOC 2200** (contrôleur
HBA), expander **SXP 24G SAS**. En 2026, Microchip reste le second fournisseur
de silicium SAS avec Broadcom. Pour un appel d'offres : **toujours mettre
Broadcom et Microchip en concurrence**.

## 34. Quand HBA vs RAID hardware : matrice

| Cas | Choix |
|---|---|
| Ceph / ZFS / SDS | HBA IT, toujours |
| Boot / OS | RAID 1 matériel ou miroir logiciel |
| Serveur applicatif simple, pas de SDS | RAID hardware 1/10 |
| Base de données sur NVMe direct | pas de carte du tout (NVMe direct CPU) |
| JBOD externe capacitaire | HBA externe 16e |

## 35. Queue depth et file d'attente

Une HBA 9500 encaisse des profondeurs de file énormes (3 M IOPS annoncés),
mais le goulot se déplace : CPU (interruptions), puis réseau. En pratique
Linux : vérifie `nr_requests`, le scheduler (`none` pour NVMe), et
l'affinité IRQ. Un NVMe Gen5 à 14 Go/s sature un cœur moderne en traitement
d'interruptions — prévois **2+ cœurs par SSD NVMe très sollicité**.

## 36. Multipath et MPIO

En SAS double-domaine (dual-path) vers un JBOD, chaque disque apparaît deux fois :
il faut **multipath** (Linux `multipathd` / `mpath`) pour n'en voir qu'un.
Sans multipath : corruption assurée dès qu'un chemin flappe. Règle : dual-path
= multipath configuré **et testé** (débranche un câble en burn-in, vérifie).

## 37. Firmware : la discipline

- Note la version firmware HBA **et** expander au déballage.
- Ne mets à jour qu'avec le package du **constructeur du serveur** (pas le
  générique Broadcom) sauf consigne contraire — les OEM verrouillent parfois.
- Après flash : revérifie le mode IT/IR (certains flashs réinitialisent).
- Garde l'ancien firmware en rollback 48 h après toute MaJ en production.

## 38. Outils : storcli, LSA, sas3ircu

| Outil | Usage |
|---|---|
| `storcli` | CLI : état, firmware, locate, rebuild (vérifié : livré avec les 9500/9600) |
| LSI Storage Authority | GUI web de supervision |
| HII (UEFI) | config pré-boot |
| `sas3flash` / `storcli` | flash firmware |

Automatise : un script qui collecte `storcli /c0 show all` par nœud et l'envoie
à ta supervision (Zabbix — voir ton `zabbix_guide.md`). Une HBA qui disparaît
du PCIe = nœud à isoler immédiatement.

## 39. NVMe direct vs Tri-Mode

| Approche | Câblage | Cas |
|---|---|---|
| NVMe direct CPU (backplane PCIe) | MCIO / SlimSAS x4 par baie | perf max, pas de carte |
| Tri-Mode HBA | SFF-8654 vers backplane UBM | mixte SAS/SATA/NVMe, souplesse |

Le NVMe direct supprime la carte (moins de latence, moins de watts) mais fige
le châssis en « tout NVMe ». Le tri-mode garde la flexibilité. Pour Ceph full
NVMe neuf : **NVMe direct si le châssis le permet**, tri-mode sinon.

## 40. Arbre de décision HBA

```
Disques = HDD/SATA capacitaires ? ──oui──▶ HBA 12G (9500-16e / 9400 recond.)
          │
          non
          ▼
Disques = NVMe ? ──oui──▶ Backplane NVMe direct ? ──oui──▶ pas de HBA
          │                        │
          │                       non──▶ HBA tri-mode 9500-16i
          non
          ▼
Mixte / évolutif ? ──oui──▶ HBA tri-mode 9500/9600
```

## 41. eHBA : le mode hybride (9600)

Les 9600 existent en version **eHBA** : accélération RAID pour quelques volumes
+ présentation JBOD/HBA pour le reste. Utile en transition (boot RAID 1 + data
JBOD sur la même carte). À manier avec prudence : un mauvais profil et tes
disques Ceph se retrouvent derrière du RAID sans que tu le voies.

## 42. Pièges HBA (rappel, détail en section 217+)

1. Firmware IR livré au lieu d'IT. 2. Câble SFF-8643 branché sur port SFF-8654
sans adaptateur (ça ne rentre pas — forcer = broches pliées). 3. HBA x8 dans un
slot x4 électrique : la moitié du débit. 4. Pas de flux d'air sur la HBA :
les 9500/9600 chauffent, prévoir le déflecteur d'air du châssis.

---

# PARTIE D — SAS : GÉNÉRATIONS, CÂBLAGE, LIMITES

## 43. Générations SAS : 6G et 12G

| Génération | Norme | Débit par voie | Statut 2026 |
|---|---|---|---|
| SAS-2 | 6 Gb/s | 600 Mo/s | legacy, SATA-like |
| SAS-3 | 12 Gb/s | 1 200 Mo/s | **standard capacitaire actuel** |
| SAS-4 | 24 Gb/s (22,5 effectifs) | ~2 250 Mo/s | **disponible (vérifié)** |

Le 12G reste le standard des baies HDD/SSD SAS neuves en 2026 : mature,
pas cher, suffisant pour des HDD (250-300 Mo/s unitaires).

## 44. SAS-4 / 24G : état vérifié au 27/09/2026

Le 24G SAS **existe et se vend** : contrôleurs Broadcom série 9600/9670
(SFF-8654/SFF-8674) en stock chez les revendeurs, interopérabilité 24G
démontrée Microchip+KIOXIA, FEC (correction d'erreur) intégrée à la norme.
Mais : les **disques 24G** restent rares et chers, et le gain réel sur HDD
est nul (un HDD ne sature pas le 12G). Le 24G a du sens pour : SSD SAS
rapides, agrégation d'expanders, et préparation de l'avenir. Pour du HDD
pur en 2026 : **reste en 12G**.

## 45. Câblage SFF-8087 / 8088 (legacy)

Mini-SAS interne (8087) / externe (8088), génération 6G. Encore présents sur
le matériel d'occasion. Règle : ne pas mélanger avec du 12G+ sans vérifier
la qualité du câble — un vieux 8087 peut négocier en 6G seulement et
brider une chaîne 12G.

## 46. SFF-8643 / 8644 : Mini SAS HD (12G)

Le standard 12G : **SFF-8643** (interne, 4 voies) et **SFF-8644** (externe).
Un câble 8644 = 4 voies × 12G = 48 Gb/s = 4,8 Go/s. C'est ce qui relie une
HBA 9500-16e à un JBOD. Longueur cuivre passive : ~2-4 m selon qualité
(au-delà : câbles actifs ou optiques — à vérifier par référence).

## 47. SFF-8654 SlimSAS 24G (vérifié)

Le connecteur 24G : **SFF-8654** « SlimSAS », 8 voies (x8) en interne.
Vérifié sur les fiches 9500-16i (2× SFF-8654) et 9670-24i (3× SFF-8654).
Attention : le 8654 existe en versions 4 voies et 8 voies, et en brochage
SAS vs PCIe — **le câble doit matcher le brochage du backplane** (SFF-9402).
Un mauvais câble SlimSAS = lien qui ne monte pas ou négocie en PCIe
au lieu de SAS.

## 48. SFF-8674 externe (vérifié)

