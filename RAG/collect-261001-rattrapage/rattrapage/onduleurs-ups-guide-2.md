---
id: collect-261001-rattrapage/rattrapage/onduleurs-ups-guide-2
title: "Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "sol"]
source: docs/RAG/collect-261001-rattrapage/onduleurs_ups_guide.md
source_anchor: ""
source_lines: [182, 368]
sha256: 082ca0774459673a55da8b3a7537ab23b2d2a80b378751e558f9eb89805447d2
---

# Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)

### 8.3 Les paramètres de charge (VRLA)
- **Floating** : 2,25–2,27 V/élément (13,5–13,6 V/bloc 12 V) à 20 °C.
- **Boost/égalisation** : 2,35–2,40 V/élément (occasionnel).
- **Compensation température** : −3 mV/°C/élément (indispensable
  en climat chaud — sinon surcharge et gonflement).
- **Tension de coupure** : ~1,67 V/élément (ne jamais descendre
  en dessous durablement).
---

## 9. Dimensionnement batteries et autonomie — les calculs

### 9.1 La formule pratique (méthode énergie)
```
Capacité (Ah) = P_charge_W × autonomie_h × 1000
                ─────────────────────────────────
                V_bus_DC × η_onduleur × η_batterie
```
- η_onduleur ≈ 0,95 ; η_batterie ≈ 0,85–0,9.
- **Exemple** : 30 kW, 15 min (0,25 h), bus 384 V :
  Ah = 30 000 × 0,25 / (384 × 0,95 × 0,85) ≈ **24 Ah**
  → 2 strings de 32 blocs 12 V / 12 Ah (ou 1 string de 24 Ah).
- ⚠️ C'est une **estimation** : validez avec les tables de décharge
  à puissance constante du fabricant de batteries ou le
  configurateur constructeur (Schneider UPS Selector, etc.).

### 9.2 Nombre de blocs — les standards
- Bus **384 V** = 32 blocs de 12 V (192 éléments de 2 V).
- Bus **480 V** = 40 blocs de 12 V.
- Plage acceptée par l'UPS : vérifiez (ex. 32–40 blocs) — le nombre
  influence l'autonomie et le courant de décharge par bloc.

### 9.3 Strings en parallèle
- 2 strings = 2× l'autonomie (ou 2× le courant admissible).
- **Chaque string a son disjoncteur** + fusibles.
- Strings **identiques** (même marque, même âge, même capacité) —
  jamais de mélange neuf/vieux.

### 9.4 Autonomies recommandées
- **Avec groupe électrogène** : 10–15 min (le temps du démarrage +
  marge).
- **Sans groupe** : 30–60 min (le temps d'arrêter proprement).
- Data center critique : 15 min + groupe avec contrat carburant.

---

## 10. Armoires batteries et câblage DC

- Armoire **ventilée**, à proximité de l'UPS (câbles DC courts =
  moins de chute de tension).
- Câblage DC : **bipolaire + et −**, section selon courant max
  (ex. 40 kVA/384 V/36 kW → ~110 A → 35 mm² cuivre, à vérifier).
- **Disjoncteur batterie** magnétothermique DC calibré (1,25× le
  courant max de décharge).
- Ordre de raccordement : **d'abord l'armoire** (blocs en série),
  vérifier la tension totale **au multimètre**, **puis** fermer le
  disjoncteur vers l'UPS.
- Étiquetage : + / −, tension totale, date de mise en service,
  « DANGER DC ».

---

## 11. Installation électrique — disjoncteurs, câbles, terre

### 11.1 Les 4 protections d'un UPS triphasé
1. **Disjoncteur entrée redresseur** (amont réseau).
2. **Disjoncteur entrée bypass** (souvent commun ou séparé).
3. **Disjoncteur de sortie** (vers la charge / TGBT).
4. **Disjoncteur batterie** (DC).

### 11.2 Calibres indicatifs (400 V tri, cuivre, à valider NF C 15-100 / CEI 60364)

| UPS | Courant ≈ | Disjoncteur | Câble (ordre de grandeur) |
|---|---|---|---|
| 10 kVA | 16 A | 20–25 A | 4 mm² |
| 20 kVA | 32 A | 40 A | 6–10 mm² |
| 40 kVA | 64 A | 80 A | 16–25 mm² |
| 60 kVA | 96 A | 125 A | 35 mm² |
| 120 kVA | 192 A | 250 A | 70–95 mm² |

- Courant calculé : I = S / (√3 × 400). Ajoutez la **distance**
  (chute de tension < 3–5 %) et les **dératations** (température,
  groupement).
- ⚠️ Ces valeurs sont des **ordres de grandeur** : le calcul
  définitif suit la norme et l'étude du bureau (note de calcul).

