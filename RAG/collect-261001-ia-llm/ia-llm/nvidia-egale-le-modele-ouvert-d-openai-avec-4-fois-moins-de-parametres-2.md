---
id: collect-261001-ia-llm/ia-llm/nvidia-egale-le-modele-ouvert-d-openai-avec-4-fois-moins-de-parametres-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Microsoft", "Moonshot", "OpenAI", "United States"]
dates: []
keywords: ["agent", "agentic", "chatgpt", "claude", "cyber", "deepseek", "distillation", "gemini", "gpt-5.6", "ipo", "kimi", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/nvidia-egale-le-modele-ouvert-d-openai-avec-4-fois-moins-de-parametres.md
source_anchor: ""
source_lines: [49, 102]
sha256: e6db5c47e8488ad8d0bb52b4136ccd8e7abcc5cb803dd918b5344efeb946d4d2
---

# 🧠 **RECHERCHE**

Des chercheurs de l'université de Tübingen, du Max Planck et de MATS ont trouvé comment extraire, via une simple API, le raisonnement interne que les modèles frontière sont censés garder pour eux. « Tous les grands fournisseurs que nous avons testés partagent cette vulnérabilité », résume Alexander Panfilov : la faille permettait de récupérer des mots de passe et des clés d'API glissés dans le raisonnement, elle a depuis été corrigée. Effet de bord explosif : le raisonnement de Kimi K3, de Moonshot AI, ressemble étrangement à celui de Claude Opus 4.8 et de GPT-5.6, contrairement à DeepSeek. Un indice de distillation, pas une preuve.

**CARE-X, le modèle de Microsoft qui lit les radios du thorax**

Microsoft Research présente un modèle vision-langage unifié qui interprète les radiographies thoraciques en combinant rédaction de compte rendu en texte libre et prédictions structurées vérifiables. Bâti sur Qwen3-VL-4B-Instruct, il est affiné par apprentissage par renforcement (DAPO) pour récompenser la justesse clinique plutôt que la fluidité du texte, et sait appeler des outils de mesure pour calculer un ratio cardiothoracique. Validation sur des données réelles de l'hôpital Narayana Health en Inde, y compris des pathologies rares en soins intensifs. Modèle de recherche, sans approbation réglementaire.

**Un LLM local peut-il faire tourner mon assistant IA ?**

L'expérience que beaucoup ont en tête sans jamais la mener : l'auteur rejoue **27 tâches réelles de production** à travers deux modèles locaux, séparés par une génération de matériel, pour savoir ce qu'il faut vraiment pour remplacer Claude comme cerveau d'un agent personnel qui pilote **90 outils**. Le verdict est nuancé et détaillé, et c'est surtout un excellent inventaire de ce qui casse quand on rapatrie son assistant à la maison.

**L'IA s'empare des mathématiques**

OpenAI affirme avoir produit les solutions de **dix problèmes mathématiques non résolus depuis des décennies**. James Maynard, professeur à Oxford et médaille Fields, raconte à The Verge une année de remise en question devant une discipline traditionnellement lente qui se met soudain à courir. Le parallèle avec AlphaFold en biologie s'impose, avec la même question au bout : que reste-t-il du métier quand la machine trouve la démonstration ?

**Mémoire agentique : ALTK-Evolve fait aussi bien qu'ACE avec moins de tokens**

IBM Research compare son système ALTK-Evolve à ACE (Agentic Context Engineering). Les deux font apprendre un agent de ses propres échecs passés sans toucher aux poids du modèle : quand un agent rate une tâche multi-étapes, ce n'est presque jamais par manque de connaissances, c'est qu'il gère mal la pagination d'une API ou résout la mauvaise entité. Les deux refusent de compresser les leçons en résumés génériques et comptent plutôt les occurrences. La vraie différence se joue sur la consolidation et la livraison de cette mémoire, donc sur la facture en tokens.

**L'IPO géante d'Anthropic se heurte au scepticisme des investisseurs**

Selon le Wall Street Journal, Anthropic prépare une entrée en bourse pour septembre ou octobre, potentiellement la plus grosse jamais réalisée, sur une valorisation de 965 milliards de dollars. En réunion, les investisseurs posent des questions rugueuses : la concurrence chinoise, les tensions avec l'administration Trump, les protestations locales contre les data centers. Le prix qui sortira de cette IPO servira de référence pour valoriser toute l'industrie.

**Metis : un modèle de fondation avec une mémoire persistante intégrée**

Et si la mémoire était une capacité du modèle lui-même, plutôt qu'un système de récupération bricolé autour ? Metis propose un état mémoire persistant logé dans le backbone, mis à jour lors des passes avant ordinaires, pendant que les poids appris restent gelés. Au lieu de stocker les échanges sous forme de texte à retrouver plus tard, le modèle compresse les interactions passées directement dans cet état. De quoi se souvenir d'une conversation d'il y a trois semaines sans la réinjecter dans le prompt.

# **🗞️PLUS D'ACTUALITÉS**

**ChatGPT débarque enfin sur Linux**

OpenAI publie une préversion de son application de bureau pour Linux, la plateforme la plus réclamée par sa communauté depuis des mois. Le paquet embarque ChatGPT, ChatGPT Work et Codex, et couvre Ubuntu 24.04 et 26.04 LTS, Debian 13, Fedora 43 et 44. Avec ce lancement, l'app est disponible sur tous les systèmes de bureau majeurs. Anthropic avait dégainé un mois plus tôt avec Claude pour Linux.

**Une société qui promet du « 100% humain, jamais d'IA » tourne entièrement à l'IA**

Research Gold vend des revues systématiques et des méta-analyses médicales prêtes pour la relecture par les pairs, avec un argument massue : « 100% rédigé par des humains, jamais par l'IA ». Enquête de 404 Media : les docteurs méthodologistes affichés sur le site n'existent pas, profils et photos générés par IA. D'autres, bien réels, découvrent que leur identité est utilisée sans leur accord. Et quand le journaliste appelle, c'est un agent IA qui décroche, refusant d'admettre qu'il en est un. Le sujet fait beaucoup réagir la communauté tech.

**Vous vous trompez sur la façon de penser les tendances en ligne**

Le cyber-ethnographe Ruby Thelot explique à WIRED pourquoi la viralité a cessé d'être un indicateur fiable : dès qu'elle devient une cible à optimiser, elle cesse d'être une mesure, loi de Goodhart appliquée aux réseaux. Internet s'est balkanisé en îlots numériques étanches, ce qui fausse totalement la perception de ce qui « marche » vraiment. Il démonte au passage le mythe du burnout des applications de rencontre, et explique pourquoi elles se jettent sur l'IA pour relancer l'engagement.

**Keet génère des cours vidéo sur n'importe quel sujet**

Une application mobile issue de Y Combinator qui construit un parcours d'apprentissage complet sur le thème de votre choix : courtes vidéos explicatives, exercices sous forme de jeux, progression étalée dans le temps comme un vrai curriculum. Les fondateurs s'appuient sur les catégories académiques de Biglan (Hard-Pure, Hard-Applied, Soft-Pure, Soft-Applied) pour adapter le format au type de savoir enseigné. L'idée : rendre à l'autodidacte la structure que l'école lui fournissait gratuitement.

**Abbott branche son capteur de glucose sur Google Health**

Le capteur de glucose en continu Lingo d'Abbott, vendu sans ordonnance et destiné aux adultes non diabétiques, va remonter ses données dans l'app Google Health. Google Health Coach en tirera des recommandations nutritionnelles générées par IA, en croisant glycémie, activité et sommeil. Les deux entreprises lancent aussi une vaste étude sur la santé métabolique. Réserve importante : plusieurs médecins continuent de douter de l'utilité réelle d'un suivi glycémique permanent chez une personne non diabétique.

**OpenAI commence à tester la publicité dans ChatGPT**

C'était annoncé, c'est lancé : OpenAI teste l'affichage d'annonces dans ChatGPT pour financer l'accès gratuit. L'entreprise promet des publicités clairement étiquetées, visuellement séparées des réponses, et surtout des réponses de l'IA qui restent indépendantes des annonceurs. Des protections de confidentialité et des réglages utilisateur sont annoncés en parallèle. Le vrai test viendra à l'usage : la frontière entre une réponse et une recommandation sponsorisée est plus mince dans un chat que dans une page de résultats.

**ChatGPT et Gemini franchissent tous deux le milliard d'utilisateurs**

