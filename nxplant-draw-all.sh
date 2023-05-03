zipFile=$1
project=$(basename $1 .zip)

tmpDir=$project/tmp
outDir=$project/out

rm -r $tmpDir
rm -r $outDir
mkdir -p $tmpDir
mkdir -p $outDir

open $outDir

extensionsPath=OSGI-INF/extensions.xml
extensionsFile=$tmpDir/$extensionsPath
typesFile=$tmpDir/types.json
schemasFile=$tmpDir/schemas.json


unzip $zipFile '*.jar' -d $tmpDir
unzip $tmpDir/*.jar $extensionsPath -d $tmpDir

curl -su Administrator:Administrator 'http://localhost:8080/nuxeo/api/v1/config/types/' > $typesFile
curl -su Administrator:Administrator 'http://localhost:8080/nuxeo/api/v1/config/schemas/' > $schemasFile

export PLANTUML_LIMIT_SIZE=20480

nxplant -extensions $extensionsFile > $outDir/diagram-1-folder-structure.pu
nxplant -folders $extensionsFile > $outDir/diagram-2-folder-structure.pu
nxplant -schemas $schemasFile -types $typesFile > $outDir/diagram-3-schemas-doctypes.pu
nxplant -schemas $schemasFile > $outDir/diagram-4-schemas.pu
nxplant -types $typesFile > $outDir/diagram-5-doctypes.pu

rm -r $tmpDir

java -jar ./bin/plantuml.jar -tpng $outDir
