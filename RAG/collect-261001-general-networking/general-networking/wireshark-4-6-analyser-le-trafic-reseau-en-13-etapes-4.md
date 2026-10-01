---
id: collect-261001-general-networking/general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes-4
title: "Sur Fedora / RHEL / Rocky Linux"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/wireshark-4-6-analyser-le-trafic-reseau-en-13-etapes.md
source_anchor: ""
source_lines: [196, 257]
sha256: c910c5cb736a46acfc365c07d2c48819568d09072b0968a29de737a26516e61b
---

# Sur Fedora / RHEL / Rocky Linux

Ce projet reproduit un scénario réaliste sur un labo isolé (deux machines virtuelles sur un réseau privé, jamais sur une infrastructure de production). L’objectif : capturer une tentative de brute-force SSH, la confirmer dans Wireshark, puis l’automatiser via TShark pour une détection reproductible.

Sur la machine cible, lancez la capture ciblée sur le port 22 :

`tshark -i eth0 -f "tcp port 22" -w ssh-bruteforce.pcapng`
Depuis la machine “attaquante” du labo, générez du trafic de test avec un outil comme Hydra ou un simple script de connexions SSH répétées (uniquement dans ce labo isolé, jamais contre une cible tierce). Après quelques dizaines de tentatives, arrêtez la capture et ouvrez le fichier dans Wireshark. Appliquez le filtre d’affichage suivant pour isoler les tentatives de connexion :

`tcp.port==22 and tcp.flags.syn==1 and tcp.flags.ack==0`
Résultat attendu : une succession de paquets SYN depuis la même adresse IP source, espacés de quelques centaines de millisecondes à peine, chacun suivi d’un échange TCP complet puis d’une clôture rapide (paquets FIN ou RST) une fois l’authentification SSH refusée. Un utilisateur humain qui se trompe de mot de passe génère au maximum trois à cinq tentatives par minute. Un script en génère des dizaines par seconde : la différence de fréquence est le signal le plus fiable, bien plus que le contenu chiffré lui-même, invisible côté SSH.

Pour automatiser ce comptage sans ouvrir l’interface graphique à chaque fois, un script TShark suffit :

```
tshark -r ssh-bruteforce.pcapng \
  -Y "tcp.port==22 and tcp.flags.syn==1 and tcp.flags.ack==0" \
  -T fields -e ip.src -e frame.time_relative \
  | awk '{print $1}' | sort | uniq -c | sort -rn
# Exemple de sortie :
#   47 192.168.56.20
#    2 192.168.56.30
```
Une adresse avec 47 tentatives SYN en quelques secondes, contre 2 pour une autre (probablement un test légitime), confirme sans ambiguïté la nature de l’attaque. Dans un contexte réel, ce script s’intègre à un cron ou à un pipeline de supervision qui alerte au-delà d’un seuil défini, par exemple 10 tentatives par minute et par IP, en complément d’un blocage automatique via Fail2ban.

## 5 erreurs fréquentes à éviter

- **Capturer sans filtre sur un lien saturé.** Une capture de dix minutes sans filtre sur un serveur à fort trafic peut générer plusieurs gigaoctets et rendre l’analyse dans l’interface graphique quasi impossible. Filtrez toujours à la capture, pas seulement à l’affichage.
- **Lancer Wireshark en root en permanence sous Linux.** C’est le raccourci le plus tentant et le plus risqué : toute faille dans le moteur de décodage (et il y en a eu, historiquement) s’exécute alors avec les droits root. Configurez le groupe`wireshark` une bonne fois pour toutes.
- **Confondre filtre de capture et filtre d’affichage.** Taper une syntaxe de filtre d’affichage (comme`tcp.port==80` ) dans le champ de capture, ou l’inverse, provoque une erreur ou un filtre silencieusement ignoré. Les deux langages ne sont pas interchangeables.
- **Ignorer les colonnes d’heure relative.** Par défaut, Wireshark affiche l’heure depuis le début de la capture, pas l’heure absolue. Pour corréler avec des logs serveur, basculez sur Affichage > Format d’heure > Date et heure du jour.
- **Analyser du trafic TLS sans exporter les clés de session au préalable.** Il est impossible de déchiffrer une capture TLS après coup sans avoir configuré`SSLKEYLOGFILE` avant la capture : pensez-y en amont, pas après avoir fermé le navigateur.

