---
id: collect-261001-rattrapage/rattrapage/confidentialite-de-github-copilot-protections-et-guide-de-depannage-1
title: "confidentialite-de-github-copilot-protections-et-guide-de-depannage"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot", "agent", "attribution", "training"]
source: docs/RAG/collect-261001-rattrapage/confidentialite-de-github-copilot-protections-et-guide-de-depannage.md
source_anchor: ""
source_lines: [1, 85]
sha256: 4aeafcabe906649171fbeb3601fe0b811d11067287f1e2bcfc62f1ce17d2e275
---

# confidentialite-de-github-copilot-protections-et-guide-de-depannage

Cours

Dès que vous activez GitHub Copilot pour une équipe, les questions de configuration s'enchaînent : quelles données quittent l'IDE ? Quels dépôts doivent être exclus des suggestions ? Que se passe-t-il lorsqu'une suggestion correspond à du code public ?

Que vous soyez développeur ou administrateur déployant Copilot à l’échelle de l’organisation, maîtriser les aspects de confidentialité et de sécurité fait partie d’une bonne utilisation de l’outil.

Dans cet article, nous passons en revue la gestion de vos données par Copilot et la manière de configurer les paramètres de confidentialité, les exclusions de contenu et des garde-fous comme le filtre de duplication, ainsi que les étapes de dépannage les plus efficaces lorsque quelque chose ne fonctionne plus.

## Comment GitHub Copilot gère vos données

Comme nous l'expliquons dans notre tutoriel GitHub Copilot : bonnes pratiques, Copilot envoie un instantané du contexte de code environnant aux serveurs de GitHub dès que vous tapez dans votre IDE. Le modèle traite ce contexte et renvoie une suggestion : ce sont ces données d'interaction qui alimentent l'expérience.

Pour les utilisateurs des offres individuelles (Free, Pro et Pro+), la déclaration de confidentialité de GitHub autorise l'utilisation de ces données d'interaction pour l'entraînement des modèles. Les utilisateurs peuvent s'y opposer à tout moment dans leurs paramètres de confidentialité personnels. Les offres Business et Enterprise relèvent de conditions contractuelles distinctes excluant totalement l'utilisation des données d'interaction pour l'entraînement — aucune action requise de la part des utilisateurs.

### Ce qui relève des données d'interaction

Le code des dépôts privés stocké au repos n'est pas utilisé. En revanche, les données d'interaction générées lorsque vous utilisez activement Copilot dans un dépôt privé peuvent servir à l'entraînement, sauf si vous vous y opposez. Mais que recouvrent précisément les « données d'interaction » ?

Lorsque vous utilisez Copilot, le système collecte plusieurs types de signaux pour améliorer son assistance :

- **Entrées et invites :** Toutes les commandes ou questions envoyées à Copilot Chat ou au CLI.
- **Sorties :** Les suggestions de code ou réponses textuelles du modèle, y compris votre acceptation ou refus.
- **Contexte de code :** Les extraits entourant la position du curseur et le contenu des fichiers ouverts utilisés pour fournir des suggestions pertinentes.
- **Métadonnées et structure :** Noms de fichiers, structure du dépôt et votre navigation dans l'IDE.
- **Retour utilisateur :** Vos évaluations (pouce levé/baissé) et vos commentaires.

### Différences de traitement selon l'offre

GitHub propose plusieurs niveaux de compte : Free, Pro, Business et Enterprise. Le traitement des données varie selon le type de compte, comme indiqué dans le tableau ci-dessous.

| **Dimension** | **Free / Pro / Pro+** | **Business** | **Enterprise** | 
| Utilisation pour l'entraînement des modèles | Nécessite une opposition | Non. Exclu contractuellement | Non. Exclu contractuellement | 
| Code de dépôt privé au repos | Non utilisé | Non utilisé | Non utilisé | 
| Rétention des invites/sorties | IDE : non conservées. Hors IDE : 28 jours | IDE : non conservées. Hors IDE : 28 jours | IDE : non conservées. Hors IDE : 28 jours | 
| Contrôles admin | Individuels uniquement | Politiques au niveau de l'organisation et gestion des licences | Tous les contrôles Business, plus héritage des politiques à l'échelle de l'entreprise et journaux d'audit | 
| Exclusions de contenu | Non disponible | Disponible au niveau repo et organisation | Disponible à l'échelle de l'entreprise | 
| Protection PI | Non incluse | Oui, avec le filtre de duplication activé | Oui, avec le filtre de duplication activé | 

