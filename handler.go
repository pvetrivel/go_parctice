package test

import (
	"bitbucket.org/tekion/erratum"
	"bitbucket.org/tekion/tbaas/apiContext"
	mMgr "bitbucket.org/tekion/tbaas/mongoManager"
	"bitbucket.org/tekion/tbaas/tapi"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

// Ping Pong
func sayPong(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, `{ "ping": "pong" }`)
}

//Get Dealer
func getDealer(w http.ResponseWriter, r *http.Request) {

	ctx := apiContext.UpgradeCtx(r.Context())
	fmt.Println("dealerID--->", ctx.DealerID)
	findQ := bson.M{"isActive": true, "_id": ctx.DealerID}
	var dealerMaster Dealer
	err := mMgr.ReadOne(ctx.Tenant, DealerCollection, findQ, nil, &dealerMaster)
	if err != nil {
		tapi.HTTPErrorResponse(ctx, w, "Dealer", erratum.DefaultErrorCode, err)
		return
	}
	fmt.Println("Object--->", dealerMaster)
	tapi.HTTPResponse(ctx, w, http.StatusOK, "dealer by id", dealerMaster)
}

//Post Dealer
func addNewDealer(w http.ResponseWriter, r *http.Request) {
	var dealerIn DealerInput
	ctx := apiContext.UpgradeCtx(r.Context())
	err := json.NewDecoder(r.Body).Decode(&dealerIn)
	if err != nil {
		log.Println(err, " Error while decoding")
	}
	dealerMaster := createTemplate(dealerIn)
	err = mMgr.Create(ctx.Tenant, DealerCollection, dealerMaster)
	if err != nil {
		tapi.HTTPErrorResponse(ctx, w, "Dealer", erratum.DefaultErrorCode, err)
		return
	}
	tapi.HTTPResponse(ctx, w, http.StatusOK, "dealer by id", dealerMaster)
}

//PUT rollOut

func editByDealer(w http.ResponseWriter, r *http.Request) {
	updateData := updateReqBody{}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		log.Println(err, " Error while decoding")
	}
	fmt.Println("update--->", updateData)
	//u, _ := url.Parse(string(r.URL.Path))
	//nm := string(u.Path)
	//_, keys := path.Split(nm)
	//filter := bson.M{"_id": keys}
	//for i,k:=range updateMap{
	//	bson.M{i:k}
	//}
	//err=mMgr.Update(ctx.Tenant,DealerCollection,filter,)
}

//Post Dealer Communication
func addCommunication(w http.ResponseWriter, r *http.Request) {
	var dealerIn DealerInput
	ctx := apiContext.UpgradeCtx(r.Context())
	err := json.NewDecoder(r.Body).Decode(&dealerIn)
	fmt.Println("dealerIn--->", dealerIn.DealerName)
	if err != nil {
		log.Println(err, " Error while decoding")
	}
	var comm Communication
	err = mMgr.Create(ctx.Tenant, CommunicationCollection, &comm)
	if err != nil {
		tapi.HTTPErrorResponse(ctx, w, "Dealer", erratum.DefaultErrorCode, err)
		return
	}
	tapi.HTTPResponse(ctx, w, http.StatusOK, "dealer by id", comm)
}

func getCommunication(w http.ResponseWriter, r *http.Request) {

	ctx := apiContext.UpgradeCtx(r.Context())
	findQ := bson.M{"dealerID": ctx.DealerID}
	var forText Communication
	err := mMgr.ReadOne(ctx.Tenant, CommunicationCollection, findQ, nil, &forText)
	if err != nil {
		tapi.HTTPErrorResponse(ctx, w, "Dealer", erratum.DefaultErrorCode, err)
		return
	}
	tapi.HTTPResponse(ctx, w, http.StatusOK, "dealer by id", forText)
}

//func listSequence(w http.ResponseWriter,r *http.Request)  {
//	var seqList []Sequence
//	var seq map[string]string
//	tenantName:="cacar"
//	dealerName:="Sequence"
//	filter:=bson.M{}
//	seqList=mongoFind(tenantName,dealerName,filter)
//
//
//}
//func getSequence(ctx ,sequenceName string) string{
//	var seq Sequence
//
//	findQ:=bson.M{"sequenceName":sequenceName}
//	err:=mMgr.ReadOne(ctx.Tenant,SequenceColl,nil,findQ,&seq)
//	if err!=nil{
//		comm=communicationTemplate(dealerIn,1)
//
//	}else {
//		comm=communicationTemplate(dealerIn,int(seq.Next))
//		seq.Next++
//		update:=bson.M{"next":seq.Next}
//		err=mMgr.Update(ctx.Tenant,SequenceColl,findQ,update)
//	}
//}
