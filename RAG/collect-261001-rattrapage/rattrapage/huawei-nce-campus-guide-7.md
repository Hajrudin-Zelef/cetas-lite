---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-7
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "license", "licenses", "training"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [492, 583]
sha256: 5699361d8ba6f5f16d52dd33b8ad3cbfabac25d6855814a5962d0e891d41343a
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

- **Périmètre** : > 500 équipements, dizaines de sites, milliers d'utilisateurs.
- **Déploiement** : grappe NCE dimensionnée avec l'outil officiel, composants d'authentification distribués (jusqu'à 20), CampusInsight pour l'analyse à l'échelle, intégration ITSM (tickets) via API.
- **Gouvernance** : multi-tenant (entité centrale + entités locales), RBAC fin, processus de changement outillé (templates versionnés, fenêtres de maintenance).
- **À ce niveau, NCE n'est plus un outil mais un programme** : chef de projet, intégrateur, formation, runbook d'exploitation. Budgéter les services d'intégration au même ordre de grandeur que les licences (à chiffrer avec le partenaire).

---

# PARTIE 4 — ÉDITIONS ET LICENCES

## 26. Modèle de licence — principe général (device-day)

Le modèle de licence d'iMaster NCE-Campus documenté dans les guides d'usage (MSP Training Manual — License Usage Guide V300R022C00 / V300R024C00) repose sur la notion de **device-day** (équipement-jour) :

- Chaque équipement géré **consomme** des « resource items » de licence au fil du temps (un équipement géré pendant un jour = un device-day).
- Les licences sont **poolées** : un stock commun de device-days dans lequel puisent tous les équipements.
- Il existe une **licence commune** (ex : 30 000 device-days tous types d'équipements confondus dans le scénario MSP documenté), consommée **en priorité**, puis la licence commerciale une fois la commune épuisée.
- La licence commune **ne peut pas** être mise dans un package réallouable (détail du mode MSP).

C'est un modèle de **souscription à la consommation** : on paie pour une capacité de gestion sur une durée, pas pour un droit perpétuel par équipement. Avantage : flexibilité (le pool absorbe les variations de parc). Inconvénient : il faut **piloter** la consommation (voir section 29-30) — un parc qui grandit sans suivi = dépassement.

## 27. Licence plateforme vs licence équipement

Distinguer deux niveaux (logique observable dans la documentation et les offres) :

1. **Licence plateforme / software platform license** : le droit d'utiliser le logiciel NCE-Campus lui-même. Dans le scénario MSP documenté, une licence plateforme d'un an est fournie **gratuitement** (avec services logiciels de management uniquement, **sans garantie/warranty**) — les conditions exactes varient selon le mode commercial : **à vérifier** sur l'offre.
2. **Licences d'équipements gérés (device management licenses)** : la capacité à gérer N équipements, mesurée en device-days, **poolée**, avec **cotermination** possible (alignement des dates de fin — « resource collaboration »).

À cela s'ajoutent les **souscriptions CampusInsight** (section 31) : package de base obligatoire + packages optionnels, tarifés selon le type et la quantité de NE.

## 28. Licences par équipement géré — comment elles se consomment

Règles pratiques de consommation (d'après le License Usage Guide) :

- Tout équipement **enregistré et géré** par le contrôleur consomme des device-days, **quel que soit son type** dans le cas de la licence commune (les licences commerciales peuvent distinguer les types — à vérifier sur l'offre).
- Les équipements **non gérés / non collectés** ne consomment pas (un équipement découvert mais non pris en gestion ne doit pas être compté — vérifier le comportement exact de la version).
- Quand le stock restant passe en **négatif**, une **période de grâce de 30 jours** s'enclenche (section 29).
- **Pas de licence d'essai** fournie aux tenants dans le mode MSP documenté — prévoir la licence dès le pilote.

Conséquence pour le dimensionnement : compter **tous** les équipements qui seront gérés (switches, AP, AR, USG, WAC), ajouter une **marge de croissance** (10-20 %/an selon la dynamique du parc), et aligner la durée de souscription sur la stratégie budgétaire (1 an vs 3 ans — le 3 ans est souvent plus intéressant : à négocier).

## 29. Période de grâce et comportement en cas de dépassement

Quand le compteur de licences passe sous zéro :

- **Période de grâce de 30 jours** : le système continue de fonctionner — c'est le délai pour régulariser (acheter/recharger des licences).
- Au-delà : **à vérifier** le comportement exact de la version (dégradation de fonctions, blocage des nouveaux onboardings, etc.) — ne jamais tester en production : simuler en maquette.

Bonnes pratiques d'un chef de service :

- **Alerte à 80 % et 90 %** de consommation (si la fonction existe dans la version — sinon suivi via API/rapports).
- Revue **trimestrielle** du compteur en comité d'exploitation.
- Prévoir au budget annuel une **enveloppe de croissance** de licences.
- Documenter la procédure d'achat en urgence (qui commande ? quel délai partenaire ?).

Voir aussi le cas pratique 7 (licence dépassée).

## 30. Mode MSP — multi-tenant et licence commune

Le mode **MSP** (Managed Service Provider) permet à un prestataire de gérer plusieurs clients (tenants) depuis un NCE-Campus mutualisé :

- Le MSP achète la licence sur l'**ESDP** (plateforme de distribution logicielle Huawei), importe le fichier de licence dans NCE-Campus en tant qu'administrateur et le valide.
- L'administrateur crée des **packages combinés de resource items** et les **alloue** aux sous-MSP/tenants (mode « réallocation capable » : un sous-MSP peut à son tour allouer à des tenants de niveau inférieur).
- Politique de **suspension** configurable par niveau (un MSP peut suspendre un tenant en défaut).
- Licence commune de 30 000 device-days fournie, consommée en priorité, **non réallouable** en package.

Pour Zelef : le mode MSP ne le concerne que s'il envisage de **fournir** du réseau managé à des entités sœurs/filiales — sinon, c'est le mode entreprise classique (mono-tenant) qui s'applique. Mais comprendre le mode MSP aide à négocier avec un intégrateur qui l'utilise.

## 31. iMaster NCE-CampusInsight — souscriptions associées

CampusInsight se souscrit **en plus** de NCE-Campus (datasheet V100R025C00) :

| Module | Type | Principe de tarif |
|---|---|---|
| Package de base d'analyse réseau intelligente | **Obligatoire** | Selon le type et la quantité de NE |
| Package valeur ajoutée — analyse applicative | Optionnel | Selon le type et la quantité de NE |
| Package valeur ajoutée — optimisation et auto-guérison | Optionnel | Souscription logicielle |
| Package valeur ajoutée — analyse de consommation énergétique | Optionnel | Selon le type et la quantité de NE |
| Package valeur ajoutée — copilote O&M réseau | Optionnel | 1 licence pour 1 ensemble CampusInsight |

Note intéressante pour un chef de service énergies : le package **d'analyse de consommation énergétique** existe — à évaluer s'il apporte une vraie valeur (corrélation conso/réseau) ou s'il fait doublon avec la GTE.

## 32. Licence de 60 devices incluse dans eSight — comparaison de logique

Pour comparer honnêtement les modèles économiques :

- **eSight** : licence plateforme (ex : AppBase-Professional **incluant 60 devices**) + licences incrémentales **perpétuelles** par device (88034GEE), par AP (88034GEF), par ONU (88034GEL), + souscription annuelle SnS (support). Logique : investissement initial puis extensions à l'unité.
- **NCE-Campus** : plateforme (+ éventuellement offerte 1 an en mode MSP) + **souscription device-day** récurrente + CampusInsight en option. Logique : **OPEX récurrent**.

Conséquence : sur 5 ans, comparer le **TCO** (section 121), pas le prix facial de l'année 1. Un NCE « pas cher la première année » peut coûter plus cher qu'eSight sur 5 ans — ou l'inverse si l'automatisation fait économiser des ETP d'exploitation.

## 33. Coût — ordre de grandeur et pourquoi on ne donne pas de prix ici

