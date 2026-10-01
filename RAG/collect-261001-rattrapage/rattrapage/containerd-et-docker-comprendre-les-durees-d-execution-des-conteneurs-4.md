---
id: collect-261001-rattrapage/rattrapage/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs-4
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "sandbox"]
source: docs/RAG/collect-261001-rattrapage/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [214, 290]
sha256: f16e9326ee723c1bf6a543dfaad5dee3bf4e9fcb59b3339c4f0fa0417e967c1d
---

# containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs

### Pour les environnements spécialisés et minimalistes

Bien que Kubernetes soit le cas d'utilisation le plus courant en production, la nature légère de containerd ouvre la voie à des scénarios de déploiement où Docker serait peu pratique.

L'informatique en périphérie et les appareils IoT fonctionnent souvent dans des conditions de ressources très limitées. La conception légère de Containerd le rend viable dans ces environnements où la pile complète de Docker serait prohibitive. Chaque mégaoctet de mémoire et chaque cycle CPU sont importants lors de l'exécution sur du matériel embarqué.

Les scénarios de sécurité avancés bénéficient de l'architecture d'exécution modulaire de containerd. Les organisations peuvent intégrer des environnements d'exécution sandboxés pour les charges de travail nécessitant des limites de sécurité supplémentaires. Voici quelques exemples :

- **gVisor : assure** une isolation efficace du noyau
- **Kata Containers :** exécute des conteneurs dans des machines virtuelles légères.

Ces intégrations s'intègrent à containerd sans nécessiter de modifications importantes.

## Transition de Docker vers Containerd

Il est important de comprendre quand utiliser chaque outil, mais il est tout aussi crucial de savoir comment passer de l'un à l'autre. La transition de Docker vers containerd dans les environnements existants nécessite une planification minutieuse, mais le processus est bien documenté et simple.

### Migration des nœuds Kubernetes

Les étapes opérationnelles pour migrer les nœuds Kubernetes de Docker vers containerd suivent un modèle standard :

1. 
**Veuillez isoler le nœud :** Empêcher la planification de nouveaux pods (`kubectl cordon` )
2. 
**Vider les pods existants :** Transférez les charges de travail vers d'autres nœuds (`kubectl drain` )
3. 
**Mettre à jour la configuration Kubelet :** Veuillez pointer vers le socket CRI de containerd à l'adresse suivante :`/run/containerd/containerd.sock`
4. 
**Vérifier les plugins CNI :** Veuillez vous assurer que les plugins réseau nécessaires sont installés sur containerd.
5. 
**Veuillez redémarrer Kubelet :** Veuillez vous enregistrer avec le nouveau runtime et rejoindre à nouveau le cluster.

Tester la migration sur des nœuds hors production permet d'identifier les problèmes spécifiques à l'environnement avant de déployer les modifications à l'échelle du cluster. Afin d'éviter les erreurs courantes, veuillez respecter les bonnes pratiques suivantes :

- **Modifications du chemin d'accès au journal :** Docker et containerd utilisent des emplacements de journaux par défaut différents. Veuillez mettre à jour votre infrastructure de journalisation en conséquence.
- **Veuillez installer les plugins CNI manquants :** Containerd nécessite les binaires du plugin CNI pour la mise en réseau ; ceux-ci ne sont pas toujours installés par défaut.
- **Veuillez prêter attention aux différences de tirage d'image :** Les paramètres d'authentification et de registre peuvent nécessiter des ajustements.
- **Veuillez prendre garde aux incompatibilités entre les pilotes de stockage :** Veuillez vous assurer que vos volumes persistants sont compatibles avec le snapshotter de containerd.

### Comprendre les différences entre les interfaces CLI

Une fois votre infrastructure migrée, les développeurs devront adapter leurs processus quotidiens afin de travailler avec le nouveau runtime.

