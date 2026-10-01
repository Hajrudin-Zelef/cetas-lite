---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-6
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [787, 965]
sha256: 41bfea1c0145afe8f1f022ce014cb9aa693c0cd035ac269eb14a43cff837e463
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

**Délai :** le compte est actif en quelques minutes ; certaines fonctions
(téléchargements logiciels) peuvent nécessiter une validation supplémentaire
liée à vos contrats.

## 48. Les rôles et la gestion multi-utilisateurs

Pour une équipe (chef de service + techniciens) :

- Créez **un compte par personne** (traçabilité des tickets et des
  téléchargements).
- Désignez un **administrateur du compte entreprise** qui gère les accès.
- À chaque départ d'un collaborateur : révoquer son accès et transférer ses
  tickets ouverts (ne jamais partager un compte unique : en cas d'audit,
  impossible de savoir qui a fait quoi).

## 49. Enregistrer ses équipements (par numéro de série)

C'est l'étape qui débloque les droits (téléchargements, garantie, TAC) :

1. Dans votre espace : **My Devices / Mes équipements** → **Register**.
2. Saisissez le **numéro de série** (S/N) de l'équipement — celui de
   l'étiquette physique ou de `display device` / `display esn`.
3. Ajoutez : modèle, site d'installation, date de mise en service.
4. Répétez pour **tout le parc** (AP361, AP761, S310, AR720, USG6000...).

**Astuce :** faites-le **à la réception** du matériel (checklist section 114),
pas six mois plus tard quand vous aurez besoin d'un firmware en urgence.

## 50. Le tableau de bord : la vue d'ensemble du parc

Une fois les équipements enregistrés, le portail affiche typiquement :

| Vue | Utilité pour le chef de service |
|---|---|
| Liste des équipements | Inventaire officiel, exportable |
| Statut de garantie par S/N | Repérer les fins de garantie (revue annuelle) |
| Contrats associés (Hi-Care) | Vérifier la couverture de chaque site |
| Alertes (bulletins de sécurité) | Anticiper les patchs critiques |
| Tickets ouverts | Suivi sans appeler la hotline |

## 51. Télécharger firmwares et patchs

1. Recherchez votre modèle (ex. : `USG6680`, `S310-24P`).
2. Onglet **Software** : versions VRP, patchs, notes de version.
3. **Lisez toujours la Release Note avant de télécharger** : prérequis,
   incompatibilités, procédure de mise à jour.
4. Téléchargez via une connexion stable ; vérifiez la somme de contrôle
   (hash) si fournie.

> **Point de vigilance :** le téléchargement de certaines versions logicielles
> exige un contrat de support logiciel actif. Sans contrat : accès refusé
> (cas vécu n°105). C'est contractuel, pas un bug du site.

## 52. Choisir la bonne version : matrice de compatibilité

Avant toute montée de version :

- [ ] Consulter la **matrice de compatibilité** (version VRP ↔ modèle ↔
      fonctionnalités) sur le portail.
- [ ] Vérifier que vos licences restent valides sur la version cible
      (rarement un problème, mais à contrôler).
- [ ] Lire les « upgrade paths » : certaines versions exigent un palier
      intermédiaire (ex. : V500 → V600 puis V700, pas de saut direct).
- [ ] Tester sur un équipement non critique ou en maquette.

## 53. Télécharger la documentation

Le portail donne accès aux guides d'installation, de configuration, de
maintenance et aux références de commandes par modèle et par version.
**Réflexe :** toujours télécharger la doc de **votre version exacte**
(la syntaxe CLI change entre versions VRP).

## 54. HedEx : présentation

**HedEx** (Huawei Electronic Documentation Explorer) est le lecteur de
documentation électronique de Huawei :

- Format propriétaire **.hdx** : documentation structurée, consultable
  hors ligne après téléchargement.
- Alternative/complément au PDF : recherche plein texte, navigation par
  arborescence, index.
- Le lecteur HedEx se télécharge depuis le portail support.

Usage typique : embarquer toute la doc d'une gamme sur le PC portable du
technicien qui part sur un site isolé sans internet.

