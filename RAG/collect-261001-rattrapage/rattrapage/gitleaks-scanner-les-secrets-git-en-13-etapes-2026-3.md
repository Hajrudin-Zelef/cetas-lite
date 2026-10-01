---
id: collect-261001-rattrapage/rattrapage/gitleaks-scanner-les-secrets-git-en-13-etapes-2026-3
title: "macOS via Homebrew"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Stripe"]
dates: []
keywords: ["aws", "cyber", "exploit", "open source"]
source: docs/RAG/collect-261001-rattrapage/gitleaks-scanner-les-secrets-git-en-13-etapes-2026.md
source_anchor: ""
source_lines: [200, 279]
sha256: 254c0fb15e7447997bba645fdfb2dfae7f8902dd2a568016f1c28ccea36e1517
---

# macOS via Homebrew

```
# Scan complet de tout l'historique du dépôt
gitleaks detect --source . --log-opts="--all" --verbose
# Scan limité aux 12 derniers mois
gitleaks detect --source . --log-opts="--since=12.months" --verbose
# Scan d'une branche spécifique uniquement
gitleaks detect --source . --log-opts="feature/nouvelle-api" --verbose
```
Sur un dépôt actif depuis plusieurs années, ce scan initial fait souvent remonter des dizaines, parfois des centaines de résultats. Ne cédez pas à la tentation de tout corriger en urgence : classez les résultats par gravité (secret probablement encore valide, secret révoqué depuis, faux positif) avant d'agir, comme détaillé à l'étape suivante.

## Étape 10 : Réagir face à un secret détecté (rotation immédiate)

Dès qu'un secret est confirmé, la règle de base de l'industrie est simple : considérez-le comme compromis, quelle que soit la probabilité réelle qu'il ait été exploité, surtout s'il a été poussé un jour vers un dépôt distant, même privé. La procédure de remédiation suit toujours le même ordre :

- **Révoquer** immédiatement la clé ou le mot de passe compromis dans la console du fournisseur concerné (AWS IAM, GitHub Settings, tableau de bord Stripe, etc.)
- **Générer** un nouveau secret et mettre à jour toutes les applications et services qui en dépendent
- **Vérifier les journaux d'accès** du fournisseur pour détecter une éventuelle utilisation frauduleuse pendant la fenêtre d'exposition
- **Rechercher les réutilisations** du même secret dans d'autres dépôts ou environnements, une pratique malheureusement fréquente
- **Remplacer** le secret en dur dans le code par une variable d'environnement ou un appel à un gestionnaire de secrets externe comme HashiCorp Vault

La rotation seule ne suffit pas : tant que le secret révoqué reste visible dans l'historique Git, il continue à polluer les scans futurs et peut induire en erreur un auditeur externe. C'est l'objet de l'étape suivante.

## Étape 11 : Purger l'historique avec git-filter-repo

Supprimer un secret du dernier commit ne l'efface pas de l'historique : il reste récupérable via n'importe quel clone antérieur. Deux outils dominent la réécriture d'historique Git : BFG Repo-Cleaner, rapide et écrit en Java, et git-filter-repo, l'outil désormais recommandé officiellement par le projet Git en remplacement de git filter-branch, jugé trop lent et sujet aux erreurs. Le choix entre les deux dépend surtout de l'environnement disponible : git-filter-repo ne nécessite qu'un interpréteur Python, tandis que BFG requiert une machine virtuelle Java installée sur le poste qui effectue la réécriture.

```
# Installation de git-filter-repo
pip install git-filter-repo
# Créer un fichier listant le texte exact du secret à purger
echo "AKIAEXEMPLEDECLEEXPOSEE==>SUPPRIME" > secrets-a-purger.txt
# Réécriture de l'historique complet
git filter-repo --replace-text secrets-a-purger.txt
# Alternative avec BFG Repo-Cleaner (nécessite Java)
java -jar bfg.jar --replace-text secrets-a-purger.txt mon-depot.git
# Forcer la mise à jour du dépôt distant après réécriture
git push origin --force --all
git push origin --force --tags
```
Cette opération réécrit tous les hashes de commit après le point d'insertion du secret, ce qui casse tous les clones locaux existants chez vos collaborateurs. Prévenez systématiquement toute l'équipe avant de forcer le push, et demandez-leur de recloner le dépôt plutôt que de tenter un pull sur leur copie locale devenue incompatible.

