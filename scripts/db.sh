#!/bin/sh

SEPARATOPION='+++++++++++++++++++++++++++'
START_MESSAGE='controle local db container.'
# CURRENT=${PWD}
# プロジェクトごとにパスを調整する
TARGET=${HOME}/dev/local-db
echo ${SEPARATOPION}
echo ${START_MESSAGE}

# ls ${HOME}/dev
cd ${TARGET} && make serve

echo ${SEPARATOPION}
echo 'Please "make db" again if you would start or stop local db!'
