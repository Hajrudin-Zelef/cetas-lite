---
id: collect-261001-rattrapage/rattrapage/principales-alternatives-a-docker-en-2026-un-guide-complet-4
title: "Buildah scripting approach with CI integration"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/principales-alternatives-a-docker-en-2026-un-guide-complet.md
source_anchor: ""
source_lines: [191, 236]
sha256: b0f8e0757c8305d82bc7f478cc6ca02264f8fb2317f5099888a4a92bd4ea5224
---

# Buildah scripting approach with CI integration

Les conteneurs sans racine éliminent le principal risque de sécurité des déploiements Docker traditionnels : le démon racine. Podman a été le premier à adopter cette approche en exécutant les conteneurs entièrement sous les privilèges utilisateur, en utilisant les espaces de noms utilisateur Linux pour mapper l'utilisateur root du conteneur à un identifiant utilisateur non privilégié sur le système hôte.

La mise en œuvre repose sur des plages d'identifiants d'utilisateurs et de groupes subordonnés (`/etc/subuid` et `/etc/subgid`) qui permettent aux utilisateurs non privilégiés de créer des espaces de noms isolés. Lorsqu'un processus conteneur pense s'exécuter en tant que root (UID 0), le noyau le mappe à votre ID utilisateur réel (par exemple, UID 1000) sur l'hôte. Même si un attaquant parvient à s'échapper du conteneur, il ne peut pas dépasser les autorisations de votre utilisateur.

Containerd a mis en œuvre des fonctionnalités similaires sans accès root grâce à son mode sans accès root, qui utilise les mêmes techniques de mappage d'espace de noms utilisateur. Le runtime peut démarrer des conteneurs, gérer des images et gérer le réseau sans nécessiter de privilèges root sur le système hôte.

Kubernetes prend désormais en charge le fonctionnement sans root de manière native grâceau projet Kubernetes-in-Rootless-Docker (KIND) et à l'intégration de containerd sans root. Cela signifie que des clusters Kubernetes entiers peuvent fonctionner sans privilèges root, ce qui réduit considérablement la surface d'attaque pour les environnements multi-locataires et les déploiements périphériques où les modèles de sécurité traditionnels ne s'appliquent pas.

L'impact sur la sécurité va au-delà de la prévention de l'escalade des privilèges. Les conteneurs sans racine ne peuvent pas se connecter aux ports privilégiés (inférieurs à 1024), n'ont pas accès à la plupart des systèmes de fichiers `/proc` et `/sys`, et ne peuvent pas effectuer d'opérations nécessitant des capacités du noyau. Cela crée des limites naturelles qui permettent de contenir les éventuelles failles de sécurité.

### Intégration Seccomp/BPF

Les environnements d'exécution modernes intègrent eBPF (extended Berkeley Packet Filter) pour l'application en temps réel de politiques de sécurité qui vont au-delà des contrôles d'accès traditionnels. Les programmes eBPF s'exécutent dans l'espace noyau et peuvent surveiller, filtrer ou modifier les appels système au fur et à mesure qu'ils se produisent, offrant ainsi une visibilité et un contrôle sans précédent sur le comportement des conteneurs.

Les profils Seccomp (Secure Computing) utilisent BPF pour filtrer les appels système au niveau du noyau. Au lieu d'autoriser aux conteneurs l'accès à l'ensemble des plus de 300 appels système Linux, les profils seccomp définissent précisément les appels autorisés. Le profil seccomp par défaut de Docker bloque les appels système potentiellement dangereux, tandis que les profils personnalisés peuvent être encore plus restrictifs en fonction des exigences de l'application.

L'intégration avancée de l'eBPF permet une surveillance comportementale en temps réel. Des outils tels que Falco utilisent des programmes eBPF pour détecter les comportements anormaux des conteneurs, tels que des connexions réseau inhabituelles, des modèles d'accès aux fichiers inattendus ou des tentatives d'utilisation d'appels système bloqués. Ces détections se produisent en temps réel avec une surcharge minimale en termes de performances, car la surveillance s'effectue dans l'espace noyau.

