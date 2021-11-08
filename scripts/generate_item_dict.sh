#/usr/bin/env bash

curl 'https://ddragon.leagueoflegends.com/cdn/11.13.1/data/en_US/item.json' | jq '.data | to_entries[] | {(.key): .value.name}' | grep [0-9] | sed 's/\([[:digit:]]*\)/\1/;s/  \(.*\)/\1,/'

