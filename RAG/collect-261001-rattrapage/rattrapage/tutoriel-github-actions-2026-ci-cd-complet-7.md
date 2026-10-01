---
id: collect-261001-rattrapage/rattrapage/tutoriel-github-actions-2026-ci-cd-complet-7
title: ".github/dependabot.yml"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/tutoriel-github-actions-2026-ci-cd-complet.md
source_anchor: ""
source_lines: [807, 819]
sha256: 85b42d168134211ec148352d2bb4c9e2857ea6e281286db011ad7c0a3c1e9527
---

# .github/dependabot.yml

Un job individuel peut s'exécuter pendant 6 heures maximum (360 minutes). Un workflow complet peut durer jusqu'à 35 jours avec les jobs en attente d'approbation. Depuis avril 2026, une nouvelle limite s'applique également au nombre de relances : les reruns d'un même workflow sont désormais plafonnés à 50 par workflow, comme le résume Zenn. Définissez toujours `timeout-minutes` pour éviter les surprises de facturation.

**Comment migrer de Jenkins vers GitHub Actions ?**

GitHub propose un outil officiel, GitHub Actions Importer, qui convertit automatiquement les Jenkinsfiles en workflows YAML. La migration typique prend 2 à 4 semaines pour un projet de taille moyenne, en comptant les tests et la validation.

**Les runners auto-hébergés sont-ils toujours gratuits ?**

Non. Depuis mars 2026, GitHub applique des frais de gestion pour les runners auto-hébergés. Cette modification impacte les organisations qui utilisaient des runners self-hosted pour optimiser les coûts. Évaluez soigneusement le rapport coût-bénéfice entre les runners hébergés par GitHub et les runners auto-hébergés.

### Articles Connexes

Pour approfondir vos connaissances en DevOps et développement, consultez ces ressources complémentaires :
