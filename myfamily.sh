

#!/bin/bash

curl  https://01.nextera.education/assets/superhero/all.json | jq -r ".[] | select(.id == ${HERO_ID}) | .connections.relatives" |tr '\n' ' ' 
