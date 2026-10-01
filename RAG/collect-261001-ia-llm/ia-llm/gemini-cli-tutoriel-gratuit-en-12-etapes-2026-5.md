---
id: collect-261001-ia-llm/ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026-5
title: "Vérifier les versions actuelles"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft"]
dates: []
keywords: ["agent", "agents", "apache", "arr", "claude", "copilot", "distribution", "gemini", "mcp", "open source", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/gemini-cli-tutoriel-gratuit-en-12-etapes-2026.md
source_anchor: ""
source_lines: [363, 425]
sha256: 8c5dd12d7299d1fcce4c1fab8d4d3156c798106d97ba90c5d7c8c6911049c80a
---

# Vérifier les versions actuelles

| Problème | Cause probable | Solution | 
|---|---|---|
| `gemini: command not found` | Binaire npm global absent du PATH | Réinstaller globalement, ou utiliser `npx @google/gemini-cli` | 
| Le navigateur ne s’ouvre pas (SSH) | Machine distante sans interface graphique | Basculer sur l’authentification par `GEMINI_API_KEY` | 
| Erreur 429 / quota dépassé | Limite de requêtes atteinte | Attendre, utiliser `/compress` , ou passer à une clé API facturée | 
| Blocage derrière un proxy d’entreprise | Trafic HTTPS filtré | Définir `HTTPS_PROXY` et`HTTP_PROXY` | 
| Erreur de version Node | Node < 20 | `nvm install 22 && nvm use 22` | 
| Serveur MCP absent de `/mcp` | Chemin ou syntaxe `settings.json` erronés | Vérifier le JSON, relancer la CLI, tester `/mcp` | 
| « docker not found » en sandbox | Moteur de conteneurs non installé | Installer Docker/Podman, ou désactiver le sandbox | 
| Réponses lentes ou hors-sujet | Contexte saturé | `/compress` puis`/clear` ; réduire la portée de`@` | 

Deux conseils transverses. Si un comportement vous semble anormal, tapez `/about` pour connaître votre version exacte et `/bug` pour ouvrir un rapport sur GitHub avec le contexte pré-rempli. Et pensez à mettre à jour régulièrement : `npm update -g @google/gemini-cli`. Le rythme de publication est soutenu depuis le lancement – la CLI avait déjà atteint la v0.7.0 dès le 22 septembre 2025 au fil des mises à jour hebdomadaires, puis la v0.20.0 du 1er décembre 2025 avait ajouté le glisser-déposer multi-fichiers – et de nouveaux correctifs continuent d’arriver chaque semaine sur le canal stable.

## Astuces avancées pour aller plus loin

Une fois les bases maîtrisées, plusieurs fonctionnalités avancées vous distingueront. Elles transforment Gemini CLI d’un assistant ponctuel en un véritable copilote d’ingénierie sur mesure.

- **Commandes slash personnalisées.** Créez vos propres commandes dans des fichiers`.toml` placés dans`~/.gemini/commands` (globales) ou`.gemini/commands` (projet). Une commande`/test` peut par exemple encapsuler un prompt de génération de tests complet, réutilisable en une frappe.
- **Sous-agents et compétences (skills).** Les commandes`/agents` et`/skills` permettent de déléguer des tâches spécialisées à des agents dédiés et de charger des « compétences » réutilisables, pour orchestrer des workflows complexes.
- **Intégration IDE.** Via`/ide` , connectez Gemini CLI à VS Code pour partager le contexte de l’éditeur (fichiers ouverts, sélection) avec l’agent du terminal.
- **Gestion fine du contexte.** Sur les très longues sessions, alternez`/compress` (résumé) et`/chat save` (point de contrôle nommé) pour reprendre exactement là où vous vous étiez arrêté.
- **Mode plan systématique.** Prenez l’habitude de démarrer toute tâche risquée par`/plan` : vous validez la stratégie avant l’exécution, un réflexe qui évite bien des mauvaises surprises.

Ces mécanismes d’extension – commandes `.toml`, MCP, sous-agents, hooks de cycle de vie – font de Gemini CLI une plateforme d’automatisation plutôt qu’un simple outil. Les équipes les plus avancées versionnent leur configuration complète (`GEMINI.md`, `settings.json`, commandes personnalisées) dans un dépôt partagé, garantissant un comportement homogène de l’agent sur tous les postes. Pour une vue d’ensemble de l’écosystème des outils de codage IA, notre dossier sur le boom du vibe coding en 2026 replace Gemini CLI dans le paysage plus large.

## Gemini CLI face à la concurrence : quel agent choisir ?

Gemini CLI n’évolue pas seul. Le marché des agents de codage en terminal s’est densifié en 2025-2026, avec Claude Code d’Anthropic, GitHub Copilot (et son interface CLI), Cursor côté éditeur, ou encore des solutions open source comme Cline et Aider. Comment se situe l’outil de Google ?

Sa force cardinale reste le **rapport puissance-prix** : aucun concurrent n’offre un palier gratuit aussi généreux (1 000 requêtes par jour, un million de tokens de contexte) doublé d’un code entièrement ouvert. Pour un étudiant, un indépendant ou une startup européenne qui débute avec l’IA agentique, c’est un point d’entrée imbattable. La nature open source et la compatibilité Vertex AI en font aussi un choix rationnel pour les organisations soucieuses de gouvernance et de souveraineté des données.

En contrepartie, certains développeurs jugent Claude Code plus abouti sur les tâches d’ingénierie logicielle très complexes, tandis que Cursor séduit ceux qui préfèrent une expérience intégrée à l’éditeur. La bonne approche consiste souvent à **combiner les outils** : Gemini CLI pour l’exploration, l’automatisation et le travail par lots à coût nul ; un agent premium pour les refactorisations critiques. Nos guides Cursor en 12 étapes et GitHub Copilot en 12 étapes vous aideront à monter en compétence sur les alternatives et à composer votre propre pile.

### À lire également sur Tech Insider

## Foire aux questions sur Gemini CLI

### Gemini CLI est-il vraiment gratuit ?

Il l’a été jusqu’au 18 juin 2026 : la connexion avec un compte Google personnel donnait alors 60 requêtes par minute et 1 000 requêtes par jour sans carte bancaire, avec une fenêtre de contexte d’un million de tokens. Depuis cette date, Google réserve ce niveau de service aux formules payantes Google AI Pro (19,99 $/mois, 1 500 requêtes par jour) et Ultra ; les comptes Google personnels gratuits ne donnent plus un accès complet. Le code de l’outil reste toutefois open source (Apache 2.0). Pour des quotas plus élevés ou un accès prioritaire aux modèles Pro, vous pouvez activer la facturation sur une clé API (Gemini 3.1 Pro Preview facturé environ 2,00 $ par million de tokens en entrée) ou passer par Vertex AI / Code Assist Enterprise (54 $/utilisateur/mois, 2 000 requêtes par jour).

### Quelle version de Node.js faut-il pour Gemini CLI ?

La version minimale requise est Node.js 20.0.0. Nous recommandons une version LTS active (Node 22 ou 24) pour la stabilité et la sécurité. Le gestionnaire `nvm` est la façon la plus simple d’installer et de basculer entre versions sans problème de permissions.

### Quelle différence entre Gemini CLI et Gemini Code Assist ?

Gemini CLI est un agent autonome qui vit dans le terminal et exécute des actions (fichiers, shell, web). Gemini Code Assist est l’assistant intégré aux IDE (complétion, chat en contexte d’édition). Les deux partagent la technologie Gemini et peuvent se compléter ; Gemini CLI se distingue par son autonomie agentique et sa scriptabilité.

### Mes données de code sont-elles envoyées à Google ?

Les requêtes sont traitées par les modèles Gemini côté Google. Les conditions de confidentialité varient selon le mode : compte gratuit, clé API payante ou Vertex AI (qui offre des garanties de traitement en entreprise, avec choix de régions, y compris dans l’UE). Pour un projet sensible, privilégiez Vertex AI et consultez les politiques applicables via la commande `/privacy`.

### Gemini CLI fonctionne-t-il sous Windows ?

Oui, nativement via Node.js pour Windows, mais l’expérience la plus fluide passe souvent par WSL 2, car l’agent exécute des commandes shell de type POSIX. Installez Node 22 LTS dans votre distribution WSL et suivez les instructions Linux de ce tutoriel.

### Comment empêcher l’agent de modifier des fichiers sans mon accord ?

Par défaut, Gemini CLI demande une confirmation avant toute action à effet de bord. Ne désactivez pas ces confirmations (évitez le mode auto-approbation). Pour une sécurité renforcée, lancez l’agent avec `--sandbox` afin d’isoler l’exécution dans un conteneur, et utilisez `/plan` pour un mode lecture seule.

### Qu’est-ce que le protocole MCP et pourquoi l’utiliser ?

