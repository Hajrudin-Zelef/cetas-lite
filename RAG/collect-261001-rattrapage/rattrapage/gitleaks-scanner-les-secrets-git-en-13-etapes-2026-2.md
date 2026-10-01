---
id: collect-261001-rattrapage/rattrapage/gitleaks-scanner-les-secrets-git-en-13-etapes-2026-2
title: "macOS via Homebrew"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["attention", "aws"]
source: docs/RAG/collect-261001-rattrapage/gitleaks-scanner-les-secrets-git-en-13-etapes-2026.md
source_anchor: ""
source_lines: [52, 199]
sha256: 692a07a4b99c6d73406dcbd93bbe333ff6e5119dfa202d0380733393e6cfc8d4
---

# macOS via Homebrew

Placez-vous à la racine d’un dépôt Git existant et lancez la commande de détection. GitLeaks analyse par défaut l’ensemble de l’historique des commits, pas seulement l’état actuel des fichiers, ce qui est essentiel puisqu’un secret supprimé du code reste présent dans l’historique tant qu’il n’est pas purgé.

```
cd mon-projet/
gitleaks detect --source . --verbose
# Pour scanner uniquement l'état actuel (sans historique)
gitleaks detect --source . --no-git --verbose
# Pour générer un rapport au format JSON
gitleaks detect --source . --report-format json --report-path gitleaks-report.json
```
Le code de sortie de la commande est déterminant pour l’automatisation : GitLeaks retourne 0 si aucun secret n’est trouvé, et 1 dès qu’au moins un secret est détecté. C’est ce code de sortie qui permet de bloquer un commit, une pull request ou un pipeline CI en cas de fuite.

## Étape 3 : Lire et interpréter le rapport GitLeaks

Chaque secret détecté génère une entrée détaillée dans le rapport, contenant le fichier concerné, le numéro de ligne, le hash du commit, l’auteur, la date, la règle qui a déclenché l’alerte et un extrait tronqué du secret (jamais affiché en clair dans les logs par défaut, pour éviter une double fuite via les journaux CI). Voici un exemple de sortie console typique lors d’un scan qui trouve une clé AWS oubliée dans un fichier de configuration :

```
Finding:     AKIA************WXYZ
Secret:      AKIA****************
RuleID:      aws-access-token
Entropy:     3.912345
File:        config/settings.py
Line:        42
Commit:      a1b2c3d4e5f6...
Author:      [email protected]
Date:        2026-06-14T09:12:00Z
10:34AM INF 1 commits scanned.
10:34AM INF scan completed in 842ms
10:34AM WRN leaks found: 1
```
Chaque ligne du rapport doit être triée en deux catégories : un vrai secret actif à révoquer immédiatement, ou un faux positif (une clé de test, une donnée d’exemple dans la documentation, un placeholder). C’est cette distinction qui justifie l’étape suivante, la configuration d’une allowlist.

## Étape 4 : Personnaliser la configuration avec .gitleaks.toml

La configuration par défaut de GitLeaks couvre la grande majorité des cas d’usage, mais chaque organisation possède ses propres formats de secrets internes (jetons d’API maison, identifiants de bases de données propriétaires). Créez un fichier .gitleaks.toml à la racine du dépôt pour étendre les règles existantes.

```
title = "Configuration GitLeaks — Projet Interne"
[extend]
useDefault = true
[[rules]]
id = "jeton-interne-acme"
description = "Jeton d'API interne ACME Corp"
regex = '''acme_[a-zA-Z0-9]{32}'''
tags = ["api", "interne"]
[[rules]]
id = "chaine-connexion-postgres"
description = "Chaîne de connexion PostgreSQL avec mot de passe"
regex = '''postgres(?:ql)?:\/\/[^:]+:[^@]+@[^\/\s]+'''
tags = ["database", "credentials"]
```
La directive useDefault = true conserve toutes les règles fournies nativement par GitLeaks tout en ajoutant vos propres motifs. Relancez ensuite gitleaks detect –source . –config .gitleaks.toml pour valider que la nouvelle règle fonctionne, idéalement contre un fichier de test contenant volontairement un faux jeton correspondant au format attendu.

## Étape 5 : Gérer les faux positifs avec une allowlist

