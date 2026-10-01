---
id: collect-261001-rattrapage/rattrapage/energie-solaire-guide-3
title: "Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/energie_solaire_guide.md
source_anchor: ""
source_lines: [350, 522]
sha256: 5161052f13cb1e2288efe4ca8563500852acd0786b32cd3ba8a8351e4ef48b5b
---

# Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)

- **Nettoyage** : un panneau poussiéreux perd 5–15 %. Eau claire,
  éponge douce, **jamais** de jet haute pression ni de produits abrasifs.
- **Monitoring** : onduleurs connectés (WiFi/4G) → applis (Victron
  VRM, Deye Cloud, etc.). Une chute de production = alerte = on agit
  avant que le client s'en aperçoive.

### 9.2 Dépannage express

| Symptôme | Causes probables | Action |
|---|---|---|
| Production nulle, onduleur éteint | Sectionneur DC ouvert, fusible HS | Vérifier coffret DC, polarités |
| Production faible durable | Poussière, ombre nouvelle (arbre poussé), panneau HS | Nettoyer, relever ombres, mesurer Voc/Isc par string |
| Une string à zéro | Diode bypass HS, MC4 fondu, panneau fissuré | Thermographie / mesure string par string |
| Onduleur en défaut « isolement » | Défaut d'isolement DC (câble abîmé, humidité) | Tester l'isolement au mégohmmètre, sécher/réparer |
| Batteries ne chargent plus | BMS en protection, régulateur mal configuré | Lire le BMS (Bluetooth), vérifier seuils |
| Bascule réseau intempestive | Seuil de tension batterie trop haut, surcharge | Relever les logs, ajuster les seuils |

---

## 10. Économie et rentabilité

### 10.1 Calculer le retour sur investissement
```
Coût total installé (€)
÷ Économies annuelles (€/an)  =  temps de retour (ans)
```
- **Économies** : kWh solaires autoconsommés × prix du kWh réseau
  (+ coût du gazole évité si groupe).
- En Afrique de l'Ouest : kWh réseau ~0,15–0,25 €, gazole groupe
  ~0,35–0,50 €/kWh → **retour typique 3–6 ans** en hybride, 2–4 ans
  si ça remplace un groupe.
- **LCOE** (coût actualisé du kWh solaire sur 25 ans) : aujourd'hui
  ~0,05–0,10 €/kWh en Afrique — imbattable face au réseau et au groupe.

### 10.2 Modèles économiques
- **Achat direct** : le client paie, ROI en 3–6 ans, 20 ans de
  quasi-gratuité ensuite.
- **Leasing / PAYG** (pay-as-you-go) : très développé en Afrique
  (M-KOPA, etc.) — le client paie au mois via mobile money,
  l'installateur garde la propriété jusqu'au solde.
- **Autoconsommation + vente du surplus** : là où le rachat existe
  (compteur bidirectionnel) ; sinon, dimensionnez pour
  l'autoconsommation (le surplus perdu = argent perdu).
- **Centrale villageoise / mini-grid** : modèle opérateur (kiosques,
  abonnements) — le plus impactant en zone rurale.

### 10.3 Financement
- Banques locales : de plus en plus de lignes « énergie verte ».
- Bailleurs (Banque mondiale, AFD, UE) : programmes d'électrification
  rurale — les installateurs certifiés y accèdent via appels d'offres.
- **Attention aux arnaques** : acompte > 50 % avant livraison,
  matériel sans garantie écrite, « panneaux 1000 Wc » à prix cassé
  (physiquement impossible en taille standard) → fuyez.

---

## 11. Réglementation et normes (repères)

- **Électrique** : NFC 15-100 (ou équivalent local) pour la partie AC ;
  **UTE C 15-712** (France) = la référence pour le photovoltaïque
  (dimensionnement, protections, mise à la terre).
- **Raccordement réseau** : déclaration auprès du distributeur
  (CIE en Côte d'Ivoire, etc.), conformité anti-îlotage (VDE 0126 /
  EN 50549 pour les onduleurs).
- **Urbanisme** : déclaration de travaux souvent requise au-delà
  d'une certaine puissance/surface.
- **Environnement** : recyclage des panneaux (filières DEEE) et des
  batteries (filière agréée — **jamais** de plomb à la décharge).
- Les règles varient par pays : vérifiez auprès du ministère de
  l'énergie / de l'agence d'électrification rurale locale.

---

## 12. Pompage solaire (cas particulier très demandé)

