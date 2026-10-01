---
id: collect-261001-general-networking/general-networking/ipython-ou-jupyter-1
title: "Assign the result to `ls`"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [1, 66]
sha256: 9fbd32983c13a1fc2513d76969eea1313302ce98ff829575e1b5ae408ca186ef
---

# Assign the result to `ls`

*Un grand merci à Brian Granger, Fernando Pérez et Robert Kern pour leurs contributions à cet article !*

Pour les apprenants comme pour les data scientists confirmés, Jupyter Notebook compte parmi les outils phares de la data science : son environnement interactif est idéal pour enseigner, apprendre, partager ses travaux avec ses pairs et garantir la reproductibilité des recherches. Pourtant, en découvrant le notebook, vous croiserez très souvent IPython.

Dans certains cas, les deux semblent synonymes et vous conviendrez que cela devient vite déroutant dès que l’on veut creuser : les « magics » relèvent-ils de Jupyter ou d’IPython ? L’enregistrement et le chargement des notebooks sont-ils des fonctionnalités d’IPython ou de Jupyter ?

La liste des questions peut être longue.

L’objectif de cet article est d’exposer clairement quelques différences fondamentales entre les deux, en repartant de leurs origines pour expliquer leurs liens, puis en passant en revue des fonctionnalités propres à l’un ou l’autre, afin de vous aider à mieux faire la distinction.

Pensez aussi à consulter le guide ultime de Jupyter Notebook de DataCamp pour des astuces, bonnes pratiques, exemples et bien plus encore.

## Les origines d’IPython / Jupyter

Pour bien comprendre ce qu’est Jupyter Notebook et en quoi il se distingue d’IPython, il est utile de revenir sur la place de ces deux projets dans l’histoire (et l’avenir) des notebooks computationnels.

### Les débuts des notebooks computationnels : MATLAB, Mathematica et Maple

Au milieu des années 1980, MATLAB est lancé par The MathWorks, fondé par Jack Little, Steve Bangert et Cleve Moler.

Fin des années 1980, 1987 pour être précis, Theodore Gray commence à travailler sur ce qui deviendra l’interface notebook de Mathematica, publiée un an plus tard. Cette interface graphique permet de créer et modifier de manière interactive des documents notebook contenant du code joliment mis en forme, du texte, et de nombreuses fonctionnalités comme des formules composées, des graphiques, des composants GUI, des tableaux et des sons. On y trouve les fonctions classiques d’un traitement de texte, dont la correction orthographique multilingue en temps réel. Les documents peuvent être diffusés en mode diaporama pour des présentations.

La structure de ces notebooks reposait sur une hiérarchie de cellules facilitant plan et sections des documents — un principe que l’on retrouve aujourd’hui dans les notebooks Jupyter.

En 1989, Maple introduit également sa première interface de type notebook, incluse avec la version 4.3 pour Macintosh. Des versions pour X11 et Windows suivent en 1990. Ces notebooks précurseurs ont inspiré et posé les bases de ce que l’on appellera plus tard les « notebooks de data science ».

### L’essor des notebooks de data science

De nombreux notebooks computationnels ont vu le jour après Maple et Mathematica. Cette section se concentre toutefois sur ceux qui ont contribué à l’essor des notebooks de data science, dont certains restent très populaires auprès des data scientists, débutants comme confirmés.

#### Sage Notebook

Le notebook Sage, en tant que système accessible via navigateur, paraît au milieu des années 2000. En 2007, une nouvelle version plus puissante voit le jour, avec gestion des comptes utilisateurs et possibilité de publier des documents. Son interface rappelle celle de Google Docs, car la mise en page du notebook Sage s’inspire de Google Notebooks.

Les créateurs de Sage ont confirmé avoir été des utilisateurs assidus des notebooks Mathematica et des feuilles Maple. D’autres facteurs ont compté dans le développement de Sage : d’une part, des échanges étroits avec l’équipe d’IPython, Sage en mode terminal s’appuyant sur IPython ; d’autre part, « une tentative avortée » de deux étudiants pour proposer une interface graphique à IPython ; enfin, la montée des applications web « AJAX » (Asynchronous JavaScript And XML), qui évitent de recharger toute la page à chaque action.

