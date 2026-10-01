---
id: collect-261001-rattrapage/rattrapage/gitlab-vs-github-2026-le-comparatif-definitif-5
title: ".gitlab-ci.yml"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot", "open source"]
source: docs/RAG/collect-261001-rattrapage/gitlab-vs-github-2026-le-comparatif-definitif.md
source_anchor: ""
source_lines: [278, 354]
sha256: 2f40ceee431b1e367d8df701ea1051f6ffecf1609ccbb8c27ce09cee615e3b4a
---

# .gitlab-ci.yml

- **Étape 1 – Import des dépôts** : GitLab propose un importeur GitHub intégré qui transfère le code, les branches, les tags, les issues, les pull requests (converties en merge requests), les labels et les milestones. Accessible depuis Nouveau Projet → Importer depuis GitHub.
- **Étape 2 – Migration CI/CD** : Convertir les fichiers`.github/workflows/*.yml` en`.gitlab-ci.yml` . La syntaxe diffère mais les concepts sont similaires. GitLab fournit un guide de correspondance des concepts (jobs, steps/script, matrix/parallel).
- **Étape 3 – Migration des secrets** : Transférer les secrets GitHub (Settings → Secrets) vers les variables CI/CD GitLab (Settings → CI/CD → Variables). Marquer les variables sensibles comme « protégées » et « masquées ».
- **Étape 4 – Configuration de la sécurité** : Activer les templates de sécurité GitLab (SAST, DAST, dependency scanning) dans le fichier CI/CD pour remplacer les fonctionnalités GHAS/Dependabot.
- **Étape 5 – Formation de l’équipe** : Prévoir 1-2 semaines d’adaptation. Les concepts sont similaires mais la terminologie et l’interface diffèrent (Pull Request → Merge Request, Actions → CI/CD, Projects → Groups).

**Migration de GitLab vers GitHub :**

- **Étape 1 – Import des dépôts** : GitHub propose un importeur GitLab (github.com/new/import) qui transfère le code, les branches et les tags. Les issues et merge requests peuvent être migrées via l’API ou des outils tiers comme`gl2gh` .
- **Étape 2 – Migration CI/CD** : Convertir le`.gitlab-ci.yml` en workflows GitHub Actions. Identifier les actions du marketplace qui remplacent les fonctionnalités GitLab natives (Docker build, Kubernetes deploy, security scanning).
- **Étape 3 – Remplacement des outils intégrés** : Identifier les alternatives pour les fonctionnalités GitLab que GitHub ne fournit pas nativement : registre de conteneurs (→ GHCR), pages (→ GitHub Pages), monitoring (→ outil tiers).
- **Étape 4 – Configuration des protections de branches** : Recréer les branch protection rules et les CODEOWNERS pour maintenir les politiques de revue de code.

Dans les deux cas, prévoyez une **période de transition de 2 à 4 semaines** avec les deux plateformes fonctionnant en parallèle. Utilisez des miroirs Git bidirectionnels pendant la transition pour éviter toute perte de données. La migration CI/CD est généralement la partie la plus chronophage, comptez 1 à 2 jours par pipeline complexe.

## Avantages et Inconvénients : Le Bilan

Après cette analyse approfondie, voici un résumé structuré des forces et faiblesses de chaque plateforme en mars 2026.

**GitHub – Avantages :**

- Communauté la plus large (100M+ utilisateurs), effet réseau incomparable pour l’open source
- Interface utilisateur rapide, intuitive et moderne
- GitHub Copilot, l’assistant IA le plus avancé pour le code
- Marketplace Actions avec 20 000+ intégrations
- Prix d’entrée très bas (Team à 4 $/utilisateur/mois)
- Intégration native dans l’écosystème Microsoft (Azure, VS Code, Teams)
- Courbe d’apprentissage douce, recrutement facilité

**GitHub – Inconvénients :**

- Code source propriétaire, appartient à Microsoft (CLOUD Act)
- Sécurité avancée (GHAS) en supplément coûteux (+49 $/utilisateur/mois)
- CI/CD (Actions) moins mature que GitLab pour les pipelines complexes
- Auto-hébergement limité et payant (Enterprise Server uniquement)
- Dépendance forte à l’écosystème tiers pour les fonctionnalités DevOps avancées
- Gestion de projet moins complète (pas d’Epics, Roadmaps limités)

