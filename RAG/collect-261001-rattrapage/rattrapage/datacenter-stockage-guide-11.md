---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-11
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "arr", "capex", "gpu", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [1746, 1920]
sha256: 2eeb55ab4a61f8daa83b348f6dbba908c3b46593efa1ff385fc48861eadde650
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Les SSD datacenter cités (D7-PS1010, CD9P-R, D5-P5430) sont garantis
**5 ans** (vérifié 27/09/2026), comme les HDD Exos (5 ans, vérifié).
En dessous (Barracuda 24 To : 1-2 ans selon canal — vérifié), c'est du
grand public : pas de datacenter. Aligne la **durée d'amortissement**
du projet sur la garantie : 5 ans.

## 182. Checklist fiabilité

- [ ] Burn-in 48-72 h avant production, résultats archivés
- [ ] smartd/nvme-cli → supervision, seuils écrits
- [ ] Scrub planifié et surveillé
- [ ] Stock de rechange : 2-5 % du parc, modèles identiques
- [ ] Politique de remplacement proactif écrite
- [ ] Firmware : versions inventoriées, flash par vagues

---

# PARTIE L — ÉNERGIE : CONSO, REFROIDISSEMENT, ALIMENTATION

## 183. Conso par disque : le tableau (vérifié le 27/09/2026)

| Disque | Actif | Idle | Source |
|---|---|---|---|
| HDD 3,5" 7200 tr/min (20-30 To) | 8-10 W | 5-7 W | ordre de grandeur (à vérifier par modèle) |
| SSD NVMe Gen5 7,68 To (D7-PS1010) | 17-18 W | 5 W | TechPowerUp (vérifié) |
| SSD NVMe E3.S (D7-PS1010, états) | 5-25 W (5 états) | 5 W | TweakTown/Solidigm (vérifié) |
| SSD QLC 30,72 To (D5-P5430) | jusqu'à 25 W | 5 W | spec Solidigm (vérifié) |
| KIOXIA CD9P (Gen5) | ~20-23 W typ. | — | product brief (vérifié, ordre de grandeur) |
| HBA 9500-16e | 8,74 W | — | fiche produit (vérifié) |
| HBA 9500-8e | 6,12 W | — | fiche produit (vérifié) |

Message clé : **un NVMe Gen5 consomme ~2× un HDD** à l'unité, mais
stocke autant pour 3-5× moins de baies. Le bilan se fait **au To
et au nœud**, jamais au disque.

## 184. Bilan nœud 60 baies HDD (chiffré)

| Poste | Calcul | Puissance |
|---|---|---|
| 60 HDD × 9 W | 540 W | 540 W |
| 2 CPU + carte mère + RAM | — | ~300 W |
| HBA + NIC + ventilateurs | — | ~100 W |
| **Total nœud** | | **~940 W** |
| + alim Titanium (96 %) | / 0,96 | ~980 W au mur |

Soit **~16 W par baie** et ~0,8 W/To (20 To/disque). Pour 10 nœuds :
~10 kW IT, ~14-15 kW avec PUE 1,5.

## 185. Bilan nœud 12 NVMe Gen5 (chiffré)

| Poste | Calcul | Puissance |
|---|---|---|
| 12 NVMe × 18 W | 216 W | 216 W |
| 2 CPU + 128 Go DDR5 | — | ~400 W |
| 4× NIC 100G (~20 W chacun) | 80 W | 80 W |
| Ventilateurs + divers | — | ~100 W |
| **Total nœud** | | **~800 W** |

Par To utile (25 To/nœud utiles en réplica 3) : **~32 W/To utile** —
~40× plus que le HDD capacitaire. Normal : ce n'est pas le même service
(latence < 1 ms vs ~10 ms).

## 186. Spin-down : principe et pièges

Le spin-down (arrêt de rotation des HDD inactifs) économise ~5-7 W par
disque. Pièges :
- **Ceph/ZFS réveillent les disques en permanence** (scrub, heartbeat,
  métadonnées) : le spin-down ne sert quasiment que sur de l'archivage
  pur (backup immuable, WORM).
- Les cycles spin-up/down **usent** la mécanique : un disque qui
  cycle 10×/jour meurt plus vite qu'un disque qui tourne.
- Le spin-up simultané de 60 disques = **pic de courant** (~2× la conso
  nominale pendant 10-20 s) : à intégrer au dimensionnement onduleur.
Verdict : utile en archivage froid, contre-productif en Ceph actif.

## 187. Refroidissement : le calcul

Règle (section 25) : ~1,7 m³/h par watt pour ΔT = 2 °C. Nœud 60 baies
à 980 W → **~1 660 m³/h**. Vérifie aussi la **température d'entrée** :
les specs disques sont données pour 0-60 °C ambiant boîtier, mais l'AFR
grimpe au-delà de 40 °C internes. Allée froide à 24-27 °C (ASHRAE A1) :
le bon compromis énergie/fiabilité.

## 188. PUE : le rappel pour le stockage

