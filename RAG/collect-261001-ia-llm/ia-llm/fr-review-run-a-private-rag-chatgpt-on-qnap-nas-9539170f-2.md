---
id: collect-261001-ia-llm/ia-llm/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f-2
title: "fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f"
domain: ia-llm
role: reference
task: reference
actors: ["Nvidia", "OpenAI", "TensorRT-LLM"]
dates: []
keywords: ["chatgpt", "arr", "gpu", "nvidia", "tensorrt"]
source: docs/RAG/collect-261001-ia-llm/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f.md
source_anchor: ""
source_lines: [32, 74]
sha256: a3aec99df5faa0d620ed7d8718cc689f4aed9444e3cde3473e84ce1f0057fbce
---

# fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f

Après avoir installé les pilotes, vérifiez que le GPU est reconnu et fonctionne correctement au sein de la VM. Vous pouvez utiliser le gestionnaire de périphériques sous Windows ou les outils de ligne de commande pertinents sous Linux pour vérifier l'état du GPU.
Dépannage et astuces
- Compatibilité: Consultez les sites Web des fabricants de QNAP et de GPU pour connaître les notes de compatibilité spécifiques ou les mises à jour du micrologiciel susceptibles d'affecter la fonctionnalité de relais.
- Performance : Surveillez les performances de votre VM et ajustez les allocations de ressources si nécessaire. Assurez-vous que votre NAS dispose de suffisamment d'espace pour le refroidissement, surtout après l'ajout d'un GPU hautes performances.
- Mise en réseau et stockage : Optimisez les paramètres réseau et les configurations de stockage pour éviter les goulots d'étranglement qui pourraient avoir un impact sur les performances des applications VM.
Chat NVIDIA avec RTX – Chat privéGPT
Bien qu'il soit facile de s'arrêter là (créer une machine virtuelle Windows avec accès GPU), nous avons poussé plus loin dans cette expérience pour offrir aux entreprises un moyen unique de tirer parti de l'IA en toute sécurité, en exploitant les performances du NAS basé sur NVMe. Dans notre cas, la VM exploitait un stockage protégé par RAID5 qui offrait des performances de 9.4 Go/s en lecture et 2.1 Go/s en écriture.
NVIDIA a récemment lancé un logiciel nommé Chat with RTX . Chat with RTX révolutionne l'interaction avec l'IA en offrant une expérience personnalisée grâce à l'intégration d'un modèle de langage étendu (LLM) basé sur GPT et d'un ensemble de données local et unique. Il permet notamment de traiter des documents, des notes, du contenu multimédia, des vidéos YouTube, des listes de lecture, et bien plus encore.
Cette application clé en main exploite la puissance de la génération augmentée par récupération (RAG), combinée à l'efficacité du LLM optimisé par TensorRT et aux capacités à grande vitesse de l'accélération RTX. Ceux-ci fournissent des réponses contextuelles, à la fois rapides et très pertinentes. Fonctionnant directement sur votre bureau ou poste de travail Windows RTX, cette configuration garantit un accès rapide aux informations et un haut degré de confidentialité et de sécurité, car tous les traitements sont gérés localement.
La mise en œuvre d'un LLM avec les capacités RAG offre une excellente solution pour les professionnels et les utilisateurs expérimentés qui donnent la priorité à la confidentialité, à la sécurité et à l'efficacité personnalisée. Contrairement aux modèles publics tels que ChatGPT, qui traitent les requêtes sur Internet, un LLM local fonctionne entièrement dans les limites de votre NAS QNAP.
Cette fonctionnalité hors ligne garantit que toutes les interactions restent privées et sécurisées. Cela permet aux utilisateurs de personnaliser la base de connaissances de l'IA en fonction de leurs besoins spécifiques, qu'il s'agisse de documents d'entreprise confidentiels, de bases de données spécialisées ou de notes personnelles. Cette approche améliore considérablement la pertinence et la rapidité des réponses de l'IA, ce qui en fait un outil précieux pour ceux qui ont besoin d'informations immédiates et contextuelles sans compromettre la confidentialité ou la sécurité des données.
A noter également, et cela peut paraître évident, l'ajout d'un GPU au NAS simplifie directement le lien entre les données d'une entreprise et le LLM. Il n'est pas nécessaire de déplacer les données pour profiter de ce modèle particulier, et le processus est aussi simple et rentable que d'installer un GPU de milieu de gamme dans le NAS. De plus, à l’heure actuelle, tous ces logiciels sont gratuits, ce qui démocratise grandement le potentiel de l’IA pour les petites organisations.
Chat avec RTX est encore un programme bêta et au moment de la rédaction, nous utilisions la version 0.2. Mais la facilité de l’installer et de faire fonctionner l’interface Web était rafraîchissante. Quiconque sait comment télécharger et installer une application peut désormais obtenir un LLM local avec RAG en quelques clics.
Activation de l'accès à distance pour discuter avec RTX via une URL universellement accessible
Nous avons fait passer notre scénario au niveau supérieur et l'avons rendu disponible pour l'ensemble du bureau.
Étape 1 : localisez le fichier de configuration
Commencez par vous diriger vers le dossier contenant le fichier de configuration :
- Chemin du fichier: C:\Users\{YourUserDir}\AppData\Local\NVIDIA\ChatWithRTX\RAG\trt-llm-rag-windows-main\ui\user_interface.py
Étape 2 : mettre à jour le code de lancement
Ouvrez le user_interface.py fichier et Ctrl-F pour interface.launch Localisez le segment correct, qui apparaîtra par défaut comme suit :
interface.launch(
    favicon_path=os.path.join(os.path.dirname(__file__), 'assets/nvidia_logo.png'),
    show_api=False,
    server_port=port
)
Pour activer l'accès au réseau, vous devez ajouter share=True ainsi:
interface.launch(
    favicon_path=os.path.join(os.path.dirname(__file__), 'assets/nvidia_logo.png'),
    show_api=False,
    share=True,
    server_port=port
)
Enregistrez les modifications dans le user_interface.py déposer. Ensuite, lancez Chat with RTX via le menu Démarrer, qui ouvrira une fenêtre d'invite de commande et activera l'interface.
Étape 3 : Recherche de l'URL publique
La fenêtre d'invite de commande affichera à la fois une URL locale et publique. Pour créer une URL publique fonctionnelle accessible depuis n'importe quel appareil, fusionnez les éléments des deux URL. Il serait préférable que vous preniez l'URL publique et que vous ajoutiez les informations du cookie local à la fin :
- URL publique : https://62e1db9de99021560f.gradio.live
- URL locale avec paramètres : http://127.0.0.1:16852?cookie=4a56dd55-72a1-49c1-a6de-453fc5dba8f3&__theme=dark
Votre URL combinée devrait ressembler à ceci, avec le ?cookie ajouté à l'URL publique :
https://62e1db9de99021560f.gradio.live?cookie=4a56dd55-72a1-49c1-a6de-453fc5dba8f3&__theme=dark
Cette URL permet d'accéder à Chat avec RTX depuis n'importe quel appareil de votre réseau, étendant ainsi sa convivialité au-delà des contraintes locales.
Réflexions finales
Nous sommes fans du leadership de QNAP en matière de conception de matériel NAS depuis longtemps, mais les clients de QNAP ont bien plus de valeur à leur disposition qu'ils ne le pensent probablement. À vrai dire, Virtualization Station est un excellent point de départ, mais pourquoi ne pas passer au niveau supérieur et essayer GPU Passthrough ? À tout le moins, les organisations peuvent fournir une machine virtuelle haut de gamme alimentée par GPU sans avoir à configurer un poste de travail dédié. Il existe également les avantages apparents d’une VM placée à côté d’un énorme pool de stockage interne avec des niveaux de performances natifs. Dans ce cas, nous avions partagé des performances de stockage de près de 10 Go/s, sans nous soucier d’une seule connexion ou d’un seul commutateur 100 GbE, tout cela parce que la VM accélérée par GPU se trouvait à l’intérieur du NAS lui-même.
Pourquoi ne pas aller encore plus loin pour réaliser les avantages de l’IA pour l’organisation ? Nous avons montré qu'ajouter un GPU décent à un NAS QNAP est relativement simple et peu coûteux. Nous avons mis un A4000 au travail, et avec un prix public d'environ 1050 XNUMX $, ce n'est pas mal si l'on considère que Virtualization Station est gratuit et que NVIDIA Chat avec RTX est disponible gratuitement. Être capable d'orienter en toute sécurité ce puissant LLM vers les données privées d'une entreprise devrait fournir des informations exploitables tout en rendant l'entreprise plus dynamique.
