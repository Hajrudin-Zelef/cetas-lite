---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-smb-switches-cisco-350-series-managed-switches-smb5673-conf-d2c2c77a-1
title: "c-fr-ca-support-docs-smb-switches-cisco-350-series-managed-switches-smb5673-conf-d2c2c77a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution", "diffusion", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-smb-switches-cisco-350-series-managed-switches-smb5673-conf-d2c2c77a.md
source_anchor: ""
source_lines: [1, 44]
sha256: 0911cb664b0546e6842d84f92d58c94715185a5490e1051fae9d3131a2f837ca
---

# c-fr-ca-support-docs-smb-switches-cisco-350-series-managed-switches-smb5673-conf-d2c2c77a

Un réseau local virtuel (VLAN) vous permet de segmenter logiquement un réseau local (LAN) en différents domaines de diffusion. Dans les scénarios où des données sensibles peuvent être diffusées sur un réseau, des VLAN peuvent être créés pour améliorer la sécurité en désignant une diffusion à un VLAN spécifique. Seuls les utilisateurs appartenant à un VLAN peuvent accéder aux données de ce VLAN et les manipuler. Les VLAN peuvent également être utilisés pour améliorer les performances en réduisant la nécessité d’envoyer des diffusions et des multidiffusions vers des destinations inutiles.
Note: Pour savoir comment configurer les paramètres VLAN de votre commutateur via l'utilitaire Web, cliquez ici. Pour obtenir des instructions basées sur l'interface de ligne de commande, cliquez ici.
Les périphériques réseau sur lesquels plusieurs protocoles s’exécutent ne peuvent pas être regroupés dans un VLAN commun. Les périphériques non standard sont utilisés pour transmettre le trafic entre différents VLAN afin d’inclure les périphériques participant à un protocole spécifique. Pour cette raison, vous ne pouvez pas tirer parti des nombreuses fonctionnalités du VLAN.
Les groupes de VLAN sont utilisés pour équilibrer la charge du trafic sur un réseau de couche 2. Les paquets sont distribués par rapport à différentes classifications et sont affectés aux VLAN. Il existe de nombreuses classifications différentes, et si plusieurs schémas de classification sont définis, les paquets sont affectés au VLAN dans l'ordre suivant :
1. Créez les VLAN. Pour savoir comment configurer les paramètres VLAN de votre commutateur via l'utilitaire Web, cliquez ici. Pour obtenir des instructions basées sur l'interface de ligne de commande, cliquez ici.
2. Configurer les interfaces avec les VLAN Pour obtenir des instructions sur l'attribution d'interfaces aux VLAN via l'utilitaire Web de votre commutateur, cliquez ici. Pour obtenir des instructions basées sur l'interface de ligne de commande, cliquez ici.
Note: Si l'interface n'appartient pas au VLAN, le paramètre de configuration des groupes de sous-réseaux vers VLAN ne prend pas en compte.
3. Configurez des groupes VLAN basés sur des protocoles. Pour obtenir des instructions sur la configuration des groupes de VLAN basés sur des protocoles via l'utilitaire Web de votre commutateur, cliquez ici.
4. (Facultatif) Vous pouvez également configurer les éléments suivants :
Des groupes de protocoles peuvent être définis, puis liés à un port. Une fois que le groupe de protocoles est lié à un port, chaque paquet provenant d'un protocole du groupe se voit attribuer un VLAN configuré dans les groupes basés sur le protocole.
Le transfert de paquets en fonction de leur protocole nécessite la configuration de groupes de protocoles, puis le mappage de ces groupes sur des VLAN. Cet article explique comment définir des groupes de protocoles et configurer des groupes basés sur des protocoles sur VLAN.
Étape 1. Connectez-vous à la console du commutateur. Le nom d'utilisateur et le mot de passe par défaut sont cisco/cisco. Si vous avez configuré un nouveau nom d'utilisateur ou mot de passe, saisissez plutôt les informations d'identification.
Note: Les commandes peuvent varier en fonction du modèle exact de votre commutateur. Dans cet exemple, le commutateur SG350X est accessible via Telnet.
Étape 2. À partir du mode d’exécution privilégié du commutateur, passez en mode de configuration globale en entrant ce qui suit :
Étape 3. En mode de configuration globale, configurez une règle de classification basée sur un protocole en entrant ce qui suit :
Étape 4. Pour mapper un protocole à un groupe de protocoles, saisissez ce qui suit :
Les options sont les suivantes :
- IP : trame Ethernet V2 comportant un paquet IPv4. Le numéro de protocole est 0x0800.
- IPX : trame Ethernet V2 avec IPX (Internetwork Packet Exchange). Les numéros de protocole vont de 0x8137 à 0x8138.
- IPv6 : trame Ethernet V2 qui possède un paquet IPv6. Le numéro de protocole est 0x86DD.
- ARP : trame Ethernet V2 avec un paquet ARP (Address Resolution Protocol). Le numéro de protocole est 0x0806.
- Défini par l'utilisateur : vous pouvez saisir une valeur de protocole en hexadécimal, avec une longueur de quatre chiffres.
- Ethernet : ce paramètre fait référence au paquet de données sur une liaison Ethernet. Il s’agit de l’encapsulation par défaut. Si la valeur d’encapsulation n’est pas définie, Ethernet sera utilisé comme type d’encapsulation.
- rfc1042 - Ce paramètre fait référence au contrôle de liaison logique avec le protocole LLC-SNAP (Sub-Network Access Protocol). Ces protocoles fonctionnent ensemble pour garantir que les données sont transmises efficacement à l’intérieur du réseau.
- llcother : ce paramètre fait référence à la méthode LLC (Logical Link Control). Il s’agit de la sous-couche de la couche liaison de données, qui agit comme une interface entre la sous-couche de contrôle d’accès au support et la couche réseau.
Note: Dans cet exemple, les groupes VLAN basés sur des protocoles 100 et 200 sont créés. Le groupe 100 filtre le protocole Ethernet IP, tandis que le groupe 200 filtre le protocole Ethernet IPv6.
Étape 5. Pour quitter le contexte de configuration d'interface, saisissez ce qui suit :
Vous devez maintenant avoir configuré les groupes VLAN basés sur des protocoles sur votre commutateur via l'interface de ligne de commande.
Étape 1. En mode de configuration globale, entrez le contexte de configuration d'interface en saisissant ce qui suit :
Note: Dans cet exemple, l'interface ge1/0/20 est utilisée.
Étape 2. Dans le contexte de configuration d'interface, utilisez la commande switchport mode pour configurer le mode d'appartenance VLAN :
Étape 3. (Facultatif) Pour rétablir le port sur le VLAN par défaut, saisissez ce qui suit :
Étape 4. Pour configurer une règle de classification basée sur un protocole, saisissez ce qui suit :
Note: Dans cet exemple, l'interface est affectée au groupe basé sur le protocole 100 qui est mappé au VLAN 20.
Étape 6. (Facultatif) Pour supprimer la règle de classification du port ou de la plage de ports, saisissez ce qui suit :
Étape 7. (Facultatif) Répétez les étapes 1 à 6 pour configurer des ports plus généraux et affecter les groupes de VLAN basés sur le protocole correspondants.
Note: Dans cet exemple, les interfaces qui vont de ge1/0/31 à 35 sont affectées au groupe basé sur le protocole 200 et affectées au VLAN 30.
Étape 8. Entrez la commande end pour revenir au mode d’exécution privilégié :
Vous devez maintenant avoir mappé des groupes de VLAN basés sur des protocoles aux VLAN de votre commutateur via l'interface de ligne de commande.
Étape 1. Pour afficher les protocoles qui appartiennent aux règles de classification basées sur les protocoles définies, entrez ce qui suit en mode d’exécution privilégié :
Étape 2. (Facultatif) Pour afficher les règles de classification d'un port spécifique sur le VLAN, saisissez ce qui suit :
Note: Chaque mode de port a sa propre configuration privée. La commande show interfaces switchport affiche toutes ces configurations, mais seule la configuration du mode de port qui correspond au mode de port actuel affiché dans la zone Mode d'administration est active.
Note: Dans cet exemple, les états d'administration et d'exploitation de l'interface ge1/0/20 sont affichés. Le tableau des règles de classification indique que l'interface a été mappée à un groupe VLAN 100 basé sur le protocole et que le trafic sera transféré au VLAN 20.
Étape 3. (Facultatif) Dans le mode d’exécution privilégié du commutateur, enregistrez les paramètres configurés dans le fichier de configuration initiale, en saisissant ce qui suit :
