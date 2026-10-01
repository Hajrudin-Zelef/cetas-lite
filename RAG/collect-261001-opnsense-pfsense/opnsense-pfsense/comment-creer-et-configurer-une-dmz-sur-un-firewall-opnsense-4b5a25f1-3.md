---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/comment-creer-et-configurer-une-dmz-sur-un-firewall-opnsense-4b5a25f1-3
title: "Comment créer et configurer une DMZ sur un firewall OPNsense ?"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/comment-creer-et-configurer-une-dmz-sur-un-firewall-opnsense-4b5a25f1.md
source_anchor: ""
source_lines: [52, 57]
sha256: c1570e9d942e420e80abbce3b5025ded54685f54767eaf2ae613f3bf5348be8c
---

# Comment créer et configurer une DMZ sur un firewall OPNsense ?

Il est impératif de tester les règles mises en place pour vérifier qu'elles répondent à nos besoins. Ceci implique de tester une connexion sur le serveur Web depuis le LAN, mais aussi d'essayer de se connecter sur d'autres ports, ou depuis la DMZ vers le LAN. Il faut s'assurer que les autorisations sont opérationnelles, mais aussi les refus.
La mise en place du serveur Web n'est pas abordée dans ce tutoriel. Si besoin, référez-vous à ces articles :
Ajouter une DMZ à un firewall OPNsense et par extension à votre système d'information est une étape cruciale pour renforcer la sécurité de votre réseau, notamment pour isoler les services exposés à l'extérieur, sur Internet.
En suivant les étapes décrites dans ce tutoriel, vous avez appris à configurer une interface dédiée et à définir les règles de pare-feu appropriées. Il ne vous reste plus qu'à adapter cette configuration à vos besoins.
Pour ajouter une couche de sécurité supplémentaire d'un point de vue applicatif, sachez qu'OPNsense peut assurer le rôle de reverse proxy.
Et si vous avez une question, pensez à poster un commentaire sur cet article ou à venir en discuter sur notre Discord.
