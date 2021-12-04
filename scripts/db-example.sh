#!/bin/sh

# CURRENT_DIR=$(cd $(dirname $0); pwd)
SEPARATOPION='---------------------------'
START_MESSAGE='controle local db container.'

# プロジェクトごとにパスを調整する
TARGET=${HOME}/path/local-db

# @param {string} message
showMessage() {
  echo ${SEPARATOPION}
  echo $1
}

showMessage ${START_MESSAGE}

# ls ${HOME}/path
cd ${TARGET} && make serve

showMessage 'Please "make db" again if you would start or stop local db!'
