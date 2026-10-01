---
id: collect-261001-general-networking/general-networking/nftables-pare-feu-linux-en-13-etapes-2026-3
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/nftables-pare-feu-linux-en-13-etapes-2026.md
source_anchor: ""
source_lines: [184, 282]
sha256: 265e8fad4ace0fc4583f5601db7b6ae22a96692516b06e16713d9472c6f08f21
---

# Debian / Ubuntu

## Étape 11 : Migrer depuis iptables avec la couche de compatibilité nft

Si vous gérez déjà un serveur avec des règles iptables historiques, il existe un chemin de migration qui évite de tout réécrire à la main. L’utilitaire `iptables-translate`, fourni avec le paquet nftables, convertit une règle iptables ligne par ligne vers sa syntaxe nft équivalente.

```
# Convertir une règle iptables existante
iptables-translate -A INPUT -p tcp --dport 22 -j ACCEPT
# Convertir un fichier iptables-save complet
iptables-restore-translate -f /etc/iptables/rules.v4
```
La traduction automatique fonctionne bien pour les règles simples, mais elle ne réorganise pas la logique en tables et sets optimisés propres à nftables. Pour une migration de production, considérez le résultat comme un premier jet à nettoyer, pas comme une configuration finale prête à déployer telle quelle.

## Étape 12 : Tester le pare-feu depuis l’extérieur avec nmap

Une configuration qui semble correcte en local peut se comporter différemment vue depuis Internet. Testez systématiquement depuis une machine externe, jamais uniquement depuis le serveur lui-même.

```
# Depuis une machine externe (pas le serveur lui-même)
nmap -Pn -p 22,80,443,3389,8080 203.0.113.10
```
Exemple de sortie attendue sur un serveur correctement filtré :

```
PORT     STATE    SERVICE
22/tcp   open     ssh
80/tcp   open     http
443/tcp  open     https
3389/tcp filtered ms-wbt-server
8080/tcp filtered http-proxy
Nmap done: 1 IP address (1 host up) scanned in 4.21 seconds
```
Les ports explicitement autorisés apparaissent en `open`, tandis que les ports non listés dans vos règles remontent en `filtered`, ce qui signifie que les paquets sont abandonnés sans réponse plutôt que rejetés activement. C’est le comportement attendu d’une politique drop bien configurée : un attaquant qui scanne votre serveur ne peut même pas distinguer un port fermé d’un port inexistant.

## Étape 13 : Superviser et auditer vos règles en continu

Une configuration figée le jour de sa création devient obsolète avec le temps : nouveaux services déployés, ports oubliés qui restent ouverts, règles temporaires jamais retirées. Prenez l’habitude d’auditer périodiquement l’état réel du pare-feu.

```
# Lister toutes les règles actives avec compteurs de paquets
sudo nft list ruleset -a
# Voir uniquement les compteurs pour repérer les règles jamais déclenchées
sudo nft -s list table inet filtre
```
Une règle avec un compteur à zéro après plusieurs semaines mérite d’être questionnée : sert-elle encore à quelque chose, ou peut-elle être supprimée pour alléger la configuration ? Des outils de supervision comme Wazuh peuvent aussi centraliser ces logs nftables pour une analyse de sécurité à l’échelle de tout un parc de serveurs.

## Projet complet : fichier nftables.conf prêt pour la production

Voici l’assemblage complet des règles vues dans ce tutoriel, prêt à copier dans `/etc/nftables.conf` sur un serveur web type LAMP ou LEMP avec accès SSH restreint.

```
#!/usr/sbin/nft -f
flush ruleset
table inet filtre {
    set ports_autorises {
        type inet_service
        elements = { 22, 80, 443 }
    }
    set ip_bannies {
        type ipv4_addr
        flags timeout
    }
    chain entree {
        type filter hook input priority 0; policy drop;
        iif lo accept
        ct state established,related accept
        ct state invalid drop
        ip saddr @ip_bannies drop
        ip protocol icmp accept
        ip6 nexthdr icmpv6 accept
        tcp dport 22 ct state new limit rate 4/minute accept
        tcp dport 22 ct state new log prefix "SSH-DROP: " drop
        tcp dport @ports_autorises accept
        log prefix "NFT-DROP: " flags all counter drop
    }
    chain sortie {
        type filter hook output priority 0; policy accept;
    }
    chain transit {
        type filter hook forward priority 0; policy drop;
    }
}
```
Appliquez-le avec `sudo nft -c -f /etc/nftables.conf` pour valider la syntaxe, puis `sudo systemctl restart nftables`. Ce socle couvre déjà l’essentiel : politique restrictive par défaut, protection SSH contre le brute-force, gestion d’une liste noire dynamique, et journalisation des paquets rejetés. Adaptez les ports, la plage d’IP SSH et les seuils de limitation à votre contexte réel avant tout déploiement en production.

## Erreurs fréquentes à éviter avec nftables

Certaines erreurs reviennent systématiquement chez les administrateurs qui découvrent nftables, souvent parce qu’ils transposent des réflexes iptables qui ne s’appliquent plus de la même façon.

- **Appliquer une politique drop sans autoriser lo et les connexions établies au préalable** , ce qui coupe immédiatement l’accès SSH en cours et exige une intervention physique ou via console pour récupérer l’accès.
- **Laisser coexister iptables legacy et nftables sur les mêmes chaînes** , provoquant des comportements incohérents où une règle semble ignorée alors qu’elle est simplement écrasée par l’autre backend.
- **Oublier d’activer ip_forward avant de configurer du NAT** , ce qui fait échouer silencieusement toute règle de masquerade sans message d’erreur explicite.
- **Ne jamais sauvegarder la configuration dans /etc/nftables.conf** après avoir testé des règles en ligne de commande, perdant tout le travail au prochain redémarrage.
- **Placer la règle de journalisation avant les règles d’acceptation** plutôt qu’à la fin de la chaîne, ce qui génère des volumes de logs énormes pour du trafic pourtant légitime.
- **Ouvrir SSH au monde entier par réflexe** plutôt que de restreindre l’accès à une plage d’IP connue ou un VPN, alors que la limitation de débit seule ne suffit pas contre des botnets distribués sur des milliers d’IP différentes.

## Dépannage : problèmes courants et leurs solutions

Voici les situations les plus fréquemment rencontrées lors de la configuration de nftables, avec la cause probable et la solution associée.

