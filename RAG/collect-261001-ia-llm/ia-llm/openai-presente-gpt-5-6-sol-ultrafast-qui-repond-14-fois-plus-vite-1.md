---
id: collect-261001-ia-llm/ia-llm/openai-presente-gpt-5-6-sol-ultrafast-qui-repond-14-fois-plus-vite-1
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Ant", "Anthropic", "Cerebras", "DeepSeek", "Google", "Hugging Face", "Moonshot", "Nvidia", "OpenAI"]
dates: []
keywords: ["agent", "agents", "arr", "attention", "bedrock", "benchmark", "chatgpt", "claude", "deepseek", "fable 5", "fp4", "fp8"]
source: docs/RAG/collect-261001-ia-llm/openai-presente-gpt-5-6-sol-ultrafast-qui-repond-14-fois-plus-vite.md
source_anchor: ""
source_lines: [1, 64]
sha256: 4b3483ab4ee69227afb2249155fdf7da59d8bd467953565b1b3a2d6e2bb0d588
---

# 🧠 **RECHERCHE**

OpenAI ouvre en preview un nouveau palier de son API, baptisé Ultrafast, qui exécute GPT-5.6 Sol jusqu'à **14 fois plus vite** que la version standard, à **750 tokens de sortie par seconde**. Ce n'est pas une version allégée ni distillée : c'est le même GPT-5.6 Sol, avec le même niveau d'intelligence, simplement posé sur un autre type de puce.

**Ce qu'il faut retenir :**

- Le gain ne vient pas du logiciel mais du matériel. Les puces **Cerebras** Wafer-Scale Engine embarquent **44 Go de SRAM** chacune et gardent les poids du modèle en mémoire locale, au lieu d'aller les rechercher dans un stockage externe à chaque token généré. Le goulot d'étranglement classique de l'inférence sur GPU, la bande passante mémoire, disparaît.

- Sur Humanity's Last Exam et ses **2 500 questions**, Ultrafast boucle l'épreuve en **11h11**, contre **78h27** pour Claude Fable 5, selon les mesures publiées par Cerebras.

- Même source : **5,6 fois** plus rapide sur le benchmark GDP-Val, **5 fois** la vitesse d'Opus 4.8 en mode Fast.

- Accès ouvert depuis le **13 août** à un groupe restreint de clients de l'API, avec extension progressive selon la capacité disponible. Ce n'est pas encore dans ChatGPT, et le prix n'est pas communiqué.

- Les cas d'usage mis en avant par OpenAI : réponse aux incidents, support client, analyse de marchés financiers, e-commerce. Tout ce qui supporte mal l'attente.

**Pourquoi ça compte :** depuis deux ans, la course se jouait sur l'intelligence des modèles, et l'attente devant un curseur clignotant faisait partie du décor. Ici, rien ne bouge côté raisonnement, seule la latence s'effondre. C'est précisément ce qui manquait pour qu'un agent IA passe de tâche de fond à outil temps réel : un assistant qui répond en une seconde ne s'utilise pas comme un assistant qui répond en quinze. Accessoirement, le calcul ne vient plus de Nvidia, et c'est le premier déploiement de cette ampleur chez un labo de premier plan.

DeepSeek publie en developer preview **Harness**, le framework open source sous licence MIT sur lequel tournent ses propres agents, code source inclus. Le principe tient dans une équation posée par l'entreprise : « Agent = Model + Harness ». Le modèle apporte l'intelligence, le harness gère tout le reste, c'est-à-dire l'environnement, les outils, les sessions et la capacité à travailler longtemps sans dérailler.

**En détail :**

- Tout est plugin : modèles, outils, compétences, sessions, sandboxes, stockage, boucles, ordonnancement, et jusqu'à l'interface. Un noyau nommé **Cordis** gère le montage, le démontage et les dépendances entre ces briques, qu'on remplace en modifiant la configuration, sans jamais toucher au code source.

