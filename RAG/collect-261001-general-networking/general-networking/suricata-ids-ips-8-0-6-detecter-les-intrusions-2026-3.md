---
id: collect-261001-general-networking/general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026-3
title: "repérez le nom de votre interface, par exemple eth0 ou ens18"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2026-30-08"]
keywords: ["agent", "open source"]
source: docs/RAG/collect-261001-general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026.md
source_anchor: ""
source_lines: [168, 275]
sha256: fcc7b322eb79fb04331e6d654866c5e81ae78e526f59cf61db2a6c65c95954d2
---

# repérez le nom de votre interface, par exemple eth0 ou ens18

`sudo suricata -T -c /etc/suricata/suricata.yaml -v`
Un message `Configuration provided was successfully loaded` confirme que tout est prêt. Si une erreur apparaît, elle indique généralement le numéro de ligne du fichier YAML ou de la règle fautive, ce qui facilite grandement le débogage.

## Étape 8 : Démarrer le service et vérifier son fonctionnement

```
sudo systemctl enable suricata
sudo systemctl start suricata
sudo systemctl status suricata
```
Vérifiez que le processus capture bien du trafic en consultant les statistiques internes, mises à jour toutes les 8 secondes par défaut dans le fichier `stats.log` :

`sudo tail -f /var/log/suricata/stats.log | grep -E "capture.kernel_packets|decoder.pkts"`
Exemple de sortie attendue si la capture fonctionne correctement :

```
capture.kernel_packets       | Total                     | 128453
decoder.pkts                 | Total                     | 128453
decoder.bytes                | Total                     | 98234112
flow.total                   | Total                     | 3021
```
Un compteur `capture.kernel_drops` élevé signale que Suricata ne parvient pas à traiter tout le trafic en temps réel, généralement parce que le nombre de threads est insuffisant ou que le CPU sature. C’est le signal qu’il faut revoir le dimensionnement matériel ou ajuster les paramètres de capture.

## Étape 9 : Déclencher et lire une alerte de test

Pour vérifier que la chaîne de détection fonctionne de bout en bout, utilisez la règle de test officielle intégrée à Emerging Threats, qui se déclenche sur une requête HTTP contenant une chaîne spécifique inoffensive.

`curl http://testmynids.org/uid/index.html`
Consultez ensuite le fichier d’alertes au format lisible :

`sudo tail -20 /var/log/suricata/fast.log`
Vous devriez voir une ligne du type :

`08/30/2026-11:42:03.128341 [**] [1:2100498:7] GPL ATTACK_RESPONSE id check returned root [**] [Classification: Potentially Bad Traffic] [Priority: 2] {TCP} 192.168.1.50:51422 -> 217.160.0.1:80`
Si cette ligne apparaît, votre pipeline de détection est opérationnel de bout en bout : capture réseau, moteur de règles et journalisation fonctionnent ensemble.

## Optimiser les performances : threads, mémoire et capture

Sur un lien réseau chargé, la configuration par défaut de Suricata montre vite ses limites. Le paramètre `threads: auto` dans la section af-packet laisse Suricata déterminer lui-même le nombre de threads de capture en fonction des cœurs disponibles, mais sur un serveur partagé avec d’autres services, il est souvent préférable de fixer une valeur explicite pour ne pas monopoliser tout le CPU.

```
af-packet:
  - interface: eth0
    threads: 6
    ring-size: 200000
    block-size: 1048576
```
Le paramètre `ring-size` définit le nombre de paquets que le noyau peut mettre en mémoire tampon avant que Suricata ne les traite : une valeur trop basse provoque des pertes de paquets sous forte charge, une valeur trop haute consomme de la RAM inutilement. Sur un serveur avec 8 Go de RAM dédiés à Suricata, une valeur entre 100 000 et 200 000 constitue un bon point de départ, à ajuster ensuite en observant le compteur `capture.kernel_drops` dans les statistiques.

Pensez également à ajuster la limite de mémoire allouée au flow tracking, qui garde en mémoire l’état de chaque connexion réseau active. Sur un réseau avec beaucoup de connexions simultanées (proxy, passerelle NAT), cette limite peut être atteinte rapidement et provoquer l’abandon prématuré du suivi de certains flux.

