---
id: collect-261001-rattrapage/rattrapage/glpi-guide-13
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [2127, 2292]
sha256: 0030a67b4bab6b216e80ff865c30fe14a5ac004f1e68b856438b52c97e6d7e3a
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

1. **Création** : la secrétaire appelle la hotline → ticket créé avec le gabarit
   (catégorie « Bourrage papier », élément associé = la fiche du copieur,
   compteur N&B relevé au téléphone, code affiché J-0511).
2. **Qualification (N1, < 30 min)** : la hotline consulte la base de connaissances
   → article « J-0511 bac 2 » → guide la secrétaire (retirer le papier, vérifier
   les galets). Échec → assignation au groupe `Techniciens terrain`, priorité
   calculée (urgence haute, impact moyen → priorité haute).
3. **Planification** : le superviseur planifie une tâche le lendemain 9h,
   technicien référent du client.
4. **Intervention** : sur site, diagnostic = galets d'entraînement usés.
   Remplacement (pièce sortie du stock → décrément + coût imputé au contrat).
   Suivi rédigé avec photo du compteur. Article KB lié.
5. **Résolution** : type de solution « Remplacement pièce », description complète.
   Notification à la déclarante.
6. **Validation** : la secrétaire approuve (tout fonctionne) → ticket **Clos**.
   Sondage envoyé → note 5/5.
7. **Capitalisation** : le superviseur vérifie que l'article KB est à jour
   (référence du kit galets ajoutée).

**Ce que le chef de service voit :** temps de prise en compte, respect du SLA,
coût de l'intervention, pièce consommée, satisfaction — tout est tracé.

---

## 61. Cas pratique 3 — planning des techniciens

### Outils

- `Assistance > Planning` : vue semaine par technicien/groupe des **tâches
  planifiées** des tickets + tickets récurrents.
- Chaque intervention planifiée = une **tâche** dans le ticket (date, durée,
  technicien).

### Rituel hebdomadaire (vendredi 16h, 30 min)

1. Ouvrir le planning de la semaine suivante.
2. Placer les tickets récurrents (préventif) en priorité — ce sont des
   engagements contractuels.
3. Répartir les tickets « En cours (Planifié) » selon : compétences (habilitations
   constructeur), secteur géographique (tournées optimisées), charge (équilibrage).
