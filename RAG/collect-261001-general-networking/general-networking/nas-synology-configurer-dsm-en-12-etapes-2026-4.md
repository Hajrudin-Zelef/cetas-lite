---
id: collect-261001-general-networking/general-networking/nas-synology-configurer-dsm-en-12-etapes-2026-4
title: "Exemple de planification de tâche Hyper Backup (via l'interface DSM)"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agents", "ethernet", "mai"]
source: docs/RAG/collect-261001-general-networking/nas-synology-configurer-dsm-en-12-etapes-2026.md
source_anchor: ""
source_lines: [162, 225]
sha256: c052af145cf8ac32707087cc93a4ca1951de0404349eb8986404f0769d08a67d
---

# Exemple de planification de tâche Hyper Backup (via l'interface DSM)

Pour rendre ce tutoriel concret, voici la configuration complète telle qu’assemblée de bout en bout sur un DS224+, pensée pour un foyer de 3 à 4 personnes avec photos, documents et un peu de partage multimédia :

- Boîtier : Synology DS224+, deux disques 4 To dédiés NAS en SHR-1 (soit environ 3,6 To exploitables après formatage)
- Système de fichiers : Btrfs, pour les snapshots et la protection contre la corruption silencieuse
- Sécurité : compte admin renommé, 2FA activée, Auto Block resserré à 3 tentatives
- Accès distant : QuickConnect activé uniquement pour Synology Photos et Drive, DSM lui-même accessible seulement via VPN local
- Sauvegarde : Hyper Backup quotidien vers Synology C2 Storage (palier 1 To, chiffré, rétention 7/4/6)
- Applications : Synology Drive pour les documents partagés, Synology Photos avec albums familiaux partagés
- Snapshots locaux : toutes les 4 heures sur le dossier “Documents”, conservés 30 jours, via le paquet Snapshot Replication

Cette configuration coûte, matériel compris (NAS + 2 disques 4 To), environ 650 à 700 € à l’achat, plus 60 € par an pour la sauvegarde cloud C2. Sur 5 ans, cela reste comparable au coût cumulé d’un abonnement cloud premium à 2 To, mais avec un contrôle total sur les données et des fonctionnalités de partage familial bien plus riches.

## Résultats attendus et exemples de sortie

Une fois la configuration terminée, voici ce que l’on doit observer dans le tableau de bord DSM (Panneau de contrôle > Ressource système) :

```
État du volume : Sain (Btrfs, SHR-1)
Espace utilisé : 2 % (installation initiale)
Statut Hyper Backup : Dernière sauvegarde réussie, il y a 2 heures
QuickConnect : Connecté (ID : nas-famille-xxxx)
2FA : Activée sur le compte administrateur
Mises à jour DSM : À jour (7.4.1-90080)
Ventilateurs : Normal, 45 % de charge
Température disques : 32-36°C
```
Sur un transfert initial de 500 Go via câble Ethernet Gigabit, comptez environ 1h15 à 1h30 en conditions réelles (débit soutenu autour de 90-100 Mo/s), contre plusieurs heures voire une journée entière en Wi-Fi selon la qualité du signal.

## 5 pièges fréquents à éviter

**1. Utiliser des disques de bureau grand public plutôt que des disques NAS.** Le manque de tolérance aux vibrations et l’absence d’optimisation pour un fonctionnement 24/7 réduisent significativement la durée de vie dans un boîtier multi-disques.

**2. Choisir ext4 sans réfléchir, puis regretter l’absence de snapshots.** La conversion a posteriori demande une sauvegarde complète, une suppression du volume et une restauration : autant bien choisir dès le départ.

**3. Confondre RAID et sauvegarde.** Un RAID protège contre la panne d’un disque, pas contre un vol, un incendie, une erreur de manipulation ou un ransomware qui chiffre le volume entier. Sans copie externe via Hyper Backup, les données restent exposées à un point de défaillance unique.

**4. Exposer DSM directement sur internet sans 2FA ni restriction d’accès.** Les campagnes de ransomware ciblant spécifiquement les NAS grand public exploitent en priorité les comptes admin par défaut et l’absence d’authentification à deux facteurs.

**5. Ignorer les avis de sécurité Synology une fois le NAS configuré.** Un NAS “qui marche” n’est pas un NAS “à jour” : sans mises à jour automatiques activées, un appareil peut rester vulnérable à des failles corrigées depuis des mois.

