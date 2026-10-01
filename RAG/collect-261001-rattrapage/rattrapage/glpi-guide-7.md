---
id: collect-261001-rattrapage/rattrapage/glpi-guide-7
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [922, 1120]
sha256: e0aa1112fdc00c10b7742c535d99fc56abd60048af1425b81758f8eb38594e2b
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

### Règles d'assignation (`Administration > Règles > Règles métier`)

Exemples qui font gagner un temps fou :

| Condition | Action |
|---|---|
| Catégorie = « Bourrage papier » | Assigner au groupe `Hotline N1` |
| Catégorie commence par « Panne matérielle » + entité = Client X | Assigner au technicien référent du client |
| Urgence = Très haute | Assigner au groupe `Superviseurs` + priorité max |
| Demandeur = direction@client.lan | Ajouter observateur = chef de service |
| Titre contient « onduleur » | Catégorie = Énergies / Onduleurs |

### Actions automatiques (déjà vues section 9)

Complétez avec des règles de **clôture automatique** : ticket « Résolu » sans
réponse du déclarant depuis 15 jours → « Clos » (avec notification).

> **Prudence :** testez chaque règle sur une entité de test avant production.
> Une règle mal écrite peut réassigner des centaines de tickets en boucle.

---

## 21. SLA : niveaux de service, escalades, pénalités

### Concepts

- **SLA (Service Level Agreement)** : engagement contractuel (ex. prise en compte
  sous 4 h ouvrées, résolution sous 2 jours ouvrés).
- **TTO (Time To Own)** : délai de **prise en compte** (statut → En cours).
- **TTR (Time To Resolve)** : délai de **résolution** (statut → Résolu).
- **Calendrier** : les délais se calculent en **heures ouvrées** (calendrier
  8h-18h lun-ven, hors jours fériés).

### Configuration pas à pas

1. `Configuration > Calendriers` : créez « Heures ouvrées SAV » (lun–ven 8h–18h),
   ajoutez les jours fériés (plages d'exclusion).
2. `Assistance > SLA` → `+` : nom (« SLA Standard Copieurs »), TTO = 4 h,
   TTR = 48 h, calendrier associé.
3. `Assistance > Niveaux de SLA` : définissez les **niveaux d'escalade** :
   - Niveau 1 : à 50 % du TTO → rappel au technicien ;
   - Niveau 2 : TTO dépassé → escalade au superviseur + notification ;
   - Niveau 3 : TTR dépassé → alerte chef de service.
4. Affectez le SLA aux tickets via **règles métier** (ex. tout ticket de l'entité
   « Client Mairie » → SLA « Collectivités ») ou manuellement.

### Pénalités et pilotage

GLPI calcule automatiquement les **dépassements**. En revue hebdomadaire
(section 64), sortez : % de tickets dans les SLA par technicien, par client,
par catégorie. Si un contrat prévoit des pénalités financières, ces chiffres
sont votre **preuve contractuelle**.

> **Point de vigilance :** le statut « En attente » **suspend** le compteur SLA
> (selon paramétrage). Formez les techniciens : un ticket en attente sans motif
> valable = un SLA artificiellement « tenu » = statistiques faussées.

---

## 22. Suivis, tâches, solutions : bien documenter

### Suivis

Chaque intervention s'écrit en **suivi** (public = visible du déclarant ;
privé = interne équipe). Règles :

- Un suivi **à chaque étape** : diagnostic, action, résultat.
- Style factuel : « Remplacement du kit de fusion (réf. FK-8350). Compteur :
  212 458 pages. Test OK, 20 copies de contrôle. »
- Joignez la **photo** du compteur ou du défaut (pièce jointe du ticket).

### Tâches

Pour découper une intervention (planification, durée, technicien) :
`Tâches` dans le ticket → « Planifier » (date, durée prévue) → le technicien
la retrouve dans son **planning** (`Assistance > Planning`).

### Solutions

À la résolution : choisissez un **type de solution** (liste : « Remplacement
pièce », « Nettoyage », « Paramétrage », « Formation », « Non reproduit »...),
décrivez la solution, liez éventuellement un **article de la base de
connaissances**. Le déclarant peut **approuver ou refuser** (avec motif).

> **Capitalisation :** une solution bien rédigée + article KB lié = la prochaine
> panne identique se résout en 10 minutes par un N1 au lieu d'un déplacement.

