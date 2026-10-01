---
id: collect-261001-general-networking/general-networking/durcir-windows-server-2025-credential-guard-2026-4
title: "Vérifier la virtualisation matérielle activée dans le firmware"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agents", "incident"]
source: docs/RAG/collect-261001-general-networking/durcir-windows-server-2025-credential-guard-2026.md
source_anchor: ""
source_lines: [207, 246]
sha256: 5c83762e0ea7563b117e4613d303f0c97a819cdb2886cdb15f653be1c2a6904e
---

# Vérifier la virtualisation matérielle activée dans le firmware

Plusieurs erreurs reviennent systématiquement dans les retours d’expérience des équipes qui déploient Credential Guard et VBS sur un parc Windows Server 2025.

- **Pilotes non signés ou obsolètes.** Les pilotes en mode noyau anciens, en particulier certains agents antivirus, clients VPN legacy ou pilotes de stockage tiers, échouent au chargement une fois HVCI activé, provoquant un écran bleu au démarrage.
- **Virtualisation imbriquée mal configurée.** Sur une VM hébergée par un hyperviseur tiers, VBS ne démarre pas si l’hôte n’expose pas correctement les extensions de virtualisation à l’invité, sans message d’erreur explicite.
- **Rupture de l’authentification legacy.** Les baselines désactivent progressivement WDigest, NTLMv1 et RC4. Une application interne encore dépendante de ces protocoles cesse de fonctionner sans avertissement préalable.
- **Activation forcée sur un contrôleur de domaine.** Microsoft exclut explicitement les contrôleurs de domaine de l’activation par défaut de Credential Guard. Forcer cette activation peut casser certains flux Kerberos spécifiques aux DC.
- **Sous-dimensionnement matériel.** VBS et HVCI consomment de la mémoire et du CPU. Sur un serveur SQL ou RDS déjà proche de la saturation, l’activation sans marge matérielle dégrade les performances perçues par les utilisateurs.
- **Verrouillage UEFI oublié en cas d’urgence.** Le mode « Activé avec verrouillage UEFI » empêche toute désactivation à distance de Credential Guard, y compris en cas de besoin légitime urgent. Documentez la procédure de désactivation physique avant d’en avoir besoin, pas pendant un incident.

## Dépannage : problèmes fréquents et solutions

Voici les incidents les plus rencontrés lors du déploiement, avec la cause probable et la correction à appliquer.

| Symptôme | Cause probable | Solution | 
|---|---|---|
| VBS reste inactif après application de la GPO | Secure Boot désactivé ou firmware en mode Legacy | Repasser le firmware en UEFI natif et activer Secure Boot avant de relancer la GPO | 
| Credential Guard indique « configuré » mais pas « en cours d’exécution » | Redémarrage manquant après application de la stratégie | Redémarrer le serveur ; Credential Guard ne s’active jamais à chaud | 
| Écran bleu au démarrage après activation de HVCI | Pilote en mode noyau incompatible | Démarrer en mode sans échec, identifier le pilote via l’observateur d’événements, mettre à jour ou retirer le pilote | 
| Get-Tpm renvoie TpmPresent : False | TPM désactivé dans le firmware ou en mode 1.2 non reconnu | Activer le TPM dans le BIOS/UEFI ou activer le TPM virtuel sur une VM Generation 2 | 
| Install-Module OSConfig échoue | Absence d’accès à PowerShell Gallery ou politique d’exécution restrictive | Configurer un proxy sortant vers PowerShell Gallery ou héberger le module sur un dépôt interne NuGet | 
| BitLocker demande la clé de récupération à chaque démarrage | Mesures TPM modifiées (mise à jour firmware, changement de configuration de boot) | Suspendre BitLocker avant toute mise à jour firmware planifiée, puis réactiver | 
| Une application métier cesse de fonctionner après la baseline | Dépendance à NTLM, WDigest ou RC4 désormais bloqués | Auditer les journaux d’authentification, migrer l’application vers Kerberos AES, ou créer une exception documentée et temporaire | 
| Get-OSConfigCompliance retourne « non conforme » après application | Réglage modifié manuellement après l’application de la baseline, ou dérive locale | Réappliquer la baseline avec Invoke-OSConfigBaseline et investiguer la source de la dérive | 
| Hyper-V refuse de s’installer | Hyperviseur tiers déjà présent ou virtualisation imbriquée non exposée | Désinstaller l’hyperviseur concurrent ou activer l’exposition de la virtualisation au niveau de l’hôte | 
| Un pilote EDR tiers plante après activation de RunAsPPL | Le pilote n’est pas signé pour s’exécuter en tant que Protected Process | Contacter l’éditeur pour une version compatible LSA Protection ou activer une exception ciblée documentée | 

