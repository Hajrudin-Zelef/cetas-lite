---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/comment-ajouter-une-dmz-a-un-firewall-opnsense-2
title: "comment-ajouter-une-dmz-a-un-firewall-opnsense"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-opnsense-pfsense/comment-ajouter-une-dmz-a-un-firewall-opnsense.md
source_anchor: ""
source_lines: [103, 139]
sha256: 6f12b9e908a44ac50c8b98243b2e86405f1d71fd195e23abddc2c0db3e90a14f
---

# comment-ajouter-une-dmz-a-un-firewall-opnsense

**Ici, il n'y a qu'un seul réseau local, donc c'est simple**. S'il y a plusieurs réseaux, plutôt que de créer plusieurs règles, vous pouvez créer un groupe de réseaux via la **fonction d'Alias** (Aliases). Il suffira de sélectionner l'Alias par son nom lors de la création de la règle. Cette fonction se situe dans : **Pare-feu > Alias.**

**Remarque** : une autre façon de procéder, c'est de déclarer les réseaux privés des différentes classes (10.0.0.0/8 - 172.16.0.0/12 - 192.168.0.0/16), pour être certain d'isoler votre DMZ. Si vous gérez une liste manuelle, ceci vous oblige à la maintenir.


### C. Récapitulatif des règles

La création de règles est une action répétitive, bien que vous puissiez dupliquer les règles grâce au bouton mis en évidence sur l'image ci-dessous. Ceci vous permettra de gagner du temps (mais attention aux erreurs !).

Pour vous aiguiller, voici les règles que vous devez avoir sur les interfaces "**LAN**" et "**DMZ**". Dans ma configuration, le NAT sortant est en mode automatique, donc OPNsense a créé des règles de lui-même pour les réseaux correspondants à mon LAN et à ma DMZ.

- **Règles sur l'interface DMZ**

- **Règles sur l'interface LAN**

Vous pouvez consulter la documentation officielle en complément :

### D. Tester les règles

Il est impératif de tester les règles mises en place pour vérifier qu'elles répondent à nos besoins. Ceci implique de tester une connexion sur le serveur Web depuis le LAN, mais aussi d'essayer de se connecter sur d'autres ports, ou depuis la DMZ vers le LAN. Il faut s'assurer que les autorisations sont opérationnelles, mais aussi les refus.

La mise en place du serveur Web n'est pas abordée dans ce tutoriel. Si besoin, référez-vous à ces articles :

- Installer un serveur Web Apache2 sur Debian
- Installer un serveur Web Caddy sur Debian
- Installer un serveur Web Nginx sur Debian
- Installer un serveur Web IIS sur Windows Server

## V. Conclusion

Ajouter une DMZ à un firewall OPNsense et par extension à votre système d'information est une étape cruciale pour renforcer la sécurité de votre réseau, notamment pour isoler les services exposés à l'extérieur, sur Internet.

En suivant les étapes décrites dans ce tutoriel, vous avez appris à configurer une interface dédiée et à définir les règles de pare-feu appropriées. Il ne vous reste plus qu'à adapter cette configuration à vos besoins.

Pour ajouter une couche de sécurité supplémentaire d'un point de vue applicatif, sachez qu'OPNsense peut assurer le rôle de reverse proxy.

Et si vous avez une question, pensez à poster un commentaire sur cet article ou à venir en discuter sur notre Discord.
