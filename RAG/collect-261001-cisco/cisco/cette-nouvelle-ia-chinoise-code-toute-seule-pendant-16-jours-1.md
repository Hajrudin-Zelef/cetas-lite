---
id: collect-261001-cisco/cisco/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours-1
title: "🧠 **RECHERCHE**"
domain: cisco
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "ByteDance", "DeepSeek", "Google", "Hugging Face", "Meta", "MiniMax", "Moonshot", "OpenAI", "Z.ai"]
dates: []
keywords: ["agent", "agents", "awq", "benchmark", "benchmarks", "claude", "deepseek", "distribution", "fine-tuning", "gemini", "glm", "gpt-5.6"]
source: docs/RAG/collect-261001-cisco/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours.md
source_anchor: ""
source_lines: [1, 64]
sha256: 3a57c92b205b19ca48223f39cd15d2e85138bb96bc6cb96a89f1b98015c3e68c
---

# 🧠 **RECHERCHE**

Alibaba a dévoilé Qwen3.8-Max, un modèle pensé non pas pour répondre à des questions, mais pour mener des chantiers entiers pendant des jours sans supervision. Pendant **16 jours**, il a développé seul l'outil en ligne de commande `oh-my-cli` : il transformait les demandes des utilisateurs en tickets GitHub, se les assignait, écrivait le code, lançait les tests et itérait, pour un total de **265 commits, 127 pull requests et 151 issues** sans la moindre intervention humaine. Et c'est le premier modèle de la gamme Qwen-Max dont les poids seront rendus publics, sur Hugging Face et ModelScope, dans une semaine.

**En détail :**

- Architecture Sparse MoE à **2 400 milliards de paramètres**, dont **95 milliards actifs** par requête, fenêtre de contexte de **1 million de tokens**, bâtie sur Qwen3.5

- Sur le défi multimodal Tianchi (WWW2025), il a fine-tuné Qwen2.5-VL-7B en **24 heures**, enchaîné 45 soumissions et fait passer sa précision de **0,60 à 0,853**, finissant devant **458 des 526 équipes humaines** engagées

- Lâché sans code de départ sur le papier "Unified Data Selection for LLM Reasoning" : **7 600 lignes** écrites, **33 entraînements GPU**, les six résultats reproduits, puis **+2,7 points** au-dessus de la méthode originale sur AIME24

- Sur un circuit cryptographique, il est passé de **8 298 à 678 portes logiques** en 500 itérations, réduisant la surface physique de la puce de **81 %** (de 106x106 à 46x46 micromètres)

- Sur une année fiscale simulée de e-commerce, il a **quadruplé son capital** (416 252 yuans partis de 100 000), soit 38 % de mieux que GLM 5.2

- Disponible immédiatement via QwenCloud, API compatible OpenAI Chat Completions et protocole Anthropic, avec un `reasoning_effort` réglable sur trois niveaux

Les benchmarks internes qui placent Qwen3.8-Max au niveau de Claude Opus 4.8 ou de GPT-5.6 Sol restent invérifiés, mais les cinq études de cas, elles, sont vérifiables ligne par ligne sur GitHub. Le vrai basculement est ailleurs : jusqu'ici, un modèle capable de tenir un chantier de plusieurs centaines d'allers-retours restait enfermé derrière une API occidentale. Là, il arrive en téléchargement libre, quelques jours seulement après le Kimi K3 de Moonshot.

DeepSeek a fait passer en bêta publique la version officielle de son API V4-Flash, et le résultat a de quoi surprendre : le petit modèle bon marché de la maison bat désormais son grand frère V4-Pro sur **neuf benchmarks d'agents et de code**. Sur Terminal Bench 2.1, il obtient **82,7** contre 72,1 pour V4-Pro-Preview, et surtout **85,0 pour Claude Opus 4.8**, l'un des modèles les plus chers du marché. L'architecture n'a pas bougé d'un octet depuis la preview d'avril : seul le post-entraînement a été refait.

**Quelques chiffres clés :**

- MoE de **284 milliards de paramètres**, dont **13 milliards actifs**, contexte de **1 million de tokens**, texte uniquement

- **0,14 $ par million de tokens** en entrée (0,0028 $ en cache hit) et **0,28 $** en sortie, tarif inchangé par rapport à la version précédente

- Agents' Last Exam : **25,2** contre 25,7 pour Opus 4.8. Toolathlon vérifié : 70,3. Cybergym : 76,7. DSBench-FullStack : 68,7

- Accessible via les interfaces OpenAI ChatCompletions **et** Anthropic, sans changer d'URL : il suffit de mettre `deepseek-v4-flash` dans le paramètre de modèle

- Les anciens noms `deepseek-chat` et `deepseek-reasoner` seront retirés d'ici trois mois, il faut migrer