```
flow:
  memcap: 512mb
  hash-size: 65536
  prealloc: 10000
```
Après chaque changement de configuration, relancez le test de validation puis surveillez les statistiques pendant au moins une heure de trafic représentatif avant de considérer le réglage comme définitif. Un dimensionnement correct dès le départ évite des heures de diagnostic plus tard, lorsque le trafic augmente en production.

## Étape 10 : Exploiter le format eve.json pour l’analyse structurée

Le fichier `fast.log` est utile pour un test rapide, mais en production, tout doit passer par `eve.json`, qui structure chaque événement en JSON exploitable par un outil d’analyse. Vérifiez que le module `eve-log` est actif dans `suricata.yaml` :

```
outputs:
  - eve-log:
      enabled: yes
      filetype: regular
      filename: eve.json
      types:
        - alert
        - dns
        - tls
        - http
        - flow
        - ssh
```
Vous pouvez interroger ce flux directement avec `jq` pour un premier niveau d’analyse sans outil externe :

`sudo tail -f /var/log/suricata/eve.json | jq 'select(.event_type=="alert") | {time:.timestamp, src:.src_ip, dst:.dest_ip, sig:.alert.signature}'`
## Étape 11 : Centraliser les alertes dans un SIEM

Un IDS/IPS isolé, dont personne ne consulte les logs, ne sert à rien en pratique. L’étape suivante consiste à envoyer le flux `eve.json` vers un outil de centralisation et de corrélation. Deux options open source dominent l’écosystème francophone : Wazuh, qui combine SIEM et EDR, et Graylog, plus orienté centralisation de logs pure.

Si vous utilisez déjà Wazuh (voir notre tutoriel Wazuh SIEM Open Source), ajoutez simplement Suricata comme source de log dans le fichier `ossec.conf` de l’agent :

```
<localfile>
  <log_format>json</log_format>
  <location>/var/log/suricata/eve.json</location>
</localfile>
```
Si vous préférez Graylog pour la centralisation, notre guide pour centraliser ses logs de sécurité avec Graylog 7.1 détaille la configuration d’un input GELF ou Syslog compatible avec la sortie JSON de Suricata.

## Étape 12 : Compléter la défense avec un pare-feu et un scanner de vulnérabilités

Suricata détecte et peut bloquer du trafic malveillant, mais il ne remplace pas un pare-feu périmétrique ni un scanner de vulnérabilités. Une architecture de sécurité réseau cohérente combine généralement les trois. Si vous n’avez pas encore de pare-feu open source en place, notre tutoriel pour installer OPNsense couvre le filtrage périmétrique, qui peut d’ailleurs intégrer nativement Suricata comme moteur IDS/IPS via son propre plugin.

Pour identifier les vulnérabilités présentes sur vos serveurs avant qu’un attaquant ne les exploite, complétez votre chaîne avec un scanner comme celui présenté dans notre guide pour configurer OpenVAS/Greenbone. Et pour cartographier votre réseau et identifier les services exposés avant même de déployer Suricata, notre tutoriel Nmap reste la première étape logique de tout audit réseau.

## Étape 13 : Renforcer contre le brute-force et automatiser le blocage

Suricata excelle dans la détection de trafic malveillant complexe, mais pour bloquer automatiquement les tentatives de brute-force sur SSH ou vos services exposés, un outil complémentaire comme CrowdSec ou Fail2ban reste plus adapté et plus léger. Notre tutoriel CrowdSec montre comment bloquer les attaques répétées grâce à une base de réputation communautaire, en complément direct des alertes générées par Suricata.

## Suricata et le RGPD : ce qu’il faut savoir avant de capturer du trafic

Déployer un IDS/IPS qui inspecte le trafic réseau, y compris les métadonnées HTTP et DNS de vos utilisateurs, soulève des questions de conformité au RGPD qu’il ne faut pas négliger en France. Une adresse IP est considérée comme une donnée à caractère personnel par la CNIL, et les journaux générés par Suricata contiennent systématiquement des adresses IP source et destination, parfois associées à des noms de domaine ou des chemins d’URL révélateurs de l’activité d’un utilisateur.

