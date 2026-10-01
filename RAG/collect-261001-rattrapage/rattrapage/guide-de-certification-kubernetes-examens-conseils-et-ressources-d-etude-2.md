---
id: collect-261001-rattrapage/rattrapage/guide-de-certification-kubernetes-examens-conseils-et-ressources-d-etude-2
title: "guide-de-certification-kubernetes-examens-conseils-et-ressources-d-etude"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-rattrapage/guide-de-certification-kubernetes-examens-conseils-et-ressources-d-etude.md
source_anchor: ""
source_lines: [118, 247]
sha256: 75c49709117198e0693f2850378082fcebec7a97a3edb73ff3974d257b06d930
---

# guide-de-certification-kubernetes-examens-conseils-et-ressources-d-etude

C'est le meilleur point de départ pour les débutants qui veulent apprendre Kubernetes sans la complexité de l'administration pratique des clusters.

La certification KCNA couvre :

- Principes fondamentaux de Kubernetes et des concepts cloud-native.
- Comprendre l'architecture et les composants de Kubernetes
- Connaissances de base en matière de conteneurisation(Docker, OCI)
- Introduction aux pratiques DevOps natives pour le cloud.

Détails de l'examen :

- Format de l'examen : Examen à choix multiples (en ligne)
- La durée de l'opération est de deux ans : 90 minutes
- Note de passage : 75%
- Prérequis : Pas de conditions préalables formelles
- Coût : 250 $ (comprend une reprise gratuite)
- Recertification : Obligatoire tous les 2 ans

### Associé en sécurité Kubernetes et Cloud-Native (KCSA)

Lacertification Kubernetes and Cloud-Native Security Associate (KCSA)est la plus récente de toutes les certifications et constitue un titre d'entrée de gamme pour les professionnels qui se concentrent sur la sécurité de Kubernetes.

Contrairement au CKS, qui est très avancé et nécessite un CKA, le KCSA fournira une perspective de sécurité fondamentale pour les débutants.

Il s'agit d'un point de départ pour les professionnels soucieux de la sécurité avant de passer au CKS.

La certification KCSA couvre

- Bases de la sécurité cloud-native et de la sécurité des composants de cluster Kubernetes.
- Compréhension des fondamentaux de la sécurité de Kubernetes.
- Aperçu du modèle de menace de Kubernetes et de la sécurité de la plateforme.
- Introduction aux cadres de conformité et de sécurité

Détails de l'examen :

- Format de l'examen : Examen à choix multiples (en ligne)
- La durée de l'opération est de deux ans : 90 minutes
- Note de passage : 75%
- Prérequis : Pas de conditions préalables formelles
- Coût : 250 $ (comprend une reprise gratuite)
- Recertification : Obligatoire tous les 2 ans

## Comment obtenir la certification Kubernetes

Je vais maintenant vous guider sur la façon d'obtenir la certification Kubernetes dans les prochains sous-chapitres.

### Préparation à l'examen

La première étape consiste à comprendre le format et lesexigences de l'examen. Visitez la page officielle de la certification CNCF et lisez les objectifs de l'examen.

Les examens CKA, CKAD et CKS sont tous des tests pratiques basés sur la performance, tandis que les examens KCNA et KCSA sont des examens à choix multiples.

Les examens sont passés en ligne via un système de surveillance à distance et sont limités dans le temps.

Vous pouvez utiliser la documentation officielle de Kubernetes pendant les examens pratiques, mais vous devez savoir où trouver rapidement les informations pertinentes.

Vous devez choisir les bonnes ressources d'apprentissage !

La documentation officielle de Kubernetes est la ressource la plus essentielle pour les examens. Puisqu'il est autorisé pendant les examens pratiques, vous devriez vous entraîner à naviguer efficacement pour trouver les réponses rapidement. Il est fortement recommandé de mettre en signet les sections clés telles que Pods, Deployments, Services, RBAC et Troubleshooting.

Vous pouvez améliorer votre apprentissage en suivant des cours en ligne. Par exemple, regardez les cours Introduction à Kubernetes et à la conteneurisation et Virtualisation avec Docker et Kubernetes. La page web officielle de la CNCF, qui propose également des formations sur Kubernetes, est une autre ressource pour les cours.

