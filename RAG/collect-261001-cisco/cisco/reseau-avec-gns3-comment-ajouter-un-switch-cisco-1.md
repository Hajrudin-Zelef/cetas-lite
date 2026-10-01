---
id: collect-261001-cisco/cisco/reseau-avec-gns3-comment-ajouter-un-switch-cisco-1
title: "reseau-avec-gns3-comment-ajouter-un-switch-cisco"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["amd", "intel"]
source: docs/RAG/collect-261001-cisco/reseau-avec-gns3-comment-ajouter-un-switch-cisco.md
source_anchor: ""
source_lines: [1, 137]
sha256: c83680f3ce0bcccc5a320ae2d3e10efb2ad10a9a0fabfd8c3cde7d7dc1f9d703
---

# reseau-avec-gns3-comment-ajouter-un-switch-cisco

## I. Présentation

**Dans le premier article d'introduction à GNS3, nous avons vu que GNS3 est un outil qui nous permet de maquetter des infrastructures en y intégrant, en autres, des équipements réseaux. Dans ce second article, nous allons voir comment ajouter un commutateur Cisco à GNS3**. Cisco a développé une version virtuelle de son système d’exploitation IOS pour commutateur baptisée IOSv déclinée en deux versions :

- **IOSvL2** : prends en charge la plupart des fonctionnalités présentes dans les commutateurs CISCO de niveau 2 et de Niveau 3
- **IOSv** : prends en charge la plupart des fonctionnalités présentes dans les routeurs CISCO

Vous pouvez retrouver le premier article en suivant le lien ci-dessous :

## II. Importer une Image Cisco IOSvL2 dans GNS3

Tout d'abord, nous devons installer l’image CISCO IOSvL2 et nous devons récupérer l’appliance sur le site internet. Dans GNS3, cela est appelé GNS3, mais dans les faits, c’est plus un outil qui nous propose un assistant d’importation de l’IOS. En effet, l’IOSv est disponible à l’achat sur le site internet de CISCO. Sur le site internet de GNS3, nous avons accès à un « marketplace », depuis la page d’accueil, nous cliquerons sur le menu « **Marketplace** ».

Après avoir cliqué sur « **Appliances** », nous téléchargerons l’**appliance Cisco IOSvL2**. Il s'agit d'un fichier au format « **gns3a** ».

Une fois le fichier téléchargé, nous allons l’importer dans GNS3. Ouvrez l'application. Nous déroulerons le menu « **File** » où nous sélectionnerons « **Import appliance** ».

Nous sélectionnerons le fichier que nous avons téléchargé depuis le site internet de GNS3, à savoir "**cisco-iosvl2.gns3a**".

L’assistant d’importation de l’IOSvL2 se lance...

A la première étape, nous sommes invités à sélectionner où nous souhaitons l’installer. Comme nous l’avons précisé dans le premier article, les commutateurs et les routeurs CISCO ne sont pas basés sur des architectures processeur x64 ou x32. Nous devons donc émuler le système pour faire fonctionner le système d’exploitation de l’IOS pour switch de chez CISCO.

Nous utiliserons l’**émulateur QEMU présent dans l’appliance GNS3 exécutée dans VMware Workstation**, c’est ce seul choix qui nous est proposé. Nous pouvons cliquer sur le bouton « **Next** ».

Nous sélectionnerons les binaires QEMU que nous utiliserons, ses binaires sont présents dans la machine virtuelle. Nous utiliserons la version par défaut, qui est aussi la dernière version la plus récente présente dans la machine virtuelle GNS3 VM.

**version IOSvL2 version 15.2(4.0.55)E**, nous sélectionnerons donc l’IOS de cette version. Ainsi, nous obtiendrons l'accès aux boutons « **Import** » et « **Download** ».

Nous cliquerons sur le bouton « **Import** », puis nous sélectionnerons **le fichier de notre image CISCO**, puis dans la boîte de dialogue, nous sélectionnerons notre fichier « **vios_l2-adventerprisek9-m.vmdk.SSA.152-4.0.55.E** ».

Nous déroulerons ensuite le nœud de la version que nous venons d’installer, nous sélectionnons le nom de fichier de notre IOSv, puis nous cliquerons sur le bouton "**Next**".

**Yes** » pour confirmer l’installation.

Nous aurons les instructions pour l’utilisation de notre Switch. Nous cliquerons sur le bouton « **Finish** ».

Une fois l’opération terminée, nous avons une boîte de dialogue, nous confirmant l’installation de notre Switch avec un simple clic sur "**OK**".

