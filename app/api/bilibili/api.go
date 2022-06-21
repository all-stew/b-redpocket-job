package bilibili

import (
	"bilibili-redpocket-job/pkg/request"
	"encoding/json"
	"net/url"
)

const (
	cookie = "LIVE_BUVID=AUTO5216555653027085; buvid3=83A4A538-BFE4-405C-957E-9EE25D131F2D167641infoc; b_lsid=5521D9BD_1817761BC49; _uuid=10110D2888-CC18-66F1-BB22-71ACBAC3347702862infoc; buvid4=15D50C4D-69BB-6852-F190-0C025F6D533B03101-022061823-kgAZ16FD5EZe35pU/S1NtA%3D%3D; b_timer=%7B%22ffp%22%3A%7B%22444.7.fp.risk_83A4A538%22%3A%221817761BD06%22%2C%22444.8.fp.risk_83A4A538%22%3A%221817761C429%22%2C%22333.42.fp.risk_83A4A538%22%3A%221817761CB82%22%7D%7D; fingerprint=22d6e6994df5387a5c00dd9366b7bbab; buvid_fp_plain=undefined; SESSDATA=ee576477%2C1671117346%2Caa334%2A61; bili_jct=03114be6b43a8f68822d62321f07b276; DedeUserID=1924897562; DedeUserID__ckMd5=0f1dee9b2ac5e857; sid=j6tdg0lk; _dfcaptcha=5b398ed5f4cff8d2557629aa0d54402e; buvid_fp=22d6e6994df5387a5c00dd9366b7bbab; Hm_lvt_8a6e55dbd2870f0f5bc9194cddf32a02=1655565395; PVID=25; Hm_lpvt_8a6e55dbd2870f0f5bc9194cddf32a02=1655565783"
)

type RedPocketPlayingResponse struct {
	TsRpcReturn struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Code    int             `json:"code"`
			Message string          `json:"message"`
			List    []RedPocketInfo `json:"list"`
		} `json:"data"`
	} `json:"_ts_rpc_return_"`
}

type RedPocketInfo struct {
	LotId     string `json:"lotId"`
	Ruid      string `json:"ruid"`
	RoomId    string `json:"roomId"`
	CountDown int    `json:"countDown"`
}

// GetRedPocketRoomInfo 获取当前有那些房间正在发放红包
func GetRedPocketRoomInfo() (*RedPocketPlayingResponse, error) {
	rawUrl := "https://api.live.bilibili.com/xlive/fuxi-interface/JuneRedPacket2022Controller/redPocketPlaying"

	header := map[string]string{
		"Cookie":          cookie,
		"Accept":          "application/json",
		"Accept-Language": "en-US,en;q=0.5",
		"Connection":      "keep-alive",
	}

	var param url.Values
	param = make(url.Values)
	param.Set("_ts_rpc_args_", "[]")
	resp, err := request.Get(rawUrl, header, param)
	if err != nil {
		return nil, err
	}

	var redPocketPlayingResp RedPocketPlayingResponse
	err = json.Unmarshal(resp, &redPocketPlayingResp)
	if err != nil {
		return nil, err
	}
	return &redPocketPlayingResp, nil
}
