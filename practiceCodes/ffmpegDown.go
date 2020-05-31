package main

import (
	"log"

	"golift.io/ffmpeg"
)

func main() {
	/* Example non-transcode direct-save from securityspy. */

	securitypsy := "rtsp://admin:ULTRA%40786@marsadubai.dyndns.org:558/LiveChannel/1/media.smp"
	output := "test.mp4"
	c := &ffmpeg.Config{
		FFMPEG: "/usr/bin/ffmpeg",
		Copy:   true,  // do not transcode
		Audio:  false, // retain audio stream
		Time:   10,    // 10 seconds
	}
	encode := ffmpeg.Get(c)
	cmd, out, err := encode.SaveVideo(securitypsy, output, "test")
	log.Println("Command Used:", cmd)
	log.Println("Command Output:", out)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Saved file from", securitypsy, "to", output)

	/* Example transcode from a Dahua IP camera. */

	// dahua := "rtsp://admin:password@192.168.1.12/live"
	// output = "/tmp/dahua_captured_file.m4v"
	// f := ffmpeg.Get(&ffmpeg.Config{
	// 	Audio:  true, // retain audio stream
	// 	Time:   10,   // 10 seconds
	// 	Width:  1920,
	// 	Height: 1080,
	// 	CRF:    23,
	// 	Level:  "4.0",
	// 	Rate:   5,
	// 	Prof:   "baseline",
	// })
	// cmd, out, err = f.SaveVideo(dahua, output, "DahuaVideoTitle")
	// log.Println("Command Used:", cmd)
	// log.Println("Command Output:", out)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println("Saved file from", dahua, "to", output)
}
