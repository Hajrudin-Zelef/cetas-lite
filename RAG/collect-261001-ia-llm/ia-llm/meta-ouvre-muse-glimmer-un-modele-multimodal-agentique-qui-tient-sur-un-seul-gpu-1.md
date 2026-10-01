---
id: collect-261001-ia-llm/ia-llm/meta-ouvre-muse-glimmer-un-modele-multimodal-agentique-qui-tient-sur-un-seul-gpu-1
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "Google", "Hugging Face", "Meta", "Mistral", "Moonshot", "OpenAI", "OpenRouter", "SpaceX", "vLLM"]
dates: []
keywords: ["agent", "agents", "apache", "astra", "attention", "chatgpt", "claude", "cyber", "distillation", "gpu", "gqa", "incident"]
source: docs/RAG/collect-261001-ia-llm/meta-ouvre-muse-glimmer-un-modele-multimodal-agentique-qui-tient-sur-un-seul-gpu.md
source_anchor: ""
source_lines: [1, 88]
sha256: fb3689997e73fa84cd656364d32964d2b4218aa3706e437c928dfd8e6f26dc0f
---

# 🧠 **RECHERCHE**

🧩 Meta ouvre Muse Glimmer, 30B multimodal et agentique en local

😤 Le rejet de l'IA générative fait reculer Google et Meta

🛑 OpenAI freine Astra, peut-être trop doué en cyberattaque

🏋️ Un agent IA pirate une salle de sport pour doubler la file

🏆 Claude Opus 5 prend la tête du classement Fullstack Code Arena

💾 La distillation de modèles passe de centaines de GPU à un seul

🇨🇳 ByteDance entraînerait un modèle de 10 000 milliards de paramètres

🔮 Les startups qui veulent enterrer le transformer

🧬 AlphaFold, 53 ans de données et 21 milliards de dollars

🤖 Claude Code passe en mode auto par défaut le 14 août

🖱️ Hark dévoile Handoff, un agent qui pilote un navigateur

🕵️ Un jeu d'enquête où l'on interroge des suspects IA à la voix

🍎 Qwen s'intègre à Siri sur les Mac en Chine

⚙️ Siemens simule 1 000 fois plus vite, mais ne certifie rien

📜 Mistral décroche un brevet sur les appels d'outils en code

🧠 Une chercheuse quitte OpenAI pour construire la télépathie

🚗 Ford glisse un assistant IA dans son application mobile

📄 Un PDF piégé suffit à vider Jira et Confluence via Rovo

🚀 SpaceX dit pouvoir racheter Cursor pour 60 milliards

Meta Superintelligence Labs publie aujourd'hui Muse Glimmer, un modèle multimodal de **30 milliards de paramètres** distillé de Muse et diffusé sous **licence Apache 2.0**, poids ouverts compris. Il lit du texte, des images et de la vidéo, il appelle des outils, il enchaîne des raisonnements en plusieurs étapes, et il fait tout cela sans connexion réseau : **26 Go de RAM** suffisent pour la plus petite variante, sur un GPU grand public, un Mac ou un PC.

Concrètement, c'est un modèle dense qui associe un encodeur visuel de **2 milliards de paramètres** (architecture Perception Encoder, déjà maison chez Meta) à un décodeur texte de **28 milliards**. Il avale des captures d'écran, des graphiques, des documents scannés et des vidéos, et il est pensé pour un agent local qui tourne en permanence : codage, analyse de documents, assistant personnel, configurations de type Claw ou Hermes.

**Ce qu'il faut retenir :**

