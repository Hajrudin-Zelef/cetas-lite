---
id: collect-261001-cisco/cisco/configurer-le-port-mirroring-sous-cisco-commutation-routage-it-connect
title: "configurer-le-port-mirroring-sous-cisco-commutation-routage-it-connect"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/configurer-le-port-mirroring-sous-cisco-commutation-routage-it-connect.md
source_anchor: ""
source_lines: [1, 107]
sha256: b7c688f23ac7cfca9761575524e1eb58892fa3569066979ecc0188966626c96b
---

# configurer-le-port-mirroring-sous-cisco-commutation-routage-it-connect

## I. Présentation

Le *port-mirroring* est une technique qui permet de copier tout le trafic d'un ou de plusieurs port(s) source(s) vers un port de destination.

Cela peut s'avérer très utile pour analyser les flux réseaux qui transitent sur certains ports ou sur un VLAN, notamment dans une phase de dépannage.

Ainsi, pas besoin de commencer à positionner un hub entre le commutateur et l'équipement qui vous intéresse. Grâce à cette méthode, le trafic sera dupliqué et vous pourrez le récupérer et l'analyser directement depuis un outil comme **Wireshark** ou encore **TcpDump**.

Le *port-mirroring* est également appelé SPAN : **Switched Port Analyzer**

## II. Port-mirroring sur port source

Commençons par mettre en place le *port-mirroring* sur un port source, en dupliquant le trafic vers un port de destination.

Commencez par passer en mode **privilégié** puis en mode **configure terminal** :

Switch>enable
password:
Switch#configure terminal

On choisit ensuite le port source **fastEthernet 0/1** et on copie aussi bien les paquets reçus que transmis (**both** - les deux). Si l'on souhaite prendre que les paquets reçus on précisera **rx**, pour ceux transmis ce sera **tx**.

Switch(config)#monitor session 1 source interface fa 0/1 both

L'exemple ci-dessus configure la **session n°1**, il est possible de configurer plusieurs sessions que vous pouvez gérer grâce à la numérotation. Si l'on veut prendre plusieurs ports sources, on réitère la commande en modifiant l'interface source, si l'on garde le même numéro de session cela sera ajouté aux paramètres en cours (pas de remplacement).

Pour finir la configuration, on indique vers quel port de destination on souhaite dupliquer le trafic en provenance du port source. Exemple avec **fastEthernet 0/2**.

Switch(config)#monitor session 1 destination interface fa0/2

Concernant la commande monitor session, voici sa syntaxe complète :

```
monitor session session_number
{destination
   {interface interface_id[,|-]
   [encapsulation {dot1q}]
   [ingress vlan vlan_id]
   } |
   remote vlan vlan_id reflector-port interface_id
} |
{source
  {interface interface_id[,|-][both|tx|rx]} |
  remote vlan vlan_id
}
```
Vérifions que la configuration est bien prise en compte :

Switch#show monitor session 1

Pour ma part, j'ai ensuite utilisé le logiciel Wireshark sur une machine cliente connectée au port de destination (**fastEthernet 0/2**) afin de récupérer le trafic. Mon hôte dispose de l'adresse IP **169.254.23.168** (adresse APIPA), on remarque que les pings entre les hôtes **192.168.1.1** et **192.168.1.2** sont bien visibles sur cette analyse. Le trafic est bien dupliqué.

Si vous remarquez que votre carté réseau ne tient pas compte du trafic dupliqué qui n'est donc pas adressé directement à votre machine, vous devez activer le mode promiscious sur votre carte réseau. WireShark active ce mode sur la carte qu'il utilise, mais la modification peut être requise manuellement avec d'autres logiciels.

## III. Port-mirroring sur un VLAN

Même principe de configuration que le cas précédent à la différence que là on duplique tout le trafic d'un VLAN en particulier. Par exemple, on prend tout le trafic du VLAN en source à destination du port fastEthernet 0/2.

Switch(config)#monitor session 2 source vlan 10 both

La commande de destination reste la même (sauf numéro de session) :

Switch(config)#monitor session 1 destination interface fa0/2

## IV. Supprimer une session port-mirroring

La suppression d'une session s'effectue grâce à son numéro, en précédant la commande par "no" comme on le fait généralement sous Cisco.

Switch(config)#no monitor session 1

La configuration du *port-mirroring* est désormais terminée et opérationnelle.

Bonjour,

Félicitations pour le partage. Pour un trafic SIP, en procédant à la technique du port morroring

je pense qu’on aura aussi bien la signalisation que les flux RTP. Or, un opérateur n’acceptera pas dans le cadre d’un contrôle du trafic qu’on miroite la voix vers un serveur quelconque.

ou bien alors il faudra un équipement intermédiaire qui filtre seulement la signalisation.

Bonne réception!

Bonjour

J’essaie de réaliser un port mirroring avec votre tuto mais impossible car la commande monitor est toujours refus

Avez vous un exemple sous packet tracer

merci d’avance

cordialement

Configurer le port-mirroring sous Cisco http://www.it-connect.fr/configurer-le-port-mirroring-sous-cisco/

bonjour,bon commentaire. Jai une préoccupation particulière aimerais que vous puissiez m’apporter quelque solutions.

aimerais savoir quelles genre d’outils Cisco pourrais je utilisé pour dubliquer un reseau et de voir en temps réel ce qui ce passe c’est a dire j’ai un organisme A qui gere mon réseau de fibre optique et me fais le des rapport trimestrielle je doute pas de sa sincérité mais aimerais juste etre au courant de l’activité sans toute fois intervenir puisqu’il a concession de l’ex^loitation du reseau.

esperant que vous pourriez bien m aider merci a vous

merci pour ce tuto qui est très simple à comprendre. puis je avoir un tuto sur VTP (VLAN Trunking Protocole)?

Bonjour,

peut on faire du port mirorring sur 2 switchs différents ? c’est à dire par exemple capturer le traffic du port 24 du switch 1 pour le dupliquer et l’envoyer sur un laptop qui serait sur le port xx du switch 2 ( sachant que les switchs sont liés entre eux via un port channel en trunk, et sur un vlan identique?

Et sinon, comment capturer les échanges entre 2 téléphones IP sur le même réseau? Pour du trouble shooting, je voudrais voir ce qui se passe car ça semble marcher que dans un sens
