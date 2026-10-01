---
id: collect-261001-general-networking/general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026-5
title: "repérez le nom de votre interface, par exemple eth0 ou ens18"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "distribution", "incident", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/suricata-ids-ips-8-0-6-detecter-les-intrusions-2026.md
source_anchor: ""
source_lines: [351, 410]
sha256: e2860e3c71d1eb3b35d19ea19b3b29fcbcd44e3e0307022e53edf3055c8193fd
---

# repérez le nom de votre interface, par exemple eth0 ou ens18

1. Un serveur Debian 13 avec une interface dédiée en mode SPAN/mirroring recevant une copie du trafic du switch principal
2. Suricata 8.0.6 installé et configuré en mode IDS, avec HOME_NET défini sur la plage interne réelle
3. Le ruleset Emerging Threats Open, mis à jour quotidiennement via cron et suricata-update
4. La sortie eve.json envoyée en local vers un agent Wazuh qui transmet les événements au serveur central
5. Un tableau de bord Wazuh affichant les alertes Suricata classées par sévérité, IP source et signature
6. Une règle d’alerte critique déclenchant une notification par e-mail ou webhook vers un canal de messagerie interne dès qu’une signature de sévérité 1 (haute priorité) est détectée

Ce montage, réalisable en une demi-journée pour une équipe technique déjà familière avec Linux, constitue une base de détection réseau crédible pour toute structure qui doit démontrer une capacité de détection d’incident dans le cadre de ses obligations réglementaires.

## Cas d’usage concret : repérer une exfiltration de données

Pour illustrer l’intérêt pratique de tout ce déploiement, prenons un scénario courant en entreprise : un poste de travail compromis par un malware tente d’envoyer des documents internes vers un serveur externe via une requête HTTPS chiffrée, un comportement qu’un pare-feu classique laisserait passer sans broncher puisque le port 443 est légitimement ouvert vers l’extérieur.

Suricata, grâce à l’inspection TLS activée sur le module `tls-log`, peut extraire le nom de domaine du certificat présenté lors de la négociation TLS (SNI), même sans déchiffrer le contenu du trafic. Si ce domaine ne correspond à aucun service métier connu, ou s’il vient d’être enregistré depuis quelques jours seulement, cela déclenche une alerte de type “flux TLS suspect vers domaine récemment enregistré”, souvent activée par les règles Emerging Threats spécialisées dans la détection de commande et contrôle (C2). Combiné à un volume de données sortant anormalement élevé sur une courte période, détectable via les compteurs de flux Suricata, ce signal permet à une équipe sécurité d’isoler le poste concerné avant qu’une fuite de données ne devienne massive.

Ce type de détection illustre pourquoi un IDS/IPS reste pertinent même dans un monde où l’essentiel du trafic est chiffré : il ne s’agit pas de lire le contenu des communications, mais d’analyser les métadonnées, les motifs de comportement et les indicateurs de compromission connus, une approche que ni un pare-feu ni un antivirus de poste ne couvrent seuls.

## Suricata face à ses alternatives : Snort et Zeek

Suricata n’est pas le seul moteur IDS/IPS open source disponible. Snort, développé historiquement par Sourcefire puis racheté par Cisco, reste très utilisé mais avec un moteur mono-thread historiquement moins performant sur des débits élevés, même si Snort 3 a rattrapé une partie de ce retard. Zeek, de son côté, n’est pas à proprement parler un IDS basé sur signatures, mais un framework d’analyse de trafic orienté scripting, souvent déployé en complément de Suricata plutôt qu’en remplacement, comme le fait justement Security Onion qui embarque les deux moteurs simultanément (Suricata 8.0.5 et Zeek 8.0.8 dans la version 3.1.0 de mai 2026, en attendant une future mise à jour vers le moteur Suricata 8.0.7 publié par l’OISF en septembre 2026).

| Outil | Type de détection | Performance multi-cœur | Cas d’usage typique | 
|---|---|---|---|
| Suricata 8.0.6 | Signatures + DPI + extraction de fichiers | Native, multi-threadée | IDS/IPS généraliste, forte inspection applicative | 
| Snort 3 | Signatures | Améliorée mais historiquement plus limitée | Environnements Cisco, règles Talos | 
| Zeek 8.0.8 | Analyse comportementale scriptée | Native | Investigation forensique, complément à Suricata | 

## Foire aux questions

**Suricata est-il gratuit et open source ?**

Oui. Suricata est distribué sous licence GPLv2 par l’Open Information Security Foundation, une fondation à but non lucratif. Aucune licence n’est requise pour l’utiliser en production, y compris en entreprise.

**Faut-il utiliser Suricata en mode IDS ou en mode IPS ?**

Commencez systématiquement en mode IDS (détection passive) pendant plusieurs semaines pour observer le comportement des règles sur votre trafic réel, avant d’envisager un passage en mode IPS actif sur des flux non critiques.

**Que se passe-t-il si je reste sur Suricata 7 après juillet 2026 ?**

Vous ne recevrez plus aucun correctif de sécurité pour le moteur lui-même, ce qui expose potentiellement votre infrastructure de détection à des vulnérabilités non corrigées. La migration vers la branche 8.0.x est fortement recommandée par l’OISF.

**Suricata peut-il remplacer un pare-feu ?**

Non. Suricata complète un pare-feu comme OPNsense en apportant une inspection de contenu applicative, mais ne gère pas nativement le NAT, le routage ou les politiques de filtrage de base qu’assure un pare-feu périmétrique.

**Combien de RAM et de CPU faut-il pour un déploiement en production ?**

Pour un lien à quelques centaines de Mbit/s, 4 cœurs et 8 Go de RAM suffisent généralement. Au-delà de 1 Gbit/s soutenu, prévoyez davantage de cœurs et envisagez des méthodes de capture accélérées comme PF_RING ou DPDK.

**Quelle est la différence entre Suricata et Security Onion ?**

Suricata est un moteur de détection unique. Security Onion est une distribution Linux complète qui intègre Suricata, Zeek, Elasticsearch et un tableau de bord de visualisation, prête à l’emploi pour un SOC.

**Les règles Emerging Threats Open suffisent-elles ?**

Pour une PME, le ruleset gratuit ET Open couvre déjà une large partie des menaces courantes. Les organisations avec des exigences de détection plus poussées peuvent envisager les règles ET Pro, payantes, mises à jour plus fréquemment.

**Comment tester que Suricata fonctionne sans attendre une vraie attaque ?**

Utilisez la règle de test officielle en effectuant une requête vers `testmynids.org/uid/index.html`, qui déclenche une alerte inoffensive prévue à cet effet, puis vérifiez son apparition dans les logs.
