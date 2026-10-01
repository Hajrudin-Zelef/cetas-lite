---
id: collect-261001-ia-llm/ia-llm/une-ia-vient-de-traduire-des-tablettes-vieilles-de-plus-de-4000-ans-1
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "CISA", "Google", "Hugging Face", "Meta", "Moonshot", "OpenAI", "Stripe", "Z.ai"]
dates: []
keywords: ["agent", "agents", "benchmark", "chatgpt", "claude", "gemini", "glm", "kimi", "open source", "open weights", "opus 5"]
source: docs/RAG/collect-261001-ia-llm/une-ia-vient-de-traduire-des-tablettes-vieilles-de-plus-de-4000-ans.md
source_anchor: ""
source_lines: [1, 81]
sha256: 06405ede7a5b0fe26b6cd667f10e98256c6c754ee516d955274150edece329ee
---

# 🧠 **RECHERCHE**

## **Aujourd'hui:**

📜 Une IA traduit le cunéiforme dans les deux sens, en libre accès

🇨🇳 GLM-5.3 prend la tête des modèles ouverts, mais les poids sont repoussés

🌀 Stripe invoque la singularité pour rester hors de la bourse

🔒 Anthropic refuse de sortir « Model 2 », plus puissant que Claude

🏺 Les rouleaux d'Herculanum livrent des livres inconnus de Philodème

🤖 Des bras robotiques qui apprennent en regardant une vidéo

🧬 Claude conçoit des protéines avec 35% de réussite

🎓 Google transforme Search et Gemini en salle de révision

💧 Les watermarks invisibles de Claude contournés en quatre heures

⚡ Codex avale en deux semaines une migration prévue sur cinq ans

🍎 Meta AI débarque sur Mac

🩸 La FDA approuve un robot qui fait les prises de sang

🛡️ La NSA alerte sur des exploits industriels écrits par IA

☢️ Le réacteur de Bill Gates se branche sur les data centers IA

Un site d'agent immobilier. Un site d'association de village. Un logiciel pour apprendre l'espagnol. Un logiciel pour gérer vos finances. Tout écrit en français, sans une ligne de code. C'est la seule compétence IA dont on peut dire ça.

Dans cette nouvelle mise à jour je vous apprends en détail Codex, qui est inclus dans votre abonnement ChatGPT. Si vous n'avez pas chatGPT, je vous montre GLM 5.2, qui est l'équivalent mais gratuit. Sinon si vous avez Claude, ça marche aussi sur Claude Code. Bref, rien à acheter en plus.

Et ce module n'est qu'une partie de la formation.

✅ Formation complète sur les IA (LLM, IA Marketing, génération d’images, sons et voix, automatisation avec n8n, premiers pas avec les agents, etc.)

✅ Leçon complète pour créer vos propres agents IA (n8n) et automatiser sans limites

✅ Leçon complète sur Claude Code, meilleur outil IA en 2026

✅ Accès à un réseau de professionnels qualifiés

✅ Des mises à jour régulières pour rester à la pointe de l’IA

➡️ Si vous rejoignez maintenant, vous verrouillez à vie le tarif de 49€ (paiement unique).

Peu importe jusqu’où le prix grimpera, et il atteindra bientôt 100€ ou plus, vu la demande, vous ne paierez pas un centime de plus.

Un seul paiement. Accès à vie.

Sur ce dernier point, la preuve est sous vos yeux : le module vibe coding datait de fin 2025, je l'ai supprimé et refait de zéro, et les inscrits n'ont pas payé un centime de plus. C'est ce que veut dire « mises à jour incluses ». Le prix montera, mais jamais pour ceux qui sont déjà entrés.

Plus de 13 000 personnes s'y forment déjà. Leurs avis sont sur la page.

On se retrouve à l’intérieur.

Un chercheur de l'USC Viterbi School of Engineering, Zhaohui Wang, a publié TabletCraft, un modèle de traduction entraîné sur **116 000 exemples** d'akkadien vers l'anglais. Il atteint un score BLEU de **49,1**, l'indicateur standard de qualité en traduction automatique, contre 37,5 pour le meilleur système précédent. Surtout, il fait quelque chose que personne n'avait publié avant : il traduit aussi dans l'autre sens, de l'anglais vers l'akkadien, et rend le résultat en signes cunéiformes.

