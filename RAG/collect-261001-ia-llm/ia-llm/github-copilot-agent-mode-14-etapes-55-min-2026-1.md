---
id: collect-261001-ia-llm/ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026-1
title: ".github/copilot-instructions.md"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft"]
dates: []
keywords: ["copilot", "agent", "benchmark", "claude", "mai"]
source: docs/RAG/collect-261001-ia-llm/github-copilot-agent-mode-14-etapes-55-min-2026.md
source_anchor: ""
source_lines: [1, 44]
sha256: bee40c2a34d56b532a64fefa64166a25a4a3032c4b4f0d0fbf14020ee3d97910
---

# .github/copilot-instructions.md

GitHub Copilot a changé de visage en 2026. À l’assistant qui complétait vos lignes de code s’est ajouté un agent capable de lire un projet entier, de planifier une série de modifications, de les appliquer dans plusieurs fichiers à la fois, puis de lancer lui-même les commandes nécessaires pour vérifier que tout fonctionne. Ce mode s’appelle l’Agent Mode, parfois désigné sous le nom de Copilot App pour son interface dédiée à l’agent : selon un état des lieux publié par News Creeta en juin 2026, cette interface agentique est désormais incluse sur l’ensemble des paliers payants, de Pro (10 $/mois) à Enterprise (39 $/utilisateur/mois), en passant par Pro+, Max et Business. Il tourne directement dans Visual Studio Code et change la façon d’aborder une tâche de développement ordinaire.

Ce tutoriel montre comment l’installer, le configurer et l’utiliser correctement, avec des exemples testés sur un vrai petit projet du début à la fin. Vous verrez aussi les erreurs les plus fréquentes, les solutions aux blocages courants, et comment garder le contrôle sur les coûts liés aux crédits d’utilisation, un sujet devenu central depuis le passage à la facturation à l’usage en juin 2026. Comptez environ 55 minutes pour suivre les 14 étapes, en partant d’un poste vierge.

## Qu’est-ce que GitHub Copilot Agent Mode ?

L’Agent Mode est le troisième niveau d’interaction proposé par GitHub Copilot dans la vue Chat de VS Code, après Ask (poser une question) et Edit (modifier un fichier ciblé). Plutôt que de répondre à une question ou de retoucher un seul fichier, l’agent prend en charge une tâche complète formulée en langage naturel. Il explore le code du projet, construit un plan d’action, écrit ou modifie plusieurs fichiers, exécute des commandes dans le terminal intégré, observe les résultats, puis corrige ce qui ne fonctionne pas. Depuis l’indexation par recherche sémantique introduite en mars 2026, cette phase d’exploration du code est devenue 50 % plus rapide sur les projets volumineux, une amélioration détaillée dans le Copilot Digest de GitHub publié en avril 2026. Le tout sans que vous ayez à copier-coller la moindre ligne. GitHub détaille ce fonctionnement dans son explicatif officiel sur l’Agent Mode, publié sur le blog de l’entreprise.

Concrètement, la différence avec l’ancien Copilot tient à l’autonomie du raisonnement. Un prompt du type « ajoute un endpoint de recherche paginée et les tests associés » suffit à déclencher une séquence complète : lecture du schéma de données existant, écriture du code, génération des tests, exécution de la suite de tests, puis correction automatique si un test échoue. Vous gardez la main à chaque étape grâce à un système de validation des modifications et des commandes, détaillé plus loin dans ce tutoriel.

GitHub met en avant des chiffres significatifs sur sa propre page produit : les développeurs qui utilisent Copilot se disent jusqu’à 55 % plus productifs pour écrire du code, sans perte de qualité annoncée, et jusqu’à 75 % plus satisfaits de leur travail que ceux qui ne l’utilisent pas. L’entreprise cite aussi le cas du groupe brésilien **Grupo Boticário**, qui rapporte une hausse de productivité développeur de 94 % après l’adoption de l’outil. Sur le plan technique, l’agent propulsé par Claude affichait déjà un score de 56 % sur le benchmark SWE-bench Verified dès avril 2025, une performance rappelée par Dev.to dans son bilan d’avril 2026 et souvent citée comme référence pour juger de la fiabilité de l’Agent Mode sur des tâches réelles de correction de bugs. Ces chiffres proviennent en grande partie de GitHub lui-même et méritent d’être lus comme des données commerciales plutôt que comme une étude indépendante, mais ils donnent une idée de l’ampleur du pari que l’entreprise a fait sur l’Agent Mode.

