# nxplant
A plantuml generator for Nuxeo Document Types



## Quick start
The following steps will automate the creation of 5 diagrams on a MacOS environment.

1. Run this to download and install the required files
``` sh
curl https://raw.githubusercontent.com/arturom/nxplant/main/setup.sh | sh
```

2. Start a nuxeo server on localhost:8080
3. Download the studio zip package
4. Run the script to generate the diagrams
``` sh
cd nxplant
./nxplant-draw-all.sh {/path/to/studio/package.zip}
```