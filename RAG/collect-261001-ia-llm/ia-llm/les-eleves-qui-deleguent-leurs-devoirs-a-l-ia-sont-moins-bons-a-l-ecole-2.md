---
id: collect-261001-ia-llm/ia-llm/les-eleves-qui-deleguent-leurs-devoirs-a-l-ia-sont-moins-bons-a-l-ecole-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "Huawei", "OpenAI", "Samsung"]
dates: []
keywords: ["apache", "benchmark", "embeddings", "gpu", "gqa", "llama", "mai", "nano-texture", "transcription"]
source: docs/RAG/collect-261001-ia-llm/les-eleves-qui-deleguent-leurs-devoirs-a-l-ia-sont-moins-bons-a-l-ecole.md
source_anchor: ""
source_lines: [79, 127]
sha256: 3583b33350b97ba2bfbb91649c299baa828bfdabe81c9f24da1c37db9b9f287e
---

# 🧠 **RECHERCHE**

**Trois versions** : v6, le modèle principal, réservé aux abonnés Pro et Premier ;**v6-wild** , plus expérimental et imprévisible, également payant ;**v6-mini** , gratuit pour tous, optimisé pour la vitesse.
La vraie nouveauté fonctionnelle est **l'édition ciblée** : vous pouvez demander en langage naturel de modifier uniquement le refrain ou une ligne de paroles, sans régénérer tout le morceau. Fini le tout ou rien.
**Génération multimodale** : le modèle accepte du texte, de l'audio et des images en entrée, et sait combiner des éléments issus de sources différentes.
L'entraînement mêle **musique Warner sous licence** et données utilisateurs de Suno, les catalogues de Believe, TuneCore et BMG étant en cours d'intégration. Les montants des accords ne sont pas divulgués.
Suno **refuse de préciser** quels catalogues et quel volume ont servi à l'entraînement, et a demandé le placement de ces informations sous scellé devant le tribunal fédéral (détail des accords).

Le paysage juridique se scinde en deux. D'un côté Warner, BMG et Believe, qui ont choisi de monétiser plutôt que de bloquer. De l'autre Universal et Sony, qui maintiennent leurs poursuites, pendant que Suno traîne une condamnation pour violation de droits d'auteur prononcée en Allemagne en juillet 2026. C'est le scénario qui se dessine partout dans l'IA générative : les ayants droit assez gros pour négocier deviennent actionnaires du problème, les autres continuent au tribunal. Le déploiement est en cours sur suno.com/app, et v6-mini permet de tester la chose sans payer.

L'iPhone Duo dévoilé lors du keynote « Surprise and Shine » est le premier téléphone pliable d'Apple, avec un **écran interne de 7,6 pouces** et un **écran externe de 5,4 pouces** utilisable replié. Mais la partie la plus intéressante est invisible : la charnière est conçue et fabriquée à l'aide d'algorithmes d'IA, unité par unité, sur la chaîne de production.

Selon Johny Srouji, directeur matériel d'Apple, des **algorithmes d'IA apparient chaque charnière individuelle avec le boîtier qui lui convient le mieux** , pour garantir un alignement parfait malgré les tolérances de fabrication.
Un **laser confocal scanne la topologie de chaque exemplaire** , puis une imprimante dépose**jusqu'à 25 micro-couches d'un photopolymère sur mesure** pour éliminer les ondulations résiduelles. Chaque appareil reçoit donc une correction qui lui est propre.
Le problème visé est connu : un pliable subit **deux fois plus d'usure mécanique** qu'un smartphone classique, et c'est là que la concurrence s'est cassé les dents.
S'y ajoutent une finition **nano-texture antireflet** et une stratégie de lamination multicouche.
Tarifs et calendrier : **2 339 €** en 256 Go, 2 589 € en 512 Go, 3 089 € en 1 To et 3 839 € en 2 To, précommandes le**16 octobre à 14h** , sortie le**23 octobre** (tarifs détaillés).

Apple arrive des années après Samsung et Huawei sur ce marché, et mise tout sur la durabilité comme différenciateur. Le point notable pour qui suit l'IA n'est pas le téléphone, c'est le déplacement du curseur : l'IA quitte le logiciel pour entrer dans la micro-fabrication de précision, avec une correction personnalisée pour chaque unité sortant de la chaîne. Peu de constructeurs communiquent sur cet usage. Reste l'inconnue habituelle : Apple ne publie aucun chiffre de cycles de pliage garantis, ni de comparaison de durabilité avec les pliables concurrents. Il faudra attendre les premiers appareils vieillis pour trancher.