## Conseils avancés pour aller plus loin

Une fois la base solide, plusieurs réglages permettent de tirer davantage parti d’un NAS Synology. Le paquet **Snapshot Replication** permet de répliquer en temps quasi réel les snapshots vers un second NAS distant, ce qui offre une redondance géographique sans passer par le cloud pour les structures qui disposent de deux sites. Le paquet **Active Backup for Business**, gratuit, permet de sauvegarder des postes Windows, des machines virtuelles VMware/Hyper-V et même des instances Microsoft 365 directement sur le NAS, sans licence supplémentaire — une option souvent ignorée alors qu’elle remplace des solutions de sauvegarde d’entreprise payantes pour les petites structures. Pour les structures qui pensent déjà à la suite, Synology a présenté à Computex **DSM Enterprise 1.0**, avec un support annoncé à partir de mai 2026 et une maintenance prévue jusqu’en mai 2028, ainsi qu’une feuille de route dévoilée le 4 juin 2026 autour d’agents d’IA privés et d’un **Cluster Manager** destinés à orchestrer de l’IA on-premise gouvernée sur plusieurs nœuds NAS.

Pour les utilisateurs qui gèrent aussi des SSD ou des disques dans d’autres machines (PC de jeu, PS5), surveiller leur état de santé reste tout aussi important que pour les disques du NAS ; notre tutoriel sur CrystalDiskInfo détaille comment lire les indicateurs SMART qui annoncent une panne avant qu’elle ne survienne. Côté extension de stockage, l’installation d’un SSD M.2 NVMe suit une logique similaire à ce que nous avions documenté pour l’installation d’un SSD M.2 sur PS5, avec les mêmes précautions de manipulation antistatique.

Enfin, pensez à documenter votre configuration : notez le modèle de NAS, la disposition RAID, les identifiants QuickConnect et la procédure de restauration Hyper Backup dans un endroit accessible même si le NAS tombe en panne. C’est le genre de détail qu’on néglige au moment de l’installation et qu’on regrette amèrement le jour où il faut restaurer une sauvegarde en urgence, sous pression, sans savoir par où commencer.

## Dépannage : 8 problèmes courants et leurs solutions

**Le NAS n’apparaît pas dans Synology Assistant ou sur find.synology.com.** Vérifiez que l’ordinateur et le NAS sont bien sur le même sous-réseau, désactivez temporairement le pare-feu du PC, et essayez de brancher le câble Ethernet sur un autre port du routeur. Sur un réseau avec VLAN, la découverte automatique échoue souvent : utilisez alors l’adresse IP directe.

**L’installation de DSM échoue ou reste bloquée.** Un téléchargement de firmware interrompu est la cause la plus fréquente. Retirez et réinsérez les disques, redémarrez le NAS, et relancez l’installation en vérifiant la stabilité de la connexion internet.

**La création du volume échoue avec une erreur de disque.** Un des disques est probablement défectueux ou mal reconnu. Vérifiez son état dans Gestionnaire de stockage > HDD/SSD, et testez-le individuellement si le message d’erreur persiste.

**QuickConnect affiche “Non connecté” ou une erreur de relais.** Vérifiez que le NAS a bien accès à internet, que l’horloge système est synchronisée (Panneau de configuration > Heure), et réessayez après quelques minutes : les serveurs de relais Synology peuvent être temporairement surchargés.

**Hyper Backup échoue avec une erreur “Chiffrement de LUN” vers un disque externe.** Rappel du point technique de l’étape 8 : les LUN iSCSI ne peuvent être sauvegardés que vers un disque externe formaté en ext4, pas en Btrfs ni en NTFS/exFAT.

**Les transferts de fichiers sont anormalement lents.** Vérifiez d’abord le câble Ethernet (un câble Cat 5 limite le débit bien en dessous du Gigabit), puis le statut de la liaison agrégée si configurée, et enfin l’activité en arrière-plan (indexation Photos, scrub du volume) qui peut temporairement consommer les ressources disque.

**Le NAS chauffe ou les ventilateurs tournent en continu à pleine vitesse.** Vérifiez l’aération autour du boîtier (au moins 10 cm d’espace libre à l’arrière), dépoussiérez les grilles de ventilation, et consultez le journal des températures dans Panneau de contrôle > Ressource système.

