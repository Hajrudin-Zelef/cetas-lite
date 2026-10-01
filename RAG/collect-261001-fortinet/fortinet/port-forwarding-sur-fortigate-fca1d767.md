---
id: collect-261001-fortinet/fortinet/port-forwarding-sur-fortigate-fca1d767
title: "port-forwarding-sur-fortigate-fca1d767"
domain: fortinet
role: reference
task: reference
actors: ["Apple", "Microsoft", "United States"]
dates: []
keywords: ["acquisition", "attention", "sol"]
source: docs/RAG/collect-261001-fortinet/port-forwarding-sur-fortigate-fca1d767.md
source_anchor: ""
source_lines: [1, 71]
sha256: fa0f97923f702e852d82246cdfbd6fdc74c29e1331662d6d7fad9d762b11cd91
---

# port-forwarding-sur-fortigate-fca1d767

Port forwarding sur Fortigate
I. Présentation
Comme ceci est un premier article sur la gamme de produits Fortigate, je vais vous parler un peu du produit. Le Fortigate (de la compagnie Fortinet) est un UTM (Unified Threat Management) ou en français, Gestion unifiée des menaces.
Mais c’est quoi ça? Ok, alors je vais être plus précis, un UTM rassemble dans un boitier (Appliance) ou une machine virtuelle plusieurs fonctions reliées à la sécurité. Avec en général les fonctions suivantes :
- Pare-feu (Firewall)
- Filtrage antipourriel (antispam)
- Antivirus
- Système de détection ou de prévention d'intrusion (IDS ou IPS)
- Filtrage de contenu applicatif (filtrage URL)
- VPN (SSL ou IPSEC)
Il existe plusieurs compagnies qui fabriquent des UTM et c’est très populaire, surtout auprès des petites et moyennes entreprises.
La compagne Fortinet est classée parmi les leaders de l’industrie depuis plusieurs années consécutives par Gartner dans la catégorie UTM.
II. Autres produits
Hormis les Fortigate que la compagnie construit, Fortinet possède aussi d’autres produits. Les plus populaires sont bien le Fortigate, mais aussi le FortiMail qui est l'antispam et le Fortimanager qui permet de gérer nos équipements. À noter aussi qu'il existe un client gratuit, le Forticlient qui nous sert de client VPN, firewall et antivirus.
Il est à noter que tous les Fortigate utilisent le même système d’exploitation (FortiOS) mais seul le CPU, la mémoire ou l’espace disque changent.
Le prix du produit commence à environ 338 $ US pour un Fortigate 30 D à plus de 100 000 $US pour les modèles très haut de gamme (Fortigate 5101C). Le prix varie aussi selon le temps de support, etc.
L’achat d’un Fortigate se fait à travers un revendeur que l’on peut trouver sur cette page : ici
III. Infrastructure
Aujourd'hui je voudrais vous présenter la façon de faire une redirection de port (ou port following) afin de permet la publication d’un serveur (WEB, d'email, FTP, etc..) derrière un Fortigate.
Dans cette présentation, je vais faire la publication d'un serveur d'email (port 25) afin de recevoir des emails de l'extérieur.
Attention l’adresse IP 175.20.20.15 est fictive.
IV. Configuration du serveur
Pour faire ceci, il bien sûr se connecter sur notre Fortigate via l’interface web de gestion.
La première étape va être de déclarer l’adresse IP de notre serveur interne (serveur d’email)
Sur le Fortigate, aller dans l’onglet Firewall objets et création d’une IP virtuelle :
Cliquez sur "Create new"
Puis nous donnons l’adresse IP publique redirigée vers le serveur d’email avec le port externe et le port vers l’interne (dans ce cas, le port 25 externe vers le port 25 vers l’interne)
Le nom de notre nouvelle adresse IP virtuelle va nous servir à la création de la règle de pare-feu.
Ce qui nous donne ceci comme résultat :
Une fois l’adresse IP virtuelle créée, la seconde étape va être de créer une règle de pare-feu (firewall) pour autoriser le trafic.
Pour ceci, allez dans la partie de droite, policy – règles et création d’une nouvelle règle.
Comme la règle ne concerne pas les VPN, nous choisissons Pare-feu
C’est une règle d’adresse IP seulement, il n’y a pas d’identification, nous choisissons "Adresse". La règle permet la communication de l’extérieur, donc nous choisissons le port relié à Internet (port WAN1 dans notre cas)
Vers le port ou est relier notre serveur d’email (dans notre cas le port 2) avec comme adresse de destination, le nom de notre adresse IP virtuelle
- Le service concerné (SMTP pour les emails entrant).
- Et bien sûr autoriser le trafic (accept) puisque c’est l’action désirée.
- Surtout ne pas activer le NAT.
Le reste de la configuration est l’activation des Logs (si désiré) et l’activation de l’antivirus et de l’antispam. Mais ses options seront vu dans un prochain article.
V. Conclusion
Voilà notre serveur d’email est publié sur Internet et prêt à recevoir ses premiers courriels. Il ne vous reste plus qu’à tester le tout par l’envoi d’un email de l’extérieur vers notre serveur interne. En cas d’erreur, il faudra aller voir les logs, mais cela fera partie d’un autre tutoriel.
bonjour merci pour l’article, est-ce qu’il y a des licence spécifique pour l’antivirus et de l’antispam et le web feltering ?
Bonjour Abdelati
Lors de l’achat de ton Fortigate, tu as part défaut la mise a jour de l’antivirus et Antispam et filtrage web pour 1 an (tu as l’option de payer pour plus longtemps)
Apres c’est effectivement une licence a renouveler . Le prix dépend du modèle de Fortigate que tu possède (plus le Fortigate est gros, plus c’est cher.
Si ta licence expire, l’appareil fonctionne encore mais les mises a jour ne ce font plus simplement.
Sur ce site, tu peux voir un exemple de prix (en dollars US) . Ça te donnera une idée, mais la politique de prix peux changer selon le pays
http://www.avfirewalls.com/products.asp
Chouette article, mais il me semble que l’on dit « port forwarding » et non « following »
Bon article, je rajouterai à la fin pour les utilisateurs qui ne connaissent pas bien Fortigate que son gros défaut c’est qu’il ne sait pas faire du reverse proxy et que la seule solution est le port forwarding :/.
Ce qui peut être un inconvénient quand on a plusieurs sites web à publier, c’est pas forcément le top d’accéder à un serveur ou un service en devant rajouter un numéro de port à la fin d’une url 🙂
Vous avez tout à fait raison, le Fortigate ne sait pas faire de reverse proxy, Mais a ma connaissance, il n’existe que deux produits qui savent le faire (corrigez-moi si je me trompe) L’ancien TMG de Microsoft ou le Sophos UTM.
Pour ceux qui ne connaissent pas Sophos, voici une vidéo qui présente le produit:
https://www.frbmg.com/utm-sophos/
Bonjour,
Ça vous prends un Fortiweb derrière le Fortigate pour effectuer cela;
http://docs-legacy.fortinet.com/fweb/admin_hlp/4-4-6/index.html#page/FortiWeb%20Help/topology_for_reverse_proxy_mode.html
Le Fortigate peut lui alléger le fardeau des serveurs web en arrière en mettant en cache une partie des requètes venant des visiteurs (web caching).
Bonjour Kevin
Tu as tout a fait raison. Le FortiWeb peux faire aussi ce travail. Mais ca sortais un peu du cadre de mon article 🙂 et c’est surtout un Appliance de plus.
Mais bien vue.
Bonjour Francis,
Désolé de répondre aussi tardivement je n’étais pas revenu sur votre article depuis mon dernier commentaire
Pour information la gamme Sonicwall de chez Dell fait du reverse proxy, mais je ne suis pas à l’aise avec l’interface, je préfère les produits Fortinet
De mon coté nous venons de faire l’acquisition du Fortiweb pour compléter notre solution Fortigate mais je n’ai pas encore eu le temps de la mettre en place
Bonjour, qu’en est il des licences VPN (ipsec/ssl), elles sont toutes activées par défaut, ou y’a seulement un nombre d’utilisateur qui est activé par défaut?
Bonjour Hakimmani
Avec les Fortigate, les licences sont assez facile. Toutes les options sauf les filtrages sont incluses. Tu n’as pas de limite de fonction ou d’utilisateurs (mais cela dépend de la taille de ton Fortigate)
La partie payante est la mise à jour des définitions (anti-virus, ips, et..) et le filtrage web. Toute tu n’es pas vraiment de limite logiciel pour le VPN. Sauf la taille de la boite va faire la différence.
Bonne journée
Francis
Bonjour la communauté, je suis très impressionné par les tutoriels it-connect. je vais un tutoriel qui me permet d’apprendre très rapidement les config de sur fortigate. merci de me joindre sur armandlongo@gmail.com
