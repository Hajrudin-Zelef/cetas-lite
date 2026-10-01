---
id: collect-261001-ia-llm/ia-llm/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Apple", "DeepSeek", "Google", "Hugging Face", "Microsoft", "OpenAI", "Xiaomi", "xAI"]
dates: []
keywords: ["agent", "agents", "benchmarks", "chatgpt", "deepseek", "gemini", "gemini 3.8", "grok", "grok 4", "llama", "open source", "open weights"]
source: docs/RAG/collect-261001-ia-llm/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors.md
source_anchor: ""
source_lines: [81, 128]
sha256: 1aca86cf587d3fba913e5a666d506d60e75570ceb3ba0a14003dd75e82a44c1f
---

# 🧠 **RECHERCHE**

- **Date limite de dépôt : 21 décembre 2026**, confirmée par ConsumerAffairs. Le règlement ne vaut pas reconnaissance de faute

Aucun acheteur français n'est concerné, le dispositif est strictement américain. Ce qui compte ici, c'est le précédent : une promesse de fonctionnalité IA non tenue vient de recevoir un prix, chiffré, payé par appareil vendu. Tous les fabricants qui écoulent aujourd'hui du matériel sur la foi de capacités IA « à venir » ont désormais un ordre de grandeur pour évaluer le risque, et Siri reste, deux ans après l'annonce, une fonction incomplète.

# 🧠 **RECHERCHE**

### **MiMo-V2.6-Pro : Xiaomi publie le meilleur modèle open weights du monde**

Le constructeur d'électronique et de voitures électriques décroche **46 sur l'Intelligence Index d'Artificial Analysis**, à égalité avec Grok 4.7 sorti le même jour, devant Gemini 3.8 Flash et DeepSeek V4.1. Le modèle est sous **licence MIT**, téléchargeable gratuitement sur Hugging Face et exécutable sur votre propre matériel, avec **1 million de tokens de contexte** et des entrées texte, image, audio et vidéo. En API, comptez 0,435 $ par million de tokens en entrée et 0,87 $ en sortie, contre 0,14 $ et 0,28 $ pour la version Flash. Une déclinaison Pro-UltraSpeed promet une génération jusqu'à 20 fois plus rapide.

**Le modèle qui arbitre les décisions des agents IA se retourne avec une phrase**

Jev, le modèle de décision de TypeSafe adopté en quelques jours par Vercel, Cloudflare, LangChain et Langfuse, ne rédige rien : il tranche (autoriser une action, choisir un outil, classer une entrée) en 70 à 500 millisecondes pour 0,042 $ par million de tokens. Le souci, c'est qu'il considère les données qu'on lui transmet comme fiables : un test d'Octomind a fait chuter la probabilité de bloquer la commande destructrice `rm -rf ~/.ssh` de **0,76 à 0,48** en y glissant une fausse mention d'approbation préalable. LangChain exclut désormais les sorties d'outils de ce que voit le classifieur, pour qu'un contenu récupéré par l'agent ne puisse pas s'auto-autoriser, et Pydantic insiste : ce type de garde-fou complète des vérifications déterministes, il ne les remplace jamais.

**RetroChimera : l'IA de Microsoft qui écrit la recette de vos molécules**

Publié dans Nature et ouvert en open source, code et poids compris, RetroChimera propose automatiquement des itinéraires de synthèse pour fabriquer une molécule cible à partir de composants achetables. Il fusionne deux modèles de rétrosynthèse complémentaires et apprend à hiérarchiser leurs propositions, ce qui le rend meilleur que chacun pris isolément. En test à l'aveugle, des chimistes de niveau doctorat ont préféré ses prédictions à celles des modèles antérieurs, et même à des réactions publiées dans la littérature scientifique. Il retrouve des types de réactions rares et se transfère sans réentraînement sur des jeux de données propriétaires.

**Le « moment ChatGPT » des cerveaux de robots serait pour 2027**

Le fondateur de la startup chinoise Spirit AI, qui développe les modèles embarqués servant de cerveau aux robots humanoïdes, situe dès l'an prochain la percée équivalente à ce que ChatGPT a été pour le texte. Si la prédiction se vérifie, les humanoïdes passeraient des démonstrations scriptées à des machines capables de comprendre une scène réelle et d'agir dessus. La Chine met les moyens sur le sujet et vise ouvertement le leadership face aux acteurs américains.

