#!/bin/bash

#VERSION="97e9fe30219e60092e107651abb379a38b342921"
#VERSION="v1.2.116"
VERSION="master"

datafile(){
	url="$1"
	name="$2"

curl --progress-bar "$url" | awk -v name="$name" -v url="$url" '
BEGIN{
	printf("package crawlerdetect\n\n", name)
	printf("// auto-generated based on %s\n", url)
	printf("var %s = []string{\n", name)
}
{
	printf("	`%s`,\n", $0)
}
END{
	printf("}\n")
}
' > "${name}.go"
}
echo "Updating crawlers.go"
datafile "https://raw.githubusercontent.com/JayBizzle/Crawler-Detect/${VERSION}/raw/Crawlers.txt" "crawlers"
echo "Updating exclusions.go"
datafile "https://raw.githubusercontent.com/JayBizzle/Crawler-Detect/${VERSION}/raw/Exclusions.txt" "exclusions"

echo "Patching files for re2 engine"
# XXX: golang re2 based regex engine does not support negative lookahead Yandex(?!Search)
# as workaround add YandexSearch to the exclusions list and add Yandex to the crawlers list
sed -i "" -e 's/`Yandex(?!Search)`/`Yandex`/' crawlers.go
sed -i "" -e '/^}$/i\
	`YandexSearch`,' exclusions.go

echo "Updating testdata/crawlers.txt"
curl --progress-bar -o testdata/crawlers.txt https://raw.githubusercontent.com/JayBizzle/Crawler-Detect/${VERSION}/tests/data/user_agent/crawlers.txt
echo "Updating testdata/devices.txt"
curl --progress-bar -o testdata/devices.txt https://raw.githubusercontent.com/JayBizzle/Crawler-Detect/${VERSION}/tests/data/user_agent/devices.txt

echo "Updating completed"
