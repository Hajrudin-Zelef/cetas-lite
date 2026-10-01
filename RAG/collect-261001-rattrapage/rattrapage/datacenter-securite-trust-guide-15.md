---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-15
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "arr", "incident"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [1801, 1942]
sha256: f9d06372b294f3ff2c8960f64a5d7ed14cbdb8c26008cbf08f10776c59b75858
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

- [ ] Mot de passe admin changé (20+ caractères, unique), comptes inutiles désactivés.
- [ ] **2FA** activée sur l'interface web (si supportée).
- [ ] Réseau : VLAN management dédié, pas de routage vers Internet/production.
- [ ] Firewall : seules les IP des postes d'admin autorisées.
- [ ] Protocoles : HTTPS + Redfish uniquement ; SSH si besoin justifié ; **IPMI 2.0
      chiffré** seulement si indispensable (jamais 1.5).
- [ ] Certificat TLS remplacé (PKI interne).
- [ ] SNMP : v3 uniquement, community par défaut supprimée.
- [ ] Syslog vers SIEM, alertes sur connexions.
- [ ] Firmware BMC suivi : abonnement aux bulletins de sécurité du constructeur,
      patch sous 30 jours pour les critiques.
- [ ] Montage d'images ISO : désactivé par défaut, activé ponctuellement avec traçabilité.

## 131. Checklist : firmware et secure boot

- [ ] Registre des versions firmware par serveur (automatisé via Redfish, revu
      mensuellement).
- [ ] Secure Boot **activé et vérifié** sur 100 % du parc (audit trimestriel — un
      serveur « en panne de secure boot » est un incident).
- [ ] Clés Secure Boot : inventaire des PK/KEK/db ; rotation planifiée si clés maison.
- [ ] Valeurs de référence **measured boot** à jour après chaque mise à jour firmware.
- [ ] Procédure de mise à jour : test sur pilote → déploiement par vagues → rollback
      documenté.
- [ ] Firmwares téléchargés **uniquement** depuis le constructeur (hash vérifié).

## 132. Checklist : mise en service d'un HSM

- [ ] Réception : scellés intacts, PV avec photos.
- [ ] Installation en baie verrouillée, **double alimentation** sur PDU A+B secourues
      (voir partie 8).
- [ ] Réseau : VLAN crypto dédié, IP statiques, firewall (section 35).
- [ ] Initialisation : SO PIN généré et stocké en MofN, jamais réutilisé ensuite.
- [ ] Partitions créées selon le plan (une par usage), CO nommés par partition.
- [ ] **Cérémonie** de génération des clés critiques (section 133), PV signé.
- [ ] Sauvegarde : réplication vers le HSM backup, test de restauration.
- [ ] Supervision : SNMP/syslog vers SIEM, alertes batterie/température/échecs d'auth.
- [ ] Documentation : schéma réseau, mots de passe au coffre, procédure d'urgence.
- [ ] Firmware à jour, validation FIPS du modèle **vérifiée** sur la liste CMVP.

## 133. Checklist : cérémonie de génération de clés

**Avant** :

- [ ] Script de cérémonie écrit, relu par 2 personnes, version figée.
- [ ] Participants convoqués : officiants (quorum), témoin/auditeur, scribe.
- [ ] Salle sécurisée réservée, HSM vérifié (scellés, firmware, initialisé).
- [ ] Supports de sauvegarde prêts (HSM backup, smartcards PED numérotées).

**Pendant** :

