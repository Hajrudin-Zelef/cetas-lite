---
id: collect-261001-ia-llm/ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026-5
title: "Windows (winget)"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Microsoft", "OpenAI", "xAI"]
dates: []
keywords: ["agent", "agents", "claude", "copilot", "gemini", "grok", "grok 4", "mcp", "model context protocol", "opus 4", "sonnet 5"]
source: docs/RAG/collect-261001-ia-llm/cursor-ai-tutoriel-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [340, 412]
sha256: 236cecf008409fed94b3d47d21458171005e4101e800dc3e6fff76e930f1f285
---

# Windows (winget)

- **Rester en mode Agent en permanence.** Pour une simple question, utilisez le mode*Ask* : il n’écrit rien et évite les modifications non désirées. Réservez l’Agent aux tâches qui doivent modifier des fichiers.
- **Négliger les règles de projet.** Sans`.cursor/rules` , l’IA ignore vos conventions et vous répétez sans cesse les mêmes consignes. Investir cinq minutes dans les règles fait gagner des heures.
- **Donner des consignes vagues.** « Améliore ce code » produit des résultats erratiques. Précisez l’objectif, les contraintes et le fichier concerné.
- **Ignorer l’indexation.** Poser des questions @Codebase avant la fin de l’indexation donne des réponses incomplètes sur les gros dépôts. Attendez la fin de l’indexation.
- **Faire une confiance aveugle au code généré.** L’IA se trompe, invente parfois des API ou introduit des failles. Relisez, testez et utilisez Bugbot avant de fusionner.

## Dépannage : 8 problèmes courants et leurs solutions

Voici les incidents les plus fréquemment rencontrés avec **Cursor** et la marche à suivre pour les résoudre rapidement.

| Problème | Cause probable | Solution | 
|---|---|---|
| « You’ve hit your usage limit » | Crédits d’usage épuisés | Vérifier le tableau de bord, passer à un plan supérieur ou attendre le renouvellement mensuel | 
| Tab ne propose rien | Fonction désactivée ou fichier non pris en charge | Activer Cursor Tab dans les paramètres ; vérifier la connexion Internet | 
| @Codebase donne des réponses incomplètes | Indexation non terminée | Attendre la fin de l’indexation ; relancer l’indexation dans les paramètres | 
| L’agent ne peut pas lancer le terminal | Auto-run désactivé ou permissions | Autoriser l’exécution ; vérifier les paramètres de sécurité de l’agent | 
| Extensions VS Code absentes | Import non effectué | Réimporter la config VS Code via la palette de commandes | 
| Réponses lentes | Modèle lourd sélectionné | Basculer sur Composer 2 ou un modèle « Fast » | 
| Serveur MCP non détecté | Erreur dans mcp.json | Valider le JSON ; vérifier la commande et le chemin ; consulter les logs MCP | 
| Cursor CLI introuvable | PATH non mis à jour | Rouvrir le terminal ou ajouter le binaire au PATH | 

**Piège n°5 :** si Cursor devient lent sur un très gros projet, ajustez votre `.cursorignore` pour exclure les dossiers volumineux (dépendances, artefacts de build, données). Une indexation plus légère améliore nettement la réactivité et la pertinence des réponses.

## Astuces avancées pour les développeurs exigeants

Une fois les bases maîtrisées, ces techniques font passer votre usage de **Cursor** au niveau supérieur.

- **Des règles par contexte.** Créez plusieurs fichiers`.mdc` dans`.cursor/rules/` : un pour le backend, un pour les tests, un pour la documentation. Chaque règle peut s’appliquer à des motifs de fichiers précis (glob).
- **Le mode plan avant l’action.** Demandez d’abord un plan (« propose un plan sans coder »), validez-le, puis lancez l’exécution. Vous gardez le contrôle sur les tâches complexes.
- **Exploiter @Docs.** Indexez la documentation officielle de vos frameworks pour que l’IA cite les API réelles plutôt que d’en inventer – un rempart contre les « hallucinations ».
- **Multi-agents pour explorer des variantes.** Lancez plusieurs agents sur une même tâche avec des modèles différents (Composer 2 vs Claude Opus 4.8) et comparez les diffs.
- **Bugbot systématique.** Activez la revue automatique sur chaque pull request pour attraper les régressions avant la fusion.
- **Cursor CLI en CI.** Intégrez`cursor-agent` dans vos pipelines pour automatiser la revue de code ou la génération de changelog.

