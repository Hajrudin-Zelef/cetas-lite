---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026-5
title: "AGENTS.md"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "attention", "cloud agent", "incident", "mcp"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [281, 338]
sha256: 7bd9f6e69147694a2b202f10ffd4ff4e57b2b4f8b4a63ebf3ce530984f304166
---

# AGENTS.md

## Déployer Cursor AI à l'échelle d'une équipe

Adopter Cursor AI pour un développeur seul prend une heure. Le déployer proprement sur une équipe de vingt personnes demande une méthode, sous peine de se retrouver avec vingt configurations différentes et zéro règle partagée.

1. **Commencez par un groupe pilote de trois à cinq développeurs** plutôt qu'un déploiement en une fois. Ce groupe teste les règles de projet, identifie les frictions et évite qu'une mauvaise configuration se propage à toute l'équipe.
2. **Versionnez les fichiers .cursor/rules et AGENTS.md dans le dépôt Git** comme n'importe quel autre fichier de configuration. Ils doivent passer en revue de code au même titre qu'un changement d'architecture, puisqu'ils influencent directement le comportement de l'IA pour toute l'équipe.
3. **Centralisez la facturation sur la formule Teams** plutôt que de laisser chaque développeur payer sa propre formule Pro. Le contrôle d'accès et le suivi de consommation deviennent alors visibles pour un seul administrateur au lieu d'être dispersés.
4. **Fixez un budget mensuel de crédits par équipe et surveillez-le dès la première semaine.** Les pics de consommation liés au mode Agent et aux Cloud Agents peuvent surprendre une organisation qui découvre l'outil, surtout si plusieurs développeurs lancent des tâches lourdes le même jour.
5. **Formez les nouvelles recrues sur les règles existantes avant de les laisser utiliser le mode Agent sans supervision.** Un développeur junior qui ne connaît pas encore les conventions du projet aura plus de mal à repérer un diff problématique généré par l'agent.

Un dernier point mérite l'attention des équipes françaises et européennes : la formule Teams inclut un mode confidentialité que beaucoup d'organisations laissent désactivé par défaut faute de savoir qu'il existe. Vérifiez ce réglage avant le déploiement plutôt qu'après un premier incident.

## Pièges courants à éviter avec Cursor AI

- **Laisser le mode Agent tourner sans relire les diffs.** L'agent peut produire du code qui compile et passe les tests tout en introduisant une faille de sécurité ou une régression fonctionnelle discrète. La relecture manuelle reste non négociable, même quand le résultat semble propre.
- **Oublier que les crédits Pro, Pro+ et Ultra ne sont pas illimités.** Un usage intensif du mode Agent sur de gros refactors peut consommer un quota mensuel en quelques jours. Surveillez la consommation dans le tableau de bord dès la première semaine d'usage.
- **Multiplier les fichiers .mdc sans jamais les auditer.** Des règles contradictoires entre deux fichiers produisent un comportement imprévisible de l'agent, qui doit arbitrer silencieusement entre deux instructions opposées.
- **Connecter un serveur MCP avec des clés API trop permissives.** Un serveur MCP relié à une base de données de production avec des droits d'écriture complets transforme une erreur de prompt en incident réel. Limitez les permissions au strict nécessaire.
- **Confondre AGENTS.md et .cursor/rules et dupliquer les instructions.** Les deux mécanismes cohabitent, mais des consignes contradictoires entre les deux fichiers créent de la confusion pour l'agent comme pour l'équipe.
- **Ignorer la confidentialité des données pour du code sous NDA ou soumis au RGPD.** Sans mode confidentialité activé, du code sensible peut transiter par des serveurs tiers. Vérifiez les réglages de rétention des données avant de connecter un dépôt sensible.

## Dépannage : les problèmes les plus fréquents et leurs solutions

