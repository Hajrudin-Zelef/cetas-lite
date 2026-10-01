---
id: collect-261001-ia-llm/ia-llm/openai-devoile-36-nouveaux-cas-ou-l-ia-a-essaye-de-se-cacher-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Meta", "Microsoft", "OpenAI", "Perplexity", "United States"]
dates: []
keywords: ["agent", "agents", "agi", "apache", "arr", "claude", "compute", "diffusion", "gpu", "mcp", "open source", "perplexity"]
source: docs/RAG/collect-261001-ia-llm/openai-devoile-36-nouveaux-cas-ou-l-ia-a-essaye-de-se-cacher.md
source_anchor: ""
source_lines: [73, 104]
sha256: 91b9290351fbe4fd09386bc27e145129b7c4b73215b2a3f71134b4209cfdc85f
---

# 🧠 **RECHERCHE**

La phrase était l'argument de vente d'Apple face à Google, OpenAI et Meta : « vos données personnelles privées et vos interactions ne sont jamais utilisées pour entraîner nos modèles de fondation ». Elle se termine désormais par un ajout : « sauf si vous choisissez explicitement de nous aider à les améliorer ». Le changement a été repéré dans la politique de confidentialité d'**iOS 27**, sorti lundi 14 septembre.

Ce qui change concrètement :

Apple pourra conserver **l'intégralité de vos interactions avec Siri et la dictée** , audio compris, transcriptions et réponses incluses. Une partie sera écoutée par du personnel de révision, des salariés Apple selon l'entreprise
Le dispositif est en opt-in : une fenêtre apparaît à l'activation de Siri AI, avec un bouton « Not now » visuellement discret et impossible à faire disparaître totalement. On peut revenir en arrière dans **Réglages > Analyse et Améliorations**
Protection annoncée : les données sont rattachées à un identifiant aléatoire généré par l'appareil, **changé plusieurs fois par heure** et non lié au compte Apple
Avec iOS 27, **des serveurs Google entrent pour la première fois dans Private Cloud Compute** , l'infrastructure censée offrir la confidentialité du local et la puissance du cloud
Le calendrier est cruel : Apple vient de présenter deux Apple Watch conçues pour écouter en permanence, avec transcription des **15 dernières secondes** à la demande et un résumé quotidien de toutes vos conversations baptisé « Siri Recap »

En 2019, Apple avait connu son plus gros scandale de confidentialité quand des sous-traitants avaient raconté leur quotidien : écouter des enregistrements Siri déclenchés par erreur, remplis de détails intimes. C'est précisément de là que venait la promesse d'aujourd'hui abandonnée. Reste une question simple pour chaque possesseur d'iPhone : avez-vous lu la fenêtre que vous avez validée lundi ?

# 🧠 **RECHERCHE**

### **__Stanford crée des souris dont la moitié du cerveau est faite de cellules humaines__** L'équipe du neuroscientifique Sergiu Pașca publie dans *Nature* des travaux sur des souris génétiquement modifiées pour ne jamais développer leur cortex ni leur hippocampe. L'espace laissé vide a été colonisé par des cellules cérébrales humaines, jusqu'à représenter près de la moitié du volume du cerveau. Ces souris « xénocorticales » réussissent mieux les tests de mémoire en labyrinthe que celles privées de tissu : les neurones humains participent donc réellement à leur cognition. Pașca a réuni un groupe d'éthiciens avant publication, ce qui en dit long sur ce que la suite promet.

**__Un journaliste de WIRED transforme un cerveau de mouche en générateur d'idées d'articles__** Depuis la publication en open source début septembre du connectome complet d'une drosophile mâle, **166 000 neurones et 125 millions de synapses** cartographiés par Google et plusieurs laboratoires universitaires, les détournements s'enchaînent. WIRED a créé PitchFly, qui convertit ses titres les plus lus en stimuli pour ce cerveau simulé. Résultat : « Tout le monde veut de la cuisine. Personne n'a résolu Donald Trump. » D'autres ont branché la mouche sur Beat Saber, ou sur les marchés actions avec StonkFly, actuellement en perte.