Enfin, gardez à l’esprit que l’IA est un copilote, pas un pilote automatique. Les développeurs les plus productifs avec **Cursor AI** ne sont pas ceux qui délèguent tout, mais ceux qui décomposent le travail en tâches claires, vérifient chaque étape et conservent une solide compréhension de leur code. Si vous souhaitez aussi faire tourner des modèles en local pour des raisons de coût ou de confidentialité, notre tutoriel Ollama est un excellent complément.

### Related Coverage

## FAQ : vos questions sur Cursor AI

### Cursor est-il gratuit ?

Oui, Cursor propose une offre gratuite *Hobby*, sans carte bancaire, incluant environ 2 000 complétions par mois selon MobileAppDaily, suffisante pour découvrir l’outil. Les fonctions agentiques intensives nécessitent toutefois une offre payante : Pro à 20 $/mois, Pro+ à 60 $/mois ou Ultra à 200 $/mois, avec 20 % de réduction en facturation annuelle ; côté équipes, Teams se décline désormais en Standard (40 $/utilisateur/mois) et Premium (120 $/utilisateur/mois), selon NoCode MBA.

### Cursor remplace-t-il VS Code ?

Cursor est un IDE autonome dérivé de VS Code : il peut le remplacer entièrement, en important vos extensions, thèmes et raccourcis. Vous n’avez pas besoin d’installer VS Code au préalable ; Cursor est un logiciel indépendant à télécharger sur cursor.com.

### Quels modèles d’IA Cursor utilise-t-il ?

En 2026, Cursor donne accès à son modèle maison Composer 2 ainsi qu’aux grands modèles du marché : Claude Opus 4.8 et Sonnet 5 (Anthropic), GPT-5.5 (OpenAI), Gemini 3.1 Pro (Google) et Grok 4.3 (xAI), entre autres. Vous changez de modèle à la volée selon la tâche.

### Qu’est-ce que le modèle Composer de Cursor ?

Composer est le premier modèle de codage maison d’Anysphere, dévoilé avec Cursor 2.0 le 29 octobre 2025. C’est un modèle « mixture-of-experts » spécialisé, annoncé comme quatre fois plus rapide que les modèles agentiques comparables, qui boucle la plupart de ses tours en moins de 30 secondes.

### Cursor est-il meilleur que GitHub Copilot ?

Ce sont deux philosophies différentes. Copilot est une extension légère intégrée à votre éditeur, idéale pour l’autocomplétion et l’écosystème GitHub. Cursor est un IDE complet centré sur l’IA, avec édition multi-fichiers, mode agent et orchestration multi-agents. Cursor va plus loin dans l’automatisation ; Copilot est plus simple et moins cher à l’entrée.

### Cursor fonctionne-t-il sous Linux et macOS ?

Oui. Cursor est disponible sur Windows (10+), macOS (10.15+, dont Apple Silicon) et Linux (paquets .deb, .rpm et AppImage). Une connexion Internet est nécessaire car les modèles s’exécutent dans le cloud.

### Le code généré par Cursor est-il fiable ?

Le code généré est souvent de bonne qualité, mais il n’est jamais infaillible : l’IA peut introduire des bugs, inventer des API ou des failles de sécurité. Relisez toujours les diffs, écrivez des tests et utilisez Bugbot pour la revue automatisée avant de fusionner.

### Qu’est-ce que MCP dans Cursor ?

Le Model Context Protocol (MCP) est un standard ouvert qui connecte Cursor à des outils externes (bases de données, gestionnaires de tickets, documentation). Un serveur MCP expose des « outils » que l’agent peut appeler. La configuration se fait dans un fichier `mcp.json`. N’installez que des serveurs de sources fiables, car ils exécutent du code sur votre machine.

*Article publié le 02 juin 2026. Les tarifs et versions évoluant rapidement, vérifiez toujours les informations à jour sur la documentation officielle de Cursor.*
