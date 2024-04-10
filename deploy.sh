#!/bin/bash

INSTALL_PREFIX="/opt/halo"

# declare -a push_server_instances=("hostname1" "hostname2")
source hosts.sh

deploy_push_server_config() {
	ROLE="push_server"
	INSTALL_DIR="${INSTALL_PREFIX}/${ROLE}"
	for i in "${push_server_instances[@]}"
	do
		if [ "$1" != "all" ] && [ "$1" != "$i" ] ; then continue; fi
		printf "\nuploading to [$i]\n"
		ssh $i sudo apt install -y sqlite3 docker.io
		ssh $i "sudo mkdir -p $INSTALL_DIR/cert && sudo chmod -R 777 $INSTALL_DIR \
		        && sudo chmod -R 777 $INSTALL_DIR"
		scp * $i:$INSTALL_DIR
		scp cert/* $i:$INSTALL_DIR/cert
		ssh $i "sudo chown -R root.root $INSTALL_DIR"
		done
}

instruction() {
	printf  "\nYou can specify a host name :\n\n"
	for i in "${push_server_instances[@]}"
	do
		echo "./deploy.sh $i"
	done
}

TARGET="$1"
if [ "${TARGET}" == "" ]
then
	instruction
	exit
fi

deploy_push_server_config ${TARGET}
