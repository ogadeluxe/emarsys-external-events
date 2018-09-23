package library

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ogadeluxe/emarsys-external-events/config"
	"github.com/ogadeluxe/emarsys-external-events/model"
)

const letterStreams = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Response format received from Emarsys API
type Response struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

// Contact : is
type Contact struct {
	ExternalID string `json:"external_id"`
	Data       struct {
		Wishlist []*model.Wishlist `json:"wishlist"`
	} `json:"data"`
}

// WishlistPayload : structure of wl payload
type WishlistPayload struct {
	KeyID    int8       `json:"key_id"`
	Contacts []*Contact `json:"contacts"`
}

// Generates Nonce: used on X-WSSE header
func generateNonce(length int) string {
	arrBytes := make([]byte, length)
	for i := range arrBytes {
		arrBytes[i] = letterStreams[rand.Int63()%int64(len(letterStreams))]
	}
	return string(arrBytes)
}

// SetHeader :  Is used to set request header
// being sent to Emarsys API
func setHeader(request *http.Request) {
	timestamp := time.Now().Format(time.RFC3339)
	nonce := generateNonce(36)

	// generates password digest
	// (nonce + timestamp + secret)
	credential := (nonce + timestamp + config.Items.Secret)

	hexIns := sha1.New()
	hexIns.Write([]byte(credential))
	sha1 := hex.EncodeToString(hexIns.Sum(nil))
	digest := base64.StdEncoding.EncodeToString([]byte(sha1))

	xWsse := string(" UsernameToken Username=\"" + config.Items.User +
		"\",PasswordDigest=\"" + digest +
		"\",Nonce=\"" + nonce +
		"\",Created=\"" + timestamp + "\"")

	request.Header.Set("X-WSSE", xWsse)
	request.Header.Set("Content-Type", "application/json")
}

// TriggerEvent : used for triggering Emarsys external event
// through request sent to its API
func TriggerEvent(method string, event string, body string) Response {
	// build request
	url := config.Items.BaseURL + event
	request, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	// set request header
	setHeader(request)

	// set client and send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	status := response.Status
	responseBody, _ := ioutil.ReadAll(response.Body)
	return Response{status, string(responseBody)}
}

// BuildWLPayload : build payload to be sent to Emarsys API
func BuildWLPayload(dataStore model.Datastore) string {
	wl, err := dataStore.AllWishlist()
	if err != nil {
		log.Panic(err.Error())
	}
	payload := WishlistPayload{}
	payload.KeyID = config.Items.KeyID

	contact := Contact{}
	contact.ExternalID = "mootaro89@yahoo.com"
	contact.Data.Wishlist = wl
	payload.Contacts = append(payload.Contacts, &contact)

	data, err := json.Marshal(payload)
	if err != nil {
		log.Panic("error on Marshal: ", err.Error())
		return ""
	}

	return string(data)
}
