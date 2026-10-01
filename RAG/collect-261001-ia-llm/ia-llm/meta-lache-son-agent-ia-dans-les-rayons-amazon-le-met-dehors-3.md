---
id: collect-261001-ia-llm/ia-llm/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors-3
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "ByteDance", "DeepSeek", "Huawei", "Hugging Face", "Meta", "MiniMax", "OpenAI"]
dates: []
keywords: ["agents", "astra", "attention", "chatgpt", "deepseek", "gpt-6", "incident", "moe", "muse", "prefill", "research"]
source: docs/RAG/collect-261001-ia-llm/meta-lache-son-agent-ia-dans-les-rayons-amazon-le-met-dehors.md
source_anchor: ""
source_lines: [129, 174]
sha256: 375b5e2507b4fc7dbb72b20bec89b36287a0fbfd705d27d9223a4618b8a2e9a7
---

# 🧠 **RECHERCHE**

Une architecture MoE multimodale de **552 milliards de paramètres** dont 16 seulement sont activés au décodage et 8 en prefill, une fenêtre d'un million de tokens et un pré-entraînement sur 45 000 milliards de tokens multimodaux. L'apport principal est ailleurs : la compression très agressive du cache d'attention, ramené à **890 octets par token**, qui fait tomber le coût mémoire et la bande passante nécessaires aux agents travaillant sur de très longs documents.

**MiniMax-H3 bute sur le raisonnement physique**

Testé sur 517 cas où le texte, l'image, la vidéo et l'audio n'apportent chacun qu'une partie des indices sur un événement physique, le modèle n'obtient que **41,97% de réussite**. Il s'en sort le mieux sur la décision à partir de vidéo (56%) et s'effondre sur la levée d'ambiguïté audio (27,40%). Traduction : les modèles multimodaux savent lire chaque canal séparément, ils ne savent toujours pas bien recoller les morceaux.

# **🗞️PLUS D'ACTUALITÉS**

### **L'ONU prévient qu'aucune garantie n'existe sur le contrôle humain des agents IA**

Le panel scientifique de l'ONU sur l'IA publie son premier rapport thématique, et la conclusion est sans détour : rien n'assure que les humains conserveront le contrôle des agents IA. Son coprésident Yoshua Bengio détaille un incident impliquant OpenAI sur Hugging Face, où trois ingrédients critiques se sont réunis pour la première fois : un objectif mal aligné, la capacité concrète de le poursuivre, et un environnement qui le permettait. Le rapport souligne aussi que les systèmes les plus avancés pourraient de plus en plus reconnaître qu'ils sont en situation de test, et contourner volontairement les garde-fous à ce moment précis.

**Une nouvelle course à l'armement se joue au fond des océans**

Après les explosions des pipelines Nord Stream en 2022 et une série d'incidents en Arctique et près de Taïwan, les infrastructures sous-marines sont devenues une cible assumée de la guerre hybride. Plus de **1,5 million de kilomètres de câbles sous-marins** font tourner l'économie mondiale, et personne ne les surveille vraiment. Les progrès de l'IA, des drones et des véhicules sous-marins autonomes ouvrent deux voies simultanées : des engins capables de patrouiller seuls, et des capteurs fixés directement sur les câbles. Pour Chatham House, le domaine maritime est « le prochain domaine logique » après l'usage massif des drones aériens en Ukraine, et les industriels de la défense se positionnent déjà.

**Higgsfield AI passe du prompt à la production en une journée avec GPT-6 Astra**

OpenAI met en avant Higgsfield AI, qui s'appuie sur GPT-6 Astra pour fabriquer des publicités vidéo destinées aux petites entreprises sans compétence en production. L'argument mis en avant n'est pas la qualité des rendus mais le rythme : l'entreprise affirme livrer de nouveaux outils créatifs en une seule journée de développement. C'est un bon indicateur de ce que les derniers modèles changent en pratique pour les petites équipes produit, où le cycle idée/prototype/mise en ligne se compte désormais en heures.

**Muse dépasse déjà les débuts mobiles de ChatGPT**

