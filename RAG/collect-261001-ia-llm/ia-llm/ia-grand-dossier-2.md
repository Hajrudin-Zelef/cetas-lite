---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-2
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Cohere", "Google", "Microsoft", "Nvidia", "OpenAI"]
dates: []
keywords: ["attention", "benchmarks", "cohere", "compute", "diffusion", "exploit", "gpu", "nvidia", "research", "training"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [98, 154]
sha256: b9fd318be06ca72659be00fd408c6171eb47d9bc0f53237b5f2ca02537bce6b4
---

# IA — Le grand dossier

**Pourquoi c'est une révolution, pas juste un record :**

- **ReLU au lieu de sigmoïde/tanh.** La fonction d'activation ReLU (max(0, x)) ne sature pas pour les valeurs positives : l'entraînement est plusieurs fois plus rapide. Détail technique, conséquence stratégique : on peut enfin entraîner profond.
- **Dropout.** Pendant l'entraînement, on « éteint » aléatoirement une partie des neurones : cela force le réseau à ne pas dépendre d'un petit nombre de neurones et réduit le surapprentissage. Technique proposée par Hinton et son équipe en 2012.
- **Deux GPU en parallèle.** Le modèle (60 millions de paramètres) est découpé sur deux cartes GTX 580. C'est la démonstration que **le parallélisme matériel est la condition du deep learning moderne** — un point qui reste vrai en 2026 avec les clusters de dizaines de milliers de GPU.
- **Augmentation de données.** Les images d'entraînement sont artificiellement variées (recadrages, miroirs) : le réseau voit « plus » de données qu'il n'y en a.

**Conséquence immédiate.** En quelques années, les CNN écrasent la vision par ordinateur : VGG (Oxford, Simonyan & Zisserman, 2014) montre que la profondeur seule suffit ; GoogLeNet (Szegedy et al., Google, 2014) gagne ImageNet 2014 ; **ResNet** (He et al., Microsoft Research, 2015) introduit les **connexions résiduelles** (skip connections) qui permettent d'entraîner des réseaux de 152 couches, puis 1000 couches. Le taux d'erreur ImageNet tombe sous les 5 %, puis sous les 3 % — meilleur que l'humain (~5 % d'erreur top-5 estimée). En 2015-2016, la vision par ordinateur est considérée comme « résolue » pour la classification d'images.

**Pour le RAG de Zelef :** les CNN et ResNet restent pertinents pour l'OCR, la détection d'équipements sur photos (lecture de plaques, d'étiquettes d'onduleurs), et les modèles de vision multimodaux actuels descendent tous de cette lignée.

### 1.3. Le détour par les jeux : le deep reinforcement learning (2013-2017)

Parallèlement, une autre révolution se joue dans les jeux :

- **2013 — DQN (Deep Q-Network), DeepMind.** Mnih et al. apprennent à une IA à jouer à des jeux Atari directement à partir des pixels, par apprentissage par renforcement profond. Première démonstration qu'un même algorithme peut maîtriser des dizaines de jeux sans règles codées en dur.
- **Mars 2016 — AlphaGo bat Lee Sedol 4-1.** Le champion du monde de go (jeu au nombre de positions astronomique, ~10^170) est battu par le système de DeepMind combinant réseaux de neurones et recherche arborescente Monte-Carlo. Les experts pensaient cet exploit à « des décennies » de distance. C'est l'événement qui fait passer l'IA du statut de curiosité académique à celui de sujet de une des journaux télévisés mondiaux.
- **2017 — AlphaGo Zero / AlphaZero.** Le système apprend **sans aucune donnée humaine**, uniquement par auto-jeu (self-play), et bat la version précédente 100-0. Leçon philosophique majeure : **les données humaines ne sont pas toujours le plafond ; le calcul + la recherche peuvent les dépasser**. Cette leçon reviendra en force avec les modèles de raisonnement (2024-2025).
- **2020 — AlphaFold 2.** DeepMind résout le problème du repliement des protéines (50 ans de biologie structurale), avec une précision proche de l'expérimentation. En octobre **2024**, Demis Hassabis et John Jumper reçoivent le **prix Nobel de chimie** (partagé avec David Baker) pour ces travaux. C'est la première fois qu'une IA « pure » vaut un Nobel à ses créateurs — signal institutionnel fort.

