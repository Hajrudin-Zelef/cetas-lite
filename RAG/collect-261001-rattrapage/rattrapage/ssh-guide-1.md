---
id: collect-261001-rattrapage/rattrapage/ssh-guide-1
title: "Guide SSH approfondi"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/ssh_guide.md
source_anchor: ""
source_lines: [1, 171]
sha256: 79d366c285281a70ff072af9d1f7a57c6de965d01486c6693f6d72d0336ea232
---

# Guide SSH approfondi

**Pour Zelef — chef de service systèmes, sysadmin.**
SSH est ton outil quotidien : ce guide va du protocole aux cas de production,
du `~/.ssh/config` bien tenu au durcissement complet d'un serveur, avec
dépannage, supervision, sauvegarde des clés et cas pratiques commentés.

- **Public :** administrateurs systèmes, usage quotidien intensif.
- **Plateformes :** exemples principalement Debian/Ubuntu, transposables RHEL.
- **Convention :** `client$` = machine locale, `serveur#` = machine distante (root),
  `serveur$` = machine distante (utilisateur normal).
- **Avertissement :** les clés, empreintes et adresses IP de ce guide sont
  **fictives**. Ne partage **jamais** une clé privée : ni par mail, ni dans un
  ticket, ni dans un dépôt Git. Une clé privée qui a fuité = une clé à révoquer.

> Mode d'emploi : les sections 1 à 8 posent les bases, 9 à 40 couvrent l'usage
> client avancé, 41 à 53 le serveur, 54 à 59 la supervision et la sauvegarde,
> 60 la mise à jour, 61 à 74 le dépannage, 75 à 79 les cas pratiques,
> 80 à 85 les aide-mémoire, le glossaire, le quiz et les ressources.

---

## Sommaire

1. Objectifs du guide
2. SSH : le protocole en notions
3. Chiffrement : ce qu'il faut retenir
4. Handshake : établissement d'une session (notions)
5. Méthodes d'authentification
6. Installation d'OpenSSH
7. Premier contact : `ssh`, clés d'hôte, `known_hosts`
8. Vérifier l'empreinte d'un hôte (TOFU)
9. `ssh_config` client : structure
10. `Host` et alias
11. `HostName`, `User`, `Port`
12. `IdentityFile` : une clé par hôte
13. Options client par hôte : les plus utiles
14. `Match`, `Include` et organisation du fichier
15. Générer une clé ed25519
16. RSA 4096 : quand s'en servir encore
17. Passphrase : la choisir et la gérer
18. `ssh-agent` et `ssh-add`
19. Lancer l'agent au démarrage (systemd, keychain, Windows)
20. Agent forwarding : fonctionnement et **dangers**
21. `ssh-copy-id` et déploiement des clés
22. `authorized_keys` : format
23. Restreindre par origine : `from=`
24. Clés à usage unique : `command=`, `no-pty`
25. `no-agent-forwarding`, `no-X11-forwarding`, `no-port-forwarding`
26. `ProxyJump` et bastions
27. Chaînes de sauts : plusieurs bastions
28. Multiplexage : `ControlMaster`, `ControlPath`, `ControlPersist`
29. Mesurer le gain du multiplexage
30. Tunnels locaux (`-L`)
31. Tunnels distants (`-R`)
32. Tunnels dynamiques (`-D`, SOCKS)
33. Cas concret : accéder à une UI web interne
34. Cas concret : exposer un service local vers l'extérieur
35. Tunnels : bonnes pratiques et pièges
36. SCP : copies simples
37. SFTP : transferts interactifs et batch
38. rsync over SSH
39. Optimiser les transferts (compression, algorithmes)
40. X11 forwarding : usage et prudence
41. `sshd_config` : structure et rechargement
42. Durcissement : authentification (`PermitRootLogin`, `PasswordAuthentication`)
43. Durcissement : `AllowUsers` / `AllowGroups`
44. Durcissement : `MaxAuthTries`, `MaxSessions`, `LoginGraceTime`
45. Algorithmes modernes : Ciphers, MACs, KexAlgorithms
46. Désactiver ce qui ne sert pas (X11, tunnels sélectifs)
47. Bannières et informations divulguées
48. fail2ban pour SSH
49. 2FA par TOTP
50. Clé + 2FA combinés (`AuthenticationMethods`)
51. Certificats SSH : tour d'horizon
52. Certificats : créer une CA et signer des clés
53. Certificats côté serveur : confiance et révocation
54. Supervision : les logs d'authentification
55. Audit des connexions : `last`, `who`, `ss`
56. Alertes sur connexions suspectes
57. Sauvegarder ses clés (critique)
58. Rotation des clés
59. Inventaire des clés autorisées sur le parc
60. Mettre à jour OpenSSH
61. Dépannage : la méthode (`-v`, `-vvv`)
62. Cas 1 : `Permission denied (publickey)`
63. Cas 2 : `Host key verification failed`
64. Cas 3 : la connexion freeze / timeout
65. Cas 4 : l'agent ne transmet pas la clé (forwarding)
66. Cas 5 : connexion lente à s'établir (DNS, GSSAPI)
67. Cas 6 : `Too many authentication failures`
68. Cas 7 : `WARNING: UNPROTECTED PRIVATE KEY FILE`
69. Cas 8 : `ssh-copy-id` échoue
70. Cas 9 : le tunnel ne répond pas / port déjà utilisé
71. Cas 10 : multiplexage, `ControlPath` trop long
72. Cas 11 : X11 forwarding sans `DISPLAY`
73. Cas 12 : SFTP en chroot qui échoue
74. Les erreurs classiques (tableau récapitulatif)
75. Cas pratique 1 : bastion d'équipe
76. Cas pratique 2 : clé restreinte pour déploiement (`command=`)
77. Cas pratique 3 : accès à une UI d'administration via tunnel
78. Cas pratique 4 : backup distant avec rsync over SSH
79. Cas pratique 5 : jump host multi-sites
80. Pense-bête de poche : commandes
81. Pense-bête : `ssh_config` minimal propre
82. Pense-bête : `sshd_config` durci minimal
83. Glossaire
84. Quiz (10 questions + réponses)
85. Pour aller plus loin

