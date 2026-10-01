---
id: collect-261001-rattrapage/rattrapage/energie-solaire-guide-4
title: "Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sol", "distribution"]
source: docs/RAG/collect-261001-rattrapage/energie_solaire_guide.md
source_anchor: ""
source_lines: [523, 707]
sha256: c2f15fc5b6fa9b31428004fbd29ffa87e48fbf8e80b7f1f578f7dc6a3080c5c1
---

# Énergie solaire — Guide ultra-complet (théorie, matériel, déploiement)

### 16.1 Installation triphasée
- 3 onduleurs hybrides monophasés (un par phase) **ou** 1 onduleur
  triphasé (Deye 12 kW tri, etc.).
- **Équilibrage des phases** : répartissez les charges (jamais tout
  sur une phase). Un déséquilibre > 30 % = disjonctions et mauvais
  rendement.
- Moteurs triphasés (pompes, clim centrales) : **démarrage progressif**
  (variateur) obligatoire — le courant de démarrage direct (6–8× In)
  fait disjoncter l'onduleur.

### 16.2 Couplage avec groupe électrogène
- Le groupe se branche sur l'entrée **GEN/AC-in** de l'onduleur hybride
  (jamais en direct sur les charges avec le solaire sans inverseur).
- **Démarrage automatique (ATS)** : l'onduleur démarre le groupe quand
  batteries < seuil ET pas de soleil (contact sec + démarreur auto).
