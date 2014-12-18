#!/bin/zsh

OUTPUT=./siteconfig.json
LOG=./siteconfig.json.log
KEYS="author body date title"

declare -A site

echo "" > $LOG
echo "{" > $OUTPUT

for siteconfig in site_config/t* ; do
  usable="yes"
  foo=${siteconfig%.txt}
  site[url]=${foo#site_config/}
  for key in $(echo $KEYS) ; do
    tmp=""
    tmp=$(grep "^$key: " $siteconfig)
    tmp=${tmp#$key: }
    if [ -z $tmp ]; then
      echo "$site[url] is missing $key" >> $LOG
      usable="no"
    fi
    site[$key]=$tmp
  done
  if [[ $usable == "yes" ]] ; then
    echo "  '$site[url]': {" >> $OUTPUT
    for key in $(echo $KEYS) ; do
      echo "    '$key': '$site[$key]'," >> $OUTPUT
    done
    echo "}," >> $OUTPUT
  fi

done

echo "}" >> $OUTPUT
exit 0
