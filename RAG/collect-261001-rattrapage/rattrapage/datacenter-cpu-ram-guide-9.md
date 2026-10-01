---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-9
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["arr", "benchmarks", "datacenter", "gpu"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [1285, 1447]
sha256: eafef355ac0348dc5f919d85f871f9b1348e9648999410f7dc8725e8f998d74b
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

1. **Baie murale ou meuble fermé** : si l'air chaud de l'arrière est
   réaspiré à l'avant (recirculation), la température d'entrée grimpe en
   boucle → emballement thermique. Symptôme : serveurs qui ventilent à fond
   dans une salle « fraîche ».
2. **Ventilateurs montés à l'envers** après maintenance : certains modules
   sont réversibles physiquement mais pas logiquement. Un fan wall qui
   souffle à contre-sens = points chauds invisibles au BMC (les sondes
   d'entrée/sortie sont placées pour le sens normal).
3. **Cache-baies (blanking panels)** : toute unité vide **doit** être
   obturée. Un U vide sans cache = court-circuit aéraulique : l'air frais
   passe à côté des serveurs au lieu de les traverser.

Schéma allées :

```
 ALLEE FROIDE (18-27 C)          ALLEE CHAUDE (35-45 C)
  [clim souffle]                      [clim aspire]
        ||                                   ^
  +-----+-----+                         +----+-----+
  | SERVEUR   |-- air chaud -->          | SERVEUR  |
  +-----+-----+                         +----+-----+
        ||                                   ^
  (jamais l'inverse)
```

## 64. Déflecteurs d'air (air shrouds)

Le capot plastique qui coiffe CPU et RAM n'est pas un emballage : c'est un
**déflecteur aéraulique** qui force l'air à travers les ailettes des
dissipateurs au lieu de le laisser passer au-dessus.

- **Ne jamais faire tourner un serveur sans son shroud** (même « pour
  tester ») : les températures CPU peuvent grimper de 15 à 25 °C.
- Après toute intervention (ajout de RAM, de carte), **remettez le shroud
  en place** : c'est l'oubli n°1 en maintenance.
- Certains shrouds sont spécifiques à la configuration (1 CPU vs 2 CPU,
  avec/sans GPU) : vérifiez la référence lors d'un upgrade.

## 65. Refroidissement liquide : panorama 2026

Avec des CPU à 500 W et des GPU à 700-1000 W, l'air atteint ses limites.
État des technologies (vérifié le 27/09/2026 dans leurs principes, détails
par constructeur à valider) :

| Technologie | Principe | Maturité |
|---|---|---|
| Dissipateur air renforcé | ailettes + heat-pipes, fans rapides | standard jusqu'à ~350-400 W/CPU |
| Direct-to-chip (DLC) | plaque froide sur CPU/GPU + boucle eau tiède | **standard HPC/IA**, se généralise |
| Immersion (bain d'huile) | serveurs plongés dans diélectrique | niche (mining, HPC expérimental) |
| Plaques arrière de porte | échangeur eau/air en sortie de baie | retrofit de salles existantes |

Le **direct-to-chip à eau tiède (30-40 °C)** est la technologie de
référence pour les déploiements CPU 500 W+ et GPU : pas besoin d'eau
glacée, le « froid » peut être du free cooling presque toute l'année
sous nos latitudes.

## 66. Direct-to-chip vs immersion : repères de décision

| Critère | Direct-to-chip | Immersion |
|---|---|---|
| Efficacité (PUE atteignable) | 1,1-1,3 | 1,05-1,15 |
| Retrofit salle existante | oui (boucle + CDU) | non (bacs dédiés) |
| Maintenance | standard (plaques rapides) | contraignante (égouttage) |
| Compatibilité matériel | kits par serveur/GPU | châssis spécifiques |
| Coût d'entrée (indicatif) | moyen (CDU + plomberie) | élevé (bacs + fluide) |

Pour un chef de service : **le DLC est l'option réaliste** pour une salle
existante qui accueille des serveurs 500 W ou des GPU. L'immersion reste
un choix de conception neuve. Dans les deux cas, l'eau tiède permet de
**valoriser la chaleur** (chauffage de bureaux) — voir section 92.

## 67. Températures cibles et throttling

Repères (indicatif, varient selon CPU ; seuils exacts dans la fiche
technique du processeur) :

| Sonde | Zone normale | Alerte | Critique |
|---|---|---|---|
| CPU (Tcase/Tctl) | 50-75 °C en charge | ~85 °C | ~95-100 °C (throttle) |
| DIMM DDR5 | 40-60 °C | ~75 °C | ~85 °C (erreurs) |
| Entrée d'air (ASHRAE A2) | 18-27 °C | 30 °C | 35 °C+ |
| Sortie d'air | 30-45 °C | — | — |

- **Throttling** : au-delà du seuil, le CPU baisse sa fréquence pour
  survivre. Symptôme : performances en dents de scie, jamais de crash franc.
  Si vos benchmarks sont instables, **regardez les températures avant
  d'accuser le code**.
- Les DIMM DDR5/MRDIMM chauffent plus que la DDR4 : surveillez leurs
  sondes, surtout en 2 DPC et avec MRDIMM (section 17 : +5 °C mesurés).

## 68. Maintenance des ventilateurs : routine

- **Trimestrielle** : contrôle visuel (poussière sur les grilles avant),
  vérification des alertes BMC, écoute des bruits anormaux (roulement).
- **Annuelle** : dépoussiérage (air comprimé sec, serveur arrêté si
  possible ; sinon par l'avant), vérification des courbes PWM après toute
  mise à jour BIOS/BMC (les réglages reviennent parfois aux défauts).
- **Au remplacement** : notez les heures de fonctionnement si le BMC les
  donne ; remplacez par la **référence exacte** (débit/pression, pas
  seulement la taille).
- **Stock** : 1-2 modules de rechange par modèle (section 59).
- **Ne jamais** : bloquer un ventilateur au doigt pour « tester »
  (contre-électromotrice, casse du roulement), ni lubrifier un roulement
  scellé.

---

## 69. Formats rack : 1U, 2U, 4U — panorama

1U = 44,45 mm de hauteur. Le format détermine le refroidissement possible,
donc le TDP CPU acceptable, donc le choix du processeur.

| Format | Hauteur | TDP/CPU typique (air) | Usages types |
|---|---|---|---|
| 1U | 44 mm | jusqu'à ~350 W (500 W = limite extrême) | densité, web, hyperconvergé |
| 2U | 89 mm | jusqu'à 500 W confortablement | standard : virtualisation, DB, HPC |
| 4U | 178 mm | 500 W+ / GPU | GPU, stockage dense, 4S |

Règle : **le format se choisit après le CPU, pas avant.** Un 2× 500 W
dans un 1U « parce qu'on manque de place » est une faute thermique qui se
paie en throttling, en bruit et en durée de vie. Si la densité l'impose,
prévoyez le DLC (section 65) dès l'achat.

## 70. 1U : contraintes thermiques et mécaniques

Le 1U est le format le plus exigeant :
- **Thermique** : 44 mm de hauteur = dissipateurs plats, ventilateurs de
  40 mm à très haut régime (bruyants, gourmands — section 62). 2× 500 W en
  1U air = à la limite du raisonnable ; au-delà, DLC obligatoire.
- **Mécanique** : cartes d'extension via risers (cartes couchées) ;
  nombre de slots PCIe et de baies disque limité (souvent 8-10 NVMe max).
- **Alimentation** : blocs compacts, redondance parfois en N+0 sur les
  configs extrêmes (section 59).
- **Poids** : 15-25 kg équipé — manipulable à une personne avec rails.

Quand le choisir : densité maximale (hébergeurs, scale-out), workloads
homogènes, salle bien climatisée avec allées strictes. **Jamais par
défaut** : le 2U est plus tolérant pour 1U de « perdu ».

## 71. 2U : le format polyvalent

Le 2U est le standard du datacenter d'entreprise, et ce n'est pas un
hasard :
- Dissipateurs de 2U (hauteur d'ailettes doublée) : 2× 400-500 W en air,
  sans héroïsme.
- 8 à 24 baies disque (2,5" ou 3,5"), plusieurs slots PCIe pleine hauteur,
  parfois 2-4 GPU double largeur.
- Ventilateurs de 60-80 mm : moins bruyants et plus efficaces qu'en 1U.
- Redondance fans et alimentations systématique sur les gammes pro.

**Recommandation de ce guide : si vous hésitez entre 1U et 2U, prenez 2U.**
Le surcoût d'un U se rentabilise en fiabilité thermique, en silence et en
marge d'évolution (ajout de cartes, de disques, upgrade CPU).

## 72. 4U : GPU, stockage dense et cas particuliers

| Usage 4U | Configuration type |
|---|---|
| Serveur GPU (IA) | 4-8 GPU double largeur + 2 CPU host |
| Stockage dense | 24-60 baies 3,5" (jusqu'à ~1 Po brut) |
| Quad-socket | 4 CPU + RAM massive (niche, section 27) |
| Tour de calcul | CPU 500 W + DLC + gros VRM |

