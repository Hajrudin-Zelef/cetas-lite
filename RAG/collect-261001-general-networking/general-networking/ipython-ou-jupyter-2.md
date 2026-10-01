---
id: collect-261001-general-networking/general-networking/ipython-ou-jupyter-2
title: "Assign the result to `ls`"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["apache", "attention", "exploit"]
source: docs/RAG/collect-261001-general-networking/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [67, 135]
sha256: 4ce89f7d5bbbe04749c2246498c4266a24c90d21f9ea8d115442316535d3d48d
---

# Assign the result to `ls`

Après la création de Jupyter, les éléments agnostiques au langage d’IPython — format de notebook, protocole de messages, console Qt, application web du notebook, etc. — sont transférés dans le projet Jupyter. L’organisation GitHub principale de Jupyter se trouve ici.

Dans les communautés Jupyter et IPython, on parle de « The Big Split ».

IPython n’a plus que deux rôles : servir de backend Python au notebook Jupyter (le kernel) et fournir un shell Python interactif. Mais ce n’est pas tout : l’écosystème IPython inclut aussi un framework de calcul parallèle, dont nous parlerons plus loin.

Comme IPython, Project Jupyter est un nom-parapluie pour plusieurs projets : les trois applications principales sont le Notebook, une Console et une console Qt, mais on trouve aussi des sous-projets comme JupyterHub pour le déploiement des notebooks, nbgrader pour l’enseignement, etc. Un aperçu de l’architecture Jupyter est disponible ici.

Cette évolution explique la confusion de nombreux pythonistes face à IPython et Jupyter : l’un est issu de l’autre, et récemment. Certains peinent encore à employer les bons termes. Plus déroutant encore, l’évolution elle-même : l’héritage commun entraîne un chevauchement notable des fonctionnalités d’IPython et de Jupyter Notebook, parfois difficile à démêler. Les sections suivantes éclaireront ces distinctions. Pour en savoir plus sur l’histoire d’IPython, lisez les témoignages de Fernando Pérez et de William Stein sur leurs notebooks.

#### Notebooks R

R Markdown et Jupyter Notebook partagent l’objectif d’un flux de travail reproductible, en tissant code, résultats et texte dans un même document, avec support de widgets interactifs et export multi-formats.

Mais ils diffèrent aussi : R Markdown privilégie l’exécution par lot reproductible, la représentation en texte brut, le contrôle de version, des sorties de production et l’usage du même éditeur/outillage que pour les scripts R. Les notebooks, eux, affichent le résultat au fil du code, mettent en cache les sorties entre sessions, et facilitent le partage du code et des résultats dans un fichier unique. Ils mettent l’accent sur l’exécution interactive. Ils n’utilisent pas une représentation en texte brut, mais un format de données structuré, comme JSON.

C’est ce qui motive l’application « notebook » de RStudio : elle marie les atouts de R Markdown et ceux des notebooks computationnels.

Pour apprendre à travailler avec les notebooks R et comprendre précisément les différences entre Jupyter et les notebooks R Markdown en matière de partage, gestion de projet, contrôle de version, etc., lisez l’article de DataCamp Jupyter et R : notebooks avec R.

#### Autres notebooks de data science

D’autres notebooks méritent votre attention en data science. Ces dernières années, de nombreuses alternatives ont émergé : Beaker Notebook, Apache Zeppelin, Spark Notebook, DataBricks Cloud, etc., mais aussi des outils comme l’IDE Rodeo ou nteract, qui rendent vos analyses interactives et reproductibles. Notez que nteract se distingue en s’appuyant sur l’architecture Jupyter (protocoles et formats).

### L’avenir des notebooks

Les notebooks sont là pour durer. La nouvelle génération de Jupyter Notebook a récemment été introduite : JupyterLab. L’application Notebook intègre non seulement le support des notebooks, mais aussi un gestionnaire de fichiers, un éditeur de texte, un terminal, un moniteur des processus Jupyter, un gestionnaire de clusters IPython et un pager d’aide.

