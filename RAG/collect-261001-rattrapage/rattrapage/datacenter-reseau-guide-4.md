---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-4
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["United States"]
dates: ["2026-09-27"]
keywords: ["datacenter", "attention", "serdes"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [387, 550]
sha256: 007a5a188950486604512c2a6aa53eea0f4016637c2f714bb4e5cf937dae95d1
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| Débit / type | OM3 | OM4 | OM5 (SWDM) | OS2 |
|---|---|---|---|---|
| 10G SR | 300 m | 400 m | — | — |
| 25G SR | 70 m | 100 m | — | — |
| 100G SR4 | 70 m | 100 m | — | — |
| 100G BiDi (SWDM4) | — | — | 100-150 m (2 brins) | — |
| 100G DR | — | — | — | 500 m |
| 100G FR1 | — | — | — | 2 km |
| 100G LR4 | — | — | — | 10 km |
| 400G SR8 | 60 m | 100 m | — | — |
| 400G DR4 | — | — | — | 500 m |
| 400G FR4 | — | — | — | 2 km |
| 400G LR4 | — | — | — | 10 km |
| 800G SR8 | 60 m | 100 m | — | — |
| 800G DR8 | — | — | — | 500 m |

**En pratique datacenter** : 90 % des liens font < 30 m (intra-rack/rack
adjacent) → DAC/AOC ; les liens leaf↔spine inter-rangées (30-100 m) → SR8/AOC ;
au-delà → DR sur OS2. **Le 800G SR8 à 60 m sur OM3** : ne pas réutiliser de
vieux tronçons OM3 pour du 800G sans mesurer.

## 20. Budgets optiques : le calcul qui valide un lien

Budget (dB) = Puissance Tx (dBm) − Sensibilité Rx (dBm).
Marge = Budget − (affaiblissement fibre + connecteurs + épissures).
**Règle : marge ≥ 3 dB** sur un lien neuf.

| Élément | Affaiblissement typique (⚠️) |
|---|---|
| Fibre OM4 @850 nm | ~3 dB/km (négligeable < 100 m) |
| Fibre OS2 @1310 nm | ~0,4 dB/km |
| Connecteur LC/MPO apparié | 0,3-0,5 dB (0,75 dB max norme) |
| Épissure fusion | 0,05-0,1 dB |
| Jarretière sale | **1-5 dB** ← voir §24 |

**Exemple** : lien 400G-DR4, budget module 4 dB (⚠️ typique) :
fibre 80 m OS2 (0,03 dB) + 4 connecteurs MPO (4×0,4 = 1,6 dB) → perte ~1,7 dB,
**marge ~2,3 dB < 3 dB** : limite. Une jarretière sale fait tomber le lien.
D'où le §24.

**Mesure** : un lien qui « monte » n'est pas un lien « bon ». Mesurer la
puissance Rx (`show interfaces transceiver` / `ethtool -m`) et la comparer à
la sensibilité : si on est à < 2 dB de marge, planifier le nettoyage ou le
remplacement **avant** la panne.

## 21. Connecteurs : LC, MPO/MTP, SN, CS

| Connecteur | Fibres | Usage | Remarque |
|---|---|---|---|
| LC duplex | 2 | SR/DR/FR/LR classiques | Le standard, détrompage simple |
| MPO-12 | 12 (8 utilisées en DR4) | 100G SR4, 400G DR4 | Broches mâle/femelle : **ne pas mélanger** |
| MPO-16 | 16 | 400G/800G SR8, DR8 | 2 rangées de 8 |
| MPO-24 | 24 | Trunks haute densité | Brassage uniquement |
| SN / CS | 2-4, très compact | 800G 2×FR4, face avant dense | Nouveau, outillage dédié |
| MTP® | = MPO premium | Trunks | Marque US Conec, meilleure précision |

**Polarité MPO** (méthodes TIA) :
- **Méthode A** : droit 1↔1 (key-up/key-down).
- **Méthode B** : croisé 1↔12 — la plus courante en datacenter.
- **Méthode C** : paires croisées.
Un trunk méthode B avec des jarretières méthode A = lien muet. **Étiqueter la
méthode sur le trunk** (voir §85). C'est le P8 (§107).

## 22. DAC vs AOC vs transceivers — le vrai comparatif économique

| Critère | DAC (cuivre passif) | ACC/AEC (cuivre actif) | AOC (fibre active) | Transceiver + fibre |
|---|---|---|---|---|
| Portée 100G | 1-5 m | 5-7 m | 3-100 m | 100 m - 10 km |
| Portée 400G | 1-3 m | 3-5 m | 3-100 m | 100 m - 10 km |
| Portée 800G | 1-2 m (⚠️) | 3-6 m (⚠️) | 3-30 m (⚠️) | 100 m - 2 km |
| Conso / extrémité | ~0,1-0,5 W | ~2-4 W | ~1-2 W | 2,5-18 W (selon débit) |
| Latence | ~0 (ns) | ~100 ns | ~100 ns | ~100 ns |
| Prix indicatif 100G 3 m (⚠️) | 50-120 € | 150-300 € | 200-400 € | 300-800 € (SR4) |
| Prix indicatif 400G 3 m (⚠️) | 150-350 € | 400-800 € | 600-1 200 € | 1 500-4 000 € (DR4) |
| Prix indicatif 800G (⚠️) | 400-900 € | 1 000-2 000 € | 1 500-3 000 € | 4 000-10 000 € (DR8) |
| Avantage | Prix, zéro conso | Portée cuivre | Léger, fin | Interopérabilité, distance |
| Inconvénient | Rigide, encombrant | Conso, prix | Non réparable | Prix, conso, chaleur |

**Doctrine d'achat** :
1. Intra-rack (< 3 m) → **DAC** toujours, sauf 800G où l'AEC devient pertinent.
2. Rack adjacent (3-30 m) → **AOC** (ou AEC en 800G).
3. Inter-rangées → transceivers SR8 (OM4) ou DR (OS2).
4. Ne jamais mélanger DAC et transceiver sur le même lien : les seuils
   électriques diffèrent, l'auto-négociation échoue (P9, §107).

## 23. Câbles breakout : multiplier les ports sans multiplier les switches

Un port 400G peut servir 4 serveurs 100G (ou 2×200G), un port 800G → 2×400G /
4×200G / 8×100G.

| Breakout | Câble | Cas d'usage |
|---|---|---|
| 400G → 4×100G | QSFP-DD → 4× QSFP28 DAC/AOC | Leaf 400G vers serveurs 100G |
| 400G → 2×200G | QSFP-DD → 2× QSFP56 | Migration douce |
| 800G → 2×400G | OSFP → 2× QSFP112 | Leaf 800G vers nœuds 400G |
| 800G → 8×100G | OSFP → 8× SFP112 | Densité max (⚠️ limites par port, voir SN5600 : le 8×100G sur port impair consomme le port pair adjacent — ✅ vérifié 27/09/2026) |

**Attention** : tous les ports d'un switch ne supportent pas tous les modes
de breakout (ressources SerDes partagées). **Valider la matrice de breakout
dans le guide de câblage du switch** avant d'acheter les câbles — sinon des
ports restent inutilisables (P10, §107).

## 24. Nettoyage fibre — le piège n°1 du datacenter

**Une particule de 1 µm sur une férule = jusqu'à 5 dB de perte.**
Le cœur d'une fibre monomode fait 9 µm : un grain de poussière le couvre
entièrement. 80 % des pannes optiques « inexpliquées » viennent d'un
connecteur sale — pas d'un module HS.

**Procédure (à afficher dans la salle) :**
1. **Inspecter** (microscope, §25) → 2. **Nettoyer** → 3. **Ré-inspecter**.
   Ne jamais brancher sans inspecter : on transfère la saleté d'un
   connecteur à l'autre.
2. Nettoyage à sec d'abord (cassette(click), lingettes) ; humide (alcool
   isopropylique + séchage) seulement si le sec échoue.