## Étape 12 : Centraliser les rapports SARIF dans GitHub Security

Pour les organisations gérant des dizaines de dépôts, consulter chaque rapport individuellement n'est pas tenable. Le format SARIF (Static Analysis Results Interchange Format), standard ouvert utilisé par la plupart des scanners de sécurité statique, permet de centraliser les résultats de GitLeaks directement dans l'onglet Security de GitHub, aux côtés des alertes Dependabot et CodeQL.

```
- name: Générer le rapport SARIF
  run: gitleaks detect --source . --report-format sarif --report-path results.sarif --exit-code 0
- name: Publier dans GitHub Security
  uses: github/codeql-action/upload-sarif@v3
  with:
    sarif_file: results.sarif
```
Notez l'usage de --exit-code 0 dans cette étape spécifique : on souhaite que le job publie le rapport même en présence de secrets, plutôt que d'échouer immédiatement et de sauter l'étape de publication. Le blocage effectif de la pull request est alors géré séparément, en amont, via le job de scan classique décrit à l'étape 7.

## Étape 13 : Documenter la conformité NIS2 et Cyber Resilience Act

Pour les entités concernées par la directive NIS2 ou le Cyber Resilience Act, la mise en place technique de GitLeaks ne suffit pas : il faut également produire une trace documentaire démontrable lors d'un audit. Consignez systématiquement la politique de scan de secrets (fréquence, périmètre, outils), la procédure de remédiation en cas de détection, ainsi que l'historique des incidents traités avec leur délai de résolution. Ces éléments alimentent directement le dossier de conformité NIS2 relatif à la gestion des vulnérabilités de la chaîne d'approvisionnement logicielle, ainsi que la documentation de cycle de développement sécurisé exigée par le Cyber Resilience Act pour tout produit numérique commercialisé dans l'Union européenne.

Un point de vigilance mérite d'être souligné : GitLeaks est un contrôle technique parmi d'autres, pas une solution de conformité à lui seul. Il doit s'inscrire dans une politique plus large de gestion des secrets, incluant idéalement un gestionnaire dédié comme HashiCorp Vault pour éliminer à la racine le besoin de stocker des identifiants en dur dans le code.

## GitLeaks vs TruffleHog vs GitGuardian : lequel choisir

GitLeaks n'est pas le seul outil du marché. TruffleHog, également open source, repose sur une combinaison similaire de règles regex et de détection par entropie, avec un accent mis sur la vérification en direct de la validité de certains types de secrets. GitGuardian, de son côté, est une plateforme SaaS commerciale qui surveille en continu l'intégralité de GitHub public à l'échelle d'Internet, en plus des dépôts privés d'un organisme, et propose un tableau de bord de remédiation centralisé.

| Critère | GitLeaks | TruffleHog | GitGuardian | 
|---|---|---|---|
| Modèle | Open source (MIT) | Open source + offre commerciale | SaaS commercial | 
| Prix | Gratuit | CLI gratuit, plans payants pour la plateforme | Offre gratuite limitée + plans entreprise | 
| Moteur | Regex + entropie, config TOML | Regex + entropie, vérification live | Règles propriétaires + validation | 
| Périmètre | Dépôts Git, fichiers, stdin | Dépôts Git et autres sources | GitHub public + dépôts privés, CI/CD, conteneurs, IaC | 
| Hébergement | Auto-hébergé, CLI | Auto-hébergé ou hébergé | Plateforme cloud centralisée | 
| Usage typique | Pre-commit et CI/CD par dépôt | Pre-commit et CI/CD par dépôt | Surveillance de flotte à l'échelle de l'organisation | 

Dans la pratique, ces outils ne s'excluent pas : une organisation peut très bien déployer GitLeaks localement dans les hooks pre-commit et les pipelines CI de chaque dépôt, tout en s'abonnant à GitGuardian pour une surveillance transversale et une visibilité centralisée sur des dizaines, voire des centaines de dépôts.

## Types de secrets détectés et niveaux de risque

Toutes les catégories de secrets ne présentent pas le même niveau de risque en cas d'exposition. Le tableau suivant résume les types les plus couramment détectés par GitLeaks et leur impact potentiel.

