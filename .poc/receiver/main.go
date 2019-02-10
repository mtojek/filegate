package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pions/webrtc"
	"github.com/pions/webrtc/examples/util"
	"github.com/pions/webrtc/pkg/datachannel"
	"github.com/pions/webrtc/pkg/ice"
)

func main() {
	// Prepare the configuration
	config := webrtc.RTCConfiguration{
		IceServers: []webrtc.RTCIceServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.New(config)
	util.Check(err)

	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState ice.ConnectionState) {
		fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())

		fmt.Println(peerConnection.CurrentRemoteDescription)
	})

	peerConnection.OnSignalingStateChange(func(state webrtc.RTCSignalingState) {
		fmt.Println("OnSignalingStateChange", state)
	})

	peerConnection.OnTrack(func(track *webrtc.RTCTrack) {
		fmt.Println("OnTrack", track)
	})

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.RTCDataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label, d.ID)

		// Register channel opening handling
		d.OnOpen(func() {
			fmt.Printf("Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", d.Label, d.ID)

			for range time.NewTicker(5 * time.Second).C {
				message := util.RandSeq(15)
				fmt.Printf("Sending %s \n", message)

				err := d.Send(datachannel.PayloadString{Data: []byte(message)})
				util.Check(err)
			}
		})

		// Register message handling
		d.OnMessage(func(payload datachannel.Payload) {
			switch p := payload.(type) {
			case *datachannel.PayloadString:
				fmt.Printf("Message '%s' from DataChannel '%s' payload '%s'\n", p.PayloadType().String(), d.Label, string(p.Data))
			case *datachannel.PayloadBinary:
				fmt.Printf("Message '%s' from DataChannel '%s' payload '% 02x'\n", p.PayloadType().String(), d.Label, p.Data)
			default:
				fmt.Printf("Message '%s' from DataChannel '%s' no payload \n", p.PayloadType().String(), d.Label)
			}
		})
	})

	fmt.Println("Receiver's offer: ")
	offer, err := peerConnection.CreateOffer(nil)
	util.Check(err)
	var b bytes.Buffer
	json.NewEncoder(&b).Encode(&offer)
	fmt.Println(b.String())

	// pasted below:
	fmt.Println("Enter sender's answer:")
	reader := bufio.NewReader(os.Stdin)
	remoteDescription, _ := reader.ReadString('\n')

	var receiverOffer webrtc.RTCSessionDescription
	json.NewDecoder(bytes.NewReader([]byte(remoteDescription))).Decode(&receiverOffer)
	peerConnection.SetRemoteDescription(receiverOffer)

	// Block forever
	select {}
}