| Problème | Cause probable | Solution | 
|---|---|---|
| Tab ne propose plus de suggestions | Modèle désactivé ou quota mensuel atteint | Vérifier Settings > Models et le solde de crédits restant | 
| Les règles .mdc ne s'appliquent pas | Extension .md utilisée au lieu de .mdc, ou frontmatter mal formé | Renommer le fichier en .mdc et valider la syntaxe YAML de l'en-tête | 
| L'agent boucle sur la même erreur | Contexte insuffisant ou règles contradictoires entre deux fichiers | Simplifier le prompt et vérifier les règles marquées alwaysApply | 
| L'import des paramètres VS Code échoue | Version de VS Code trop ancienne ou profil corrompu | Réinstaller les extensions manuellement depuis le marketplace intégré | 
| Le serveur MCP ne se connecte pas | Chemin de commande incorrect ou variable d'environnement manquante | Tester la commande du serveur en local, hors Cursor, avant de la déclarer | 
| Facturation Bugbot inattendue | Fonction à l'usage activée par défaut sur la formule Teams | Désactiver l'option dans les réglages d'équipe ou fixer un plafond mensuel | 
| Le Cloud Agent ne se déclenche pas depuis Slack | Intégration Slack non autorisée par l'administrateur du workspace | Réautoriser l'application Cursor dans les paramètres Slack de l'organisation | 
| Perte de contexte sur une longue conversation | Limite de fenêtre de contexte du modèle atteinte | Démarrer une nouvelle conversation et résumer l'essentiel dans une règle .mdc | 
| Synchronisation GitHub bloquée | Jeton d'accès expiré ou révoqué | Reconnecter le compte GitHub depuis les réglages du compte Cursor | 
| Latence élevée en mode Agent | Charge serveur élevée aux heures de pointe ou modèle premium saturé | Basculer temporairement sur un modèle alternatif dans Settings > Models | 

## Astuces avancées pour les développeurs expérimentés

- **Combinez règles globales et règles par dossier sur les monorepos.** Une règle à la racine fixe les conventions générales, des règles dans chaque sous-projet affinent le contexte sans surcharger l'ensemble.
- **Réservez le modèle le plus coûteux aux tâches qui le justifient.** Basculez sur un modèle plus léger pour la génération de tests ou la documentation, et gardez le modèle premium pour les refactors complexes ou l'architecture.
- **Scriptez le déclenchement de Cloud Agents via l'API pour les tâches récurrentes.** Mise à jour de dépendances, génération de changelog ou synchronisation de traductions peuvent tourner sans intervention humaine, sur un déclencheur planifié.
- **Combinez Cursor AI et l'exécution locale via Ollama pour les portions de code sensibles.** Gardez le cloud pour la majorité du travail et un modèle local pour les fichiers qui ne doivent jamais sortir de votre infrastructure.
- **Auditez régulièrement les serveurs MCP connectés.** Un serveur oublié avec des identifiants encore valides reste une porte d'entrée non surveillée sur vos systèmes.
- **Activez le mode confidentialité d'équipe avant tout onboarding.** Il est plus simple de l'imposer dès le départ que de revenir en arrière une fois que plusieurs développeurs ont déjà connecté des dépôts sensibles.
- **Documentez les prompts qui fonctionnent bien pour les tâches récurrentes.** Un prompt qui a donné un bon résultat sur une migration de base de données mérite d'être conservé dans un fichier partagé plutôt que réinventé à chaque fois par chaque développeur.

## Foire aux questions

**Cursor AI est-il gratuit ?**

La formule Hobby est gratuite et ne demande pas de carte bancaire, mais elle limite le nombre de requêtes Agent et de complétions Tab. Un usage professionnel quotidien nécessite en pratique une formule Pro à partir de 20 $/mois.

**Cursor AI est-il disponible en français ?**

L'interface est principalement en anglais, mais le chat et l'agent comprennent et répondent en français sans difficulté particulière. Vous pouvez rédiger vos règles de projet et vos prompts directement en français.

**Quelle est la différence entre .cursor/rules et AGENTS.md ?**

