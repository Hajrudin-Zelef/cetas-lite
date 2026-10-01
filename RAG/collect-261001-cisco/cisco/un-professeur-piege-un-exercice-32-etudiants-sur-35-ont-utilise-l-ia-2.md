---
id: collect-261001-cisco/cisco/un-professeur-piege-un-exercice-32-etudiants-sur-35-ont-utilise-l-ia-2
title: "🧠 **RECHERCHE**"
domain: cisco
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "United States"]
dates: []
keywords: ["agent", "agents", "attention", "benchmark", "cyber", "diffusion", "distillation", "gpu", "kimi", "mai", "mistral", "nvidia"]
source: docs/RAG/collect-261001-cisco/un-professeur-piege-un-exercice-32-etudiants-sur-35-ont-utilise-l-ia.md
source_anchor: ""
source_lines: [66, 118]
sha256: 05433a024c39eb2602e0a2520237fad451bbae67d54f4a5a6d22edb7cfdaae20
---

# 🧠 **RECHERCHE**

**FeyNoBg : un modèle de détourage d'arrière-plan et sa bibliothèque d'entraînement, en open source**

La startup Feyn publie FeyNoBg, un modèle de suppression automatique d'arrière-plan, et **NoBg**, la bibliothèque Python open source qui sert à l'entraîner et à le faire tourner. La démarche est inhabituelle : plutôt que d'empiler des paramètres, l'équipe a disséqué l'architecture BiRefNet pour identifier l'étage qui concentre l'essentiel de l'information utile (repérage du sujet et tracé fin des contours : cheveux, flou de mouvement, camouflage), puis a agrandi cet étage précis. Entraîné sur 26 100 exemples issus de 10 jeux de données. Démo sur Hugging Face, code sur GitHub, 101 points sur Hacker News.

**NVIDIA Cosmos-H-Dreams : un simulateur génératif temps réel pour la robotique chirurgicale**

Les scènes chirurgicales sont un cauchemar à simuler : tissus déformables, surfaces réfléchissantes, sutures, aiguilles, fumée, occlusions. NVIDIA prend le problème par l'autre bout avec Cosmos-H-Dreams, un modèle qui **génère** les conséquences visuelles des gestes du robot au lieu de les calculer. Il reçoit une image de départ et le flux cinématique du robot, puis produit la suite de la scène en continu, en boucle fermée, **sur un seul GPU RTX PRO 6000**. Le modèle a été distillé depuis Cosmos-H-Surgical-Simulator via la librairie FlashDreams, et son intégration a été démontrée sur la plateforme chirurgicale Versius de CMR Surgical.

**Un candidat médicament en 9 mois au lieu de 4 ans et demi**

Insilico Medicine, cotée à Hong Kong, affirme ramener à **environ un an** (13 mois en moyenne, **9 mois** pour son programme le plus rapide) le temps nécessaire pour identifier un candidat médicament, contre environ 4 ans et demi par les méthodes conventionnelles. L'IA générative propose les cibles biologiques et dessine les molécules, ce qui réduit à 60-200 le nombre de composés à synthétiser et tester. Attention à ce que ça ne couvre pas : essais cliniques, fabrication et validation réglementaire restent des étapes séparées, non accélérées. Le CEO Alex Zhavoronkov attribue une partie du gain à l'écosystème chinois, la R&D IA étant menée à Montréal et Abou Dabi et les tests biologiques à Shanghai.

**Le vrai goulot d'étranglement de la découverte de médicaments par IA, ce sont les données**

Depuis les années 1950, le coût de développement d'un nouveau médicament double environ tous les neuf ans, un phénomène baptisé **loi d'Eroom**. Aujourd'hui : 10 à 15 ans, entre 1 et 2,5 milliards de dollars, et plus de 90% d'échecs. L'IA fait basculer l'industrie du criblage physique de molécules vers la conception prédictive avant toute manipulation en laboratoire. Mais l'article insiste sur le frein réel : la qualité et l'authenticité des données de paillasse, et leur intégration effective dans les systèmes de R&D existants.

**Dario Amodei : "nous n'avons jamais plaidé pour une interdiction des modèles à poids ouverts"**

Le patron d'Anthropic a publié lundi un billet de blog pour couper court aux accusations montant dans l'industrie. Sa position de rechange tient en trois points : contrôler l'accès aux puces avancées, lutter contre la distillation industrielle (il cite une attaque qu'il attribue à Alibaba/Qwen), et imposer des tests de sécurité à tout modèle suffisamment puissant, ouvert ou fermé. Une clarification qui arrive après la publication d'une lettre collective anti-restrictions qu'Anthropic n'a pas signée.

