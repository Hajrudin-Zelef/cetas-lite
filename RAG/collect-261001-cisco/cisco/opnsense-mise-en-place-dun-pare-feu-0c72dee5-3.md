---
id: collect-261001-cisco/cisco/opnsense-mise-en-place-dun-pare-feu-0c72dee5-3
title: "OPNsense : mise en place d’un pare-feu"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-cisco/opnsense-mise-en-place-dun-pare-feu-0c72dee5.md
source_anchor: ""
source_lines: [90, 110]
sha256: 4a7f614c4db2755c980248b342ae06e04716be0f9b8e21d5658c11345624aee7
---

# OPNsense : mise en place d’un pare-feu

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
Lancer un navigateur Internet et essayer de surfer, si tout se passe bien, vous devriez avoir une erreur de certificat :
Cette erreur est normale car nous n’avons pas installé le certificat que l’on a créé, car le trafic HTTPS ayant était déchiffrer par le proxy et chiffré de nouveau avec son certificat, mais cela prouve le bon fonctionnement.
Exporter le certificat depuis l’interface Web :
Sur l’ordinateur client, installer le certificat au niveau de l’ordinateur dans le magasin Autorités de certification racines de confiance.
Pour un déploiement à l’échelle d’un parc informatique, le déploiement peut être fait par GPO : GPO : déployer un certificat.
Une fois le certificat installé, la page devrait s’afficher normalement :
Si on affiche le certificat du site Internet, on peut voir que celui-ci a été émis avec l’autorité de certification que l’on a créé dans OPNsense que l’on a défini pour le proxy.
Ce premier tutoriel sur OPNsense s’arrête ici. D’autre devrait suivre bientôt.
Vous pouvez maintenant déployer OPNsense pour installer un firewall gratuit, mais qui reste très limité par rapport à des solutions « propriétaire »‘.