4. Identifier les tickets « En attente » avec relance dépassée → action.
5. Publier le planning (impression PDF ou partage d'écran en réunion).

### Règles d'équipe

- Aucune intervention sans **tâche planifiée** (sauf urgence critique).
- Le technicien pointe le **temps réel** (début/fin) dans la tâche → fiabilité
  des coûts et des SLA.
- Astreinte : un seul technicien d'astreinte, groupe `Astreinte`, téléphone
  dédié, tickets urgents uniquement.

---

## 62. Cas pratique 4 — stock pièces et consommables

### Organisation du magasin

- **Zones** : consommables courants (toners), pièces d'usure (kits, tambours),
  pièces détachées (cartes, moteurs), parc de prêt.
- Chaque référence = modèle de cartouche/consommable dans GLPI avec **seuil
  d'alerte** et **fournisseur habituel**.

### Processus

| Étape | Action GLPI |
|---|---|
| Réception livraison | Entrée de stock (+ bon de livraison en pièce jointe) |
| Intervention | Sortie depuis le ticket (quantité, imputation client/contrat) |
| Seuil atteint | Alerte auto → commande (demande d'achat = ticket interne) |
| Inventaire trimestriel | Comptage physique → régularisation motivée |
| Pièce défectueuse | Retour fournisseur (suivi dans le ticket d'origine) |

### Indicateurs

- **Taux de rupture** : nombre d'interventions retardées faute de pièce
  (objectif : < 2 %).
- **Rotation** : références dormantes depuis > 12 mois → déstockage.
- **Coût pièces par client** : renégociation des contrats déficitaires.

---

## 63. Cas pratique 5 — migration depuis Excel/ancien outil

### Méthode en 5 étapes

1. **Inventaire des données** : lister les fichiers (parc.xls, tickets.mdb...),
   colonnes, qualité (doublons ? séries manquantes ?).
2. **Nettoyage** : normaliser marques/modèles, dédupliquer, compléter les
   n° de série (un passage atelier peut être nécessaire).
3. **Correspondance** : mapper chaque colonne vers un champ GLPI
   (tableau de correspondance écrit et validé).
4. **Import** : `Administration > Maintenance > Import` (CSV) par lots de
   200 lignes, **vérification d'un échantillon après chaque lot**.
5. **Historique** : les vieux tickets se résument en 1 suivi « Historique repris
   de l'ancien système » sur la fiche du matériel — pas d'import ligne à ligne
   (trop coûteux, peu de valeur).

### Pièges

- Importer des doublons → prévoir la passe de déduplication (section 36).
- Vouloir tout importer → ne migrez que l'**actif** (matériels en service,
  contrats en cours, tickets ouverts + 12 mois d'historique).
- Oublier de **former** avant l'import → les utilisateurs remplissent mal
  dès le premier jour.

---

## 64. Cas pratique 6 — revue hebdomadaire du chef de service

**Rituel : lundi 9h, 45 minutes, avec les superviseurs.** Support : dashboard
« Pilotage SAV » (section 40) + rapport PDF.

### Ordre du jour type

1. **File active** (5 min) : tickets ouverts par ancienneté — tout ticket
   > 15 jours doit avoir une explication et un plan.
2. **SLA** (10 min) : % TTO/TTR tenus par client et par technicien ; chaque
   dépassement → cause + action corrective.
3. **Qualité** (10 min) : sondages < 3/5 → analyse ; tickets rouverts
   (résolu → nouveau) → problème de diagnostic ?
4. **Préventif** (5 min) : tickets récurrents en retard, compteurs non relevés.
5. **Stock & contrats** (10 min) : alertes stock, contrats à J-90, budgets à 80 %.
6. **Actions** (5 min) : 3 actions max, un responsable, une échéance —
   notées dans un ticket interne ou un projet GLPI.

> **Modèle de rapport** en annexe E. La régularité du rituel vaut plus que sa
> sophistication : c'est lui qui fait vivre GLPI.

---

## 65. Cas pratique 7 — intégrer les guides existants dans la base de connaissances

L'entreprise possède déjà des guides (copieurs Kyocera, onduleurs, Debian/Ubuntu,
Proxmox, langages...) : en faire des **articles KB exploitables par les
techniciens**.

### Méthode de découpage

Un guide de 1600 lignes ≠ un article KB. Découpez en **articles ciblés** :

| Guide source | Articles KB à créer (exemples) |
|---|---|
| Guide copieurs Kyocera | « U000 — imprimer le rapport », « C6000 — rupture de chauffe : diagnostic », « Remplacer le kit de fusion », « Codes d'accès maintenance par famille » |
| Guide onduleurs | « Dimensionner un onduleur 20 kVA », « Remplacer les batteries : procédure », « Alarmes courantes et conduite à tenir » |
| Guide Proxmox | « Redémarrer proprement un nœud », « Restaurer une VM depuis PBS » |
| Guide Debian/Ubuntu | « Durcir SSH : checklist », « Étendre un volume LVM » |

### Règles d'intégration

1. **Un article = un problème/une procédure** (section 24).
2. Catégorie KB = catégorie de ticket correspondante.
3. Ajoutez en bas de chaque article : « Source : Guide <nom>, section X —
   mise à jour le <date> ».
4. Liez les articles aux **fiches matériels** concernées et aux **gabarits**
   de tickets (suggestion automatique).
5. Nommez un **référent par famille d'articles** (ex. le spécialiste Kyocera
   relit les articles copieurs chaque année).

---

# PARTIE XIII — RÉFÉRENCES RAPIDES

## 66. Pense-bête de poche (une page)

**À imprimer et afficher à l'atelier.**