## Exemple de sortie attendue après un déploiement réussi

Sur un serveur correctement durci, la commande `Get-ItemProperty -Path 'HKLM:\SYSTEM\CurrentControlSet\Control\Lsa' -Name LsaCfgFlags` doit renvoyer la valeur `1` (activé avec verrouillage UEFI). L’outil `msinfo32` ou la commande `Get-ComputerInfo | Select-Object DeviceGuardSecurityServicesRunning` doit lister à la fois `CredentialGuard` et `HypervisorEnforcedCodeIntegrity` comme services actifs. Une exécution de `Get-OSConfigCompliance -Name 'WindowsServer2025-MemberServer'` conforme affiche un statut global « Compliant » avec le détail des quelque 300 réglages contrôlés, sans écart non justifié. Si l’un de ces indicateurs manque après redémarrage, reprenez la vérification des prérequis matériels avant de recommencer les étapes de configuration.

## Conseils avancés pour la production et projet complet clé en main

En production, ne déployez jamais Credential Guard et VBS directement sur l’ensemble du parc. Constituez d’abord un groupe pilote représentatif (serveurs de fichiers, serveurs applicatifs, serveurs SQL) et laissez tourner la configuration en mode audit OSConfig pendant au moins deux semaines avant de basculer en application forcée. Utilisez le rapport de conformité pour identifier les écarts avant qu’ils ne deviennent des incidents en heures ouvrées.

Pour les environnements soumis à la directive NIS2, documentez chaque étape de ce durcissement dans votre dossier de conformité : Credential Guard et les baselines OSConfig répondent directement aux exigences de gestion des accès privilégiés et de réduction de la surface d’attaque exigées par le texte. Combinez ce travail avec une gestion rigoureuse des secrets applicatifs, par exemple via un coffre-fort centralisé, pour éviter que des identifiants de service ne circulent en clair dans des scripts de déploiement.

Prévoyez également une procédure de retour arrière documentée avant chaque vague de déploiement. Si un serveur critique refuse de redémarrer après application de la baseline OSConfig, la restauration passe par un démarrage en mode sans échec avec réseau, suivi d’une désactivation temporaire de la clé `EnableVirtualizationBasedSecurity` le temps d’identifier la cause. Conservez toujours une sauvegarde d’état système récente et testez la procédure de restauration sur l’environnement pilote avant de l’appliquer à des serveurs de production critiques. Sur le plan budgétaire, ce durcissement ne nécessite aucun investissement matériel si vos serveurs datent de moins de cinq ans : la quasi-totalité des serveurs vendus depuis 2021 intègrent nativement UEFI, Secure Boot et un TPM 2.0. Le coût réel se situe dans le temps d’ingénierie nécessaire pour tester la compatibilité applicative, généralement estimé entre deux et quatre semaines pour un parc de taille moyenne, en fonction du nombre d’applications legacy à auditer.

Voici un script consolidé qui reprend l’essentiel des étapes de ce tutoriel pour un déploiement automatisé sur un serveur membre, à adapter selon votre politique interne et à tester impérativement en préproduction avant tout usage en production :

