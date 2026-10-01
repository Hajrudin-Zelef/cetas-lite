---
id: collect-261001-rattrapage/rattrapage/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux-2
title: "Git extrait un commit intermédiaire ; vous testez, puis :"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux.md
source_anchor: ""
source_lines: [100, 217]
sha256: 1f427fd665153e64bdf15762483b5c21b0b583d3ade4088d43d4b10097ba2356
---

# Git extrait un commit intermédiaire ; vous testez, puis :

`git stash` stocke temporairement des modifications locales qui ne sont pas prêtes à être commitées. Il permet d’enregistrer vos travaux en cours sans les valider dans le dépôt.

Le stash est utile lorsque vous changez de branche sans vouloir committer ni perdre vos changements. Plus tard, vous pourrez appliquer le stash à votre espace de travail ou le dépiler pour reprendre là où vous vous êtes arrêté.

### Qu’est-ce que git reflog ?

git reflog affiche les journaux de références, qui enregistrent les changements du pointeur HEAD et l’historique des commits consultés dans le dépôt. Il fournit une liste chronologique des actions récentes (commits, checkouts, merges, resets).

Le reflog est précieux pour retrouver des commits ou des branches perdus et comprendre la séquence des actions menées dans le dépôt.

### Comment faire suivre une branche locale par une branche distante existante ?

Pour faire suivre une branche locale par une branche distante, utilisez `git branch` avec l’option `--set-upstream-to` ou `-u`, suivie du nom de la branche distante.

La syntaxe est la suivante :

`git branch --set-upstream-to=<remote-name>/<branch-name>`
ou

`git branch -u <remote-name>/<branch-name>`
## Questions Git avancées

### Comment gérez-vous plusieurs configurations selon les projets dans Git ?

Utilisez `git config` avec les indicateurs `--global`, `--system` ou `--local` pour ajuster les réglages à différents niveaux. Vous pouvez aussi employer `includeIf` dans la configuration Git pour inclure des paramètres spécifiques en fonction du chemin du dépôt.

### Comment gérer les gros fichiers avec Git ?

Les gros fichiers peuvent alourdir le dépôt et dégrader les performances. Utilisez Git LFS pour stocker ces fichiers en dehors du dépôt Git tout en gardant des pointeurs légers dans l’historique. Cela réduit la taille du dépôt et améliore les performances. Git LFS prend en charge divers fournisseurs de stockage et s’intègre naturellement aux workflows Git.

### À quoi sert git submodule et comment en mettre un à jour ?

La commande `git submodule` permet de gérer des dépendances externes au sein d’un dépôt Git. Elle vous autorise à inclure des dépôts externes comme sous-modules dans votre dépôt principal, pratique pour intégrer du code tiers tout en le gardant séparé de votre base de code.

Pour mettre à jour un sous-module :

1. 
Placez-vous dans le répertoire du sous-module dans le dépôt principal.
2. 
Utilisez `git fetch` pour récupérer les derniers changements du dépôt distant du sous-module.
3. 
Pour avancer vers le dernier commit de la branche suivie par le sous-module, utilisez `git pull` .
4. 
Sinon, pour viser un commit ou une branche spécifique, faites `git checkout` avec le hash ou le nom de branche voulu.
5. 
Une fois à l’état désiré, validez dans le dépôt principal pour enregistrer la nouvelle révision du sous-module.

### Qu’est-ce que git cherry-pick et quand l’utiliser ?

`git cherry-pick` applique un commit précis d’une branche sur une autre, sans fusionner l’ensemble de la branche.

`git cherry-pick <commit-hash>``main` mais vous avez aussi besoin du correctif sur une branche `release` : vous pouvez ne récupérer que ce commit plutôt que de fusionner toute la branche `main` dans `release`.
Utile aussi lorsqu’un commit a été fait par erreur sur la mauvaise branche : cherry-pickez-le sur la bonne, puis revertissez-le de celle où il n’a rien à faire.

### Qu’est-ce que git bisect et à quoi sert-il ?

