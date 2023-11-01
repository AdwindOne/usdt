package usdtapi

import (
	"log"
	"usdt/rpc"
)

// OmniClient The client for rpc of omni core
type OmniClient struct {
	ConnCfg *rpc.ConnConfig
}

// NewOmniClient returns a new client instance.
func NewOmniClient(connCfg *rpc.ConnConfig) *OmniClient {
	return &OmniClient{connCfg}
}

// GetBalance Returns the token balance for a given address and property.
//
// Arguments:
//
// 1. address              (string, required) the address
//
// 2. propertyid           (number, required) the property identifier
//
// Result:
//
// "balance" : "n.nnnnnnnn",   (string) the available balance of the address
// "reserved" : "n.nnnnnnnn"   (string) the amount reserved by sell offers and accepts
func (o *OmniClient) GetBalance(address string, propertyid uint32) (balance, reserved string) {

	c := rpc.NewClient(o.ConnCfg)
	var result balanceResult

	if err := c.Call(&result, "omni_getbalance", address, propertyid); err != nil {
		log.Printf("Call %v", err)
	}

	return result.Balance, result.Reserved
}

// ListTransactions List wallet transactions, optionally filtered by an address and block boundaries.
//
// Arguments:
//
// 1. address              (string, optional) address filter (default: "*")
//
// 2. count                (number, optional) show at most n transactions (default: 10)
//
// 3. skip                 (number, optional) skip the first n transactions (default: 0)
//
// 4. startblock           (number, optional) first block to begin the search (default: 0)
//
// 5. endblock             (number, optional) last block to include in the search (default: 999999)
//
// Result:
//
// Array of Tx objects
func (o *OmniClient) ListTransactions(args ...interface{}) (result []Tx) {
	c := rpc.NewClient(o.ConnCfg)

	if err := c.Call(&result, "omni_listtransactions", args...); err != nil {
		log.Printf("Call %v", err)
	}

	return
}

// Send Create and broadcast a simple send transaction.
//
// Arguments:
//
// 1. fromaddress          (string, required) the address to send from
//
// 2. toaddress            (string, required) the address of the receiver
//
// 3. propertyid           (number, required) the identifier of the tokens to send
//
// 4. amount               (string, required) the amount to send
//
// Result:
//
// "hash"                  (string) the hex-encoded transaction hash
func (o *OmniClient) Send(fromaddress, toaddress string, propertyid uint32, amount string) (string, error) {
	c := rpc.NewClient(o.ConnCfg)
	var hash string
	if err := c.Call(&hash, "omni_send", fromaddress, toaddress, propertyid, amount); err != nil {
		//log.Printf("Call %v", err)
		return hash, err
	}

	return hash, nil
}

func (o *OmniClient) GetBlockCount() int64 {
	c := rpc.NewClient(o.ConnCfg)

	var count int64 = 0
	if err := c.Call(&count, "getblockcount"); err != nil {
		log.Printf("Call %v", err)
	}

	return count
}

func (o *OmniClient) GetTransaction(txid string) *Tx {

	c := rpc.NewClient(o.ConnCfg)
	var result Tx

	if err := c.Call(&result, "omni_gettransaction", txid); err != nil {
		log.Printf("Call %v", err)
	}

	return &result
}