Sans allowlist correctement configurée, un projet contenant de la documentation avec des exemples de clés, des fixtures de tests ou des fichiers de vendoring tiers génère rapidement des dizaines de fausses alertes, ce qui use la vigilance des équipes et pousse à ignorer les vrais résultats. GitLeaks permet d’exclure des chemins entiers, des regex spécifiques ou des commits précis.

```
[allowlist]
description = "Exclusions globales du projet"
paths = [
  '''node_modules''',
  '''vendor''',
  '''(.*)_test\.go''',
  '''docs/exemples/'''
]
regexes = [
  '''EXEMPLE_[A-Z0-9]{16}''',
  '''your-api-key-here'''
]
commits = [
  "a1b2c3d4e5f6789012345678901234567890abcd"
]
```
Attention à ne pas transformer l’allowlist en solution de facilité : chaque exclusion doit être documentée et justifiée dans une revue de code, sinon elle devient elle-même un angle mort de sécurité, un vrai secret pouvant se glisser discrètement dans un chemin exclu à tort.

## Étape 6 : Bloquer les secrets avant le commit avec un hook pre-commit

La détection la plus efficace intervient avant même que le secret n’atteigne le dépôt distant. Le framework pre-commit permet d’exécuter GitLeaks automatiquement à chaque tentative de commit local, en scannant uniquement les fichiers modifiés (staged). Installez le framework puis déclarez GitLeaks dans le fichier de configuration.

```
# Installation du framework pre-commit
pip install pre-commit
# Fichier .pre-commit-config.yaml à la racine du projet
cat > .pre-commit-config.yaml << 'EOF'
repos:
  - repo: https://github.com/gitleaks/gitleaks
    rev: v8.30.1
    hooks:
      - id: gitleaks
EOF
# Activation du hook dans le dépôt local
pre-commit install
```
À partir de cet instant, chaque git commit déclenche automatiquement gitleaks protect sur les fichiers indexés. Si un secret est détecté, le commit est bloqué et un message d'erreur détaillé s'affiche dans le terminal, obligeant le développeur à retirer le secret avant de poursuivre.

## Étape 7 : Automatiser le scan dans GitHub Actions

Un hook local peut être contourné avec l'option --no-verify, volontairement ou par mégarde. La ligne de défense suivante s'exécute donc côté serveur, dans le pipeline CI, sur chaque push et chaque pull request. Créez le fichier .github/workflows/gitleaks.yml suivant.

```
name: Scan GitLeaks
on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]
jobs:
  gitleaks:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Exécuter GitLeaks
        uses: gitleaks/gitleaks-action@v2
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          GITLEAKS_CONFIG: .gitleaks.toml
```
L'option fetch-depth: 0 est indispensable : sans elle, GitHub Actions ne récupère qu'un clone superficiel du dépôt et GitLeaks ne peut analyser que le dernier commit, laissant passer un secret introduit plus tôt dans la même pull request. Le pipeline échoue automatiquement (statut rouge sur la pull request) dès qu'un secret est détecté, empêchant la fusion tant que le problème n'est pas résolu.

## Étape 8 : Intégrer GitLeaks à GitLab CI

Pour les équipes hébergées sur GitLab, le principe est identique mais la syntaxe change. Ajoutez le job suivant à votre fichier .gitlab-ci.yml, en utilisant l'image Docker officielle pour éviter toute installation manuelle du binaire dans le runner.

```
scan-secrets:
  stage: test
  image:
    name: zricethezav/gitleaks:v8.30.1
    entrypoint: [""]
  script:
    - gitleaks detect --source . --report-format sarif --report-path gitleaks.sarif --exit-code 1
  artifacts:
    reports:
      sast: gitleaks.sarif
    when: always
  rules:
    - if: '$CI_PIPELINE_SOURCE == "merge_request_event"'
    - if: '$CI_COMMIT_BRANCH == "main"'
```
L'option --exit-code 1 garantit que le job échoue en cas de détection, ce qui bloque la merge request dans les règles de protection de branche si celles-ci exigent un pipeline vert. Le format de sortie SARIF permet également d'alimenter directement le tableau de bord de sécurité intégré à GitLab Ultimate.

## Étape 9 : Scanner tout l'historique Git en profondeur

Un scan CI classique porte généralement sur les commits récents d'une pull request, pas sur l'intégralité de l'historique du dépôt. Or un audit initial, avant la mise en place de GitLeaks, doit couvrir tous les commits depuis la création du projet. Utilisez l'option --log-opts pour cibler une plage précise, ou omettez-la pour scanner tout l'historique.

