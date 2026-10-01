---
id: collect-261001-rattrapage/rattrapage/gitleaks-scanner-les-secrets-git-en-13-etapes-2026-1
title: "macOS via Homebrew"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Google", "Stripe"]
dates: []
keywords: ["aws", "cyber", "open source"]
source: docs/RAG/collect-261001-rattrapage/gitleaks-scanner-les-secrets-git-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 51]
sha256: a7fffea8585e5e1aace514774f81f03aaa9402aad06d9449d2ab050e9caf978b
---

# macOS via Homebrew

En 2025, GitGuardian a détecté **28,65 millions de nouveaux secrets** (clés API, mots de passe, jetons d’accès) codés en dur dans les dépôts publics de GitHub, soit une hausse de 34 % sur un an, la plus forte progression jamais enregistrée par l’éditeur dans son rapport State of Secrets Sprawl 2026. Pire encore : parmi les identifiants fuités en 2022, 64 % étaient toujours valides en janvier 2026. Autrement dit, la grande majorité des secrets qui fuitent dans le code ne sont jamais révoqués. Pour les équipes techniques françaises et européennes soumises à la directive NIS2 et au Cyber Resilience Act, ce chiffre n’est plus une simple statistique de conférence sécurité : c’est un risque de conformité direct. Ce tutoriel explique comment déployer GitLeaks, l’outil open source de référence pour scanner un dépôt Git à la recherche de secrets exposés, du premier scan jusqu’à l’intégration complète dans une chaîne CI/CD.

Ce sujet n’a rien de théorique pour les organisations françaises. Les incidents récents impliquant des fuites massives de données personnelles ou administratives rappellent régulièrement qu’un identifiant oublié dans un dépôt de code, un fichier de configuration ou un script d’automatisation peut servir de porte d’entrée initiale à un attaquant, bien avant toute exploitation de vulnérabilité logicielle complexe. Les équipes qui gèrent des applications exposées sur Internet, des API publiques ou des pipelines de déploiement automatisés ont donc tout intérêt à traiter le scan de secrets comme un contrôle de base, au même titre qu’un pare-feu applicatif ou une authentification à double facteur.

## Pourquoi le scan de secrets est devenu incontournable en 2026

Le phénomène a un nom dans l’industrie : le “secrets sprawl”, la dissémination incontrôlée de secrets dans le code source. Selon le rapport State of Secrets Sprawl 2026 de GitGuardian, plus de 4 012 054 dépôts publics GitHub contenaient au moins un secret en 2025, avec une moyenne de 0,32 secret par dépôt. La croissance des secrets exposés (+152 % depuis 2021) dépasse largement celle du nombre de développeurs actifs sur la plateforme (+98 % sur la même période). Un facteur aggrave la tendance depuis 2025 : l’essor du code généré par IA. Les clés d’API liées aux services d’intelligence artificielle ont bondi de 81 % en un an, avec plus de 1,27 million de secrets détectés pour ces seuls services. Les outils de complétion de code assistés par IA font fuiter des secrets à un rythme environ deux fois supérieur à la moyenne observée sur l’ensemble de GitHub.

Cette explosion coïncide avec un durcissement réglementaire en Europe. La directive NIS2 impose aux entités essentielles et importantes de mettre en place des mesures techniques de gestion des vulnérabilités sur toute leur chaîne d’approvisionnement logicielle, ce qui inclut la gestion des identifiants dans le code source. Le Cyber Resilience Act, de son côté, exige des pratiques de développement sécurisé et une gestion continue des vulnérabilités pendant tout le cycle de vie d’un produit numérique. Un secret oublié dans un commit peut désormais être qualifié de faille de sécurité au sens réglementaire, avec des obligations de notification et de remédiation associées. Le scan de secrets, longtemps considéré comme une bonne pratique optionnelle, devient une brique de conformité difficile à ignorer pour toute organisation qui publie du code en Europe. L’ENISA, l’agence de l’Union européenne pour la cybersécurité, classe d’ailleurs la gestion des secrets et des identifiants parmi les mesures techniques prioritaires de sécurisation du cycle de développement logiciel dans ses recommandations aux États membres.

## Qu’est-ce que GitLeaks et comment il détecte les secrets

