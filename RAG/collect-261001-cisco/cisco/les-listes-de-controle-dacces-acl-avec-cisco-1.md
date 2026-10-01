---
id: collect-261001-cisco/cisco/les-listes-de-controle-dacces-acl-avec-cisco-1
title: "les-listes-de-controle-dacces-acl-avec-cisco"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-cisco/les-listes-de-controle-dacces-acl-avec-cisco.md
source_anchor: ""
source_lines: [1, 80]
sha256: 64d64bc915ac415c0f14e55787db4ae4b8f72f150a6ce73539c0d8915075e878
---

# les-listes-de-controle-dacces-acl-avec-cisco

## I. Présentation

**Dans ce tutoriel, nous allons aborder la notion d'ACL, c'est-à-dire les listes de contrôle d'accès, avec quelques exemples pratiques à mettre en œuvre sur Cisco.**

### Une ACL, c’est quoi et pourquoi ?

Les **ACL**, pour ***A**ccess **C**ontrol **L**ist*, sont des règles appliquées aux trafics transitant via les interfaces du routeur que ce soit en entrée (*in*) ou en sortie (*out*). Les ACL filtrent le trafic en demandant aux interfaces d’acheminer ou non les paquets qui y transitent. Pour ce faire, le routeur lit l'en-tête de chaque paquet afin de déterminer s'il doit être acheminé ou non en fonction des conditions définies dans la liste de contrôle d’accès ACLs.

Les avantages de les mettre en place sont nombreux, on peut citer par exemple :

- Maitriser le réseau en déterminant quel type de trafic sera acheminé ou bloqué.
- Augmenter le niveau de sécurité d’accès réseau de manière basique en accordant ou non l’accès à un segment du réseau.
- Optimiser le réseau à d'autres fins que la sécurité, par exemple pour contrôler la bande-passante, restreindre le contenu des mises à jour de routage ou identifier et classer le trafic par fonctionnalités de qualité de service (QoS).

**Les ACL s'appliquent selon un ordre séquentiel en évaluant les paquets au début de la liste d’instruction**. Si le paquet répond aux critères de la première instruction (appelées ACE pour *Access Control Entry*), il ignore le reste des règles. Pour reproduire ce tutoriel, vous pouvez utiliser du matériel virtuel Cisco sous Packet Tracer.

## II. ACL : entrante ou sortante

Une question très importante que l'on peut se poser lorsque l’on veut mettre en place l’ACLs, c'est savoir sur quelle interface faut-il appliquer les ACLs ? L’interface entrante ou sortante du routeur ? Répondons à cette question dans cette première partie de l'article.

Les ACLs peuvent être associés à une interface particulière et pour une direction du flux (en entrée ou en sortie). En outre, ces règles de filtrages peuvent être appliquées avant que le routeur ne prenne sa décision de routage (interface en entrée), ce qui est un bon moyen d'économiser les ressources matérielles du routeur, ou après que le routeur ait pris sa décision de transfert et déterminé l’interface de sortie à utiliser pour acheminer le paquet.

Un petit exemple :

Dans la topologie ci-dessus, l’interface entrante de **R1** est **Gi 1/0,** et l’interface sortante est **Gi 0/0**.

L’interface entrante de **R2** est **Gi 0/1** et l’interface sortante est **Gi 0/3**. Si par exemple, nous avons activé une ACL sur l’interface **Gi 0/2** de **R2**, cette ACL ne pourrait pas filtrer les paquets envoyés de PC2 au serveur S1, car les paquets ne passent pas l’interface **Gi 0/2**.

**En somme, pour filtrer un paquet, on doit activer une ACL sur l’interface qui traite le paquet.**

Si on active l’ACL sur **R1** pour les paquets entrants sur l’interface Gi1/0, R1 va **comparer chaque paquet entrant sur cette interface aux entrées de l’ACL** **afin de décider le sort de ce paquet** : continuer sans changement et acheminer le paquet ou le rejeter.

## III. La correspondance des paquets

C’est la manière dont on configure les ACL, on indique les informations qui seront utilisées par le processus de filtrage, cela peut être **une adresse IP ou bien un protocole réseau**. En se basant sur ces informations, le routeur déterminera s’il doit autoriser on refuser les paquets.

