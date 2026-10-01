---
id: collect-261001-huawei/huawei/openshift-vs-kubernetes-explication-des-principales-differences-3
title: "openshift-vs-kubernetes-explication-des-principales-differences"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/openshift-vs-kubernetes-explication-des-principales-differences.md
source_anchor: ""
source_lines: [152, 205]
sha256: 919f2a38394924d9d68d8604b5785640dc2b575a7f9b0699faeb70a17c916f9e
---

# openshift-vs-kubernetes-explication-des-principales-differences

Kubernetes et OpenShift prennent tous deux en charge le RBAC, ce qui vous permet de contrôler qui a accès à quoi au sein du cluster.

- Kubernetes fournit un système RBAC hautement configurable, vous permettant de définir des rôles et de les attribuer à des utilisateurs ou à des comptes de service. Cependant, c'est à l'administrateur de la grappe de définir les rôles correctement, et il n'existe pas de structure imposée sur la manière dont l'accès doit être défini.
- OpenShift s'appuie sur le modèle RBAC de Kubernetes, mais inclut des rôles prédéfinis et un modèle d'accès basé sur les projets qui facilitent la gestion cohérente des accès. Chaque projet OpenShift (à peu près équivalent à un espace de noms Kubernetes) est accompagné de son propre ensemble de permissions, ce qui aide les équipes à isoler les charges de travail et à gérer l'accès de manière plus intuitive.

OpenShift est donc mieux adapté aux équipes qui souhaitent mettre en place un contrôle d'accès sans partir de zéro.

### Conformité et surveillance intégrées

- Dans Kubernetes, la surveillance et la conformité sont généralement des modules complémentaires. Vous pouvez installer des outils tels que Prometheus, Grafana, Falco ou des systèmes de journalisation d'audit. Cependant, ces derniers nécessitent une configuration et une maintenance manuelles. Cela vous donne de la flexibilité mais augmente les frais généraux.

- OpenShift comprend une pile de surveillance intégrée (Prometheus, Alertmanager et Grafana), une journalisation centralisée et des outils tels que le Compliance Operator. Ils sont intégrés à la plateforme et pris en charge dès le départ, ce qui permet aux équipes de savoir ce qui est en cours d'exécution et si cela correspond aux normes de sécurité.

Pour les équipes travaillant dans des secteurs réglementés, tels que la finance, la santé et l'administration, cela permet de gagner beaucoup de temps et d'efforts et de réduire le risque de passer à côté de quelque chose de crucial.

## Développeur et expérience utilisateur

Une part importante du choix entre Kubernetes et OpenShift se résume à la sensation qu'ils procurent à l'utilisation.

Alors que Kubernetes fournit des blocs de construction puissants, OpenShift ajoute une couche raffinée qui simplifie les flux de travail et améliore la productivité.

Cette section examine le processus de développement, de déploiement et de gestion des applications sur chaque plateforme.

### Interface et facilité d'utilisation

- Kubernetes est principalement piloté par la ligne de commande, avec `kubectl` comme interface principale. Il propose également un tableau de bord en ligne, qui est relativement basique et souvent désactivé par défaut pour des raisons de sécurité. La plupart des utilisateurs interagissent avec Kubernetes par le biais de fichiers de configuration et de commandes CLI. Cela vous donne beaucoup de contrôle, mais s'accompagne également d'une courbe d'apprentissage abrupte, en particulier pour les nouveaux arrivants.
- OpenShift, en revanche, comprend une console web complète conçue à la fois pour les développeurs et les administrateurs de clusters. Vous pouvez créer et gérer des projets, déployer des applications à partir de dépôts Git ou d'images de conteneurs, afficher des journaux et des mesures, et même déclencher des constructions sans toucher à la ligne de commande. L'interface utilisateur est bien intégrée et conviviale, ce qui abaisse la barrière d'entrée pour les équipes qui ne disposent pas d'une expertise approfondie de Kubernetes.

J'aime travailler avec l'interface utilisateur d'OpenShift, car elle vous permet de gérer vos applications facilement et rapidement. En particulier, le débogage est devenu beaucoup plus rapide grâce à l'utilisation de l'interface utilisateur au lieu de s'appuyer entièrement sur les commandes de l'interface de programmation.

Je vous recommande de lire What's New in the OpenShift 4.4 Web Console Developer Experience pour ensavoir plus sur les fonctionnalités et la conception de la console web d'OpenShift.

OpenShift prend également en charge son propre outil CLI (`oc`), qui étend `kubectl` et fournit des commandes supplémentaires adaptées aux fonctionnalités d'OpenShift. Les utilisateurs plus expérimentés disposent ainsi du même niveau de contrôle que celui qu'ils attendent de Kubernetes.

### Intégration des outils DevOps

- Kubernetes offre une intégration avec des outils DevOpstels que Jenkins, Argo CD, Tekton et Flux. Mais leur mise en place implique généralement un travail supplémentaire. Vous devez les installer et les configurer manuellement, gérer les identifiants et construire vos propres pipelines CI/CD.
- OpenShift simplifie cela en incluant OpenShift Pipelines (basé sur Tekton) et OpenShift GitOps (basé sur Argo CD) comme composants natifs. Ceux-ci sont étroitement intégrés à la plateforme, de sorte que vous pouvez construire, tester et déployer directement à partir de l'interface utilisateur ou du CLI sans dépendre de plugins externes. OpenShift prend également en charge les configurations de construction, les flux d'images et les déploiements automatiques déclenchés par Git ou les changements d'image. Ces fonctionnalités facilitent la livraison en continu.

### Ecosystème et extensibilité

- Kubernetes dispose d'un écosystème massif. Vous trouverez des milliers d'outils open-source, des graphiques Helm et des modules complémentaires pris en charge par la communauté pour tout ce qui concerne la surveillance et le traçage, les maillages de services et les charges de travail d'intelligence artificielle. Il est flexible et extensible, mais cette flexibilité signifie que vous devez prendre plus de décisions et maintenir plus de composants vous-même.
- OpenShift prend également en charge Helm, les CRD et les extensions natives de Kubernetes, mais gère son écosystème par l'intermédiaire d' OperatorHub. OperatorHubfournit un catalogue centralisé, soutenu par Red Hat, qui facilite la recherche et le déploiement de composants prêts pour la production.

En bref, Kubernetes vous donne les éléments de base pour assembler votre plateforme. OpenShift vous en donne plus dès le départ, avec des outils déjà intégrés et pris en charge, ce qui permet de gagner du temps et de réduire la complexité.

## Soutien et écosystème

Le succès à long terme d'une plateforme dépend souvent non seulement de ses fonctionnalités, mais aussi de son écosystème et du type de soutien qu'elle offre. Ceci est particulièrement important lorsque vous travaillez dans une grande entreprise, car vous avez besoin d'une assistance complète pour votre environnement de production.

Kubernetes et OpenShift disposent tous deux de communautés florissantes et du soutien des fournisseurs, mais ils adoptent des approches très différentes en matière de gouvernance, de responsabilité des fournisseurs et de préparation des entreprises.

### Soutien de la communauté et des fournisseurs

