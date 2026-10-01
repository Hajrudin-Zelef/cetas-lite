---
id: collect-261001-rattrapage/rattrapage/onduleurs-ups-guide-4
title: "Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/onduleurs_ups_guide.md
source_anchor: ""
source_lines: [580, 780]
sha256: 32190187b7769187c012f49f86557e2250bc3a5e3c1df51195fd4008f3197eef
---

# Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)

### 24.1 Consignation (la règle d'or)
```
1. Séparer (ouvrir TOUS les disjoncteurs : entrée, bypass, sortie, batterie)
2. Condamner (cadenas + étiquette nominative)
3. Vérifier l'absence de tension (VAT, AC **et** DC)
4. Attendre la décharge du bus DC (5–10 min), re-vérifier
5. Travailler, puis déconsigner en sens inverse
```
- **Le bus DC reste chargé après coupure** : c'est le piège n°1.

### 24.2 Batteries — dangers spécifiques
- **Court-circuit** : un string 384 V peut débiter des **kA** →
  outils **isolés**, une seule main si possible, pas de bijoux.
- **Hydrogène** : ventilation, pas de flamme/étincelle à proximité.
- **Acide** : kit de neutralisation + douche oculaire à proximité.
- **Poids** : un bloc 12 V/100 Ah = ~30 kg → manutention à 2.

### 24.3 Arc flash
- Tableau TGBT / armoire UPS : risque d'arc électrique.
- EPI : gants isolants classe adaptée, écran facial, vêtements
  ignifugés pour les interventions sous tension (à éviter au
  maximum).
- Distances de sécurité affichées sur les armoires.

### 24.4 Habilitation électrique
- En Afrique francophone : référez-vous à la norme locale
  (inspirée NF C 18-510 : B1V, B2V, BR, BC…).
- **Seul le personnel habilité** intervient dans les armoires.

---

## 25. Dépannage express

| Symptôme | Vérifier d'abord |
|---|---|
| L'UPS ne démarre pas | Disjoncteurs, séquence de phases, neutre |
| Passe sur batterie sans coupure visible | Réseau hors tolérance (mesurer !) |
| Autonomie très faible | Batteries (impédance, âge, température) |
| Bascule sur bypass intempestive | Surcharge, surchauffe, défaut onduleur |
| Alarme ventilateur | Remplacer le ventilateur |
| Ne communique plus (NMC) | IP, câble, mot de passe, reboot carte |
| Bruit anormal | Ventilateur, transformateur (charge ?) |
| Odeur d'œuf pourri | **Batterie en emballement → ventiler, consigner, remplacer** |
| Disjoncte à l'enclenchement | Court-circuit aval, sélectivité |
| Le groupe ne prend pas la charge | Slew rate, dimensionnement (§15) |

---

## 26. Cas pratiques commentés (10 cas terrain)

**Cas 1 — Autonomie de 4 min au lieu de 15 (40 kVA)**
Batteries de 4 ans dans un local à 33 °C, jamais testées. Impédance
+60 %. → Remplacement complet + clim du local + baseline.

**Cas 2 — L'UPS passe sur batterie tous les soirs à 19h**
Creux de tension du réseau (pointe du soir). → Mesure avec
enregistreur, signalement au distributeur, élargir les seuils
d'entrée si possible.

**Cas 3 — Casse onduleur après orage**
Pas de parafoudre en tête. → Remplacement carte + **parafoudre
Type 1+2** + vérification de la terre.

**Cas 4 — Le groupe démarre mais l'UPS reste sur batteries**
Rampe de fréquence trop rapide. → Réglage du régulateur du groupe
(slew rate), test mensuel en charge.

**Cas 5 — Emballement thermique d'un bloc**
Bloc gonflé, odeur, température 55 °C. → **Ventiler, consigner,
remplacer le string**, vérifier la compensation de température du
chargeur.

**Cas 6 — Bypass de maintenance oublié enclenché**
La charge n'est plus protégée depuis 3 mois. → Procédure écrite +
checklist de déconsignation + alarme « en bypass » supervisée.

**Cas 7 — Surcharge à 110 % le lundi matin**
Tout le monde allume en même temps. → Délestage (clim non
prioritaire sur un autre départ), ou montée en puissance.

**Cas 8 — NMC inaccessible après changement d'IP**
Mot de passe perdu, IP en DHCP. → Reset carte (procédure
constructeur), IP fixe, mot de passe en coffre.