Pour les développeurs habitués aux commandes Docker, nerdctl offre une expérience similaire. Les commandes telles que ` `nerdctl run``, ` `nerdctl build`` et ` `nerdctl compose up` ` fonctionnent exactement comme leurs équivalents Docker, ce qui rend la transition transparente.

Pour le débogage, il est essentiel de comprendre comment mapper les workflows de débogage Docker vers containerd. Si vous utilisez `docker inspect` pour examiner un conteneur, `ctr containers info` fournit des informations similaires, mais dans un format différent. De même, `ctr tasks list` affiche les conteneurs en cours d'exécution.

La plupart des développeurs constatent qu' `nerdctl` élimine la nécessité d'apprendre la syntaxe d' `ctr` pour les tâches quotidiennes. L' `ctr`, de bas niveau, reste utile pour le dépannage de problèmes spécifiques à l'exécution ou lorsque l'on travaille directement avec les API de containerd.

## Conclusion

La relation entre Docker et containerd illustre parfaitement la réussite de la conception modulaire dans l'infrastructure logicielle. Docker demeure l'outil optimal pour les développeurs qui écrivent du code, car il offre une expérience intégrée, un écosystème complet et des interfaces conviviales qui rendent la création d'applications conteneurisées productive et agréable.

Containerd, quant à lui, se distingue comme le runtime idéal pour les machines exécutant du code, offrant la stabilité, les performances et le minimalisme requis pour les plateformes d'orchestration de production. Le fait que Docker Engine utilise containerd en arrière-plan démontre que ces deux outils se complètent plutôt qu'ils ne se font concurrence.

Je recommande une approche pragmatique pour la plupart des organisations : continuer à utiliser Docker sur les ordinateurs portables des développeurs, où ses outils accélèrent les workflows de développement, mais envisager de migrer les clusters Kubernetes de production vers containerd directement pour bénéficier des avantages opérationnels que sont la réduction des frais généraux et la simplification des piles d'exécution.

Que vous optiez pour la plateforme complète de Docker ou pour le runtime spécialisé de containerd, ces deux solutions restent des composants essentiels de l'écosystème moderne des conteneurs, chacune étant optimisée pour différentes étapes du cycle de vie des applications.

Pour continuer à vous former, nous vous invitons à vous inscrire à notre cursus de compétences « Conteneurisation et virtualisation avec Docker et Kubernetes ».

## FAQ sur Containerd et Docker

### Est-il possible d'utiliser des images Docker avec containerd ?

**Oui, tout à fait. Containerd prend en charge toutes les images de conteneur conformes à la norme OCI, y compris celles créées avec Docker. Étant donné que Docker crée des images conformes à la norme OCI, celles-ci fonctionnent parfaitement avec containerd et tous les autres environnements d'exécution compatibles OCI. Vous pouvez utiliser `docker build` localement et exécuter ces images avec containerd en production sans aucun problème de compatibilité.**

### Docker utilise-t-il containerd en arrière-plan ?

**Oui, Docker Engine utilise containerd comme moteur d'exécution de conteneurs principal. À partir de la version 1.11, Docker a intégré containerd pour gérer les opérations liées au cycle de vie des conteneurs, telles que la création, l'exécution et la gestion. Lorsque vous exécutez `docker run`, le démon Docker (`dockerd`) délègue l'exécution effective du conteneur à containerd, qui utilise ensuite runc pour interagir avec le noyau Linux.**

### Pourquoi Kubernetes a-t-il supprimé la prise en charge de Docker ?

**En 2022, Kubernetes a supprimé dockershim (la couche de compatibilité Docker) afin d'éliminer une couche de traduction superflue. Docker est antérieur à l'interface d'exécution de conteneurs (CRI), donc Kubernetes avait besoin de dockershim pour assurer la conversion entre ses API et Docker. En communiquant directement avec containerd via CRI, Kubernetes atteint de meilleures performances, une plus grande stabilité et une pile d'exécution plus simple. Les images Docker continuent de fonctionner parfaitement dans Kubernetes.**

### Devrais-je passer de Docker à containerd pour le développement local ?

