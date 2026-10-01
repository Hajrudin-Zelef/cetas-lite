---
id: collect-261001-rattrapage/rattrapage/tutoriel-github-actions-2026-ci-cd-complet-6
title: ".github/dependabot.yml"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Google", "Microsoft"]
dates: []
keywords: ["attention", "aws", "copilot", "gpu", "open source"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-github-actions-2026-ci-cd-complet.md
source_anchor: ""
source_lines: [731, 806]
sha256: 7e618697777b2ff341bcb702741578e81e7bd9a3b29a56f57330139fe8f6b4f7
---

# .github/dependabot.yml

Par conception, les secrets ne sont pas transmis aux workflows déclenchés par des PRs de forks externes. C'est une mesure de sécurité fondamentale. Pour les projets open source, utilisez `pull_request_target` avec prudence ou exécutez les tests nécessitant des secrets uniquement après merge sur une branche protégée.

**Problème 8 : Déploiements bloqués en « Waiting »**

Si un déploiement reste en attente indéfiniment, vérifiez les règles de protection de l'environnement. Un reviewer doit approuver manuellement. Si le reviewer n'est plus disponible, un administrateur du dépôt peut bypasser la protection temporairement. Configurez aussi un timeout d'approbation pour éviter les blocages prolongés.

**Problème 9 : OIDC « Not authorized to perform sts:AssumeRoleWithWebIdentity »**

L'erreur OIDC avec AWS provient généralement d'une politique de confiance mal configurée. Vérifiez que le champ `sub` correspond exactement à `repo:OWNER/REPO:ref:refs/heads/BRANCH`. Les wildcards sont possibles mais doivent être utilisées avec précaution. Vérifiez aussi que l'audience est `sts.amazonaws.com`.

**Problème 10 : Coûts de minutes qui explosent**

Surveillez votre consommation dans Settings → Billing → Actions. Les causes fréquentes : workflows sans concurrency qui s'empilent, tests matriciels trop larges, runners macOS (10x plus chers), builds Docker sans cache. Ajoutez des filtres `paths` pour ne déclencher les workflows que lorsque les fichiers pertinents changent.

## Astuces Avancées pour Maîtriser GitHub Actions en 2026

Voici les techniques avancées que les équipes expérimentées utilisent pour tirer le maximum de GitHub Actions.

### Fuseaux Horaires pour les Workflows Planifiés

Depuis 2026, GitHub Actions supporte les fuseaux horaires IANA pour les planifications cron. Au lieu d'être limité à l'UTC, vous pouvez spécifier `timezone: "Europe/Paris"` pour que vos workflows de nettoyage nocturne s'exécutent effectivement à l'heure française. Cette fonctionnalité simplifie considérablement la gestion des workflows planifiés pour les équipes européennes.

### Workflows Requis pour la Conformité

Disponibles en GA depuis février 2026, les **required workflows** permettent aux administrateurs d'organisation d'imposer des workflows spécifiques sur tous les push et pull requests. C'est idéal pour les scans de sécurité obligatoires, la vérification de licences ou les contrôles de qualité de code. Combinés avec les bonnes pratiques de sécurité GitHub, ils constituent une base solide pour la gouvernance CI/CD.

### Optimisation des Builds Docker Multi-Stage

Combinez le cache GitHub Actions avec les builds Docker multi-stage pour des performances optimales. Utilisez `cache-from: type=gha` et `cache-to: type=gha,mode=max` dans l'action `docker/build-push-action`. Le mode `max` cache toutes les couches intermédiaires, pas seulement la couche finale, ce qui accélère les rebuilds partiels de 40 à 70 %.

Pour les projets monorepo, utilisez la directive `paths` dans les triggers pour ne construire que les services modifiés. Combiné avec la directive de concurrency, cela réduit drastiquement la consommation de minutes.

Selon Kelsey Hightower, ancien Staff Developer Advocate chez Google, « l'automatisation CI/CD n'est pas un luxe, c'est une hygiène de base du développement logiciel. Les équipes qui investissent dans des pipelines robustes livrent 4 fois plus souvent avec 3 fois moins de bugs en production. » Cette observation s'applique particulièrement à GitHub Actions, dont l'intégration native avec GitHub réduit la friction et accélère l'adoption.

Pour comparer les différentes plateformes DevOps avant de faire votre choix, consultez notre comparatif GitLab vs GitHub 2026.

## Tarification GitHub Actions : Comparatif des Plans 2026

Comprendre la tarification de GitHub Actions est essentiel pour optimiser vos coûts CI/CD. Voici un comparatif détaillé des plans disponibles, à jour en août 2026 selon la documentation officielle GitHub : le plan Free inclut désormais 2 000 minutes Actions par mois et 500 Mo de stockage d'artefacts, tandis que GitHub Enterprise Cloud offre 50 000 minutes Actions par mois et 50 Go d'artefacts.

| Plan | Prix/mois | Minutes incluses | Stockage | Runners disponibles | 
|---|---|---|---|---|
| Free | 0 $ | 2 000 min | 500 Mo | Linux, Windows, macOS standard | 
| Pro | 4 $ | 3 000 min | 1 Go | Linux, Windows, macOS standard | 
| Team | 4 $/utilisateur | 10 000 min | 2 Go | Tous (incluant Arm64, larger) | 
| Enterprise | 21 $/utilisateur | 50 000+ min | 50 Go | Tous (incluant GPU) | 
| Dépôts publics | Gratuit | Illimité | Illimité | Tous les runners standard | 

Le point important à noter est le multiplicateur de minutes selon le système d'exploitation. Les minutes Windows comptent double (1 minute = 2 minutes du quota), et les minutes macOS comptent ×10 (1 minute = 10 minutes du quota). Cela signifie que 2 000 minutes gratuites représentent en réalité 200 minutes macOS. Planifiez vos tests multiplateformes en conséquence et privilégiez Linux pour les tâches non spécifiques à un OS.

À noter également : depuis mars 2026, les runners auto-hébergés ne sont plus gratuits. GitHub facture désormais un coût de gestion pour les runners self-hosted, ce qui modifie le calcul économique pour les organisations qui hébergeaient leurs propres runners pour éviter les coûts.

## FAQ : Questions Fréquentes sur GitHub Actions en 2026

**Quelle est la différence entre GitHub Actions et Jenkins ?**

GitHub Actions est une solution CI/CD cloud-native intégrée à GitHub, sans infrastructure à gérer. Jenkins est un serveur CI/CD auto-hébergé open source qui offre plus de flexibilité mais nécessite une maintenance significative. En 2026, GitHub Actions domine le marché grâce à son marketplace de plus de 15 000 actions et son intégration transparente avec l'écosystème GitHub.

**GitHub Actions est-il gratuit ?**

Oui, pour les dépôts publics, GitHub Actions est entièrement gratuit avec des minutes illimitées. Pour les dépôts privés, le plan Free inclut 2 000 minutes par mois. Au-delà, les minutes supplémentaires sont facturées 0,008 $ par minute pour les runners Linux. Attention toutefois : à partir du 1er juin 2026, GitHub Copilot code review commencera lui aussi à consommer des minutes Actions, d'après le changelog officiel de GitHub, ce qui peut accélérer l'atteinte de votre quota mensuel si vous automatisez les revues de code.

**Comment déboguer un workflow GitHub Actions qui échoue ?**

Activez le logging détaillé en ajoutant le secret `ACTIONS_RUNNER_DEBUG` avec la valeur `true` dans votre dépôt. Utilisez `set -euxo pipefail` dans vos scripts shell. Vous pouvez également relancer un job en échec avec le logging de diagnostic activé directement depuis l'interface.

**Peut-on utiliser GitHub Actions avec des dépôts GitLab ou Bitbucket ?**

Non, GitHub Actions est exclusif à GitHub. Pour des alternatives multiplateformes, consultez GitLab CI/CD ou CircleCI. Cependant, vous pouvez déclencher des workflows GitHub Actions via des webhooks externes, ce qui permet une intégration partielle avec d'autres plateformes.

**Comment sécuriser mes secrets dans GitHub Actions ?**

Utilisez les secrets chiffrés à trois niveaux (organisation, dépôt, environnement). Privilégiez l'authentification OIDC pour les providers cloud au lieu des clés statiques. Épinglez les actions tierces par SHA. Définissez des permissions minimales sur le GITHUB_TOKEN. Et ne jamais utiliser `echo` pour afficher des valeurs sensibles dans les logs.

**Quelle est la limite de temps d'exécution d'un workflow ?**

