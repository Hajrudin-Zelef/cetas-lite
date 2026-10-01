---
id: collect-261001-general-networking/general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026-3
title: "Vérifier l'espace disque disponible (minimum 20 Go recommandé)"
domain: general-networking
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-general-networking/openvas-greenbone-installer-un-scanner-de-failles-2026.md
source_anchor: ""
source_lines: [171, 250]
sha256: 57840879abc74f4a143090975de84f4b22ca673e5960e44ebccf10cd21856290
---

# Vérifier l'espace disque disponible (minimum 20 Go recommandé)

```
# Via l'API GMP, exporter un rapport au format CSV
gmp.get_report(report_id="abc-123", filter_string="apply_overrides=0 min_qod=70", format_id="c1645568-627a-11e3-a660-406186ea4fc5")
```
Conservez ces rapports au moins 12 mois dans un stockage séparé du serveur Greenbone lui-même, idéalement chiffré et avec un accès restreint. En cas d'audit NIS2 ou de contrôle de la CNIL suite à un incident, la capacité à démontrer un historique continu de scans et de remédiation pèse souvent plus lourd que le résultat d'un seul rapport isolé. Beaucoup d'organisations découvrent trop tard, au moment de l'audit, qu'elles n'ont conservé aucune trace des scans réalisés l'année précédente.

## Étape 10 : configurer des scans authentifiés (credentialed scans)

Un scan non authentifié ne voit que la surface exposée depuis le réseau. Pour détecter les failles internes (versions de paquets installés, correctifs Windows manquants, mauvaises configurations locales), configurez des identifiants dans Configuration → Credentials, puis associez-les à votre cible.

- Pour Linux : créez un couple SSH clé publique/privée dédié au scan, avec un compte à droits limités plutôt que root
- Pour Windows : utilisez un compte de service avec accès WMI/SMB, jamais un compte administrateur de domaine
- Stockez ces identifiants uniquement dans le coffre-fort intégré de Greenbone, jamais en clair dans un script

Les scans authentifiés remontent en moyenne deux à trois fois plus de vulnérabilités que les scans réseau seuls, car ils inspectent directement l'état du système plutôt que de deviner ses failles depuis l'extérieur. C'est cette profondeur qui distingue un vrai programme de gestion des vulnérabilités d'un simple scan de surface, et c'est aussi le réglage que les auditeurs NIS2 vérifient en priorité lorsqu'ils examinent la méthodologie de scan d'une entité essentielle ou importante.

## Étape 11 : intégrer OpenVAS avec Proxmox VE et les environnements virtualisés

Le classement Wiz 2026 souligne la prise en charge d'OpenVAS pour les infrastructures virtualisées type Proxmox VE 8.0 et supérieur, ou Huawei FusionCompute 8.0. Dans ce contexte, créez une cible groupée par VLAN de gestion plutôt que de scanner chaque hyperviseur individuellement, et excluez les interfaces de migration à chaud (souvent non chiffrées et sensibles aux interruptions) de la portée du scan.

```
# Exemple de plage ciblée pour un cluster Proxmox (3 nœuds)
10.10.20.11-10.10.20.13
# Exclure l'interface de migration à chaud (souvent sur un VLAN dédié)
# via le champ "Exclude Hosts" du formulaire Target
```
## Étape 12 : mettre à jour vers OPENVAS SCAN 25.0.6 ou migrer depuis 24.10

Si votre installation tourne encore sur la branche 24.10.x, planifiez une migration vers 25.0.6 lors d'une fenêtre de maintenance. Notez que certains dépôts d'installation communautaires sur GitHub annoncent depuis août 2025 la prise en charge d'une mise à niveau automatique (auto upgrade), une option à évaluer avec prudence en production avant de vous y fier entièrement. Sauvegardez d'abord la base PostgreSQL et le volume de configuration :

```
# Sauvegarde avant migration
docker compose exec pg-gvm pg_dumpall -U gvmd > backup_gvmd_$(date +%F).sql
# Mise à jour de l'image et redémarrage
docker compose pull
docker compose up -d
# Vérifier la version après migration
docker compose exec gvmd gvmd --version
```
Consultez la documentation officielle de Greenbone TechDoc avant toute migration en production, car chaque montée de version majeure peut modifier le schéma de base de données.

