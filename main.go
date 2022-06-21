package main

import (
	"bilibili-redpocket-job/app/api/bilibili"
	"encoding/json"
	"fmt"
	"github.com/imroc/req"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	cookie = "b_lsid=10FEFB4AC_181817E03D9; _uuid=3389E588-5172-9398-EE4E-E727B7423510328349infoc; LIVE_BUVID=AUTO8916557349285250; buvid3=489B8A0E-A7BB-4295-9430-E123D68B0520167638infoc; buvid4=2CA5EE1D-F217-646D-9C3D-61EF17713B6928772-022062022-ecjYyAkrfygRX3mNCqZx1M2AVWe5cFB1uY4IxU4/m+1TUTLWOiKEmg%3D%3D; fingerprint=22d6e6994df5387a5c00dd9366b7bbab; buvid_fp_plain=undefined; SESSDATA=ad29a25f%2C1671286956%2C5fac6%2A61; bili_jct=b1c511dedcae6fdd22caa1b852ba3665; DedeUserID=1924897562; DedeUserID__ckMd5=0f1dee9b2ac5e857; sid=c85n9ee4; _dfcaptcha=826f5d9149aa6f620a4d81e3a4bf35ff; buvid_fp=22d6e6994df5387a5c00dd9366b7bbab; b_timer=%7B%22ffp%22%3A%7B%22444.7.fp.risk_489B8A0E%22%3A%22181817E04C2%22%2C%22333.42.fp.risk_489B8A0E%22%3A%22181817E1779%22%2C%22444.8.fp.risk_489B8A0E%22%3A%22181817EA1BC%22%7D%7D; PVID=8"
	//cookie           = "b_lsid=10FEFB4AC_181817E03D9; _uuid=3389E588-5172-9398-EE4E-E727B7423510328349infoc; LIVE_BUVID=AUTO8916557349285250; buvid3=489B8A0E-A7BB-4295-9430-E123D68B0520167638infoc; buvid4=2CA5EE1D-F217-646D-9C3D-61EF17713B6928772-022062022-ecjYyAkrfygRX3mNCqZx1M2AVWe5cFB1uY4IxU4/m+1TUTLWOiKEmg%3D%3D; fingerprint=22d6e6994df5387a5c00dd9366b7bbab; buvid_fp_plain=undefined; SESSDATA=ad29a25f%2C1671286956%2C5fac6%2A61; bili_jct=b1c511dedcae6fdd22caa1b852ba3665; DedeUserID=1924897562; DedeUserID__ckMd5=0f1dee9b2ac5e857; sid=c85n9ee4; _dfcaptcha=826f5d9149aa6f620a4d81e3a4bf35ff; buvid_fp=22d6e6994df5387a5c00dd9366b7bbab; b_timer=%7B%22ffp%22%3A%7B%22444.7.fp.risk_489B8A0E%22%3A%22181817E04C2%22%2C%22333.42.fp.risk_489B8A0E%22%3A%22181817E1779%22%2C%22444.8.fp.risk_489B8A0E%22%3A%22181817EA1BC%22%7D%7D; PVID=7"
	RoomUrl          = "https://api.live.bilibili.com/xlive/fuxi-interface/JuneRedPacket2022Controller/redPocketPlaying?_ts_rpc_args_=[]"
	PostRedPocketUrl = "https://api.live.bilibili.com/xlive/lottery-interface/v1/popularityRedPocket/RedPocketDraw"
	//https://live-trace.bilibili.com/xlive/rdata-interface/v1/heartbeat/webHeartBeat?hb=NXwyMjUzMzU2OHwxfDA%3D&pf=web
	//https://api.live.bilibili.com/xlive/lottery-interface/v1/lottery/getLotteryInfoWeb?roomid=${ROOM_ID}
	//https://api.bilibili.com/x/relation/tag/user?fid=${ROOM_USER_ID}&jsonp=jsonp&_=${Date.now()}
	//https://api.live.bilibili.com/xlive/web-room/v1/giftPanel/giftConfig?platform=pc&room_id=${ROOM_ID}
)

