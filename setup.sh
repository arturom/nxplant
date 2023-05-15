mkdir -p nxplant
cd nxplant

curl -sf http://goblin.barelyhuman.xyz/github.com/arturom/nxplant | sh
# curl https://gobinaries.com/arturom/nxplant | sh

curl -LO https://raw.githubusercontent.com/arturom/nxplant/main/nxplant-draw-all.sh
chmod +x ./nxplant-draw-all.sh

curl -LO https://github.com/plantuml/plantuml/releases/latest/download/plantuml.jar