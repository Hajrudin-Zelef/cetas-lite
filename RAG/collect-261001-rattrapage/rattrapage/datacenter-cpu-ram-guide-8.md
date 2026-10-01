---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-8
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["compute", "datacenter", "dram", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [1114, 1284]
sha256: c284e3a19de464dd53ac9eb2695360245f0d0a1ecef5e77707c6ae5c2de876e8
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

| Fonction | Rôle |
|---|---|
| ECC SECDED | corrige 1 bit, détecte 2 bits (section 41) |
| SDDC / Chipkill | survit à la panne **complète d'une puce DRAM** |
| DDDC (Intel) | double protection puce (certaines plateformes) |
| PPR (Post-Package Repair) | remplace une ligne défectueuse par une ligne de secours, à chaud |
| Memory sparing/mirroring | bascule sur barrette de secours ou recopie (coûte 50 % de capacité) |

En pratique :
- Supervisez les **compteurs d'erreurs ECC corrigées** via le BMC/IPMI :
  une barrette qui corrige de plus en plus est une barrette à remplacer
  **avant** qu'elle ne tue un job ou une VM.
- Le mirroring mémoire divise la capacité par 2 : ne l'activez que pour
  les workloads qui l'exigent vraiment (certaines DB critiques).
- Après tout remplacement de barrette, **revérifiez la symétrie**
  (section 45) : une barrette de capacité différente casse l'entrelacement.

## 56. CXL : la mémoire désagrégée (état vérifié le 27/09/2026)

CXL (Compute Express Link) permet d'ajouter de la mémoire via PCIe, avec
cohérence de cache : extension, pooling entre hôtes, et à terme
composition de ressources à l'échelle du rack.

| Version | Capacité clé | Statut (vérifié 27/09/2026) |
|---|---|---|
| CXL 1.1 | mémoire attachée au device | mature |
| CXL 2.0 | pooling, switch, hot-plug géré | **en production** (Azure depuis 11/2025) |
| CXL 3.0/3.1 | fabric, cohérence multi-hôtes | silicium de switch en échantillonnage |
| CXL 4.0 | sur PCIe 7.0 | IP uniquement (attendu ~2026) |

Ordres de grandeur de latence : DDR5 locale 80-90 ns ; CXL 2.0 direct
170-250 ns ; CXL 3.x commuté 250-500 ns (analyses publiques 2026).

Positionnement 2026 : **le CXL n'est pas encore un produit d'achat courant
pour une PME/ETI.** C'est une technologie d'hyperscalers et d'appliances
(extension mémoire pour l'IA, tiering). Prévoyez des plateformes **CXL-ready**
(CPU + BIOS : Turin, Xeon 6, Venice le sont) sans acheter de matériel CXL
aujourd'hui — sauf besoin avéré d'extension mémoire.

---

## 57. Pourquoi le refroidissement est un sujet CPU

Un CPU moderne ne « consomme » pas : il **convertit** l'électricité en
chaleur, à 99 %+ (le 1 % restant part en signaux). Refroidir un EPYC 9965,
c'est évacuer **500 W thermiques** d'une surface de quelques cm² — soit une
densité de flux comparable à une plaque de cuisson.

Chaîne thermique complète :

```
 CPU (500 W) -> dissipateur -> air du chassis -> ventilateurs -> allee chaude
   -> climatisation salle -> groupe froid -> facture electrique (x2 a x3)
```

Chaque watt CPU engendre 0,3 à 1 W de climatisation selon le PUE
(section 81). **Choisir un CPU 500 W au lieu de 300 W, c'est choisir
~300 W de froid en plus.** Le refroidissement n'est pas un accessoire :
c'est la moitié du sujet énergie.

## 58. Fan wall : le principe du mur de ventilateurs

Dans un serveur rack, pas de « ventilateur de CPU » individuel : un **mur
de ventilateurs** (fan wall, 4 à 8 modules) pousse l'air d'avant en arrière
à travers tout le châssis.

```
 Vue de dessus (flux avant -> arriere) :

  [entree d'air frais] 
        |
  v +---+---+---+---+---+---+ v
    | F | F | F | F | F | F |    <- fan wall (6 modules)
  v +---+---+---+---+---+---+ v
        |
  [CPU0][RAM ][CPU1][RAM ][PCIe][NVMe]  <- composants en travers du flux
        |
  [sortie air chaud] -> allee chaude
```

Caractéristiques :
- Modules **hot-swap** : remplaçables serveur allumé (avec redondance N+1,
  section 59).
- Pilotage **PWM** commun par le BMC selon les sondes (CPU, DIMM, entrée).
- Les ventilateurs serveur sont des modèles **contrarotatifs** (deux
  hélices) : forte pression statique pour traverser dissipateurs et
  déflecteurs.

## 59. Redondance N+1 des ventilateurs

N+1 : le serveur peut perdre **un** module ventilateur sans surchauffe.
En pratique :
- 6 modules installés, 5 suffisent : si un tombe, les autres accélèrent.
- Le BMC alerte (voyant, SNMP, IPMI) : remplacez le module **sans attendre**,
  car la redondance est alors consommée.
- Certains châssis 1U denses sont en N+0 sur les configurations extrêmes
  (2× 500 W) : **vérifiez la fiche technique** — un 1U sans redondance fan
  est un pari thermique.

Règle d'exploitation : gardez **1 à 2 modules de rechange** par modèle de
serveur en stock (coût indicatif : quelques dizaines d'euros pièce). C'est
l'assurance la moins chère du datacenter.

## 60. Courbes PWM : lecture et réglage

Le BMC pilote les ventilateurs en PWM (0-100 %) selon des courbes
température/vitesse. Deux profils types :

| Profil | Logique | Usage |
|---|---|---|
| Acoustique / éco | fans au minimum jusqu'à ~70 °C CPU | bureaux, salles non critiques |
| Performance / datacenter | fans agressifs dès 60 °C | salles climatisées, HPC |

Points clés :
- En dessous de ~30 % PWM, beaucoup de ventilateurs serveur **ne démarrent
  pas** (seuil de démarrage) : ne réglez jamais une courbe sous ce seuil.
- Le mode « plein régime » permanent double ou triple la consommation des
  ventilateurs (loi cubique, section 62) pour un gain thermique marginal :
  laissez le BMC réguler, sauf diagnostic.
- Après un changement de CPU (upgrade TDP), **revérifiez les seuils** :
  une courbe calibrée pour 200 W laisse un 400 W throttler.

## 61. Bruit : dB(A) et réalité terrain

Un serveur 2U « normal » : 45-60 dB(A) à 1 mètre en charge. Un 1U dense
2× 500 W : **70-80 dB(A)** — niveau d'une tondeuse à gazon, intenable pour
un humain à proximité prolongée.

| Situation | Repère |
|---|---|
| Bureau open-space | < 45 dB(A) exigé → serveurs tour silencieux uniquement |
| Local technique fermé | 60-70 dB(A) tolérable en intervention courte |
| Salle datacenter | 75-85 dB(A) : protections auditives pour séjours longs |

Conséquences :
- **Jamais de serveur rack 1U/2U dans un bureau** : le bruit seul le
  disqualifie, avant même la chaleur.
- Les courbes « silencieuses » des BIOS grand public n'existent pas sur
  les serveurs : le BMC privilégie toujours la survie du matériel.
- En edge (armoire de rue, atelier), prévoyez des châssis à ventilateurs
  lents ou du refroidissement passif renforcé — et mesurez.

## 62. Consommation des ventilateurs : la loi cubique

Physique des ventilateurs : **P ∝ vitesse³**. Doubler la vitesse =
8× la puissance.

| Régime | Vitesse relative | Puissance relative |
|---|---|---|
| 30 % PWM | 0,3 | ~3 % |
| 50 % PWM | 0,5 | ~12 % |
| 70 % PWM | 0,7 | ~34 % |
| 100 % PWM | 1,0 | 100 % |

Ordres de grandeur (indicatif) : un module de fan wall 2U consomme ~10 W
à mi-régime et ~30-40 W à plein régime. Un serveur 2U avec 6 modules à
plein régime = **~200 W de ventilation seule** — autant qu'un petit CPU.

Implications :
- Un datacenter qui laisse ses serveurs en « fans à fond » permanent
  gaspille 5 à 10 % de sa facture serveurs.
- Le free cooling / l'augmentation de la température de consigne salle
  (27 °C au lieu de 22 °C, norme ASHRAE A2) **réduit** la facture de froid
  mais **augmente** la consommation des ventilateurs : l'optimum se calcule
  (section 92).

## 63. Le piège du sens du flux : avant/arrière

Standard datacenter : **entrée d'air en face avant, sortie en face
arrière** (allée froide / allée chaude). Les pièges :

