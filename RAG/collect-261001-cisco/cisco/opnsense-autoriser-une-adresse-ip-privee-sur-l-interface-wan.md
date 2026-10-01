---
id: collect-261001-cisco/cisco/opnsense-autoriser-une-adresse-ip-privee-sur-l-interface-wan
title: "opnsense-autoriser-une-adresse-ip-privee-sur-l-interface-wan"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/opnsense-autoriser-une-adresse-ip-privee-sur-l-interface-wan.md
source_anchor: ""
source_lines: [1, 24]
sha256: 7ec7bb1b3e7ea97f0ab7593344ccd1451222dad8fb330a4f7e4929d9d38d08f9
---

# opnsense-autoriser-une-adresse-ip-privee-sur-l-interface-wan

## Introduction

Dans ce tutoriel, je vais vous expliquer comment autoriser une adresse IP privée sur l’interface WAN d’OPNsense.

Ce cas se présente fréquemment lorsqu’on installe un pare-feu OPNsense derrière une box opérateur ou un autre routeur. Dans ce type de configuration, l’interface WAN d’OPNsense ne reçoit pas directement une adresse publique, mais une adresse privée (par exemple 192.168.x.x). Par défaut, OPNsense considère cela comme une erreur de configuration et bloque l’interface WAN pour des raisons de sécurité.

Nous allons voir comment activer l’option permettant d’accepter ce type d’adresse et ainsi utiliser OPNsense dans un environnement où il est placé derrière un équipement déjà connecté à Internet.

## Autoriser une adresse IP privée sur l’interface WAN sur OPNsense

Depuis l’interface Web du Pare-feu, aller sur Interface / [WAN], qui permet d’afficher la configuration de celle-ci.

Comme on peut le voir ci-dessus, Bloquer les réseaux privés est activé par défaut sur l’interface WAN.


Décocher les cases : Bloquer les réseaux privés et Bloquer les adresses bogon (non attribuées par l’IANA) 1 et cliquer sur le bouton Sauvegarder 2 en bas de la page.

Cliquer sur Appliquer les changements 1 pour prendre en compte les modifications.

## Conclusion

Vous savez maintenant comment autoriser une adresse IP privée sur l’interface WAN d’OPNsense.

Cette option est indispensable lorsque le pare-feu est installé derrière une box opérateur ou un autre routeur, afin d’éviter le blocage par défaut du WAN. Vous pouvez ainsi déployer OPNsense dans ce type d’architecture sans rencontrer d’erreurs liées à l’utilisation d’adresses privées.
