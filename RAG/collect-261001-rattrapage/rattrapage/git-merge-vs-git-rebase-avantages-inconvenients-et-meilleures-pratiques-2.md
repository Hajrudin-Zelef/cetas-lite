---
id: collect-261001-rattrapage/rattrapage/git-merge-vs-git-rebase-avantages-inconvenients-et-meilleures-pratiques-2
title: "On main, bring in feature-xyz:"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/git-merge-vs-git-rebase-avantages-inconvenients-et-meilleures-pratiques.md
source_anchor: ""
source_lines: [102, 171]
sha256: 7e6accfc72db8a57fdab54fc0cd1bc0ab955f636ee0a1bbbac8a14260b885dac
---

# On main, bring in feature-xyz:

Les équipes axées sur le rebasage adoptent des pratiques de rebasage continu où les développeurs mettent régulièrement à jour leurs branches de fonctionnalités avec les dernières modifications de la branche principale. Au lieu d'attendre l'achèvement des fonctionnalités, les membres de l'équipe refont leur travail plusieurs fois par jour pour rester au courant des développements en cours.

Ces équipes mettent fortement l'accent sur l' écrasement des engagements et sur la conservation de l'historique de . Avant de fusionner tout travail, les développeurs combinent les livraisons connexes, réécrivent les messages de livraison pour plus de clarté et s'assurent que leur branche raconte une histoire cohérente du développement de la fonctionnalité.

Le projet du noyau Linux est lee parfait exemple d'une utilisation efficace de la base de données à grande échelle. Avec des milliers de contributeurs dans le monde entier, le projet maintient un historique des livraisons remarquablement propre grâce à des normes strictes de rebasage. Linus Torvalds lui-même préconise le rebasage des branches de fonctionnalités avant leur soumission, arguant du fait qu'un historique conservé rend le débogage et la révision du code nettement plus efficaces.

Les équipes orientées vers la rebase font souvent état de cycles de révision du code plus rapides, car les réviseurs peuvent suivre une progression logique des changements plutôt que de déchiffrer le chaos chronologique du développement collaboratif.

**> Êtes-vous un débutant cherchant à garder vos dépôts t***idy ? La commande Git clean est tout ce dont vous avez besoin.*

## Apprenez les bases de Git dès aujourd'hui

## Dynamique de résolution des conflits

Lorsque les branches divergent et modifient le même code, les conflits deviennent inévitables, quelle que soit votre stratégie d'intégration. Cela dit, la fusion et la refonte gèrent ces conflits de différentes manières qui affectent à la fois votre flux de travail immédiat et la maintenance à long terme du projet.

### Fusionner les caractéristiques du conflit

Git merge présente tous les conflits en une seule session de résolution. Lorsque vous lancez `git merge feature-branch`, Git identifie chaque fichier conflictuel et marque toutes les sections problématiques en même temps, ce qui vous permet de voir l'étendue complète des défis d'intégration en amont.

Cette approche permet de préserver le contexte lors de la résolution des conflits. Yous pouvez voir exactement quelles modifications proviennent de quelle branche, quand elles ont été effectuées et qui en est l'auteur. Le commit de fusion qui résulte de votre résolution de conflit devient un enregistrement permanent de la façon dont vous réconciliez les changements concurrents.

La fusion des conflits permet de conserver un suivi clair du processus décisionnel. Si, des mois plus tard, vous devez réexaminer les raisons pour lesquelles un certain code a été préféré à d'autres, le commit de fusion montre à la fois les versions conflictuelles d'origine et votre résolution finale. Cette piste d'audit s'avère inestimable pour le débogage et la compréhension de l'évolution du projet.

Toutefois, les fusions complexes comportant de nombreux conflits peuvent devenir accablantes. Vous pouvez être amené à résoudre des dizaines de fichiers conflictuels au cours d'une seule session, ce qui vous permet de passer à côté de problèmes d'intégration subtils ou d'introduire de nouveaux bogues au cours du processus de résolution.

```
git checkout main
git merge feature-xyz
# ← conflicts in file foo.js
# Resolve in your editor, then:
git add foo.js
git commit
```
### Rebase caractéristiques du conflit

Git rebase vous oblige à résoudre les conflits de manière itérative, un commit à la fois, en reproduisant l'historique de votre branche. Lorsque des conflits surviennent pendant le rebasement, Git s'arrête à chaque livraison problématique et vous demande de résoudre les conflits avant de passer à la livraison suivante.

Cette approche granulaire permet de décomposer les problèmes d'intégration complexes en éléments gérables. Au lieu de faire face à tous les conflits simultanément, vous les traitez dans l'ordre logique dans lequel ils ont été créés, ce qui rend souvent le processus de résolution plus intuitif et moins sujet aux erreurs.

La nature itérative permet de maintenir l'intégrité historique en veillant à ce que chaque engagement individuel reste cohérent et fonctionnel. Comme vous résolvez les conflits dans le contexte de changements spécifiques, vous avez moins de chances de mélanger accidentellement des modifications sans rapport ou de créer des modifications qui cassent la compilation.

Cependant, la résolution des conflits par rebase peut devenir fastidieuse pour les branches de fonctionnalités de longue durée. Vous pouvez rencontrer le même conflit plusieurs fois si des modifications similaires ont été effectuées sur plusieurs commits, ce qui nécessite la résolution répétée de problèmes essentiellement identiques.

```
git checkout feature-xyz
git rebase main
# ← stops at first bad commit:
# Resolve foo.js, then:
git add foo.js
git rebase --continue
```
## Préservation de l'histoire et préservation de l'environnement Lisibilité

La tension fondamentale entre la fusion et la refonte est centrée sur la question de savoir si vous donnez la priorité à des enregistrements historiques authentiques ou à des récits de projet propres et lisibles. Ce choix a une incidence sur la manière dont les futurs développeurs comprendront votre base de code et résoudront les problèmes au fil du temps.

### Fusionner la fidélité historique

La stratégie de fusion préserve l'authenticité temporelle en maintenant l'ordre chronologique exact des activités de développement. Vous pouvez voir quand les développeurs ont réellement écrit du code, quand ils ont apporté des modifications et comment les différentes fonctionnalités ont évolué en parallèle plutôt que de manière isolée.

Cette approche permet de saisir un contexte de collaboration précieux qui se perd dans d'autres méthodes d'intégration. Vous verrez le va-et-vient des révisions de code, les validations expérimentales qui ont conduit à des percées, et les faux départs qui ont finalement permis de trouver de meilleures solutions. La réalité désordonnée du développement de logiciels fait partie de votre dossier permanent.

Les commits de fusion servent de timestamps qui marquent le moment où des fonctionnalités spécifiques ont rejoint la base de code principale. Si vous avez besoin de comprendre à quoi ressemblait l'application à un moment donné, l'historique des fusions vous donne des points d'intégration précis plutôt que des séquences artificiellement construites.

Cependant, cette fidélité se fait au détriment de la cohérence narrative. Votre graphique de livraison devient un réseau complexe de branches entrelacées qui peut submerger les développeurs essayant de comprendre le développement des fonctionnalités. La chronologie authentique masque souvent la progression logique des idées, ce qui rend plus difficile le suivi du raisonnement qui sous-tend les changements majeurs.

### Clarté narrative du rebasement

Rebase construit un historique qui présente le développement des fonctionnalités comme une séquence logique de changements intentionnels. Au lieu de montrer la réalité désordonnée du développement, l'histoire remaniée raconte ce qui aurait dû se passer si les promoteurs avaient été parfaitement prévoyants.