- Dimensionnez le groupe à **1,5× la puissance de l'onduleur**
  (un groupe sous-chargé s'encrasse, surchargé disjoncte).
- Le groupe charge les batteries via l'onduleur (chargeur intégré) :
  réglez le courant de charge groupe à ~80 % de sa capacité.

---

## 17. Câblage — abaques et calculs

### 17.1 Chute de tension : la formule
```
ΔU (V) = 2 × L (m) × I (A) × ρ / S (mm²)    avec ρ cuivre = 0,0172
% chute = ΔU / U × 100    →  viser < 3 % en DC, < 5 % en AC
```

### 17.2 Abaque DC (aller-retour 20 m, 48 V)
| Courant | Section | Chute |
|---|---|---|
| 20 A | 6 mm² | 2,3 % |
| 40 A | 16 mm² | 1,7 % |
| 60 A | 25 mm² | 1,6 % |
| 80 A | 35 mm² | 1,6 % |

- **Exemple** : parc 48 V, 60 A, 15 m aller → `2×15×60×0,0172/25 =
  1,24 V` → 2,6 % : OK en 25 mm², limite en 16 mm² (4 %).
- En **400 V DC** (string), les chutes sont 8× moindres à puissance
  égale : d'où l'intérêt des strings haute tension.

### 17.3 Protections : calibres
- **Fusible DC par string** : 1,56 × Isc (ex. Isc 14 A → fusible 20 A
  DC 1000 V, type gPV).
- **Sectionneur DC** : calibre ≥ 1,25 × Isc total, **pouvoir de
  coupure DC** réel (un sectionneur AC fond en DC).
- **Disjoncteur AC sortie onduleur** : In = P / 230 V × 1,25.
- **Différentiel 30 mA** type A (ou B si onduleur sans séparation
  galvanique — lisez la doc).

---

## 18. Structures — tenir face au vent

- **Prise au vent** : un panneau de 2,3 m² à 120 km/h subit ~150 kg
  de poussée. Une rangée de 10 panneaux = 1,5 tonne à ancrer.
- **Toiture tôle** : vis tirefond dans les **pannes** (pas dans la
  tôle seule !), joint EPDM, rail alu. Espacement des crochets
  selon doc fabricant (tous les 1–1,2 m).
- **Terrasse / au sol** : triangles acier galvanisé, **lest béton**
  (plots) ou pieux battus. Lest indicatif : 80–120 kg/panneau en
  zone ventée (à faire valider par un BE pour les grandes centrales).
- **Corrosion** : bord de mer (Abidjan !) → aluminium anodisé +
  visserie **inox A4**, jamais d'acier brut. L'air salin tue
  l'acier galvanisé en 3 ans.
- **Dilatation** : laissez du jeu dans les rails (l'alu se dilate
  de ~2 mm/m entre nuit et plein soleil).

---

## 19. Mini-grids villageoises (électrification rurale)

Le modèle qui change la donne en zone non raccordée :

### 19.1 Architecture type (village 200 foyers)
- **Production** : 30–60 kWc au sol + onduleurs triphasés 30–50 kW.
- **Stockage** : 100–200 kWh LiFePO4 (2 jours d'autonomie).
- **Distribution** : réseau BT 400/230 V sur poteaux (quelques km),
  compteurs **prépayés** (STS) par foyer.
- **Tarification** : abonnement + kWh (mobile money), typiquement
  0,30–0,50 €/kWh — cher au kWh mais sans investissement pour le foyer.

### 19.2 Points critiques
- **Foisonnement** : jamais 200 foyers à pleine puissance en même
  temps → dimensionnez sur la pointe **probable** (compteurs
  communicants pour l'apprendre).
- **Vol et vandalisme** : clôture, gardiennage, panneaux boulonnés
  (vis antivol), GPS sur les onduleurs de valeur.
- **Maintenance locale** : formez 2 techniciens villageois (nettoyage,
  lecture des défauts, fusibles). Sans eux, la centrale meurt en 2 ans.
- **Modèle économique** : l'opérateur doit couvrir O&M + renouvellement
  batteries (provisionnez **dès le jour 1** : ~15 % du CA).

---

## 20. Froid solaire et climatisation

- **Chambres froides solaires** (poisson, vaccins, produits agricoles) :
  le froid se stocke **thermiquement** (plaques eutectiques, glace) —
  bien moins cher que des batteries électriques.
- **Climatisation solaire directe** : clims à compresseur DC alimentées
  en direct par les panneaux (sans batteries, elles tournent quand il y
  a du soleil — exactement quand on a chaud). Marques : Gree, Midea
  (gammes « solar hybrid »).
- **Réfrigérateurs DC** (12/24 V, Steca, Sundanzer) : 3–4× moins
  gourmands qu'un frigo AC via onduleur. Pour le médical et l'off-grid,
  c'est le standard.

---

## 21. Eau chaude solaire (thermique — pas photovoltaïque)

- **Chauffe-eau solaire** (capteurs plans + ballon) : 60–80 % d'économie
  sur l'eau chaude, retour en 3–5 ans. Technologie mûre et simple.
- **Ne pas** chauffer l'eau avec des panneaux PV + résistance : le
  rendement est 3–4× inférieur au thermique (le PV fait de
  l'électricité, pas de la chaleur efficacement).
- En Afrique de l'Ouest : un simple **thermosiphon** (ballon au-dessus
  des capteurs, circulation naturelle) suffit — pas de pompe, pas de
  panne.

---

## 22. Devenir installateur — le métier

### 22.1 Outillage indispensable
- Multimètre TRMS + pince ampèremétrique **DC** (AC/DC !).
- Testeur d'isolement (mégohmmètre 1000 V).
- Sertisseuse MC4 + dénudeur solaire.
- Perceuse, visseuse à chocs, scie, niveau laser.
- Harnais + longe (toiture), EPI (gants isolés 1000 V, lunettes).
- Caméra thermique (ou louez-la pour les audits).

### 22.2 Devis type (ce qu'il doit contenir)
1. Bilan de consommation (tableau signé par le client).
2. Descriptif matériel **avec références exactes** (marque, modèle,
   puissance, garanties).
3. Schéma de principe + plan d'implantation.
4. Planning, garantie d'installation (2 ans mini), contrat de
   maintenance proposé.
5. Prix **détaillé** (matériel / main-d'œuvre / transport), pas un
   forfait opaque.
6. Production estimée (kWh/an) et temps de retour calculé.

### 22.3 Certifications et crédibilité
- Formations fabricants (Victron, Deye…) : gratuites ou peu chères,
  elles ouvrent les garanties pro.
- Qualifications type **QualiPV** (France) ou équivalents locaux ;
  en Afrique de l'Ouest, les appels d'offres des agences
  d'électrification rurale exigent des références.
- **Assurance décennale / RC pro** : indispensable dès que vous touchez
  à la toiture d'un client.

---

## 23. Études de cas détaillées

### 23.1 Clinique rurale (froid vaccins + éclairage + petit labo)
- **Besoin** : 8 kWh/j, criticité maximale (chaîne du froid).
- **Kit** : 8 × 550 Wc (4,4 kWc), hybride 5 kW, 2 × LiFePO4 48 V
  200 Ah (19 kWh, 2+ jours d'autonomie), **2 réfrigérateurs DC**
  médicaux en redondance, monitoring 4G avec alertes SMS.
- **Particularité** : double chaîne de froid (si un frigo lâche,
  l'autre prend le relais), groupe en secours auto.

### 23.2 École (salles + bureautique, usage diurne)
- **Besoin** : 12 kWh/j, **80 % le jour** → peu de batteries.
- **Kit** : 10 × 550 Wc (5,5 kWc), onduleur hybride 8 kW,
  1 × LiFePO4 48 V 200 Ah (9,6 kWh, juste le soir).
- **Astuce** : le profil diurne divise par 2 le coût batteries —
  toujours analyser la **courbe de charge**, pas seulement le total.

### 23.3 Maquis / restaurant (froid + sono le soir)
- **Besoin** : 15 kWh/j, pointe le soir (sono, éclairage, congélateurs).
- **Kit** : 12 × 550 Wc (6,6 kWc), hybride 8 kW, 2 × LiFePO4 48 V
  200 Ah (19 kWh).
- **Astuce** : les congélateurs sont des « batteries thermiques » —
  on les surgèle le jour (solaire abondant), ils tiennent la nuit.

---

## 24. Monitoring avancé et supervision multi-sites

