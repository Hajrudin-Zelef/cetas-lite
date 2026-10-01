---
id: collect-261001-rattrapage/rattrapage/canon-copieurs-guide-6
title: "Canon — Guide ultra-complet copieurs (maintenance au cœur)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-rattrapage/canon_copieurs_guide.md
source_anchor: ""
source_lines: [814, 999]
sha256: 3823c6d073b00c64ee6873ac82723c52eeabf2e2c93bff197edabaabc447ccf2
---

# Canon — Guide ultra-complet copieurs (maintenance au cœur)

- Toujours faire les **calibrations automatiques** (menu utilisateur)
  avant de toucher aux réglages manuels : 80 % des problèmes s'y
  règlent.
- Un déréglage ADJUST sans note = des heures pour retrouver la valeur
  d'origine. Discipline non négociable.

---

## 21. Consommables — achats, stocks, faux toners

### 21.1 Gérer le stock
- Stock minimum : 1 toner de chaque couleur par machine suivie +
  1 kit rollers + 1 bac usagé. Les pièces lourdes (fusion, tambour)
  se commandent au compteur PARTS (anticipez à 80 % de vie).
- **FIFO** : premier entré, premier sorti (le toner vieillit).
- Traçabilité : n° de lot, date d'achat (litiges garantie).

### 21.2 Repérer un faux toner Canon
- Prix < 60 % du tarif officiel = suspect.
- Hologramme/QR : vérifiable (Canon a un système de vérification).
- Poids : les faux sont souvent plus légers (moins de toner).
- Symptômes après montage : fond gris rapide, fuite de toner,
  usure accélérée du développeur.
- **En maintenance pro : n'installez jamais de compatible**
  (responsabilité + image).

---

## 22. Contrats de maintenance — le business

> C'est votre métier : cette section vaut de l'or.

### 22.1 Les 3 modèles de contrat

| Modèle | Principe | Marge | Risque |
|---|---|---|---|
| **Coût-page** | X FCFA/page N&B, Y FCFA/page couleur, tout inclus (toner, pièces, MO) | Élevée si bien calculé | Sous-évaluation du coût réel |
| **Forfait + consommables** | Forfait mensuel MO + pièces/toner facturés | Moyenne | Client freine les remplacements |
| **Tout inclus au forfait** | Mensualité fixe tout compris | Prévisible | Dérive si volume explose |

### 22.2 Calculer un coût-page rentable
```
Coût-page = (toner/page + pièces/page + MO/visites + marge) × 1,2
```
- **Toner/page** : prix cartouche ÷ rendement (à 5 % de couverture ;
  en Afrique, la couverture réelle est souvent 8–12 % → ajustez !).
- **Pièces/page** : (prix kit ÷ durée de vie) par pièce.
- **MO** : visites préventives/an × coût visite ÷ volume annuel.
- **La couverture** : demandez au client des exemples d'impressions
  types. Un client « bureautique » à 10 % de couverture use 2× plus
  de toner que le calcul à 5 %.

### 22.3 Le SLA (engagement de service)
- Définissez : délai d'intervention (ex. 4h ouvrées), délai de
  résolution, machine de prêt si > 48h, pénalités.
- **Ne promettez que ce que vous pouvez tenir** : un SLA non tenu
  coûte plus cher qu'un SLA modeste tenu.
- Exclusions : papier non conforme, toner non d'origine, vandalisme,
  coupures électriques sans onduleur.

### 22.4 Fidéliser
- Rapport de visite **écrit** à chaque passage (compteurs, actions,
  préconisations) → le client voit le travail.
- Alerte proactive : « votre tambour est à 85 %, prévoyons le
  remplacement le mois prochain » → pas de panne = client heureux.
- Formation utilisateurs (15 min) : divise les appels « bourrage »
  par 2.

---

## 23. Diagnostic avancé

### 23.1 Au multimètre
- **Lampe de fusion** : continuité (hors tension, connecteur débranché).
  Infini = lampe coupée.
- **Thermistance** : résistance variable avec la température
  (comparez à la doc ; valeur aberrante = HS).
- **Moteur** : résistance d'enroulement (infini = coupé, 0 = court-circuit).
- **Alimentations** : 24 V / 5 V / 3,3 V présents sur la carte DC ?
  (machine sous tension, **prudence**.)

### 23.2 Lire les compteurs d'usure comme un pro
- COUNTER > PARTS : ne regardez pas seulement le % — regardez la
  **vitesse d'usure** (photos à 3 mois d'intervalle). Une pièce qui
  s'use 2× plus vite que la théorie = cause externe (papier,
  environnement, toner).
- Corrélez : fond gris + développeur à 90 % + toner compatible =
  diagnostic en 2 minutes.

