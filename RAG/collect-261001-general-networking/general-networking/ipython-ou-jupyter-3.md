---
id: collect-261001-general-networking/general-networking/ipython-ou-jupyter-3
title: "Assign the result to `ls`"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [136, 240]
sha256: 113f1b8b91f92be6a4cac64d613e31498e7bf2c8a97ca4fc6e3991c99f57a5c1
---

# Assign the result to `ls`

IPython peut être adapté à un usage de type shell système grâce à l’« escape » shell : les lignes commençant par `!` sont passées directement au shell. Par exemple, !ls exécutera `ls` dans le répertoire courant. Vous pouvez affecter le résultat d’une commande système à une variable Python avec la syntaxe `myfiles=!ls`. Pour afficher explicitement le résultat de ls sous forme de liste de chaînes sans l’assigner, utilisez deux points d’exclamation (`!!ls`) ou la commande magique `%sx` sans assignation.

```
# Assign the result to `ls`
ls = !ls
 
# Explicit `ls`
!!ls
 
# Or with magics
%sx
 
# Assign magics result
ls = %sx
```
Notez que les commandes `!!` ne peuvent pas être assignées à une variable, alors que le résultat d’une magic (si elle renvoie une valeur) peut l’être.

IPython permet aussi d’étendre la valeur des variables Python dans les appels système : il suffit d’englober variables ou expressions entre accolades (`{}`). Dans une commande avec `!` ou `!!`, toute variable Python préfixée par `$` est développée. Dans l’extrait ci-dessous, vous affichez l’attribut argv de la variable sys. Vous pouvez également utiliser les syntaxes `$`/`$$` pour référencer des variables Python depuis la sortie du système, puis vous en servir dans vos scripts.

Pour transmettre un `$` littéral au shell, utilisez un double `$$`. Vous en aurez besoin pour accéder aux variables d’environnement comme $PATH :

```
# Import and initialize
import math
x = 4
 
# System call with variable
!echo {math.factorial(x)}
 
# Expand a variable
!echo $sys.argv
 
# Use $$ for a literal $
!echo "A system variable: $$HOME"
```
En savoir plus ici.

Outre IPython, d’autres kernels proposent aussi des syntaxes spéciales pour exécuter des lignes au shell ! Ce ne sont pas forcément des « magics » au sens strict, car le nom et la richesse peuvent varier selon l’implémentation.

Vous pouvez également définir des **alias** pour des commandes système : ce sont des raccourcis vers des commandes bash. Un alias est un tuple : (« showTheDirectory », « ls »). Lancez `%alias?` pour plus d’infos !

**Astuce** : utilisez `%rehashx` pour charger tout votre $PATH comme alias IPython.

### Magics ?

Si vous avez parcouru le guide ultime de Jupyter Notebook de DataCamp ou si vous avez déjà utilisé Jupyter, vous connaissez sans doute les « magic commands ». Les magics reposent sur un élément de syntaxe invalide dans le langage sous-jacent et un mot qui évoque une commande. Sous le capot, ce sont des fonctions Python.

Le kernel IPython utilise le symbole % (qui n’est pas unitaire valide en Python). Les lignes commençant par %% marquent une « cell magic » : elles prennent en arguments non seulement le reste de la ligne, mais aussi toutes les lignes suivantes du bloc courant. Les cell magics peuvent modifier librement l’entrée reçue, qui n’a même pas besoin d’être du Python valide. Elles reçoivent le bloc entier comme une unique chaîne.

Les magics sont spécifiques aux kernels qui les fournissent et sont conçues pour rendre votre travail dans Jupyter Notebook plus interactif. Leur disponibilité dépend des développeurs du kernel et varie d’un kernel à l’autre. Vous l’aurez compris : les magics sont une fonctionnalité de kernel.

Quand vous utilisez le backend Python de Jupyter, IPython (le kernel), voici quelques astuces pour accéder à des fonctionnalités qui accélèrent, simplifient et enrichissent votre programmation. La liste n’est pas exhaustive. Consultez la liste complète des magics intégrées.

#### Visualisation

Une fonctionnalité majeure du kernel IPython est l’affichage des graphiques produits par les cellules de code. Il est conçu pour fonctionner de manière fluide avec la bibliothèque matplotlib. Pour l’activer, utilisez la magic `%matplotlib`.

