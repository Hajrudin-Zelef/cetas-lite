---
id: collect-261001-rattrapage/rattrapage/glpi-guide-12
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "memory"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [1954, 2126]
sha256: 361778252849a6d509886c1c3ad5fc40598539c8f033aabf9d0d86048dd7040c
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

`Configuration > Générale > Mode débogage` : à activer **temporairement** pour
reproduire l'erreur avec un compte de test — jamais en production ouverte.

---

## 55. Dépannage — le cron ne tourne pas

**Symptômes :** pas de courriels, SLA figés, collecteur muet, tickets récurrents
non créés.

### Checklist

1. Le cron système existe-t-il ?
   ```bash
   cat /etc/cron.d/glpi
   sudo grep CRON /var/log/syslog | grep -i glpi | tail
   ```
2. Test manuel :
   ```bash
   sudo -u www-data php /var/www/glpi/front/cron.php; echo "code=$?"
   ```
   - code ≠ 0 → lire l'erreur affichée.
   - silencieux mais rien ne change → voir 3.
3. `Configuration > Actions automatiques` : chaque action a-t-elle une
   **dernière exécution récente** ? Une action en « erreur » bloque parfois les
   suivantes → consultez son journal, corrigez, puis « Exécuter ».
4. Conflit de **verrou** : `SELECT * FROM glpi_crontasks WHERE state=2;`
   (2 = en cours) — un cron tué brutalement laisse un verrou ; repassez à 1.
5. PHP CLI vs PHP web : vérifiez que le **CLI** a les bonnes extensions
   (`php -m`) et le bon `php.ini` (`php --ini`).

> **Piège :** deux crons (système + pseudo-cron web) qui se marchent dessus.
> Choisissez **un seul** mode (section 9).

---

## 56. Dépannage — les courriels ne partent pas

### Checklist ordonnée

