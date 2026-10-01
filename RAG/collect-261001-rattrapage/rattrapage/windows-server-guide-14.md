---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-14
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "benchmarks", "datacenter", "distribution"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [2232, 2308]
sha256: 9e83e2b9beb2930c04a48ccac9f9adc73d9a61acda406dd4caea6e010aa5dd9a
---

# Windows Server en entreprise — Guide technique ultra-complet

- **NTP** : protocole de synchronisation de l'heure (vital pour Kerberos/AD).
- **OCSP** : vérification en ligne du statut de révocation d'un certificat.
- **OSE** (Operating System Environment) : instance de système d'exploitation couverte par une licence.
- **PKI** : infrastructure à clés publiques (AC, certificats, CRL).
- **Quorum** : mécanisme de vote qui évite le split-brain dans un cluster.
- **RDS** : services Bureau à distance (ex-TSE).
- **RPO / RTO** : perte de données max acceptable / durée max d'interruption acceptable.
- **RSAT** : outils d'administration à distance (sur poste Windows).
- **SMB** : protocole de partage de fichiers Windows.
- **Spooler** : service gérant les files d'impression.
- **VSS** : clichés instantanés (versions précédentes).
- **WAC** (Windows Admin Center) : console web d'administration des serveurs.
- **WID** : base de données interne Windows (utilisée par WSUS).
- **WSB** : Windows Server Backup.
- **WSUS** : service de distribution des mises à jour Microsoft.

---

## 83. Quiz : 10 questions

1. Un hôte Hyper-V fait tourner 10 VM de production. Quelle édition de Windows Server est la plus économique, et pourquoi ?
2. Quelle est la différence fondamentale entre un checkpoint standard et un checkpoint de production ?
3. Un utilisateur a "Accès refusé" sur un partage alors qu'il est dans le bon groupe AD depuis ce matin. Quelle est la cause la plus probable ?
4. Pourquoi faut-il combiner DFS-N avec des chemins UNC directs vers les serveurs ?
5. Quelle est la durée du délai de grâce RDS sans serveur de licences, et que se passe-t-il après ?
6. Dans une PKI à deux niveaux, quel est le rôle de l'AC racine et combien de temps doit-elle rester allumée ?
7. Que signifie 3-2-1 et quelle erreur commet-on le plus souvent avec les sauvegardes ?
8. KMS ou ADBA : lequel choisir pour un parc 100 % joint au domaine, et pourquoi ?
9. Un client WSUS n'apparaît pas dans la console. Citez 3 vérifications dans l'ordre.
10. Pourquoi ne faut-il jamais exposer le port 3389 (RDP) directement sur Internet ?

---

## 84. Quiz : réponses commentées

1. **Datacenter.** Chaque licence Standard ne couvre que 2 VM ; pour 10 VM il faudrait 5 licences Standard, plus cher qu'une Datacenter (VM illimitées). Seuil de rentabilité ≈ 6 VM (§2).
2. Le checkpoint **standard** fige l'état mémoire exact (restauration à l'identique, plutôt lab) ; le checkpoint de **production** utilise VSS dans la VM pour une image cohérente au niveau applicatif (adapté à la prod). Ni l'un ni l'autre n'est une sauvegarde (§20).
3. **Le jeton d'accès** : l'appartenance au groupe a été ajoutée après l'ouverture de session. Il faut fermer/rouvrir la session (ou `klist purge` + reverrouillage) pour que le nouveau groupe soit pris en compte (§37).
4. Il ne faut **pas** utiliser de chemins directs : DFS-N fournit un chemin logique stable (`\\entreprise.lan\data\...`) qui survit aux migrations de serveurs. Les UNC directs (`\\SRV-01\...`) cassent à chaque migration (§33).
5. **120 jours.** Après, les connexions RDS sont refusées ("aucun serveur de licences disponible"). C'est la panne RDS la plus classique (§45).
6. L'AC **racine** signe uniquement l'AC subordonnée (1×/an ou à l'émission), puis est **éteinte et mise au coffre**. Elle ne doit jamais être en ligne en permanence : c'est l'AC subordonnée d'entreprise qui émet au quotidien (§47-48).
7. **3 copies, 2 supports différents, 1 hors site.** L'erreur la plus courante : ne jamais tester la restauration — une sauvegarde non testée est une hypothèse (§54).
8. **ADBA** : pas de seuil d'activation (KMS exige 25 postes/5 serveurs), activation persistante tant que la machine est jointe au domaine, rien à maintenir (§57).
9. (1) La GPO pointe-t-elle vers `http://srv-wsus:8530` ? (registre `WUServer`) ; (2) connectivité `Test-NetConnection -Port 8530` ; (3) `SusClientId` en doublon (clone sans sysprep) puis `wuauclt /resetauthorization /detectnow` et `UsoClient StartScan` (§29).
10. Le RDP exposé subit un **brute-force permanent** (bots) ; une faille ou un mot de passe faible = porte d'entrée (rançongiciels). Passer par VPN ou passerelle RDS en 443, avec NLA et comptes nominatifs (§60-61).

---

## 85. Pour aller plus loin : 20 ressources et chantiers

**Chantiers à planifier :**
1. Migrer les derniers 2019 vers 2022/2025 (support étendu = sécurité seule).
2. Déployer Windows LAPS natif sur tout le parc (§63).
3. Passer les hôtes Hyper-V en Server Core si ce n'est pas fait (§10).
4. Mettre en place la PKI à deux niveaux si absente (§47-48).
5. Tester une restauration bare metal par trimestre (§76).
6. Documenter le PRA et le faire valider par la direction.
7. Isoler les équipements exigeant SMBv1 sur un VLAN dédié (§60).
8. Remplacer les UNC directs par DFS-N partout (§33).
9. Auditer l'AD avec PingCastle (gratuit).
10. Chiffrer les disques de sauvegarde en rotation (BitLocker, §54).

**Ressources :**
11. Microsoft Learn — documentation Windows Server (référence officielle, à jour par version).
12. Microsoft Security Compliance Toolkit — baselines GPO de durcissement.
13. PingCastle — audit de la sécurité Active Directory.
14. Les guides compagnons : `onduleurs_ups_guide.md` (coupures → arrêt propre des serveurs : NUT), `proxmox_guide.md`, `zabbix_guide.md`, `debian_ubuntu_guide.md`, `kyocera_copieurs_guide.md` (métier copieurs ↔ serveur d'impression §38-42).
15. Veeam — guides de dimensionnement et best practices (même sans licence, la doc est pédagogique).
16. NIST SP 800-53 / CIS Benchmarks Windows Server — référentiels de durcissement.
17. TechNet Gallery / GitHub — scripts PowerShell communautaires (à relire avant usage).
18. Wireshark — comprendre SMB/Kerberos/LDAP quand le dépannage coince.
19. Laboratoire : 1 hôte Hyper-V + 3 VM (DC, fichiers, WSUS) pour tester chaque section de ce guide sans risque.
20. Ce guide lui-même : le relire une fois par an, cocher ce qui est appliqué, planifier le reste.

---

*Fin du guide — Windows Server en entreprise. Document de travail : adapte les noms de serveurs, IP et domaines à ton infrastructure réelle avant toute exécution.*
