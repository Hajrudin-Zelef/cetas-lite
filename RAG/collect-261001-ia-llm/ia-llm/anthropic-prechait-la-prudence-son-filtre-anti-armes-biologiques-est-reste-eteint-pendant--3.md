---
id: collect-261001-ia-llm/ia-llm/anthropic-prechait-la-prudence-son-filtre-anti-armes-biologiques-est-reste-eteint-pendant--3
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Hugging Face", "Meta", "Nvidia", "OpenAI", "OpenRouter", "Stripe", "Z.ai"]
dates: []
keywords: ["agent", "agents", "chatgpt", "claude", "glm", "lean", "mai", "muse", "muse spark", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/anthropic-prechait-la-prudence-son-filtre-anti-armes-biologiques-est-reste-eteint-pendant-11-mois.md
source_anchor: ""
source_lines: [109, 140]
sha256: a6cb66f17a13cd661bac5696d48313993596eef148f44e93aca72f24711d23ac
---

# 🧠 **RECHERCHE**

Sept méthodes de raisonnement comparées sur les modèles Qwen2.5 de 1,5 à 7 milliards de paramètres, avec cette fois une contrainte honnête : le même budget de tokens pour toutes. Le gagnant est la méthode la plus bête, résoudre plusieurs fois le même problème de façon indépendante et garder la réponse qui revient le plus souvent. Elle bat l'auto-critique façon Self-Refine ou Reflexion. Traduction pratique : faire réessayer un modèle vaut souvent mieux que lui demander de réfléchir à sa propre réponse.

Quand un agent plante, accuser le modèle envoie souvent vers le mauvais correctif. Scale AI propose de commencer par localiser la première défaillance non rattrapée : le modèle, le contexte, la mémoire, la couche d'outils, un autre agent, l'évaluateur ou l'environnement. Une instruction ignorée peut signifier que le modèle l'a vue et n'a pas su la suivre, ou que la compaction du contexte l'avait déjà effacée. Le papier organise 41 types d'échecs selon leur point d'origine.

La loi d'échelle Chinchilla, qui sert à prédire la performance d'un modèle selon sa taille et son volume de données d'entraînement, paraît quasi parfaite à l'intérieur de la grille où elle a été calibrée, et dérape dès qu'on extrapole au delà. Meta FAIR montre pourquoi : Chinchilla suppose que taille et données agissent séparément, alors que chacune modifie l'utilité de l'autre. Leur correctif, Skaling, ajoute un unique terme de couplage et réduit l'erreur de prédiction d'un facteur 1,5 à 3.

# **🗞️PLUS D'ACTUALITÉS**

En juillet, lors d'un exercice de cybersécurité, un agent autonome d'OpenAI est sorti de l'environnement isolé dans lequel il était censé rester confiné, a accédé à internet et a compromis la plateforme Hugging Face. Ni simulation, ni scénario de film : une machine qui déborde du périmètre qu'on lui avait tracé, pendant un test encadré. The Verge en fait le point de bascule du débat sur les agents autonomes, qui vient officiellement de quitter le rayon science-fiction. À noter que l'équipe d'OpenAI dissoute fin juillet avait précisément pour mission d'anticiper ce type de scénario.

X Square Robot, basée à Shenzhen, a diffusé en direct son système WALL-B en train de trier **1 816 colis par heure avec 98% de précision**, en piochant directement dans un tas de colis en vrac pour les poser sur un tapis en mouvement. En face, trois humanoïdes de Figure AI tournant 200 heures d'affilée plafonnent à environ **1 250 colis par heure à eux trois**. Un seul bras spécialisé, sans jambes, sans tête et sans mains à cinq doigts, fait donc nettement mieux que trois robots à forme humaine. De quoi reposer la question qui fâche : à quoi sert exactement la silhouette humanoïde dans un entrepôt ?

Anthropic a annoncé que tous les modèles Claude vont bientôt marquer le contenu qu'ils génèrent, texte compris, pour se conformer à la réglementation européenne. L'annonce ne disait pas comment. John Gruber révèle le procédé : une forme de stéganographie qui oriente le choix des mots au moment de la génération, de façon à laisser une empreinte statistiquement détectable. Autrement dit, le texte n'est pas simplement tatoué en marge, il est légèrement altéré dans sa formulation même, ce qui contredit la promesse initiale d'un marquage sans effet sur le sens ni la qualité. Le sujet fait beaucoup réagir la communauté tech.

Beaucoup de startups achètent des crédits d'inférence qu'elles ne consomment jamais. Un écosystème de courtiers s'est monté autour de ce constat : les "token brokers" rachètent ces stocks dormants et les revendent avec des remises annoncées **jusqu'à 40%**, en particulier sur les crédits Anthropic. Des places de marché dédiées existent déjà (AI Credits, AICreditMart, CheapCredits, Tokvana, Neokens), certaines affichant une conformité RGPD, et le reste se négocie sur Telegram, Reddit et dans des groupes privés de fondateurs. L'auteur, lui-même actif dans le secteur, s'interroge sur l'origine réelle de certaines remises.

Vous écrivez un problème mathématique en langage courant, l'agent le convertit en théorème Lean 4, écrit une preuve candidate, la compile, corrige ses erreurs et recommence jusqu'à passer. Le tout repose sur un serveur Lean maintenu en mémoire qui compile en **0,4 seconde environ au lieu de 30**, une bibliothèque de théorèmes réutilisables et une recherche de lemmes déjà vérifiés dans Mathlib via leansearch.net et Loogle. Les théorèmes complexes sont découpés en sous-objectifs prouvés en parallèle, et les dépendances entre preuves se visualisent dans un graphe Obsidian. Il faut macOS en arm64 ou Linux x86_64, plus la CLI codex.

Dans un essai de **6 500 mots** intitulé "The Future Is for Everyone", Mark Zuckerberg promet à chacun un agent personnel "exceptionnellement capable", qui comprend vos objectifs et tout ce qui compte pour vous. Deux modèles doivent l'incarner : Glimmer, l'agent léger et permanent, et Muse Spark, le modèle plus puissant et payant censé générer les revenus. Les journalistes de TechCrunch ne mordent pas et rappellent le passif de Meta, quinze ans de promesses de connexion humaine qui ont surtout produit du ragebait et de la publicité ciblée. Leur collègue Russell Brandom résume : ce manifeste est "exactement la raison pour laquelle les gens n'aiment pas l'IA".

Prouver mathématiquement qu'un programme fait exactement ce qu'il annonce passait pour une curiosité de laboratoire, réservée aux systèmes critiques. Le code généré par IA a changé la donne : quand plus personne n'a écrit le programme ligne à ligne, la relecture humaine ne suffit plus comme garantie. L'essai ressort un article de 1979 qui enterrait la vérification formelle, au motif qu'une preuve mathématique ne vaut que par le processus social qui la valide, et teste si l'argument tient toujours face à Lean et aux outils d'aujourd'hui. Will Wilson, d'Antithesis, estime que la discipline est en train de devenir mainstream.

OpenRouter est le point d'entrée unique qui permet d'appeler **plus de 400 modèles d'IA** avec une seule intégration, et de basculer de l'un à l'autre selon le prix ou la tâche, sans dépendre d'un fournisseur unique. Selon Bloomberg, Stripe a finalisé son rachat. La startup, qui revendique 8 millions d'utilisateurs, était valorisée 1,3 milliard de dollars lors de sa série B en mai. Pour tous ceux qui font tourner leurs propres outils dessus, la vraie question est de savoir si le catalogue et la tarification resteront aussi ouverts sous pavillon Stripe.

Nvidia s'associe à Apollo, BlackRock, Blackstone, Brookfield, Goldman Sachs et KKR pour créer des plateformes de financement dédiées aux infrastructures de calcul IA, avec l'objectif de mobiliser **plus de 500 milliards de dollars** de capitaux tiers. Le montage est aussi une opération de langage : Jensen Huang présente désormais le calcul Nvidia comme une classe d'actifs à part entière, que des fonds peuvent acheter et amortir comme une autoroute ou un réseau électrique. Une façon de faire porter par la finance le coût d'une infrastructure que les clients de Nvidia n'ont plus les moyens de payer comptant.

Un site d'agent immobilier. Un site d'association de village. Un logiciel pour apprendre l'espagnol. Un logiciel pour gérer vos finances. Tout écrit en français, sans une ligne de code. C'est la seule compétence IA dont on peut dire ça.

Dans cette nouvelle mise à jour je vous apprends en détail Codex, qui est inclus dans votre abonnement ChatGPT. Si vous n'avez pas chatGPT, je vous montre GLM 5.2, qui est l'équivalent mais gratuit. Sinon si vous avez Claude, ça marche aussi sur Claude Code. Bref, rien à acheter en plus.

Et ce module n'est qu'une partie de la formation.