L'application des politiques réseau via eBPF permet un contrôle granulaire du trafic au niveau des paquets. Cilium, un CNI Kubernetes très apprécié, utilise eBPF pour mettre en œuvre des politiques réseau capables de filtrer le trafic en fonction des protocoles de la couche application, et pas seulement des adresses IP et des ports. Cela signifie que vous pouvez créer des politiques telles que « autoriser les requêtes HTTP GET vers `/api/v1/users` mais bloquer les requêtes POST » directement dans le noyau.

La sécurité basée sur eBPF permet également une surveillance tenant compte des conteneurs, qui comprend la relation entre les processus, les conteneurs et les pods Kubernetes. Les outils de surveillance traditionnels examinent les processus individuels, tandis que les programmes eBPF peuvent établir une corrélation entre les appels système et les métadonnées des conteneurs afin de fournir des informations de sécurité contextuelles.

Ces fonctionnalités transforment la sécurité, qui passe d'une approche réactive de correction à une approche proactive d'application des politiques, où les comportements suspects sont automatiquement bloqués avant qu'ils ne puissent causer des dommages.

## Stratégies d'optimisation des performances

Les performances des conteneurs sont particulièrement importantes lorsque vous exécutez des centaines, voire des milliers de conteneurs sur votre infrastructure. Explorons les stratégies qui minimisent la surcharge des ressources et la latence au démarrage.

### Réduction des démarrages à froid

Le temps de démarrage à froid, c'est-à-dire le délai entre la demande d'un conteneur et sa disponibilité pour traiter le trafic, a un impact direct sur l'expérience utilisateur et l'efficacité des ressources. Des techniques ont été développées afin de minimiser ces délais dans différentes architectures d'exécution.

Les images pré-extraites éliminent le temps de téléchargement en conservant les images de conteneurs fréquemment utilisées en cache sur les nœuds. Les DaemonSets Kubernetes peuvent pré-télécharger des images critiques, tandis que les registres tels que Harbor prennent en charge la réplication d'images vers des emplacements périphériques. Cette technique permet de réduire le temps de démarrage à froid de quelques secondes à quelques millisecondes pour les images mises en cache.

L'optimisation des couches d'image réduit la quantité de données à transférer et à extraire. Les constructions en plusieurs étapes permettent d'obtenir des images finales plus petites, tandis que des outils tels que dive permettent d'identifier les couches inutiles. Les images Distroless de Google éliminent les gestionnaires de paquets et les shells, ce qui réduit souvent considérablement la taille des images.

Le chargement différé avec des projets tels que Stargz permet aux conteneurs d' s de démarrer avant que l'image entière ne soit téléchargée. Le runtime ne récupère que les fichiers nécessaires au démarrage initial, téléchargeant les couches supplémentaires à la demande. Cela peut réduire le temps de démarrage à froid de plusieurs secondes à moins d'une seconde pour les images de grande taille.

L'optimisation de l'exécution varie selon l'implémentation. L'implémentation Rust de Youki présente des performances de démarrage améliorées par rapport à runC grâce à une meilleure gestion de la mémoire. Crun, écrit en C, permet d'obtenir des améliorations similaires en éliminant la surcharge liée au ramasse-miettes de Go lors de la création de conteneurs.

Le partage de snapshots dans containerd permet à plusieurs conteneurs de partager des snapshots de systèmes de fichiers en lecture seule, réduisant ainsi la charge de stockage et de mémoire. Lorsque vous démarrez plusieurs conteneurs à partir de la même image, seules les couches inscriptibles nécessitent une allocation distincte.

L'optimisation du processus d'initialisation peut réduire le temps de démarrage en utilisant des systèmes d'initialisation légers tels que tini ou en concevant soigneusement les séquences de démarrage des applications afin de minimiser le travail d'initialisation.

### Efficacité de la mémoire