# 🧠 **RECHERCHE**

### **Un développeur entraîne seul un modèle de 3,8 milliards de paramètres pour 998 dollars**

Hugo Vergnes a entraîné depuis zéro un LLM de **3,8 milliards de paramètres** sur **65 milliards de tokens**, en **43 heures de GPU loués**, pour un coût total de **998 dollars**. Résultat : **0,384 au benchmark CORE**, au-dessus de nanochat de Karpathy à budget comparable. Deux enseignements pratiques : les B200 loués offrent un meilleur rapport coût sur travail que les H100 pour ce type de run, et l'auteur a construit little-lm, un framework où chaque expérience tient en quelques lignes de YAML. L'architecture est de type Llama, avec RMSNorm, RoPE, GQA, QK-norm et value embeddings à la ResFormer. La zone entre le jouet pédagogique et le laboratoire de recherche est plus accessible qu'on ne le croit.

**IBM ouvre un modèle de prévision de séries temporelles sous licence commerciale**

Granite Time Series PatchTST-FM-r2 est un modèle de fondation de **385 millions de paramètres** qui produit des prévisions sans aucun entraînement préalable sur vos données. Il gère un contexte de **8 192 points**, sort des prévisions probabilistes via **99 quantiles** et sait combler les valeurs manquantes. Au 8 septembre, il est **2e du classement GIFT-Eval** parmi les modèles zero-shot reproductibles, et **1er parmi ceux sous licence permissive**. Poids, architecture et code d'inférence sont publiés sous double licence Apache 2.0 et OpenMDW 1.0, donc utilisables en production : prévision de demande, de prix, de consommation d'énergie ou de trafic.

**Anthropic publie un modèle économique qui classe les prédictions de son propre PDG dans le scénario extrême**

Anthropic a modélisé trois trajectoires pour l'économie américaine jusqu'en 2030. Dans le scénario le plus violent, la production **double tous les 4,5 ans** et le chômage des travailleurs du savoir atteint **17,9 %**. Or ce sont précisément les chiffres que Dario Amodei avançait publiquement en mai dernier. Son propre laboratoire vient donc de ranger la prédiction de son dirigeant dans la case des hypothèses minoritaires, ce qui en dit long sur l'écart entre la communication publique du secteur et ses modèles internes.

**Les lecteurs préfèrent-ils les histoires écrites par une IA ?**

Cambridge University Press pose publiquement la question, qui devient sérieuse à mesure que les modèles produisent de la fiction lisible. À ce stade, seul l'intitulé du billet est accessible : ni la méthodologie ni les résultats de l'étude ne sont consultables, et il vaut mieux attendre la publication complète avant d'en tirer une conclusion. Le sujet à surveiller reste celui de la perception, un même texte n'étant pas jugé de la même façon selon qu'on croit ou non qu'une machine l'a écrit.

**Les batteries battent un nouveau record d'installation aux États-Unis**

**20,2 GWh** de stockage installés au deuxième trimestre 2026, l'équivalent de la consommation quotidienne de **700 000 foyers**, avec sept chantiers de plus d'un gigawattheure mis en service. Le pays est sur une trajectoire de **71 GWh sur l'année**, en hausse de 20 %. Les data centers représentent environ **trois quarts** des nouvelles batteries du segment commercial, pendant que le résidentiel recule de **16 %** après la fin d'un crédit d'impôt. La quasi-totalité des cellules vient encore de Chine.

# **🗞️PLUS D'ACTUALITÉS**

### **L'Apple Watch écoute désormais en permanence, et transcrit**

Trois fonctions arrivent sur la montre : **Audio Intelligence**, qui détecte sirènes, alarmes, sonnettes et pleurs de bébé **en local, sans iPhone à proximité**, **Live Rewind**, qui transcrit les **15 dernières secondes** d'une conversation sur une double pression de la Digital Crown et enregistre le texte dans la nouvelle app Siri, et Siri Recap. La détection de sons relève clairement de l'accessibilité. La transcription rétroactive, beaucoup moins : elle pose la question du consentement de votre interlocuteur, qui ne sait pas qu'il vient d'être retranscrit. Venant de l'entreprise qui a bâti son marketing sur la vie privée, le virage est notable, et il ressemble à une réponse anticipée aux appareils d'écoute passive d'OpenAI, Friend ou Bee.

**Un spécialiste des protéines démonte le scénario du supervirus conçu par IA**