Par exemple, lorsque vous soumettez un formulaire sur un site : sans AJAX, vous êtes redirigé vers une nouvelle page renvoyée par le serveur. Avec AJAX, JavaScript envoie la requête, récupère la réponse et met à jour l’écran, sans rechargement ni redirection.

Parmi les applications AJAX connues : Gmail, Google Maps, Facebook, Twitter, etc. Presque tout repose désormais sur AJAX.

#### IPython et Jupyter Notebook

Fin 2001, environ vingt ans après que Guido van Rossum a commencé à travailler sur Python au CWI (Pays-Bas), Fernando Pérez démarre le développement d’IPython. Comme Sage et bien d’autres projets, IPython est fortement influencé par les notebooks Mathematica et les feuilles Maple.

En 2005, une première tentative de notebook voit le jour avec Wx, une boîte à outils de widgets pour créer des interfaces graphiques multiplateformes. Deux étudiants du Google Summer of Code travaillent sur un prototype sous la direction de Robert Kern et Fernando Pérez. Robert pousse ensuite plus loin le chantier. Il s’agit moins d’un prototype de notebook que d’un nettoyage du cœur d’IPython pour faciliter l’écriture d’un frontal wxPython au shell IPython. Ces refontes ont contribué à rendre possible un notebook propre, et certaines ont été intégrées progressivement à IPython lorsque l’effort, cette fois réussi, s’est concrétisé.

Le deuxième prototype d’un notebook IPython est réalisé à l’été 2006 par Min Ragan-Kelley, conseillé par Brian Granger. Basé sur le web avec un backend SQL, il est finalement abandonné car trop complexe pour les technologies web de l’époque.

Le troisième prototype arrive en octobre 2010, développé par un tiers lors d’un hackathon de quelques jours. Enfin, au printemps-été 2011, Brian Granger travaille à temps plein sur un prototype de notebook web, s’appuyant sur ses travaux de 2010, lorsqu’il a créé avec Fernando Pérez l’architecture du kernel IPython et la spécification des messages, ainsi que PyZMQ avec Min Ragan Kelley.

PyZMQ est une bibliothèque qui fournit des bindings Python pour ZeroMQ : nécessaire pour les fonctionnalités de calcul parallèle, la console Qt et le notebook d’IPython.

PyZMQ et les websockets sont les technologies clés qui rendent le notebook possible. À l’automne, d’autres contributeurs (dont Matthias Bussonnier, Min et Fernando) rejoignent l’effort. Le 21 décembre 2011 sort la première version d’IPython Notebook (0.12).

Les années suivantes, l’équipe reçoit des prix (par exemple, l’Advancement of Free Software pour Fernando Pérez le 23 mars 2013, le Jolt Productivity Award) et des financements (Fondation Alfred P. Sloan, entre autres).

En 2014, naît finalement Project Jupyter, issu d’IPython.

La dernière version d’IPython avant la scission regroupait dans un seul dépôt le shell interactif, le serveur de notebook, la console Qt, etc. Le projet devenait vaste, avec des composants de plus en plus distincts.

Mais la taille n’est pas la seule raison de la création de Jupyter : de 2011 à 2014, l’IPython Notebook commence à fonctionner avec d’autres langages. Julia arrive en premier, puis R. Le fait que l’« IPython » Notebook — suggérant un notebook exclusivement Python — fonctionne aussi avec des kernels pour Julia et R déroute.

Pourquoi garder « IPython » si d’autres langages sont pris en charge ?

Le nom « Jupyter » s’inspire des principaux langages ouverts pour la science (Julia, Python et R). Même s’il peut sembler acronyme, il n’a jamais signifié que les autres langages n’étaient pas les bienvenus. Surtout, il représentait mieux le projet et faisait clin d’œil à ses racines scientifiques.