Outre l'opt-out pour les offres individuelles, il faut retenir que les utilisateurs Business et Enterprise bénéficient des exclusions de contenu, de la protection PI et de contrôles admin au niveau de l'organisation, absents des offres gratuites et individuelles. Pour une comparaison détaillée au-delà de la gestion des données, consultez notre guide GitHub Copilot : offres.

**Quelques précisions utiles :** Les exclusions de contenu ne s'appliquent pas encore au mode Edit, au mode Agent dans les chats IDE, au GitHub Copilot CLI, ni à l'agent cloud. La couverture de protection PI exige que le filtre de détection de duplication soit activé et que la suggestion soit utilisée sans modification.

## Configurer les paramètres de confidentialité de Copilot

La confidentialité des données est devenue un enjeu crucial pour toute entreprise. Voyons les paramètres qui déterminent si vos interactions servent à entraîner les modèles lorsque vous utilisez GitHub Copilot.

### Opt-out pour Free, Pro et Pro+

Pour votre compte individuel, vous pouvez refuser l'utilisation de vos données à des fins d'entraînement en allant dans GitHub Settings et en positionnant « Allow GitHub to use my data for AI model training » sur Disabled, comme illustré ci-dessous.

Cette opposition stoppe les collectes futures et ne réduit pas les fonctionnalités de Copilot. Toutefois, GitHub ne peut pas garantir le retrait des données déjà utilisées pour des entraînements antérieurs ; vos données préalablement collectées peuvent donc subsister dans des jeux d'entraînement existants.

### Politiques d'organisation et d'entreprise

Les utilisateurs Business et Enterprise sont déjà exclus de l'entraînement des modèles, mais les administrateurs doivent tout de même revoir les politiques de partage des données pour contrôler les fonctionnalités Copilot activées dans l'organisation :

- **Settings** de l'organisation >**Copilot** >**Policies** permet de gérer les fonctionnalités, l'attribution des licences et la sélection des modèles pour tous les membres.
- Les politiques au niveau org priment sur les préférences individuelles : tout réglage défini ici s'applique à tous.
- Les propriétaires Enterprise peuvent définir des politiques héritées par plusieurs organisations et auditer l'état depuis un tableau de bord unique.

## Utiliser les exclusions de contenu de Copilot

Vous pouvez empêcher Copilot d'accéder à certains contenus. Depuis les paramètres du dépôt, définissez les éléments que Copilot doit ignorer.

### Fonctionnement des exclusions de contenu

Pour les fichiers exclus :

- Les suggestions inline ne seront pas proposées.
- Leur contenu ne servira pas à générer des suggestions dans d'autres fichiers.
- Leur contenu ne sera pas utilisé par GitHub Copilot Chat pour ses réponses.
- La relecture de code par Copilot ne s'appliquera pas à ces fichiers.

Les exclusions peuvent être configurées par les administrateurs de dépôts, les propriétaires d'organisation et les propriétaires Enterprise.

### Configurer des exclusions au niveau du dépôt et de l'organisation

Au niveau repo, ouvrez Settings > Copilot > Content Exclusion, et indiquez des chemins à l'aide de motifs glob. Exemples courants : `**"**/secrets/**"**` pour exclure tout chemin contenant un répertoire secrets, et `**"*.env"**` pour exclure tous les fichiers d'environnement. 

L'API REST offre une option programmatique si vous gérez des exclusions à grande échelle et souhaitez versionner votre configuration.

Au niveau org, le chemin est : Org Settings > Copilot > Content Exclusion. Les règles définies ici s'appliquent à tous les dépôts de l'organisation.

Les règles au niveau org et repo sont additives : elles s'appliquent simultanément. Les règles au niveau Enterprise priment sur celles au niveau org et repo.