**À retenir pour la suite :** DeepMind a toujours privilégié la voie « RL + recherche + compute », distincte de la voie « pré-entraînement sur texte » d'OpenAI. Les deux voies convergent en 2024-2026 avec les modèles de raisonnement (RL appliqué aux LLM).

### 1.4. Les GAN : la génération d'images avant la diffusion (2014-2019)

**2014 — Ian Goodfellow** (alors doctorant sous Yoshua Bengio à l'Université de Montréal) invente les **GAN (Generative Adversarial Networks)** : deux réseaux s'affrontent, un générateur qui fabrique des images et un discriminateur qui essaie de détecter les faux. Le papier « Generative Adversarial Nets » (Goodfellow et al., NeurIPS 2014) ouvre l'ère de la génération d'images réalistes.

Progrès marquants : DCGAN (2015), puis **StyleGAN** (Karras et al., NVIDIA, 2018-2019) qui génère des visages photoréalistes impossibles à distinguer pour un humain non averti. Les GAN dominent la génération d'images jusqu'en 2021-2022, avant d'être détrônés par la diffusion (section 1.8).

### 1.5. Révolution n° 2 — Le Transformer : « Attention Is All You Need » (2017)

**Le fait.** Le 12 juin 2017, huit chercheurs de Google (équipes Google Brain / Google Research) publient sur arXiv le papier **« Attention Is All You Need »** (présenté à NeurIPS 2017). Auteurs, dans l'ordre : **Ashish Vaswani, Noam Shazeer, Niki Parmar, Jakob Uszkoreit, Llion Jones, Aidan N. Gomez, Łukasz Kaiser, Illia Polosukhin**.

**L'idée.** Jusqu'alors, le traitement du langage reposait sur des réseaux récurrents (RNN, LSTM) qui lisent le texte mot à mot, séquentiellement — donc lentement et avec une mémoire limitée. Le Transformer **supprime entièrement la récurrence** et la remplace par la **self-attention** : chaque mot « regarde » tous les autres mots de la phrase en une seule opération parallélisable.

La formule au cœur du papier, restée inchangée depuis :

```
Attention(Q, K, V) = softmax(QK^T / √d_k) V
```

Q (queries), K (keys), V (values) sont trois projections apprises des représentations des mots. En clair : le modèle apprend, pour chaque mot, **à quels autres mots prêter attention** et avec quel poids.

**Pourquoi c'est la révolution la plus importante du lot :**

1. **Parallélisme total.** Plus de lecture séquentielle : tout le texte est traité d'un coup. Sur GPU, c'est 10 à 100× plus rapide à entraîner que les LSTM. Sans ça, entraîner sur des milliards de mots serait impraticable.
2. **Dépendances longue distance.** Dans un LSTM, l'information du début d'un long texte se dilue ; en self-attention, le mot n° 5000 peut directement « voir » le mot n° 3.
3. **Généricité.** La même architecture sert pour le texte (GPT, BERT), les images (ViT, 2020), l'audio, les protéines, le code. C'est devenu **l'architecture universelle** de l'IA moderne.

**Destinées croisées des huit auteurs (vérifié) :** fait remarquable, **aucun des huit ne travaille plus chez Google** en 2026. La plupart ont fondé ou rejoint des startups/labos : Aidan Gomez a cofondé **Cohere** (2020), Llion Jones et Jakub Pachocki... (à vérifier pour le détail individuel — voir section 2). Illia Polosukhin a cofondé **NEAR Protocol**. Cette diaspora illustre la diffusion du savoir-faire Transformer dans tout l'écosystème.

### 1.6. 2018 : l'année où le langage devient « pré-entraînable »

Deux papiers fondateurs, deux philosophies :

- **Juin 2018 — GPT-1 (OpenAI, Radford et al.).** « Improving Language Understanding by Generative Pre-Training » : on pré-entraîne un Transformer **décodeur** (qui prédit le mot suivant) sur du texte brut du web, puis on l'affine (fine-tune) sur chaque tâche. 117 millions de paramètres. C'est modeste, mais le principe est posé : **le pré-entraînement génératif non supervisé transfère aux tâches en aval**.
- **Octobre 2018 — BERT (Google, Devlin et al.).** « BERT: Pre-training of Deep Bidirectional Transformers » : Transformer **encodeur** bidirectionnel entraîné à prédire des mots masqués. 340 millions de paramètres (version Large). BERT écrase les benchmarks de compréhension (GLUE, SQuAD) et devient le standard industriel 2019-2021 pour la classification, la recherche sémantique, les chatbots d'entreprise.