- **Principe** : panneaux → **variateur de pompe solaire**
  (onduleur dédié avec MPPT) → pompe immergée AC ou DC. **Pas de
  batteries** : on pompe quand il y a du soleil, on stocke **l'eau**
  (château d'eau), pas l'électricité.
- **Dimensionnement** : débit (m³/h) × HMT (hauteur manométrique
  totale) → puissance pompe → puissance crête = P_pompe × 1,3.
- **Exemple** : besoin 10 m³/jour, HMT 40 m → pompe ~1,5 kW →
  ~2 kWc de panneaux + variateur 2,2 kW.
- **Protection** : marche à sec (sonde de niveau), parafoudre,
  filtration en entrée de pompe.
- C'est le système solaire **le plus rentable** (pas de batteries =
  pas de remplacements).

---

## 13. Erreurs classiques des débutants

1. **Sous-dimensionner les batteries** (« on ajoutera plus tard » —
   on ne mélange jamais des batteries d'âges différents).
2. **Surdimensionner les panneaux sans augmenter les batteries**
   (l'énergie excédentaire est perdue, les batteries restent le
   goulot).
3. **Négliger l'ombrage** (un simple câble qui passe devant…).
4. **Câbles trop fins / trop longs** (chutes de tension = pertes
   silencieuses de 5–10 %).
5. **Pas de parafoudres** en zone orageuse.
6. **Onduleur « modified sine »** pour économiser 100 € → moteurs
   grillés.
7. **Batteries au soleil** dans un coin du local technique.
8. **Aucun monitoring** → panne découverte 3 mois plus tard.
9. **Acheter au prix, pas à la fiche technique** (panneaux
   contrefaits).
10. **Oublier la maintenance dans le devis** (le client doit savoir
    nettoyer ses panneaux).

---

## 14. Glossaire express

| Terme | Signification |
|---|---|
| Wc (watt-crête) | Puissance max d'un panneau en conditions STC |
| STC | 1000 W/m², 25 °C, AM1.5 — conditions de test |
| MPPT | Régulateur qui cherche le point de puissance max |
| Voc / Isc | Tension à vide / courant de court-circuit |
| DoD | Profondeur de décharge (50 % plomb, 80 %+ LiFePO4) |
| BMS | Électronique de protection des batteries lithium |
| AFCI | Détection d'arcs électriques (sécurité DC) |
| LCOE | Coût du kWh sur la durée de vie |
| Autoconsommation | Consommer ce qu'on produit |
| Îlotage | Fonctionnement isolé du réseau (interdit en grid-tie) |
| HMT | Hauteur manométrique totale (pompage) |
| PAYG | Paiement à l'usage via mobile money |

---

---

## 15. Onduleurs hybrides — réglages et modes (le cœur du système)

Les marques dominantes en Afrique : **Deye/Sunsynk, Victron, Growatt,
Must, Voltronic (Axpert)**. Principes communs :

### 15.1 Les 3 modes de fonctionnement
- **Priorité solaire** (Solar first) : le solaire alimente les charges
  et charge les batteries ; le réseau n'intervient qu'en secours.
  → Le mode standard en zone de délestages.
- **Priorité réseau** (Utility first) : le réseau alimente, le solaire
  en appoint. → Utile si le kWh réseau est subventionné et fiable.
- **SBU** (Solar → Battery → Utility) : ordre de priorité explicite.
  C'est le réglage le plus fin : solaire d'abord, batteries ensuite,
  réseau en dernier recours.

### 15.2 Réglages critiques (à configurer à la mise en service)
| Paramètre | Valeur type (48 V LiFePO4) | Pourquoi |
|---|---|---|
| Tension bulk/absorption | 56,8 V | Charge complète sans stress |
| Tension float | 54,4 V | Maintien |
| Coupure basse (back to grid) | 48–49 V (~20–30 % restants) | Protège les batteries |
| Retour solaire (back to battery) | 52–53 V | Évite les oscillations |
| Courant de charge max | Selon parc (0,5C max) | 100 Ah → 50 A max |
| Puissance de sortie max | Bridée si besoin | Évite les surcharges |

- **0,5C** : ne chargez jamais une LiFePO4 à plus de la moitié de sa
  capacité en ampères (100 Ah → 50 A). Au-delà = vieillissement accéléré.
- Activez la **compensation en température** si le capteur est fourni
  (surtout en plomb).

### 15.3 Parallélisme
- La plupart des hybrides 5 kW se mettent en **parallèle** (jusqu'à 6
  unités) pour monter à 30 kW. Câbles de communication entre onduleurs
  (CAN/RS485) **obligatoires**, même longueur de câbles batteries.
- En triphasé : 3 onduleurs (un par phase) — voir §16.

---

## 16. Triphasé et couplage groupe électrogène

