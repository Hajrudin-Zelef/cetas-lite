---
id: collect-261001-ia-llm/ia-llm/nvidia-egale-le-modele-ouvert-d-openai-avec-4-fois-moins-de-parametres-1
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Meta", "OpenAI", "xAI"]
dates: []
keywords: ["agent", "agents", "apache", "astra", "benchmarks", "claude", "gemini", "gpt-5.6", "mai", "memory", "omni", "open source"]
source: docs/RAG/collect-261001-ia-llm/nvidia-egale-le-modele-ouvert-d-openai-avec-4-fois-moins-de-parametres.md
source_anchor: ""
source_lines: [1, 48]
sha256: 8652da75942558f29af6adbc131224fd9f2e9a5b047ca2f7200cf5a408d9af7e
---

# 🧠 **RECHERCHE**

Sundar Pichai a confirmé ce matin que **Koray Kavukcuoglu** prend la tête de Google DeepMind avec le titre de Senior Vice President, en reportant directement au PDG d'Alphabet. Demis Hassabis, cofondateur et patron historique, quitte la direction opérationnelle pour devenir président de DeepMind et Chief Scientist d'Alphabet. Derrière le communiqué policé, c'est un laboratoire sous tension qui change de main.

Kavukcuoglu n'est pas un parachuté : arrivé en 2012, deux ans avant le rachat par Google, docteur en informatique de NYU sous la direction de Yann LeCun, il a contribué à DQN, AlphaGo et WaveNet. Il supervise désormais les modèles Gemini, la recherche frontière, l'app Gemini et les plateformes développeurs.

Google n'a pas sorti de modèle frontière depuis **Gemini 3.1 Pro en février** , pendant qu'Anthropic (Mythos) et OpenAI (GPT-5.6) ont pris l'avantage
**Gemini 3.5 Pro a raté trois dates de sortie** : juin, mi-juillet, puis août, et Gemini 3.6 Flash est jugé derrière OpenAI, Anthropic, xAI, Meta et plusieurs labos chinois sur les benchmarks d'intelligence
L'hémorragie de talents est spectaculaire : Jeff Dean part après **27 ans** chez Google monter sa propre boîte, Noam Shazeer (co-lead de Gemini) est passé chez OpenAI en juin, et le prix Nobel John Jumper est parti chez Anthropic avec Jonas Adler et Alexander Pritzel, trois des artisans d'AlphaFold
Climat interne dégradé selon Fortune : semaines de 60 heures, moral bas, et un mouvement syndical lancé en mai
Une équipe interne baptisée **Code Strike** a été montée pour rattraper le retard sur le code, terrain où Morningstar juge Anthropic et OpenAI « à des kilomètres devant »

L'audience n'est pas le problème : l'app Gemini vient de franchir le milliard d'utilisateurs mensuels et les modèles Gemma cumulent plus de 900 millions de téléchargements. Le problème est la frontière technologique, et les analystes lisent cette nomination comme un recentrage : moins de projets académiques et de world models chers à Hassabis, plus de LLM et d'outils pour développeurs. Pour vous, rien à tester aujourd'hui, mais le prochain Gemini dira si le virage a fonctionné.

Alibaba a publié **Qwen-MM-Plugins**, un dépôt open source sous licence Apache-2.0 qui ajoute des capacités multimodales aux agents de code que vous utilisez déjà. L'idée est maligne : plutôt que de vous forcer à adopter Qwen Code, Qwen transforme sa vision et son audio en une **couche d'outils réutilisable par n'importe quel agent**.

Une fois installé, votre agent gagne des fonctions qu'il découvre et enchaîne tout seul : lire une image ou une vidéo, faire de l'OCR, localiser un objet dans une photo, segmenter, transcrire de l'audio, recadrer. Vous pointez un fichier avec `@` et vous demandez en français ce que vous voulez.

