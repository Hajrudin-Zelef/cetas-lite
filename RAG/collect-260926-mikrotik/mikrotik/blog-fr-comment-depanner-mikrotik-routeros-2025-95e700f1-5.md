---
id: collect-260926-mikrotik/mikrotik/blog-fr-comment-depanner-mikrotik-routeros-2025-95e700f1-5
title: "Exemple : Tester si le pare-feu bloque le trafic"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/blog-fr-comment-depanner-mikrotik-routeros-2025-95e700f1.md
source_anchor: ""
source_lines: [491, 623]
sha256: e3f77c26b361842df9288b6b1a4fd7bdff5e3c5bd503d42808c6f1635d7b2c83
---

# Exemple : Tester si le pare-feu bloque le trafic

1. 
**Capture et Analyse de Paquets :**
  - Capturer le trafic sur des interfaces spécifiques :
plaintext```
/tool sniffer set filter-interface=ether1 filter-ip-address=192.168.1.100/32
/tool sniffer start
```
  - Exporter les captures pour analyse dans Wireshark :
plaintext```
/tool sniffer save file=capture.pcap
```
  - Analyser des protocoles spécifiques :
plaintext```
/tool sniffer set filter-interface=ether1 filter-port=53
```
  - Se concentrer sur les problèmes d'établissement de connexion :
plaintext```
/tool sniffer set filter-interface=ether1 filter-tcp-flags=syn
```
2. Capturer le trafic sur des interfaces spécifiques :
3. 
**Analyse Avancée des Journaux :**
  - Configurer une journalisation détaillée pour des sujets spécifiques :
plaintext```
/system logging add topics=firewall,debug action=memory
```
  - Filtrer les journaux pour des motifs spécifiques :
plaintext```
/log print where message~"failed"
```
  - Exporter les journaux pour analyse externe :
plaintext```
/log print file=journaux-detaillés
```
  - Configurer la journalisation à distance :
plaintext```
/system logging add topics=system,critical action=remote remote=192.168.1.5
```
4. Configurer une journalisation détaillée pour des sujets spécifiques :
5. 
**Scripting pour les Diagnostics Automatisés :**
  - Créer un script de diagnostic complet :
plaintext```
/system script add name=diagnostics source={
  :log info "Démarrage des diagnostics"
  :log info "Ressources système :"
  /system resource print
  :log info "État de l'interface :"
  /interface print status
  :log info "Table de routage :"
  /ip route print
  :log info "Connexions actives :"
  /ip firewall connection print count-only
  :log info "État DNS :"
  /tool dns-lookup name=google.com
}
```
  - Planifier une exécution régulière :
plaintext```
/system scheduler add name=diagnostics-quotidiens interval=1d on-event=diagnostics
```
  - Créer des scripts de récupération conditionnelle :
plaintext```
/system script add name=récupérer-internet source={
  :if ([/ping 8.8.8.8 count=3] = 0) do={
    :log warning "Internet HS, réinitialisation du WAN"
    /interface disable ether1
    :delay 5s
    /interface enable ether1
  }
}
```
6. Créer un script de diagnostic complet :
7. 
**Diagnostic des Problèmes Intermittents :**
  - Implémenter une surveillance continue :
plaintext```
/tool netwatch add host=8.8.8.8 interval=30s up-script=":log info up" down-script=":log warning down"
```
  - Créer des graphiques de bande passante pour une analyse à long terme :
plaintext```
/tool graphing interface add interface=ether1
```
  - Configurer des alertes e-mail automatiques pour les événements critiques :
  - Utiliser la fonction de surveillance de l'état de santé :
plaintext```
/system health print
```
8. Implémenter une surveillance continue :
9. 
**Diagnostics au Niveau Matériel :**
  - Rechercher les erreurs matérielles :
plaintext```
/system routerboard print
```
  - Tester la stabilité de l'alimentation électrique :
plaintext```
/system health print
```
  - Surveiller la température :
plaintext```
/system health print
```
  - Effectuer des tests de contrainte pour identifier les problèmes matériels :
plaintext```
/tool bandwidth-test address=mikrotik-distant duration=1h direction=both
```
10. Rechercher les erreurs matérielles :

**Résumé de la Section :** Les techniques de dépannage avancées vous permettent d'approfondir les problèmes MikroTik complexes qui ne cèdent pas aux diagnostics de base. En maîtrisant l'analyse de paquets, le scriptage et la surveillance systématique, vous pouvez résoudre même les problèmes réseau les plus difficiles et développer des mesures proactives pour prévenir leur récurrence.

**Mini-FAQ :**

Utilisez la capture de paquets lorsque vous avez besoin de comprendre exactement ce qui se passe au niveau du protocole, en particulier lors du dépannage de problèmes spécifiques à l'application, de l'investigation de problèmes de sécurité ou du diagnostic de problèmes où les symptômes n'indiquent pas clairement la cause. C'est particulièrement précieux pour les problèmes intermittents que d'autres outils ne parviennent pas à identifier.

Configurez des scripts planifiés pour exécuter des diagnostics pendant les périodes problématiques, configurez une journalisation détaillée axée sur les composants suspects et mettez en œuvre une surveillance continue avec des outils comme Netwatch et Graphing. La collecte de données sur plusieurs occurrences révèle souvent des modèles qui pointent vers la cause première.

**Introduction à la Section :** Un dépannage efficace ne consiste pas seulement à connaître des techniques individuelles, il s'agit d'avoir une approche systématique qui mène à une résolution efficace des problèmes. Cette section vous aidera à développer un flux de travail structuré pour résoudre les problèmes MikroTik.

**Explication :** Un processus de dépannage méthodique permet de s'assurer qu'aucune cause potentielle n'est négligée et évite de perdre du temps sur des solutions inefficaces. Il facilite également le transfert de connaissances et la documentation.

**Détails Techniques :** Nous explorerons un cadre de dépannage étape par étape spécialement adapté aux environnements MikroTik, y compris la définition du problème, la collecte d'informations, le test d'hypothèses et la mise en œuvre de la solution.

**Avantages et Applications :**

- Réduire le temps moyen de résolution des problèmes réseau
- Assurer une qualité de dépannage constante entre les membres de l'équipe
- Prévenir les problèmes récurrents grâce à une analyse des causes premières appropriée
- Constituer une base de connaissances de solutions pour référence future
- Minimiser le risque d'aggraver les situations pendant le dépannage

**Instructions Étape par Étape pour la Mise en Œuvre d'un Flux de Travail de Dépannage :**