Enfin et surtout, passez des examens blancs ! C'est essentiel, car les examens de certification sont limités dans le temps, et si vous ne savez pas créer rapidement des ressources, vous perdrez le temps nécessaire. Vous pouvez utiliser le simulateur d'examen officiel de la CNCF, Killer.sh.

### Pratique

Comme la plupart des certifications Kubernetes sont pratiques, l'expérience pratique est le facteur le plus critique pour réussir !

Commencez par configurer votre propre cluster Kubernetes local à l'aide d'outils comme Minikube ou un cluster Kubernetes basé sur le cloud à partir de l'un des fournisseurs de cloud, comme AWS, GCP ou Azure. Ensuite :

- Déployez des applications réelles dans votre cluster et jouez avec les ressources.
- Déployer et gérer des pods, des déploiements, des services et des contrôleurs d'entrée.
- Configurez les ConfigMaps, les Secrets et les PersistentVolumes.
- Configurez le contrôle d'accès basé sur les rôles (RBAC) pour les autorisations des utilisateurs.
- Mettre en œuvre des politiques de réseau et résoudre les problèmes liés aux clusters.

Consultezle tutoriel Kubernetes pour apprendreà configurer votre cluster local avec minikube et déployer un petit serveur web dans ce cluster.

### Inscription à l'examen

Vous pouvez vous inscrire pour votre certification sur la page officielle des certifications de la CNCF.

Planifiez votre examen dans les 12 mois suivant l'achat.

Comme les examens Kubernetes sont passés en ligne avec un surveillant qui vous suit en direct pendant l'examen, vous devez vous assurer que vous avez.. :

- Une webcam et un microphone en état de marche pour la vérification de l'identité.
- Un environnement calme et sans distractions.
- Le navigateur sécurisé de l'ISP est installé.
- Une connexion internet stable (filaire de préférence).

### Conseils pour le jour de l'examen

Chaque question a un poids différent. Consacrez plus de temps aux tâches les plus importantes.

Si une question est trop difficile, marquez-la pour révision et passez à la suivante. Essayez de terminer avec au moins 10 minutes pour réviser les questions signalées.

Soyez rapide avec la documentation Kubernetes ! C'est important, car vous pouvez perdre beaucoup de temps à chercher dans la documentation.

Un autre conseil important est de connaître les ordres de `kubectl`. Mémorisez les plus importantes !

Utilisez l'argument `--dry-run=client -o yaml` pour obtenir les manifestes d'une commande `kubectl` et rediriger la sortie vers un fichier que vous pourrez manipuler par la suite. C'est un gain de temps considérable !

Par exemple :

`kubectl -n namespace create deploy app --image=nginx:latest --replicas=1 --dry-run=client -o yaml > nginx-deployment.yaml`
Pensez à utiliser des alias de commande pour accélérer la saisie. J'ai utilisé le site `alias k=kubectl` pour gagner du temps.

Enfin, ne paniquez pas si vous êtes bloqué ! Utilisez les ressources disponibles et concentrez-vous sur la résolution des problèmes que vous pouvez résoudre.

## Que faire après l'examen ?

Cette section décrit brièvement ce qui vous attend après l'examen. Il déterminera rapidement quand attendre les résultats et comment procéder si vous échouez au premier essai.

### Réception des résultats

Les résultats de l'examen sont généralement disponibles dans un délai de 24 à 36 heures. Si vous réussissez, vous recevrez un certificat numérique et un badge CNCF, que vous pourrez partager sur LinkedIn, GitHub et votre CV.

### Recertification

Si vous avez échoué à votre première tentative, ne vous inquiétez pas ; vous avez droit à un nouvel essai gratuit. Planifiez-le judicieusement après avoir passé en revue vos points faibles ! Concentrez-vous sur vos points faibles et repassez des examens blancs avant de réessayer.

## Meilleures pratiques pour préparer la certification Kubernetes

Se préparer à une certification Kubernetes nécessite plus que d'étudier les parties théoriques. Elle nécessite un apprentissage pratique et structuré ainsi qu'une approche stratégique.

Les meilleures pratiques suivantes vous aideront à maximiser vos chances de réussir votre examen du premier coup.

### Commencez par l'essentiel