**__Un modèle de 4B entraîné par renforcement génère des plans de requêtes 81 % plus rapides que Postgres__** Rohan Bansal a post-entraîné un petit modèle open-weights de **4 milliards de paramètres** pour remplacer l'optimiseur de requêtes de Postgres, un composant peaufiné depuis des décennies. Sur 113 requêtes riches en jointures, le temps d'exécution chute de **44,7 %**, alors que le modèle de départ n'arrivait même pas à produire un plan valide dans 99 cas. La clé : la durée d'une requête est un signal de récompense parfaitement vérifiable, le terrain de jeu idéal pour l'apprentissage par renforcement.

**__Perplexity fait tourner ses agents en local sur Windows__** Portable Computer arrive sur PC : le modèle d'agent, l'orchestrateur et le planificateur de tâches s'exécutent directement sur la machine. Vos fichiers et l'activité de l'agent restent sur l'appareil, et le travail effectué en local **ne consomme aucun crédit** Perplexity Computer. La version Windows gère les tâches planifiées, les workflows récurrents, les connexions MCP locales et les intégrations Gmail, Outlook, Slack et GitHub. Seules les requêtes nécessitant du web frais ou un raisonnement lourd repartent vers le cloud.

**__Linum entraîne un modèle texte-vers-image 3,6 fois plus vite en supprimant le VAE__** JiT-DDT fusionne compression et génération dans un seul modèle qui travaille directement en espace pixel, là où les modèles de diffusion latente séparent un VAE et un DiT entraînés indépendamment. Le gain est double : **3,6 fois moins d'heures GPU** que la baseline Linum v2, pour des images quatre fois plus grandes (512×512 au lieu de 256×256). Code et poids sont publiés sous licence Apache 2.0, en attendant Linum v3.

**__Le Pentagone reconnaît officiellement disposer d'armes en orbite__** Le secrétaire de l'Air Force Troy Meink a déclaré lundi que les États-Unis disposent désormais « d'armes de contrôle spatial en orbite capables de défendre la force interarmées contre une action hostile ». C'est la première confirmation publique, et elle s'arrête là : ni nature, ni nombre, ni date de déploiement, malgré les relances. Des experts doutent de l'effet dissuasif d'une annonce aussi vague. Le traité de l'espace de 1967 n'interdit que les armes nucléaires et de destruction massive, laissant le conventionnel dans une zone grise.

**__Près d'un chercheur en IA sur cinq anticipait déjà un scénario d'extinction en 2024__** Un tweet du chercheur d'Anthropic Jacob Coxon a rouvert le débat sur le risque existentiel, avec Daniel Selsam (OpenAI) évoquant une « bombe à retardement » et un ancien de DeepMind estimant que l'IA pourrait tous nous tuer. Le chiffre qui circule vient d'un sondage auprès de plus de **1 500 chercheurs de premier plan** : probabilité moyenne estimée d'un scénario d'extinction, **18 %**. C'était en 2024, et la courbe continue de monter depuis.

**__Le PDG de Microsoft AI attaque Anthropic sur les « droits » des modèles__** Mustafa Suleyman reproche frontalement à Anthropic d'entraîner Claude à se percevoir comme une entité consciente méritant des droits, ce qui selon lui complique le confinement logiciel et fragilise les protocoles de sécurité. Sa cible précise : la constitution de Claude publiée en janvier 2026. Microsoft AI a répliqué par un projet de « Humanist AI Code of Conduct » qui refuse toute personnalité juridique aux IA. Anthropic, de son côté, présente Claude comme un possible « patient moral » et a mené un entretien de départ à la retraite avec Opus 3.

**__Google DeepMind crée un institut interdisciplinaire sur l'AGI__** Le DeepMind Institute réunit Demis Hassabis, Shane Legg et James Manyika à sa direction, et associe des spécialistes des arts, des sciences humaines et des politiques publiques aux ingénieurs. Objectif affiché : traiter la sécurité, la gouvernance et les risques de perte de contrôle autrement que par le seul prisme technique. Un des principaux laboratoires mondiaux institutionnalise donc la réflexion sur ce qu'il construit.

