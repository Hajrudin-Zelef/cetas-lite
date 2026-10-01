---
id: collect-261001-rattrapage/rattrapage/glpi-guide-9
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [1315, 1508]
sha256: 385a7319a4573eae6801c3688f137647d8ee4452a6abc778eaea52fd70005b0b
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

- SNMPv3 (auth + chiffrement) dès que l'équipement le supporte.
- Ne scannez que les plages utiles, la nuit, pour ne pas saturer le réseau.
- Mettez les équipements inventoriés en **verrou** sur les champs critiques
  (n° de série) pour éviter les écrasements.

---

## 36. Règles d'import et de déduplication

`Administration > Inventaire > Règles` : GLPI décide si un inventaire remonté
= **nouvel objet** ou **mise à jour** d'un existant. Critères par type :

| Type | Critère de rapprochement conseillé |
|---|---|
| Ordinateur | N° de série + UUID |
| Imprimante/Copieur | N° de série, puis adresse MAC |
| Matériel réseau | N° de série, puis IP |

- **Verrous** (`Verrous` sur la fiche) : empêchent l'agent d'écraser les champs
  saisis manuellement (lieu, utilisateur, inventaire).
- **Dictionnaires** : normalisent les marques/modèles remontés
  (« HEWLETT-PACKARD » → « HP »).
- Purgez les doublons via `Administration > Maintenance > Vérification` puis
  fusion manuelle si besoin.

> **Erreur classique :** réinstaller un poste sans conserver l'UUID → doublon.
> Procédure : avant réinstallation, noter l'UUID ; après, forcer le rapprochement
> sur le n° de série.

---

# PARTIE VI — CONTRATS, BUDGETS, PROJETS

## 37. Gestion des contrats et des fournisseurs

`Gestion > Contrats` + `Gestion > Fournisseurs` : le suivi contractuel.

### Fiche contrat

| Champ | Exemple |
|---|---|
| Nom | `Maintenance copieurs — Mairie (2026)` |
| Fournisseur | (si sous-traité) ou « Interne » |
| Type | Forfait / Coût par copie / Garantie |
| Dates | Début, fin, **préavis de reconduction** |
| Coût | Annuel, modalités |
| SLA lié | « Collectivités » (section 21) |
| Éléments couverts | Les 12 copieurs (liaison) |
| Documents | Contrat signé (PDF) |
| Alertes | Fin de contrat à J-90 / J-30 |

### Notifications critiques

- **Fin de garantie** des matériels (J-30) → renouvellement ou remplacement.
- **Fin de contrat** (J-90/J-30) → renégociation ; un contrat oublié reconduit
  tacitement = de l'argent perdu.
- **Seuils de consommables** liés au contrat (inclus ou facturés ?).

### Fournisseurs

Fiche par fournisseur : contacts, délais constatés, tickets liés
(« assigné au fournisseur » quand la panne relève de la garantie constructeur).
Le **taux de respect des délais** par fournisseur = argument de renégociation.

---

## 38. Budgets et lignes budgétaires

`Gestion > Budgets` : enveloppes annuelles (ex. « Budget SAV 2026 — 180 k€ »),
découpées en **lignes** (pièces détachées, consommables, sous-traitance,
renouvellement parc).

