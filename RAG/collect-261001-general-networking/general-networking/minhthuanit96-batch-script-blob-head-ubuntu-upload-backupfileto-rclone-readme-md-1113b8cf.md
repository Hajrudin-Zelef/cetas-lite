---
id: collect-261001-general-networking/general-networking/minhthuanit96-batch-script-blob-head-ubuntu-upload-backupfileto-rclone-readme-md-1113b8cf
title: "minhthuanit96-batch-script-blob-head-ubuntu-upload-backupfileto-rclone-readme-md-1113b8cf"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/minhthuanit96-batch-script-blob-head-ubuntu-upload-backupfileto-rclone-readme-md-1113b8cf.md
source_anchor: ""
source_lines: [1, 44]
sha256: a2f63a737d60ae5222c199ff95ea1c8b26724bb3e85f0c79e4b28382eeb84a2d
---

# minhthuanit96-batch-script-blob-head-ubuntu-upload-backupfileto-rclone-readme-md-1113b8cf

This script automates the process of backing up a UniFi Network Controller, uploading the backups to a remote storage provider using Rclone, and sending notifications to a Telegram channel.
- Automated Backups: Automatically detects new UniFi backup files (.unf ).
- Rclone Integration: Securely uploads backups to any cloud storage provider supported by Rclone.
- Telegram Notifications: Sends real-time notifications to a Telegram chat to monitor the backup status.
- Automated Cleanup: Keeps the two newest backups and deletes older ones, both locally and on the remote storage, to save space.
- Logging: Maintains a detailed log file for debugging and auditing purposes.
- Error Handling: Exits immediately if any command fails and sends a failure notification.
- Easy Configuration: All settings are managed through a simple .env file.
Before using this script, ensure you have the following tools installed on your system:
- Rclone: For uploading files to remote storage.
- rsync: For efficiently syncing files.
- curl: For sending Telegram notifications.
You can check if these tools are installed by running:
command -v rclone
command -v rsync
command -v curl
- 
Clone the repository: git clone https://github.com/your-username/your-repository.git
cd your-repository
- 
Create the configuration file: Rename the .env.txt file to.env and customize the variables:mv .env.txt .env
- 
Edit the .env file:
Open the.env file and set the following variables:
  - BACKUP_DIR : The absolute path to the directory where the UniFi Controller saves its automatic backups (e.g.,/var/lib/unifi/backup/autobackup ).
  - UPLOAD_FOLDER : A temporary local directory to stage the backups before uploading (e.g.,/home/user/unifi-backups ).
  - REMOTE_DIR : The destination directory on your Rclone remote (e.g.,gdrive:UniFi-Backups ).
  - LOG_FILE : The absolute path to the log file (e.g.,/var/log/unifi-backup.log ).
  - TELEGRAM_BOT_TOKEN : Your Telegram bot token.
  - CHAT_ID : The ID of the Telegram chat where you want to receive notifications.
  - MESSAGE_THREAD_ID : (Optional) The message thread ID if you are using topics in your Telegram group.
To run the backup script manually, execute the following command:
bash backup-unifi-script.sh
To automate the backup process, you can schedule the script to run at regular intervals using a cron job.
- 
Open the crontab editor: crontab -e
- 
Add a new cron job: Add the following line to run the script every day at 2:00 AM: 0 2 * * * /path/to/your/repository/backup-unifi-script.sh Replace /path/to/your/repository/ with the actual path to the script.
- The script is executed by a cron job or manually.
- It checks for new .unf files in theBACKUP_DIR .
- If new backup files are found, they are copied to the UPLOAD_FOLDER .
- The script then uses Rclone to upload the .unf files from theUPLOAD_FOLDER to theREMOTE_DIR .
- After a successful upload, the script performs a cleanup, deleting older backups from both the UPLOAD_FOLDER and theREMOTE_DIR , keeping only the two most recent backups.
- A notification is sent to the configured Telegram chat, indicating the success or failure of the backup process.