---

## 1. Objectifs du guide

Ce guide a trois objectifs :

1. **Maîtriser l'outil** : configurer un client SSH rapide, fiable et agréable
   (alias, multiplexage, bastions, tunnels) pour un usage quotidien intensif.
2. **Sécuriser** : durcir les serveurs (clés uniquement, 2FA, fail2ban,
   algorithmes modernes), gérer les clés comme des secrets de production.
3. **Exploiter** : superviser les accès, sauvegarder les clés, dépanner vite
   avec une méthode et des cas documentés.

À la fin, tu dois pouvoir : écrire un `~/.ssh/config` d'équipe propre,
durcir un `sshd_config` de zéro, diagnostiquer 90 % des pannes SSH en moins
de 5 minutes, et expliquer à ton équipe pourquoi l'agent forwarding est
dangereux.

## 2. SSH : le protocole en notions

SSH (Secure Shell) est un protocole réseau chiffré qui fournit :

- un **shell distant sécurisé** (le cas d'usage historique, remplaçant Telnet,
  rlogin, rsh qui transmettaient tout en clair, mots de passe inclus) ;
- un **transport générique** : copie de fichiers (SCP/SFTP), tunnels TCP,
  proxy SOCKS, X11, rsync, Git, etc. ;
- une **authentification forte** : clés cryptographiques, certificats,
  second facteur.

Architecture en trois couches (modèle du RFC 4251) :

| Couche | Rôle | Exemple |
|---|---|---|
| Transport (RFC 4253) | Chiffrement, intégrité, authentification du **serveur** | Négociation des algorithmes, échange de clés |
| Authentification utilisateur (RFC 4252) | Authentification du **client** | publickey, password, keyboard-interactive |
| Connexion (RFC 4254) | Multiplexage de **canaux** logiques | shell, exec, sftp, tcpip-forward |

Point clé : **une seule connexion TCP** transporte plusieurs canaux
(shell + SFTP + tunnels simultanés). C'est ce qui rend le multiplexage
(section 28) et les tunnels (sections 30-32) possibles.

Port par défaut : **22/TCP**. Le changer (section 42) réduit le bruit des
robots mais n'est pas une mesure de sécurité à lui seul.

## 3. Chiffrement : ce qu'il faut retenir

Sans entrer dans la cryptographie, voici le minimum opérationnel :

- **Chiffrement symétrique** (rapide, pour les données) : `chacha20-poly1305`,
  `aes256-gcm`, `aes128-ctr`. Négocié à chaque connexion.
- **Échange de clés** (pour se mettre d'accord sur la clé de session sans
  l'exposer) : `curve25519-sha256` (moderne, rapide), `diffie-hellman-group16-sha512`.
  Fournit la **confidentialité persistante** (perfect forward secrecy) : même si
  la clé privée du serveur fuit un jour, les sessions passées restent illisibles.
- **Clé d'hôte** (authentification du serveur) : le serveur prouve son identité
  avec sa clé privée d'hôte (`/etc/ssh/ssh_host_*`) ; le client vérifie avec
  l'empreinte stockée dans `~/.ssh/known_hosts`.
- **Signatures** (authentification du client par clé) : le client signe un
  défi avec sa clé privée ; le serveur vérifie avec la clé publique de
  `authorized_keys`. La clé privée **ne voyage jamais** sur le réseau.