**GitLab – Avantages :**

- Plateforme DevOps complète tout-en-un (CI/CD, sécurité, monitoring, registres)
- Open-core : Community Edition gratuite et auto-hébergeable
- DevSecOps intégré nativement (SAST, DAST, fuzz testing, compliance)
- Meilleur TCO pour les grandes équipes avec sécurité avancée
- Souveraineté des données : auto-hébergement gratuit, instances UE
- Intégration Kubernetes et Terraform native de premier ordre
- GitLab Duo IA inclus dans Ultimate (pas de surcoût)

**GitLab – Inconvénients :**

- Prix par utilisateur plus élevé sur les plans payants (29-99 $ vs 4-21 $)
- Interface parfois plus lente et moins intuitive que GitHub
- Communauté plus petite (40M vs 100M), moins d’effet réseau
- Courbe d’apprentissage plus raide du fait de la richesse fonctionnelle
- Marketplace d’intégrations moins fourni
- Copilot reste supérieur à Duo en qualité brute de suggestions IA

## Avis des Experts et de la Communauté

Les opinions des experts tech en 2026 convergent vers un consensus nuancé : il n’y a pas de gagnant absolu, mais un meilleur choix selon le contexte.

**MKBHD**, dans sa série sur les outils de productivité tech, observe que « GitHub est devenu le réseau social des développeurs. Même si GitLab offre techniquement plus de fonctionnalités, la valeur du réseau GitHub pour le recrutement et la collaboration open source est impossible à ignorer. » Cette perspective met en lumière un aspect souvent sous-estimé : GitHub n’est pas juste un outil, c’est un **réseau professionnel** pour les développeurs.

**Fireship**, connu pour ses analyses techniques concises, résume le débat ainsi : « GitHub pour la vitesse et la communauté, GitLab pour le contrôle et la sécurité. Si vous hésitez, commencez par GitHub et migrez vers GitLab quand la complexité de votre infrastructure le justifie. » Cette approche pragmatique est partagée par de nombreux experts de la communauté DevOps.

**ThePrimeagen**, dans une analyse approfondie de l’écosystème DevOps 2026, souligne un point crucial pour les développeurs européens : « Avec NIS2 et le RGPD, les entreprises européennes n’ont souvent pas le luxe de choisir uniquement sur la base des fonctionnalités. GitLab CE auto-hébergé est parfois la seule option qui passe les audits de conformité sans trois mois de paperasse supplémentaire. » Ce point résonne particulièrement avec le contexte réglementaire français, où les exigences de cybersécurité nationales sont de plus en plus strictes.

Du côté de la communauté Stack Overflow, le **Developer Survey 2025** montre que 82,8 % des développeurs utilisent GitHub contre 37 % pour GitLab (les chiffres ne sont pas exclusifs – beaucoup utilisent les deux). Cependant, parmi les développeurs qui travaillent en entreprise avec des exigences de sécurité élevées, GitLab affiche un taux de satisfaction supérieur, principalement grâce à son approche DevSecOps intégrée.

## Conformité Européenne : RGPD, NIS2 et AI Act

La conformité réglementaire est devenue un critère de choix majeur pour les organisations européennes en 2026. Trois réglementations clés influencent directement le choix entre GitLab et GitHub.

Le **RGPD** exige que les données personnelles des citoyens européens soient protégées avec des garanties appropriées, y compris lors des transferts hors UE. GitHub, en tant que filiale de Microsoft, est soumis au **CLOUD Act américain**, qui autorise les autorités américaines à demander l’accès aux données stockées par des entreprises américaines, même si ces données sont hébergées en Europe. Bien que Microsoft ait mis en place des mécanismes de protection (chiffrement, clauses contractuelles types), certaines organisations considèrent que cela crée un risque juridique inacceptable.

GitLab, en tant qu’entreprise cotée au NASDAQ mais avec une architecture permettant l’auto-hébergement complet, offre une solution à ce dilemme. En déployant **GitLab CE ou EE sur une infrastructure européenne**, les organisations gardent le contrôle total sur la localisation et le traitement des données. Aucune donnée ne transite par des serveurs américains, éliminant le risque lié au CLOUD Act.