Le stockage dense a un **bon PUE apparent** (beaucoup d'IT par m²) mais
un mauvais **W/To** en NVMe. Quand tu présentes un projet : donne les
deux. Et n'oublie pas que la clim d'un nœud 60 baies à 1 kW, c'est
~0,5 kW de froid avec un PUE de 1,5 : le nœud coûte **1,5 kW** au
compteur, pas 1 kW.

## 189. Dimensionnement électrique : exemple 4 nœuds NVMe

4 nœuds × 800 W = 3,2 kW IT. Avec PUE 1,5 : 4,8 kW. Ajoute : 2 switchs
100G (~300 W chacun), marge 30 %, pic spin-up N/A (pas de HDD).
**Onduleur : 10 kVA mini** pour ce seul cluster (voir ton guide
onduleurs : dimensionnement 1,5-2×, autonomie selon besoin).
Câblage : 2 PDU par rack (A/B), chaque nœud en double alimentation
répartie A/B.

## 190. Onduleur et séquence d'arrêt

Le stockage est **le dernier à s'éteindre et le premier à redémarrer** :
1. Arrêt des VM/applis (Proxmox),
2. Arrêt des clients Ceph,
3. `ceph osd set noout` puis arrêt des OSD/nœuds,
4. Arrêt des JBOD,
5. Arrêt onduleur.
Redémarrage inverse. **Automatise avec NUT** (voir ton `proxmox_guide.md`,
section NUT) et **teste la séquence** 2×/an. Un arrêt brutal répété,
c'est des OSD à reconstruire et des SSD sans power-loss qui corrompent.

## 191. NUT : le lien avec ton guide onduleurs

NUT (Network UPS Tools) : l'onduleur signale la coupure, le serveur
maître NUT ordonne l'arrêt séquentiel. Pour un cluster Ceph : le
délai d'arrêt d'un nœud avec 12 OSD peut dépasser 10 min (flush) —
règle l'autonomie et les temporisations NUT en conséquence, pas avec
les valeurs par défaut d'un serveur unique.

## 192. Rendement Titanium : le calcul (vérifié)

Alimentations Titanium : **96 %** à 50 % de charge (vérifié : specs
SSG-6049P). Sur un nœud 940 W : perte 38 W vs 75 W en Gold (92 %).
Pour 10 nœuds, 24/7 : (75-38) × 10 × 8760 h ≈ **3 240 kWh/an**
économisés, soit ~650 €/an à 0,20 €/kWh — sans compter la clim
évitée (~×1,5 avec PUE). L'écart de prix Gold→Titanium se rembourse
généralement en < 2 ans.

## 193. C-states, ASPM, power caps

- Active les C-states CPU et l'ASPM PCIe en BIOS : -5 à -15 % sur un
  nœud peu chargé, sans impact mesurable sur Ceph (à valider en bench).
- `powercap` RAPL : plafonne un nœud à X W en cas de contrainte
  électrique (canicule, groupe en maintenance).
- Sur NVMe : les états de puissance (5 états sur D7-PS1010, vérifié)
  se pilotent via `nvme set-feature` — à réserver aux nœuds peu
  sollicités (froid), jamais sur du chemin critique latence.

## 194. Refroidissement liquide des SSD (vérifié)

Deux réalités 2026 (vérifié 27/09/2026) : KIOXIA NX1 (E1.S, liquide
direct) et Solidigm D7-PS1010 E1.S (cold plate, co-développé avec
NVIDIA pour serveurs GPU fanless). Si tu densifies en E1.S Gen5/Gen6 :
**la boucle d'eau n'est plus optionnelle**, elle fait partie du design
mécanique du rack. Prévois les CDU et les débits avec ton équipe
énergie — c'est ton terrain.

## 195. Températures relevées (vérifié)

Test TweakTown (vérifié) : KIOXIA CD9P-R 7,68 To E3.S à **53 °C** en
écriture séquentielle soutenue avec refroidissement à air classique
(contre ~70 °C redoutés). Traduction : l'E3.S bien ventilé tient l'air ;
l'E1.S dense et le Gen6 demanderont plus. Sonde les SSD (`nvme smart-log`,
`temperature`) comme les HDD : alerte à 65 °C, critique à 75 °C.

## 196. Airflow : EDSFF vs U.2

L'argument thermique n°1 de l'EDSFF : les règles verticales laissent
**passer l'air entre elles**, là où un mur de 24 U.2 15 mm en façade
le bloque. Sur un 2U32 E3.S, le flux traverse le châssis ; sur un
vieux 2U24 U.2, il contourne. À densité égale, l'EDSFF demande moins
de pression statique → ventilateurs moins rapides → **moins de watts
de ventilation et moins de bruit**.

## 197. Coût énergie : €/an par To (exemple)

Nœud 60 baies : 980 W au mur × 8760 h = 8 585 kWh/an. Avec PUE 1,5 :
12 877 kWh/an. À 0,20 €/kWh : **~2 575 €/an** pour ~870 To utiles
(EC 8+3 sur 60×20 To) → **~3 €/To/an** d'énergie. Nœud NVMe : 800 W →
~2 100 €/an pour 25 To utiles → **~84 €/To/an**. L'énergie creuse
l'écart HDD/NVMe : intègre-la au TCO 5 ans, pas seulement le CAPEX.

## 198. Dimensionnement clim : exemple rangée