- Chaque **coût de ticket** (pièces + main-d'œuvre saisie) s'impute sur une ligne.
- Chaque **achat** (fiche matériel : valeur d'achat) aussi.
- Tableau de bord : consommé vs prévu, par ligne, avec alertes à 80 % / 100 %.

> **Rituel du chef de service :** le 5 de chaque mois, 15 minutes sur le budget
> GLPI avant la revue de direction. Aucune surprise en fin d'année.

---

## 39. Projets et tâches : piloter les chantiers

`Outils > Projets` : pour les chantiers (pas les incidents) — ex. « Déploiement
de 20 copieurs — Groupe Scolaire », « Migration du serveur GLPI », « Refonte du
plan de maintenance préventive ».

- **Tâches** : découpage, responsables, dates, dépendances, avancement %.
- **Tickets liés** : les incidents survenus pendant le projet restent tracés.
- **Coûts** : temps passé × taux horaire → coût réel du projet.
- Diagramme de Gantt intégré pour la présentation en réunion.

**Distinction nette à faire respecter :** incident (ticket) vs projet.
« Le copieur est en panne » = ticket. « Remplacer les 12 copieurs du client » =
projet (avec tickets d'installation rattachés).

---

# PARTIE VII — PILOTAGE : RAPPORTS, RECHERCHE, SUPERVISION

## 40. Rapports et statistiques : le tableau de bord du chef de service

`Assistance > Statistiques` + `Outils > Rapports` : le pilotage chiffré.

### Les 10 indicateurs à suivre chaque semaine

| # | Indicateur | Où le trouver |
|---|---|---|
| 1 | Tickets ouverts / résolus / en retard | Statistiques globales |
| 2 | % SLA tenus (TTO/TTR) par client | Statistiques par SLA |
| 3 | Ancienneté du stock (tickets > 15 jours) | Recherche : statut + date |
| 4 | Charge par technicien (tickets en cours) | Statistiques par technicien |
| 5 | Top 10 catégories (panne la plus fréquente) | Statistiques par catégorie |
| 6 | Top 10 matériels (copieur le plus en panne) | Statistiques par élément |
| 7 | Consommation toner par client | Rapport cartouches |
| 8 | Coût SAV par client / contrat | Rapports + budgets |
| 9 | Note satisfaction moyenne | Sondages |
| 10 | Tickets récurrents en retard (préventif) | Recherche tickets récurrents |

### Tableaux de bord personnalisés

`Outils > Tableaux de bord` : assemblez vos indicateurs en widgets
(nombre de tickets, graphiques, listes). Créez un dashboard **« Pilotage SAV »**
affiché par défaut aux superviseurs, et un **« Mon activité »** pour les techniciens.

### Export

Toutes les listes s'exportent en **CSV/PDF** (icône en bas de liste) →
réutilisation dans le rapport mensuel de direction (modèle en annexe E).

---

## 41. Recherche, filtres et vues enregistrées

La **recherche avancée** (`Assistance > Tickets` → « Rechercher ») est d'une
puissance redoutable : critères multiples (ET/OU), tri, colonnes choisies.

### Recherches à enregistrer (marque-pages)

- « Mes tickets en cours » (assigné à = moi, statut ≠ résolu/clos).
- « Tickets en attente depuis > 7 jours » (relance).
- « Copieurs — compteurs non relevés ce mois ».
- « Toners sous le seuil d'alerte ».
- « Contrats se terminant dans < 90 jours ».
- « Tickets sans catégorie / sans élément associé » (contrôle qualité).

`Outils > Notes` : pensez à documenter chaque recherche enregistrée
(à quoi elle sert, qui l'utilise).

> **Contrôle qualité hebdo (15 min) :** lancez « tickets sans élément associé »
> et « tickets en attente sans date de relance ». Zéro résultat = saisie
> disciplinée. Sinon, rappel aux techniciens concernés.

---

# PARTIE VIII — SÉCURITÉ

## 42. Sécurité — comptes et mots de passe

La sécurité de GLPI = la sécurité de **toutes les données du SAV** (coordonnées
clients, contrats, mots de passe parfois notés dans les tickets...).

### Politique des mots de passe

`Configuration > Générale > Sécurité` :

| Paramètre | Valeur conseillée |
|---|---|
| Longueur minimale | 12 caractères |
| Complexité | Majuscules + minuscules + chiffres + spéciaux |
| Durée de validité | 180 jours (comptes locaux) |
| Historique | 5 derniers interdits |
| Verrouillage | 5 échecs → blocage 15 min |
| Expiration | Notification à J-7 (action `sendpasswordexpires`) |

### Règles

- **Comptes nominatifs uniquement** : jamais de compte partagé « sav » ou « admin ».
  Chaque action est tracée par utilisateur (audit).
- Privilégiez l'**authentification AD/LDAP** (section 12) : un départ = désactivation
  dans l'AD = accès GLPI coupé automatiquement.
- **Double authentification (2FA)** : GLPI 10.x supporte le TOTP natif
  (`Préférences > Sécurité`) — imposez-la aux profils Admin et Superviseur.
- Revoyez les comptes **trimestriellement** : désactivez les inactifs depuis 90 jours
  (recherche : dernière connexion).
- **Jamais de mot de passe en clair** dans un ticket, un suivi ou la base de
  connaissances. Utilisez un coffre (KeePass/Vault) et ne stockez que la référence.

---

## 43. Sécurité — durcissement des profils

Reprenez la matrice de la section 11 et appliquez le **moindre privilège** :

