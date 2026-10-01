---
id: collect-261001-rattrapage/rattrapage/energie-solaire-guide-5
title: "Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention", "mai"]
source: docs/RAG/collect-261001-rattrapage/energie_solaire_guide.md
source_anchor: ""
source_lines: [708, 845]
sha256: 1737d6f5d3997ddbcb95e2fd1f03ecd0f510b69255f64895b99078e11f1e468e
---

# Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)

- **Victron VRM**, **Deye Cloud**, **Growatt Shine** : production,
  SoC batteries, alertes — gratuit avec l'onduleur connecté.
- **Installateur multi-sites** : créez un compte pro, ajoutez les
  installations clients → maintenance **prédictive** (vous voyez la
  panne avant l'appel du client = argument commercial massif).
- **Alertes utiles** : production < 70 % de l'attendu, SoC < 20 %,
  température batterie > 40 °C, onduleur en défaut.
- **4G** : un routeur 4G à 30 € + carte SIM suffit pour connecter
  un site isolé (quelques Mo/mois).

---

## 25. Recyclage et fin de vie

- **Panneaux** : 95 % recyclables (verre, alu, silicium). Filières
  DEEE en Europe ; en Afrique, filières émergentes — **ne jamais**
  les abandonner (cadmium sur certaines couches minces).
- **Batteries plomb** : filière plomb très active (les ferrailleurs
  les rachètent — attention aux fonderies sauvages, très polluantes).
- **Batteries lithium** : filière en structuration ; stockez-les en
  lieu sûr en attendant (jamais à la décharge, risque d'incendie).
- Prévoyez une **provision démantèlement** dans les gros projets
  (c'est parfois une obligation légale).

---

## 26. FAQ express

**Le solaire marche-t-il quand il pleut ?**
Oui, à 10–25 % de la puissance (lumière diffuse). D'où les jours
d'autonomie des batteries.

**Faut-il laver les panneaux ?**
Oui en saison sèche/harmattan (poussière = −5 à −15 %). Eau claire,
tôt le matin.

**Combien de temps vit une installation ?**
Panneaux 25–30 ans, onduleur 10–15 ans, batteries LiFePO4 10–15 ans,
structure 25 ans. Prévoyez le remplacement onduleur + batteries une
fois dans la vie de l'installation.

**Peut-on climatiser avec le solaire ?**
Oui, mais c'est le poste le plus gourmand : comptez 1,5 kWc par
clim de 12 000 BTU en usage diurne. Isolation d'abord, clim ensuite.

**Solaire + parafoudre : obligatoire ?**
En zone orageuse (toute l'Afrique de l'Ouest de mai à octobre) : oui,
DC + AC. Un seul impact proche sans protection peut tuer l'onduleur.

**Que faire en cas d'ouragan/tempête ?**
Inclinaison faible (10°) = moins de prise au vent. Au-delà de
150 km/h prévus, aucune structure standard ne garantit rien :
assurance multirisque.

---

## 27. Pompage solaire — dimensionnement détaillé

### 27.1 Calcul de la HMT (hauteur manométrique totale)
```
HMT = hauteur géométrique (niveau nappe → réservoir)
    + pertes de charge (tuyaux, coudes, filtre)
    + pression résiduelle (0,5–1 bar au robinet = 5–10 m)
```
- **Pertes de charge** : comptez ~10–15 % de la hauteur géométrique en
  première approximation (affinez avec l'abaque du tuyau : diamètre,
  longueur, débit).
- **Exemple** : forage à 30 m, château à 6 m, 50 m de tuyau 50 mm →
  HMT ≈ 36 + 5 (pertes) + 5 (pression) ≈ **46 m**.

### 27.2 De l'eau à l'électricité
```
P hydraulique (W) = Q (m³/s) × HMT (m) × 9,81 × 1000
P électrique = P hydraulique / rendement pompe (0,4–0,6)
```
- **Exemple** : 2 m³/h (0,00056 m³/s), HMT 46 m →
  P_hydr = 0,00056 × 46 × 9810 ≈ 252 W →
  P_élec ≈ 252 / 0,5 ≈ **504 W** → pompe 0,75 kW.
- **Champ PV** : P_pompe × 1,3 à 1,5 (le variateur a besoin de marge
  le matin/soir) → **1–1,2 kWc** + variateur solaire 1,5 kW.

### 27.3 Règles du pompage
- **Stockez l'eau, pas l'électricité** : château d'eau = batterie
  gratuite et inusable. Dimensionnez le réservoir à 2–3 jours de
  consommation.
- **Sonde de niveau** bas (marche à sec = mort de la pompe) et haut
  (arrêt quand le château est plein).
- **Démarrage progressif** : le variateur rampe en fréquence —
  indispensable (le pic de démarrage direct grillerait le champ PV).
- **Filtration** : crépine + filtre à sable en entrée, surtout en
  forage sableux.

---

## 28. Auditer une installation existante (checklist)

Utile avant rachat, reprise de maintenance ou diagnostic :

- [ ] **Visuel** : panneaux (fissures, délamination, jaunissement),
  structure (corrosion, desserrage), câbles (rongeurs, UV).
- [ ] **Électrique** : Voc/Isc par string vs fiche technique (±10 %),
  isolement DC au mégohmmètre (> 1 MΩ), serrage des bornes.
- [ ] **Thermographie** : points chauds (MC4, diodes, cellules mortes).
- [ ] **Batteries** : âge, technologie, tensions par élément/bloc
  (écart > 0,2 V = déséquilibre), historique BMS (cycles, défauts).
- [ ] **Production** : comparer 12 mois de monitoring à la théorie
  (P_c × ensoleillement × 0,75). Écart > 20 % = problème.
- [ ] **Documents** : schémas, réglages, garanties, factures —
  une installation sans doc = décote de 20 %.
- [ ] **Sécurité** : parafoudres présents et < 5 ans, terre mesurée
  (< 100 Ω, idéalement < 10 Ω), extincteur présent.

---

## 29. Unités et conversions (pense-bête)

| Grandeur | Unité | Repère |
|---|---|---|
| Puissance instantanée | W / kW | Ce que tire un appareil **maintenant** |
| Énergie | Wh / kWh | Puissance × temps (facture, dimensionnement) |
| Puissance crête | Wc / kWc | Puissance max du panneau (STC) |
| Capacité batterie | Ah | Ah × V = Wh (ex. 200 Ah × 48 V = 9,6 kWh) |
| Irradiation | kWh/m²/jour | 5 à Abidjan ≈ 5 h de « plein soleil » |
| Courant | A | P = U × I (60 A en 48 V = 2 880 W) |

- **Ne confondez jamais W et Wh** : un panneau de 550 Wc ne produit
  pas 550 Wh par heure toute la journée — il produit 550 W **au zénith**,
  soit ~2 kWh sur la journée.
- **1 kWc à Abidjan ≈ 1 400–1 600 kWh/an** (retenez 1 500 pour vos
  devis rapides).

---

*Fin du guide — 800+ lignes. Retenez l'essentiel : dimensionnez sur
la consommation réelle, achetez TOPCon + LiFePO4 + hybride sinusoïdal
pur, protégez (parafoudres, fusibles DC, terre), nettoyez, monitorez.
Le solaire est simple ; c'est la rigueur d'exécution qui fait la
différence entre 5 ans et 25 ans de service.*