**Des drones sous-marins pour couvrir 98% des océans d'ici 2028**

L'Ifremer a validé deux prototypes de flotteurs Deep-6000 capables de descendre à **6 000 mètres**, où la pression est 600 fois celle de la surface, ce qui fait de la France le troisième pays au monde à atteindre ces profondeurs après la Chine et les États-Unis, après cinq plongées réussies entre le 11 janvier et le 2 février. Trente flotteurs abyssaux rejoindront d'ici 2028 la flotte internationale Argo et ses 4 000 appareils. L'angle mort qu'ils comblent est de taille : **10% du réchauffement océanique se produit sous les 4 000 mètres** que la plupart des flotteurs actuels n'atteignent pas. Chaque engin pèse 40 kg, tient sept ans sans recharge et suit un cycle de dix jours, neuf de dérive en profondeur puis trois heures de profilage.

**Amazon veut concevoir des anticorps par ordinateur**

Amazon Bio Discovery publie trois travaux sur la conception d'anticorps thérapeutiques, dont **MochiBind**, bâti sur le modèle de langage protéique ESM-2, qui compare la force de liaison de deux anticorps au lieu de prédire une valeur absolue difficilement comparable d'une expérience à l'autre. Un troisième travail confie à un agent IA le choix des sites de liaison et a produit des anticorps validés en laboratoire contre une cible cancéreuse inédite. L'objectif : raboter les six à douze mois habituellement nécessaires pour passer d'une cible biologique à un candidat médicament. Amazon pointe au passage la faiblesse des benchmarks actuels, fiables sur les cibles connues, beaucoup moins sur les nouvelles.

**Et si on homologuait l'IA médicale comme un médicament ?**

Des chercheurs de l'université de Bristol proposent un cadre baptisé « Learning Ensemble » pour évaluer les systèmes d'IA utilisés en médecine, calqué sur les procédures d'approbation des médicaments. Trois axes de vérification : les limites du système, l'équité entre groupes de patients, l'adéquation au contexte clinique réel. L'idée est d'attraper les modèles qui brillent sur le papier et se trompent gravement au lit du patient. Argument central des auteurs : la médecine sait depuis longtemps encadrer des mécanismes qu'elle ne comprend pas entièrement.

**Grok Voice Transcribe 2.0 : deux fois plus précis, au même prix**

xAI double la précision de son modèle de transcription sans toucher au tarif : **0,10 $ par heure d'audio** en traitement par lots, 0,20 $ en streaming. Il est entraîné sur des enregistrements réels, bruités et multilingues, et vise l'audio difficile : sur les commandes vocales courtes, le taux d'erreur par mot tombe de **20,6% à 6,8%**. Il gère des dizaines de langues avec détection et bascule automatiques en cours d'enregistrement, et se classe premier en précision parmi 32 modèles en streaming sur le classement public d'Artificial Analysis. Les intégrations existantes de l'API récupèrent le gain sans une ligne de code à modifier.

**Compresser un LLM comme un physicien**

Multiverse Computing reformule la suppression de blocs entiers d'un transformer, la méthode la plus brutale pour accélérer un modèle, comme un problème de verre de spins d'Ising. L'intuition de départ : l'effet du retrait d'un bloc dépend des autres blocs retirés, donc les noter isolément, ce que font les méthodes classiques, est une erreur de raisonnement. L'énergie du système sert ensuite de proxy rapide pour prédire la performance du modèle compressé sans repasser les benchmarks. Résultat annoncé : **près de 23 points de MMLU** gagnés face à la meilleure méthode concurrente sur un Llama-3.3-70B-Instruct compressé de moitié.

**Valorisée 1,4 milliard de dollars avant d'avoir publié le moindre modèle**

Naive AI, fondée en février à Pékin par Jifeng Dai, professeur à l'université Tsinghua, aurait levé environ 400 millions de dollars en trois tours avec **moins de 100 salariés** et aucun modèle publié à ce jour. Sa stratégie tranche avec celle de ses concurrents : ne pas pré-entraîner depuis zéro, mais repartir d'un modèle chinois open-weight existant pour concentrer ses moyens sur le midtraining, le post-training, l'apprentissage par renforcement et l'architecture. Un modèle en open weights pourrait sortir dès ce mois-ci.

**DeepSeek-V4.1-Flash comprime sa mémoire à 890 octets par token**