- **Quatre modes d'exécution** : Standard (tout l'outillage), Code (l'agent écrit du code pour orchestrer plusieurs tours d'appels d'outils), Minimal (un shell et un éditeur de fichiers, pour évaluer un modèle sans béquilles) et Creator (inspection du runtime en direct, test de plugins en mémoire, fabrication de nouveaux modes).

- Chaque session produit un journal d'événements en ajout seul qui enregistre tout ce que le modèle voit : prompts système, raisonnement, appels d'outils et résultats, planification de sous-agents, injections de contexte. On peut reprendre, dupliquer, chercher et rejouer une exécution à partir de ce même flux.

- **Compatible avec n'importe quel modèle** : DeepSeek, Anthropic, OpenAI, Bedrock, Vertex, Azure et tout endpoint compatible OpenAI.

- Ça s'essaie tout de suite : installez Node.js, lancez `npx @deepseek-ai/dsh web`, et l'interface tourne en local. Le dépôt complet est sur GitHub sous `deepseek-ai/deepseek-harness`.

**Le contexte :** c'est le premier rival open source sérieux de Claude Code, et il arrive le jour où DeepSeek sort aussi son modèle **V4-Pro**, dont les tarifs d'API augmentent le **16 août** avec une tarification différenciée heures pleines et heures creuses en Chine. Le sujet fait beaucoup réagir la communauté tech, moins pour des performances brutes encore non chiffrées que pour ce qu'il signifie : la partie difficile d'un agent, ce n'est plus le modèle, c'est la plomberie autour, et DeepSeek vient de la mettre sur la table.

Google DeepMind sort Gemini 3.7 Flash **trois semaines seulement** après Gemini 3.6 Flash, et le lance à **la moitié du prix** de son prédécesseur. Ce modèle multimodal (texte, image, audio, vidéo) vise le codage, le développement web et les workflows d'agents, avec des progrès nets sur l'ensemble des évaluations publiées.

**Quelques chiffres clés :**

- Tarif d'introduction jusqu'au 31 décembre 2026 : **0,75 $ par million de tokens en entrée**, **3,75 $ en sortie**. Au 1er janvier 2027, on passe à 1,50 $ et 7,50 $.

- Fenêtre de contexte de **1 million de tokens** en entrée, sortie plafonnée à **64 000 tokens**, effort de réflexion réglable pour arbitrer entre qualité, coût et latence. Connaissances arrêtées à mars 2026.

- Face à 3.6 Flash : FrontierCode 1.1 à **43,6 %** contre 34,4 %, DeepSWE v1.1 à **65,3 %** contre 49,0 %, WebDev Arena à **1588 Elo** contre 1538, GDP.pdf (traitement de documents complexes) à **34,0 %** contre 22,0 %, AutomationBench (workflows métier réels) à **30,4 %** contre 17,0 %.

- Ça ne suffit pas partout : sur Terminal-bench, GPT-5.6 Terra garde la tête avec **87,4 %** contre 85,8 %.

- Où l'essayer : Google AI Studio, l'API Gemini, Android Studio, Antigravity, et dès aujourd'hui dans Gemini Spark pour les abonnés Google AI Pro et Ultra de plus de **160 pays**.

Google accompagne le lancement de démonstrations parlantes : un jeu 3D jouable généré depuis un simple prompt texte, avec Nano Banana qui fabrique personnages, objets et textures en temps réel, ou un rapport annuel en PDF transformé en page web interactive avec graphiques vivants.

**Ce que ça change :** le rythme, surtout. Trois semaines entre deux générations, avec un prix divisé par deux et des scores en hausse de dix à quinze points, cela veut dire une chose très concrète pour vous : le coût de l'IA « suffisamment bonne » baisse plus vite que ses performances ne montent. Ce qui coûtait cher à automatiser il y a un mois ne coûte plus grand-chose aujourd'hui.

inclusionAI, le laboratoire IA d'Ant Group, publie Ling 3.0 Flash sous licence MIT, poids téléchargeables sur Hugging Face. C'est un modèle Mixture-of-Experts : **124 milliards de paramètres au total**, mais seulement **5,1 milliards actifs** pour chaque token généré. Autrement dit, le coût de calcul d'un petit modèle avec une bonne partie du savoir d'un gros.

**Les points essentiels :**

- Architecture hybride BailingMoeV3 : **512 experts** dont 8 activés par token plus un expert partagé, attention linéaire Kimi Delta combinée à de la Multi-head Latent Attention, fenêtre de contexte native de **256 000 tokens**.

- **38 points** sur l'Artificial Analysis Intelligence Index, à hauteur de Qwen3.6 27B, mais encore loin du leader ouvert DeepSeek V4 Flash et ses 52 points. Aucun modèle plus petit n'atteint son score.

- Le vrai progrès est ailleurs : sur le test AA Omniscience, le taux d'hallucination tombe de **97 % à 44 %** par rapport à la version précédente. Le modèle refuse désormais bien plus souvent de répondre quand il n'a pas de réponse fiable.

- Gains également nets sur les tâches d'agent, notamment le benchmark t3-Bench Banking.

- Disponible en fp8, fp4, int4 et **GGUF** sur Hugging Face, hébergé sur DeepInfra et l'API inclusionAI, gratuit pour l'instant sur Kilo Code. Coût par token inférieur à Qwen3.6 27B, même s'il consomme davantage de tokens sur les tâches complexes.

