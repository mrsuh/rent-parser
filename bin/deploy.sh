#!/bin/sh

cd "$(cd `dirname $0` && pwd)/.."

for f in 'config/config.yml' 'tomita/price/config.proto' 'tomita/type/config.proto';  do
if [ ! -f $f ]; then
    cp $f.dist $f
    echo "File $f created from $f.dist"
fi
done

file=$(pwd)/bin/tomita
if [ ! -f $file ]; then
 wget -O $file.bz2 'http://download.cdn.yandex.net/tomita/tomita-linux64.bz2'
 bzip2 -d $file.bz2
 chmod u+x $file
fi