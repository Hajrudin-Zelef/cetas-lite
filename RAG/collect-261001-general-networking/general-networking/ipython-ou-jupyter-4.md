---
id: collect-261001-general-networking/general-networking/ipython-ou-jupyter-4
title: "Assign the result to `ls`"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [241, 326]
sha256: 133cff9cf9ad9478bf6f46c1c16cc7bcbe1c7dcc76ed26f23a2800be86d8f6c3
---

# Assign the result to `ls`

```
# Type Information
%type 1
 
# Library Management
%libraryDependencies
%update
```
Comme indiqué plus haut, la bibliothèque `sparkmagic` fournit aussi des kernels Scala et Python pour se connecter automatiquement à un cluster Spark distant, exécuter du code et des requêtes SQL, gérer la configuration de votre serveur Livy et de vos jobs Spark, et générer des visualisations automatiques — sans écrire de code supplémentaire.

Par exemple, vous pouvez lancer des requêtes SparkSQL avec `%%sql` ou accéder aux informations et journaux d’application Spark via la magic `%%info`.

Si vous utilisez un autre kernel et vous demandez s’il supporte les magics, sachez que certains kernels s’appuient sur le projet metakernel et reprennent, dans bien des cas, les mêmes magics que le kernel IPython. La liste des magics metakernel est ici. metakernel est un modèle de kernel Jupyter/IPython incluant des magics de base.

Exemples :

- Le kernel MATLAB `matlab_kernel` ,
- Le kernel Octave `octave_kernel` ,
- Le kernel Java9 `java9_kernel` ,
- Le kernel Wolfram `wolfram_kernel` ,
- Le kernel SAS. … Et bien d’autres !

Ainsi, avec le kernel MATLAB, vous disposez par exemple des magics suivantes :

```
Available line magics:
%cd  %connect_info  %download  %edit  %get  %help  %html  %install  %install_magic  %javascript  %kernel  %kx  %latex  %load  %ls  %lsmagic  %magic  %parallel  %plot  %pmap  %px  %python  %reload_magics  %restart  %run  %set  %shell  %spell
 
Available cell magics:
%%debug  %%file  %%help  %%html  %%javascript  %%kx  %%latex  %%processing  %%px  %%python  %%shell  %%show  %%spell
```
Si vous comparez avec les magics disponibles par défaut dans le kernel IPython, vous noterez des recoupements :

```
Available line magics:
%alias  %alias_magic  %autocall  %automagic  %autosave  %bookmark  %cat  %cd  %clear  %colors  %config  %connect_info  %cp  %debug  %dhist  %dirs  %doctest_mode  %ed  %edit  %env  %gui  %hist  %history  %killbgscripts  %ldir  %less  %lf  %lk  %ll  %load  %load_ext  %loadpy  %logoff  %logon  %logstart  %logstate  %logstop  %ls  %lsmagic  %lx  %macro  %magic  %man  %matplotlib  %mkdir  %more  %mv  %notebook  %page  %pastebin  %pdb  %pdef  %pdoc  %pfile  %pinfo  %pinfo2  %popd  %pprint  %precision  %profile  %prun  %psearch  %psource  %pushd  %pwd  %pycat  %pylab  %qtconsole  %quickref  %recall  %rehashx  %reload_ext  %rep  %rerun  %reset  %reset_selective  %rm  %rmdir  %run  %save  %sc  %set_env  %store  %sx  %system  %tb  %time  %timeit  %unalias  %unload_ext  %who  %who_ls  %whos  %xdel  %xmode
 
Available cell magics:
%%!  %%HTML  %%SVG  %%bash  %%capture  %%debug  %%file  %%html  %%javascript  %%js  %%latex  %%perl  %%prun  %%pypy  %%python  %%python2  %%python3  %%ruby  %%script  %%sh  %%svg  %%sx  %%system  %%time  %%timeit  %%writefile
```
Pour discerner les magics spécifiques à IPython de celles exploitables ailleurs, posez-vous une question simple : la fonctionnalité est-elle propre à Python ou générale au langage que vous utilisez ?

Par exemple, `%pdb` (débogueur Python) ou `%matplotlib` sont spécifiques à Python et n’ont pas de sens avec un kernel JavaScript. En revanche, changer de répertoire avec `%cd` est assez général et devrait fonctionner dans tout environnement, sous réserve que le kernel gère les magics.

### Conversion et formatage des notebooks ?

