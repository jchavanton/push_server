//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	prid, present := query["prid"]
	if !present || len(prid) == 0 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "prid not found or invalid %v", prid)
		return
	}
	bundle_id, present := query["bundle_id"]
	if !present || len(bundle_id) == 0 {
		w.WriteHeader(404)
		fmt.Fprintf(w, "bundle_id not found or invalid %v", prid)
		return
	}
	bundle_id[0] = bundle_id[0] + ".voip"
	var from, to, callId string
	tmp := r.URL.Query()["from"]
	if len(tmp) > 0 {
		from = tmp[0]
	}
	tmp = r.URL.Query()["ci"]
	if len(tmp) > 0 {
		callId = tmp[0]
	}
	tmp = r.URL.Query()["to"]
	if len(tmp) > 0 {
		to = tmp[0]
	}
	push(prid[0], bundle_id[0], from, to,  callId, w)
}

func main() {
	version := "1.0.2"
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal("invalid http port:", err)
	}
	listen := fmt.Sprintf(":%d", port)
	fmt.Printf("push_server starting version:%s, http listening on port:%d\n", version, port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listen, nil))
}

func push(prid string, bundle_id string, from string, to string, callId string, w http.ResponseWriter) {
	authKey, err := token.AuthKeyFromFile(os.Getenv("APPLE_AUTH_KEY_FILE")) // ".p8"
	if err != nil {
		log.Fatal("token error:", err)
	}

	token := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID: os.Getenv("APPLE_KEY_ID"),
		// TeamID from developer account (View Account -> Membership)
		TeamID: os.Getenv("APPLE_TEAM_ID"),
	}

	uuid := uuid.NewString()

	payload := fmt.Sprintf(`
{
  "aps": {
    "sound": "default",
    "alert": {
      "title": "INVITE",
      "subtitle": "",
      "body": ""
    },
    "uuid": "%s",
    "from": "%s",
    "to": "%s",
    "callId": "%s"
  }
}
`, uuid, from, to, callId)
	fmt.Println(payload)
	notification := &apns2.Notification{}
	notification.DeviceToken = prid
	notification.Topic = bundle_id
	notification.Payload = []byte(payload)
	notification.PushType = apns2.PushTypeVOIP
	notification.Expiration = time.Now().Add(time.Second * 6)

	// If you want to test push notifications for builds running directly from XCode (Development), use
	// client := apns2.NewClient(cert).Development()
	// For apps published to the app store or installed as an ad-hoc distribution use Production()

	client := apns2.NewTokenClient(token).Production()

	res, err := client.Push(notification)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Printf("topic[%s] device_token[%v] apns_id[%v] %v %v\n",
		notification.Topic, notification.DeviceToken, res.ApnsID, res.StatusCode, res.Reason)

	w.WriteHeader(res.StatusCode)
	fmt.Fprintf(w, "pushed[%v][%v][%v][%v]", prid, bundle_id, res.ApnsID, res.Reason)
}