`git bisect` est un outil de débogage qui utilise la recherche binaire pour trouver le commit qui a introduit un bug. Plutôt que de tester des commits un à un, vous indiquez à Git un commit "bon" (sans bug) et un commit "mauvais" (avec bug) ; Git va alors checkout des commits intermédiaires, divisant l’espace de recherche par deux jusqu’à trouver le responsable.

```
git bisect start
git bisect bad                # le commit courant contient le bug
git bisect good <commit-hash> # cet ancien commit était sain
# Git extrait un commit intermédiaire ; vous testez, puis :
git bisect good   # ou git bisect bad
# répétez jusqu'à identification du premier mauvais commit
git bisect reset  # retour à l'état initial
```
C’est bien plus rapide que des tests manuels dans un grand dépôt.

### Que sont les hooks Git et comment les utiliser ?

Les hooks Git sont des scripts exécutés automatiquement à des moments clés du workflow Git. Ils résident dans le répertoire `.git/hooks/` d’un dépôt et peuvent être écrits dans n’importe quel langage de script.

On distingue deux types :

- 
**Côté client** : exécutés en local — par exemple`pre-commit` (avant la création d’un commit) ou`commit-msg` (validation du format du message de commit).
- 
**Côté serveur** : exécutés sur le dépôt distant — par exemple`pre-receive` (avant d’accepter des commits poussés).

Un usage courant est un hook `pre-commit` qui lance automatiquement un linter ou une suite de tests avant d’autoriser un commit, afin d’imposer des standards de qualité.

Notez que les hooks ne sont pas copiés lors d’un clonage ; les équipes les partagent donc via un script dédié ou un outil comme `pre-commit` (le paquet Python).

## Questions sur des concepts Git souvent confondus

### Quelle est la différence entre git fetch et git pull ?

La principale différence entre git fetch et git pull tient à leur effet sur le dépôt local.

`git fetch` récupère les changements d’un dépôt distant et met à jour les branches de suivi à distance (par ex. origin/master) sans modifier votre répertoire de travail ni fusionner quoi que ce soit dans la branche courante. Vous pouvez ainsi examiner les nouveautés sans impacter votre travail.

`git pull` récupère aussi les changements, mais va plus loin : il enchaîne un fetch puis un merge dans votre branche courante, intégrant directement les mises à jour distantes.

### À quoi sert git reset ?

La commande `git reset` repositionne HEAD sur un état donné. Elle permet d’annuler des changements, de retirer des fichiers de l’index ou de déplacer HEAD vers un autre commit. Trois modes principaux existent :

- `--soft` : déplace HEAD vers un commit spécifique en conservant les changements dans l’index. Les fichiers restent modifiés et prêts à être re-commités.

- `--mixed` : déplace HEAD et retire les changements de l’index. Les fichiers restent modifiés dans l’espace de travail, mais ne sont plus en scène.

- `--hard` : déplace HEAD et supprime toutes les modifications dans l’espace de travail et l’index. À utiliser avec précaution : les changements non commités sont définitivement perdus.

**Important :** n’utilisez jamais `git reset --hard` sur des commits déjà poussés sur une branche partagée. Cela réécrit l’historique et posera de sérieux problèmes à vos collègues. Préférez `git revert` pour des commits publics.

### Pourquoi privilégier git push --force-with-lease à git push --force ?

`git push --force-with-lease` est une façon plus prudente de forcer un push que `git push --force` car elle évite d’écraser par inadvertance le travail d’autrui sur le dépôt distant.

Avec `git push --force`, vous forcez la mise à jour sans vérifier si la branche distante a été modifiée depuis votre dernier fetch, ce qui peut effacer le travail d’autres développeurs.

À l’inverse, `git push --force-with-lease` vérifie que la branche distante n’a pas évolué depuis votre dernière récupération. Si c’est le cas, le push est refusé, empêchant l’écrasement involontaire des changements des autres.

### Qu’est-ce que git rebase et en quoi diffère-t-il de git merge ?

git rebase et `git merge` intègrent des changements d’une branche dans une autre, mais de manière différente.

