echo "Setting up the network.."

echo "Creating channel genesis block.."

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=BuyerBankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyerBank.tf.com/users/Admin@buyerBank.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.buyerBank.tf.com:7051" cli peer channel create -o orderer.tf.com:7050 -c lcchannel -f /etc/hyperledger/configtx/tfchannel.tx


sleep 5

echo "Channel genesis block created."

echo "peer0.buyerBank.tf.com joining the channel..."
# Join peer0.B=bank.tf.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=BuyerBankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyerBank.tf.com/users/Admin@buyerBank.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.buyerBank.tf.com:7051" cli peer channel join -b lcchannel.block

docker exec -e "CORE_PEER_LOCALMSPID=SellerBankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/sellerBank.tf.com/users/Admin@sellerBank.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.sellerBank.tf.com:7051" cli peer channel join -b lcchannel.block

echo "peer0.bank.tf.com joined the channel"

echo "peer0.buyer.tf.com joining the channel..."

# Join peer0.buyer.tf.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=BuyerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyer.tf.com/users/Admin@buyer.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.buyer.tf.com:7051" cli peer channel join -b lcchannel.block

echo "peer0.buyer.tf.com joined the channel"

echo "peer0.seller.tf.com joining the channel..."
# Join peer0.seller.tf.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=SellerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/seller.tf.com/users/Admin@seller.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.seller.tf.com:7051" cli peer channel join -b lcchannel.block
sleep 5

echo "peer0.seller.tf.com joined the channel"

echo "Following is the docker network....."

docker exec -e "CORE_PEER_LOCALMSPID=ShippingMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shipping.tf.com/users/Admin@shipping.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.shipping.tf.com:7051" cli peer channel join -b lcchannel.block
docker ps