## Dépannage : 8 problèmes courants et leurs solutions

| Problème | Cause probable | Solution | 
|---|---|---|
| “You don’t have permission to capture” au démarrage | Droits insuffisants sur l’interface ou le pilote de capture | Ajouter l’utilisateur au groupe wireshark (Linux) ou réinstaller Npcap avec les droits par défaut (Windows) | 
| Aucune interface ne s’affiche dans la liste | Npcap ou libpcap non installé, ou service arrêté | Réinstaller Npcap, ou vérifier `systemctl status NetworkManager` sur Linux | 
| La capture reste vide malgré du trafic actif | Mauvaise interface sélectionnée, ou filtre de capture trop restrictif | Vérifier avec `ip addr` quelle interface est active, retirer temporairement le filtre | 
| Filtre d’affichage en rouge, refusé | Erreur de syntaxe (opérateur, nom de champ) | Utiliser l’autocomplétion intégrée ou “Apply as Filter” depuis un paquet existant | 
| Impossible de voir le trafic d’une autre machine sur le switch | Réseau commuté moderne : chaque port ne voit que son propre trafic | Configurer un port mirroring (SPAN) sur le switch, ou capturer directement sur la machine cible | 
| Trafic TLS toujours chiffré malgré SSLKEYLOGFILE configuré | Variable définie après le lancement de l’application, ou mauvais chemin de fichier | Redémarrer l’application après avoir exporté la variable, vérifier le chemin dans Préférences > TLS | 
| Wireshark se fige ou consomme toute la RAM sur une grosse capture | Fichier de plusieurs Go ouvert directement dans l’interface graphique | Découper le fichier avec `editcap -c 100000 gros.pcapng partie.pcapng` ou filtrer via TShark en amont | 
| Les noms de domaine ne s’affichent pas, seulement des IP | Résolution de noms désactivée pour limiter les requêtes DNS parasites | Activer temporairement Affichage > Résolution de noms, en sachant que cela génère du trafic DNS supplémentaire | 

## Astuces avancées pour aller plus loin

Une fois les bases maîtrisées, plusieurs fonctions accélèrent nettement le travail d’analyse au quotidien. Les colonnes personnalisées (clic droit sur un champ > “Apply as Column”) permettent d’afficher directement dans la liste principale une information précise, par exemple le code de statut HTTP ou la durée de round-trip TCP, sans avoir à déplier les détails de chaque paquet. Les profils de configuration (Édition > Profils de configuration) autorisent de basculer instantanément entre un jeu de colonnes et de filtres adapté au débogage web et un autre adapté à l’analyse d’intrusion, sans reconfigurer l’interface à chaque fois.

Pour le trafic USB, Wireshark peut capturer directement les échanges d’un périphérique via le module USBPcap (Windows) ou usbmon (Linux), utile pour analyser un keylogger matériel suspect ou déboguer un firmware. Pour le Wi-Fi, le mode moniteur (nécessitant une carte compatible et, sous Linux, l’outil `airmon-ng`) permet de capturer les trames 802.11 brutes, y compris les paquets de gestion, avant même l’association à un point d’accès.

Enfin, l’outil `editcap`, fourni avec Wireshark, permet de découper, fusionner ou anonymiser des captures en ligne de commande, une étape souvent nécessaire avant de partager un fichier .pcapng avec un tiers dans le cadre d’un support technique, sans exposer d’adresses IP ou de données internes sensibles.

## Wireshark face aux autres outils d’analyse réseau

| Outil | Type | Interface | Cas d’usage principal | 
|---|---|---|---|
| Wireshark | Analyseur de paquets | Graphique (GTK/Qt) | Analyse détaillée manuelle, décodage protocolaire | 
| TShark | Analyseur de paquets | Ligne de commande | Automatisation, capture sur serveur distant | 
| tcpdump | Capteur de paquets | Ligne de commande | Capture rapide et légère, sans décodage avancé | 
| Suricata | IDS/IPS | Ligne de commande + règles | Détection automatique de signatures et blocage temps réel | 
| Nmap | Scanner réseau | Ligne de commande | Découverte d’hôtes et de ports, pas d’analyse de flux | 

