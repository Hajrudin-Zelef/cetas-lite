---
id: collect-261001-general-networking/general-networking/fr-review-unlocking-the-power-of-nas-exploring-qnaps-top-applications-7c4f1527-2
title: "fr-review-unlocking-the-power-of-nas-exploring-qnaps-top-applications-7c4f1527"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "gpu"]
source: docs/RAG/collect-261001-general-networking/fr-review-unlocking-the-power-of-nas-exploring-qnaps-top-applications-7c4f1527.md
source_anchor: ""
source_lines: [28, 42]
sha256: 82ec781b8f89d5012c70afc702816eca70ed2dd942e9f7d978a8f3bce98dab9a
---

# fr-review-unlocking-the-power-of-nas-exploring-qnaps-top-applications-7c4f1527

L'exécution de cette application directement sur le NAS, par opposition à l'intérieur d'une autre machine virtuelle avec Docker ou une autre plate-forme de conteneurisation, offre de nombreuses fonctionnalités, telles que la possibilité de modifier les configurations en ligne et de mettre à jour en temps réel en recréant des conteneurs en cours d'exécution ou arrêtés. Les utilisateurs peuvent également enregistrer les commandes de conteneur fréquemment utilisées pour une gestion efficace. De plus, l'une des fonctionnalités les plus utiles est que les administrateurs sont capables de télécharger des fichiers image et YAML à partir de référentiels cloud ou du NAS local sur lequel Container Station s'exécute activement.
Passage PCIe
Dans le passé, les appareils NAS n'offraient qu'un relais GPU discret pour améliorer la fonctionnalité des appareils. Cependant, QNAP prend désormais en charge PCIe Passthrough pour une gamme de périphériques, notamment les GPU, les adaptateurs Fibre Channel et les cartes d'extension USB 3.1.
L'avantage de l'utilisation de périphériques PCIe Passthrough avec des machines virtuelles est qu'elle permet aux utilisateurs de décharger les ressources informatiques du processeur lors de l'exécution de ces périphériques dans un environnement de VM. Les utilisateurs peuvent se connecter directement à des équipements externes, ce qui augmente les vitesses de transfert du réseau et des données et fournit même des options d'accélération matérielle pour les applications exécutées dans les machines virtuelles. En ajoutant cette fonctionnalité, QNAP a considérablement amélioré la convivialité de la gestion et du fonctionnement de sa machine virtuelle, notamment avec la prise en charge du GPU.
Sauvegarde des données avec QNAP
L'emplacement principal de gestion des données dans n'importe quel NAS QNAP est l'application Stockage et instantanés. Cela fournit un aperçu complet des disques et des données actuellement stockés sur l'appareil. Avec cette application, les administrateurs peuvent voir l'état des disques installés, gérer les périphériques de stockage externes et bien d'autres fonctionnalités.
Cette application permet également de créer des instantanés , qui peuvent être stockés directement sur les disques ou enregistrés ailleurs sur un support différent pour une protection accrue en cas de panne complète du disque.
Synchronisation de sauvegarde hybride
L'application QNAP la plus connue pour la sécurité des données est sans doute Hybrid Backup Sync , désormais disponible en version 3. Solution de stockage principale pour la sauvegarde et la restauration, HBS combine sauvegarde, restauration et synchronisation des données dans une interface compacte et intuitive. Idéale pour la reprise après sinistre, HBS est compatible avec les NAS QNAP, les serveurs distants et de nombreux fournisseurs de stockage cloud.
L'écran principal de HBS affiche des informations sur toutes les tâches en cours d'exécution, les alertes survenues et offre à l'administrateur un portail simple pour créer un nouvel espace de stockage dans lequel configurer de nouvelles tâches de sauvegarde ou de synchronisation avec des services cloud ou des serveurs distants. .
L'écran Espace de stockage est l'endroit où les administrateurs peuvent gérer toutes les connexions et le stockage actuels sur les services externes et cloud.
Lorsque vous sélectionnez Ajouter un nouvel espace de stockage, vous pouvez voir clairement combien d'options sont disponibles pour vous connecter de manière native.
Réflexions de clôture
Alors que beaucoup considèrent les systèmes NAS comme des périphériques de stockage, QNAP possède plus de 65 de ses propres applications et autres outils qui peuvent ajouter une tonne de valeur, la plupart à peu ou pas de frais. Lorsque vous achetez presque tous les QNAP du portefeuille, vous avez accès à cette vaste bibliothèque. Les intégrateurs de systèmes et les revendeurs à valeur ajoutée de canal le comprennent : leurs clients sont régulièrement ravis lorsqu'ils découvrent qu'un NAS peut être si polyvalent, répondant non seulement au besoin de stockage partagé d'une organisation, mais bien plus encore.
Que vous cherchiez à étendre vos capacités de stockage actuelles, à consolider les multiples disques externes empilés sur votre bureau ou que vous ayez simplement besoin de quelque chose qui remplisse de nombreux rôles différents, de la virtualisation au serveur multimédia en passant par la sauvegarde et la synchronisation, les périphériques NAS QNAP méritent d'être étudiés. Nous venons d'aborder ici certaines des applications les plus populaires, rendez-vous service et explorez la bibliothèque d'applications QNAP.
