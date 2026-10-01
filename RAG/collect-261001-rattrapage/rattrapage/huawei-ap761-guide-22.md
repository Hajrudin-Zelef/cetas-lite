---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-22
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [2176, 2287]
sha256: 468bcffeac8726c1359c2f2c9efd96d04c8d7bde85f5a84cebb442231936388a
---

# Guide ultra-complet — Huawei eKit AP761

**Lecture :** à Wi-Fi égal, l'AP761 coûte **~3× le prix** de l'AP361 : tu paies l'IP68, le métal, les antennes directionnelles, le SFP, la plage de température. **Ne jamais mettre un AP761 en intérieur** (surcoût inutile + secteur directionnel inadapté) **ni un AP361 en extérieur** (il mourra à la première pluie — et la garantie ne suivra pas).

## 97. Quand choisir l'AP361 plutôt que l'AP761

**Choisir l'AP361 quand :**
- ✅ La zone est **intérieure** (bureau, salle de réunion, commerce, hall).
- ✅ Le budget est serré : ~3× moins cher que l'AP761 pour le même Wi-Fi 6.
- ✅ La consommation compte : 8.8 W vs 17.7 W — sur un gros parc, ça se voit sur le budget PoE et l'onduleur.
- ✅ Il faut une couverture **à 360°** (open space) : ses smart antennas omnidirectionnelles sont faites pour ça.
- ✅ Le montage est au **plafond** (format disque discret).