## 55. HedEx : recherche et navigation efficaces

- Utilisez la **recherche plein texte** avec des mots-clés CLI exacts
  (`display license`, `license active`) plutôt que des phrases.
- Naviguez par **arborescence produit → version → guide** pour éviter les
  docs d'une autre version.
- Mettez en **favoris** les pages utilisées en intervention (procédures
  d'urgence, codes d'erreur).

## 56. Le centre de licences (lien avec les parties B–C)

Depuis le portail support, la rubrique **Licenses** (ou via le portail
dédié, adresse constatée `https://app.huawei.com/isdp` — **à vérifier sur
le portail officiel**) permet de :

- Générer les fichiers `.dat` (sections 20–21).
- Consulter l'historique des générations (quel ESN, quand, quel droit).
- Suivre les consommations d'Entitlements (droits restants sur un lot).
- Initier des transferts / réémissions (section 32).

## 57. Vérifier la garantie d'un équipement par S/N

1. Équipements enregistrés (section 49) → fiche de l'équipement.
2. La fiche affiche : **date de début / fin de garantie**, niveau de service
   (standard ou Hi-Care avec son palier).
3. En cas d'écart avec votre contrat : capture d'écran + ticket au
   partenaire (c'est lui qui fait corriger les données).

**À faire une fois par an** pour tout le parc (checklist section 115).

## 58. La communauté et la base de connaissances

Avant d'ouvrir un ticket TAC (partie E) :

1. Recherchez le message d'erreur exact dans la base de connaissances.
2. Consultez les **cas de maintenance** publiés (souvent des pannes
   identiques déjà résolues).
3. Posez la question sur la **communauté technique** Huawei pour les
   sujets non urgents (S3/S4).

Un ticket ouvert avec « j'ai vérifié la base de connaissances, cas
n°XXX, sans succès » est traité plus vite : vous avez fait la moitié du
diagnostic.

## 59. L'application mobile et les autres canaux

Huawei propose des canaux mobiles (application de support, chatbot
intelligent — noms exacts **à vérifier sur le portail officiel**) permettant
de :

- Suivre les tickets en déplacement.
- Recevoir les notifications de bulletins de sécurité.
- Accéder à la documentation.

Utile pour le chef de service d'astreinte : le suivi d'un ticket S1/S2
depuis le téléphone, un dimanche soir.

## 60. Bonnes pratiques du portail (à afficher dans le bureau)

- [ ] Un compte **par personne**, email professionnel, 2FA activée.
- [ ] Équipements enregistrés **à la réception** (pas « quand on aura le temps »).
- [ ] Droits d'admin du compte entreprise : **2 personnes minimum**
      (chef + adjoint) pour éviter le blocage en cas d'absence.
- [ ] Export trimestriel de la liste des équipements + garanties.
- [ ] Ne jamais partager les identifiants ; révoquer les accès des partants.
- [ ] En cas de changement de partenaire : vérifier que les enregistrements
      et les droits suivent (nouveau partenaire = nouvelles preuves de droit).

---

# E. OUVRIR UN TICKET TAC

## 61. Le TAC : qu'est-ce que c'est, comment y accéder

Le **TAC** (Technical Assistance Center) est le centre d'assistance technique
de Huawei. Il fournit de l'assistance à distance (diagnostic, contournement,
résolution) **24h/24, 7j/7** pour les clients couverts (garantie ou contrat
Hi-Care — voir partie F).

Canaux d'accès constatés (d'après la documentation Hi-Care) :
- **Hotlines TAC** (numéros par région — voir section 117).
- **Email** (adresses régionales, ex. : support Europe).
- **Portail web** : ouverture et suivi des « Service Requests ».
- **Application mobile** de support.

> Le délai de réponse contractuel court **à partir de l'acceptation de votre
> demande par le TAC** jusqu'au premier contact d'un ingénieur — pas à partir
> de l'envoi de votre email. D'où l'intérêt du portail (horodatage fiable).

## 62. Avant d'appeler : la checklist pré-ticket

Un ticket bien préparé se résout 2 à 3 fois plus vite. Avant tout contact :

