#!/bin/sh

SEPARATOPION='---------------------------'
START_MESSAGE='controle local db container.'
# CURRENT=${PWD}
# プロジェクトごとにパスを調整する
TARGET=${HOME}/path/local-db
echo ${SEPARATOPION}
echo ${START_MESSAGE}

# ls ${HOME}/path
cd ${TARGET} && make serve

echo ${SEPARATOPION}
echo 'Please "make db" again if you would start or stop local db!'
