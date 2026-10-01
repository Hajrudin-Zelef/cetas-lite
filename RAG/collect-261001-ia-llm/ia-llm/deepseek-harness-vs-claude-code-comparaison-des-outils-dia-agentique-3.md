---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique-3
title: "deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "OpenAI", "vLLM", "xAI"]
dates: []
keywords: ["agent", "claude", "deepseek", "agents", "grok", "sandbox", "vllm"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique.md
source_anchor: ""
source_lines: [199, 299]
sha256: 16404c39977269dab4e241b8cd11a98bacfad45750afabc1ed584e8353c07a02
---

# deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique

Les scores publiés par DeepSeek utilisaient le mode Minimal, et non le mode Standard utilisé dans cette comparaison. Des tests indépendants avec un autre harness ont rapporté des résultats inférieurs. Je ne les ai pas reproduits. Les résultats en mode Minimal n’établissent pas les performances du mode Standard et ne mesurent pas Claude Code.

### Coût : appels d’API vs abonnements Claude

DeepSeek Harness n’a pas de frais de licence, mais les coûts de modèles et d’infrastructure s’appliquent. L’API de DeepSeek utilise des tarifs de pointe et hors pointe. Claude Code est inclus dans les offres Claude payantes. Consultez la page tarifs d’Anthropic avant de budgéter.

L’exécution avec Harness a coûté environ 0,11 $ aux tarifs publics de l’API Anthropic. Claude Code affichait 0,349 $ tout en utilisant des identifiants d’abonnement. Comme ces chiffres proviennent de systèmes de facturation différents, ce n’est pas une comparaison directe de prix.

Les offres individuelles de Claude listent Pro à 20 $ par mois et Max à 100 $ ou 200 $ par mois, tandis que DeepSeek Harness n’a pas d’abonnement équivalent ; c’est le fournisseur de modèle qui facture. Les modèles locaux éliminent les frais d’API mais mobilisent tout de même les ressources de la machine.

Les équipes doivent également prendre en compte les bacs à sable hébergés et le calcul autohébergé lorsque ces services sont facturés en dehors du modèle. Ces coûts n’apparaissent pas dans la licence Harness.

### Stabilité : aperçu développeur vs produit en production

Harness était encore en release candidate lors des tests. Les équipes devraient figer les versions des paquets, car des mises à jour peuvent casser des profils ou des sessions enregistrées ; Claude Code est déjà en production.

## Limites de DeepSeek Harness vs Claude Code

Aucun des deux systèmes ne couvre tous les workflows. Leurs limites proviennent de choix de conception différents.

### Limites de DeepSeek Harness

Harness expose davantage de composants du runtime, ce qui laisse aussi plus de configuration et de tests à la charge de l’utilisateur :

- 
Le projet reste en aperçu développeur, et des mises à jour peuvent casser des profils ou des sessions enregistrées
- 
Les identifiants fournisseur, IDs de modèles et règles d’endpoint nécessitent une configuration manuelle
- 
Son bac à sable `workspace-write` n’a pas démarré dans cette étude de cas Windows et a exigé trois approbations`danger-full-access`
- 
Le SDK Python actuel n’a pas de wheel Windows

Le constat sur le bac à sable vaut uniquement pour l’exécution Windows décrite plus haut. Je n’ai pas répété le test sous Linux ou macOS.

Les notes de version v0.1.0-rc.7 mentionnent une correction de latence Bash persistante et une mise à jour de `node-pty` pour une meilleure compatibilité PTY. Cette observation est donc à la fois spécifique à la version et à Windows.

### Limites de Claude Code

Claude Code prend en charge davantage de décisions de runtime, mais ses frontières fixes empêchent certains types d’expérimentations :

- 
Il exécute des modèles Claude et ne fournit pas de voie pour d’autres familles de modèles
- 
Sa boucle de session intégrée n’est pas exposée comme composant remplaçable
- 
Claude Code nécessite un accès payant ou des crédits d’API facturés séparément, avec des limites d’usage
- 
Les longues sessions peuvent compacter l’ancien contexte ; placez donc les règles persistantes dans `CLAUDE.md`

Ces limites comptent surtout lorsqu’une tâche requiert un autre fournisseur de modèles ou une boucle personnalisée.

## DeepSeek Harness vs Claude Code : lequel choisir ?

Mettez les noms des modèles de côté un instant. Votre tâche consiste-t-elle à terminer le code applicatif, ou à modifier le runtime qui fait le travail ?

Choisissez DeepSeek Harness si le runtime lui-même est le projet. Il convient aux tests multi-fournisseurs, aux boucles ou bacs à sable sur mesure, et aux travaux nécessitant des traces d’exécution détaillées. Son statut de préversion implique aussi d’assumer les changements de version et la validation de configuration.

Choisissez Claude Code si vous travaillez sur une base de code existante et souhaitez que l’agent explore les fichiers, modifie le code et exécute les contrôles du projet avec moins de configuration de runtime. Dans ce test Windows, il a exécuté la suite d’origine sans les trois approbations d’accès total exigées par Harness.

Utilisez les deux si cette répartition reflète votre réalité. Une équipe peut utiliser Claude Code pour le quotidien et Harness pour expérimenter sur les prompts, les outils ou la boucle d’agent.

## Conclusion

Les deux outils savent inspecter un dépôt, modifier des fichiers et exécuter des commandes. La différence tient à la maîtrise du runtime : Harness expose ses composants sous forme de plugins, tandis que Claude Code packe la boucle intégrée et vous permet d’étendre le workflow autour.

Si je devais choisir un point de départ sur la base de ce test, je commencerais par Claude Code pour le travail applicatif courant. Je partirais sur Harness quand il s’agit de modifier ou d’inspecter le runtime, parce que c’est l’objet même de la tâche. Il s’agit d’une étude de cas Windows avec Harness en préversion, pas d’un classement définitif.

Pour aller plus loin, consultez nos articles sur Grok Build, Cursor, OpenCode, Codex et d’autres alternatives à Claude Code.

## DeepSeek Harness vs Claude Code : FAQ

### DeepSeek Harness peut-il réutiliser les instructions de projet de Claude Code ?

**Oui. Harness lit `CLAUDE.md` et `AGENTS.md`, de sorte que les instructions du dépôt peuvent être réutilisées. Les Skills de Claude Code restent séparées.**

### DeepSeek Harness peut-il exécuter Claude Code ou Codex comme sous-agent ?

**Oui. Les builds de préversion incluent des bundles de fournisseur de sous-agent optionnels pour Claude Code et Codex. Figez la version de Harness avant de vous reposer sur leur configuration.**

### Claude Code permet-il de remplacer le modèle par un autre ?

Non, il exécute uniquement Claude. Harness peut héberger Claude, mais Claude Code ne peut pas héberger DeepSeek.

### Quelles plateformes prennent actuellement en charge le SDK Python de DeepSeek Harness ?

**La wheel actuelle prend en charge Linux en x64 ou arm64 ainsi que les versions récentes de macOS arm64. Il n’existe pas de wheel Windows, et l’exemple de composition utilise un accès complet au système de fichiers.**

### Claude Code a-t-il des restrictions de plateforme comme le SDK Python de Harness ?

Pas exactement. Il n’a pas de bac à sable Windows natif, donc il nécessite WSL2 pour `/sandbox` sous Windows, mais la CLI elle-même fonctionne sous macOS, Linux et Windows.

### Le mode Code de DeepSeek Harness peut-il exécuter des programmes Python ?

**Pas avec le backend fourni. Le mode PTC (aussi appelé Code) reconnaît les langages de programmation, mais le runtime documenté propose actuellement un backend TypeScript.**

### DeepSeek Harness peut-il exécuter des modèles locaux ?

**Oui. Ajoutez un fournisseur personnalisé exposant un endpoint compatible OpenAI, tel qu’un serveur local Ollama ou vLLM, puis enregistrez l’ID de modèle dans DeepSeek Harness.**

Je suis ingénieur de données et créateur de communautés. Je travaille sur les pipelines de données, le cloud et les outils d'IA, tout en rédigeant des tutoriels pratiques et percutants pour DataCamp et les développeurs émergents.
