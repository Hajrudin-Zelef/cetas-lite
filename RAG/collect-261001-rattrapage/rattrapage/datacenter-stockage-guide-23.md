---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-23
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Samsung"]
dates: ["2026-09-27"]
keywords: ["datacenter", "arr", "capex"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [3825, 4005]
sha256: abc2c4af20e02ea50ea9fb39797496e33b78128e850338c156e3821eb59b65f2
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

## 392. Tableau : prix indicatifs (rappel, à vérifier)

| Produit | Prix constaté sept. 2026 | Statut |
|---|---|---|
| HDD Exos M 30 To | ~565-600 $ (~19 $/To) | vérifié |
| NVMe entreprise Gen5 | ~350-600 €/To | ordre de grandeur |
| HBA 9500-16i/e | ~250-450 € | ordre de grandeur |
| SSD QLC 30,72 To | ~250 €/To | ordre de grandeur |
| Switch 32×100G | ~15-30 k€ | ordre de grandeur |

Marché en allocation tendue : **devis datés exigés** (section 213).

## 393. Index des schémas ASCII du guide

- Section 3 : paysage DAS/NAS/SAN/SDS
- Section 40 : arbre de décision HBA
- Section 75 : HBA → JBOD simple
- Section 76 : cascade (daisy chain)
- Section 77 : double domaine redondant
- Section 104 : matrice RAID/EC/JBOD
- Section 105 : architecture Ceph
- Section 276 : topologie réseau Ceph
- Section 322 : pyramide des tiers

## 394. Index des tableaux de dimensionnement

- Sections 115-119 : 100 To utiles NVMe (disques, nœuds, réseau, RAM/CPU)
- Section 120 : BOM nœud NVMe chiffrée
- Section 140 : 500 To hybrides chiffrés
- Section 147 : prix hybride
- Sections 184-185 : bilans de puissance
- Sections 237-242 : 4 cas (1 Po backup, IA, PME, 2 Po S3)
- Section 243 : comparatif €/To/an
- Section 292 : gabarit TCO

## 395. Anti-sèche : les 12 chiffres à connaître par cœur

1. 100-200 PG/OSD (Ceph)
2. 1 OSD = 1 disque
3. 2,5 % db (4 % RGW)
4. ≤ 15 HDD par NVMe de db
5. 22,5 Gb/s effectifs par voie 24G
6. 1024 SAS / 32 NVMe par HBA
7. 70 % alerte / 85 % critique (remplissage)
8. 5 ans de garantie = durée d'amortissement
9. 96 % = rendement Titanium
10. 1,7 m³/h par watt (ΔT 2 °C)
11. 19 $/To = HDD 30 To (sept. 2026)
12. 3-2-1 = la règle des backups

## 396. Ce que ce guide ne couvre pas (pistes)

- Le tuning CephFS très large échelle (> 1 Po de metadata),
- Le détail des firmwares par constructeur (errata),
- La tarification cloud (S3 managé vs self-hosted),
- Le stockage Kubernetes (Rook, Longhorn, OpenEBS) — piste de guide,
- La conformité sectorielle détaillée (HDS, DORA),
- Le dimensionnement des salles (cf. guides onduleurs/Proxmox).

## 397. Historique des vérifications

| Date | Objet |
|---|---|
| 27/09/2026 | SAS 24G (Broadcom 9600/9670, fiches revendeurs) |
| 27/09/2026 | HBA 9500 (specs, conso 6,12/8,74 W) |
| 27/09/2026 | SSD : D7-PS1010, CD9P-R, CD8P, D5-P5430, PM9D3 |
| 27/09/2026 | EDSFF E1.S/E3.S, Micron 9650 Gen6, KIOXIA NX1 |
| 27/09/2026 | CXL 2.0/3.x, Marvell Structera, Samsung/SK/Micron |
| 27/09/2026 | PLC (non trouvé en produit), SSD 245 To (Samsung BM1773) |
| 27/09/2026 | HDD Exos M 30 To (~19 $/To), marché en allocation |
| 27/09/2026 | Ceph PG (docs), block.db sizing, alternatives SDS |
| 27/09/2026 | Supermicro SSG-6049P-E1CR60H, gamme EDSFF |

## 398. Comment maintenir ce guide à jour

À chaque projet : re-vérifie les sections 392 (prix), 207 (roadmaps
SSD), 213 (marché), 358 (versions Ceph). Le fond (formules, topologies,
runbooks) vieillit lentement ; les références produits vite. Note la
date de chaque re-vérification dans le tableau 397. Un guide technique
non daté est un piège ; celui-ci est daté : **27/09/2026**.

---

**FIN DU GUIDE — 398 sections numérotées.**
**Fichier : `~/workspace/user/files/datacenter_stockage_guide.md`**
**Vérifications web au 27/09/2026 — faits marqués « vérifié », « à vérifier » ou « non trouvé ».**
**Guides liés : `proxmox_guide.md`, `onduleurs_ups_guide.md`, `zabbix_guide.md`, `debian_ubuntu_guide.md`.**

---

# PARTIE Y — FICHES RÉFLEXES (IMPRIMABLES)

## 399. Fiche réflexe : câblage SAS

```
12G : SFF-8643 (interne) / SFF-8644 (externe) — 4 voies = 4,8 Go/s
24G : SFF-8654 (interne) / SFF-8674 (externe) — 4 voies = 9,6 Go/s
Débits négociés : 22,5 / 12 / 6 Gb/s par voie
Règles : 2 câbles/domaine mini, étiquettes A1/A2/B1/B2,
         vis serrées, spare de chaque type dans la baie,
         jamais de cascade triple, multipath si dual-path.
Diagnostic : storcli → vitesse négociée par phy.
```

## 400. Fiche réflexe : Ceph au quotidien

```
ceph -s ; ceph health detail          # santé (chaque matin)
ceph osd df tree                      # remplissage par OSD
Alerte : 70 % warn / 85 % crit / PG degraded = immédiat
PG/OSD cible : 100-200 (autoscaler on, pools bulk:true)
Runbook panne disque : out → HEALTH_OK → purge → remplace → recrée
Ne jamais : purger avant out, toucher aux PG en upgrade,
            redémarrer tous les MON « pour voir ».
```

## 401. Fiche réflexe : énergie

```
Nœud 60 baies HDD : ~980 W au mur → ~1 470 W à climatiser (PUE 1,5)
Nœud 12 NVMe Gen5 : ~800 W au mur
Règles : Titanium imposé, PDU A/B à 100 % chacune,
         pic spin-up = 2× pendant 10-20 s (HDD),
         NUT : arrêt séquentiel testé 2×/an (arrêt ET redémarrage),
         sondes disques → supervision (45 °C warn / 50 °C crit HDD).
TCO : kWh × PUE × prix × 5 ans — toujours, pas CAPEX seul.
```

## 402. Fiche réflexe : achats

```
Checklist bon de commande :
[ ] HBA en IT (écrit), baies UBM/E3.S, alim Titanium
[ ] Disques CMR (pas SMR), 5 ans de garantie, PLP (SSD)
[ ] 2 fournisseurs qualifiés par famille
[ ] Devis daté, délai ferme, pénalités (allocation 2026)
[ ] Spares : max(2, 5 %) par modèle + 1 HBA + câbles
[ ] Burn-in 48-72 h prévu au planning (non compressible)
Neuf 2026 : E3.S Gen5 (perf) / Exos 30 To (froid) / 9500 tri-mode.
```

## 403. Fiche réflexe : pannes (top 5)

```
1. Tous les disques d'un domaine disparaissent → HBA/câble/expander
   (lspci, storcli, domaine B en attendant)
2. Latence p99 qui grimpe sans panne → SSD lent (slow ops), réseau
   (CRC), ou horloge (chrony) — dans cet ordre
3. HEALTH_WARN Ceph → ceph health detail, tout lire, ne pas rebooter
   en masse
4. Rebuild RAID qui échoue → URE (RAID 5 sur gros disques = à bannir)
5. Après coupure : ne pas redémarrer tous les nœuds d'un coup
   (pic spin-up → onduleur qui disjoncte) — séquentiel !
```

## 404. Note de version du guide

```
Version : 1.0 — 27/09/2026
Sections : 404 numérotées (398 + 6 fiches réflexes)
Vérifications : web, 27/09/2026 (tableau section 397)
Prochaine re-vérification conseillée : prix/dispos (trimestrielle),
roadmaps SSD et versions Ceph (semestrielle).
Auteur : Leo pour Zelef — chef de service systèmes & énergies.
```

## 405. Pour aller plus loin : la feuille de route 6 mois

```
Mois 1-2 : dimensionnement (formules section 339) + 2 devis par famille
           + qualification (fiches 397)
Mois 3   : commande (délais 4-12 sem. en allocation 2026) + préparation
           salle (élec, clim, PDU A/B, NUT)
Mois 4   : réception + burn-in 48-72 h + câblage (photos, étiquettes)
Mois 5   : bench + game day n°1 (kill OSD) + runbooks + supervision
Mois 6   : production progressive + revue (TCO réel, AFR, saturation)
En continu : capacity planning trimestriel, game day trimestriel,
             revue annuelle (section 302), veille mensuelle (379).
Règle d'or : on ne compresse jamais le burn-in ni les tests —
ce sont les semaines les plus rentables du projet.
```
