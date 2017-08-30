#!/bin/sh

cd "$( cd `dirname $0` && pwd )/.."

for f in  'config/config.yml'
do
    if [ ! -f $f ]; then
        cp $f.dist $f
        echo "File created from $f"
    fi
done

files[1]='tomita/area/config.proto'
files[2]='tomita/contact/config.proto'
files[3]='tomita/price/config.proto'
files[4]='tomita/type/config.proto'

for f in "${files[@]}";  do
if [ ! -f $f ]; then
    cp $f.dist $f
    echo "Created $f"
fi
done

file=$(pwd)/bin/tomita
if [ ! -f $file ]; then
 wget -O $file.bz2 'http://download.cdn.yandex.net/tomita/tomita-linux64.bz2'
 bzip2 -d $file.bz2
 chmod u+x $file
fi