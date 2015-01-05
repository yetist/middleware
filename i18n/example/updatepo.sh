#!/bin/bash
LANGS=('en_US' 'zh_CN' 'zh_TW')
DOMAIN=example
SOURCES=*.go
KEYWORDS="--keyword=__ --keyword=N_ --keyword=C_:1c,2 --keyword=NC_:1c,2 --keyword=Q_:1g"
xgettext -d $DOMAIN -s $SOURCES -o ${DOMAIN}.pot -L c++ -i $KEYWORDS
for l in ${LANGS[@]}; do
	[ -d locale/$l/LC_MESSAGES ] || mkdir -p locale/$l/LC_MESSAGES
	if [ -f locale/$l/LC_MESSAGES/${DOMAIN}.po ]; then
		msgmerge -U locale/$l/LC_MESSAGES/${DOMAIN}.po ${DOMAIN}.pot
	else
		msginit -l ${l}.utf8 -o locale/$l/LC_MESSAGES/${DOMAIN}.po -i ${DOMAIN}.pot
	fi
	msgfmt -c -v -o  locale/$l/LC_MESSAGES/${DOMAIN}.mo locale/$l/LC_MESSAGES/${DOMAIN}.po
done