**8 modules installables séparément** : core (images, vidéos, documents, fichiers 3D), api (vision, OCR, grounding, transcription, segmentation via les modèles Qwen VL et Omni), search (recherche web, extraction de pages, recherche inversée d'image), video-memory (mémoire hiérarchique pour vidéos longues), video-edit (génération et montage image, vidéo, audio), blender (modélisation et rendu 3D), freecad (CAO paramétrique, formats STEP, STL, FEM) et edu-agent
**Six agents compatibles** : Claude Code, Codex, Qoder, OpenClaw, Qwen Code et Gemini CLI, avec un fichier de configuration partagé
Un seul script fait tout : `curl -fsSL https://raw.githubusercontent.com/QwenLM/Qwen-MM-Plugins/main/install.sh | bash` installe, configure, vérifie et désinstalle
**Le module core ne demande aucune clé API.** Les autres réclament une clé DashScope, et Serper pour la recherche web
Gratuit à installer, vous ne payez que les appels d'API tiers que vous déclenchez

C'est le genre de brique qui change la nature d'un outil sans changer d'outil. Un agent de code qui savait seulement lire du texte peut désormais ouvrir une facture scannée, décrire une vidéo, sortir un rendu Blender ou manipuler un plan CAO. Et pendant qu'OpenAI et Anthropic vendent des modèles multimodaux fermés, Qwen distribue gratuitement de quoi rendre leurs propres agents voyants.

Depuis le 11 août, les **Meta Glasses sont interdites dans tous les tribunaux d'Angleterre et du Pays de Galles**. Le His Majesty's Courts & Tribunals Service, qui gère l'ensemble des juridictions pénales, civiles et familiales, les récupère à l'entrée et les rend à la sortie. Détail savoureux : le smartphone, lui, reste toléré, tant qu'il ne filme pas.

La raison est évidente une fois posée. Filmer sans autorisation dans un tribunal est un outrage à la cour, passible de poursuites. Or on ne colle pas discrètement un téléphone sur son visage, alors qu'une paire de lunettes à caméra intégrée permet d'enregistrer une audience entière sans que personne ne s'en rende compte.

**7 millions de paires** vendues en 2025, avec une caméra 12 mégapixels, de la vidéo jusqu'en 3K à 30 images par seconde, à partir de**299 dollars**
L'État de New York a devancé Londres : interdiction des lunettes connectées (Meta, Xreal, RayNeo) dans ses **1 240 tribunaux depuis le 20 juillet** , pour empêcher l'identification secrète de jurés et de témoins
La faille qui inquiète les institutions : la LED censée signaler l'enregistrement peut être neutralisée par un mod à 60 dollars, documenté par 404 Media
Le déclencheur remonte à février : lors d'un procès en Californie, un juge a dû menacer l'entourage de Mark Zuckerberg d'outrage à la cour pour qu'il retire ses lunettes
Un tribunal londonien avait déjà eu affaire aux lunettes connectées, mais pour un tout autre usage : un homme est accusé d'avoir été soufflé ses réponses en contre-interrogatoire via les haut-parleurs intégrés

Le rejet dépasse largement les prétoires : la chaîne de pubs Wetherspoons, des restaurants, des théâtres, des salles de concert et des conventions britanniques les bannissent aussi, après une vague de vidéos de personnes filmées sans leur accord et parfois identifiées en ligne. Meta a refusé de commenter. Quand un produit se fait confisquer à l'entrée des tribunaux d'un pays entier, ce n'est plus un problème d'image, c'est un problème de conception.

# 🧠 **RECHERCHE**

**L'IA n'est plus artificielle : des cerveaux cultivés en labo pour remplacer les LLM**

Des biologistes fabriquent, à partir de cellules souches, des organoïdes cérébraux humains : des amas de quelques millions de neurones vivants qui émettent de vraies ondes cérébrales après huit mois de culture. À UC San Diego, ces amas guident des robots dans des labyrinthes ; à Johns Hopkins, ils servent de base à du biocomputing ; une startup de Melbourne leur fait jouer à Pong et à Doom. Le pari : une intelligence vivante plutôt qu'artificielle, en alternative frontale aux grands modèles de langage.

**AMIE, l'IA médicale de Google, mène des consultations vidéo en temps réel**

Google Research et Google DeepMind font passer AMIE du chat texte à la visio. Construit sur Gemini et Project Astra avec une architecture multi-agents, le système interprète les signaux visuels et sonores, guide un examen physique à distance et raisonne sur le diagnostic en direct. Dans une étude randomisée avec des acteurs patients et des généralistes, AMIE est bien évalué sur l'anamnèse, la précision diagnostique et la qualité de communication, et les patients simulés ont préféré la vidéo au texte. Ça reste un système de recherche, non déployé cliniquement.

**Une nouvelle technique révèle les pensées cachées des modèles d'IA**

