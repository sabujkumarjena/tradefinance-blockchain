/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Trade Finance Use Case - WORK IN  PROGRESS
 */

package main


import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define the letter of credit
type LetterOfCredit struct {
	LCId			string		`json:"lcId"`
	BLId			string		`json:"BLId"`
	PurchaseId      string		`json:"pId"`
	BLOwner			string		`json:"blOwner"`
	ExpiryDate		string		`json:"expiryDate"`
	Buyer    string   `json:"buyer"`
	BuyerBank		string		`json:"buyerBank"`
	SellerBank		string		`json:"sellerBank"`
	Seller		string		`json:"seller"`
	Amount			int		`json:"amount,int"`
	Status			string		`json:"status"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

id, err := cid.New(APIstub)
if err != nil {
// Handle error
return shim.Error(err.Error())
}
mspid, err := id.GetMSPID()
if err != nil {
// Handle error
return shim.Error(err.Error())
}

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "requestLC" &&  mspid == "BuyerMSP" { // buyer will request to buyer bank 
		return s.requestLC(APIstub, args)
	} else if function == "issueLC" && mspid == "BuyerBankMSP" { //buyer bank will issue lc
		return s.issueLC(APIstub, args)
	} else if function == "acceptLC"  && mspid == "SellerBankMSP"{ //accept lc by seller's bank
		return s.acceptLC(APIstub, args)
	} else if function == "issueBL" && mspid == "ShippingMSP" { //by shipping company
		return s.issueBL(APIstub, args)
	} else if function == "transferBL" { // input is lcid and new owner
		return s.transferBL(APIstub, args)
	} else if function == "getLC" {
		return s.getLC(APIstub, args)
	}else if function == "getLCHistory" {
		return s.getLCHistory(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}





// This function is initiate by Buyer 
func (s *SmartContract) requestLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {


	LC := LetterOfCredit{}

	err  := json.Unmarshal([]byte(args[0]),&LC)
if err != nil {
		return shim.Error("Not able to parse args into LC")
	}
	LC.Status = "Requested"
	LCBytes, err := json.Marshal(LC)
    APIstub.PutState(LC.LCId,LCBytes)
	fmt.Println("LC Requested -> ", LC)

	

	return shim.Success(nil)
}

// This function is initiate by Seller
func (s *SmartContract) issueLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcID := struct {
		LcID  string `json:"lcID"`
	}{}
	err  := json.Unmarshal([]byte(args[0]),&lcID)
	if err != nil {
		return shim.Error("Not able to parse args into LCID")
	}
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(lcID.LcID)

	var lc LetterOfCredit

	err = json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with LC json unmarshaling")
	}

   // check status to be 
   if(lc.Status == "Requested") {
	LC := LetterOfCredit{LCId: lc.LCId, BLId: lc.BLId, PurchaseId: lc.PurchaseId, BLOwner: lc.BLOwner, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, BuyerBank: lc.BuyerBank, SellerBank: lc.SellerBank, Seller: lc.Seller, Amount: lc.Amount, Status: "Issued"}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with LC json marshaling")
	}

    APIstub.PutState(lc.LCId,LCBytes)
	fmt.Println("LC Issued -> ", LC)
   

	return shim.Success(nil)
   }
   fmt.Println("LC is not requestd -> ")
   return shim.Success(nil)

}

func (s *SmartContract) issueBL(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	input := struct {
		LcId  string `json:"lcId"`
		BLId  string `json:"BLId"`
		BLOwner string `json:"blOwner"`
	}{}
	err  := json.Unmarshal([]byte(args[0]),&input)
	if err != nil {
		return shim.Error("Not able to parse args ")
	}
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(input.LcId)

	var lc LetterOfCredit

	err = json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with LC json unmarshaling")
	}


	LC := LetterOfCredit{LCId: lc.LCId, BLId: input.BLId, PurchaseId: lc.PurchaseId, BLOwner: input.BLOwner,  ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, BuyerBank: lc.BuyerBank, SellerBank: lc.SellerBank, Seller: lc.Seller, Amount: lc.Amount, Status: lc.Status}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with LC json marshaling")
	}

    APIstub.PutState(lc.LCId,LCBytes)
	fmt.Println("BL No Generated -> ", LC)


	return shim.Success(nil)
}

func (s *SmartContract) transferBL(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	input := struct {
		LcId  string `json:"lcId"`
		BLId  string `json:"BLId"`
		BLOwner string `json:"blOwner"`
	}{}
	err  := json.Unmarshal([]byte(args[0]),&input)
	if err != nil {
		return shim.Error("Not able to parse args ")
	}
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(input.LcId)

	var lc LetterOfCredit

	err = json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with LC json unmarshaling")
	}


	LC := LetterOfCredit{LCId: lc.LCId, BLId: input.BLId, PurchaseId: lc.PurchaseId, BLOwner: input.BLOwner,  ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, BuyerBank: lc.BuyerBank, SellerBank: lc.SellerBank, Seller: lc.Seller, Amount: lc.Amount, Status: lc.Status}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with LC json marshaling")
	}

    APIstub.PutState(lc.LCId,LCBytes)
	fmt.Println("BL is transfered -> ", LC)


	return shim.Success(nil)
}

func (s *SmartContract) acceptLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcID := struct {
		LcID  string `json:"lcID"`
	}{}
	err  := json.Unmarshal([]byte(args[0]),&lcID)
	if err != nil {
		return shim.Error("Not able to parse args into LC")
	}

	LCAsBytes, _ := APIstub.GetState(lcID.LcID)

	var lc LetterOfCredit

	err = json.Unmarshal(LCAsBytes, &lc)

	if err != nil {
		return shim.Error("Issue with LC json unmarshaling")
	}


	//LC := LetterOfCredit{LCId: lc.LCId, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, Bank: lc.Bank, Seller: lc.Seller, Amount: lc.Amount, Status: "Accepted"}
	LC := LetterOfCredit{LCId: lc.LCId, BLId: lc.BLId, PurchaseId: lc.PurchaseId, BLOwner: lc.BLOwner, ExpiryDate: lc.ExpiryDate, Buyer: lc.Buyer, BuyerBank: lc.BuyerBank, SellerBank: lc.SellerBank, Seller: lc.Seller, Amount: lc.Amount, Status: "Accepted"}
	LCBytes, err := json.Marshal(LC)

	if err != nil {
		return shim.Error("Issue with LC json marshaling")
	}

    APIstub.PutState(lc.LCId,LCBytes)
	fmt.Println("LC Accepted -> ", LC)


	

	return shim.Success(nil)
}

func (s *SmartContract) getLC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0];
	
	// if err != nil {
	// 	return shim.Error("No Amount")
	// }

	LCAsBytes, _ := APIstub.GetState(lcId)

	return shim.Success(LCAsBytes)
}

func (s *SmartContract) getLCHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	lcId := args[0];
	
	

	resultsIterator, err := APIstub.GetHistoryForKey(lcId)
	if err != nil {
		return shim.Error("Error retrieving LC history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("Error retrieving LC history.")
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getLCHistory returning:\n%s\n", buffer.String())

	

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