- [ ] Identité des participants vérifiée, feuille de présence signée.
- [ ] Génération **dans** le HSM (jamais d'import pour une racine), paramètres lus à
      voix haute et vérifiés (algo, taille, usage, dates).
- [ ] Empreinte de la clé publique relevée et notée par 2 personnes indépendamment.
- [ ] Sauvegarde immédiate (duplication ou MofN), vérification de la sauvegarde.
- [ ] Zéro document sensible laissé sur table ; photos interdites sauf procédure.

**Après** :

- [ ] PV rédigé sous 48 h, signé par tous les participants.
- [ ] PV archivé (2 copies, lieux séparés), n° de série et empreintes dans l'inventaire.
- [ ] Smartcards remises aux custodians contre signature.

## 134. Checklist : mise en service d'une carte FPGA

- [ ] Serveur hôte : slots PCIe x16 libres, **airflow** vérifié (carte passive !),
      alimentation dimensionnée (+190 W par V80).
- [ ] Drivers XRT/OFS installés, version épinglée dans l'image OS.
- [ ] **Bitstream signé** : vérification de la signature avant flash, version notée.
- [ ] Chiffrement/authentification du bitstream activé sur la carte.
- [ ] Tests : design de validation du constructeur OK, température en charge < seuils.
- [ ] Monitoring : température, utilisation, erreurs ECC → supervision.
- [ ] Procédure de mise à jour du bitstream écrite, avec **rollback**.
- [ ] Bitstream source versionné (git + hash), accès restreint (propriété
      intellectuelle).

## 135. Checklist : baie datacenter sécurisée

- [ ] Baie **verrouillée** (serrure par baie, pas de clé unique), panneaux obturateurs.
- [ ] Double alimentation : PDU A + PDU B sur **circuits secourus différents**
      (voir section 141).
- [ ] Câblage : étiqueté, chemins séparés courant fort/courant faible.
- [ ] Sondes température/humidité en haut et bas de baie, alertes.
- [ ] Extinction : détection incendie de la salle, consignes affichées.
- [ ] Accès : seuls les habilités, registre des interventions.
- [ ] Plan de baie à jour (U par U), affiché + version numérique.

## 136. Runbook : incident « HSM indisponible »

**Symptômes** : les applications remontent des erreurs PKCS#11, les signatures s'arrêtent.

1. **Qualifier** (5 min) : un seul HSM ou tout le cluster ? Panne réseau (VLAN crypto)
   ou panne boîtier ? Consulter la supervision (SNMP/syslog).
2. **Basculer** : si cluster HA, vérifier que le client a basculé sur le 2e membre ;
   sinon bascule manuelle documentée.
3. **Isoler** : ne pas redémarrer « pour voir » avant d'avoir les logs (un redémarrage
   peut effacer des traces, voire déclencher une zeroization si la batterie est en cause).
4. **Si panne matérielle** : appliquer la procédure constructeur (remplacement, avec
   support sous contrat — d'où l'importance du support 15-20 %, section 47).
5. **Restaurer** : recharger les clés depuis le HSM backup (procédure testée, section 138).
6. **Post-mortem** : cause racine, RTO réel vs objectif, actions correctives.
7. **Ne jamais** : restaurer une clé depuis une sauvegarde non vérifiée, ni réutiliser
   un HSM dont l'intégrité physique est douteuse sans expertise.

## 137. Runbook : compromission suspectée d'une clé

1. **Confiner** : révoquer les certificats associés (CRL/OCSP), désactiver la clé dans
   le HSM/KMS (pas la détruire tout de suite — besoin forensique).
2. **Évaluer** : périmètre (quelles données sous cette clé ? quelle période ?),
   vecteur probable (fuite de sauvegarde ? accès abusif ? malware ?).
3. **Régénérer** : nouvelle clé (cérémonie si racine), **re-chiffrer** les données
   exposées avec la nouvelle clé.
4. **Notifier** : selon les obligations (clients, autorités — avec le juriste ; voir
   l'avertissement partie 6).
5. **Post-mortem** : comment la clé a-t-elle fuité ? Corriger la cause (souvent
   organisationnelle, pas technique — voir les erreurs de la section 119).

## 138. Plan de test DR crypto : le test annuel

Une fois par an minimum, sur un **environnement isolé** :

- [ ] Restauration d'une partition HSM depuis la sauvegarde → clés utilisables ?
- [ ] Restauration du quorum MofN : réunir M custodians, reconstituer, vérifier.
- [ ] Démarrage de Vault en auto-unseal après **coupure électrique simulée** (ordre
      de démarrage HSM → Vault → applications).
- [ ] Révocation + réémission d'un certificat (la chaîne CRL/OCSP fonctionne ?).
- [ ] Déchiffrement d'une sauvegarde ancienne avec la clé archivée (la clé existe
      toujours ?).
- [ ] Chronométrer le **RTO réel** et comparer à l'objectif.
- [ ] PV de test archivé, écarts corrigés sous 90 jours.

Un DR crypto jamais testé est une fiction — et c'est le test qui révèle les
sauvegardes pourries, les mots de passe perdus et les procédures obsolètes.

# PARTIE 8 — LIEN ÉNERGIE : QUAND LA CRYPTO DÉPEND DU COURANT

> Le credo de Zelef (chef de service systèmes & énergies) : la sécurité cryptographique
> est un **service électrique** comme un autre. Un HSM sans courant, c'est une PKI à
> l'arrêt. Cette partie chiffre le lien.

## 139. Consommations : le tableau HSM/FPGA/serveurs