Le diagramme ci-dessous illustre mes propos :

Lors du traitement de l’ACL, le routeur traite le paquet comme suit :

**Les ACLs utilisent la logique de première correspondance.** Une fois qu’un paquet correspond à une ligne dans l’ACL, le routeur exécute cette règle puis le processus s’arrête.

L’exemple ci-dessous nous permettra de voir ce que cela signifie précisément :

Sur l'exemple ci-dessous, l'ACL est appliquée sur l’interface **Gi1/3** et elle contient trois règles :

- Si la source = 10.1.1.1, on autorise le paquet (*permit* )
- Si la source = 10.1.1.X, on refuse le paquet (*deny* )
- Si la source = 10.X.X.X, on autorise le paquet (*permit* )

On considère qu’un paquet est envoyé par le **PC1 (10.1.1.1)** au serveur **S1**, le routeur **R1** compare ce paquet à l’ACL correspondant à la première ligne de l'ACL, le paquet sera autorisé, car l'adresse IP du PC correspond à celle définie dans l'ACL.

Ensuite, considérons qu’un paquet est envoyé par le **PC2** **(10.1.1.2)** au serveur **S1.** À l’arrivée du paquet, le routeur va poursuivre la même logique de recherche en comparant le paquet à la première ligne de notre ACL. Il ne fera pas une correspondance, car 10.1.1.2 (adresse IP de PC2) n'est pas égal à l'adresse IP définie dans la première règle (10.1.1.1). Donc, R1 passe ensuite à la deuxième instruction à savoir "*10.1.1.x deny*", où **x** correspond à n’importe quelle valeur qui peut exister dans le dernier octet. En ne comparant que les trois premiers octets, R1 réalise que notre paquet émis depuis PC2 a une adresse IP source qui correspond à "10.1.1". De ce fait, R1 considère qu’il s’agit bien d’une correspondance à la deuxième règle de notre ACL et applique l’action comme indiqué dans la règle, à savoir refuser le paquet (*deny*). **Le routeur R1 arrête également le traitement ACL sur ce paquet, en ignorant la troisième règle de l’ACL.**


**Si jamais un paquet ne correspond à aucune des règles de la liste ACLs, le paquet est rejeté. On parle d'un refus implicite.**

Il est important de savoir que le traitement des ACL en mode séquentiel, en lisant les règles dans l'ordre, s’applique sur tout type d’ACL.

## IV. Les types d'ACLs

Il existe plusieurs types d'ACL, que nous allons découvrir ensemble sans plus attendre.

### A. Les ACL standards

Dans ce type, l’ACL ne peut être liée qu'à l’adresse IP source du paquet. Ces ACLs sont identifiables par identifiant correspondant à un nombre allant de **1 à 99** et de **1300 à 1999**.

Nous pourrons utiliser ce type d'ACL pour autoriser ou interdire un segment du réseau ou l'adresse IP d’une machine à communiquer avec un autre segment de réseau ou une autre machine.

Afin de concrétiser la notion d’ACL standard, nous allons les mettre en place sur un routeur Cisco.

Dans l’exemple ci-dessus, les trois ordinateurs, situés sur des segments réseau différents, communiquent entre eux. Le but est d’interdire au réseau « **10.1.1.0/24 »** de communiquer avec le réseau « **10.1.2.0/24 »** tout en ayant la possibilité de communiquer avec le réseau « **10.1.3.0/24 »** pour ce faire, nous allons, sur l’interface **Gi0/2**, interdire les paquets provenant du réseau « **10.1.1.0/24 »**


1 - Après s’être connecté sur le routeur en mode de configuration globale en tapant les commandes "*enable"* puis "*configuration terminal",* on commence par la création de la règle :

Router(config)#access-list 1 deny 10.1.2.0 0.0.0.255

En précisant "access-list 1" on attribue **un ID à notre ACL**, puis ensuite on précise que l'on veut refuser avec "deny", et enfin on précise l'adresse IP de destination (10.1.2.0) et le masque au format inversé appelé *wildcards mask* (0.0.0.225).

