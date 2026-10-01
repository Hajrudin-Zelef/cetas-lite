---
id: collect-261001-general-networking/general-networking/ox-alpha-tombe-le-masque-le-modele-chinois-qui-code-a-prix-casse-1
title: "🧠 **RECHERCHE**"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Google", "Hugging Face", "Nvidia", "OpenAI", "OpenRouter", "Z.ai"]
dates: []
keywords: ["agent", "agents", "agi", "chatgpt", "claude", "embeddings", "gemini", "glm", "gpt-5.6", "mixture of experts", "multimodal", "nvidia"]
source: docs/RAG/collect-261001-general-networking/ox-alpha-tombe-le-masque-le-modele-chinois-qui-code-a-prix-casse.md
source_anchor: ""
source_lines: [1, 101]
sha256: f45f0ac56531d7469930bff5ea4c606bb1adfb4d58b755762f3413e1f2e28815
---

# 🧠 **RECHERCHE**

## **Aujourd'hui:**

🕵️ Ox Alpha révèle son identité et ses 320 milliards de paramètres

🎙️ Gemini 3.5 Transcribe nettoie et structure la parole en temps réel

⚡ Qwen3.8-Flash-Next réduit drastiquement le calcul nécessaire

🤗 Nvidia s’apprête à racheter Hugging Face

🎬 LAION ouvre 80 millions de vidéos à la recherche

🔓 OpenAI explique comment un agent a échappé à son environnement

🧠 Sam Altman prévoit son AGI avant la fin de 2026

🤖 Jetson Orin Nano 2 fait tourner l’IA directement dans les robots

☕ Un café JD.com fonctionne avec un seul robot

🧩 Les assistants IA se perdent dans leurs propres marques

📱 Apple fixe son rendez-vous iPhone et Siri AI au 9 septembre

🖥️ Le Mac mini M5 Pro reste plus puissant que le nouveau M6

🩸 La FDA autorise le premier robot autonome de prise de sang

🏃 Les humanoïdes chinois battent de nouveaux records

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

Le mystérieux Ox Alpha est en réalité **GLM-5.3-Flash**, un modèle de Z.ai conçu pour programmer, manipuler des outils et conduire des automatisations pendant de longues sessions. Il possède **320 milliards de paramètres**, mais n’en active que **18 milliards** à chaque étape, avec des entrées texte, image et vidéo et un contexte dépassant **un million de tokens**.

Le modèle peut analyser un dépôt logiciel, modifier du code, travailler dans un terminal, interpréter des documents ou des graphiques et enchaîner des appels d’outils. Ses poids sont disponibles sous licence MIT, tandis que son API affiche un prix particulièrement agressif.

### **Ce qu’il faut retenir :**

L’architecture Mixture of Experts compte **320 milliards de paramètres, dont 18 milliards actifs** , répartis sur 45 couches et entraînés avec**30 000 milliards de tokens** .
La fenêtre de contexte atteint **1 048 576 tokens** , avec une sortie maximale de**131 072 tokens** , assez pour traiter de grands dépôts ou de longues chaînes d’actions.
Z.ai publie un score de **84,3 sur Terminal Bench 2.1** , contre 85,0 pour Claude Opus 4.8 et 87,4 pour GPT-5.6 Terra. Sur Toolathlon Verified, GLM-5.3-Flash atteint**78,4** , devant Claude Opus 4.8 à 76,2.
OpenRouter facture **0,075 dollar par million de tokens entrants** et**0,25 dollar par million de tokens sortants** .

### **Pourquoi ça compte**

Le prix de l’API rend le modèle immédiatement testable pour des agents de code ou des automatisations volumineuses. L’exécution locale reste cependant réservée à des serveurs très équipés, car publier les poids d’un modèle de 320 milliards de paramètres ne le transforme pas en modèle domestique. Z.ai apporte surtout une nouvelle pression tarifaire sur OpenAI et Anthropic, avec des performances suffisamment proches de leurs meilleurs modèles pour rendre l’écart de prix difficile à ignorer.

Google DeepMind lance **Gemini 3.5 Transcribe**, un modèle speech-to-text propriétaire capable de transcrire plus de **85 langues**, même lorsqu’un utilisateur change de langue au milieu d’une phrase. Il supprime les hésitations, comprend les autocorrections orales, ajoute la ponctuation et peut reconnaître jusqu’à **1 000 expressions personnalisées**.

Le service existe en deux versions. `gemini-3.5-transcribe-live` traite un flux audio par WebSocket avec une latence inférieure à une seconde, tandis que `gemini-3.5-transcribe` analyse des fichiers et ajoute l’identification des locuteurs ainsi que des horodatages mot par mot.

### **En détail :**

Le modèle repose sur Gemini 3 Pro, accepte des entrées audio et texte et dispose d’un contexte de **96 000 tokens** , avec une sortie maximale de**32 000 tokens** .
Google annonce un taux d’erreur moyen de **4,0 % en direct** et de**2,6 % sur les enregistrements** , selon Artificial Analysis.
Le délai nécessaire pour obtenir une transcription finale serait réduit de **70 % par rapport à Chirp 3** . Sur FLEURS, les taux d’erreur annoncés sont de 5,50 % en streaming et 5,04 % hors streaming.
Les fichiers peuvent durer jusqu’à **une heure** , ou 30 minutes lorsque l’identification des locuteurs et les horodatages détaillés sont activés.
La préversion est disponible dans Gemini API et Google AI Studio, avec un palier gratuit puis environ **0,009 dollar par minute en direct** et**0,005 dollar par minute sur fichier** .

### **Ce que ça change**

Un développeur peut désormais créer un outil de sous-titrage, une dictée professionnelle, un compte rendu de réunion ou un agent vocal sans gérer lui-même de modèle audio. Le vocabulaire personnalisé sera notamment utile pour les noms de produits, les termes médicaux ou juridiques et les références alphanumériques. Les poids ne sont toutefois pas téléchargeables, et Google documente encore des risques d’hallucination, d’expiration et de ralentissement.


Alibaba présente **Qwen3.8-Flash-Next**, un modèle multimodal à poids ouverts qui préfigure l’architecture de Qwen4. Son réseau principal contient **125 milliards de paramètres**, mais seulement **6 milliards** travaillent sur chaque token, ce qui réduit fortement le calcul nécessaire pour le raisonnement, le code et les agents.

Le modèle reçoit du texte, des images et des vidéos, puis génère du texte sur un contexte natif de **262 144 tokens**, extensible à un million. Il peut explorer un dépôt logiciel, utiliser des outils, analyser une vidéo ou exécuter de longues tâches bureautiques.

### **Les points essentiels :**

L’architecture comprend **512 experts** , dont 10 sont sélectionnés dynamiquement et un reste partagé. Elle ajoute 51 milliards de paramètres d’embeddings N-gram et 4 milliards consacrés à la prédiction multi-token.
Alibaba revendique un coût d’entraînement équivalent à environ **un neuvième de celui de Qwen3.7-Plus** .
Les résultats publiés atteignent **62,5 sur SWE-bench Pro** , 58,7 sur DeepSWE, 73,9 sur CoWorkBench et**91,7 sur GPQA Diamond** . Aucune évaluation indépendante consolidée n’a encore été trouvée.
L’API QwenCloud coûte **0,15 dollar par million de tokens entrants** et**0,47 dollar par million de tokens sortants** , avec function calling, recherche web et interpréteur de code.

