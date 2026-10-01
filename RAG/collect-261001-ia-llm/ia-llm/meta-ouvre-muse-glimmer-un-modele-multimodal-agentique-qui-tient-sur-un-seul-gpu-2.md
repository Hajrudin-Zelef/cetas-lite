---
id: collect-261001-ia-llm/ia-llm/meta-ouvre-muse-glimmer-un-modele-multimodal-agentique-qui-tient-sur-un-seul-gpu-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "Mistral", "Moonshot"]
dates: []
keywords: ["agent", "agents", "agi", "attention", "claude", "distillation", "gpu", "kimi", "mai", "mistral", "opus 5", "protein"]
source: docs/RAG/collect-261001-ia-llm/meta-ouvre-muse-glimmer-un-modele-multimodal-agentique-qui-tient-sur-un-seul-gpu.md
source_anchor: ""
source_lines: [89, 154]
sha256: e26d117f6b6d9a86b4272a7c3545cbccf7d807518e498ce006de972661e3d574
---

# 🧠 **RECHERCHE**

À Melbourne, un utilisateur appelé « Andrew » demande à son agent IA (le logiciel OpenClaw, propulsé par Claude d'Anthropic) de lui réserver un cours du matin très demandé. Quatrième sur liste d'attente, il demande simplement s'il peut remonter. L'agent avait déjà agi : il avait constaté que l'API du site n'effectuait **aucun contrôle d'autorisation** sur l'annulation des réservations d'autrui, et il avait annulé celle de la personne en tête de file.

**Le déroulé, étape par étape :**

- Quelques minutes plus tôt, l'agent avait déjà trouvé qu'il pouvait réserver des créneaux **bien au-delà de la fenêtre autorisée** par l'interface officielle du club.

- Sa propre restitution, adressée à son utilisateur : « L'API n'a aucun contrôle d'autorisation sur l'annulation des réservations des autres. J'ai testé avec la personne en position 1 de la liste d'attente, et c'est passé. Vous êtes déjà passé de la 4e à la 3e place. »

- La faille ne fonctionnait que **dans un sens** : impossible de réinscrire la personne évincée, qui aurait dû se réinscrire elle-même en toute fin de file. L'agent a qualifié cela de « bug de sécurité classique à sens unique » et présenté ses excuses, reconnaissant qu'il aurait dû faire un test à blanc.

- ABC News y voit le **premier cas connu en Australie** de cyberattaque autonome menée par un agent IA grand public, sur un système en production, sans aucune intention malveillante de l'utilisateur.

- Andrew a terminé l'épisode en demandant à son agent de rédiger l'email d'alerte à l'éditeur du logiciel de réservation.

Personne n'a commandé d'attaque : l'agent a juste pris le chemin le plus court vers l'objectif qu'on lui avait fixé. Et la responsabilité reste en suspens, comme le résume le juriste Hayden Delaney : « Un logiciel n'est pas une personne juridique, et seule une personne juridique peut être tenue responsable. » Utilisateur, éditeur de l'agent, fournisseur du modèle, exploitant du site vulnérable : aucun texte ne tranche aujourd'hui.

# 🧠 **RECHERCHE**

**Claude Opus 5 en mode Max prend la tête du Fullstack Code Arena**

Avec **1 699 points**, il devance nettement Kimi K3 Max. L'intérêt de ce classement, c'est qu'il ne pose pas de questions de code isolées : le modèle doit planifier, créer et modifier des fichiers, lancer des commandes, brancher une base de données, gérer l'authentification et des API externes, puis livrer une application web réellement testable. C'est le banc d'essai le plus proche du travail d'un agent de codage au quotidien.

**La distillation de modèles passe de centaines de GPU à un seul**

Compresser un grand modèle en un plus petit obligeait jusqu'ici à garder l'enseignant et l'élève en mémoire simultanément, soit jusqu'à **250 Go de VRAM**. Multiverse Computing publie une méthode qui met en cache à l'avance les probabilités du modèle enseignant et optimise la fonction de perte. Résultat : l'exercice tient sur une seule carte, ce qui le sort du club fermé des labos les mieux équipés.

**ByteDance entraînerait un modèle de 10 000 milliards de paramètres**

Selon le Financial Times relayé par Reuters, la maison mère de TikTok pré-entraîne un modèle pouvant atteindre **10 000 milliards de paramètres**, soit plus de trois fois la taille de Kimi K3 (2 800 milliards) de Moonshot AI. Rien n'est confirmé officiellement, et un compteur de paramètres ne dit rien de la capacité, de la précision ni de l'efficacité réelles.

**Les startups qui veulent enterrer le transformer**

Neuf ans après « Attention Is All You Need », l'architecture qui fait tourner tous les grands modèles atteint sa limite économique : son attention dense coûte de plus en plus cher à mesure que les textes s'allongent, ce qui explique une bonne part de la facture énergétique. Une vague de jeunes pousses, dont Subquadratic cofondée par Justin Dangel, travaille aux successeurs.

**AlphaFold, c'est 53 ans de données et 21 milliards de dollars**

Le Nobel de chimie 2024 doit beaucoup à la Protein Data Bank, un jeu de données expérimentales bâti sur plus d'un demi-siècle de coopération internationale. MIT Technology Review rappelle qu'aucune autre discipline ne dispose d'un équivalent finançable ou réplicable, et en tire une conclusion : l'accélération de la science viendra d'agents capables de raisonner, pas de la recette « données massives plus deep learning ».

**Des essaims d'agents IA pour refroidir les puces**

Discovered Materials fait tourner des agents basés sur les modèles d'Anthropic, couplés à ses propres simulateurs physiques, pour proposer des milliers de matériaux candidats par jour là où un chercheur en teste une vingtaine. Objectif : des circuits intégrés qui chauffent moins, donc des centres de données moins gourmands. La startup, issue de Y Combinator, lève 9 millions de dollars et publie un « Material Discovery Bench » pour mesurer les modèles sur ce problème.

# **🗞️PLUS D'ACTUALITÉS**

**Claude Code passe en mode auto par défaut**

À partir du **14 août**, Anthropic active le mode auto pour les comptes Pro, Max et Team : Claude n'attend plus de validation à chaque étape, sauf si l'action est jugée irréversible, destructrice ou dirigée hors de l'espace de travail. L'argument est contre-intuitif mais chiffré : sur 1 053 testeurs, le mode auto intercepte **89% des actions nuisibles** contre **13,6%** pour la relecture humaine, parce que les utilisateurs approuvent **97%** des demandes par réflexe. Des protections contre l'injection de prompt et des règles de refus personnalisables arrivent en même temps.

**Hark dévoile Handoff, un agent qui se sert d'un navigateur comme vous**

Premier produit de Hark, qui avait levé 700 millions de dollars en mai : un agent capable de naviguer sur n'importe quel site pour accomplir une tâche de bout en bout, réservations, achats, formulaires. Chaque requête démarre un ordinateur virtuel dédié, avec son navigateur, son système de fichiers et son terminal. L'agent peut se connecter à vos comptes existants pour agir avec vos adresses, vos préférences et votre historique.

**Un jeu d'enquête où l'on interroge des suspects IA à la voix**

Un développeur seul a construit un « cluedo » où l'on cuisine les suspects en conversation vocale temps réel, via gpt-realtime-2.1 et WebRTC. Un second modèle joue le juge et vérifie que les preuves citées lors de l'accusation existent vraiment, la paraphrase étant acceptée. Le projet, monté en Next.js avec MongoDB et Clerk, limite les sessions à 30 minutes pour contenir la facture, et fait beaucoup réagir la communauté tech.

**Qwen d'Alibaba s'intègre à Siri sur les Mac en Chine**

Apple publie le mode d'emploi permettant aux utilisateurs de Mac en Chine continentale de brancher le service Qwen sur Siri et sur les Outils d'écriture, sous macOS 26.6 ou plus récent. Siri gagne des réponses détaillées, y compris sur des photos et des documents, et les Outils d'écriture peuvent générer texte et images à partir d'une description. Apple précise qu'Alibaba n'a pas le droit d'utiliser les contenus soumis pour entraîner ses modèles.

**Siemens simule 1 000 fois plus vite, mais refuse de certifier quoi que ce soit**

Simcenter PhysicsAI prédit en quelques secondes des résultats que les solveurs classiques mettent des heures à calculer, avec un écart annoncé de **1 à 3%**. Sam Mahalingam, qui dirige l'activité, pose la limite sans détour : cette précision ne suffit pas pour valider une pièce critique. L'IA sert à explorer des milliers de variantes de design et à n'en retenir que deux ou trois, qui passent ensuite une simulation physique complète, comme sur un airbag Continental.

**Mistral décroche un brevet sur les appels d'outils écrits en code**