GitLeaks est un outil en ligne de commande, écrit en Go, entièrement gratuit et open source, distribué sous licence MIT via son dépôt GitHub officiel. Il scanne des dépôts Git, des chemins de fichiers arbitraires ou un flux stdin à la recherche de motifs correspondant à des secrets. Son moteur de détection combine deux approches complémentaires. La première repose sur des règles regex prédéfinies, chacune ciblant un format de secret précis : les clés d’accès AWS commencent toujours par un préfixe reconnaissable, les jetons GitHub personnels suivent un format documenté, de même pour les clés d’API Stripe, Slack ou Google Cloud. La seconde approche s’appuie sur l’entropie : GitLeaks calcule le niveau d’aléatoire d’une chaîne de caractères et la signale comme suspecte si elle dépasse un seuil, en particulier lorsqu’elle est assignée à une variable nommée de façon évocatrice comme API_KEY ou SECRET_TOKEN.

La version stable actuelle est GitLeaks v8.30.1, publiée le 21 mars 2026. Toute la configuration de l’outil passe par un fichier au format TOML, généralement nommé .gitleaks.toml, qui définit les règles actives, les listes d’exclusion (allowlists) par fichier, par chemin ou par auteur de commit, ainsi que le format des rapports générés (JSON, SARIF, CSV). GitLeaks embarque une configuration par défaut couvrant plusieurs centaines de types de secrets connus, que chaque équipe peut ensuite étendre avec ses propres règles internes.

## Prérequis : versions et environnement nécessaires

Avant de commencer, vérifiez que votre environnement dispose des éléments suivants. Ce tutoriel a été testé sur Ubuntu 24.04 LTS, macOS 15 et Windows 11 avec WSL2.

- **Git** version 2.30 ou supérieure (nécessaire pour l’historique complet et les hooks)
- **GitLeaks** v8.30.1 (dernière version stable, 21 mars 2026)
- **Go** 1.21 ou supérieur, uniquement si vous compilez GitLeaks depuis les sources
- **Docker** 24.0 ou supérieur, pour l’exécution en conteneur (optionnel mais recommandé en CI)
- **Python** 3.9 ou supérieur et le framework**pre-commit** 3.6 ou supérieur, pour les hooks locaux
- Un compte **GitHub** avec GitHub Actions activé, ou un compte**GitLab** avec GitLab CI/CD activé
- Un accès administrateur au dépôt cible, indispensable si vous devez réécrire l’historique Git
- **git-filter-repo** version 2.45 ou supérieure, pour la purge de l’historique (voir étape 11)

Comptez environ 60 minutes pour parcourir l’intégralité des 13 étapes sur un dépôt de taille moyenne, en incluant les tests d’intégration CI/CD. La réécriture de l’historique sur un très gros dépôt (plusieurs dizaines de milliers de commits) peut prendre nettement plus longtemps.

## Étape 1 : Installer GitLeaks selon votre système

GitLeaks propose plusieurs méthodes d’installation selon votre plateforme. Sur macOS, la méthode la plus rapide passe par Homebrew. Sur Linux, vous pouvez télécharger le binaire précompilé depuis la page de releases GitHub ou utiliser un gestionnaire de paquets tiers. Sur Windows, l’exécutable est également disponible en téléchargement direct, et fonctionne nativement dans WSL2.

```
# macOS via Homebrew
brew install gitleaks
# Linux : téléchargement du binaire v8.30.1
curl -sSL https://github.com/gitleaks/gitleaks/releases/download/v8.30.1/gitleaks_8.30.1_linux_x64.tar.gz -o gitleaks.tar.gz
tar -xzf gitleaks.tar.gz
sudo mv gitleaks /usr/local/bin/
# Vérification de l'installation
gitleaks version
# Alternative via Docker (sans installation locale)
docker pull zricethezav/gitleaks:v8.30.1
```
Une fois l’installation terminée, la commande gitleaks version doit afficher 8.30.1. Si vous préférez une approche totalement isolée sans installer de binaire sur votre machine, l’image Docker officielle zricethezav/gitleaks est mise à jour à chaque release et constitue l’option recommandée pour les environnements CI/CD partagés.

## Étape 2 : Lancer votre premier scan avec gitleaks detect