### 11.3 Terre et régime de neutre
- **TN-S** recommandé (PE séparé du neutre jusqu'au TGBT).
- Résistance de terre : **< 5 Ω** (mesurée, pas supposée).
- Liaisons équipotentielles : armoire UPS, armoire batteries,
  chemins de câbles, masses métalliques.
- Parafoudre **Type 1+2** en tête d'installation (orages !).

### 11.4 Sélectivité
- Les protections amont/aval doivent être **sélectives** : un
  défaut sur un départ ne doit pas faire tomber l'UPS.
- Courbes et calibres : étude de sélectivité pour 60 kVA et +.

---

## 12. Le local technique — climatisation et accès

- **Température** : 20–25 °C idéal (batteries !). Jamais > 30 °C
  en continu.
- **Climatisation redondée** (N+1) pour 40 kVA et + : une clim en
  panne un week-end = batteries cuites.
- **Ventilation** : extraction d'air chaud de l'UPS (l'avant aspire,
  l'arrière souffle — respecter les distances du manuel, ~50 cm).
- **Hydrogène** : ventilation du local batteries (H₂ < 1 %,
  limite d'explosivité 4 %).
- **Accès** : porte ≥ 80 cm, espace de maintenance devant/derrière
  (manuel), éclairage, extincteur **CO₂** (jamais d'eau sur
  électrique).
- **Sol** : charge admissible (un 120 kVA + batteries = 1–2 tonnes),
  plancher technique ou dalle.

---

## 13. Mise en service (commissioning) — checklists

### 13.1 Checklist commune (toutes puissances)
```
□ Réception : chocs, accessoires, documentation, n° de série
□ Local : clim, ventilation, éclairage, extincteur, accès
□ Câblage : serrage au couple, repérage, terre < 5 Ω
□ Batteries : tension totale, polarité, disjoncteur ouvert
□ Protections : calibres conformes à la note de calcul
□ Parafoudre : présent et raccordé
□ Carte réseau (NMC) : installée, IP fixée
```

### 13.2 Séquence de démarrage
```
1. Fermer disjoncteur bypass → vérifier affichage
2. Fermer disjoncteur entrée → le redresseur démarre
3. Fermer disjoncteur batterie → vérifier floating
4. Démarrer l'onduleur (menu) → vérifier 400 V / 50 Hz en sortie
5. Fermer disjoncteur sortie → charge alimentée
6. TEST COUPURE : ouvrir l'entrée → 0 ms, autonomie OK
7. TEST BYPASS : forcer le bypass → retour onduleur sans coupure
8. Configurer : seuils, alarmes, arrêt programmé, supervision
```

### 13.3 Points spécifiques par puissance
- **10–20 kVA** : souvent en armoire existante → vérifier la
  ventilation et le disjoncteur amont (souvent sous-dimensionné).
- **40 kVA** : local dédié recommandé, note de calcul électrique
  exigée.
- **60–120 kVA** : étude complète (sélectivité, plancher, groupe),
  mise en service **avec le constructeur** si possible, banc de
  charge pour le test (§23).

---

## 14. Bypass — statique et maintenance

- **Bypass statique** : automatique (surcharge > 125 %, défaut
  onduleur). La charge passe sur le réseau **sans coupure**.
  Retour automatique après défaut (selon config).
- **Bypass de maintenance** : manuel, pour intervenir sur l'UPS.
  Procédure : forcer le bypass statique d'abord, **puis** basculer
  le bypass manuel, **puis** isoler l'UPS (entrée + batterie).
- ⚠️ En bypass de maintenance, **aucune protection** : planifiez
  l'intervention (pas pendant un orage !).
- Testez les deux bypass **à chaque visite annuelle**.

---

## 15. Groupes électrogènes et UPS — le couple critique

- **Dimensionnement** : groupe = **1,5 à 2×** la puissance de l'UPS
  (kVA). Un groupe trop juste : fréquence instable → l'UPS refuse
  le réseau et vide les batteries.
- **Problèmes classiques** :
  - Rampe de fréquence trop rapide → régler le slew rate du groupe.
  - THDi élevé du groupe → UPS IGBT moderne = OK ; vieux UPS =
    prévoir.
  - Démarrage à froid : 10–30 s → c'est l'autonomie batterie qui
    couvre (§9.4).
- **Test mensuel** : démarrage du groupe **en charge** (pas à vide),
  avec l'UPS qui bascule dessus. Un groupe qui ne démarre pas =
  discovered pendant la coupure = catastrophe.
- Contrat carburant : niveau mini, rotation du stock (gazole
  vieilli = panne).
---

## 16. Parallèle et redondance N+1

