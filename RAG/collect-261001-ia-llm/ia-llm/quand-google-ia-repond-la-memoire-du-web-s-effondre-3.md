---
id: collect-261001-ia-llm/ia-llm/quand-google-ia-repond-la-memoire-du-web-s-effondre-3
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Apple", "Google", "OpenAI", "Z.ai"]
dates: []
keywords: ["agent", "agents", "aws", "claude", "compute", "cyber", "diffusion", "glm", "gpt-5.6", "gpu", "lora", "mai"]
source: docs/RAG/collect-261001-ia-llm/quand-google-ia-repond-la-memoire-du-web-s-effondre.md
source_anchor: ""
source_lines: [133, 178]
sha256: 25c75a910e8ccc3f73a9ebca7ce49567444ddbb8a221d96e0e4e844f6fa69713
---

# 🧠 **RECHERCHE**

Lors d'une rencontre du programme AI2050 financé par Eric et Wendy Schmidt, des professeurs racontent un déplacement du centre de gravité de la recherche vers Anthropic, OpenAI et Google, qui gardent secrets les détails d'entraînement de leurs modèles. Sans budget pour acheter des GPU ni pour interroger massivement ces systèmes, des chercheurs de UC Berkeley et Johns Hopkins réorientent leurs travaux vers les questions que les entreprises n'ont aucun intérêt à explorer elles-mêmes.

**AWS lance une compétition pour co-concevoir modèles et puces**

Amazon invite laboratoires académiques et industriels à entraîner un modèle de langage d'environ **50 millions de paramètres** depuis zéro sur ses puces Trainium, avec liberté totale sur l'architecture, l'optimiseur et les noyaux de calcul. L'idée de fond est intéressante : les architectures actuelles ont été façonnées par les contraintes des GPU, et Trainium propose un profil différent (plus de SRAM sur puce, contrôle logiciel explicite des mouvements de données). À quoi ressemblerait un modèle conçu pour ce matériel-là ?

**Macaron-V1, des agents qui continuent d'apprendre après leur déploiement**

Cette famille de modèles open source repose sur une architecture **Mixture-of-LoRA** : un modèle de base gelé, plusieurs adaptateurs spécialisés (conversation, agent, code, interface générative), et un seul sélectionné à chaque tour de conversation. Le modèle phare, Macaron-V1-Venti, associe une base **GLM-5.2 de 744 milliards de paramètres** à ces adaptateurs, tandis que la version Tall (**50B**, base Qwen3.6) vise le déploiement local. Le tout est piloté par une boucle d'auto-amélioration où l'expérience d'une version sert à construire la suivante.

**Des modèles vidéo qui apprennent la physique au lieu d'imiter les pixels**

Les modèles de diffusion vidéo produisent des images plausibles sans jamais modéliser comment les pixels se transforment dans le temps, d'où les objets qui traversent les murs. La méthode **Latent Dynamics Reasoning** traite la transition entre images comme une intégration cinématique explicite, et généralise à des situations jamais vues : entraînée sur des balles rouges, elle prédit correctement la trajectoire d'un carré bleu. Bonus non négligeable, elle est **26 fois plus légère et 143 fois plus rapide** que les modèles de référence.

# **🗞️PLUS D'ACTUALITÉS**

**Une chaîne de pharmacies débranche son assistant vocal IA après des centaines de plaintes**

Kinney Drugs avait lancé en mai « Burt », un assistant vocal IA baptisé du prénom du fondateur de la chaîne, chargé de gérer les appels des patients sur leurs ordonnances et renouvellements. Les clients ont rapidement signalé des conversations incohérentes, des **dosages erronés communiqués au téléphone** et des notifications de renouvellement jamais envoyées. L'entreprise fait marche arrière : retour au bon vieux serveur vocal à touches pour les appels entrants, l'IA n'étant conservée que pour des SMS sortants avec consentement explicite. La déclaration du dirigeant mérite d'être citée : réussir la confidentialité et la sécurité ne signifie pas avoir réussi l'expérience, et là, c'était raté.

