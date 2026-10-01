---
id: collect-261001-rattrapage/rattrapage/onduleurs-ups-guide-1
title: "Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "distribution"]
source: docs/RAG/collect-261001-rattrapage/onduleurs_ups_guide.md
source_anchor: ""
source_lines: [1, 181]
sha256: 3b654e9b6f9530c142a4a30d25a1ca2e542045ba2da21c929af595e4049cd06e
---

# Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)

> Tout sur les onduleurs triphasés 10, 20, 40, 60 et 120 kVA :
> topologies, dimensionnement, Easy UPS et concurrents, batteries,
> installation électrique, mise en service, maintenance préventive
> totale, alarmes, sécurité, groupes électrogènes, monitoring,
> management d'équipe et contrats — **angle chef de service
> systèmes & énergies**.
>
> Rédigé à partir des connaissances du domaine (2025-2026). Les
> valeurs exactes (tensions batterie, seuils, codes) varient selon
> les modèles : **validez toujours sur la fiche technique et le
> manuel du modèle installé**.

---

## 1. Le rôle du chef de service systèmes & énergies

- Vous êtes responsable de la **continuité électrique** : un UPS qui
  tombe = data center, bureaux, ateliers à l'arrêt.
- Vos 4 piliers : **dimensionner juste, installer propre, maintenir
  en préventif, superviser en continu**.
- Ce guide couvre les 4 — et donne à votre équipe des procédures
  exécutables, pas de la théorie.

---

## 2. Les 3 topologies d'onduleurs (norme IEC 62040-3)

| Topologie | Principe | Temps de transfert | Usage |
|---|---|---|---|
| **VFD** (offline/standby) | Le réseau alimente la charge ; l'onduleur ne démarre qu'en coupure | 4–10 ms | Petits bureaux, < 2 kVA |
| **VI** (line-interactive) | Régulateur (AVR) + bascule sur batterie en coupure | 2–4 ms | 1–10 kVA, bureautique |
| **VFI** (online double conversion) | La charge est **toujours** alimentée par l'onduleur | **0 ms** | 10–120 kVA et + : data centers, industrie |

- **Toutes les puissances de ce guide (10–120 kVA) sont en VFI
  double conversion.** C'est le seul choix professionnel à ces
  puissances : zéro micro-coupure, tension et fréquence régénérées.
- Classification IEC : `VFI-SS-111` (le code complet figure sur la
  plaque — sachez le lire).

---

## 3. La double conversion en détail — les 4 blocs

```
RÉSEAU ──► [REDRESSEUR] ──► BUS DC ──► [ONDULEUR] ──► CHARGE
                              │                        ▲
                         [BATTERIES]                   │
                              │                   [BYPASS STATIQUE]
                         (secours DC)                      │
                                              RÉSEAU ───────┘
```

### 3.1 Le redresseur (AC → DC)
- Convertit le 400 V AC en DC (~384–480 V selon modèles).
- **IGBT** sur les modèles récents : facteur de puissance d'entrée
  > 0,99, THDi < 3 % → ne pollue pas le réseau, compatible groupes.
- Anciens modèles à thyristors : THDi 25–30 % → prévoir filtres.

### 3.2 Le bus DC
- Le « cœur » : tension continue stabilisée (~384–480 V).
- **Condensateurs électrolytiques** : pièce d'usure n°1 après les
  batteries (5–7 ans en chaud). Surveillance au §21.

### 3.3 L'onduleur (DC → AC)
- Régénère un 400 V / 50 Hz **parfait**, isolé des perturbations
  réseau (creux, surtensions, harmoniques, variations de fréquence).
- Rendement double conversion : **95–96 %**. Mode ECO : jusqu'à
  **99 %** (la charge repasse sur le réseau filtré — à n'utiliser
  que si le réseau est de bonne qualité).

### 3.4 Les batteries
- Branchées sur le bus DC via un **disjoncteur batterie**.
- En régime normal : maintenues en floating. En coupure : elles
  alimentent le bus → **aucune interruption** (0 ms).

### 3.5 Le bypass
- **Bypass statique** (thyristors) : bascule automatique sur le
  réseau en cas de surcharge ou défaut onduleur (sans coupure).
- **Bypass de maintenance** (manuel) : isole totalement l'UPS pour
  intervenir **sans couper la charge**.
- ⚠️ En bypass, la charge n'est **plus protégée** : à n'utiliser
  que le temps de l'intervention.

---

## 4. kVA vs kW — le facteur de puissance

