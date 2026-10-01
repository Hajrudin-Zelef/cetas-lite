---
id: collect-261001-rattrapage/rattrapage/glpi-guide-6
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "incident"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [721, 921]
sha256: a220f2f630d221fd79b88e123570d27557710b6d916001205da7424578af1ab1
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

```
##ticket.id##            → n° du ticket
##ticket.title##          → titre
##ticket.content##        → description
##ticket.url##            → lien direct (dépend de l'URL appli, section 8.3 !)
##ticket.status##         → statut
##user.name##             → destinataire
```

### 14.3 Qui reçoit quoi (matrice minimale)

| Événement | Déclarant | Technicien assigné | Groupe assigné | Superviseur |
|---|---|---|---|---|
| Nouveau ticket | ✅ | ✅ | ✅ | — |
| Nouveau suivi | ✅ | ✅ | — | — |
| Prise en compte | ✅ | — | — | — |
| Solution proposée | ✅ | — | — | — |
| Ticket résolu | ✅ | ✅ | — | — |
| Dépassement SLA (alerte) | — | ✅ | ✅ | ✅ |

### 14.4 File d'attente et collecteur

- Les courriels partent via la file `queuednotification` traitée par le cron
  (délai normal : < 2 minutes).
- Le **collecteur** (`Configuration > Collecteurs`) lit une boîte IMAP et crée
  des tickets depuis les courriels reçus — idéal pour une adresse `sav@entreprise.lan`.

---

# PARTIE III — TICKETS : LE CŒUR DU SAV

## 15. Le ticket : cycle de vie complet

### Les statuts (GLPI 10)

```
NOUVEAU ──▶ EN COURS (Attribué) ──▶ EN COURS (Planifié) ──▶ RÉSOLU ──▶ CLOS
   │               │                       │
   │               ▼                       ▼
   └─────▶ EN ATTENTE ◀────────────────────┘
```

| Statut | Signification | Qui l'utilise |
|---|---|---|
| **Nouveau** | Ticket créé, non qualifié | Déclarant, collecteur mail |
| **En cours (Attribué)** | Assigné à technicien/groupe | Superviseur, technicien |
| **En cours (Planifié)** | Intervention planifiée (date) | Technicien |
| **En attente** | Bloqué (pièce, client, fournisseur) | Technicien (+ motif obligatoire) |
| **Résolu** | Solution proposée, en attente validation | Technicien |
| **Clos** | Validé par le déclarant ou auto-clôture | Système / déclarant |

### Règles de gestion conseillées

1. **Tout ticket naît « Nouveau »** — la hotline le qualifie sous 30 min ouvrées.
2. **« En attente » exige un motif** (liste déroulante : attente pièce, attente
   client, attente fournisseur) + date de relance prévue.
3. **« Résolu » ≠ « Clos »** : le déclarant valide la solution (ou clôture
   automatique après 15 jours sans réponse — paramétrable).
4. Un ticket **ne retourne jamais** à « Nouveau » une fois traité (traçabilité).

### Champs clés d'un ticket

Titre, description, catégorie, urgence/impact/priorité, SLA, demandeur,
observateur(s), assigné à (technicien/groupe/fournisseur), éléments associés
(le copieur concerné !), statut, solution, coût, durée.

---

## 16. Catégories de tickets : la taxonomie du SAV

`Assistance > Catégories de tickets` : arborescence **hiérarchique**. Une bonne
taxonomie = des statistiques exploitables. Exemple pour la maintenance copieurs :

```
Panne matérielle
├── Bourrage papier
├── Qualité d'impression (traces, pâleur, taches)
├── Code erreur affiché
├── Bac / alimentation papier
└── Scanner / chargeur
Consommables
├── Toner à remplacer
├── Kit de maintenance (fusion, tambour)
└── Papier
Demande
├── Installation / déménagement
├── Formation utilisateur
├── Relevé de compteur
└── Devis / commande
Réseau & logiciels
├── Imprimante hors ligne
├── Pilote / file d'impression
└── Scan vers dossier / e-mail
Maintenance préventive
└── Visite planifiée
```

### Règles d'or

- **2 à 3 niveaux maximum** : au-delà, les techniciens choisissent « Autre ».
- Chaque catégorie peut avoir son **gabarit** (section 18) et son **groupe
  assigné par défaut** (règle métier, section 20).
- La catégorie est **obligatoire** à la création (gabarit) : pas de ticket
  « sans catégorie » = pas de statistique fausse.
