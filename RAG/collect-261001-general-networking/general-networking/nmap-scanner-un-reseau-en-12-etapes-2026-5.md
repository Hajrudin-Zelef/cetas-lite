---
id: collect-261001-general-networking/general-networking/nmap-scanner-un-reseau-en-12-etapes-2026-5
title: "Nmap version 7.991 ( https://nmap.org )"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "open source"]
source: docs/RAG/collect-261001-general-networking/nmap-scanner-un-reseau-en-12-etapes-2026.md
source_anchor: ""
source_lines: [318, 371]
sha256: 1aee15416cb258f3aac6e64b81bc0bedd5639afdf4e0de69edc8e4287f303b37
---

# Nmap version 7.991 ( https://nmap.org )

- **« nmap: command not found »** après installation : la variable PATH n’a pas été rechargée. Fermez et rouvrez le terminal, ou lancez`hash -r` sous bash.
- **« You requested a scan type which requires root privileges »** : ajoutez`sudo` devant la commande sous Linux/macOS, ou lancez le terminal en administrateur sous Windows.
- **Le scan reste bloqué indéfiniment sur une IP** : ajoutez`--host-timeout 5m` pour forcer un abandon après cinq minutes et passer à l’hôte suivant.
- **Tous les ports apparaissent « filtered »** : un pare-feu intermédiaire bloque probablement les réponses RST. Essayez un scan ACK (`-sA` ) pour distinguer les règles de filtrage.
- **L’hôte semble éteint alors qu’il répond via navigateur** : le ping ICMP est bloqué. Utilisez`-Pn` pour forcer le scan sans phase de découverte.
- **Un script NSE dépasse le temps imparti** : ajoutez`--script-timeout 30s` pour limiter la durée de chaque script individuellement.
- **Le scan est anormalement lent sur un grand réseau** : augmentez le parallélisme avec`--min-parallelism 50` ou passez à un template de timing plus agressif (`-T4` ).
- **Détection de version incorrecte ou incomplète** : montez l’intensité des sondes avec`--version-intensity 9` , la valeur par défaut (7) n’interroge pas toujours les services les plus rares.

## Astuces avancées pour aller plus loin

Une fois les bases maîtrisées, plusieurs options méritent d’être explorées pour professionnaliser vos audits. La combinaison `-sV --version-intensity 9 --script default,vuln` donne un bon compromis entre exhaustivité et temps d’exécution pour un audit hebdomadaire planifié. Pensez aussi à exclure systématiquement les hôtes critiques d’un scan agressif avec `--exclude 192.168.1.1,192.168.1.254` pour éviter de perturber une passerelle ou un contrôleur industriel fragile. Pour les environnements IPv6, ajoutez simplement `-6` à votre commande : la couverture IPv6 de Nmap s’est nettement améliorée depuis l’introduction du résolveur DNS parallèle en version 7.96, qui a permis de résoudre un million de noms d’hôtes en environ une heure contre 49 heures auparavant selon le changelog officiel du projet.

Enfin, pour les équipes qui gèrent plusieurs segments réseau, envisagez de déployer un scanner Nmap par segment (via des agents Wazuh distincts par exemple), plutôt qu’un scan centralisé unique qui traverse des pare-feux internes et fausse les résultats de latence et de filtrage. Pour la documentation complète des options, référez-vous au manuel officiel de Nmap.

Dernier conseil, souvent négligé : versionnez vos scripts d’audit Nmap dans un dépôt Git privé, au même titre que du code de production. Un script d’automatisation qui scanne l’ensemble d’un système d’information mérite une revue de code, un historique de modifications et une procédure de restauration en cas d’erreur de configuration, exactement comme n’importe quel outil interne critique.

## FAQ : vos questions sur Nmap

**Nmap est-il gratuit ?**

Oui, Nmap est un logiciel open source distribué sous une licence propre inspirée de la GPL, entièrement gratuit pour un usage personnel comme professionnel. Le code source est disponible sur nmap.org.

**Puis-je utiliser Nmap sans les droits administrateur ?**

Oui pour un scan Connect (-sT) ou un scan de découverte simple, mais les scans SYN (-sS) et la détection d’OS (-O) nécessitent des privilèges élevés pour manipuler les paquets réseau bruts.

**Quelle est la différence entre Nmap et Zenmap ?**

Zenmap est simplement l’interface graphique officielle de Nmap. Il exécute les mêmes commandes en arrière-plan mais facilite la visualisation des résultats et la sauvegarde de profils de scan.

**Scanner ma propre box internet est-il légal ?**

Oui, scanner votre propre réseau domestique ou une infrastructure que vous possédez ne pose aucun problème légal en France. La question se pose uniquement pour des systèmes appartenant à des tiers.

**Nmap peut-il détecter tous les appareils IoT sur mon réseau ?**

Dans la majorité des cas oui, via le ping scan (-sn) et la détection de service, mais certains objets connectés à très faible consommation ne répondent qu’à des sondes spécifiques et peuvent nécessiter des scripts NSE dédiés.

**Combien de temps prend un scan complet d’un réseau /24 ?**

Avec un template -T4 et les ports par défaut, comptez entre 2 et 10 minutes selon le nombre d’hôtes actifs. Un scan complet des 65 535 ports par hôte peut prendre plusieurs heures.

**Nmap peut-il être détecté par un pare-feu ou un antivirus ?**

Oui, la plupart des solutions de sécurité réseau modernes, dont Wazuh, Suricata ou les pare-feux nouvelle génération, détectent les signatures de scan Nmap et peuvent déclencher une alerte ou un blocage automatique de l’adresse IP source.

**Faut-il préférer Nmap ou un scanner de vulnérabilités comme OpenVAS ?**

Ce ne sont pas des concurrents directs. Nmap excelle en découverte et cartographie réseau rapide, tandis qu’OpenVAS ou Greenbone conviennent mieux à un audit de conformité formel avec base CVE intégrée. Beaucoup d’équipes utilisent les deux, l’un après l’autre.

**Nmap fonctionne-t-il sur un réseau IPv6 ?**

Oui, il suffit d’ajouter l’option `-6` à votre commande. La prise en charge d’IPv6 s’est nettement améliorée ces dernières versions, notamment grâce au nouveau résolveur DNS parallèle introduit en 7.96, mais certaines techniques de scan (comme la détection d’OS) restent moins fiables qu’en IPv4.