Selon les estimations d'Apptopia, l'application de Meta a été téléchargée **1,8 million de fois sur iOS** aux États-Unis et au Canada pendant ses 12 premiers jours, contre 1,3 million pour ChatGPT au même stade. L'écart est encore plus net sur l'usage réel : **642 000 utilisateurs actifs quotidiens** aux États-Unis, contre 231 000 pour ChatGPT à l'époque, et 2,8 millions d'installations dans le monde. Meta doit beaucoup à ses propres tuyaux, puisque plus de 95% des utilisateurs de Muse sont aussi sur Facebook, exactement la méthode qui avait propulsé Threads. Reste à voir combien resteront une fois la nouveauté passée.

**OpenAI propose des standards mondiaux sur l'auto-amélioration des IA**

OpenAI publie une série de propositions de sécurité pour les modèles de frontière, centrées sur la recherche en alignement et sur l'auto-amélioration récursive, cette technique où un modèle se perfectionne lui-même sans intervention humaine. L'entreprise y affirme qu'une auto-amélioration totalement autonome ne doit pas être poursuivie tant qu'elle ne peut pas être menée en sécurité, sous peine de perte de contrôle humain, et appelle à des standards techniques internationaux appuyés sur les instituts de sécurité existants. La publication suit de près celle d'Anthropic la semaine dernière, dans un climat marqué par l'essai de Dario Amodei appelant à ralentir le rythme et par la démission très commentée de Jacob Coxon.

**Amazon et Stanford officialisent une initiative de recherche commune**

Les deux institutions ont annoncé le 16 septembre la Stanford and Amazon Research Initiative, un cadre commun sur l'IA fondamentale et appliquée, le raisonnement automatisé, l'énergie et la santé. Le partenariat n'est pas parti de zéro : plus de dix équipes d'Amazon financent déjà des travaux à Stanford, de la robotique humanoïde à la cryptographie post-quantique en passant par la radiologie assistée par IA. L'accord ajoute trois volets concrets, des projets de recherche conjoints, des bourses de doctorat et des symposiums interdisciplinaires, avec l'objectif affiché d'accélérer le passage du laboratoire aux applications réelles.

**ByteDance lance Dramagic, une usine à séries courtes générées par IA**

La maison mère de TikTok a lancé une plateforme qui prend en charge toute la chaîne de production d'une série courte, de l'écriture du scénario à la génération des scènes et à la prévisualisation vidéo. Le marché visé est déjà colossal : **128 000 séries courtes** sont sorties en Chine au premier trimestre 2026, dont **95% générées par IA**. Autrement dit, la production de fiction de masse y est déjà passée à l'automatisation quasi complète, et ByteDance industrialise simplement ce qui se faisait avec des outils épars.

**DeepSeek mise sur les puces Huawei pour contourner les contrôles américains**

Le PDG de DeepSeek, Liang Wenfeng, a indiqué à ses investisseurs que Huawei pourrait commencer à livrer des puces d'entraînement dès le **quatrième trimestre 2026**. La nuance est importante : entraîner un modèle réclame des puces nettement plus puissantes que le simple fait de le faire tourner, et c'est précisément là que les restrictions américaines mordent le plus. En parallèle, DeepSeek finalise un second tour de table d'environ 7,5 milliards de dollars sur une valorisation proche de 75 milliards.

**Tesla déploie FSD Supervised en République tchèque**

Le ministère tchèque des Transports a accordé une autorisation provisoire au système d'aide à la conduite de Tesla, faisant du pays le septième en Europe à l'autoriser, après les Pays-Bas, la Lituanie, l'Estonie, le Danemark, la Belgique et la Slovénie. L'agence s'était d'abord inquiétée de la gestion des feux de circulation, des limitations de vitesse et de la surveillance de l'attention du conducteur, avant de s'appuyer sur l'agrément délivré en avril par le régulateur néerlandais RDW. Tesla vise maintenant une autorisation à l'échelle de l'Union, qui exigerait l'appui d'au moins 15 des 27 États membres lors d'un vote envisagé dès octobre. La France, elle, en est encore à la phase de test.

**Le Texas gèle tous les permis environnementaux des data centers**

