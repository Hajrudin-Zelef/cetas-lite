---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-17
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["datacenter", "arr", "capex", "compute", "gpu"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [2951, 3131]
sha256: bfe7fe13060c0ffd29f6ce947983bc2ef31ca0f5c0e16c3b3657ce2f130e40a1
---

# Datacenter Builds — Le guide des BOMs

1. **RAM « au cas où » en 2026** : 384 Go inutiles = 8 280 € brûlés (§32).
2. **DB sur Ceph sans QoS** : la latence réseau tue les TPS — NVMe local (§38).
3. **RAID 5/6 SSD pour une DB** : pénalité ×4 + GC = latence imprévisible (§38).
4. **Ceph rempli à > 85 %** : plus de rééquilibrage, perfs en chute libre (§46).
5. **Back Ceph en 10G** : le premier rebuild sature tout (§51).
6. **HDD SMR en Ceph** : inutilisable — CMR uniquement (§59).
7. **VDI sans test boot storm** : tout s'écroule le lundi 8h55 (§69).
8. **Licences vGPU oubliées** : 250 k€ sur 5 ans pour 200 users (§65).
9. **Acheter le GPU avant de tester le modèle** : louez 10 h d'abord (§89).
10. **Châssis non qualifié 600 W/GPU** pour RTX PRO 6000 (§99).
11. **Rack 1 000 mm** : les serveurs + PDU ne ferment pas (§136).
12. **Une seule PDU / phases déséquilibrées** : pas de maintenance sans arrêt,
    neutre qui chauffe (§136).
13. **Clim « au total » pas par rack** : le rack du fond surchauffe (§158).
14. **Onduleur en kVA** : 60 kVA ≠ 60 kW, ça disjoncte (§166).
15. **Burn-in sauté** : le disque qui meurt à J+15 coûte 10× plus cher (§174).

---

## Conclusion — la méthode en une page

1. **Mesurez** (working set, EPS, débit, IOPS) avant d'acheter.
2. **Calculez** (les 5 formules §10–14) — la RAM et les watts d'abord.
3. **Choisissez** la BOM S/M/quad du workload correspondant.
4. **Additionnez** les P_réaliste (§125) → onduleur (§162) + froid (§150).
5. **Vérifiez** les prix et délais **le jour du devis** (marché 2026 volatil).
6. **Burn-in 72 h**, firmware uniformes, as-built complet.
7. **Suivez** les 5 métriques €/unité (§124) chaque trimestre.

Un datacenter n'est pas une collection de serveurs : c'est une **chaîne
électrique, thermique et réseau** dimensionnée d'un seul tenant. Les BOMs de
ce guide sont les maillons ; les parties C et D sont la chaîne.

*Prix et références vérifiés le 27/09/2026 sauf mentions « à vérifier ».*
*Guide rédigé pour Zelef — chef de service systèmes & énergies.*

---

# PARTIE D — APPROFONDISSEMENTS

---

## 181. Cas chiffré global : salle de 8 racks, 60 kW IT

Hypothèse : PME/ETI qui construit sa salle. Contenu : 40 serveurs mixtes
(compute, DB, Ceph, backup, VDI), 2 racks réseau, onduleurs, froid.

| Lot | Détail | Budget |
|---|---|---|
| Serveurs (40×, mix §120 ×5) | 354 k€ × 5 | 1 770 000 € |
| 8× racks 42U équipés (§138) | 9 300 € × 8 | 74 400 € |
| Réseau : 16× ToR + 2× agg 100G | — | 220 000 € |
| Câblage inter-racks + fibre | lot | 35 000 € |
| Froid : 8× RDHx + CDU + boucle | §155 adapté | 220 000 € |
| Onduleurs 2× 80 kW N+1 + batteries 15 min | — | 90 000 € |
| Groupe électrogène 200 kVA + inverseur | §183 | 55 000 € |
| Sécurité incendie (gaz, détection) | §190 | 40 000 € |
| Supervision (sondes, DCIM) | §188 | 25 000 € |
| Main-d'œuvre intégration (10 %) | — | 250 000 € |
| **CAPEX total** | | **≈ 2 780 000 €** |
| Électricité/an (60 kW × 8 760 × 0,20 × PUE 1,4) | — | **≈ 147 000 €/an** |
| Maintenance/an (8 % CAPEX) | — | ≈ 222 000 €/an |

**Le bâtiment et le génie civil ne sont pas inclus** (dalle, locaux, TGBT :
compter 200–500 k€ selon l'existant — à chiffrer avec un BET).

---

## 182. Plan de salle ASCII — 8 racks

```
+----------------------------------------------------------+
|  ENTREE (porte CF 1h)                                    |
|                                                          |
|  [Rack R1] [Rack R2] [Rack R3] [Rack R4]   ALLÉE FROIDE  |
|                                                          |
|  [Rack R5] [Rack R6] [Rack R7] [Rack R8]   ALLÉE FROIDE  |
|                                                          |
|  +--------+  +--------+                                  |
|  | UPS-1  |  | UPS-2  |   local batteries (ventilé)      |
|  +--------+  +--------+                                  |
|  [TGBT] [Groupe: extérieur, local technique]             |
|  [Baie brassage général] [Console]                       |
+----------------------------------------------------------+
  Allées chaudes entre les rangées (confinement)
  Distance rack/mur : ≥ 1 m (maintenance arrière)
  Chemins de câbles au plafond, boucle d'eau en caniveau
```

Règles : allées froides face aux façades avant, allées chaudes confinées vers
les RDHx/CTA, **1 m libre** derrière chaque rack (sortie d'un serveur 2U en
rails).

---

## 183. Groupe électrogène : dimensionnement

```
P_groupe(kVA) = (P_IT + P_froid_secouru) × 1,5 / 0,8
```

- 60 kW IT + 20 kW froid secouru = 80 kW → 80 × 1,5 / 0,8 = **150 kVA** →
  standard **200 kVA**.
- Prix : ≈ 40 000–60 000 € posé (à vérifier), + inverseur de source 15 k€.
- **Test mensuel en charge** (banc de charge ou délestage réel) : un groupe
  qui ne tourne jamais ne démarre jamais.
- Fioul : 200 L pour 8 h à 75 % de charge (ordre de grandeur — voir fiche
  groupe). Contrat de livraison fioul en astreinte.

---

## 184. Disjoncteurs et TGBT : le dimensionnement

| Circuit | Protection | Courbe | Pourquoi |
|---|---|---|---|
| PDU rack 32 A tri | 32 A tétrapolaire | C ou D | appel PSU |
| Onduleur 80 kW | 160 A | D | fort appel |
| Groupe | 250 A | D | — |
| Prises de maintenance | 16 A | C + 30 mA | personnel |

- **Courbe D** pour les charges à fort appel (onduleurs, groupes, gros
  serveurs) : la courbe C déclenche au démarrage simultané.
- **Sélectivité** : l'aval doit déclencher avant l'amont — étude à faire par
  l'électricien (pas au doigt mouillé).
- Terre : **< 5 Ω** (voir guide onduleurs), contrôle annuel.

---

## 185. TCO on-premise vs cloud — comparatif 3 ans chiffré

Hypothèse : 100 VMs (4 vCPU / 16 Go), 50 To stockage, 3 ans.

| Poste | On-premise (6× Compute M) | Cloud (ordre de grandeur) |
|---|---|---|
| CAPEX | 6 × 57 k€ = 342 k€ | 0 |
| Électricité 3 ans | 6 × 0,95 kW × 8 760 × 0,20 × 3 = 30 k€ | inclus |
| Maintenance 3 ans | 82 k€ | incluse |
| Licences (Proxmox) | 3 k€ | — |
| **Total 3 ans** | **≈ 457 k€** | — |
| Cloud équivalent | — | ≈ 25–35 k€/mois → **900 k€–1,26 M€ / 3 ans** |

**L'on-premise coûte 2–3× moins cher à 3 ans** pour une charge stable et
prévisible. Le cloud gagne quand : charge variable, besoin < 18 mois,
pas d'équipe infra. Refaites ce calcul à chaque projet — c'est le tableau
que la direction comprend.

---

## 186. Niveau Tier : ce que ça veut dire (et coûte)

| Tier (Uptime Institute) | Dispo | Redondance | Surcoût vs Tier I |
|---|---|---|---|
| I (basique) | 99,671 % | N | base |
| II (composants redondants) | 99,741 % | N+1 partiel | +15–25 % |
| III (maintenable sans arrêt) | 99,982 % | N+1, 2 chemins (1 actif) | +40–60 % |
| IV (tolérant aux pannes) | 99,995 % | 2N | +80–120 % |

Une PME vise **Tier II+** : onduleurs N+1, 2 PDU, froid N+1, groupe.
Le Tier III (tout maintenable sans arrêt) double presque le lot électrique —
à réserver aux applis qui le justifient vraiment.

---

## 187. PUE par design : les leviers

| Levier | Gain PUE | Coût |
|---|---|---|
| Allées confinées | −0,1 à −0,2 | faible |
| Free cooling | −0,3 à −0,5 | moyen (conception) |
| Température +3 °C (24→27 °C) | −0,05 à −0,1 | 0 € (ASHRAE le permet) |
| RDHx vs sur-clim | −0,1 | moyen |
| Onduleurs ecolo-mode (VFD) | −0,02 | 0 € (réglage) |

ASHRAE 2021 : **18–27 °C** en entrée serveur, classe A1. Faire tourner à
24 °C au lieu de 21 °C ne tue aucun serveur et économise 10–15 % de froid.
Le 21 °C « parce qu'on a toujours fait comme ça » est une habitude chère.

---

## 188. Supervision et DCIM : le minimum vital