Rien de neuf ? En réalité, JupyterLab permet d’exploiter tous ces blocs de construction du calcul interactif de manières inédites.

À lire ici.

L’arsenal d’outils de Jupyter Notebook a grandi organiquement, guidé par les besoins des utilisateurs et des développeurs. JupyterLab apporte une architecture de nouvelle génération pour les réunir, avec une interface flexible et réactive, offrant une mise en page pilotée par l’utilisateur.

## IPython ou Jupyter ?

L’évolution du projet et la « Big Split » qui s’ensuit posent les bases pour comprendre les vraies différences entre les deux. Mais comme ils sont intimement liés, le doute persiste parfois sur ce qui appartient à quoi.

La section suivante passe en revue des fonctionnalités relevant soit de l’écosystème IPython, soit du projet Jupyter.

À vous d’identifier la bonne réponse et d’en apprendre plus sur chaque fonction !

### Kernels ?

Même si les kernels apparaissent en bonne place dans l’application Jupyter Notebook, le premier outil à utiliser un kernel et un protocole complets fut la console Qt, antérieure au Notebook. Aujourd’hui, les kernels sont exploités de façon souple par le Notebook, historiquement comme architecturalement, et relèvent de Jupyter et non du seul Notebook. Autrement dit, les kernels sont bien plus qu’une fonctionnalité : ils constituent une abstraction centrale de l’architecture Jupyter, utilisée par des outils hors-notebook comme la console texte, la console Qt, Thebe d’O’Reilly, Kernel Gateway, l’éditeur Hydrogen de nteract. Enfin, dans JupyterLab, un kernel peut aussi se connecter à tout, d’un notebook à une console web, voire un simple fichier texte.

Un kernel est un programme qui exécute et introspecte le code de l’utilisateur : il assure le calcul et la communication avec les interfaces front-end comme les notebooks. L’application Jupyter Notebook propose trois kernels principaux : IPython, IRkernel et IJulia.

Étant donné que « Jupyter » s’inspire de Julia, Python et R, rien d’étonnant. Le kernel IPython est maintenu par l’équipe Jupyter, conséquence logique de l’évolution du projet.

Vous pouvez cependant faire tourner bien d’autres langages dans l’application Jupyter Notebook : Scala, JavaScript, Haskell, Ruby, etc. Il s’agit de kernels maintenus par la communauté.

### Déploiement de notebooks ?

Le déploiement des notebooks est typiquement un sujet que vous rencontrerez avec Jupyter. Plusieurs packages de l’écosystème Jupyter vous y aident.

En voici quelques-uns :

- `docker-stacks` : des piles d’applications Jupyter et de kernels sous forme de conteneurs Docker.
- `ipywidgets` : des widgets HTML & JavaScript interactifs (curseurs, cases à cocher, champs texte, graphiques…) pour l’architecture Jupyter, reliant contrôles front-end et kernel.
- `jupyter-drive` : permet à IPython d’utiliser Google Drive pour la gestion des fichiers.
- `jupyter-sphinx-theme` : ajoute un thème Sphinx Jupyter à vos notebooks pour créer une documentation soignée et intelligente.
- `kernel_gateway` : un serveur web qui prend en charge différents mécanismes de lancement et de communication avec des kernels Jupyter. Des cas d’usage ici.
- `nbviewer` : pour partager vos notebooks. Galerie ici.
- `tmpnb` : crée des serveurs Jupyter Notebook temporaires via Docker. À essayer ici.
- `traitlets` : un framework qui dote les classes Python d’attributs avec vérification de type, valeurs par défaut dynamiques et callbacks « on change ». Utile aussi pour la configuration (fichiers, arguments de ligne de commande). traitlets alimente le système de configuration d’IPython et de Jupyter ainsi que l’API déclarative des widgets interactifs IPython.

### Utilisation du shell système ?

