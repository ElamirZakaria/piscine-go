curl https://01.nextera.education/assets/superhero/all.json | jq ".[] | select (.id == 170) |.name,.powerstats.power,.appearance.gender"
