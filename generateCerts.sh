rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile TFOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile TFOrgChannel -outputCreateChannelTx ./config/tfchannel.tx -channelID lcchannel
