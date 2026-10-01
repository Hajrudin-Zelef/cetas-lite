---
id: collect-261001-ia-llm/ia-llm/google-prend-du-retard-face-a-la-concurrence-gemini-3-5-pro-repousse-une-3e-fois-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Hugging Face", "Meta", "Nvidia", "OpenAI", "Perplexity", "Samsung", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "agents", "apache", "arr", "benchmark", "chatgpt", "claude", "diffusion", "fine-tuning", "glm", "gpt-5.6", "gpt-live"]
source: docs/RAG/collect-261001-ia-llm/google-prend-du-retard-face-a-la-concurrence-gemini-3-5-pro-repousse-une-3e-fois.md
source_anchor: ""
source_lines: [85, 153]
sha256: 0df8111223154cddf2b4372fa5f9871c72ae54e112f864416d69ab8e55eba6cd
---

# 🧠 **RECHERCHE**

**Le contexte :** Anthropic invoque une demande « difficile à gérer » et le coût de sa capacité serveur. Mais le calendrier trahit la vraie pression : **GPT-5.6 Sol** d'OpenAI offrirait des performances comparables pour environ **un tiers du prix**, et les modèles chinois cassent les tarifs. Signe des tensions sur le calcul, Anthropic serait même en discussion pour **louer la capacité de Meta** jusqu'à 10 milliards de dollars. Résultat : le modèle le plus capable du marché devient aussi le plus cher et le moins accessible, au risque de faire fuir la communauté qui l'a adopté.

# 🧠 **RECHERCHE**

### **Colibrì fait tourner un modèle de 744 milliards de paramètres sur un PC avec 25 Go de RAM**

Colibrì est un runtime open source qui exécute **GLM-5.2**, un modèle MoE de **744 milliards** de paramètres, sur du matériel grand public. L'astuce : chaque token n'active qu'environ **40 milliards** de paramètres. Le runtime garde en RAM les tenseurs denses quantifiés en int4 (~9,9 Go) et laisse les **experts** (~370 Go) sur un SSD NVMe, chargés à la demande. C'est **lent, mais ça fonctionne**, et ça met un modèle de pointe à portée d'un PC costaud.

**Grok 4.5 bat Claude Opus 4.8 comme orchestrateur, pour deux fois moins cher**

Sur le benchmark **WANDR** (tâches de recherche complexes en plusieurs étapes) au sein de Perplexity Computer, **Grok 4.5** décroche le meilleur score (**0,328**) pour **4,76 $** l'essai. Il devance **Opus 4.8** (0,254 pour 9,46 $) et **GPT-5.6** (0,289 pour 2,64 $). Un bon rappel que, pour l'orchestration multi-agents, le meilleur rapport performance/prix ne vient pas toujours du modèle le plus prestigieux.

**NVIDIA et Hugging Face ouvrent le fine-tuning des modèles vidéo à grande échelle**

Les deux acteurs publient une intégration **open source** (Apache 2.0) entre **NeMo Automodel** et **Diffusers**. On peut désormais entraîner et fine-tuner des modèles de diffusion image et vidéo (**FLUX.1-dev, Wan 2.1, HunyuanVideo**) sans conversion de checkpoint, et passer d'un seul GPU à des centaines en changeant une simple ligne de configuration. De quoi démocratiser l'entraînement de modèles génératifs vidéo.

**Planifier un week-end entier en parlant à GPT-Live**

Ce guide montre comment se servir de **GPT-Live**, le nouveau modèle vocal derrière ChatGPT Voice, pour organiser un voyage de vive voix : on l'interrompt, on réoriente ses recherches en temps réel, puis on transforme la conversation en plan d'action (hébergement, itinéraire, ordre des réservations). Méthode en six étapes. **GPT-Live-1** est réservé aux comptes payants, une version mini est offerte aux gratuits.


# **🗞️PLUS D'ACTUALITÉS**

**xAI passe son agent de codage Grok Build en open source**

xAI a publié sous licence **Apache 2.0** le code **Rust** de Grok Build, son agent de codage en terminal : boucle d'agent, outils fichiers, exécution shell, recherche web, interface. On peut le compiler soi-même et y brancher une **inférence locale** via config.toml. Le dépôt a déjà dépassé **1 900 étoiles**, même si les contributions externes restent fermées.

**Amazon Zoox rappelle 105 robotaxis après un passage en pleine fumée**