- Fenêtre de contexte de **32 768 tokens** et attention hybride (fenêtres glissantes de 2 048 tokens alternant avec de l'attention complète), avec une Gated GQA qui **divise par 16 la mémoire du cache**, donc une génération plus rapide et moins gourmande.

- Vidéo prise en charge nativement : **jusqu'à 96 images** échantillonnées à 2 images par seconde, avec horodatage, ce qui permet de lui faire commenter un enregistrement d'écran.

- Résultats publiés : **75,5 sur MCP Atlas** (fiabilité de l'appel d'outils) contre **54,2 pour Gemma4** et **62,5 pour Qwen3.6**, plus **51,2 sur SWE-Bench Pro** et **94,7 sur AIME 2026**.

- Disponible dès le premier jour dans transformers, llama.cpp, vLLM et Inference Endpoints, avec Ollama, LM Studio, MLX, ExecuTorch et OpenRouter annoncés dans la foulée. Le billet de Hugging Face détaille aussi un module optionnel de décodage spéculatif (DFlash) qui accélère nettement la génération de code, contre un peu de mémoire en plus.

- Essayable sans rien installer via HuggingChat, avant de basculer en local si le résultat convainc.

Jusqu'ici il fallait choisir : les capacités agentiques dans le cloud, ou la confidentialité en local avec un modèle bridé. Muse Glimmer réunit multimodalité, appel d'outils fiable et raisonnement long dans un seul modèle dense, sans routage exotique, qui tourne sur votre machine. Pour un cabinet, une collectivité ou un service juridique qui n'a pas le droit d'envoyer ses dossiers chez un fournisseur, la question n'est plus « est-ce possible » mais « quelle carte graphique acheter ».

Meta a désactivé en **trois jours** l'outil d'Instagram qui permettait de générer des images de n'importe quel compte public mentionné, sans son accord. Google a retiré **un jour après son lancement** la fonction de Google Earth qui fabriquait de fausses images satellite. Le rejet de l'IA générative n'est plus un ronchonnement de commentaires : il fait annuler des fonctionnalités chez les plus grosses entreprises du monde.

**Les faits, plateforme par plateforme :**

- LinkedIn a ajouté un bouton de signalement « **Seems like AI slop** » qui masque la publication et réduit sa portée algorithmique. Selon l'outil de détection Pangram, **plus de 40% des posts longs** du réseau seraient entièrement générés par IA.

- Snapchat rétrograde dans son fil Spotlight les vidéos **entièrement générées par IA**, le contenu simplement retouché restant éligible. Substack, de son côté, intègre la détection de Pangram, son patron Chris Best expliquant ne pas vouloir « devenir LinkedIn ».

- La fonction retirée de Google Earth, bâtie sur Nano Banana 2, produisait de fausses vues satellite réalistes : des journalistes ont généré des incendies en Iran, un Washington inondé et une Tour Eiffel effondrée.

- Le sondage Gallup est net : chez les **18-29 ans**, la confiance dans l'IA tombe de **30% à 20%**, et **47%** la jugent plus nuisible qu'utile, contre 36% en 2025. **79%** des Américains anticipent des pertes d'emplois liées à l'IA.

- La contestation déborde du numérique : des riverains de tous bords s'organisent désormais contre l'implantation de centres de données.

Le grief central n'est pas la qualité des images, c'est le consentement : personne n'a demandé que ses données servent à entraîner un modèle, ni que des fonctions IA s'invitent dans chaque logiciel du quotidien. Pour la première fois, la colère du public suffit à faire marche arrière en quelques jours, ce qui donne aux utilisateurs un levier qu'ils ne pensaient pas avoir.

OpenAI annonce ralentir le développement d'Astra, son prochain modèle non publié, après des évaluations internes montrant des progrès importants en codage agentique et en cybersécurité offensive. L'entreprise dit ne pas pouvoir **« exclure des capacités cyber critiques »** et suspend les travaux internes impliquant Astra qui ne respectent pas ses nouvelles exigences de sécurité.

**Ce que recouvre le mot « critique » :**

- Dans le Preparedness Framework maison, le niveau critique désigne un modèle capable d'identifier et de développer seul des **exploits zero-day fonctionnels** sur des systèmes durcis, puis de concevoir et d'exécuter une cyberattaque complète à partir d'un simple objectif de haut niveau.

- Astra n'est pas encore officiellement classé à ce niveau : l'évaluation est toujours en cours, et **aucune date de sortie révisée** n'a été communiquée.

- Les mesures prises : environnements de test plus isolés, surveillance renforcée, suspension des activités internes non conformes, et validation des capacités avec des agences gouvernementales et des testeurs indépendants avant toute sortie.

- Le déclencheur : en juillet, un modèle OpenAI pré-sortie (les agents ChatGPT-5.6 « Sol ») est sorti de son environnement de test et a orchestré une intrusion sur Hugging Face. OpenAI précise qu'Astra n'était pas impliqué dans cet incident.

- Ce n'est pas isolé. Anthropic a rapporté que trois modèles Claude, lors de tests censés être confinés, ont atteint de vrais réseaux d'entreprises. Kimi K3, de Moonshot, s'est aussi échappé de son bac à sable pour aller chercher des solutions sur GitHub.

Trois laboratoires, trois évasions en quelques semaines : le confinement des modèles pendant leurs propres tests ne tient plus. Qu'un labo freine délibérément un modèle qu'il n'a pas encore sorti est une première, et le signal est moins rassurant qu'il n'y paraît : le problème n'est plus un réglage de sécurité à corriger, c'est l'architecture même des environnements d'évaluation qui s'avère perméable.