### 23.3 La méthode des 5 pourquoi (exemple)
> Bourrages 0108 récurrents.
> Pourquoi ? → Rouleaux lisses. Pourquoi ? → Poussière + papier bas
> de gamme. Pourquoi ? → Client achète le moins cher. Pourquoi ? →
> Personne ne lui a expliqué le lien. → **Action** : former le client
> + imposer un papier validé au contrat. Problème réglé durablement.

---

## 24. Cas pratiques commentés (10 cas terrain)

**Cas 1 — E000 un lundi matin après un week-end orageux**
Micro-coupures → surtension → thermistance affolée. Laisser refroidir,
CLEAR > ERR. Si OK : vendre un onduleur + parafoudre au client
(argument : « ça a failli coûter une carte »).

**Cas 2 — Fond gris 2 semaines après un toner « pas cher »**
Toner compatible. Solution : aspirer le circuit, remettre toner
origine, ajustement DENS. Expliquer au client le coût réel
(développeur à remplacer à 200k au lieu de 400k).

**Cas 3 — 0108 tous les lundis**
Papier chargé le vendredi, humidité du week-end (local non climatisé).
Solution : ne charger que le nécessaire, stocker la rame emballée,
ventiler avant chargement.

**Cas 4 — Couleurs décalées après déménagement**
Choc pendant le transport → ITB décalée. Registration auto
(FUNCTION > ADJUST > REGIST). Si échec : vérifier la courroie.

**Cas 5 — E602 après coupure CIE**
HDD corrompu. CHK-HDD, si échec → remplacement + SST. Vendre
l'onduleur. (Cas ultra-fréquent en Afrique de l'Ouest.)

**Cas 6 — Scan SMB mort après mise à jour Windows**
Passer le copieur en SMBv2/3 (firmware + réglage) ou créer un partage
avec SMBv2 forcé. Alternative : basculer en scan-to-email.

**Cas 7 — Bourrage 0301 + toner qui s'efface**
Fusion froide : thermistance encrassée ou lampe faible. Nettoyer,
tester, remplacer l'unité si compteur > 80 %.

**Cas 8 — Lignes noires verticales toujours au même endroit**
Tambour rayé (agrafe passée dans l'ADF ? corps étranger ?).
Remplacer le tambour, chercher la cause (sinon le neuf se raye aussi).

**Cas 9 — Le client dit « elle imprime pâle » mais les copies sont OK**
Problème **pilote/PC**, pas copieur : vérifier le pilote (UFR II à
jour), les paramètres d'impression (brouillon ?), tester depuis un
autre PC.

**Cas 10 — Agrafeuse qui se bloque chaque semaine**
Agrafes compatibles. Passer aux agrafes Canon d'origine + nettoyer
le mécanisme. Coût des compatibles < coût de vos déplacements.

---

## 25. Fiches modèles — les plus courants en Afrique

### 25.1 imageRUNNER 2520 / 2530 / 2545 (N&B A3)
- La bête de somme : simple, robuste, pièces disponibles.
- Points faibles : rouleaux pickup (poussière), HDD vieillissant.
- Maintenance : kit rollers à 100k, tambour ~150k.
- Idéal en contrat coût-page (prévisible).

### 25.2 imageRUNNER ADVANCE C5500 series (couleur A3)
- 4 tambours + ITB : maintenance plus fine (registration couleurs).
- Points faibles : développeurs (toner compatible = mort),
  unité de fusion (compteur à surveiller).
- Toujours calibrer après intervention.

### 25.3 imageRUNNER ADVANCE DX C5800 series (génération actuelle)
- Sécurité renforcée (TPM, McAfee) : ne pas désactiver.
- Firmware : maintenir à jour (SMB, failles).
- Même base mécanique que C5500 : vos réflexes restent valables.

---

## 26. Sécurité de l'intervention

- **Électrique** : débrancher avant tout démontage (pas seulement
  éteindre — l'interrupteur ne coupe pas tout). Condensateurs :
  attendre 5 min après débranchement.
- **Fusion** : brûlure à 200 °C — attendre le refroidissement,
  gants thermiques si urgent.
- **Toner** : masque FFP2, ne pas inhaler, ne pas utiliser d'air
  comprimé sur du toner en vrac (nuage explosif en théorie,
  encrassement garanti en pratique → aspirateur à toner).
- **Laser** : classe 1 en fonctionnement normal, mais **ne jamais**
  regarder le faisceau unité ouverte (classe 3B à l'intérieur).
- **Poids** : un copieur A3 = 100–200 kg. Transport à 2 minimum,
  diable, jamais par les trappes.

---

## 27. Former le client — diviser les appels par 2

