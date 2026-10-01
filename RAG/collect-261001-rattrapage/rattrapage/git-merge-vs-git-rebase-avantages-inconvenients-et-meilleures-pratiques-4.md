---
id: collect-261001-rattrapage/rattrapage/git-merge-vs-git-rebase-avantages-inconvenients-et-meilleures-pratiques-4
title: "On main, bring in feature-xyz:"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/git-merge-vs-git-rebase-avantages-inconvenients-et-meilleures-pratiques.md
source_anchor: ""
source_lines: [232, 244]
sha256: 464af75b39334574b1abb3f2f3f92dc097c0c5f7df8c99ca1f5033dc93e73246
---

# On main, bring in feature-xyz:

**Utilisez git merge lorsque vous travaillez dans des environnements collaboratifs où plusieurs développeurs contribuent à la même base de code et où vous devez préserver le contexte dans lequel les fonctionnalités ont été développées ensemble. Merge est essentiel pour les projets sensibles à l'audit qui nécessitent une traçabilité complète des modifications de code à des fins de conformité. C'est également le choix le plus sûr lorsque vous travaillez avec des branches publiques sur lesquelles d'autres membres de l'équipe ont basé leur travail, car la fusion ne réécrit pas l'historique des livraisons existantes. En outre, la fusion fonctionne bien pour les équipes qui préfèrent des limites de fonctionnalités claires et des cycles d'intégration périodiques plutôt qu'une conservation continue de l'historique.**

### Quand dois-je utiliser git rebase au lieu de git merge ?

**Git rebase est idéal pour les branches privées où vous êtes le seul développeur à effectuer des changements et où vous souhaitez présenter une séquence logique et propre de livraisons à votre équipe. Utilisez rebase lorsque vous avez besoin d'éliminer les livraisons expérimentales, de corriger les messages de livraison ou de réorganiser les modifications en une histoire plus cohérente avant de partager votre travail. Il est particulièrement utile pendant les phases de développement actives, lorsque vous souhaitez conserver un historique linéaire du projet, facile à suivre et à déboguer. Rebase convient également aux équipes qui privilégient l'efficacité de l'examen du code et préfèrent l'historique des livraisons à l'authenticité chronologique.**

### Puis-je utiliser à la fois git merge et git rebase dans le même projet ?

**Oui, de nombreuses équipes performantes adoptent des flux de travail hybrides qui combinent les deux stratégies en fonction du contexte et de la phase de développement. Une approche courante consiste à utiliser rebase pour le nettoyage du développement local afin d'organiser et de peaufiner vos commits, puis à passer à merge pour l'intégration au niveau de l'équipe afin de préserver le contexte de collaboration. Vous pouvez rebaser votre branche de fonctionnalités pour créer une séquence de validation propre, puis la fusionner avec la branche principale pour maintenir des limites claires dans l'historique partagé. Cette approche hybride vous permet de bénéficier des avantages des deux stratégies tout en évitant leurs inconvénients respectifs.**

### Que se passe-t-il si je rebase une branche publique sur laquelle d'autres travaillent ?

**Rebaser une branche publique sur laquelle d'autres ont basé leur travail crée de sérieux problèmes de coordination et peut dupliquer des commits dans votre historique partagé. Lorsque vous rebasez, vous modifiez les hachages SHA des commits, ce qui signifie que les branches des autres développeurs n'auront plus les commits parents corrects. Cela oblige les membres de l'équipe à effectuer des opérations de récupération complexes ou à perdre potentiellement leur travail lorsqu'ils essaient de fusionner leurs modifications. Si vous devez absolument rebaser une branche publique, coordonnez à l'avance avec tous les membres de l'équipe concernés et envisagez d'utiliser `git rebase --onto` ou d'autres techniques avancées similaires pour minimiser les interruptions.**