**Des startups robotiques américaines ramènent leurs pièces chinoises dans leurs valises**

The Information raconte une scène devenue routinière : des ingénieurs de startups américaines de robots humanoïdes prennent l'avion pour Shenzhen, écument le marché électronique de Huaqiangbei et repartent avec les composants dans leurs bagages à main. Les pièces concernées ne sont pas anodines, ce sont notamment les actionneurs, le cœur mécanique d'un humanoïde. La raison est prosaïque : le fret prend des semaines que des petites équipes en course contre la montre n'ont pas. L'anecdote résume assez bien le rapport de force matériel entre les deux pays.

**Apple préparerait un certificat d'authenticité pour vos photos**

9to5Mac a repéré dans une bêta d'iOS 27 les traces d'une fonctionnalité nommée **Reference Image**, capable de prouver qu'une photo a bien été capturée par un iPhone. Le principe : activer un mode Reference dans l'appareil photo, puis envoyer les métadonnées du capteur à Private Cloud Compute pour vérification. C'est l'approche inverse des filigranes type SynthID de Google, qui marquent les images générées : ici, on certifie le réel plutôt que de tatouer le synthétique. Rien ne garantit encore que la fonction survivra jusqu'à la sortie publique cet automne.

**Google ajoute des agents IA à Google Ads et Analytics**

Google déploie une série de fonctions IA dans ses outils marketing, disponibles dès maintenant dans les comptes. Au menu : des résumés automatiques de performance directement sur la page d'accueil d'Analytics, avec des alertes personnalisables par notification, la **création de rapports visuels à partir d'une simple phrase** en langage naturel, et un comparatif de vos performances publicitaires face à des entreprises similaires. Le tout s'appuie sur l'agent Ask Advisor. Pour quiconque passe ses matinées à fabriquer des tableaux de bord à la main, c'est le genre de mise à jour qui se teste en dix minutes.

**Claude va marquer ses contenus pour se conformer à l'AI Act européen**

Anthropic a signé le Code de bonnes pratiques de l'article 50(2) de l'AI Act sur la transparence des contenus générés par IA. Concrètement, **tous les nouveaux modèles Claude lancés dans l'UE après le 2 août 2026** intégreront dès leur sortie un marquage lisible par machine : filigranes dans le texte, métadonnées de provenance signées dans les fichiers. Le dispositif s'appliquera partout où Claude est utilisé, de l'API à Claude Code en passant par Claude Cowork. Anthropic promet une documentation technique pour permettre aux tiers de détecter ces marques, ce qui sera le vrai test de l'utilité du système.

**La FCC veut interdire des drones qu'elle avait elle-même approuvés**

Le régulateur américain des télécoms propose d'étendre son interdiction aux drones équipés de LiDAR, d'imagerie thermique ou de capacités de dispersion d'aérosols, y compris des modèles qu'il avait déjà validés. Plusieurs DJI grand public sont visés (Air 3S, Avata 360, Mini 5 Pro), ainsi que des modèles professionnels utilisés en agriculture et en recherche-sauvetage. Les appareils déjà vendus ne seraient pas confisqués, mais l'accès aux pièces détachées et aux mises à jour logicielles pourrait se tarir. DJI dénonce un revirement complet et la consultation publique reste ouverte jusqu'au 2 septembre.

**OpenAI lance GPT-5.6-Cyber, un modèle dédié à la cybersécurité**

OpenAI étend son programme Daybreak avec un modèle spécialisé, **GPT-5.6-Cyber**, accessible via la plateforme Daybreak Red. Il est destiné à la recherche de vulnérabilités autorisée, à la validation d'exploits et aux tests de sécurité, c'est-à-dire à des usages offensifs encadrés au service de la défense. L'entreprise justifie le lancement par le rétrécissement de la fenêtre entre le moment où une faille devient exploitable par une IA et celui où les défenseurs peuvent la corriger. C'est aussi le signe d'une stratégie de modèles verticaux, taillés par métier plutôt que généralistes.

**OpenAI écrit au gouverneur du Texas sur ses data centers**