Un robotaxi Zoox sans passager a foncé dans une **scène d'incendie enfumée** à Las Vegas le 20 juin avant de freiner et s'arrêter. La filiale d'Amazon rappelle **105** véhicules pour corriger le bug logiciel qui les empêchait de détecter la fumée dense. La NHTSA vient justement d'exiger des constructeurs autonomes qu'ils règlent d'ici fin juillet les interférences avec les secours.

**Grève chez Hyundai contre les robots humanoïdes**

Des milliers d'ouvriers de l'usine d'Ulsan (Corée du Sud) ont débrayé après l'échec des négociations sur le déploiement de robots **Atlas** de Boston Dynamics. C'est le premier arrêt de travail de l'automobile lié à la peur des humanoïdes : Hyundai veut en déployer **plus de 25 000**. Chaque Atlas coûte environ **130 000 $**, rentabilisable en deux ans, avec un coût d'exploitation bientôt inférieur au salaire minimum fédéral américain.

**Des infirmières de Kaiser dénoncent une IA qui note leur empathie**

Chez Kaiser Permanente, un logiciel d'IA mesure la durée des appels et évalue jusqu'au **ton de voix** et à l'empathie des infirmières. Au-delà de **15 minutes** de conversation, même pour un patient suicidaire, la sanction guette. Le syndicat négocie un contrat pour 25 000 soignants pendant que la Californie étudie des lois de protection. (495 points sur Hacker News.)

**TikTok teste un outil de détection des ressemblances générées par IA**

TikTok expérimente un outil **opt-in** qui laisse les créateurs signaler les vidéos exploitant leur apparence sans autorisation. Le test est limité à certains créateurs américains, avec vérification d'identité (selfie en temps réel plus pièce d'identité via **Jumio**). YouTube propose déjà un dispositif similaire à tous ses utilisateurs adultes.

**Patreon bloque désormais activement les robots d'IA**

Fini le fichier robots.txt poliment ignoré : Patreon s'appuie sur **Cloudflare** (AI Crawl Control) pour bloquer les crawlers qui entraînent des IA sur le travail des créateurs. Les tentatives d'accès hebdomadaires seraient tombées de **plusieurs milliers à zéro**. Les robots d'indexation qui renvoient du trafic, eux, restent autorisés.

**Claude Code : anatomie d'une fonctionnalité ratée**

Anthropic avait glissé dans Claude Code (v2.1.198) un **minuteur de 60 secondes** au bout duquel l'agent continue seul si l'humain ne répond pas, sans le mentionner au changelog. L'auteur pointe le manque de transparence et le risque d'actions non supervisées. Anthropic a corrigé le comportement quelques jours après sa mise au jour. (139 points sur Hacker News.)

**La faim de mémoire des data centers IA fait grimper le prix des smartphones**

Samsung, SK Hynix et Micron réorientent leur production vers la mémoire **HBM** des accélérateurs IA, bien plus rentable. Résultat en Inde, 2e marché mondial : les expéditions de smartphones ont chuté de **10 %** au 2e trimestre, la pire baisse en six ans. Le segment sous 150 $ s'effondre de **45 %**, et les gens gardent leur téléphone plus longtemps.

**L'UE force Google à ouvrir Android aux assistants IA rivaux**

Au nom du **DMA**, la Commission européenne ordonne à Google d'ouvrir **11 fonctions** d'Android (recherche vocale, réservations) aux assistants IA concurrents d'ici juillet 2027, et de partager des données de recherche anonymisées dès janvier 2027. Google prévient que ces règles pourraient créer des risques de confidentialité et de cybersécurité.

**Le Pentagone : la lenteur d'adoption de l'IA plus risquée qu'un alignement imparfait**

La Marine américaine signe une stratégie pour une flotte **« IA-first »**, avec des grands modèles de langage tournant **directement sur les navires** de guerre et un « conseil de guerre IA » pour prioriser les scénarios de mission. Message assumé de la doctrine : tarder à adopter l'IA militaire serait plus dangereux que déployer des systèmes imparfaitement alignés.

**OpenAI propose un « scorecard » pour mesurer le vrai ROI de l'IA**

La directrice financière d'OpenAI, **Sarah Friar**, propose un cadre pour sortir du marketing flou : mesurer le **travail utile** produit, le **coût par tâche réussie**, la fiabilité et le retour sur le calcul investi. Objectif : donner aux entreprises des métriques tangibles pour juger si l'IA leur rapporte vraiment, plutôt que des promesses.

**Meta pourrait louer son calcul IA excédentaire à Anthropic**