**Cas 9 — Condensateur explosé (60 kVA, 6 ans)**
Pas de remplacement préventif. → Kit condensateurs + 2 jours
d'arrêt. Leçon : planifier à 5 ans.

**Cas 10 — Faux « défaut batterie »**
Disjoncteur batterie **vibré ouvert** (mauvais serrage). →
Resserrage au couple + contrôle thermographique trimestriel.

---

## 27. Contrats de maintenance — côté chef de service

### 27.1 Les niveaux de contrat
- **Bronze** : 1 visite annuelle + astreinte.
- **Silver** : 2 visites + supervision + pièces (hors batteries).
- **Gold** : 4 visites + supervision 24/7 + pièces + batteries +
  SLA 4 h.
- Adaptez au **risque client** : un hôpital = Gold, un bureau =
  Silver.

### 27.2 Chiffrage
```
Coût annuel = visites × (MO + déplacement)
            + pièces prévisionnelles (batteries/5 ans, ventilos/4 ans, condos/5 ans)
            + supervision
            + marge (20–30 %) + risque
```
- **Les batteries sont le poste n°1** : provisionnez-les dès la
  signature (sinon à 4 ans vous n'avez plus de marge).

### 27.3 Le SLA
- GTR (temps de rétablissement) : 4 h (Gold), 8 h (Silver),
  24 h (Bronze).
- Astreinte : planning, téléphone dédié, **pièces critiques en
  stock** (sinon le SLA est un mensonge).
- Pénalités : ne les acceptez que si vous maîtrisez le stock.

---

## 28. Gestion d'équipe et planning

- **1 chef + 2 techniciens** minimum pour un parc de 20–30 UPS :
  un senior (diagnostic, HT), un junior (visites, relevés).
- Planning : visites mensuelles/trimestrielles planifiées à
  l'année, astreinte tournante.
- **Fiche d'intervention** obligatoire (modèle §30) : sans écrit,
  pas de traçabilité, pas de garantie.
- Réunion hebdo 30 min : revue des alarmes, planning, stock.

---

## 29. Stock pièces — le minimum vital

| Pièce | Stock conseillé |
|---|---|
| Blocs batterie 12 V (le modèle suivi) | 1 string d'avance |
| Ventilateurs (par modèle) | 2 par modèle |
| Kit condensateurs bus DC | 1 par modèle critique |
| Disjoncteurs (entrée/sortie/batterie) | 1 de chaque calibre |
| Carte NMC de secours | 1 |
| Fusibles DC | jeu complet |
| Câbles, cosses, gaine thermo | assortiment |

- Délai Afrique : **4–8 semaines** → le stock n'est pas une option.

---

## 30. KPI et reporting — piloter le service

- **Disponibilité** : objectif 99,9 % (8 h d'arrêt/an max).
- **MTTR** : temps moyen de réparation (objectif < 4 h Gold).
- **Batteries** : % du parc testé, âge moyen, remplacements prévus.
- **Alarmes** : nombre/semaine, temps de traitement.
- **Rapport mensuel direction** : 1 page — disponibilité, incidents,
  actions, budget, risques (batteries à remplacer, etc.).
- **Rapport annuel par site** : état complet + plan d'investissement
  (batteries, clim, groupe).

### Modèle de fiche d'intervention UPS
```
Date : ___  Site : ___  UPS : ___ (marque/modèle/n° série)
Puissance : ___ kVA  Charge : ___ %  Batteries : ___ V / ___ °C
Alarmes : ___
Actions : ___
Mesures (tensions, impédances) : ___
Pièces : ___
Prochaine échéance : ___
Signatures : technicien ___ / client ___
```

---

## 31. Formation de l'équipe (plan 4 semaines)

- **S1** : sécurité électrique + consignation (§24), topologies (§2–3).
- **S2** : batteries — mesures, impédance, remplacement (§8–10, §20).
- **S3** : installation — disjoncteurs, câbles, mise en service (§11–13).
- **S4** : supervision, alarmes, dépannage (§17–18, §25) + tournées
  accompagnées.

---

## 32. Fin de vie et recyclage

- **Batteries plomb** : filière agréée (le plomb se recycle à 95 %+,
  et il a une valeur — négociez la reprise).
- **UPS** : DEEE — cartes (métaux), condensateurs, transformateur
  (cuivre).
- **Avant mise au rebut** : effacer les configs (mots de passe
  réseau !), récupérer la carte NMC.
- Traçabilité : bordereaux de suivi des déchets (exigence clients
  institutionnels).

---

## 33. Glossaire