## III. Créer ses deux premiers Switch Cisco dans GNS3

Il est temps de tester afin de vérifier le bon fonctionnement notre image CISCO. Pour cela, nous allons créer un nouveau projet dans lequel nous ajouterons deux switchs Cisco.

Pour créer un nouveau projet, nous cliquerons sur le menu déroulant « **File** » puis sur « **New Blank Project** ».

**un lab est appelé « projet »**, nous devons lui donner un nom et un emplacement où seront installés les différents fichiers utilisés pour les maquettes. Une fois ces informations renseignées, nous cliquerons sur le bouton « **OK** ». Voici un exemple :

Dans GNS3, au sein de la catégorie commutateurs, nous pouvons visualiser notre switch CISCO.

Nous glisserons le switch identifié avec le nom **Cisco IOSvL2 15.2(4.0.55)E** sur l’espace de travail, à droite. Grâce à cette action, on crée un switch.

Nous recommençons cette opération pour que nous ayons **deux commutateurs sur l’espace de travail** puis nous cliquerons sur le **triangle vert pour démarrer les deux équipements**.

Nous lierons les deux équipements à l’aide de l’outil « **link** » accessible dans le menu latéral.

Pour lier, avec l’outil « **link** », nous cliquerons sur le premier switch. Là, nous aurons la liste des ports. Nous sélectionnerons le **port Gi0/0**.

Nous cliquerons ensuite sur **le second switch**, nous sélectionnerons également le même port. **Nous venons de créer un lien entre les deux switchs.**

**terminal pour la configuration**. Nous attribuerons une adresse IP au VLAN 1 sur chacun des équipements et nous effectuerons un ping entre les deux commutateurs.

Sur le **premier switch**, on attribue l'adresse IP "**192.168.0.1/24**" :

Switch>enable
Switch#configure terminal
Switch(config)#interface vlan 1
Switch(config-if)#ip address 192.168.0.1 255.255.255.0
Switch(config-if)#no shutdown

En image :

Sur le **second switch**, on attribue l'adresse IP "**192.168.0.2/24**" :

Switch>enable
Switch#configure terminal
Switch(config)#interface vlan 1
Switch(config-if)#ip address 192.168.0.1 255.255.255.0
Switch(config-if)#no shutdown

En image :

Il ne reste plus qu'à tester un ping dans les deux sens : vous verrez que vos switchs sont en mesure de communiquer !

## IV. Conclusion

**Félicitations ! Vous venez d'importer une image CISCO dans GNS3, de créer vos premiers switchs et de les faire communiquer ensemble !**

bonjour

petite question

es ce que ce sont des switch niv3 ?

et merci pour ce tuto très interessant qui nous offre une alternative a packet tracer .

Hello,

Oui, Il en a. Ces son ceux qui ont dans le nom L3

Bonjour Loïc,

Merci pour ton intérêt, la version de l’IOS (vIOSL2) intègre les fonctionnalités des switch de niveau 3.

Bonjour Monsieur,

Pourriez-vous s’il vous plaît rédiger un tutoriel pour installer la fonctionnalité de GSN3 « One server, multiples clients » ? En effet, cette fonctionnalité peut-être très utile pour les projets de groupes ! 🙂

Je bloque dessus et il n’y a pas tant de tutos que ça

Bonne journée

Bonjour Youssef,

C’est tout à fait envisageable de faire un article sur GNS3 multi utilisateur. Je le mets dans ma to do list.

BOnjour,

Au moment d’importer le fichier ,’ai une erreur indiquant que le serveur ne supporte pas KVM, j’ai bien utilisé le fichier OVA téléchargé lors de l’installation.

https://i.imgur.com/WHG4sNo.png

Une idée ?

Bonjour Thomas,

Au niveau de la machine virtuelle GNS3 VM, est ce que le processeur virtuel prend bien en charge la virtualisation. Pour vous en assurer vérifiez que la case virtualize intel vt-x/ept or amd-v/rvi soit bien cochée dans VMware Workstation.

https://ibb.co/zSWgL1V

Ensuite, sur la Machine virtuelle de GNS3 après qu’elle est fini de démarrer, vérifier que « KVM support available » soit bien à TRUE.

https://ibb.co/qsKJf6K

Je vous laisse revenir vers moi pour votre retour.

Bonjour comment avez-vous fait pour télécharger le fichier vios_l2-adventerprisek9-m.vmdk.SSA.152-4.0.55.E

