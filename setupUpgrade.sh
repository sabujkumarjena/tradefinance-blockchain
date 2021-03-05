
echo "Installing tf chaincode to peer0.bank.tf.com..."
echo $COMPOSE_PROJECT_NAME
# install chaincode
# Install code on bank peer
docker exec -e "CORE_PEER_LOCALMSPID=BuyerBankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyerBank.tf.com/users/Admin@buyerBank.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.buyerBank.tf.com:7051" cli peer chaincode install -n tf -v 6.0 -p github.com/tf/go -l golang

docker exec -e "CORE_PEER_LOCALMSPID=SellerBankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/sellerBank.tf.com/users/Admin@sellerBank.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.sellerBank.tf.com:7051" cli peer chaincode install -n tf -v 6.0 -p github.com/tf/go -l golang

echo "Installed tf chaincode to peer0.bank.tf.com"

echo "Installing tf chaincode to peer0.buyer.tf.com...."

# Install code on buyer peer
docker exec -e "CORE_PEER_LOCALMSPID=BuyerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyer.tf.com/users/Admin@buyer.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.buyer.tf.com:7051" cli peer chaincode install -n tf -v 6.0 -p github.com/tf/go -l golang

echo "Installed tf chaincode to peer0.buyer.tf.com"

echo "Installing tf chaincode to peer0.seller.tf.com..."
# Install code on seller peer
docker exec -e "CORE_PEER_LOCALMSPID=SellerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/seller.tf.com/users/Admin@seller.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.seller.tf.com:7051" cli peer chaincode install -n tf -v 6.0 -p github.com/tf/go -l golang

sleep 5
# docker exec -e "CORE_PEER_LOCALMSPID=SellerBankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/sellerBank.tf.com/users/Admin@sellerBank.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.sellerBank.tf.com:7051" cli peer chaincode install -n tf -v 6.0 -p github.com/tf/go -l golang

sleep 5

sleep 5
docker exec -e "CORE_PEER_LOCALMSPID=ShippingMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/shipping.tf.com/users/Admin@shipping.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.shipping.tf.com:7051" cli peer chaincode install -n tf -v 6.0 -p github.com/tf/go -l golang

echo "Installed tf chaincode to peer0.buyer.tf.com"

echo "Instantiating tf chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=BuyerBankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/buyerBank.tf.com/users/Admin@buyerBank.tf.com/msp" -e "CORE_PEER_ADDRESS=peer0.buyerBank.tf.com:7051" cli peer chaincode upgrade -o orderer.tf.com:7050 -C lcchannel -n tf -l golang -v 6.0 -c '{"Args":[""]}' -P "OR ('ShippingMSP.member','SellerBankMSP.member','BuyerBankMSP.member','BuyerMSP.member','SellerMSP.member')"

echo "Instantiated tf chaincode."

echo "Following is the docker network....."

docker ps
