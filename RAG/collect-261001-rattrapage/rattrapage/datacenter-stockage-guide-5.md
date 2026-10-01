---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-5
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia", "Samsung"]
dates: ["2026-09-27"]
keywords: ["datacenter", "arr", "hbm", "nand", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [669, 843]
sha256: 49f07df053eba3a95ec81ad4464af56dadd66a014b59aa09ecfde5523930dde1
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Deux faits vérifiés le 27/09/2026 :
- **KIOXIA NX1** : premier E1.S avec **refroidissement liquide direct**,
  +38 % d'écriture séquentielle, conforme OCP Datacenter NVMe SSD 2.6, FDP.
- **Solidigm D7-PS1010 E1.S** : version **cold plate** (plaque froide)
  co-développée avec NVIDIA, états de puissance 5-25 W.
Message : à partir de Gen5, le SSD est un composant **thermique** autant
qu'électronique. Prévois le flux d'air (ou la boucle d'eau) **avant** la
densité.

## 70. Gammes SSD datacenter 2026 (vérifié)

| Modèle | Interface | Format | Capacités | Positionnement |
|---|---|---|---|---|
| Solidigm D7-PS1010 | Gen5 x4 | U.2 / E3.S / E1.S | 1,92-15,36 To | performance (vérifié) |
| KIOXIA CD9P-R | Gen5 x4 | E3.S | 1,92-30,72 To | lecture intensive, 1 DWPD (vérifié) |
| KIOXIA CD8P | Gen5 x4 | U.2 / E3.S | 1,6-12,8 To | mixte, 3 DWPD (vérifié) |
| Samsung PM9D3 | Gen5 x4 | E3.S | jusqu'à 128 To | datacenter (vérifié) |
| Solidigm D5-P5430 | Gen4 x4 | U.2 / E3.S | jusqu'à 30,72 To | QLC capacitaire, 0,58 DWPD (vérifié) |
| Micron 9650 | **Gen6** | E1.S / E3.S | à vérifier | premier Gen6 (vérifié) |

Prix : ~350-600 €/To constatés pour du NVMe entreprise PCIe 5.0
(ordre de grandeur public, allocation tendue — **à vérifier** sur devis).

## 71. Le marché NAND en septembre 2026 (vérifié)

Fusion Worldwide Greensheet (sept. 2026, vérifié) : **SSD entreprise
PCIe 5.0 / haute capacité = allocation contrainte, prix en hausse**,
comme HBM et DDR5 haute densité. Le NAND spot grand public reste plus
souple. Traduction achats : commandes fermes tôt, stocks tampons sur les
références critiques, et second sourcing (2 fournisseurs qualifiés).

## 72. Tableau comparatif des formats NVMe

| Format | Idéal pour | Densité | Thermique | Maturité 2026 |
|---|---|---|---|---|
| U.2 15 mm | existant, Gen4 | 24/2U | correcte | maximale |
| U.3 | mixte tri-mode | 24/2U | correcte | haute |
| E1.S | 1U dense, Gen5/6 | 24/1U | bonne (+liquide) | croissante |
| E3.S | 2U all-flash neuf | 32/2U | très bonne | **standard neuf** |
| E1.L/E3.L | capacité extrême | variable | à étudier | émergente |

---

# PARTIE F — JBOD : TOPOLOGIES, CÂBLAGE, LIMITES

## 73. Principe du JBOD

JBOD = « Just a Bunch Of Disks » : un châssis **sans carte mère ni CPU**,
juste des baies, des expanders, des alimentations et des ventilateurs.
Il se pilote depuis un serveur via une HBA externe. Usage : étendre un
nœud de stockage au-delà de ses baies internes, ou construire du capacitaire
piloté par un petit serveur.

## 74. JBOD vs serveur de stockage

| Critère | JBOD | Serveur de stockage |
|---|---|---|
| CPU/RAM | non (tête externe) | oui |
| €/baie | moins cher | plus cher |
| Flexibilité | dépend de la tête | autonome |
| Point de panne | tête unique | nœud complet |

Pour Ceph : un JBOD derrière un nœud Ceph, c'est 60 OSD sur **un seul**
domaine de panne (la tête). Mieux vaut souvent 2 nœuds de 30 baies que
1 nœud + 1 JBOD de 60.

## 75. Topologie simple : HBA → JBOD

```
 ┌──────────┐   SFF-8644/8674    ┌──────────────────────┐
 │ Serveur  │────────────────────▶│ JBOD 60 baies        │
 │ HBA 16e  │◀────────────────────│ Expander A           │
 └──────────┘   (2 câbles = 8     └──────────────────────┘
                 voies = 9,6 Go/s)
```

Deux câbles externes = 8 voies 12G = 9,6 Go/s théoriques pour 60 disques.
Pour 60 HDD (~15 Go/s si tous à fond — irréaliste) c'est OK ; pour des
SSD SATA (60 × 550 Mo/s = 33 Go/s), **c'est un goulot** : ajoute des
câbles ou passe en 24G.

## 76. Topologie cascade (daisy chain)

```
 Serveur (HBA 16e)
    │ 8 voies
    ▼
 ┌──────┐  4 voies ┌──────┐
 │JBOD-1│─────────▶│JBOD-2│
 │60    │◀─────────│60    │
 └──────┘          └──────┘
   120 disques sur 8 voies = 9,6 Go/s partagés
```

La cascade économise des ports HBA mais **divise la bande passante** et
ajoute un domaine de panne (JBOD-1 down = JBOD-2 isolé, sauf double
domaine). Limite raisonnable : 2 JBOD en cascade, jamais 3+ en production.

## 77. Topologie redondante double domaine

```
              ┌──────────┐
              │ Serveur  │
              │ HBA 16e  │
              └──┬───┬───┘
                 │   │  (domaine A / domaine B)
        ┌────────▼┐ ┌▼────────┐
        │ Expdr A │ │ Expdr B │   ◀── deux expanders par JBOD
        └────────┬┘ └┬────────┘
                 │   │
          disques dual-port (SAS)
```

Chaque disque SAS dual-port est vu par les deux domaines. Un câble ou un
expander qui meurt = bascule transparente (avec multipath, section 36).
**Obligatoire** pour tout JBOD en production sérieuse. Ne fonctionne
**qu'avec des disques SAS** (le SATA n'a qu'un port).

## 78. Zoning d'expander

Le zoning découpe un expander en groupes isolés : par exemple, 1 JBOD de
60 baies partagé entre **2 serveurs** (30 baies chacun), chacun ne voyant
que ses disques. Utile en cluster 2 nœuds actif/actif ou pour isoler un
domaine de test. Mal configuré, c'est aussi le meilleur moyen de rendre
la moitié des disques invisibles — documente le plan de zoning.

## 79. SES-3 : le monitoring du JBOD

SCSI Enclosure Services : températures, ventilateurs, alimentations, état
des baies, LEDs de localisation. Exposé via la HBA (`storcli`, `sg_ses`).
**Branche le SES à ta supervision** : un ventilateur de JBOD mort un
vendredi soir = 60 disques qui cuisent le week-end. Seuil d'alerte :
température interne > 45 °C.

## 80. Limite n°1 : la bande passante partagée (calcul)

Formule : `bande_passante_amont / nombre_de_disques_actifs = débit/disque`.
Exemple : 8 voies 12G = 9,6 Go/s ; 60 HDD en rebuild simultané à 250 Mo/s
chacun = 15 Go/s demandés → **le lien amont sature à 64 % du besoin**.
En régime normal (quelques disques actifs), pas de problème. En rebuild
généralisé ou en resilver ZFS, le temps de reconstruction **double**.
Dimensionne l'amont pour le **pire cas**, pas le cas nominal.

## 81. Limite n°2 : le nombre de disques

Au-delà de ~120-200 disques par HBA, le firmware de la HBA et les temps
d'inventaire (découverte SAS) deviennent pénibles : boot de plusieurs
minutes, `storcli show all` qui rame. Règle : **1 HBA 16 voies pour
60-120 disques max** en production.

## 82. Latence ajoutée par l'expander

Chaque traversée d'expander ajoute quelques microsecondes. En cascade
(HBA → exp1 → exp2 → disque), on peut ajouter 10-20 µs aller-retour.
Négligeable pour du HDD (ms), **mesurable** pour du NVMe/SAS-SSD à
faible latence (< 100 µs). Pour du flash performant : pas de cascade,
backplane direct.

## 83. Alimentations et refroidissement du JBOD

Un JBOD 60 baies : 2+ alimentations redondantes (souvent 1200-1600 W),
ventilateurs hot-swap. Le JBOD **ne s'arrête pas proprement** tout seul :
en cas de coupure, c'est la tête (serveur) qui doit gérer l'arrêt via
NUT/stonith. Prévois le JBOD dans le plan d'arrêt électrique (lien avec
ton guide onduleurs : même onduleur, séquence d'arrêt documentée).

## 84. Poids et manutention

Rappel section 26, aggravé : un JBOD 60 baies plein, c'est le même poids
qu'un serveur 4U (~60 kg) **sans les poignées ni l'équilibrage d'un
serveur**. Certains JBOD n'ont pas de rails coulissants : maintenance =
sortir le châssis. Vérifie ce point **avant** l'achat.

## 85. Firmware d'expander