//暂不实现
//== Omni layer (configuration) ==
//omni_setautocommit flag
//
//== Omni layer (data retrieval) ==
//omni_getactivations
//omni_getactivecrowdsales
//omni_getactivedexsells ( address )
//omni_getallbalancesforaddress "address"
//omni_getallbalancesforid propertyid
//omni_getbalance "address" propertyid
//omni_getbalanceshash propertyid
//omni_getcrowdsale propertyid ( verbose )
//omni_getcurrentconsensushash
//omni_getfeecache ( propertyid )
//omni_getfeedistribution distributionid
//omni_getfeedistributions propertyid
//omni_getfeeshare ( address ecosystem )
//omni_getfeetrigger ( propertyid )
//omni_getgrants propertyid
//omni_getinfo
//omni_getmetadexhash propertyId
//omni_getorderbook propertyid ( propertyid )
//omni_getpayload "txid"
//omni_getproperty propertyid
//omni_getseedblocks startblock endblock
//omni_getsto "txid" "recipientfilter"
//omni_gettrade "txid"
//omni_gettradehistoryforaddress "address" ( count propertyid )
//omni_gettradehistoryforpair propertyid propertyid ( count )
//omni_gettransaction "txid"
//omni_listblocktransactions index
//omni_listpendingtransactions ( "address" )
//omni_listproperties
//omni_listtransactions ( "address" count skip startblock endblock )
//
//== Omni layer (payload creation) ==
//omni_createpayload_cancelalltrades ecosystem
//omni_createpayload_canceltradesbypair propertyidforsale propertiddesired
//omni_createpayload_canceltradesbyprice propertyidforsale "amountforsale" propertiddesired "amountdesired"
//omni_createpayload_changeissuer propertyid
//omni_createpayload_closecrowdsale propertyid
//omni_senddexaccept propertyid "amount"
//omni_createpayload_dexsell propertyidforsale "amountforsale" "amountdesired" paymentwindow minacceptfee action
//omni_createpayload_disablefreezing propertyid
//omni_createpayload_enablefreezing propertyid
//omni_createpayload_freeze "toaddress" propertyid amount
//omni_createpayload_grant propertyid "amount" ( "memo" )
//omni_createpayload_issuancecrowdsale ecosystem type previousid "category" "subcategory" "name" "url" "data" propertyiddesired tokensperunit deadline earlybonus issuerpercentage
//omni_createpayload_issuancefixed ecosystem type previousid "category" "subcategory" "name" "url" "data" "amount"
//omni_createpayload_issuancemanaged ecosystem type previousid "category" "subcategory" "name" "url" "data"
//omni_createpayload_revoke propertyid "amount" ( "memo" )
//omni_createpayload_sendall ecosystem
//omni_createpayload_simplesend propertyid "amount"
//omni_createpayload_sto propertyid "amount" ( distributionproperty )
//omni_createpayload_trade propertyidforsale "amountforsale" propertiddesired "amountdesired"
//omni_createpayload_unfreeze "toaddress" propertyid amount
//
//== Omni layer (raw transactions) ==
//omni_createrawtx_change "rawtx" "prevtxs" "destination" fee ( position )
//omni_createrawtx_input "rawtx" "txid" n
//omni_createrawtx_multisig "rawtx" "payload" "seed" "redeemkey"
//omni_createrawtx_opreturn "rawtx" "payload"
//omni_createrawtx_reference "rawtx" "destination" ( amount )
//omni_decodetransaction "rawtx" ( "prevtxs" height )
//
//== Omni layer (transaction creation) ==
//omni_send "fromaddress" "toaddress" propertyid "amount" ( "redeemaddress" "referenceamount" )
//omni_sendall "fromaddress" "toaddress" ecosystem ( "redeemaddress" "referenceamount" )
//omni_sendcancelalltrades "fromaddress" ecosystem
//omni_sendcanceltradesbypair "fromaddress" propertyidforsale propertiddesired
//omni_sendcanceltradesbyprice "fromaddress" propertyidforsale "amountforsale" propertiddesired "amountdesired"
//omni_sendchangeissuer "fromaddress" "toaddress" propertyid
//omni_sendclosecrowdsale "fromaddress" propertyid
//omni_senddexaccept "fromaddress" "toaddress" propertyid "amount" ( override )
//omni_senddexsell "fromaddress" propertyidforsale "amountforsale" "amountdesired" paymentwindow minacceptfee action
//omni_senddisablefreezing "fromaddress" propertyid
//omni_sendenablefreezing "fromaddress" propertyid
//omni_sendfreeze "fromaddress" "toaddress" propertyid amount
//omni_sendgrant "fromaddress" "toaddress" propertyid "amount" ( "memo" )
//omni_sendissuancecrowdsale "fromaddress" ecosystem type previousid "category" "subcategory" "name" "url" "data" propertyiddesired tokensperunit deadline ( earlybonus issuerpercentage )
//omni_sendissuancefixed "fromaddress" ecosystem type previousid "category" "subcategory" "name" "url" "data" "amount"
//omni_sendissuancemanaged "fromaddress" ecosystem type previousid "category" "subcategory" "name" "url" "data"
//omni_sendrawtx "fromaddress" "rawtransaction" ( "referenceaddress" "redeemaddress" "referenceamount" )
//omni_sendrevoke "fromaddress" propertyid "amount" ( "memo" )
//omni_sendsto "fromaddress" propertyid "amount" ( "redeemaddress" distributionproperty )
//omni_sendtrade "fromaddress" propertyidforsale "amountforsale" propertiddesired "amountdesired"
//omni_sendunfreeze "fromaddress" "toaddress" propertyid amount
