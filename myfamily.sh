curl https://01.nextera.education/assets/superhero/all.json | jq --arg id "$HERO_ID" " .[] | select (.id == HERO_ID)|.connections.relatives" | tr -d '"'