Pour les équipes basées en France et en Europe, l’intérêt dépasse le simple gain de vitesse. Les offres Business et Enterprise ajoutent des options pensées pour les organisations soumises au RGPD (authentification unique, journaux d’audit, politiques de rétention), abordées plus loin dans la section consacrée à la sécurité. Si vous cherchez plutôt un panorama des autres assistants IA pour développeurs, notre rubrique logiciels regroupe l’ensemble de nos tutoriels sur les outils de développement.

## Prérequis avant de commencer

Avant de lancer la première étape, vérifiez que votre environnement correspond à ce tableau. Rien d’exotique : l’essentiel tient dans une version récente de VS Code et un abonnement GitHub Copilot actif, même sur le palier gratuit pour les premiers tests.

| Élément | Version ou condition minimale | Remarque | 
|---|---|---|
| Système d’exploitation | Windows 10/11, macOS ou Linux | Aucune restriction propre à l’Agent Mode | 
| Visual Studio Code | Version récente (1.99 ou ultérieure) | Vérifiez avec `code --version` | 
| Compte GitHub | Actif, avec abonnement Copilot | Free suffit pour découvrir, Pro recommandé pour un usage réel | 
| Extensions VS Code | GitHub Copilot + GitHub Copilot Chat | Installables depuis le Marketplace intégré | 
| Git | Dernière version stable | Indispensable pour suivre les modifications de l’agent | 
| Python | 3.11 ou plus récent | Utilisé pour le projet complet de ce tutoriel | 
| Node.js | 20 LTS ou plus récent (optionnel) | Utile si vous adaptez le projet en JavaScript | 
| Connexion réseau stable | Requise | L’agent interroge le cloud à chaque étape de raisonnement | 

Un dernier conseil avant de vous lancer : travaillez toujours dans un dépôt Git initialisé, même pour un projet de test. L’Agent Mode modifie parfois plusieurs fichiers en une seule séquence, et pouvoir revenir en arrière avec `git diff` ou `git checkout` change tout en cas de résultat inattendu.

## Étape 1 – Choisir le bon plan GitHub Copilot

Dès le 27 avril 2026, BigHat Group rapportait que GitHub s’apprêtait à généraliser un système de crédits IA mensuels à l’ensemble des paliers payants à compter du 1er juin 2026, tout en maintenant les prix d’abonnement en l’état : 10 $ pour Pro, 19 $ pour Business et 39 $ pour Enterprise. Ce système de crédits remplace le mécanisme de requêtes premium mis en place dès mai 2025, quand GitHub avait fixé un quota de 300 requêtes premium mensuelles pour les abonnés Pro et facturait 0,04 $ chaque requête supplémentaire au-delà de ce plafond. Plus vous consommez de requêtes vers des modèles premium, plus vous puisez dans ce crédit. Comprendre cette grille avant de commencer évite les mauvaises surprises en fin de mois, surtout si vous comptez faire tourner l’Agent Mode plusieurs heures par jour.

| Palier | Prix | Complétions de code | Crédits mensuels inclus | Accès à l’Agent Mode | 
|---|---|---|---|---|
| Free | 0 $/mois | 2 000/mois, 50 requêtes chat | Modèles de base uniquement | Limité | 
| Pro | 10 $/mois | Illimitées | 15 $ de crédits | Oui | 
| Pro+ | 39 $/mois | Illimitées | 70 $ de crédits, modèles premium (dont Opus) | Oui, étendu | 
| Max | 100 $/mois | Illimitées | 200 $ de crédits, accès prioritaire aux nouveautés | Oui, complet | 
| Business | 19 $/utilisateur/mois | Illimitées | Géré par l’organisation | Oui, avec politiques admin | 
| Enterprise | Sur devis | Illimitées | Personnalisé | Oui, avec options de conformité | 