**Ne pas choisir l'AP361 quand :**
- ❌ Extérieur ou local humide/poussiéreux (pas d'IP68).
- ❌ Il faut de la **distance directionnelle** (cour de 100 m : l'AP361 arrosera partout faiblement au lieu de porter loin dans l'axe).
- ❌ Il faut un **uplink fibre** (pas de SFP).
- ❌ Environnement industriel chaud/froid extrême.

## 98. Quand choisir l'AP761 plutôt que l'AP361

**Choisir l'AP761 quand :**
- ✅ **Extérieur**, sans hésitation : c'est son métier (IP68, −40 à +65 °C, 6 kV).
- ✅ Zone à couvrir en **secteur** (cour, parking, terrasse) : les antennes directionnelles 65° portent loin et proprement.
- ✅ Distance **> 100 m** jusqu'au local : le SFP fibre règle le problème.
- ✅ Environnement **hostile** : poussière (entrepôt ouvert, carrière), embruns (bord de mer — rincer le sel lors des visites, chap. 102), variations thermiques fortes.
- ✅ Besoin de **BLE** pour l'onboarding/maintenance de proximité (5.2) — à vérifier pour l'AP361.

**Ne pas choisir l'AP761 quand :**
- ❌ Intérieur : surcoût, secteur directionnel inadapté, esthétique « pavé industriel » au plafond d'un bureau.
- ❌ Le besoin est du **Wi-Fi 7** : voir chap. 99.
- ❌ Le budget PoE est déjà tendu : 17.7 W par AP, ça se planifie (chap. 25).

## 99. Et le Wi-Fi 7 outdoor ? AP771 vs AP772E

Quand le besoin Wi-Fi 7 outdoor est avéré (chap. 95), les deux candidats eKit (valeurs vérifiées le 27/09/2026 — détails **à vérifier sur la fiche du modèle exact**) :

| Critère | AP771 | AP772E |
|---|---|---|
| Wi-Fi | 7 (802.11be), bi-bande 2.4/5 GHz | 7 (802.11be), bi-bande 2.4/5 GHz |
| Débit max | 3.57 Gbps | **6.45 Gbps** |
| Ports | à vérifier | 1× **2.5GE** PoE + 1× **10GE SFP+** |
| PoE | à vérifier | **802.3af/at/bt** ou DC |
| Protection | à vérifier (outdoor) | IP68, −40 à +70 °C, 6 kA |
| Portée indicative | ~130 m | ~250 m (optimale) |
| Cas d'usage | Outdoor Wi-Fi 7 **économique**, zones modérées | Outdoor Wi-Fi 7 **haute capacité**, zones denses |

**Ce que le choix implique (au-delà du prix de l'AP) :**
- **AP772E :** switch **2.5G PoE++ (802.3bt)** ou injecteur bt, câblage **Cat6A** conseillé pour la 10G sur SFP+ (fibre en pratique), budget PoE **recalculé** (un AP Wi-Fi 7 consomme nettement plus que 17.7 W — valeur exacte à vérifier sur la fiche du modèle exact).
- **AP771 :** marche probable sur infra existante 1G/PoE+ — à vérifier, mais c'est son positionnement « Wi-Fi 7 accessible ».
- **Les deux :** pas de 6 GHz, pas de 320 MHz (chap. 16) — le Wi-Fi 7 outdoor eKit, c'est du MLO + 4K-QAM + efficacité, pas des canaux géants.

## 100. Tableau décisionnel global : quel AP eKit pour quel besoin

| Besoin | Choix | Pourquoi |
|---|---|---|
| Bureau PME, Wi-Fi 6, petit budget | **AP361** | Le moins cher, 8.8 W, largement suffisant |
| Bureau dense, Wi-Fi 6 | **AP362/AP362E** | 2.975 Gbps, plus de capacité |
| Open space très dense / auditorium | **AP661** | 6.575 Gbps, tri-radio |
| **Cour, parking, terrasse — Wi-Fi 6** | **AP761** | **IP68, directionnel, SFP : c'est lui** |
| Cour/parking — Wi-Fi 7, budget serré | **AP771** | Wi-Fi 7 outdoor accessible — à vérifier |
| Zone dense outdoor — Wi-Fi 7 | **AP772E** | 6.45 Gbps, 2.5G/10G, bt — à vérifier |
| Indoor — Wi-Fi 7 économique | **AP371** | 3.57 Gbps, 2.5GE |
| Indoor — Wi-Fi 7 max (avec 6 GHz) | **AP673** | Tri-bande, 320 MHz — à vérifier |
| Chambre d'hôtel | **AP160/AP162E** | Format mural, discret |

**Règle d'or :** on choisit d'abord **l'environnement** (indoor/outdoor), puis **la capacité** (débit/densité), puis **la génération** (Wi-Fi 6 vs 7). Choisir la génération d'abord, c'est acheter un AP772E à 600 € pour couvrir un parking vide.

## 101. Maintenance préventive : planning annuel

Un AP extérieur s'entretient comme tout équipement exposé. Le planning type pour un parc d'AP761 :

| Fréquence | Action | Chapitre |
|---|---|---|
| **Mensuelle** (15 min, à distance) | Revue console cloud : AP en ligne, alertes, utilization canal, clients | 66–67 |
| **Trimestrielle** (à distance) | Tendances : un site se remplit-il ? Firmware à jour ? Sauvegardes OK ? | 72, 74 |
| **Semestrielle** (sur site) | Contrôle visuel : fixation, câble, presse-étoupes, parafoudre (voyant), nettoyage | 102 |
| **Annuelle** (sur site) | Contrôle électrique : PoE, terre, serrages. Contrôle radio : mini-survey, canaux, puissances | 103–104 |
| **Annuelle** (bureau) | Revue du dossier de site, test de restauration d'une sauvegarde, revue des alertes | 73 |
| **Après chaque orage violent** | Contrôle parafoudres + état des AP exposés | 26 |
| **Après travaux à proximité** | Vérifier orientation, obstacles nouveaux, câbles | 62 |

**Le dossier de site** (le vrai livrable de la maintenance) :
```
dossier-site/
├── plan-implantation.pdf (AP positionnés, secteurs tracés)
├── photos-avant-apres/
├── releves-survey/ (fiches de mesure chap. 60)
├── configs/ (sauvegardes chap. 72)
├── firmwares/ (versions + release notes)
├── interventions.log (date, qui, quoi, pourquoi)
└── contacts/ (astreinte, SAV, distributeur)
```
Un site sans dossier à jour = un site qu'on redécouvre à chaque panne. **Le dossier, c'est la moitié de la maintenance.**

## 102. Contrôle visuel et nettoyage (semestriel)

**Checklist visite semestrielle (par AP, 10 min) :**
- [ ] **Fixation :** étrier/colliers serrés ? Pas de jeu ? (Revérifier le serrage après l'hiver — le gel/dégel desserre.)
- [ ] **Orientation :** l'AP pointe-t-il toujours vers la zone ? (Tempête, travaux, vandalisme.)
- [ ] **Câble :** pas de frottement, pas de gaine craquelée par les UV, boucle d'égouttage en place.
- [ ] **Presse-étoupes :** serrés, pas de trace d'eau, joints souples.
- [ ] **Parafoudre :** voyant d'état OK (chap. 26). Le remplacer si doute.
- [ ] **Environnement :** végétation qui a poussé devant le secteur ? Nouvel obstacle (chantier, enseigne, camion permanent) ?
- [ ] **Propreté :** dépoussiérer le boîtier (chiffon humide, **jamais de nettoyeur haute pression** — l'IP68 a des limites face au karcher).
- [ ] **Bord de mer :** rincer le sel au jet doux — la corrosion adore les fixations.
- [ ] **Étiquetage :** l'étiquette d'identification est-elle toujours lisible ? (Sinon : ré-étiqueter.)

**Noter dans le dossier :** date, technicien, observations, photos si anomalie. Une anomalie notée = une anomalie suivie.

## 103. Contrôle électrique et PoE (annuel)

Ton domaine, Zelef — le Wi-Fi n'est que la partie émergée.