**Anthropic, seul grand labo IA à ne pas avoir signé la lettre pro open weight**

Nvidia, Microsoft, Meta, Palantir, Mistral, Google et finalement OpenAI, après une hésitation, ont signé vendredi une lettre ouverte demandant à Washington de ne pas restreindre les modèles à poids ouverts. Jensen Huang y a consacré **son tout premier post sur X** : *"Les modèles ouverts renforcent la sécurité et la cybersécurité, accélèrent l'innovation et la diffusion, et permettent la souveraineté."* Le texte vise en creux les mesures qu'envisagerait l'administration Trump contre des acteurs comme Moonshot AI, l'éditeur de Kimi K3. Anthropic est le seul grand nom absent.

**Microsoft lance MAI-Cyber-1-Flash, mais appelle toujours OpenAI pour les cas difficiles**

Microsoft dévoile un modèle compact spécialisé en cybersécurité qui atteint **96% sur le benchmark CyberGym** une fois intégré à son système multi-agents MDASH. L'intérêt est économique autant que technique : un routage intelligent n'envoie que les cas les plus durs vers GPT-5.4, ce qui ferait chuter les coûts de **50%** par rapport à un usage direct de modèles frontière. La dépendance à OpenAI pour le raisonnement complexe reste, elle, entière.

**METR chiffre en dollars le moment où un agent IA coûte plus cher qu'un humain**

Le laboratoire METR propose l'**"expenditure horizon"**, une métrique qui compare directement le coût en dollars d'un agent IA et celui d'un humain pour résoudre le même problème. Premiers résultats sur le speedrun NanoGPT : peu flatteurs pour les agents actuels. METR reconnaît des angles morts dans la mesure, et note que la génération de modèles qui arrive pourrait rebattre les cartes.

**Et si la prochaine étape n'était pas des modèles plus gros, mais des agents qui se coordonnent**

Cisco, via sa division Outshift, défend une architecture à deux couches, un *"Internet of Agents"* pour la connectivité et un *"Internet of Cognition"* pour le sens partagé, afin que des agents spécialisés collaborent réellement sur un objectif commun au lieu de simplement s'échanger des données. L'argument chiffré fait mal : une étude citée mesure un taux d'échec de **41 à 87%** sur sept systèmes multi-agents open source évalués. La thèse : passer du scaling vertical au scaling horizontal. À noter, il s'agit d'un contenu sponsorisé.

**Des lasers pour extraire du combustible nucléaire de vieux déchets d'enrichissement**

Près de Paducah, dans le Kentucky, des milliers de cylindres stockent les résidus d'une usine d'enrichissement fermée. Global Laser Enrichment veut les retraiter avec une technique qui ne sépare plus les isotopes par la masse, comme les centrifugeuses, mais les cible sélectivement grâce aux **signatures vibratoires propres à l'U-235**. Objectif : produire du combustible à la concentration d'un minerai naturel, y compris pour les réacteurs avancés. Le nucléaire représente aujourd'hui environ **9% de l'électricité mondiale**.

**Quand une IA devrait-elle vous contredire ?**

Cette étude soutient que la complaisance des modèles est un phénomène plus fin que la simple flagornerie. Trois facteurs déterminent si un modèle change de jugement moral : la distance entre l'avis exprimé et sa position initiale, l'identité de celui qui exprime le désaccord, et le fait qu'un groupe le soutienne ou non. L'enjeu pratique : savoir distinguer une reconsidération légitime d'une soumission à la pression sociale, dans les conversations où ça compte vraiment.


# **🗞️PLUS D'ACTUALITÉS**

**Alibaba publie open-code-review, un outil gratuit de revue de code par IA**

Alibaba met en open source l'outil de revue de code qu'il utilise en interne, sur des bases de code à son échelle. L'architecture est hybride : des pipelines déterministes pour ce qui se vérifie mécaniquement, un agent LLM pour le reste, avec des commentaires précis ligne par ligne. Le ruleset intégré est affiné pour détecter les NullPointerException, les problèmes de thread-safety, les failles XSS et les injections SQL. Compatible avec les API OpenAI et Anthropic, donc branchables sur le modèle de votre choix. Déjà plus de 15 000 étoiles sur GitHub.

**Hugging Face a un problème massif de deepfakes nus non consentis**

