---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-14
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["datacenter", "arr", "distribution", "gpu", "sol"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [2417, 2617]
sha256: 782562ddd6295abbf42df879aace53c9723eec822fd9d41905bd9dfa8675788e
---

# Datacenter Builds — Le guide des BOMs

Repères vérifiés le 27/09/2026 (AFCOM 2026, analyses industrie) : 36 % des
opérateurs ont déployé du liquide, 28 % prévoient sous 12–24 mois ; le DTC
monophasé = 55 % du marché direct-to-chip. **Au-delà de 20 kW/rack, le liquide
doit être étudié ; au-delà de 50 kW, il n'y a plus de débat.**

---

## 150. Calcul thermique simplifié : kW → débit d'air

```
Débit_air (m³/h) = P(kW) × 860 / ΔT(°C) / 1,2
```

Avec ΔT = 12 °C (entrée 24 °C → sortie 36 °C, standard) :
- **1 kW ≈ 300 m³/h** d'air.
- Rack 5,6 kW (§120) → **≈ 1 700 m³/h**.
- Rack 25 kW (Ceph §47) → **≈ 7 500 m³/h** — c'est énorme : 2 ventilos
  industriels. D'où le liquide.

Vérifiez que la **CTA/clim** de la salle fournit ce débit PAR RACK, pas au
total de la salle. L'erreur classique : une clim de 50 kW pour 4 racks de
15 kW = 60 kW → ça ne passe pas.

---

## 151. Rear-door heat exchanger (RDHx) : la porte qui refroidit

- Principe : porte arrière du rack avec batterie à eau — l'air chaud des
  serveurs est refroidi **avant** de sortir dans la salle. La salle reste à
  température normale.
- **Actif** (avec ventilos) ou passif ; capacité 20–50 kW/porte.
- Avantages : pas de modification des serveurs, retrofit possible, pas de
  risque d'eau DANS le serveur.
- Inconvénients : il faut une **boucle d'eau** (CDU ou eau glacée du bâtiment),
  +1 porte = +15 cm de profondeur.
- Prix : 8 000–20 000 €/porte (à vérifier). **Le meilleur rapport €/kW pour
  passer de 15 à 40 kW/rack.**

---

## 152. DLC direct-to-chip : le standard > 100 kW

- Principe : **cold plates** en cuivre sur CPU/GPU, eau glycolée (75 % eau /
  25 % glycol) en circuit fermé, raccords **quick-disconnect** anti-goutte.
- Le rack NVL72 (> 120 kW) ne fonctionne QUE comme ça.
- Il faut : CDU (Coolant Distribution Unit) par rangée, manifold par rack,
  boucle bâtiment (dry cooler ou groupe froid).
- Les composants non refroidis (RAM, disques, alims) restent à l'air → les
  racks DLC gardent souvent un **RDHx en complément** (hybride).
- Fournisseurs : Supermicro (DLC intégré, 3 000+ racks/mois en 2026), Dell
  (XE9780L + rRDHx), Vertiv/Schneider (pods modulaires).

---

## 153. CDU et boucles : l'hydraulique du datacenter

```
[ Dry cooler / groupe froid ] ← boucle bâtiment (eau glacée)
            |
      [ CDU redondante ]  ← échangeur plaques, sépare les boucles
            |
   [ Manifold rack 1 ] [ Manifold rack 2 ] ...  ← boucle techno (eau pure)
            |
   [ Cold plates CPU/GPU ] (quick-disconnect)
```

- **2 boucles séparées** : la boucle bâtiment (sale, traitée) n'entre jamais
  dans les serveurs ; la boucle techno (eau déminéralisée) reste propre.
- Détection de fuite : câble détecteur au sol + vannes automatiques.
  **Testez la détection** (un seau d'eau suffit pour le test).
- Qualité d'eau : conductivité < 5 µS/cm, traitement anti-légionelles sur la
  boucle bâtiment.

---

## 154. Immersion : l'état en 2026

- Principe : serveurs plongés dans un **diélectrique** (mono ou diphasique).
- Le diphasique (ébullition) est freiné par la réglementation **PFAS**
  (reporting EPA dès janvier 2027, bans locaux) → le **monophasé** domine les
  nouveaux projets.
- Usage 2026 : hyperscalers, mining reconverti, quelques HPC. **Pas encore
  pour la PME** : maintenance (sortir un serveur du bain), compatibilité
  composants, coût du fluide.
- À surveiller : si vos racks dépassent 200 kW un jour, l'immersion reviendra
  sur la table (section 179).

---

## 155. Cas chiffré : salle de 100 kW IT

| Poste | Calcul | Résultat |
|---|---|---|
| Charge IT | 4 racks × 25 kW | 100 kW |
| Climatisation (PUE 1,4) | 100 × 0,4 | 40 kW froid |
| Débit d'air équivalent | 100 × 300 m³/h | 30 000 m³/h |
| Solution | 4× RDHx 30 kW + boucle eau | — |
| Budget froid | 4 portes × 15 k€ + CDU 30 k€ + boucle 50 k€ | **≈ 140 k€** |
| Électricité froid/an | 40 kW × 8 760 × 0,20 € | **70 k€/an** |

Le froid coûte **autant que les serveurs** sur 5 ans. C'est pour ça que le
PUE est la métrique reine (voir guide onduleurs/énergie).

---

## 156. Free cooling : le climat comme allié

- En France, **8–10 mois/an** l'air extérieur suffit à refroidir (économiseur).
- Dry cooler + free cooling : PUE 1,15–1,25 atteignable (vs 1,6–2,0 en
  détente directe classique).
- Contrainte : humidité contrôlée (pas d'air extérieur direct sur les
  serveurs — échangeur obligatoire).
- **Dimensionnez le free cooling dès la conception** : l'ajouter après coûte 3×.

---

## 157. Surveillance thermique : les sondes qui sauvent

| Sonde | Où | Seuil d'alerte |
|---|---|---|
| Température entrée rack (bas/milieu/haut) | 3/rack | > 27 °C |
| Température sortie rack | 1/rack | > 40 °C |
| Humidité salle | 2/salle | < 20 % ou > 60 % |
| Débit boucle eau / pression | CDU | variation > 10 % |
| Détection fuite | sol + manifold | contact |

Remontez tout en SNMP vers Zabbix (voir guide zabbix) : une dérive lente de
+1 °C/mois = filtre encrassé ou porte mal fermée.

---

## 158. Pièges terrain — Refroidissement

1. **Clim dimensionnée « au total »** et pas par rack : le rack du fond
   surchauffe pendant que celui de devant gèle.
2. **Allées chaude/froide non respectées** : un serveur monté à l'envers
   recycle son air chaud → emballement thermique local.
3. **RDHx sans boucle d'eau prévue** : la porte arrive, l'eau n'existe pas.
   L'hydraulique se décide AVANT les serveurs.
4. **Eau dans la salle sans détection de fuite** : un raccord qui goutte sur
   un PDU = court-circuit. Câble détecteur obligatoire.
5. **Filtres jamais changés** : −20 % de débit en 6 mois en environnement
   poussiéreux.
6. **Oublier la redondance du froid** : N+1 sur les groupes froids comme sur
   les onduleurs. Une clim en panne en août = arrêt.

---

## 159. Tableau de décision rapide

| Votre rack fait… | Faites… |
|---|---|
| < 10 kW | Air, allées chaude/froide, obturateurs |
| 10–20 kW | Air + confinement d'allée |
| 20–40 kW | **RDHx** |
| 40–100 kW | RDHx + DLC partiel (GPU) |
| > 100 kW | DLC complet + CDU |

---

## 160. Lien avec l'énergie : le PUE

```
PUE = énergie_totale_datacenter / énergie_IT
```

- PUE 2,0 : pour 1 kW de serveurs, 1 kW de froid/pertes → **l'énergie double**.
- PUE 1,3 (free cooling) : +30 % seulement.
- Sur 100 kW IT à 0,20 €/kWh : PUE 2,0 = 350 k€/an ; PUE 1,3 = 228 k€/an.
  **122 k€/an d'écart** — le refroidissement EST un sujet financier.
  (Détail dans le guide onduleurs de Zelef — lien §161.)

---

## 161. Onduleurs : lien avec le guide onduleurs de Zelef (pas de doublon)

Ce guide ne duplique PAS le guide `onduleurs_ups_guide.md` (topologies VFI/VI/VFD,
batteries VRLA/Li-ion, maintenance, contrats — tout y est). Ici : **la méthode
pour dimensionner la chaîne onduleur à partir des BOMs ci-dessus**.

Rappel utile : un onduleur se dimensionne en **kVA et en kW** (kW = kVA × fp,
fp ≈ 0,9–1,0 en 2026). C'est le kW qui compte pour des serveurs à PFC actif.

---

## 162. Méthode : des BOMs à l'onduleur en 5 étapes

1. **Σ P_réaliste** de tous les serveurs + réseau + froid critique
   (tableau §125). Exemple §120 : 5,6 kW IT + 0,6 kW réseau = **6,2 kW**.
2. **× 1,25** (marge + croissance) → 7,75 kW.
3. **+ froid secouru** si la clim doit tenir sur onduleur (souvent non en PME —
   les serveurs tiennent 10 min sans froid, le temps d'un arrêt propre).
4. Choisissez l'onduleur : **P_onduleur_kW ≥ résultat**, en N+1 si critique
   (2 onduleurs en parallèle redondant).
5. **Autonomie** : batteries pour 10–15 min (arrêt propre) ou 1–2 h (tenir
   une coupure) — voir §164.

---

## 163. Cas chiffré : 40 kW IT (salle de 6–8 racks)

