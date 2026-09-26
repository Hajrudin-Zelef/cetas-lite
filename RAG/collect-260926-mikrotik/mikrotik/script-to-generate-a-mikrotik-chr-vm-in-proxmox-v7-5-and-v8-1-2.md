---
id: collect-260926-mikrotik/mikrotik/script-to-generate-a-mikrotik-chr-vm-in-proxmox-v7-5-and-v8-1-2
title: "script-to-generate-a-mikrotik-chr-vm-in-proxmox-v7-5-and-v8-1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/chr-proxmox/script-to-generate-a-mikrotik-chr-vm-in-proxmox-v7-5-and-v8-1.md
source_anchor: ""
source_lines: [19, 217]
sha256: 10e660ba33ce762caacee515d56d80680543aa964a066da51b58a6344ef9b159
---

# script-to-generate-a-mikrotik-chr-vm-in-proxmox-v7-5-and-v8-1

  |  | #!/usr/bin/env bash | 
|  | function header_info { | 
|  | cat <<"EOF" | 
|  | __ ____ __ __ _ __ ____ __ ____ _____ ________ ______ | 
|  | / \|/ (_) /___________ / /_(_) /__ / __ \____ __ __/ /____ _____/ __ \/ ___/ / ____/ / / / __ \ | 
|  | / /\|_/ / / //_/ ___/ __ \/ __/ / //_/ / /_/ / __ \/ / / / __/ _ \/ ___/ / / /\__ \ / / / /_/ / /_/ / | 
|  | / / / / / ,< / / / /_/ / /_/ / ,< / _, _/ /_/ / /_/ / /_/ __/ / / /_/ /___/ / / /___/ __ / _, _/ | 
|  | /_/ /_/_/_/\|_/_/ \____/\__/_/_/\|_\| /_/ \|_\|\____/\__,_/\__/\___/_/ \____//____/ \____/_/ /_/_/ \|_\| | 
|  | EOF | 
|  | } | 
|  | clear | 
|  | header_info | 
|  | echo -e "Loading..." | 
|  | GEN_MAC=$(echo '00 60 2f'$(od -An -N3 -t xC /dev/urandom) \| sed -e 's/ /:/g' \| tr '[:lower:]' '[:upper:]') | 
|  | NEXTID=$(pvesh get /cluster/nextid) | 
|  | YW=$(echo "\033[33m") | 
|  | BL=$(echo "\033[36m") | 
|  | HA=$(echo "\033[1;34m") | 
|  | RD=$(echo "\033[01;31m") | 
|  | BGN=$(echo "\033[4;92m") | 
|  | GN=$(echo "\033[1;92m") | 
|  | DGN=$(echo "\033[32m") | 
|  | CL=$(echo "\033[m") | 
|  | BFR="\\r\\033[K" | 
|  | HOLD="-" | 
|  | CM="${GN}✓${CL}" | 
|  | set -o errexit | 
|  | set -o errtrace | 
|  | set -o nounset | 
|  | set -o pipefail | 
|  | shopt -s expand_aliases | 
|  | alias die='EXIT=$? LINE=$LINENO error_exit' | 
|  | trap die ERR | 
|  | trap cleanup EXIT | 
|  | function error_exit() { | 
|  | trap - ERR | 
|  | local reason="Unknown failure occurred." | 
|  | local msg="${1:-$reason}" | 
|  | local flag="${RD}‼ ERROR ${CL}$EXIT@$LINE" | 
|  | echo -e "$flag $msg" 1>&2 | 
|  | [ ! -z ${VMID-} ] && cleanup_vmid | 
|  | exit $EXIT | 
|  | } | 
|  | function cleanup_vmid() { | 
|  | if $(qm status $VMID &>/dev/null); then | 
|  | if [ "$(qm status $VMID \| awk '{print $2}')" == "running" ]; then | 
|  | qm stop $VMID | 
|  | fi | 
|  | qm destroy $VMID | 
|  | fi | 
|  | } | 
|  | function cleanup() { | 
|  | popd >/dev/null | 
|  | rm -rf $TEMP_DIR | 
|  | } | 
|  | TEMP_DIR=$(mktemp -d) | 
|  | pushd $TEMP_DIR >/dev/null | 
|  | if ! pveversion \| grep -Eq "pve-manager/(7.4-[1][3-9]\|8.1.[1-9])"; then | 
|  | echo "⚠ This version of Proxmox Virtual Environment is not supported" | 
|  | echo -e "Requires PVE7 Version 7.4-13 or later, or PVE8 Version 8.1.1 or later." | 
|  | echo "Exiting..." | 
|  | sleep 3 | 
|  | exit | 
|  | fi | 
|  | if (whiptail --backtitle "Proxmox VE Helper Scripts" --title "Mikrotik RouterOS CHR VM" --yesno "This will create a New Mikrotik RouterOS CHR VM. Proceed?" 10 58); then | 
|  | echo "User selected Yes" | 
|  | else | 
|  | clear | 
|  | echo -e "⚠ User exited script \n" | 
|  | exit | 
|  | fi | 
|  | function msg_info() { | 
|  | local msg="$1" | 
|  | echo -ne " ${HOLD} ${YW}${msg}..." | 
|  | } | 
|  | function msg_ok() { | 
|  | local msg="$1" | 
|  | echo -e "${BFR} ${CM} ${GN}${msg}${CL}" | 
|  | } | 
|  | function default_settings() { | 
|  | echo -e "${DGN}Using Virtual Machine ID: ${BGN}$NEXTID${CL}" | 
|  | VMID=$NEXTID | 
|  | echo -e "${DGN}Using Hostname: ${BGN}mikrotik-routeros-chr${CL}" | 
|  | HN=mikrotik-routeros-chr | 
|  | echo -e "${DGN}Allocated Cores: ${BGN}1${CL}" | 
|  | CORE_COUNT="2" | 
|  | echo -e "${DGN}Allocated RAM: ${BGN}256${CL}" | 
|  | RAM_SIZE="512" | 
|  | echo -e "${DGN}Using Bridge: ${BGN}vmbr0${CL}" | 
|  | BRG="vmbr0" | 
|  | echo -e "${DGN}Using MAC Address: ${BGN}$GEN_MAC${CL}" | 
|  | MAC=$GEN_MAC | 
|  | echo -e "${DGN}Using VLAN: ${BGN}Default${CL}" | 
|  | VLAN="" | 
|  | echo -e "${DGN}Using Interface MTU Size: ${BGN}Default${CL}" | 
|  | MTU="" | 
|  | echo -e "${DGN}Start VM when completed: ${BGN}no${CL}" | 
|  | START_VM="no" | 
|  | echo -e "${BL}Creating a Mikrotik RouterOS CHR VM using the above default settings${CL}" | 
|  | } | 
|  | function advanced_settings() { | 
|  | VMID=$(whiptail --backtitle "Proxmox VE Helper Scripts" --inputbox "Set Virtual Machine ID" 8 58 $NEXTID --title "VIRTUAL MACHINE ID" 3>&1 1>&2 2>&3) | 
|  | exitstatus=$? | 
|  | if [ $exitstatus = 0 ]; then | 
|  | echo -e "${DGN}Using Virtual Machine ID: ${BGN}$VMID${CL}" | 
|  | else | 
|  | exit | 
|  | fi | 
|  | VM_NAME=$(whiptail --backtitle "Proxmox VE Helper Scripts" --inputbox "Set Hostname" 8 58 mikrotik-routeros-chr --title "HOSTNAME" 3>&1 1>&2 2>&3) | 
|  | exitstatus=$? | 
|  | if [ $exitstatus = 0 ]; then | 
|  | HN=$(echo ${VM_NAME,,} \| tr -d ' ') | 
|  | echo -e "${DGN}Using Hostname: ${BGN}$HN${CL}" | 
|  | else | 
|  | exit | 
|  | fi | 
|  | CORE_COUNT=$(whiptail --backtitle "Proxmox VE Helper Scripts" --inputbox "Allocate CPU Cores" 8 58 2 --title "CORE COUNT" 3>&1 1>&2 2>&3) | 
|  | exitstatus=$? | 
|  | if [ $exitstatus = 0 ]; then | 
|  | echo -e "${DGN}Allocated Cores: ${BGN}$CORE_COUNT${CL}" | 
|  | else | 
|  | exit | 
|  | fi | 
|  | RAM_SIZE=$(whiptail --backtitle "Proxmox VE Helper Scripts" --inputbox "Allocate RAM in MiB" 8 58 512 --title "RAM" 3>&1 1>&2 2>&3) | 
|  | exitstatus=$? | 
|  | if [ $exitstatus = 0 ]; then | 
|  | echo -e "${DGN}Allocated RAM: ${BGN}$RAM_SIZE${CL}" | 
|  | else | 
|  | exit | 
|  | fi | 
|  | BRG=$(whiptail --backtitle "Proxmox VE Helper Scripts" --inputbox "Set a Bridge" 8 58 vmbr0 --title "BRIDGE" 3>&1 1>&2 2>&3) | 
|  | exitstatus=$? | 
|  | if [ $exitstatus = 0 ]; then | 
|  | echo -e "${DGN}Using Bridge: ${BGN}$BRG${CL}" | 
|  | else | 
|  | exit | 
|  | fi | 
|  | MAC1=$(whiptail --backtitle "Proxmox VE Helper Scripts" --inputbox "Set a MAC Address" 8 58 $GEN_MAC --title "MAC ADDRESS" 3>&1 1>&2 2>&3) | 
|  | exitstatus=$? | 
|  | if [ $exitstatus = 0 ]; then | 
|  | MAC="$MAC1" | 
|  | echo -e "${DGN}Using MAC Address: ${BGN}$MAC1${CL}" | 
|  | else | 
|  | exit | 
|  | fi | 
|  | VLAN1=$(whiptail --backtitle "Proxmox VE Helper Scripts" --inputbox "Set a Vlan(leave blank for default)" 8 58 --title "VLAN" 3>&1 1>&2 2>&3) | 
|  | exitstatus=$? | 
|  | if [ $exitstatus = 0 ]; then | 
|  | if [ -z $VLAN1 ]; then | 
|  | VLAN1="Default" VLAN="" | 
|  | echo -e "${DGN}Using Vlan: ${BGN}$VLAN1${CL}" | 
|  | else | 
|  | VLAN=",tag=$VLAN1" | 
|  | echo -e "${DGN}Using Vlan: ${BGN}$VLAN1${CL}" | 
|  | fi | 
|  | fi | 
|  | MTU1=$(whiptail --backtitle "Proxmox VE Helper Scripts" --inputbox "Set Interface MTU Size (leave blank for default)" 8 58 --title "MTU SIZE" --cancel-button Exit-Script 3>&1 1>&2 2>&3) | 
|  | exitstatus=$? | 
|  | if [ $exitstatus = 0 ]; then | 
|  | if [ -z $MTU1 ]; then | 
|  | MTU1="Default" MTU="" | 
|  | echo -e "${DGN}Using Interface MTU Size: ${BGN}$MTU1${CL}" | 
|  | else | 
|  | MTU=",mtu=$MTU1" | 
|  | echo -e "${DGN}Using Interface MTU Size: ${BGN}$MTU1${CL}" | 
|  | fi | 
|  | fi | 
|  | if (whiptail --backtitle "Proxmox VE Helper Scripts" --title "START VIRTUAL MACHINE" --yesno "Start Mikrotik RouterOS CHR VM when completed?" 10 58); then | 
|  | echo -e "${DGN}Start Mikrotik RouterOS CHR VM when completed: ${BGN}yes${CL}" | 
|  | START_VM="yes" | 
|  | else | 
|  | echo -e "${DGN}Start Mikrotik RouterOS CHR VM when completed: ${BGN}no${CL}" | 
|  | START_VM="no" | 
|  | fi | 
|  | if (whiptail --backtitle "Proxmox VE Helper Scripts" --title "ADVANCED SETTINGS COMPLETE" --yesno "Ready to create Mikrotik RouterOS VM?" 10 58); then | 
|  | echo -e "${RD}Creating Mikrotik RouterOS CHR VM using the above advanced settings${CL}" | 
|  | else | 
|  | clear | 
|  | header_info | 
|  | echo -e "${RD}Using Advanced Settings${CL}" | 
|  | advanced_settings | 
|  | fi | 
|  | } | 
|  | function start_script() { | 
|  | if (whiptail --backtitle "Proxmox VE Helper Scripts" --title "SETTINGS" --yesno "Use Default Settings?" --no-button Advanced 10 58); then | 
|  | clear | 
|  | header_info | 
|  | echo -e "${BL}Using Default Settings${CL}" | 
|  | default_settings | 
|  | else | 
|  | clear | 
|  | header_info | 
|  | echo -e "${RD}Using Advanced Settings${CL}" | 
|  | advanced_settings | 
|  | fi | 
|  | } | 
|  | start_script | 
|  | msg_info "Validating Storage" | 
|  | while read -r line; do | 