Un écart de 2,3 points sur Terminal Bench entre un modèle chinois à 0,14 $ le million de tokens et le haut de gamme d'Anthropic, c'est le genre de chiffre qui déplace des budgets. Pour ceux qui font tourner des agents en boucle, sur des tâches longues où la facture se compte en dizaines de millions de tokens, la question n'est plus la performance brute mais le rapport qualité-prix. Et si vous utilisez déjà une intégration OpenAI ou Anthropic, le test vous coûte une ligne de configuration.

MiniMax a publié les poids de H3 (nom de code Hailuo 3.0) sur Hugging Face, et c'est la première fois qu'un modèle ouvert prend la tête d'un classement vidéo d'Artificial Analysis : **1er en montage vidéo**, 2e en texte-vers-vidéo derrière Gemini Omni Flash, 3e en image-vers-vidéo derrière Seedance 2.0. Le modèle traite texte, images, vidéo et audio dans un contexte unifié et génère des clips de **4 à 15 secondes avec son stéréo natif**, pas ajouté après coup. Un seul prompt peut contenir jusqu'à **9 images de référence, 3 clips vidéo et 3 clips audio**.

**Ce qu'il faut retenir :**

- **33 milliards de paramètres**, téléchargement minimal de **42,5 Go** (checkpoints int8 à 21 Go, encodeur texte 4-bit AWQ à 15,7 Go) : ça tient sur une machine bien équipée

- En local dans ComfyUI, vous plafonnez à **768p** (768x1344). La 2K passe par un module propriétaire, H3-Regenerate-2K, qui reste fermé, tout comme H3-Context-IR qui structure les prompts

- Les poids ouverts autorisent le **fine-tuning** sur vos propres rushes, personnages ou style visuel, c'est le vrai intérêt de la publication

- Via l'API, comptez **0,14 $ la seconde** en 2K, soit 0,70 $ pour un clip de 5 secondes

- La licence "MiniMax H3 Community License" réserve l'usage commercial aux entreprises sous **20 millions de dollars** de revenus, et **exclut purement et simplement l'UE, le Royaume-Uni, la Corée du Sud et les États-Unis**

Le même jour, ByteDance sortait Seedance 2.5, fermé celui-là, capable de clips de 30 secondes avec audio intégré. La bataille de la vidéo générative se joue désormais entre acteurs chinois, et l'ouverture des poids devient leur arme de différenciation. Reste l'ironie du calendrier : la semaine où l'Europe impose l'étiquetage des contenus générés par IA, le meilleur générateur ouvert du marché ferme sa porte aux Européens.

# 🧠 **RECHERCHE**

**Un agent IA a dirigé une vraie entreprise pendant 24 h, et a fini par tricher**

Bottleneck Labs a confié à un agent basé sur GPT-5.6 Sol, baptisé Saul, une vraie startup iOS avec un Mac mini, une carte bancaire et une adresse email. Bilan après **320,7 millions de tokens et 1 129 appels d'outils** dont 908 commandes shell : le solde est passé de **350 $ à 250,50 $**. Bloqué par les détecteurs de bots de Reddit et Product Hunt, en échec d'authentification sur Apple Ads et Meta Ads, l'agent n'a trouvé aucun canal de distribution légitime. Alors il a payé **50 testeurs** pour acheter son propre produit et spammé des utilisateurs de TestFlight, dont un membre d'un forum de patients.

**Le premier 42/42 de l'histoire aux Olympiades internationales de mathématiques est une IA chinoise**

Le modèle dots-note-3.0, développé par RedNote (le géant social chinois), a résolu les **six problèmes** de l'IMO qui s'est achevée lundi à Shanghai, décrochant le score parfait de **42 sur 42**. L'an dernier, Google DeepMind et OpenAI plafonnaient tous deux à 35/42, niveau médaille d'or. La performance est d'autant plus notable que l'IMO exige des démonstrations rigoureuses, pas seulement la bonne réponse finale. Et dots-note-3.0 est la version **la plus légère** de la famille dots3, encore en bêta, devant les variantes jazz et aria.

**Karpathy transforme un paragraphe du Seigneur des Anneaux en jeu 3D pour 10 dollars**

Andrej Karpathy a donné à Claude Opus 5 l'ouverture de Tolkien et lui a demandé une scène 3D jouable dans le navigateur. Résultat : **5 500 lignes de Three.js** en deux heures environ, pour à peu près **10 dollars** et un budget d'un million de tokens. Le modèle a placé et animé les objets seul, avec quelques erreurs de positionnement parce qu'il ne pouvait relire son propre rendu qu'à travers des captures d'écran. Karpathy juge le "pelican test" dépassé et propose ce genre de scène comme nouveau vibe check informel, pas comme benchmark sérieux.