Architecture basée sur **ByT5** , un modèle de traduction spécialisé et non un grand modèle de langage généraliste. Le nombre de paramètres n'a pas été publié.
Validé sur le jeu de test Akkademia (**2 812 exemples** ) :**49,1** de BLEU en akkadien vers anglais,**48,5** dans le sens inverse, une première.
Un convertisseur de signes intégré couvre **14 240 correspondances** , soit**95,3%** du répertoire, et affiche la traduction en signes Unicode puis en image de tablette d'argile virtuelle.
Publié en open source sous licence Creative Commons BY 4.0, avec code, modèle pré-entraîné, interface en ligne de commande et démo web sur GitHub. L'installation tient en une commande : `pip install cuneiscribe` .
Le travail a été présenté en juillet 2026 et accepté au workshop C3NLP de la conférence ACL 2026, avec un papier détaillé.

Environ **500 000 tablettes** cunéiformes dorment dans les réserves des musées du monde, et seule une fraction infime a été traduite. Le goulot d'étranglement n'a jamais été technique : il tient au fait que quelques centaines de personnes seulement savent encore lire couramment ce système d'écriture. Un modèle installable en une ligne de commande sur une machine personnelle change la nature du problème. N'importe quel curieux, professeur ou étudiant peut désormais interroger un texte vieux de cinq millénaires sans passer par un assyriologue. Le sens inverse, écrire en cunéiforme, relève davantage du plaisir que de la science, mais il en dit long : l'outil s'adresse au grand public autant qu'aux chercheurs.

Z.ai, la startup chinoise anciennement connue sous le nom de Zhipu AI, a lancé GLM-5.3 le 14 août. Le modèle décroche **60 points** à l'Artificial Analysis Intelligence Index, à égalité avec Kimi K3 en tête des modèles ouverts et **7 points** devant son prédécesseur. Mais la publication des poids, c'est à dire la partie qui rend un modèle réellement ouvert, a été repoussée : GLM-5.3 s'est révélé trop efficace pour détecter des failles de sécurité dans du code.

Même architecture de base que GLM-5.2, les paramètres ne sont pas publiés. Les gains viennent entièrement d'un post-entraînement amélioré. Le modèle est orienté texte et code au lancement, sans nouvelle modalité.
Score Elo de **1770** sur le benchmark agentique GDPval-AA v2, contre 1524 pour GLM-5.2. C'est la deuxième place mondiale, derrière le seul Claude Opus 5 (1855), qui est propriétaire.
**0,68 $ par tâche** , contre 0,84 $ pour Kimi K3, son rival ouvert direct, soit**19% moins cher** . GLM-5.2 restait à 0,44 $.
Accessible dès maintenant via l'API Z.ai, sur un endpoint compatible OpenAI, et via le GLM Coding Plan à **18 $ par mois** en Lite, 80 $ en Pro, 168 $ en Max, avec une remise annuelle qui descend le Lite à 12,60 $. La variante`glm-5.3[1m]` ouvre une fenêtre de contexte de**1 million de tokens** , utilisable dans Claude Code.
Les poids ouverts sont attendus vers le 28 août, réservés pour l'instant à des partenaires sécurité. Aucun dépôt Hugging Face officiel, donc aucune installation locale possible à ce stade.

Pour la première fois, un modèle facturé moins cher que ses concurrents ouverts talonne le meilleur modèle propriétaire du marché sur les tâches d'agents. Mais le vrai signal est ailleurs : un laboratoire chinois vient de retarder la publication de ses poids parce que son modèle trouve trop bien les vulnérabilités logicielles. La promesse « open weights égale liberté totale » rencontre sa première limite pratique, et elle ne vient pas de Washington ni de Bruxelles.


Dans une lettre envoyée mercredi à ses investisseurs, Stripe fixe au **1er janvier 2026** le « début de la singularité » et en fait un argument pour ne pas entrer en bourse. L'entreprise reconnaît dans le même document que le terme est « flou et peut-être déjà galvaudé ». Toute la presse tech a repris l'histoire dans la journée.

