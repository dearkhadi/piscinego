curl -s https://01.tomorrow-school.ai/assets/superhero/all.json | jq --arg hero_id "$HERO_ID" '.[] | select(.id == ($hero_id|tonumber)) | .connections.relatives' | sed 's/"//g'