## Étape 13 : durcir l'accès à l'interface web GSA

L'interface GSA donne accès à l'ensemble des identifiants de scan, des rapports de vulnérabilités et des cibles de votre réseau : c'est une cible de choix si elle reste exposée sans protection supplémentaire. Trois réglages simples réduisent fortement le risque. D'abord, limitez l'accès au port 9392 par pare-feu à une liste d'IP autorisées plutôt que de l'ouvrir largement, même en interne.

```
# Restreindre l'accès au port GSA à un sous-réseau d'administration
sudo ufw allow from 192.168.10.0/24 to any port 9392 proto tcp
sudo ufw deny 9392/tcp
```
Ensuite, remplacez le certificat TLS auto-signé par un certificat émis par une autorité interne ou par Let's Encrypt si l'interface est exposée derrière un reverse proxy comme Nginx ou Traefik. Enfin, désactivez le compte `admin` par défaut une fois un second compte administrateur nominatif créé, et activez la journalisation des connexions pour repérer toute tentative d'accès suspecte. Ces trois réglages combinés transforment une interface d'administration ouverte en un point d'accès correctement cloisonné, conforme aux attentes d'un audit NIS2 sur le contrôle des accès.

## Quelle fréquence de scan choisir selon votre secteur d'activité

La fréquence idéale dépend directement de la criticité des systèmes exposés et du niveau de réglementation applicable. Une collectivité soumise à NIS2 n'a pas les mêmes obligations qu'un hébergeur SecNumCloud ou qu'une PME sans exigence sectorielle particulière, mais dans tous les cas, un scan trimestriel isolé ne suffit plus face au rythme de publication des nouvelles CVE.

| Profil d'organisation | Fréquence recommandée | Portée | 
|---|---|---|
| Entité essentielle NIS2 (santé, énergie, transport) | Hebdomadaire ou continue | Tous les actifs exposés + scan authentifié mensuel | 
| Collectivité territoriale | Hebdomadaire | Serveurs exposés Internet, mensuel pour le parc interne | 
| PME sans obligation réglementaire | Mensuelle | Serveurs de production, sites web publics | 
| Hébergeur / infogéreur | Continue (à chaque changement) | Ensemble du parc client, déclenché par la CI/CD | 
| Environnement de test / laboratoire | À la demande | Selon les campagnes de test | 

Dans la pratique, la meilleure approche combine un scan léger et fréquent (hebdomadaire, orienté détection de nouvelles failles critiques) avec un scan complet et authentifié mensuel qui couvre l'intégralité du parc, y compris les correctifs internes. Cette cadence à deux vitesses évite de saturer les équipes tout en gardant une visibilité proche du temps réel sur les vulnérabilités critiques nouvellement publiées.

## 5 erreurs fréquentes lors de l'installation d'OpenVAS

Ces pièges reviennent le plus souvent dans les forums communautaires et les retours d'administrateurs qui déploient Greenbone pour la première fois.

- **Lancer un scan avant la fin de la synchronisation du feed.** Le premier scan sur un feed incomplet donne des résultats faussement rassurants, car de nombreux NVT ne sont pas encore chargés.
- **Scanner sans autorisation écrite.** Même sur son propre réseau d'entreprise, l'absence de mandat clair peut poser problème en cas d'incident pendant le scan.
- **Sous-dimensionner la RAM.** En dessous de 4 Go, le conteneur PostgreSQL et le moteur de scan se disputent la mémoire, provoquant des scans qui échouent à 0 % (un problème régulièrement signalé sur le forum communautaire Greenbone en août 2026, notamment sur Kali Linux).
- **Ignorer les faux positifs.** Un scan non authentifié peut signaler une faille sur un service dont la bannière de version a été modifiée volontairement (hardening), sans que la faille existe réellement. Vérifiez toujours manuellement les résultats critiques.
- **Oublier de renouveler le certificat TLS de l'interface GSA.** Le certificat auto-signé par défaut expire au bout de 12 mois et bloque l'accès à l'interface sans message d'erreur explicite dans certains navigateurs.

## Dépannage : 8 problèmes courants et leurs solutions

Voici les incidents les plus signalés lors du déploiement de Greenbone Community Edition, avec la commande ou le réglage qui permet généralement de les résoudre.