type RedPocketPlayingResponse struct {
	TsRpcReturn TsRpcReturn `json:"_ts_rpc_return_"`
}

type TsRpcReturn struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    LivingList `json:"data"`
}

type LivingList struct {
	List []Living `json:"list"`
}

type Living struct {
	LotId     string `json:"lotId"`
	Ruid      string `json:"ruid"`
	RoomId    string `json:"roomId"`
	CountDown int64  `json:"countDown"`
}

type HeartBeatResp struct {
	Data struct {
		NextInterval int `json:"next_interval"`
	}
}

func main() {
	request := http.Request{Header: map[string][]string{}}
	request.Header.Set("cookie", cookie)
	//uid, err := request.Cookie("DedeUserID")
	//if err != nil {
	//	panic("cookie error")
	//}
	csrf, err := request.Cookie("bili_jct")
	if err != nil {
		panic("cookie error")
	}
	//num, err := strconv.ParseInt(uid.Value, 10, 32)
	//if err != nil {
	//	panic("uid parse err")
	//}

	var test map[string]int64
	test = make(map[string]int64)
	go func() {
		for {
			header := req.Header{
				"Cookie":          cookie,
				"Accept":          "application/json",
				"Accept-Language": "en-US,en;q=0.5",
				"Connection":      "keep-alive",
			}
			get, err := req.Get("https://live-trace.bilibili.com/xlive/rdata-interface/v1/heartbeat/webHeartBeat", header)
			_, _ = req.Get("https://api.live.bilibili.com/relation/v1/Feed/heartBeat", header)
			if err != nil {
				return
			}
			var resp HeartBeatResp
			err = json.Unmarshal([]byte(get.String()), &resp)
			time.Sleep(time.Duration(resp.Data.NextInterval) * time.Millisecond)
		}

	}()
	for {
		randomSleepTime := 5000 + rand.Intn(1000)
		time.Sleep(time.Duration(randomSleepTime) * time.Millisecond)

		resp, err := bilibili.GetRedPocketRoomInfo()

		if err != nil {
			return
		}
		if len(resp.TsRpcReturn.Data.List) == 0 {
			continue
		}
		now := uint64(time.Now().Unix())
		for _, roomInfo := range resp.TsRpcReturn.Data.List {
			LotId, err := strconv.ParseInt(roomInfo.LotId, 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			Ruid, err := strconv.ParseInt(roomInfo.Ruid, 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			roomId, err := strconv.ParseInt(roomInfo.RoomId, 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			header := req.Header{
				"Cookie":          cookie,
				"Accept":          "application/json",
				"Accept-Language": "en-US,en;q=0.5",
				"Connection":      "keep-alive",
			}

			requestJson := map[string]interface{}{
				"lot_id":     LotId,
				"ruid":       Ruid,
				"room_id":    roomId,
				"spm_id":     "444.8.red_envelope.extract",
				"jump_from":  "",
				"session_id": "",
				"csrf_token": csrf.Value,
				"csrf":       csrf.Value,
				"visit_id":   "",
			}
			key := fmt.Sprintf("%d:%d:%d", LotId, Ruid, roomId)
			value, ok := test[key]
			if ok && value > time.Now().Unix() {
				continue
			}
			bodyJSON := req.BodyJSON(requestJson)
			if now+uint64(roomInfo.CountDown) < uint64(time.Now().Unix()) {
				continue
			}

			resp, err := req.Post(PostRedPocketUrl, header, bodyJSON)

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(resp)
			test[key] = time.Now().Unix() + int64(roomInfo.CountDown)
			randomSleepTime := 5000 + rand.Intn(1000)
			time.Sleep(time.Duration(randomSleepTime) * time.Millisecond)
		}
	}
}