3. **Bouchons** : tout connecteur débranché = bouchon remis immédiatement.
   Un transceiver sans bouchon dans un tiroir = poubelle optique.
4. Ne jamais toucher la férule, ne jamais souffler dessus (salive),
   ne jamais regarder dans une fibre active (laser classe 1M, risque
   rétinien en monomode).

**Kit minimum par salle** : 1 microscope d'inspection (sonde vidéo),
cassettes de nettoyage LC + MPO, lingettes, bouchons de rechange.
Coût : ~1 500-3 000 € (⚠️) — le prix d'**un seul** module 800G. Rentabilité
immédiate.

## 25. Inspection : seuils IEC 61300-3-35

La norme définit des zones (cœur, cladding, ferrule) et des seuils de
défauts (rayures, piqûres) par zone. En pratique :
- Les microscopes modernes donnent un ** verdict PASS/FAIL automatique** —
  l'exiger, ne pas juger « à l'œil ».
- Rejeter toute férule avec défaut dans la zone cœur/cladding.
- Documenter : photo avant/après pour les liens critiques (preuve en cas de
  litige fournisseur).

**Test terrain** : si un lien 400G+ affiche des erreurs FEC corrigées qui
augmentent (> 1e-9 pré-FEC, voir §27), **nettoyer avant** de changer le
module. Dans 4 cas sur 5, le « module défectueux » retourné au fournisseur
est en fait sale (constat d'atelier, ⚠️ retour d'expérience).

## 26. Polarité et genre MPO — ne pas mélanger

- **Mâle** = broches (pins) qui dépassent ; **femelle** = trous.
  Règle : **un lien = un mâle + une femelle**. Deux mâles = casse des
  broches ; deux femelles = mauvais contact.
- Les transceivers sont généralement **mâles**, les jarretières **femelles**
  vers le module — vérifier à la commande.
- **Schéma d'un lien MPO méthode B** :

```
  [Module Tx] ----fibre 1----> ----fibre 1----> [Module Rx]
  (MPO mâle)      croisée B      (MPO femelle)
  position 1  ==============>  position 12
  position 2  ==============>  position 11
  ...
```

- **Étiquette obligatoire** sur chaque trunk : méthode (A/B/C), genre des
  extrémités, longueur, référence. Sans ça, le brassage devient un jeu de
  devinettes à 3h du matin.

## 27. FEC : la correction d'erreur qui rend le PAM4 viable

