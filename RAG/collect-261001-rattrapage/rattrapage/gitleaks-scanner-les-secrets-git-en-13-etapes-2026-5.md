---
id: collect-261001-rattrapage/rattrapage/gitleaks-scanner-les-secrets-git-en-13-etapes-2026-5
title: "macOS via Homebrew"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "cyber", "incident", "open source"]
source: docs/RAG/collect-261001-rattrapage/gitleaks-scanner-les-secrets-git-en-13-etapes-2026.md
source_anchor: ""
source_lines: [360, 398]
sha256: 5fd4234bbdd182478aa1e6cdee4b541bdfc43225335f60d9c557b0c5728c50dc
---

# macOS via Homebrew

## Combien coûte réellement la mise en place d'un programme de scan de secrets

Un argument revient souvent pour justifier de repousser ce chantier : le manque de budget ou de temps disponible côté équipe sécurité. Dans le cas de GitLeaks, cet argument tient mal, puisque l'outil lui-même est entièrement gratuit, sans limite de dépôts, de développeurs ou de scans mensuels. Le coût réel se situe ailleurs : le temps d'ingénierie nécessaire pour intégrer l'outil dans les pipelines existants, traiter le stock initial de faux positifs, et surtout gérer la charge de travail générée par un audit initial qui révèle souvent des dizaines de secrets historiques à révoquer.

Pour une équipe de taille moyenne gérant une dizaine de dépôts actifs, l'expérience de terrain montre qu'il faut généralement compter une demi-journée pour l'installation et la configuration initiale décrite dans ce tutoriel, puis entre une et trois journées supplémentaires pour trier et traiter les résultats du premier audit complet de l'historique, selon l'ancienneté et l'activité du dépôt. Cet investissement ponctuel se rentabilise rapidement : le coût moyen d'une intrusion liée à un identifiant cloud compromis, incluant la remédiation, la communication et les éventuelles amendes réglementaires, dépasse largement de simples journées d'ingénieur. Les organisations qui préfèrent une solution clé en main plutôt qu'un déploiement interne peuvent se tourner vers l'offre gratuite de GitGuardian pour la surveillance de leurs dépôts publics, ou vers ses plans payants pour une couverture complète incluant les dépôts privés et les conteneurs, en complément et non en remplacement de GitLeaks en local.

## Questions fréquentes

**GitLeaks est-il vraiment gratuit pour un usage commercial ?**

Oui. GitLeaks est distribué sous licence MIT, sans restriction d'usage commercial, sans limite de nombre de dépôts scannés et sans version payante cachée. Le projet vit de contributions open source, à la différence de GitGuardian qui propose une plateforme SaaS payante en complément.

**GitLeaks peut-il scanner des dépôts privés hébergés en interne ?**

Oui, GitLeaks fonctionne en local sur n'importe quel dépôt Git accessible depuis la machine ou le runner CI, qu'il soit hébergé sur GitHub, GitLab, Bitbucket, ou sur un serveur Git auto-hébergé de type Gitea ou Forgejo.

**Faut-il utiliser gitleaks detect ou gitleaks protect ?**

gitleaks detect analyse l'historique complet des commits d'un dépôt, adapté à un audit ou un scan CI sur une pull request. gitleaks protect ne scanne que les modifications indexées (staged) non encore commitées, ce qui en fait le choix approprié pour un hook pre-commit local.

**Un secret détecté par GitLeaks doit-il toujours être révoqué ?**

Par précaution, oui, dès lors que le secret a été poussé vers un dépôt distant, même privé. Un dépôt privé peut être exposé par une mauvaise configuration de permissions, un clone local compromis ou un ancien collaborateur ayant toujours accès. La révocation systématique élimine ce risque résiduel.

**GitLeaks ralentit-il significativement les pipelines CI/CD ?**

Un scan ciblé sur les commits d'une pull request s'exécute généralement en quelques secondes. Seul le scan complet de l'historique sur de très gros dépôts (plusieurs dizaines de milliers de commits) peut prendre plusieurs minutes, raison pour laquelle il est recommandé de réserver ce scan complet à une tâche planifiée hebdomadaire plutôt qu'à chaque pull request.

**GitLeaks suffit-il pour être conforme à NIS2 ou au Cyber Resilience Act ?**

Non. GitLeaks est un contrôle technique qui contribue à démontrer une gestion des vulnérabilités de la chaîne d'approvisionnement logicielle, mais la conformité NIS2 et Cyber Resilience Act exige une politique documentée plus large, incluant la gestion des accès, la gestion des correctifs et une procédure de notification d'incident.

**Comment tester GitLeaks sans risquer de casser un dépôt de production ?**

Clonez le dépôt cible dans un répertoire temporaire et lancez gitleaks detect en lecture seule. Cette commande n'écrit jamais dans le dépôt, contrairement à git-filter-repo ou BFG Repo-Cleaner, qui doivent impérativement être testés sur un clone jetable avant toute application sur le dépôt réel.

**Quelle est la différence entre une règle regex et la détection par entropie dans GitLeaks ?**

Une règle regex cible un format de secret connu et documenté, comme le préfixe caractéristique d'une clé AWS. La détection par entropie, elle, repère des chaînes de caractères statistiquement aléatoires sans connaître leur format à l'avance, ce qui permet de détecter des secrets génériques ou propriétaires qu'aucune règle prédéfinie ne couvre.