---

## 23. Sondages de satisfaction

`Assistance > Sondages` : questionnaire envoyé automatiquement à la résolution
(paramétrable : délai, cible). Questions types :

1. Note globale de l'intervention (1–5).
2. Délai d'intervention satisfaisant ? (Oui/Non)
3. Technicien courtois et compétent ? (1–5)
4. Commentaire libre.

**Exploitation :** taux de réponse, note moyenne par technicien/client.
Une note < 3 déclenche une **alerte au superviseur** (notification sur événement
« sondage négatif »). Présentez la synthèse en revue mensuelle : c'est l'indicateur
« qualité perçue » qui complète les SLA.

---

## 24. Base de connaissances : rédiger des articles utiles

`Assistance > Base de connaissances` : le wiki interne du SAV. **C'est ici que
les guides existants de l'entreprise prennent vie** (voir cas pratique 7,
section 65).

### Structure d'un bon article

```
Titre : [Marque] [Modèle] — Symptôme / Procédure
Ex. : « Kyocera TASKalfa 4054ci — Bourrage récurrent bac 2 (code J-0511) »

1. Symptômes
2. Causes probables (par ordre de fréquence)
3. Diagnostic pas à pas
4. Résolution / procédure
5. Pièces et références (toner, kits)
6. Prévention
7. Liens (manuel constructeur, article lié)
```

### Règles éditoriales

- **Un article = un problème = une solution.** Pas de pavé fourre-tout.
- Catégories miroir des catégories de tickets (section 16) : l'article se
  retrouve facilement depuis le ticket.
- **Cible** : rédigez pour le technicien N1 qui ne connaît pas la machine.
- Chaque article a un **responsable de mise à jour** et une **date de révision**
  (revue annuelle).
- Marquez les articles « FAQ publique » pour les déclarants (ex. « Changer le
  toner vous-même ») vs « Interne » (procédures SAV).

### Workflow de publication

Technicien rédige (brouillon) → Superviseur relit/valide → Publication.
Un article validé peut être **lié aux tickets** (solution type) et proposé
automatiquement lors de la création d'un ticket de même catégorie.

---

# PARTIE IV — LE PARC (CMDB)

## 25. Le parc : vision d'ensemble (CMDB)

`Parc` : l'inventaire. Chaque **type d'objet** a sa fiche : ordinateurs, moniteurs,
imprimantes, matériels réseau, téléphones, onduleurs... et tout objet personnalisé.

### Principes

- **Un objet = une fiche**, avec : état (neuf, en service, en panne, au rebut),
  type, marque, modèle, n° de série, n° d'inventaire, entité, lieu, utilisateur,
  groupe responsable, fournisseur, contrat(s), dates (achat, mise en service,
  garantie), valeur, composants, liaisons, documents, historique, tickets liés.
- Les **liaisons** font la CMDB : « ce PC est branché sur ce switch, utilise cette
  imprimante, appartient à cet utilisateur ».
- L'**historique** de chaque fiche est automatique : qui a modifié quoi, quand.

### États et cycle de vie du matériel

```
Commandé → Livré → En service ──▶ En maintenance ──▶ En service
                            └──▶ Au rebut / Sorti (don, reprise)
```

> **Discipline :** aucun matériel n'entre ou ne sort sans fiche GLPI à jour.
> Le technicien qui installe un copieur crée/met à jour la fiche **le jour même**.

---

## 26. Ordinateurs : fiches, composants, liaisons

Fiche ordinateur : système d'exploitation, processeur, RAM, disques, carte réseau
(adresses MAC/IP), logiciels installés, périphériques liés (moniteur, imprimante),
utilisateur, dernière modification.

**Alimentation automatique :** l'agent FusionInventory (sections 33-36) remonte
tout cela sans saisie manuelle. La saisie manuelle ne concerne que les champs
« gestion » (utilisateur, lieu, inventaire, garantie).

### Champs critiques à renseigner (même avec inventaire auto)

| Champ | Pourquoi |
|---|---|
| N° d'inventaire interne | Étiquette physique ↔ fiche GLPI |
| Utilisateur / Groupe | Responsabilité |
| Lieu (bâtiment, bureau) | Retrouver la machine |
| Statut | En service / en panne / stock |
| Date de fin de garantie | Alertes automatiques (section 37) |

---

## 27. Moniteurs, imprimantes et périphériques