- **kVA** = puissance apparente (ce que l'UPS peut fournir).
- **kW** = puissance réelle consommée = kVA × cos φ.
- Les UPS modernes affichent un **PF de sortie de 0,9 à 1,0**
  (unity) : un 40 kVA / PF 0,9 = **36 kW** utiles.
- **Règle de dimensionnement** : ne jamais dépasser **80 %** de
  charge en nominal → marge pour les pics et l'avenir.
  - Exemple : besoin 28 kW → 28 / 0,9 / 0,8 ≈ 39 kVA → **40 kVA**.

---

## 5. Dimensionnement — méthode et exemples par puissance

### 5.1 La méthode en 5 étapes
1. **Inventaire des charges** (kW + cos φ de chaque équipement).
2. **Somme** → P totale.
3. **kVA = P / PF_sortie_UPS**, puis **÷ 0,8** (marge 80 %).
4. **Arrondir** au standard supérieur (10/20/40/60/120).
5. **Autonomie batterie** : définir (10–15 min standard = le temps
   que le groupe démarre ; 30–60 min si pas de groupe).

### 5.2 Exemples types

| UPS | Charge cible (≈80 %) | Usage typique |
|---|---|---|
| **10 kVA** (~9 kW) | 7 kW | Petite salle serveurs, agence bancaire |
| **20 kVA** (~18 kW) | 14 kW | Salle serveurs PME, petit data center |
| **40 kVA** (~36 kW) | 29 kW | Data center PME, site industriel |
| **60 kVA** (~54 kW) | 43 kW | Data center moyen, hôpital (partiel) |
| **120 kVA** (~108 kW) | 86 kW | Data center, gros site, N+1 de 60 kVA |

### 5.3 Ne pas oublier dans l'inventaire
- **Climatisation** de la salle (souvent le plus gros poste !).
- Courants d'appel (démarrage moteurs, inrush serveurs).
- Croissance **3 ans** : +20–30 % de marge.
- Charges non linéaires (harmoniques) : prévoir le déclassement si
  l'UPS est ancien.

---

## 6. Schneider Electric Easy UPS — la gamme

> La gamme « Easy » = le bon rapport qualité/prix de Schneider pour
> l'Afrique. Vérifiez les fiches exactes — les déclinaisons évoluent.

| Gamme | Puissance | Positionnement |
|---|---|---|
| **Easy UPS 3S** | 10, 15, 20, 30, **40 kVA** (tri/tri 400 V) | Data centers PME, industrie légère |
| **Easy UPS 3M** | **60** à **120 kVA** (tri/tri 400 V) | Sites moyens, N+1 |
| **Easy UPS 3L** | 250–600 kVA | Gros data centers (hors sujet ici) |
| Smart-UPS (APC) | jusqu'à 10 kVA (mono) | Bureaux, petites salles |

- Points forts Easy : double conversion IGBT, rendement ~96 %,
  carte réseau en option, prix contenu, réseau de distribution
  large en Afrique.
- **À vérifier à l'achat** : PF de sortie exact, plage batteries
  acceptée (32–40 blocs ?), carte NMC incluse ou en option,
  garantie et disponibilité pièces locale.

---

## 7. Les concurrents — repères

| Marque | Gamme équivalente | Note |
|---|---|---|
| **Eaton** | 93E / 93PS (10–120 kVA) | Très bon rendement, robuste |
| **Vertiv** (ex-Emerson) | Liebert EXS / ITA2 | Data centers, monitoring avancé |
| **Riello** | Sentryum / Multi Sentry | Bon rapport qualité/prix |
| **Socomec** | MODULYS / DELPHYS | Modulaire (hot-swap) |
| **Huawei** | UPS5000-A (30–120 kVA) | Modulaire, prix agressif |

- En maintenance multi-marques : les **principes sont identiques**
  (bus DC, floating, bypass) — seuls les menus, codes et pièces
  changent.

---

## 8. Batteries — VRLA vs Lithium

### 8.1 VRLA AGM (plomb étanche) — le standard
- Blocs 12 V (6 éléments de 2 V), design life **3–5 ans**
  (standard) ou **10–12 ans** (longue vie, Eurobat 10+).
- **En Afrique (30–35 °C) : divisez par 2** (loi d'Arrhenius :
  +10 °C = vie / 2). D'où l'importance de la clim du local.
- Pas d'entretien d'électrolyte, mais **surveillance obligatoire**.

### 8.2 Lithium-ion (LFP)
- Vie 10–15 ans, 2× plus d'énergie au m³, pas de floating
  permanent, monitoring par BMS intégré.
- Coût initial 2–3× supérieur, mais TCO souvent meilleur.
- ⚠️ Compatibilité : vérifier que l'UPS accepte le Li-ion
  (courbes de charge, communication BMS).