1. **File d'attente** : `Configuration > Notifications > File d'attente des
   notifications` — des courriels en attente ? Leur âge ? Un message d'erreur
   SMTP s'affiche par ligne (authentification, connexion refusée...).
2. **Cron** : `queuednotification` tourne-t-il ? (section 55)
3. **Paramètres SMTP** : hôte, port, chiffrement, identifiants — refaites le
   **courriel de test** (section 8.5).
4. **Réseau** : depuis le serveur,
   ```bash
   nc -zv smtp.entreprise.lan 587
   openssl s_client -connect smtp.entreprise.lan:587 -starttls smtp
   ```
5. **Notifications activées ?** `Configuration > Notifications > Notifications` :
   l'événement est-il actif **pour ce profil/entité** ? Un modèle désactivé =
   silence total.
6. **Anti-spam** : vérifiez les spams du destinataire + SPF/DKIM du domaine
   d'envoi (`glpi@entreprise.lan` doit être autorisé à envoyer).
7. **Collecteur (sens inverse)** : si les mails **entrants** ne créent pas de
   tickets → tester la connexion IMAP dans le collecteur, vérifier
   `mailcollector` dans les actions automatiques.

---

## 57. Dépannage — 7 autres cas concrets

### Cas 1 — « Impossible de se connecter » après import LDAP

- Vérifiez le DN de connexion et le mot de passe du compte de service
  (test : `ldapsearch -x -D "CN=..." -W -h ad.entreprise.lan -b "DC=..."`).
- Champ identifiant : `sAMAccountName` (AD) vs `uid` (OpenLDAP).
- L'utilisateur existe-t-il dans la base LDAP ciblée par le filtre ?

### Cas 2 — Lenteurs générales

- Activez OPcache (déjà en annexe A) ; `memory_limit` à 256M mini.
- Purgez les vieux journaux (`Administration > Maintenance > Purge`).
- Vérifiez la taille de `glpi_queuednotifications` et `glpi_logs` (tables
  géantes = requêtes lentes) ; optimisez : `mysqlcheck -o glpidb`.
- Trop d'onglets/acteurs sur les tickets ? Limitez les notifications.

### Cas 3 — Pièces jointes impossibles à téléverser

- `upload_max_filesize` / `post_max_size` (php.ini) vs taille du fichier.
- Droits d'écriture sur `files/` (`www-data`).
- Quota disque du volume.

### Cas 4 — PDF illisibles / caractères bizarres

- Vérifiez `utf8mb4` sur la base **et** la connexion (`config_db.php`).
- Après import d'un dump en latin1 : convertissez
  (`ALTER TABLE ... CONVERT TO CHARACTER SET utf8mb4`).

### Cas 5 — Sessions qui expirent sans arrêt

- `session.gc_maxlifetime` trop court ; espace `/var/lib/php/sessions` plein.
- Plusieurs onglets + bascule HTTP/HTTPS (cookie `secure`) → restez en HTTPS.

### Cas 6 — Doublons d'inventaire après réinstallation de postes

- Voir section 36 : verrous + règles de rapprochement sur n° de série.
- Fusion manuelle des doublons, puis verrouillage.

### Cas 7 — « Vous n'avez pas les droits » alors que tout semble bon

- Profil **non récursif** sur la bonne entité ? L'utilisateur a-t-il le profil
  sur **l'entité du ticket** ?
- Videz le cache des droits : déconnexion/reconnexion.
- Vérifiez les **règles d'affectation** qui auraient écrasé le profil.

---

## 58. 10 erreurs classiques (et comment les éviter)

| # | Erreur | Conséquence | Prévention |
|---|---|---|---|
| 1 | Mots de passe par défaut conservés | Compromission triviale | Section 8.2 — changer dès l'install |
| 2 | Pas de cron configuré | Aucun mail, SLA morts | Section 9 — cron dès le jour 1 |
| 3 | Dossier `install/` oublié | Réinstallation / fuite d'infos | Section 8.1 |
| 4 | URL appli fausse (http/ip) | Liens morts dans les mails | Section 8.3 |
| 5 | Tickets sans élément associé | Stats inutilisables | Gabarit obligatoire (section 18) |
| 6 | Catégories en vrac (« Autre » à 60 %) | Pilotage impossible | Taxonomie sobre + revue annuelle (16) |
| 7 | SLA calculés en heures calendaires | Engagements intenables la nuit | Calendrier heures ouvrées (21) |
| 8 | Sauvegarde jamais testée | Découverte du désastre le jour J | Test trimestriel (49) |
| 9 | Plugins installés à l'aveugle | Page blanche à la MAJ | Règles de la section 51 |
| 10 | Mots de passe notés dans les tickets | Fuite de credentials | Coffre dédié, règle d'équipe (42) |

---

# PARTIE XII — CAS PRATIQUES

## 59. Cas pratique 1 — organiser le SAV d'une entreprise de maintenance copieurs

**Contexte :** 8 techniciens, 2 agences, 350 copieurs chez 120 clients, contrats
au coût/copie et au forfait.

### Semaine 1 — fondations

1. Installer GLPI (partie I), sécuriser (partie VIII), configurer le cron.
2. Créer l'arborescence d'entités : racine → 2 agences → 5 gros clients en
   sous-entités (les autres restent au niveau agence avec un champ « client »).
3. Créer les profils : Technicien SAV, Superviseur, Déclarant client, Magasinier.
4. Importer les utilisateurs depuis l'AD ; règles d'affectation automatiques.

### Semaine 2 — référentiels

5. Catégories de tickets (section 16) + gabarit « Intervention copieur » (annexe D).
6. Créer les fiches des 350 copieurs : import CSV
   (`Administration > Maintenance > Import`) — colonnes : marque, modèle, série,
   client, compteur, contrat. **Vérifier un échantillon de 20 fiches à la main.**
7. Créer les contrats et les lier aux copieurs ; renseigner les SLA.
8. Créer les modèles de cartouches + stock initial (comptage physique).

### Semaine 3 — processus

9. Règles d'assignation automatiques (section 20).
10. Tickets récurrents : relevés de compteurs mensuels, visites préventives
    trimestrielles.
11. Modèles de notification relus et testés ; collecteur `sav@entreprise.lan`.
12. Formation : 2 h techniciens (créer/traiter un ticket), 1 h déclarants clients.

### Semaine 4 — pilote puis généralisation

13. Pilote sur 1 agence / 10 clients pendant 2 semaines.
14. Revue : catégories utilisées ? champs remplis ? SLA réalistes ?
15. Ajustements, puis généralisation. **Date butoir** : l'ancien système
    (cahier, Excel) s'arrête — pas de double saisie au-delà d'un mois.

---

## 60. Cas pratique 2 — tickets d'intervention copieurs de A à Z

**Scénario :** la mairie (gros client) signale un bourrage récurrent sur son
Kyocera TASKalfa 4054ci (bureau 204).

