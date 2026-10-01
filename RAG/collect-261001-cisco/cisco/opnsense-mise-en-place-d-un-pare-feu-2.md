---
id: collect-261001-cisco/cisco/opnsense-mise-en-place-d-un-pare-feu-2
title: "opnsense-mise-en-place-d-un-pare-feu"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "open source"]
source: docs/RAG/collect-261001-cisco/opnsense-mise-en-place-d-un-pare-feu.md
source_anchor: ""
source_lines: [213, 324]
sha256: d39788e749ea6e3bdfa725108da7eff642d055b643bddb99f8c6f332144e03bc
---

# opnsense-mise-en-place-d-un-pare-feu

Remplir tous les champs obligatoire 1 et cliquer sur le bouton Sauvegarder 2.

L’autorité est créée.

#### Configuration du Proxy Web

On va maintenant passer à la configuration du Proxy.

Pour être franc, ce OPNsense ce n’est pas une partie de plaisir par rapport à d’autre solution …


On va commencer par activer le Proxy Web, depuis le menu aller sur Services / Proxy Web / Administration.

Sur l’onglet Réglages Proxy généraux, cocher la case Activer le proxy 1.

Pour accéder à plus d’option, activé le mode avancé 1.

Aller ensuite sur l’onglet Forward Proxy 1, cocher les cases suivantes : Activer le proxy HTTP Transparent 2 et Activer l’inspection SSL 3. Sélectionner l’autorité de certificat 4 que l’on a créé précédemment puis valider les paramètres en cliquant sur Appliquer 5.

Sur les autres Firewall que j’ai l’habitude de manipuler, normalement la configuration est terminée mais pas sur OPNsense, il faut maintenant créer deux règles pour faire passer le trafic dans le proxy.


Pour la création des règles de redirection de trafic, on a de la chance, cela va faire « automatiquement ».

Cliquer sur l’icone d’information 1 du champ Activer le proxy HTTP Transparent puis sur le lien Add a new firewall rule 2.

Vous êtes redirigé vers une page de création de règle NAT, qui normalement est déjà configurée, cliquer sur le bouton Sauvegarder 1 en bas de page.

A la création de la règle, vous êtes redirigé sur la page qui affiche la liste des règles. Retourner sur la page de configuration du Proxy Web.

On va refaire la même manipulation pour l’inspection SSL, cliquer sur l’icone d’information 1 puis sur le lien Add new firewall rule 2.

Valider la création de la règle en cliquant sur le bouton Sauvegarder en bas de page.

Les deux règles sont créées 1. Cliquer sur le bouton Appliquer les changements 2.

Ce n’est pas fini, il faut désactiver la règle par défaut qui autorise tout le trafic vers Internet. Aller sur Pare-feu / Règles / LAN et cliquer sur l’icone vert 1 pour désactiver la désactiver.

On va maintenant créer une règle pour autorisé le trafic DNS, sinon pas de résolution nom, cliquer sur l’icone Ajouter 1.

Je ne vais pas rentrer dans le détail de la création de la règle, mais le but, c’est d’autorisé le trafic du LAN vers des serveurs DNS (53/UDP) qui sont sur Internet.

Voici la règle :

Cliquer ensuite sur Appliquer les changements 1.

#### Tester le proxy Web depuis un ordinateur

Lancer un navigateur Internet et essayer de surfer, si tout se passe bien, vous devriez avoir une erreur de certificat :

Cette erreur est normale car nous n’avons pas installé le certificat que l’on a créé, car le trafic HTTPS ayant était déchiffrer par le proxy et chiffré de nouveau avec son certificat, mais cela prouve le bon fonctionnement.

Exporter le certificat depuis l’interface Web :

Sur l’ordinateur client, installer le certificat au niveau de l’ordinateur dans le magasin Autorités de certification racines de confiance.

Pour un déploiement à l’échelle d’un parc informatique, le déploiement peut être fait par GPO : GPO : déployer un certificat.


Une fois le certificat installé, la page devrait s’afficher normalement :

Si on affiche le certificat du site Internet, on peut voir que celui-ci a été émis avec l’autorité de certification que l’on a créé dans OPNsense que l’on a défini pour le proxy.

## Conclusion

Ce premier tutoriel sur OPNsense s’arrête ici. D’autre devrait suivre bientôt.

Vous pouvez maintenant déployer OPNsense pour installer un firewall gratuit, mais qui reste très limité par rapport à des solutions « propriétaire »‘.

Bonjour,

Merci pour ce tutoriel intéressant !

Quand vous dites : « Vous pouvez maintenant déployer OPNsense pour installer un firewall gratuit, mais qui reste très limité par rapport à des solutions « propriétaire » », à quelles solutions pensez-vous ? J’avais l’impression qu’OPNsense était plutôt permissif/flexible ?!

Merci.

Bonjour Etienne,

Je pensais à des solutions comme Fortigate ou Firewall Sophos qui permettent des configurations plus poussées au niveau filtrage Web, règles de VPN et intégration natives de IPS / AV …

OPNsense comme la plupart des firewalls « gratuit » sont très limités en configuration par rapport à des solutions propriétaires.

Bonjour Romain,

Je me permets de vous appeler par votre prénom.

Je suis un peu mitigé sur la formulation : « reste très limité par rapport à des solutions ‘propriétaires’ » pour un produit open source et gratuit. Ce dernier propose des addons et remplit bien le cahier des charges.

J’en conviens, il est probable que certaines fonctionnalités soient moins développées que dans une solution payante. Cependant, il est essentiel de tout analyser : l’usage prévu, les connaissances techniques de la personne qui installe ou utilise le produit.

Pour ma part, je l’utilise dans un cadre privé et j’en suis très satisfait.

Je referme ici la parenthèse sur cette remarque, que je trouve un peu maladroite (et cela n’engage que moi). Cela n’enlève en rien à la qualité de l’article et de votre travail en général, que je trouve excellent.

J’espère que vous ne m’en tiendrez pas rigueur.

Pour ceux qui sont curieux, il existe une version « home » de Sophos. Pour en savoir plus, je vous invite à taper « Sophos Firewall Édition familiale » dans votre moteur de recherche préféré.

Cordialement,

John

Bonjour John,

Pas de soucis pour ton retour, comme tu le dis, il y a des addons qui permettent d’ajouter des fonctionnalités, comme souvent dans les produits gratuits ou basés sur du libre au niveau des pares-feu, ceci est développé gratuitement par des personnes lambas, qui souvent au bout d’un certain temps, passe à autre chose et on se retrouve avec des addons qui ne sont plus maintenu. Après si on prend l’aspect technique, je préfère encore une solution payante où j’aurais un contrat de maintenance et / ou je pourrais trouver un prestataire qualifié en cas de problème, pour moi encore un point de plus pour une solution propriétaire.

Pour moi ce type de firewall peut être très bien pour monter un vpn site à site où « sécurisé » un réseau domestique, mais comme tu le souligne, je préférai utilise Sophos Home qui est un très bon produit.

Cordialement,

Romain
