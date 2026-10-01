---
id: collect-261001-general-networking/general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes-8
title: "git version 2.47.1"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agents"]
source: docs/RAG/collect-261001-general-networking/tutoriel-gitlab-ci-cd-2026-pipeline-en-12-etapes.md
source_anchor: ""
source_lines: [1018, 1030]
sha256: f2d89d03017a75bc740cf839c42c2417f2dc89d3f9e14a6f36cb4abf57e5b7c0
---

# git version 2.47.1

Oui, les runners GitLab sont disponibles pour Linux, Windows et macOS. Les runners SaaS partagés sont sous Linux uniquement. Pour Windows et macOS, vous devez enregistrer vos propres runners avec les executors Shell ou VirtualBox. Les runners macOS sont indispensables pour les builds iOS avec Xcode.

### Comment migrer de Jenkins vers GitLab CI/CD ?

GitLab propose un outil de migration automatique qui convertit les fichiers Jenkinsfile en .gitlab-ci.yml. Pour une migration manuelle, mappez chaque stage Jenkins à un stage GitLab, convertissez les plugins en jobs natifs et remplacez les agents Jenkins par des runners GitLab. Le processus prend généralement 1 à 3 jours pour un pipeline complexe.

### Quelle est la limite de taille des artefacts ?

Sur GitLab SaaS, la limite par défaut est de 1 Go par artefact. Pour les instances self-managed, cette limite est configurable dans les paramètres d'administration. Les artefacts volumineux (rapports de couverture, builds compilés) doivent avoir un `expire_in` court pour éviter de saturer le stockage.

### Comment sécuriser les pipelines contre les attaques supply chain ?

Activez le Secret Detection pour bloquer les commits contenant des secrets. Utilisez le Dependency Scanning pour détecter les vulnérabilités dans les packages NPM/PyPI. Limitez l'accès aux variables protégées aux branches protégées uniquement. Pour les images Docker, utilisez des images de base signées et scannez-les avec le Container Scanning de GitLab Ultimate.