Par défaut, le graphique s’ouvre dans une fenêtre séparée. Vous pouvez aussi préciser un backend (inline, qt, etc.) pour afficher les tracés en ligne ou via une autre interface GUI. Plus d’infos ici.

#### Navigation dans le système de fichiers

Les magics du kernel offrent aussi des moyens de naviguer dans votre système de fichiers. Les magics `%cd` et `%bookmark` permettent de changer de répertoire ou d’ajouter des favoris pour accéder plus vite aux dossiers fréquents.

#### Accès au débogueur

Vous pouvez invoquer un débogueur Python avec `%pdb` à chaque exception non interceptée. Il vous guide dans la portion de code à l’origine de l’erreur pour accélérer l’analyse.

La magic `%run` avec l’option -d exécute un script sous contrôle du débogueur Python et place automatiquement des points d’arrêt initiaux. La magic `%debug` offre un accès encore plus direct.

#### Extensions IPython

La magic `%load_ext` charge une extension IPython par son nom de module. Les extensions IPython sont des modules Python qui modifient le comportement du shell : elles peuvent enregistrer des magics, définir des variables et, plus généralement, enrichir l’espace utilisateur pour offrir de nouvelles fonctions dans les cellules. Exemples :

- `%load_ext oct2py.ipython` : appeler de façon transparente des M-files et fonctions Octave depuis Python,
- `%load_ext rpy2.ipython` : interface vers R embarqué dans un processus Python,
- `%load_ext Cython` : utiliser le compilateur Python→C,
- `sympy.init_printing()` : jolis rendus automatiques des objets Sympy Basic,
- `%load_ext fortranmagic` : utiliser Fortran dans vos sessions interactives.

… Et bien d’autres ! Vous pouvez créer et publier vos propres extensions IPython sur PyPI : il existe donc beaucoup d’extensions et de magics communautaires. Exemple : `ipython_unittest`, et consultez aussi cet index des extensions.

Gardez aussi un œil sur sparkmagic, un ensemble d’outils pour travailler de façon interactive avec des clusters Spark distants via Livy (serveur REST Spark) dans des notebooks Jupyter. La bibliothèque `sparkmagic` fournit une magic `%%spark` pour exécuter facilement du code contre un cluster Spark distant depuis un notebook IPython classique.

```
# Load in sparkmagic
%load_ext sparkmagic.magics
 
# Set the endpoint
%manage_spark
 
# Ask for help
%spark?
```
Rendez-vous ici pour plus d’exemples d’utilisation de ces magics avec un cluster Spark.

Notez qu’IPython propose deux autres magics, `%reload_ext` et `%unload_ext`, pour recharger et décharger des extensions directement depuis Jupyter Notebook.

#### Différents kernels, autres magics

Dans d’autres langages, l’élément de syntaxe utilisé par les magics peut déjà avoir un sens. Le kernel R, IRkernel, n’a pas de système de magics. Pour exécuter des commandes bash, vous utiliserez par exemple des fonctions R comme `system()` pour invoquer des commandes OS, par exemple `system("head -5 *.csv", intern=TRUE)`. En incluant l’argument `intern`, vous indiquez que vous souhaitez capturer la sortie de la commande en vecteur de caractères R. Pour afficher du Markdown, utilisez `display_markdown()`, à qui vous passez le code Markdown en vecteur de caractères. De même, le kernel Julia IJulia n’utilise pas les « magics ». D’autres syntaxes, plus naturelles en Julia, fonctionnent en dehors des cellules IJulia et sont souvent plus puissantes. Les développeurs d’IJulia ont toutefois prévu qu’en entrant une magic IPython dans une cellule IJulia, un message d’aide indique la manière d’obtenir un effet équivalent en Julia, lorsque c’est possible.

Par exemple, l’équivalent de `%load` d’IPython en IJulia est `IJulia.load()`.

À l’inverse, certains kernels comme IScala (Scala) prennent en charge des magics, de façon similaire à IPython, mais avec un jeu adapté à Scala et à la JVM. Une magic commence par `%` suivi d’un identifiant et, éventuellement, d’une saisie. Exemples notables :