- Revoyez la taxonomie **une fois par an** avec les techniciens.

---

## 17. Urgence, impact, priorité : la matrice

GLPI calcule la **priorité** à partir de l'**urgence** (ressenti du demandeur)
et de l'**impact** (étendue réelle), selon une matrice paramétrable
(`Configuration > Générale > Assistance`).

### Matrice par défaut (à adapter)

| Urgence \ Impact | 1-Mineur | 2-Moyen | 3-Majeur | 4-Critique |
|---|---|---|---|---|
| **1-Basse** | 1-Très basse | 1-Très basse | 2-Basse | 2-Basse |
| **2-Moyenne** | 2-Basse | 2-Basse | 3-Moyenne | 3-Moyenne |
| **3-Haute** | 3-Moyenne | 3-Moyenne | 4-Haute | 4-Haute |
| **4-Très haute** | 4-Haute | 4-Haute | 5-Très haute | 5-Très haute |

### Définitions métier (exemple SAV copieurs)

| Impact | Définition |
|---|---|
| 4-Critique | Tout un site client à l'arrêt (ex. unique copieur d'une mairie) |
| 3-Majeur | Service entier perturbé (ex. copieur du service comptabilité en clôture) |
| 2-Moyen | Quelques utilisateurs gênés |
| 1-Mineur | Gêne légère, contournement possible |

> **Conseil :** laissez le déclarant fixer l'**urgence**, mais faites valider
> l'**impact** par la hotline. La priorité affichée doit refléter la réalité
> opérationnelle, pas l'émotion du moment.

---

## 18. Gabarits de tickets et champs obligatoires

Les **gabarits** (`Configuration > Gabarits > Gabarits de tickets`) définissent,
**par catégorie**, quels champs sont visibles, obligatoires ou pré-remplis.
C'est l'outil n°1 de la qualité des tickets.

### Gabarit « Intervention copieur » (modèle complet en annexe D)

| Champ | Réglage | Pourquoi |
|---|---|---|
| Catégorie | Pré-remplie, verrouillée | Évite les erreurs |
| Éléments associés | **Obligatoire** | Le copieur concerné, toujours |
| Urgence | Visible, obligatoire | — |
| Type | Pré-rempli « Incident » | — |
| Compteur actuel | Champ personnalisé, obligatoire | Facturation / suivi |
| Code erreur affiché | Champ personnalisé, facultatif | Diagnostic |
| Description | Obligatoire, aide à la saisie | « Décrivez les symptômes... » |

### Champs personnalisés

`Configuration > Champs personnalisés` (plugin ou natif selon version) :
ajoutez « Compteur N&B », « Compteur couleur », « Code erreur », « N° de contrat ».
Ces champs deviennent interrogeables dans les rapports et l'API.

### Aides à la saisie

Chaque champ peut afficher un texte d'aide (« Ex. : bourrage bac 2, code C-0202 »).
Rédigez-les avec les mots **des utilisateurs**, pas le jargon technique.

---

## 19. Tickets récurrents : la maintenance planifiée

`Assistance > Tickets récurrents` : des tickets générés automatiquement selon un
calendrier — l'équivalent GMAO de GLPI. **Parfait pour la maintenance préventive
des copieurs** (le métier !).

### Exemples de tickets récurrents

| Ticket | Périodicité | Assigné à |
|---|---|---|
| Relevé compteurs copieurs — Client Mairie | Mensuel, le 1er | Technicien référent |
| Nettoyage + contrôle copieurs — Clinique | Trimestriel | Groupe Techniciens terrain |
| Remplacement kit de fusion (préventif) | Selon compteur (> 200 000 pages) | Atelier |
| Vérification onduleurs salle serveurs | Semestriel | Groupe Énergies |
| Test groupe électrogène | Mensuel | Groupe Énergies |

### Configuration

1. Créez un **modèle de ticket** (titre, catégorie « Maintenance préventive »,
   gabarit, éléments associés si fixes).
2. `Tickets récurrents` → `+` : choisissez le modèle, la périodicité
   (quotidien, hebdo, mensuel, annuel), la date de début, la **création anticipée**
   (ex. créer le ticket 7 jours avant pour préparer la tournée).
3. Le cron génère les tickets automatiquement.

> **Astuce chef de service :** la liste des tickets récurrents = votre
> **plan de maintenance préventive** officiel. Présentez-la en revue de direction.

---

## 20. Actions automatiques et règles métier sur tickets