La conversion et le formatage des notebooks relèvent de l’écosystème Jupyter. Deux outils typiques : `nbconvert` et `nbformat`.

Le premier convertit les notebooks vers d’autres formats pour présenter l’information sous des formes familières, publier des travaux, intégrer des notebooks dans des articles, collaborer et partager plus largement.

Le second définit le format des notebooks Jupyter : des documents JSON simples contenant des métadonnées (kernel, infos de langage), la version du format (majeure et mineure) et les cellules où sont stockés texte, code, etc.

### Enregistrement et chargement des notebooks ?

L’enregistrement et le chargement des notebooks sont des fonctionnalités de l’application Jupyter Notebook. Vous pouvez charger des notebooks (fichiers .ipynb) créés par d’autres en les téléchargeant puis en les ouvrant dans l’application Jupyter. Concrètement, créez un nouveau notebook, puis ouvrez le fichier via l’onglet « File », « Open », et sélectionnez votre notebook téléchargé.

Inversement, enregistrez vos propres notebooks via le même onglet « File » en choisissant « Download as » pour récupérer le fichier, ou « Save and Checkpoint » pour poser un jalon. Très utile pour une forme légère de contrôle de versions et revenir à un état antérieur si besoin. Bien sûr, vos modifications sont sauvegardées automatiquement toutes les quelques minutes, donc l’action n’est pas toujours nécessaire.

Notez que vous pouvez aussi préserver un notebook original en travaillant sur une copie et en y enregistrant vos changements !

### Raccourcis clavier et multi‑curseur ?

Sélection de cellules multiples, affichage/masquage des sorties, insertion de nouvelles cellules, etc. : autant d’actions disposant de raccourcis clavier dans Jupyter Notebook. La liste est accessible via le menu du haut : onglet « Help », puis « Keyboard Shortcuts ».

Le support multi‑curseur est également une fonctionnalité de Jupyter Notebook !

### Calcul parallèle distribué ?

Le réseau de calcul parallèle faisait partie d’IPython, mais depuis la version 4.0, il est scindé en un package autonome, `ipyparallel`. Ce package regroupe des scripts CLI pour contrôler des clusters Jupyter.

Bien que séparé, il demeure un composant puissant de l’écosystème IPython, souvent sous-estimé : au lieu d’un seul kernel Python, vous pouvez démarrer de nombreux kernels distribués sur plusieurs machines.

Cas d’usage typiques : exécuter un modèle de très nombreuses fois pour estimer la distribution de ses sorties ou leur sensibilité aux paramètres d’entrée. Lorsque les runs sont indépendants, on gagne du temps en les parallélisant sur plusieurs machines d’un cluster. Pensez à l’entraînement distribué de modèles ou à des simulations.

### Terminal ?

Cette fonctionnalité appartient à l’écosystème Jupyter : il existe Jupyter Console et une application de terminal Jupyter. Depuis l’origine, IPython désigne toutefois le terminal interactif initial pour Python, offrant une boucle REPL enrichie, particulièrement adaptée au calcul scientifique. C’était la référence avant 2011, date à laquelle le Notebook a introduit une interface web moderne et puissante pour Python.

On disposait aussi d’une console IPython, qui lançait deux processus : le shell terminal IPython et le profil/kernel par défaut (Python si non précisé). La console IPython est désormais dépréciée et, pour retrouver un usage similaire, il faut utiliser Jupyter Console, un frontal en terminal pour les kernels Jupyter. Il reprend l’expérience interactive d’IPython au terminal, mais permet de se connecter à n’importe quel kernel Jupyter, pas seulement à IPython. Vous pouvez ainsi tester tout kernel Jupyter installé, sans ouvrir un Notebook complet. La Console permet des interactions en terminal avec IJulia, IRKernel, etc.

Enfin, l’application Jupyter Notebook inclut aussi une application Terminal : un shell bash simple qui s’exécute dans votre navigateur. Vous le trouverez facilement en démarrant l’application puis en créant un nouveau terminal via le menu déroulant.

### Console Qt ?

La console Qt faisait partie d’IPython, mais elle a rejoint le projet Jupyter. C’est une application légère qui ressemble à un terminal tout en offrant des améliorations propres à une GUI : figures en ligne, véritable édition multilignes avec coloration syntaxique, aides contextuelles graphiques, etc. La console Qt peut utiliser n’importe quel kernel Jupyter.

