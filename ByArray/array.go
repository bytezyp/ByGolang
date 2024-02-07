package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// pangle native ad
	r.HandleFunc("/mock/868/get_pangle_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82","seatbid":[{"bid":[{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82u1110-1683699610810402","impid":"1","price":1.23,"nurl":"https://www.baidu.com","burl":"","lurl":"","adm":"{\"native\":{\"ver\":\"1.2\",\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://p0.ipstatp.com/origin/ad-union-i18n-api/7a22905c6326ce2f74e408b590876660\",\"w\":640,\"h\":640}},{\"id\":2,\"img\":{\"url\":\"https://p0.ipstatp.com/origin/ad-site-i18n-sg/202011185d0d9ab68f774be742af8aab\",\"w\":200,\"h\":200}},{\"id\":3,\"title\":{\"text\":\"Pangle Test Ads\"}},{\"id\":4,\"data\":{\"value\":\"Pangle - Where Apps Thrive\"}},{\"id\":5,\"data\":{\"value\":\"View now\"}}],\"link\":{\"url\":\"https://www.pangleglobal.com/?ttclid=E.C.P.CsIB3IzFbzp9lNqqnJnnPEMVwrB_-ixVa9GaAICWOnyP-F4NEwJYEjlnbWx3bpZiWVZx2LwnHoJ_bS3TwQmp438pUhNzEvC5rLQob3uMJz9Kd_BVqUxYsp4K1myMA_b-UISN14hXHRtvsTpIBiMLKfkrHC5gYPxisBjskwmbof_m5Qhj3bKOwT8FX0LHn-X8B0BQA5Fi-LaPOgWqv2qx6tqO5o_0J71is41oq9EDydAULEmhZTwA3_5KFfdfYXlt4XvnQDsSBHYyLjAaIIGPdARrIalPL11BqR7latkVlKDlGmNZ_fMl7kvC8NaE\",\"clicktrackers\":[\"https://api16-access-sg.pangle.io/api/ad/union/redirect/?active_extra=%2FX4pe384QEpmUKEumKZcRnneyhHIQtBCSEA35gbeAL8%3D&call_back=he%2BxOgO0v%2BhfjXsu8CphBLk2aDvcu5VqmmLYMSTf6SQ%3D&click_source=11&extra=nT%2FQBQrkQerKkxHf4VpaDmIxi%2FV3%2FhrlrBLmSzlQwi8ExJKSkXThZNLU4iSaFXM2FrTW69tIn1MImPD5SXZ4uoIOUHzTRuX3fBbjXPwS7UPL8xMz7r4gl7sdRODVSHZpKcVZRtG00%2FaMUdJ78n%2BbbTpuZrafjQbgg7ODWhEdBBun9fh2FYveB%2F5EqigcTBzZDdwbW5rZ2zmA5VRFSrhJirvTmtvm2MrJnyBg4OR5ZgJz0GPvcSf8tvuG5rcHmXoXpenqEdO17SfVAh%2FnP31oO4lzy9A4O5ZGt7U2vmR2MjQtVObVBN8cTJclWzwl2Z%2B22GRBJT%2BbIH4aGniLT1PBJizDGRzwRDM39r3VzFSmPCjhPYsdNjq%2FY%2BG2df1VUw1rF0bdolDUTkRFoHVo5F5fmQr%2FNdZr9qO0wF0G1%2Fzpw4bwXWMuQqqziSb1Nq5KSy%2Fjy1d6i5P3N2WabxpD65sGAQMzlpM%2BZI%2B2ZDQaSk2cZXgVOBaL%2BlbXBFZgjuucCecb6aw6dN4vSVBlkVYz8%2F13zNh3hf7jCNRUwmy4%2FNRX9bGqBB0MNiy6sJtPf7yqmub3VLfWVhO1V06iqOfcHhxlte3r1PSVGhkvTGHiaE0ZascNkH9VRMOgoY0ScSTKrtf0vBzvcvEBFfUqWH%2BO2Wwl46pbpAxhcvZAeiHANfQu6vdiTEz1YKdJy3bVzjc2wUvwIn2NJF1od2dLa2%2BTHZA7Ev8HmRoTdSFfnSOUlQyfKWv65SzCpeoOUtfVYYaB6T793UqIKnk8FGKvGPvlC0Z6jc7o9ZHQLOSfdj3zvcPkpGcxYYm%2F%2BHwlcTmPnJ2CP4wF6%2BvhHfYW0dcjrGYjuffHXg%3D%3D&is_dsp=1&pack_time=1662021479.39&req_id=6e0bbc2c-6974-4719-a5fd-91b280317b82u1110&rit=980239539&source_type=1&ttdsp_adx_index=229&use_pb=1\"],\"fallback\":\"https://www.pangleglobal.com/?ttclid=E.C.P.CsIB3IzFbzp9lNqqnJnnPEMVwrB_-ixVa9GaAICWOnyP-F4NEwJYEjlnbWx3bpZiWVZx2LwnHoJ_bS3TwQmp438pUhNzEvC5rLQob3uMJz9Kd_BVqUxYsp4K1myMA_b-UISN14hXHRtvsTpIBiMLKfkrHC5gYPxisBjskwmbof_m5Qhj3bKOwT8FX0LHn-X8B0BQA5Fi-LaPOgWqv2qx6tqO5o_0J71is41oq9EDydAULEmhZTwA3_5KFfdfYXlt4XvnQDsSBHYyLjAaIIGPdARrIalPL11BqR7latkVlKDlGmNZ_fMl7kvC8NaE\"}}}","adid":"1683699610810402","adomain":["pangleglobal.com"],"iurl":"https://p0.ipstatp.com/origin/ad-union-i18n-api/7a22905c6326ce2f74e408b590876660","cid":"1683699610810402","crid":"1683699610810402","cat":["IAB3-1"],"attr":[4],"w":640,"h":640,"ext":{}}],"seat":"pangle"}],"bidid":"6e0bbc2c-6974-4719-a5fd-91b280317b82u1110","cur":"USD"}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_pangle_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
					"id":"c7044da583784d458b6e53e4f3f6e5c7",
					"seatbid":[
						{
							"bid":[
								{
									"id":"c7044da583784d458b6e53e4f3f6e5c7u368-1750816935764001",
									"impid":"1",
									"price":0.01862,
									"nurl":"https://api16-access-sg.pangle.io/api/ad/union/openrtb/win/?req_id=c7044da583784d458b6e53e4f3f6e5c7u368\u0026ttdsp_adx_index=229\u0026rit=980290222\u0026extra=4AfhtQRp22DGOBvA5zI0le73cXGNbQsEqhuOTfYNdq2CEwnO9Vp6qtVncz8hML2Zh1T9vcBlRM925wRaQLwxLbJqatUvLeMBjXPAiYb5%2FwMbRvICSA6%2Fk4FAJLtYE5ioW%2B07PiGh9puSFjEXu9ptLDcSn%2FN8BOHX%2FuFjkL9Zrtqjrw8BJpjund99P0WptyIezjRzDYiAQ5%2BpvIgsTRCy5B6tuODAT5ra7fUK6FPlbnJ6xIoBdwjE8sWPuDpfh6NZTKx67AKBXLk5gP5HIBfCdwPBGwSbTtaPBeI7d7oY0xOc9CP0PcMR3MTQLC6ApFoY%2BOQpA%2FoCZQ0C57Ma70w%2FEmRB6%2BJdpPNO2cZrUeqKf7kA3cn5DeO7Sz%2Fhow%2BA06GRI4u0%2BXmcjY3%2FlZcqrzRwcCbTnM1PPnRqawLo7oyty7wMkw4laGnKc%2BkUN0kO1q7GWkDSl%2BvfnQD3hk0q82mBNb8pQrhZl7qoYbESs2b5J%2Fy0Vy41Ya5A3fkNdX562GrHb8dFxQtfjhqVo2THjUIA%2FemoZ0IKoeZViHL8Sw7kJz5BSp3%2BRFJGKoxMUIZuV1jbIOpJnDrg3ye7LxD3oDKoXG%2FjLaAMnSnGijAes0mzBvMOrJNGOz868%2FcDISVuUorBbOAGn1JG%2FJoyL%2B0xH0MIbnfpuvxeJozjeLysUhMblTb79WNkkkAY2X0blxJb2knB8HYx7apoWTcGiIUtZg0joX1ia3PCJX53sW%2FI9aXdBFp17I%2FkWA0wFRIIxTG2T89zL8crlamPyDnpIKxWA7SL%2FrG6Rh8Q7yDHTwItBRRPIvcnL46806j0i74tziR6hC12CEXyRd7Hez%2F4rDSg6CHdljgUhyyM%2BNnF%2F2Wy%2Fg8a7i6kd%2BjdHx9qn8fJmGpvslAA5ZDoxC8u2OxT3MBix3SsS9efY5A71iGATgLEtY3COm%2Btb44fL3%2F8AeB4ES4AYGMfgFynrLgBXQGuonPK2%2FdxojewANAJwJU%2B1geCkKpo9oTL0an3kI9%2F5illb8Cg4Ak9%2FD5p9BZ%2FCIMPDOktRSEXjQTUJAGSrXB3TlNJOIit2wWU7lIYYR1LMtoczfmhmzXOvMLMuVLU8H61ahHcEuTYigv3DxHYqUYstX753ZyHCsi9DF0pQoyCqIcZOAPy93wV6V6hCYafix36M0hcAOoEgrc7tjnJHKeOxJkpoIM18DwGUmuSjMz6tztWU3Xt90l5wawLMZKiH1Zd3xH3e9hO4TSzlzdo7ZMMiWZFKsR99G2vNaZWijAYvciGMJuwS756FQFEjfoqcqSgLyvq87uVoPPtP%2FplxRoOb%2B0E1rzfwvQ%2Fnb7NHyEtWZ%2BG8THfL5J%2F%2Byhgme%2B0YWr2QkIT8hWyDYcJP8GpiI3eBGdBr5oOOSdS30WYHveZB3CedkDfVc0irHVxs4yh8EhD%2BqbTsvTY9cJI785PDv8xxiUaSM24hnHqvShEfvUupankrloUIuxYE%2BZpe9hbb%2Fhyt%2BdNVLk%2F2w1e2J65X178OXpR9%2Fv1mC09ivky3lwaEWkIgYz6MGvY0y8J%2F3RwVUZX6szgkh6NhbavMJbhjzm7Ria7fKMSddaUSkqNBTUmXmbgkJiktpAR0nPjzxidXuX1HIOJqMnX%2BmbGKMXj0Ne6eSQu2lnyNnj%2FNaSe9P7aO2RL0%2FrQCY1nHNFQAB2HGWjtUsX3TTs6wtYDl7QCMG7R9q1KClPyhIezNdoQ%2F9U4hn6WwFBbbFc8WRdxpo4nJ1oSZQd5qojZIeZo3hGRKE2jkRvSC5p3Lu10hghRKVt%2BMSt%2BW034O3gZjGYxcH8seOv5hcufyMWKcPzbqQlqkoBvr0040tLIhwkvIuFIznYEWrkX9UUjwH7YnTaaLvoFAqUoYbNcSWKckVeiM2LH9X2tzNpi5Xu5lZuAwGGTNM5Ht1M1tf9ARp%2BV\u0026win_price=${AUCTION_PRICE}\u0026auction_mwb=${AUCTION_BID_TO_WIN}",
									"burl":"https://api16-event-sg.pangle.io/api/ad/union/show_event/?req_id=c7044da583784d458b6e53e4f3f6e5c7u368\u0026ttdsp_adx_index=229\u0026extra=4AfhtQRp22DGOBvA5zI0le73cXGNbQsEqhuOTfYNdq2CEwnO9Vp6qtVncz8hML2Zh1T9vcBlRM925wRaQLwxLbJqatUvLeMBjXPAiYb5%2FwMbRvICSA6%2Fk4FAJLtYE5ioW%2B07PiGh9puSFjEXu9ptLDcSn%2FN8BOHX%2FuFjkL9Zrtqjrw8BJpjund99P0WptyIezjRzDYiAQ5%2BpvIgsTRCy5B6tuODAT5ra7fUK6FPlbnJ6xIoBdwjE8sWPuDpfh6NZTKx67AKBXLk5gP5HIBfCdwPBGwSbTtaPBeI7d7oY0xOc9CP0PcMR3MTQLC6ApFoY%2BOQpA%2FoCZQ0C57Ma70w%2FEmRB6%2BJdpPNO2cZrUeqKf7kA3cn5DeO7Sz%2Fhow%2BA06GRI4u0%2BXmcjY3%2FlZcqrzRwcCbTnM1PPnRqawLo7oyty7wMkw4laGnKc%2BkUN0kO1q7GWkDSl%2BvfnQD3hk0q82mBNb8pQrhZl7qoYbESs2b5J%2Fy0Vy41Ya5A3fkNdX562GrHb8dFxQtfjhqVo2THjUIA%2FemoZ0IKoeZViHL8Sw7kJz5BSp3%2BRFJGKoxMUIZuV1jbIOpJnDrg3ye7LxD3oDKoXG%2FjLaAMnSnGijAes0mzBvMOrJNGOz868%2FcDISVuUorBbOAGn1JG%2FJoyL%2B0xH0MIbnfpuvxeJozjeLysUhMblTb79WNkkkAY2X0blxJb2knB8HYx7apoWTcGiIUtZg0joX1ia3PCJX53sW%2FI9aXdBFp17I%2FkWA0wFRIIxTG2T89zL8crlamPyDnpIKxWA7SL%2FrG6Rh8Q7yDHTwItBRRPIvcnL46806j0i74tziR6hC12CEXyRd7Hez%2F4rDSg6CHdljgUhyyM%2BNnF%2F2Wy%2Fg8a7i6kd%2BjdHx9qn8fJmGpvslAA5ZDoxC8u2OxT3MBix3SsS9efY5A71iGATgLEtY3COm%2Btb44fL3%2F8AeB4ES4AYGMfgFynrLgBXQGuonPK2%2FdxojewANAJwJU%2B1geCkKpo9oTL0an3kI9%2F5illb8Cg4Ak9%2FD5p9BZ%2FCIMPDOktRSEXjQTUJAGSrXB3TlNJOIit2wWU7lIYYR1LMtoczfmhmzXOvMLMuVLU8H61ahHcEuTYigv3DxHYqUYstX753ZyHCsi9DF0pQoyCqIcZOAPy93wV6V6hCYafix36M0hcAOoEgrc7tjnJHKeOxJkpoIM18DwGUmuSjMz6tztWU3Xt90l5wawLMZKiH1Zd3xH3e9hO4TSzlzdo7ZMMiWZFKsR99G2vNaZWijAYvciGMJuwS756FQFEjfoqcqSgLyvq87uVoPPtP%2FplxRoOb%2B0E1rzfwvQ%2Fnb7NHyEtWZ%2BG8THfL5J%2F%2Byhgme%2B0YWr2QkIT8hWyDYcJP8GpiI3eBGdBr5oOOSdS30WYHveZB3CedkDfVc0irHVxs4yh8EhD%2BqbTsvTY9cJI785PDv8xxiUaSM24hnHqvShEfvUupankrloUIuxYE%2BZpe9hbb%2Fhyt%2BdNVLk%2F2w1e2J65X178OXpR9%2Fv1mC09ivky3lwaEWkIgYz6MGvY0y8J%2F3RwVUZX6szgkh6NhbavMJbhjzm7Ria7fKMSddaUSkqNBTUmXmbgkJiktpAR0nPjzxidXuX1HIOJqMnX%2BmbGKMXj0Ne6eSQu2lnyNnj%2FNaSe9P7aO2RL0%2FrQCY1nHNFQAB2HGWjtUsX3TTs6wtYDl7QCMG7R9q1KClPyhIezNdoQ%2F9U4hn6WwFBbbFc8WRdxpo4nJ1oSZQd5qojZIeZo3hGRKE2jkRvSC5p3Lu10hghRKVt%2BMSt%2BW034O3gZjGYxcH8seOv5hcufyMWKcPzbqQlqkoBvr0040tLIhwkvIuFIznYEWrkX9UUjwH7YnTaaLvoFAqUoYbNcSWKckVeiM2LH9X2tzNpi5Xu5lZuAwGGTNM5Ht1M1tf9ARp%2BV\u0026source_type=1\u0026pack_time=1670219535.53\u0026openrtb_adx_id=229\u0026pc=BacTkYhDa%2F5FRvOICyJvj2afV%2FjcG%2FID2AgSFwMkc6M%3D\u0026ttdsp_price=${AUCTION_PRICE}",
									"lurl":"https://api16-access-sg.pangle.io/api/ad/union/openrtb/loss/?req_id=c7044da583784d458b6e53e4f3f6e5c7u368\u0026ttdsp_adx_index=229\u0026rit=980290222\u0026extra=4AfhtQRp22DGOBvA5zI0le73cXGNbQsEqhuOTfYNdq2CEwnO9Vp6qtVncz8hML2Zh1T9vcBlRM925wRaQLwxLbJqatUvLeMBjXPAiYb5%2FwMbRvICSA6%2Fk4FAJLtYE5ioW%2B07PiGh9puSFjEXu9ptLDcSn%2FN8BOHX%2FuFjkL9Zrtqjrw8BJpjund99P0WptyIezjRzDYiAQ5%2BpvIgsTRCy5B6tuODAT5ra7fUK6FPlbnJ6xIoBdwjE8sWPuDpfh6NZTKx67AKBXLk5gP5HIBfCdwPBGwSbTtaPBeI7d7oY0xOc9CP0PcMR3MTQLC6ApFoY%2BOQpA%2FoCZQ0C57Ma70w%2FEmRB6%2BJdpPNO2cZrUeqKf7kA3cn5DeO7Sz%2Fhow%2BA06GRI4u0%2BXmcjY3%2FlZcqrzRwcCbTnM1PPnRqawLo7oyty7wMkw4laGnKc%2BkUN0kO1q7GWkDSl%2BvfnQD3hk0q82mBNb8pQrhZl7qoYbESs2b5J%2Fy0Vy41Ya5A3fkNdX562GrHb8dFxQtfjhqVo2THjUIA%2FemoZ0IKoeZViHL8Sw7kJz5BSp3%2BRFJGKoxMUIZuV1jbIOpJnDrg3ye7LxD3oDKoXG%2FjLaAMnSnGijAes0mzBvMOrJNGOz868%2FcDISVuUorBbOAGn1JG%2FJoyL%2B0xH0MIbnfpuvxeJozjeLysUhMblTb79WNkkkAY2X0blxJb2knB8HYx7apoWTcGiIUtZg0joX1ia3PCJX53sW%2FI9aXdBFp17I%2FkWA0wFRIIxTG2T89zL8crlamPyDnpIKxWA7SL%2FrG6Rh8Q7yDHTwItBRRPIvcnL46806j0i74tziR6hC12CEXyRd7Hez%2F4rDSg6CHdljgUhyyM%2BNnF%2F2Wy%2Fg8a7i6kd%2BjdHx9qn8fJmGpvslAA5ZDoxC8u2OxT3MBix3SsS9efY5A71iGATgLEtY3COm%2Btb44fL3%2F8AeB4ES4AYGMfgFynrLgBXQGuonPK2%2FdxojewANAJwJU%2B1geCkKpo9oTL0an3kI9%2F5illb8Cg4Ak9%2FD5p9BZ%2FCIMPDOktRSEXjQTUJAGSrXB3TlNJOIit2wWU7lIYYR1LMtoczfmhmzXOvMLMuVLU8H61ahHcEuTYigv3DxHYqUYstX753ZyHCsi9DF0pQoyCqIcZOAPy93wV6V6hCYafix36M0hcAOoEgrc7tjnJHKeOxJkpoIM18DwGUmuSjMz6tztWU3Xt90l5wawLMZKiH1Zd3xH3e9hO4TSzlzdo7ZMMiWZFKsR99G2vNaZWijAYvciGMJuwS756FQFEjfoqcqSgLyvq87uVoPPtP%2FplxRoOb%2B0E1rzfwvQ%2Fnb7NHyEtWZ%2BG8THfL5J%2F%2Byhgme%2B0YWr2QkIT8hWyDYcJP8GpiI3eBGdBr5oOOSdS30WYHveZB3CedkDfVc0irHVxs4yh8EhD%2BqbTsvTY9cJI785PDv8xxiUaSM24hnHqvShEfvUupankrloUIuxYE%2BZpe9hbb%2Fhyt%2BdNVLk%2F2w1e2J65X178OXpR9%2Fv1mC09ivky3lwaEWkIgYz6MGvY0y8J%2F3RwVUZX6szgkh6NhbavMJbhjzm7Ria7fKMSddaUSkqNBTUmXmbgkJiktpAR0nPjzxidXuX1HIOJqMnX%2BmbGKMXj0Ne6eSQu2lnyNnj%2FNaSe9P7aO2RL0%2FrQCY1nHNFQAB2HGWjtUsX3TTs6wtYDl7QCMG7R9q1KClPyhIezNdoQ%2F9U4hn6WwFBbbFc8WRdxpo4nJ1oSZQd5qojZIeZo3hGRKE2jkRvSC5p3Lu10hghRKVt%2BMSt%2BW034O3gZjGYxcH8seOv5hcufyMWKcPzbqQlqkoBvr0040tLIhwkvIuFIznYEWrkX9UUjwH7YnTaaLvoFAqUoYbNcSWKckVeiM2LH9X2tzNpi5Xu5lZuAwGGTNM5Ht1M1tf9ARp%2BV\u0026reason=${AUCTION_LOSS}\u0026auction_mwb=${AUCTION_PRICE}",
									"adm":"{\"native\":{\"ver\":\"1.2\",\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://p16-ttam-va.ibyteimg.com/origin/ad-site-i18n-sg/202211075d0db6a615185db148548148\",\"w\":250,\"h\":250}},{\"id\":2,\"img\":{\"url\":\"https://p16-ttam-va.ibyteimg.com/origin/app-market-data-sg/68fa0a7402620d72ad27f9e842858d06\",\"w\":200,\"h\":200}},{\"id\":3,\"title\":{\"text\":\"Storm Device Cooler\"}},{\"id\":4,\"data\":{\"value\":\"فون کی جگہ صاف کریں۔\"}},{\"id\":5,\"data\":{\"value\":\"Download\"}}],\"link\":{\"url\":\"https://play.google.com/store/apps/details?id=com.storm.device.cool.ects.cooler\",\"clicktrackers\":[\"https://api16-event-sg.pangle.io/api/ad/union/redirect/?active_extra=ta%2F5ajlJMOlsPXR0ZWErWjOiDs4RSeYXesm3s9A%2B3woh4voag5mgTe%2FxujGUMMeK2zT9liTxR7BrVwPg3WINTg%3D%3D\u0026call_back=jLuLQ0Ck9ylk%2FQdnmY%2BBeBWxFQZuFGTr9SLbNZYhA%2FTRT8h%2B3AqnhfixHvEJ9FgGPYxz08POeyv5Vz0BdtAgBeXQzIkfwSnCT9UuMOWn1GJmn1f43BvyA9gIEhcDJHOj\u0026click_source=11\u0026extra=X8eY4x5BVsGa591fOmbtOPao9%2BnWegyMKAzROBlUTBmrjzYd1pk5ko5yru2nj01lVtq3nsY%2FGv5GCNXZdgKeQIIOUHzTRuX3fBbjXPwS7UPWa223LIsBACLcVEcpBdsGVCio%2BxDXJvCpwoDL%2FxA46iuH8hw84rN7ceZJsmRYlh60nv6YCengBHUIBnULQi%2BFVHq0YWCJhP74D%2FRXH7ZXWxaLbawbp5aQNIxOMuZwBE0%2FNs8igrFUp5rrezxBPtgueqkzUJJhayRkQYYm%2FTNUOkXD%2F%2FUe1w5FI3HAVmQRnhQrs%2BkrPe4hOHLN4CU2abniu8ZgPTOX%2BunkpCBnTYOOib%2FpKDXYb4m6y4XtPxIPLXv1XkVnXMBnjSyL%2BRswnPCRhLAFUNAtpt3frvNtqp1RUz3Q1cfNPveP1e4HO0pi73FJhxHwtuJusojv5VHEBA0%2FPNTjRw5LRtye6Ni6zs1JcpuHo9yGjSRgECR8t%2FeH5qMUn4VdSsIcptFVmsF16QRQkp9fjeltccMANy%2Bo9mKsq0tyPQUfHBL8x5uSb6%2B6Wa0HImNOkQqY%2FALx7T1s0QwIzJe5lcvX3w8CY%2FsyWnWAYOUy22FHd0qhYiz71RkXzZFkLoSxRyWPncEsoEil%2BYoejGEr28U9WY09gcRfM6YmM43Q%2B5O%2BAtP5h5MnSnGHULbTFRXymEpPGyACxSgJxxCQQagdEIecjvX9yyQh0XoTbzUMztkAJ%2FRaRTOd0iorU5ly6MuFkxc950j9w8%2BwkTriikoz8FizkzlMbZS3I7FWSckXwP%2F%2F3zaDObZESPMl4v7AmU4AQa2smrVX5OW8RhzENtKVtqiHDJgTi%2BrFM9oyBUa7vNtyhXXQK2FqrX90sL4MzpJpSG0CLCyDjeJ33RWYjqkT9%2FxNrRiYomIkGqb9YUezETc1vmJQnwJKIiRbSB0%3D\u0026is_dsp=1\u0026pack_time=1670219535.53\u0026req_id=c7044da583784d458b6e53e4f3f6e5c7u368\u0026rit=980290222\u0026source_type=1\u0026ttdsp_adx_index=229\u0026use_pb=1\"],\"fallback\":\"https://play.google.com/store/apps/details?id=com.storm.device.cool.ects.cooler\"}}}",
									"adid":"1750816935764001",
									"adomain":[
										"supersallytool.com"
									],
									"bundle":"com.storm.device.cool.ects.cooler",
									"iurl":"https://p16-ttam-va.ibyteimg.com/origin/ad-site-i18n-sg/202211075d0db6a615185db148548148",
									"cid":"1750816935764001",
									"crid":"1750816935764001",
									"cat":[
										"IAB19"
									],
									"attr":[
										4
									],
									"w":250,
									"h":250,
									"ext":{
				
									}
								}
							],
							"seat":"pangle"
						}
					],
					"bidid":"c7044da583784d458b6e53e4f3f6e5c7u368",
					"cur":"USD"
				}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	// webeye native ad
	r.HandleFunc("/mock/868/get_webeye_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"id":"9960824f-7b6a-4bbb-92ef-fe18cf2dd98b","bidid":"9960824f-7b6a-4bbb-92ef-fe18cf2dd98b","cur":"USD","seatbid":[{"seat":"2","group":0,"bid":[{"id":"1","impid":"1","price":0.07963397334907299,"w":640,"h":640,"adm":"{\"ver\":\"\",\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://p0.ipstatp.com/origin/ad-union-i18n-api/28fb62df1b0c9dd96dd1a2206cf40344\",\"w\":640,\"h\":640}},{\"id\":2,\"img\":{\"url\":\"https://p0.ipstatp.com/origin/app-market-data-sg/7d26e659c52bc7a9292c96eb556d3153\",\"w\":200,\"h\":200}},{\"id\":3,\"title\":{\"text\":\"Tantan\"}},{\"id\":5,\"data\":{\"value\":\"Dapatkan\"}},{\"id\":4,\"data\":{\"value\":\"Dia mengirimi permintaan pertemanan\"}}],\"link\":{\"url\":\"https://play.google.com/store/apps/details?id=com.p1.mobile.putong\",\"clicktrackers\": [\"https://api16-access-sg.pangle.io/api/ad/union/redirect/?active_extra=y4e2n95wYka1cKfWPPXsB2BUX3ni6QZCsJUc1jCkvIUhuhEuEVDtwNapidu8%2FoCGb6rIkxCbI8qcXz52wkOWzg%3D%3D&call_back=jLuLQ0Ck9ylk%2FQdnmY%2BBeBWxFQZuFGTr9SLbNZYhA%2FTRT8h%2B3AqnhfixHvEJ9FgGAw9E7CLoVY8rrsyKOEyHQcZxWqvgZ0q29v5NN%2FRqlT9mn1f43BvyA9gIEhcDJHOj&click_source=11&extra=drwpzxX%2FX9eWz2PsxZhQFDYuMCxZ9P9IWy%2BcH9lZmqVVOrfRkRBxnAmLus0eyYYZIkSKJFoZoAILT%2FvtL5DGr6CpkDvt38XCnfw%2B8lqxbyw%2BwO%2FRl6JDAOKL5z4WGF59oILaUamO92cOJDwVt6oJPICFU4KQ7OhIaWNiBA%2BSMJVtNptOkUuIMCRXju42U81BhnJ%2BvM%2Fw8EOYvuCVqOzw4yomUlTAQKfSMNm95ARWyFU7er5hor0aIhmMcXJD84pEonEPumm2D8Q8ZEendR5nUt5rKurvNDdTuEQxQKB%2BSbRJo%2FiF9ay57PtRM%2Fv4ih2TQj5jZy%2BaGx4IoXAd07mmAmIkKWfZuJ%2BDc5XFDyVYy%2BZmLeLaagFDtFmrK8nFmEK%2B6g%2FFTRHtGZFZMmbaNPawo4kfjEIT6o%2Bi%2FGcDs62w0%2BUV4Wz%2BZCmZlt6m0lU6F3VB%2BMRqnnFOdju5hUHDQaqDN1L62ux5%2FydA0ega3xv0Zmn3DFuDOkYP1X%2BRR3xznUr0PrLL%2BFOgZyeYOcJ0LpS5XVuZgs5LKb%2BDuJyKgkNPMmMpiYIbJVaBzMIe6VFu5XQX%2Byt2fSHo7oaUtdl427%2BgO1azRHW6FrNTcqw3GsZwNtOLq9RJZNa89%2BYQ97smFHmI78OdCVwv%2FOCxulTkpExyEm9HNBUGOHnjv5R6VemNG5A7JqccWBmmfUp7h4jgGcy6H9I%2F%2Bbik7HSKhUogjU02ID2wMgW%2BTtuQriNjubhzvrIey3hQ17MZCC31f8WMhbeLP0ipdC1aTfoIz0fVQguYQJA9QUHVYAp%2FF7lBCbABSZYbaDjk36qnpS%2B7PzHqWhFxPJT6a%2F%2F4D%2FTUSbTbmU98xru2zJR7oF7DcDxOwpNSvKH3sc6eqKS1nTtxhnBpVTu3CGbo4MAIT0mNttslge0fy7r3%2Bag3EPBqzVZWd8D%2Ba1M%3D&is_dsp=1&pack_time=1662019247.51&req_id=cc86dbug1t0ueuf6elrgu6081&rit=926901987&source_type=1&ttdsp_adx_index=101&use_pb=1\",\"https://rtbtrack-asia.rtblab.net/clk?p=271&d=80&t=wcIm3rKLUAeTTiiptPNypq8S4aGKjSiX2c00pNrZx1rg3x8J6JFd_UlNbkawVsluAPKku6gyuqDpyuQ4lKM4Adnq2telLHixtIm-cw2CXUUxqZxgp0Vnu6y4QAsUQV6EBIkS1RNjbz5N36dqg79prISxTFa7cJFrf5UG0FyUG2UhCJXhvt-XKevL65_-lAcgGQ4i1igwGV1b889Mzweo8OEge_dDupQATeGdrRTBWHjB6oes6vKZDZ3g8OuzSB2rGZTyghPj7ejOaSTHuQ2yrR2dfUU37sj-8XJeIPSOwRilyNxkf9XXAcClyuKW7Of3wzHbw-oW7T1T2HuPb1ncXlblWTbk1neDepSxiTfqNzhwXYjSqgEa0Lh6vPeAX9zB44frdEp8pJQtuLyzbAoxjtuzq_1WG7twVpRWhiQs91PX1ft6mVQkRaO6mFY_X24ww35csc8FzUMtEav1NghIr8W46ZA2Fot_8Nulec4kbELmxR_do7innoKDOc2R058QJeFekvj2ci8O660VBhQeKIFT9xlGrpoTtmgeMj8WGLY\"]},\"imptrackers\": [\"https://impression.appsflyer.com/com.p1.mobile.putong?pid=bytedanceglobal_int&c=Tantan_tiktok_ID_android_Pangle_regis2.5_0909&af_adset=ID_adr_pangle_regis2.5_idv257&af_ad=257-IDv_C_1x1-2207%E5%88%B0%E6%9C%9F.mp4_004&af_siteid=900000000&af_ad_type=50&af_channel=Pangle&af_c_id=1710424731716657&af_adset_id=1710424903089202&af_ad_id=1710424861130786&af_viewthrough_lookback=1d&clickid=E.C.P.CskBh3VlpnS2LG9Gvkys9Dlprig40xlL0_JRlZvOezWiuj3CfGZKBe6opaLqdowoLxMbgtwu23Cn5x6wGuXYAsDiQGZz0ypNjw4Su7AdRB4SrpNMMRSdmot1DvrKgH_BjNEYMTt6UC_b3Mo3fgKSogAHrPTaJy91-xXK3IP9cTb6AtPF-qLBlH0v6eUrXuFQXN8xkrxTs9LwjJqnIPATLhbSelnXMOA9l7LWJ-TspPARU0CcdJRSe26mOSeKBxg4Hv1V9hUvWshzsDk1EgR2Mi4wGiA2eZpXTkZlHcvX_kXAcjAWhbREUqDEgny5F_1CLu1WHw&advertising_id=5ef57c14-d642-447f-ace9-543262c458bc&idfa=&os=0&af_ip=114.124.182.13&af_ua=Mozilla%2F5.0+%28Linux%3B+Android+9%3B+SM-A207F+Build%2FPPR1.180610.011%3B+wv%29+AppleWebKit%2F537.36+KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F74.0.3729.136+Mobile+Safari%2F537.36&af_lang=id&redirect=false\",\"https://rtbtrack-asia.rtblab.net/imp?p=271&d=80&t=wcIm3rKLUAeTTiiptPNypq8S4aGKjSiX2c00pNrZx1rg3x8J6JFd_UlNbkawVsluAPKku6gyuqDpyuQ4lKM4Adnq2telLHixtIm-cw2CXUUxqZxgp0Vnu6y4QAsUQV6EBIkS1RNjbz5N36dqg79prISxTFa7cJFrf5UG0FyUG2UhCJXhvt-XKevL65_-lAcgGQ4i1igwGV1b889Mzweo8OEge_dDupQATeGdrRTBWHjB6oes6vKZDZ3g8OuzSB2rGZTyghPj7ejOaSTHuQ2yrR2dfUU37sj-8XJeIPSOwRilyNxkf9XXAcClyuKW7Of3wzHbw-oW7T1T2HuPb1ncXlblWTbk1neDepSxiTfqNzhwXYjSqgEa0Lh6vPeAX9zB44frdEp8pJQtuLyzbAoxjtuzq_1WG7twVpRWhiQs91PX1ft6mVQkRaO6mFY_X24ww35csc8FzUMtEav1NghIr8W46ZA2Fot_8Nulec4kbELmxR_do7innoKDOc2R058QJeFekvj2ci8O660VBhQeKIFT9xlGrpoTtmgeMj8WGLY&ns=l84an9h49lG1MtsHPlulbrwptsvHvoSsq2Va6k9EyuAnVAiDCuN-f-BF57BfUqNFJxfTSflyV5pSwe5zBs_dpa1RVo9BryIhw1K0GuXlVT2Fe_5d9x3pGRekbftUsPZG2rauroL-fhGPaWO4IMTzZWKel6pjbjMlrcCysHYkeApJpkeAo7KrxcFwTUz-o1wqzXaiBy4jNIB4PjUbwnib9XKQGb8Xx93HLcHk4pznt8smg32a_kMFNCYlyN9HVAmsfUuyxEpW9ci8Dz37zPcO4ghNsKD4NJrxkT2W4NXrlQyEAC4p6xkdjtVGV8kLXUdxRyI-XYpsStE9NTG3EwrHy_4sfrUCpOf33B7vpTIVawVfKmPpvOXdNBgETFieu5bCUkI4THamkyauE-LJ00GDFqOjoy1H6DIF7jko8bGLvFB96nBViSnpN7unNEoD54aRUyPLGm48iMT4ECdEJoraE9ZxUS7bJSIm2zVvSHQCms4jhGw_me_ZtTwHS0uakGywRxeVKk_aRC7qz7sxEz2zAktt1lgBGlsODMna8mXdzz-dqKG5-YVMc69iqEbhPoS_TMdf4VSz02xLgMPrQT4mXhWBhEPr3f5SdevRfVgZjrIce8_2AktZ3KuxxwmQ1VYjtdj7Acspu64g8SgJfkOsEY_KGalW-SMDj2rcJbDKvxitMBs7uQd95Gjs9efAdjrXbcilBCpwtcCpw4BEXrmkeDHAztkmeuJ43-wKlRoz28twiUzgTJDWQ7ShxqHS6v-GheNf8DtH2aWT1PvtTW7W19BHlEWh-omLIqnzCVlPY5lsL4WERSmT_u3gbexGD3fT7YkR4L4jaNDXRuyaUfo94O7RmUvRohrQ63t7y1u7wRqghGWRZTRIjKaDpaIF6KDjThaiLtdLp4R5Goc6t4KEL7FyofcB8FhGMDee_n-yQbL1rLMcuBWzR8SG8KQ6BZFQMKFpDk4sj9hS7qCR0DJIB51_-nFIQgHD-31ukkrYpCNyMjpXDkJAoL5DKeYr32adhI3pFxrpWIkz-kBGvxNH_9l_oxNvPHYUUT9xGqJAuRG_VWSfyuvpxUQiB4LAKowMA8Vl2WHa8k_D8FsiMbKzrzmiLR7bT0Trr1U_EHcWzzNOHlsZFeQS8E0QSbqwLVKyGCWmh58UoIV7MHgnKiKtp0dD1KJO9yQVybVw1j1LPTT-R2hOz0CLxmNvtksngzWDmzKIUQ2nRI0ol29QAv--PUDSbkuxdPp4SLKm4wWbGY_Ruxa8p3OP1UbzZHJiCQMRYK77uc-LAw8_eq4XatnEEa07YClWKXKXqYKfAsz0kM8OLqouQqpk5ttvkjKiSvNrSTnCv93Jv34FxDu_IXbHTqpX_VG85TxHO0MvMkmzQP71KkhVnaaLBUtUHvyZp_x8KyA_xH9yd8tCRtbr1N88fJ7_oSBlUXA42fSCS2uDMBvErofKy5qTum-zXwmfqJ717T7f3rdBKT9J4alz09czvnzrz7VchsI6ynvPWTIL6V3oKrtOHmor6ItunZunwAIET2W2T9CzT5YWaDqoBJbby4meZ_4008B62gJdD0trD7Na6fR1cYdFhwmXTei13uv5RxK2_Dq0WPqJpdqboAB6kudaVd4293LQ_Hj77Lu-EGfnnaLgSa3eVeH_tRIhxpHMw1q4Pkat2lvAJwN4XKaevZT4E-A2B_MFl8TlXLNRu7gjSYK9hkJIvGkKlbA3EHNmyhK6uuV8xyhD4ph-TZWX73KUYtp8wZ9Tny-ZsgYGMe-KUaUI7-Mo9s38-uRGqpkL8JV93i2goxIUr0S7TsT9HTyPBmd9mIjGAMOqvD52E4nh-rZjxEG4EuFecbJkFsPxfkJZMAW4jHRme4iV6lP1tQIWiwMveySTZHCvuZ7UWx8tRZNUfYZEtQqwMLoU1z_Scs56EiecvcAg_Wk8G0N3Bi8c7Jf4FFmVuOEbg6VcfsDQ394Y4UFGPCkERb7lLZ9gWuDCCloiNeqD--onmzM9dYVwv50h9_CMio17t988OGuvqpAyNdYhJm5AX1kJ39uLyqrsf6efRhEHhKhHahIoVN2d3OFECDRERvvwgjEk1JLL7ccdwbo9CeT3q1OApmX0uBm0Hi-ljVFt9gXdl8pEfh_lztz3yQysYERr0fQYTR7kC5Y-Xcbeszhjg7VFTDxcC9hj93tE6btyDRQCi1CsHjo8WJqV-fg2hK4n9V0gsIzyWNDEhh0LXkFCokrntU1EutC8H8GtuGcQxwuaoPqJ_hnCg2psKjoav7wCgiGbtVrRpMN98XYXxaWY0ehGS1F80WjsEY3he9TKh39MeR_iD-dH4VaX7-j5pjkHHRUd2lpy39OYggO3Wi8Hk6XTlZrDkJDLz7GNZulribRapBT9AwL3zJBr0PQrJfQr59zlBweprmuCb0L-W606JE_m6AgAkG5Kt9oPT8d5MCyDulP79WpxJc2CnWhmO5zeaLdVo3LJgr9NntsLERrx9U1TQi3rg9M68w89ZJeRrEdbDS5hfA7fdzYyBrMN_5jOUgUz8HMKxCYScaNPruR1JXl6uQZNVkDynQfo_sUNNwrltShQM3A54tv19-n3sPoMhvrSetsxFJCd-qbjeHvD2bkLXnWfttnsZKQIPTS1e1enhTlQJuRpa9-jzx9wR5YOmxQD__PXl-zR2qvQWWrlUy7k7-kgYbVOzBrH5Owu8FGA1vfcaeNggRlWnotd2SYg4KPMcEqpFXOk5wVa1sEkjoTwXLMRSiGjEcDw5oybd7TOkZXdsZE80cyGdBn8OiWIDBM29kshT2pesjc5wRd9BoxrehuLq9QbGn7PnW2fsHP5k90t-BRu_jDspjA0UyNp_klqrV54-omMqHkC7fS-jDLiVCl-lvmjzVeEy5wYqVVXpksFVamwSlhjsrq79INu5Q3M9HYB5SAbvEijeWROGBU2ULg-7JQwlVstLSULPmAmoFRptxLDKDWyi9TdqsBtdnbLR9CwCllfWdWvZKAfvGqR6N228oqeBsI62KHuR8dHOYfN6a4QWEw7ofVCDud0klGpZgPdu1ZPKuyIzkcswJFVp3M_72zhE4UM_ytZzWJQTotDoil5OQyeLdOOT1I1TwppkFmxXd6qhN0FqHeF6oM54kBsLAeicPDUZ4IM_3uGkl5sq8QbYr4_ObPKA1LMf_qzyV73GTMwSuxXn7MgqS9uPiPnm4xZCOpq_k0t9sG8Zyrl-IEkTkI2xTRdWk2wT_5XGacUAgLJCxLvl0xbcknWdoSlxDmok2sB5G85Ie6JDns3ElilJ7mCMsJEzvmOaR-2EkKhGfsdxDHnh6olX1SttdXEDPWbecRRbmumm268o2WAcFPFxyJ7b36WmnAAR50H6jOv5Yz0y_XnsjcnpE0-aURmbKFHzhYwmIG9MR7zWNp6BpPiFGDnYUVKpUilSt7wHkt-FgMBqWBCY7VeOHNmPAy9paANljF_QxUoGm2ZZXZUiShz5IBBP6E_n6E8tFy9T6HVC86Uxz63jWoX5HAkUwbzWjd62PRq-6eu2rXMChqSioEWIEN8UW_I5pbK83-YLus3RlWSNeIfP2JWwz6FVgGiszHDDc9o_7LNBxbBaKNRpJB8lA8xXbLlNJZWhn3A8T81ed5cgPNz4pBdwlhsVzn5ic1ijyU9evePaDy8Ycb7j2t8mA15VZ2P9ORCnoAR13MnpoUqGDYJ6JU_uEZWjdUFOyNtoKjafw8mIvGjKhw8lTwznYoMasCGj2CKZM3zijdea1YpKSx1OrU4tXGT9XusmgeHYqcpi7AdHF2rPSGbsLN2Xym9zgL2L2IB62VThEQtNxu6AHU2ZB8WDynXSEEHorGcJ5sI6FsDtLxHyVJidh0MXSolUaXHF0LyDzMLdSepe3A4pr_q99QlH8XS_8OaI9W98BUFP5OalQoFhe0gjSiQCLD-_ATyhNKAUcsJp1xGTz648GBAwzY3GCtrZ1zmEcM9CGVpzGZOj5f7xZWhCeB1Ie_vUfPjEDpvvJ-XfL98cC8nH-CYt-oLEnCEtuPqJOEIvwpZXWuoUyhnh-e8ECiYSj1TdccNr5xby6WrfDRN3iOuW_n94u91fhcRV4-2sMsIqeaIYdkXRclDIRgaQNJQR2F41AaeEsq5vahtxjdCfDO9BE8enexu88umyMqSDQVt7bhoc3b1JWqgoJrFwwxffunHKkkGhtN3JCQxghflVwRg8DWPj58PtkETG_tpyHsuHI2IoOyEaVTVCtPfE7svz3gJBMk6IHXABYFZJaazT_wUbcWWXdPZ6w4biVx0&bs=l84an9h49lG1MtsHPlulbrwptsvHvoSsq2Va6k9EyuDWc3vuy2eNFU8TNgSitfShHPFRc09j-nLeyptp4tptNQ_x4RqnvslwK78ToCZ-NGo0xbJjU5ovzYL5Zkx7Wqau3lSlh27MqCtob-ENwMCmaM_tCb-1h0ag-39-5Fut543rJMBN8T3yjh6l2sZzWNoqnerlbX8pnXYlcxghc7sPUi-CNBKwPQi73AJiSAB7fcDSrsE3wtQ3eu6JCD_q9nTvWTVVMhmtZvyGZNkMD3GiMXjfxa7xK3CgGSrvrXuXJcEqyvecHispyJVpU-oVye5-nCuwDYVR16TaICe3qSI7h5O14zubWSxVh08WOnaFxdGQtkDtrk-qi5SRv2VIF6SaI7L7bKww8amRekH1l7EZ4ebpMb60jeOfj8DM-LdlrpfkGCsj8ZvI9ieyp2zzlJ9Y5WbpjTaofDxjstAcsJt28xOzb8WLCCebq78IH7fGj--P5J5oT7PxK26SRNMc6CS4HQY-kLxamzaiFqVgi-waqVLlHRwntbpBkUQjM9m6yzK_6gkEyiAJxUY92GER-LTxq6PF9SF_NZwGNCp7ULkU-EsEQiofWZdLZk3NUjPbjO0xm8mBiuRz4J1R1Ds0MKiCnDZfk5a6US3UOqNXjwYX8Gpw0kRmqLaTJ96s6_GZMqZS2vBhRwMRykm4mAic0L4FfGhacxrhDUNBzODdoH0jxL0UKflatgVHXP2fcfA1rC7f1VHQcXjDiyohwSkyhCbPXuGKefTPWRauYO2cJV5Q0ASgeHmI4HPhJgBqFsSvS1jHUWyo1c1TWmk5xzDF3WpKY1v1vtg05Pj5XiiX5uNZvsZBjpyYNME7s19rbphZXIIoyQjZOOdc2EDaL_ilS745odV8l7xZfSCYGcVuIU7uXyMgC7Oiu5r52MkiB4X8scF7UxqRRkU0YOdBoeiNHTHoGS38Juk14BlxLAWRxaz2_UoUTYKQhSeJWKtswyB22c4oa_FXrPV0snzW4ztUPJ62R46BZtp6dwpLoiWwDH951goVkaWXrJhFpBh6yDB3rRwQUCb2ojL8t4haCCLj4OTfsxt3D69EnMEjPsRMBKCWp3l0yGjb_2W7eVjD8pSM3ZN57h5XEsomnLfYm2gkDW2CqFSp3vourg6UL9004G78hYk0zruWL6BELNJTXIV1rsXzoCLRdyJaPBOlzcm93x6iAVlJHD7O18XiwmaFOLStYdke-SOyRUCQfR9ogqXTNaiN75S-sfOM-5pKxBNmnCbKv3mw40qo5ngKSe8fM2qpAjYN0DRTtt03VsfWSLbHTw9mRmPTSSyk1di1he2Uwikm_Qx41XI_JZLHf1qvjJbO9ZivCYhrYf3_k1zg0me8kLWSWOhQb6E7cOFdyjNdGPj20taLhKQVosR1ZVYBe2ORDhERbXe0vdnS_hIu1g1bIPXq3zR12ptkPgErCGF3Mg4mp386vj6PNTPoama-4GHn8UtMFgo7SpxYa2KbUl3CkojWYNwsJ8cIOGLWmAPLlqZgaeH29Fr0LnPudsDxzGujJ7re01EO7D3fAAe0nwQuVaUQee4m8pzairz6YjnIhs5xE-pw5xlEn3urVlSNonwz6ZMhNhbaw8WgUHJk3MiTTdlsQZxr9DW7guJN-y4kHdpmittCUIUOZSzmWWoMG2im8YXg_AgL2V7Dzarl1yWne-yZ95DJIdX4buX715_w0v8bWO-_PlGwoT0sahkPBsAM0QBS-0Mxwui-p8z_JYA_L3Q_PqLZRpGTwrr0i8j_ZhcAKH9wZId0p-dD9PjVQK80GZw9t9MpSsbV2G2zvNX_L5qxPl-Cz99DoIXSKuHd5oAI-NPBpBu7EARkNctfPVNRGvpLLwoNxoit4WwV2D9igAKeIrpq_LuQAlITWUHXCvVNJ7Cryi-uHgsTi7lZE6_5WCn6g8dW2rLe9eVze7eUuIqh61eptWIKIHevMZfA6UqZkQXYZempFIgiyf1jzofdmBSUMg69_iXIDSj-Sr9HHXWXc2hliVZkNZmB9iA1Oq_MBAxvUN4yqY_8-LmGXrWw7_nBt1XWCjxLwogdH7Haa0XE-x1gEl1kgGukvPmrjZKrKq73HnwJPLhaP9bUc9Id8HNLzqSotuvALL1D1EKtPFkhdC3RDwK44pQR-UIsv2Cic5qBlmXRCZGuPnAz6tKahwzu63jN9GmTJWa8fL8Vww8IKho26zSMYcxgOdrNvvJ6hsStofRR5EkN7lfvOa82zJlVo4-VQAWWB8ah5419L2EGhjLesQcohKex7p4ChCdWKCnahDE_oe__xZO8ax7gxXx3q2o5I340EZ6Z_4zxJ8Ep6SvB8oH34GzelmpuoeX7U7mK9zVZbp2pyxorelKUDuydZjHmc7yTftEJKknrB55LqVrkz5Wtve89EBY3gfSrlAtMZ1dPOpw2fRCuiRI1KmEYj71my5JhU02E24BOasvgyEHChLk6yeDGY0EZBc3G-A_936BDEBiNiigqqLmDygn6eH_i0YcH07PQTDtkTYZlT2zGAqgsOIJUO5Si6yTs7XrXxUbAclZFFVyyBLtiQGGKhNeg3u-wRbrp_jvyzZyfaZa51LfvQg2ev-q9mKnyZQ3RXvbCynKGTwCv4BJsFmI-0oJfQArWzO3I_WZgjBnVgUnQddbhAESSh4W_144hni3o3nktwu5oHBi42LAJXHPbs23CRHefqhrH8qdjUN7vCba1WHW-661qNglAaNB3gtb1fMlusOMPT2_Al95d2Nt7Zuu8wovrhprBBxMGP3ViCbXQ5sZDl5t_75WpcnoHYv7gxALDMzCl4xuj5NnAqbTr1VUmtrdwHTpRQE1udSeW9xB35eEVQA6a3FlyUFQP-bGQjFyirkiiH8NOyKoJwIlgYgx8E-BWsnP37Hblq3UPRpBt68KsWFlBk5ly3H_TCwCRty5TyBMc0atB1001gXcX7IRI5RuOuQvN4rpcep3K-Xt7DMmh7__G5kTIoEab-QnPxjHxkEy5ETfezJNzM_T5YFuFhAfs6rANiOrQvb3993QKpZQa12cs4wQ1macYhVXHEGlsp6j0ykjoIeUkDqhILmowPutkBqDQjbOpu25wgNWYizt8OdzvehvJOfWvwSx6EjWaIQ1UNdxcKSC994EWcJZJZYdFQx_OS_EE3X1dRj0aUNNlIQFYcoTAlZCplwkXHCfGPYLDVVBu93oNf0EY4hYFo99b8YdxeCSXJEx9IrKEI_CsfrflbuR3K4kUcm3Z9TSmq9j8ABUuSSmXlc8FH9Mx7faAMufWaXAV3euoyBcFBBhbxeiSOD2Yn-E7DpPQOx7p1gMIjeTyJvfy21KOq3MCgu9c5A4Dgt3b76hSJU_El7JeMp0-lhvv6K1wcEVDqczfXeZHAWjM3BO1pNXV3rRiCFqsTFi2QPoumuLq4Qd12Qh5hwM8sJzgWuNxpl--gMrTk-DZMR2eYIc0DRcLemRz-_tejXR63MFXn-wlvdmSV60Mxv0Lp-6JoBFaguGvAKcWLWyJtkQhaXdfuNGJg5yayiqzpKYg6Y0xQSJ__cm2pCnB7-57FaDkMYqQZ9ne0zfEn9ADZUReopkHKPZEYCrAwq4erajpON2jBfUZWtsdHmyH_ZjJ4Zw_M2_XkVRrkRfwQlS6mb2k-FwoI0DItJVZYpr7EFSgQTTrBonIVA9y9csXQKhAYe0lNYcDz_PvjoZNRoEcFIb46Fjacou5R261eB-3T4rViSv0ATbXUjGAMYiHA1Mbsrb2vIcL9tsy8C-WbIFhdFkFDuRobxf2vJgEUfnrBfa5eE0JHGEap01cgKP6r30LFFlICGRizwy0bLUbxxiV8yiLgEdv9Ve9mTkUuFIZvb5YWOlx-S7j6d2TO8C5Q88brUqXDDEdv8DL02tS4XOvrn5EELy3qp_peE9T0_aV5T_w0SkkzegZj5RNM_enVLzyKkvvGKssv62SCJ-DtvhSinskOF1V59j7SGHU7k4t0lTzZ5bBlHZR8VyfNc1fSpFDPCTlaXYvV_1Pttnacrjv0OVvgW1z8UpLhyKa1pTBHs_szL6VRMJ1ORTE7UtB5eJZCDfMay60dAkGz7CSHdbGDzFZo8rjbYuSKfVAFn5hHkp8MWiFoxK18sTyQOxeJjmDdSfRADvMkOvdt5PE6zA2qNR-5dmJBJ6rBY7TnuScfpTA7OfJhe_-9s9S-2tBz6H-I-XUvMmpcv3SZ6USQvYXLo1oxwTOW_BIL3LZpmO52T0WELFj5rraaoGV63UgzMecYoTpyXdZHGU9oKde0XzKIMKsqwL5-FKuWv6u3EZVk3ZhcSdOKqIBNs0GU_lHWd4J9US__pmaHB2SiIL5vE0Di6lOHKdA1w&ap={AUCTION_PRICE}\"]}","bundle":"com.p1.mobile.putong","iurl":"https://p0.ipstatp.com/origin/ad-union-i18n-api/28fb62df1b0c9dd96dd1a2206cf40344","cid":"1710424861130786","adid":"1710424861130786","crid":"1710424861130786","lurl":"https://rtbtrack-asia.rtblab.net/simpleloss?p=271&d=80&loss={AUCTION_LOSS}&l_b={AUCTION_ID}&l_a={AUCTION_AD_ID}&l_p={AUCTION_PRICE}&l_m={AUCTION_MBR}","nurl":"https://rtbtrack-asia.rtblab.net/win?p=271&d=80&t=wcIm3rKLUAeTTiiptPNypq8S4aGKjSiX2c00pNrZx1rg3x8J6JFd_UlNbkawVsluAPKku6gyuqDpyuQ4lKM4Adnq2telLHixtIm-cw2CXUUxqZxgp0Vnu6y4QAsUQV6EBIkS1RNjbz5N36dqg79prISxTFa7cJFrf5UG0FyUG2UhCJXhvt-XKevL65_-lAcgGQ4i1igwGV1b889Mzweo8OEge_dDupQATeGdrRTBWHjB6oes6vKZDZ3g8OuzSB2rGZTyghPj7ejOaSTHuQ2yrR2dfUU37sj-8XJeIPSOwRilyNxkf9XXAcClyuKW7Of3wzHbw-oW7T1T2HuPb1ncXlblWTbk1neDepSxiTfqNzhwXYjSqgEa0Lh6vPeAX9zB44frdEp8pJQtuLyzbAoxjtuzq_1WG7twVpRWhiQs91PX1ft6mVQkRaO6mFY_X24ww35csc8FzUMtEav1NghIr8W46ZA2Fot_8Nulec4kbELmxR_do7innoKDOc2R058QJeFekvj2ci8O660VBhQeKIFT9xlGrpoTtmgeMj8WGLY","cat":["IAB7-45"],"adomain":["tantanapp.com"],"attr":[4],"ext":{"crtype":""}}]}],"ext":{"source":""}}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	// webeye banner ad
	r.HandleFunc("/mock/868/get_webeye_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"DspId":100095,"WinPrice":0.05716640627817538,"DspName":"webeye-DSP","FloorPrice":0.01,"ValidImpDuration":0,"BidPriceCoefficient":0,"id":"0d85aafd-407b-4b8a-8660-a68e423f7af5","seatbid":[{"bid":[{"id":"1","impid":"1","price":0.05716640627817538,"nurl":"https://rtbtrack-asia.rtblab.net/win?p=271&d=80&t=6pNRXg2e2TSwJlAYX9223raysQ55--RhBRBhheFyjKbF5v_6YuMsJ9iMz5QUCobECHjA1OoAGhB3wR1A_aguhvVuXTf2-Q27sxbze8ejFIRnKigzKjChJmZFVVQTpk8U-LV4DesTSHAKO1xG4TC_GdtZm9TZDZRMB2ErIb6bTIaCujFhY8a9v58SuHL0gR06sqOmpE1XnBNWpZPiCQ6rUkeAumNxZ2oWl4yrPeKm1i75yNINmxbni3e3uK543zCsryxpWISNbDawRBY-Zk64sFe4R0W2u16hhtSA28cDHWj7KtTSdpuFVV5Kx4vK5y7OtswCiBqG_O6G2HsZLbAqslrPq95vUkPKYHyGux2JVk7ngdgDozEv909WrQzflSYyuDaxc9tfbEAJ6UmpLCviWqrHFWhD461rWvYBCXEE3NA7Mg-QrFIUkdkc5c0qOcKxJfQyGFIWY_Bq9h6CL_bT37ArupIFhmBQ8ZByVqufkQ0bxBquDKaVKV-k367HtR7jpDYTlcOTaiioBxJhP6oGvXprMAEK5jw-6bC1BeXZr0E","lurl":"https://rtbtrack-asia.rtblab.net/simpleloss?p=271&d=80&loss={AUCTION_LOSS}&l_b={AUCTION_ID}&l_a={AUCTION_AD_ID}&l_p={AUCTION_PRICE}&l_m={AUCTION_MBR}","adm":"<link rel=\\\"stylesheet\\\" href=\\\"https://sf-tb-sg.ibytedtos.com/obj/ad-pattern-sg/static/images/fallback/style.css\\\" /> <script> window._track_show_ = JSON.parse(\\\"[\\\\u0022https:\\\\/\\\\/impression.appsflyer.com\\\\/com.grabtaxi.passenger?pid=bytedanceglobal_int\\\\u0026af_adset_id=1737755010331665\\\\u0026af_siteid=900000000\\\\u0026af_c_id=1732593610523698\\\\u0026c=ID_TT0361_CLUSTERALL-CLUSTERALL_PAX_AG_ALL_120522_CONV-MRTA-APPC_TTL_%3CID22OGPRO2AGCHURNREG%3E\\\\u0026af_ad=ID_TT0361_CLUSTERALL-CLUSTERALL_PAX_AG_ALL_230622_CONV-MRTA-APPC_TTL_%3CID22OGPRO2AGCHURNREG%3E_customaudience_VideoAd-15s_GM25rbyukgrab-sVAkt-\\\\u0026af_ad_type=15\\\\u0026af_channel=Pangle\\\\u0026af_ad_id=1737755016617010\\\\u0026af_adset=ID_TT0361_CLUSTERALL-CLUSTERALL_PAX_AG_ALL_120522_CONV-MRTA-APPC_TTL_%3CID22OGPRO2AGCHURNREG%3E_customaudience-1\\\\u0026is_retargeting=true\\\\u0026af_viewthrough_lookback=1d\\\\u0026af_reengagement_window=30d\\\\u0026af_inactivity_window=0d\\\\u0026clickid=E.C.P.CqIBtBeGiWbqzg9PMTu_wNrQNIKFAHjAHkmAA7lV48pOz-7Kfk9l3er1LKepP8yWmWMr5TuWHPnG-K4FV-EZ37sFTkLOlszvqDAmCbRz8mHOGg6-zGYvAOUcW9tXBBoFK5gbvPPhjKsgB2WZDuJOEiapnREnsu9jrFeBHgtRnidpP9ntOfD5f16XIznBaIDClhMVMt3DBVni7_XXWyfre293YtpSEgR2Mi4wGiB8CDz7xt8cJZ723bWpPq5f6RatyThwDOyVBzMTeRtkxQ\\\\u0026advertising_id=2f78cf99-5444-4a61-8d12-b3936bb7d91d\\\\u0026idfa=\\\\u0026os=0\\\\u0026af_ip=203.30.237.98\\\\u0026af_ua=Mozilla%2F5.0\\\\u002b%28Linux%3B\\\\u002bAndroid\\\\u002b10%3B\\\\u002bCPH1823\\\\u002bBuild%2FQP1A.190711.020%3B\\\\u002bwv%29\\\\u002bAppleWebKit%2F537.36\\\\u002b%28KHTML%2C\\\\u002blike\\\\u002bGecko%29\\\\u002bVersion%2F4.0\\\\u002bChrome%2F104.0.5112.97\\\\u002bMobile\\\\u002bSafari%2F537.36\\\\u0026af_lang\\\\u0022]\\\" || \\\"[]\\\"); window._track_download_event_ = JSON.parse(\\\"[\\\\u0022https:\\\\/\\\\/api16-access-sg.pangle.io\\\\/api\\\\/ad\\\\/union\\\\/redirect\\\\/?req_id=cc9bnrjo2gd8v35jou1gu2304\\\\u0026ttdsp_adx_index=101\\\\u0026use_pb=1\\\\u0026rit=926901987\\\\u0026call_back=jLuLQ0Ck9ylk%2FQdnmY%2BBeBWxFQZuFGTr9SLbNZYhA%2FTRT8h%2B3AqnhfixHvEJ9FgGykFveaZV9GXMoDPQ59bn4WcutmKLY08ArxsIyZHTcvFmn1f43BvyA9gIEhcDJHOj\\\\u0026extra=A2E7G6dKvpI1XIR3UUZWansHJlSWhEORIZoqbjxQKIYqxHta7HHCDkLFd%2BAXRU6tLL%2BtWqjlFZJJbZBgB9Y39P1GBeJzyL6rQy5Wqr3f8MAz%2FWuWprGpUxvTJnbZ3dmyT4GEsZoITgXKl1e0cX0z%2BPgK1PdhDqGxMfraTe%2BvuOkcrNj8pvn8jeFSGC%2B%2FQqr%2BWF3De3erIn9PJ1mF3KWXNyrzGgPeCc2WAf9JMeLsVwNyZcb3oh%2F3yDtBf3m4LSviJhdLBShxrXUXzh301bUEYRjRMe6DZnfNAn9USR8QeFouquNZSFb5jfdPdOOwhQJws4QD%2B168D649KkJI58NN7UjUChjC%2FoYeS%2FmTF8u0G0W1Zqgy4YDUley585vFa%2FbzO3Pdy9cX3GBgCUBBIGEfx6zMaCkTh6ZDJQ313ogeDEnGXuR%2F44DvXDmlhSJ0oGYCx2QqkRcob1ZCI0Jwa5vSUPpzKmq3gr3VqiZ27lce6p9epbWcS06l9qcYhT4FjIexHvjxZPm%2B8j19wN6%2BQCnowlq200k%2BCostFB5ed8ZqsufHmjztMMvT6Cq590xJfYMsWc2tHtQ4KQgTxYrBuYKapXRZCTR0jpk6fKH4p10mwPTAFqMQjLAKcWK3NxgbMRZQsIVhuKq3%2Fa2GFzXonsNQdb%2FtF9pbXVlhC6x9bXe%2FQRn2psL%2BnSU7QJt0thpUC2GgAPvwiajUdrwLw4qcV2tnUIWgoQmbHD1DsdBeh8zv3CrnBjhPdx%2B9KXRxWmR%2F6ogZEtXbvWAvmiE2R0T%2B3E3G01ulacqujWJu6IjBau8B8Bkn0QSNzb6HQ50Ipw1xIYTfZp9X%2BNwb8gPYCBIXAyRzow%3D%3D\\\\u0026source_type=1\\\\u0026pack_time=1662172143.00\\\\u0026active_extra=1Ij5ALNNOiSs381yIub9rwEDvbdJoPQCPmWZGIRI%2FDIWX5Af5SW5RpygFcDhxCsPkNP88TRPxKSfGuXMyjvyYw%3D%3D\\\\u0026is_dsp=1\\\\u0022,\\\\u0022https:\\\\/\\\\/api16-access-sg.pangle.io\\\\/api\\\\/ad\\\\/union\\\\/redirect\\\\/?req_id=cc9bnrjo2gd8v35jou1gu2304\\\\u0026ttdsp_adx_index=101\\\\u0026use_pb=1\\\\u0026rit=926901987\\\\u0026call_back=jLuLQ0Ck9ylk%2FQdnmY%2BBeBWxFQZuFGTr9SLbNZYhA%2FTRT8h%2B3AqnhfixHvEJ9FgGykFveaZV9GXMoDPQ59bn4WcutmKLY08ArxsIyZHTcvFmn1f43BvyA9gIEhcDJHOj\\\\u0026extra=A2E7G6dKvpI1XIR3UUZWansHJlSWhEORIZoqbjxQKIYqxHta7HHCDkLFd%2BAXRU6tLL%2BtWqjlFZJJbZBgB9Y39P1GBeJzyL6rQy5Wqr3f8MAz%2FWuWprGpUxvTJnbZ3dmyT4GEsZoITgXKl1e0cX0z%2BPgK1PdhDqGxMfraTe%2BvuOkcrNj8pvn8jeFSGC%2B%2FQqr%2BWF3De3erIn9PJ1mF3KWXNyrzGgPeCc2WAf9JMeLsVwNyZcb3oh%2F3yDtBf3m4LSviJhdLBShxrXUXzh301bUEYRjRMe6DZnfNAn9USR8QeFouquNZSFb5jfdPdOOwhQJws4QD%2B168D649KkJI58NN7UjUChjC%2FoYeS%2FmTF8u0G0W1Zqgy4YDUley585vFa%2FbzO3Pdy9cX3GBgCUBBIGEfx6zMaCkTh6ZDJQ313ogeDEnGXuR%2F44DvXDmlhSJ0oGYCx2QqkRcob1ZCI0Jwa5vSUPpzKmq3gr3VqiZ27lce6p9epbWcS06l9qcYhT4FjIexHvjxZPm%2B8j19wN6%2BQCnowlq200k%2BCostFB5ed8ZqsufHmjztMMvT6Cq590xJfYMsWc2tHtQ4KQgTxYrBuYKapXRZCTR0jpk6fKH4p10mwPTAFqMQjLAKcWK3NxgbMRZQsIVhuKq3%2Fa2GFzXonsNQdb%2FtF9pbXVlhC6x9bXe%2FQRn2psL%2BnSU7QJt0thpUC2GgAPvwiajUdrwLw4qcV2tnUIWgoQmbHD1DsdBeh8zv3CrnBjhPdx%2B9KXRxWmR%2F6ogZEtXbvWAvmiE2R0T%2B3E3G01ulacqujWJu6IjBau8B8Bkn0QSNzb6HQ50Ipw1xIYTfZp9X%2BNwb8gPYCBIXAyRzow%3D%3D\\\\u0026source_type=1\\\\u0026pack_time=1662172143.00\\\\u0026active_extra=1Ij5ALNNOiSs381yIub9rwEDvbdJoPQCPmWZGIRI%2FDIWX5Af5SW5RpygFcDhxCsPkNP88TRPxKSfGuXMyjvyYw%3D%3D\\\\u0026is_deeplink=1\\\\u0022]\\\" || \\\"[]\\\"); window._download_link_ = \\\"https:\\\\/\\\\/play.google.com\\\\/store\\\\/apps\\\\/details?id=com.grabtaxi.passenger\\\"; window._image_url_ = \\\"https:\\\\/\\\\/p0.ipstatp.com\\\\/origin\\\\/ad-union-i18n-api\\\\/0a43666d328dade0b97811d8fea5151e\\\"; window._deeplink_ = \\\"\\\"; window._event_domain_ = \\\"\\\"; window._creative_extra_ = \\\"\\\"; window._media_request_size_ = {}; window._ttclid_ = \\\"\\\"; window._pixel_urls_ = JSON.parse(\\\"[\\\\u0022\\\\u0022]\\\" || \\\"[]\\\"); </script> <script type=\\\"text/javascript\\\">if(window._track_download_event_&&Array.isArray(window._track_download_event_)){window._track_download_event_.push(\\\"https://rtbtrack-asia.rtblab.net/clk?p=271&d=80&t=6pNRXg2e2TSwJlAYX9223raysQ55--RhBRBhheFyjKbF5v_6YuMsJ9iMz5QUCobECHjA1OoAGhB3wR1A_aguhvVuXTf2-Q27sxbze8ejFIRnKigzKjChJmZFVVQTpk8U-LV4DesTSHAKO1xG4TC_GdtZm9TZDZRMB2ErIb6bTIaCujFhY8a9v58SuHL0gR06sqOmpE1XnBNWpZPiCQ6rUkeAumNxZ2oWl4yrPeKm1i75yNINmxbni3e3uK543zCsryxpWISNbDawRBY-Zk64sFe4R0W2u16hhtSA28cDHWj7KtTSdpuFVV5Kx4vK5y7OtswCiBqG_O6G2HsZLbAqslrPq95vUkPKYHyGux2JVk7ngdgDozEv909WrQzflSYyuDaxc9tfbEAJ6UmpLCviWqrHFWhD461rWvYBCXEE3NA7Mg-QrFIUkdkc5c0qOcKxJfQyGFIWY_Bq9h6CL_bT37ArupIFhmBQ8ZByVqufkQ0bxBquDKaVKV-k367HtR7jpDYTlcOTaiioBxJhP6oGvXprMAEK5jw-6bC1BeXZr0E\\\")};</script><section class=\\\"pangle-banner-wrapper\\\"> </section> <script src=\\\"https://sf-tb-sg.ibytedtos.com/obj/ad-pattern-sg/static/dsp/zepto.min.js\\\"> </script> <script src=\\\"https://sf-tb-sg.ibytedtos.com/obj/ad-pattern-sg/static/images/fallback/banner-1.0.4-5.js\\\"> </script><img src=\\\"https://rtbtrack-asia.rtblab.net/imp?p=271&d=80&t=6pNRXg2e2TSwJlAYX9223raysQ55--RhBRBhheFyjKbF5v_6YuMsJ9iMz5QUCobECHjA1OoAGhB3wR1A_aguhvVuXTf2-Q27sxbze8ejFIRnKigzKjChJmZFVVQTpk8U-LV4DesTSHAKO1xG4TC_GdtZm9TZDZRMB2ErIb6bTIaCujFhY8a9v58SuHL0gR06sqOmpE1XnBNWpZPiCQ6rUkeAumNxZ2oWl4yrPeKm1i75yNINmxbni3e3uK543zCsryxpWISNbDawRBY-Zk64sFe4R0W2u16hhtSA28cDHWj7KtTSdpuFVV5Kx4vK5y7OtswCiBqG_O6G2HsZLbAqslrPq95vUkPKYHyGux2JVk7ngdgDozEv909WrQzflSYyuDaxc9tfbEAJ6UmpLCviWqrHFWhD461rWvYBCXEE3NA7Mg-QrFIUkdkc5c0qOcKxJfQyGFIWY_Bq9h6CL_bT37ArupIFhmBQ8ZByVqufkQ0bxBquDKaVKV-k367HtR7jpDYTlcOTaiioBxJhP6oGvXprMAEK5jw-6bC1BeXZr0E&ns=l84an9h49lG1MtsHPlulbrwptsvHvoSsq2Va6k9EyuAnVAiDCuN-f-BF57BfUqNFJxfTSflyV5pSwe5zBs_dpXm-UGmEIIDqIQXjoC_QyTt6NX3WQ5ANluQqqGHgsUeo1gzlafPin3q7n-CrjzNZdicnmH3ZuqvodVeBWgOE9v1em0BEhlPqBz7abGyOOpmoRA6YxQHI5dnYQoUr10aCAfMPZi4pjHyKhaaH7_TjoWtJMiePEhXxnnAvH6PpWgUUgr_3HPeR3OAaLq3qaEVZtYr7MPoq5yH_j-nuR9indvVp8pb653Qg0kvYuIe6KW6k4vC3-D9ajRZcSXvCpv5KsHfavf8eLfcZFOYmlfwJJzn1Z32GZMVGDwc4BeqpCcSE9yQQdEcMfY264OTYEZ_8WTtxoiAaWZHUSx0NNxWxgA7hHi2A9X7Ev-N7TKk_62MWAyHRmO2B0DqS9YGDBJCitZ7qReO-qKHI90GZzLTQsMoVLGRndLD1LaU9TJZVn2HV0KWps28NBZUBvtQz_xrjBn7X5DEKmKuTdDieDC2Axsm_AFeP7ue0PbhjhDUZCHPBx-0NsHJXSyoc87ZcfDO7jTlbq16sm2jDPBOY8TLpkvKwO9Ngrni8ChbgUD5fjmuPF3Bq0JoIHGLnr4XVkfreZ-L0Kg4GikQ-cMR0RfXiiGRw_tSKmelbhdUOp9tlrabs-myRBPwnUkEel63vyBcINdCEwv5PYB8fiZt3FoQcFk43DpwAikuz8xmcb6FEtb2lw_kgQ8iojEszuqciE4S4AexIB6gWcp7nUmz6SP1h7vMCN_IMKmsk9qgpZzlpqjfHlUywJUIQgaj3KaepqOzy0PJFPTjSge9f0CSDBFG7nQocMsSi0xI2JN_eauGD0HmntpjNwHNAjYv8L9rAn0TfH_o4jOdy6XpTXgb3ukhwzyNdyFJp9jxXwx9X30vuDN088oDpf1hwq9ECsuEkZasu3vnT6VhWOO300Eg9mLEeks9LRmkgUkqs_6hCh9OVK2lyuMkCsk5MkBQIvva02UjsjPGUmPQn1yWjZr-raV6L-ZaIvHK7MzJ_ZpzDupaL4CCdFgIdiD9o1OND3GY9LupTbB8T8i6JomHaKfq-UwYL2a7_YIMqCjI0Fw-VpCItF6mxQgK7NcrGUHmwQd7ISV-6X-FA8F-cBx-Ouw35Hshfmms1yBYlzudaDGFdizTyY7Jm9bep3xKU28RwS9QzFQkrOsSJfwTPUo8y91NOH_9Xi84cVPC75UmReGMqvZmZ3p2BIN6BhDIqhDtaYpOk1WnADPsUeAb46RBKLBzHjRckgetrPIY_C3kllK8AlKbeGqf669-2xsCtaZU9h86ShLxHj0O5K0geWgs8KZhj-_k45cS7PhEdvvmTRXcdWFsLMPBCmOAxdgufj6FI5p6sGGx9hSlMN8Byd1i5RiDHL8hqUr-9AnZ5hHWrhsiPciGXwYDfQwV8bYgpfaAXsqC83Rs00eN-X-nE1CEo6JhnIYZxQl3rxNM6R4PXsi9p8ANb46X86uggosXShMAToj70zmzAqJITG9CJa5X2wLq_PwwUFRj2nroF2_YTq3NiyBTIGx9TICQUFO9fmlStzVR8CvzTLwd0C5KvOVHPkGjdmr6dmGO6Y3CHr1nqBzPp-n9OIKQ_lzwMZ3dyJLjasVjGffVHgO2Yfn9mU8Z9i0WPZ2YWq0lid8SYvLyBdNqeXgxY1aLMVCBNpHI3cZ7av2SksEubqMjWsGXt1J-2pvTcH7AivDSCtRjtqh2xhqSEzcVntHtUKyblWD_ltkdaNND6yMLa_w7t2wl8lcCQTlzlbf3rJ8GpJF4zzfEVF_I8wOVwx8H2mBU6mrmP7F6Lcc8DMeVc_lWM2XRv8kRoIDQfrFh1ql7xDj1Kp5hREIqhCGnUZgitiagM6hFYtICNUlPIl5vFaAAxntbO0Usa2VBM_ziTIroloVVSufYlJtbk8UGKVUeyR8RwCLiBvbwfKrlzzzKXc4HcBL4M6muJUwdC5Dywpn7O1Wl0jN7irM8DejIMtnegAbfgHycYRmIxCx1p6ueJS7epGDCiKdcFCq52KjAyH_RGRtD1b1x6di5kcqErUB7pediXh4A_UMtfxkobJz7meoYqMtGKsC9LtZIaJxyM1nX6crEbHtW7CdVWWZ-NsShmBTMoFJnm1an1_HmP29ravzPWn-rP68At7RgUTz4w9Z6OBH7chdIRZ0_T2FnW1gvfCNRReYqC-yEzIMCtowuG45XsNOJ0pb2qYvFQQZt_X83UOtdiHt6ZkU6912aBk5qXS_7uG11VBKLeLdtstb-rogbZQkdpHP-W09h5GnG36ImixiCdDSSvJLrRCRYAd0dYUKkaNXOlKxO2xwrdgop2Jtm5fkKm7Cl_66N_ps0-acbs9taEQ9MleIXWLIDgroVhM2mTunOsu30ShbiWIzLIX9A-TTZF2XZgzKOGaqihvYstyG-eo0mda8suhP0SA7PL1VEMPHrRdMCcgUr9zt1oXLwpUO3KBDnpkEl_WOSH8dwH5nuqcFAK10_lRySR-v_qDedGuN4xF1l6kmIRyxWLkAcsk0MLwNJMy0afYWgEK9GdmAbOzH5bdJN9wljWAGJ8-qYfn_zDNRCkBLN0aiG070JxTznI8qLiy1CF9anh9Ys3hYygbd7mhGmoL04vPPxf9cICZnncHUN_CikPEje33M2BhsARutcUT5sTEnEnFhSkKy5mXQuKUoow4ys9A4HF38qa-egpuyuOlTJrHesNSM5e6Oua54xrL60gBiac2AAo_1o5XQ7WbCSSUi4GUVRCdG6hZrdPEZg5qbaSdZrJsX8Yy3L9SvJoQjO50prRQfsx80NYZcKvA3_otKXFCvJZBzAIk1pOAQctsnKIQ2dX2pY35zgLPeqB1UcHEqR833YlpdmqjEhCax7Ln3kyCgODncObymbHYmFMmldFAcNfBCHCfHC893yjjMfW27m7Qn0dNPvwDb4m7ry9mI_hYys3CrAyd1U9UvBH2OH5Jvfj4TtWxFPsIdvN7G1X-cG1QcrF0H7BKnpVSiLUijslDc9D6_qFLn_9jobVzyZJ2WPLrI4L01ewvr9CpQO3VjytuArjmXe9XPhRwF__ZVOQBiHE4-xqVf4Lmxcv3V8-QA1TOgcaP3UVNhncsAj6ZRJ0LKXHy3EHaBVMI1GhWwpuPLkdgiRGwmZP7OKf0VpoQPCkcvHTuKMQywkOLuzSjbMDn-DHoUkS0nQq6JnRKSvnxJZrX_GlVDAj8fPiD-KSU2xMOvMg4LP7mZ6RIYM509Opqx9wRO0dXwfrfYcXWGUdEJY0q6qQJ8Mh6hbwqCs0c-uaX_W0BiF7hKCwc6Gjxl6ntQfroN09qCcsMRqhE4gUJjsvVFaAnY8ffPGYR-ymtMuLQGtK-n-kDsy4tOboVW6Bjub4RhWXHLLo9Cr2vGT-XaL6IPvaNy6ulsWTxrmV0hZp0CwRx0BBQ0dYZDwYjBimVSA1FyhGesJMvdPuteoyKXQBgfFEENg_Q1cGZRKThJjS5xZlgnrZRJ-UupekI9jQiukgyKuQiRWmvOZ3NSdA4auPBu-uz7Yc_kXWokG_zVx2nV9V-vwBHchl5pmQzvH9dslu2eswdcmBFGy1hGmWeibgJ_tuw3feb52HW0Yw06O1K4KfAGZv0RHZez-HvWyf7ZBc99ZDrR1JtuZy1Pc1p7nEz38mjkg_83BXMqaZu6ehoFMKrwpXM4vnCiPFW-6OweyOIe7oZOM3XTkG3DgluhsfC8lqonIU-w4-qzQfTNkq_8MDnp2Ie0Q0mupdwX27LeQ&bs=l84an9h49lG1MtsHPlulbrwptsvHvoSsq2Va6k9EyuDWc3vuy2eNFU8TNgSitfShHPFRc09j-nLeyptp4tptNbWfc3L0PMyoEMymnP7Z_IX_B-1nX1jIBOx8KCi95pVAmeZKeoqNEtzBzCk4e1Txfryj4fJBt9ghN9pmSXKw-x2sD66gCZcissMtRkXUYySvsEny7BCz-aFWpoGt6qUqdqrbRrS-DCq5h9SyqmmgrLqI0jnHkkjtJPTV5Jj2nER-UMcq4k-alkqCRunsjY0VATT7kC8wYu3TgVdU_yk-3c4GokR2_0g8JC-gUx_ZwUHpM2EZZ7Ak7oOEBE6kc01BFIGPAIZxcbZkLLC16gmQUn-WzY3-mUDJ3i7Tm1eTXhYi5xDsVtvQ-h2ZHj3YNt1vBI6l7Qe6pb1JsYbdM91oX76PHGr3rRTZDo1f9m1nIz0V6bq77tTwwkhbxsylC8-NOLGpn1oY1836UwHY0ZiGGV4nqC9tPDoeecd8-GcTeVj7RIhkpR9uvAT0k63ATtpQ5jnRJUNzWz9odO-PqkYQDYUpWgsQXgH_0sEINq9p-Xe7gCL5RoxhLgfojs45_VSX1k9A1RJ3nqdUixWD31-DNXOZwze9l_vUa74Dp_C4j-FsTfUCb47-xj4KiG839xLP3YiUaSqupDtjZGRwNtrdStatHhA4gEgd_7_YWeGLtf6VplAmOF61Na_mdkk_iSziPfb9wSAc0Z8jlUvmQ-CPW1brK4UgZgEYKO5kfBh3MfOWDW-wcHpzSlKMFXtSFvz8_SRS5vrf8K8ixWjOPhZS9xC9QMI6jSAHqqHOM-eH0a08iGMSCK161tqwS_pHSOvqWAcXfmhsQKhS09l-kMAD0jMzaqVS0JVjQkox-_HM5cV5BKzhFWLl71RhEl-C_Pk41d5GeRcW9VYOiDZQ9An4ek0ZUkfQ1f877XOQjvOvtSlpGIPYmaR_yR20tjmca6P_7g9zTNHdN3ytllvJYRtcnaBOp6Um07SOWA_wmTAydyPz_cD2TR17ncDk-06G27JziT7FWgvI8PN0z1YkAjNykNB35y3ZBDSs9N5w-Oi1qcEo8K8y9jRpAecSde4evySrzVanAlx2rt_PvSrQcgm2haLD1ZMqXaHlpDcrsiaUOTS6gSdaZIK79IRC-rdRCjrnxrRa6wU04XNo78n0Bqd1hQ-90ACNqcDMA8p_rMg1H6HpgCj9JeHihxk-vFZzc_qkfd0xVfSYkg_FK2kN75pKP0TyRsrV0QY6rZobNMMiCTFXRCntvCzQJ9jjJDXdcD1g1Wtd8w-X_e_HCdU8JAdBx_tMHcQdpFLO6lQLpGxsIkuq8F0uFjmgvpl3cs-OwXqXmpJ9Oa0Mn9dxcNknasriYRSopWAAICep_w9EMIeljjY8HmhWYYqbltlxN71w2haNt4McSx9YbCpiem_25NeoqdmO4VJ5f13lUmxOE7zhn3HQFvTxbfy8bvASeb-rzNxuiH-rij2Syw3I-sQ-MfM5Ex_u1z1xWXwARayVp2thgdfQJ3TyKHRwu9cz5zpC1ZXBeI_6VrX2Oal10YlKXegQKdD-1pgzhsHGOwqEIhOyHJPaxim1H1MIdP2C7cv_i-41POOLp8PSyA0qd9xd0oA0hgq1T05GAnyCDYCl22KQPEoxEZwNmJOG6s8kwPsyaH9ssKH1qDOJQBCZjh4gt1llKazAXbS9QKCQNbYRB3Y_NrpqnI70zqSlcfvBrygHM0Yl4WdcWHbyzSoV2vNgHPrI6H7Tsr9-bEuQnM6HURyzHKR65uabJ9_n2blJlJrrzKGtEXBBb-JgIR00MCR78aYmAXhSNmYYFfPlhMtQ09PV8_FT2q2kq9ApNrdUyuzDf-F1sAMIGH9EFBgy47UY98kZeszhyg-ECCB9-6F6g4L5cM6fBQJphMUxVm_jYc0vBEcNvK4fNK8GXqhPPnK0HHmbqA9_892vwpak3rCxwLQgFgLbn42aGo404pWgaLIrsER3Mu1GC5cqgoxL-sihXxFcqQtiIcyfEXDs1ViSuYDOnNlpHWJ_4kEZoiQDUGwKAmewAFHFwnoYwC42Dq4iaI7OhSSfWpIKEEQIYW_CKzqxVnx56PSoT_qBn2FiSLmZXI5MwOIGNwfsJNBFy3V9YNukqUWjQdMJbM8J3itNG9mIfFMyT8VozUGKhdtNyh9ebmA_ISAVRgJlLX2KYLZGZiQAJI-iJZvUz_OLpXD6BPOy2XokR6MSHON2U_yyVpNwtBEGsUGBBPvz_yi-BFiDT4phnWL4ssv8T4SohuQtVqwtCUj_pmMD5lxI5St-WyzZh2V0g7h7ZGzJ2C-F_fchWWT8j8RnCcd3MzRwZ3RopyoEApomJE9RVoo-S2m3lG0VAxJBWFzSzh7y7CsVR0HEiWq0bD65ntfYP0LaA4Ja-euTUDiIBdCAHo5NIKAEnD5lVGoytzeFKanCb8AwjqSbj8yqJ1zApoaYyxTje1JzUvljWNMhBFG30O4blclB4IA9V3odX4ZmWMfZKq3_wzcuG-Ecxopg5jsj9svpbA3CQ64AXDk-4HXclpst506tQ04tyZfDcqumIfD091WozyCcJhHVsybqAW2ugkMO0K8pMKEOSQlKChcC3dm4Z9Idwq1x7OzTL_1VLgrVMjEKU4_rpSShLKyEtPmAGNKNVPvPRt8oWdc-FnY-WXAbMQAIlq6ndYoWpDuZcOZYRyxRYf5XSuLFC7PpiuI0Q82-kTbOkQSakjw9TdIF0weo_wlPldxRaOz0HlL3LU84LkPjlZI3_w7JC9jEjDgl0AaqQ__Kt6LaDKpqEqrOPJqGnh0I4rNQZbcznanmXrGANMgOy6sAz-zaDRsdr77zpu1t0hra3SAvWmXhSBJ7zSRptytDKmMVozgmZN7P3DKuFv-J2Uhn7x5n3LClbV7C0Fi1q9A9sl6sJCqLo9GdYG5rh5hYTbdzrq1TvMct1Ak7PkXxzfaXwoEAzgjXpIlYAn0_aXPa_dvdTigmdmcxblFooJqlmKo5YOVm-28K_7jB4DjWChfM9txSoA83RhJB6Y2XhD9bgpLFRWkMYOrmOk9EetW5AzienUar0GfEWUIVCwvpMtH30UYYVdEjgxlOuV25d8hq4py6XHfawQBrAFGp4Om8eR67x-xegeufYkEKr-6H2BPegqSFbgDj5VQ613Co4xeK-cq58RRhYUmGFErVzTEbAXMqZqMWHp6J5YLrHlljlxSqw8gPgfPEV19MfpmO4VR_vG1cU20E7Y2VJHhdjWK8kI4IL5-CES2WNLRJDJHo4HL_v191DGnctGoVnhutecYCKUx9hZoCpVqqT-Y0O7Q8bgMHJHNdc-h6SRvhTmb_GkabtV0nxJDTnla5zR-TKQhTEm6j-F3-tCzCSm_1IcFJxATALNPm4JzDM-tkJphcFJddm9x5TWOweIeWSrdDQ3as-vRaPAJShi5b1p3UDiCfx5THKRm-KUnJjzEi5YpEYDCEH9Vs6jdLE6pMHjLsBKb5lLwnA3D9j4JpWvNlnN2hNsdsV_okwGf3l6XLbANo6zRIuXkz1YApih0_IiOnFblGeEAzlyEAoolwxEpXVbwsfmBrUr3btYwWKqsDfEDbXr3qP5JZfBtIHckXCGeKKMWpqQCqbII_E719jR4M1mfJjoydMhjDAAdwxA28IBtcZcKKzkhdimAua6yTFjpG_40BlfdxgjmEh2yxZXjv5Av-6JPSIX0Bd4l3XyyrCEK8sNft_JmGMU-xWaYuN37kKFKjeuqvlvLbComOKWNbGhR_dFzF2mol3oTHaivwf7y8jz4jJyHxQZqIv5bvmj-5_hHbdsKm387M65Wdc1w7NVtaq5N12DccziTUJHWwbSlfgWeOKSMK6uTd1-n5Bj2svFdBhspbLC56&ap={AUCTION_PRICE}\\\" style=\\\"display:none\\\" /><img src=\\\"https://adrta.com/i?clid=we&paid=we&avid=80&plid=1737755016617010&caid=1737755016617010&siteId=0&kv1=300x250&kv3=2f78cf99-5444-4a61-8d12-b3936bb7d91d&kv4=203.30.237.98&publisherId=271&kv7=271&kv11=0d85aafd-407b-4b8a-8660-a68e423f7af5&kv12=Vungle_SSP&kv16=-7.803500&kv17=110.364600&kv18=com.snaptube.premium&kv19=2f78cf99-5444-4a61-8d12-b3936bb7d91d&kv26=Android&kv28=OPPO&kv24=Mobile_InApp&kv15=&kv23=&kv62=0&kv63=&kv25=Vungle_SSP&kv27=Mozilla%2F5.0+%28Linux%3B+Android+10%3B+CPH1823+Build%2FQP1A.190711.020%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F104.0.5112.97+Mobile+Safari%2F537.36\\\" style=\\\"display:none\\\" /><script>var imgLoad=!1;function weValidClick(b,e,d){try{if(!window.bannerPlugin.utils.isValidClick(b))return Math.abs(b.up_x-b.down_x)<e&&Math.abs(b.up_y-b.down_y)<d}catch(c){return!1}}function plugin(b,e){var d=document.querySelector(\\\".pangle-banner-wrapper img\\\"),c={};d&&(imgLoad=!0,clearInterval(weTimer),d.addEventListener(\\\"touchstart\\\",function(a){a.stopPropagation();a=a.touches[0];c.down_x=a.clientX;c.down_y=a.clientY}),d.addEventListener(\\\"touchend\\\",function(a){a.stopPropagation();a=a.changedTouches[0];c.up_x=a.clientX;c.up_y=a.clientY;weValidClick(c,b,e)&&(bannerPlugin&&bannerPlugin.emitClickTrack(),bannerPlugin&&bannerPlugin.userLink())}))}var weTimer=setInterval(function(){plugin(150,125)},200);</script>","adid":"1737755016617010","adomain":["grab.com"],"bundle":"com.grabtaxi.passenger","iurl":"https://p0.ipstatp.com/origin/ad-union-i18n-api/0a43666d328dade0b97811d8fea5151e","cid":"1737755016617010","crid":"1737755016617010","cat":["IAB24"],"attr":[4],"w":300,"h":250,"ext":{}}],"seat":"2"}],"cur":"USD","ext":{"bid_sign":""},"vaststr":"","bidid":"0d85aafd-407b-4b8a-8660-a68e423f7af5","expire_time":0}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	// webeye video ad
	r.HandleFunc("/mock/868/get_webeye_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"id":"a24fe97c-a63e-416d-a357-8f2e3c679b29","bidid":"a24fe97c-a63e-416d-a357-8f2e3c679b29","cur":"USD","seatbid":[{"seat":"2","group":0,"bid":[{"id":"1","impid":"1","price":0.255,"w":720,"h":1280,"lurl":"https://rtbtrack-asia.rtblab.net/simpleloss?p=339&d=65&loss={AUCTION_LOSS}&l_b={AUCTION_ID}&l_a={AUCTION_AD_ID}&l_p={AUCTION_PRICE}&l_m={AUCTION_MBR}","nurl":"https://rtbtrack-asia.rtblab.net/win?p=339&d=65&t=55bW10qLPxlzQAXbMZUcWsf-OhubisbBjtjwjqdktKIJr2dmjrh-OmixU60az2HDosUCkB_GdRKL0zk60bPraXxkcuFOontU8dzS2Paqo_YKQa3DNJAh50yUvRWiacOSfOxGfUDUVhYtGQ2Ab7dIdxVZmCvBgmyZKNbyBjvQqdl0QsOEF0wals32i2chByx2fJicH9S5WC8D09hUuJLxkQ96S1pteutWmGp8Csv6C62OOzBToqZBk3nqBrJuIZTfbUoUrEYH7kQ8SeL6ltrcNHIUcp8sS6OdXHkirwusfPbJ1ExBXYw-D-RoojdsZcwIeNx7SKhPHeOqCr3FoS1ziFmtMvkvODPVJGTVxVmHSh6Yjy8b4fPZ8cmvgjRH4hsQ-d8Y3c3rtSIrYdgoptniln2pCXNfMxtrJnWtd1BIeI_IogEQVgztsHhn1O19KhKV7KxmEh0Nsf2H0kSTDI-yaYQtHGegEkJHITBnedGLwTsLaE3ri_pl296s4wZNNs7UKkOTtNOd1QiNoLs5foAl_bpXb3ESxOfsyWFlAN_dSinbpgVWyJmS_TektfONub_N","adm":"<?xml version=\\\"1.0\\\" encoding=\\\"UTF-8\\\"?><VAST version=\\\"2.0\\\"><Ad><InLine><AdSystem><![CDATA[TaurusX]]></AdSystem><AdTitle>com.zhiliaoapp.musically</AdTitle><Impression><![CDATA[http://34.126.85.24:4000/imp?p=webeye&d=minidsp&b=com.zhiliaoapp.musically]]></Impression><Impression><![CDATA[https://rtbtrack-asia.rtblab.net/imp?p=339&d=65&t=55bW10qLPxlzQAXbMZUcWsf-OhubisbBjtjwjqdktKIJr2dmjrh-OmixU60az2HDosUCkB_GdRKL0zk60bPraXxkcuFOontU8dzS2Paqo_YKQa3DNJAh50yUvRWiacOSfOxGfUDUVhYtGQ2Ab7dIdxVZmCvBgmyZKNbyBjvQqdl0QsOEF0wals32i2chByx2fJicH9S5WC8D09hUuJLxkQ96S1pteutWmGp8Csv6C62OOzBToqZBk3nqBrJuIZTfbUoUrEYH7kQ8SeL6ltrcNHIUcp8sS6OdXHkirwusfPbJ1ExBXYw-D-RoojdsZcwIeNx7SKhPHeOqCr3FoS1ziFmtMvkvODPVJGTVxVmHSh6Yjy8b4fPZ8cmvgjRH4hsQ-d8Y3c3rtSIrYdgoptniln2pCXNfMxtrJnWtd1BIeI_IogEQVgztsHhn1O19KhKV7KxmEh0Nsf2H0kSTDI-yaYQtHGegEkJHITBnedGLwTsLaE3ri_pl296s4wZNNs7UKkOTtNOd1QiNoLs5foAl_bpXb3ESxOfsyWFlAN_dSinbpgVWyJmS_TektfONub_N&ns=&bs=&ap={AUCTION_PRICE}]]></Impression><Impression><![CDATA[https://adrta.com/i?clid=we&paid=we&avid=65&plid=1337026&caid=345135&siteId=0&kv1=720x1280&kv3=b3f1d810-952d-4b06-beff-2bdf0672d5e3&kv4=114.122.73.251&publisherId=339&kv7=339&kv11=a24fe97c-a63e-416d-a357-8f2e3c679b29&kv12=Vungle_SSP&kv16=-6.921700&kv17=107.607100&kv18=mobi.mangatoon.novel&kv19=b3f1d810-952d-4b06-beff-2bdf0672d5e3&kv26=Android&kv28=Vivo&kv24=Mobile_InApp&kv15=&kv23=&kv62=0&kv63=&kv25=Vungle_SSP&kv27=Mozilla%2F5.0+%28Linux%3B+Android+9%3B+vivo+1904+Build%2FPPR1.180610.011%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F103.0.5060.71+Mobile+Safari%2F537.36]]></Impression><Error><![CDATA[https://rtbtrack-asia.rtblab.net/vasterror?p=339&d=65&c=1337026&b=a24fe97c-a63e-416d-a357-8f2e3c679b29&error=[ERRORCODE]]]></Error><Creatives><Creative><Linear><Duration>00:00:13</Duration><TrackingEvents></TrackingEvents><VideoClicks><ClickThrough><![CDATA[https://play.google.com/store/apps/details?id=com.zhiliaoapp.musically]]></ClickThrough><ClickTracking><![CDATA[https://rtbtrack-asia.rtblab.net/clk?p=339&d=65&t=55bW10qLPxlzQAXbMZUcWsf-OhubisbBjtjwjqdktKIJr2dmjrh-OmixU60az2HDosUCkB_GdRKL0zk60bPraXxkcuFOontU8dzS2Paqo_YKQa3DNJAh50yUvRWiacOSfOxGfUDUVhYtGQ2Ab7dIdxVZmCvBgmyZKNbyBjvQqdl0QsOEF0wals32i2chByx2fJicH9S5WC8D09hUuJLxkQ96S1pteutWmGp8Csv6C62OOzBToqZBk3nqBrJuIZTfbUoUrEYH7kQ8SeL6ltrcNHIUcp8sS6OdXHkirwusfPbJ1ExBXYw-D-RoojdsZcwIeNx7SKhPHeOqCr3FoS1ziFmtMvkvODPVJGTVxVmHSh6Yjy8b4fPZ8cmvgjRH4hsQ-d8Y3c3rtSIrYdgoptniln2pCXNfMxtrJnWtd1BIeI_IogEQVgztsHhn1O19KhKV7KxmEh0Nsf2H0kSTDI-yaYQtHGegEkJHITBnedGLwTsLaE3ri_pl296s4wZNNs7UKkOTtNOd1QiNoLs5foAl_bpXb3ESxOfsyWFlAN_dSinbpgVWyJmS_TektfONub_N]]></ClickTracking></VideoClicks><MediaFiles><MediaFile height=\\\"1280\\\" width=\\\"720\\\" type=\\\"video/mp4\\\" delivery=\\\"progressive\\\"><![CDATA[https://storage.googleapis.com/taurusx/adx-minidsp/creatives/com.zhiliaoapp.musically_720x1280.mp4]]></MediaFile></MediaFiles></Linear></Creative><Creative><CompanionAds><Companion width=\\\"720\\\" height=\\\"1280\\\"><CompanionClickThrough><![CDATA[https://storage.googleapis.com/taurusx/adx-minidsp/creatives/com.zhiliaoapp.musically_endcard_720x1280.jpg]]></CompanionClickThrough><TrackingEvents></TrackingEvents><CompanionClickTracking><![CDATA[https://rtbtrack-asia.rtblab.net/clk?p=339&d=65&t=55bW10qLPxlzQAXbMZUcWsf-OhubisbBjtjwjqdktKIJr2dmjrh-OmixU60az2HDosUCkB_GdRKL0zk60bPraXxkcuFOontU8dzS2Paqo_YKQa3DNJAh50yUvRWiacOSfOxGfUDUVhYtGQ2Ab7dIdxVZmCvBgmyZKNbyBjvQqdl0QsOEF0wals32i2chByx2fJicH9S5WC8D09hUuJLxkQ96S1pteutWmGp8Csv6C62OOzBToqZBk3nqBrJuIZTfbUoUrEYH7kQ8SeL6ltrcNHIUcp8sS6OdXHkirwusfPbJ1ExBXYw-D-RoojdsZcwIeNx7SKhPHeOqCr3FoS1ziFmtMvkvODPVJGTVxVmHSh6Yjy8b4fPZ8cmvgjRH4hsQ-d8Y3c3rtSIrYdgoptniln2pCXNfMxtrJnWtd1BIeI_IogEQVgztsHhn1O19KhKV7KxmEh0Nsf2H0kSTDI-yaYQtHGegEkJHITBnedGLwTsLaE3ri_pl296s4wZNNs7UKkOTtNOd1QiNoLs5foAl_bpXb3ESxOfsyWFlAN_dSinbpgVWyJmS_TektfONub_N&ec=1]]></CompanionClickTracking></Companion></CompanionAds></Creative></Creatives></InLine></Ad></VAST>","bundle":"com.zhiliaoapp.musically","iurl":"https://static-conversant-assist.rtblab.cc/R4zEsf0zjcfD1eC.jpg","cid":"345135","adid":"345135_1337026","crid":"1337026","cat":["IAB18"],"adomain":["musically.zhiliaoapp.com"],"attr":[4],"ext":{"crtype":"VAST 2.0"}}]}],"ext":{"source":""}}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_mytarget_vidio_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"cur":"RUB","seatbid":[{"bid":[{"price":1.24,"id":"1267238136","adid":"22247072","impid":"57f06495-325b-45b9-9466-e49c 1a557246","adm":"{\"imptrackers\":[\"http://rs. mail.ru/pixel/CrACCKDtzQoQmLCXzQQdz2jNVyAsKAFqswEBAOFoNvNjIOhs E1cMYuludYG87vNgQ2XqxviMe27qWmXVDdZEqnfEin9_XndBaD_HMfYBmwAWKU vswo7FbhugB0qPSd_x4ocwHuKVULdrZ53BZ1FA9VlaUkUeHgzzcE1I8A8C6NkN a9KYFSXtp7fjUzC_bK_Z-NrGR0CbxB9N2jbkkbkR9fhzZOGnv7FDdvm7HtfuQw wxvweb1F9u47rdgzMuroaVtRnOhgJJiS5c3UiYAYgBA5IBCC9rO0a2PwEAqgEL CAgQHhiB-ZMDIAiwAY6TPbgBRtoBCKM3Y-fOnSZBqgIQRaj0XC2lQSmTPx7oEZ 2unsICIec79QcHywHGAQ_EAYMHOgSBB6IEzAQBAQYD9wRMAwEBAQ.gif \"],\"link\":{\"url\":\"https://play.google.com/store/apps/det ails?id=com.mapswithme.maps.pro\",\"clicktrackers\":[\"http:// rs.mail.ru/pixel/CtYCCIXpoQoQ1s_h4AUdYTlZVyAhKAFqiwEBANFJIt_D1 0J9MHwp5-FiroeQUJsr_XR2YrhCYkg9vwmUK45R5N8LPtYpbq2rS7PwoPlovAx Ec-xd3qdwp5OvmsN9VT49swuM7L1QAdh9RPV9av5H3yl_Fd0uoppCn729btzRf HQ7bwQ883Cfbevdkg_tVPVIM9bsyWv304XDm7kiUtEg_Qt1bogBiAEDqgEGCAg QgKMFuAFGwAEB0gEQL5UrqdveShCIyKN-S6lA7-oBWAC7lnQn4Yvt19709C4Ye k7Ro8hqQo0cM5iYWu--rfaNHmJhKtrq32jR5RS4K31WVeDq32jR5RS5YQILfaN HlFLgIvW-0aPKKXEi1G32jR5RS4oV6bXg13iYAuDcAcICKc05AxINAfcBxgYCw QQBBA6BCQTbBA4BkwIEoQTaA3MCBgPDBQMBAQEoyAIBGhUB_jSKyNwnp_qzsSk MSdoVuKyzQFw.gif\"]},\"assets\":[{\"title\":{\"text\":\"[0+] M APS.ME — Оффлайн карты\"},\"id\":1},{\"img\":{\"url\":\"http s://r.mradx.net/img/7E/4B910B.png\",\"h\":256,\"w\":256},\"id \":2},{\"img\":{\"url\":\"http://r.mradx.net/img/C3/184E30.jpg\",\"h\":607,\"w\":1080},\"id\":3},{\"data\":{\"value\":\"Рекл ама\"},\"id\":4},{\"data\":{\"value\":\"Хорошая детализация в России. Отмечены даже пешие тропинки. Качайте бесплатно! \"},\"id\":5},{\"data\":{\"value\":\"4.5\"},\"id\":6},{\"data \":{\"value\":\"575251\"},\"id\":7},{\"data\":{\"value\":\"Уст ановить\"},\"id\":9},{\"video\":{\"vasttag\":\"<VAST xmlns:xsi =\\\"https://www.w3.org/2001/XMLSchema-instance\\\" version= \\\"2.0\\\" xsi:noNamespaceSchemaLocation=\\\"vast.xsd\\\"><Ad ><InLine><AdSystem>Advertising Technology Mail.ru Group</AdSys tem><Error><![CDATA[https://rs.mail.ru/pixel/CowBCMS07wpACMABC sgB1A8Q6-HgRh3lic1XIBooAqABoIq65w24AUb6ARBmY2YyZTgzNjFiMzFhYzF iuAKz_vyW6QOQA6CKuudN8gJBCCAQBBgBIAMoAjAlOAFAA1ABWAFgEGgBcAR4A eADAboEH2h0dHBzOi8vdmsuY29tL3ZpZGVvcy01MTE4OTcwNj8.gif]]></Err or><Impression id=\\\"mytarget\\\"><![CDATA[https://rs.mail.r u/pixel/CowBCMS07wpACMABCsgB0Q8Q6-HgRh3lic1XIBooAqABoIq65w24AU b6ARBmY2YyZTgzNjFiMzFhYzFiuAKz_vyW6QOQA6CKuudN8gJBCCAQBBgBIAMo AjAlOAFAA1ABWAFgEGgBcAR4AeADAboEH2h0dHBzOi8vdmsuY29tL3ZpZGVvcy 01MTE4OTcwNj8.gif]]></Impression><Creatives><Creative id=\\\"6 754883\\\" AdID=\\\"22796868\\\"><Linear><AdParameters /><Dura tion>00:00:15</Duration><MediaFiles><MediaFile id=\\\"1\\\" de livery=\\\"progressive\\\" type=\\\"video/mp4\\\" bitrate= \\\"651\\\" width=\\\"640\\\" height=\\\"360\\\" codec=\\\"h26 4\\\"><![CDATA[https://r.mradx.net/img/BD/F3CF81.mp4]]></Media File></MediaFiles><TrackingEvents><Tracking event=\\\"start \\\"><![CDATA[https://rs.mail.ru/pixel/CoUBCMS07wpACBDr4eBGHeW JzVcgGigCoAGgirrnDbgBRvoBEGZjZjJlODM2MWIzMWFjMWK4ArP-_JbpA5ADo Iq6503yAkEIIBAEGAEgAygCMCU4AUADUAFYAWAQaAFwBHgB4AMBugQfaHR0cHM 6Ly92ay5jb20vdmlkZW9zLTUxMTg5NzA2Pw.gif]]></Tracking><Tracking event=\\\"firstQuartile\\\"><![CDATA[https://rs.mail.ru/pixel/ CowBCMS07wpACMABCsgB6QcQ6-HgRh3lic1XIBooAqABoIq65w24AUb6ARBmY2 YyZTgzNjFiMzFhYzFiuAKz_vyW6QOQA6CKuudN8gJBCCAQBBgBIAMoAjAlOAFA A1ABWAFgEGgBcAR4AeADAboEH2h0dHBzOi8vdmsuY29tL3ZpZGVvcy01MTE4OT cwNj8.gif]]></Tracking><Tracking event=\\\"midpoint\\\"><![CDA TA[https://rs.mail.ru/pixel/CowBCMS07wpACMABCsgB6gcQ6-HgRh3lic 1XIBooAqABoIq65w24AUb6ARBmY2YyZTgzNjFiMzFhYzFiuAKz_vyW6QOQA6CK uudN8gJBCCAQBBgBIAMoAjAlOAFAA1ABWAFgEGgBcAR4AeADAboEH2h0dHBzOi 8vdmsuY29tL3ZpZGVvcy01MTE4OTcwNj8.gif]]></Tracking><Tracking e vent=\\\"thirdQuartile\\\"><![CDATA[https://rs.mail.ru/pixel/CowBCMS07wpACMABCsgB6wcQ6-HgRh3lic1XIBooAqABoIq65w24AUb6ARBmY2Y yZTgzNjFiMzFhYzFiuAKz_vyW6QOQA6CKuudN8gJBCCAQBBgBIAMoAjAlOAFAA 1ABWAFgEGgBcAR4AeADAboEH2h0dHBzOi8vdmsuY29tL3ZpZGVvcy01MTE4OTc wNj8.gif]]></Tracking><Tracking event=\\\"complete\\\"><![CDAT A[https://rs.mail.ru/pixel/CowBCMS07wpACMABCsgB7AcQ6-HgRh3lic1 XIBooAqABoIq65w24AUb6ARBmY2YyZTgzNjFiMzFhYzFiuAKz_vyW6QOQA6CKu udN8gJBCCAQBBgBIAMoAjAlOAFAA1ABWAFgEGgBcAR4AeADAboEH2h0dHBzOi8 vdmsuY29tL3ZpZGVvcy01MTE4OTcwNj8.gif]]></Tracking><Tracking ev ent=\\\"mute\\\"><![CDATA[https://rs.mail.ru/pixel/CowBCMS07wp ACMABCsgB7QcQ6-HgRh3lic1XIBooAqABoIq65w24AUb6ARBmY2YyZTgzNjFiM zFhYzFiuAKz_vyW6QOQA6CKuudN8gJBCCAQBBgBIAMoAjAlOAFAA1ABWAFgEGg BcAR4AeADAboEH2h0dHBzOi8vdmsuY29tL3ZpZGVvcy01MTE4OTcwNj8.gif]] ></Tracking><Tracking event=\\\"unmute\\\"><![CDATA[https://r s.mail.ru/pixel/CowBCMS07wpACMABCsgB7gcQ6-HgRh3lic1XIBooAqABoI q65w24AUb6ARBmY2YyZTgzNjFiMzFhYzFiuAKz_vyW6QOQA6CKuudN8gJBCCAQ BBgBIAMoAjAlOAFAA1ABWAFgEGgBcAR4AeADAboEH2h0dHBzOi8vdmsuY29tL3 ZpZGVvcy01MTE4OTcwNj8.gif]]></Tracking><Tracking event=\\\"clo se\\\"><![CDATA[https://rs.mail.ru/pixel/CowBCMS07wpACMABCsgB9 AcQ6-HgRh3lic1XIBooAqABoIq65w24AUb6ARBmY2YyZTgzNjFiMzFhYzFiuAK z_vyW6QOQA6CKuudN8gJBCCAQBBgBIAMoAjAlOAFAA1ABWAFgEGgBcAR4AeADA boEH2h0dHBzOi8vdmsuY29tL3ZpZGVvcy01MTE4OTcwNj8.gif]]></Trackin g></TrackingEvents><VideoClicks><ClickTracking id=\\\"mytarget \\\"><![CDATA[https://rs.mail.ru/pixel/CowBCMS07wpACMABCsgB8Ac Q6-HgRh3lic1XIBooAqABoIq65w24AUb6ARBmY2YyZTgzNjFiMzFhYzFiuAKz_ vyW6QOQA6CKuudN8gJBCCAQBBgBIAMoAjAlOAFAA1ABWAFgEGgBcAR4AeADAbo EH2h0dHBzOi8vdmsuY29tL3ZpZGVvcy01MTE4OTcwNj8.gif]]></ClickTrac king></VideoClicks></Linear></Creative></Creatives></InLine></ Ad></VAST>\"},\"id\":10}],\"ext\":{\"agerestrictions\":\"0+ \"}}"}]}],"id":"e88b8b18-299b-45b2-9041-5ad0bd1977c4"}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	// yandex
	r.HandleFunc("/mock/868/get_yandex_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82","seatbid":[{"bid":[{"price":0.0105,"id":"6145348938983074026","cid":"1_172609473","impid":"1","nurl":"https://an.yandex.ru/metacount/1GZy3E-D0Tq100000000U9mpVFNYnGUVLHwXbbUZ_98PWC0J9X9QcVbCFTOzcOqXbH4edldgXSHDa7WfY5URhyx0qCe88cSoVc9c08E5Z45Grah6BYO8Qo-ZWpmRmr4m_oeZIEjTHWO90QQvJ23sKYOT8BMMwHUGVPRfFn0yPRfPl8Fv5PF0xy85UQzUC32ORks8z7hDol2NYHEeg9KPP4ClCyZqgSmW8-Ooou8iP6PoK6bW501Qo9ocDPNnAbZkBfXyodpml2VNJsO79twmarQmV6BVmF8V1XXt8B4F8B5liXY3mMlxXondpu306_ktB22F_69XCJ8rCcbYbKpEpzK3ujpY_DAYv5IzfDJ4qfBAWfpIufEipuqvmTPzPGMPzekLqz3SoCp0qiJSk87juFpbGHri3ImBs7PlUlYcwyMdF-iPLxB1B7d4nfjHvcILLXoOw-bkbcB-aWqIbX_iVy9P4zbPYnyPkagzjT_0tj30iOETyGIxo0FshxVMJUs-VEvDsW_szMm0rN-jRm00?ssp-cur-price={AUCTION_PRICE}&ssp-cur=USD","crid":"1_172609473_300x250","iurl":"https://bs-metadsp.yandex.ru/resource/spacer.gif","adomain":["promotion-course.ru"],"w":"300","adid":"yabs.NzIwNTc2MDU5NjQ5MTk3MjA=","h":"250","cat":["IAB5","IAB3-1"],"adm":"<iframe frameborder=\\\"no\\\" scrolling=\\\"no\\\" marginwidth=\\\"0\\\" marginheight=\\\"0\\\" src=\\\"https://yandexadexchange.net/rtb_cache?rtb-cache-id=6145348938983074026&scr-height=250&h=250&scr-width=300&w=300&enabled-ad-type=text&ty=text&domain-hash=172609473&dh=172609473&yt=1&ssp-id=62419929&uniformat=true&layout-config=%7B%22width%22%3A300%2C%22height%22%3A250%7D&use-server-side-rendering=true&click_macro=&click-macro=\\\" style=\\\"height:250px;width:300px;\\\"></iframe>"}]}],"cur":"USD"}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_yandex_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82","seatbid":[{"bid":[{"price":0.0105,"id":"6144776674602970422","cid":"1_179376143","impid":"1","nurl":"https://an.yandex.ru/metacount/1LNENN-30T4100000000U9mpD7jxmV7AGhTGIxyWlsG68F24YOJs_ixLYdLFPcD8PGIAPvgzNQ8E4Yy5yMhprHaOMXb1v4p6zymCOB2GCGZAMidOnGHXxMNCKcLZE0hcNsK4QRtBo20Xm3IN2SG-o_GV29uotIpUmVoAIQ0qkEAHsLi6XciTMTSzNxCo_6NY12fgPGQPsiiCCd4gCqZRN6PPa2NCJ0TK6XW5W9Om-qzUpd0LB7SNJ3vbvXVOoFMJsS697-oC5IplsJUm_8S1nXr8x0D8x9iinc3msluXozap8F36_YqBoF6_M5YCJ8rCcbWba_DpzO0ujpY_j2WvLIyfDR6qfB8WfxJu9EjpOqwmTLzPGQRzOkKqDBSqDZSpD35kOEjVS2IqWws1PG7RxaqFttIz-FHdFUDADfWbZzXuWroShIhoCutANCnIp5-oGSBlW_rFE6iYEsln8qDNgRVs6zWR6ZXsi3EUO1TvmFvkwySNdxiwTgJ-O5_F1W0HHQVc?ssp-cur-price={AUCTION_PRICE}&ssp-cur=USD&call-rtb-count=1","crid":"1_179376143_0x0","iurl":"https://bs-metadsp.yandex.ru/resource/spacer.gif","adomain":["mcruises.ru"],"adid":"yabs.NzIwNTc2MDY5MjA2NDMyNjg=","cat":["IAB20"],"adm":"{\"native\":{\"imptrackers\":[\"https://an.yandex.ru/count/WK4ejI_zO9u0lGa0T10Vb8leBM2ppWK0dW4nJb9EOm00000uZeqpY081y0BFvi6H0_050Q06tfq1gGSsxQGdrA56LVW70T080gWAsUVX7NLJUE8sxG00reXgrqdhau6OW8G8g0-ssiEEnBgeerMG4FkGvU3LiQIxjW6e4S24FV0LhP-4jPoMuTure1QmfQgQ1iaMy3-15z0O0lWOeOQ3x9g8oCBS0O0PmOhsxAEFlFnZYHddCma00000090P0RWP____0T8P4dbXOdDVSsLoTcLoBt8tCpKjCUWPWC83y1c0mWE270rIH4CwU6zfCdP4DcGtwHm0y3-07Vz_-1w07zmJW220W0Y8807G8V__0TKY__z__mu0HREAcprmMGeQZEt2fOLaaiJ0HZGZdHULBfGn5sYqSrHtu2WbedgKUYCD0me0~1=WLSejI_zO5u0RGe0T1C-tAsVNW6oxe_VqzNAtya1W041Y07olVl8cW6G0OYBzVtVW8200fW1Y8lr_L-u0VoWiyybs06Um8Yh0U01pAxd0kW1hWFu0QI_thu1e0AUpRmAi0C2w0JQKeW5yS81a0NC-W6m1SOnk0NRFy05ffKEo0M1pGFG1OOMg0RUdG6f1pRjf2VKeKPLk0U01U07XeA01k08pwQU1UW91u0A0VWAWBKOw0m3o130i9220Q4HdvcPcPcPoOWI0P0I0T0KpBln5UWKZ0AW5h2bgfe6oHRmFvWNojnyk1S1m1UrrW6W6TwT0RWP____0O4Q__zB3SmEJcAe7W6m7m787u-Wn5Mf87eJbMqRSS4_a2BEhYYG8jAkA90YrAuea2BLhYYG8l78A90YyiWea2BurYYG8ldMADKY__z__mK0BDTeHH9ZtC5SIEB6aA1Wcu03~1\"],\"link\":{\"url\":\"https://an.yandex.ru/count/WSOejI_zODq0vGq0f1aVb8lelc8ONmK0tG4nJb9EOm00000uZesoxe_VqzNAtya1W041Y07olVl8cW6G0OYBzVtVW8200fW1Y8lr_L-u0VoWiyybm042s06Um8Yh0U01pAxd0kW1hWFu0QIUpRmAy0BFvi6H0x031EW4sbBu1F720OW5yS81a0NC-W6W1Rexg0N6CR05nZ6u1Ti_m0McbGx81O7D0z05XXRe1G6W1jwT0QW6tfq1gGSsxQGdrA56LRW7W0NW1uOAWe06u0ZFffu5w0a7W0e1g0hPv-4Tw0m3u0s2W9E1c8242CWGmB2GWW6X4P-PcPcPcSc84W6G4W7G5CoxyHNe58m2e1QmfQgQ1iaMy3-O5yhSVBWN0S0NjTO1e1dUdG6u6V___m616l__ImtC3avYi1gxik7iYkZo_Ay1WXmDKaH3EdXlQJ9sH3PaDwWU0VWU-v3buDMnfBks0O0Vt1Em7m787u-Wn5Mf87eJbMqRSS4_a2BEhYYG8jAkA90YrAuea2BLhYYG8l78A90YyiWea2BurYYG8ldMADKY__z__mW0qjTeaHDZ56OkINd7h4LdShtDSuuYg4EN0um28-81~1\"},\"assets\":[{\"title\":{\"text\":\"Новогодние круизы по Мальдивам!Скидки,…\"},\"id\":3},{\"data\":{\"value\":\"Круизные пакеты с авиаперелетом. Новогодние заезды и каникулы! Бронируй сейчас!\"},\"id\":4},{\"data\":{\"value\":\"Go\"},\"id\":5},{\"img\":{\"url\":\"https://avatars.mds.yandex.net/get-direct/4872497/vwsZ7c9bhz2mFhl_-8naMg/y450\",\"h\":450,\"w\":600},\"id\":1},{\"img\":{\"url\":\"https://avatars.mds.yandex.net/get-direct/4872497/vwsZ7c9bhz2mFhl_-8naMg/y450\",\"h\":450,\"w\":600},\"id\":2}]}}"}]}],"cur":"USD"}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_yandex_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"id":"32f94228-52d3-4bce-be88-5dd1d799ae46","seatbid":[{"bid":[{"price":0.0105,"id":"6145412198853052875","cid":"1_174992724","impid":"1","nurl":"https://an.yandex.ru/metacount/1RSoejUK0U8100000000U9mpFCtvzzZLnnwXbiUZ_P8PWC0J9X8wvxPGTAnxCXj3AYDGFFCqf-nQa7WfY5URhyx0qCe88cSowh0mWC52nXPc26ibOmSJXBMNqK6UZU4ec7-L4QJrBYE3X03JN2OGUocJ0v1Qo_G0I9vb-Wy4Jvbk5kzWViKai7Xv9VjfJaO6xrUlqhrUipByPU84QcXb1jdGommoVIepo8YvpB8W2vcP75GQ60M05Z9p2Yrdk0gMkmicdpA7wEPn_MHsSE87Esa5ongE3Ip_OO1n1nBx01Bx9ikicI3mqluXozap8F36_YqBo2E_M5YCp0oCcbWba_DpzO0ujpY_j2WvLIyfDR6qfB8WfxJu9EjpOqwmTHTO6oOmVMK56VcBbTFGtCJIqiZSoAI1xSgj7S4TR0qi2zXsRthufkl5fp_h6LUomMnwnCPROkPgPHXQQHRbRfPY_f8D4g0Vx7_2MHFPMOiV6RfAlRNVmDxGmB63dV44kyW3zkbtp_mu-lpNYeGVxElP031ChR00?ssp-cur-price={AUCTION_PRICE}&ssp-cur=USD","crid":"1_174992724_320x480","iurl":"https://bs-metadsp.yandex.ru/resource/spacer.gif","adomain":["play.google.com"],"w":"320","adid":"yabs.NzIwNTc2MDYyODQ1OTAwMjU=","h":"480","cat":["IAB19"],"adm":"<iframe frameborder=\\\"no\\\" scrolling=\\\"no\\\" marginwidth=\\\"0\\\" marginheight=\\\"0\\\" src=\\\"https://yandexadexchange.net/rtb_cache?rtb-cache-id=6145412198853052875&scr-height=480&h=480&scr-width=320&w=320&enabled-ad-type=text&ty=text&domain-hash=174992724&dh=174992724&yt=1&ssp-id=62419929&click_macro=\\\" style=\\\"height:480px;width:320px;\\\"></iframe>"}]}],"cur":"USD"}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_yandex_native_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
					"id":"e657d364-1164-49bf-b8cc-184e04737288",
					"bidid":"c6ca008e9088c951f3ad52e5dc4960f0",
					"seatbid":[
						{
							"bid":[
								{
									"id":"a6fe364e662e032e7d91e01eeb83fdbb",
									"impid":"1",
									"price":0.11284,
									"nurl":"http://b1.dcntr-ads.com/?win=nurl&sp=${AUCTION_PRICE}&t=native&uniq=511763b40a8546344164ad9895147f71",
									"adm":"{\"native\":{\"imptrackers\":[\"https://an.yandex.ru/count/WQmejI_zO1G15Gq0v1Lf6VkPFqaTN0K0508nmBIkOm00000uZeqpY081y0AapjUW1l050Q067ia6S_7Y0knRwcsf1xiCmSToUa5i20Ae2jdduHrnFr52kmm101inaeq8yPE1c804yWqaD68sC6GrEJKjCpTbDoqqPc9bBJbbOZKjPZKoCJSnOsCmPM4pg0-xcUIHhSxUe6oG48gmeOE8XSYy8gWHm8GzwH8mt15T863eXF0I__y1-182a1JpuekNmFsMzMtm5Psyuw6Unus1pm6W5f3oovm6oHRmFu4Ng1S4q1W2-1YLpyFJsz_AfLc06S6AzkoZZxpyOuaPqGYG6G6u6V__0T8P4dbXOdDVSsLoTcLoBt8tDp0jDEWPWC83y1c0mWEO6lEaAi8Qi1ink1i2WXmDCKf5Ed1bScnmKNHbD-aSW1t_V_WUW1_L3e0WW808Y201q27___y1rIB__t__WIC00000003mFnG0GuUQsKBmo9YH4ICczkc1vOU-DKWBoLWcXfyj7o0NCrCAbVUXDGFwTu20k8OTPIop9EqS9HuuYege7hZlIDc00W00~1=WP4ejI_zOAa0FGm0L1RjPiqRgG7SsiU6uTVPWvm1W041Y069pykBdW6G0OIunfhYW8200fW1XBZ6ccAu0R2Qg_abs06Wa-oe0U01pDgxe07e0T01-07Edzw-0PW2fgc55w02_hG2i0Fv98W5ZDKLa0Mdtq-m1OF11RW5YgCKm0MWoyi2o0NKfbJG1Twv0gW67ga7kmp1nt9wGMou1u05u0U62j08keZ0WS22W0RW2B-Q0UW91u0A0VWAWBKOw0oN0iWGmB0GeUaICDmHNI1Ww8I0582G0j0KvVYd4kWKZ0AW5f3oovm6oHRG5XUGjFWjk1S1m1UrrW6W6Hwu6V___m616l__Onl4TY_Fg1u1o1u3i1y1o1_GYRPKcY0FSdKkOcLoTIvXRcHoRsbagI2bapBdxgxEFxWWhUce0v0Yr9uga2BMdYgm8W7L8l__V_-18uaZUnIkH-5wt3-G8uoqgSkeufhHdW6O8wl-riZrXeQdYW490DGV9AhT4NncD9qGX0uCmq8eYSyoJp7KgPJECTWcYW00~1\"],\"link\":{\"url\":\"https://an.yandex.ru/count/Wa8ejI_zO9O1nH40921f6VkP7MICOGK0bW8nmBIkOm00000uZetSsiU6uTVPWvm1W041Y069pykBdW6G0OIunfhYW8200fW1XBZ6ccAu0R2Qg_abm0600TW1e9Fig07W0SpQkw01w07G0VW1pf-O0gQfXHUW0lwq0l02fCxNeFaa-0ICrHM81OpL5P05fzzFe0M4pWQe1OF11R05Wy45k0MAenJ01Q3BomB81TIcLD05tha2S_7Y0knRwcsf1xiCmSToUa5ik0U01U07XWhG2Bg8m870We06u0Y_cW7e2GU02W6e2jdduHte39S2u0s2W9E1c804yWqaD68sC6GrEJKjCpTbDoqqPc9bBJbbOZKjPZKoCJSnOsCmPM4pmB0GeUaICDmHNI1Ww8I0582G0f0Ky-ABby3zblLjq1Jb-ASIw1IC0g0MaFBBd0R95l0_q1REdzw-0PWNaBJuBQWN1BWN0S0NjTO1e1aUk1d___y1WHh__sCRn7Olpx0QkxBXx8heylol0S8QWXmDCKf5Ed1bScnmKNHbDwWU0SWU0_WUYh2XWuY5oBmYW1_L3h0V0SWVq8csL9eW3t9rBc9bSdKkOMvaSczfPAaWfPCov-wkpZ-u8Atfg0EG8jIUAf0Yrfugi281rIB__t__WIE98tiKhaVXUjm_a2ECjAdBgEAQqPu1c2Eh_jR8zOQ6fue13W0brOY8edbkA6Au2XMZA0t3N2a98teAGgPa2QTtOYYiyuCL9QuVR7mhN4IQm0OO1KR40m00~1\"},\"assets\":[{\"title\":{\"text\":\"Товары ИКЕА уже в продаже на Яндекс Маркете\"},\"id\":1},{\"data\":{\"value\":\"[16+] Оригиналы от бренда ИКЕА доступны в приложении Яндекс Маркет\"},\"id\":2},{\"data\":{\"value\":\"Скачать\"},\"id\":3},{\"img\":{\"url\":\"https://avatars.mds.yandex.net/get-direct/5249943/OzLwWUXiiRDTwQcHG-PtGg/y450\",\"h\":450,\"w\":253},\"id\":5},{\"img\":{\"url\":\"https://avatars.mds.yandex.net/get-google-play-app-icon/40549/f4eb3947f9ad0dc310444343a3a47ed5/icon-xld-retina\",\"h\":300,\"w\":300},\"id\":4},{\"data\":{\"value\":\"4.2\"},\"id\":\"7\"},{\"video\":{\"vasttag\":\"<?xml version=\\\"1.0\\\" encoding=\\\"utf-8\\\"?><VAST version=\\\"3.0\\\"><Ad id=\\\"15001\\\"><InLine><AdSystem>Yandex</AdSystem><Error><![CDATA[https://advideo.dsp.yandexadexchange.net/errorlog/spacer.gif?bidid=c66c2578-0d53-46b2-bef9-3f3b969d201e&vl=Y29udj9jcmlkPTFfMF8weDAmdz0wJmg9MCZpbWd1cmw9&ssp-id=62419929&adid=yabs.NzIwNTc2MDcyMTkxMDUyODI%3D&code=[ERRORCODE]]]></Error><AdTitle><![CDATA[Товары ИКЕА уже в продаже на Яндекс Маркете]]></AdTitle><Description><![CDATA[[16+] Оригиналы от бренда ИКЕА доступны в приложении Яндекс Маркет]]></Description><Extensions><Extension type=\\\"yandex_ad_info\\\"><![CDATA[{\\\"assets\\\":{},\\\"text\\\":{}}]]></Extension></Extensions><Survey>https://www.yandex.ru</Survey><Creatives><Creative id=\\\"Linear\\\"><Linear><Duration>00:00:05</Duration><TrackingEvents><Tracking event=\\\"start\\\"><![CDATA[https://an.yandex.ru/count/WQmejI_zO1G15Gq0v1Lf6VkPFqaTN0K0508nmBIkOm00000uZeqpY081y0AapjUW1l050Q067ia6S_7Y0knRwcsf1xiCmSToUa5i20Ae2jdduHrnFr52kmm101inaeq8yPE1c804yWqaD68sC6GrEJKjCpTbDoqqPc9bBJbbOZKjPZKoCJSnOsCmPM4pg0-xcUIHhSxUe6oG48gmeOE8XSYy8gWHm8GzwH8mt15T863eXF0I__y1-182a1JpuekNmFsMzMtm5Psyuw6Unus1pm6W5f3oovm6oHRmFu4Ng1S4q1W2-1YLpyFJsz_AfLc06S6AzkoZZxpyOuaPqGYG6G6u6V__0T8P4dbXOdDVSsLoTcLoBt8tDp0jDEWPWC83y1c0mWEO6lEaAi8Qi1ink1i2WXmDCKf5Ed1bScnmKNHbD-aSW1t_V_WUW1_L3e0WW808Y201q27___y1rIB__t__WIC00000003mFnG0GuUQsKBmo9YH4ICczkc1vOU-DKWBoLWcXfyj7o0NCrCAbVUXDGFwTu20k8OTPIop9EqS9HuuYege7hZlIDc00W00~1=WP4ejI_zOAa0FGm0L1RjPiqRgG7SsiU6uTVPWvm1W041Y069pykBdW6G0OIunfhYW8200fW1XBZ6ccAu0R2Qg_abs06Wa-oe0U01pDgxe07e0T01-07Edzw-0PW2fgc55w02_hG2i0Fv98W5ZDKLa0Mdtq-m1OF11RW5YgCKm0MWoyi2o0NKfbJG1Twv0gW67ga7kmp1nt9wGMou1u05u0U62j08keZ0WS22W0RW2B-Q0UW91u0A0VWAWBKOw0oN0iWGmB0GeUaICDmHNI1Ww8I0582G0j0KvVYd4kWKZ0AW5f3oovm6oHRG5XUGjFWjk1S1m1UrrW6W6Hwu6V___m616l__Onl4TY_Fg1u1o1u3i1y1o1_GYRPKcY0FSdKkOcLoTIvXRcHoRsbagI2bapBdxgxEFxWWhUce0v0Yr9uga2BMdYgm8W7L8l__V_-18uaZUnIkH-5wt3-G8uoqgSkeufhHdW6O8wl-riZrXeQdYW490DGV9AhT4NncD9qGX0uCmq8eYSyoJp7KgPJECTWcYW00~1]]></Tracking></TrackingEvents><MediaFiles></MediaFiles><VideoClicks><ClickThrough id=\\\"AW\\\"><![CDATA[https://an.yandex.ru/count/Wa8ejI_zO9O1nH40921f6VkP7MICOGK0bW8nmBIkOm00000uZetSsiU6uTVPWvm1W041Y069pykBdW6G0OIunfhYW8200fW1XBZ6ccAu0R2Qg_abm0600TW1e9Fig07W0SpQkw01w07G0VW1pf-O0gQfXHUW0lwq0l02fCxNeFaa-0ICrHM81OpL5P05fzzFe0M4pWQe1OF11R05Wy45k0MAenJ01Q3BomB81TIcLD05tha2S_7Y0knRwcsf1xiCmSToUa5ik0U01U07XWhG2Bg8m870We06u0Y_cW7e2GU02W6e2jdduHte39S2u0s2W9E1c804yWqaD68sC6GrEJKjCpTbDoqqPc9bBJbbOZKjPZKoCJSnOsCmPM4pmB0GeUaICDmHNI1Ww8I0582G0f0Ky-ABby3zblLjq1Jb-ASIw1IC0g0MaFBBd0R95l0_q1REdzw-0PWNaBJuBQWN1BWN0S0NjTO1e1aUk1d___y1WHh__sCRn7Olpx0QkxBXx8heylol0S8QWXmDCKf5Ed1bScnmKNHbDwWU0SWU0_WUYh2XWuY5oBmYW1_L3h0V0SWVq8csL9eW3t9rBc9bSdKkOMvaSczfPAaWfPCov-wkpZ-u8Atfg0EG8jIUAf0Yrfugi281rIB__t__WIE98tiKhaVXUjm_a2ECjAdBgEAQqPu1c2Eh_jR8zOQ6fue13W0brOY8edbkA6Au2XMZA0t3N2a98teAGgPa2QTtOYYiyuCL9QuVR7mhN4IQm0OO1KR40m00~1]]></ClickThrough></VideoClicks><Icons><Icon program=\\\"icon\\\" height=\\\"80\\\" width=\\\"80\\\" xPosition=\\\"left\\\" yPosition=\\\"bottom\\\"><StaticResource creativeType=\\\"image/png\\\"><![CDATA[https://yastatic.net/pcode/static/icons/icon-globe-192x192.png]]></StaticResource></Icon><Icon program=\\\"AdChoices\\\" height=\\\"15\\\" width=\\\"15\\\" xPosition=\\\"right\\\" yPosition=\\\"bottom\\\"><StaticResource creativeType=\\\"image/png\\\"><![CDATA[https://yastatic.net/pcode/static/icons/icon-globe-64x64.png]]></StaticResource><IconClicks><IconClickThrough><![CDATA[https://yandex.ru/tune/adv]]></IconClickThrough></IconClicks></Icon></Icons></Linear></Creative></Creatives></InLine></Ad></VAST>\"},\"id\":10}]}}",
									"adomain":[
										"badoo.com"
									],
									"cat":[
										"IAB14-1"
									],
									"adid":"55dbdb1c9686",
									"iurl":"http://b1.dcntr-ads.com/?t=preview&k=55dbdb1c9686",
									"cid":"49_885_161950",
									"crid":"49_885_304461"
								}
							],
							"seat":"49"
						}
					],
					"cur":"USD"
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_yandex_native1_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"e657d364-1164-49bf-b8cc-184e04737288",
						"bidid":"c6ca008e9088c951f3ad52e5dc4960f0",
						"seatbid":[
							{
								"bid":[
									{
										"id":"a6fe364e662e032e7d91e01eeb83fdbb",
										"impid":"1",
										"price":0.11284,
										"nurl":"http://b1.dcntr-ads.com/?win=nurl&sp=${AUCTION_PRICE}&t=native&uniq=511763b40a8546344164ad9895147f71",
										"adm":"{\"native\":{\"imptrackers\":[\"https://an.yandex.ru/count/WRGejI_zO381LGq0f1Owpc__C30McmK0CW8nQsDSP000000uZeqpY081y0BcywQA3BRyqJPF9fCygGTP3K65w0Fm8lW70gWAsUVX7P60DY5P3G40V803l6Bxau6OW8G8yWqaOpCuCMPXPJCjCsHZE2qqPJ5cBJXYDp4jDJCqDMKrEJ8tC6Hag0_PclIAYVs0-2AG4CxLy871jhtP_G6e4S24FUaIb_h-mO2hfShm4Wg84m6G4pAG5BRvnxRpoSd9FF0LdyhlgEdy-ytR0PWMaFEeeWQWoHQ15wWN1D0O0lWO-hV5nRQpu_Z80O0PmOhsxAEFlFnZYHbJbGi00000090P0RWP_m7I6H9vOM9pNtDbSdPbSYzoE34vBJ7e6O320_0PWC83c1hKmrF26h0RLBWR0u8S3LfCHJepD6PFI70tQJVf780T_t_u7e0W0eWW3D0X____0TKY__z__u4Ze2C1g2C1i2C1k2C1yYE8906e98O1i2IP3nC0lcighbxomVo4Di4nOqXF33ZOfyc-Y6wH9hH8MsN86DBQOK9vPZp-UCs2XhkZpZ2A0QrbMC-C4JRkNMWo~1=WU0ejI_zO0a1jGu0b1eyog5_2GBAhDMltCZAegm1W06Ql_-8bexqpJw80T2XsAUc0P01oipJt-U0W802c07ApDFVPxW1gl33noRO0P36qh41u07Qy9cL0SwVthu1e0BmW8SRi0Fo88W5uU0Oa0NHwpMm1S-b1hW5ZjWDm0Nbv4B81RRz4D05uiG2g0Qg0wa7MGr1XUW3y2Au1u05q0YwYC21m8201-08kC333UW91u0A0VWAWBKOw0oN0fWDnj4umR0Gc16gpIwXkUaIb_h-mO2hfSg05820W81Wq1G4w1IC0fWMaFEeeWQWoHRG5iwVthu1c1UtaCKhk1S1m1UrrW6W6Qe3k1d_0O4Q__-Nmv7_7_6e7W787a2m7m787yk9s6QQ821ZRsqkUM5kP6LuBcrlOcbiPIvmR7LpBcTXRMKkOsbqUQaWLCHv7olNwZ-u8DEKBf0Ybyqka2AQpIwG8gFDBf0YfCqka2AgpIxL8l__V_-18uaZUnIkH-5wt3-G8vpEyj7DlxgVMPWZrg6BfxAnWOBt0V8Z4960DY5P3G40V803l6Bx1G2u9E41mYG1CSea0p8oDGm0BiT8V7l9O5gaGMNGAJ5CXLpOR9vCSYnGaZEButOD6-sUK1uDV5Z-Lk40~1\"],\"link\":{\"url\":\"https://an.yandex.ru/count/Wf8ejI_zOFq1HHG0P2Gwpc__UbQMG0K0_G8nQsDSP000000uZeqpogpLhzp8ogAi0O01ch__Y9QEzCq-Y07GeTYdfW6G0ShCqz_dW8200fW1oipJtsUu0QhmmyScm0600TW1aCRIiG7W0ThmcPK10Q02y8276_02vlEcYl8W-0JXu1Y81U7W6905qUire0NVnWUe1S-b1h05pwK6k0MEs0t01UNaGiW5jlqGq0NYn0As_D4sJoQJFAa7MGr1XUW3y2Au1u05q0YwYC21m8201-08kC333UW91u0A0QWAsUVX7UWCbmAO3SRHEE0DWe2JWPY0X0Zo3IHZCpWnPc5bCoqpP6CuBJHbCMOjE68tCIqrCpGrPJKvCZSmP6IG4CxLy871jhtP_S6m49WHgiqkeRdf4fVw_i60gwNAW1I0W820O90Kjld7jlF9oSayq1G4w1IC0fWMaFEeeWQWoHRG5iwVthu1c1UtaCKhg1S4k1S1m1UrrW6W6Qe3k1d_0O4Q__-Nmv7_7_6m6hkouUoAwFByhm726e8S3LfCHJepD6PFI70tQJUe7W787a3u7h0V0SWVoudOPfeW86DlRIvvOMvaPNWkRMzYQMnbBd1iTNCkPs5jPIvZQNHvgI1Kn7aVAzVgFxWWqvGka2ANpIwG8fhDBf0Yeyqka2AapIwG8ghDBjKY__z__u4ZYIDx5Av7uNhSFv0ZdCxoqSs_kfzPc2FMeOkdih61WlS1yYCGaO0s8LaD0G1yW0EyOli50AWaXW6m99aFk2JX0S8a0J7A90CoCZKI09pRI6Hksr3NF2mU8TbyaJQLeJR5MBda8eJqP5SqrEbsQSTODWM9GnedDw8Wt93lZDx_5QPp8E6ZCKm_Vwju~1\"},\"assets\":[{\"title\":{\"text\":\"[0+] Играй в Плюс Сити и получай баллы Плюса\"},\"id\":1},{\"data\":{\"value\":\"Строй свой виртуальный город, играй в мини-игры и получай реальные баллы Плюса!\"},\"id\":4},{\"data\":{\"value\":\"Реклама\"},\"id\":8},{\"data\":{\"value\":\"Открыть\"},\"id\":5},{\"img\":{\"url\":\"https://avatars.mds.yandex.net/get-direct/3950780/V5HGi9dLpbFDfvil9LeldQ/y450\",\"h\":450,\"w\":450},\"id\":2},{\"img\":{\"url\":\"https://avatars.mds.yandex.net/get-google-play-app-icon/3936215/48e8c8600169e8cf2ac5025fea3e677b/icon-xld-retina\",\"h\":300,\"w\":300},\"id\":3},{\"data\":{\"value\":\"4.1\"},\"id\":\"6\"}]}}",
										"adomain":[
											"badoo.com"
										],
										"cat":[
											"IAB14-1"
										],
										"adid":"55dbdb1c9686",
										"iurl":"http://b1.dcntr-ads.com/?t=preview&k=55dbdb1c9686",
										"cid":"49_885_161950",
										"crid":"49_885_304461"
									}
								],
								"seat":"49"
							}
						],
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//flatads
	r.HandleFunc("/mock/868/get_flatads_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
			"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
			"bidid":"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364",
			"cur":"USD",
			"seatbid":[
				{
					"seat":"adx",
					"group":0,
					"bid":[
						{
							"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366",
							"impid":"1",
							"price":0.09000000000000001,
							"adm":"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<VAST version=\"3.0\"><Ad id=\"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101365\"><InLine><AdSystem version=\"1.0\">flat-ads</AdSystem><AdTitle><![CDATA[Moto Bike Attack Race]]></AdTitle><Impression><![CDATA[https://api.flat-ads.com/api/tracker/tracking/impression?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=imp.23951.676725849872.526032182772.1]]></Impression><Creatives><Creative adID=\"111594\" id=\"111594\"><Linear><Duration>00:00:26</Duration><MediaFiles><MediaFile delivery=\"progressive\" type=\"video/mp4\" width=\"720\" height=\"1280\"><![CDATA[http://dsp-adcreative.mobshark.net/adshark_dsp/1666339183940_zd.mp4]]></MediaFile></MediaFiles><TrackingEvents><Tracking event=\"start\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=start]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=firstQuartile]]></Tracking><Tracking event=\"midpoint\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=midpoint]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=thirdQuartile]]></Tracking><Tracking event=\"complete\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=complete]]></Tracking></TrackingEvents><VideoClicks><ClickThrough><![CDATA[https://game.86753.com/vid/1013/MotoBikeAttackRace/uid_12307201/etr_vid_ad/index.html?browserType=2&entry=sdk_ads]]></ClickThrough><ClickTracking><![CDATA[https://api.flat-ads.com/api/tracker/tracking/click?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=click.88712.271095740792.121392083592.1&af_siteid=166_155354_2292]]></ClickTracking></VideoClicks></Linear><CreativeExtensions></CreativeExtensions></Creative></Creatives><Description><![CDATA[Choose Your Favourite Moto and Weapon!]]></Description><Extensions><Extension><ImageWidth>320</ImageWidth><ImageHeight>480</ImageHeight><ImageUrl>http://dsp-adcreative.mobshark.net/adshark_dsp/1660294465587.jpg?x-oss-process=style/hq</ImageUrl></Extension></Extensions></InLine></Ad></VAST>",
							"nurl":"https://api.flat-ads.com/api/tracker/tracking/win_notice?log_source=win&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=imp.23951.676725849872.526032182772.1",
							"crid":"111594",
							"ext":null,
							"cid":"2292",
							"language":"",
							"attr":null,
							"cat":null
						}
					]
				}
			]
		}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_flatads_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
				"bidid":"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667290360.21.51882279",
				"cur":"USD",
				"seatbid":[
					{
					"seat":"adx",
					"group":0,
					"bid":[
						{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667290360.21.51882280",
						"impid":"1",
						"price":0.09000000000000001,
						"adm":"<html><head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"></head><body><a id=\"adx_banner_1\" style=\"display: block; width: 100%;  margin: 0 auto; position: relative;padding:0px;\" > <img style=\"width: 100%;display:block;\" id=\"adx_banner_1_img\" title=\"adx banner img\" alt=\"adx banner img\" onload=\"var adxScript = document.getElementById('adx_banner_1_script'); if (!adxScript || !adxScript.innerHTML) {return;} var scriptTxt = adxScript.innerHTML; eval(scriptTxt);\" src=\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAIAAAACAQMAAABIeJ9nAAAAAXNSR0IB2cksfwAAAAlwSFlzAAALEwAACxMBAJqcGAAAAANQTFRF////p8QbyAAAAAxJREFUeJxjYGBgAAAABAAB9hc4VQAAAABJRU5ErkJggg==\" /></a><script id=\"adx_banner_1_script\"> (function () { function init() { if (window['adx_banner_1_done'] === '1') { return; } else { window['adx_banner_1_done'] = '1'; } var bannerDom = document.getElementById( 'adx_banner_1'); var bannerObj={\"title\":\"Moto Bike Attack Race\",\"desc\":\"Choose Your Favourite Moto and Weapon!\",\"link\":\"https://game.86753.com/vid/1013/MotoBikeAttackRace/uid_12307201/etr_vid_ad/index.html?browserType=2&entry=sdk_ads\",\"fileurls\":[\"http://dsp-adcreative.mobshark.net/adshark_dsp/1659585098405.jpg?x-oss-process=style/hq\"],\"w\":320,\"h\":50,\"clicktrackers\":[\"https://api.flat-ads.com/api/tracker/tracking/click?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667290360.21.51882279&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667290360.21.51882280&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=97798&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155353&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_97798_155353_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=click.79801.180006092995.127903326795.1&af_siteid=166_155353_2292\"],\"imptrackers\":[\"https://api.flat-ads.com/api/tracker/tracking/impression?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667290360.21.51882279&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667290360.21.51882280&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=97798&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155353&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_97798_155353_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=imp.88424.48737602355.42438335635.1\"],\"action\":\"browser\",\"eventtrackers\":null}; var bannerImg = document.getElementById( 'adx_banner_1_img'); try { var imgSrc = bannerObj.fileurls[0]; var imptrackers = bannerObj.imptrackers; var isSafari = (/Safari/.test(navigator.userAgent) && !/Chrome/.test(navigator.userAgent)); if (isSafari && imgSrc.indexOf('?x-oss-process=style/hq') !== -1) {imgSrc = imgSrc.replace('?x-oss-process=style/hq', '?x-oss-process=image/format,png');} bannerImg.src = imgSrc; createReportImg({ srcList: imptrackers, type: 'tracker' }); initEvent(bannerObj); } catch (error) { console.error(error); bannerImg.src = ''; bannerImg.setAttribute('alt', 'error: render data fail'); } } function initEvent(bannerObj) { var bannerDom = document.getElementById( 'adx_banner_1'); var clicktrackers = bannerObj.clicktrackers; var goUrl = bannerObj.link; var appIcon = bannerObj.app_icon; var action = bannerObj.action; var title = bannerObj.title; if (bannerDom) { bannerDom.onclick = function () { goUrlFn({ goUrl: goUrl, appIcon: appIcon, action: action, title: title, }); createReportImg({ srcList: clicktrackers, type: 'click' }); }; } } function createReportImg(opt) { var srcList = opt.srcList; var type = opt.type; for (var i = 0; i < srcList.length; i++) { new Image().src = srcList[i]; } } function goUrlFn(opt) { var u=navigator.userAgent;var isAndroid=u.indexOf('Android')>-1||u.indexOf('Adr')>-1;var isiOS=(/Safari/.test(u) && !/Chrome/.test(u)); if (opt.goUrl) { if ( opt.action == 'apk' && window.vbrowser && window.vbrowser.addTask) { window.vbrowser.addTask( opt.goUrl, opt.title || '', opt.appIcon || '', 0, 'apk', 'jssdk'); } else if(opt.action == 'market'){if(isAndroid){window.location.href=opt.goUrl;}else if(isiOS){window.location.href=opt.goUrl; }else{window.location.href=opt.goUrl; }} else { window.location.href = opt.goUrl; } } } init(); })(); </script></body></html>",
						"nurl":"https://api.flat-ads.com/api/tracker/tracking/win_notice?log_source=win&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667290360.21.51882279&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667290360.21.51882280&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=97798&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155353&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_97798_155353_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=imp.88424.48737602355.42438335635.1",
						"crid":"97798",
						"ext":null,
						"cid":"2292",
						"language":"",
						"attr":null,
						"cat":null
						}
					]
					}
				]
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_flatads_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
				"bidid":"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424609",
				"cur":"USD",
				"seatbid":[
					{
					"seat":"adx",
					"group":0,
					"bid":[
						{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424610",
						"impid":"1",
						"price":0.09000000000000001,
						"adm":"{\"native\":{\"assets\":[{\"id\":1,\"required\":1,\"img\":{\"type\":3,\"url\":\"http://dsp-adcreative.mobshark.net/adshark_dsp/1620899840630.jpg?x-oss-process=style/hq\",\"w\":1200,\"h\":627}}],\"link\":{\"url\":\"http://well.456halogames.com/subway-surfers/index.html?entry=sdk_ads\",\"clicktrackers\":[\"https://api.flat-ads.com/api/tracker/tracking/click?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424609&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424610&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=88876&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155356&dspid=100&srcid=3815373&cou=IDN&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&ag=6&ac=108&clickid=v1_2292_12606_88876_155356_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_108&et=click.47363.872080077674.927287201574.1&af_siteid=166_155356_2292\"],\"deeplink\":\"\"},\"imptrackers\":[\"https://api.flat-ads.com/api/tracker/tracking/impression?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424609&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424610&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=88876&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155356&dspid=100&srcid=3815373&cou=IDN&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&ag=6&ac=108&clickid=v1_2292_12606_88876_155356_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_108&et=imp.03621.059429549382.104726872282.1\"],\"eventtrackers\":null}}",
						"nurl":"https://api.flat-ads.com/api/tracker/tracking/win_notice?log_source=win&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424609&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424610&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=88876&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155356&dspid=100&srcid=3815373&cou=IDN&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&ag=6&ac=108&clickid=v1_2292_12606_88876_155356_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_108&et=imp.03621.059429549382.104726872282.1",
						"crid":"88876",
						"ext":null,
						"cid":"2292",
						"language":"",
						"attr":null,
						"cat":null
						}
					]
					}
				]
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//opera
	r.HandleFunc("/mock/868/get_opera_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "id": "418e3726-0184-1000-cfce-fe30abf60088",
				  "seatbid": [
					{
					  "bid": [
						{
						  "id": "0004074a5d1dcb016701",
						  "impid": "1",
						  "price": 1.6940549286200808,
						  "nurl": "https://t-odx.op-mobile.opera.com/win?a=a4377737133120&adomain=shopee.com&b=com.cyg.novafly.gp&cc=MY&cid=bBUU0HCkLUKCREwcy33F&cm=1&crid=vvfGCVU3KFe5fZ7P52t0&ct=&dbbf=1.000&dbmb=0.000&dpl=1&dvt=PHONE&ep=ep6871964013312&ext=u3wITA6z5bVQjAI-ZQAJ-1i7mRP4cGVeujjQlu60lvCNwjhYpeU01WJCC3yxrkJjDULTrY1uCcRirMG3m7LSccjfLgHbP1eXsVOul1Npl-MAnZ46b7uQD6Ns1vJzj_Fkwr4U8UT84evbkFCwJ4M6crH6qejlszvKv7HrPxSgaia-26R1tW7vhfk2oKPgMiiQ&gi=2c1c2290-b2a9-4dc7-b1d6-c81c57c9d4c2&impr_dl=1667557950234&it=1&m=m4377737133312&msz=360x640&om=1&ot=ANDROID&pc=1&pubId=pub6871903319744&radv=rtbhouse&rip=183.171.20.5&s=s6871968295872&said=30ecee4c84ae475ba0c97dc16c03b7fd&sc=2&schp=1-2&se=0004074a5d1dcb016701&spid=fee43d609cce4da5827d086f331a3590&srid=418e3726-0184-1000-cfce-fe30abf60088&ss=2&u=7a38888bf39d48ba&va=a4377737133120&vm=m4377737133312&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}",
						  "burl": "https://t-odx.op-mobile.opera.com/billing?a=a4377737133120&adomain=shopee.com&b=com.cyg.novafly.gp&bl=1&cc=MY&cid=bBUU0HCkLUKCREwcy33F&cm=1&cnu=H4sIAAAAAAAC_wTA2ZJzSgAA4LeZOzOYyHKqpk4hdqETiQ43Sro1bWuxhqf_v2Icu-G_n5-Btt-oz9KRzhnC7Tdizc9CW65lIyXr_2P1By1H2n_KpdwDftylymLR4tzxpp4PYPe0WYwH59d5L41vdfh4Ey6mJa872rvFS_NIUq-dxEsmOHlVvI_cJpw2SUU86ItFimCVbltWywNsSxkCWPbOouUvZhNHdIYmzC73bVGXITjxL-kS1IK7S-wgf5fbXkHcvs-ehkreLEwB0132uKZmc185MU9QPT0eWV3EaARRVEDAZr6fBM4oovZ58FnoHzPzWh34TzIflBqI8QeSCLDu7L2dlc6yYnC9KpF722s3WFaCS_0gvm3rc7pMGAoMYQPX_sfyABtPJN0akpNysIknCjJ-3YzDBZmJKLQrjdJ-xEc_tn_vVyWSt5km79-XyvH5NTbmU2ZpMjsBpYetKIqMd01ac-Q82JciTWrgvRpDsiN937S-0kLpqBn4CR5ZGhpldm_czLGKimA8-Kl4CgWOnWN_DgyPUSJoj269p2degahopq5jqaCtjVFjQy0DSPJsK9ctDz56pU_E33QjqPtpC_i-lG8qwrdQWDnV1xHHRKQfudj2PgOEnMwfUvlr6f7Eb0E4SPvj14jR30DbfwEAAP___2GgW1QCAAA&crid=vvfGCVU3KFe5fZ7P52t0&ct=&dbbf=1.000&dbmb=0.000&dpl=1&dvt=PHONE&ep=ep6871964013312&ext=u3wITA6z5bVQjAI-ZQAJ-1i7mRP4cGVeujjQlu60lvCNwjhYpeU01WJCC3yxrkJjDULTrY1uCcRirMG3m7LSccjfLgHbP1eXsVOul1Npl-MAnZ46b7uQD6Ns1vJzj_Fkwr4U8UT84evbkFCwJ4M6crH6qejlszvKv7HrPxSgaia-26R1tW7vhfk2oKPgMiiQ&gi=2c1c2290-b2a9-4dc7-b1d6-c81c57c9d4c2&impr_dl=1667557950234&it=1&iv=1&m=m4377737133312&msz=360x640&om=1&ot=ANDROID&pc=1&pubId=pub6871903319744&radv=rtbhouse&rip=183.171.20.5&s=s6871968295872&said=30ecee4c84ae475ba0c97dc16c03b7fd&sc=2&schp=1-2&se=0004074a5d1dcb016701&spid=fee43d609cce4da5827d086f331a3590&srid=418e3726-0184-1000-cfce-fe30abf60088&ss=2&u=7a38888bf39d48ba&va=a4377737133120&vm=m4377737133312&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}",
						  "lurl": "https://t-odx.op-mobile.opera.com/loss?a=a4377737133120&adomain=shopee.com&b=com.cyg.novafly.gp&cc=MY&cid=bBUU0HCkLUKCREwcy33F&cm=1&crid=vvfGCVU3KFe5fZ7P52t0&ct=&dbbf=1.000&dbmb=0.000&dpl=1&dvt=PHONE&ep=ep6871964013312&ext=u3wITA6z5bVQjAI-ZQAJ-1i7mRP4cGVeujjQlu60lvCNwjhYpeU01WJCC3yxrkJjDULTrY1uCcRirMG3m7LSccjfLgHbP1eXsVOul1Npl-MAnZ46b7uQD6Ns1vJzj_Fkwr4U8UT84evbkFCwJ4M6crH6qejlszvKv7HrPxSgaia-26R1tW7vhfk2oKPgMiiQ&gi=2c1c2290-b2a9-4dc7-b1d6-c81c57c9d4c2&impr_dl=1667557950234&it=1&m=m4377737133312&msz=360x640&om=1&ot=ANDROID&pc=1&pubId=pub6871903319744&radv=rtbhouse&rip=183.171.20.5&s=s6871968295872&said=30ecee4c84ae475ba0c97dc16c03b7fd&sc=2&schp=1-2&se=0004074a5d1dcb016701&spid=fee43d609cce4da5827d086f331a3590&srid=418e3726-0184-1000-cfce-fe30abf60088&ss=2&u=7a38888bf39d48ba&va=a4377737133120&vm=m4377737133312&al=${AUCTION_LOSS}",
						  "adm": "<script type=\"text/javascript\">\n  function changeURLArg(url, arg, arg_val) {\n    var pattern = arg + '=([^&]*)';\n    var replaceText = arg + '=' + arg_val;\n    if (url.match(pattern)) {\n      var tmp = '/(' + arg + '=)([^&]*)/gi';\n      tmp = url.replace(eval(tmp), replaceText);\n      return tmp;\n    } else {\n      if (url.match('[\\?]')) {\n        return url + '&' + replaceText;\n      } else {\n        return url + '?' + replaceText;\n      }\n    }\n  }\n  function tracking(url) {\n    try {\n      navigator.sendBeacon(url);\n    } catch (error) {\n    }\n  }\n  function clickTracks(clkType) {\n    var urls = \"https://t-odx.op-mobile.opera.com/click?a=a4377737133120&adomain=shopee.com&b=com.cyg.novafly.gp&cc=MY&cid=bBUU0HCkLUKCREwcy33F&cm=1&crid=vvfGCVU3KFe5fZ7P52t0&ct=&dbbf=1.000&dbmb=0.000&dpl=1&dvt=PHONE&ep=ep6871964013312&ext=u3wITA6z5bVQjAI-ZQAJ-1i7mRP4cGVeujjQlu60lvCNwjhYpeU01WJCC3yxrkJjDULTrY1uCcRirMG3m7LSccjfLgHbP1eXsVOul1Npl-MAnZ46b7uQD6Ns1vJzj_Fkwr4U8UT84evbkFCwJ4M6crH6qejlszvKv7HrPxSgaia-26R1tW7vhfk2oKPgMiiQ&gi=2c1c2290-b2a9-4dc7-b1d6-c81c57c9d4c2&it=1&m=m4377737133312&msz=360x640&om=1&ot=ANDROID&pc=1&pubId=pub6871903319744&radv=rtbhouse&rip=183.171.20.5&s=s6871968295872&said=30ecee4c84ae475ba0c97dc16c03b7fd&sc=2&schp=1-2&se=0004074a5d1dcb016701&spid=fee43d609cce4da5827d086f331a3590&srid=418e3726-0184-1000-cfce-fe30abf60088&ss=2&u=7a38888bf39d48ba&va=a4377737133120&vm=m4377737133312&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}&ir=${IS_PREVIEW}\";\n    urls = urls.split(\",\");\n    for (var i = 0; i < urls.length; i++) {\n      var url = changeURLArg(urls[i], 'ct', clkType);\n      tracking(url);\n    }\n  }\n  var clickHandler = function() {\n    clickTracks(1)\n  }\n  var blurHandler = function() {\n    if (document.activeElement.tagName === \"IFRAME\") {\n      clickTracks(2)\n    }\n  }\n  var touchstartHandler = function() {\n    clickTracks(3)\n  }\n\n  document.addEventListener(\"DOMContentLoaded\", function () {\n    document.body.addEventListener('click', clickHandler);\n    document.body.addEventListener('touchstart', touchstartHandler);\n    window.addEventListener(\"blur\", blurHandler);\n  });\n\n  document.addEventListener(\"unload\", function () {\n    document.body.removeEventListener('click', clickHandler);\n    document.body.removeEventListener('touchstart', touchstartHandler);\n    window.removeEventListener(\"blur\", blurHandler);\n  });\n\n</script><head><style type='text/css'>body {margin:auto auto;text-align:center;} </style></head><div id=\"-oadx_0004074a5d1dcb016701m43777371333121667547150235122498\"><div><iframe src=\"https://sin.creativecdn.com/imp-delivery?tk=WIK56xjwj6P0t4aBwIihDp0HFgsP4XJoZdsK3KqwmOIpd8R1MHIAy4irLhbENf_lyp505HP9NkZ6YLmVuz5Cc0Prhw5YWkazzelAsWnjAWPWjrKwEgboJfK2KsmVeMTzwCwsS90b5MSl1L4_JSgqjz6Bc-6reXGCfqoVaPoFLoUQaHmTy-2g_cluUUelhZctPYYhWPov0ru1-GhYnX7OoVO8eHQk70x_v7BlP2ZxWfYPopDNqKyivABG-rC5fTnrERWjk1LiOSZRzyXuMudW1ocdGdlOxINPot9fazmfgfjsJfN21AdbRG7McH_21nyiYartd8OZJ3TQBYAzvi_q3bC-0gQZGv9eIEAo9PBrWn222o0LHil-fDsJMha_lPNbmG5JYF6mnOBnW58EGdXPUeaVGjeTmLeKIhkfddsOa29V1-oDZOvSGNoif1EUpyTaD0BWchmuppoa1EymGldGCjSWfgezjyzgSxFkFufOzFGSlruzS0rjARCcdRV1y-COFc-o2cF8-ZJNxsWW-A07aA&amp;curl=https%3A%2F%2Fsin.creativecdn.com%2Fclicks%3Fid%3D20221104_b2aZMgw007zXJAkYZj6S%26%7BEXTRA_CLICK_PARAMS%7D&amp;tdc=sin\" width=\"360\" height=\"640\" scrolling=\"no\" frameBorder=\"0\"  loading=\"eager\" style=\"margin: auto; display: block;\"></iframe></div><img id='adxImpressionTrackingPixel0' alt=\"\" src=\"https://t-odx.op-mobile.opera.com/impr?a=a4377737133120&adomain=shopee.com&b=com.cyg.novafly.gp&cc=MY&cid=bBUU0HCkLUKCREwcy33F&cm=1&crid=vvfGCVU3KFe5fZ7P52t0&ct=&dbbf=1.000&dbmb=0.000&dpl=1&dvt=PHONE&ep=ep6871964013312&ext=u3wITA6z5bVQjAI-ZQAJ-1i7mRP4cGVeujjQlu60lvCNwjhYpeU01WJCC3yxrkJjDULTrY1uCcRirMG3m7LSccjfLgHbP1eXsVOul1Npl-MAnZ46b7uQD6Ns1vJzj_Fkwr4U8UT84evbkFCwJ4M6crH6qejlszvKv7HrPxSgaia-26R1tW7vhfk2oKPgMiiQ&gi=2c1c2290-b2a9-4dc7-b1d6-c81c57c9d4c2&impr_dl=1667557950234&it=1&m=m4377737133312&msz=360x640&om=1&ot=ANDROID&pc=1&pubId=pub6871903319744&radv=rtbhouse&rip=183.171.20.5&s=s6871968295872&said=30ecee4c84ae475ba0c97dc16c03b7fd&sc=2&schp=1-2&se=0004074a5d1dcb016701&spid=fee43d609cce4da5827d086f331a3590&srid=418e3726-0184-1000-cfce-fe30abf60088&ss=2&u=7a38888bf39d48ba&va=a4377737133120&vm=m4377737133312&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}\" style='width:0px;height:0px;border:0px;margin:0px;float:left;'/><img id='adxImpressionTrackingPixel1' alt=\"\" src=\"https://creativecdn.com/cm-notify?pi=opera\" style='width:0px;height:0px;border:0px;margin:0px;float:left;'/></div>",
						  "adomain": [
							"shopee.com"
						  ],
						  "cid": "bBUU0HCkLUKCREwcy33F",
						  "crid": "vvfGCVU3KFe5fZ7P52t0",
						  "h": 640,
						  "w": 360,
						  "exp": 10800,
						  "ext": {
							"deeplink":"www.baidu.com",
							"fallback":"www.google.com"
							}
						}
					  ],
					  "seat": "adv3198545450624"
					}
				  ],
				  "bidid": "0004074a5d1dcb016701",
				  "cur": "USD"
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_opera_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "id": "cdic11ql7v5ro693jfig",
				  "seatbid": [
					{
					  "bid": [
						{
						  "id": "0000074a5d9409b7e9c1",
						  "impid": "1",
						  "price": 10,
						  "nurl": "https://t.adx.opera.com/win?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vm=m7005769600192&vp=5&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}",
						  "burl": "https://t.adx.opera.com/billing?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&bl=1&cbu=H4sIAAAAAAAC_0yPzU7DMBCE38a3prGd2AlShEDijnoCLpZjb5OlrdesnSLx9Kg9II7zzY80a625POz3eMkMpSClHWxMGZozHisdjw3SnjKw_xd59NEtTFt2GCfZWzMaEVafEpzvRGkRGHzFK9y0GrTsBhHhiuEGXFm9nKw0c98pE9QI8-hVHLUdYBzmrtNgdZTW9qNujSi0cQDnc3alEt8nA12apOs38ampUBmLmDG6WiZpjO07q6xUUgu_hYqUbpVX8D8fdH174nA4vLw_nz6_-M_PjAGmttGCGBdM0-qX5Qy8u19vVfsbAAD__zZAuC8oAQAA&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&iv=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vm=m7005769600192&vp=5&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}",
						  "lurl": "https://t.adx.opera.com/loss?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vm=m7005769600192&vp=5&al=${AUCTION_LOSS}",
						  "adm": "<VAST version=\"3.0\"><Ad id=\"47062a33-b046-43eb-8c11-ef52b91371fb\"><InLine><AdSystem><![CDATA[Liftoff]]></AdSystem><AdTitle><![CDATA[283148]]></AdTitle><Impression><![CDATA[https://impression-europe.liftoff.io/opera/beacon?ad_group_id=157696&channel_id=123&creative_id=283148&auction_id=PeazZovXArcRREYBkjqr&origin=haggler-opera020]]></Impression><Impression><![CDATA[https://t.adx.opera.com/impr?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vm=m7005769600192&vp=5]]></Impression><Impression><![CDATA[https://opr.adx.opera.com/i?clid=opr&paid=opr&dvid=v&avid=adv7005106418624&caid=o7005106419136&plid=m7005769600192-283148&siteId=app4007015262912&kv7=pub4007008646336&publisherId=70098&kv1=320x480&kv4=94.181.222.81&kv3=2577e4d7f21fd166&kv10=&kv12=s4139511258944&kv25=Tetris%C2%AE&kv15=RU&kv16=0.00000000&kv17=0.00000000&kv18=com.n3twork.tetris&kv19=83f32e3e-25d4-4910-aef4-110f0ed4ace3&kv28=hry-lx1&kv26=android&kv23=&kv11=0000074a5d9409b7e9c1&kv27=Mozilla%2F5.0+%28Linux%3B+Android+10%3B+HRY-LX1+Build%2FHONORHRY-L21%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F106.0.5249.126+Mobile+Safari%2F537.36&lineItemId=a7005769600000&kv29=[ERRORCODE]&kv30=[CONTENTPLAYHEAD]_[ADPLAYHEAD]&kv33=[ASSETURI]&kv34=[VASTVERSIONS]&kv35=[IFATYPE]&kv36=[IFA]&kv37=[CLIENTUA]&kv38=[SERVERUA]&kv39=[DEVICEUA]&kv40=[DEVICEIP]&kv41=[LATLONG]&kv42=[DOMAIN]&kv43=[PAGEURL]&kv44=&kv45=[PLAYERSIZE]&kv46=[REGULATIONS]&kv47=[ADTYPE]&kv48=[TRANSACTIONID]&kv49=[BREAKPOSITION]&kv50=[APPNAME]&kv51=[PLACEMENTTYPE]&kv54=[LAT]&kv5=OpenRTB&kv24=Mobile_InApp_Video]]></Impression><Creatives><Creative id=\"crid:a7005769600000:7af62c16004c392b1b50f34556a0cc6c\"><Linear><Duration>00:00:25</Duration><TrackingEvents><Tracking event=\"start\"><![CDATA[https://adexp.liftoff.io/event/vast/start/f4dcd1CIDQCRIZMDAwMDA3NGE1ZDk0MDliN2U5YzFfMDE5MhjYnsCMxDAgeyiMpBEwliE6EmNvbS5uM3R3b3JrLnRldHJpc0IeY3JzZXJ2ZS1pb3MtcG9kLWZvcmNlZC1jb250cm9sQiBjcnNlcnZlLXVtYnJlbGxhLXRlc3QtZXhwZXJpbWVudEoKYzc1ZTEwZTRiM1AGWgNSVVNgAmgEcgxldS1jZW50cmFsLTHgAQOAAXuSAQJydZgBAqEBAAAAAAAAsD-qAQg3MjB4MTI4MLIBBlB1enpsZboBCFRldHJpc8KuwgEZdmFzdC05MGU0MDFhN2UzM2FkNGRiYTRiNsoBAwEDAtIBBjYxMDAwMNoBBXZpZGVv?playhead=[CONTENTPLAYHEAD]&sr=1]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://adexp.liftoff.io/event/vast/firstQuartile/f4dcd1CIDQCRIZMDAwMDA3NGE1ZDk0MDliN2U5YzFfMDE5MhjYnsCMxDAgeyiMpBEwliE6EmNvbS5uM3R3b3JrLnRldHJpc0IeY3JzZXJ2ZS1pb3MtcG9kLWZvcmNlZC1jb250cm9sQiBjcnNlcnZlLXVtYnJlbGxhLXRlc3QtZXhwZXJpbWVudEoKYzc1ZTEwZTRiM1AGWgNSVVNgAmgEcgxldS1jZW50cmFsLTHgAQOAAXuSAQJydZgBAqEBAAAAAAAAsD-qAQg3MjB4MTI4MLIBBlB1enpsZboBCFRldHJpc8KuwgEZdmFzdC05MGU0MDFhN2UzM2FkNGRiYTRiNsoBAwEDAtIBBjYxMDAwMNoBBXZpZGVv?playhead=[CONTENTPLAYHEAD]&sr=1]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://adexp.liftoff.io/event/vast/midpoint/f4dcd1CIDQCRIZMDAwMDA3NGE1ZDk0MDliN2U5YzFfMDE5MhjYnsCMxDAgeyiMpBEwliE6EmNvbS5uM3R3b3JrLnRldHJpc0IeY3JzZXJ2ZS1pb3MtcG9kLWZvcmNlZC1jb250cm9sQiBjcnNlcnZlLXVtYnJlbGxhLXRlc3QtZXhwZXJpbWVudEoKYzc1ZTEwZTRiM1AGWgNSVVNgAmgEcgxldS1jZW50cmFsLTHgAQOAAXuSAQJydZgBAqEBAAAAAAAAsD-qAQg3MjB4MTI4MLIBBlB1enpsZboBCFRldHJpc8KuwgEZdmFzdC05MGU0MDFhN2UzM2FkNGRiYTRiNsoBAwEDAtIBBjYxMDAwMNoBBXZpZGVv?playhead=[CONTENTPLAYHEAD]&sr=1]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://adexp.liftoff.io/event/vast/thirdQuartile/f4dcd1CIDQCRIZMDAwMDA3NGE1ZDk0MDliN2U5YzFfMDE5MhjYnsCMxDAgeyiMpBEwliE6EmNvbS5uM3R3b3JrLnRldHJpc0IeY3JzZXJ2ZS1pb3MtcG9kLWZvcmNlZC1jb250cm9sQiBjcnNlcnZlLXVtYnJlbGxhLXRlc3QtZXhwZXJpbWVudEoKYzc1ZTEwZTRiM1AGWgNSVVNgAmgEcgxldS1jZW50cmFsLTHgAQOAAXuSAQJydZgBAqEBAAAAAAAAsD-qAQg3MjB4MTI4MLIBBlB1enpsZboBCFRldHJpc8KuwgEZdmFzdC05MGU0MDFhN2UzM2FkNGRiYTRiNsoBAwEDAtIBBjYxMDAwMNoBBXZpZGVv?playhead=[CONTENTPLAYHEAD]&sr=1]]></Tracking><Tracking event=\"complete\"><![CDATA[https://adexp.liftoff.io/event/vast/complete/f4dcd1CIDQCRIZMDAwMDA3NGE1ZDk0MDliN2U5YzFfMDE5MhjYnsCMxDAgeyiMpBEwliE6EmNvbS5uM3R3b3JrLnRldHJpc0IeY3JzZXJ2ZS1pb3MtcG9kLWZvcmNlZC1jb250cm9sQiBjcnNlcnZlLXVtYnJlbGxhLXRlc3QtZXhwZXJpbWVudEoKYzc1ZTEwZTRiM1AGWgNSVVNgAmgEcgxldS1jZW50cmFsLTHgAQOAAXuSAQJydZgBAqEBAAAAAAAAsD-qAQg3MjB4MTI4MLIBBlB1enpsZboBCFRldHJpc8KuwgEZdmFzdC05MGU0MDFhN2UzM2FkNGRiYTRiNsoBAwEDAtIBBjYxMDAwMNoBBXZpZGVv?playhead=[CONTENTPLAYHEAD]&sr=1]]></Tracking><Tracking event=\"creativeView\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=creativeview]]></Tracking><Tracking event=\"start\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=start]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=firstquartile]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=midpoint]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=thirdquartile]]></Tracking><Tracking event=\"complete\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=complete]]></Tracking><Tracking event=\"mute\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=mute]]></Tracking><Tracking event=\"unmute\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=unmute]]></Tracking><Tracking event=\"pause\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=pause]]></Tracking><Tracking event=\"resume\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=resume]]></Tracking><Tracking event=\"rewind\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=rewind]]></Tracking><Tracking event=\"fullscreen\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=fullscreen&isfullscreen=1]]></Tracking><Tracking event=\"exitFullscreen\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=fullscreen&isfullscreen=0]]></Tracking><Tracking event=\"acceptInvitationLinear\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=acceptinvitationlinear]]></Tracking><Tracking event=\"closeLinear\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=closelinear]]></Tracking><Tracking event=\"skip\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=skip]]></Tracking><Tracking event=\"progress\" offset=\"00:00:02\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=progress&offset=2s]]></Tracking><Tracking event=\"progress\" offset=\"00:00:15\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=progress&offset=15s]]></Tracking><Tracking event=\"progress\" offset=\"97%\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=progress&offset=97p]]></Tracking></TrackingEvents><VideoClicks><ClickThrough><![CDATA[https://play.google.com/store/apps/details?id=com.zm.watersort&referrer=adjust_external_click_id%3Dv.2_g.157696_a.0000074a5d9409b7e9c1_0192_c.123_t.ua_u.716b5426c29eb9a2%26utm_campaign%3Dtest-param1%26utm_content%3DTest%2BSource%2BApp%26utm_source%3Dliftoff%26utm_term%3DTest%2BCreative_320x50]]></ClickThrough><ClickTracking><![CDATA[https://click.liftoff.io/v1/campaign_click/0MO4F1fxskxmMDCTX3cZhodMBMiMW7JQgnKQFFCD30L5pumDT39MfANSG3gWdF7sngAQTZ9YLyBqmv4o4zmrqk1sv_yKFqzIMsOFyn4yRKJh6d-oSY62xqMgPzlrHNomcL3bsY9N4hC1q2I87kQRceb_r3svGj7sMMbdLycDkj6F1I1GFn_DRGmAXrhSw5iNksLg2eUsrlRJJtbUqgRMf0qOVPYMHTNk_oHzccXiGaxRxXhwKWjV0zJaNDgpWXg0WoHZiAn0Ri9HtPVVQFN-2ugNgI2IdNSHHMuN7zVmUhCsZ4qRKtm5qhBcS1qiaxeXpSAFgVkkFQx4oWjOG0YRRV84pFKW9jyLwX3rkY73KlVc4fZzXDrZFSvYln3uexe8alh28o99jhqM0LWttwAwEE458_WezUiUNhu92ZK2WxBd-z2jqT4vIQTgz5RW-wBRWbMYZcjTDz2juoHbqRBhwv5rtcDqfd-QdqdH8XONUNEwo87sBjxoWGDLZe4kak1K0wxtB0SIzdbgcfSfGafBLQS0xC3QkAXoUI6eAKZpT_UeLz1wGc8qPhSRSi4xd5YaPI587nkz1Pq0_ls?vast_el=1]]></ClickTracking><ClickTracking><![CDATA[https://t.adx.opera.com/click?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5]]></ClickTracking></VideoClicks><MediaFiles><MediaFile delivery=\"progressive\" type=\"video/mp4\" bitrate=\"1050\" width=\"360\" height=\"640\" scalable=\"true\" maintainAspectRatio=\"true\"><![CDATA[https://cdn.liftoff-creatives.io/customers/cba62ed85f/videos/mobile/3a71f7d2630de9d5a8a7.mp4]]></MediaFile></MediaFiles></Linear></Creative><Creative><CompanionAds><Companion width=\"320\" height=\"480\" assetWidth=\"1080\" assetHeight=\"1920\"><CompanionClickThrough><![CDATA[https://play.google.com/store/apps/details?id=com.zm.watersort&referrer=adjust_external_click_id%3Dv.2_g.157696_a.0000074a5d9409b7e9c1_0192_c.123_t.ua_u.716b5426c29eb9a2%26utm_campaign%3Dtest-param1%26utm_content%3DTest%2BSource%2BApp%26utm_source%3Dliftoff%26utm_term%3DTest%2BCreative_320x50]]></CompanionClickThrough><CompanionClickTracking><![CDATA[https://click.liftoff.io/v1/campaign_click/0MO4F1fxskxmMDCTX3cZhodMBMiMW7JQgnKQFFCD30L5pumDT39MfANSG3gWdF7sngAQTZ9YLyBqmv4o4zmrqk1sv_yKFqzIMsOFyn4yRKJh6d-oSY62xqMgPzlrHNomcL3bsY9N4hC1q2I87kQRceb_r3svGj7sMMbdLycDkj6F1I1GFn_DRGmAXrhSw5iNksLg2eUsrlRJJtbUqgRMf0qOVPYMHTNk_oHzccXiGaxRxXhwKWjV0zJaNDgpWXg0WoHZiAn0Ri9HtPVVQFN-2ugNgI2IdNSHHMuN7zVmUhCsZ4qRKtm5qhBcS1qiaxeXpSAFgVkkFQx4oWjOG0YRRV84pFKW9jyLwX3rkY73KlVc4fZzXDrZFSvYln3uexe8alh28o99jhqM0LWttwAwEE458_WezUiUNhu92ZK2WxBd-z2jqT4vIQTgz5RW-wBRWbMYZcjTDz2juoHbqRBhwv5rtcDqfd-QdqdH8XONUNEwo87sBjxoWGDLZe4kak1K0wxtB0SIzdbgcfSfGafBLQS0xC3QkAXoUI6eAKZpT_UeLz1wGc8qPhSRSi4xd5YaPI587nkz1Pq0_ls?vast_el=2]]></CompanionClickTracking><CompanionClickTracking><![CDATA[https://t.adx.opera.com/click?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=2&vm=m7005769600192&vp=5]]></CompanionClickTracking><TrackingEvents><Tracking event=\"creativeView\"><![CDATA[https://adexp.liftoff.io/event/vast/staticEndscreenView/f4dcd1CIDQCRIZMDAwMDA3NGE1ZDk0MDliN2U5YzFfMDE5MhjYnsCMxDAgeyiMpBEwliE6EmNvbS5uM3R3b3JrLnRldHJpc0IeY3JzZXJ2ZS1pb3MtcG9kLWZvcmNlZC1jb250cm9sQiBjcnNlcnZlLXVtYnJlbGxhLXRlc3QtZXhwZXJpbWVudEoKYzc1ZTEwZTRiM1AGWgNSVVNgAmgEcgxldS1jZW50cmFsLTHgAQOAAXuSAQJydZgBAqEBAAAAAAAAsD-qAQg3MjB4MTI4MLIBBlB1enpsZboBCFRldHJpc8KuwgEZdmFzdC05MGU0MDFhN2UzM2FkNGRiYTRiNsoBAwEDAtIBBjYxMDAwMNoBBXZpZGVv?playhead=[CONTENTPLAYHEAD]&sr=1]]></Tracking><Tracking event=\"creativeView\"><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=2&vm=m7005769600192&vp=5&evt=companionView]]></Tracking></TrackingEvents><StaticResource creativeType=\"image/jpeg\"><![CDATA[https://cdn.liftoff-creatives.io/customers/cba62ed85f/image/lambda_jpg_89/cefef889ab09e4adc03f.jpg]]></StaticResource></Companion></CompanionAds></Creative></Creatives><Description></Description><Survey></Survey><Error><![CDATA[https://t.adx.opera.com/video?a=a7005769600000&adomain=zephyrmobile.com&b=com.n3twork.tetris&cc=RU&cid=157696&cm=1&crid=283148&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep4007249252672&ext=T5CgZGNSzrYaWm5mp3Zh84Uz9x8peYWFb-H0vZGzjLwOWQOC7u4HD7bsU5HJQ--N1sIPpZ1A1MDtYslw1nwMmOj3U1YQ_IGUZ2JOMYPo74bviArUBr71RY9C9wDUuFyOIclVXE5kvNrPP2UUot3kfQ%3D%3D&gi=83f32e3e-25d4-4910-aef4-110f0ed4ace3&iabCat=IAB9-5&impr_dl=1667557271259&it=1&m=m7005769600192&msz=320x480&ot=ANDROID&pubId=pub4007008646336&rip=94.181.222.81&rw=1&s=s4139511258944&said=137918&sc=3&schp=1-3&se=0000074a5d9409b7e9c1&spid=70098&srid=cdic11ql7v5ro693jfig&ss=2&u=2577e4d7f21fd166&va=a7005769600000&vd=30&vel=1&vm=m7005769600192&vp=5&evt=error]]></Error></InLine></Ad></VAST>",
						  "adomain": [
							"zephyrmobile.com"
						  ],
						  "bundle": "com.zm.watersort",
						  "iurl": "https://cdn.liftoff.io/customers/1145/previews/844858dab9-ios-56e8d31488.png",
						  "cid": "157696",
						  "crid": "283148",
						  "cat": [
							"IAB9-5"
						  ],
						  "attr": [
							6
						  ],
						  "h": 480,
						  "w": 320,
						  "exp": 10000,
						  "ext": {}
						}
					  ],
					  "seat": "adv7005106418624"
					}
				  ],
				  "bidid": "0000074a5d9409b7e9c1",
				  "cur": "USD"
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_opera_native1.2_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
    "id":"d8b356ae-5c40-4d77-90b5-1401b2751254",
    "seatbid":[
        {
            "bid":[
                {
                    "id":"0004074a5d6c86804c80",
                    "impid":"0",
                    "price":0.2454296,
                    "nurl":"https://t-odx.op-mobile.opera.com/win?a=a7005771013696&adomain=snap.com&b=com.mi.android.globalminusscreen&cc=IN&cid=160529&cm=1&crid=139087&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep5865235689408&ext=gpJ1cA2ilf-HT1DWiq0QvE6nd5yi_KW490CgkIYdYJN3WsS349HZQUturdfrOlM6euSgiZkNqMbYBOx0YNTkoaf9JPkSMV_2xEG5XS1uMsJuLw1QgljD3l2Sk4NC5WFy24ivHa0dyUT1H7TbOiKiPi4lkYlDTMlyB5OwCzpjg7A3D&gi=e3d4066d-c42b-4afd-bc8d-1b3b061269ba&iabCat=IAB1&impr_dl=1667554430843&m=m7005771014016&msz=1200x627&ot=ANDROID&pubId=pub5865193350528&rip=157.38.134.34&s=s5865212277120&said=337&sc=2&schp=1-2&se=0004074a5d6c86804c80&spid=637&srid=d8b356ae-5c40-4d77-90b5-1401b2751254&ss=2&u=1af1ac59ea3aff7b&va=a7005771013696&vm=m7005771014016&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}",
                    "lurl":"https://t-odx.op-mobile.opera.com/loss?a=a7005771013696&adomain=snap.com&b=com.mi.android.globalminusscreen&cc=IN&cid=160529&cm=1&crid=139087&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep5865235689408&ext=gpJ1cA2ilf-HT1DWiq0QvE6nd5yi_KW490CgkIYdYJN3WsS349HZQUturdfrOlM6euSgiZkNqMbYBOx0YNTkoaf9JPkSMV_2xEG5XS1uMsJuLw1QgljD3l2Sk4NC5WFy24ivHa0dyUT1H7TbOiKiPi4lkYlDTMlyB5OwCzpjg7A3D&gi=e3d4066d-c42b-4afd-bc8d-1b3b061269ba&iabCat=IAB1&impr_dl=1667554430843&m=m7005771014016&msz=1200x627&ot=ANDROID&pubId=pub5865193350528&rip=157.38.134.34&s=s5865212277120&said=337&sc=2&schp=1-2&se=0004074a5d6c86804c80&spid=637&srid=d8b356ae-5c40-4d77-90b5-1401b2751254&ss=2&u=1af1ac59ea3aff7b&va=a7005771013696&vm=m7005771014016&al=${AUCTION_LOSS}",
                    "adm":"{\"native\":{\"ver\":\"1.2\",\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://gdl.bytegle.tech/as/web-source/4hc/1jFJyG.png\",\"w\":640,\"h\":640}},{\"id\":2,\"img\":{\"url\":\"https://gdl.bytegle.tech/as/web-source/4hb/1OTPkQ.png\",\"w\":128,\"h\":128}},{\"id\":3,\"title\":{\"text\":\"Easy Paisa \"}},{\"id\":4,\"data\":{\"value\":\"Enjoy upto 40% discount\\u0026 Rs 150 cashback\"}},{\"id\":5,\"data\":{\"value\":\"Download\"}}],\"link\":{\"url\":\"https://play.google.com/store/apps/details?id=pk.com.telenor.phoenix\\u0026amp;hl=en\\u0026amp;gl=US\",\"clicktrackers\":[\"https://track.lacunads.com/clk?viewid=5th4nW7N_Vceu_iStyezex0BUxXapYWayiGyUcUMJkJyQ0yC0cVh5-Ix3qZdDHlPQxenbkZIDS3NgUHQIBUa7VG8NjMQwlhHY_YGLYoHs4KjSSfyqZuWYvqMKhsR4G0PbjUoJVqhmDkkcz6MX2lf2XKzcoGi0xOEyEK_ZFGfKDN_L8DLMlTWQ1-tSa23bY84e8t_Pn0LuaOl6JQgXA1Y67Oml98_FTR3iqMd-19BwL7aesJvdLee0PyJtu6acN5uXKLWeFfriS3ZF9H_zqmxmeMkAW2Ilb_DfuuVxzeKt6oOSxmbb7696iYlDQwrswaxANQE_XPGdEMqOgMFiPFBR2ZjaFI4wpWLXvgq__zYztdP8ZqlHQtwfDGA6TVW5noqGp7-DcdsBSeo6C88xxJYLSJzFurrdYPjvMYBKGor5_LYpdwHrYvatxNEzStKHrXOv6dXzELQQoPSIdNFuo_PAhvcflSXpoRtWgitFFXS3kcr5SkHewWoVeo2vuD5VPsnk3t2e9mlbMcdTfWkkAyGrznYc6kC5G4kkg_4r3bOm7S62ot_5PCL1JHMJq3botQrTP0nMh6KLFUXuWKgorW0TjWtkYmMno7NnpdDsJqba-mMe2WjIwgBXuvWm3kgwfDP7f6NOJe3xt1NnKesOLIXnI3hIIElefIHgz0r1KT96CsRjdsCmNsP1PVfFYmUM6GEaDkMi4o50en-mNjEB10Otnx9dctIWSHJpP-iuQ0nH5XbMvuY7QTUqFGVDdzzjZv1sZok1cM0qSNpx-P1lAl71Yx7xzNyc07krfJ2D_NCtMjqXjl0RlggJUnIPBbWBBHGFhvcfjCmS8Hj0a5WD1UH7JbSloXxlhY6qwk1rRPKdyxBkWjwBXA64HQBGAcUPFHCoBjFgWc-7gIcBXBmSUuETRQtxWpFt7mVCvjWnB1EVEroS-PXSRoCvYpED9fmW-JJkEnHj36olPzM-8dJ3lVvUi_7WThEiVmReOd5qMohUWps2jl4nZ559nJSmDQpxGWkTgDi5PsIz30uEiHBAQyru9jsCer5ZoC5x6sot2iN4XJmLTTcmtKlGYe9N59pdBchkv44lz4NIL33NaaqkT-HEUP7_heBSRLlXmxs1D7xd5NgiptXKksi1LsrUOuKSk9FK9K819HbllIGa2KbLiMQUgAprEMusCl7ZISGf1thyEleKsFiwXDozh4YJJx7_T_R0qQ91lgkS-RHDdGvcIYH50c6iE4vqoHhIXglS9pJDsIoBvnYFb_mGzIAZSCrDJqFU-kGrTQkyc6ealI6AY9jw1F2oKn9Xj9YcnmI07KKjGnG-lDuYdjMDX5_JD1lJS5rwOWkMZ9iP-CR938qGPKYbv1QVZDzk18EBHR9LX1-oDMAISiQHaKV9X6qcSvjXxIqSf6M_gQeSsueU79CTvSbKRiCKnvT_MBXsWcadNxKfpWIexbfWaRB7SkyFhm-UM9QakBcL38SoGq_JYJe8C6hr79NopAAYZVqvmxwisFEn-punH_1kc97dg5VVxhpJrFDksLhdC9JRnvsYbNDlzNtAr3iOyO7nvcxdl3CyFwF06uFdmsgLR5ik3EHsb3tHPsfaKTla_PfYCxlvq1WIyAfpqCxSyBwAO8PmdNq2Gk_TXbfkLNDk_BbIwsua967QGmsaRsjNOFSqjXFw6M1zrGhmt2d9N5JtwSPrd1sRZVwZ2LUTyjB_n_xpUpjR6ddl1JtI6PWQF3bfFkgSSlOC2c3J05ztHJuep_OU-G_CmiOWjiJnED7gI_NY5cMtLwAE1Pj-LWNp9rLA_C2iPsVeepyKZdJzqDu-tjWio-kVY5ln9CQG68n1ThIsD2a75oJO3f3t6jDpwX-FfsXkrNoVbGMo7wfOVaZGtykat4pigTV_AHbPYaHTml6CxxAmojq4GJIgFR4LHi0s8O8X0RiPVFjRH1icS_yb1AoWWq-DSKPynPX7Ov5zqSO2JTs-VB_LbHzJVVkGaRJZGR3m-WquCdO7cQcX_rqgerxwNp2MYRBBFzEKdGyV5IF8TteBGDrCi7HblgxB0dAR6znIM8_brRsWAHKtIzY4t5M6pmSrJP5NC2NUQ2M_PlerxSE22ricYQCCiGbxVEG1oGfeRlzWIzwOxufceckqHaA_6X101ZB2rP5a3nACW7oV9QZzLIzQw_Tjwuo4nrJXlIbWXxPj47cqGGHREItnG9qxQu3LaK1Z_XoFGdOfjmXpyTTOPF01LXygvlso6fHvBAV94YPG9bBvjWAwCBbAsyjfwgkD45Ztf1U7rpVJCuWtVcq2jhrTFHND1Fiwiz_Tk7jTa4ncLhgCS3gd7TkiaZYrcHhB_qLeCI-vPaocVowWzlJpDPmuhxato0JMzzO0V03CV0j3aCjZav2twIsP-GEAovr704GF9ZdEfb3cfWaU3IfIv_Ci62H9k8Y8858kBQRYqKICWKRyVakpWQfaIzcty_2gg9CwIYIQ5cY--QwxeCxSheiKcMi3n9OJp7j3FnvUVE3VgaCVhI7OpJUHU2rjjpM3PdJTnYNmH3nasgFvhEpUFj9Y31yF1nugSYXhKro3COD1wVECFe76AVwM1wOmsZ9zAoKaDWbz8A60ZFvr6pmt5HL_20FZ0y6G3jUQ8Hzcs5M9eFDf30ZpNMqdQX2KgvdyGafDDjIsY4ezcZq-gbOPG0geW55N1mLr7_dK1_NwouPaK--1AC2iDKO8A5W0BLXBL7fAJxx1raqAGvYI1APfToERWjP_4nVLregdf2TBBDt-nxd65rcdoYPKH-fwH9fIT9elSFbkYToNyxalsxEo8Uzpe_83HVfChNj4Yv3pyXbdnFKgEeY3JEU6MsTbiVcys3UrdZn-U66TtFmMJ1QAx1_uyNnJFbZkJ2js_GHY-AQMYa_qZGsbVXzAz3YzwZOh2PScd7m2sSchPTc3z49Eiscyb9rjh4llYbLuxggnCoytjs0mTjAxmef61uOeeYq_NpzT0NuIpmT-NyVDxI36mxynePaChjBHWOShvfBdSRphFHUcRS_QitvIZJfQPkhbBfVrgcIXV04jeKS5t1OEvdtDdMNtZfHk1UV6N-hBi7vzQ2thBidG4mLgj_GMZkr51Cias0PGC_Gkaa1otce-iPFiht5RV8RHY4JmirDq35jLNA-VKlZ0JlCIX_bikKVQpIS1scDYoRsMEFHnWirSq8rDcDls6oOMlNdsvsORT9dyO-MxBIkgPWwlANiemQ5CYLJZe9NEqs8S4mmHFU5shB7qsyIC9cwWEk9gtxRIDqnqji97kGC1rqjPYZIitgLlHiUJURhz48ZUCBZikMbmjUge4AOn3FDCVVw97mZwR6Jo11kdAqX5OvW5wI_V58qwoCASFO2ORt38dgD9EFsE6tjxSb6JpUBiqf7ylJglY2id9PomirKkzj5O450De80fBkVu0wD0P0A7yUTQB5AO9Ap383TqH3NwNwa761cXV1Gz4vkbm-exrGr_L4FMFOmhmS_RFTbdXuGWxnaDrtlMo2RB7Jp8nTsMqyxZEvMdIX5BwM27hX6Z1jCWT2mb6c7ymXSQEO3mzgWvmp6rpkkSpot9LUjZ8RhYZF2aG_X2j2mJq_W8FHUw2fA38i91ff0zA==\\u0026enc=1\",\"https://api.gov-img.site/Ad/AdxEvent?sid=665196088243873080\\u0026sslot=1999\\u0026adtype=1\\u0026s_p_r=0.999900\\u0026slot=100053-100742\\u0026event=click\\u0026os_ver=30\\u0026publisher_type=0\\u0026sspid=1122\\u0026gaid=2bbfb411-afb0-4150-87c4-079c6c4c52c2\\u0026os=android\\u0026adid=33116684178986496\\u0026tgt_pkg_name=pk.com.telenor.phoenix\\u0026pkg=com.meetchat\\u0026pr=10627429\\u0026adxtype=1\\u0026area=asia\\u0026identity_type=0\\u0026req_id=80d7dc6f-1f16-4200-b624-41d8d95084f0\\u0026auc_mode=0\\u0026country=pk\\u0026ts=1673331106\",\"https://api.bytegle.site/ad/trackingclk?sign=Bzf1fY6exCfMpTAACUGAptYA6VYdDU4Hbduqou\\u0026click_flag=0\\u0026gaid=2bbfb411-afb0-4150-87c4-079c6c4c52c2\\u0026target_id=\\u0026action_type=\\u0026click_type=\\u0026click_prop=\\u0026timestamp_ms=\\u0026pkg_name=com.meetchat\\u0026tgt_pkg_name=pk.com.telenor.phoenix\\u0026sdk_vc=0\\u0026disp_ts=1673331105\\u0026ad_id=33116684178986496\\u0026adset_id=43116684176627200\\u0026sid=665196087534050112\\u0026series_id=53000620001205760\\u0026org_id=21999255440560384\\u0026place_id=com.meetchat_1999\\u0026account_id=126306\\u0026placement_id=com.meetchat_1999\\u0026ad_type=1\\u0026adx_type=1\\u0026slot=1999\\u0026ad_style=203\\u0026bt=1\\u0026cf=426-52-1\\u0026os=android\\u0026country=pk\\u0026pro=eyJhZ2UiOiIwIiwiY2l0eSI6InBrLXB1bmphYi1sYWhvcmUiLCJjb3VudHJ5IjoicGsiLCJnZW5kZXIiOiIwIiwiaXNwIjoidGVsZW5vciIsImxhbiI6ImVuIiwibW8yMrYRHREktzM0tzS8CBYNjg4QiIsIm5ldCI6IndpZmkiLCJvcyI6ImFuZHJvaWQiLCJvc182MLczkR0RMrcRFhE3ua9dmVyIjoiMzAuMCIsInN0YXRlIjoicGstcHVuamFiIiwidmVuZG85ER0RJKcjJKckrBAmp6EkpiSqLJAmJKakqiKiETd\\u0026pp=eyJhYmZsYWdzIjoiMTY0MDFfYmFzZV8YGRYYmxoZHC9YmFzZV8YlhkaGJwbr7K8HAwLDE4NTE4X2V4cF8YFhkaGJsbr7K8HBfMCwyNDAzNV8yvDgYr5iWGRmbmpovsTEYFhkaGBkbr7ExGJYZGZyZGC9ZXhwMSIsImFkX2FiZmxhZ3MiOiIiLCJwaXhlbF80shEdERE9B\\u0026source_id=1122\\u0026source_type=1\\u0026access_type=1\\u0026identity_type=0\\u0026publisher_type=0\\u0026preview=0\\u0026dispatch_cache=1\\u0026src_region=hk\\u0026sign_end=end\"]},\"imptrackers\":[\"https://track.lacunads.com/imp?viewid=5th4nW7N_Vceu_iStyezex0BUxXapYWayiGyUcUMJkJyQ0yC0cVh5-Ix3qZdDHlPQxenbkZIDS3NgUHQIBUa7VG8NjMQwlhHY_YGLYoHs4KjSSfyqZuWYvqMKhsR4G0PbjUoJVqhmDkkcz6MX2lf2XKzcoGi0xOEyEK_ZFGfKDN_L8DLMlTWQ1-tSa23bY84e8t_Pn0LuaOl6JQgXA1Y67Oml98_FTR3iqMd-19BwL7aesJvdLee0PyJtu6acN5uXKLWeFfriS3ZF9H_zqmxmeMkAW2Ilb_DfuuVxzeKt6oOSxmbb7696iYlDQwrswaxANQE_XPGdEMqOgMFiPFBR2ZjaFI4wpWLXvgq__zYztdP8ZqlHQtwfDGA6TVW5noqGp7-DcdsBSeo6C88xxJYLSJzFurrdYPjvMYBKGor5_LYpdwHrYvatxNEzStKHrXOv6dXzELQQoPSIdNFuo_PAhvcflSXpoRtWgitFFXS3kcr5SkHewWoVeo2vuD5VPsnk3t2e9mlbMcdTfWkkAyGrznYc6kC5G4kkg_4r3bOm7S62ot_5PCL1JHMJq3botQrTP0nMh6KLFUXuWKgorW0TjWtkYmMno7NnpdDsJqba-mMe2WjIwgBXuvWm3kgwfDP7f6NOJe3xt1NnKesOLIXnI3hIIElefIHgz0r1KT96CsRjdsCmNsP1PVfFYmUM6GEaDkMi4o50en-mNjEB10Otnx9dctIWSHJpP-iuQ0nH5XbMvuY7QTUqFGVDdzzjZv1sZok1cM0qSNpx-P1lAl71Yx7xzNyc07krfJ2D_NCtMjqXjl0RlggJUnIPBbWBBHGFhvcfjCmS8Hj0a5WD1UH7JbSloXxlhY6qwk1rRPKdyxBkWjwBXA64HQBGAcUPFHCoBjFgWc-7gIcBXBmSUuETRQtxWpFt7mVCvjWnB1EVEroS-PXSRoCvYpED9fmW-JJkEnHj36olPzM-8dJ3lVvUi_7WThEiVmReOd5qMohUWps2jl4nZ559nJSmDQpxGWkTgDi5PsIz30uEiHBAQyru9jsCer5ZoC5x6sot2iN4XJmLTTcmtKlGYe9N59pdBchkv44lz4NIL33NaaqkT-HEUP7_heBSRLlXmxs1D7xd5NgiptXKksi1LsrUOuKSk9FK9K819HbllIGa2KbLiMQUgAprEMusCl7ZISGf1thyEleKsFiwXDozh4YJJx7_T_R0qQ91lgkS-RHDdGvcIYH50c6iE4vqoHhIXglS9pJDsIoBvnYFb_mGzIAZSCrDJqFU-kGrTQkyc6ealI6AY9jw1F2oKn9Xj9YcnmI07KKjGnG-lDuYdjMDX5_JD1lJS5rwOWkMZ9iP-CR938qGPKYbv1QVZDzk18EBHR9LX1-oDMAISiQHaKV9X6qcSvjXxIqSf6M_gQeSsueU79CTvSbKRiCKnvT_MBXsWcadNxKfpWIexbfWaRB7SkyFhm-UM9QakBcL38SoGq_JYJe8C6hr79NopAAYZVqvmxwisFEn-punH_1kc97dg5VVxhpJrFDksLhdC9JRnvsYbNDlzNtAr3iOyO7nvcxdl3CyFwF06uFdmsgLR5ik3EHsb3tHPsfaKTla_PfYCxlvq1WIyAfpqCxSyBwAO8PmdNq2Gk_TXbfkLNDk_BbIwsua967QGmsaRsjNOFSqjXFw6M1zrGhmt2d9N5JtwSPrd1sRZVwZ2LUTyjB_n_xpUpjR6ddl1JtI6PWQF3bfFkgSSlOC2c3J05ztHJuep_OU-G_CmiOWjiJnED7gI_NY5cMtLwAE1Pj-LWNp9rLA_C2iPsVeepyKZdJzqDu-tjWio-kVY5ln9CQG68n1ThIsD2a75oJO3f3t6jDpwX-FfsXkrNoVbGMo7wfOVaZGtykat4pigTV_AHbPYaHTml6CxxAmojq4GJIgFR4LHi0s8O8X0RiPVFjRH1icS_yb1AoWWq-DSKPynPX7Ov5zqSO2JTs-VB_LbHzJVVkGaRJZGR3m-WquCdO7cQcX_rqgerxwNp2MYRBBFzEKdGyV5IF8TteBGDrCi7HblgxB0dAR6znIM8_brRsWAHKtIzY4t5M6pmSrJP5NC2NUQ2M_PlerxSE22ricYQCCiGbxVEG1oGfeRlzWIzwOxufceckqHaA_6X101ZB2rP5a3nACW7oV9QZzLIzQw_Tjwuo4nrJXlIbWXxPj47cqGGHREItnG9qxQu3LaK1Z_XoFGdOfjmXpyTTOPF01LXygvlso6fHvBAV94YPG9bBvjWAwCBbAsyjfwgkD45Ztf1U7rpVJCuWtVcq2jhrTFHND1Fiwiz_Tk7jTa4ncLhgCS3gd7TkiaZYrcHhB_qLeCI-vPaocVowWzlJpDPmuhxato0JMzzO0V03CV0j3aCjZav2twIsP-GEAovr704GF9ZdEfb3cfWaU3IfIv_Ci62H9k8Y8858kBQRYqKICWKRyVakpWQfaIzcty_2gg9CwIYIQ5cY--QwxeCxSheiKcMi3n9OJp7j3FnvUVE3VgaCVhI7OpJUHU2rjjpM3PdJTnYNmH3nasgFvhEpUFj9Y31yF1nugSYXhKro3COD1wVECFe76AVwM1wOmsZ9zAoKaDWbz8A60ZFvr6pmt5HL_20FZ0y6G3jUQ8Hzcs5M9eFDf30ZpNMqdQX2KgvdyGafDDjIsY4ezcZq-gbOPG0geW55N1mLr7_dK1_NwouPaK--1AC2iDKO8A5W0BLXBL7fAJxx1raqAGvYI1APfToERWjP_4nVLregdf2TBBDt-nxd65rcdoYPKH-fwH9fIT9elSFbkYToNyxalsxEo8Uzpe_83HVfChNj4Yv3pyXbdnFKgEeY3JEU6MsTbiVcys3UrdZn-U66TtFmMJ1QAx1_uyNnJFbZkJ2js_GHY-AQMYa_qZGsbVXzAz3YzwZOh2PScd7m2sSchPTc3z49Eiscyb9rjh4llYbLuxggnCoytjs0mTjAxmef61uOeeYq_NpzT0NuIpmT-NyVDxI36mxynePaChjBHWOShvfBdSRphFHUcRS_QitvIZJfQPkhbBfVrgcIXV04jeKS5t1OEvdtDdMNtZfHk1UV6N-hBi7vzQ2thBidG4mLgj_GMZkr51Cias0PGC_Gkaa1otce-iPFiht5RV8RHY4JmirDq35jLNA-VKlZ0JlCIX_bikKVQpIS1scDYoRsMEFHnWirSq8rDcDls6oOMlNdsvsORT9dyO-MxBIkgPWwlANiemQ5CYLJZe9NEqs8S4mmHFU5shB7qsyIC9cwWEk9gtxRIDqnqji97kGC1rqjPYZIitgLlHiUJURhz48ZUCBZikMbmjUge4AOn3FDCVVw97mZwR6Jo11kdAqX5OvW5wI_V58qwoCASFO2ORt38dgD9EFsE6tjxSb6JpUBiqf7ylJglY2id9PomirKkzj5O450De80fBkVu0wD0P0A7yUTQB5AO9Ap383TqH3NwNwa761cXV1Gz4vkbm-exrGr_L4FMFOmhmS_RFTbdXuGWxnaDrtlMo2RB7Jp8nTsMqyxZEvMdIX5BwM27hX6Z1jCWT2mb6c7ymXSQEO3mzgWvmp6rpkkSpot9LUjZ8RhYZF2aG_X2j2mJq_W8FHUw2fA38i91ff0zA==\\u0026enc=1\\u0026auction_price=0.10626366\",\"https://api.gov-img.site/Ad/AdxEvent?sid=665196088243873080\\u0026sslot=1999\\u0026adtype=1\\u0026s_p_r=0.999900\\u0026slot=100053-100742\\u0026event=impression\\u0026os_ver=30\\u0026publisher_type=0\\u0026sspid=1122\\u0026gaid=2bbfb411-afb0-4150-87c4-079c6c4c52c2\\u0026os=android\\u0026adid=33116684178986496\\u0026tgt_pkg_name=pk.com.telenor.phoenix\\u0026pkg=com.meetchat\\u0026pr=10627429\\u0026adxtype=1\\u0026area=asia\\u0026identity_type=0\\u0026req_id=80d7dc6f-1f16-4200-b624-41d8d95084f0\\u0026auc_mode=0\\u0026country=pk\\u0026ts=1673331106\",\"https://api.bytegle.site/ad/trackingimp?sign=Bzf1fY6exCfMpTAACUGAptYA6VYdDU4Hbduqou\\u0026gaid=2bbfb411-afb0-4150-87c4-079c6c4c52c2\\u0026timestamp_ms=\\u0026pkg_name=com.meetchat\\u0026tgt_pkg_name=pk.com.telenor.phoenix\\u0026sdk_vc=0\\u0026disp_ts=1673331105\\u0026ad_id=33116684178986496\\u0026adset_id=43116684176627200\\u0026sid=665196087534050112\\u0026series_id=53000620001205760\\u0026org_id=21999255440560384\\u0026place_id=com.meetchat_1999\\u0026account_id=126306\\u0026placement_id=com.meetchat_1999\\u0026ad_type=1\\u0026adx_type=1\\u0026slot=1999\\u0026ad_style=203\\u0026bt=1\\u0026cf=426-52-1\\u0026os=android\\u0026country=pk\\u0026pro=eyJhZ2UiOiIwIiwiY2l0eSI6InBrLXB1bmphYi1sYWhvcmUiLCJjb3VudHJ5IjoicGsiLCJnZW5kZXIiOiIwIiwiaXNwIjoidGVsZW5vciIsImxhbiI6ImVuIiwibW8yMrYRHREktzM0tzS8CBYNjg4QiIsIm5ldCI6IndpZmkiLCJvcyI6ImFuZHJvaWQiLCJvc182MLczkR0RMrcRFhE3ua9dmVyIjoiMzAuMCIsInN0YXRlIjoicGstcHVuamFiIiwidmVuZG85ER0RJKcjJKckrBAmp6EkpiSqLJAmJKakqiKiETd\\u0026pp=eyJhYmZsYWdzIjoiMTY0MDFfYmFzZV8YGRYYmxoZHC9YmFzZV8YlhkaGJwbr7K8HAwLDE4NTE4X2V4cF8YFhkaGJsbr7K8HBfMCwyNDAzNV8yvDgYr5iWGRmbmpovsTEYFhkaGBkbr7ExGJYZGZyZGC9ZXhwMSIsImFkX2FiZmxhZ3MiOiIiLCJwaXhlbF80shEdERE9B\\u0026source_id=1122\\u0026source_type=1\\u0026access_type=1\\u0026identity_type=0\\u0026publisher_type=0\\u0026preview=0\\u0026dispatch_cache=1\\u0026src_region=hk\\u0026sign_end=end\"]}}",
                    "adomain":[
                        "snap.com"
                    ],
                    "bundle":"com.snapchat.android",
                    "iurl":"https://cdn.liftoff.io/customers/1378/previews/320d7bbbdb-ios-fad4aced99.png",
                    "cid":"160529",
                    "crid":"139087",
                    "cat":[
                        "IAB1"
                    ],
                    "h":627,
                    "w":1200,
                    "exp":7200,
                    "ext":{

                    }
                }
            ],
            "seat":"adv7005106418624"
        }
    ],
    "bidid":"0004074a5d6c86804c80",
    "cur":"USD"
}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_opera_native1.0_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				"id":"0000074e00471161ed80_9520",
				"bidid":"0000074e00471161ed80_9520",
				"cur":"USD",
				"seatbid":[
					{
						"bid":[
							{
								"id":"20221107_mBWs8sQoVd99Ge4UISRX",
								"impid":"0000074e00471161ed80_9520",
								"price":1.3,
								"adid":"PNL6nzigKzHJ0aWWd67H",
								"nurl":"https://ams.creativecdn.com/win-notify?tk=GFxk2OqGqqqh1ry1Od_UXpmD93s2xDNKx8iCMSWI6SDX8g7b81Ch82CGyuPwjOj1IgWhLOCsnvrozZTOaU_NZr3JLnEv8Mtu1afYeD0R2Sx6s2SoxRwmK8c2at1N9jYNmTgbEpot1kyS_6fTqWxXSBkA5BTfnfIIltViFfhKmkYWS270mVWzJ4Srw1InZft5mzvBFH4JobWVPgeOBtxiaFzaJJ0C_91WSMgB7Jqx5lomuNCOHU0q5eTHk50YP9DI6mdGyr_GkAYSizGAL0-Nwf0KLdMwK5euBnSqxSEsSDEWMd7tUaJ9VdUBc6N47Tom5MZE5OfH2VnOG4V7467WYBuI48lkhitHcaWhto5WynZe2_LOH4XEk4hE1zNEK9PquIQxUOzUJAFRCMvfAp4_eI-AHKGE4Sg59rmL7HWo1F-JAsbFAF652ZbeCIKRJN3on3rHUxeESUR5UF0GVXL-DmcdLcCUy1iIKfAKYLaCx4YCZw7DFDhRhhRv0-UpBaOvTv-gLRHPjBAmmG8wfTkJIg&wp={AUCTION_PRICE}&tdc=ams",
								"adm":"{\"ver\":\"1.1\",\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://ams.creativecdn.com/images?id=3da8c2301273199507386a3b5a2f8438a04685a0&w=1200&h=627&o=62489743449&fid=MeMnnO3Z4y1ZJNJMYuRf\",\"w\":1200,\"h\":627}}],\"link\":{\"url\":\"https://ams.creativecdn.com/clicks?id=20221107_mBWs8sQoVd99Ge4UISRX&tk=9QFwdW1J7Qkt4eayBhLFYmO6shv-Ok3PZ_zCMf6xsNCu74EEMcVyw1gxnYtjGR65AUsAcDutSR6JG4ucfsZgvyMs3Ef28d3ynheaerHfzUGv1mRJimmbRf8fH3wizIBAj0fce4bDUFzUxXeBvCF5RLhFsKWg1FbGnqSa72KShRlqwxtETZROsYpT2gOmP86Yjvpo6KjWUTXeFLCq4xT8by-iJDuppiusc2zC_4m_lvg1U9ZablVjfmZbcUXEmVUmkDfMpD3s2JMbMyxbEDJnKZiynBFtqtsUiHrvwGbt_Ax0zy1jKbAkPlU4ITKlSwSdxyK1aAt9gDAYQypqzIggjzuGwqQI44VKSllEYXKagu5_hw-1efRicqK7Nrkc3zgki9WEiZt0DDXu8KsaikrBYZvU1zHlVRjwVKMjl26Llb_Oqh8ft8VN1rQkul8kEjDL6_C6Mu4rs8nyyynbLHKLzLZOqm-9DMt1gkVxzz7qF-VFPirz6M6rb_22LTVnBewHkycV0mLJK59VM-UApFDbyZ2RarvicujaVMu4l4Igh0c\"},\"imptrackers\":[\"https://ams.creativecdn.com/win-notify?tk=GFxk2OqGqqqh1ry1Od_UXpmD93s2xDNKx8iCMSWI6SDX8g7b81Ch82CGyuPwjOj1IgWhLOCsnvrozZTOaU_NZr3JLnEv8Mtu1afYeD0R2Sx6s2SoxRwmK8c2at1N9jYNmTgbEpot1kyS_6fTqWxXSBkA5BTfnfIIltViFfhKmkYWS270mVWzJ4Srw1InZft5mzvBFH4JobWVPgeOBtxiaFzaJJ0C_91WSMgB7Jqx5lomuNCOHU0q5eTHk50YP9DI6mdGyr_GkAYSizGAL0-Nwf0KLdMwK5euBnSqxSEsSDEWMd7tUaJ9VdUBc6N47Tom5MZE5OfH2VnOG4V7467WYBuI48lkhitHcaWhto5WynZe2_LOH4XEk4hE1zNEK9PquIQxUOzUJAFRCMvfAp4_eI-AHKGE4Sg59rmL7HWo1F-JAsbFAF652ZbeCIKRJN3on3rHUxeESUR5UF0GVXL-DmcdLcCUy1iIKfAKYLaCx4YCZw7DFDhRhhRv0-UpBaOvTv-gLRHPjBAmmG8wfTkJIg&wp={AUCTION_PRICE}&tdc=ams\"]}",
								"adomain":[
									"alibaba.com"
								],
								"cid":"IBmPijklpco9ip1TYV3O",
								"crid":"PNL6nzigKzHJ0aWWd67H",
								"attr":[
			
								],
								"mtype":4
							}
						]
					}
				]
			}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_midas_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"eaff116e682c41f188d771df2aa1f581",
							"seatbid":[
								{
									"bid":[
										{
											"id":"715978735127455068",
											"impid":"1",
											"price":1,
											"nurl":"https://api.gov-img.site/Ad/AdxEvent?ap={AUCTION_PRICE}&sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&adtype=1&s_p_r=0.900000&slot=100199-101467&event=win&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&scu=aHR0cHM6Ly9hcGkuYnl0ZWdsZS5zaXRlL2JpZ29hZC9hZHdpbj9zaWduPWJKZm5acXgyYmxlWTE1TWlGWHd5dHZaaDEzenA5cXpuZndwVWgmZ2FpZD0xNGMyMjhkZC03Y2U2LTQwM2ItYjdiZS0yNTQyZWI4ZTAwMDgmcmFua19pcD0mc2lnbl9zdGFydD0xJnBrZ19uYW1lPWNvbS5sZW5vdm8uYW55c2hhcmUuZ3BzJnRndF9wa2dfbmFtZT1jb20ubGF6YWRhLmFuZHJvaWQmc2RrX3ZjPTAmZGlzcF90cz0xNjg1NDM4NjMyJmFkX2lkPTMzMjc2ODM5MjAwNjkxNDU2JmFkc2V0X2lkPTQzMjc2ODM5MjAwMjk4MjQwJnNpZD03MTU5Nzg3MzUwNzYxNDc3NzYmc2VyaWVzX2lkPTUzMjc2ODMyMzUyNTQwNDE2Jm9yZ19pZD0yMjI4NjYxNDAxMjE3MDQ5NiZwbGFjZV9pZD1jb20ubGVub3ZvLmFueXNoYXJlLmdwc19hZDpsYXllcl9wX2RyX21wMSZhY2NvdW50X2lkPTEyMDQ4NCZwbGFjZW1lbnRfaWQ9Y29tLmxlbm92by5hbnlzaGFyZS5ncHNfYWQ6bGF5ZXJfcF9kcl9tcDEmYWRfdHlwZT0xJmFkeF90eXBlPTEmc2xvdD1hZDpsYXllcl9wX2RyX21wMSZhZF9zdHlsZT0xMDYmYnQ9MSZjZj00MTctNTItMSZvcz1hbmRyb2lkJmNvdW50cnk9cGgmcHJvPWV5SmhaMlVpT2lJd0lpd2lZMmwwZVNJNkluQm9MWEJ5YjNacGJtTmxYMjh6TDdnd3RyZ3d0ek93bHJDM003SzJNcm12c2JTNlBKRVdFVEczdXJjNk9UeVJIUkU0TkJFV0VUT3l0ekl5dVJFZEVSZ1JGaEUwdWJnUkhSRVlFUllSTmpDM0VSMFJNcmNSRmhFMnQ3SXl0aEVkRVNHb0pCaWJtUnlSRmhFM01yb1JIUkU3dExNMGtSWVJON21SSFJFd3R6STVON1N5RVJZUk43bXZ0akMzTTVFZEVUSzNFUllSTjdtdnV6SzVFUjBSRzVjWWtSWVJPYm93dWpLUkhSRTROQmE0T1RlN05MY3hzcTliMlpmY0dGdGNHRnVaMkVpTENKMlpXNWtiM0lpT2lKUFVGQlBJbk4mcHA9ZXlKaFltWnNZV2R6SWpvaU1UWTBNREZmZEdWemRGOFlHNVlaR3hrY0c2OVpYaHdYekFzTWpZNU16WmZZbUZ6WlY4WWxoa2JHcGdZcjdFd3ViS3ZtQllaR3hzWkc2OVpYaHdYekFzTWpVMk9UZGZZbUZ6WlY4WUZoa2JHWnViTDdFd3ViS3ZtQllaR3h3Wm02OVlXRXhMREkyTnpVeFgySmlNQ3d5TmpReE5WOHdzSmlXR1JzY21CeXZzcnc0TDVpWUZoa2JISm9jTDdFd3ViS1lFUllSTUxJdnNMRXpOakN6dVpFZEVSRVdFVGcwdkRLMkw3U3lFUjBSRVRkJnNvdXJjZV9pZD0xMTIyJnNvdXJjZV90eXBlPTEmYWNjZXNzX3R5cGU9MSZpZGVudGl0eV90eXBlPTAmcHVibGlzaGVyX3R5cGU9MCZwcmV2aWV3PTAmZGlzcGF0Y2hfY2FjaGU9MSZzcmNfcmVnaW9uPWhrJnNpZ25fZW5kPWVuZA%3D%3D&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632",
											"burl":"https://api.gov-img.site/Ad/AdxEvent?ap={AUCTION_PRICE}&sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&adtype=1&s_p_r=0.900000&slot=100199-101467&event=billing&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632",
											"lurl":"https://api.gov-img.site/Ad/AdxEvent?ap={AUCTION_PRICE}&sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&lr={AUCTION_LOSS}&adtype=1&s_p_r=0.900000&slot=100199-101467&event=loss&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632",
											"adm":"{\"native\":{\"assets\":[{\"id\":1,\"required\":1,\"img\":{\"type\":3,\"w\":720,\"h\":1280,\"url\":\"http://static-dev.rqmob.com/test/sa/20221108/12878e41e7f42119cf614e165e58563f__FILE_CM____FILE_CM480_720____WEBP__.png\"}}],\"link\":{\"url\":\"https://c.lazada.com.ph/t/c.0JuzgV?rta_event_id=715978735076147776&rta_token=&os=android&gps_adid=14c228dd-7ce6-403b-b7be-2542eb8e0008&device_model=CPH1729&device_make=OPPO&sub_id1=715978735076147776&sub_id2=14c228dd-7ce6-403b-b7be-2542eb8e0008\",\"clicktrackers\":[\"https://api.gov-img.site/Ad/AdxEvent?sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&adtype=1&s_p_r=0.900000&slot=100199-101467&event=click&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632\",\"https://pb.mobshark.net/api/affiliate/postback/click?click_id=715978735076147776&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&offer_id=10278&pub=2VMOuO2NWwe3jzz6Far8pQ==&sub=120484&redirect=true\",\"https://api.bytegle.site/ad/trackingclk?sign=bJfnZqx2bleY15MiFXwytvZh13zp9qznfwpUh&click_flag=0&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&target_id=&action_type=&click_type=&click_prop=&sign_start=1&pkg_name=com.lenovo.anyshare.gps&tgt_pkg_name=com.lazada.android&sdk_vc=0&disp_ts=1685438632&ad_id=33276839200691456&adset_id=43276839200298240&sid=715978735076147776&series_id=53276832352540416&org_id=22286614012170496&place_id=com.lenovo.anyshare.gps_ad:layer_p_dr_mp1&account_id=120484&placement_id=com.lenovo.anyshare.gps_ad:layer_p_dr_mp1&ad_type=1&adx_type=1&slot=ad:layer_p_dr_mp1&ad_style=106&bt=1&cf=417-52-1&os=android&country=ph&pro=eyJhZ2UiOiIwIiwiY2l0eSI6InBoLXByb3ZpbmNlX28zL7gwtrgwtzOwlrC3M7K2MrmvsbS6PJEWETG3urc6OTyRHRE4NBEWETOytzIyuREdERgRFhE0ubgRHREYERYRNjC3ER0RMrcRFhE2t7IythEdESGoJBibmRyRFhE3MroRHRE7tLM0kRYRN7mRHREwtzI5N7SyERYRN7mvtjC3M5EdETK3ERYRN7mvuzK5ER0RG5cYkRYRObowujKRHRE4NBa4OTe7NLcxsq9b2ZfcGFtcGFuZ2EiLCJ2ZW5kb3IiOiJPUFBPInN&pp=eyJhYmZsYWdzIjoiMTY0MDFfdGVzdF8YG5YZGxkcG69ZXhwXzAsMjY5MzZfYmFzZV8YlhkbGpgYr7EwubKvmBYZGxsZG69ZXhwXzAsMjU2OTdfYmFzZV8YFhkbGZubL7EwubKvmBYZGxwZm69YWExLDI2NzUxX2JiMCwyNjQxNV8wsJiWGRscmByvsrw4L5iYFhkbHJocL7EwubKYERYRMLIvsLEzNjCzuZEdEREWETg0vDK2L7SyER0RETd&source_id=1122&source_type=1&access_type=1&identity_type=0&publisher_type=0&preview=0&dispatch_cache=1&src_region=hk&sign_end=end\"]},\"imptrackers\":[\"https://api.gov-img.site/Ad/AdxEvent?sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&adtype=1&s_p_r=0.900000&slot=100199-101467&event=impression&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632\",\"https://c.lazada.com.ph/i/c.0JuzgV?rta_event_id=715978735076147776&rta_token=&os=android&gps_adid=14c228dd-7ce6-403b-b7be-2542eb8e0008&device_model=CPH1729&device_make=OPPO&sub_id1=715978735076147776&sub_id2=14c228dd-7ce6-403b-b7be-2542eb8e0008\",\"https://api.bytegle.site/ad/trackingimp?sign=bJfnZqx2bleY15MiFXwytvZh13zp9qznfwpUh&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&sign_start=1&pkg_name=com.lenovo.anyshare.gps&tgt_pkg_name=com.lazada.android&sdk_vc=0&disp_ts=1685438632&ad_id=33276839200691456&adset_id=43276839200298240&sid=715978735076147776&series_id=53276832352540416&org_id=22286614012170496&place_id=com.lenovo.anyshare.gps_ad:layer_p_dr_mp1&account_id=120484&placement_id=com.lenovo.anyshare.gps_ad:layer_p_dr_mp1&ad_type=1&adx_type=1&slot=ad:layer_p_dr_mp1&ad_style=106&bt=1&cf=417-52-1&os=android&country=ph&pro=eyJhZ2UiOiIwIiwiY2l0eSI6InBoLXByb3ZpbmNlX28zL7gwtrgwtzOwlrC3M7K2MrmvsbS6PJEWETG3urc6OTyRHRE4NBEWETOytzIyuREdERgRFhE0ubgRHREYERYRNjC3ER0RMrcRFhE2t7IythEdESGoJBibmRyRFhE3MroRHRE7tLM0kRYRN7mRHREwtzI5N7SyERYRN7mvtjC3M5EdETK3ERYRN7mvuzK5ER0RG5cYkRYRObowujKRHRE4NBa4OTe7NLcxsq9b2ZfcGFtcGFuZ2EiLCJ2ZW5kb3IiOiJPUFBPInN&pp=eyJhYmZsYWdzIjoiMTY0MDFfdGVzdF8YG5YZGxkcG69ZXhwXzAsMjY5MzZfYmFzZV8YlhkbGpgYr7EwubKvmBYZGxsZG69ZXhwXzAsMjU2OTdfYmFzZV8YFhkbGZubL7EwubKvmBYZGxwZm69YWExLDI2NzUxX2JiMCwyNjQxNV8wsJiWGRscmByvsrw4L5iYFhkbHJocL7EwubKYERYRMLIvsLEzNjCzuZEdEREWETg0vDK2L7SyER0RETd&source_id=1122&source_type=1&access_type=1&identity_type=0&publisher_type=0&preview=0&dispatch_cache=1&src_region=hk&sign_end=end\"]}}",
											"adid":"33276839200691456",
											"adomain":[
												"c.lazada.com.ph"
											],
											"bundle":"com.lazada.android",
											"crid":"6d7538d5112f1e93ea96a4c72be9a6b9_1",
											"ext":{
						
											}
										}
									]
								}
							],
							"cur":"USD"
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_midas1_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"eaff116e682c41f188d771df2aa1f581",
						"seatbid":[
							{
								"bid":[
									{
										"id":"715978735127455068",
										"impid":"1",
										"price":1,
										"nurl":"https://api.gov-img.site/Ad/AdxEvent?ap={AUCTION_PRICE}&sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&adtype=1&s_p_r=0.900000&slot=100199-101467&event=win&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&scu=aHR0cHM6Ly9hcGkuYnl0ZWdsZS5zaXRlL2JpZ29hZC9hZHdpbj9zaWduPWJKZm5acXgyYmxlWTE1TWlGWHd5dHZaaDEzenA5cXpuZndwVWgmZ2FpZD0xNGMyMjhkZC03Y2U2LTQwM2ItYjdiZS0yNTQyZWI4ZTAwMDgmcmFua19pcD0mc2lnbl9zdGFydD0xJnBrZ19uYW1lPWNvbS5sZW5vdm8uYW55c2hhcmUuZ3BzJnRndF9wa2dfbmFtZT1jb20ubGF6YWRhLmFuZHJvaWQmc2RrX3ZjPTAmZGlzcF90cz0xNjg1NDM4NjMyJmFkX2lkPTMzMjc2ODM5MjAwNjkxNDU2JmFkc2V0X2lkPTQzMjc2ODM5MjAwMjk4MjQwJnNpZD03MTU5Nzg3MzUwNzYxNDc3NzYmc2VyaWVzX2lkPTUzMjc2ODMyMzUyNTQwNDE2Jm9yZ19pZD0yMjI4NjYxNDAxMjE3MDQ5NiZwbGFjZV9pZD1jb20ubGVub3ZvLmFueXNoYXJlLmdwc19hZDpsYXllcl9wX2RyX21wMSZhY2NvdW50X2lkPTEyMDQ4NCZwbGFjZW1lbnRfaWQ9Y29tLmxlbm92by5hbnlzaGFyZS5ncHNfYWQ6bGF5ZXJfcF9kcl9tcDEmYWRfdHlwZT0xJmFkeF90eXBlPTEmc2xvdD1hZDpsYXllcl9wX2RyX21wMSZhZF9zdHlsZT0xMDYmYnQ9MSZjZj00MTctNTItMSZvcz1hbmRyb2lkJmNvdW50cnk9cGgmcHJvPWV5SmhaMlVpT2lJd0lpd2lZMmwwZVNJNkluQm9MWEJ5YjNacGJtTmxYMjh6TDdnd3RyZ3d0ek93bHJDM003SzJNcm12c2JTNlBKRVdFVEczdXJjNk9UeVJIUkU0TkJFV0VUT3l0ekl5dVJFZEVSZ1JGaEUwdWJnUkhSRVlFUllSTmpDM0VSMFJNcmNSRmhFMnQ3SXl0aEVkRVNHb0pCaWJtUnlSRmhFM01yb1JIUkU3dExNMGtSWVJON21SSFJFd3R6STVON1N5RVJZUk43bXZ0akMzTTVFZEVUSzNFUllSTjdtdnV6SzVFUjBSRzVjWWtSWVJPYm93dWpLUkhSRTROQmE0T1RlN05MY3hzcTliMlpmY0dGdGNHRnVaMkVpTENKMlpXNWtiM0lpT2lKUFVGQlBJbk4mcHA9ZXlKaFltWnNZV2R6SWpvaU1UWTBNREZmZEdWemRGOFlHNVlaR3hrY0c2OVpYaHdYekFzTWpZNU16WmZZbUZ6WlY4WWxoa2JHcGdZcjdFd3ViS3ZtQllaR3hzWkc2OVpYaHdYekFzTWpVMk9UZGZZbUZ6WlY4WUZoa2JHWnViTDdFd3ViS3ZtQllaR3h3Wm02OVlXRXhMREkyTnpVeFgySmlNQ3d5TmpReE5WOHdzSmlXR1JzY21CeXZzcnc0TDVpWUZoa2JISm9jTDdFd3ViS1lFUllSTUxJdnNMRXpOakN6dVpFZEVSRVdFVGcwdkRLMkw3U3lFUjBSRVRkJnNvdXJjZV9pZD0xMTIyJnNvdXJjZV90eXBlPTEmYWNjZXNzX3R5cGU9MSZpZGVudGl0eV90eXBlPTAmcHVibGlzaGVyX3R5cGU9MCZwcmV2aWV3PTAmZGlzcGF0Y2hfY2FjaGU9MSZzcmNfcmVnaW9uPWhrJnNpZ25fZW5kPWVuZA%3D%3D&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632",
										"burl":"https://api.gov-img.site/Ad/AdxEvent?ap={AUCTION_PRICE}&sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&adtype=1&s_p_r=0.900000&slot=100199-101467&event=billing&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632",
										"lurl":"https://api.gov-img.site/Ad/AdxEvent?ap={AUCTION_PRICE}&sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&lr={AUCTION_LOSS}&adtype=1&s_p_r=0.900000&slot=100199-101467&event=loss&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632",
										"adm":"{\"native\":{\"assets\":[{\"id\":2,\"required\":1,\"img\":{\"type\":1,\"w\":160,\"h\":160,\"url\":\"http://cdn-3s.mobvista.com/upload/default/6346851971e95.png\"}},{\"id\":3,\"required\":1,\"title\":{\"text\":\"Shopee 10.10 Brands Festival\"}},{\"id\":4,\"data\":{\"value\":\"Beli produk original dari brand favoritmu di Shopee 10.10 Brands Festival!\"}},{\"id\":5,\"data\":{\"value\":\"Download\"}}],\"link\":{\"url\":\"https://c.lazada.com.ph/t/c.0JuzgV?rta_event_id=715978735076147776&rta_token=&os=android&gps_adid=14c228dd-7ce6-403b-b7be-2542eb8e0008&device_model=CPH1729&device_make=OPPO&sub_id1=715978735076147776&sub_id2=14c228dd-7ce6-403b-b7be-2542eb8e0008\",\"clicktrackers\":[\"https://api.gov-img.site/Ad/AdxEvent?sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&adtype=1&s_p_r=0.900000&slot=100199-101467&event=click&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632\",\"https://pb.mobshark.net/api/affiliate/postback/click?click_id=715978735076147776&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&offer_id=10278&pub=2VMOuO2NWwe3jzz6Far8pQ==&sub=120484&redirect=true\",\"https://api.bytegle.site/ad/trackingclk?sign=bJfnZqx2bleY15MiFXwytvZh13zp9qznfwpUh&click_flag=0&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&target_id=&action_type=&click_type=&click_prop=&sign_start=1&pkg_name=com.lenovo.anyshare.gps&tgt_pkg_name=com.lazada.android&sdk_vc=0&disp_ts=1685438632&ad_id=33276839200691456&adset_id=43276839200298240&sid=715978735076147776&series_id=53276832352540416&org_id=22286614012170496&place_id=com.lenovo.anyshare.gps_ad:layer_p_dr_mp1&account_id=120484&placement_id=com.lenovo.anyshare.gps_ad:layer_p_dr_mp1&ad_type=1&adx_type=1&slot=ad:layer_p_dr_mp1&ad_style=106&bt=1&cf=417-52-1&os=android&country=ph&pro=eyJhZ2UiOiIwIiwiY2l0eSI6InBoLXByb3ZpbmNlX28zL7gwtrgwtzOwlrC3M7K2MrmvsbS6PJEWETG3urc6OTyRHRE4NBEWETOytzIyuREdERgRFhE0ubgRHREYERYRNjC3ER0RMrcRFhE2t7IythEdESGoJBibmRyRFhE3MroRHRE7tLM0kRYRN7mRHREwtzI5N7SyERYRN7mvtjC3M5EdETK3ERYRN7mvuzK5ER0RG5cYkRYRObowujKRHRE4NBa4OTe7NLcxsq9b2ZfcGFtcGFuZ2EiLCJ2ZW5kb3IiOiJPUFBPInN&pp=eyJhYmZsYWdzIjoiMTY0MDFfdGVzdF8YG5YZGxkcG69ZXhwXzAsMjY5MzZfYmFzZV8YlhkbGpgYr7EwubKvmBYZGxsZG69ZXhwXzAsMjU2OTdfYmFzZV8YFhkbGZubL7EwubKvmBYZGxwZm69YWExLDI2NzUxX2JiMCwyNjQxNV8wsJiWGRscmByvsrw4L5iYFhkbHJocL7EwubKYERYRMLIvsLEzNjCzuZEdEREWETg0vDK2L7SyER0RETd&source_id=1122&source_type=1&access_type=1&identity_type=0&publisher_type=0&preview=0&dispatch_cache=1&src_region=hk&sign_end=end\"]},\"imptrackers\":[\"https://api.gov-img.site/Ad/AdxEvent?sid=715978735127455068&sslot=ad%3Alayer_p_dr_mp1&adtype=1&s_p_r=0.900000&slot=100199-101467&event=impression&os_ver=7.1&publisher_type=0&sspid=1122&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&os=android&adid=33276839200691456&tgt_pkg_name=com.lazada.android&pkg=com.lenovo.anyshare.gps&pr=1228317&adxtype=1&area=asia&identity_type=0&req_id=eaff116e682c41f188d771df2aa1f581&auc_mode=0&country=ph&ts=1685438632\",\"https://c.lazada.com.ph/i/c.0JuzgV?rta_event_id=715978735076147776&rta_token=&os=android&gps_adid=14c228dd-7ce6-403b-b7be-2542eb8e0008&device_model=CPH1729&device_make=OPPO&sub_id1=715978735076147776&sub_id2=14c228dd-7ce6-403b-b7be-2542eb8e0008\",\"https://api.bytegle.site/ad/trackingimp?sign=bJfnZqx2bleY15MiFXwytvZh13zp9qznfwpUh&gaid=14c228dd-7ce6-403b-b7be-2542eb8e0008&sign_start=1&pkg_name=com.lenovo.anyshare.gps&tgt_pkg_name=com.lazada.android&sdk_vc=0&disp_ts=1685438632&ad_id=33276839200691456&adset_id=43276839200298240&sid=715978735076147776&series_id=53276832352540416&org_id=22286614012170496&place_id=com.lenovo.anyshare.gps_ad:layer_p_dr_mp1&account_id=120484&placement_id=com.lenovo.anyshare.gps_ad:layer_p_dr_mp1&ad_type=1&adx_type=1&slot=ad:layer_p_dr_mp1&ad_style=106&bt=1&cf=417-52-1&os=android&country=ph&pro=eyJhZ2UiOiIwIiwiY2l0eSI6InBoLXByb3ZpbmNlX28zL7gwtrgwtzOwlrC3M7K2MrmvsbS6PJEWETG3urc6OTyRHRE4NBEWETOytzIyuREdERgRFhE0ubgRHREYERYRNjC3ER0RMrcRFhE2t7IythEdESGoJBibmRyRFhE3MroRHRE7tLM0kRYRN7mRHREwtzI5N7SyERYRN7mvtjC3M5EdETK3ERYRN7mvuzK5ER0RG5cYkRYRObowujKRHRE4NBa4OTe7NLcxsq9b2ZfcGFtcGFuZ2EiLCJ2ZW5kb3IiOiJPUFBPInN&pp=eyJhYmZsYWdzIjoiMTY0MDFfdGVzdF8YG5YZGxkcG69ZXhwXzAsMjY5MzZfYmFzZV8YlhkbGpgYr7EwubKvmBYZGxsZG69ZXhwXzAsMjU2OTdfYmFzZV8YFhkbGZubL7EwubKvmBYZGxwZm69YWExLDI2NzUxX2JiMCwyNjQxNV8wsJiWGRscmByvsrw4L5iYFhkbHJocL7EwubKYERYRMLIvsLEzNjCzuZEdEREWETg0vDK2L7SyER0RETd&source_id=1122&source_type=1&access_type=1&identity_type=0&publisher_type=0&preview=0&dispatch_cache=1&src_region=hk&sign_end=end\"]}}",
										"adid":"33276839200691456",
										"adomain":[
											"c.lazada.com.ph"
										],
										"bundle":"com.lazada.android",
										"crid":"6d7538d5112f1e93ea96a4c72be9a6b9_1",
										"ext":{
					
										}
									}
								]
							}
						],
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//telecoming
	r.HandleFunc("/mock/868/get_telecoming_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "DspId":100038,
				  "WinPrice":2,
				  "DspName":"telecoming",
				  "FloorPrice":0.01,
				  "ValidImpDuration":0,
				  "BidPriceCoefficient":0,
				  "id":"21790c9288aa47e090b08ef513b3a6d3",
				  "seatbid":[
					{
					  "bid":[
						{
						  "id":"21790c9288aa47e090b08ef513b3a6d3",
						  "impid":"1",
						  "price":2,
						  "nurl":"https://shareiteu.lynxio.net/win/?rtb_seat_id=shareiteu\u0026auid={AUCTION_ID}\u0026aupr={AUCTION_PRICE}\u0026adid=32a7f088706db9c5243dbeaf233592a8:ZA:65501:008bc49f08bae95340bd50ae8ab74148:5c36065aad22ef36c55bb7e3497678e9:6905:47759\u0026uid=b266e3c171d31c3f0b3e69557ec21882",
						  "burl":"https://shareiteu.lynxio.net/imp/?rtb_seat_id=shareiteu\u0026auid={AUCTION_ID}\u0026aupr={AUCTION_PRICE}\u0026adid=32a7f088706db9c5243dbeaf233592a8:ZA:65501:008bc49f08bae95340bd50ae8ab74148:5c36065aad22ef36c55bb7e3497678e9:6905:47759\u0026uid=b266e3c171d31c3f0b3e69557ec21882",
						  "adm":"{\"native\":{\"link\":{\"url\":\"https://shareiteu.lynxio.net/click/?h=32a7f088706db9c5243dbeaf233592a8:ZA:65501:008bc49f08bae95340bd50ae8ab74148:5c36065aad22ef36c55bb7e3497678e9:6905:47759\u0026uid=b266e3c171d31c3f0b3e69557ec21882\u0026printts=1668484378\u0026auid=21790c9288aa47e090b08ef513b3a6d3\"},\"ver\":\"1.1\",\"imptrackers\":[\"https://shareiteu.lynxio.net/imp/?rtb_seat_id=shareiteu\u0026auid={AUCTION_ID}\u0026aupr={AUCTION_PRICE}\u0026adid=32a7f088706db9c5243dbeaf233592a8:ZA:65501:008bc49f08bae95340bd50ae8ab74148:5c36065aad22ef36c55bb7e3497678e9:6905:47759\u0026uid=b266e3c171d31c3f0b3e69557ec21882\"],\"assets\":[{\"id\":1,\"required\":1,\"img\":{\"type\":3,\"url\":\"https://d77ulgzzk2dyh.cloudfront.net/8lG4FoqHiN006905047759_1200x627.jpg\",\"h\":627,\"w\":1200}},{\"id\":2,\"required\":1,\"img\":{\"type\":1,\"url\":\"https://d77ulgzzk2dyh.cloudfront.net/KlPLMV0DNe006905047759_80x80.png\",\"h\":80,\"w\":80}},{\"id\":3,\"required\":1,\"title\":{\"text\":\"PUT YOUR FINGER HERE!\"}},{\"id\":4,\"required\":1,\"data\":{\"type\":2,\"value\":\"Unlock THE BEST GAMES for your smartphone! (T\u0026C Apply - R10/day)\"}},{\"id\":5,\"required\":1,\"data\":{\"type\":12,\"value\":\"PLAY\"}}]}}",
						  "adomain":[
							"teensup.net"
						  ],
						  "cid":"6905",
						  "crid":"47759",
						  "cat":[
							"IAB1",
							"IAB19-29"
						  ]
						}
					  ]
					}
				  ],
				  "cur":"USD",
				  "ext":null,
				  "vaststr":"",
				  "bidid":"",
				  "expire_time":0
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_tele_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"a89fn23absb3227",
						"cur":"USD",
						"seatbid":[
							{
								"bid":[
									{
										"id":"a89fn23absb3228",
										"impid":"bha9h9234bvks344",
										"price":0.2,
										"adomain":[
											"sn.vidbites.net"
										],
										"cid":"7087",
										"crid":"48168",
										"burl":"https://[telecoming_integration_host]/imp/?rtb_seat_id=[telecoming_integration_id]&auid=${AUCTION_ID}&aupr=${AUCTION_PRICE}&adid=[telecoming_bid_id]&uid=[telecoming_user_id]",
										"nurl":"https://[telecoming_integration_host]/win/?rtb_seat_id=[telecoming_integration_id]&auid=${AUCTION_ID}&aupr=${AUCTION_PRICE}&adid=[telecoming_bid_id]&uid=[telecoming_user_id]",
										"cat":[
											"IAB1",
											"IAB19-29"
										],
										"adid":"c81e728d9d4c2f636f067f89cc14862c:SN:60801:bcfa98f8646876b5d5e70f36211b2173:94c9298b94dc161216483cfaf7ad8333:7087:48168",
										"adm":"<a href=\"https://[telecoming_integration_host]/click/?h=c81e728d9d4c2f636f067f89cc14862c:SN:60801:bcfa98f8646876b5d5e70f36211b2173:94c9298b94dc161216483cfaf7ad8333:7087:48168&uid=04e8e79b00930d99f2454c5ea6c761b8&printts=1670255978&auid=a89fn23absb3227\"> <img src =\"http://d77ulgzzk2dyh.cloudfront.net/s3C8i1WOGk007087048168_320x50.jpg\"></a>",
										"h":50,
										"w":320
									}
								]
							}
						]
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_tele_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"a89fn23absb3222",
						"cur":"USD",
						"seatbid":[
							{
								"bid":[
									{
										"id":"a89fn23absb3222",
										"impid":"1",
										"price":0.2,
										"adomain":[
											"sn.kiddies-tube.com"
										],
										"cid":"7054",
										"crid":"48053",
										"burl":"https://[telecoming_integration_host]/imp/?rtb_seat_id=[telecoming_integration_id]&auid=${AUCTION_ID}&aupr=${AUCTION_PRICE}&adid=[telecoming_bid_id]&uid=[telecoming_user_id]",
										"nurl":"https://[telecoming_integration_host]/win/?rtb_seat_id=[telecoming_integration_id]&auid=${AUCTION_ID}&aupr=${AUCTION_PRICE}&adid=[telecoming_bid_id]&uid=[telecoming_user_id]",
										"cat":[
											"IAB1",
											"IAB19-29"
										],
										"protocol":2,
										"adid":"c81db8fc8f086c6bbba8361384e51706:SN:60801:4bd4e7b3f7b01c60a6c564e23fa32a2a:d76d76273584cca0c27e4ffc29463aff:7054:48053",
										"adm":"<?xml version=\"1.0\" encoding=\"UTF-8\"?><VAST version=\"2.0\"><Ad id=\"48053\"><InLine><AdSystem>TLC</AdSystem><AdTitle>LIT-48053</AdTitle><Impression><![CDATA[https://[telecoming_integration_host]/imp/?rtb_seat_id=[telecoming_integration_id]&auid=${AUCTION_ID}&aupr=${AUCTION_PRICE}&adid=[telecoming_bid_id]&uid=[telecoming_user_id]]]></Impression><Creatives><Creative><Linear><Duration>00:00:17</Duration><VideoClicks><ClickThrough><![CDATA[https://[telecoming_integration_host]/click/?h=c81db8fc8f086c6bbba8361384e51706:SN:60801:4bd4e7b3f7b01c60a6c564e23fa32a2a:d76d76273584cca0c27e4ffc29463aff:7054:48053&uid=262d95a4dd66bc2575de3c008fa79465&printts=1669106875&auid=a89fn23absb3222]]></ClickThrough></VideoClicks><MediaFiles><MediaFile delivery=\"progressive\" height=\"480\" width=\"320\" type=\"video/mp4\"><![CDATA[http://d77ulgzzk2dyh.cloudfront.net/S1kzunWPve007054048053_video_320x480.mp4]]></MediaFile></MediaFiles></Linear></Creative></Creatives></InLine></Ad></VAST>",
										"h":480,
										"w":320,
										"attr":[
											6
										]
									}
								]
							}
						]
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//bigo
	r.HandleFunc("/mock/868/get_bigo_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				"DspId":100062,
				"WinPrice":0.18330325,
				"DspName":"bigo",
				"FloorPrice":0.01,
				"ValidImpDuration":0,
				"BidPriceCoefficient":0,
				"id":"15d1bb3e-cc59-4024-add3-1b5019be5fd4",
				"seatbid":[
					{
						"bid":[
							{
								"id":"644841349853570256",
								"impid":"1",
								"price":0.18330325,
								"nurl":"https://api.gov-img.site/Ad/AdxEvent?bnurl=aHR0cHM6Ly9hcGkuYnl0ZWdsZS5zaXRlL2JpZ29hZC9hZHdpbj9zaWduPUIwNWtmbkpaRU5OZjJlODk0UUVhQ3NWVlBkUHlqU3FsY0w1aGdiJmdhaWQ9ZjAyZTc4YjAtZDRkNy00M2EyLWFiMTItNTAxZmQ2NTY3ZmQ2JnJhbmtfaXA9MTAuMTUyLjcyLjQzOjIyOTk2JnRpbWVzdGFtcF9tcz0wJnBrZ19uYW1lPWNvbS5sZW5vdm8uYW55c2hhcmUuZ3BzJnRndF9wa2dfbmFtZT0mc2RrX3ZjPTAmZGlzcF90cz0xNjY4NDc4MTU4JmFkX2lkPTMyNTM1ODAwMDE3MzI3MzYwJmFkc2V0X2lkPTQyNTM1ODAwMDE2NTczNjk2JnNpZD02NDQ4NDEzNDk0Mjg5NTMyNDkmc2VyaWVzX2lkPTUyMzU3NTIyMzk5NTk3ODI0Jm9yZ19pZD0yMjIwMTk0NTM2MzY0ODc2OCZwbGFjZV9pZD1jb20ubGVub3ZvLmFueXNoYXJlLmdwc19hZDpsYXllcl9wX2xyMSZhY2NvdW50X2lkPTExMjc0MiZwbGFjZW1lbnRfaWQ9Y29tLmxlbm92by5hbnlzaGFyZS5ncHNfYWQ6bGF5ZXJfcF9scjEmYWRfdHlwZT0xJmFkeF90eXBlPTEmc2xvdD1hZDpsYXllcl9wX2xyMSZhZF9zdHlsZT0xMDUmYnQ9MSZjZj0xNDMtMTEmb3M9YW5kcm9pZCZjb3VudHJ5PXV6JnBybz1leUpoWjJVaU9pSXdJaXdpWTJsMGVTSTZJaUlzSW1OdmRXNTBjbmtpT2lKMWVpSXNJbWRsYm1SbGNpSTZJakFpTENKcGMzQWlPaUoxZWkxdGRITWlMQ0pzWVc0aU9pSnlkU0lzSW0xdlpHVnNJam9pTWpJd01URXhOMVJISWl3aWJtVjBJam9pZDJsbWFTSXNJbTg1a1IwUk1MY3lPVGUwc2hFV0VUZTVyN1l3dHpPUkhSRTVPcEVXRVRlNXI3c3l1UkVkRVJpWWx4Z1JGaEU1dWpDNk1wRWRFUkVXRVRzeXR6STN1UkVkRVN3MHNMZTJ0SkU5QiZwcD1leUpoWW1ac1lXZHpJam9pTVRZME1qaGZZbUZ6WlY4WkZoa1ptSmljcjdLOEhCZk1Dd3hPRFV4T0Y4eXZEZ3ZtQllaR1JnYW15OVltRnpaVjhZbGhrWkhCbWFyN0N3bUJFV0VUQ3lMN0N4TXpZd3M3bVJIUkVSRmhFNE5Md3l0aTlhV1FpT2lJNU1ESXpOREF6TXpjeE5UTXpPREEyTURnaWZCJnNvdXJjZV9pZD0xMTIyJnNvdXJjZV90eXBlPTEmYWNjZXNzX3R5cGU9MSZpZGVudGl0eV90eXBlPTAmcHVibGlzaGVyX3R5cGU9MCZwcmV2aWV3PTAmc3JjX3JlZ2lvbj1oayZzaWduX2VuZD1lbmQ%3D\u0026ap={AUCTION_PRICE}\u0026sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=win\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158",
								"burl":"https://api.gov-img.site/Ad/AdxEvent?ap={AUCTION_PRICE}\u0026sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=billing\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158",
								"lurl":"https://api.gov-img.site/Ad/AdxEvent?ap=${AUCTION_PRICE}\u0026sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026lr={AUCTION_LOSS}\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=loss\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158",
								"adm":"{\"native\":{\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://gdl.bytegle.tech/as/web-source/4h8/1ZxYEX.jpeg\",\"w\":1200,\"h\":627}},{\"id\":2,\"img\":{\"url\":\"https://gdl.bytegle.tech/as/web-source/4h8/25saMJ.png\",\"w\":128,\"h\":128}},{\"id\":3,\"title\":{\"text\":\"Действенно\"}},{\"id\":4,\"data\":{\"value\":\"Как похудеть к отпуску за 2 недели\"}},{\"id\":5,\"data\":{\"value\":\"Узнать больше\"}}],\"link\":{\"url\":\"https://api.imotech.video/ad/pixelfile.html?url=https%3A%2F%2Fnews-sphere.com%2F151527-vrach-nazvala-deistvennyi-metod-poxudet-k-otpusku-za-dve-nedeli.html%3Fsource%3D1350%26id%3D85894%26tid%3D85894%26bbg%3DChdjb20ubGVub3ZvLmFueXNoYXJlLmdwcxIkZjAyZTc4YjAtZDRkNy00M2EyLWFiMTItNTAxZmQ2NTY3ZmQ2GKGByYmn9N38QgiETMyNTM1ODAwMDE3MzI3MzYwKhI5MDIzNDAzMzcxNTMzODA2MDhiBDExMjJoAXAB\u0026bbg=Chdjb20ubGVub3ZvLmFueXNoYXJlLmdwcxIkZjAyZTc4YjAtZDRkNy00M2EyLWFiMTItNTAxZmQ2NTY3ZmQ2GKGByYmn9N38QgiETMyNTM1ODAwMDE3MzI3MzYwKhI5MDIzNDAzMzcxNTMzODA2MDhiBDExMjJoAXAB\u0026pixel_id=902340337153380608\",\"clicktrackers\":[\"https://api.gov-img.site/Ad/AdxEvent?sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=click\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158\",\"https://api.bytegle.site/ad/trackingclk?sign=B05kfnJZENNf2e894QEaCsVVPdPyjSqlcL5hgb\u0026click_flag=0\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026target_id=\u0026action_type=\u0026click_type=\u0026click_prop=\u0026timestamp_ms=\u0026pkg_name=com.lenovo.anyshare.gps\u0026tgt_pkg_name=\u0026sdk_vc=0\u0026disp_ts=1668478158\u0026ad_id=32535800017327360\u0026adset_id=42535800016573696\u0026sid=644841349428953249\u0026series_id=52357522399597824\u0026org_id=22201945363648768\u0026place_id=com.lenovo.anyshare.gps_ad:layer_p_lr1\u0026account_id=112742\u0026placement_id=com.lenovo.anyshare.gps_ad:layer_p_lr1\u0026ad_type=1\u0026adx_type=1\u0026slot=ad:layer_p_lr1\u0026ad_style=105\u0026bt=1\u0026cf=143-11\u0026os=android\u0026country=uz\u0026pro=eyJhZ2UiOiIwIiwiY2l0eSI6IiIsImNvdW50cnkiOiJ1eiIsImdlbmRlciI6IjAiLCJpc3AiOiJ1ei1tdHMiLCJsYW4iOiJydSIsIm1vZGVsIjoiMjIwMTExN1RHIiwibmV0Ijoid2lmaSIsIm85kR0RMLcyOTe0shEWETe5r7YwtzORHRE5OpEWETe5r7syuREdERiYlxgRFhE5ujC6MpEdEREWETsytzI3uREdESw0sLe2tJE9B\u0026pp=eyJhYmZsYWdzIjoiMTY0MjhfYmFzZV8ZFhkZmJicr7K8HBfMCwxODUxOF8yvDgvmBYZGRgamy9YmFzZV8YlhkZHBmar7CwmBEWETCyL7CxMzYws7mRHRERFhE4NLwyti9aWQiOiI5MDIzNDAzMzcxNTMzODA2MDgifB\u0026source_id=1122\u0026source_type=1\u0026access_type=1\u0026identity_type=0\u0026publisher_type=0\u0026preview=0\u0026src_region=hk\u0026sign_end=end\"]},\"imptrackers\":[\"https://api.gov-img.site/Ad/AdxEvent?sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=impression\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158\",\"https://api.bytegle.site/ad/trackingimp?sign=B05kfnJZENNf2e894QEaCsVVPdPyjSqlcL5hgb\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026timestamp_ms=\u0026pkg_name=com.lenovo.anyshare.gps\u0026tgt_pkg_name=\u0026sdk_vc=0\u0026disp_ts=1668478158\u0026ad_id=32535800017327360\u0026adset_id=42535800016573696\u0026sid=644841349428953249\u0026series_id=52357522399597824\u0026org_id=22201945363648768\u0026place_id=com.lenovo.anyshare.gps_ad:layer_p_lr1\u0026account_id=112742\u0026placement_id=com.lenovo.anyshare.gps_ad:layer_p_lr1\u0026ad_type=1\u0026adx_type=1\u0026slot=ad:layer_p_lr1\u0026ad_style=105\u0026bt=1\u0026cf=143-11\u0026os=android\u0026country=uz\u0026pro=eyJhZ2UiOiIwIiwiY2l0eSI6IiIsImNvdW50cnkiOiJ1eiIsImdlbmRlciI6IjAiLCJpc3AiOiJ1ei1tdHMiLCJsYW4iOiJydSIsIm1vZGVsIjoiMjIwMTExN1RHIiwibmV0Ijoid2lmaSIsIm85kR0RMLcyOTe0shEWETe5r7YwtzORHRE5OpEWETe5r7syuREdERiYlxgRFhE5ujC6MpEdEREWETsytzI3uREdESw0sLe2tJE9B\u0026pp=eyJhYmZsYWdzIjoiMTY0MjhfYmFzZV8ZFhkZmJicr7K8HBfMCwxODUxOF8yvDgvmBYZGRgamy9YmFzZV8YlhkZHBmar7CwmBEWETCyL7CxMzYws7mRHRERFhE4NLwyti9aWQiOiI5MDIzNDAzMzcxNTMzODA2MDgifB\u0026source_id=1122\u0026source_type=1\u0026access_type=1\u0026identity_type=0\u0026publisher_type=0\u0026preview=0\u0026src_region=hk\u0026sign_end=end\"]}}",
								"adid":"32535800017327360",
								"adomain":[
									"news-sphere.com"
								],
								"crid":"8ef0e46c217b9d4b27008379e161e4a1_1",
								"ext":{
			
								}
							}
						]
					}
				],
				"cur":"USD",
				"ext":null,
				"vaststr":"",
				"bidid":"",
				"expire_time":0
			}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//pubnative
	r.HandleFunc("/mock/868/get_pubnative_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "DspId":100006,
				  "WinPrice":0.303,
				  "DspName":"pubnative",
				  "FloorPrice":0.1,
				  "ValidImpDuration":0,
				  "BidPriceCoefficient":0,
				  "id":"c4082554-f18e-41d3-bc9e-d14de375d9eb",
				  "seatbid":[
					{
					  "bid":[
						{
						  "id":"ec717f104fa34b0eba209248a82dd150",
						  "impid":"1",
						  "price":0.303,
						  "nurl":"https://got.pubnative.net/dspv1/winnotice?ap={AUCTION_PRICE}\u0026t=-x8jOfOps7jVYQ2VzD80OkRmSiLTvYT6i74ivCtgC2YtnX-UibErikoda-QZXTOSkDMMv2FnMZbENsgGxHS9zYaVYKeRgG27Wj8yJRs4ma5H6_rEFyKX_u-gwIP7lNj4o_kWa38OMI2NNko7lbrExm8iVACiVOhMBovF4tXv_W1vPt1-PExiyA",
						  "adm":"{\"native\":{\"assets\":[{\"id\":2,\"required\":1,\"img\":{\"url\":\"https://zem.outbrainimg.com/p/srv/sha/57/2d/fe/185558655c379bea6384b62d614a195406.jpg?w=250\u0026h=250\u0026fit=crop\u0026crop=center\u0026fm=jpg\",\"w\":250,\"h\":250,\"type\":1}},{\"id\":3,\"required\":1,\"title\":{\"text\":\"GB WhatsApp downloads review\"}},{\"id\":4,\"required\":1,\"data\":{\"label\":\"description\",\"value\":\"GB WhatsApp, features and downloads review 2022\"}},{\"id\":5,\"required\":1,\"data\":{\"label\":\"cta\",\"value\":\"Download\"}},{\"id\":1,\"required\":1,\"img\":{\"url\":\"https://zem.outbrainimg.com/p/srv/sha/10/ff/7f/f1bff817cf934fddba1e3644319e487a5d.jpg?w=1200\u0026h=627\u0026fit=crop\u0026crop=top\u0026fm=jpg\",\"w\":1200,\"h\":627,\"type\":3}}],\"link\":{\"url\":\"https://r1-usc1.zemanta.com/rp2/b1_pubnative/25083954/86804516/M3HPLHBNUO2IR7JGRNOVPHK53X4SYO2NJDOENFTNZ4J7YTZKS3E5BB2RNQMYDMF75UC56NMHS4KW2IFVXDEM3RWTFDV43SPTHPDTHPW2XBKKD6MGSXGA25GEEESRL762TV3WKJVWGHC3QVYKZ65K6FQPWY5IRUWCUER6NUYDG43NTEAXGS7MDVO4UXNM73ACDODRDP5XC2OUX54RPQRT6XT5JGBZCOAYO7LRXZ3QG3EKXR6HKOI3FGF543O4NNDUBKUUJGEEB3OWEWTUV3OEBOFBW4UYJMOM32JOMUYGEG2BQNEM22ZGZQTF7ISF67AETOFN2INJQG2CJL4TKOGWBHLPXE5OV7J2LEWNDPHMAUZQHKUVNW5TKYFQLGC7JHHZBWBBS6HG6KWLV2ZOLCKU2QXSHABPBJZZMSM6GDCPLE6MIEPFDRCHERY3ZMXAU7EC7CW4LVBDDWKQKJHRBRLCYVRX3CNWBBMN37J347IKAKM6HN7NUS43DFOBEJ5CN24CVMQEFHXL2EXMBG2Q44LQEIBQNB4DU4YVCE2TVKFFTN4GN7SYV2VZWPZ2UUJV2PNJVLPJBFKKDGXTZC7ZKSEJLX4JHWOO6YZ5K6CFEPFJHNZCPGA7F2KLIMVGZFUI44HHIA7H7F3BNTOJERD5NGYM54R7AO6RI627FYTWP5JAM4RGFLXQXNCPWCAT7N7SWNDWHFSGGMIHASJZR52YSB5P6K7J6RFFNVI3RRAEVX6GQ3JGYTK6UJWO2VCYIEFJ5KD33QNOPF3A7ZJ72F5SOZH6E7IXTTTTEJLWRQE7SXA/\",\"clicktrackers\":[\"https://got.pubnative.net/click/rtb?aid=1157065\u0026t=b5T4FR0RkJBFz_71ICekb9keLSMqBeqlKhS5frPcLKCBO6R4YUiw98wBEHItXdrlO6DiYqDOrVIJvxZP4gSr6r92zDkKg5nvBdJO-UIpYqq1sD5A4JZ1SJfFWJD4Lwcrw4fK4c0GVdCuOGGFopH8fXjvD2Cn-UyND_ZF1mmOAqdH7u4k35ule_tqAgOBSme40LjkgFWQreZhf95ob-RNnO70MKAaZQsuRVN4pWI2b5RTvmS9pGhbTpv3NJEecYpkcT0VsCMT281y1Lzs7fyZ8LwcBN7RepxRz4-9crtK5hymNZbzYG1OB6XWKM514kjDuprI3qPgqT9q8EgnXyxJ9DejqOSCCxPgK_LCpGnjRTU72Mp49wZAl4bNufi5jLi5R-EuzzJdsNsTALmcONLSVSrKXD4BCr-zDW5H2e8PA_Axf0LsuNS-wL0JoU8qj-HW5e0SBzveDr6ewr09jQVVTOWdoF_tQUQTqSqANMTbyCz3Z345XPMaODM2x0NOQFIt3PrMdMSR1X27tfLM2pvOLbzDbNF__pX1LtDSNFjCJTyInAc5Gb92354iaRP2Kfwp6RwblVKf7bsD0Hxj_9capRee971aWfC4CqkauD9O3UAbi07yLtFNDFCxeM7_vWBE_19wJ80QMnDRzKwNkjrFFw32DQUxMa3iA444F8cOeNY8OentHP6oahEqLNOmKNkmSBWVFgxOeTx7yDQhTnCYMpd_aPsqn8UPbPy8Q1-op_laJjmGsSISVxKH1qlqtB8WSGW-JIRnjLVDMl8aU2ydqUFaqsYrfehIhpLnWnyx6EU7xJMOlL04EZ-VcnFvGo7gLYuZVOHTHm-oxYPF4-OJikVLSs5ZWjQzsyoAtZEz_0G1VdeGiQuRkLY5tPig9MZW1GQyOf43t3fD_Cnr8IW_FpIcFF4OqyRWN51xlR43rEyMpha6QvuhAwZbJpuX7ovknx6tNPGofQTHADCo5_N8HGv8607ToJQrQ3uHWJ4sPQZaoCqzuHkYh6qa5WsUBieR7o1wcDp_tAR70KW2JHVxy9oJWP5cr7gU_3ATbormNpMuSmY69nGdpiTOClwqaamczTzuLKInJjdzKNUL1myRlGFEdZvkYzyBL4xOEMr_6b5lYtb92LvhB4PuAIfG365nm8BpOYoa6TwthuyCCz8vIVUOveP2rKW67rWQag72nMzSqnfHjwAAuWvQmxSWvh5Y0HtlNWRIklo2YJZOyX4LInm4XhyEWDWrFKKVem7k9d85DVkAD2ftnQPXn7xn0qIy3l7yjsA8czZf-6lK4xx60SuF5wCMgEwXLy_t1VNig_0DMJjClDpGSTvmn5u-25pmC8erBM91CmJN_sYt7bJzEKVKq0Hzqina1QBJgNWb5J1oBbYOYPtXiW7g_w89E9MdZ7e2sVrgF3qAsIS-k0lviXEqA1Uq9m6sD2iCbD7svvJowrQgzbZpiJh4I6lRAjXhVR7OXlfkb1nb1LLbUbDx5d-f5c72jhpUEQP7a45eyCYLTy9FuymeI4wwDMVmbasSm0wqbdcdarZLd0mVNJXv87z07dGi3L9XfagZk_tOoikMUwcFIlR-Arfiv6cdOz7mNe9O_KbOAy2ldvIIzDajnnYoS1rmTPjeWoFWwVJg4tkNRsOxTIiWeg39dGN_nmkMGv5Uje7wp8dExwoiVg8MVjrsdzSP7mpwhtHLzb0A_D99hvFmA33bztUY8-xm97IaARm2kYzrTAhGl0EEGlCedAZKCpd1J6aRPMF19RZAnBug9OCv_UgFj7CTKEAFj6myDk9EQLPDd6rsc3iBgN5tg6DKhY9rumtJE_w1tSZpjkNg4we-E5khVntlE7XZEtvy3CJG687iZQUQzg_IYqOlDqpM5aNwJtI-Mt7f4l-E3MBIhB9SeHL3nk3JGbUENGdwHBeSjwLVrpu7q2ANMHKfhUlQHlJliQd0LRl8sZi6oSJNtAlXXwzzH0potYB7g0XCoe9kbhAT-NYBJkLqimBShqXslXUNk6OAGOEPM1w\"]},\"imptrackers\":[\"https://got.pubnative.net/impression?aid=1157065\u0026t=VFiBBbWWOxDdJFoslbuBWucNoSsPi-AN_ryqoG6FCS7c-zu1YD_tnz4oBXPNWwRkBKEVxP_0uhrYL3S9ClaMAwjBFhYn2KUQvEhp404swgW4-cJBuwQ4ykl1IXQSNzOeAkET3Og4g9ofeuFyViQlrcOUvi1WDGqRORpzbKhqPmjiZrHawEGXGGmWqt0jnePKaeICXLYmcQik0HekytlB8UMoybF45vkVUFsypVI-2skxssLRzMPkfrP5XN-Cx2R2-SHWsRJodQBJjPrnN6N46h7_b1sB9qqHbdGgw4Upz1nywog5n_N4OPHTP4lEWqDKYBap3DR07ipTxP6lRpWcA1A2ifG41r_bklqIVCfXOAQLDqgCqA9LF1flpczqWSKoIrlYeZMA0EuD2XeFbl6ibMioXodBD7xIqPuAYuYaTuEEzFdAZxC5Tfrwyf73XBuztbrKeFuPqnTQhS1_UpFwOYjfjDgfY2y6vJn9oEYe8grag99bUCmTMXd7FgcLKItJO_P8gltyrWbXeaYS7QYYatW8u8smA7cDrtw8A1O43v1Mt38826C4ANy46e17TyJ7HQ9L0tiGqbM6Hr-tRaYTTe9aH91UT_fqi1WBLIDfXyToGnYsniiHYAC7ykxT59Qy-n7FN6G5TiMvZqrC_fJuBGJrBd-lpKyAn6Ppq7Y0vhpzoh7gu1EaMx8KT87tZuSGJHB5EVpGqQs4_JXw1a2BWCL4XrDorKSwek9FCbPVSOkRcP_L8FOpjPaTjjq1KSVdjBJAJSmHmaWtp32N4GOIa9jZovtDkNMvuSkXmfdrDEez4P5JDldJT-FbKe94Ru2-aTHOuO19ZK9wxQQutUm0CdIIyageIaS5T2Qyh5pDXdyg8VHJcYZ7uVew0UFFM4osb3tZLmiht-w0pXb7re-pcAPtdYQpq2QhAeVM--mJcJdx2UFry6CtptXkyGm8zwdjHpby9BLF4zZ8ZpGFcJCNj7Xr2DGYPwoZ54GEL1CHwx7QG9HZx2I208x8q3qeYVi8qr4mxVfs80oAnLlsckk4V90jheStNwYz_rouc0Y59LoHbr05aJACGPYasEs97zmSwhXk2SPUHPE51tWZP10Loy_pamgMDHnYtuCGL1faN9LQRgj1HVEM0nfhyrhmYSX5izFuh4Z6KqNgyzEboFm1v4CkGVBDnSJeRF-xep4d0nFj6NjCEX0KbKfh_gv1mG8NIWd_QyIj_GCSmkauhXhZ7mRrrA8uA1f_IiXoZeTFGO8qArUY7L4rr3omk47eMT2Vr51YfOlPnkgBg-SNMi231I_DnbJ2pFhAUDGwei3vw2JHUYph-p8WtJxl0hmt2q9Gw1EJs8-v6soBnv3A0rmIqQrbCF2s6gPAlZ2qSAG0jLyha-eCFVNEgXbPTpJFqXTbnNInXxCz2uLRYO8PSms6jQcbRMvsHCC5mQA9KwLb7YN0uAWPooSEzpBENZot4ERnJuNBhEM8lQ6Z5O9j_97jzbNtGNNApiB9EZo4q5Sf7yAc0RrvrZQWIbHBBgTMjOzLGyBKOWKgesmIOczlcW_uTBS6GllmW2pv3Wo10RP2loT1gRzzbidZ9q6PMoS95oUHetzK_2nXWgaCUIEm2vRJJvJqGzdzq1Rb0uiWaAHJLPjPIrOetNpg3cAL_XcArSSMTnezh8fPEeK2lK-iJfdkJylj2C5WqUiz1VRd_jOv6BpIN0HclS1vo7olb-x2eppHh6UHreHNhYjy8fEYcjCoaC253T0-WuXgaOVtN12g8Qu219jq2He2NWwZCrVZhY0iKVYeaxXHR_bjweRBlrKFKbNLt3OPyQWzUGhR9H2zYlGeiaN4ExVXkbF_QVf7G__MqWKRwV3rkbn1FOD5WhqviIG6JIMRTIDtQjLaepAJhFs7xNxwGo40G5b5WfFMP45NbJw1otYdcc0cg_Xq0MfPyVwvwQfYSLeb3BBaiZpEbtQooCcS_fh87WUaM5F7MkcEaKPiRC-ZukXKxIuqjZ6G9yPPg61SiD-xmAABWagqy6clCSXX-IadZbjbnT0sN14IaMLUKUklUOok1NSEsi4ipyWFoSgBq_Tf843Ja1K3ZQ93fIb1rqz0DoBbs8amCQnL4_2hm4gEjLAR7e3hA4PxzlNUzQr8UfAGBZKf_ggrbcAoVQJwNCU8oNDtsDC8QwGvTNVPuw4M1sbI9YPVHExD03fcjriJ3Vt0K7y1XTPDPv-GEvI-yfpi7EC3yEaKfH5EHtmCaecUXk43M30Gu0W0T18_yFWCwsE79cJ2Z5SoHX4ph9nggz3l7aXQnBtCdI57BvHdY8x-ZvKOIIE2cJzTCOWpskjJS3QutPUVizBQ83faG3716Q34pQT5A2UX5znUuGGO6D9RnqK6uB3A3dEtNZurnGmxZQ\u0026px=1\u0026ap={AUCTION_PRICE}\",\"https://b1t-sindc1.zemanta.com/t/imp/impression/Y7CSMQCVDWSNHCMK3HAMPXSIUGP7PJOMV53JMNHKO4ZGQ7QDQJWW4NOL5YNAU6P2KPSB3QXAKIAT2Y7WASJN47R3525OW33SQ7TH2NGELB3OWHSAXCCPA4KKQPIHTN7XY54D6BM56RWDJ2ZGBNWHQ5B67AY3H473LWROZZ3OWD5AQU2GZXJQJCU5HFIBD27AHLLXGNNJICG7EVWRC26O6R4BIPBZEHUDQEYW4TRIM5MPHN6P7WYAMPWELL5GTOXCAWWLN3Y6XDK3HX3OPRMEMEX33EV3KY6TNN47J7WQVIH5DA62J4K3EJEXA2KNPJUEGQTQK4X7FPVLQL3YQ4BRU44HD5PYBJRSYFDXLIUP7N5ZNKLSE356L7NOEOOZ4FT4/?\"]}}\n",
						  "adid":"244934316",
						  "adomain":[
							"baseplay.co"
						  ],
						  "cid":"25083954",
						  "crid":"86804516",
						  "cat":[
							"IAB10-2"
						  ]
						}
					  ]
					}
				  ],
				  "cur":"USD",
				  "ext":null,
				  "vaststr":"",
				  "bidid":"",
				  "expire_time":0
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//smadex
	r.HandleFunc("/mock/868/get_smadex_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "DspId":100022,
				  "WinPrice":0.16249,
				  "DspName":"smadex",
				  "FloorPrice":0.01,
				  "ValidImpDuration":0,
				  "BidPriceCoefficient":0,
				  "id":"64a2648d-3f5e-45ed-877a-4b7534b961ad",
				  "seatbid":[
					{
					  "bid":[
						{
						  "id":"64a2648d-3f5e-45ed-877a-4b7534b961ad",
						  "impid":"1",
						  "price":0.16249,
						  "nurl":"https://geo-tracker.smadex.com/hyperad/rtb/163469/bidwon/14-H4sIAAAAAAAAAJVRPWgUQRTePI4j6BEjWqQRUkS7WefNzuzeHsZLCkPAH4KnKdIcb_dmcot3u8vuXiSNpVhIsBAVWy3Eyiq2FhYWFmIjQRAR7SVYiCg6F6PYOvBmvpl5P9_3Hlz_cAQcBxopVcmGtmj3xwSco5yV2ajqayorhq7gfFi6PPZICS59qfwwRF_DKg3WsyKp-kPW39zQLNeFyYohpbFmcU6sHEVlXCR5lWTp_zk_2rnxtOHA9rs6nH79ZXtlabftwOHZ1bnOR962hMfrcxsAHWsAUIP6451PE05tbRGOjvJBYqouDQbwMJ8_8ObB1_YfgVBDZUWuaCml6kUR7xkjPe1hGCH5XMacm2bIYx1TLJoUOzCDXLhChq5ymx7UdApTHRqWo3Qdpjvn2aLiagmmKO0VWdKDBqKLLodlHiguSGumAy9kstkUjELJmTHGE9L0BMkAjpZDKqq8n6UaHDhV_s7bLYddsmlNd0MX2LV9oRSRW9rOvnILV41oGe33TBNDYWn76MnA09YQFTc-RTwgiXHgBZGNA5gZlX0qdFK5cTaE8eYOdJptZC6lm3tf7npewr0XLyccqCM2bcxUZ3nx4pmkgoXQj7RRZDj6oZLkCV_FKoqIhzISiD7MUq81oE1ddPPuMM8jhEaSdinPt97ffH5ijzrc2arB9yfH7AzQ3r7dPgmvDv2dZQN9T_rhuKqvFOOCw9Tl9EqaXU3_AWDfhWCIDNUl7rUEtjzPxUCtLYcUeaEdFvNlKK2H7rEwiCOGhpCIowyiJhwfwf3Jt_PPYO7uz1vXDrYXoH527cKYBExKztGe02M8ruLskd1v-C-1IOi5KQMAAA/1/9ab39102-6494-11ed-97cb-1fa1aa0147b8/{AUCTION_PRICE}/notify?cid=240702",
						  "adm":"{\"native\":{\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://static-content-1.smadex.com/uploads/banners/8e812906a5d42149f3781fef9a429e5223eb45b5b4d814ed641a247bc183e4f8.png\",\"w\":1200,\"h\":627}},{\"id\":2,\"img\":{\"url\":\"https://static-content-1.smadex.com/uploads/banners/e0874c445bc145bd0ff36d79481de8370215e01088af9258a9c65bbb471b0dc4.png\",\"w\":80,\"h\":80}},{\"id\":3,\"title\":{\"text\":\"Upgrade WhatsApp\"}},{\"id\":4,\"data\":{\"value\":\"Get the latest Upgrade for WhatsApp\"}},{\"id\":5,\"data\":{\"value\":\"Upgrade NOW\"}}],\"link\":{\"url\":\"http://optin.telkomsdp.co.za/service/13?direct=true\u0026gv_id=326\u0026cid=600004\u0026utm_medium=web\u0026utm_source=SMDX\u0026utm_campaign=MyStatus\u0026utm_channel=Main\u0026utm_term=GenWhats\u0026utm_action=UpgrArrow\u0026click_id=9ab39102-6494-11ed-97cb-1fa1aa0147b8\",\"clicktrackers\":[\"https://geo-tracker.smadex.com/ct?q=9d2d71c060b48850936fdacb443493a3ef20d47fbdcbe8a640c9f75f3d3811f9104c966d6875a8ce6ac3d7e5387b0fe8d062151c908838983613ee3fcc869ae1607af941c02138fec6f65ed4e565c7699fa7d169b16c02080530bb4a188deab294f6b34b01ee19ab085c6e4d4c1c64ed462ef1dd8255f77d73a5aff37c3d322ee7564ac780276e8a058e0ae3a4c7b8d31ac4ea628ab8a6663e8dca0d7ab1341e31d996c43fb73ecd8454cdf942b5e583ad27d7fa2d504ba568436118506aa26882a373c085e3d7b5a6c6fc8d7d3e9c4f\"]},\"imptrackers\":[\"https://geo-tracker.smadex.com/hyperad/rtb/163469/impression/14-H4sIAAAAAAAAAJVRPWgUQRTePI4j6BEjWqQRUkS7WefNzuzeHsZLCkPAH4KnKdIcb_dmcot3u8vuXiSNpVhIsBAVWy3Eyiq2FhYWFmIjQRAR7SVYiCg6F6PYOvBmvpl5P9_3Hlz_cAQcBxopVcmGtmj3xwSco5yV2ajqayorhq7gfFi6PPZICS59qfwwRF_DKg3WsyKp-kPW39zQLNeFyYohpbFmcU6sHEVlXCR5lWTp_zk_2rnxtOHA9rs6nH79ZXtlabftwOHZ1bnOR962hMfrcxsAHWsAUIP6451PE05tbRGOjvJBYqouDQbwMJ8_8ObB1_YfgVBDZUWuaCml6kUR7xkjPe1hGCH5XMacm2bIYx1TLJoUOzCDXLhChq5ymx7UdApTHRqWo3Qdpjvn2aLiagmmKO0VWdKDBqKLLodlHiguSGumAy9kstkUjELJmTHGE9L0BMkAjpZDKqq8n6UaHDhV_s7bLYddsmlNd0MX2LV9oRSRW9rOvnILV41oGe33TBNDYWn76MnA09YQFTc-RTwgiXHgBZGNA5gZlX0qdFK5cTaE8eYOdJptZC6lm3tf7npewr0XLyccqCM2bcxUZ3nx4pmkgoXQj7RRZDj6oZLkCV_FKoqIhzISiD7MUq81oE1ddPPuMM8jhEaSdinPt97ffH5ijzrc2arB9yfH7AzQ3r7dPgmvDv2dZQN9T_rhuKqvFOOCw9Tl9EqaXU3_AWDfhWCIDNUl7rUEtjzPxUCtLYcUeaEdFvNlKK2H7rEwiCOGhpCIowyiJhwfwf3Jt_PPYO7uz1vXDrYXoH527cKYBExKztGe02M8ruLskd1v-C-1IOi5KQMAAA/1/9ab39102-6494-11ed-97cb-1fa1aa0147b8/{AUCTION_PRICE}/notify?cid=240702\"],\"privacy\":\"https://smadex.com/data-subject-rights-policy\"}}",
						  "adid":"9ab39102-6494-11ed-97cb-1fa1aa0147b8",
						  "adomain":[
							"telkomsdp.co"
						  ],
						  "iurl":"https://static-content-1.smadex.com/uploads/banners/8e812906a5d42149f3781fef9a429e5223eb45b5b4d814ed641a247bc183e4f8.png",
						  "cid":"240702",
						  "crid":"1881685",
						  "cat":[
							"IAB9-30"
						  ]
						}
					  ],
					  "seat":"161921"
					}
				  ],
				  "cur":"",
				  "ext":null,
				  "vaststr":"",
				  "bidid":"9ab39102-6494-11ed-97cb-1fa1aa0147b8",
				  "expire_time":0
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//mytarget
	r.HandleFunc("/mock/868/get_mytarget_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "DspId":100002,
				  "WinPrice":0.28932,
				  "DspName":"mytarget",
				  "FloorPrice":0.1,
				  "ValidImpDuration":0,
				  "BidPriceCoefficient":0,
				  "id":"58683a27-844b-4c29-9fda-622ac3f57a6b",
				  "seatbid":[
					{
					  "bid":[
						{
						  "id":"2:99:63730d15277d244f",
						  "impid":"1",
						  "price":21.12,
						  "nurl":"https://rs.mail.ru/pixel/AADQ9AGRRayx2VutJ2toKR9ObHMFAjk-5Eeq-DJJCHYF7N1c3zk7y0VmTd4AzhOR41SVBseelKXl_WBwpCSMCbXB2ZwRkVXwrcBcAeHHH8xkeldnV4T4RUhoTzJszcNC3RzT0Q9zLuBdLLsI2op2y6Ph3dQ1JjehVoaSx5-jV7Oar7ZhKy8qkBF0AQAAz5_7A9kUZ3cUGIk_K2dXqLXgWccrtRRSlEaJl6riCX18AXN7cT7qfn87WctSla4tqN0Vz7WeeoTj61x4YwG6lQ3tKvVMh4Tqbhtmfib0Tndt_1iwB77phydlHjxFRYlphgxEaKxN2NaIMOSSRsZwgyCvXEUV6nxAor5xnl3k-PJxutugRNu0t8n6BbkPfnEzxngQytb9ZnBlvf6vlFAYZ5zJXoLZXGdWhRR46r1_Wn3gNvkF5Hj-VQrSoUduv01UR92edZMOf-cmdCyq9unpEfB4FtXKFYR-r-0CnzpILOTOA6wvgwSuo7IhZtqcWHmtfdXPkHEZIwzr97_RUSfTFJrvICzHAVNIGjYtR4fHELW9f2veoaCLX3QMZ674Ux-pG65ekIEclsH99_8ClW6YZMqzcnHx20xFflHgQ7CcMajXfj0FQG89F3vyLdAza_7rA7hEmbt163ilKMbTruhVeJWydM5c-tEGKvTIB9kFuXP8-9RH85esE7ulKf8t_wF-.gif",
						  "burl":"https://rs.mail.ru/pixel/AADQ9AGRRayx2VutJ2toKR9ObHMFAjk-5Eeq-DJJCHYF7N1c3zk7y0VmTd4AzhOR41SVBseelKXl_WBwpCSMCbXB2ZwRkVXwrcBcAeHHH8xkeldnV4T4RUhoTzJszcNC3RzT0Q9zLuBdLLsI2op2y6Ph3dQ1JjehVoaSx5-jV7Oar7ZhK-moYRN0AQAAtqT7P9O8wVUgkyAVq1_dQ3ElAd0lxTPfEEmipFPOFbEmunpmnBS5maRh5YEASt1QHlHolDsHboEU9F42Ut6C04I_scT3AGaDuGUP-DXfDL648LU8E1hqeqNGAwrH7Bmuirwi6t3bU8W5WCioML08xa0nlgrZqaTpYXKAzOzx8SfI6Qztz384SV-2BjYLGZoCDu_eEhTqNOV8ubBXVE3LguriGv1hENmtFx6l-8nASN4wpjmp40ImXTpZowfx5pKHT1nIgn7n289QIpESQsKP1qoEjhtMMpMldzq8iAlLtIxQCFInu1YJVttB6ZLf7teJtAvRDQGboDY0OE6I0WljaDr4mHsTbk5F8o0txJM3T2a-viQG_x5x--2OLQiB4-jGAHbjvO0B-sFMYxl2udo21f2XCVGogzIiVV8akJDIO_HcGPNX_c5HyKck6NXwsurcTdq6lRJ49JsSBBmNwr0lX5zFz04u-GWNd6D0OSSp6k-tbl3n5MBba_QIO4IRRJRl.gif",
						  "lurl":"https://rs.mail.ru/pixel/AADQ9AGRRayx2VutJ2toKR9ObHMFAjk-5Eeq-DJJCHYF7N1c3zk7y0VmTd4AzhOR41SVBseelKXl_WBwpCSMCbXB2ZwRkVXwrcBcAeHHH8xkeldnV4T4RUhoTzJszcNC3RzT0Q9zLuBdLLsI2op2y6Ph3dQ1JjehVoaSx5-jV7Oar7ZhK5i2KiPWAQAABLlu5xA_7Uj2Q6hCpJM7orgUrIPvxV9tu9VzJgYKAsTjIOoobJ4X38hMqWlojgmgKzl5NsQ41u6jHwkUERT09cHObcRt8mcTKbvHE1vDJ3SH-fyxYdNb4lhRs9FMHyqy2dd_RUDECZwL4HmXXiI5heS3bIqfC1EAbX1kxsrVp3RdVBExCZJGSor_GEgqmS85Cqwu5enW3ZfpR4CGJaKVNNtCikMIDvrH-NZ2Ynzpd-F0F8rO5ukzUMo4ggKHZmzk52_l2xR6D6m5xjB18B8hkSuUGEF81-4o99wFTHWY-wOub0172rOHyRPjVS2iM5FrQXQ6_J2tomRAAmIz6PDXJMGVkE1UbsPfrObklK2xRzJAIf6CSJ9fiqXW61f_IGR9yhH0j09_ZzVMMtK-y07mBmBSNcXHDE-54d19sr8BCBWhycOxVcF0Z0hqWiJ2cHGkcw-Qb4qFzzVfQf0uI9JX1aTkgddmn_FnJVJ-cg35c0-hg7ZWqpIj-F61DPE9vxawZrwrJA6_79MBw85o25B4f4b7tqOpVqIWQkDPdXO2jAFMV4cLazZTHCDgDUaz9Gs70eiIJ2I1SFEpph6qiSz6Gk0x3j7e3312e1qQ1nTsDXTrianZL8JMwX0CFUh7bChS.gif",
						  "adm":"{\"native\":{\"ver\":\"1.1\",\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://cdn.liftoff-creatives.io/customers/c140390104/image/lambda_jpg_89/9053bdb427c0cff599e6.jpg\",\"w\":1200,\"h\":627}},{\"id\":2,\"img\":{\"url\":\"https://cdn.liftoff-creatives.io/customers/c140390104/image/lambda_jpg_89/825e0f98c1d052bac358.jpg\",\"w\":150,\"h\":150}},{\"id\":3,\"title\":{\"text\":\"YAPPY\"}},{\"id\":4,\"data\":{\"value\":\"Каждый день новые видео!\"}},{\"id\":5,\"data\":{\"value\":\"УСТАНОВИТЬ\"}}],\"link\":{\"url\":\"https://click.liftoff.io/v1/campaign_click/HIHtBwc0VkWS6_xOS6O53_VHsRnFcQN4N148seq45AzNFbbod9lq4mPpbT-fn_Fvmyk77kZpZJ0AAshtpJtOajpSkUxP96sLpMRqRSMJN3NfOIcvQ_4Q4cgvLz_t6tjES4M3bgA3T6kpOfSjmwOPykfPJz2FJIW7orvZk2sEyrYAJ8dXCs3wIYMsAJlT-C85-k7ewbNrZq5mi_OROWLDNwQZdR9MJuHsb7CTbG5Qqtf8jrQtrqAuqrpY2FQhb-MWrYJclP76Z4_NX4nEG56ktrls9DUq10ysb1wKFrSynXJMf1bhEIRnTc_hRXnQMxyerhUUUUyJNj1QZXjhGnPesD6x87YepkqaeBHpiW1hHRZ6JXErBHXWWuro2j_nRTEJVOQEUoCcFUHCNhAbqkPuIijszNgmnGD0GspUK6QeJtRCffEPDORLjbO7ewLfOcjTDOV17Iwzjn8-GtVz27DGJR3836SAxGNNp5G_C63w8_r-qw9rdQcDYgpjaAQmz5KoNgaf2nMGPH9Gq-MS2qTEbABNTkLdRUS5o9Y6TLowIBJXocpywwDO8DGCx-8K8qGz02VsN3mWdsJl3NMKgQYoEMM9Xqv3VIVksySNUQWpuyeTop30xNLMjoz88g\"},\"imptrackers\":[\"https://rs.mail.ru/pixel/AADQ9AGRRayx2VutJ2toKR9ObHMFAjk-5Eeq-DJJCHYF7N1c3zk7y0VmTd4AzhOR41SVBseelKXl_WBwpCSMCbXB2ZwRkVXwrcBcAeHHH8xkeldnV4T4RUhoTzJszcNC3RzT0Q9zLuBdLLsI2op2y6Ph3dQ1JjehVoaSx5-jV7Oar7ZhKz3zDm-nAgAACoG-EgvLlHg9KCN4OTU3jp9r0oW4XnKDjNH5boRk37x4_pjqoCSJdF_VkrdRw20A_C4nw7mUZjbb1XBw4o4WE-edlNP0mjZkbO6HsRwL7ZDOo7IiAkrdBFUP-509Zywm1K61wNNNsgM4cBUpM7cLW3zHLtQ_4bToRubb5bCdOgAWELvfLtXB-fhQMchsu14qZ1Kg1gNJX0LcBSxZI1L5g7cAfCkPBiPJrlDPcAJrG0jknCnjVt53Gq0X-malvHEcgPdGbOH_lguRP0cLjs5QN7l9zsagOPlT2Zq38Gqau8F_pXpPax5wl5O7H2b7sSkKt5Brg032uFTM0dTnaGo1bLrx1Xht2-co2cwyDgNAwh3fQ3m32Gek6mE6CXCCvJSMbuGjn_njyowuUcEI3LHRntPTZWLu64Xx39tGO-USSFYDC4J31-Ia_YUHh_JvRU6Z4Nl2UojaeacUrU28_4Eagix-4knbYE43cv_N1sV4abI1Tx6HrR4HQWd8OQvqsggwcZ16oAFF-bpf3HFSTva6-yUFdYmApmRmIMReNyfWSIHlTNHBo95cZuxzIly-h7El83va7YGH67ARP39rQNRiKEnJOkGw_a56s8kWj6BUJvVo0BqAR6pfeMr6GhNMAAX4MdYBmW4Q-IG_jEGK96H3n231wAbDlCacE9xDeN_uJAKzqEk8CWrjaC-eXFYldJP8005HBPYJVYr4PwAKrR-bhh5Q21_22D66WE4pXWxo0llGa2ZTAT6RcY3xH9giOyUBiqBwSlDbSL6hMjq78ZiVHDI1Di4NPKsM9OpR005c7BhgbzU8VkPykcy1Svz-ChVVOzQQnnOnBAw7nXsniJwwYiArQcn81M1IkCVvG-l6_CO3hw5nbKi5iQX1nG3dCQVs0V0aGUrGTLwckqJPHHeOtg.gif\"],\"ext\":{\"bid_sign\":\"\"}}}",
						  "adomain":[
							"ritm.com"
						  ],
						  "bundle":"ru.ritmmedia.yappy",
						  "crid":"99:288465",
						  "cat":[
							"IAB19-29"
						  ]
						}
					  ]
					}
				  ],
				  "cur":"RUB",
				  "ext":null,
				  "vaststr":"",
				  "bidid":"58683a27-844b-4c29-9fda-622ac3f57a6b",
				  "expire_time":0
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//Rtbhouse
	r.HandleFunc("/mock/868/get_rtbhouse_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "DspId":100068,
				  "WinPrice":0.0456473162306253,
				  "DspName":"RTBhouse_APAC",
				  "FloorPrice":0.01,
				  "ValidImpDuration":0,
				  "BidPriceCoefficient":0,
				  "id":"a42b43ef-02f7-470a-9dc6-1ffc0958dad7",
				  "seatbid":[
					{
					  "bid":[
						{
						  "id":"20221115_oM2tcCJTyQIHZ7DR56Qt",
						  "impid":"1",
						  "price":0.0456473162306253,
						  "nurl":"https://sin.creativecdn.com/win-notify?tk=LN96k6sZT69LaxDJlavO9Gcat5U5B9rmSe9IFgiL5wofjHhg3PniiQf55fD8PeJpj3PKY03_ZUoS7C0ZVWITF10lGZhq5J1XitDbbBxVUPkkHJK0VMz5VYMGHYKSfFtgeZyOdEYaFuT2lBqSqtRtmPsfZnX0FAizQNGu3v1YsxfToMp1Qjcc2Vmw17Ief_60C5x-5sIV2W1L7n00XSuqInE0h2W0K299dw-wOwtGL8Ufa3j1ETTCoRaZaRX1sDXe7EYnhjhcpgdonXVu5wPAb9WQj61LvBn0Rm67yCJL6ljLagbe3DBlJWe-9qAlVYcaHMmJhDfjmDNqi2PLuthZU01KWUmGStiMb6Iu2Y2N2n5Z42aKcVFZwIlCJ39PTsb6GyLKu56alXjAeYzQ7dovZHa_BvNEMNNNl1JBk5y9YX9XLbHaR_qx4Wc_iNImhQO0qSFM9Bwmgm_HFhm9RpolbyknhCdesyjX1yk9QcaMNcY\u0026wp={AUCTION_PRICE}\u0026tdc=sin",
						  "adm":"{\"native\":{\"ver\":\"1.1\",\"assets\":[{\"id\":3,\"title\":{\"text\":\"Free Shipping for All\"}},{\"id\":4,\"data\":{\"value\":\"RM22.80 - Ready Stock Men Sandal Outdoor Selipar Lelaki Black Strap Sandals\"}},{\"id\":5,\"data\":{\"value\":\"Buy Now\"}},{\"id\":1,\"img\":{\"url\":\"https://sin.creativecdn.com/images?id=e384ce2963204b8eddf01580b1546a9edb10ecaa\u0026w=1200\u0026h=627\u0026o=9938242930\u0026fid=R8Jr7HKuqEotOpvi3IiJ\",\"w\":1200,\"h\":627}},{\"id\":2,\"img\":{\"url\":\"https://statics.creativecdn.com/7CAbBVMVjPllbWkAJn5e/vn_shopee.png\",\"w\":250,\"h\":250}}],\"link\":{\"url\":\"https://sin.creativecdn.com/clicks?id=20221115_oM2tcCJTyQIHZ7DR56Qt\u0026tk=V6st2ORDbXHk3Jb7BVKM9Xxc-yoWwuMrLg74ByEZy2Vg3fLfx22ZwdSh1M7zjSh1SJxCqd6Y1-Asy_lex80twb1g4VsfcEhJgHOexqTL5xUFfv8r8e4IBoh2KY16cWOuN0LhYH5by2JAclbIPqLkDeTZp7TRkU7uG9ss6fNqj6K63OPJXBWbNQHF4wcBTX6lnx3tYiR6t_FpzofO8cm2J8rWXwjxL7VjYBbpJs2ShAOgLDnTWNxLQ3ku888Jq6YO-AchtTXngJVFmpTwqWWQkr9lyf0GtUJCralHJM8Bl3eEOSilCQ3XrXcrlhCgYuiPHn97EHKesPQuovZyorLUo3VBuSq8oz1D3Gv-TwS50XQwnkiEliRMEwLWD6Nk8ay9Gc83VBLCnFJzqR7Al9T17i4Wpb0i6-mRFsi9t3oY8bXv_Kj5hVEVMa1msDddm-0KiCFG6oGDl_CBwvsbqWKxK3uEeTZLOMHWX1edCvOYmQE\"},\"imptrackers\":[\"https://sin.creativecdn.com/win-notify?tk=LN96k6sZT69LaxDJlavO9Gcat5U5B9rmSe9IFgiL5wofjHhg3PniiQf55fD8PeJpj3PKY03_ZUoS7C0ZVWITF10lGZhq5J1XitDbbBxVUPkkHJK0VMz5VYMGHYKSfFtgeZyOdEYaFuT2lBqSqtRtmPsfZnX0FAizQNGu3v1YsxfToMp1Qjcc2Vmw17Ief_60C5x-5sIV2W1L7n00XSuqInE0h2W0K299dw-wOwtGL8Ufa3j1ETTCoRaZaRX1sDXe7EYnhjhcpgdonXVu5wPAb9WQj61LvBn0Rm67yCJL6ljLagbe3DBlJWe-9qAlVYcaHMmJhDfjmDNqi2PLuthZU01KWUmGStiMb6Iu2Y2N2n5Z42aKcVFZwIlCJ39PTsb6GyLKu56alXjAeYzQ7dovZHa_BvNEMNNNl1JBk5y9YX9XLbHaR_qx4Wc_iNImhQO0qSFM9Bwmgm_HFhm9RpolbyknhCdesyjX1yk9QcaMNcY\u0026wp={AUCTION_PRICE}\u0026tdc=sin\"]}}",
						  "adid":"BzMCIp1WY54hGSn20CJe",
						  "adomain":[
							"shopee.com"
						  ],
						  "cid":"bBEILNzj0Dc94sdiLaAQ",
						  "crid":"BzMCIp1WY54hGSn20CJe"
						}
					  ],
					  "seat":"rtbhouse"
					}
				  ],
				  "cur":"USD",
				  "ext":null,
				  "vaststr":"",
				  "bidid":"a42b43ef-02f7-470a-9dc6-1ffc0958dad7",
				  "expire_time":0
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//smatto
	r.HandleFunc("/mock/868/get_smatto_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "image":null,
				  "richmedia":null,
				  "native":{
					"ver":"1.1",
					"assets":[
					  {
						"id":1,
						"required":0,
						"title":{
						  "text":"NBA",
						  "len":0,
						  "ext":null
						},
						"img":null,
						"video":null,
						"data":null,
						"link":null,
						"ext":null
					  },
					  {
						"id":2,
						"required":0,
						"title":null,
						"img":{
						  "type":1,
						  "url":"https://static-content-1.smadex.com/uploads/banners/14520264786f3fa6c472bcdce8c7b3748317a17a7115480220cd68ad8798bf64.jpeg",
						  "w":80,
						  "h":80,
						  "ext":null
						},
						"video":null,
						"data":null,
						"link":null,
						"ext":null
					  },
					  {
						"id":3,
						"required":0,
						"title":null,
						"img":{
						  "type":3,
						  "url":"https://static-content-1.smadex.com/uploads/banners/bbe4c5c57829f1f7684d69b385586d7765bf2c2c0cb7cff19daa4453fdaacc4f.jpeg",
						  "w":1200,
						  "h":627,
						  "ext":null
						},
						"video":null,
						"data":null,
						"link":null,
						"ext":null
					  },
					  {
						"id":4,
						"required":0,
						"title":null,
						"img":null,
						"video":null,
						"data":{
						  "type":2,
						  "len":0,
						  "value":"Get now the NBA league Pass",
						  "ext":null
						},
						"link":null,
						"ext":null
					  },
					  {
						"id":5,
						"required":0,
						"title":null,
						"img":null,
						"video":null,
						"data":{
						  "type":12,
						  "len":0,
						  "value":"GET IT NOW",
						  "ext":null
						},
						"link":null,
						"ext":null
					  }
					],
					"assetsurl1":"",
					"dcourl":"",
					"link":{
					  "url":"http://za.leaguepass.tv/cheerleaders-sm?ts=SM\u0026refid=SM\u0026click_id=f3a56c49-6496-11ed-a2f5-cfaba5dbd58b",
					  "clicktrackers":[
						"http://ets-ap-southeast-1.track.smaato.net/v1/click?sessionId=a082f090-ee97-3a18-4c9a-9f6778884fdb\u0026adSourceId=c9e7eeca-62e6-adc0-333c-d6c53583df7c\u0026originalRequestTime=1668483501388\u0026e=soma",
						"https://geo-tracker.smadex.com/ct?q=1625e0b1f7d386735771294758a8a97c58e0dd96a0c8f6e869eac40f10506cdeffc4ffc7dfb19d93f0eba79521e68f69f709861b5f39e8104d1c0ad88b306b838b118c883ef16a7418a43d5bd2872acc132e3f0401046577dbd565f5197b42779383812ae0f659a9b7a060dc0714b14fa88e52bdd461811431aaf6a6fc7709909956f1be401e5727dad09eb88cbc0ee669038ab7516175755cadb1ab3901871ced947b59bc58abe7a352d1e06f48a828db94a9d62bf72cd9e2f63badcb525f325428422be64d06a24814edbe816613f7"
					  ],
					  "fallback":"",
					  "ext":null
					},
					"imptrackers":null,
					"jstracker":"",
					"eventtrackers":[
					  {
						"event":1,
						"method":1,
						"url":"https://geo-tracker.smadex.com/hyperad/rtb/14363/impression/14-H4sIAAAAAAAAAG1RO2sUURSeHJYQw7IESaUkf0DvcN-7EyGbSJT1ibqSIs1y7525uxN3HszMLlhaBLSwMFhYaCGInWBnKgliIUGwSCtorSCCVoLg3fgq9NzinAvn-853vgN72wvgeVBPVRWPI1d9-DYL51WOymxUDSJVVoj4FOOk9LG0RkmOm0rYAEdYglbDflbE1SBBeWKyNI1MhZKor6waDrUy15DJFSpHujRFnFdxlsL6_yClRXE6_rf56ZM7z-oe7L08CjuLWx8v73xqe3Do-IPbe8dOtJ3sSXxuA2APgABADaZ3bryZ8mobqzA_yoexrXpOCUygi1tf2r_XhBoRbtVLYSQJsRHmNFBNwUJjsbEqUESzwFBpCZcGB1h5cIQTnzCfcOxTEkAtSqHRVUk5Svsw172AznKCT0NDpWGRxSHMtnziY-gERDeZ1hGSphkhzghGAdYaNU0rtCZkxoYU5stEFVU-yNIIPGiVP2l7ZdLbdKz93jgqSM_5MlIW6t1EqSrzXKP3K2Dd0iXGqZVhKEyAhQwo5oRbIkgUUi6pdq-prdRMO5Mc1DdZ4g-jNBtnvkqvlwNVRH4_L-Heq9dTMHNm9STz4DChmASUUu4wjW5n9cqpuIIVYbhpScfMgiZzZkkiQmedc8ZJaJrA4dyaArewgHqc9lSeP97ff9Q4kAwvbtXg7v0FdwDiftvvd6fg6_epP6ecJZxJNpknhUDrawTqa6NCq_RvBkwxpYgQRMRVzJZYa4kSnwVso2OZEtLwAEkeSNcRhUhRK5C7qVYi1KFoaehsPgyL58u7cDN-97aXLK_A9LmNixMFMMMxJi7PTerJFO9A6y-ffwAp8yxFKwMAAA/1/f3a56c49-6496-11ed-a2f5-cfaba5dbd58b/0.89781/notify?cid=239370",
						"customdata":null,
						"ext":null
					  },
					  {
						"event":1,
						"method":1,
						"url":"http://ets-ap-southeast-1.track.smaato.net/v1/view?sessionId=a082f090-ee97-3a18-4c9a-9f6778884fdb\u0026adSourceId=c9e7eeca-62e6-adc0-333c-d6c53583df7c\u0026originalRequestTime=1668483501388\u0026e=soma\u0026expires=1668487101388\u0026dpid=d2kgqXUzxN60WygogqbGCQ%3D%3D%7CEHJM0Gy5iilNzxjG5ZScUw%3D%3D\u0026winurl=Q6LOu1LWC6z5D2Qc1_LW6YDbvktLP3ONQIyq19r_gm4dJaqk7m38y_GLV-z748fBoU_094zeqC9X3uJl8SMANy5-GMOmlSUluIhNy3YBx4E8wqS384MTmqVVwKeZlnvPgrtIParD6aa4paN949vYLQqsCun4b-PmSvHcAIuPN2ysO3yXrfcc4TOn42SP2x2aUvq-dNMqIMt3TVZ5yhg_XmZTbbg55IomQgIILn9D-GifBE3yQYF2yqtAT6I4G57z9Q1FD7fI46QFfAsJ7ljz47twVCp58L2XuassOfNuYMMMplJxkVdTpsuK05jRj4zN-WtlusnVo9Yz8w-L3D1E45cN1l2C6FqLRggoeISFe2OBwHzBolNfPut3hsOLMYjuEUapHusfoJI4jjTu_Cwzjn26N-JFk1NjMj5yOp4dQLF7QOS8ZPR9dJRO4lko3THxcpxI3Nnxng2MYQsH8WTz72I1IAMs9dPA0JohGFMjJjYsn7S3-TqOuWZid-6zi3N3WJgBFjwc5kluACVVNKMzJ7e4giYAEJ4PYZ5_E5UH0-SBLibne_03_uwUiYNImdsJh3t_u1qg_DEczon_r3zN5TPOMSMHYSxYhhjoxJ8LsfHAmHep1h_6piMyydVbd5LgRX8PS7G1jz6FODIqtq8XzLO-yQ7JWBmrbgu_vKxFfcVg7Lz6krwNfc5dkDffAWYFWLUJsWT4CQtM6vu822oYT_jvmmR10cgJHtcnS59zgJv299RaspWUwjHK86lNB55EuNyfG3vz-z106RkF0IdQqUrOzbwu3uafci84-t-7Mplsly694rp5hrkBdVHtzspTfC3KXkB4KqQaZCJavF1bwJHxRog1Qh2ZUol6qKk304_BCayFyG3RTFiNJM1wHYEWtNCBPnONgJNVFy-ZFbyvWC2Xzm_oTB7itoxnXS0MbK3N_5pdjsp0yjay5v_MrXorIdTDgrnrVT2Wz83QjtNJPs57TZ-5DvzhBViQuoB0JuDIxM8Z8Z9ZrHd2kVg3ONkGNuKoxqtt7893ZWnestCCbEXtKSZexBnhDkO7xdhkAsvJFJ_EG-zkE8Z00aCxpzHWBRJrInlK0lcHkOllggFcdF0WkbmwjKWy47uwR6LVEjKsgb9ITQO29h1pURQIIs5Pf2y2eh5C718gIse6zpnYfThXfkdNC_Df3YjLuoPs7m93taCHO41VYV_loDBKYrQVD9MTfHUwqp3hW0EhDtq1C7q9rc08g7zC2y1othJtSucCLsnBnNDIVxX-dixzqZztUbQrSg0WnTmDd6rzkM0iMaePk1ZcMMa7MBW27y1ZxdI%3D%7Cl3iQqc6exdieoia3pX6nsw%3D%3D",
						"customdata":null,
						"ext":null
					  }
					],
					"privacy":"https://smadex.com/data-subject-rights-policy",
					"ext":null
				  },
				  "ValidImpDuration":0
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//yeahmobi
	r.HandleFunc("/mock/868/get_yeahmobi_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"DspId":100062,
						"WinPrice":0.18330325,
						"DspName":"bigo",
						"FloorPrice":0.01,
						"ValidImpDuration":0,
						"BidPriceCoefficient":0,
						"id":"15d1bb3e-cc59-4024-add3-1b5019be5fd4",
						"seatbid":[
							{
								"bid":[
									{
										"id":"644841349853570256",
										"impid":"1",
										"price":0.18330325,
										"nurl":"https://api.gov-img.site/Ad/AdxEvent?bnurl=aHR0cHM6Ly9hcGkuYnl0ZWdsZS5zaXRlL2JpZ29hZC9hZHdpbj9zaWduPUIwNWtmbkpaRU5OZjJlODk0UUVhQ3NWVlBkUHlqU3FsY0w1aGdiJmdhaWQ9ZjAyZTc4YjAtZDRkNy00M2EyLWFiMTItNTAxZmQ2NTY3ZmQ2JnJhbmtfaXA9MTAuMTUyLjcyLjQzOjIyOTk2JnRpbWVzdGFtcF9tcz0wJnBrZ19uYW1lPWNvbS5sZW5vdm8uYW55c2hhcmUuZ3BzJnRndF9wa2dfbmFtZT0mc2RrX3ZjPTAmZGlzcF90cz0xNjY4NDc4MTU4JmFkX2lkPTMyNTM1ODAwMDE3MzI3MzYwJmFkc2V0X2lkPTQyNTM1ODAwMDE2NTczNjk2JnNpZD02NDQ4NDEzNDk0Mjg5NTMyNDkmc2VyaWVzX2lkPTUyMzU3NTIyMzk5NTk3ODI0Jm9yZ19pZD0yMjIwMTk0NTM2MzY0ODc2OCZwbGFjZV9pZD1jb20ubGVub3ZvLmFueXNoYXJlLmdwc19hZDpsYXllcl9wX2xyMSZhY2NvdW50X2lkPTExMjc0MiZwbGFjZW1lbnRfaWQ9Y29tLmxlbm92by5hbnlzaGFyZS5ncHNfYWQ6bGF5ZXJfcF9scjEmYWRfdHlwZT0xJmFkeF90eXBlPTEmc2xvdD1hZDpsYXllcl9wX2xyMSZhZF9zdHlsZT0xMDUmYnQ9MSZjZj0xNDMtMTEmb3M9YW5kcm9pZCZjb3VudHJ5PXV6JnBybz1leUpoWjJVaU9pSXdJaXdpWTJsMGVTSTZJaUlzSW1OdmRXNTBjbmtpT2lKMWVpSXNJbWRsYm1SbGNpSTZJakFpTENKcGMzQWlPaUoxZWkxdGRITWlMQ0pzWVc0aU9pSnlkU0lzSW0xdlpHVnNJam9pTWpJd01URXhOMVJISWl3aWJtVjBJam9pZDJsbWFTSXNJbTg1a1IwUk1MY3lPVGUwc2hFV0VUZTVyN1l3dHpPUkhSRTVPcEVXRVRlNXI3c3l1UkVkRVJpWWx4Z1JGaEU1dWpDNk1wRWRFUkVXRVRzeXR6STN1UkVkRVN3MHNMZTJ0SkU5QiZwcD1leUpoWW1ac1lXZHpJam9pTVRZME1qaGZZbUZ6WlY4WkZoa1ptSmljcjdLOEhCZk1Dd3hPRFV4T0Y4eXZEZ3ZtQllaR1JnYW15OVltRnpaVjhZbGhrWkhCbWFyN0N3bUJFV0VUQ3lMN0N4TXpZd3M3bVJIUkVSRmhFNE5Md3l0aTlhV1FpT2lJNU1ESXpOREF6TXpjeE5UTXpPREEyTURnaWZCJnNvdXJjZV9pZD0xMTIyJnNvdXJjZV90eXBlPTEmYWNjZXNzX3R5cGU9MSZpZGVudGl0eV90eXBlPTAmcHVibGlzaGVyX3R5cGU9MCZwcmV2aWV3PTAmc3JjX3JlZ2lvbj1oayZzaWduX2VuZD1lbmQ%3D\u0026ap=${AUCTION_PRICE}\u0026sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=win\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158",
										"burl":"https://api.gov-img.site/Ad/AdxEvent?ap=${AUCTION_PRICE}\u0026sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=billing\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158",
										"lurl":"https://api.gov-img.site/Ad/AdxEvent?ap=${AUCTION_PRICE}\u0026sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026lr={AUCTION_LOSS}\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=loss\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158",
										"adm":"{\"native\":{\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://gdl.bytegle.tech/as/web-source/4h8/1ZxYEX.jpeg\",\"w\":1200,\"h\":627}},{\"id\":2,\"img\":{\"url\":\"https://gdl.bytegle.tech/as/web-source/4h8/25saMJ.png\",\"w\":128,\"h\":128}},{\"id\":3,\"title\":{\"text\":\"Действенно\"}},{\"id\":4,\"data\":{\"value\":\"Как похудеть к отпуску за 2 недели\"}},{\"id\":5,\"data\":{\"value\":\"Узнать больше\"}}],\"link\":{\"url\":\"https://api.imotech.video/ad/pixelfile.html?url=https%3A%2F%2Fnews-sphere.com%2F151527-vrach-nazvala-deistvennyi-metod-poxudet-k-otpusku-za-dve-nedeli.html%3Fsource%3D1350%26id%3D85894%26tid%3D85894%26bbg%3DChdjb20ubGVub3ZvLmFueXNoYXJlLmdwcxIkZjAyZTc4YjAtZDRkNy00M2EyLWFiMTItNTAxZmQ2NTY3ZmQ2GKGByYmn9N38QgiETMyNTM1ODAwMDE3MzI3MzYwKhI5MDIzNDAzMzcxNTMzODA2MDhiBDExMjJoAXAB\u0026bbg=Chdjb20ubGVub3ZvLmFueXNoYXJlLmdwcxIkZjAyZTc4YjAtZDRkNy00M2EyLWFiMTItNTAxZmQ2NTY3ZmQ2GKGByYmn9N38QgiETMyNTM1ODAwMDE3MzI3MzYwKhI5MDIzNDAzMzcxNTMzODA2MDhiBDExMjJoAXAB\u0026pixel_id=902340337153380608\",\"clicktrackers\":[\"https://api.gov-img.site/Ad/AdxEvent?sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=click\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158\",\"https://api.bytegle.site/ad/trackingclk?sign=B05kfnJZENNf2e894QEaCsVVPdPyjSqlcL5hgb\u0026click_flag=0\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026target_id=\u0026action_type=\u0026click_type=\u0026click_prop=\u0026timestamp_ms=\u0026pkg_name=com.lenovo.anyshare.gps\u0026tgt_pkg_name=\u0026sdk_vc=0\u0026disp_ts=1668478158\u0026ad_id=32535800017327360\u0026adset_id=42535800016573696\u0026sid=644841349428953249\u0026series_id=52357522399597824\u0026org_id=22201945363648768\u0026place_id=com.lenovo.anyshare.gps_ad:layer_p_lr1\u0026account_id=112742\u0026placement_id=com.lenovo.anyshare.gps_ad:layer_p_lr1\u0026ad_type=1\u0026adx_type=1\u0026slot=ad:layer_p_lr1\u0026ad_style=105\u0026bt=1\u0026cf=143-11\u0026os=android\u0026country=uz\u0026pro=eyJhZ2UiOiIwIiwiY2l0eSI6IiIsImNvdW50cnkiOiJ1eiIsImdlbmRlciI6IjAiLCJpc3AiOiJ1ei1tdHMiLCJsYW4iOiJydSIsIm1vZGVsIjoiMjIwMTExN1RHIiwibmV0Ijoid2lmaSIsIm85kR0RMLcyOTe0shEWETe5r7YwtzORHRE5OpEWETe5r7syuREdERiYlxgRFhE5ujC6MpEdEREWETsytzI3uREdESw0sLe2tJE9B\u0026pp=eyJhYmZsYWdzIjoiMTY0MjhfYmFzZV8ZFhkZmJicr7K8HBfMCwxODUxOF8yvDgvmBYZGRgamy9YmFzZV8YlhkZHBmar7CwmBEWETCyL7CxMzYws7mRHRERFhE4NLwyti9aWQiOiI5MDIzNDAzMzcxNTMzODA2MDgifB\u0026source_id=1122\u0026source_type=1\u0026access_type=1\u0026identity_type=0\u0026publisher_type=0\u0026preview=0\u0026src_region=hk\u0026sign_end=end\"]},\"imptrackers\":[\"https://api.gov-img.site/Ad/AdxEvent?sid=644841349853570256\u0026sslot=ad%3Alayer_p_lr1\u0026adtype=1\u0026s_p_r=0.999900\u0026slot=100053-100742\u0026event=impression\u0026publisher_type=0\u0026sspid=1122\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026os=android\u0026adid=32535800017327360\u0026pkg=com.lenovo.anyshare.gps\u0026pr=18332159\u0026adxtype=1\u0026area=asia\u0026identity_type=0\u0026req_id=15d1bb3e-cc59-4024-add3-1b5019be5fd4\u0026auc_mode=0\u0026country=uz\u0026ts=1668478158\",\"https://api.bytegle.site/ad/trackingimp?sign=B05kfnJZENNf2e894QEaCsVVPdPyjSqlcL5hgb\u0026gaid=f02e78b0-d4d7-43a2-ab12-501fd6567fd6\u0026timestamp_ms=\u0026pkg_name=com.lenovo.anyshare.gps\u0026tgt_pkg_name=\u0026sdk_vc=0\u0026disp_ts=1668478158\u0026ad_id=32535800017327360\u0026adset_id=42535800016573696\u0026sid=644841349428953249\u0026series_id=52357522399597824\u0026org_id=22201945363648768\u0026place_id=com.lenovo.anyshare.gps_ad:layer_p_lr1\u0026account_id=112742\u0026placement_id=com.lenovo.anyshare.gps_ad:layer_p_lr1\u0026ad_type=1\u0026adx_type=1\u0026slot=ad:layer_p_lr1\u0026ad_style=105\u0026bt=1\u0026cf=143-11\u0026os=android\u0026country=uz\u0026pro=eyJhZ2UiOiIwIiwiY2l0eSI6IiIsImNvdW50cnkiOiJ1eiIsImdlbmRlciI6IjAiLCJpc3AiOiJ1ei1tdHMiLCJsYW4iOiJydSIsIm1vZGVsIjoiMjIwMTExN1RHIiwibmV0Ijoid2lmaSIsIm85kR0RMLcyOTe0shEWETe5r7YwtzORHRE5OpEWETe5r7syuREdERiYlxgRFhE5ujC6MpEdEREWETsytzI3uREdESw0sLe2tJE9B\u0026pp=eyJhYmZsYWdzIjoiMTY0MjhfYmFzZV8ZFhkZmJicr7K8HBfMCwxODUxOF8yvDgvmBYZGRgamy9YmFzZV8YlhkZHBmar7CwmBEWETCyL7CxMzYws7mRHRERFhE4NLwyti9aWQiOiI5MDIzNDAzMzcxNTMzODA2MDgifB\u0026source_id=1122\u0026source_type=1\u0026access_type=1\u0026identity_type=0\u0026publisher_type=0\u0026preview=0\u0026src_region=hk\u0026sign_end=end\"]}}",
										"adid":"32535800017327360",
										"adomain":[
											"news-sphere.com"
										],
										"crid":"8ef0e46c217b9d4b27008379e161e4a1_1",
										"ext":{
					
										}
									}
								]
							}
						],
						"cur":"USD",
						"ext":null,
						"vaststr":"",
						"bidid":"",
						"expire_time":0
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//criteo
	r.HandleFunc("/mock/868/get_criteo_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
				  "requestid":"1ef23679-b858-424e-8686-c79b886e18ba",
				  "seatbid":[
					{
					  "seat":"criteo-global",
					  "bid":[
						{
						  "impid":"1ef23679-b858-424e-8686-c79b886e18ba",
						  "traceid":"63732c7d283d09425dc62d04fb5b7900",
						  "lurl":"",
						  "bidprice":0.02,
						  "currency":"USD",
						  "dealid":"",
						  "adm":"{\"native\":{\"assets\":[{\"id\":3,\"title\":{\"text\":\"Combo 10 củ huệ (10 màu đặc biệt) tặng 1 củ\",\"len\":100}},{\"id\":1,\"img\":{\"type\":3,\"url\":\"https://pix.as.criteo.net/img/img?c=3\\u0026cq=256\\u0026h=400\\u0026m=0\\u0026partner=34657\\u0026q=80\\u0026r=0\\u0026u=http%3A%2F%2Fcf.shopee.vn%2Ffile%2F50b81e4850e2d27043de2d3cdb9c9f3b\\u0026ups=1\\u0026v=3\\u0026w=400\\u0026s=0M8jubcPZsAjYJxriGQB05iS\",\"w\":660,\"h\":346}},{\"id\":2,\"img\":{\"type\":1,\"url\":\"https://pix.as.criteo.net/img/img?h=1200\\u0026m=0\\u0026partner=34657\\u0026q=80\\u0026r=0\\u0026u=http%3A%2F%2Fstatic.sg1.as.criteo.net%2Fdesign%2Fdt%2F34652%2F201215%2F6c7d301ff5534031a06726ac11d5160b_native.png\\u0026v=3\\u0026w=1200\\u0026s=RBjvRjFTIGEtJX8ZEdvtWKpD\"}},{\"id\":4,\"data\":{\"type\":2,\"value\":\"nhà vườn bán thanh lý củ hoa huệ. giá 8k/1 củ thời gian bắt đầu 12/6-30/6 giá sốc ưu đãi cho khách hàng. củ khỏe đẹp ra hoa trong khoảng 2 tháng sau khi trồng. củ đặc biệt rất dễ trồng và dễ chăm s...\"}},{\"id\":5,\"data\":{\"type\":12,\"value\":\"XEM NGAY ▸\"}}],\"link\":{\"url\":\"https://cat.sg1.as.criteo.com/delivery/ckn.php?cppv=3\\u0026cpp=VL9-8hlWy3rX0gAVP-apR2D7_6rbfqcsYwYSr398wKU2yMdFjry9xzwZs8F9a2hBAmbajFCOPvg0qL3FqOW4NbA6CAFcjrE1aDWwgHnRIfshcrZtE9qVpR5C-PsP6aDzcGdTkO0OT_wt5Kzw61_rxo6FT9XU-9JGoe_fTqbeeiNmh1stk6aou--yu0hw7GzFs3fVqqXkBUpsevag3lIGVZNIC2XMLmCMX4TKcTwbT9OqLXpZOKZnSlSISJNwzn75ghvEez_N00P14GtGwDUefOa91xVMW_NreRKsIjyjuUouAzQj1LRHSXMd3OsaP92a19mvnPUoD47d3IylMSPPYZwCbN-mPAWTERPFDL2f9aRPomxV4_AB8YVz8f7NWQ-6U7IyN_NPtAoGHlETfGCxXfEdmkRypeFccjw9bu3Ox77mnb6mnNtgDM6UivbpPTWQde_M7hi2OGuDIIPRhSh-UJ7x6AG7YorGIXJog6s-3k1iWTNv4ambmi3FgzHzBcUx58__Gw\\u0026maxdest=market%3A%2F%2Fdetails%3Fid%3Dcom.shopee.vn%26referrer%3Dcriteonew_int_63732c7d283d09425dc62d04fb5b7900\"},\"imptrackers\":[\"https://cat.sg1.as.criteo.com/delivery/lgn.php?cppv=3\\u0026cpp=RBqBSBlWy3rX0gAVP-apR2D7_6rbfqcsYwYSr398wKU2yMdFjry9xzwZs8F9a2hBAmbajFCOPvg0qL3FqOW4NbA6CAFcjrE1aDWwgHnRIfshcrZtE9qVpR5C-PsP6aDzcGdTkO0OT_wt5Kzw61_rxo6FT9WKVbp519AKA-Fx39OkfiC4-S5YPCd58SUul3rLA8Fd8G8wEzCPL4rDOfT60f5CT8i0S69Hh2yf_-k6fjNsRJkJlh2RYD7V3iQ_ZthDYfQ9NaRikHGYJ7B5SwqGtoqy7hFKhP5j9ILGSZxG9am3sAc_u3xAtOQAGMP_gwISoWhz_jq6DFY4RyPeeTQe-kAvwoaGUQisn38JIh3vrbNwdC8s-wtYx9spT956IRcjDr4vmhDB0oK4SdKaUmKWTamvP7Q6jCH_YXisdyIcLH31QhvVc2G99jgWs6ZbQTFbKP7iVBDoAD4cTqwkM7JUT6htlQM\",\"https://impression.appsflyer.com/com.shopee.vn?c=VN_ANDROID_UA_CPI\\u0026pid=criteonew_int\\u0026af_viewthrough_lookback=18h\\u0026advertising_id=793c5a06-5fa7-4877-a537-7099c7774303\\u0026clickid=y_fnRHxqUWxkQ1djd3BIMUFCL0tNMFBBaTNlVGRGRTcrOXJWNEpvRWdpSHZRSklQNUphSUQ5WHN2d1h5d0hvU01NOWFvMUt4NzRtSElGUnRJWW5UZ1hvWnVxUjBXNENnRnE1V09SQU91Z1ljRGRNbz18\\u0026af_c_id=221350\\u0026af_siteid=com.lenovo.anyshare.gps\"]}}",
						  "adomain":"shopee.vn",
						  "crid":"10790209",
						  "fillimages":0,
						  "w":0,
						  "h":0,
						  "creative":null,
						  "payload":"WWkAYYeSPl7FHP79vobsZWdezgvBiO0ubSTAtwNt_6K96-ctmGwEMb_RxIzXFFCs3iuEKyz4Dyxs5xgxgdHvLfwiar2AWMOda9i4xFfQ4SnxUs60Hcd3Hwht04Ub1xfcihPB5KkPI7aNY9KO6-EzqcL6Hut4KaDrEJdXCdxl2o4RXEKj1eD7u2cL-aFdZx-pchO1oRgG2xUqJso11TXVmklCRba7wTGBiAaxqqpDJ3pxxpkUkVcH77yAC4hV7u_GfRt-qwK0IbuFk0sKq2aCIr7fox-agxJuN9WFq_jfbcBWVANb7UAkEb11IgDYosg6tzE0_6rpoCa8b5SQCpfvk2pq2qQgqn-UK1VEV_Lw_O6wbFbQojJqE0G4z5mnM62gz4QrXUGXxTNMI5UwXMphsKjikuRAjcQHw3oyuLN62aYYC7z5OOwjLl4Go5JIoK9q2Qpfd4lKZcjolQDa4Aa_4Ft24nnYpvvONZ1xKrh5GJimORFulZR0VvP1FVJZhecQUPYBRIWasg9tFSngX21OTNgOcxYUrRQGx9v8CILW0h5PFmu1rHwAYCglgL6gAELItqJ-F4zg8YCHoURBHA4-1UH3wcZ0NtMO-vb_yvlb3QNidHVRTrKOaseX6q10NOr6TIjjpG6PYsB95T12XU3i2gEZXIDph1Mx4wM1s-CjA2rYWFt0sUm1Ajw1HY_nl881JdEWFryGRRVoUssvOkXa6_AYpbjAR1YnT9vQy-gMXIuMX5hyjaMg_PgAyt_0tFHy2l3ycbNHvlcnvMVteVQXjHdfYGoW0dTa4vIb14tY2KrbPUPiRiFgarlKqQCL_Fr_",
						  "privacy":null,
						  "assetsurl":"https://rdi.sg1.as.criteo.com/delivery/r/native?u=%7C1skYvcDLpubqLUXofC9Pdaq0hf0cEF7F5t%2BA7wR8rfo%3D%7C\u0026zn=\u0026callback=\u0026c1=Dcz_gsP0hEuJH1VnunqGy_FuLiVNXMEiQz0xpqHWOrICeFz0kuZuJbO98rEdpK0qxgrzccpE5iOvCmdQoEciYJsN63QjiTf7PX8Dj6AzJtZZeFU9ju-uh_W1HBn5ag68IfIS5sKGdMKlVsvYNH3RR0zJJ7JO5dydmxJuack6-I-iVBVpTFwvmaNGPrbw7_6Vm3J7tjCBZ8Dsp_Ko_nBzjtS_GME9omJ_GnXQgZm0jfgPQ4QWGjmF70Q-xJ1TTgEEPAOvK17z4DWejt2zfGU13LU4Y7nPVhalAcP8gZQZCdCvGc7HAFqWH9L9xuroQR6tL72eI-a6zEeRlM2EcPdDu-axU49fAzvHfwVO_BP0tfI_dqkAcSf1FoEB5oENHa_JbpLdwyQQL_laiWcR7-BYGe9WS-eKkMpCO_c_i6amgd1YuHeY01dPDVPNQuSxjzlq8qYWnu4qyKPUYqt25gwx0Y1qUgfYuM5160Q0x0qY7F6s9L03QCYa7MYYgnXai_9KTavOQZ_aV-DuH5ol9_ol8Nb4y4iV8KNQ5lwpPRi8vfhZhwZ8OuaQsk00-RQaQzFv9WAioxhxL4LuszH1uIOyCxL2LtxCHt8_mSIo78lxw4dOg1zwiWXnvdza7qDXbFljWV4NhM4DeivJuzpPEgIb9rziaeyLg_NxQJ8zBwjQNLkLDfd3XWu9TFkz_K5FKRMdAO3RTNHHkPsDbe0dN1cU1KLkq8AKxn06bUr58iGBKQpnX6jmaz44fg"
						}
					  ]
					}
				  ]
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//ecom
	r.HandleFunc("/mock/868/get_ecom_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"bidid":"3d162b01-0b36-42cb-9583-3b8a4e69283e",
						"seatbid":[
							{
								"bid":[
									{
										"id":"3d162b01-0b36-42cb-9583-3b8a4e69283e",
										"impid":"1",
										"price":1,
										"nurl":"https://test-track.lacunads.com/win?viewid=5th4nW7N_Vcfu_iStyeze2GWRz8xXrx-3kk4Bh45kQQXV5wi6RZk61IpfSbTeI0og_Pjcd5NjgKJLF8xbUQpfPMnBpkvFIfv3t9Q6kdDvJUwo9KEWjhu07bP3CxPDxnh7EzUJAMH1DayUWdDG8y6e6lq-2HqNu9mi1bvIpUtf3EejnINfoqlobIVqHkOQRW4vF_A78PELz9G9AvmMUTex4iXGHzJ5iP-aWsJRr4oSg0Xvso6W8Hlgn_DFP3dC412j_r1jmJQ1-qsOb4hS4-faFgcTTkTc4AX4rrB3uCV6U-kFHgnAjHXDiQmGy-OWlhtg7FZWMo-kjO06r7q4fjb4ihbVa2Fqw3NlI4BHhKF__6lQVsonU6zpQn02Hmx8KmGtg8sjeSZpn2WamnEmg7TdBYY1w6vdyDWir09w18sRqUBead-XBXIHlnZTjnLArMjSQExsr_0NAARlZ17GIUnOC3FvIGE4KCt3VenFqDD67KtWQAEnqZ6uHVItt7eRn9oj8pHLSyd2Sn5MDIGTOVp10mHidwHjJdg23sBou0MIrnRYjH7D011TxH-GHvR_7xO1jcFaJSRcXlq_S7TQpXx3TM9RHVd4KBdeWy6SrZbXwOIBg7gZUmVDm4rZes2LEXstnZe_YjKlA4G4XTZkjZFRcoIF8JKG9J6Rru2ni4L7JAj1t-8gGh4oMSleMl_Ca1HMhAqOtKIhzzWE98cer9SI44MfhKyZXCjicj9cc86BbYieU3uHYlc7lCCrXEvVTZxAMtm4qUtRvHDQNd30ljYEs192nxsVTts3J7xX3dbtfK67y4TJOuljIIKhTVQse-Y1sBJn9FMwM5EIHsALMfF2LJRB9HICfS91dnvdL8ifdiVJOF53bH8TpZ3d4PbuzETAXDZtZp2cVikG516B7EA2cYUt1CqfXNSy1__UFX8Yxa1ZoNRjMel5pi9dPp9teocy8hwqp9uEaVvRReUFmSxJd7jY03_KcfqDQyiDZMejH5KDS5cZ0mrTdpvxblu-wJFxqF3B1rx5eS3MjBZqKqWb-smLHa0IHg_uoanZ0FN3F8Qbpm20wgGq4kX0vWDU0KP9a8AMkHOxWUzA5MSrabdrvNBVAP6yMd7ki4Kqj-ApCU1uIMb5QrGDrrxwccdjjhD64OYs5QYsZovsKd94ANdORBplr-_g4DzkcXZ50bjAofjdUrwDUXsHKRBooL0x0agZ7yHML5kEJ9gQ2LJ9k7LLpyU6fawaWg3uZ-sHJyzlLkA0VnXZ5tBp3m2kAKP12cn05oZYzOc1fD9eWS1EbQxynDcAR8B-aJ_7E5O3OQ9_ObCdlW0MjKTLy7kVw8ndT3owbLhEhyEqxnWvmQIXlGrdutvPY4K3dRbSfy6o7Wy7pPREcGC8mBFQIn56bNAhvmJFvpts_mxx9A9CCBviqznhDz7TjnE0B8wToeRGK0QIBwk6vbj15rL3C6e6x2anFuDEk-oxMy_qZMVbDjCI6-pIDNRH-QvQmV-PgpiKYxNU3slOg8bc_8DjQweRmLo6dNd_c5zarFrQ7d3nJDQVHxGUUA9CggW7OBoAaJqPXwCNVMYpx8bTlgWlAijnvskkgOqCfkAM_SEGBf7IAKzcEAjhi1DdcfNPs_c5wfvzCTa7qBHWVXKGiVu_oYeq_3I4CwN5gyCdZ8KBjtTrKcsjUpg3b9hyiBeIqZm0zV-MKvaKBmbmXr8w_kslO97RHLfQqJ5Ipqb85fbl5X85ex7_9T1NDJfWrfFOShF3HE2gveKXV0G-vAZHKOA6Tp-92JgpJJgWqaOUid3mkGBVktSOwG1pTac7ECrb6WoP4dA25YNbFpySH4Wxd7-CXfmU2Cvr94PgNhg6cE77LXYxX6C9oAN6ZXH5_WQziCJI98iQCvS-l2bWfn4LwdvGCyDXk9_ZimidkwO3aJx6XKWSskosS9YgFrU0kRREuow4h6bpUleL3rRwZAvGZG395DbAyfs6X1cm48WLraFSEjb9EkiMfz5fhMBEeB91OeayFs80eLjc3EhS_dgebnJlpjrpKiTY4QRlCNfZob-WqwsU319F6L1HJUyBcNksbj3AXmYiwrPSBUhEhoNOvbRRtDiyEbIhzu9HE7FLmecdhjUjQPFdNV1zZFte58fusgaPA5FEBpB5h96LEyzFyKLl6Gi3xe9QJ3jmoez9NvkS1OWp9z-eh1dJ3AYN8nS9XcYGzUGFze2uBsGG5PDyRQB1QWy-dDcfZtUG8LJKSqccyAVj3pI3E4CDmlYfby4wfX2qTWCjI1z5dFlQDNkrrX11gvSo0r0zzUtDiSw9UcVLmtua5tqOybetCd9CEstA-3Qkj0TpNYpJzchVg5GU5iUiZZlYmvsjjW1CUUJbqDwBsSvZeAvyLHSY86ZoyKeaYa0b61_Ymr-YA7iHHTXw1bVsOgH_0IEIPx98aU-TttUYYYiYqOtH6WOWrjH-XA1HYFuKdQn8EQGkEIKMXhJ7LiexRQhZzzzOIgha580XQgOtVdHdntAV50H1UJus_bHKo0RSzc7AhmA_fL1D6D3bPHujdsYKyvnre_ivVy5gzz_ZWqO-YyHpva1VTfd6AKijlrXv2OKRIbgQC83MZQ9TKEstT3aMXHG48NUYbyZeaNO04wAqbk_cL6sN3qBmpSxi4qrimag_PmW95Q6W3uSWAoCAVcWZDp23RLpqSL2W28Jk-WZVxiyNGYIXsIyFnkQG8BLnJVQa_yktwnuouPmCjbJYqshFHear0dZ2idJr3M7Y4-4DNIrj9Q5w6hbDhkOa8DOV2x0QPLsW6m50NzqFx4rvJVe5zUngOn-at_GTlfyj34JyEhiRJg6a1ndExthaIntFWqcwkSwLFeUdBVKZ3E4Qf1y9FzbeoRQ2-fUAFlgnC6ndVvBcayPc5lMRxRbWMaOYGDI8gzO2eD0XYubqarMu_fxpmqmPzgQz_f1bhqIuha9jm31sHRkv8Vc59oHkFXErMxzka87HMawCm7JCgc_U__DPdHn7RggcfhTxxnqJDad0xKSBx8WYJrN0hWCq2pUAKqORBkRn8reEVxWO0otUuG-oELznVlWQWfi32hhuyo3gdJZG1J_CL9f095JwoQ7Lv4mZxw17lEI0AOwhVaSwnpW4jf6cwZuLnv35oXTZw6Yn_QXgbogrB4IFV9J7hXKxX__aLAV7PdTF9r-YHKUzXlRFWASQyR4alxQYF5DpZAt3pg0WUEHUGZ83Zn00rmFjpfqr1sxMwwwgVYsbAijtb-RTmaIMjgk85dE85Jc-1X31xITGGRtGgfJdL7gh9rbzLkSNDcrbcaUjLAQE5EQV_XW8CdlZQ6tntqulJCM-k9QF3T1PSIBtMlWKr8qwrQ8aadOZkYK8u0DgAZYtvhzoyyF179dFIdiXeAhMAUfLuz8LdnY8RusQGYwg2BkZP-vMYwi0mYpMy5YsX0JMRC9NcSIWL3ye2OoxseNES5miI0r5oVBDBQqpoZwIFx0ZW5K9SanJHdtOskiDQ7XoWr9yWqug8rz4p7FX48VsPSwa9NVWQMHniR1Nd_8c7nTk671cX48DJSjbLClJx41wKlgdl6LDrVGHvhiOcSJ44q5ECaQ4DKk8JIxqT1ayFWFyG3YbX7U6F1j1LGXT4JAaoMyJMOnWNtIEwOHAx-jPn7T1xTOJtdzhWlb1a7jVwS7OJNQksBamXV50ehS-XuR-5zD2nhYooQiiJ8n-2muSfcYVggzi_IT42xbFXb4UEXHErmmsSCj_SjGiQRAu3S0EgUDb9W24YJZjdRQjr66kIA9O7mwNqfD9Ra9i5r89i9m1EBG7o2l_9PJ3bPn9rAfVX6Nja4SPwl7OIDswM-LxlNMpQDloby-__Qa2If0dTvt61zvJ4lJ6HGfUqs-q8pzY1jA_VJzCNkUR0xPgIdmI7fTJr2rGJFBJh0xw6V8bB5NLqoyjb7kRzpnrBjSoDMDmxnHeYVBfdE8cqwxmrUkGF-82M64GwwcgC2k15ylscZ_pkIcP9bsHKWOERcjjn4TV8RafsIB0ZG_lquV3tZjLDDYmfbFMbh3WluHqlATB9lh6XiYXPm9AfRdDYmh75alTb4FvQbSc_FCo3m0oxJcU3eIJdhxbYo_e-SIH-ScZPVAoxyfSZyyjNJCVHRqnuDqaRLyjPho-IcCfcUiu-tw_GqwQkH3kCD7g1_W3MYKo-lkZ7r817DBAJURdvYEnrpEaqOpZy6foEco5fmJNu8OFaRgbnncOkAKXFHoiKRb7Ky-CFLeHFxEXm2_Jk4W4NLGKWiwqDlkwcRpBZqtMes5LULDyllUylZRmSwKUIimNaJuy-SApIziUzalbur4ku44EbEpe1v2xnZZkvjcTPkcXP7KWJaw0pSLlJGEl6k_CWXW8H7Pd3-uzP25yVyVj-yPhd3PCT7IvsyQ9ukt5vtcQ-rF0kuhDDBW9mLBPkKVtdoSyWcaBsbGeDhZGXtm7AsYlXj8vHiGkCnEchiFQDXkuG8sQsSOEGLrhLlbfeJF3oICsNNEFnTCag030xtKcKw5PQQfUaxpNHQVcgiy_--uQqeNnitUZMBXxbl7vKne5aK3Xppo9CGMVptcLwMS9TZ3mTFwH93iCHCrXv-6KPy1jtZzgk9J19Xmc1XTwo7cR9X_EIifPcC6tbONY5HbJ_u9vnSdfEhD5xtv_w0-whYgq7gHOZx2AWthXhtk7p-3O9jXYoXu-7K-gQ0TsGFaS8O09bC3reEO26xASLIbXTy6nj_bwCT-qRn93mWsrWBChm4uci3CoanwauZZ7i6iiFU95iXOk4X5zEt6Ff7b7Uziq4gDj29EkXsN14SA59-mCtwEvPKKtCZfUIBpURyKqqm88Lo_98uFYc3iMtWdDXYfDU1PMv0Nw5kTU4rhhgPVT-N3D6mcWN57YBrJZSob1SbBHeQ6-Aa8Pb1vGQzbudCz1V9iEd4A6801RIP4jqCa2-YV-8TIBzmWLmD6pi-Y5FR-TGDmaars8FtfevNlZRlAKATW6r3yMuh_kE0elUL4wf1puVBK23h03zCl7VstrO9hBzUAM0MhUisE0ihlymHe8tq7W_OMnsRsXhn7ROTJ7AFnaUVEugD39YbazWmhWKgGdtFSFocEFLk_jSwGQ6prHcCu5Zeo9GhmGca4rUm3NzghrhRsLpCqKGuYApFFVBo2aTQe5QE8176xNjbP&enc=1&auction_price={AUCTION_PRICE}",
										"adm":"<VAST version=\"3.0\"><Ad id=\"1825467\" sequence=\"0\"><InLine><Impression><![CDATA[https://test-track.lacunads.com/imp?viewid=5th4nW7N_Vcfu_iStyeze2GWRz8xXrx-3kk4Bh45kQQXV5wi6RZk61IpfSbTeI0og_Pjcd5NjgKJLF8xbUQpfPMnBpkvFIfv3t9Q6kdDvJUwo9KEWjhu07bP3CxPDxnh7EzUJAMH1DayUWdDG8y6e6lq-2HqNu9mi1bvIpUtf3EejnINfoqlobIVqHkOQRW4vF_A78PELz9G9AvmMUTex4iXGHzJ5iP-aWsJRr4oSg0Xvso6W8Hlgn_DFP3dC412j_r1jmJQ1-qsOb4hS4-faFgcTTkTc4AX4rrB3uCV6U-kFHgnAjHXDiQmGy-OWlhtg7FZWMo-kjO06r7q4fjb4ihbVa2Fqw3NlI4BHhKF__6lQVsonU6zpQn02Hmx8KmGtg8sjeSZpn2WamnEmg7TdBYY1w6vdyDWir09w18sRqUBead-XBXIHlnZTjnLArMjSQExsr_0NAARlZ17GIUnOC3FvIGE4KCt3VenFqDD67KtWQAEnqZ6uHVItt7eRn9oj8pHLSyd2Sn5MDIGTOVp10mHidwHjJdg23sBou0MIrnRYjH7D011TxH-GHvR_7xO1jcFaJSRcXlq_S7TQpXx3TM9RHVd4KBdeWy6SrZbXwOIBg7gZUmVDm4rZes2LEXstnZe_YjKlA4G4XTZkjZFRcoIF8JKG9J6Rru2ni4L7JAj1t-8gGh4oMSleMl_Ca1HMhAqOtKIhzzWE98cer9SI44MfhKyZXCjicj9cc86BbYieU3uHYlc7lCCrXEvVTZxAMtm4qUtRvHDQNd30ljYEs192nxsVTts3J7xX3dbtfK67y4TJOuljIIKhTVQse-Y1sBJn9FMwM5EIHsALMfF2LJRB9HICfS91dnvdL8ifdiVJOF53bH8TpZ3d4PbuzETAXDZtZp2cVikG516B7EA2cYUt1CqfXNSy1__UFX8Yxa1ZoNRjMel5pi9dPp9teocy8hwqp9uEaVvRReUFmSxJd7jY03_KcfqDQyiDZMejH5KDS5cZ0mrTdpvxblu-wJFxqF3B1rx5eS3MjBZqKqWb-smLHa0IHg_uoanZ0FN3F8Qbpm20wgGq4kX0vWDU0KP9a8AMkHOxWUzA5MSrabdrvNBVAP6yMd7ki4Kqj-ApCU1uIMb5QrGDrrxwccdjjhD64OYs5QYsZovsKd94ANdORBplr-_g4DzkcXZ50bjAofjdUrwDUXsHKRBooL0x0agZ7yHML5kEJ9gQ2LJ9k7LLpyU6fawaWg3uZ-sHJyzlLkA0VnXZ5tBp3m2kAKP12cn05oZYzOc1fD9eWS1EbQxynDcAR8B-aJ_7E5O3OQ9_ObCdlW0MjKTLy7kVw8ndT3owbLhEhyEqxnWvmQIXlGrdutvPY4K3dRbSfy6o7Wy7pPREcGC8mBFQIn56bNAhvmJFvpts_mxx9A9CCBviqznhDz7TjnE0B8wToeRGK0QIBwk6vbj15rL3C6e6x2anFuDEk-oxMy_qZMVbDjCI6-pIDNRH-QvQmV-PgpiKYxNU3slOg8bc_8DjQweRmLo6dNd_c5zarFrQ7d3nJDQVHxGUUA9CggW7OBoAaJqPXwCNVMYpx8bTlgWlAijnvskkgOqCfkAM_SEGBf7IAKzcEAjhi1DdcfNPs_c5wfvzCTa7qBHWVXKGiVu_oYeq_3I4CwN5gyCdZ8KBjtTrKcsjUpg3b9hyiBeIqZm0zV-MKvaKBmbmXr8w_kslO97RHLfQqJ5Ipqb85fbl5X85ex7_9T1NDJfWrfFOShF3HE2gveKXV0G-vAZHKOA6Tp-92JgpJJgWqaOUid3mkGBVktSOwG1pTac7ECrb6WoP4dA25YNbFpySH4Wxd7-CXfmU2Cvr94PgNhg6cE77LXYxX6C9oAN6ZXH5_WQziCJI98iQCvS-l2bWfn4LwdvGCyDXk9_ZimidkwO3aJx6XKWSskosS9YgFrU0kRREuow4h6bpUleL3rRwZAvGZG395DbAyfs6X1cm48WLraFSEjb9EkiMfz5fhMBEeB91OeayFs80eLjc3EhS_dgebnJlpjrpKiTY4QRlCNfZob-WqwsU319F6L1HJUyBcNksbj3AXmYiwrPSBUhEhoNOvbRRtDiyEbIhzu9HE7FLmecdhjUjQPFdNV1zZFte58fusgaPA5FEBpB5h96LEyzFyKLl6Gi3xe9QJ3jmoez9NvkS1OWp9z-eh1dJ3AYN8nS9XcYGzUGFze2uBsGG5PDyRQB1QWy-dDcfZtUG8LJKSqccyAVj3pI3E4CDmlYfby4wfX2qTWCjI1z5dFlQDNkrrX11gvSo0r0zzUtDiSw9UcVLmtua5tqOybetCd9CEstA-3Qkj0TpNYpJzchVg5GU5iUiZZlYmvsjjW1CUUJbqDwBsSvZeAvyLHSY86ZoyKeaYa0b61_Ymr-YA7iHHTXw1bVsOgH_0IEIPx98aU-TttUYYYiYqOtH6WOWrjH-XA1HYFuKdQn8EQGkEIKMXhJ7LiexRQhZzzzOIgha580XQgOtVdHdntAV50H1UJus_bHKo0RSzc7AhmA_fL1D6D3bPHujdsYKyvnre_ivVy5gzz_ZWqO-YyHpva1VTfd6AKijlrXv2OKRIbgQC83MZQ9TKEstT3aMXHG48NUYbyZeaNO04wAqbk_cL6sN3qBmpSxi4qrimag_PmW95Q6W3uSWAoCAVcWZDp23RLpqSL2W28Jk-WZVxiyNGYIXsIyFnkQG8BLnJVQa_yktwnuouPmCjbJYqshFHear0dZ2idJr3M7Y4-4DNIrj9Q5w6hbDhkOa8DOV2x0QPLsW6m50NzqFx4rvJVe5zUngOn-at_GTlfyj34JyEhiRJg6a1ndExthaIntFWqcwkSwLFeUdBVKZ3E4Qf1y9FzbeoRQ2-fUAFlgnC6ndVvBcayPc5lMRxRbWMaOYGDI8gzO2eD0XYubqarMu_fxpmqmPzgQz_f1bhqIuha9jm31sHRkv8Vc59oHkFXErMxzka87HMawCm7JCgc_U__DPdHn7RggcfhTxxnqJDad0xKSBx8WYJrN0hWCq2pUAKqORBkRn8reEVxWO0otUuG-oELznVlWQWfi32hhuyo3gdJZG1J_CL9f095JwoQ7Lv4mZxw17lEI0AOwhVaSwnpW4jf6cwZuLnv35oXTZw6Yn_QXgbogrB4IFV9J7hXKxX__aLAV7PdTF9r-YHKUzXlRFWASQyR4alxQYF5DpZAt3pg0WUEHUGZ83Zn00rmFjpfqr1sxMwwwgVYsbAijtb-RTmaIMjgk85dE85Jc-1X31xITGGRtGgfJdL7gh9rbzLkSNDcrbcaUjLAQE5EQV_XW8CdlZQ6tntqulJCM-k9QF3T1PSIBtMlWKr8qwrQ8aadOZkYK8u0DgAZYtvhzoyyF179dFIdiXeAhMAUfLuz8LdnY8RusQGYwg2BkZP-vMYwi0mYpMy5YsX0JMRC9NcSIWL3ye2OoxseNES5miI0r5oVBDBQqpoZwIFx0ZW5K9SanJHdtOskiDQ7XoWr9yWqug8rz4p7FX48VsPSwa9NVWQMHniR1Nd_8c7nTk671cX48DJSjbLClJx41wKlgdl6LDrVGHvhiOcSJ44q5ECaQ4DKk8JIxqT1ayFWFyG3YbX7U6F1j1LGXT4JAaoMyJMOnWNtIEwOHAx-jPn7T1xTOJtdzhWlb1a7jVwS7OJNQksBamXV50ehS-XuR-5zD2nhYooQiiJ8n-2muSfcYVggzi_IT42xbFXb4UEXHErmmsSCj_SjGiQRAu3S0EgUDb9W24YJZjdRQjr66kIA9O7mwNqfD9Ra9i5r89i9m1EBG7o2l_9PJ3bPn9rAfVX6Nja4SPwl7OIDswM-LxlNMpQDloby-__Qa2If0dTvt61zvJ4lJ6HGfUqs-q8pzY1jA_VJzCNkUR0xPgIdmI7fTJr2rGJFBJh0xw6V8bB5NLqoyjb7kRzpnrBjSoDMDmxnHeYVBfdE8cqwxmrUkGF-82M64GwwcgC2k15ylscZ_pkIcP9bsHKWOERcjjn4TV8RafsIB0ZG_lquV3tZjLDDYmfbFMbh3WluHqlATB9lh6XiYXPm9AfRdDYmh75alTb4FvQbSc_FCo3m0oxJcU3eIJdhxbYo_e-SIH-ScZPVAoxyfSZyyjNJCVHRqnuDqaRLyjPho-IcCfcUiu-tw_GqwQkH3kCD7g1_W3MYKo-lkZ7r817DBAJURdvYEnrpEaqOpZy6foEco5fmJNu8OFaRgbnncOkAKXFHoiKRb7Ky-CFLeHFxEXm2_Jk4W4NLGKWiwqDlkwcRpBZqtMes5LULDyllUylZRmSwKUIimNaJuy-SApIziUzalbur4ku44EbEpe1v2xnZZkvjcTPkcXP7KWJaw0pSLlJGEl6k_CWXW8H7Pd3-uzP25yVyVj-yPhd3PCT7IvsyQ9ukt5vtcQ-rF0kuhDDBW9mLBPkKVtdoSyWcaBsbGeDhZGXtm7AsYlXj8vHiGkCnEchiFQDXkuG8sQsSOEGLrhLlbfeJF3oICsNNEFnTCag030xtKcKw5PQQfUaxpNHQVcgiy_--uQqeNnitUZMBXxbl7vKne5aK3Xppo9CGMVptcLwMS9TZ3mTFwH93iCHCrXv-6KPy1jtZzgk9J19Xmc1XTwo7cR9X_EIifPcC6tbONY5HbJ_u9vnSdfEhD5xtv_w0-whYgq7gHOZx2AWthXhtk7p-3O9jXYoXu-7K-gQ0TsGFaS8O09bC3reEO26xASLIbXTy6nj_bwCT-qRn93mWsrWBChm4uci3CoanwauZZ7i6iiFU95iXOk4X5zEt6Ff7b7Uziq4gDj29EkXsN14SA59-mCtwEvPKKtCZfUIBpURyKqqm88Lo_98uFYc3iMtWdDXYfDU1PMv0Nw5kTU4rhhgPVT-N3D6mcWN57YBrJZSob1SbBHeQ6-Aa8Pb1vGQzbudCz1V9iEd4A6801RIP4jqCa2-YV-8TIBzmWLmD6pi-Y5FR-TGDmaars8FtfevNlZRlAKATW6r3yMuh_kE0elUL4wf1puVBK23h03zCl7VstrO9hBzUAM0MhUisE0ihlymHe8tq7W_OMnsRsXhn7ROTJ7AFnaUVEugD39YbazWmhWKgGdtFSFocEFLk_jSwGQ6prHcCu5Zeo9GhmGca4rUm3NzghrhRsLpCqKGuYApFFVBo2aTQe5QE8176xNjbP&enc=1&auction_price={AUCTION_PRICE}]]></Impression><Impression><![CDATA[https://adrta.com/i?clid=srit&paid=srit&kv29=[ERRORCODE]&kv30=[CONTENTPLAYHEAD]_[ADPLAYHEAD]&kv33=[ASSETURI]&kv34=[VASTVERSIONS]&kv35=[IFATYPE]&kv36=[IFA]&kv37=[CLIENTUA]&kv38=[SERVERUA]&kv39=[DEVICEUA]&kv40=[DEVICEIP]&kv41=[LATLONG]&kv42=[DOMAIN]&kv43=[PAGEURL]&kv44=autoplay&kv45=[PLAYERSIZE]&kv46=[REGULATIONS]&kv47=[ADTYPE]&kv48=[TRANSACTIONID]&kv49=[BREAKPOSITION]&kv50=[APPNAME]&kv51=[PLACEMENTTYPE]&kv52=[SSAI]&kv54=[LAT]&avid=69&caid=bb57db42f77807a9c5823bd8c2d9aaef&dvid=v&kv1=0x0&kv11=6e0bbc2c-6974-4719-a5fd-91b280317b82&kv15=IND&kv16=21.000000&kv17=78.000000&kv18=share.sharekaro.pro&kv19=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&kv24=Mobile_InApp_Video&kv26=Android&kv34=3.0&kv4=114.124.182.18&plid=23397081&publisherId=31&siteId=0]]></Impression><Creatives><Creative id=\"\" AdID=\"\" sequence=\"0\"><CompanionAds><Companion width=\"720\" height=\"1280\"/><StaticResource creativeType=\"image/gif\"><![CDATA[https://static.rqmob.com/sa/20221118/a1fcd18847d88fe185f1c8925905c938__FILE_CM____FILE_CM480_720__.gif]]></StaticResource></CompanionAds></Creative><Creative sequence=\"0\"><Linear><Duration>00:00:05</Duration><TrackingEvents><Tracking event=\"start\"><![CDATA[https://prod-ping.swat-adping.com/ping_server/play?viewid=313F111995C936548E538DB71BC0C541D3742981BEB7583E4A06BCB6D8383119A85F29C31F0F4040BA7D3AFE310345FDDBA68B92BC7112CC60D5EDAC1E08FA3BD4C0264805A63A24FD77E927083F832A49FC88F30D033A95DB30DB765AF1803B27F2AB75468FC0FE0C322AB2BEB0AAB3968E449B4A1776B7E905C83DFD95D3EA3C9CB16B383139D008BB1A9F1E47F7B1B94324517EAB8BD0C515BA6A33C06BE24846C602DEC38028554A7477F30E1C070F27B7002E56619DC469BD108C097324B9EABB92AA09B26490BC323E4C85D816964562D48142A978BFDD7CA937DB20BC71820D7F679625DCB8FC11CBA24947AA727EBBE85C5F20A5EA81C86AF66E7F3D4BA8D3DE946769B5DDED02299B953DD279F36BE010543DE27BECCD6C79D466954B2018FCCEA80F4304F68F7515D116EB50D86DDAD45E9DB269D21E90104B7C7DA5AC99C1327E87F29E4F49278209E79833A38B842E99A346B5E8B7B2421B37F84D5BE4759A7F9A6850B1623246B8DA6A531313E4F23F14714B4CA10A703FCCBB1F7DDD358FCC416EB378A8A5C5D1C00637A766B04BE8F70672D620BE08CC11398346B1177FADD235BA8094C8FE0CCCB7103F02847972DEF797482A97A536441882E10F506D27C1B608AF86F4215D43B3DEECA7F4949E9F09E6C849099C544855C36E9F8ED36F47FE1C82B1E03511DC7DED98A2E905DD36226229AD4B5E098B820F15296888A751B95FAE79F78D5FAA9A62F8996FF5FA560942F49B71A990E0C00617511D0EF46CBE2B284D854988FAB8FB97EC432D8CCE1A70962BE36C545022DE9B8287DB55B1B72DB1D45BB500C2FCE6272B77332A8D41EB087A0DBC05BA78492690D2056A8CF4AC78262A6B4D4B5AE69B6986DEFC0F84A2CC4A0CB4DEBE660BA61D832E1093BAB32E4A6A4E314A87B6F3CE068FCB2D8CD3A5338ED3B1A8F3CA17BF2605ED833119AD73EEDEC131F42B835E213D0B7A207BB234AD0ECB6A812EECA946E9D52B37737B026495EE08F6D7C0440ACE4C1C6FC2DFA46AA6DBF02C03481BC86FE197BA7ED9331EBFF2D2C8D150B549F23486BE3859802158E086D1D1E68EF395703791CADF0BEC1F9F2A2DBC08238407F0FFBE63D13D1F779B6337609C1EA346597D177D0B4CD49AFE18A490EFF9F39205F8032F3E1F94418E0C296151AB025F463C24E4114808730DDBAD68CFBFA2BA94F54390D9A9183365475E1C836D201526DA7C200803702D7F1DC4838680D8E6B9EFF6AEAB49E366AC53DD0CDA887AC2721B64BAF9ADB8DBE6732B1865FE86AF2B4B3F713314735752D1F9EC05B6AA6CFB4EA1267AEF4EAAEEC30FD0598A6DE3BB6C45BBD091253DA9F2AAA77001D7089EA12712D2B675FC43E66EC8604025BB0EF18F60D704D4DC26AD84D4A4BC2B6E9E6A6C435F6CF87D90EAA3AAD6B3EA74F5CDF45872C73128C5B2F17D3D57125613A9611C53DCC26915DF140D98B10615CB5767A1E025BBE363CDF4A2F5CB52B2F0CC0201BD171F77D83B345B2686B967333896039A0FAB42C1B2E3FD0F26C3F59A3C51526C67DB34E823DA8894C2F1CC1BF48BB6986CDD84788D3AA481226D892688C333D7DF558EC978BD027A542ED131B5D673EDCA7BC327E3B05E5AF22A4557AE4FFF2C888CA5DF9C22633621791C10FE543E8D4AAB3FE516B3A394DBCCFA2914A26F01E8BA8C7F8F75A3C554D18CBD6A2F7B29A5682889C68EE5DF96C3FEC79B213FFBBE611B1CAB72A3BB1A52A09FAFD3201D958FD05835398A1EB9151ECCBF8AF244E306993A32E95CA4F39EB38C4FB39A1A0235131CFD95FED9EEC543D95F1AE8051421721F95A1491BC4EF2FA55149DAE8B4D94E25899D70152037A305112E080E6084D3C3F03CA89FBAA643B01B14992A36F887F78B1A2C5A11BB7827CBB1DD322BA40727FBAA9DAD1C9003B46ED6077B279258E07133D60517104005D8FD9AE8057CDC36D68BEE75A4ECCD366D24E82D7201D80C485577B2FA1D184BB29B578C85A5F4F49F17EC3A0E9199B1C439F341A2879B0C70FC305843F51F579F1F89C5BCB27514EB311A50D33E767688A9CD2C4F1D77DB5B4E55A9005583C2D5CEBE8413FA4A8E60BD9B4C45082138D227F1549C69F3978E607C8ABA8CC5F75A91D8BD4D43DDFF8B5155691D2E27A22CC92CB6E3BD4A9D25A028F9EC051DCA18D71484BB4DC39B12AC96B74A9F8D521E39BEB4094309C33500AD8CCE011E08307386574A6A9702EC0DD345DE81DE904112B80055E0AFFC2ED71D046F8D8859E5FE106EAD9E2C2733D0B112B4358B8071FFB11918913B1F3E22367BB756BCAFD37EFB99143E343AF80B&enc=1&play_duration={PLAYDURATION}&sid=__SID__]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://prod-ping.swat-adping.com/ping_server/play?viewid=313F111995C936548E538DB71BC0C541D3742981BEB7583E4A06BCB6D8383119A85F29C31F0F4040BA7D3AFE310345FDDBA68B92BC7112CC60D5EDAC1E08FA3BD4C0264805A63A24FD77E927083F832A49FC88F30D033A95DB30DB765AF1803B27F2AB75468FC0FE0C322AB2BEB0AAB3968E449B4A1776B7E905C83DFD95D3EA3C9CB16B383139D008BB1A9F1E47F7B1B94324517EAB8BD0C515BA6A33C06BE24846C602DEC38028554A7477F30E1C070F27B7002E56619DC469BD108C097324B9EABB92AA09B26490BC323E4C85D816964562D48142A978BFDD7CA937DB20BC71820D7F679625DCB8FC11CBA24947AA727EBBE85C5F20A5EA81C86AF66E7F3D4BA8D3DE946769B5DDED02299B953DD279F36BE010543DE27BECCD6C79D466954B2018FCCEA80F4304F68F7515D116EB50D86DDAD45E9DB269D21E90104B7C7DA5AC99C1327E87F29E4F49278209E79833A38B842E99A346B5E8B7B2421B37F84D5BE4759A7F9A6850B1623246B8DA6A531313E4F23F14714B4CA10A703FCCBB1F7DDD358FCC416EB378A8A5C5D1C00637A766B04BE8F70672D620BE08CC11398346B1177FADD235BA8094C8FE0CCCB7103F02847972DEF797482A97A536441882E10F506D27C1B608AF86F4215D43B3DEECA7F4949E9F09E6C849099C544855C36E9F8ED36F47FE1C82B1E03511DC7DED98A2E905DD36226229AD4B5E098B820F15296888A751B95FAE79F78D5FAA9A62F8996FF5FA560942F49B71A990E0C00617511D0EF46CBE2B284D854988FAB8FB97EC432D8CCE1A70962BE36C545022DE9B8287DB55B1B72DB1D45BB500C2FCE6272B77332A8D41EB087A0DBC05BA78492690D2056A8CF4AC78262A6B4D4B5AE69B6986DEFC0F84A2CC4A0CB4DEBE660BA61D832E1093BAB32E4A6A4E314A87B6F3CE068FCB2D8CD3A5338ED3B1A8F3CA17BF2605ED833119AD73EEDEC131F42B835E213D0B7A207BB234AD0ECB6A812EECA946E9D52B37737B026495EE08F6D7C0440ACE4C1C6FC2DFA46AA6DBF02C03481BC86FE197BA7ED9331EBFF2D2C8D150B549F23486BE3859802158E086D1D1E68EF395703791CADF0BEC1F9F2A2DBC08238407F0FFBE63D13D1F779B6337609C1EA346597D177D0B4CD49AFE18A490EFF9F39205F8032F3E1F94418E0C296151AB025F463C24E4114808730DDBAD68CFBFA2BA94F54390D9A9183365475E1C836D201526DA7C200803702D7F1DC4838680D8E6B9EFF6AEAB49E366AC53DD0CDA887AC2721B64BAF9ADB8DBE6732B1865FE86AF2B4B3F713314735752D1F9EC05B6AA6CFB4EA1267AEF4EAAEEC30FD0598A6DE3BB6C45BBD091253DA9F2AAA77001D7089EA12712D2B675FC43E66EC8604025BB0EF18F60D704D4DC26AD84D4A4BC2B6E9E6A6C435F6CF87D90EAA3AAD6B3EA74F5CDF45872C73128C5B2F17D3D57125613A9611C53DCC26915DF140D98B10615CB5767A1E025BBE363CDF4A2F5CB52B2F0CC0201BD171F77D83B345B2686B967333896039A0FAB42C1B2E3FD0F26C3F59A3C51526C67DB34E823DA8894C2F1CC1BF48BB6986CDD84788D3AA481226D892688C333D7DF558EC978BD027A542ED131B5D673EDCA7BC327E3B05E5AF22A4557AE4FFF2C888CA5DF9C22633621791C10FE543E8D4AAB3FE516B3A394DBCCFA2914A26F01E8BA8C7F8F75A3C554D18CBD6A2F7B29A5682889C68EE5DF96C3FEC79B213FFBBE611B1CAB72A3BB1A52A09FAFD3201D958FD05835398A1EB9151ECCBF8AF244E306993A32E95CA4F39EB38C4FB39A1A0235131CFD95FED9EEC543D95F1AE8051421721F95A1491BC4EF2FA55149DAE8B4D94E25899D70152037A305112E080E6084D3C3F03CA89FBAA643B01B14992A36F887F78B1A2C5A11BB7827CBB1DD322BA40727FBAA9DAD1C9003B46ED6077B279258E07133D60517104005D8FD9AE8057CDC36D68BEE75A4ECCD366D24E82D7201D80C485577B2FA1D184BB29B578C85A5F4F49F17EC3A0E9199B1C439F341A2879B0C70FC305843F51F579F1F89C5BCB27514EB311A50D33E767688A9CD2C4F1D77DB5B4E55A9005583C2D5CEBE8413FA4A8E60BD9B4C45082138D227F1549C69F3978E607C8ABA8CC5F75A91D8BD4D43DDFF8B5155691D2E27A22CC92CB6E3BD4A9D25A028F9EC051DCA18D71484BB4DC39B12AC96B74A9F8D521E39BEB4094309C33500AD8CCE011E08307386574A6A9702EC0DD345DE81DE904112B80055E0AFFC2ED71D046F8D8859E5FE106EAD9E2C2733D0B112B4358B8071FFB11918913B1F3E22C9F47DF5510FD3319458964E00D09F6B&enc=1&play_duration={PLAYDURATION}&sid=__SID__]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://prod-ping.swat-adping.com/ping_server/play?viewid=313F111995C936548E538DB71BC0C541D3742981BEB7583E4A06BCB6D8383119A85F29C31F0F4040BA7D3AFE310345FDDBA68B92BC7112CC60D5EDAC1E08FA3BD4C0264805A63A24FD77E927083F832A49FC88F30D033A95DB30DB765AF1803B27F2AB75468FC0FE0C322AB2BEB0AAB3968E449B4A1776B7E905C83DFD95D3EA3C9CB16B383139D008BB1A9F1E47F7B1B94324517EAB8BD0C515BA6A33C06BE24846C602DEC38028554A7477F30E1C070F27B7002E56619DC469BD108C097324B9EABB92AA09B26490BC323E4C85D816964562D48142A978BFDD7CA937DB20BC71820D7F679625DCB8FC11CBA24947AA727EBBE85C5F20A5EA81C86AF66E7F3D4BA8D3DE946769B5DDED02299B953DD279F36BE010543DE27BECCD6C79D466954B2018FCCEA80F4304F68F7515D116EB50D86DDAD45E9DB269D21E90104B7C7DA5AC99C1327E87F29E4F49278209E79833A38B842E99A346B5E8B7B2421B37F84D5BE4759A7F9A6850B1623246B8DA6A531313E4F23F14714B4CA10A703FCCBB1F7DDD358FCC416EB378A8A5C5D1C00637A766B04BE8F70672D620BE08CC11398346B1177FADD235BA8094C8FE0CCCB7103F02847972DEF797482A97A536441882E10F506D27C1B608AF86F4215D43B3DEECA7F4949E9F09E6C849099C544855C36E9F8ED36F47FE1C82B1E03511DC7DED98A2E905DD36226229AD4B5E098B820F15296888A751B95FAE79F78D5FAA9A62F8996FF5FA560942F49B71A990E0C00617511D0EF46CBE2B284D854988FAB8FB97EC432D8CCE1A70962BE36C545022DE9B8287DB55B1B72DB1D45BB500C2FCE6272B77332A8D41EB087A0DBC05BA78492690D2056A8CF4AC78262A6B4D4B5AE69B6986DEFC0F84A2CC4A0CB4DEBE660BA61D832E1093BAB32E4A6A4E314A87B6F3CE068FCB2D8CD3A5338ED3B1A8F3CA17BF2605ED833119AD73EEDEC131F42B835E213D0B7A207BB234AD0ECB6A812EECA946E9D52B37737B026495EE08F6D7C0440ACE4C1C6FC2DFA46AA6DBF02C03481BC86FE197BA7ED9331EBFF2D2C8D150B549F23486BE3859802158E086D1D1E68EF395703791CADF0BEC1F9F2A2DBC08238407F0FFBE63D13D1F779B6337609C1EA346597D177D0B4CD49AFE18A490EFF9F39205F8032F3E1F94418E0C296151AB025F463C24E4114808730DDBAD68CFBFA2BA94F54390D9A9183365475E1C836D201526DA7C200803702D7F1DC4838680D8E6B9EFF6AEAB49E366AC53DD0CDA887AC2721B64BAF9ADB8DBE6732B1865FE86AF2B4B3F713314735752D1F9EC05B6AA6CFB4EA1267AEF4EAAEEC30FD0598A6DE3BB6C45BBD091253DA9F2AAA77001D7089EA12712D2B675FC43E66EC8604025BB0EF18F60D704D4DC26AD84D4A4BC2B6E9E6A6C435F6CF87D90EAA3AAD6B3EA74F5CDF45872C73128C5B2F17D3D57125613A9611C53DCC26915DF140D98B10615CB5767A1E025BBE363CDF4A2F5CB52B2F0CC0201BD171F77D83B345B2686B967333896039A0FAB42C1B2E3FD0F26C3F59A3C51526C67DB34E823DA8894C2F1CC1BF48BB6986CDD84788D3AA481226D892688C333D7DF558EC978BD027A542ED131B5D673EDCA7BC327E3B05E5AF22A4557AE4FFF2C888CA5DF9C22633621791C10FE543E8D4AAB3FE516B3A394DBCCFA2914A26F01E8BA8C7F8F75A3C554D18CBD6A2F7B29A5682889C68EE5DF96C3FEC79B213FFBBE611B1CAB72A3BB1A52A09FAFD3201D958FD05835398A1EB9151ECCBF8AF244E306993A32E95CA4F39EB38C4FB39A1A0235131CFD95FED9EEC543D95F1AE8051421721F95A1491BC4EF2FA55149DAE8B4D94E25899D70152037A305112E080E6084D3C3F03CA89FBAA643B01B14992A36F887F78B1A2C5A11BB7827CBB1DD322BA40727FBAA9DAD1C9003B46ED6077B279258E07133D60517104005D8FD9AE8057CDC36D68BEE75A4ECCD366D24E82D7201D80C485577B2FA1D184BB29B578C85A5F4F49F17EC3A0E9199B1C439F341A2879B0C70FC305843F51F579F1F89C5BCB27514EB311A50D33E767688A9CD2C4F1D77DB5B4E55A9005583C2D5CEBE8413FA4A8E60BD9B4C45082138D227F1549C69F3978E607C8ABA8CC5F75A91D8BD4D43DDFF8B5155691D2E27A22CC92CB6E3BD4A9D25A028F9EC051DCA18D71484BB4DC39B12AC96B74A9F8D521E39BEB4094309C33500AD8CCE011E08307386574A6A9702EC0DD345DE81DE904112B80055E0AFFC2ED71D046F8D8859E5FE106EAD9E2C2733D0B112B4358B8071FFB11918913B1F3E2226F12CCC7E49997BA8330A89D655CD1A&enc=1&play_duration={PLAYDURATION}&sid=__SID__]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://prod-ping.swat-adping.com/ping_server/play?viewid=313F111995C936548E538DB71BC0C541D3742981BEB7583E4A06BCB6D8383119A85F29C31F0F4040BA7D3AFE310345FDDBA68B92BC7112CC60D5EDAC1E08FA3BD4C0264805A63A24FD77E927083F832A49FC88F30D033A95DB30DB765AF1803B27F2AB75468FC0FE0C322AB2BEB0AAB3968E449B4A1776B7E905C83DFD95D3EA3C9CB16B383139D008BB1A9F1E47F7B1B94324517EAB8BD0C515BA6A33C06BE24846C602DEC38028554A7477F30E1C070F27B7002E56619DC469BD108C097324B9EABB92AA09B26490BC323E4C85D816964562D48142A978BFDD7CA937DB20BC71820D7F679625DCB8FC11CBA24947AA727EBBE85C5F20A5EA81C86AF66E7F3D4BA8D3DE946769B5DDED02299B953DD279F36BE010543DE27BECCD6C79D466954B2018FCCEA80F4304F68F7515D116EB50D86DDAD45E9DB269D21E90104B7C7DA5AC99C1327E87F29E4F49278209E79833A38B842E99A346B5E8B7B2421B37F84D5BE4759A7F9A6850B1623246B8DA6A531313E4F23F14714B4CA10A703FCCBB1F7DDD358FCC416EB378A8A5C5D1C00637A766B04BE8F70672D620BE08CC11398346B1177FADD235BA8094C8FE0CCCB7103F02847972DEF797482A97A536441882E10F506D27C1B608AF86F4215D43B3DEECA7F4949E9F09E6C849099C544855C36E9F8ED36F47FE1C82B1E03511DC7DED98A2E905DD36226229AD4B5E098B820F15296888A751B95FAE79F78D5FAA9A62F8996FF5FA560942F49B71A990E0C00617511D0EF46CBE2B284D854988FAB8FB97EC432D8CCE1A70962BE36C545022DE9B8287DB55B1B72DB1D45BB500C2FCE6272B77332A8D41EB087A0DBC05BA78492690D2056A8CF4AC78262A6B4D4B5AE69B6986DEFC0F84A2CC4A0CB4DEBE660BA61D832E1093BAB32E4A6A4E314A87B6F3CE068FCB2D8CD3A5338ED3B1A8F3CA17BF2605ED833119AD73EEDEC131F42B835E213D0B7A207BB234AD0ECB6A812EECA946E9D52B37737B026495EE08F6D7C0440ACE4C1C6FC2DFA46AA6DBF02C03481BC86FE197BA7ED9331EBFF2D2C8D150B549F23486BE3859802158E086D1D1E68EF395703791CADF0BEC1F9F2A2DBC08238407F0FFBE63D13D1F779B6337609C1EA346597D177D0B4CD49AFE18A490EFF9F39205F8032F3E1F94418E0C296151AB025F463C24E4114808730DDBAD68CFBFA2BA94F54390D9A9183365475E1C836D201526DA7C200803702D7F1DC4838680D8E6B9EFF6AEAB49E366AC53DD0CDA887AC2721B64BAF9ADB8DBE6732B1865FE86AF2B4B3F713314735752D1F9EC05B6AA6CFB4EA1267AEF4EAAEEC30FD0598A6DE3BB6C45BBD091253DA9F2AAA77001D7089EA12712D2B675FC43E66EC8604025BB0EF18F60D704D4DC26AD84D4A4BC2B6E9E6A6C435F6CF87D90EAA3AAD6B3EA74F5CDF45872C73128C5B2F17D3D57125613A9611C53DCC26915DF140D98B10615CB5767A1E025BBE363CDF4A2F5CB52B2F0CC0201BD171F77D83B345B2686B967333896039A0FAB42C1B2E3FD0F26C3F59A3C51526C67DB34E823DA8894C2F1CC1BF48BB6986CDD84788D3AA481226D892688C333D7DF558EC978BD027A542ED131B5D673EDCA7BC327E3B05E5AF22A4557AE4FFF2C888CA5DF9C22633621791C10FE543E8D4AAB3FE516B3A394DBCCFA2914A26F01E8BA8C7F8F75A3C554D18CBD6A2F7B29A5682889C68EE5DF96C3FEC79B213FFBBE611B1CAB72A3BB1A52A09FAFD3201D958FD05835398A1EB9151ECCBF8AF244E306993A32E95CA4F39EB38C4FB39A1A0235131CFD95FED9EEC543D95F1AE8051421721F95A1491BC4EF2FA55149DAE8B4D94E25899D70152037A305112E080E6084D3C3F03CA89FBAA643B01B14992A36F887F78B1A2C5A11BB7827CBB1DD322BA40727FBAA9DAD1C9003B46ED6077B279258E07133D60517104005D8FD9AE8057CDC36D68BEE75A4ECCD366D24E82D7201D80C485577B2FA1D184BB29B578C85A5F4F49F17EC3A0E9199B1C439F341A2879B0C70FC305843F51F579F1F89C5BCB27514EB311A50D33E767688A9CD2C4F1D77DB5B4E55A9005583C2D5CEBE8413FA4A8E60BD9B4C45082138D227F1549C69F3978E607C8ABA8CC5F75A91D8BD4D43DDFF8B5155691D2E27A22CC92CB6E3BD4A9D25A028F9EC051DCA18D71484BB4DC39B12AC96B74A9F8D521E39BEB4094309C33500AD8CCE011E08307386574A6A9702EC0DD345DE81DE904112B80055E0AFFC2ED71D046F8D8859E5FE106EAD9E2C2733D0B112B4358B8071FFB11918913B1F3E223B193A4263168FF42100B98FF118D967E549F3F8DADAFF1C5676E8276335B376&enc=1&play_duration={PLAYDURATION}&sid=__SID__]]></Tracking><Tracking event=\"complete\"><![CDATA[https://prod-ping.swat-adping.com/ping_server/play?viewid=313F111995C936548E538DB71BC0C541D3742981BEB7583E4A06BCB6D8383119A85F29C31F0F4040BA7D3AFE310345FDDBA68B92BC7112CC60D5EDAC1E08FA3BD4C0264805A63A24FD77E927083F832A49FC88F30D033A95DB30DB765AF1803B27F2AB75468FC0FE0C322AB2BEB0AAB3968E449B4A1776B7E905C83DFD95D3EA3C9CB16B383139D008BB1A9F1E47F7B1B94324517EAB8BD0C515BA6A33C06BE24846C602DEC38028554A7477F30E1C070F27B7002E56619DC469BD108C097324B9EABB92AA09B26490BC323E4C85D816964562D48142A978BFDD7CA937DB20BC71820D7F679625DCB8FC11CBA24947AA727EBBE85C5F20A5EA81C86AF66E7F3D4BA8D3DE946769B5DDED02299B953DD279F36BE010543DE27BECCD6C79D466954B2018FCCEA80F4304F68F7515D116EB50D86DDAD45E9DB269D21E90104B7C7DA5AC99C1327E87F29E4F49278209E79833A38B842E99A346B5E8B7B2421B37F84D5BE4759A7F9A6850B1623246B8DA6A531313E4F23F14714B4CA10A703FCCBB1F7DDD358FCC416EB378A8A5C5D1C00637A766B04BE8F70672D620BE08CC11398346B1177FADD235BA8094C8FE0CCCB7103F02847972DEF797482A97A536441882E10F506D27C1B608AF86F4215D43B3DEECA7F4949E9F09E6C849099C544855C36E9F8ED36F47FE1C82B1E03511DC7DED98A2E905DD36226229AD4B5E098B820F15296888A751B95FAE79F78D5FAA9A62F8996FF5FA560942F49B71A990E0C00617511D0EF46CBE2B284D854988FAB8FB97EC432D8CCE1A70962BE36C545022DE9B8287DB55B1B72DB1D45BB500C2FCE6272B77332A8D41EB087A0DBC05BA78492690D2056A8CF4AC78262A6B4D4B5AE69B6986DEFC0F84A2CC4A0CB4DEBE660BA61D832E1093BAB32E4A6A4E314A87B6F3CE068FCB2D8CD3A5338ED3B1A8F3CA17BF2605ED833119AD73EEDEC131F42B835E213D0B7A207BB234AD0ECB6A812EECA946E9D52B37737B026495EE08F6D7C0440ACE4C1C6FC2DFA46AA6DBF02C03481BC86FE197BA7ED9331EBFF2D2C8D150B549F23486BE3859802158E086D1D1E68EF395703791CADF0BEC1F9F2A2DBC08238407F0FFBE63D13D1F779B6337609C1EA346597D177D0B4CD49AFE18A490EFF9F39205F8032F3E1F94418E0C296151AB025F463C24E4114808730DDBAD68CFBFA2BA94F54390D9A9183365475E1C836D201526DA7C200803702D7F1DC4838680D8E6B9EFF6AEAB49E366AC53DD0CDA887AC2721B64BAF9ADB8DBE6732B1865FE86AF2B4B3F713314735752D1F9EC05B6AA6CFB4EA1267AEF4EAAEEC30FD0598A6DE3BB6C45BBD091253DA9F2AAA77001D7089EA12712D2B675FC43E66EC8604025BB0EF18F60D704D4DC26AD84D4A4BC2B6E9E6A6C435F6CF87D90EAA3AAD6B3EA74F5CDF45872C73128C5B2F17D3D57125613A9611C53DCC26915DF140D98B10615CB5767A1E025BBE363CDF4A2F5CB52B2F0CC0201BD171F77D83B345B2686B967333896039A0FAB42C1B2E3FD0F26C3F59A3C51526C67DB34E823DA8894C2F1CC1BF48BB6986CDD84788D3AA481226D892688C333D7DF558EC978BD027A542ED131B5D673EDCA7BC327E3B05E5AF22A4557AE4FFF2C888CA5DF9C22633621791C10FE543E8D4AAB3FE516B3A394DBCCFA2914A26F01E8BA8C7F8F75A3C554D18CBD6A2F7B29A5682889C68EE5DF96C3FEC79B213FFBBE611B1CAB72A3BB1A52A09FAFD3201D958FD05835398A1EB9151ECCBF8AF244E306993A32E95CA4F39EB38C4FB39A1A0235131CFD95FED9EEC543D95F1AE8051421721F95A1491BC4EF2FA55149DAE8B4D94E25899D70152037A305112E080E6084D3C3F03CA89FBAA643B01B14992A36F887F78B1A2C5A11BB7827CBB1DD322BA40727FBAA9DAD1C9003B46ED6077B279258E07133D60517104005D8FD9AE8057CDC36D68BEE75A4ECCD366D24E82D7201D80C485577B2FA1D184BB29B578C85A5F4F49F17EC3A0E9199B1C439F341A2879B0C70FC305843F51F579F1F89C5BCB27514EB311A50D33E767688A9CD2C4F1D77DB5B4E55A9005583C2D5CEBE8413FA4A8E60BD9B4C45082138D227F1549C69F3978E607C8ABA8CC5F75A91D8BD4D43DDFF8B5155691D2E27A22CC92CB6E3BD4A9D25A028F9EC051DCA18D71484BB4DC39B12AC96B74A9F8D521E39BEB4094309C33500AD8CCE011E08307386574A6A9702EC0DD345DE81DE904112B80055E0AFFC2ED71D046F8D8859E5FE106EAD9E2C2733D0B112B4358B8071FFB11918913B1F3E22F7AB878FA15B51E89A341F297009B5AC&enc=1&play_duration={PLAYDURATION}&sid=__SID__]]></Tracking><Tracking event=\"resume\"><![CDATA[https://prod-ping.swat-adping.com/ping_server/play?viewid=313F111995C936548E538DB71BC0C541D3742981BEB7583E4A06BCB6D8383119A85F29C31F0F4040BA7D3AFE310345FDDBA68B92BC7112CC60D5EDAC1E08FA3BD4C0264805A63A24FD77E927083F832A49FC88F30D033A95DB30DB765AF1803B27F2AB75468FC0FE0C322AB2BEB0AAB3968E449B4A1776B7E905C83DFD95D3EA3C9CB16B383139D008BB1A9F1E47F7B1B94324517EAB8BD0C515BA6A33C06BE24846C602DEC38028554A7477F30E1C070F27B7002E56619DC469BD108C097324B9EABB92AA09B26490BC323E4C85D816964562D48142A978BFDD7CA937DB20BC71820D7F679625DCB8FC11CBA24947AA727EBBE85C5F20A5EA81C86AF66E7F3D4BA8D3DE946769B5DDED02299B953DD279F36BE010543DE27BECCD6C79D466954B2018FCCEA80F4304F68F7515D116EB50D86DDAD45E9DB269D21E90104B7C7DA5AC99C1327E87F29E4F49278209E79833A38B842E99A346B5E8B7B2421B37F84D5BE4759A7F9A6850B1623246B8DA6A531313E4F23F14714B4CA10A703FCCBB1F7DDD358FCC416EB378A8A5C5D1C00637A766B04BE8F70672D620BE08CC11398346B1177FADD235BA8094C8FE0CCCB7103F02847972DEF797482A97A536441882E10F506D27C1B608AF86F4215D43B3DEECA7F4949E9F09E6C849099C544855C36E9F8ED36F47FE1C82B1E03511DC7DED98A2E905DD36226229AD4B5E098B820F15296888A751B95FAE79F78D5FAA9A62F8996FF5FA560942F49B71A990E0C00617511D0EF46CBE2B284D854988FAB8FB97EC432D8CCE1A70962BE36C545022DE9B8287DB55B1B72DB1D45BB500C2FCE6272B77332A8D41EB087A0DBC05BA78492690D2056A8CF4AC78262A6B4D4B5AE69B6986DEFC0F84A2CC4A0CB4DEBE660BA61D832E1093BAB32E4A6A4E314A87B6F3CE068FCB2D8CD3A5338ED3B1A8F3CA17BF2605ED833119AD73EEDEC131F42B835E213D0B7A207BB234AD0ECB6A812EECA946E9D52B37737B026495EE08F6D7C0440ACE4C1C6FC2DFA46AA6DBF02C03481BC86FE197BA7ED9331EBFF2D2C8D150B549F23486BE3859802158E086D1D1E68EF395703791CADF0BEC1F9F2A2DBC08238407F0FFBE63D13D1F779B6337609C1EA346597D177D0B4CD49AFE18A490EFF9F39205F8032F3E1F94418E0C296151AB025F463C24E4114808730DDBAD68CFBFA2BA94F54390D9A9183365475E1C836D201526DA7C200803702D7F1DC4838680D8E6B9EFF6AEAB49E366AC53DD0CDA887AC2721B64BAF9ADB8DBE6732B1865FE86AF2B4B3F713314735752D1F9EC05B6AA6CFB4EA1267AEF4EAAEEC30FD0598A6DE3BB6C45BBD091253DA9F2AAA77001D7089EA12712D2B675FC43E66EC8604025BB0EF18F60D704D4DC26AD84D4A4BC2B6E9E6A6C435F6CF87D90EAA3AAD6B3EA74F5CDF45872C73128C5B2F17D3D57125613A9611C53DCC26915DF140D98B10615CB5767A1E025BBE363CDF4A2F5CB52B2F0CC0201BD171F77D83B345B2686B967333896039A0FAB42C1B2E3FD0F26C3F59A3C51526C67DB34E823DA8894C2F1CC1BF48BB6986CDD84788D3AA481226D892688C333D7DF558EC978BD027A542ED131B5D673EDCA7BC327E3B05E5AF22A4557AE4FFF2C888CA5DF9C22633621791C10FE543E8D4AAB3FE516B3A394DBCCFA2914A26F01E8BA8C7F8F75A3C554D18CBD6A2F7B29A5682889C68EE5DF96C3FEC79B213FFBBE611B1CAB72A3BB1A52A09FAFD3201D958FD05835398A1EB9151ECCBF8AF244E306993A32E95CA4F39EB38C4FB39A1A0235131CFD95FED9EEC543D95F1AE8051421721F95A1491BC4EF2FA55149DAE8B4D94E25899D70152037A305112E080E6084D3C3F03CA89FBAA643B01B14992A36F887F78B1A2C5A11BB7827CBB1DD322BA40727FBAA9DAD1C9003B46ED6077B279258E07133D60517104005D8FD9AE8057CDC36D68BEE75A4ECCD366D24E82D7201D80C485577B2FA1D184BB29B578C85A5F4F49F17EC3A0E9199B1C439F341A2879B0C70FC305843F51F579F1F89C5BCB27514EB311A50D33E767688A9CD2C4F1D77DB5B4E55A9005583C2D5CEBE8413FA4A8E60BD9B4C45082138D227F1549C69F3978E607C8ABA8CC5F75A91D8BD4D43DDFF8B5155691D2E27A22CC92CB6E3BD4A9D25A028F9EC051DCA18D71484BB4DC39B12AC96B74A9F8D521E39BEB4094309C33500AD8CCE011E08307386574A6A9702EC0DD345DE81DE904112B80055E0AFFC2ED71D046F8D8859E5FE106EAD9E2C2733D0B112B4358B8071FFB11918913B1F3E22890E6D20C3A4B107BB6037664546CCB9&enc=1&play_duration={PLAYDURATION}&sid=__SID__]]></Tracking></TrackingEvents><MediaFiles><MediaFile delivery=\"progressive\" type=\"video/mp4\" width=\"720\" height=\"1280\"><![CDATA[http://static.rqmob.com/sa/20221115/afcfa40ecc5ec9c0a9fad272ea442b7c.mp4]]></MediaFile></MediaFiles><VideoClicks><ClickThrough id=\"\"><![CDATA[http://midas.hellay.net/]]></ClickThrough><ClickTracking id=\"\"><![CDATA[https://prod-ping.swat-adping.com/ping_server/clk?viewid=313F111995C936548E538DB71BC0C541D3742981BEB7583E4A06BCB6D8383119A85F29C31F0F4040BA7D3AFE310345FDDBA68B92BC7112CC60D5EDAC1E08FA3BD4C0264805A63A24FD77E927083F832A49FC88F30D033A95DB30DB765AF1803B27F2AB75468FC0FE0C322AB2BEB0AAB3968E449B4A1776B7E905C83DFD95D3EA3C9CB16B383139D008BB1A9F1E47F7B1B94324517EAB8BD0C515BA6A33C06BE24846C602DEC38028554A7477F30E1C070F27B7002E56619DC469BD108C097324B9EABB92AA09B26490BC323E4C85D816964562D48142A978BFDD7CA937DB20BC71820D7F679625DCB8FC11CBA24947AA727EBBE85C5F20A5EA81C86AF66E7F3D4BA8D3DE946769B5DDED02299B953DD279F36BE010543DE27BECCD6C79D466954B2018FCCEA80F4304F68F7515D116EB50D86DDAD45E9DB269D21E90104B7C7DA5AC99C1327E87F29E4F49278209E79833A38B842E99A346B5E8B7B2421B37F84D5BE4759A7F9A6850B1623246B8DA6A531313E4F23F14714B4CA10A703FCCBB1F7DDD358FCC416EB378A8A5C5D1C00637A766B04BE8F70672D620BE08CC11398346B1177FADD235BA8094C8FE0CCCB7103F02847972DEF797482A97A536441882E10F506D27C1B608AF86F4215D43B3DEECA7F4949E9F09E6C849099C544855C36E9F8ED36F47FE1C82B1E03511DC7DED98A2E905DD36226229AD4B5E098B820F15296888A751B95FAE79F78D5FAA9A62F8996FF5FA560942F49B71A990E0C00617511D0EF46CBE2B284D854988FAB8FB97EC432D8CCE1A70962BE36C545022DE9B8287DB55B1B72DB1D45BB500C2FCE6272B77332A8D41EB087A0DBC05BA78492690D2056A8CF4AC78262A6B4D4B5AE69B6986DEFC0F84A2CC4A0CB4DEBE660BA61D832E1093BAB32E4A6A4E314A87B6F3CE068FCB2D8CD3A5338ED3B1A8F3CA17BF2605ED833119AD73EEDEC131F42B835E213D0B7A207BB234AD0ECB6A812EECA946E9D52B37737B026495EE08F6D7C0440ACE4C1C6FC2DFA46AA6DBF02C03481BC86FE197BA7ED9331EBFF2D2C8D150B549F23486BE3859802158E086D1D1E68EF395703791CADF0BEC1F9F2A2DBC08238407F0FFBE63D13D1F779B6337609C1EA346597D177D0B4CD49AFE18A490EFF9F39205F8032F3E1F94418E0C296151AB025F463C24E4114808730DDBAD68CFBFA2BA94F54390D9A9183365475E1C836D201526DA7C200803702D7F1DC4838680D8E6B9EFF6AEAB49E366AC53DD0CDA887AC2721B64BAF9ADB8DBE6732B1865FE86AF2B4B3F713314735752D1F9EC05B6AA6CFB4EA1267AEF4EAAEEC30FD0598A6DE3BB6C45BBD091253DA9F2AAA77001D7089EA12712D2B675FC43E66EC8604025BB0EF18F60D704D4DC26AD84D4A4BC2B6E9E6A6C435F6CF87D90EAA3AAD6B3EA74F5CDF45872C73128C5B2F17D3D57125613A9611C53DCC26915DF140D98B10615CB5767A1E025BBE363CDF4A2F5CB52B2F0CC0201BD171F77D83B345B2686B967333896039A0FAB42C1B2E3FD0F26C3F59A3C51526C67DB34E823DA8894C2F1CC1BF48BB6986CDD84788D3AA481226D892688C333D7DF558EC978BD027A542ED131B5D673EDCA7BC327E3B05E5AF22A4557AE4FFF2C888CA5DF9C22633621791C10FE543E8D4AAB3FE516B3A394DBCCFA2914A26F01E8BA8C7F8F75A3C554D18CBD6A2F7B29A5682889C68EE5DF96C3FEC79B213FFBBE611B1CAB72A3BB1A52A09FAFD3201D958FD05835398A1EB9151ECCBF8AF244E306993A32E95CA4F39EB38C4FB39A1A0235131CFD95FED9EEC543D95F1AE8051421721F95A1491BC4EF2FA55149DAE8B4D94E25899D70152037A305112E080E6084D3C3F03CA89FBAA643B01B14992A36F887F78B1A2C5A11BB7827CBB1DD322BA40727FBAA9DAD1C9003B46ED6077B279258E07133D60517104005D8FD9AE8057CDC36D68BEE75A4ECCD366D24E82D7201D80C485577B2FA1D184BB29B578C85A5F4F49F17EC3A0E9199B1C439F341A2879B0C70FC305843F51F579F1F89C5BCB27514EB311A50D33E767688A9CD2C4F1D77DB5B4E55A9005583C2D5CEBE8413FA4A8E60BD9B4C45082138D227F1549C69F3978E607C8ABA8CC5F75A91D8BD4D43DDFF8B5155691D2E27A22CC92CB6E3BD4A9D25A028F9EC051DCA18D71484BB4DC39B12AC96B74A9F8D521E39BEB4094309C33500AD8CCE011E08307386574A6A9702EC0DD345DE81DE904112B80055E0AFFC2ED71D046F8D8859E5FE106EAD9E2C2733D00A4E8D4A5778550E9041087809902330C9FA062339422DC2334AB9ABCC052E72&enc=1&auction_price=0.000000&sid=__SID__]]></ClickTracking><ClickTracking><![CDATA[https://test-track.lacunads.com/clk?viewid=5th4nW7N_Vcfu_iStyeze2GWRz8xXrx-3kk4Bh45kQQXV5wi6RZk61IpfSbTeI0og_Pjcd5NjgKJLF8xbUQpfPMnBpkvFIfv3t9Q6kdDvJUwo9KEWjhu07bP3CxPDxnh7EzUJAMH1DayUWdDG8y6e6lq-2HqNu9mi1bvIpUtf3EejnINfoqlobIVqHkOQRW4vF_A78PELz9G9AvmMUTex4iXGHzJ5iP-aWsJRr4oSg0Xvso6W8Hlgn_DFP3dC412j_r1jmJQ1-qsOb4hS4-faFgcTTkTc4AX4rrB3uCV6U-kFHgnAjHXDiQmGy-OWlhtg7FZWMo-kjO06r7q4fjb4ihbVa2Fqw3NlI4BHhKF__6lQVsonU6zpQn02Hmx8KmGtg8sjeSZpn2WamnEmg7TdBYY1w6vdyDWir09w18sRqUBead-XBXIHlnZTjnLArMjSQExsr_0NAARlZ17GIUnOC3FvIGE4KCt3VenFqDD67KtWQAEnqZ6uHVItt7eRn9oj8pHLSyd2Sn5MDIGTOVp10mHidwHjJdg23sBou0MIrnRYjH7D011TxH-GHvR_7xO1jcFaJSRcXlq_S7TQpXx3TM9RHVd4KBdeWy6SrZbXwOIBg7gZUmVDm4rZes2LEXstnZe_YjKlA4G4XTZkjZFRcoIF8JKG9J6Rru2ni4L7JAj1t-8gGh4oMSleMl_Ca1HMhAqOtKIhzzWE98cer9SI44MfhKyZXCjicj9cc86BbYieU3uHYlc7lCCrXEvVTZxAMtm4qUtRvHDQNd30ljYEs192nxsVTts3J7xX3dbtfK67y4TJOuljIIKhTVQse-Y1sBJn9FMwM5EIHsALMfF2LJRB9HICfS91dnvdL8ifdiVJOF53bH8TpZ3d4PbuzETAXDZtZp2cVikG516B7EA2cYUt1CqfXNSy1__UFX8Yxa1ZoNRjMel5pi9dPp9teocy8hwqp9uEaVvRReUFmSxJd7jY03_KcfqDQyiDZMejH5KDS5cZ0mrTdpvxblu-wJFxqF3B1rx5eS3MjBZqKqWb-smLHa0IHg_uoanZ0FN3F8Qbpm20wgGq4kX0vWDU0KP9a8AMkHOxWUzA5MSrabdrvNBVAP6yMd7ki4Kqj-ApCU1uIMb5QrGDrrxwccdjjhD64OYs5QYsZovsKd94ANdORBplr-_g4DzkcXZ50bjAofjdUrwDUXsHKRBooL0x0agZ7yHML5kEJ9gQ2LJ9k7LLpyU6fawaWg3uZ-sHJyzlLkA0VnXZ5tBp3m2kAKP12cn05oZYzOc1fD9eWS1EbQxynDcAR8B-aJ_7E5O3OQ9_ObCdlW0MjKTLy7kVw8ndT3owbLhEhyEqxnWvmQIXlGrdutvPY4K3dRbSfy6o7Wy7pPREcGC8mBFQIn56bNAhvmJFvpts_mxx9A9CCBviqznhDz7TjnE0B8wToeRGK0QIBwk6vbj15rL3C6e6x2anFuDEk-oxMy_qZMVbDjCI6-pIDNRH-QvQmV-PgpiKYxNU3slOg8bc_8DjQweRmLo6dNd_c5zarFrQ7d3nJDQVHxGUUA9CggW7OBoAaJqPXwCNVMYpx8bTlgWlAijnvskkgOqCfkAM_SEGBf7IAKzcEAjhi1DdcfNPs_c5wfvzCTa7qBHWVXKGiVu_oYeq_3I4CwN5gyCdZ8KBjtTrKcsjUpg3b9hyiBeIqZm0zV-MKvaKBmbmXr8w_kslO97RHLfQqJ5Ipqb85fbl5X85ex7_9T1NDJfWrfFOShF3HE2gveKXV0G-vAZHKOA6Tp-92JgpJJgWqaOUid3mkGBVktSOwG1pTac7ECrb6WoP4dA25YNbFpySH4Wxd7-CXfmU2Cvr94PgNhg6cE77LXYxX6C9oAN6ZXH5_WQziCJI98iQCvS-l2bWfn4LwdvGCyDXk9_ZimidkwO3aJx6XKWSskosS9YgFrU0kRREuow4h6bpUleL3rRwZAvGZG395DbAyfs6X1cm48WLraFSEjb9EkiMfz5fhMBEeB91OeayFs80eLjc3EhS_dgebnJlpjrpKiTY4QRlCNfZob-WqwsU319F6L1HJUyBcNksbj3AXmYiwrPSBUhEhoNOvbRRtDiyEbIhzu9HE7FLmecdhjUjQPFdNV1zZFte58fusgaPA5FEBpB5h96LEyzFyKLl6Gi3xe9QJ3jmoez9NvkS1OWp9z-eh1dJ3AYN8nS9XcYGzUGFze2uBsGG5PDyRQB1QWy-dDcfZtUG8LJKSqccyAVj3pI3E4CDmlYfby4wfX2qTWCjI1z5dFlQDNkrrX11gvSo0r0zzUtDiSw9UcVLmtua5tqOybetCd9CEstA-3Qkj0TpNYpJzchVg5GU5iUiZZlYmvsjjW1CUUJbqDwBsSvZeAvyLHSY86ZoyKeaYa0b61_Ymr-YA7iHHTXw1bVsOgH_0IEIPx98aU-TttUYYYiYqOtH6WOWrjH-XA1HYFuKdQn8EQGkEIKMXhJ7LiexRQhZzzzOIgha580XQgOtVdHdntAV50H1UJus_bHKo0RSzc7AhmA_fL1D6D3bPHujdsYKyvnre_ivVy5gzz_ZWqO-YyHpva1VTfd6AKijlrXv2OKRIbgQC83MZQ9TKEstT3aMXHG48NUYbyZeaNO04wAqbk_cL6sN3qBmpSxi4qrimag_PmW95Q6W3uSWAoCAVcWZDp23RLpqSL2W28Jk-WZVxiyNGYIXsIyFnkQG8BLnJVQa_yktwnuouPmCjbJYqshFHear0dZ2idJr3M7Y4-4DNIrj9Q5w6hbDhkOa8DOV2x0QPLsW6m50NzqFx4rvJVe5zUngOn-at_GTlfyj34JyEhiRJg6a1ndExthaIntFWqcwkSwLFeUdBVKZ3E4Qf1y9FzbeoRQ2-fUAFlgnC6ndVvBcayPc5lMRxRbWMaOYGDI8gzO2eD0XYubqarMu_fxpmqmPzgQz_f1bhqIuha9jm31sHRkv8Vc59oHkFXErMxzka87HMawCm7JCgc_U__DPdHn7RggcfhTxxnqJDad0xKSBx8WYJrN0hWCq2pUAKqORBkRn8reEVxWO0otUuG-oELznVlWQWfi32hhuyo3gdJZG1J_CL9f095JwoQ7Lv4mZxw17lEI0AOwhVaSwnpW4jf6cwZuLnv35oXTZw6Yn_QXgbogrB4IFV9J7hXKxX__aLAV7PdTF9r-YHKUzXlRFWASQyR4alxQYF5DpZAt3pg0WUEHUGZ83Zn00rmFjpfqr1sxMwwwgVYsbAijtb-RTmaIMjgk85dE85Jc-1X31xITGGRtGgfJdL7gh9rbzLkSNDcrbcaUjLAQE5EQV_XW8CdlZQ6tntqulJCM-k9QF3T1PSIBtMlWKr8qwrQ8aadOZkYK8u0DgAZYtvhzoyyF179dFIdiXeAhMAUfLuz8LdnY8RusQGYwg2BkZP-vMYwi0mYpMy5YsX0JMRC9NcSIWL3ye2OoxseNES5miI0r5oVBDBQqpoZwIFx0ZW5K9SanJHdtOskiDQ7XoWr9yWqug8rz4p7FX48VsPSwa9NVWQMHniR1Nd_8c7nTk671cX48DJSjbLClJx41wKlgdl6LDrVGHvhiOcSJ44q5ECaQ4DKk8JIxqT1ayFWFyG3YbX7U6F1j1LGXT4JAaoMyJMOnWNtIEwOHAx-jPn7T1xTOJtdzhWlb1a7jVwS7OJNQksBamXV50ehS-XuR-5zD2nhYooQiiJ8n-2muSfcYVggzi_IT42xbFXb4UEXHErmmsSCj_SjGiQRAu3S0EgUDb9W24YJZjdRQjr66kIA9O7mwNqfD9Ra9i5r89i9m1EBG7o2l_9PJ3bPn9rAfVX6Nja4SPwl7OIDswM-LxlNMpQDloby-__Qa2If0dTvt61zvJ4lJ6HGfUqs-q8pzY1jA_VJzCNkUR0xPgIdmI7fTJr2rGJFBJh0xw6V8bB5NLqoyjb7kRzpnrBjSoDMDmxnHeYVBfdE8cqwxmrUkGF-82M64GwwcgC2k15ylscZ_pkIcP9bsHKWOERcjjn4TV8RafsIB0ZG_lquV3tZjLDDYmfbFMbh3WluHqlATB9lh6XiYXPm9AfRdDYmh75alTb4FvQbSc_FCo3m0oxJcU3eIJdhxbYo_e-SIH-ScZPVAoxyfSZyyjNJCVHRqnuDqaRLyjPho-IcCfcUiu-tw_GqwQkH3kCD7g1_W3MYKo-lkZ7r817DBAJURdvYEnrpEaqOpZy6foEco5fmJNu8OFaRgbnncOkAKXFHoiKRb7Ky-CFLeHFxEXm2_Jk4W4NLGKWiwqDlkwcRpBZqtMes5LULDyllUylZRmSwKUIimNaJuy-SApIziUzalbur4ku44EbEpe1v2xnZZkvjcTPkcXP7KWJaw0pSLlJGEl6k_CWXW8H7Pd3-uzP25yVyVj-yPhd3PCT7IvsyQ9ukt5vtcQ-rF0kuhDDBW9mLBPkKVtdoSyWcaBsbGeDhZGXtm7AsYlXj8vHiGkCnEchiFQDXkuG8sQsSOEGLrhLlbfeJF3oICsNNEFnTCag030xtKcKw5PQQfUaxpNHQVcgiy_--uQqeNnitUZMBXxbl7vKne5aK3Xppo9CGMVptcLwMS9TZ3mTFwH93iCHCrXv-6KPy1jtZzgk9J19Xmc1XTwo7cR9X_EIifPcC6tbONY5HbJ_u9vnSdfEhD5xtv_w0-whYgq7gHOZx2AWthXhtk7p-3O9jXYoXu-7K-gQ0TsGFaS8O09bC3reEO26xASLIbXTy6nj_bwCT-qRn93mWsrWBChm4uci3CoanwauZZ7i6iiFU95iXOk4X5zEt6Ff7b7Uziq4gDj29EkXsN14SA59-mCtwEvPKKtCZfUIBpURyKqqm88Lo_98uFYc3iMtWdDXYfDU1PMv0Nw5kTU4rhhgPVT-N3D6mcWN57YBrJZSob1SbBHeQ6-Aa8Pb1vGQzbudCz1V9iEd4A6801RIP4jqCa2-YV-8TIBzmWLmD6pi-Y5FR-TGDmaars8FtfevNlZRlAKATW6r3yMuh_kE0elUL4wf1puVBK23h03zCl7VstrO9hBzUAM0MhUisE0ihlymHe8tq7W_OMnsRsXhn7ROTJ7AFnaUVEugD39YbazWmhWKgGdtFSFocEFLk_jSwGQ6prHcCu5Zeo9GhmGca4rUm3NzghrhRsLpCqKGuYApFFVBo2aTQe5QE8176xNjbP&enc=1]]></ClickTracking></VideoClicks></Linear></Creative></Creatives><Impression><![CDATA[https://prod-ping.swat-adping.com/ping_server/impl?viewid=313F111995C936548E538DB71BC0C541D3742981BEB7583E4A06BCB6D8383119A85F29C31F0F4040BA7D3AFE310345FDDBA68B92BC7112CC60D5EDAC1E08FA3BD4C0264805A63A24FD77E927083F832A49FC88F30D033A95DB30DB765AF1803B27F2AB75468FC0FE0C322AB2BEB0AAB3968E449B4A1776B7E905C83DFD95D3EA3C9CB16B383139D008BB1A9F1E47F7B1B94324517EAB8BD0C515BA6A33C06BE24846C602DEC38028554A7477F30E1C070F27B7002E56619DC469BD108C097324B9EABB92AA09B26490BC323E4C85D816964562D48142A978BFDD7CA937DB20BC71820D7F679625DCB8FC11CBA24947AA727EBBE85C5F20A5EA81C86AF66E7F3D4BA8D3DE946769B5DDED02299B953DD279F36BE010543DE27BECCD6C79D466954B2018FCCEA80F4304F68F7515D116EB50D86DDAD45E9DB269D21E90104B7C7DA5AC99C1327E87F29E4F49278209E79833A38B842E99A346B5E8B7B2421B37F84D5BE4759A7F9A6850B1623246B8DA6A531313E4F23F14714B4CA10A703FCCBB1F7DDD358FCC416EB378A8A5C5D1C00637A766B04BE8F70672D620BE08CC11398346B1177FADD235BA8094C8FE0CCCB7103F02847972DEF797482A97A536441882E10F506D27C1B608AF86F4215D43B3DEECA7F4949E9F09E6C849099C544855C36E9F8ED36F47FE1C82B1E03511DC7DED98A2E905DD36226229AD4B5E098B820F15296888A751B95FAE79F78D5FAA9A62F8996FF5FA560942F49B71A990E0C00617511D0EF46CBE2B284D854988FAB8FB97EC432D8CCE1A70962BE36C545022DE9B8287DB55B1B72DB1D45BB500C2FCE6272B77332A8D41EB087A0DBC05BA78492690D2056A8CF4AC78262A6B4D4B5AE69B6986DEFC0F84A2CC4A0CB4DEBE660BA61D832E1093BAB32E4A6A4E314A87B6F3CE068FCB2D8CD3A5338ED3B1A8F3CA17BF2605ED833119AD73EEDEC131F42B835E213D0B7A207BB234AD0ECB6A812EECA946E9D52B37737B026495EE08F6D7C0440ACE4C1C6FC2DFA46AA6DBF02C03481BC86FE197BA7ED9331EBFF2D2C8D150B549F23486BE3859802158E086D1D1E68EF395703791CADF0BEC1F9F2A2DBC08238407F0FFBE63D13D1F779B6337609C1EA346597D177D0B4CD49AFE18A490EFF9F39205F8032F3E1F94418E0C296151AB025F463C24E4114808730DDBAD68CFBFA2BA94F54390D9A9183365475E1C836D201526DA7C200803702D7F1DC4838680D8E6B9EFF6AEAB49E366AC53DD0CDA887AC2721B64BAF9ADB8DBE6732B1865FE86AF2B4B3F713314735752D1F9EC05B6AA6CFB4EA1267AEF4EAAEEC30FD0598A6DE3BB6C45BBD091253DA9F2AAA77001D7089EA12712D2B675FC43E66EC8604025BB0EF18F60D704D4DC26AD84D4A4BC2B6E9E6A6C435F6CF87D90EAA3AAD6B3EA74F5CDF45872C73128C5B2F17D3D57125613A9611C53DCC26915DF140D98B10615CB5767A1E025BBE363CDF4A2F5CB52B2F0CC0201BD171F77D83B345B2686B967333896039A0FAB42C1B2E3FD0F26C3F59A3C51526C67DB34E823DA8894C2F1CC1BF48BB6986CDD84788D3AA481226D892688C333D7DF558EC978BD027A542ED131B5D673EDCA7BC327E3B05E5AF22A4557AE4FFF2C888CA5DF9C22633621791C10FE543E8D4AAB3FE516B3A394DBCCFA2914A26F01E8BA8C7F8F75A3C554D18CBD6A2F7B29A5682889C68EE5DF96C3FEC79B213FFBBE611B1CAB72A3BB1A52A09FAFD3201D958FD05835398A1EB9151ECCBF8AF244E306993A32E95CA4F39EB38C4FB39A1A0235131CFD95FED9EEC543D95F1AE8051421721F95A1491BC4EF2FA55149DAE8B4D94E25899D70152037A305112E080E6084D3C3F03CA89FBAA643B01B14992A36F887F78B1A2C5A11BB7827CBB1DD322BA40727FBAA9DAD1C9003B46ED6077B279258E07133D60517104005D8FD9AE8057CDC36D68BEE75A4ECCD366D24E82D7201D80C485577B2FA1D184BB29B578C85A5F4F49F17EC3A0E9199B1C439F341A2879B0C70FC305843F51F579F1F89C5BCB27514EB311A50D33E767688A9CD2C4F1D77DB5B4E55A9005583C2D5CEBE8413FA4A8E60BD9B4C45082138D227F1549C69F3978E607C8ABA8CC5F75A91D8BD4D43DDFF8B5155691D2E27A22CC92CB6E3BD4A9D25A028F9EC051DCA18D71484BB4DC39B12AC96B74A9F8D521E39BEB4094309C33500AD8CCE011E08307386574A6A9702EC0DD345DE81DE904112B80055E0AFFC2ED71D046F8D8859E5FE106EAD9E2C2733D005D2598B2C4B49CB21E4DED63B8F7D058472B2341E7D82389F17FBB6548FCC61&enc=1&auction_price=0.000000&sid=__SID__]]></Impression><Survey><![CDATA[]]></Survey><Advertiser></Advertiser><Description><![CDATA[]]></Description><AdTitle><![CDATA[]]></AdTitle><AdSystem><![CDATA[ecom]]></AdSystem></InLine></Ad></VAST>",
										"adid":"1825467",
										"adomain":[
											"http://midas.hellay.net/"
										],
										"cid":"bb57db42f77807a9c5823bd8c2d9aaef",
										"crid":"23397081"
									}
								]
							}
						],
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_ecom_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
					"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
					"bidid":"de208481-248d-4bcc-8c98-7d14e43d5954",
					"seatbid":[
						{
							"bid":[
								{
									"id":"de208481-248d-4bcc-8c98-7d14e43d5954",
									"impid":"1",
									"price":1,
									"nurl":"https://test-track.lacunads.com/win?viewid=5th4nW7N_Vceu_iStyezex0BXgDdrNDI4nfhOYNRDQFo0N6qmsIhumL8HOP0CgcwV9MXiIzbyt1Djimz5Ql7CGSOHC-fg5jzOe8IHBXHVGa6iY_WVfCpqinOVdoeXs2GEMKW5yNlaME2deWsQfdGrrWvDOQEZrplkoQ3x0B9wUOY7ylJ0AVBhs1f3NyG0-PN2dOtERr8mzDqxzUMQ0ms9VgZeYPmzTLFD4xDZvCGBbsI-aMo_jgzCAHgeDEgo_PoGP3NizXz35m4azjrXeC9lyBCbQyn0CFMxFxJm0Oulsg0XjD5s3KFmYPxPZbMzWfQfCeGtXfH91uKokpz4W3cPq__SP_6UwH5q1duvLbs4XgxcKlk0McOCeGZpI2bRLziriA25Yk2GLJlzQOvjk9-mKwxgw8exaN8B9_jBpj9yaQXEbiluTaQvz_p3J_Sz2rhVNcc5s3G3SGyt6w8HV4VhSh1y5LynhkRyzid8Xxbipy2vcenxBq33QuTqdiSUD5FbDH6cg38Zp3hqFLHKSGsK5fhw38L56xhdm18qrwkZbWyUKSuwE3I8a3GjG5TnYrPXIZHzAt7tlCPCNcoad3Z2irExkJNH2KxrfIULRxdvO3Y9N9qbNWP2YkIUNaJJJ3lKXjbDooMkZBq0aht5GKvMul3WvKyujvuP6EIHLEj-Dk6K8LbynXYNBr8NSlbD_a-ZoveGoJpnC_3J8IV1UEjwpmtVxetT4jO1HlmP6_HR2uQf8cyS2nAU0OpmE4_YR2aqwxh6wmwSy1SyT4nO6c26htnW23tlkXNAn359ClKjcfRmN07hKOBuS1V1KDfAmKbdZcL3nKmq7G9Q9vylTFqmIiajZUFCpmJbFRE2kbDpYD4_RLWReNee5HNdsO52gTs9ikuMPNdVFH-UAKnA-gdgNm92axyLYFKYXNI-LLGPhjucwtmOeTQZ0b_ZUpIgub6sf2BhDINBVMvrYkRBE_4Rggyce_o19T5IKutun44XkmS-vrZWHr1etTlIfBdc__D67rAW-UG59AMnTRcWYjoLvT_fkWF7YMtZhzdRHzud2X3zeFOAibH62jLuzeI0cg4zdGvq7sWlqQv1E7b2IyDwhQUJNdqZgLew_Ksp4YMdyl2gCVJS8LyWqfQYEYRTHYhqYKRmVzcDKnsy4gOlJYQNZMaHI3Gwp08SzDgwAdxxYjaXbLOxl0VwTFsKBsCvzLPlsdb0GnlUg_RG-KKDxbho4nJ3vuNguKlzFOeDzftBHNJmTli6mNQZCg3szHfiAhFrtN7zG7mjXc8Sm8TizMuI7BBtLBHnHYDE0Oand2Mo2seT8sqliRtshprR2m6UHOIb2mjcU0ASIoAafQ4LyBiW60vc08TZHCydRjlPaRe_kOdu04mPjA1R7vmB3g7BY4oxhPBeydqH5YNpn2sYUC94m3ACrs33Qr_yI1uf-vuuPqisXpjNwZWVWEQUdlDOe4CqiHLUuZ3CQ_7rb99o4n0YRcX_BYrCqxZFxPWki9Aw_HFNHHtGn9aKNtoGZ2nQaV2LGclzOYTFwzpK9XsdgCqFwDXo4vmvIVQeEnD8li3Iaz8wkCZYIFtl5bA3laBInyCtnZkvOY1NQQOu9CQSVCP4EWFOJo9u0mGgrKXTaNTC4BSepw4hvPT79LVGy_JjWu5i3pN8VeOWiDHZBttMfvflHH4nw91S0ySWmzUkof_rO-zajKXL4QVEKBzviEL2B5zQKuvnlB_HuL7cHGBXIYNZbfHbSzd5XFaySEh3hydqbh5qTkNd2UDLEJVF8sacQ4CEeAd-14Jhf7tdTzpZuiU5xCpP72IbxXSUH7BLU8z4q9iEfDCAfhTz1Si9pxZWmcAg05ChUIzYUrwdi6r88YRo0hEYMPDCQJGO9gfczvKp10Y0ls8fxhrclE6Q4ruxy0CBaJvlE8eXWGwU4CFlsZa8HDVd9yLrPYHeGiEFR1D102LxQ6nSYae8eDGjefhjHjq1yGv0F_LxNKWUuiHbc-OaXQKFDHzT2NUehaugaTSsKTrd-8HhV6CokodnnsuS1Snc6ozsvtkmYz9jXZHw-GwpvAIOWT1FrYo8IAkZ4Y9u3ITrKtz9qwwox2b64d418umRk2o__6jDo7vwaZ63W47TrMBGoxa_dZP23QgnI8cRju1xmmKDlNofhyBsbcG2LWYhCbEZE7K67pKW-V3VXa4UImA0zL_Nf9DV2i0iPIZHhgeLg6voBY2YTjJR-GEK9642LfIjzASpAiikmUKzQPHgLokA2M0yJQOVpWI3POshDG_bgSfbWkxW8HrlKAyRXHwe1yFfT5P4HJXLRsGJX9kUiw8ENalSEx4mRMbijepMddEChofedAYKPMgq03gUz2MRRZ5CshNk--FOUgNO_d8c3Rn6GOOdJ6_fV-M-346yMXUZHgjAFiP6QHXJsMI6HfkZktwjTUjUNE5mBxefz7B-0C8jpOrMsZYlwqYDTaP7oY1UBSUUvnmHU7LjfXXgDhPcYNHjjzTgQMdhVpsYWkDydMe9NkA5XCTVskdcqziZoqaBQKFK-qY6lWIsXahyKLOa_Y3AG-XCeBRFV_qvBcXtAc8cfF1FgCysln7ppgUmaOoRDppoPRmKhWmXwi7qvF35NeHapdTUvujDIG3-_-ROIbvtvV6uYUQBKbhXkxWtovsRPtXrWIWvbP63-orVU31p42XtVbB-zGp7Ovv-_qk6gmKihbJYkDNXEL-0-onzFC4qLDIsSaQrKILcazfN8MQsbp0vPs80PPt4PliBbMhcHr3Gq30GC203rEmTplq3l-LBdQg9tPQrIo3wJy8XvhOocsg04YgLvGsPMQ0v6kntRjVe6WXqQgort6HPTmyhl3YYE-_ZGlX2Fn1aDOsJrogPRRGuwThbt2mN5eGdXv3qn4g8bs1sETNu-2OmgCS_BFsHcxYqN7YzsgL24RF1RBNt3YN7teR9YphuHt8E2XolbToJGDNUKOAmj3-9OtVHVTaD1z4pJmnbehUKF8Ww2FdCCYn14PdcFJwzi8aiKIHVcGs9kcmPGItbEMDnaEWlqin9SBG1RAwPuCeKRTSYumWRcHOpI0Iwj_6AvvCmboLvXuhkhq_11-g2lEszotd_HVZGhQ2oU3xwNGW1jsL3c3uoaGnmMjkLqliEYbuKvzzAhrMQmMO8_5oIUs9d_diSxZAdNcgvHojTmKBLIyFB4Ukaf8EK4OsUcxkZonSPj6ssmPlHcEoXxc3iGdducT1r8Z64FFsp2vmxfnvKFZu2fBZd3DuBA-J-sY4XkbKOeTO2rL36JLGtP81XK8X6qNl42xX6qNb7orvfe3-dud9LchNtAwH63HoXOi_ctTTesDQVZhend58STYBIgcFHwejnx0o_p1fddQACqmQ_vmxDxgAKCIs93Rqa1TyTmxKVmhSPlvR3_0R67qUJ-QjLCflZxIDFhlypbtwnHpp6na7MSvlCABMh_0ZNR2LJAaANU8fH0ngWhfgNUfhoMQyR67FjgYSa_2zp-WD_GYQ0fvH2VIXIWBHUI_2tgcH-ulCKT_We1vr6b-t3Gu6dHsNEo8MmxeXCrTZgwMNxtPU74SkRxAzXMNVhDP20ScPowu9dyhSFJH7VSBoXRjmFKpDUr4_3RviB8Ez5SsKFyIAoY4YDqgujtLEQOqYiOZafcdjIBgWEuXqAKSyPzE4BvejTMq7XBhaTBpwvOuXjEGFB5PORJeRWKpk1wDfcyLIKfB4bW7gCt3_ZY7tPWyj7DxIJe2RvC-2ToR8MVXdv2Lu5uQEXqc3bV2Iyu5Qm9yMBxvgdHxWCUmuo5FHP3rRsCfdtHY-sUgSZfACpqC8rlKkEyvvWUGahk7g9zXBZ2UMhKdzkYvV57wXouhS8Jvf3_OI44jWZya6K_NH_YRCD3KGF2HA0-vNu7pixbe7XGperlwFY9yq9kqRXs-ImeQcbd1qj1Re3kOngmPimfHETBp3uHeOF-BhBk1uypPtfKnEmYxFG4zrTsi7bmEjj29ZNFLi-tClt_mlQeLFcYEo-30KCP_a_KXDNyVf4Al3jaGdpSLxYRE3RvkVTXOsS8yUsce7jMlrXJW4HxX84bumMCdNIj5TYOUmVsaz9owTv6ATUFxDzw_voq34JJ7ONhgzacAtkKKafjATBZhmiA8eO4QW299XzttlIfPwUiKjYMRgJZjzvybwFeiJwwDlurv_UWcy_sjIV7_4EOO-XN7PzOV-jfPzONBe_kI9jR2lGVv8WAAf9VcwF4g6IpRmDThWe004QEvz01Ig6bq4fLW2OHb9OlXRl_dWinilxVqN2-0ykk95n6orKkUy3YkvirOJHYS4Vx_ZsAGBjCkt69Ta3Lgv2YH4RNWc-a4uCkLAKs2iKZ3FJMfXYAiOPxpB1SEBFpBWt0hj75I54p3MNyetoYt3yfi4zbWqt_SOr-p67CQJjzk1saq7wbuChN2A7pnczWlSdUilP62R1gP0N0PbZO0dfJZ5fjREDJVELmhrMRS2SKm5HrtVRIJEzWV6S1mTX1Qn4WlDz3YL6Y8IMIhfGYbe3Kh3JBKr_WwQK9jFM2-V1QWoTm9aqfpo-n7pjpPeFf2_-sDrAZ15IqdF8C5VcWO5pUOFv42HYkGruL9F8rFcTcL-WOYUn0_OGclutBilBWvWYB-kq4353YzYYw5e2ZL4XnGVhUoKHFlhCVsLOksvrf1mMX_uQ6XuinI_mGx3pV7526WxW_LvNG4Vmiw3KzPsbiRrLFarwKjl7_LTXLQ_gVJn-mBPS-oFUQT9sLc21pp7kU59XcqqbMq4ZB5g9cFOrRYbBfJIgyzqRqkAUIUwkShLwNno62glfNkEUsSu-z-6s7nWJn1wbgh1YfZdGK_tXHGnm2Iz9cEqqW5d1vdIgWYeAYl9Jew2koFchmDLzXT_UwqQLNmB27wRc1_65LBiwyS3aqbMxS9_zZXDGDZVSLRn5IXesVrYbr1ydLpWGnrwqT8mo2Xw15SddjZV8p3PhQchtca0u9YgdJuNea5LW8F3eyjPBLbb-uX3Xn6K3ZEbEVhcIQPiTwCnV5q-2dF5Pds3f5AtnKnWzjYzrzbn6ZQHM9bEaQNmdqAUUXpeoZ0UEUe8dWxWsSGuhfsFI9Oh1nDDlTLhE3opa4b3ujgylpmQi3ajZDHdRK9KMlSaj26KDdn2ytL0gbWFJiPBmu_5PuJDWDw=&enc=1&auction_price={AUCTION_PRICE}",
									"adm":"{\"native\":{\"ver\":\"1.0\",\"assets\":[{\"id\":1,\"img\":{\"type\":3,\"url\":\"https://static.rqmob.com/sa/20221118/e29f275fd2f535f17a6d779c87bb082c__FILE_CM____FILE_CM480_720__.gif\",\"w\":660,\"h\":346}}],\"link\":{\"url\":\"http://midas.hellay.net/\",\"clicktrackers\":[\"https://test-track.lacunads.com/clk?viewid=5th4nW7N_Vceu_iStyezex0BXgDdrNDI4nfhOYNRDQFo0N6qmsIhumL8HOP0CgcwV9MXiIzbyt1Djimz5Ql7CGSOHC-fg5jzOe8IHBXHVGa6iY_WVfCpqinOVdoeXs2GEMKW5yNlaME2deWsQfdGrrWvDOQEZrplkoQ3x0B9wUOY7ylJ0AVBhs1f3NyG0-PN2dOtERr8mzDqxzUMQ0ms9VgZeYPmzTLFD4xDZvCGBbsI-aMo_jgzCAHgeDEgo_PoGP3NizXz35m4azjrXeC9lyBCbQyn0CFMxFxJm0Oulsg0XjD5s3KFmYPxPZbMzWfQfCeGtXfH91uKokpz4W3cPq__SP_6UwH5q1duvLbs4XgxcKlk0McOCeGZpI2bRLziriA25Yk2GLJlzQOvjk9-mKwxgw8exaN8B9_jBpj9yaQXEbiluTaQvz_p3J_Sz2rhVNcc5s3G3SGyt6w8HV4VhSh1y5LynhkRyzid8Xxbipy2vcenxBq33QuTqdiSUD5FbDH6cg38Zp3hqFLHKSGsK5fhw38L56xhdm18qrwkZbWyUKSuwE3I8a3GjG5TnYrPXIZHzAt7tlCPCNcoad3Z2irExkJNH2KxrfIULRxdvO3Y9N9qbNWP2YkIUNaJJJ3lKXjbDooMkZBq0aht5GKvMul3WvKyujvuP6EIHLEj-Dk6K8LbynXYNBr8NSlbD_a-ZoveGoJpnC_3J8IV1UEjwpmtVxetT4jO1HlmP6_HR2uQf8cyS2nAU0OpmE4_YR2aqwxh6wmwSy1SyT4nO6c26htnW23tlkXNAn359ClKjcfRmN07hKOBuS1V1KDfAmKbdZcL3nKmq7G9Q9vylTFqmIiajZUFCpmJbFRE2kbDpYD4_RLWReNee5HNdsO52gTs9ikuMPNdVFH-UAKnA-gdgNm92axyLYFKYXNI-LLGPhjucwtmOeTQZ0b_ZUpIgub6sf2BhDINBVMvrYkRBE_4Rggyce_o19T5IKutun44XkmS-vrZWHr1etTlIfBdc__D67rAW-UG59AMnTRcWYjoLvT_fkWF7YMtZhzdRHzud2X3zeFOAibH62jLuzeI0cg4zdGvq7sWlqQv1E7b2IyDwhQUJNdqZgLew_Ksp4YMdyl2gCVJS8LyWqfQYEYRTHYhqYKRmVzcDKnsy4gOlJYQNZMaHI3Gwp08SzDgwAdxxYjaXbLOxl0VwTFsKBsCvzLPlsdb0GnlUg_RG-KKDxbho4nJ3vuNguKlzFOeDzftBHNJmTli6mNQZCg3szHfiAhFrtN7zG7mjXc8Sm8TizMuI7BBtLBHnHYDE0Oand2Mo2seT8sqliRtshprR2m6UHOIb2mjcU0ASIoAafQ4LyBiW60vc08TZHCydRjlPaRe_kOdu04mPjA1R7vmB3g7BY4oxhPBeydqH5YNpn2sYUC94m3ACrs33Qr_yI1uf-vuuPqisXpjNwZWVWEQUdlDOe4CqiHLUuZ3CQ_7rb99o4n0YRcX_BYrCqxZFxPWki9Aw_HFNHHtGn9aKNtoGZ2nQaV2LGclzOYTFwzpK9XsdgCqFwDXo4vmvIVQeEnD8li3Iaz8wkCZYIFtl5bA3laBInyCtnZkvOY1NQQOu9CQSVCP4EWFOJo9u0mGgrKXTaNTC4BSepw4hvPT79LVGy_JjWu5i3pN8VeOWiDHZBttMfvflHH4nw91S0ySWmzUkof_rO-zajKXL4QVEKBzviEL2B5zQKuvnlB_HuL7cHGBXIYNZbfHbSzd5XFaySEh3hydqbh5qTkNd2UDLEJVF8sacQ4CEeAd-14Jhf7tdTzpZuiU5xCpP72IbxXSUH7BLU8z4q9iEfDCAfhTz1Si9pxZWmcAg05ChUIzYUrwdi6r88YRo0hEYMPDCQJGO9gfczvKp10Y0ls8fxhrclE6Q4ruxy0CBaJvlE8eXWGwU4CFlsZa8HDVd9yLrPYHeGiEFR1D102LxQ6nSYae8eDGjefhjHjq1yGv0F_LxNKWUuiHbc-OaXQKFDHzT2NUehaugaTSsKTrd-8HhV6CokodnnsuS1Snc6ozsvtkmYz9jXZHw-GwpvAIOWT1FrYo8IAkZ4Y9u3ITrKtz9qwwox2b64d418umRk2o__6jDo7vwaZ63W47TrMBGoxa_dZP23QgnI8cRju1xmmKDlNofhyBsbcG2LWYhCbEZE7K67pKW-V3VXa4UImA0zL_Nf9DV2i0iPIZHhgeLg6voBY2YTjJR-GEK9642LfIjzASpAiikmUKzQPHgLokA2M0yJQOVpWI3POshDG_bgSfbWkxW8HrlKAyRXHwe1yFfT5P4HJXLRsGJX9kUiw8ENalSEx4mRMbijepMddEChofedAYKPMgq03gUz2MRRZ5CshNk--FOUgNO_d8c3Rn6GOOdJ6_fV-M-346yMXUZHgjAFiP6QHXJsMI6HfkZktwjTUjUNE5mBxefz7B-0C8jpOrMsZYlwqYDTaP7oY1UBSUUvnmHU7LjfXXgDhPcYNHjjzTgQMdhVpsYWkDydMe9NkA5XCTVskdcqziZoqaBQKFK-qY6lWIsXahyKLOa_Y3AG-XCeBRFV_qvBcXtAc8cfF1FgCysln7ppgUmaOoRDppoPRmKhWmXwi7qvF35NeHapdTUvujDIG3-_-ROIbvtvV6uYUQBKbhXkxWtovsRPtXrWIWvbP63-orVU31p42XtVbB-zGp7Ovv-_qk6gmKihbJYkDNXEL-0-onzFC4qLDIsSaQrKILcazfN8MQsbp0vPs80PPt4PliBbMhcHr3Gq30GC203rEmTplq3l-LBdQg9tPQrIo3wJy8XvhOocsg04YgLvGsPMQ0v6kntRjVe6WXqQgort6HPTmyhl3YYE-_ZGlX2Fn1aDOsJrogPRRGuwThbt2mN5eGdXv3qn4g8bs1sETNu-2OmgCS_BFsHcxYqN7YzsgL24RF1RBNt3YN7teR9YphuHt8E2XolbToJGDNUKOAmj3-9OtVHVTaD1z4pJmnbehUKF8Ww2FdCCYn14PdcFJwzi8aiKIHVcGs9kcmPGItbEMDnaEWlqin9SBG1RAwPuCeKRTSYumWRcHOpI0Iwj_6AvvCmboLvXuhkhq_11-g2lEszotd_HVZGhQ2oU3xwNGW1jsL3c3uoaGnmMjkLqliEYbuKvzzAhrMQmMO8_5oIUs9d_diSxZAdNcgvHojTmKBLIyFB4Ukaf8EK4OsUcxkZonSPj6ssmPlHcEoXxc3iGdducT1r8Z64FFsp2vmxfnvKFZu2fBZd3DuBA-J-sY4XkbKOeTO2rL36JLGtP81XK8X6qNl42xX6qNb7orvfe3-dud9LchNtAwH63HoXOi_ctTTesDQVZhend58STYBIgcFHwejnx0o_p1fddQACqmQ_vmxDxgAKCIs93Rqa1TyTmxKVmhSPlvR3_0R67qUJ-QjLCflZxIDFhlypbtwnHpp6na7MSvlCABMh_0ZNR2LJAaANU8fH0ngWhfgNUfhoMQyR67FjgYSa_2zp-WD_GYQ0fvH2VIXIWBHUI_2tgcH-ulCKT_We1vr6b-t3Gu6dHsNEo8MmxeXCrTZgwMNxtPU74SkRxAzXMNVhDP20ScPowu9dyhSFJH7VSBoXRjmFKpDUr4_3RviB8Ez5SsKFyIAoY4YDqgujtLEQOqYiOZafcdjIBgWEuXqAKSyPzE4BvejTMq7XBhaTBpwvOuXjEGFB5PORJeRWKpk1wDfcyLIKfB4bW7gCt3_ZY7tPWyj7DxIJe2RvC-2ToR8MVXdv2Lu5uQEXqc3bV2Iyu5Qm9yMBxvgdHxWCUmuo5FHP3rRsCfdtHY-sUgSZfACpqC8rlKkEyvvWUGahk7g9zXBZ2UMhKdzkYvV57wXouhS8Jvf3_OI44jWZya6K_NH_YRCD3KGF2HA0-vNu7pixbe7XGperlwFY9yq9kqRXs-ImeQcbd1qj1Re3kOngmPimfHETBp3uHeOF-BhBk1uypPtfKnEmYxFG4zrTsi7bmEjj29ZNFLi-tClt_mlQeLFcYEo-30KCP_a_KXDNyVf4Al3jaGdpSLxYRE3RvkVTXOsS8yUsce7jMlrXJW4HxX84bumMCdNIj5TYOUmVsaz9owTv6ATUFxDzw_voq34JJ7ONhgzacAtkKKafjATBZhmiA8eO4QW299XzttlIfPwUiKjYMRgJZjzvybwFeiJwwDlurv_UWcy_sjIV7_4EOO-XN7PzOV-jfPzONBe_kI9jR2lGVv8WAAf9VcwF4g6IpRmDThWe004QEvz01Ig6bq4fLW2OHb9OlXRl_dWinilxVqN2-0ykk95n6orKkUy3YkvirOJHYS4Vx_ZsAGBjCkt69Ta3Lgv2YH4RNWc-a4uCkLAKs2iKZ3FJMfXYAiOPxpB1SEBFpBWt0hj75I54p3MNyetoYt3yfi4zbWqt_SOr-p67CQJjzk1saq7wbuChN2A7pnczWlSdUilP62R1gP0N0PbZO0dfJZ5fjREDJVELmhrMRS2SKm5HrtVRIJEzWV6S1mTX1Qn4WlDz3YL6Y8IMIhfGYbe3Kh3JBKr_WwQK9jFM2-V1QWoTm9aqfpo-n7pjpPeFf2_-sDrAZ15IqdF8C5VcWO5pUOFv42HYkGruL9F8rFcTcL-WOYUn0_OGclutBilBWvWYB-kq4353YzYYw5e2ZL4XnGVhUoKHFlhCVsLOksvrf1mMX_uQ6XuinI_mGx3pV7526WxW_LvNG4Vmiw3KzPsbiRrLFarwKjl7_LTXLQ_gVJn-mBPS-oFUQT9sLc21pp7kU59XcqqbMq4ZB5g9cFOrRYbBfJIgyzqRqkAUIUwkShLwNno62glfNkEUsSu-z-6s7nWJn1wbgh1YfZdGK_tXHGnm2Iz9cEqqW5d1vdIgWYeAYl9Jew2koFchmDLzXT_UwqQLNmB27wRc1_65LBiwyS3aqbMxS9_zZXDGDZVSLRn5IXesVrYbr1ydLpWGnrwqT8mo2Xw15SddjZV8p3PhQchtca0u9YgdJuNea5LW8F3eyjPBLbb-uX3Xn6K3ZEbEVhcIQPiTwCnV5q-2dF5Pds3f5AtnKnWzjYzrzbn6ZQHM9bEaQNmdqAUUXpeoZ0UEUe8dWxWsSGuhfsFI9Oh1nDDlTLhE3opa4b3ujgylpmQi3ajZDHdRK9KMlSaj26KDdn2ytL0gbWFJiPBmu_5PuJDWDw=\&enc=1\",\"https://prod-ping.swat-adping.com/ping_server/clk?viewid=EE86E0B1DD0B6E55E9D6215783ACBFB2AF94041DACDAD6AA30659D8AD2657E7EA388649211E44CE0552239710DDACE1EBD89D35F4D20D6A7EDC5E058D2B9D33D9CC2267B52D8E9C5CB74CA6DE3193BA42C19A2B9C421369D01A8B9598B541BCEA69E5217D58C098C2E93767A437474FE875FE01404CF3922A41CCFF7BE8E71DC627C875341EFAE4AB12B370171583EA079ED03773C041524798DAEC1D588A558E8888CCD7542E6945B4BCFBB72F80E624D2A38911BCE1188CFB06FC10F36757AF812DF2E0F24AC9E3A47F3BE07A1EE7D3704DA5679A03BBE5C15F5EF29113E1579508043E7C0C2B59AEF14F59B07D637B89B54DCA87690CB48EB0E0CB7187F0592D6942A6145930301668866FB3E30910D0C896976919CA5440DF69B31DEBBF48EDB040DAA29AE8F8A861FF61340CB0AAB8B1CE5BA5D8CF08F1925F433703134911864AA08D1F87526E189659384CF88CE5565756F99866B1956DF2D791901F670C6D5BE96F5FE687B186E2A3126A01CE46ECE4083EBF118BF9122C84DA2AAF4741BF84E125ABB48FC2F46069B0D66E9C770A52CD3152F2B68AE4D814EB9FE6D295DA8F75B63C3109805B601F19F000CE9F58B613ED3594F6B814A78CF43AB76212E5461308D93E2EE806DB183D8A5D3EAE247E4078290C8F5305BC0229DD5EC189792B3134F2C74186C47C63D72560A3EEE74653E09E93E09772F19D6846A10CB3BDC1161C6B6BD28C54FD98AFE94D1C047152F6118D73703A4EF2B3A3E77ADEDDB5A95E758C94959B5E5FE33DA0E19DFCA64B53A097920B93B81658F0567539BE90FD3609EE710585D12500C8DDCC37483623429B6A0ED2A6202A087B1B8370DCC2ECB6DB4B45CBED5A15F33930659B61881F4E5709549267CBF5E81FBB3CE28F0FFC99D02A10CFE208D3E9C8EB820F6E18A9E7DDA6ABD7EB30B1A691322B9B64038691A4F308B4B9CDFB2336F98F565858761CB4B274C1E57E3F20C7F23DAADAD52D3EB5CA5A128F10A754D0125881D94C6733FBEF1840A9689E5CF2F860B15E31AC9DF36A8026B8C7AAB1ABC97513E50EEBC6B3FA4B8A57F2312A6DEA79DF34C46CB35D0C3C938A09B49E615E240FBE1868F0267C5B9B7DE62929D14C10C2EE68AE2B700985D71984792676093C555A27A8BD03D0DE4A22C9A39F9707B5F015F9AEDCC6A5F268A9AE9515B28D82F1C5C14A1D531BC1B15CC9FBA24A579D0B367EAE72C8A5490127DBBC078486279424DCD9F3AD82C3D201DAD92227B9459FFF59C824BF24B75B102DA6CDFD9506FCED011A7E9277DF72B97C84A3113F54701DEC64892A2C33D5B817500B510E69CF11CAB8EB6C40EB1B5D82AC3C7C21F5277D8452813BEA9D43D9059A8FD8F6947F256E4D990D3AA68BFD24865CF9681EC1C36E6A3680C99A0F78A649DDFDEBE6977FC674F73ECF422B4C01F38F4B1FBEAB2011437904E12049D055D934046F2C3C1943A559F1E595C0374785A3F4324B9752F5D8647D6A25B40D8C92B1D022CC500DECC82DBF9B365074869DAD611D78050DD2816D0F6F8DE0D4D6750627F09DECE35475A189623B3153BF04083FC5796DD1E0F15250F9612BD011E3E9EA660663D7A7B4A4AF7EDD4AD296628A458BE8C0A197A07401D02D7E19A51E03A5F0F4F75FCE223A356C6F7B31995E8FEBC534071F9F920D6A07D2FE145E4C70A606C96C3DCF716BF39A03BDB071FB39EAC45FCBCD7FF458A7A32B2ECD9679449451C74B059FE18AC7EF85A068F5BAEBE60EAC0E6849C23CB6E57E3EFF371725BEA9A198758620A3105F34078A3CAEE07FFB072DD22BBDDC2C97C66326D6035F31A3452170FDDE575B25ADB1B1B1637620DA331583A15376652CEF2B24290BFD2C358E4B69522D6E14281D11D6DF9E01C4853EBEA61A9493DDD48FE90551F74BC5DA9E3E3B165A9DBD2FBEC12D7B101C2C4DEEEE18EA35EB3436A79B84FBE1AB2560F0CBD4C10C9416710DDEAE2380C126280539399491BE5D9B8C703FA9EB54B62F927EFEA79188298A4DC2C7761F7712AA4D40955F866F941F9CB625FC3622BBC123CAFD6D0AD87763937A06CAB409DF8CDD60FDF0B552A4D9660A968E627802593579927CF1AAE6EC13E4AB93A891D412395A15E9871E7BD4DED3F022AA703982D26B3EF7EC5AA1087833BBD066EFCA82F2ACD801BD166F39C3ED5EAC1CB3A9FE44DE452A6EAE6EAA33A978C6DB8DCD1A02A22B3E226CCEB1EE3C36AE7C34AE014BF4223571D692F27F5477A357C5120FD2A74FFF75500B265991A2C8612547A5B550B9379E4E07580F6C7526114B8275043F64C79E3F46FBBD39078180A0B31588E\&enc=1\&auction_price=0.000000\&sid=__SID__\"]},\"imptrackers\":[\"https://test-track.lacunads.com/imp?viewid=5th4nW7N_Vceu_iStyezex0BXgDdrNDI4nfhOYNRDQFo0N6qmsIhumL8HOP0CgcwV9MXiIzbyt1Djimz5Ql7CGSOHC-fg5jzOe8IHBXHVGa6iY_WVfCpqinOVdoeXs2GEMKW5yNlaME2deWsQfdGrrWvDOQEZrplkoQ3x0B9wUOY7ylJ0AVBhs1f3NyG0-PN2dOtERr8mzDqxzUMQ0ms9VgZeYPmzTLFD4xDZvCGBbsI-aMo_jgzCAHgeDEgo_PoGP3NizXz35m4azjrXeC9lyBCbQyn0CFMxFxJm0Oulsg0XjD5s3KFmYPxPZbMzWfQfCeGtXfH91uKokpz4W3cPq__SP_6UwH5q1duvLbs4XgxcKlk0McOCeGZpI2bRLziriA25Yk2GLJlzQOvjk9-mKwxgw8exaN8B9_jBpj9yaQXEbiluTaQvz_p3J_Sz2rhVNcc5s3G3SGyt6w8HV4VhSh1y5LynhkRyzid8Xxbipy2vcenxBq33QuTqdiSUD5FbDH6cg38Zp3hqFLHKSGsK5fhw38L56xhdm18qrwkZbWyUKSuwE3I8a3GjG5TnYrPXIZHzAt7tlCPCNcoad3Z2irExkJNH2KxrfIULRxdvO3Y9N9qbNWP2YkIUNaJJJ3lKXjbDooMkZBq0aht5GKvMul3WvKyujvuP6EIHLEj-Dk6K8LbynXYNBr8NSlbD_a-ZoveGoJpnC_3J8IV1UEjwpmtVxetT4jO1HlmP6_HR2uQf8cyS2nAU0OpmE4_YR2aqwxh6wmwSy1SyT4nO6c26htnW23tlkXNAn359ClKjcfRmN07hKOBuS1V1KDfAmKbdZcL3nKmq7G9Q9vylTFqmIiajZUFCpmJbFRE2kbDpYD4_RLWReNee5HNdsO52gTs9ikuMPNdVFH-UAKnA-gdgNm92axyLYFKYXNI-LLGPhjucwtmOeTQZ0b_ZUpIgub6sf2BhDINBVMvrYkRBE_4Rggyce_o19T5IKutun44XkmS-vrZWHr1etTlIfBdc__D67rAW-UG59AMnTRcWYjoLvT_fkWF7YMtZhzdRHzud2X3zeFOAibH62jLuzeI0cg4zdGvq7sWlqQv1E7b2IyDwhQUJNdqZgLew_Ksp4YMdyl2gCVJS8LyWqfQYEYRTHYhqYKRmVzcDKnsy4gOlJYQNZMaHI3Gwp08SzDgwAdxxYjaXbLOxl0VwTFsKBsCvzLPlsdb0GnlUg_RG-KKDxbho4nJ3vuNguKlzFOeDzftBHNJmTli6mNQZCg3szHfiAhFrtN7zG7mjXc8Sm8TizMuI7BBtLBHnHYDE0Oand2Mo2seT8sqliRtshprR2m6UHOIb2mjcU0ASIoAafQ4LyBiW60vc08TZHCydRjlPaRe_kOdu04mPjA1R7vmB3g7BY4oxhPBeydqH5YNpn2sYUC94m3ACrs33Qr_yI1uf-vuuPqisXpjNwZWVWEQUdlDOe4CqiHLUuZ3CQ_7rb99o4n0YRcX_BYrCqxZFxPWki9Aw_HFNHHtGn9aKNtoGZ2nQaV2LGclzOYTFwzpK9XsdgCqFwDXo4vmvIVQeEnD8li3Iaz8wkCZYIFtl5bA3laBInyCtnZkvOY1NQQOu9CQSVCP4EWFOJo9u0mGgrKXTaNTC4BSepw4hvPT79LVGy_JjWu5i3pN8VeOWiDHZBttMfvflHH4nw91S0ySWmzUkof_rO-zajKXL4QVEKBzviEL2B5zQKuvnlB_HuL7cHGBXIYNZbfHbSzd5XFaySEh3hydqbh5qTkNd2UDLEJVF8sacQ4CEeAd-14Jhf7tdTzpZuiU5xCpP72IbxXSUH7BLU8z4q9iEfDCAfhTz1Si9pxZWmcAg05ChUIzYUrwdi6r88YRo0hEYMPDCQJGO9gfczvKp10Y0ls8fxhrclE6Q4ruxy0CBaJvlE8eXWGwU4CFlsZa8HDVd9yLrPYHeGiEFR1D102LxQ6nSYae8eDGjefhjHjq1yGv0F_LxNKWUuiHbc-OaXQKFDHzT2NUehaugaTSsKTrd-8HhV6CokodnnsuS1Snc6ozsvtkmYz9jXZHw-GwpvAIOWT1FrYo8IAkZ4Y9u3ITrKtz9qwwox2b64d418umRk2o__6jDo7vwaZ63W47TrMBGoxa_dZP23QgnI8cRju1xmmKDlNofhyBsbcG2LWYhCbEZE7K67pKW-V3VXa4UImA0zL_Nf9DV2i0iPIZHhgeLg6voBY2YTjJR-GEK9642LfIjzASpAiikmUKzQPHgLokA2M0yJQOVpWI3POshDG_bgSfbWkxW8HrlKAyRXHwe1yFfT5P4HJXLRsGJX9kUiw8ENalSEx4mRMbijepMddEChofedAYKPMgq03gUz2MRRZ5CshNk--FOUgNO_d8c3Rn6GOOdJ6_fV-M-346yMXUZHgjAFiP6QHXJsMI6HfkZktwjTUjUNE5mBxefz7B-0C8jpOrMsZYlwqYDTaP7oY1UBSUUvnmHU7LjfXXgDhPcYNHjjzTgQMdhVpsYWkDydMe9NkA5XCTVskdcqziZoqaBQKFK-qY6lWIsXahyKLOa_Y3AG-XCeBRFV_qvBcXtAc8cfF1FgCysln7ppgUmaOoRDppoPRmKhWmXwi7qvF35NeHapdTUvujDIG3-_-ROIbvtvV6uYUQBKbhXkxWtovsRPtXrWIWvbP63-orVU31p42XtVbB-zGp7Ovv-_qk6gmKihbJYkDNXEL-0-onzFC4qLDIsSaQrKILcazfN8MQsbp0vPs80PPt4PliBbMhcHr3Gq30GC203rEmTplq3l-LBdQg9tPQrIo3wJy8XvhOocsg04YgLvGsPMQ0v6kntRjVe6WXqQgort6HPTmyhl3YYE-_ZGlX2Fn1aDOsJrogPRRGuwThbt2mN5eGdXv3qn4g8bs1sETNu-2OmgCS_BFsHcxYqN7YzsgL24RF1RBNt3YN7teR9YphuHt8E2XolbToJGDNUKOAmj3-9OtVHVTaD1z4pJmnbehUKF8Ww2FdCCYn14PdcFJwzi8aiKIHVcGs9kcmPGItbEMDnaEWlqin9SBG1RAwPuCeKRTSYumWRcHOpI0Iwj_6AvvCmboLvXuhkhq_11-g2lEszotd_HVZGhQ2oU3xwNGW1jsL3c3uoaGnmMjkLqliEYbuKvzzAhrMQmMO8_5oIUs9d_diSxZAdNcgvHojTmKBLIyFB4Ukaf8EK4OsUcxkZonSPj6ssmPlHcEoXxc3iGdducT1r8Z64FFsp2vmxfnvKFZu2fBZd3DuBA-J-sY4XkbKOeTO2rL36JLGtP81XK8X6qNl42xX6qNb7orvfe3-dud9LchNtAwH63HoXOi_ctTTesDQVZhend58STYBIgcFHwejnx0o_p1fddQACqmQ_vmxDxgAKCIs93Rqa1TyTmxKVmhSPlvR3_0R67qUJ-QjLCflZxIDFhlypbtwnHpp6na7MSvlCABMh_0ZNR2LJAaANU8fH0ngWhfgNUfhoMQyR67FjgYSa_2zp-WD_GYQ0fvH2VIXIWBHUI_2tgcH-ulCKT_We1vr6b-t3Gu6dHsNEo8MmxeXCrTZgwMNxtPU74SkRxAzXMNVhDP20ScPowu9dyhSFJH7VSBoXRjmFKpDUr4_3RviB8Ez5SsKFyIAoY4YDqgujtLEQOqYiOZafcdjIBgWEuXqAKSyPzE4BvejTMq7XBhaTBpwvOuXjEGFB5PORJeRWKpk1wDfcyLIKfB4bW7gCt3_ZY7tPWyj7DxIJe2RvC-2ToR8MVXdv2Lu5uQEXqc3bV2Iyu5Qm9yMBxvgdHxWCUmuo5FHP3rRsCfdtHY-sUgSZfACpqC8rlKkEyvvWUGahk7g9zXBZ2UMhKdzkYvV57wXouhS8Jvf3_OI44jWZya6K_NH_YRCD3KGF2HA0-vNu7pixbe7XGperlwFY9yq9kqRXs-ImeQcbd1qj1Re3kOngmPimfHETBp3uHeOF-BhBk1uypPtfKnEmYxFG4zrTsi7bmEjj29ZNFLi-tClt_mlQeLFcYEo-30KCP_a_KXDNyVf4Al3jaGdpSLxYRE3RvkVTXOsS8yUsce7jMlrXJW4HxX84bumMCdNIj5TYOUmVsaz9owTv6ATUFxDzw_voq34JJ7ONhgzacAtkKKafjATBZhmiA8eO4QW299XzttlIfPwUiKjYMRgJZjzvybwFeiJwwDlurv_UWcy_sjIV7_4EOO-XN7PzOV-jfPzONBe_kI9jR2lGVv8WAAf9VcwF4g6IpRmDThWe004QEvz01Ig6bq4fLW2OHb9OlXRl_dWinilxVqN2-0ykk95n6orKkUy3YkvirOJHYS4Vx_ZsAGBjCkt69Ta3Lgv2YH4RNWc-a4uCkLAKs2iKZ3FJMfXYAiOPxpB1SEBFpBWt0hj75I54p3MNyetoYt3yfi4zbWqt_SOr-p67CQJjzk1saq7wbuChN2A7pnczWlSdUilP62R1gP0N0PbZO0dfJZ5fjREDJVELmhrMRS2SKm5HrtVRIJEzWV6S1mTX1Qn4WlDz3YL6Y8IMIhfGYbe3Kh3JBKr_WwQK9jFM2-V1QWoTm9aqfpo-n7pjpPeFf2_-sDrAZ15IqdF8C5VcWO5pUOFv42HYkGruL9F8rFcTcL-WOYUn0_OGclutBilBWvWYB-kq4353YzYYw5e2ZL4XnGVhUoKHFlhCVsLOksvrf1mMX_uQ6XuinI_mGx3pV7526WxW_LvNG4Vmiw3KzPsbiRrLFarwKjl7_LTXLQ_gVJn-mBPS-oFUQT9sLc21pp7kU59XcqqbMq4ZB5g9cFOrRYbBfJIgyzqRqkAUIUwkShLwNno62glfNkEUsSu-z-6s7nWJn1wbgh1YfZdGK_tXHGnm2Iz9cEqqW5d1vdIgWYeAYl9Jew2koFchmDLzXT_UwqQLNmB27wRc1_65LBiwyS3aqbMxS9_zZXDGDZVSLRn5IXesVrYbr1ydLpWGnrwqT8mo2Xw15SddjZV8p3PhQchtca0u9YgdJuNea5LW8F3eyjPBLbb-uX3Xn6K3ZEbEVhcIQPiTwCnV5q-2dF5Pds3f5AtnKnWzjYzrzbn6ZQHM9bEaQNmdqAUUXpeoZ0UEUe8dWxWsSGuhfsFI9Oh1nDDlTLhE3opa4b3ujgylpmQi3ajZDHdRK9KMlSaj26KDdn2ytL0gbWFJiPBmu_5PuJDWDw=\&enc=1\&auction_price={AUCTION_PRICE}\",\"https://adrta.com/i?clid=srit\&paid=srit\&avid=69\&caid=bb57db42f77807a9c5823bd8c2d9aaef\&kv1=0x0\&kv11=6e0bbc2c-6974-4719-a5fd-91b280317b82\&kv16=21.000000\&kv17=78.000000\&kv18=share.sharekaro.pro\&kv19=a70afac7-c1c8-856e-b729-afc030b83a49\&kv24=Mobile_InApp_Native\&kv26=Android\&kv4=114.124.182.15\&plid=23397093\&publisherId=31\&siteId=0\",\"https://prod-ping.swat-adping.com/ping_server/impl?viewid=EE86E0B1DD0B6E55E9D6215783ACBFB2AF94041DACDAD6AA30659D8AD2657E7EA388649211E44CE0552239710DDACE1EBD89D35F4D20D6A7EDC5E058D2B9D33D9CC2267B52D8E9C5CB74CA6DE3193BA42C19A2B9C421369D01A8B9598B541BCEA69E5217D58C098C2E93767A437474FE875FE01404CF3922A41CCFF7BE8E71DC627C875341EFAE4AB12B370171583EA079ED03773C041524798DAEC1D588A558E8888CCD7542E6945B4BCFBB72F80E624D2A38911BCE1188CFB06FC10F36757AF812DF2E0F24AC9E3A47F3BE07A1EE7D3704DA5679A03BBE5C15F5EF29113E1579508043E7C0C2B59AEF14F59B07D637B89B54DCA87690CB48EB0E0CB7187F0592D6942A6145930301668866FB3E30910D0C896976919CA5440DF69B31DEBBF48EDB040DAA29AE8F8A861FF61340CB0AAB8B1CE5BA5D8CF08F1925F433703134911864AA08D1F87526E189659384CF88CE5565756F99866B1956DF2D791901F670C6D5BE96F5FE687B186E2A3126A01CE46ECE4083EBF118BF9122C84DA2AAF4741BF84E125ABB48FC2F46069B0D66E9C770A52CD3152F2B68AE4D814EB9FE6D295DA8F75B63C3109805B601F19F000CE9F58B613ED3594F6B814A78CF43AB76212E5461308D93E2EE806DB183D8A5D3EAE247E4078290C8F5305BC0229DD5EC189792B3134F2C74186C47C63D72560A3EEE74653E09E93E09772F19D6846A10CB3BDC1161C6B6BD28C54FD98AFE94D1C047152F6118D73703A4EF2B3A3E77ADEDDB5A95E758C94959B5E5FE33DA0E19DFCA64B53A097920B93B81658F0567539BE90FD3609EE710585D12500C8DDCC37483623429B6A0ED2A6202A087B1B8370DCC2ECB6DB4B45CBED5A15F33930659B61881F4E5709549267CBF5E81FBB3CE28F0FFC99D02A10CFE208D3E9C8EB820F6E18A9E7DDA6ABD7EB30B1A691322B9B64038691A4F308B4B9CDFB2336F98F565858761CB4B274C1E57E3F20C7F23DAADAD52D3EB5CA5A128F10A754D0125881D94C6733FBEF1840A9689E5CF2F860B15E31AC9DF36A8026B8C7AAB1ABC97513E50EEBC6B3FA4B8A57F2312A6DEA79DF34C46CB35D0C3C938A09B49E615E240FBE1868F0267C5B9B7DE62929D14C10C2EE68AE2B700985D71984792676093C555A27A8BD03D0DE4A22C9A39F9707B5F015F9AEDCC6A5F268A9AE9515B28D82F1C5C14A1D531BC1B15CC9FBA24A579D0B367EAE72C8A5490127DBBC078486279424DCD9F3AD82C3D201DAD92227B9459FFF59C824BF24B75B102DA6CDFD9506FCED011A7E9277DF72B97C84A3113F54701DEC64892A2C33D5B817500B510E69CF11CAB8EB6C40EB1B5D82AC3C7C21F5277D8452813BEA9D43D9059A8FD8F6947F256E4D990D3AA68BFD24865CF9681EC1C36E6A3680C99A0F78A649DDFDEBE6977FC674F73ECF422B4C01F38F4B1FBEAB2011437904E12049D055D934046F2C3C1943A559F1E595C0374785A3F4324B9752F5D8647D6A25B40D8C92B1D022CC500DECC82DBF9B365074869DAD611D78050DD2816D0F6F8DE0D4D6750627F09DECE35475A189623B3153BF04083FC5796DD1E0F15250F9612BD011E3E9EA660663D7A7B4A4AF7EDD4AD296628A458BE8C0A197A07401D02D7E19A51E03A5F0F4F75FCE223A356C6F7B31995E8FEBC534071F9F920D6A07D2FE145E4C70A606C96C3DCF716BF39A03BDB071FB39EAC45FCBCD7FF458A7A32B2ECD9679449451C74B059FE18AC7EF85A068F5BAEBE60EAC0E6849C23CB6E57E3EFF371725BEA9A198758620A3105F34078A3CAEE07FFB072DD22BBDDC2C97C66326D6035F31A3452170FDDE575B25ADB1B1B1637620DA331583A15376652CEF2B24290BFD2C358E4B69522D6E14281D11D6DF9E01C4853EBEA61A9493DDD48FE90551F74BC5DA9E3E3B165A9DBD2FBEC12D7B101C2C4DEEEE18EA35EB3436A79B84FBE1AB2560F0CBD4C10C9416710DDEAE2380C126280539399491BE5D9B8C703FA9EB54B62F927EFEA79188298A4DC2C7761F7712AA4D40955F866F941F9CB625FC3622BBC123CAFD6D0AD87763937A06CAB409DF8CDD60FDF0B552A4D9660A968E627802593579927CF1AAE6EC13E4AB93A891D412395A15E9871E7BD4DED3F022AA703982D26B3EF7EC5AA1087833BBD066EFCA82F2ACD801BD166F39C3ED5EAC1CB3A9FE44DE452A6EAE6EAA33A978C6DB8DCD1A02A22B3E226CCEB1EE3C36AE7C34AE014BF4223571D692F27F5477A357C5120FD2A74FFF75500B265991A2C8612547A5B5538678D82439B70A24F30C7EE546F72DB58B77C9B59D68ECC8507CDFC4743D966\&enc=1\&auction_price=0.000000\&sid=__SID__\"]}}",
									"adid":"1825467",
									"adomain":[
										"http://midas.hellay.net/"
									],
									"cid":"bb57db42f77807a9c5823bd8c2d9aaef",
									"crid":"23397093"
								}
							]
						}
					],
					"cur":"USD"
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_ecom_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
					"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
					"bidid":"de208481-248d-4bcc-8c98-7d14e43d5954",
					"seatbid":[
						{
							"bid":[
								{
									"id":"de208481-248d-4bcc-8c98-7d14e43d5954",
									"impid":"1",
									"price":1,
									"nurl":"https://test-track.lacunads.com/win?viewid=5th4nW7N_Vceu_iStyezex0BXgDdrNDI4nfhOYNRDQFo0N6qmsIhumL8HOP0CgcwV9MXiIzbyt1Djimz5Ql7CGSOHC-fg5jzOe8IHBXHVGa6iY_WVfCpqinOVdoeXs2GEMKW5yNlaME2deWsQfdGrrWvDOQEZrplkoQ3x0B9wUOY7ylJ0AVBhs1f3NyG0-PN2dOtERr8mzDqxzUMQ0ms9VgZeYPmzTLFD4xDZvCGBbsI-aMo_jgzCAHgeDEgo_PoGP3NizXz35m4azjrXeC9lyBCbQyn0CFMxFxJm0Oulsg0XjD5s3KFmYPxPZbMzWfQfCeGtXfH91uKokpz4W3cPq__SP_6UwH5q1duvLbs4XgxcKlk0McOCeGZpI2bRLziriA25Yk2GLJlzQOvjk9-mKwxgw8exaN8B9_jBpj9yaQXEbiluTaQvz_p3J_Sz2rhVNcc5s3G3SGyt6w8HV4VhSh1y5LynhkRyzid8Xxbipy2vcenxBq33QuTqdiSUD5FbDH6cg38Zp3hqFLHKSGsK5fhw38L56xhdm18qrwkZbWyUKSuwE3I8a3GjG5TnYrPXIZHzAt7tlCPCNcoad3Z2irExkJNH2KxrfIULRxdvO3Y9N9qbNWP2YkIUNaJJJ3lKXjbDooMkZBq0aht5GKvMul3WvKyujvuP6EIHLEj-Dk6K8LbynXYNBr8NSlbD_a-ZoveGoJpnC_3J8IV1UEjwpmtVxetT4jO1HlmP6_HR2uQf8cyS2nAU0OpmE4_YR2aqwxh6wmwSy1SyT4nO6c26htnW23tlkXNAn359ClKjcfRmN07hKOBuS1V1KDfAmKbdZcL3nKmq7G9Q9vylTFqmIiajZUFCpmJbFRE2kbDpYD4_RLWReNee5HNdsO52gTs9ikuMPNdVFH-UAKnA-gdgNm92axyLYFKYXNI-LLGPhjucwtmOeTQZ0b_ZUpIgub6sf2BhDINBVMvrYkRBE_4Rggyce_o19T5IKutun44XkmS-vrZWHr1etTlIfBdc__D67rAW-UG59AMnTRcWYjoLvT_fkWF7YMtZhzdRHzud2X3zeFOAibH62jLuzeI0cg4zdGvq7sWlqQv1E7b2IyDwhQUJNdqZgLew_Ksp4YMdyl2gCVJS8LyWqfQYEYRTHYhqYKRmVzcDKnsy4gOlJYQNZMaHI3Gwp08SzDgwAdxxYjaXbLOxl0VwTFsKBsCvzLPlsdb0GnlUg_RG-KKDxbho4nJ3vuNguKlzFOeDzftBHNJmTli6mNQZCg3szHfiAhFrtN7zG7mjXc8Sm8TizMuI7BBtLBHnHYDE0Oand2Mo2seT8sqliRtshprR2m6UHOIb2mjcU0ASIoAafQ4LyBiW60vc08TZHCydRjlPaRe_kOdu04mPjA1R7vmB3g7BY4oxhPBeydqH5YNpn2sYUC94m3ACrs33Qr_yI1uf-vuuPqisXpjNwZWVWEQUdlDOe4CqiHLUuZ3CQ_7rb99o4n0YRcX_BYrCqxZFxPWki9Aw_HFNHHtGn9aKNtoGZ2nQaV2LGclzOYTFwzpK9XsdgCqFwDXo4vmvIVQeEnD8li3Iaz8wkCZYIFtl5bA3laBInyCtnZkvOY1NQQOu9CQSVCP4EWFOJo9u0mGgrKXTaNTC4BSepw4hvPT79LVGy_JjWu5i3pN8VeOWiDHZBttMfvflHH4nw91S0ySWmzUkof_rO-zajKXL4QVEKBzviEL2B5zQKuvnlB_HuL7cHGBXIYNZbfHbSzd5XFaySEh3hydqbh5qTkNd2UDLEJVF8sacQ4CEeAd-14Jhf7tdTzpZuiU5xCpP72IbxXSUH7BLU8z4q9iEfDCAfhTz1Si9pxZWmcAg05ChUIzYUrwdi6r88YRo0hEYMPDCQJGO9gfczvKp10Y0ls8fxhrclE6Q4ruxy0CBaJvlE8eXWGwU4CFlsZa8HDVd9yLrPYHeGiEFR1D102LxQ6nSYae8eDGjefhjHjq1yGv0F_LxNKWUuiHbc-OaXQKFDHzT2NUehaugaTSsKTrd-8HhV6CokodnnsuS1Snc6ozsvtkmYz9jXZHw-GwpvAIOWT1FrYo8IAkZ4Y9u3ITrKtz9qwwox2b64d418umRk2o__6jDo7vwaZ63W47TrMBGoxa_dZP23QgnI8cRju1xmmKDlNofhyBsbcG2LWYhCbEZE7K67pKW-V3VXa4UImA0zL_Nf9DV2i0iPIZHhgeLg6voBY2YTjJR-GEK9642LfIjzASpAiikmUKzQPHgLokA2M0yJQOVpWI3POshDG_bgSfbWkxW8HrlKAyRXHwe1yFfT5P4HJXLRsGJX9kUiw8ENalSEx4mRMbijepMddEChofedAYKPMgq03gUz2MRRZ5CshNk--FOUgNO_d8c3Rn6GOOdJ6_fV-M-346yMXUZHgjAFiP6QHXJsMI6HfkZktwjTUjUNE5mBxefz7B-0C8jpOrMsZYlwqYDTaP7oY1UBSUUvnmHU7LjfXXgDhPcYNHjjzTgQMdhVpsYWkDydMe9NkA5XCTVskdcqziZoqaBQKFK-qY6lWIsXahyKLOa_Y3AG-XCeBRFV_qvBcXtAc8cfF1FgCysln7ppgUmaOoRDppoPRmKhWmXwi7qvF35NeHapdTUvujDIG3-_-ROIbvtvV6uYUQBKbhXkxWtovsRPtXrWIWvbP63-orVU31p42XtVbB-zGp7Ovv-_qk6gmKihbJYkDNXEL-0-onzFC4qLDIsSaQrKILcazfN8MQsbp0vPs80PPt4PliBbMhcHr3Gq30GC203rEmTplq3l-LBdQg9tPQrIo3wJy8XvhOocsg04YgLvGsPMQ0v6kntRjVe6WXqQgort6HPTmyhl3YYE-_ZGlX2Fn1aDOsJrogPRRGuwThbt2mN5eGdXv3qn4g8bs1sETNu-2OmgCS_BFsHcxYqN7YzsgL24RF1RBNt3YN7teR9YphuHt8E2XolbToJGDNUKOAmj3-9OtVHVTaD1z4pJmnbehUKF8Ww2FdCCYn14PdcFJwzi8aiKIHVcGs9kcmPGItbEMDnaEWlqin9SBG1RAwPuCeKRTSYumWRcHOpI0Iwj_6AvvCmboLvXuhkhq_11-g2lEszotd_HVZGhQ2oU3xwNGW1jsL3c3uoaGnmMjkLqliEYbuKvzzAhrMQmMO8_5oIUs9d_diSxZAdNcgvHojTmKBLIyFB4Ukaf8EK4OsUcxkZonSPj6ssmPlHcEoXxc3iGdducT1r8Z64FFsp2vmxfnvKFZu2fBZd3DuBA-J-sY4XkbKOeTO2rL36JLGtP81XK8X6qNl42xX6qNb7orvfe3-dud9LchNtAwH63HoXOi_ctTTesDQVZhend58STYBIgcFHwejnx0o_p1fddQACqmQ_vmxDxgAKCIs93Rqa1TyTmxKVmhSPlvR3_0R67qUJ-QjLCflZxIDFhlypbtwnHpp6na7MSvlCABMh_0ZNR2LJAaANU8fH0ngWhfgNUfhoMQyR67FjgYSa_2zp-WD_GYQ0fvH2VIXIWBHUI_2tgcH-ulCKT_We1vr6b-t3Gu6dHsNEo8MmxeXCrTZgwMNxtPU74SkRxAzXMNVhDP20ScPowu9dyhSFJH7VSBoXRjmFKpDUr4_3RviB8Ez5SsKFyIAoY4YDqgujtLEQOqYiOZafcdjIBgWEuXqAKSyPzE4BvejTMq7XBhaTBpwvOuXjEGFB5PORJeRWKpk1wDfcyLIKfB4bW7gCt3_ZY7tPWyj7DxIJe2RvC-2ToR8MVXdv2Lu5uQEXqc3bV2Iyu5Qm9yMBxvgdHxWCUmuo5FHP3rRsCfdtHY-sUgSZfACpqC8rlKkEyvvWUGahk7g9zXBZ2UMhKdzkYvV57wXouhS8Jvf3_OI44jWZya6K_NH_YRCD3KGF2HA0-vNu7pixbe7XGperlwFY9yq9kqRXs-ImeQcbd1qj1Re3kOngmPimfHETBp3uHeOF-BhBk1uypPtfKnEmYxFG4zrTsi7bmEjj29ZNFLi-tClt_mlQeLFcYEo-30KCP_a_KXDNyVf4Al3jaGdpSLxYRE3RvkVTXOsS8yUsce7jMlrXJW4HxX84bumMCdNIj5TYOUmVsaz9owTv6ATUFxDzw_voq34JJ7ONhgzacAtkKKafjATBZhmiA8eO4QW299XzttlIfPwUiKjYMRgJZjzvybwFeiJwwDlurv_UWcy_sjIV7_4EOO-XN7PzOV-jfPzONBe_kI9jR2lGVv8WAAf9VcwF4g6IpRmDThWe004QEvz01Ig6bq4fLW2OHb9OlXRl_dWinilxVqN2-0ykk95n6orKkUy3YkvirOJHYS4Vx_ZsAGBjCkt69Ta3Lgv2YH4RNWc-a4uCkLAKs2iKZ3FJMfXYAiOPxpB1SEBFpBWt0hj75I54p3MNyetoYt3yfi4zbWqt_SOr-p67CQJjzk1saq7wbuChN2A7pnczWlSdUilP62R1gP0N0PbZO0dfJZ5fjREDJVELmhrMRS2SKm5HrtVRIJEzWV6S1mTX1Qn4WlDz3YL6Y8IMIhfGYbe3Kh3JBKr_WwQK9jFM2-V1QWoTm9aqfpo-n7pjpPeFf2_-sDrAZ15IqdF8C5VcWO5pUOFv42HYkGruL9F8rFcTcL-WOYUn0_OGclutBilBWvWYB-kq4353YzYYw5e2ZL4XnGVhUoKHFlhCVsLOksvrf1mMX_uQ6XuinI_mGx3pV7526WxW_LvNG4Vmiw3KzPsbiRrLFarwKjl7_LTXLQ_gVJn-mBPS-oFUQT9sLc21pp7kU59XcqqbMq4ZB5g9cFOrRYbBfJIgyzqRqkAUIUwkShLwNno62glfNkEUsSu-z-6s7nWJn1wbgh1YfZdGK_tXHGnm2Iz9cEqqW5d1vdIgWYeAYl9Jew2koFchmDLzXT_UwqQLNmB27wRc1_65LBiwyS3aqbMxS9_zZXDGDZVSLRn5IXesVrYbr1ydLpWGnrwqT8mo2Xw15SddjZV8p3PhQchtca0u9YgdJuNea5LW8F3eyjPBLbb-uX3Xn6K3ZEbEVhcIQPiTwCnV5q-2dF5Pds3f5AtnKnWzjYzrzbn6ZQHM9bEaQNmdqAUUXpeoZ0UEUe8dWxWsSGuhfsFI9Oh1nDDlTLhE3opa4b3ujgylpmQi3ajZDHdRK9KMlSaj26KDdn2ytL0gbWFJiPBmu_5PuJDWDw=&enc=1&auction_price=${AUCTION_PRICE}",
									"adm":"{\"native\":{\"ver\":\"1.0\",\"assets\":[{\"id\":1,\"img\":{\"type\":3,\"url\":\"https://static.rqmob.com/sa/20221118/e29f275fd2f535f17a6d779c87bb082c__FILE_CM____FILE_CM480_720__.gif\",\"w\":660,\"h\":346}}],\"link\":{\"url\":\"http://midas.hellay.net/\",\"clicktrackers\":[\"https://test-track.lacunads.com/clk?viewid=5th4nW7N_Vceu_iStyezex0BXgDdrNDI4nfhOYNRDQFo0N6qmsIhumL8HOP0CgcwV9MXiIzbyt1Djimz5Ql7CGSOHC-fg5jzOe8IHBXHVGa6iY_WVfCpqinOVdoeXs2GEMKW5yNlaME2deWsQfdGrrWvDOQEZrplkoQ3x0B9wUOY7ylJ0AVBhs1f3NyG0-PN2dOtERr8mzDqxzUMQ0ms9VgZeYPmzTLFD4xDZvCGBbsI-aMo_jgzCAHgeDEgo_PoGP3NizXz35m4azjrXeC9lyBCbQyn0CFMxFxJm0Oulsg0XjD5s3KFmYPxPZbMzWfQfCeGtXfH91uKokpz4W3cPq__SP_6UwH5q1duvLbs4XgxcKlk0McOCeGZpI2bRLziriA25Yk2GLJlzQOvjk9-mKwxgw8exaN8B9_jBpj9yaQXEbiluTaQvz_p3J_Sz2rhVNcc5s3G3SGyt6w8HV4VhSh1y5LynhkRyzid8Xxbipy2vcenxBq33QuTqdiSUD5FbDH6cg38Zp3hqFLHKSGsK5fhw38L56xhdm18qrwkZbWyUKSuwE3I8a3GjG5TnYrPXIZHzAt7tlCPCNcoad3Z2irExkJNH2KxrfIULRxdvO3Y9N9qbNWP2YkIUNaJJJ3lKXjbDooMkZBq0aht5GKvMul3WvKyujvuP6EIHLEj-Dk6K8LbynXYNBr8NSlbD_a-ZoveGoJpnC_3J8IV1UEjwpmtVxetT4jO1HlmP6_HR2uQf8cyS2nAU0OpmE4_YR2aqwxh6wmwSy1SyT4nO6c26htnW23tlkXNAn359ClKjcfRmN07hKOBuS1V1KDfAmKbdZcL3nKmq7G9Q9vylTFqmIiajZUFCpmJbFRE2kbDpYD4_RLWReNee5HNdsO52gTs9ikuMPNdVFH-UAKnA-gdgNm92axyLYFKYXNI-LLGPhjucwtmOeTQZ0b_ZUpIgub6sf2BhDINBVMvrYkRBE_4Rggyce_o19T5IKutun44XkmS-vrZWHr1etTlIfBdc__D67rAW-UG59AMnTRcWYjoLvT_fkWF7YMtZhzdRHzud2X3zeFOAibH62jLuzeI0cg4zdGvq7sWlqQv1E7b2IyDwhQUJNdqZgLew_Ksp4YMdyl2gCVJS8LyWqfQYEYRTHYhqYKRmVzcDKnsy4gOlJYQNZMaHI3Gwp08SzDgwAdxxYjaXbLOxl0VwTFsKBsCvzLPlsdb0GnlUg_RG-KKDxbho4nJ3vuNguKlzFOeDzftBHNJmTli6mNQZCg3szHfiAhFrtN7zG7mjXc8Sm8TizMuI7BBtLBHnHYDE0Oand2Mo2seT8sqliRtshprR2m6UHOIb2mjcU0ASIoAafQ4LyBiW60vc08TZHCydRjlPaRe_kOdu04mPjA1R7vmB3g7BY4oxhPBeydqH5YNpn2sYUC94m3ACrs33Qr_yI1uf-vuuPqisXpjNwZWVWEQUdlDOe4CqiHLUuZ3CQ_7rb99o4n0YRcX_BYrCqxZFxPWki9Aw_HFNHHtGn9aKNtoGZ2nQaV2LGclzOYTFwzpK9XsdgCqFwDXo4vmvIVQeEnD8li3Iaz8wkCZYIFtl5bA3laBInyCtnZkvOY1NQQOu9CQSVCP4EWFOJo9u0mGgrKXTaNTC4BSepw4hvPT79LVGy_JjWu5i3pN8VeOWiDHZBttMfvflHH4nw91S0ySWmzUkof_rO-zajKXL4QVEKBzviEL2B5zQKuvnlB_HuL7cHGBXIYNZbfHbSzd5XFaySEh3hydqbh5qTkNd2UDLEJVF8sacQ4CEeAd-14Jhf7tdTzpZuiU5xCpP72IbxXSUH7BLU8z4q9iEfDCAfhTz1Si9pxZWmcAg05ChUIzYUrwdi6r88YRo0hEYMPDCQJGO9gfczvKp10Y0ls8fxhrclE6Q4ruxy0CBaJvlE8eXWGwU4CFlsZa8HDVd9yLrPYHeGiEFR1D102LxQ6nSYae8eDGjefhjHjq1yGv0F_LxNKWUuiHbc-OaXQKFDHzT2NUehaugaTSsKTrd-8HhV6CokodnnsuS1Snc6ozsvtkmYz9jXZHw-GwpvAIOWT1FrYo8IAkZ4Y9u3ITrKtz9qwwox2b64d418umRk2o__6jDo7vwaZ63W47TrMBGoxa_dZP23QgnI8cRju1xmmKDlNofhyBsbcG2LWYhCbEZE7K67pKW-V3VXa4UImA0zL_Nf9DV2i0iPIZHhgeLg6voBY2YTjJR-GEK9642LfIjzASpAiikmUKzQPHgLokA2M0yJQOVpWI3POshDG_bgSfbWkxW8HrlKAyRXHwe1yFfT5P4HJXLRsGJX9kUiw8ENalSEx4mRMbijepMddEChofedAYKPMgq03gUz2MRRZ5CshNk--FOUgNO_d8c3Rn6GOOdJ6_fV-M-346yMXUZHgjAFiP6QHXJsMI6HfkZktwjTUjUNE5mBxefz7B-0C8jpOrMsZYlwqYDTaP7oY1UBSUUvnmHU7LjfXXgDhPcYNHjjzTgQMdhVpsYWkDydMe9NkA5XCTVskdcqziZoqaBQKFK-qY6lWIsXahyKLOa_Y3AG-XCeBRFV_qvBcXtAc8cfF1FgCysln7ppgUmaOoRDppoPRmKhWmXwi7qvF35NeHapdTUvujDIG3-_-ROIbvtvV6uYUQBKbhXkxWtovsRPtXrWIWvbP63-orVU31p42XtVbB-zGp7Ovv-_qk6gmKihbJYkDNXEL-0-onzFC4qLDIsSaQrKILcazfN8MQsbp0vPs80PPt4PliBbMhcHr3Gq30GC203rEmTplq3l-LBdQg9tPQrIo3wJy8XvhOocsg04YgLvGsPMQ0v6kntRjVe6WXqQgort6HPTmyhl3YYE-_ZGlX2Fn1aDOsJrogPRRGuwThbt2mN5eGdXv3qn4g8bs1sETNu-2OmgCS_BFsHcxYqN7YzsgL24RF1RBNt3YN7teR9YphuHt8E2XolbToJGDNUKOAmj3-9OtVHVTaD1z4pJmnbehUKF8Ww2FdCCYn14PdcFJwzi8aiKIHVcGs9kcmPGItbEMDnaEWlqin9SBG1RAwPuCeKRTSYumWRcHOpI0Iwj_6AvvCmboLvXuhkhq_11-g2lEszotd_HVZGhQ2oU3xwNGW1jsL3c3uoaGnmMjkLqliEYbuKvzzAhrMQmMO8_5oIUs9d_diSxZAdNcgvHojTmKBLIyFB4Ukaf8EK4OsUcxkZonSPj6ssmPlHcEoXxc3iGdducT1r8Z64FFsp2vmxfnvKFZu2fBZd3DuBA-J-sY4XkbKOeTO2rL36JLGtP81XK8X6qNl42xX6qNb7orvfe3-dud9LchNtAwH63HoXOi_ctTTesDQVZhend58STYBIgcFHwejnx0o_p1fddQACqmQ_vmxDxgAKCIs93Rqa1TyTmxKVmhSPlvR3_0R67qUJ-QjLCflZxIDFhlypbtwnHpp6na7MSvlCABMh_0ZNR2LJAaANU8fH0ngWhfgNUfhoMQyR67FjgYSa_2zp-WD_GYQ0fvH2VIXIWBHUI_2tgcH-ulCKT_We1vr6b-t3Gu6dHsNEo8MmxeXCrTZgwMNxtPU74SkRxAzXMNVhDP20ScPowu9dyhSFJH7VSBoXRjmFKpDUr4_3RviB8Ez5SsKFyIAoY4YDqgujtLEQOqYiOZafcdjIBgWEuXqAKSyPzE4BvejTMq7XBhaTBpwvOuXjEGFB5PORJeRWKpk1wDfcyLIKfB4bW7gCt3_ZY7tPWyj7DxIJe2RvC-2ToR8MVXdv2Lu5uQEXqc3bV2Iyu5Qm9yMBxvgdHxWCUmuo5FHP3rRsCfdtHY-sUgSZfACpqC8rlKkEyvvWUGahk7g9zXBZ2UMhKdzkYvV57wXouhS8Jvf3_OI44jWZya6K_NH_YRCD3KGF2HA0-vNu7pixbe7XGperlwFY9yq9kqRXs-ImeQcbd1qj1Re3kOngmPimfHETBp3uHeOF-BhBk1uypPtfKnEmYxFG4zrTsi7bmEjj29ZNFLi-tClt_mlQeLFcYEo-30KCP_a_KXDNyVf4Al3jaGdpSLxYRE3RvkVTXOsS8yUsce7jMlrXJW4HxX84bumMCdNIj5TYOUmVsaz9owTv6ATUFxDzw_voq34JJ7ONhgzacAtkKKafjATBZhmiA8eO4QW299XzttlIfPwUiKjYMRgJZjzvybwFeiJwwDlurv_UWcy_sjIV7_4EOO-XN7PzOV-jfPzONBe_kI9jR2lGVv8WAAf9VcwF4g6IpRmDThWe004QEvz01Ig6bq4fLW2OHb9OlXRl_dWinilxVqN2-0ykk95n6orKkUy3YkvirOJHYS4Vx_ZsAGBjCkt69Ta3Lgv2YH4RNWc-a4uCkLAKs2iKZ3FJMfXYAiOPxpB1SEBFpBWt0hj75I54p3MNyetoYt3yfi4zbWqt_SOr-p67CQJjzk1saq7wbuChN2A7pnczWlSdUilP62R1gP0N0PbZO0dfJZ5fjREDJVELmhrMRS2SKm5HrtVRIJEzWV6S1mTX1Qn4WlDz3YL6Y8IMIhfGYbe3Kh3JBKr_WwQK9jFM2-V1QWoTm9aqfpo-n7pjpPeFf2_-sDrAZ15IqdF8C5VcWO5pUOFv42HYkGruL9F8rFcTcL-WOYUn0_OGclutBilBWvWYB-kq4353YzYYw5e2ZL4XnGVhUoKHFlhCVsLOksvrf1mMX_uQ6XuinI_mGx3pV7526WxW_LvNG4Vmiw3KzPsbiRrLFarwKjl7_LTXLQ_gVJn-mBPS-oFUQT9sLc21pp7kU59XcqqbMq4ZB5g9cFOrRYbBfJIgyzqRqkAUIUwkShLwNno62glfNkEUsSu-z-6s7nWJn1wbgh1YfZdGK_tXHGnm2Iz9cEqqW5d1vdIgWYeAYl9Jew2koFchmDLzXT_UwqQLNmB27wRc1_65LBiwyS3aqbMxS9_zZXDGDZVSLRn5IXesVrYbr1ydLpWGnrwqT8mo2Xw15SddjZV8p3PhQchtca0u9YgdJuNea5LW8F3eyjPBLbb-uX3Xn6K3ZEbEVhcIQPiTwCnV5q-2dF5Pds3f5AtnKnWzjYzrzbn6ZQHM9bEaQNmdqAUUXpeoZ0UEUe8dWxWsSGuhfsFI9Oh1nDDlTLhE3opa4b3ujgylpmQi3ajZDHdRK9KMlSaj26KDdn2ytL0gbWFJiPBmu_5PuJDWDw=\&enc=1\",\"https://prod-ping.swat-adping.com/ping_server/clk?viewid=EE86E0B1DD0B6E55E9D6215783ACBFB2AF94041DACDAD6AA30659D8AD2657E7EA388649211E44CE0552239710DDACE1EBD89D35F4D20D6A7EDC5E058D2B9D33D9CC2267B52D8E9C5CB74CA6DE3193BA42C19A2B9C421369D01A8B9598B541BCEA69E5217D58C098C2E93767A437474FE875FE01404CF3922A41CCFF7BE8E71DC627C875341EFAE4AB12B370171583EA079ED03773C041524798DAEC1D588A558E8888CCD7542E6945B4BCFBB72F80E624D2A38911BCE1188CFB06FC10F36757AF812DF2E0F24AC9E3A47F3BE07A1EE7D3704DA5679A03BBE5C15F5EF29113E1579508043E7C0C2B59AEF14F59B07D637B89B54DCA87690CB48EB0E0CB7187F0592D6942A6145930301668866FB3E30910D0C896976919CA5440DF69B31DEBBF48EDB040DAA29AE8F8A861FF61340CB0AAB8B1CE5BA5D8CF08F1925F433703134911864AA08D1F87526E189659384CF88CE5565756F99866B1956DF2D791901F670C6D5BE96F5FE687B186E2A3126A01CE46ECE4083EBF118BF9122C84DA2AAF4741BF84E125ABB48FC2F46069B0D66E9C770A52CD3152F2B68AE4D814EB9FE6D295DA8F75B63C3109805B601F19F000CE9F58B613ED3594F6B814A78CF43AB76212E5461308D93E2EE806DB183D8A5D3EAE247E4078290C8F5305BC0229DD5EC189792B3134F2C74186C47C63D72560A3EEE74653E09E93E09772F19D6846A10CB3BDC1161C6B6BD28C54FD98AFE94D1C047152F6118D73703A4EF2B3A3E77ADEDDB5A95E758C94959B5E5FE33DA0E19DFCA64B53A097920B93B81658F0567539BE90FD3609EE710585D12500C8DDCC37483623429B6A0ED2A6202A087B1B8370DCC2ECB6DB4B45CBED5A15F33930659B61881F4E5709549267CBF5E81FBB3CE28F0FFC99D02A10CFE208D3E9C8EB820F6E18A9E7DDA6ABD7EB30B1A691322B9B64038691A4F308B4B9CDFB2336F98F565858761CB4B274C1E57E3F20C7F23DAADAD52D3EB5CA5A128F10A754D0125881D94C6733FBEF1840A9689E5CF2F860B15E31AC9DF36A8026B8C7AAB1ABC97513E50EEBC6B3FA4B8A57F2312A6DEA79DF34C46CB35D0C3C938A09B49E615E240FBE1868F0267C5B9B7DE62929D14C10C2EE68AE2B700985D71984792676093C555A27A8BD03D0DE4A22C9A39F9707B5F015F9AEDCC6A5F268A9AE9515B28D82F1C5C14A1D531BC1B15CC9FBA24A579D0B367EAE72C8A5490127DBBC078486279424DCD9F3AD82C3D201DAD92227B9459FFF59C824BF24B75B102DA6CDFD9506FCED011A7E9277DF72B97C84A3113F54701DEC64892A2C33D5B817500B510E69CF11CAB8EB6C40EB1B5D82AC3C7C21F5277D8452813BEA9D43D9059A8FD8F6947F256E4D990D3AA68BFD24865CF9681EC1C36E6A3680C99A0F78A649DDFDEBE6977FC674F73ECF422B4C01F38F4B1FBEAB2011437904E12049D055D934046F2C3C1943A559F1E595C0374785A3F4324B9752F5D8647D6A25B40D8C92B1D022CC500DECC82DBF9B365074869DAD611D78050DD2816D0F6F8DE0D4D6750627F09DECE35475A189623B3153BF04083FC5796DD1E0F15250F9612BD011E3E9EA660663D7A7B4A4AF7EDD4AD296628A458BE8C0A197A07401D02D7E19A51E03A5F0F4F75FCE223A356C6F7B31995E8FEBC534071F9F920D6A07D2FE145E4C70A606C96C3DCF716BF39A03BDB071FB39EAC45FCBCD7FF458A7A32B2ECD9679449451C74B059FE18AC7EF85A068F5BAEBE60EAC0E6849C23CB6E57E3EFF371725BEA9A198758620A3105F34078A3CAEE07FFB072DD22BBDDC2C97C66326D6035F31A3452170FDDE575B25ADB1B1B1637620DA331583A15376652CEF2B24290BFD2C358E4B69522D6E14281D11D6DF9E01C4853EBEA61A9493DDD48FE90551F74BC5DA9E3E3B165A9DBD2FBEC12D7B101C2C4DEEEE18EA35EB3436A79B84FBE1AB2560F0CBD4C10C9416710DDEAE2380C126280539399491BE5D9B8C703FA9EB54B62F927EFEA79188298A4DC2C7761F7712AA4D40955F866F941F9CB625FC3622BBC123CAFD6D0AD87763937A06CAB409DF8CDD60FDF0B552A4D9660A968E627802593579927CF1AAE6EC13E4AB93A891D412395A15E9871E7BD4DED3F022AA703982D26B3EF7EC5AA1087833BBD066EFCA82F2ACD801BD166F39C3ED5EAC1CB3A9FE44DE452A6EAE6EAA33A978C6DB8DCD1A02A22B3E226CCEB1EE3C36AE7C34AE014BF4223571D692F27F5477A357C5120FD2A74FFF75500B265991A2C8612547A5B550B9379E4E07580F6C7526114B8275043F64C79E3F46FBBD39078180A0B31588E\&enc=1\&auction_price=0.000000\&sid=__SID__\"]},\"imptrackers\":[\"https://test-track.lacunads.com/imp?viewid=5th4nW7N_Vceu_iStyezex0BXgDdrNDI4nfhOYNRDQFo0N6qmsIhumL8HOP0CgcwV9MXiIzbyt1Djimz5Ql7CGSOHC-fg5jzOe8IHBXHVGa6iY_WVfCpqinOVdoeXs2GEMKW5yNlaME2deWsQfdGrrWvDOQEZrplkoQ3x0B9wUOY7ylJ0AVBhs1f3NyG0-PN2dOtERr8mzDqxzUMQ0ms9VgZeYPmzTLFD4xDZvCGBbsI-aMo_jgzCAHgeDEgo_PoGP3NizXz35m4azjrXeC9lyBCbQyn0CFMxFxJm0Oulsg0XjD5s3KFmYPxPZbMzWfQfCeGtXfH91uKokpz4W3cPq__SP_6UwH5q1duvLbs4XgxcKlk0McOCeGZpI2bRLziriA25Yk2GLJlzQOvjk9-mKwxgw8exaN8B9_jBpj9yaQXEbiluTaQvz_p3J_Sz2rhVNcc5s3G3SGyt6w8HV4VhSh1y5LynhkRyzid8Xxbipy2vcenxBq33QuTqdiSUD5FbDH6cg38Zp3hqFLHKSGsK5fhw38L56xhdm18qrwkZbWyUKSuwE3I8a3GjG5TnYrPXIZHzAt7tlCPCNcoad3Z2irExkJNH2KxrfIULRxdvO3Y9N9qbNWP2YkIUNaJJJ3lKXjbDooMkZBq0aht5GKvMul3WvKyujvuP6EIHLEj-Dk6K8LbynXYNBr8NSlbD_a-ZoveGoJpnC_3J8IV1UEjwpmtVxetT4jO1HlmP6_HR2uQf8cyS2nAU0OpmE4_YR2aqwxh6wmwSy1SyT4nO6c26htnW23tlkXNAn359ClKjcfRmN07hKOBuS1V1KDfAmKbdZcL3nKmq7G9Q9vylTFqmIiajZUFCpmJbFRE2kbDpYD4_RLWReNee5HNdsO52gTs9ikuMPNdVFH-UAKnA-gdgNm92axyLYFKYXNI-LLGPhjucwtmOeTQZ0b_ZUpIgub6sf2BhDINBVMvrYkRBE_4Rggyce_o19T5IKutun44XkmS-vrZWHr1etTlIfBdc__D67rAW-UG59AMnTRcWYjoLvT_fkWF7YMtZhzdRHzud2X3zeFOAibH62jLuzeI0cg4zdGvq7sWlqQv1E7b2IyDwhQUJNdqZgLew_Ksp4YMdyl2gCVJS8LyWqfQYEYRTHYhqYKRmVzcDKnsy4gOlJYQNZMaHI3Gwp08SzDgwAdxxYjaXbLOxl0VwTFsKBsCvzLPlsdb0GnlUg_RG-KKDxbho4nJ3vuNguKlzFOeDzftBHNJmTli6mNQZCg3szHfiAhFrtN7zG7mjXc8Sm8TizMuI7BBtLBHnHYDE0Oand2Mo2seT8sqliRtshprR2m6UHOIb2mjcU0ASIoAafQ4LyBiW60vc08TZHCydRjlPaRe_kOdu04mPjA1R7vmB3g7BY4oxhPBeydqH5YNpn2sYUC94m3ACrs33Qr_yI1uf-vuuPqisXpjNwZWVWEQUdlDOe4CqiHLUuZ3CQ_7rb99o4n0YRcX_BYrCqxZFxPWki9Aw_HFNHHtGn9aKNtoGZ2nQaV2LGclzOYTFwzpK9XsdgCqFwDXo4vmvIVQeEnD8li3Iaz8wkCZYIFtl5bA3laBInyCtnZkvOY1NQQOu9CQSVCP4EWFOJo9u0mGgrKXTaNTC4BSepw4hvPT79LVGy_JjWu5i3pN8VeOWiDHZBttMfvflHH4nw91S0ySWmzUkof_rO-zajKXL4QVEKBzviEL2B5zQKuvnlB_HuL7cHGBXIYNZbfHbSzd5XFaySEh3hydqbh5qTkNd2UDLEJVF8sacQ4CEeAd-14Jhf7tdTzpZuiU5xCpP72IbxXSUH7BLU8z4q9iEfDCAfhTz1Si9pxZWmcAg05ChUIzYUrwdi6r88YRo0hEYMPDCQJGO9gfczvKp10Y0ls8fxhrclE6Q4ruxy0CBaJvlE8eXWGwU4CFlsZa8HDVd9yLrPYHeGiEFR1D102LxQ6nSYae8eDGjefhjHjq1yGv0F_LxNKWUuiHbc-OaXQKFDHzT2NUehaugaTSsKTrd-8HhV6CokodnnsuS1Snc6ozsvtkmYz9jXZHw-GwpvAIOWT1FrYo8IAkZ4Y9u3ITrKtz9qwwox2b64d418umRk2o__6jDo7vwaZ63W47TrMBGoxa_dZP23QgnI8cRju1xmmKDlNofhyBsbcG2LWYhCbEZE7K67pKW-V3VXa4UImA0zL_Nf9DV2i0iPIZHhgeLg6voBY2YTjJR-GEK9642LfIjzASpAiikmUKzQPHgLokA2M0yJQOVpWI3POshDG_bgSfbWkxW8HrlKAyRXHwe1yFfT5P4HJXLRsGJX9kUiw8ENalSEx4mRMbijepMddEChofedAYKPMgq03gUz2MRRZ5CshNk--FOUgNO_d8c3Rn6GOOdJ6_fV-M-346yMXUZHgjAFiP6QHXJsMI6HfkZktwjTUjUNE5mBxefz7B-0C8jpOrMsZYlwqYDTaP7oY1UBSUUvnmHU7LjfXXgDhPcYNHjjzTgQMdhVpsYWkDydMe9NkA5XCTVskdcqziZoqaBQKFK-qY6lWIsXahyKLOa_Y3AG-XCeBRFV_qvBcXtAc8cfF1FgCysln7ppgUmaOoRDppoPRmKhWmXwi7qvF35NeHapdTUvujDIG3-_-ROIbvtvV6uYUQBKbhXkxWtovsRPtXrWIWvbP63-orVU31p42XtVbB-zGp7Ovv-_qk6gmKihbJYkDNXEL-0-onzFC4qLDIsSaQrKILcazfN8MQsbp0vPs80PPt4PliBbMhcHr3Gq30GC203rEmTplq3l-LBdQg9tPQrIo3wJy8XvhOocsg04YgLvGsPMQ0v6kntRjVe6WXqQgort6HPTmyhl3YYE-_ZGlX2Fn1aDOsJrogPRRGuwThbt2mN5eGdXv3qn4g8bs1sETNu-2OmgCS_BFsHcxYqN7YzsgL24RF1RBNt3YN7teR9YphuHt8E2XolbToJGDNUKOAmj3-9OtVHVTaD1z4pJmnbehUKF8Ww2FdCCYn14PdcFJwzi8aiKIHVcGs9kcmPGItbEMDnaEWlqin9SBG1RAwPuCeKRTSYumWRcHOpI0Iwj_6AvvCmboLvXuhkhq_11-g2lEszotd_HVZGhQ2oU3xwNGW1jsL3c3uoaGnmMjkLqliEYbuKvzzAhrMQmMO8_5oIUs9d_diSxZAdNcgvHojTmKBLIyFB4Ukaf8EK4OsUcxkZonSPj6ssmPlHcEoXxc3iGdducT1r8Z64FFsp2vmxfnvKFZu2fBZd3DuBA-J-sY4XkbKOeTO2rL36JLGtP81XK8X6qNl42xX6qNb7orvfe3-dud9LchNtAwH63HoXOi_ctTTesDQVZhend58STYBIgcFHwejnx0o_p1fddQACqmQ_vmxDxgAKCIs93Rqa1TyTmxKVmhSPlvR3_0R67qUJ-QjLCflZxIDFhlypbtwnHpp6na7MSvlCABMh_0ZNR2LJAaANU8fH0ngWhfgNUfhoMQyR67FjgYSa_2zp-WD_GYQ0fvH2VIXIWBHUI_2tgcH-ulCKT_We1vr6b-t3Gu6dHsNEo8MmxeXCrTZgwMNxtPU74SkRxAzXMNVhDP20ScPowu9dyhSFJH7VSBoXRjmFKpDUr4_3RviB8Ez5SsKFyIAoY4YDqgujtLEQOqYiOZafcdjIBgWEuXqAKSyPzE4BvejTMq7XBhaTBpwvOuXjEGFB5PORJeRWKpk1wDfcyLIKfB4bW7gCt3_ZY7tPWyj7DxIJe2RvC-2ToR8MVXdv2Lu5uQEXqc3bV2Iyu5Qm9yMBxvgdHxWCUmuo5FHP3rRsCfdtHY-sUgSZfACpqC8rlKkEyvvWUGahk7g9zXBZ2UMhKdzkYvV57wXouhS8Jvf3_OI44jWZya6K_NH_YRCD3KGF2HA0-vNu7pixbe7XGperlwFY9yq9kqRXs-ImeQcbd1qj1Re3kOngmPimfHETBp3uHeOF-BhBk1uypPtfKnEmYxFG4zrTsi7bmEjj29ZNFLi-tClt_mlQeLFcYEo-30KCP_a_KXDNyVf4Al3jaGdpSLxYRE3RvkVTXOsS8yUsce7jMlrXJW4HxX84bumMCdNIj5TYOUmVsaz9owTv6ATUFxDzw_voq34JJ7ONhgzacAtkKKafjATBZhmiA8eO4QW299XzttlIfPwUiKjYMRgJZjzvybwFeiJwwDlurv_UWcy_sjIV7_4EOO-XN7PzOV-jfPzONBe_kI9jR2lGVv8WAAf9VcwF4g6IpRmDThWe004QEvz01Ig6bq4fLW2OHb9OlXRl_dWinilxVqN2-0ykk95n6orKkUy3YkvirOJHYS4Vx_ZsAGBjCkt69Ta3Lgv2YH4RNWc-a4uCkLAKs2iKZ3FJMfXYAiOPxpB1SEBFpBWt0hj75I54p3MNyetoYt3yfi4zbWqt_SOr-p67CQJjzk1saq7wbuChN2A7pnczWlSdUilP62R1gP0N0PbZO0dfJZ5fjREDJVELmhrMRS2SKm5HrtVRIJEzWV6S1mTX1Qn4WlDz3YL6Y8IMIhfGYbe3Kh3JBKr_WwQK9jFM2-V1QWoTm9aqfpo-n7pjpPeFf2_-sDrAZ15IqdF8C5VcWO5pUOFv42HYkGruL9F8rFcTcL-WOYUn0_OGclutBilBWvWYB-kq4353YzYYw5e2ZL4XnGVhUoKHFlhCVsLOksvrf1mMX_uQ6XuinI_mGx3pV7526WxW_LvNG4Vmiw3KzPsbiRrLFarwKjl7_LTXLQ_gVJn-mBPS-oFUQT9sLc21pp7kU59XcqqbMq4ZB5g9cFOrRYbBfJIgyzqRqkAUIUwkShLwNno62glfNkEUsSu-z-6s7nWJn1wbgh1YfZdGK_tXHGnm2Iz9cEqqW5d1vdIgWYeAYl9Jew2koFchmDLzXT_UwqQLNmB27wRc1_65LBiwyS3aqbMxS9_zZXDGDZVSLRn5IXesVrYbr1ydLpWGnrwqT8mo2Xw15SddjZV8p3PhQchtca0u9YgdJuNea5LW8F3eyjPBLbb-uX3Xn6K3ZEbEVhcIQPiTwCnV5q-2dF5Pds3f5AtnKnWzjYzrzbn6ZQHM9bEaQNmdqAUUXpeoZ0UEUe8dWxWsSGuhfsFI9Oh1nDDlTLhE3opa4b3ujgylpmQi3ajZDHdRK9KMlSaj26KDdn2ytL0gbWFJiPBmu_5PuJDWDw=\&enc=1\&auction_price=${AUCTION_PRICE}\",\"https://adrta.com/i?clid=srit\&paid=srit\&avid=69\&caid=bb57db42f77807a9c5823bd8c2d9aaef\&kv1=0x0\&kv11=6e0bbc2c-6974-4719-a5fd-91b280317b82\&kv16=21.000000\&kv17=78.000000\&kv18=share.sharekaro.pro\&kv19=a70afac7-c1c8-856e-b729-afc030b83a49\&kv24=Mobile_InApp_Native\&kv26=Android\&kv4=114.124.182.15\&plid=23397093\&publisherId=31\&siteId=0\",\"https://prod-ping.swat-adping.com/ping_server/impl?viewid=EE86E0B1DD0B6E55E9D6215783ACBFB2AF94041DACDAD6AA30659D8AD2657E7EA388649211E44CE0552239710DDACE1EBD89D35F4D20D6A7EDC5E058D2B9D33D9CC2267B52D8E9C5CB74CA6DE3193BA42C19A2B9C421369D01A8B9598B541BCEA69E5217D58C098C2E93767A437474FE875FE01404CF3922A41CCFF7BE8E71DC627C875341EFAE4AB12B370171583EA079ED03773C041524798DAEC1D588A558E8888CCD7542E6945B4BCFBB72F80E624D2A38911BCE1188CFB06FC10F36757AF812DF2E0F24AC9E3A47F3BE07A1EE7D3704DA5679A03BBE5C15F5EF29113E1579508043E7C0C2B59AEF14F59B07D637B89B54DCA87690CB48EB0E0CB7187F0592D6942A6145930301668866FB3E30910D0C896976919CA5440DF69B31DEBBF48EDB040DAA29AE8F8A861FF61340CB0AAB8B1CE5BA5D8CF08F1925F433703134911864AA08D1F87526E189659384CF88CE5565756F99866B1956DF2D791901F670C6D5BE96F5FE687B186E2A3126A01CE46ECE4083EBF118BF9122C84DA2AAF4741BF84E125ABB48FC2F46069B0D66E9C770A52CD3152F2B68AE4D814EB9FE6D295DA8F75B63C3109805B601F19F000CE9F58B613ED3594F6B814A78CF43AB76212E5461308D93E2EE806DB183D8A5D3EAE247E4078290C8F5305BC0229DD5EC189792B3134F2C74186C47C63D72560A3EEE74653E09E93E09772F19D6846A10CB3BDC1161C6B6BD28C54FD98AFE94D1C047152F6118D73703A4EF2B3A3E77ADEDDB5A95E758C94959B5E5FE33DA0E19DFCA64B53A097920B93B81658F0567539BE90FD3609EE710585D12500C8DDCC37483623429B6A0ED2A6202A087B1B8370DCC2ECB6DB4B45CBED5A15F33930659B61881F4E5709549267CBF5E81FBB3CE28F0FFC99D02A10CFE208D3E9C8EB820F6E18A9E7DDA6ABD7EB30B1A691322B9B64038691A4F308B4B9CDFB2336F98F565858761CB4B274C1E57E3F20C7F23DAADAD52D3EB5CA5A128F10A754D0125881D94C6733FBEF1840A9689E5CF2F860B15E31AC9DF36A8026B8C7AAB1ABC97513E50EEBC6B3FA4B8A57F2312A6DEA79DF34C46CB35D0C3C938A09B49E615E240FBE1868F0267C5B9B7DE62929D14C10C2EE68AE2B700985D71984792676093C555A27A8BD03D0DE4A22C9A39F9707B5F015F9AEDCC6A5F268A9AE9515B28D82F1C5C14A1D531BC1B15CC9FBA24A579D0B367EAE72C8A5490127DBBC078486279424DCD9F3AD82C3D201DAD92227B9459FFF59C824BF24B75B102DA6CDFD9506FCED011A7E9277DF72B97C84A3113F54701DEC64892A2C33D5B817500B510E69CF11CAB8EB6C40EB1B5D82AC3C7C21F5277D8452813BEA9D43D9059A8FD8F6947F256E4D990D3AA68BFD24865CF9681EC1C36E6A3680C99A0F78A649DDFDEBE6977FC674F73ECF422B4C01F38F4B1FBEAB2011437904E12049D055D934046F2C3C1943A559F1E595C0374785A3F4324B9752F5D8647D6A25B40D8C92B1D022CC500DECC82DBF9B365074869DAD611D78050DD2816D0F6F8DE0D4D6750627F09DECE35475A189623B3153BF04083FC5796DD1E0F15250F9612BD011E3E9EA660663D7A7B4A4AF7EDD4AD296628A458BE8C0A197A07401D02D7E19A51E03A5F0F4F75FCE223A356C6F7B31995E8FEBC534071F9F920D6A07D2FE145E4C70A606C96C3DCF716BF39A03BDB071FB39EAC45FCBCD7FF458A7A32B2ECD9679449451C74B059FE18AC7EF85A068F5BAEBE60EAC0E6849C23CB6E57E3EFF371725BEA9A198758620A3105F34078A3CAEE07FFB072DD22BBDDC2C97C66326D6035F31A3452170FDDE575B25ADB1B1B1637620DA331583A15376652CEF2B24290BFD2C358E4B69522D6E14281D11D6DF9E01C4853EBEA61A9493DDD48FE90551F74BC5DA9E3E3B165A9DBD2FBEC12D7B101C2C4DEEEE18EA35EB3436A79B84FBE1AB2560F0CBD4C10C9416710DDEAE2380C126280539399491BE5D9B8C703FA9EB54B62F927EFEA79188298A4DC2C7761F7712AA4D40955F866F941F9CB625FC3622BBC123CAFD6D0AD87763937A06CAB409DF8CDD60FDF0B552A4D9660A968E627802593579927CF1AAE6EC13E4AB93A891D412395A15E9871E7BD4DED3F022AA703982D26B3EF7EC5AA1087833BBD066EFCA82F2ACD801BD166F39C3ED5EAC1CB3A9FE44DE452A6EAE6EAA33A978C6DB8DCD1A02A22B3E226CCEB1EE3C36AE7C34AE014BF4223571D692F27F5477A357C5120FD2A74FFF75500B265991A2C8612547A5B5538678D82439B70A24F30C7EE546F72DB58B77C9B59D68ECC8507CDFC4743D966\&enc=1\&auction_price=0.000000\&sid=__SID__\"]}}",
									"adid":"1825467",
									"adomain":[
										"http://midas.hellay.net/"
									],
									"cid":"bb57db42f77807a9c5823bd8c2d9aaef",
									"crid":"23397093"
								}
							]
						}
					],
					"cur":"USD"
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//mobupps
	r.HandleFunc("/mock/868/get_mobupps_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"9a191d46-6054-4869-a547-9ef8d2426c55",
						"seatbid":[
							{
								"bid":[
									{
										"id":"0002076e700e27e2ee83",
										"impid":"1",
										"price":2.384,
										"nurl":"https://t-odx.op-mobile.opera.com/win?a=a5522269118912\u0026adomain=wendys.com\u0026b=949295212\u0026cc=US\u0026cid=40021128\u0026cm=1\u0026crid=80859_55272534\u0026ct=\u0026dbbf=1.000\u0026dbmb=0.000\u0026dvt=PHONE\u0026ep=ep7270864899840\u0026ext=Cp85eZi9P-BMmPnnTXQkoyNWkG6zxQvYxK83vbIIOtLV8b2-5i1rN_9s1h8tAP9F15_Kezbi9-11_mALy0S-z7xcitQPmr18NGda6X6gOflvDwEB6zapWeuLZXXlfYqG_nFwff_rD0LbonlzBEY3Y3FZNwYObE6MejELWJEmg2g%3D\u0026gi=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026iabCat=IAB8-9\u0026impr_dl=1669969833967\u0026it=1\u0026m=m5522269119233\u0026ot=IOS\u0026pubId=pub7226695071424\u0026rip=172.58.103.70\u0026s=s7226745164736\u0026said=f5540250981342e753b259364d1da363\u0026sc=-1\u0026schp=-1\u0026se=0002076e700e27e2ee83\u0026spid=c1be64e41556140ba4007003a6a838e5\u0026srid=9a191d46-6054-4869-a547-9ef8d2426c55\u0026ss=3\u0026u=7cc3bfc211ced95f\u0026va=a5522269118912\u0026vd=120\u0026vm=m5522269119233\u0026ac=${AUCTION_CURRENCY}\u0026ap=${AUCTION_PRICE}",
										"lurl":"https://t-odx.op-mobile.opera.com/loss?a=a5522269118912\u0026adomain=wendys.com\u0026b=949295212\u0026cc=US\u0026cid=40021128\u0026cm=1\u0026crid=80859_55272534\u0026ct=\u0026dbbf=1.000\u0026dbmb=0.000\u0026dvt=PHONE\u0026ep=ep7270864899840\u0026ext=Cp85eZi9P-BMmPnnTXQkoyNWkG6zxQvYxK83vbIIOtLV8b2-5i1rN_9s1h8tAP9F15_Kezbi9-11_mALy0S-z7xcitQPmr18NGda6X6gOflvDwEB6zapWeuLZXXlfYqG_nFwff_rD0LbonlzBEY3Y3FZNwYObE6MejELWJEmg2g%3D\u0026gi=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026iabCat=IAB8-9\u0026impr_dl=1669969833967\u0026it=1\u0026m=m5522269119233\u0026ot=IOS\u0026pubId=pub7226695071424\u0026rip=172.58.103.70\u0026s=s7226745164736\u0026said=f5540250981342e753b259364d1da363\u0026sc=-1\u0026schp=-1\u0026se=0002076e700e27e2ee83\u0026spid=c1be64e41556140ba4007003a6a838e5\u0026srid=9a191d46-6054-4869-a547-9ef8d2426c55\u0026ss=3\u0026u=7cc3bfc211ced95f\u0026va=a5522269118912\u0026vd=120\u0026vm=m5522269119233\u0026al=${AUCTION_LOSS}",
										"adm":"\u003cVAST version=\"4.0\"\u003e\u003cAd id=\"55272534\"\u003e\u003cWrapper\u003e\u003cAdSystem\u003e\u003c![CDATA[Conversant]]\u003e\u003c/AdSystem\u003e\u003cVASTAdTagURI\u003e\u003c![CDATA[https://iad-usadmm.dotomi.com/fetch/html5/player/video/vast/2_0?cg=91\u0026dtmid=391591916636903898\u0026magic=42\u0026utype=0\u0026bidServerId=6238\u0026pnid=243\u0026pid=243\u0026ms=50\u0026trid=6238068087894529189\u0026dtm_server_id=1936\u0026comId=80859\u0026msgCampId=40021128\u0026tid=55272534\u0026parentMsgId=40021128\u0026ptid=50020278\u0026icb=0\u0026cgcb=-1\u0026fpc=0\u0026dvcid=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026mwp=2.980000\u0026btcurl=TriPeaks\u0026rt=4\u0026supplyType=3\u0026ctrl_ad_id=1\u0026iblob=h-ep3fq0CKWpnf2AvITJVhCi8eeOzTAaHVRyaVBlYWtzIFNvbGl0YWlyZTogQ2FyZCBHYW1lIgYyMTcyMDYqCTk0OTI5NTIxMjAAOiQ0ODhCQjU4RC1GN0IwLTQ5NUYtOTAyOC1EN0I1RTBFQzAyNTJSBUFwcGxlWgZJUEhPTkViBDE1LjVqBW9wZXJhcgUwLjAuMXgAggENMTcyLjU4LjEwMy43MJAB5ASaAQUwLjAuMaABA6oBBjEwMDI2ObIBBElBQjmyAQZJQUI5LTeyAQdJQUI5LTMwuAEDwAHL7MyY1eTc7QTIAf___________wHQAQDgAfitN-ABoPc24AHQ4TjoAdDju6O7jPKXxQHzAQoCVVMSAlVTGP4BKAA4AEAASABQAGAAbQAAAAB1AAAAAHoRVC1NT0JJTEUgVVNBIElOQy6KAQhULU1PQklMRZIBBk1PQklMRfQB-wEKAlVTEgJVUxj-ASICVFgoLDiXAlDqBFoFNzc1MjH8AYICCDEwMjcxODgxiAL___________8BmAIBoAIAqAKC7FCwAgHAAgLKAjwyMDQ5ODEwMjExfDg5ODI5Mjg0N3wxMDkwMTE0NDY2fDIxMTk2NjcyMzN8MTM2NDUxOTg3OXwwfC0xfDDgAgDoAgbzAgjprx8QmuniufUvGgExIeF6FK5H4co_KVVVVVVVVdU_9ALzAgjDsR8Q7v3jnMEwGgExIeF6FK5H4co_KVVVVVVVVdU_9ALzAgiXxxgQ_oO77q0wGgExIQAAAAAAAAAAKVVVVVVVVdU_9AL5AgAAAAAAAADAgQNczM8NTVnmP4kDzLIngc051z-RAwAAAAAAAPC_mQMAAAAAAADwv6EDAAAAAAAA8L-pAwAAAAAAAPA_sAMA8gMDVVNE-QMAAAAAAADwP4EEAAAAAAAAMUCJBNejcD0K1wdAkQQUrkfhehT-P5kETxHn1OHA-r6gBP-7xI_KMKgE5vzGAbAElBK5BLmDTAchFMhAwQQZdv3OFm_GP-kEAAAAAAAAAADxBAAAAAAAAAAA-AQAggUDaU9TiAUBkAUPmAUZqAUAsQUAAAAAAAAAALkFAAAAAAAAAADBBQAAAAAAAPC_yQUAAAAAAAAAANAFANgFAOkFAAAAAAAAAADxBQAAAAAAAAAA-QUAAAAAAAAAAIIGAklQmAb___________8BqAYAsAYB\u0026cturl=\u0026vpaid=1\u0026apis=6\u0026vcskippable_type=2\u0026vcskippable=false\u0026vrp=8\u0026min_duration=5\u0026max_duration=120\u0026dtm_user_ip=172.58.103.70]]\u003e\u003c/VASTAdTagURI\u003e\u003cImpression\u003e\u003c![CDATA[https://event.ad.cpe.dotomi.com/cvx/event/imp?enc=eAGNVmtzssgS_i9-8FNMhptAtqxTzS2ighLBxJzashBGRLkFUNRT57-fRpP47m7V7pkPMtP99NMzPe10_6dzqGgZh53nDiczgszITL_P9WXCSbLUeejUp6tOZPi-yEoyz_QliYiEMKjLaN3k5f6Vfh5oVZstx59xnIy4qqVgGZEl_YdOGB_jKs6zFs4gf6vj2IdOmq_jhGp-7SMN-U3mZVYWWIZF-3Uczso4oEjyKEsPneKwVvOq_ln6ZZ3R0qAIII_koRMX87qMswiJ0OmjID0yhHsUCVLF4abl5yVJUQRJ6xmiQnq8LBg9mbBSTxMVQSe6SlihdVwdiiI5u-cCmTk0zmoalX6Nu_-RpTSMr5LZbRcLWrana8_wSB7bKP0gbjZ9Bg9bJH5AU5rdgsYQVmQkqQVv_UrNswo1nee6PNCHTl7Qm0uknLPzleFNJggsbu7sQ5mgYlvXRfX89BT7Ye9Q-WGaPoZ5nafxY5CnT_SIfE9--JTEGxqcg4Q-BYeyROG_wjqNw8Ffb76b-lEcDHi2e6jx-APSReQKM-VIyxVaMDLX79KriiNE7OIVza9KMxz0WU7q4gXcZqQvEUnEzBFYmZHkbloNBNLFbSESNYLcxaPgVqwqQglPCMswrISwSPXT4ldRgbmCtoTFcKGDdiGwIitwfDcO1rjFIMJPj-mW9YDvhscAEf_PRXevNxRgRg3a_OoWGRqyPNctvr5pU1w1BEf3nhMDrhtXGBOMJR0w3Xid5OvBtkcLbvNJ1PFbkW1YOJruaLFVY4nS6cUFf7h4PfsLJVm-1RfTsI_rl4Qs35Lzh5tHDmucP1RluHxjEjNani03OFva8lN192TqmoLtmidrBzCNHTLVtqqz8_hXlXmxidlMXEewvWU9deE8VRkdZcyrqxjOBc62O5orntEEL6fkLfoYefp25u4XsaLpzGS3-FTe5ObjfbQNIq-Z7OBgvUcQRbrd-p-gj8lOb6wzz1kjUASY--B84XxQoJ8rSqvXTHa6NhVFTxRnl57B-Rg5nilMXIrz8GtuNQfQtQaGE9E6LxnqBqLjmrBZ3UczdMDRItjEtd0DJZ8FLA9Dh3d3OYTa7tCfirvZ-P3kDC_g5OpiYc0hWbgvM14ZA_CgA8xBceAFYO0ADoVpf2GYvy5UxnbJaOTqXrRY2IqpJ1Pn3B-Ds_UmHjNz9on1-oFn2H_NN47Sa_Rxy69Ho4V32vVgbqoLI8onWvwOifapGLlhXwLG2g0lWJqqqmEsdsFpqkWnGCb3c60kJQVTycGETxiroqE2EA0BoskYdg3esyNMr7b6aaNFODcFaxcRm2tOlrZvLFcntrZkNxrmgLtn7R3mxsWWLNdibc07Td2Im743zUYlaI_xi7QcovUFol1RniQnPWTxYeMdXyL9ZFKjb4yFIR_kq_Hia4TeSoZJi9eqV8kRj9wus_Tmn_Hx--kUOat8Koqf5Ia_xv0adfgTvwDRTQ6gQeTYwcWSbHeRpTN-r10mZhYFRGAuvVdovnAzdZU61tdCa459XftmkCa94o6DVQUWSJGl4d3qvV9sZryOWXEblqeOFJvuAo2MmSaEveN45X6zpVu3NxP2unsaZsx0CL2yHymznngyV2NrHOnC8fICa8DsHgvKJNVcCLaGtYXGcT7CIzc10tXLrLf_8YPnOynfTvHbAwf_U57me7Ibg6fswZul4H18ggeV492Rk71xX2gKZvJtYBzOv-JsMMCODJj-AX_6waNV71e8ab7APnFSWP8hJz9hCRUslev73rYCq7gY_KVqYlEq6Xwb0wSr9cZPKixMXxWofaX9OLrW8873830vUGpJsTwe6bU7uL73q-9X-w6CsFVnhyT5odXy1I-zqvP8705Ds_BctZWs8zvWZJoktPy158BXmSVin2JXQlmRspRK3EpmOQ4d3NBmW0jakvw3yJD6yX0X9FTTMmslt6p921vo1_R6IIps2CfJMhY3jpMF_qFT3rqgG7wzd8HW4FXDLdAsKM9FTcM2hncPeRlHcYaxyaLvytl5btsn7HfKvMGbaAu-v_HLGElCesQG6IvcurZKKPWL2Cj9lLZdWBsq5oH9_b__A8AqCTc\u0026__EXTRA__=__MACROS__]]\u003e\u003c/Impression\u003e\u003cImpression\u003e\u003c![CDATA[https://t-odx.op-mobile.opera.com/impr?a=a5522269118912\u0026adomain=wendys.com\u0026b=949295212\u0026bl=1\u0026cc=US\u0026cid=40021128\u0026cm=1\u0026crid=80859_55272534\u0026ct=\u0026dbbf=1.000\u0026dbmb=0.000\u0026dvt=PHONE\u0026ep=ep7270864899840\u0026ext=Cp85eZi9P-BMmPnnTXQkoyNWkG6zxQvYxK83vbIIOtLV8b2-5i1rN_9s1h8tAP9F15_Kezbi9-11_mALy0S-z7xcitQPmr18NGda6X6gOflvDwEB6zapWeuLZXXlfYqG_nFwff_rD0LbonlzBEY3Y3FZNwYObE6MejELWJEmg2g%3D\u0026gi=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026iabCat=IAB8-9\u0026impr_dl=1669969833967\u0026it=1\u0026m=m5522269119233\u0026ot=IOS\u0026pubId=pub7226695071424\u0026rip=172.58.103.70\u0026s=s7226745164736\u0026said=f5540250981342e753b259364d1da363\u0026sc=-1\u0026schp=-1\u0026se=0002076e700e27e2ee83\u0026spid=c1be64e41556140ba4007003a6a838e5\u0026srid=9a191d46-6054-4869-a547-9ef8d2426c55\u0026ss=3\u0026u=7cc3bfc211ced95f\u0026va=a5522269118912\u0026vd=120\u0026vm=m5522269119233]]\u003e\u003c/Impression\u003e\u003cImpression\u003e\u003c![CDATA[https://opr.adx.opera.com/i?clid=opr\u0026paid=opr\u0026dvid=v\u0026avid=adv5522269118272\u0026caid=o5522269118720\u0026plid=m5522269119233-80859_55272534\u0026siteId=app7226709609152\u0026kv7=pub7226695071424\u0026publisherId=c1be64e41556140ba4007003a6a838e5\u0026kv1=\u0026kv4=172.58.103.70\u0026kv3=7cc3bfc211ced95f\u0026kv10=\u0026kv12=s7226745164736\u0026kv25=TriPeaks+Solitaire%3A+Card+Game\u0026kv15=US\u0026kv16=29.64710000\u0026kv17=-95.48710000\u0026kv18=949295212\u0026kv19=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026kv28=IPHONE\u0026kv26=ios\u0026kv23=\u0026kv11=0002076e700e27e2ee83\u0026kv27=Mozilla%2F5.0+%28iPhone%3B+CPU+iPhone+OS+15_5+like+Mac+OS+X%29+AppleWebKit%2F605.1.15+%28KHTML%2C+like+Gecko%29+Mobile%2F15E148\u0026lineItemId=a5522269118912\u0026kv29=[ERRORCODE]\u0026kv30=[CONTENTPLAYHEAD]_[ADPLAYHEAD]\u0026kv33=[ASSETURI]\u0026kv34=[VASTVERSIONS]\u0026kv35=[IFATYPE]\u0026kv36=[IFA]\u0026kv37=[CLIENTUA]\u0026kv38=[SERVERUA]\u0026kv39=[DEVICEUA]\u0026kv40=[DEVICEIP]\u0026kv41=[LATLONG]\u0026kv42=[DOMAIN]\u0026kv43=[PAGEURL]\u0026kv44=\u0026kv45=[PLAYERSIZE]\u0026kv46=[REGULATIONS]\u0026kv47=[ADTYPE]\u0026kv48=[TRANSACTIONID]\u0026kv49=[BREAKPOSITION]\u0026kv50=[APPNAME]\u0026kv51=[PLACEMENTTYPE]\u0026kv54=[LAT]\u0026kv5=OpenRTB\u0026kv24=Mobile_InApp_Video]]\u003e\u003c/Impression\u003e\u003cError\u003e\u003c![CDATA[https://t-odx.op-mobile.opera.com/video?a=a5522269118912\u0026adomain=wendys.com\u0026b=949295212\u0026cc=US\u0026cid=40021128\u0026cm=1\u0026crid=80859_55272534\u0026ct=\u0026dbbf=1.000\u0026dbmb=0.000\u0026dvt=PHONE\u0026ep=ep7270864899840\u0026ext=Cp85eZi9P-BMmPnnTXQkoyNWkG6zxQvYxK83vbIIOtLV8b2-5i1rN_9s1h8tAP9F15_Kezbi9-11_mALy0S-z7xcitQPmr18NGda6X6gOflvDwEB6zapWeuLZXXlfYqG_nFwff_rD0LbonlzBEY3Y3FZNwYObE6MejELWJEmg2g%3D\u0026gi=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026iabCat=IAB8-9\u0026impr_dl=1669969833967\u0026it=1\u0026m=m5522269119233\u0026ot=IOS\u0026pubId=pub7226695071424\u0026rip=172.58.103.70\u0026s=s7226745164736\u0026said=f5540250981342e753b259364d1da363\u0026sc=-1\u0026schp=-1\u0026se=0002076e700e27e2ee83\u0026spid=c1be64e41556140ba4007003a6a838e5\u0026srid=9a191d46-6054-4869-a547-9ef8d2426c55\u0026ss=3\u0026u=7cc3bfc211ced95f\u0026va=a5522269118912\u0026vd=120\u0026vel=1\u0026vm=m5522269119233\u0026evt=error]]\u003e\u003c/Error\u003e\u003cCreatives\u003e\u003cCreative id=\"80859_55272534\"\u003e\u003c/Creative\u003e\u003c/Creatives\u003e\u003cExtensions\u003e\u003c/Extensions\u003e\u003c/Wrapper\u003e\u003c/Ad\u003e\u003c/VAST\u003e",
										"adomain":[
											"wendys.com"
										],
										"cid":"40021128",
										"crid":"80859_55272534",
										"cat":[
											"IAB8-9"
										],
										"exp":1800,
										"ext":{
					
										}
									}
								],
								"seat":"adv5522269118272"
							}
						],
						"bidid":"0002076e700e27e2ee83",
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_mobupps_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"xp90l0uqtll0c5c6kl2",
						"seatbid":[
							{
								"bid":[
									{
										"id":"000206e636279f4cfc41",
										"impid":"1",
										"price":0.073,
										"nurl":"https://t.adx.opera.com/win?a=a7314230524032\u0026adomain=litres.ru\u0026b=turkru.tv\u0026cc=RU\u0026cid=4028988167164987656\u0026cm=1\u0026crid=4028988167164987656\u0026ct=\u0026dvt=DESKTOP\u0026ep=ep7226729523904\u0026ext=JffUhbyX6FhEBira-RWv6HNh2Ei8QkcNRN5yyqJH1rtkdBp9YKcbFALlZBFdTs6VMAPuCCUKOMEiXBavLB7F-yg189l-sy7FGRsPk5h5kx1WlXlUxQp26AXhFodngyguN8t8AyP-RWcy2oW7jNeGYBYzVLg8K0dLdhCFXMmS86A%3D\u0026impr_dl=1660827850378\u0026m=m7314230524544\u0026msz=300x250\u0026ot=WINDOWS\u0026pubId=pub7226695071424\u0026radv=1495\u0026rip=94.233.234.31\u0026s=s7226771668160\u0026said=fef65162602c7fc04d1f266b49d775e7\u0026sc=-1\u0026se=000206e636279f4cfc41\u0026sh=-1\u0026spid=31\u0026srid=xp90l0uqtll0c5c6kl2\u0026ss=3\u0026u=630a321bb1d3c443\u0026va=a7314230524032\u0026vm=m7314230524544\u0026ac=${AUCTION_CURRENCY}\u0026ap=${AUCTION_PRICE}",
										"lurl":"https://t.adx.opera.com/loss?a=a7314230524032\u0026adomain=litres.ru\u0026b=turkru.tv\u0026cc=RU\u0026cid=4028988167164987656\u0026cm=1\u0026crid=4028988167164987656\u0026ct=\u0026dvt=DESKTOP\u0026ep=ep7226729523904\u0026ext=JffUhbyX6FhEBira-RWv6HNh2Ei8QkcNRN5yyqJH1rtkdBp9YKcbFALlZBFdTs6VMAPuCCUKOMEiXBavLB7F-yg189l-sy7FGRsPk5h5kx1WlXlUxQp26AXhFodngyguN8t8AyP-RWcy2oW7jNeGYBYzVLg8K0dLdhCFXMmS86A%3D\u0026impr_dl=1660827850378\u0026m=m7314230524544\u0026msz=300x250\u0026ot=WINDOWS\u0026pubId=pub7226695071424\u0026radv=1495\u0026rip=94.233.234.31\u0026s=s7226771668160\u0026said=fef65162602c7fc04d1f266b49d775e7\u0026sc=-1\u0026se=000206e636279f4cfc41\u0026sh=-1\u0026spid=31\u0026srid=xp90l0uqtll0c5c6kl2\u0026ss=3\u0026u=630a321bb1d3c443\u0026va=a7314230524032\u0026vm=m7314230524544\u0026al=${AUCTION_LOSS}",
										"adm":"\u003cscript type=\"text/javascript\"\u003e\n  function changeURLArg(url, arg, arg_val) {\n    var pattern = arg + '=([^\u0026]*)';\n    var replaceText = arg + '=' + arg_val;\n    if (url.match(pattern)) {\n      var tmp = '/(' + arg + '=)([^\u0026]*)/gi';\n      tmp = url.replace(eval(tmp), replaceText);\n      return tmp;\n    } else {\n      if (url.match('[\\?]')) {\n        return url + '\u0026' + replaceText;\n      } else {\n        return url + '?' + replaceText;\n      }\n    }\n  }\n  function tracking(url) {\n    try {\n      navigator.sendBeacon(url);\n    } catch (error) {\n    }\n  }\n  function clickTracks(clkType) {\n    var urls = \"https://t.adx.opera.com/click?a=a7314230524032\u0026adomain=litres.ru\u0026b=turkru.tv\u0026cc=RU\u0026cid=4028988167164987656\u0026cm=1\u0026crid=4028988167164987656\u0026ct=\u0026dvt=DESKTOP\u0026ep=ep7226729523904\u0026ext=JffUhbyX6FhEBira-RWv6HNh2Ei8QkcNRN5yyqJH1rtkdBp9YKcbFALlZBFdTs6VMAPuCCUKOMEiXBavLB7F-yg189l-sy7FGRsPk5h5kx1WlXlUxQp26AXhFodngyguN8t8AyP-RWcy2oW7jNeGYBYzVLg8K0dLdhCFXMmS86A%3D\u0026m=m7314230524544\u0026msz=300x250\u0026ot=WINDOWS\u0026pubId=pub7226695071424\u0026radv=1495\u0026rip=94.233.234.31\u0026s=s7226771668160\u0026said=fef65162602c7fc04d1f266b49d775e7\u0026sc=-1\u0026se=000206e636279f4cfc41\u0026sh=-1\u0026spid=31\u0026srid=xp90l0uqtll0c5c6kl2\u0026ss=3\u0026u=630a321bb1d3c443\u0026va=a7314230524032\u0026vm=m7314230524544\u0026ac=${AUCTION_CURRENCY}\u0026ap=${AUCTION_PRICE}\u0026ir=${IS_PREVIEW}\";\n    urls = urls.split(\",\");\n    for (var i = 0; i \u003c urls.length; i++) {\n      var url = changeURLArg(urls[i], 'ct', clkType);\n      tracking(url);\n    }\n  }\n  var clickHandler = function() {\n    clickTracks(1)\n  }\n  var blurHandler = function() {\n    if (document.activeElement.tagName === \"IFRAME\") {\n      clickTracks(2)\n    }\n  }\n  var touchstartHandler = function() {\n    clickTracks(3)\n  }\n\n  document.addEventListener(\"DOMContentLoaded\", function () {\n    document.body.addEventListener('click', clickHandler);\n    document.body.addEventListener('touchstart', touchstartHandler);\n    window.addEventListener(\"blur\", blurHandler);\n  });\n\n  document.addEventListener(\"unload\", function () {\n    document.body.removeEventListener('click', clickHandler);\n    document.body.removeEventListener('touchstart', touchstartHandler);\n    window.removeEventListener(\"blur\", blurHandler);\n  });\n\n\u003c/script\u003e\u003chead\u003e\u003cstyle type='text/css'\u003ebody {margin:auto auto;text-align:center;} \u003c/style\u003e\u003c/head\u003e\u003cdiv id=\"-oadx_000206e636279f4cfc41m73142305245441660826050379648726\"\u003e\u003cdiv\u003e\u003chtml\u003e\n\u003chead\u003e\n\u003cmeta charset=\"UTF-8\"\u003e\n\u003cmeta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/\u003e\u003c/head\u003e\n\u003cbody style=\"margin: 0px; overflow:hidden;\"\u003e \n\u003cscript type=\"text/javascript\"\u003e\nwindow.tbl_trc_domain = 'am-trc.taboola.com';\nwindow._taboola = window._taboola || [];\n_taboola.push({home:'auto'});\n!function (e, f, u, i) {\nif (!document.getElementById(i)){\ne.async = 1;\ne.src = u;\ne.id = i;\nf.parentNode.insertBefore(e, f);\n}\n}(document.createElement('script'),\ndocument.getElementsByTagName('script')[0],\n'//cdn.taboola.com/libtrc/example-placements/loader.js',\n'tb_loader_script');\nif(window.performance \u0026\u0026 typeof window.performance.mark == 'function')\n{window.performance.mark('tbl_ic');}\n\u003c/script\u003e\n\n\u003cdiv id=\"taboola-below-article-thumbnails\" style=\"height: 250px; width: 300px;\"\u003e\u003c/div\u003e\n\u003cscript type=\"text/javascript\"\u003e\nwindow._taboola = window._taboola || [];\n_taboola.push({\nmode: 'SCOD_Test_300x250_1x1',\ncontainer: 'taboola-below-article-thumbnails',\nplacement: 'turkru.tv_Ppub7226695071424_Sfef65162602c7fc04d1f266b49d775e7_W300_H250_N1',\ntarget_type: 'mix',\n\"rtb-win\":{ \nbi:'000206e636279f4cfc41',\ncu:'USD',\nwp:'0.09109',\nwcb:'http://imprammp.taboola.com/st?cipid=66576435\u0026ttype=0\u0026cirid=08a8d08b-24c4-4e1d-b0e5-0b34d1e4a882\u0026cicmp=67154105\u0026cijs=1\u0026dast=V7qJoBxgHm25KoM64CXQN1XcaZDemyzAR1XcaZDemyzAUAAAADBg0HJDBcGEeG4cQtGW1Ma9HKMXJLDCvXWjDxjEaOlWlhHE7WEIGm0-Fz3etFr8vX8rqLbncAAAAAeAAou5yC6NFtXI_o0W1clwAAAACgCKj4txC4AAAAAMD4____zwBIQLjVAIDlMHCG9fKw-wMA4OEBBABAAIMEAAAAoAQAAADg5P________8_ZuAefEaGAAA298agp-DaBwBACAAAIGtIAR6OB04ZpI4oQa-IEQAAAIBXuvrH0aROqLqq_v__-60ArgAAAgRh5b8wsnp1G9e3JmgwGS6Hw8Vst5iNlsPdbDUbVwI7TZeX5y55vQtBBpvNbLRczGuCBpPhcjhczHaL2Wg53M1Ws4UBAAAAjC3o3_L3XG53jd_tsv________9_s_-zfzQhF4WeNCEZMzO1Ht3G9bVfQAAANwCAtwC4mAOwAwAAALj7____95Ow0XK1GQ2Wo-FutRgOBqvl_gZiMBrgRAyWy8lkMdmtRqvRZrgbzQYLFIjBBCdkONpMVqPdajdZDiej0Wwz2SBFq1az0WYwXM0ms91uNRwMl6MRUrRmMZtMFrPRcrcZLCejwXAy3CGEbQYLz2QxsThGno1ptBlihm0GC89kMbE4Rp6NabTZE16Wh6fDN9ERAAAg4LEhR3PJZjOXbEZzzWKTIOay2VaL2WQ2mGx8M49hNHLMLLPZxLQc-XarlW-VP7dLxqfvnbKWLgXkdU29ZFB4YyWDQhsrGRTWWMmgUDZjJYNCGisZFMpYyaAwxkoGhTNWMiiUwVjJoDCWYyWDwhiOlQwKYzdWMiiM2VjJoDBWYyWDwhiNlQwKYzNWMiiMyVjJoDAWYyWDwhiMlQwKc6xkUIhjJYNCWWwTAAAAgBMCXhffZDKbLVeD3WI0GS05wzaDhWeymFgcI8_GNNrsCS_Lw9Phm9gsGgAAgIAc!\u0026excid=93\u0026tst=1\u0026docw=0\u0026s2s=true',\nrt:'1660826050326',\nrdc:'am.taboolasyndication.com',\nti:'5648',\nex:'OperaSCoD',\nbs:'fef65162602c7fc04d1f266b49d775e7',\nbp:'pub7226695071424',\nbd:'turkru.tv',\nsi:'17916'\n} \n,\nrec: {\"trc\":{\"si\":\"cb2411907026f3e933c0b5b1034f8eca\",\"sd\":\"v2_cb2411907026f3e933c0b5b1034f8eca_a8ae62c2-4760-42fd-818c-102827ee9605-tuct9f7b742_1660826050_1660826050_CIi3jgYQvt1UGIq6g8euqtuzsAEgASgBMLkBOMnrDEDymRBIs9naA1DAnydYAGAAaKjtzdPIpeaqL3AA\",\"ui\":\"a8ae62c2-4760-42fd-818c-102827ee9605-tuct9f7b742\",\"plc\":\"DESK\",\"wi\":\"-3325698876771888143\",\"cc\":\"RU\",\"route\":\"AM:AM:V\",\"el2r\":[\"bulk-metrics\",\"debug\",\"social\",\"abtests\",\"metrics\",\"perf\",\"supply-feature\"],\"uvpw\":\"1\",\"pi\":\"1388222\",\"cpb\":\"GNO629MGIJz__________wEqGWFtLnRhYm9vbGFzeW5kaWNhdGlvbi5jb20yDXRyYy1zY29kNDAxMjM4gNSn9wVAyesMSPKZEFCz2doDWMCfJ2MI_hYQiiAYE2RjCNcWENUfGCNkYwjSAxDgBhgIZGMIlhQQmRwYGGRjCIMvENc-GAlkYwj0FBCeHRgfZGMIpCcQijUYL2RqFDAwMDIwNmU2MzYyNzlmNGNmYzQxeAGAAQKIAaHXqM8BkAEc\",\"evh\":\"217325835\",\"evi\":{\"47\":\"5028|6794\"},\"vl\":[{\"ri\":\"9c405fad219c85984fef8ea32b187f9d\",\"uip\":\"turkru.tv_Ppub7226695071424_Sfef65162602c7fc04d1f266b49d775e7_W300_H250_N1\",\"ppb\":\"CIoF\",\"estimation_method\":\"EcpmEstimationMethodType_ESTIMATION\",\"original_ecpm\":\"0.11386413645642028\",\"baseline_variant\":\"false\",\"v\":[{\"thumbnail\":\"https://cdn.taboola.com/libtrc/static/thumbnails/0152b8e34e82f950278bd82e94be9878.gif\",\"all-thumbnails\":\"https://cdn.taboola.com/libtrc/static/thumbnails/195174768ca56fdc6f0c1ec91da69be0.jpg!-#@1000x600\",\"is-gift\":\"true\",\"origin\":\"default\",\"thumb-size\":\"1000x600\",\"description\":\"\",\"placement-label\":\"turkru.tv_Ppub7226695071424_Sfef65162602c7fc04d1f266b49d775e7_W300_H250_N1\",\"title\":\"ЛитРес: подписка за 399₽/месяц\",\"type\":\"text\",\"published-date\":\"1660217303\",\"branding-text\":\"ЛитРес\",\"url\":\"https://www.litres.ru/litres_subscription/?lfrom=1025278492\u0026utm_source=taboola\u0026utm_medium=video\u0026utm_campaign=podpiska\u0026tblci=GiCbMvUTXFYb4FcvwNWK0lmgVpboQvkCjCzJ1vGuNL25ZiC19VUouYKR6J7Bofct#tblciGiCbMvUTXFYb4FcvwNWK0lmgVpboQvkCjCzJ1vGuNL25ZiC19VUouYKR6J7Bofct\",\"duration\":\"0\",\"balancedPacingValue\":\"1.0\",\"sig\":\"f4ef82007d0a712c894d19c39a0e1a169f12da4e0f01\",\"constraintType\":\"SCALE\",\"item-id\":\"~~V1~~4028988167164987656~~DxPudjVzzvbSQ2dLcDfidnsDXnQuMLP-Dj1qSg5qobzTxvAnL2wqac4MyzR7uD46gj3kUkbS3FhelBtnsiJV6MhkDZRZzzIqDobN6rWmCPA3hYz5D3PLat6nhIftiT1lStQk6jaQcGGDUiGnTk93_4FS8qr-f53vDbiQQMMys25Xvi4P24MPI4rQxfLITK8T8ywME2ZwzKigjRsQeO32V8hbD21urIhLvAJGMPjxI-zqGDWsdg__ILQP_Zbbaqe-UWSjD2vpf1qi2yEN9nyKxQB8MQ5DwSrwujBax9ayWUY\",\"adomain\":\"litres.ru\",\"uploader\":\"\",\"cta-text\":\"Купить Сейчас\",\"is-syndicated\":\"true\",\"publisher\":\"aitagetlitres-sc\",\"id\":\"~~V1~~4028988167164987656~~DxPudjVzzvbSQ2dLcDfidnsDXnQuMLP-Dj1qSg5qobzTxvAnL2wqac4MyzR7uD46gj3kUkbS3FhelBtnsiJV6MhkDZRZzzIqDobN6rWmCPA3hYz5D3PLat6nhIftiT1lStQk6jaQcGGDUiGnTk93_4FS8qr-f53vDbiQQMMys25Xvi4P24MPI4rQxfLITK8T8ywME2ZwzKigjRsQeO32V8hbD21urIhLvAJGMPjxI-zqGDWsdg__ILQP_Zbbaqe-UWSjD2vpf1qi2yEN9nyKxQB8MQ5DwSrwujBax9ayWUY\",\"views\":\"0\"}]}],\"tslt\":{\"p-video-overlay\":{\"cancel\":\"Cancel\",\"goto\":\"Go To\"},\"read-more\":{\"DEFAULT_CAPTION\":\"Read%20More\"},\"next-up\":{\"BTN_TEXT\":\"Read Next Story\"},\"vignette\":{\"openBtn\":\"Learn More\",\"closeBtn\":\"Close\"},\"time-ago\":{\"yesterday\":\"Yesterday\",\"hours\":\"{0} hours ago\",\"hour\":\"1 hour ago\",\"minutes\":\"{0} minutes ago\",\"now\":\"Now\",\"today\":\"Today\",\"days\":\"{0} days ago\"},\"explore-more\":{\"POPUP_TEXT\":\"More stories to check out before you go\",\"TITLE_TEXT\":\"Keep on reading\"},\"adchoice\":{\"adChoiceBtn.title\":\"Why do I see this item?\"},\"userx\":{\"popover.content.questionnaire.options.uninteresting\":\"Uninteresting\",\"popover.content.questionnaire.options.racy\":\"Vulgar/Racy\",\"undoBtn.label\":\"Undo\",\"popover.title.scRemoved\":\"Sponsored link removed\",\"popover.content.questionnaire.options.repetitive\":\"Repetitive\",\"popover.title.thankYou\":\"Thank You!\",\"popover.title.removed\":\"Removed!\",\"popover.content.questionnaire.options.offensive\":\"Offensive\",\"popover.content.questionnaire.options.misleading\":\"Misleading\",\"removeBtn.title\":\"Remove this item\",\"popover.content.questionnaire.tellUsWhy\":\"Tell us why?\",\"popover.content.questionnaire.options.other\":\"Other\",\"popover.content.approval\":\"We will try not to show you this content anymore.\"}},\"cpcud\":{\"upc\":\"0.0\",\"upr\":\"0.0\"},\"dcga\":{\"pubConfigOverride\":{\"border-color\":\"black\",\"font-weight\":\"bold\",\"inherit-title-color\":\"true\",\"module-name\":\"cta-lazy-module\",\"enable-call-to-action-creative-component\":\"true\",\"disable-cta-on-custom-module\":\"true\"}}}}\n});\n\u003c/script\u003e\n\n\u003cscript type=\"text/javascript\"\u003e\nwindow._taboola = window._taboola || [];\n_taboola.push({flush: true});\n\u003c/script\u003e\n\n\u003c/body\u003e\n\u003c/html\u003e\u003c/div\u003e\u003cimg id='adxImpressionTrackingPixel0' alt=\"\" src=\"https://t.adx.opera.com/impr?a=a7314230524032\u0026adomain=litres.ru\u0026b=turkru.tv\u0026bl=1\u0026cc=RU\u0026cid=4028988167164987656\u0026cm=1\u0026crid=4028988167164987656\u0026ct=\u0026dvt=DESKTOP\u0026ep=ep7226729523904\u0026ext=JffUhbyX6FhEBira-RWv6HNh2Ei8QkcNRN5yyqJH1rtkdBp9YKcbFALlZBFdTs6VMAPuCCUKOMEiXBavLB7F-yg189l-sy7FGRsPk5h5kx1WlXlUxQp26AXhFodngyguN8t8AyP-RWcy2oW7jNeGYBYzVLg8K0dLdhCFXMmS86A%3D\u0026impr_dl=1660827850378\u0026m=m7314230524544\u0026msz=300x250\u0026ot=WINDOWS\u0026pubId=pub7226695071424\u0026radv=1495\u0026rip=94.233.234.31\u0026s=s7226771668160\u0026said=fef65162602c7fc04d1f266b49d775e7\u0026sc=-1\u0026se=000206e636279f4cfc41\u0026sh=-1\u0026spid=31\u0026srid=xp90l0uqtll0c5c6kl2\u0026ss=3\u0026u=630a321bb1d3c443\u0026va=a7314230524032\u0026vm=m7314230524544\u0026ac=${AUCTION_CURRENCY}\u0026ap=${AUCTION_PRICE}\" style='width:0px;height:0px;border:0px;margin:0px;float:left;'/\u003e\u003cimg id='adxImpressionTrackingPixel1' alt=\"\" src=\"https://adrta.com/i?clid=opr\u0026paid=opr\u0026avid=adv7004556758144\u0026kv7=pub7226695071424\u0026publisherId=31\u0026plid=m7314230524544-4028988167164987656\u0026siteId=app7226716562816\u0026caid=o7004556758592\u0026lineItemId=a7314230524032\u0026kv1=300x250\u0026kv2=https%3A%2F%2Fturkru.tv\u0026kv3=630a321bb1d3c443\u0026kv4=94.233.234.31\u0026kv10=\u0026kv11=000206e636279f4cfc41\u0026kv12=s7226771668160\u0026kv15=RU\u0026kv16=43.22860000\u0026kv17=44.78190000\u0026kv18=\u0026kv19=\u0026kv28=\u0026kv23=\u0026kv25=\u0026kv26=windows\u0026kv27=Mozilla%2F5.0+%28Windows+NT+10.0%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F104.0.5112.81+Safari%2F537.36+Edg%2F104.0.1293.54\u0026kv5=OpenRTB\u0026kv24=Mobile_Web\" style='width:0px;height:0px;border:0px;margin:0px;float:left;'/\u003e\u003cimg id='adxImpressionTrackingPixel2' alt=\"\" src=\"https://sync.taboola.com/sg/OperaSCoD/1/cm\" style='width:0px;height:0px;border:0px;margin:0px;float:left;'/\u003e\u003c/div\u003e",
										"adomain":[
											"litres.ru"
										],
										"cid":"4028988167164987656",
										"crid":"4028988167164987656",
										"h":250,
										"w":300,
										"exp":1800,
										"ext":{
					
										}
									}
								],
								"seat":"adv7004556758144"
							}
						],
						"bidid":"000206e636279f4cfc41",
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_mobupps_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
					"id":"e657d364-1164-49bf-b8cc-184e04737288",
					"bidid":"c6ca008e9088c951f3ad52e5dc4960f0",
					"seatbid":[
						{
							"bid":[
								{
									"id":"a6fe364e662e032e7d91e01eeb83fdbb",
									"impid":"1",
									"price":0.11284,
									"nurl":"http://b1.dcntr-ads.com/?win=nurl&sp=${AUCTION_PRICE}&t=native&uniq=511763b40a8546344164ad9895147f71",
									"adm":"{\"native\":{\"ver\":\"1.1\",\"assets\":[{\"id\":1,\"required\":1,\"title\":{\"text\":\"Badoo\"}},{\"id\":2,\"img\":{\"url\":\"https://cdn.liftoff-creatives.io/customers/e12978da97/image/lambda_png/25d0fb7ccac186abdc31.png\",\"w\":1200,\"h\":627}},{\"id\":3,\"required\":1,\"img\":{\"url\":\"https://cdn.liftoff-creatives.io/customers/e12978da97/image/lambda_jpg_89/5b8003ee61c3a4c3114e.jpg\",\"w\":80,\"h\":80}},{\"id\":4,\"required\":1,\"data\":{\"value\":\"Otro estilo de citas. Añade tus intereses y habla antes de hacer match en Badoo.\",\"len\":80}},{\"id\":5,\"data\":{\"value\":\"INSTALAR\",\"len\":8}}],\"link\":{\"url\":\"https://play.google.com/store/apps/details?id=com.badoo.mobile&referrer=Liftoff_v.2_g.161950_a.0005076e65f4af012640_2832_c.123_t.ua_u.e5c4cb4f3d4a2167\",\"clicktrackers\":[\"https://click.liftoff.io/v1/campaign_click/Ch2McWAd1-6xzddPHmqZI6fgdFk_QUkhDlu2gVF0KFVthn--5jPDJL1uCORiMXvfSFTste9D-5snNiJl8USUrTkJfiXeLpLecjtEKa4cjpXe1Khf_-9p-jOOrSDtmGt025yj5ZuYDFcbgeFXyvz8D0-AoXlvRsdz6rV1fRafvr845eoYUu62nxVWWU58OIDug9etxabhuUsWRpC8QNu4LUu6YV3mQgFYycKCCQVkhUieyej2texwvMwBk2gx5wvIe2NFPgEX52XTaOyy-5AiQxVCTJdtLWtlxt8guUDIY7msrLkn462AqVTGVV7Y1tzVA6uAJK3BchrGLTfsqbp54ga4xaSYjGKtuOyOect_EZ11-NoaMkTPN6rdv87mqtnLDfYFBGbS9fwLauWflAas064gck6o_rIWMjlSip1ON6Z9wX1Q7TLL2aYdjgIhgdFGq9uT5slwnmt435mirs8u1zUwlJ4ju1SwTF7Ezb9shIum2WZq5t0WYJiRn8U4uamyvYKcT1SLmQw7omEP2Xx9kbTH3laHmLH7iUW2bl9exCtXRPbvtXuZytdAq-iD1qRA7e4hF1cABTfmwrgyfew7gyqffMkpuxq_P93xx89cxwjqILc\",\"https://t-odx.op-mobile.opera.com/click?a=a7005769032448&adomain=badoo.com&b=com.etermax.preguntados.lite&cc=ES&ch=third_party_sdk&cid=161950&cm=1&crid=304461&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep5738257056320&ext=awkKsd4F1X7kXA1Rd9wFUmHhnqyUMGIwf8CcUlqaL4Ioq0Bx0x13G-UL0R1l9-QejHDBC-pO0ERMNujZ9J_m6bWToxomGPKKkTzRPUqeqdIUpPXshi-eB5sTt6Be2h3C9vGeSF5arsi5HOtfG2d0Ag%3D%3D&gi=1B3E178C-313E-42D1-A607-8E3DFB6FDFE5&iabCat=IAB14-1&m=m7005769032832&msz=1200x627&ot=ANDROID&pc=1&pubId=pub4903319265728&rip=88.29.175.23&s=s4903325786432&said=af6df55abb31&sc=2&schp=1-2&se=0005076e65f4af012640&spid=194724177&srid=66e9fc2c806e0bf7b554d910a762&ss=2&stid=third_party_sdk&u=32871b0f481d9e97&va=a7005769032448&vm=m7005769032832&ac=USD&ap=0.18596&ir=${IS_PREVIEW}\"]},\"imptrackers\":[\"https://impression-europe.liftoff.io/opera/impression?ad_group_id=161950&channel_id=123&creative_id=304461&device_id_sha1=e5c4cb4f3d4a2167961759f3327f1846e2573093&source_app_store_id=com.etermax.preguntados.lite&bid_ts=1669965386480&auction_id=ocmCKsXcknnrLiQNXOWC&auction_price=0.245&origin=haggler-opera022\",\"https://opr.adx.opera.com/i?clid=opr&paid=opr&avid=adv7005106418624&caid=o7005106419136&plid=m7005769032832-304461&siteId=app4903323212480&kv7=pub4903319265728&publisherId=194724177&kv1=1200x627&kv4=88.29.175.23&kv3=32871b0f481d9e97&kv10=&kv12=s4903325786432&kv25=Trivia+Crack&kv15=ES&kv16=0.00000000&kv17=0.00000000&kv18=com.etermax.preguntados.lite&kv19=1B3E178C-313E-42D1-A607-8E3DFB6FDFE5&kv28=2201116SG&kv26=android&kv23=&kv27=Mozilla%2F5.0+%28Linux%3B+Android+11%3B+2201116SG+Build%2FRKQ1.211001.001%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F107.0.5304.141+Mobile+Safari%2F537.36&kv11=0005076e65f4af012640&lineItemId=a7005769032448&kv5=OpenRTB&kv24=Mobile_InApp_Native\",\"https://t-odx.op-mobile.opera.com/impr?a=a7005769032448&adomain=badoo.com&b=com.etermax.preguntados.lite&bl=1&cbu=H4sIAAAAAAAC_0yQu27rMBBE_4adZfEhyrqAcAuXCRKkijuBJpfSIhKXWFJGPj-wiiDlnHkUs9Say7_zGbfMUApSOsHOlKFZMVaKsUE6UwZ2fyL_XZhmpj1PGEZp5dC1wi8uJVgPorTwDK7iA55at8ZYKQI80D_BVBYnR-i88XcTdTBOSdsPVvbdELVWfZQXY0F1vW4HLQrt7GFyOU-lEh-TnrYGKvDmvpvMMO-pukClWbGCuGOYahmltcNgO32x5tIKt_uKlJ5l8tv1pdz8V0r8ih9vt_fP66-fGT2MbaNMJ4hxxjQubp5X4NNxQ6vUTwAAAP__IfgYdzQBAAA&cc=ES&ch=third_party_sdk&cid=161950&cm=1&crid=304461&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep5738257056320&ext=awkKsd4F1X7kXA1Rd9wFUmHhnqyUMGIwf8CcUlqaL4Ioq0Bx0x13G-UL0R1l9-QejHDBC-pO0ERMNujZ9J_m6bWToxomGPKKkTzRPUqeqdIUpPXshi-eB5sTt6Be2h3C9vGeSF5arsi5HOtfG2d0Ag%3D%3D&gi=1B3E178C-313E-42D1-A607-8E3DFB6FDFE5&iabCat=IAB14-1&impr_dl=1669967186547&m=m7005769032832&msz=1200x627&ot=ANDROID&pc=1&pubId=pub4903319265728&rip=88.29.175.23&s=s4903325786432&said=af6df55abb31&sc=2&schp=1-2&se=0005076e65f4af012640&spid=194724177&srid=66e9fc2c806e0bf7b554d910a762&ss=2&stid=third_party_sdk&u=32871b0f481d9e97&va=a7005769032448&vm=m7005769032832&ac=USD&ap=0.18596&ir=${IS_PREVIEW}\",\"https://ib1.bcc-ads.com/?win=impr&price=0.13715&prt=885_undefined&t=native&uniq=37e65a366d2df6ac4072918ea9367e85\",\"\",\"https://adrta.com/i?clid=deca&paid=deca&dvid=v&avid=opera_native_ron_11aug&caid=161950&publisherId=194724177&siteId=&plid=304461&kv3=a6dbe90ecdd44ec67064f771b914d44f3dbd9868&kv5=bidscube&kv16=&kv17=&kv23=Telefonica%20de%20Espana&kv1=undefinedxundefined&kv2=https%3A%2F%2Fplay.google.com%2Fstore%2Fapps%2Fdetails%3Fid%3Dcom.etermax.preguntados.lite&kv4=88.29.175.23&kv7=bidscube___dcntr&kv11=06a6e5af661c24f8f5527649ef1745f4&kv12=562488d7747b16478d2b0410a449&kv18=com.etermax.preguntados.lite&kv19=1B3E178C-313E-42D1-A607-8E3DFB6FDFE5&kv28=2201116SG&kv25=com.etermax.preguntados.lite&kv26=android&kv27=Mozilla%2F5.0%20(Linux%3B%20Android%2011%3B%202201116SG%20Build%2FRKQ1.211001.001%3B%20wv)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Version%2F4.0%20Chrome%2F107.0.5304.141%20Mobile%20Safari%2F537.36&kv24=Mobile_InApp_Native\",\"https://b1.dcntr-ads.com/?win=impr&price=${AUCTION_PRICE}&prt=49_undefined&t=native&uniq=511763b40a8546344164ad9895147f71\",\"\",\"https://adrta.com/i?clid=deca&paid=deca&dvid=v&avid=bdscb_native_mar31&caid=885_161950&publisherId=194724177&siteId=&plid=885_304461&kv3=a6dbe90ecdd44ec67064f771b914d44f3dbd9868&kv5=dcntrads&kv16=&kv17=&kv23=Telefonica%20de%20Espana&kv1=undefinedxundefined&kv2=https%3A%2F%2Fplay.google.com%2Fstore%2Fapps%2Fdetails%3Fid%3Dcom.etermax.preguntados.lite&kv4=88.29.175.23&kv7=dcntrads___mobupps&kv11=04136d991e7fb8a481411d903508b6ba&kv12=e657d364-1164-49bf-b8cc-184e04737288&kv18=com.etermax.preguntados.lite&kv19=1b3e178c-313e-42d1-a607-8e3dfb6fdfe5&kv28=2201116SG&kv25=com.etermax.preguntados.lite&kv26=android&kv27=Mozilla%2F5.0%20(Linux%3B%20Android%2011%3B%202201116SG%20Build%2FRKQ1.211001.001%3B%20wv)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Version%2F4.0%20Chrome%2F107.0.5304.141%20Mobile%20Safari%2F537.36&kv24=Mobile_InApp_Native\"]}}",
									"adomain":[
										"badoo.com"
									],
									"cat":[
										"IAB14-1"
									],
									"adid":"55dbdb1c9686",
									"iurl":"http://b1.dcntr-ads.com/?t=preview&k=55dbdb1c9686",
									"cid":"49_885_161950",
									"crid":"49_885_304461"
								}
							],
							"seat":"49"
						}
					],
					"cur":"USD"
				}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_pangle_midas_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"DspId":25,
						"WinPrice":0.17864,
						"DspName":"pangle",
						"DspGroupName":"pangle",
						"ZoneId":"",
						"FloorPrice":0,
						"ValidImpDuration":0,
						"BidPriceCoefficient":0,
						"id":"8b82dc8c-1ba7-4b42-a5e8-79f21193ac64",
						"seatbid":[
							{
								"bid":[
									{
										"id":"8b82dc8c-1ba7-4b42-a5e8-79f21193ac64u6931-1753558363522082",
										"price":0.17864,
										"nurl":"https://api16-access-sg.pangle.io/api/ad/union/openrtb/win/?req_id=8b82dc8c-1ba7-4b42-a5e8-79f21193ac64u6931\u0026ttdsp_adx_index=229\u0026rit=980239539\u0026extra=D1dHlaitCWcbuS1tBI1IbwgqChVLXaxcUeNP9IOZIaGGPS6HUO0qMQNwt%2B8G8%2F5%2FelHJMqcla6A7bFTlZShn4FMrExljLWl3aK2AKEkc1kKUOk2rwymGzveMHxezxZGWnc5gs6L0zPb0TtEVzDejGFpFm3PIdZKaMvpPc1Gxr7XYjfResgNkndoNmOyIQQFwY0BI84o37y9mbsU4ME4CMRspo8i0PmGZ17M1G5BGJbZwoRt2wSajgkjtKBVTkbhMI%2FJC6%2BYYZIqhOSOzvGzNeUtsl7rIfgfjDvLb7o7eGT3kZWY4Dty69IbIeUWngE4wTJaENmNySlP3zspxFd6GpPa%2Fu6NDIDew6ketecr4QEsOtkXjZiZlLtPvsy7%2FA%2BHkZp9X%2BNwb8gPYCBIXAyRzo2afV%2FjcG%2FID2AgSFwMkc6M%3D\u0026win_price=${AUCTION_PRICE}\u0026auction_mwb=${AUCTION_BID_TO_WIN}\u0026use_pb=1",
										"burl":"https://api16-event-sg.pangle.io/api/ad/union/show_event/?req_id=8b82dc8c-1ba7-4b42-a5e8-79f21193ac64u6931\u0026ttdsp_adx_index=229\u0026extra=D1dHlaitCWcbuS1tBI1IbwgqChVLXaxcUeNP9IOZIaGGPS6HUO0qMQNwt%2B8G8%2F5%2FelHJMqcla6A7bFTlZShn4FMrExljLWl3aK2AKEkc1kKUOk2rwymGzveMHxezxZGWnc5gs6L0zPb0TtEVzDejGFpFm3PIdZKaMvpPc1Gxr7XYjfResgNkndoNmOyIQQFwY0BI84o37y9mbsU4ME4CMRspo8i0PmGZ17M1G5BGJbZwoRt2wSajgkjtKBVTkbhMI%2FJC6%2BYYZIqhOSOzvGzNeUtsl7rIfgfjDvLb7o7eGT3kZWY4Dty69IbIeUWngE4wTJaENmNySlP3zspxFd6GpPa%2Fu6NDIDew6ketecr4QEsOtkXjZiZlLtPvsy7%2FA%2BHkZp9X%2BNwb8gPYCBIXAyRzo2afV%2FjcG%2FID2AgSFwMkc6M%3D\u0026source_type=1\u0026pack_time=1672813645.30\u0026use_pb=1\u0026openrtb_adx_id=229\u0026pc=3aR0arTEoIYYwsT%2Fvu0st2afV%2FjcG%2FID2AgSFwMkc6M%3D\u0026ttdsp_price=${AUCTION_PRICE}",
										"lurl":"https://api16-access-sg.pangle.io/api/ad/union/openrtb/loss/?req_id=8b82dc8c-1ba7-4b42-a5e8-79f21193ac64u6931\u0026ttdsp_adx_index=229\u0026rit=980239539\u0026extra=D1dHlaitCWcbuS1tBI1IbwgqChVLXaxcUeNP9IOZIaGGPS6HUO0qMQNwt%2B8G8%2F5%2FelHJMqcla6A7bFTlZShn4FMrExljLWl3aK2AKEkc1kKUOk2rwymGzveMHxezxZGWnc5gs6L0zPb0TtEVzDejGFpFm3PIdZKaMvpPc1Gxr7XYjfResgNkndoNmOyIQQFwY0BI84o37y9mbsU4ME4CMRspo8i0PmGZ17M1G5BGJbZwoRt2wSajgkjtKBVTkbhMI%2FJC6%2BYYZIqhOSOzvGzNeUtsl7rIfgfjDvLb7o7eGT3kZWY4Dty69IbIeUWngE4wTJaENmNySlP3zspxFd6GpPa%2Fu6NDIDew6ketecr4QEsOtkXjZiZlLtPvsy7%2FA%2BHkZp9X%2BNwb8gPYCBIXAyRzo2afV%2FjcG%2FID2AgSFwMkc6M%3D\u0026reason=${AUCTION_LOSS}\u0026auction_mwb=${AUCTION_PRICE}\u0026use_pb=1",
										"adm":"\u003clink rel=\"stylesheet\" href=\"https://sf-tb-sg.ibytedtos.com/obj/ad-pattern-sg/static/images/fallback/style.css\" /\u003e \u003cscript\u003e window._track_show_ = JSON.parse(\"[\\u0022https:\\/\\/s2s.adjust.com\\/impression\\/dlxzm9l?campaign=RMN_AND_AI_ZA_AutoBids%261733996132415538\\u0026adgroup=RMN_AND_AI_ZA_AutoBids_ModelDay%261753558369457154-Pangle\\u0026creative=RMN_ModelDay_916_11s_EN_Leti-camera-NewApproach-app_v01.mp4_auto%261753558363522082\\u0026s2s=1\\u0026idfa=\\u0026gps_adid=81481f8a-334f-4136-a5c9-1933d917d067\\u0026adgroup_id=1753558369457154\\u0026ip_address=41.116.232.11\\u0026user_agent=Mozilla%2F5.0\\u002b%28Linux%3B\\u002bAndroid\\u002b11%3B\\u002bNokia\\u002bC10\\u002bBuild%2FRP1A.201005.001%3B\\u002bwv%29\\u002bAppleWebKit%2F537.36\\u002b%28KHTML%2C\\u002blike\\u002bGecko%29\\u002bVersion%2F4.0\\u002bChrome%2F90.0.4430.210\\u002bMobile\\u002bSafari%2F537.36\\u0026campaign_id=1733996132415538\\u0026creative_id=1753558363522082\\u0026tiktok_placement=Pangle\\u0026external_tracker_ids=1\\u0026tiktok_callback_param=E.C.P.CrIBgaPNrBNY_VMZMaeV2boEVCJYLuCFA2gAvLbZaocazrMpYaqlC8R7fLIzdW2p4DrNKC6Tx3u4-7s_krQvpiM9VCHRYoEUKmsYIOpJ63ZWcp70N2nFxdtq-uEmmW9UdKhWng-2Ji0GlJUY2HIBThTj0KQ1ur7tMOeQQFcPKZEihclLwN17CVXXcNebYTnTuiZ_kjzHineFNKoFrLLZgC5d2jfF0ilbjfV0C4JI1uKZa6SUqRIEdjIuMBogqWxgAFKjyZpWG5mfYpb7TtbsVXWdNaNEd8frssdT2u8\\u0022]\" || \"[]\"); window._track_download_event_ = JSON.parse(\"[\\u0022https:\\/\\/api16-event-sg.pangle.io\\/api\\/ad\\/union\\/redirect\\/?req_id=8b82dc8c-1ba7-4b42-a5e8-79f21193ac64u6931\\u0026ttdsp_adx_index=229\\u0026use_pb=1\\u0026rit=980239539\\u0026call_back=jLuLQ0Ck9ylk%2FQdnmY%2BBeBWxFQZuFGTr9SLbNZYhA%2FTRT8h%2B3AqnhfixHvEJ9FgG27m0u%2FUsixHaHeWf32rDgYzOhEJT3PuuS1ha0z675Qdmn1f43BvyA9gIEhcDJHOj\\u0026extra=D1dHlaitCWcbuS1tBI1IbwgqChVLXaxcUeNP9IOZIaGGPS6HUO0qMQNwt%2B8G8%2F5%2FelHJMqcla6A7bFTlZShn4FMrExljLWl3aK2AKEkc1kKUOk2rwymGzveMHxezxZGWnc5gs6L0zPb0TtEVzDejGFpFm3PIdZKaMvpPc1Gxr7XYjfResgNkndoNmOyIQQFwY0BI84o37y9mbsU4ME4CMRspo8i0PmGZ17M1G5BGJbZwoRt2wSajgkjtKBVTkbhMI%2FJC6%2BYYZIqhOSOzvGzNeUtsl7rIfgfjDvLb7o7eGT3kZWY4Dty69IbIeUWngE4wTJaENmNySlP3zspxFd6GpPa%2Fu6NDIDew6ketecr4QEsOtkXjZiZlLtPvsy7%2FA%2BHkZp9X%2BNwb8gPYCBIXAyRzo2afV%2FjcG%2FID2AgSFwMkc6M%3D\\u0026source_type=1\\u0026pack_time=1672813645.30\\u0026active_extra=oAHOMAcDDa7qYU%2BiIZ2iNQurH4x2SYAfFsu2fPypxqjErEdHvqqf9TmmjrZSTpYTO%2FuUzxhnA1jdVKbH00ysgQ%3D%3D\\u0026is_dsp=1\\u0022]\" || \"[]\"); window._download_link_ = \"https:\\/\\/play.google.com\\/store\\/apps\\/details?id=com.bigwinepot.nwdn.international\"; window._image_url_ = \"https:\\/\\/p16-ttam-va.ibyteimg.com\\/origin\\/ad-union-i18n-api\\/ae8f34c8b3e1f57fb0819f58d686724a\"; window._deeplink_ = \"\"; window._event_domain_ = \"\"; window._creative_extra_ = \"\"; window._media_request_size_ = {}; window._ttclid_ = \"\"; window._pixel_urls_ = JSON.parse(\"[\\u0022\\u0022]\" || \"[]\"); \u003c/script\u003e \u003csection class=\"pangle-banner-wrapper\"\u003e \u003c/section\u003e \u003cscript src=\"https://sf-tb-sg.ibytedtos.com/obj/ad-pattern-sg/static/dsp/zepto.min.js\"\u003e \u003c/script\u003e \u003cscript src=\"https://sf-tb-sg.ibytedtos.com/obj/ad-pattern-sg/static/images/fallback/banner-1.0.4-5.js\"\u003e \u003c/script\u003e",
										"adid":"1753558363522082",
										"adomain":[
											"remini.ai"
										],
										"bundle":"com.bigwinepot.nwdn.international",
										"iurl":"https://p16-ttam-va.ibyteimg.com/origin/ad-union-i18n-api/ae8f34c8b3e1f57fb0819f58d686724a",
										"cid":"1753558363522082",
										"crid":"1753558363522082",
										"cat":[
											"IAB19-13"
										],
										"w":320,
										"h":50,
										"dsp_info":{
											"id":25,
											"name":"pangle",
											"group_name":"pangle",
											"endpoint":"https://api16-access-sg-nocdn.pangle.io/api/ad/union/openrtb/get_ads",
											"daily_budget":99999999999000000,
											"competable_qps_threshold":2000,
											"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBZHhJZCI6MjI5LCJleHAiOjI2MDc5OTcxMjYsImlzcyI6ImFkLmdvLXVuaW9uLmNvbW1vbi5pMThuZHNwLnRva2VuIn0.u55TgTZAjgmiw9Q4lAW8e6ELmHwVR5bjprVeZmqdGHY",
											"top_price":250000000,
											"bidfloor":10000,
											"auction_type":1,
											"tmax":400,
											"dsp_type":1
										},
										"ext":{
					
										}
									}
								],
								"seat":"pangle"
							}
						],
						"bidid":"8b82dc8c-1ba7-4b42-a5e8-79f21193ac64u6931",
						"cur":"USD",
						"vaststr":"",
						"ext":null,
						"expire_time":0
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//adsnowy
	r.HandleFunc("/mock/868/get_adsnowy_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"adid":"1234",
										"id":"bidid",
										"price":1,
										"nurl":"https:\/\/track.adsbea.com\/wn?p=${AUCTION_PRICE}&wp=MS4wMTIxMjE%3D&r=1.2&at=1&pw=320&ph=50&ad=bnVsbA%3D%3D&ig=0&wn=aHR0cDovL3Rlc3QuYWRzaS5jb20uY24vdGVzdF93aW4%2FcHJpY2U9MS4wMTIxMjE%3D&t=3&itl=0&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&dd=2&md=3&sy=1&rts=1676279103&pf=1&bd=c2hhcmUuc2hhcmVrYXJvLnBybw%3D%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D",
										"adm":"<body class=\"container\"style=\"width:100%; height:100%;overflow:hidden;\"><img src=\"https:\/\/track.adsbea.com\/imp?p=${AUCTION_PRICE}&loss=${AUCTION_LOSS}&wp=MS4wMTIxMjE%3D&rip=114.124.182.18&dvc=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&tid=&r=1.2&sk=&at=1&pw=320&ph=50&ad=bnVsbA%3D%3D&ig=0&mt=1&rwd=&t=3&itl=0&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&dd=2&md=3&sy=1&rts=1676279103&pf=1&bd=c2hhcmUuc2hhcmVrYXJvLnBybw%3D%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D\"style=\"display: none;\"\/><meta content=\"width=device-width,initial-scale=1,maximum-scale=1,minimum-scale=1,user-scalable=no\" name=\"viewport\"><meta http-equiv='Content-Type' content='textml charset=UTF-8' \/><style type='text\/css'>*{padding:0px;margin:0px;} a:link{text-decoration:none;}<\/style><a target='_blank' href='https:\/\/test.adsi.com.cn\/test'><img src='https:\/\/cdn.adsi.com.cn\/x\/img\/tbtest320x50.jpg' width='100%' height='100%' border='0' alt='' \/><\/a><script type=\"text\/javascript\">(function(window,document){var getAllA=document.getElementsByClassName('container')[0].childNodes;for(var i=getAllA.length-1;i>=0;i--){getAllA[i].addEventListener('touchstart',function(ev){var oImg=new Image();oImg.src=\"https:\/\/track.adsbea.com\/clk?p=${AUCTION_PRICE}&wp=MS4wMTIxMjE%3D&rip=114.124.182.18&dvc=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&tid=&r=1.2&sk=&at=1&pw=320&ph=50&ad=bnVsbA%3D%3D&ig=0&mt=1&rwd=&t=3&itl=0&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&dd=2&md=3&sy=1&rts=1676279103&pf=1&bd=c2hhcmUuc2hhcmVrYXJvLnBybw%3D%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D\"},false)}})(Function('return this')(),document);<\/script><\/body>"
									}
								]
							}
						]
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_adsnowy_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"adid":"adid",
										"nurl":"https:\/\/track.adsbea.com\/wn?p=${AUCTION_PRICE}&md=3&sy=3&rts=1676279685&pf=1&bd=c2hhcmUuc2hhcmVrYXJvLnBybw%3D%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D&wp=MS4wMTIxMjE%3D&itl=0&pw=720&r=1.2&ad=bnVsbA%3D%3D&ig=0&wn=aHR0cDovL3Rlc3QuYWRzaS5jb20uY24vdGVzdF93aW4%2FcHJpY2U9MS4wMTIxMjE%3D&ph=1280&at=1&t=4&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&dd=3",
										"adm":"<VAST version=\"3.0\"><Ad id=\"309CB8F2F1A86611\"><Wrapper><AdSystem version=\"1.0\">AdServer<\/AdSystem><VASTAdTagURI><![CDATA[http:\/\/ads48.adtelligent.com\/vast\/?adid=309CB8F2F1A86611&aid=654560&cmpId=547578&cb=1819771561]]><\/VASTAdTagURI><Impression><![CDATA[https:\/\/tracker-mrl.ortb.net\/imp?ie=CgExEiRjODhjYWMyNC0zNjdjLTY3OGUtMTBmYy1kNWI0NGI4NmExOWMyBTU2MTkwOP62A0IjUHJlcXVlbF9WQVNUX0lBXyQxLjhfMDcvMjgvMjAyMV9BZHRI6aHcBFIZVmxpb25fdmlkZW9faWFfMDcvMDgvMjAyMXlM27-y0qRUP4EBkst_SL99XT-IAQGQAejW04-2L5oBEGNvbS56dW1vYmkubXNuYmOiAR5odHRwOi8vYXBwcy5uYmNuZXdzLmNvbS9tb2JpbGWqAQdBbmRyb2lksgENSElHSEVORF9QSE9ORboBMW5iY19uZXdzOl9icmVha2luZ19uZXdzLF91c19uZXdzXyZhbXA7X2xpdmVfdmlkZW_IAcAC0AHgA9oBA1VTQfIBJDNlYTZlZWI5LTQ0NzAtNDA4MC1hNDMxLTNjMjk4N2ZhZjVjM5ICBVZJREVPmgIDQVBQogIDbXJssAL___________8BuAL8tgPCAiNQcmVxdWVsX1ZBU1RfSUFfJDEuOF8wNy8yOC8yMDIxX0FkdMoCCjU5Mjk5MDAyNjXyAgNBRE2BAzSAt0CCgkZAiQPzH9JvX09XwJIDCzY5LjU0LjQ2Ljk4mgOgAU1vemlsbGEvNS4wIChMaW51eDsgQW5kcm9pZCAxMTsgU00tRzk3M1UgQnVpbGQvUlAxQS4yMDA3MjAuMDEyOyB3dikgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi80LjAgQ2hyb21lLzg4LjAuNDMyNC4xODEgTW9iaWxlIFNhZmFyaS81MzcuMzaiAwpHYWxheHkgUzEwqgMHU2Ftc3VuZ7gDAcIDGVNFUlZFUl9UT19TRVJWRVJfT1BFTl9SVELKAwNUQUfSAwttaW5uZWFwb2xpc-oDCU5PTl9GUkFVRPIDBkJPVE1BTvoDAlYxugQFVklERU_CBAIxMcoED0FuZHJvaWQgd2Vidmlld9gErAHiBAdQcmVxdWVs6ASHAfIEBVZsaW9u&p=1.012121]]><\/Impression><Error><![CDATA[https:\/\/tracker-mrl.ortb.net\/tve?ve=CgExEiRjODhjYWMyNC0zNjdjLTY3OGUtMTBmYy1kNWI0NGI4NmExOWMyBTU2MTkwOP62A0IjUHJlcXVlbF9WQVNUX0lBXyQxLjhfMDcvMjgvMjAyMV9BZHRI6aHcBFIZVmxpb25fdmlkZW9faWFfMDcvMDgvMjAyMYgBD5AB6NbTj7YvmgEQY29tLnp1bW9iaS5tc25iY6IBHmh0dHA6Ly9hcHBzLm5iY25ld3MuY29tL21vYmlsZaoBB0FuZHJvaWSyAQ1ISUdIRU5EX1BIT05FugExbmJjX25ld3M6X2JyZWFraW5nX25ld3MsX3VzX25ld3NfJmFtcDtfbGl2ZV92aWRlb8gBwALQAeAD2gEDVVNB8gEkM2VhNmVlYjktNDQ3MC00MDgwLWE0MzEtM2MyOTg3ZmFmNWMzkgIFVklERU-aAgNBUFCiAgNtcmywAv___________wG4Avy2A8ICI1ByZXF1ZWxfVkFTVF9JQV8kMS44XzA3LzI4LzIwMjFfQWR0ygIKNTkyOTkwMDI2NYEDNIC3QIKCRkCJA_Mf0m9fT1fAkgMLNjkuNTQuNDYuOTiaA6ABTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDExOyBTTS1HOTczVSBCdWlsZC9SUDFBLjIwMDcyMC4wMTI7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvODguMC40MzI0LjE4MSBNb2JpbGUgU2FmYXJpLzUzNy4zNqIDCkdhbGF4eSBTMTCqAwdTYW1zdW5nuAMBwgMZU0VSVkVSX1RPX1NFUlZFUl9PUEVOX1JUQsoDA1RBR9IDC21pbm5lYXBvbGlz6gMJTk9OX0ZSQVVE8gMGQk9UTUFO-gMCVjG6BAVWSURFT8IEAjExygQPQW5kcm9pZCB3ZWJ2aWV32ASsAeIEB1ByZXF1ZWzoBIcB8gQFVmxpb24&vec=[ERRORCODE]&type=2]]><\/Error><Impression><![CDATA[http:\/\/ads48.adtelligent.com\/tre\/imp\/?adid=309CB8F2F1A86611&aid=654560&cmpId=547578]]><\/Impression><Impression><![CDATA[http:\/\/ads48.adtelligent.com\/t\/3d\/?adid=309CB8F2F1A86611&aid=654560&cmpId=547578&i=0&l=t&e=i]]><\/Impression><Creatives><Creative><Linear><TrackingEvents><Tracking event=\"creativeView\"><![CDATA[http:\/\/ads48.adtelligent.com\/t\/e\/?adid=309CB8F2F1A86611&code=52&aid=654560]]><\/Tracking><Tracking event=\"start\"><![CDATA[http:\/\/ads48.adtelligent.com\/t\/e\/?adid=309CB8F2F1A86611&code=53&aid=654560]]><\/Tracking><Tracking event=\"midpoint\"><![CDATA[http:\/\/ads48.adtelligent.com\/t\/e\/?adid=309CB8F2F1A86611&code=54&aid=654560]]><\/Tracking><Tracking event=\"firstQuartile\"><![CDATA[http:\/\/ads48.adtelligent.com\/t\/e\/?adid=309CB8F2F1A86611&code=55&aid=654560]]><\/Tracking><Tracking event=\"thirdQuartile\"><![CDATA[http:\/\/ads48.adtelligent.com\/t\/e\/?adid=309CB8F2F1A86611&code=56&aid=654560]]><\/Tracking><Tracking event=\"complete\"><![CDATA[http:\/\/ads48.adtelligent.com\/t\/e\/?adid=309CB8F2F1A86611&code=57&aid=654560]]><\/Tracking><Tracking event=\"skip\"><![CDATA[http:\/\/ads48.adtelligent.com\/t\/e\/?adid=309CB8F2F1A86611&code=66&aid=654560]]><\/Tracking><Tracking event=\"start\"><![CDATA[https:\/\/tracker-mrl.ortb.net\/tve?ve=CgExEiRjODhjYWMyNC0zNjdjLTY3OGUtMTBmYy1kNWI0NGI4NmExOWMyBTU2MTkwOP62A0IjUHJlcXVlbF9WQVNUX0lBXyQxLjhfMDcvMjgvMjAyMV9BZHRI6aHcBFIZVmxpb25fdmlkZW9faWFfMDcvMDgvMjAyMYgBD5AB6NbTj7YvmgEQY29tLnp1bW9iaS5tc25iY6IBHmh0dHA6Ly9hcHBzLm5iY25ld3MuY29tL21vYmlsZaoBB0FuZHJvaWSyAQ1ISUdIRU5EX1BIT05FugExbmJjX25ld3M6X2JyZWFraW5nX25ld3MsX3VzX25ld3NfJmFtcDtfbGl2ZV92aWRlb8gBwALQAeAD2gEDVVNB8gEkM2VhNmVlYjktNDQ3MC00MDgwLWE0MzEtM2MyOTg3ZmFmNWMzkgIFVklERU-aAgNBUFCiAgNtcmywAv___________wG4Avy2A8ICI1ByZXF1ZWxfVkFTVF9JQV8kMS44XzA3LzI4LzIwMjFfQWR0ygIKNTkyOTkwMDI2NYEDNIC3QIKCRkCJA_Mf0m9fT1fAkgMLNjkuNTQuNDYuOTiaA6ABTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDExOyBTTS1HOTczVSBCdWlsZC9SUDFBLjIwMDcyMC4wMTI7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvODguMC40MzI0LjE4MSBNb2JpbGUgU2FmYXJpLzUzNy4zNqIDCkdhbGF4eSBTMTCqAwdTYW1zdW5nuAMBwgMZU0VSVkVSX1RPX1NFUlZFUl9PUEVOX1JUQsoDA1RBR9IDC21pbm5lYXBvbGlz6gMJTk9OX0ZSQVVE8gMGQk9UTUFO-gMCVjG6BAVWSURFT8IEAjExygQPQW5kcm9pZCB3ZWJ2aWV32ASsAeIEB1ByZXF1ZWzoBIcB8gQFVmxpb24&type=10]]><\/Tracking><Tracking event=\"firstQuartile\"><![CDATA[https:\/\/tracker-mrl.ortb.net\/tve?ve=CgExEiRjODhjYWMyNC0zNjdjLTY3OGUtMTBmYy1kNWI0NGI4NmExOWMyBTU2MTkwOP62A0IjUHJlcXVlbF9WQVNUX0lBXyQxLjhfMDcvMjgvMjAyMV9BZHRI6aHcBFIZVmxpb25fdmlkZW9faWFfMDcvMDgvMjAyMYgBD5AB6NbTj7YvmgEQY29tLnp1bW9iaS5tc25iY6IBHmh0dHA6Ly9hcHBzLm5iY25ld3MuY29tL21vYmlsZaoBB0FuZHJvaWSyAQ1ISUdIRU5EX1BIT05FugExbmJjX25ld3M6X2JyZWFraW5nX25ld3MsX3VzX25ld3NfJmFtcDtfbGl2ZV92aWRlb8gBwALQAeAD2gEDVVNB8gEkM2VhNmVlYjktNDQ3MC00MDgwLWE0MzEtM2MyOTg3ZmFmNWMzkgIFVklERU-aAgNBUFCiAgNtcmywAv___________wG4Avy2A8ICI1ByZXF1ZWxfVkFTVF9JQV8kMS44XzA3LzI4LzIwMjFfQWR0ygIKNTkyOTkwMDI2NYEDNIC3QIKCRkCJA_Mf0m9fT1fAkgMLNjkuNTQuNDYuOTiaA6ABTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDExOyBTTS1HOTczVSBCdWlsZC9SUDFBLjIwMDcyMC4wMTI7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvODguMC40MzI0LjE4MSBNb2JpbGUgU2FmYXJpLzUzNy4zNqIDCkdhbGF4eSBTMTCqAwdTYW1zdW5nuAMBwgMZU0VSVkVSX1RPX1NFUlZFUl9PUEVOX1JUQsoDA1RBR9IDC21pbm5lYXBvbGlz6gMJTk9OX0ZSQVVE8gMGQk9UTUFO-gMCVjG6BAVWSURFT8IEAjExygQPQW5kcm9pZCB3ZWJ2aWV32ASsAeIEB1ByZXF1ZWzoBIcB8gQFVmxpb24&type=11]]><\/Tracking><Tracking event=\"midpoint\"><![CDATA[https:\/\/tracker-mrl.ortb.net\/tve?ve=CgExEiRjODhjYWMyNC0zNjdjLTY3OGUtMTBmYy1kNWI0NGI4NmExOWMyBTU2MTkwOP62A0IjUHJlcXVlbF9WQVNUX0lBXyQxLjhfMDcvMjgvMjAyMV9BZHRI6aHcBFIZVmxpb25fdmlkZW9faWFfMDcvMDgvMjAyMYgBD5AB6NbTj7YvmgEQY29tLnp1bW9iaS5tc25iY6IBHmh0dHA6Ly9hcHBzLm5iY25ld3MuY29tL21vYmlsZaoBB0FuZHJvaWSyAQ1ISUdIRU5EX1BIT05FugExbmJjX25ld3M6X2JyZWFraW5nX25ld3MsX3VzX25ld3NfJmFtcDtfbGl2ZV92aWRlb8gBwALQAeAD2gEDVVNB8gEkM2VhNmVlYjktNDQ3MC00MDgwLWE0MzEtM2MyOTg3ZmFmNWMzkgIFVklERU-aAgNBUFCiAgNtcmywAv___________wG4Avy2A8ICI1ByZXF1ZWxfVkFTVF9JQV8kMS44XzA3LzI4LzIwMjFfQWR0ygIKNTkyOTkwMDI2NYEDNIC3QIKCRkCJA_Mf0m9fT1fAkgMLNjkuNTQuNDYuOTiaA6ABTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDExOyBTTS1HOTczVSBCdWlsZC9SUDFBLjIwMDcyMC4wMTI7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvODguMC40MzI0LjE4MSBNb2JpbGUgU2FmYXJpLzUzNy4zNqIDCkdhbGF4eSBTMTCqAwdTYW1zdW5nuAMBwgMZU0VSVkVSX1RPX1NFUlZFUl9PUEVOX1JUQsoDA1RBR9IDC21pbm5lYXBvbGlz6gMJTk9OX0ZSQVVE8gMGQk9UTUFO-gMCVjG6BAVWSURFT8IEAjExygQPQW5kcm9pZCB3ZWJ2aWV32ASsAeIEB1ByZXF1ZWzoBIcB8gQFVmxpb24&type=12]]><\/Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https:\/\/tracker-mrl.ortb.net\/tve?ve=CgExEiRjODhjYWMyNC0zNjdjLTY3OGUtMTBmYy1kNWI0NGI4NmExOWMyBTU2MTkwOP62A0IjUHJlcXVlbF9WQVNUX0lBXyQxLjhfMDcvMjgvMjAyMV9BZHRI6aHcBFIZVmxpb25fdmlkZW9faWFfMDcvMDgvMjAyMYgBD5AB6NbTj7YvmgEQY29tLnp1bW9iaS5tc25iY6IBHmh0dHA6Ly9hcHBzLm5iY25ld3MuY29tL21vYmlsZaoBB0FuZHJvaWSyAQ1ISUdIRU5EX1BIT05FugExbmJjX25ld3M6X2JyZWFraW5nX25ld3MsX3VzX25ld3NfJmFtcDtfbGl2ZV92aWRlb8gBwALQAeAD2gEDVVNB8gEkM2VhNmVlYjktNDQ3MC00MDgwLWE0MzEtM2MyOTg3ZmFmNWMzkgIFVklERU-aAgNBUFCiAgNtcmywAv___________wG4Avy2A8ICI1ByZXF1ZWxfVkFTVF9JQV8kMS44XzA3LzI4LzIwMjFfQWR0ygIKNTkyOTkwMDI2NYEDNIC3QIKCRkCJA_Mf0m9fT1fAkgMLNjkuNTQuNDYuOTiaA6ABTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDExOyBTTS1HOTczVSBCdWlsZC9SUDFBLjIwMDcyMC4wMTI7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvODguMC40MzI0LjE4MSBNb2JpbGUgU2FmYXJpLzUzNy4zNqIDCkdhbGF4eSBTMTCqAwdTYW1zdW5nuAMBwgMZU0VSVkVSX1RPX1NFUlZFUl9PUEVOX1JUQsoDA1RBR9IDC21pbm5lYXBvbGlz6gMJTk9OX0ZSQVVE8gMGQk9UTUFO-gMCVjG6BAVWSURFT8IEAjExygQPQW5kcm9pZCB3ZWJ2aWV32ASsAeIEB1ByZXF1ZWzoBIcB8gQFVmxpb24&type=13]]><\/Tracking><Tracking event=\"complete\"><![CDATA[https:\/\/tracker-mrl.ortb.net\/tve?ve=CgExEiRjODhjYWMyNC0zNjdjLTY3OGUtMTBmYy1kNWI0NGI4NmExOWMyBTU2MTkwOP62A0IjUHJlcXVlbF9WQVNUX0lBXyQxLjhfMDcvMjgvMjAyMV9BZHRI6aHcBFIZVmxpb25fdmlkZW9faWFfMDcvMDgvMjAyMYgBD5AB6NbTj7YvmgEQY29tLnp1bW9iaS5tc25iY6IBHmh0dHA6Ly9hcHBzLm5iY25ld3MuY29tL21vYmlsZaoBB0FuZHJvaWSyAQ1ISUdIRU5EX1BIT05FugExbmJjX25ld3M6X2JyZWFraW5nX25ld3MsX3VzX25ld3NfJmFtcDtfbGl2ZV92aWRlb8gBwALQAeAD2gEDVVNB8gEkM2VhNmVlYjktNDQ3MC00MDgwLWE0MzEtM2MyOTg3ZmFmNWMzkgIFVklERU-aAgNBUFCiAgNtcmywAv___________wG4Avy2A8ICI1ByZXF1ZWxfVkFTVF9JQV8kMS44XzA3LzI4LzIwMjFfQWR0ygIKNTkyOTkwMDI2NYEDNIC3QIKCRkCJA_Mf0m9fT1fAkgMLNjkuNTQuNDYuOTiaA6ABTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDExOyBTTS1HOTczVSBCdWlsZC9SUDFBLjIwMDcyMC4wMTI7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvODguMC40MzI0LjE4MSBNb2JpbGUgU2FmYXJpLzUzNy4zNqIDCkdhbGF4eSBTMTCqAwdTYW1zdW5nuAMBwgMZU0VSVkVSX1RPX1NFUlZFUl9PUEVOX1JUQsoDA1RBR9IDC21pbm5lYXBvbGlz6gMJTk9OX0ZSQVVE8gMGQk9UTUFO-gMCVjG6BAVWSURFT8IEAjExygQPQW5kcm9pZCB3ZWJ2aWV32ASsAeIEB1ByZXF1ZWzoBIcB8gQFVmxpb24&type=14]]><\/Tracking><\/TrackingEvents><VideoClicks><ClickTracking><![CDATA[http:\/\/ads48.adtelligent.com\/t\/e\/?adid=309CB8F2F1A86611&code=71&aid=654560]]><\/ClickTracking><ClickTracking><![CDATA[https:\/\/track.adsbea.com\/clk?p=${AUCTION_PRICE}&t=4&sy=3&rts=1676279685&pf=1&bd=c2hhcmUuc2hhcmVrYXJvLnBybw%3D%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D&wp=MS4wMTIxMjE%3D&rip=114.124.182.18&dvc=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&tid=&itl=0&sk=&rwd=&pw=720&r=1.2&ad=bnVsbA%3D%3D&ig=0&ph=1280&dd=3&at=1&md=3&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&mt=1]]><\/ClickTracking><\/VideoClicks><\/Linear><\/Creative><\/Creatives><Impression><![CDATA[https:\/\/track.adsbea.com\/imp?p=${AUCTION_PRICE}&loss=${AUCTION_LOSS}&t=4&sy=3&rts=1676279685&pf=1&bd=c2hhcmUuc2hhcmVrYXJvLnBybw%3D%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D&wp=MS4wMTIxMjE%3D&rip=114.124.182.18&dvc=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&tid=&itl=0&sk=&rwd=&pw=720&r=1.2&ad=bnVsbA%3D%3D&ig=0&ph=1280&dd=3&at=1&md=3&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&mt=1]]><\/Impression><\/Wrapper><\/Ad><\/VAST>",
										"id":"bidid",
										"price":1
									}
								]
							}
						]
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_adsnowy_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"adid":"1235",
										"id":"bidid",
										"price":1,
										"nurl":"https:\/\/track.adsbea.com\/wn?p=${AUCTION_PRICE}&wp=MS4wMTIxMjE%3D&r=1.2&at=1&pw=1200&ph=627&ad=bnVsbA%3D%3D&ig=0&wn=aHR0cDovL3Rlc3QuYWRzaS5jb20uY24vdGVzdF93aW4%2FcHJpY2U9MS4wMTIxMjE%3D&t=5&itl=0&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&dd=4&md=3&sy=2&rts=1676279762&pf=1&bd=Y29tLm1vYnoudmQuaW4%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D",
										"adm":"{\"native\":{\"imptrackers\":[\"http:\\\/\\\/test.adsi.com.cn\\\/test_imp\",\"https:\\\/\\\/track.adsbea.com\\\/imp?p=${AUCTION_PRICE}&loss=${AUCTION_LOSS}&wp=MS4wMTIxMjE%3D&rip=114.124.182.18&dvc=a70afac7-c1c8-856e-b729-afc030b83a49&tid=&r=1.2&sk=&at=1&pw=1200&ph=627&ad=bnVsbA%3D%3D&ig=0&mt=1&rwd=&t=5&itl=0&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&dd=4&md=3&sy=2&rts=1676279762&pf=1&bd=Y29tLm1vYnoudmQuaW4%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D\"],\"assets\":[{\"id\":1,\"required\":1,\"img\":{\"h\":627,\"w\":1200,\"url\":\"https:\\\/\\\/digivizer.com\\\/wp-content\\\/uploads\\\/2020\\\/08\\\/FB-Image-Ads-Specs.png\"}}],\"link\":{\"clicktrackers\":[\"http:\\\/\\\/test.adsi.com.cn\\\/test_click\",\"https:\\\/\\\/track.adsbea.com\\\/clk?p=${AUCTION_PRICE}&wp=MS4wMTIxMjE%3D&rip=114.124.182.18&dvc=a70afac7-c1c8-856e-b729-afc030b83a49&tid=&r=1.2&sk=&at=1&pw=1200&ph=627&ad=bnVsbA%3D%3D&ig=0&mt=1&rwd=&t=5&itl=0&wt=1&rd=NmUwYmJjMmMtNjk3NC00NzE5LWE1ZmQtOTFiMjgwMzE3Yjgy&dd=4&md=3&sy=2&rts=1676279762&pf=1&bd=Y29tLm1vYnoudmQuaW4%3D&cou=Y2Fu&did=2&dp=MS4wMTIxMjE%3D\"],\"url\":\"http:\\\/\\\/www.baidu.com\"}}}"
									}
								]
							}
						]
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//xapads
	r.HandleFunc("/mock/868/get_xapdas_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
							"seatbid":[
								{
									"bid":[
										{
											"id":"xUxq97RUTnM_0",
											"impid":"1",
											"price":2,
											"adid":"5717180",
											"burl":"https://rtb-apac.rtbserve.io/bill?i=xUxq97RUTnM_0",
											"lurl":"https://rtb-apac.rtbserve.io/loss?i=xUxq97RUTnM_0&lr=${AUCTION_LOSS}",
											"adm":"{\"native\":{\"ver\":\"1.1\",\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://rtb-apac.rtbserve.io/thumbnail?i=xUxq97RUTnM_0\",\"w\":1200,\"h\":627}}],\"link\":{\"url\":\"https://rtb-apac.rtbserve.io/click?i=xUxq97RUTnM_0\"},\"imptrackers\":[\"https://rtb-apac.rtbserve.io/win?i=xUxq97RUTnM_0&f=imp\"]}}",
											"adomain":[
												"ostin.com"
											],
											"cid":"916450",
											"crid":"5717180"
										}
									],
									"seat":"a113067"
								}
							],
							"bidid":"xUxq97RUTnM",
							"cur":"USD"
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_xapads_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"id":"vzsQHWz3-8Q_0",
										"impid":"1",
										"price":2,
										"adid":"2474265",
										"burl":"https://rtb-apac.rtbserve.io/bill?i=vzsQHWz3-8Q_0",
										"lurl":"https://rtb-apac.rtbserve.io/loss?i=vzsQHWz3-8Q_0&lr=${AUCTION_LOSS}",
										"adm":"<?xml version=\"1.0\" encoding=\"UTF-8\"?><VAST version=\"2.0\"><Ad id=\"215891\"><Wrapper><AdSystem version=\"0.1\">AdKernel</AdSystem><VASTAdTagURI><![CDATA[https://rtb-apac.rtbserve.io/win?i=vzsQHWz3-8Q_0&f=nurlnw]]></VASTAdTagURI><Impression><![CDATA[https://rtb-apac.rtbserve.io/win?i=vzsQHWz3-8Q_0&f=imp]]></Impression><Error><![CDATA[https://rtb-apac.rtbserve.io/vast-error?i=vzsQHWz3-8Q_0&err=[ERRORCODE]]]></Error><Creatives><Creative><Linear><TrackingEvents><Tracking event=\"firstQuartile\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=first_quartile]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=midpoint]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=third_quartile]]></Tracking><Tracking event=\"complete\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=complete]]></Tracking><Tracking event=\"mute\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=mute]]></Tracking><Tracking event=\"unmute\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=unmute]]></Tracking><Tracking event=\"pause\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=pause]]></Tracking><Tracking event=\"resume\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=resume]]></Tracking><Tracking event=\"fullscreen\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=fullscreen]]></Tracking><Tracking event=\"close\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=close]]></Tracking><Tracking event=\"acceptInvitation\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=accept_invitation]]></Tracking></TrackingEvents></Linear></Creative><Creative><NonLinearAds><TrackingEvents><Tracking event=\"expand\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=overlay_expand]]></Tracking><Tracking event=\"collapse\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=overlay_collapse]]></Tracking><Tracking event=\"acceptInvitation\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=overlay_accept_invitation]]></Tracking><Tracking event=\"close\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=overlay_close]]></Tracking><Tracking event=\"complete\"><![CDATA[https://rtb-apac.rtbserve.io/vast-track?i=vzsQHWz3-8Q_0&ve=overlay_complete]]></Tracking></TrackingEvents></NonLinearAds></Creative></Creatives></Wrapper></Ad></VAST>",
										"bundle":"ru.dodopizza.app",
										"adomain":[
											"dodopizza"
										],
										"cid":"215891",
										"crid":"5718256",
										"protocol":3,
										"dur":15
									}
								],
								"seat":"a113067"
							}
						],
						"bidid":"vzsQHWz3-8Q",
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_xapads_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"id":"kCqFQhQi4uU_0",
										"impid":"1",
										"price":2,
										"adid":"5604356",
										"burl":"https://rtb-apac.rtbserve.io/bill?i=kCqFQhQi4uU_0",
										"lurl":"https://rtb-apac.rtbserve.io/loss?i=kCqFQhQi4uU_0&lr=${AUCTION_LOSS}",
										"adm":"<img src='https://rtb-apac.rtbserve.io/win?i=kCqFQhQi4uU_0&f=imp' alt=' ' style='display:none'><a href=\"https://rtb-apac.rtbserve.io/click?i=kCqFQhQi4uU_0\" target=\"_blank\"><img src=\"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg\" width=\"320\" height=\"50\" border=\"0\" ></a>",
										"bundle":"com.amazon.mShop.android.shopping",
										"adomain":[
											"amazon.com"
										],
										"iurl":"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg",
										"cid":"916450",
										"crid":"5604356",
										"h":50,
										"w":320
									}
								],
								"seat":"a113067"
							}
						],
						"bidid":"kCqFQhQi4uU",
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//request_optimize
	r.HandleFunc("/mock/868/get_optimize_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
							"seatbid":[
								{
									"bid":[
										{
											"id":"xUxq97RUTnM_0",
											"impid":"1",
											"price":2,
											"adid":"5717180",
											"burl":"https://rtb-apac.rtbserve.io/bill?i=xUxq97RUTnM_0",
											"lurl":"https://rtb-apac.rtbserve.io/loss?i=xUxq97RUTnM_0&lr=${AUCTION_LOSS}",
											"adm":"{\"native\":{\"ver\":\"1.1\",\"assets\":[{\"id\":1,\"img\":{\"url\":\"https://rtb-apac.rtbserve.io/thumbnail?i=xUxq97RUTnM_0\",\"w\":1200,\"h\":627}}],\"link\":{\"url\":\"https://rtb-apac.rtbserve.io/click?i=xUxq97RUTnM_0\"},\"imptrackers\":[\"https://rtb-apac.rtbserve.io/win?i=xUxq97RUTnM_0&f=imp\"]}}",
											"adomain":[
												"ostin.com"
											],
											"cid":"916450",
											"crid":"5717180",
											"cat":[
												"IAB5",
												"IAB26-1"
											],
											"bundle":"com.taptodate",
											"language":"en"
										}
									],
									"seat":"a113067"
								}
							],
							"bidid":"xUxq97RUTnM",
							"cur":"USD"
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//request_optimize
	r.HandleFunc("/mock/868/get_optimize_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
							"seatbid":[
								{
									"bid":[
										{
											"id":"kCqFQhQi4uU_0",
											"impid":"1",
											"price":2,
											"adid":"5604356",
											"burl":"https://rtb-apac.rtbserve.io/bill?i=kCqFQhQi4uU_0",
											"lurl":"https://rtb-apac.rtbserve.io/loss?i=kCqFQhQi4uU_0&lr=${AUCTION_LOSS}",
											"adm":"<img src='https://rtb-apac.rtbserve.io/win?i=kCqFQhQi4uU_0&f=imp' alt=' ' style='display:none'><a href=\"https://rtb-apac.rtbserve.io/click?i=kCqFQhQi4uU_0\" target=\"_blank\"><img src=\"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg\" width=\"320\" height=\"50\" border=\"0\" ></a>",
											"bundle":"com.taptodate",
											"adomain":[
												"amazon.com"
											],
											"cat":[
												"IAB5",
												"IAB26-1"
											],
											"language":"en",
											"iurl":"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg",
											"cid":"916450",
											"crid":"5604356",
											"h":50,
											"w":320
										}
									],
									"seat":"a113067"
								}
							],
							"bidid":"kCqFQhQi4uU",
							"cur":"USD"
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//terabox
	r.HandleFunc("/mock/868/get_terabox_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
							"cur":"USD",
							"bidid":"4ae4b7bbcef64b859e2d655c20d08aa8",
							"seatbid":[
								{
									"seat":"0EVZd3X5kzM",
									"bid":[
										{
											"id":"4ae4b7bbcef64b859e2d655c20d08aa8",
											"impid":"1",
											"price":0.0191,
											"adm":"{\"native\":{\"ver\":\"1.0\",\"assets\":[{\"img\":{\"w\":1200,\"h\":627,\"url\":\"https://ww0.svr-algorix.com/pic/bbba43a353016fa2dc3b6ed4278c5ec3.jpg\"},\"id\":1}],\"link\":{\"clicktrackers\":[\"https://biddingmax.apse.trk.rixserv.com/clk?crpv=3&info=5QTYzgjYwMDMjZWYtkjM3IWLlZTN40COjFzYtczYhZWYwcTY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJyV2Zh5WYt5Cdz92bi5WYlx2Yu02bj1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmITP5RXYmEjMwAjNx0DZpNnJ1MjMxYTN5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZCMx0DcmgTYhhDMkBjMjVTN2QmMllTN4IGN2YWZjJmY3IGNlFGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ae4b7bbcef64b859e2d655c20d08aa8\",\"https://apac.trk.svr-algorix.com/clk?crpv=3&info=QO0E2M4IGMzAzYmFWL5IzNi1SZ2UDOtgzYxMWL3MWYmFGM3EWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZicldWYuFWbuQ3cv9mYuFWZsNmLt92Y9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJy0Te0FmJ3EDMwYTM9QWazZSNzITM2UTO3YTM9QncmIzNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmITM4QzNx0DcmADNiJzYxcDNihDO4kzYzgjN3kDNiBDNkhTO2YjMyAzN9Enc&price=0.0238&s=160017&r=70226698d40b497683c9888b471c2b40\",\"https://d-sapse.svr-algorix.com/dsp/rtb/click?c=IuzIGWir_uT1O3h_qwDhzQumQw306zbRBiLtgysqneEBfjxj62lHNiKQaYs5QbzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5atPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3b52xB64OpDY7STLJTy9fQcM5hQ0jmxlXCclqLvpypSZyqw2I5toMjX0WI4IlzHIMfKITWwKh0kqlL2V9SsuwhLTMRzDLvtV6Vy8dPUJu5oTePDEbCSPN1hHJc9k10V3-FFbSi0r4eU8fOL7QBctRyIyJD8hgz7E8xkxTj9M9HctnvLKr-o_NBh6kDvi4IaxQHUzmjJEPFz22_zTII3o19s-knxpKItYSpiiXgN4EU5NwZe8A%3D%3D\"],\"url\":\"https://www.lazada.sg?clickid=8FB5__AD_IuzIGWir_uT1O3h_qwDhzQumQw306zbRBiLtgysqneEBfjxj62lHNiKQaYs5QbzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5atPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3b52xB64OpDY7STLJTy9fQcM5hQ0jmxlXCclqLvpypSZyqw2I5toMjX0WI4IlzHIMfKITWwKh0kqlL2V9SsuwhLTMRzDLvtV6Vy8dPUJu5oTePDEbCSPN1hHJc9k10V3-FFbSi0r4eU8fOL7QBctRyIyJD8hgz7E8xkxTj9M9HctnvLKr-o_NBh6kDvi4IaxQHUzmjJEPFz22_zTII3o19s-knxpKItYSpiiXgN4EU5NwZe8A%3D%3D_QO0E2M4IGMzAzYmFWL5IzNi1SZ2UDOtgzYxMWL3MWYmFGM3EWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZicldWYuFWbuQ3cv9mYuFWZsNmLt92Y9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJy0Te0FmJ3EDMwYTM9QWazZSNzITM2UTO3YTM9QncmIzNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmITM4QzNx0DcmADNiJzYxcDNihDO4kzYzgjN3kDNiBDNkhTO2YjMyAzN9Enc&advertising_id=a70afac7-c1c8-856e-b729-afc030b83a49&af_siteid=30012\"},\"imptrackers\":[\"https://biddingmax.apse.trk.rixserv.com/win?crpv=3&info=5QTYzgjYwMDMjZWYtkjM3IWLlZTN40COjFzYtczYhZWYwcTY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJyV2Zh5WYt5Cdz92bi5WYlx2Yu02bj1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmITP5RXYmEjMwAjNx0DZpNnJ1MjMxYTN5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZCMx0DcmgTYhhDMkBjMjVTN2QmMllTN4IGN2YWZjJmY3IGNlFGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ae4b7bbcef64b859e2d655c20d08aa8\",\"https://biddingmax.apse.trk.rixserv.com/imp?crpv=3&info=5QTYzgjYwMDMjZWYtkjM3IWLlZTN40COjFzYtczYhZWYwcTY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJyV2Zh5WYt5Cdz92bi5WYlx2Yu02bj1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmITP5RXYmEjMwAjNx0DZpNnJ1MjMxYTN5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZCMx0DcmgTYhhDMkBjMjVTN2QmMllTN4IGN2YWZjJmY3IGNlFGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ae4b7bbcef64b859e2d655c20d08aa8\",\"https://apac.trk.svr-algorix.com/win?crpv=3&info=QO0E2M4IGMzAzYmFWL5IzNi1SZ2UDOtgzYxMWL3MWYmFGM3EWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZicldWYuFWbuQ3cv9mYuFWZsNmLt92Y9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJy0Te0FmJ3EDMwYTM9QWazZSNzITM2UTO3YTM9QncmIzNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmITM4QzNx0DcmADNiJzYxcDNihDO4kzYzgjN3kDNiBDNkhTO2YjMyAzN9Enc&price=0.0238&s=160017&r=70226698d40b497683c9888b471c2b40\",\"https://apac.trk.svr-algorix.com/imp?crpv=3&info=QO0E2M4IGMzAzYmFWL5IzNi1SZ2UDOtgzYxMWL3MWYmFGM3EWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZicldWYuFWbuQ3cv9mYuFWZsNmLt92Y9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJy0Te0FmJ3EDMwYTM9QWazZSNzITM2UTO3YTM9QncmIzNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmITM4QzNx0DcmADNiJzYxcDNihDO4kzYzgjN3kDNiBDNkhTO2YjMyAzN9Enc&price=0.0238&s=160017&r=70226698d40b497683c9888b471c2b40\",\"https://d-sapse.svr-algorix.com/dsp/rtb/impress?c=IuzIGWir_uT1O3h_qwDhzQumQw306zbRBiLtgysqneEBfjxj62lHNiKQaYs5QbzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5atPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3b52xB64OpDY7STLJTy9fQcM5hQ0jmxlXCclqLvpypSZyqw2I5toMjX0WI4IlzHIMfKITWwKh0kqlL2V9SsuwhLTMRzDLvtV6Vy8dPUJu5oTePDEbCSPN1hHJc9k10V3-FFbSi0r4eU8fOL7QBctRyIyJD8hgz7E8xkxTj9M9HctnvLKr-o_NBh6kDvi4IaxQHUzmjJEPFz22_zTII3o19s-knxpKItYSpiiXgN4EU5NwZe8A%3D%3D\"]}}",
											"iurl":"https://ww0.svr-algorix.com/pic/bbba43a353016fa2dc3b6ed4278c5ec3.jpg",
											"cid":"227",
											"crid":"515_2155",
											"bundle":"com.lazada.android",
											"adomain":[
												"lazada.sg"
											],
											"w":1200,
											"h":627,
											"exp":2000,
											"ext":{
												"mediaType":"native"
											}
										}
									]
								}
							]
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_terabox_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
							"cur":"USD",
							"bidid":"4ede7a267533440fb490ca2944df08a5",
							"seatbid":[
								{
									"seat":"0EVZd3X5kzM",
									"bid":[
										{
											"id":"4ede7a267533440fb490ca2944df08a5",
											"impid":"1",
											"price":0.0191,
											"adm":"<VAST xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:noNamespaceSchemaLocation=\"vast.xsd\" version=\"2.0\"><Ad id=\"AlgoriX-78f1a8d28c91d5a038a555f62676434a\"><InLine><AdSystem>Algorix</AdSystem><AdTitle>BELI SEKARANG</AdTitle><Impression><![CDATA[https://d-sapse.svr-algorix.com/dsp/rtb/impress?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D]]></Impression><Impression><![CDATA[https://d-sapse.svr-algorix.com/alx/track?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D&event=impress]]></Impression><Error><![CDATA[https://d-sapse.svr-algorix.com/alx/track?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D&event=error&errcode=[ERRORCODE]]]></Error><Creatives><Creative><Linear><Duration>00:00:20</Duration><MediaFiles><MediaFile width=\"720\" height=\"1280\" type=\"video/mp4\" bitrate=\"440\" delivery=\"progressive\" scalable=\"true\" maintainAspectRatio=\"true\"><![CDATA[https://ww0.svr-algorix.com/pic/611991af3310ddc9eef32ccf13d363ed.mp4]]></MediaFile></MediaFiles><TrackingEvents><Tracking event=\"start\"><![CDATA[https://d-sapse.svr-algorix.com/alx/track?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D&event=start]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://d-sapse.svr-algorix.com/alx/track?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D&event=midpoint]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://d-sapse.svr-algorix.com/alx/track?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D&event=firstQuartile]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://d-sapse.svr-algorix.com/alx/track?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D&event=thirdQuartile]]></Tracking><Tracking event=\"complete\"><![CDATA[https://d-sapse.svr-algorix.com/alx/track?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D&event=complete]]></Tracking><Tracking event=\"start\"><![CDATA[https://trk.svr-algorix.com/vtrk?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a&vtype=start]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://trk.svr-algorix.com/vtrk?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a&vtype=firstQuartile]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://trk.svr-algorix.com/vtrk?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a&vtype=midpoint]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://trk.svr-algorix.com/vtrk?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a&vtype=thirdQuartile]]></Tracking><Tracking event=\"complete\"><![CDATA[https://trk.svr-algorix.com/vtrk?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a&vtype=complete]]></Tracking><Tracking event=\"start\"><![CDATA[https://biddingmax.apse.trk.rixserv.com/vtrk?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5&vtype=start]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://biddingmax.apse.trk.rixserv.com/vtrk?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5&vtype=firstQuartile]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://biddingmax.apse.trk.rixserv.com/vtrk?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5&vtype=midpoint]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://biddingmax.apse.trk.rixserv.com/vtrk?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5&vtype=thirdQuartile]]></Tracking><Tracking event=\"complete\"><![CDATA[https://biddingmax.apse.trk.rixserv.com/vtrk?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5&vtype=complete]]></Tracking></TrackingEvents><VideoClicks><ClickThrough><![CDATA[https://www.lazada.sg?clickid=8FB5__AD_IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D_AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&advertising_id=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d&af_siteid=30012]]></ClickThrough><ClickTracking><![CDATA[https://d-sapse.svr-algorix.com/alx/track?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D&event=click]]></ClickTracking><ClickTracking><![CDATA[https://d-sapse.svr-algorix.com/dsp/rtb/click?c=IuzIGWir_uT1O3h_qwDkzVinH1v06TzdVi3t0il-nbpVJW45tDhNM3XHOt5qFLzAu8HJHWBPHLSWJxx7nPUM96HZB_Bh4cdsqYz2YjsvL8L8z5eqPtZ3gnXE0qt4dcmGfAktrWbiC4WzfE_aKnSJKqOAH8wu5zK79h53Z2xAImBWDAzA1_oMopuCiwhVP4YPob16SG1VMpvtPJfEBuw2TOJmWFrPfoVgw8Zn4AXK0Cogf7koa3f52xB64OpDY7STLJTy9fQMNJQM1HunnH-MnKb9oWsRdja-ncc964KUmj57YQLSJZPPIzD8cR5rq06hDJqmo18fU5ZpBrD3Vacgo9XRJ-E_DvvYWuySb49yF4Eik1sPwOJJcTe9pYuJ67POqzZcvC3Dn8ToiwT1HMNlxjvjIoGI5EraN47jq4dz4ELzmdRfixPY0EjNBYcoyD-8XdxgoBoi6EXvkrIxRzNpzy4csAooOA%3D%3D]]></ClickTracking><ClickTracking><![CDATA[https://apac.trk.svr-algorix.com/clk?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a]]></ClickTracking><ClickTracking><![CDATA[https://biddingmax.apse.trk.rixserv.com/clk?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5]]></ClickTracking></VideoClicks></Linear></Creative></Creatives><Impression><![CDATA[https://apac.trk.svr-algorix.com/win?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a]]></Impression><Impression><![CDATA[https://apac.trk.svr-algorix.com/imp?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a]]></Impression><Error><![CDATA[https://trk.svr-algorix.com/vtrk?crpv=3&info=AZ3MWN1IWMkN2NygTLwITMh1iMjNGNtImY0gTLjNWNjR2NkFWPhZWamgTMuIDOx4CNyEjL0ETM9AXa1ZybyBnLvJXYrVmchh2cuUmchh2c9QWdiZCM9QHdyNnJx0Dd0lmYmETPtBnYmAzN9I3cyNnJx0TbmBnJwADMy0Dc4VmJw0TawFmJz0Te0FmJ3EDMwYTM9QWazZiM1QjN4gTO3YTM9QncmUjNyAjLw0TbhZSOzIDMuATPtZiN3IDMuATPtJmJwYTM9knYmckRB1zYmMTM4QzNx0DcmEWMzIjM3MzMmhzMkNmY2ImZmJGN4ADMzQTY2kTM1M2M9Enc&price=0.0238&s=160017&r=3c5196a430084bffb6bcd38f3372231a&vtype=error&errcode=[ERRORCODE]]]></Error><Impression><![CDATA[https://biddingmax.apse.trk.rixserv.com/win?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5]]></Impression><Impression><![CDATA[https://biddingmax.apse.trk.rixserv.com/imp?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5]]></Impression><Error><![CDATA[https://biddingmax.apse.trk.rixserv.com/vtrk?crpv=3&info=kdzY1UjYxQ2Y3IDOtAjMxEWLyM2Y00iYiRDOtM2Y1MGZ3QWY9EmZpZCOx4iM4EjL0ITMuQTMx0DcpVnJvJHcu8mchtWZyFGaz5SZyFGaz1DZ1JmJw0Dd0J3cmETP0RXaiZSM90GciZCM30jczJ3cmETPtZGcmADMwITPwhXZmATPpBXYmMTP5RXYmEjMwAjNx0DZpNnJyUDN2gDO5cjNx0DdyZCOzIDMuATPtFmJxkTMw4CM90mJ4MjMw4CM90mYmkTOz0TeiZyRGFUPjZSMx0DcmUTY4AjZkRDN5ITYjBTO0ImZwQDNzMTN3YjMhdTZkVGN9Enc&price=${AUCTION_PRICE}&s=160021&r=4ede7a267533440fb490ca2944df08a5&vtype=error&errcode=[ERRORCODE]]]></Error></InLine></Ad></VAST>",
											"cid":"227",
											"crid":"515_2142",
											"bundle":"com.lazada.android",
											"adomain":[
												"lazada.sg"
											],
											"w":720,
											"h":1280,
											"exp":2000,
											"ext":{
												"mediaType":"video"
											}
										}
									]
								}
							]
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_terabox_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
							"cur":"USD",
							"bidid":"b998f496c1124885bd721000ac46eab4",
							"seatbid":[
								{
									"seat":"0EVZd3X5kzM",
									"bid":[
										{
											"id":"b998f496c1124885bd721000ac46eab4",
											"impid":"1",
											"price":0.0089,
											"adm":"<script type=\"text/javascript\">!function(){var adc=function(str){return decodeURIComponent(escape(window.atob(str)))};document.write(adc(\"PHNjcmlwdCB0eXBlPSJ0ZXh0L2phdmFzY3JpcHQiIHNyYz0iaHR0cHM6Ly9hcHNlLWNyLnN2ci1hbGdvcml4LmNvbS9iP2NycHY9MyZpbmZvPVFqTXpNRE9pQnpNd01tWmgxU095Y2pZdFVtTjFnVEw0TVdNajF5TmpGbVpoQnpOaDFUWW1sbUoyRWpMNElqTHpJakx6QVRNOUFYYTFaeU15RXpieUJuTHZKWFlyVm1jaGgyY3VVbWNoaDJjOVFXZGlaQ005UUhkeU5uSngwRGQwbG1ZbUVUUHRCblltQXpOOUkzY3lObkp4MFRibUJuSndBRE94MERjNFZtSncwVGF3Rm1KeDBUZTBGbUozRURNd1lUTTlRV2F6WkNPd1lqTjRnVE8zWVRNOVFuY21ZVE14QWpMdzBUYmhaQ054RURNdUFUUHRaQ094RURNdUFUUHRKbUp3WVRNOWtuWW1ja1JCMXpZbUVUTTRRek54MERjbVlHTmhWbU5tWnpNd2dEWjBNRE1pSm1OM2tETjRnVFp6Z3pZeVlqTXdrak45RW5jJnByaWNlPTAuMDExMyZzPTE2MDAxNyZyPTY5MDI2MmM4M2U4ODQ5NzZiYjAzNGQ4MDM2ZjZlYTRmIj48L3NjcmlwdD4=\"));document.write(adc(\"PGltZyBzcmM9Imh0dHBzOi8vYmlkZGluZ21heC5hcHNlLnRyay5yaXhzZXJ2LmNvbS9pbXA/Y3Jwdj0zJmluZm89MEl6TXpnall3TURNalpXWXRrak0zSVdMbFpUTjQwQ09qRnpZdGN6WWhaV1l3Y1RZOUVtWnBaaU54NENPeTR5TXk0eU13RVRQd2xXZG1Nak14OG1jdzV5YnlGMmFsSlhZb05uTGxKWFlvTlhQa1ZuWW1BVFAwUm5jelpTTTlRSGRwSm1KeDBUYndKbUp3Y1RQeU5uY3paU005MG1ad1pDTXdnVE05QUhlbFpDTTlrR2NoWlNNOWtIZGhaU015QURNMkVUUGtsMmNtZ0RNMllETzRrek4yRVRQMEpuSnhFVE13NENNOTBXWW1rRE93QWpMdzBUYm1NVE14QWpMdzBUYmlaU081TVRQNUptSkhaVVE5TW1KNTBEY21RalloVm1OME1XWXdBRE14SXpOa0pXTjRnRE55RVRNalpUTzBZR081a2pZOUVuYyZwcmljZT0ke0FVQ1RJT05fUFJJQ0V9JnM9MTYwMDIxJnI9Yjk5OGY0OTZjMTEyNDg4NWJkNzIxMDAwYWM0NmVhYjQiIHdpZHRoPSIxIiBoZWlnaHQ9IjEiIHN0eWxlPSJkaXNwbGF5Om5vbmU7Ij48c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCI+IWZ1bmN0aW9uKCl7KG5ldyBJbWFnZSkuc3JjPSJodHRwczovL2JpZGRpbmdtYXguYXBzZS50cmsucml4c2Vydi5jb20vaW1wP2NycHY9MyZpbmZvPTBJek16Z2pZd01ETWpaV1l0a2pNM0lXTGxaVE40MENPakZ6WXRjelloWldZd2NUWTlFbVpwWmlOeDRDT3k0eU15NHlNd0VUUHdsV2RtTWpNeDhtY3c1eWJ5RjJhbEpYWW9ObkxsSlhZb05YUGtWblltQVRQMFJuY3paU005UUhkcEptSngwVGJ3Sm1Kd2NUUHlObmN6WlNNOTBtWndaQ013Z1RNOUFIZWxaQ005a0djaFpTTTlrSGRoWlNNeUFETTJFVFBrbDJjbWdETTJZRE80a3pOMkVUUDBKbkp4RVRNdzRDTTkwV1lta0RPd0FqTHcwVGJtTVRNeEFqTHcwVGJpWlNPNU1UUDVKbUpIWlVROU1tSjUwRGNtUWpZaFZtTjBNV1l3QURNeEl6TmtKV040Z0ROeUVUTWpaVE8wWUdPNWtqWTlFbmMmcHJpY2U9JHtBVUNUSU9OX1BSSUNFfSZzPTE2MDAyMSZyPWI5OThmNDk2YzExMjQ4ODViZDcyMTAwMGFjNDZlYWI0JmVycj05In0oKTs8L3NjcmlwdD4=\").replace(new RegExp(adc(\"XCR7QVVDVElPTl9QUklDRX0=\"),\"g\"),\"${AUCTION_PRICE}\"))}();</script><img src=\"https://biddingmax.apse.trk.rixserv.com/win?crpv=3&info=0IzMzgjYwMDMjZWYtkjM3IWLlZTN40COjFzYtczYhZWYwcTY9EmZpZiNx4COy4yMy4yMwETPwlWdmMjMx8mcw5ybyF2alJXYoNnLlJXYoNXPkVnYmATP0RnczZSM9QHdpJmJx0TbwJmJwcTPyNnczZSM90mZwZCMwgTM9AHelZCM9kGchZSM9kHdhZSMyADM2ETPkl2cmgDM2YDO4kzN2ETP0JnJxETMw4CM90WYmkDOwAjLw0TbmMTMxAjLw0TbiZSO5MTP5JmJHZUQ9MmJ50DcmQjYhVmN0MWYwADMxIzNkJWN4gDNyETMjZTO0YGO5kjY9Enc&price=${AUCTION_PRICE}&s=160021&r=b998f496c1124885bd721000ac46eab4\"width=\"1\"height=\"1\"style=\"display:none;\">",
											"iurl":"https://ww0.svr-algorix.com/pic/8f2f761386d6f5ed5f6fa698193cf8b5.jpg",
											"cid":"227",
											"crid":"515_2133",
											"bundle":"com.lazada.android",
											"adomain":[
												"lazada.sg"
											],
											"w":320,
											"h":50,
											"exp":1800,
											"ext":{
												"mediaType":"banner"
											}
										}
									]
								}
							]
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//pubmatic
	r.HandleFunc("/mock/868/get_pubmatic_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"id":"kCqFQhQi4uU_0",
										"impid":"1",
										"price":2,
										"adid":"5604356",
										"burl":"https://rtb-apac.rtbserve.io/bill?i=kCqFQhQi4uU_0",
										"lurl":"https://rtb-apac.rtbserve.io/loss?i=kCqFQhQi4uU_0&lr=${AUCTION_LOSS}",
										"adm":"<img src='https://rtb-apac.rtbserve.io/win?i=kCqFQhQi4uU_0&f=imp' alt=' ' style='display:none'><a href=\"https://rtb-apac.rtbserve.io/click?i=kCqFQhQi4uU_0\" target=\"_blank\"><img src=\"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg\" width=\"320\" height=\"50\" border=\"0\" ></a>",
										"bundle":"com.taptodate",
										"adomain":[
											"amazon.com"
										],
										"cat":[
											"IAB5",
											"IAB26-1"
										],
										"language":"en",
										"iurl":"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg",
										"cid":"916450",
										"crid":"5604356",
										"h":50,
										"w":320
									}
								],
								"seat":"a113067"
							}
						],
						"bidid":"kCqFQhQi4uU",
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_pubmatic_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"id":"kCqFQhQi4uU_0",
										"impid":"1",
										"price":2,
										"adid":"5604356",
										"burl":"https://rtb-apac.rtbserve.io/bill?i=kCqFQhQi4uU_0",
										"lurl":"https://rtb-apac.rtbserve.io/loss?i=kCqFQhQi4uU_0&lr=${AUCTION_LOSS}",
										"adm":"<img src='https://rtb-apac.rtbserve.io/win?i=kCqFQhQi4uU_0&f=imp' alt=' ' style='display:none'><a href=\"https://rtb-apac.rtbserve.io/click?i=kCqFQhQi4uU_0\" target=\"_blank\"><img src=\"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg\" width=\"320\" height=\"50\" border=\"0\" ></a>",
										"bundle":"com.taptodate",
										"adomain":[
											"amazon.com"
										],
										"cat":[
											"IAB5",
											"IAB26-1"
										],
										"language":"en",
										"iurl":"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg",
										"cid":"916450",
										"crid":"5604356",
										"h":50,
										"w":320
									}
								],
								"seat":"a113067"
							}
						],
						"bidid":"kCqFQhQi4uU",
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_pubmatic_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"id":"kCqFQhQi4uU_0",
										"impid":"1",
										"price":2,
										"adid":"5604356",
										"burl":"https://rtb-apac.rtbserve.io/bill?i=kCqFQhQi4uU_0",
										"lurl":"https://rtb-apac.rtbserve.io/loss?i=kCqFQhQi4uU_0&lr=${AUCTION_LOSS}",
										"adm":"<img src='https://rtb-apac.rtbserve.io/win?i=kCqFQhQi4uU_0&f=imp' alt=' ' style='display:none'><a href=\"https://rtb-apac.rtbserve.io/click?i=kCqFQhQi4uU_0\" target=\"_blank\"><img src=\"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg\" width=\"320\" height=\"50\" border=\"0\" ></a>",
										"bundle":"com.taptodate",
										"adomain":[
											"amazon.com"
										],
										"cat":[
											"IAB5",
											"IAB26-1"
										],
										"language":"en",
										"iurl":"https://static.rtbserve.io/n425/ad/320x50_DsjSp4Hr.jpg",
										"cid":"916450",
										"crid":"5604356",
										"h":50,
										"w":320
									}
								],
								"seat":"a113067"
							}
						],
						"bidid":"kCqFQhQi4uU",
						"cur":"USD"
					}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//yandex_filter_video
	r.HandleFunc("/mock/868/get_yandex_filter_video", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"1681714664067197644_5d640558-d3ff-4eac-83a9-75a466a4aee2_yeahmobi",
							"seatbid":[
								{
									"bid":[
										{
											"price":2.2719,
											"id":"1977054490448169440",
											"cid":"1_184819800",
											"impid":"1",
											"nurl":"https://an.yandex.ru/metacount/1SDnBhSZ0by200000000U9nJV9Kefl6eQAw0uGFEWprJN-HBMxvuQvF200IUC97G-I-syCUw9xCnf382nJCVS3PMWSHBGRnQ9Y3HoWWYPpAsRMO7wMHOo7YYZGtCMideVva9pDe8KltW893OooYsWPNXA5ZcAYD8-2uZWuG0mrmc47ifqmCWjPR90KYUPVeF14-PRXRlO7x59F2jRAdL0qv6XgrxzROs_MHb-Cl40hAScHb8zZA3h9dA339YBhCio6hc9WUL1aS1e8NCd4BB6Kuol_2ThdxajOlcLR2gks3oANE40WltJsO79twmCxjOL99JafDJKv8JTPDJqynrJH8DqdGjayrDTLCjKYtCBPBCJCrCZLUmj3sNsy32k8E5-G0BuqcMhCsoJYUSf_o-oJ9h9WXyj7_8ukYpxhUemKtxjomWdtvWUNNuve55T-wS4_Z5mXESaLYKcbbIbfcIcj_4AfsSdwkNdvoQcAUNN5HQdA6NcrXSabgabvwOcxg0bUKagrai5zXr5rWR9Z1zPGM5oYzMJaCB4mj3Imi3WmQs6mScH7ImDR3COFkyH6i1s-vD3zzqlVZqPptZIZQOCmwmAPzWOth4nkjgxcgLxspfNHWUaM7-4X9F31_iVy9P4zatv1G4Slc-iT_1pdY2NUO1-zOPYxzr7rs_YFq7-oQb1Hn0lw_WED4hjO7Z_3Bk7HoCjpWmukOf1mD3Wpq0oslBSW00?ssp-cur-price=${AUCTION_PRICE}&ssp-cur=USD",
											"crid":"1_184819800_480x320",
											"iurl":"http://bs-metadsp.yandex.ru/resource/spacer.gif",
											"adomain":[
												"play.google.com"
											],
											"w":"480",
											"adid":"yabs.NzIwNTc2MDc3MjgzNTg1MDc=",
											"h":"320",
											"cat":[
												"IAB19",
												"IAB8"
											],
											"adm":"<iframe frameborder=\"no\" scrolling=\"no\" marginwidth=\"0\" marginheight=\"0\" src=\"https://yandexadexchange.net/rtb_cache?rtb-cache-id=1977054490448169440&scr-height=320&h=320&scr-width=480&w=480&enabled-ad-type=text&ty=text&domain-hash=184819800&dh=184819800&yt=1&ssp-id=62419929&click_macro=\" style=\"height:320px;width:480px;\"></iframe>"
										}
									]
								}
							],
							"cur":"USD"
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//advilon
	r.HandleFunc("/mock/868/get_advilon_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"h":50,
										"price":0.261272,
										"adomain":[
											"doordash.com"
										],
										"cat":[
											"IAB8"
										],
										"iurl":"http:\/\/cdn.liftoff.io\/customers\/320\/previews\/eceec90d4a-ios-a547e6bc7b.png",
										"nurl":"http:\/\/adx-use.advlion.com\/wn?p=${AUCTION_PRICE}&t=5163&md=3130&sy=1&rts=1678781276&pf=1&bd=Y29tLnNwb3J0YnJhaW4uamV3ZWxwdXp6bGU%3D&cou=dXNh&did=142&dp=MC4zNjU3ODE%3D&itl=0&wn=&wt=1&dd=203&rd=MWZhYzE1ZGItODA3NC00ZmEwLTg1YjktYmE0YzNhODRmMjc3&ig=0&om=60&ad=ZG9vcmRhc2guY29t&ph=50&pw=320&r=1.4&at=1",
										"cid":"149639",
										"impid":"1",
										"w":320,
										"id":"1fac15db-8074-4fa0-85b9-ba4c3a84f277:haggler-advlion276.us-e.ec2.liftoff.io",
										"adm":"<body class=\"container\"style=\"width:100%; height:100%;overflow:hidden;\"><img src=\"http:\/\/track-use.advlion.com\/imp?p=${AUCTION_PRICE}&t=5163&wt=1&rd=MWZhYzE1ZGItODA3NC00ZmEwLTg1YjktYmE0YzNhODRmMjc3&r=1.4&dd=203&md=3130&sy=1&at=1&rts=1678781276&pf=1&bd=Y29tLnNwb3J0YnJhaW4uamV3ZWxwdXp6bGU=&sk=&rwd=&om=60&cou=dXNh&did=142&dp=MC4zNjU3ODE=&mt=1&wp=MC4zNjU3ODE=&rip=47.153.129.224&dvc=06f523e0-063a-44c2-88a4-bb59f2386a37&tid=&itl=0&bl=aHR0cDovL2ltcHJlc3Npb24tZWFzdC5saWZ0b2ZmLmlvL2Fkdmxpb24vaW1wcmVzc2lvbj9hZF9ncm91cF9pZD0xNDk2MzkmY2hhbm5lbF9pZD0xMjAmY3JlYXRpdmVfaWQ9MjI3NDU0JmRldmljZV9pZF9zaGExPTkyMThlMWJkNmNkZDQxYzc2ZWI5ZGVmYjE0NzViZDRlZDljZWUxNWQmc291cmNlX2FwcF9zdG9yZV9pZD1jb20uc3BvcnRicmFpbi5qZXdlbHB1enpsZSZiaWRfdHM9MTY3ODc4MTI3Njc0NCZhdWN0aW9uX2lkPUNUN0M3VEUxcUk2THN5NVRjZjExJmF1Y3Rpb25fcHJpY2U9MC4zNjU3ODEmb3JpZ2luPWhhZ2dsZXItYWR2bGlvbjI3Ng==&pw=320&ph=50&ad=ZG9vcmRhc2guY29t&ig=0\"style=\"display: none;\"\/><div>\n  <meta charset=\"utf-8\"\/>\n  <meta name=\"viewport\" content=\"width=device-width,initial-scale=1,user-scalable=no\"\/>\n  <a id=\"exchange-ping-url\" href=\"http:\/\/cdn.liftoff-creatives.io\/tracking_pixel.gif\" style=\"display:none\"><\/a>\n  <script>window.LIFTOFF_START_TS = Date.now()<\/script>\n  <script src=\"mraid.js\"><\/script>\n  <div id=\"tracking-tag-holder\" style=\"display:none\">\n    <img id=\"liftoff-beacon\" src=\"http:\/\/impression-east.liftoff.io\/advlion\/beacon?ad_group_id=149639&channel_id=120&creative_id=227454&auction_id=CT7C7TE1qI6Lsy5Tcf11&origin=haggler-advlion276\"\/>\n  <\/div>\n  <script src=\"http:\/\/cdn.liftoff-creatives.io\/resources\/polyfills-f238b7cf8f0fae0abed3.js\"><\/script>\n  <link href=\"http:\/\/cdn.liftoff-creatives.io\/resources\/liftoff-c58849acacc1df260b20.css\" rel=\"stylesheet\"\/>\n  <script src=\"http:\/\/cdn.liftoff-creatives.io\/resources\/liftoff-e1715d1076e351e38b80.js\"><\/script>\n  <script>\n    Liftoff.init(JSON.parse(decodeURIComponent('{\"creative_policy\":{\"show_chinese_watermark\":false,\"show_ad_choices\":false,\"use_custom_close\":false,\"audio_policy\":\"always-mute\",\"remove_ad_wrapper\":false,\"russian_age_rating\":null,\"close_delay\":0,\"add_border\":false,\"animation_duration_limit\":0,\"block_strobing\":false,\"block_notifications\":false},\"is_wrapper_vast_endscreen\":false,\"products\":%5B%5D,\"enable_html_return\":false,\"bid_context\":{\"dest_app_id\":4379,\"normalized_language\":\"en\",\"cdnURL\":\"http:\/\/cdn.liftoff-creatives.io\",\"creative_id\":227454,\"width\":320,\"source_app_name\":\"Jewel Sliding%E2%84%A2 Block Puzzle\",\"is_reengagement\":false,\"is_interactive\":false,\"source_app_id\":\"com.sportbrain.jewelpuzzle\",\"lopv_autoclick_behavior\":null,\"is_interstitial\":false,\"pv_autoclick_behavior\":null,\"viewclick_behavior\":null,\"exchange\":\"advlion\",\"country\":\"USA\",\"ab_tests\":%5B\"crserve-ios-pod-forced-control\",\"madlib-localize-generic-cta-control\",\"crserve-skan-pod-control\",\"crserve-deadzone-control\",\"liftoff-mraid-v2-experiment\"%5D,\"platform\":\"android\",\"height\":50},\"html_return_delay\":\"0\",\"click_config\":{\"clickthrough_url\":{\"url\":\"http:\/\/play.google.com\/store\/apps\/details%3Fid=com.doordash.driverapp%26referrer=adjust_external_click_id%253Dv.2_g.149639_a.1fac15db-8074-4fa0-85b9-ba4c3a84f277_c.120_t.ua_u.9218e1bd6cdd41c7%2526utm_campaign%253DTest%252BCampaign%2526utm_content%253DTest%252BSource%252BApp%2526utm_source%253DLiftoff%252B-%252BDasher%252BCampaign%2526utm_term%253DTest%252BCreative_320x50\",\"type\":2},\"click_tracking_urls\":%5B{\"url\":\"http:\/\/click.liftoff.io\/v1\/campaign_click\/RQHmMdzP3-o8yWZ8iKA8MvGbsHNj5ecT6gARHvZ062qizR-p15_ezWxpzh4DmJszwNmDLQaN8hwgkBNuppASZMDj4uRYgRygdru21e12ZpWxpP3l34DItUrHed7S7yobUKeBsOvup0jXb86LBXS2l-ZAwTwYrabb1ny1d5oEMlOiTCzm7wnedEVBSrSWTfA9Ab1EcP3v_Ki5UJliV1qd6fj82Efe8bPkT19OPfjBK0Vyi0GpZhLPcDGzRQu3tIfmUbDvymJHr9HdSF06aaAOxkJUHYaQJSDrB8z-hYM9TzepVCr-Y-pRrFc45RdDEHJdzmY2T0dWz9Va67FasWMvVd8vVy3IJUm0dQW4WQEXLotTEdSbHHT5cdGFjE-VZ6NNXRYQg20dYY0ddz_vaeSd91k23hsg2x-MU7SieWWxK-1pSueNur9W1ozi2F7SibyhaufjiyV5w-Ul4LIHbOCTD-ioqTRHcH6lWr1qx5mDNv51wY1PT27gFbVi0bTvJmZz_1J7iQDp-XdSL_tc1BhgrIYd3_gPnDNMFoG4AwVe9ihtPD2A-iRPgdQ0YJA6MrffnUpDpaLDYv3MnZHMyvMK991B5YWccyFiOGpF8HuppxwO6pktBuIX6xUTPbQDqhCxpW5eg7KgGEL0KF5kDN9uzg\",\"type\":1}%5D},\"app_store_data\":null,\"enable_logging\":false,\"logging_endpoint\":\"https:\/\/adexp.liftoff.io\",\"event_tracking_param\":\"aa5100CIeRCRIkMWZhYzE1ZGItODA3NC00ZmEwLTg1YjktYmE0YzNhODRmMjc3GOC-pfntMCB4KP7wDTCbIjoaY29tLnNwb3J0YnJhaW4uamV3ZWxwdXp6bGVCHmNyc2VydmUtaW9zLXBvZC1mb3JjZWQtY29udHJvbEIjbWFkbGliLWxvY2FsaXplLWdlbmVyaWMtY3RhLWNvbnRyb2xCGGNyc2VydmUtc2thbi1wb2QtY29udHJvbEIYY3JzZXJ2ZS1kZWFkem9uZS1jb250cm9sQhtsaWZ0b2ZmLW1yYWlkLXYyLWV4cGVyaW1lbnRKCjcxZmRkZDIyNGRQAloDVVNBYAJoBHIJdXMtZWFzdC0x4AEBgAF4kgECZW6YAQKhAQAAAAAAALA_qgEGMzIweDUwsgEGUHV6emxlugEdSmV3ZWwgU2xpZGluZ-KEoiBCbG9jayBQdXp6bGXCARlodG1sLWJiOWViMDY3OWRjZDZmMDM5NDg00gEFOTUwMzfaAQYzMjB4NTA\",\"app_icon\":\"http:\/\/cdn.liftoff-creatives.io\/customers\/320\/creatives\/4379-icon-250x250.png\"}')), {\n      image_originals: {\"638157\":{\"id\":638157,\"image_versions\":[{\"id\":2710436,\"version-id\":3,\"version-name\":\"lambda_png\",\"duration\":0.0,\"public-url\":\"http:\/\/cdn.liftoff-creatives.io\/customers\/5ffb825932\/image\/lambda_png\/971ea0b119df01b03517.png\",\"width\":434,\"height\":434,\"mime-type\":\"image\/png\",\"filesize\":2851}]}},\n      video_originals: {},\n    },\n      JSON.parse(decodeURIComponent('%5B%5D'))\n    );\n  <\/script>\n  <script src=\"http:\/\/cdn.liftoff-creatives.io\/resources\/madlib-82103ca2b8bdd456f9c7.js\"><\/script>\n  <link href=\"http:\/\/cdn.liftoff-creatives.io\/resources\/generic-cta-1b57f41737c4a9f3a2b3.css\" rel=\"stylesheet\"\/>\n  <link href=\"http:\/\/cdn.liftoff-creatives.io\/resources\/counter-5584fc29e9cd0e6da040.css\" rel=\"stylesheet\"\/>\n  <link href=\"http:\/\/cdn.liftoff-creatives.io\/resources\/hpbar-b7475c3bd31ecd0357f5.css\" rel=\"stylesheet\"\/>\n  <script src=\"http:\/\/cdn.liftoff-creatives.io\/resources\/generic-cta-f730e7121f98961b1c7f.js\"><\/script>\n  <script src=\"http:\/\/cdn.liftoff-creatives.io\/resources\/simple-image-838e36d4c473364aeb4a.js\"><\/script>\n  <script src=\"http:\/\/cdn.liftoff-creatives.io\/resources\/counter-0a0fad9887caf45cbcd1.js\"><\/script>\n  <script src=\"http:\/\/cdn.liftoff-creatives.io\/resources\/hpbar-68da0b528d775841c407.js\"><\/script>\n<script>\n  Madlib.load(Liftoff.snippetEl, JSON.parse(decodeURIComponent('{\"scenes\":%5B{\"id\":\"main\",\"layoutID\":43,\"transition\":null,\"layoutContent\":{\"css\":{\"CTA\":{\"font-size\":\"15px\"},\"Icon\":{\"overflow\":\"hidden\",\"border-radius\":\"20%25\"},\"Name\":{\"bottom\":\"5px\",\"font-size\":\"18px\",\"line-height\":\"20px\"},\"Caption\":{\"color\":\"rgb%28102, 102, 102%29\",\"font-size\":\"12px\",\"line-height\":\"14px\"}},\"containers\":%5B{\"id\":\"Background\",\"children\":%5B%5D},{\"id\":\"Icon\",\"children\":%5B%5D},{\"id\":\"CONTAINER - Text\",\"children\":%5B{\"id\":\"Name\",\"children\":%5B%5D},{\"id\":\"Caption\",\"children\":%5B%5D}%5D},{\"id\":\"CTA\",\"children\":%5B%5D}%5D,\"constraints\":{\"CTA\":{\"right\":\"4px\",\"height\":\"28px\",\"centerVertical\":true},\"Icon\":{\"top\":\"4px\",\"left\":\"4px\",\"bottom\":\"4px\",\"matchWidthToHeight\":true},\"Background\":{\"top\":0,\"left\":0,\"right\":0,\"bottom\":0},\"CONTAINER - Text\":{\"left\":{\"edge\":\"right\",\"anchor\":\"Icon\",\"distance\":\"5px\"},\"right\":{\"edge\":\"left\",\"anchor\":\"CTA\",\"distance\":\"5px\"},\"centerVertical\":true}},\"clickTargets\":%5B%5D,\"coordination\":%5B%5D,\"componentParams\":{\"CTA\":{\"params\":{\"blink\":{\"type\":\"boolean\",\"value\":false},\"uaCTA\":{\"type\":\"text\",\"value\":\"INSTALL\"},\"abrCTA\":{\"type\":\"text\",\"value\":\"OPEN\"}},\"component\":\"generic-cta\"},\"Icon\":{\"params\":{\"hpos\":{\"type\":\"text\",\"value\":\"center\"},\"size\":{\"type\":\"text\",\"value\":\"contain\"},\"vpos\":{\"type\":\"text\",\"value\":\"center\"},\"image\":{\"type\":\"image_original\",\"value\":638157}},\"component\":\"simple-image\"},\"Name\":{\"params\":{\"durationMS\":{\"type\":\"number\",\"value\":2000},\"countryCode\":{\"type\":\"text\",\"value\":null},\"backText\":{\"type\":\"text\",\"value\":null},\"useSeparators\":{\"type\":\"boolean\",\"value\":true},\"startValue\":{\"type\":\"number\",\"value\":null},\"currency\":{\"type\":\"string\",\"value\":\"USD\"},\"loops\":{\"type\":\"number\",\"value\":null},\"delayMS\":{\"type\":\"number\",\"value\":1500},\"type\":{\"type\":\"string\",\"value\":\"currency\"},\"endValue\":{\"type\":\"number\",\"value\":100},\"decimals\":{\"type\":\"number\",\"value\":2},\"frontText\":{\"type\":\"text\",\"value\":null}},\"component\":\"counter\"},\"Caption\":{\"params\":{\"backgroundColor\":{\"type\":\"string\",\"value\":\"%23bee0e9\"},\"durationMS\":{\"type\":\"number\",\"value\":2000},\"backText\":{\"type\":\"text\",\"value\":null},\"foregroundColors\":{\"type\":\"string\",\"value\":\"%23ae2004,%23FF3008\"},\"foregroundType\":{\"type\":\"string\",\"value\":\"gradient\"},\"height\":{\"type\":\"string\",\"value\":\"12px\"},\"image\":{\"type\":\"image_original\",\"value\":null},\"loops\":{\"type\":\"number\",\"value\":null},\"delayMS\":{\"type\":\"number\",\"value\":1500},\"direction\":{\"type\":\"string\",\"value\":\"grow\"},\"borderRadius\":{\"type\":\"number\",\"value\":25},\"wrapperStyle\":{\"type\":\"style\",\"value\":{\"transform\":\"translate%280%25,-40%25%29\"}},\"type\":{\"type\":\"string\",\"value\":\"none\"},\"frontText\":{\"type\":\"text\",\"value\":null},\"highlightColor\":{\"type\":\"string\",\"value\":null}},\"component\":\"hpbar\"}}}}%5D}')));\n<\/script>\n\n<\/div>\n<script type=\"text\/javascript\">(function(window,document){var getAllA=document.getElementsByClassName('container')[0].childNodes;for(var i=getAllA.length-1;i>=0;i--){getAllA[i].addEventListener('touchstart',function(ev){var oImg=new Image();oImg.src=\"http:\/\/track-use.advlion.com\/clk?p=${AUCTION_PRICE}&t=5163&wt=1&rd=MWZhYzE1ZGItODA3NC00ZmEwLTg1YjktYmE0YzNhODRmMjc3&dd=203&md=3130&sy=1&at=1&rts=1678781276&om=60&pf=1&bd=Y29tLnNwb3J0YnJhaW4uamV3ZWxwdXp6bGU=&cou=dXNh&did=142&tid=&itl=0&pw=320&ph=50&ad=ZG9vcmRhc2guY29t&sk=&rwd=&ig=0\"},false)}})(Function('return this')(),document);<\/script><script type=\"text\/javascript\" src=\"https:\/\/q.adrta.com\/s\/advl\/aa.js?cb=KhsXh2jqNtT2aJBF#advl;paid=advl;avid=142;caid=203;plid=227454;publisherId=3130;siteId=;priceBid=0.365781;pricePaid=${AUCTION_PRICE};lineitemID=;kv1=320x50;kv2=https%3A%2F%2Fplay.google.com%2Fstore%2Fapps%2Fdetails%3Fid%3Dcom.sportbrain.jewelpuzzle;kv3=06f523e0063a44c288a4bb59f2386a37;kv4=47.153.129.224;kv7=2098395;kv11=1fac15db-8074-4fa0-85b9-ba4c3a84f277;kv12=5163;kv14=MRAID-1;kv15=California;kv16=37.1357;kv17=-121.6502;kv18=com.sportbrain.jewelpuzzle;kv32=2098395;kv19=06f523e0-063a-44c2-88a4-bb59f2386a37;kv25=Jewel+Sliding%E2%84%A2+Block+Puzzle;kv26=android;kv27=Mozilla%2F5.0+%28Linux%3B+Android+13%3B+SM-N981U+Build%2FTP1A.220624.014%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F110.0.5481.153+Mobile+Safari%2F537.36;kv28=sm-n981u;kv44=;kv52=0;kv24=Mobile_InApp\" ><\/script><\/body>",
										"crid":"227454",
										"bundle":"com.doordash.driverapp"
									}
								]
							}
						]
					}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_advilon_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"cur":"USD",
						"bidid":"c98dacbdf0ab64779c13bcf8b1a174a2",
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"seat":"LAZADA",
								"bid":[
									{
										"price":0.08,
										"adomain":[
											"c.lazada.com.ph"
										],
										"lurl":"https:\/\/pbv2.quk.com\/v1\/callback?trace=c98dacbdf0ab64779c13bcf8b1a174a2&reqid=23143081-53d0-4785-a82e-871fac602d77&channel=vidmate_banner_native&country=PHL&adtype=native&os=android&bundle=free.tube.premium.advanced.tuber&adsize=1200x627&tagid=9881a530-affe-11ec-b852-f150f58a0c3a&slotid=4a82a8bb480d36362a3dc0b6bc42507a&device=248f0244-5702-4745-bafa-638073974744&ip=112.198.70.199&bidfloor=5&server=172.21.59.63&account=1&adgroup=648840820379750400&campaign=649644261746282496&creative=732999259124797440&linkid=24&convprice=1&deepprice=11.5&bidprice=8&skuid=3492974164&stuffid=5451&displaymanager=&instl=0&premium=0&act=loss&code=${AUCTION_LOSS}",
										"adid":"649644261746282496",
										"cid":"649644261746282496",
										"nurl":"http:\/\/adx-saas.advlion.com\/wn?p=${AUCTION_PRICE}&bd=ZnJlZS50dWJlLnByZW1pdW0uYWR2YW5jZWQudHViZXI%3D&cou=cGhs&did=123&dp=MC4wOA%3D%3D&itl=0&pw=1200&ph=627&ad=Yy5sYXphZGEuY29tLnBo&r=1&om=60&at=1&ig=0&t=5060&wn=aHR0cHM6Ly9wYnYyLnF1ay5jb20vdjEvY2FsbGJhY2s%2FdHJhY2U9Yzk4ZGFjYmRmMGFiNjQ3NzljMTNiY2Y4YjFhMTc0YTImcmVxaWQ9MjMxNDMwODEtNTNkMC00Nzg1LWE4MmUtODcxZmFjNjAyZDc3JmNoYW5uZWw9dmlkbWF0ZV9iYW5uZXJfbmF0aXZlJmNvdW50cnk9UEhMJmFkdHlwZT1uYXRpdmUmb3M9YW5kcm9pZCZidW5kbGU9ZnJlZS50dWJlLnByZW1pdW0uYWR2YW5jZWQudHViZXImYWRzaXplPTEyMDB4NjI3JnRhZ2lkPTk4ODFhNTMwLWFmZmUtMTFlYy1iODUyLWYxNTBmNThhMGMzYSZzbG90aWQ9NGE4MmE4YmI0ODBkMzYzNjJhM2RjMGI2YmM0MjUwN2EmZGV2aWNlPTI0OGYwMjQ0LTU3MDItNDc0NS1iYWZhLTYzODA3Mzk3NDc0NCZpcD0xMTIuMTk4LjcwLjE5OSZiaWRmbG9vcj01JnNlcnZlcj0xNzIuMjEuNTkuNjMmYWNjb3VudD0xJmFkZ3JvdXA9NjQ4ODQwODIwMzc5NzUwNDAwJmNhbXBhaWduPTY0OTY0NDI2MTc0NjI4MjQ5NiZjcmVhdGl2ZT03MzI5OTkyNTkxMjQ3OTc0NDAmbGlua2lkPTI0JmNvbnZwcmljZT0xJmRlZXBwcmljZT0xMS41JmJpZHByaWNlPTgmc2t1aWQ9MzQ5Mjk3NDE2NCZzdHVmZmlkPTU0NTEmZGlzcGxheW1hbmFnZXI9Jmluc3RsPTAmcHJlbWl1bT0wJmFjdD13aW4mdG9iaWRwcmljZT0wLjA4&wt=1&dd=164&rd=MjMxNDMwODEtNTNkMC00Nzg1LWE4MmUtODcxZmFjNjAyZDc3&md=3060&sy=2&rts=1678861585&pf=1",
										"id":"c3433d09-14de-44b0-bb80-0210586d9d00",
										"impid":"1",
										"adm":"{\"native\":{\"link\":{\"clicktrackers\":[\"https:\\\/\\\/pbv2.quk.com\\\/v1\\\/callback?trace=c98dacbdf0ab64779c13bcf8b1a174a2&reqid=23143081-53d0-4785-a82e-871fac602d77&channel=vidmate_banner_native&country=PHL&adtype=native&os=android&bundle=free.tube.premium.advanced.tuber&adsize=1200x627&tagid=9881a530-affe-11ec-b852-f150f58a0c3a&slotid=4a82a8bb480d36362a3dc0b6bc42507a&device=248f0244-5702-4745-bafa-638073974744&ip=112.198.70.199&bidfloor=5&server=172.21.59.63&account=1&adgroup=648840820379750400&campaign=649644261746282496&creative=732999259124797440&linkid=24&convprice=1&deepprice=11.5&bidprice=8&skuid=3492974164&stuffid=5451&displaymanager=&instl=0&premium=0&act=clk&clkid=a43b8ba9-5265-4c62-a9e9-c691a41ed659\",\"http:\\\/\\\/track.advlion.com\\\/clk?p=${AUCTION_PRICE}&t=5060&wt=1&rd=MjMxNDMwODEtNTNkMC00Nzg1LWE4MmUtODcxZmFjNjAyZDc3&dd=164&md=3060&sy=2&at=1&rts=1678861585&om=60&pf=1&bd=ZnJlZS50dWJlLnByZW1pdW0uYWR2YW5jZWQudHViZXI=&cou=cGhs&did=123&tid=&itl=0&pw=1200&ph=627&ad=Yy5sYXphZGEuY29tLnBo&sk=&rwd=&ig=0\"],\"url\":\"lazada:\\\/\\\/ph\\\/web\\\/www\\\/marketing\\\/gateway\\\/rta?null&dsource=sml&exlaz=e_CJReSJoO%2BYXGip8qo24MCZFmNCmP6gU%2FYmtqxGSrWkbCXvTKdNAtBT5gdgtVDHeQoEyKCPL4gwe4gVu3I2btRM9ngxN6tC4WD3WnASIas1Y%3D&rta_token=7ec8a8e0-26e4-4075-9818-41aab0b94a85&rta_event_id=21016c7d16788607745302581&os=android&gps_adid=248f0244-5702-4745-bafa-638073974744&android_id=&idfa=248f0244-5702-4745-bafa-638073974744&idfv=&bundle_id=free.tube.premium.advanced.tuber&device_model=&device_make=Realme&sub_id1=Y3JlYXRpdmU9NzMyOTk5MjU5MTI0Nzk3NDQwJmxpbmtpZD0yNCZjb252cHJpY2U9MSZkZWVwcHJpY2U9MTEuNSZyZXFpZD0yMzE0MzA4MS01M2QwLTQ3ODUtYTgyZS04NzFmYWM2MDJkNzcmYWR0eXBlPW5hdGl2ZSZidW5kbGU9ZnJlZS50dWJlLnByZW1pdW0uYWR2YW5jZWQudHViZXImYWRzaXplPTEyMDB4NjI3JnRhZ2lkPTk4ODFhNTMwLWFmZmUtMTFlYy1iODUyLWYxNTBmNThhMGMzYSZwcmVtaXVtPTAmaW5zdGw9MCZkaXNwbGF5bWFuYWdlcj0mc3R1ZmZpZD01NDUxJnNrdWlkPTM0OTI5NzQxNjQmY291bnRyeT1QSEwmb3M9YW5kcm9pZCZ0cmFjZT1jOThkYWNiZGYwYWI2NDc3OWMxM2JjZjhiMWExNzRhMiZiaWRwcmljZT04JmNoYW5uZWw9dmlkbWF0ZV9iYW5uZXJfbmF0aXZlJmlwPTExMi4xOTguNzAuMTk5JnNsb3RpZD00YTgyYThiYjQ4MGQzNjM2MmEzZGMwYjZiYzQyNTA3YSZkZXZpY2U9MjQ4ZjAyNDQtNTcwMi00NzQ1LWJhZmEtNjM4MDczOTc0NzQ0JmJpZGZsb29yPTUmc2VydmVyPTE3Mi4yMS41OS42MyZhY2NvdW50PTEmYWRncm91cD02NDg4NDA4MjAzNzk3NTA0MDAmY2FtcGFpZ249NjQ5NjQ0MjYxNzQ2MjgyNDk2&sub_id2=&sub_id3=&sub_id4=1107722&sub_id5=&lzdcid=&trigger_item=3492974164\",\"fallback\":\"https:\\\/\\\/c.lazada.com.ph\\\/t\\\/c.YTLOkp?rta_token=7ec8a8e0-26e4-4075-9818-41aab0b94a85&rta_event_id=21016c7d16788607745302581&os=android&gps_adid=248f0244-5702-4745-bafa-638073974744&imei=&android_id=&idfa=248f0244-5702-4745-bafa-638073974744&idfv=&bundle_id=free.tube.premium.advanced.tuber&device_model=&device_make=Realme&sub_id1=Y3JlYXRpdmU9NzMyOTk5MjU5MTI0Nzk3NDQwJmxpbmtpZD0yNCZjb252cHJpY2U9MSZkZWVwcHJpY2U9MTEuNSZyZXFpZD0yMzE0MzA4MS01M2QwLTQ3ODUtYTgyZS04NzFmYWM2MDJkNzcmYWR0eXBlPW5hdGl2ZSZidW5kbGU9ZnJlZS50dWJlLnByZW1pdW0uYWR2YW5jZWQudHViZXImYWRzaXplPTEyMDB4NjI3JnRhZ2lkPTk4ODFhNTMwLWFmZmUtMTFlYy1iODUyLWYxNTBmNThhMGMzYSZwcmVtaXVtPTAmaW5zdGw9MCZkaXNwbGF5bWFuYWdlcj0mc3R1ZmZpZD01NDUxJnNrdWlkPTM0OTI5NzQxNjQmY291bnRyeT1QSEwmb3M9YW5kcm9pZCZ0cmFjZT1jOThkYWNiZGYwYWI2NDc3OWMxM2JjZjhiMWExNzRhMiZiaWRwcmljZT04JmNoYW5uZWw9dmlkbWF0ZV9iYW5uZXJfbmF0aXZlJmlwPTExMi4xOTguNzAuMTk5JnNsb3RpZD00YTgyYThiYjQ4MGQzNjM2MmEzZGMwYjZiYzQyNTA3YSZkZXZpY2U9MjQ4ZjAyNDQtNTcwMi00NzQ1LWJhZmEtNjM4MDczOTc0NzQ0JmJpZGZsb29yPTUmc2VydmVyPTE3Mi4yMS41OS42MyZhY2NvdW50PTEmYWRncm91cD02NDg4NDA4MjAzNzk3NTA0MDAmY2FtcGFpZ249NjQ5NjQ0MjYxNzQ2MjgyNDk2&sub_id2=&sub_id3=&sub_id4=1107722&sub_id5=&lzdcid=&trigger_item=3492974164\"},\"ver\":\"1.2\",\"assets\":[{\"title\":{\"len\":13,\"text\":\"Lowest Prices\"},\"id\":1,\"required\":0},{\"img\":{\"type\":3,\"url\":\"https:\\\/\\\/cdn.getads.cn\\\/dsp\\\/vmdsp\\\/image\\\/1678191775460.a654ImhyDv-LiePl7arXt.jpg\",\"w\":1200,\"h\":627},\"id\":5,\"required\":1},{\"id\":3,\"data\":{\"len\":8,\"value\":\"SHOP NOW\"},\"required\":1},{\"id\":2,\"data\":{\"len\":38,\"value\":\"Free Shipping Nationwide, P0 min spend\"},\"required\":0}],\"imptrackers\":[\"https:\\\/\\\/c.lazada.com.ph\\\/i\\\/c.YTLOkp?rta_token=7ec8a8e0-26e4-4075-9818-41aab0b94a85&rta_event_id=21016c7d16788607745302581&os=android&gps_adid=248f0244-5702-4745-bafa-638073974744&imei=&android_id=&idfa=248f0244-5702-4745-bafa-638073974744&idfv=&bundle_id=free.tube.premium.advanced.tuber&device_model=&device_make=Realme&sub_id1=Y3JlYXRpdmU9NzMyOTk5MjU5MTI0Nzk3NDQwJmxpbmtpZD0yNCZjb252cHJpY2U9MSZkZWVwcHJpY2U9MTEuNSZyZXFpZD0yMzE0MzA4MS01M2QwLTQ3ODUtYTgyZS04NzFmYWM2MDJkNzcmYWR0eXBlPW5hdGl2ZSZidW5kbGU9ZnJlZS50dWJlLnByZW1pdW0uYWR2YW5jZWQudHViZXImYWRzaXplPTEyMDB4NjI3JnRhZ2lkPTk4ODFhNTMwLWFmZmUtMTFlYy1iODUyLWYxNTBmNThhMGMzYSZwcmVtaXVtPTAmaW5zdGw9MCZkaXNwbGF5bWFuYWdlcj0mc3R1ZmZpZD01NDUxJnNrdWlkPTM0OTI5NzQxNjQmY291bnRyeT1QSEwmb3M9YW5kcm9pZCZ0cmFjZT1jOThkYWNiZGYwYWI2NDc3OWMxM2JjZjhiMWExNzRhMiZiaWRwcmljZT04JmNoYW5uZWw9dmlkbWF0ZV9iYW5uZXJfbmF0aXZlJmlwPTExMi4xOTguNzAuMTk5JnNsb3RpZD00YTgyYThiYjQ4MGQzNjM2MmEzZGMwYjZiYzQyNTA3YSZkZXZpY2U9MjQ4ZjAyNDQtNTcwMi00NzQ1LWJhZmEtNjM4MDczOTc0NzQ0JmJpZGZsb29yPTUmc2VydmVyPTE3Mi4yMS41OS42MyZhY2NvdW50PTEmYWRncm91cD02NDg4NDA4MjAzNzk3NTA0MDAmY2FtcGFpZ249NjQ5NjQ0MjYxNzQ2MjgyNDk2&sub_id2=&sub_id3=&sub_id4=1107722&sub_id5=&lzdcid=&trigger_item=3492974164\",\"https:\\\/\\\/pbv2.quk.com\\\/v1\\\/callback?trace=c98dacbdf0ab64779c13bcf8b1a174a2&reqid=23143081-53d0-4785-a82e-871fac602d77&channel=vidmate_banner_native&country=PHL&adtype=native&os=android&bundle=free.tube.premium.advanced.tuber&adsize=1200x627&tagid=9881a530-affe-11ec-b852-f150f58a0c3a&slotid=4a82a8bb480d36362a3dc0b6bc42507a&device=248f0244-5702-4745-bafa-638073974744&ip=112.198.70.199&bidfloor=5&server=172.21.59.63&account=1&adgroup=648840820379750400&campaign=649644261746282496&creative=732999259124797440&linkid=24&convprice=1&deepprice=11.5&bidprice=8&skuid=3492974164&stuffid=5451&displaymanager=&instl=0&premium=0&act=imp&tobidprice=0.08&curc=1.00\",\"http:\\\/\\\/track.advlion.com\\\/imp?p=${AUCTION_PRICE}&t=5060&wt=1&rd=MjMxNDMwODEtNTNkMC00Nzg1LWE4MmUtODcxZmFjNjAyZDc3&r=1&dd=164&md=3060&sy=2&at=1&rts=1678861585&pf=1&bd=ZnJlZS50dWJlLnByZW1pdW0uYWR2YW5jZWQudHViZXI=&sk=&rwd=&om=60&cou=cGhs&did=123&dp=MC4wOA==&mt=1&wp=MC4wOA==&rip=112.198.70.199&dvc=248f0244-5702-4745-bafa-638073974744&tid=&itl=0&pw=1200&ph=627&ad=Yy5sYXphZGEuY29tLnBo&ig=0\"]}}",
										"crid":"732999259124797440",
										"bundle":"com.lazada.android"
									}
								]
							}
						]
					}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_advilon_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
			"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
			"bidid":"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364",
			"cur":"USD",
			"seatbid":[
				{
					"seat":"adx",
					"group":0,
					"bid":[
						{
							"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366",
							"impid":"1",
							"price":0.09000000000000001,
							"adm":"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<VAST version=\"3.0\"><Ad id=\"6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101365\"><InLine><AdSystem version=\"1.0\">flat-ads</AdSystem><AdTitle><![CDATA[Moto Bike Attack Race]]></AdTitle><Impression><![CDATA[https://api.flat-ads.com/api/tracker/tracking/impression?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=imp.23951.676725849872.526032182772.1]]></Impression><Creatives><Creative adID=\"111594\" id=\"111594\"><Linear><Duration>00:00:26</Duration><MediaFiles><MediaFile delivery=\"progressive\" type=\"video/mp4\" width=\"720\" height=\"1280\"><![CDATA[http://dsp-adcreative.mobshark.net/adshark_dsp/1666339183940_zd.mp4]]></MediaFile></MediaFiles><TrackingEvents><Tracking event=\"start\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=start]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=firstQuartile]]></Tracking><Tracking event=\"midpoint\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=midpoint]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=thirdQuartile]]></Tracking><Tracking event=\"complete\"><![CDATA[http://api.flat-ads.com/api/tracker/tracking/video?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=video.69563.292054236776.142351569576.1&event=complete]]></Tracking></TrackingEvents><VideoClicks><ClickThrough><![CDATA[https://game.86753.com/vid/1013/MotoBikeAttackRace/uid_12307201/etr_vid_ad/index.html?browserType=2&entry=sdk_ads]]></ClickThrough><ClickTracking><![CDATA[https://api.flat-ads.com/api/tracker/tracking/click?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=click.88712.271095740792.121392083592.1&af_siteid=166_155354_2292]]></ClickTracking></VideoClicks></Linear><CreativeExtensions></CreativeExtensions></Creative></Creatives><Description><![CDATA[Choose Your Favourite Moto and Weapon!]]></Description><Extensions><Extension><ImageWidth>320</ImageWidth><ImageHeight>480</ImageHeight><ImageUrl>http://dsp-adcreative.mobshark.net/adshark_dsp/1660294465587.jpg?x-oss-process=style/hq</ImageUrl></Extension></Extensions></InLine></Ad></VAST>",
							"nurl":"https://api.flat-ads.com/api/tracker/tracking/win_notice?log_source=win&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101364&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297051.22.114101366&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=111594&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155354&dspid=100&srcid=3815373&cou=IDN&gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1&ag=6&ac=281&clickid=v1_2292_12606_111594_155354_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_281&et=imp.23951.676725849872.526032182772.1",
							"crid":"111594",
							"ext":null,
							"cid":"2292",
							"language":"",
							"attr":null,
							"cat":null
						}
					]
				}
			]
		}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//younger
	r.HandleFunc("/mock/868/get_younger_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"7898359u9yuu4645r4",
						"seatbid":[
							{
								"bid":[
									{
										"id":"4039773",
										"impid":"1",
										"price":0.022222222222222223,
										"adid":"4039773",
										"nurl":"https://win.lenovoadx.weshare.live/win?mid=4016\u0026c=eyJkYXRlIjoiMjAyMzA1MDkiLCJob3VyIjoiMDYiLCJwdmVyIjoidjIuMC42Iiwic3VwcGx5X2lkIjo0MDE2LCJuYnIiOjAsImNvdW50cnkiOiJJRCIsInBsYXRmb3JtIjoyLCJ0cmFmZmljX3R5cGUiOiJhcHAiLCJpcCI6IjExNC4xMjQuMTgyLjE1IiwidWEiOiIiLCJnYWlkIjoiYWQ3ZGM1Y2MtODRiYi00Y2MyLWExMjAtODI3Y2QxYjU1YzdkMSIsImFkX2Zvcm1hdCI6OTAwMiwiYWRfc2l6ZSI6IjEyMDAqNjI3IiwiZG9tYWluIjoiIiwiYXBwX2J1bmRsZSI6InNoYXJlLnNoYXJla2Fyby5wcm8xIiwicHVibGlzaGVyX2lkIjoiIiwicGxhY2VtZW50IjoiMjUyNCIsInJlcXVlc3RfaWQiOiI3ODk4MzU5dTl5dXU0NjQ1cjQiLCJiaWRmbG9vciI6MC4wMTExMTExMTExMTExMTExMTIsImJpZGZsb29yY3VyIjoiVVNEIiwiYmlkX3ByaWNlIjowLjAyMjIyMjIyMjIyMjIyMjIyMywid2luX3ByaWNlIjowLCJyZXF1ZXN0X2JvZHkiOiIiLCJkY19sb2NhdGlvbiI6Imxlbm9hZHhzZyIsIm5hc3NldCI6ImltZzMiLCJpbnN0bCI6MCwiZGV2aWNldHlwZSI6NCwibWFrZSI6IkhVQVdFSSIsIm1vZGVsIjoiTUhBLUwyOSIsIm9zIjoiQW5kcm9pZCIsIndpbm5lciI6ODAwLCJkZW1hbmQiOlt7ImRlbWFuZF9pZCI6ODAwLCJiaWRfaWQiOiI0MDM5NzczIiwiY2FtcGFpZ25faWQiOiIyMzQ4NSIsImFkX2lkIjoiNDAzOTc3MyIsImNyZWF0aXZlX2lkIjoiNDAzOTc3MzEzIiwiZG9tYWluIjoiIiwiYXBwX2J1bmRsZSI6InJvb2NrLnRlc3QiLCJuYnIiOjAsImJpZF9wcmljZSI6MS44NDQwMTAxMTExMTExMTExLCJ3aW5fcHJpY2UiOjEuODQ0MDEwMTExMTExMTExMSwid2luX25vdGlmaWVyIjo5MDAyLCJlbmRwb2ludF9pZCI6NTksInRtdCI6ImRlYnVnIiwicmVxX3RpbWUiOjAsIm1hZ2ljX2FwcF9idW5kbGUiOiIiLCJjb252X3R5cGUiOjgwMDEsImF0dHJpYnV0aW9uX3R5cGUiOjgwMDEsImNwYyI6MCwicmVxdWVzdCI6MCwiYmlkIjowLCJ3aW4iOjAsInRpbWVvdXQiOjAsImltcHJlc3Npb24iOjAsImNsaWNrIjowLCJjb3N0IjowLCJjb252ZXJzaW9uIjowLCJjb252ZXJzaW9uX3JldmVudWUiOjAsIml1cmwiOiIiLCJjdXJsIjoiIiwibnVybCI6Imh0dHBzOi8vcmVxLmxlbm92b2FkeC53ZXNoYXJlLmxpdmUvd2luIiwibHVybCI6IiJ9XSwicmVxdWVzdCI6MCwiYmlkIjowLCJ3aW4iOjAsInRpbWVvdXQiOjAsImltcHJlc3Npb24iOjAsImNsaWNrIjowLCJjb3N0IjowLCJjb252ZXJzaW9uIjowLCJjb252ZXJzaW9uX3JldmVudWUiOjB9\u0026time=20230509061648\u0026price=${AUCTION_PRICE}\u0026bidid=${AUCTION_BID_ID}\u0026ext=800-9002",
										"lurl":"https://req.lenovoadx.weshare.live/lwin",
										"adm":"{\"native\":{\"ver\":\"1.0\",\"assets\":[{\"id\":1,\"required\":1,\"img\":{\"url\":\"https://wangmenglog.oss-ap-southeast-3.aliyuncs.com/static/20220810/2979648cd494f6c0d76796487fb891dc.jpg\",\"w\":1200,\"h\":627}}],\"link\":{\"url\":\"http://c3.sdkclickurl.com/click/?adgroup_id=4023292\\u0026user_id=699999\\u0026timeflag={timeflag}\\u0026mid={mid}\\u0026tid={click_id}\\u0026gaid={gaid}\\u0026android_id={android_id}\\u0026bundle={bundle}\\u0026pub_id={pub_id}\\u0026ip={ip}\\u0026ua={ua}\\u0026lang={lang}\",\"clicktrackers\":[\"https://click.lenovoadx.weshare.live/click?mid=4016\\u0026c=eyJkYXRlIjoiMjAyMzA1MDkiLCJob3VyIjoiMDYiLCJwdmVyIjoidjIuMC42Iiwic3VwcGx5X2lkIjo0MDE2LCJuYnIiOjAsImNvdW50cnkiOiJJRCIsInBsYXRmb3JtIjoyLCJ0cmFmZmljX3R5cGUiOiJhcHAiLCJpcCI6IjExNC4xMjQuMTgyLjE1IiwidWEiOiIiLCJnYWlkIjoiYWQ3ZGM1Y2MtODRiYi00Y2MyLWExMjAtODI3Y2QxYjU1YzdkMSIsImFkX2Zvcm1hdCI6OTAwMiwiYWRfc2l6ZSI6IjEyMDAqNjI3IiwiZG9tYWluIjoiIiwiYXBwX2J1bmRsZSI6InNoYXJlLnNoYXJla2Fyby5wcm8xIiwicHVibGlzaGVyX2lkIjoiIiwicGxhY2VtZW50IjoiMjUyNCIsInJlcXVlc3RfaWQiOiI3ODk4MzU5dTl5dXU0NjQ1cjQiLCJiaWRmbG9vciI6MC4wMTExMTExMTExMTExMTExMTIsImJpZGZsb29yY3VyIjoiVVNEIiwiYmlkX3ByaWNlIjowLjAyMjIyMjIyMjIyMjIyMjIyMywid2luX3ByaWNlIjowLCJyZXF1ZXN0X2JvZHkiOiIiLCJkY19sb2NhdGlvbiI6Imxlbm9hZHhzZyIsIm5hc3NldCI6ImltZzMiLCJpbnN0bCI6MCwiZGV2aWNldHlwZSI6NCwibWFrZSI6IkhVQVdFSSIsIm1vZGVsIjoiTUhBLUwyOSIsIm9zIjoiQW5kcm9pZCIsIndpbm5lciI6ODAwLCJkZW1hbmQiOlt7ImRlbWFuZF9pZCI6ODAwLCJiaWRfaWQiOiI0MDM5NzczIiwiY2FtcGFpZ25faWQiOiIyMzQ4NSIsImFkX2lkIjoiNDAzOTc3MyIsImNyZWF0aXZlX2lkIjoiNDAzOTc3MzEzIiwiZG9tYWluIjoiIiwiYXBwX2J1bmRsZSI6InJvb2NrLnRlc3QiLCJuYnIiOjAsImJpZF9wcmljZSI6MS44NDQwMTAxMTExMTExMTExLCJ3aW5fcHJpY2UiOjEuODQ0MDEwMTExMTExMTExMSwid2luX25vdGlmaWVyIjo5MDAyLCJlbmRwb2ludF9pZCI6NTksInRtdCI6ImRlYnVnIiwicmVxX3RpbWUiOjAsIm1hZ2ljX2FwcF9idW5kbGUiOiIiLCJjb252X3R5cGUiOjgwMDEsImF0dHJpYnV0aW9uX3R5cGUiOjgwMDIsImNwYyI6MCwicmVxdWVzdCI6MCwiYmlkIjowLCJ3aW4iOjAsInRpbWVvdXQiOjAsImltcHJlc3Npb24iOjAsImNsaWNrIjowLCJjb3N0IjowLCJjb252ZXJzaW9uIjowLCJjb252ZXJzaW9uX3JldmVudWUiOjAsIml1cmwiOiIiLCJjdXJsIjoiIiwibnVybCI6IiIsImx1cmwiOiIifV0sInJlcXVlc3QiOjAsImJpZCI6MCwid2luIjowLCJ0aW1lb3V0IjowLCJpbXByZXNzaW9uIjowLCJjbGljayI6MCwiY29zdCI6MCwiY29udmVyc2lvbiI6MCwiY29udmVyc2lvbl9yZXZlbnVlIjowfQ==\\u0026time=20230509061648\"]}}}",
										"bundle":"roock.test",
										"cid":"23485",
										"crid":"403977313",
										"attr":[
											3,
											4,
											5
										],
										"ext":{
											"lp_country":"id"
										}
									}
								],
								"seat":"roockmobile"
							}
						],
						"bidid":"9009d2a009bac110c61a13e3006d12d0",
						"cur":"USD"
					}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_younger_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82",
						"seatbid":[
							{
								"bid":[
									{
										"id":"4039773",
										"impid":"1",
										"price":0.022222222222222223,
										"adid":"4039773",
										"nurl":"https://win.lenovoadx.weshare.live/win?mid=4016\u0026c=eyJkYXRlIjoiMjAyMzA1MDkiLCJob3VyIjoiMDYiLCJwdmVyIjoidjIuMC42Iiwic3VwcGx5X2lkIjo0MDE2LCJuYnIiOjAsImNvdW50cnkiOiJKUCIsInBsYXRmb3JtIjoyLCJ0cmFmZmljX3R5cGUiOiJhcHAiLCJpcCI6IjExMC4xMjkuMTgyLjE1IiwidWEiOiIiLCJnYWlkIjoiYWQ3ZGM1Y2MtODRiYi00Y2MyLWExMjAtODI3Y2QxYjU1YzdkMSIsImFkX2Zvcm1hdCI6OTAwMSwiYWRfc2l6ZSI6IjMyMCo1MCIsImRvbWFpbiI6IiIsImFwcF9idW5kbGUiOiJzaGFyZS5zaGFyZWthcm8ucHJvIiwicHVibGlzaGVyX2lkIjoiIiwicGxhY2VtZW50IjoiMjUyNCIsInJlcXVlc3RfaWQiOiI2ZTBiYmMyYy02OTc0LTQ3MTktYTVmZC05MWIyODAzMTdiODIiLCJiaWRmbG9vciI6MC4wMTExMTExMTExMTExMTExMTIsImJpZGZsb29yY3VyIjoiVVNEIiwiYmlkX3ByaWNlIjowLjAyMjIyMjIyMjIyMjIyMjIyMywid2luX3ByaWNlIjowLCJyZXF1ZXN0X2JvZHkiOiIiLCJkY19sb2NhdGlvbiI6Imxlbm9hZHhzZyIsIm5hc3NldCI6IiIsImluc3RsIjowLCJkZXZpY2V0eXBlIjo0LCJtYWtlIjoiSFVBV0VJIiwibW9kZWwiOiJNSEEtTDI5Iiwib3MiOiJBbmRyb2lkIiwid2lubmVyIjo4MDAsImRlbWFuZCI6W3siZGVtYW5kX2lkIjo4MDAsImJpZF9pZCI6IjQwMzk3NzMiLCJjYW1wYWlnbl9pZCI6IjIzNDg1IiwiYWRfaWQiOiI0MDM5NzczIiwiY3JlYXRpdmVfaWQiOiI0MDM5NzczMTMiLCJkb21haW4iOiIiLCJhcHBfYnVuZGxlIjoicm9vY2sudGVzdCIsIm5iciI6MCwiYmlkX3ByaWNlIjowLjEyNzc4MTExMTExMTExMTEsIndpbl9wcmljZSI6MC4xMjc3ODExMTExMTExMTExLCJ3aW5fbm90aWZpZXIiOjkwMDIsImVuZHBvaW50X2lkIjo1OSwidG10IjoiZGVidWciLCJyZXFfdGltZSI6MywibWFnaWNfYXBwX2J1bmRsZSI6IiIsImNvbnZfdHlwZSI6ODAwMSwiYXR0cmlidXRpb25fdHlwZSI6ODAwMSwiY3BjIjowLCJyZXF1ZXN0IjowLCJiaWQiOjAsIndpbiI6MCwidGltZW91dCI6MCwiaW1wcmVzc2lvbiI6MCwiY2xpY2siOjAsImNvc3QiOjAsImNvbnZlcnNpb24iOjAsImNvbnZlcnNpb25fcmV2ZW51ZSI6MCwiaXVybCI6IiIsImN1cmwiOiIiLCJudXJsIjoiaHR0cHM6Ly9yZXEubGVub3ZvYWR4Lndlc2hhcmUubGl2ZS93aW4iLCJsdXJsIjoiIn1dLCJyZXF1ZXN0IjowLCJiaWQiOjAsIndpbiI6MCwidGltZW91dCI6MCwiaW1wcmVzc2lvbiI6MCwiY2xpY2siOjAsImNvc3QiOjAsImNvbnZlcnNpb24iOjAsImNvbnZlcnNpb25fcmV2ZW51ZSI6MH0=\u0026time=20230509062113\u0026price=${AUCTION_PRICE}\u0026bidid=${AUCTION_BID_ID}\u0026ext=800-9001",
										"lurl":"https://req.lenovoadx.weshare.live/lwin",
										"adm":"\u003c!DOCTYPE html\u003e\u003chead\u003e\u003cstyle\u003ebody{padding:0;margin:0;font-family:Arial,sans-serif}\u003c/style\u003e\u003c/head\u003e\u003cbody\u003e \u003cimg src='https://imp.lenovoadx.weshare.live/imp?mid=4016\u0026c=eyJkYXRlIjoiMjAyMzA1MDkiLCJob3VyIjoiMDYiLCJwdmVyIjoidjIuMC42Iiwic3VwcGx5X2lkIjo0MDE2LCJuYnIiOjAsImNvdW50cnkiOiJKUCIsInBsYXRmb3JtIjoyLCJ0cmFmZmljX3R5cGUiOiJhcHAiLCJpcCI6IjExMC4xMjkuMTgyLjE1IiwidWEiOiIiLCJnYWlkIjoiYWQ3ZGM1Y2MtODRiYi00Y2MyLWExMjAtODI3Y2QxYjU1YzdkMSIsImFkX2Zvcm1hdCI6OTAwMSwiYWRfc2l6ZSI6IjMyMCo1MCIsImRvbWFpbiI6IiIsImFwcF9idW5kbGUiOiJzaGFyZS5zaGFyZWthcm8ucHJvIiwicHVibGlzaGVyX2lkIjoiIiwicGxhY2VtZW50IjoiMjUyNCIsInJlcXVlc3RfaWQiOiI2ZTBiYmMyYy02OTc0LTQ3MTktYTVmZC05MWIyODAzMTdiODIiLCJiaWRmbG9vciI6MC4wMTExMTExMTExMTExMTExMTIsImJpZGZsb29yY3VyIjoiVVNEIiwiYmlkX3ByaWNlIjowLjAyMjIyMjIyMjIyMjIyMjIyMywid2luX3ByaWNlIjowLCJyZXF1ZXN0X2JvZHkiOiIiLCJkY19sb2NhdGlvbiI6Imxlbm9hZHhzZyIsIm5hc3NldCI6IiIsImluc3RsIjowLCJkZXZpY2V0eXBlIjo0LCJtYWtlIjoiSFVBV0VJIiwibW9kZWwiOiJNSEEtTDI5Iiwib3MiOiJBbmRyb2lkIiwid2lubmVyIjo4MDAsImRlbWFuZCI6W3siZGVtYW5kX2lkIjo4MDAsImJpZF9pZCI6IjQwMzk3NzMiLCJjYW1wYWlnbl9pZCI6IjIzNDg1IiwiYWRfaWQiOiI0MDM5NzczIiwiY3JlYXRpdmVfaWQiOiI0MDM5NzczMTMiLCJkb21haW4iOiIiLCJhcHBfYnVuZGxlIjoicm9vY2sudGVzdCIsIm5iciI6MCwiYmlkX3ByaWNlIjowLjEyNzc4MTExMTExMTExMTEsIndpbl9wcmljZSI6MC4xMjc3ODExMTExMTExMTExLCJ3aW5fbm90aWZpZXIiOjkwMDIsImVuZHBvaW50X2lkIjo1OSwidG10IjoiZGVidWciLCJyZXFfdGltZSI6MywibWFnaWNfYXBwX2J1bmRsZSI6IiIsImNvbnZfdHlwZSI6ODAwMSwiYXR0cmlidXRpb25fdHlwZSI6ODAwMSwiY3BjIjowLCJyZXF1ZXN0IjowLCJiaWQiOjAsIndpbiI6MCwidGltZW91dCI6MCwiaW1wcmVzc2lvbiI6MCwiY2xpY2siOjAsImNvc3QiOjAsImNvbnZlcnNpb24iOjAsImNvbnZlcnNpb25fcmV2ZW51ZSI6MCwiaXVybCI6IiIsImN1cmwiOiIiLCJudXJsIjoiIiwibHVybCI6IiJ9XSwicmVxdWVzdCI6MCwiYmlkIjowLCJ3aW4iOjAsInRpbWVvdXQiOjAsImltcHJlc3Npb24iOjAsImNsaWNrIjowLCJjb3N0IjowLCJjb252ZXJzaW9uIjowLCJjb252ZXJzaW9uX3JldmVudWUiOjB9\u0026time=20230509062113\u0026price=${AUCTION_PRICE}\u0026bidid=${AUCTION_BID_ID}\u0026h=lenosg002\u0026secure=8001\u0026ifa=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1\u0026bundle=share.sharekaro.pro\u0026ext=800-9001\u0026img.gif' border='0' width='1' alt='' height='1'\u003e \u003ca href=\"http://c3.sdkclickurl.com/click/?adgroup_id=4023292\u0026user_id=699999\u0026timeflag=20230509062113\u0026mid=108\u0026tid=20230509062113_adx_108_6e0bbc2c-6974-4719-a5fd-91b280317b82\u0026gaid=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1\u0026android_id=ad7dc5cc-84bb-4cc2-a120-827cd1b55c7d1\u0026bundle=roock.test\u0026pub_id=108_ttt\u0026ip=110.129.182.15\u0026ua=Mozilla%2F5.0+%28Linux%3B+Android+8.0.0%3B+SM-J810G+Build%2FR16NW%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F89.0.4389.90+Mobile+Safari%2F537.36\u0026lang=en\"\u003e\u003cimg src=\"https://wangmenglog.oss-ap-southeast-3.aliyuncs.com/static/20220810/2979648cd494f6c0d76796487fb891dc.jpg\" border=\"0\" alt=\"\" width=\"320px\" height=\"50px\"\u003e\u003c/a\u003e\u003c/body\u003e",
										"bundle":"roock.test",
										"cid":"23485",
										"crid":"403977313",
										"attr":[
											3,
											4,
											5
										],
										"w":320,
										"h":50,
										"ext":{
											"lp_country":"jp"
										}
									}
								],
								"seat":"roockmobile"
							}
						],
						"bidid":"e136ca870b860f1205ef58d2a7884e32",
						"cur":"USD"
					}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_younger_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"4a7350a4-0f2e-0f2c-421d-a90c98bdddf4",
						"seatbid":[
							{
								"bid":[
									{
										"id":"f3bdb09d-3ef2-11ed-9135-5c89237aa8f4",
										"impid":"1",
										"price":0.12345,
										"nurl":"https://win.adx.domain.com/win?mid=800&c=test&time=20220901055901&price=${AUCTION_PRICE}&bidid=${AUCTION_BID_ID}",
										"lurl":"http://lwin.adx.domain.com/lwin?mid=800&c=test&time=20220901055901&lose=${AUCTION_LOSS}&bidid=${AUCTION_BID_ID}",
										"adm":"<?xml version=\"1.0\" encoding=\"UTF-8\" ?><VAST version=\"3.0\"><Ad id=\"9199\"><Wrapper><AdSystem><![CDATA[Beeswax]]></AdSystem><VASTAdTagURI><![CDATA[https://vast.test.adtag/]]></VASTAdTagURI><Impression id=\"0\"><![CDATA[https://imp.adx.domain.com/imp?mid=800&c=test&time=20220901055901&price=${AUCTION_PRICE}&bidid=${AUCTION_BID_ID}&h=sg1&ifa=xxx&bundle=test.package&img.gif]]></Impression><Error><![CDATA[https://error.log/log/?error_code=[ERRORCODE]]]></Error><Creatives><Creative><Linear><TrackingEvents><Tracking event=\"start\"><![CDATA[https://test.tracking.log/log/act/ap?ai==&]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCACKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCADKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCAEKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"complete\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCAFKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"mute\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCAHKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"unmute\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCAIKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"pause\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCAJKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"resume\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCAKKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"fullscreen\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCALKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking><Tracking event=\"close\"><![CDATA[https://test.tracking.log/log/act/ap?ai=AFGgJhcCAMKihtaWQuRThEOThDNzQtMjc0OC00MkQ3LThBQzctQ0E2MEMyRjUzQUM2QO9HSM0BUgJhcGAAeh4SBAgDEAESBAgEEAESBAgFEAESBAgBEAESBAgCEAE=&]]></Tracking></TrackingEvents><VideoClicks><ClickTracking><![CDATA[https://test.tracking.log/?c=e&m=ck&key=639f0b397abe92d89e8b0c43e3b79334]]></ClickTracking><ClickTracking><![CDATA[https://test.tracking.log/log/clk/ap?ai=AFGgJhcCIobWlkLkU4RDk4Qzc0LTI3NDgtNDJENy04QUM3LUNBNjBDMkY1M0FDNjjvR0DNAUgBUgJhcGAAeh4SBAgEEAESBAgDEAESBAgCEAESBAgBEAESBAgFEAE=&audit_flag_wp=11.62405]]></ClickTracking></VideoClicks></Linear></Creative></Creatives><Extensions></Extensions></Wrapper></Ad></VAST>",
										"bundle":"com.test",
										"adomain":[
											"com.test"
										],
										"cid":"23485",
										"crid":"403977313",
										"cat":[
											"IAB16-5"
										]
									}
								],
								"seat":"acc-123"
							}
						],
						"bidid":"f3bdb09d-3ef2-11ed-9135-5c89237aa8f4",
						"cur":"USD"
					}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//oem-oppo
	r.HandleFunc("/mock/868/get_oem_oppo_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
					"retreival_ad_info":[
						{
							"pos_id":2800,
							"ad_info":[
								{
									"scoring":{
										"local_id":5085,
										"ecpm":35000000
									},
									"ad_info":{
										"ad_id":101694233,
										"target_page":7,
										"creative_info_map":{
											"1001694233":{
												"creative_id":1001694233,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":101694233,
											"campaign":{
												"id":101694233,
												"objective":2,
												"billing_type":4,
												"bid_price":35000000,
												"ad_app_pkg_name":"com.exness.android.pa",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=2191600&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":11694366,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"com.exness.android.pa",
													"app_name":"Exness 浜ゆ槗搴旂敤",
													"status":1,
													"gp_url":"https://play.google.com/store/apps/details?id=com.exness.android.pa&hl=en",
													"icon_url":"http://pmp-cdn.modengs.net/plutus/client/mr.png"
												},
												"apk_select_info":{
				
												},
												"order_type":1,
												"cpi_dsp_id":10005,
												"cpi_dsp_offer_id":"50683-4137853775"
											},
											"start_time":1283270400,
											"end_time":4102416000,
											"placement_ids":[
												288,
												2756,
												2800,
												2532
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":2019,
										"ecpm":11000000,
										"ad_serie_num":1
									},
									"ad_info":{
										"ad_id":101701282,
										"target_page":7,
										"creative_info_map":{
											"1001701282":{
												"creative_id":1001701282,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":101701282,
											"campaign":{
												"id":101701282,
												"objective":2,
												"billing_type":4,
												"bid_price":11000000,
												"ad_app_pkg_name":"com.fintopia.idnEasycash.google",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=2200912&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":11698745,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"com.fintopia.idnEasycash.google",
													"app_name":"com.lingyue.easycash.EasyCashApplication",
													"status":1,
													"version_name":"3.24.0",
													"version_code":"32400",
													"package_size":25386077,
													"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2022-10-25/com.fintopia.idnEasycash.google-1666679910.apk",
													"gp_url":"https://play.google.com/store/apps/details?id=com.fintopia.idnEasycash.google&hl=en",
													"silent_install_key":"101e4c33312499633493387d801563bda5e3cea3e6ac86f00fc1472b27835dc1c02a6e83df3ff3499a50e2a98c2b236278b9b70d6eaf8cf0d9865aad245b75c2",
													"icon_url":"http://pmp-cdn.modengs.net/plutus/client/mr.png",
													"pi_status":2
												},
												"apk_select_info":{
				
												},
												"order_type":1,
												"cpi_dsp_id":10004,
												"cpi_dsp_offer_id":"18173"
											},
											"start_time":1283270400,
											"end_time":4102416000,
											"placement_ids":[
												288,
												2756,
												2800,
												2532
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":856,
										"ecpm":9100000,
										"ad_serie_num":2
									},
									"ad_info":{
										"ad_id":101703132,
										"target_page":7,
										"creative_info_map":{
											"1001703132":{
												"creative_id":1001703132,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":101703132,
											"campaign":{
												"id":101703132,
												"objective":2,
												"billing_type":4,
												"bid_price":9100000,
												"ad_app_pkg_name":"com.ocbcnisp.onemobileapp",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=2203369&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":11677977,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"com.ocbcnisp.onemobileapp",
													"app_name":"ONe Mobile",
													"status":1,
													"gp_url":"https://play.google.com/store/apps/details?id=com.ocbcnisp.onemobileapp&hl=en"
												},
												"apk_select_info":{
				
												},
												"order_type":1,
												"cpi_dsp_id":10015,
												"cpi_dsp_offer_id":"1456343"
											},
											"start_time":1283270400,
											"end_time":4102416000,
											"placement_ids":[
												288,
												2756,
												2800,
												2532
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":7198,
										"ecpm":20000000,
										"ad_serie_num":3
									},
									"ad_info":{
										"ad_id":101682078,
										"target_page":7,
										"creative_info_map":{
											"1001682078":{
												"creative_id":1001682078,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":101682078,
											"campaign":{
												"id":101682078,
												"objective":2,
												"billing_type":4,
												"bid_price":20000000,
												"ad_app_pkg_name":"com.finaccel.android",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=2173857&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":11014643,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"com.finaccel.android",
													"app_name":"Kredivo",
													"status":1,
													"version_name":"3.5.7",
													"version_code":"903050721",
													"package_size":18257883,
													"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2022-01-29/com.finaccel.android-1643458522.sapk",
													"gp_url":"https://play.google.com/store/apps/details?id=com.finaccel.android&hl=en",
													"silent_install_key":"097da414c2cf4478d16954ee4517054aa29deaff14285c955bd558930ebc5f47c9881cb9ad17f58042b1b404c41655902d108a54913473499c0d12dcaacd3591",
													"icon_url":"http://amp-cdn-test.rqmob.com/apps/2022-01-29/imu4pi05g1tyyvew1qr63jhm23unib6o-1643458520.png",
													"pi_status":2
												},
												"apk_select_info":{
				
												},
												"order_type":1,
												"cpi_dsp_id":10001,
												"cpi_dsp_offer_id":"2326919"
											},
											"start_time":1283270400,
											"end_time":4102416000,
											"placement_ids":[
												288,
												2756,
												2800,
												2532
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":6536,
										"ecpm":8400000,
										"ad_serie_num":4
									},
									"ad_info":{
										"ad_id":101701886,
										"target_page":7,
										"creative_info_map":{
											"1001701886":{
												"creative_id":1001701886,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":101701886,
											"campaign":{
												"id":101701886,
												"objective":2,
												"billing_type":4,
												"bid_price":8400000,
												"ad_app_pkg_name":"com.octafx",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=2201692&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":57930648,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"com.octafx",
													"status":1,
													"icon_url":"http://pmp-cdn.modengs.net/plutus/client/mr.png"
												},
												"apk_select_info":{
				
												},
												"order_type":1,
												"cpi_dsp_id":10015,
												"cpi_dsp_offer_id":"1452122"
											},
											"start_time":1283270400,
											"end_time":4102416000,
											"placement_ids":[
												288,
												2756,
												2800,
												2532
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":2821,
										"ecpm":70000000,
										"ad_serie_num":5
									},
									"ad_info":{
										"ad_id":101681149,
										"target_page":7,
										"creative_info_map":{
											"1001681149":{
												"creative_id":1001681149,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":101681149,
											"campaign":{
												"id":101681149,
												"objective":2,
												"billing_type":4,
												"bid_price":70000000,
												"ad_app_pkg_name":"com.ticno.olymptrade",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=2167437&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":11004735,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"com.ticno.olymptrade",
													"status":1
												},
												"apk_select_info":{
				
												},
												"order_type":1,
												"cpi_dsp_id":10004,
												"cpi_dsp_offer_id":"16935"
											},
											"start_time":1283270400,
											"end_time":4102416000,
											"placement_ids":[
												288,
												2756,
												2800,
												2532
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":189,
										"ecpm":33000000,
										"ad_serie_num":6
									},
									"ad_info":{
										"ad_id":770992,
										"click_url":"https://pms.ushareit.me/index.php?m=bug&f=view&bugID=14487",
										"update_time":1675674758,
										"landing_page":"https://pms.ushareit.me/index.php?m=bug&f=view&bugID=14487",
										"creative_info_map":{
											"6388817":{
												"creative_id":6388817,
												"creative_type":20,
												"format_id":61
											}
										},
										"dsp_type":1009,
										"landing_page_info":[
											{
												"id":56294,
												"template_id":15,
												"template_type":2,
												"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1672933404281_bOzHdYAC.html?__id__=56294&__name__=biao&isAuto=false&isToNa=%7BisToNa%7D&version=7ef8f9db7633fe98543da2f694684466"
											},
											{
												"id":56295,
												"template_id":17,
												"template_type":2,
												"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1672933404408_CceGVCss.html?__id__=56295&__name__=biao&isAuto=false&isToNa=%7BisToNa%7D&version=566dc3a15efd373042a20447e18db4ee"
											},
											{
												"id":56301,
												"template_id":5,
												"template_type":1
											},
											{
												"id":56302,
												"template_id":9,
												"template_type":1
											}
										],
										"ad_set":{
											"id":7290,
											"campaign":{
												"id":12223,
												"objective":2,
												"billing_type":4,
												"attribution_tracker_url":"https://app.adjust.com/cu01w2h?campaign={campaign}&adgroup={adgroup}&creative={creative}&idfa={idfa}&gps_adid={gaid}&android_id={android_id}&install_callback=https%3A%2F%2Fping.rqmob.com%2Fpostback%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26platform%3Dadjust&reattribution_callback=https%3A%2F%2Fping.rqmob.com%2Fpostback%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26reinstalled_ts%3D%7Breinstalled_at%7D%26created_ts%3D%7Bcreated_at%7D%26is_retargeting%3D1%26platform%3Dadjust&rejected_install_callback=https%3A%2F%2Fping.rqmob.com%2Frejected%3Fclickid%3D{clickid}%26rejection_reason%3D%7Brejection_reason%7D&event_callback_wauti4=https%3A%2F%2Fping.rqmob.com%2Fevent%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26revenue_usd%3D%7Brevenue_usd%7D%26cost_type%3D%7Bcost_type%7D%26reporting_revenue%3D%7Breporting_revenue%7D%26event_token%3D%7Bevent%7D%26event_name%3D%7Bevent_name%7D%26platform%3Dadjust",
												"bid_price":33000000,
												"ad_app_pkg_name":"com.didiglobal.passenger",
												"start_time":1672933320,
												"end_time":4102416000,
												"ad_app_id":20695,
												"attribution_pi":1,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=29&campid=2193132&platform=adshonor&pkg=com.didiglobal.passenger&is_offline=1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"attribution_platform":"adjust",
												"app_install_id":11707981,
												"is_online_ad":true,
												"ad_app":{
													"id":20695,
													"package_name":"com.didiglobal.passenger",
													"app_name":"DiDi",
													"status":1,
													"version_name":"7.3.4",
													"version_code":"1107030404",
													"package_size":101286802,
													"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2022-07-21/com.didiglobal.passenger-1658389753.sapk",
													"gp_url":"https://play.google.com/store/apps/details?id=com.didiglobal.passenger&hl=en",
													"silent_install_key":"42785bd53f98f3fc70074157431b8019729f60216c364e8b8720868019d7658757699678945bb69826c98bb6f01d7d5cbd39857a0d7401b2eac57b584ca65b2b",
													"icon_url":"http://amp-cdn-test.rqmob.com/apps/2022-07-21/w3v08w9rz8nthtbv8mgp7farbh4r7gqr-1658389734.png",
													"pi_status":1
												},
												"apk_select_info":{
													"origin_id":20695,
													"strategy_id":20695
												},
												"app_category":"MAPS_AND_NAVIGATION",
												"campaign_id":"202301050009",
												"so_line_id":13,
												"order_type":1,
												"campaign_name":"鑽ｈ��adx-yym02"
											},
											"start_time":1672933320,
											"end_time":4102416000,
											"placement_ids":[
												1600,
												2532,
												2756,
												2800
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":3089,
										"ecpm":8000000,
										"ad_serie_num":7
									},
									"ad_info":{
										"ad_id":100653950,
										"target_page":7,
										"creative_info_map":{
											"1000653950":{
												"creative_id":1000653950,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":100653950,
											"campaign":{
												"id":100653950,
												"objective":2,
												"billing_type":4,
												"bid_price":8000000,
												"ad_app_pkg_name":"com.yinshan.program.banda",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=1985311&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":11000367,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"com.yinshan.program.banda",
													"status":1
												},
												"apk_select_info":{
				
												},
												"order_type":1,
												"cpi_dsp_id":10015,
												"cpi_dsp_offer_id":"1352621"
											},
											"start_time":1283270400,
											"end_time":4102416000,
											"placement_ids":[
												288,
												2756,
												2800,
												2532
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":3401,
										"ecpm":9000000,
										"ad_serie_num":8
									},
									"ad_info":{
										"ad_id":101702657,
										"target_page":7,
										"creative_info_map":{
											"1001702657":{
												"creative_id":1001702657,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":101702657,
											"campaign":{
												"id":101702657,
												"objective":2,
												"billing_type":4,
												"bid_price":9000000,
												"ad_app_pkg_name":"id.co.linebank",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=2202670&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":11000601,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"id.co.linebank",
													"app_name":"LINE Bank",
													"status":1,
													"version_name":"1.1.5",
													"version_code":"10105",
													"package_size":100350512,
													"cdn_url":"http://pmp-cdn.modengs.net/apps/2021-09-14/LINE Bank-1631657296.apk",
													"gp_url":"https://play.google.com/store/apps/details?id=id.co.linebank&hl=en&gl=ID",
													"silent_install_key":"2ce4488ab0b35365e5fc59387cff0369ffa49a65c8d6cbe010ab79a0e4036702f0cf57d53f2ae1d1a3771633d7db771f392f8630e1cf7b09ff4ec29223729480",
													"icon_url":"http://pmp-cdn.modengs.net/apps/2021-09-14/tmp_icon-1631657256.png",
													"pi_status":4
												},
												"apk_select_info":{
				
												},
												"order_type":1,
												"cpi_dsp_id":10015,
												"cpi_dsp_offer_id":"1454572"
											},
											"start_time":1283270400,
											"end_time":4102416000,
											"placement_ids":[
												288,
												2756,
												2800,
												2532
											]
										},
										"auto_download":1,
										"silently_install":true
									}
								},
								{
									"scoring":{
										"local_id":1752,
										"ecpm":8000000,
										"ad_serie_num":9
									},
									"ad_info":{
										"ad_id":101700994,
										"target_page":7,
										"creative_info_map":{
											"1001700994":{
												"creative_id":1001700994,
												"creative_type":20,
												"format_id":61,
												"SubAction":3
											}
										},
										"dsp_type":1009,
										"ad_set":{
											"id":101700994,
											"campaign":{
												"id":101700994,
												"objective":2,
												"billing_type":4,
												"bid_price":8000000,
												"ad_app_pkg_name":"com.ocbc.mobile",
												"start_time":1283270400,
												"end_time":4102416000,
												"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&campid=2200529&platform=shareit_attribution&pkg={pkg}&is_offline=-1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}",
												"app_install_id":11717946,
												"is_online_ad":true,
												"ad_app":{
													"package_name":"com.ocbc.mobile",
														"status":1
													},
													"apk_select_info":{
					
													},
													"order_type":1,
													"cpi_dsp_id":10015,
													"cpi_dsp_offer_id":"1447355"
												},
												"start_time":1283270400,
												"end_time":4102416000,
												"placement_ids":[
													288,
													2756,
													2800,
													2532
												]
											},
											"auto_download":1,
											"silently_install":true
										}
									}
								]
							}
						],
						"call_ma_success":true,
						"log_data":{
							"user_app_info":{
					
							}
						}
					}
				`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_oem_vivo_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"retreival_ad_info":[{"pos_id":2756,"ad_info":[{"scoring":{"local_id":265,"ecpm":70000000},"ad_info":{"ad_id":101681149,"target_page":7,"creative_info_map":{"1001681149":{"creative_id":1001681149,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101681149,"campaign":{"id":101681149,"objective":2,"billing_type":4,"bid_price":70000000,"ad_app_pkg_name":"com.ticno.olymptrade","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2167437&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11004735,"is_online_ad":true,"ad_app":{"package_name":"com.ticno.olymptrade","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10004,"cpi_dsp_offer_id":"16935"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":1526,"ecpm":20000000,"ad_serie_num":1},"ad_info":{"ad_id":101682078,"target_page":7,"creative_info_map":{"1001682078":{"creative_id":1001682078,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101682078,"campaign":{"id":101682078,"objective":2,"billing_type":4,"bid_price":20000000,"ad_app_pkg_name":"com.finaccel.android","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2173857&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11014643,"is_online_ad":true,"ad_app":{"package_name":"com.finaccel.android","app_name":"Kredivo","status":1,"gp_url":"https://play.google.com/store/apps/details?id=com.finaccel.android&amp;hl=en"},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10001,"cpi_dsp_offer_id":"2326919"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":4172,"ecpm":8000000,"ad_serie_num":2},"ad_info":{"ad_id":101700994,"target_page":7,"creative_info_map":{"1001700994":{"creative_id":1001700994,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101700994,"campaign":{"id":101700994,"objective":2,"billing_type":4,"bid_price":8000000,"ad_app_pkg_name":"com.ocbc.mobile","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2200529&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11717946,"is_online_ad":true,"ad_app":{"package_name":"com.ocbc.mobile","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1447355"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":3678,"ecpm":35000000,"ad_serie_num":3},"ad_info":{"ad_id":101694233,"target_page":7,"creative_info_map":{"1001694233":{"creative_id":1001694233,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101694233,"campaign":{"id":101694233,"objective":2,"billing_type":4,"bid_price":35000000,"ad_app_pkg_name":"com.exness.android.pa","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2191600&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11694366,"is_online_ad":true,"ad_app":{"package_name":"com.exness.android.pa","app_name":"Exness 浜ゆ槗搴旂敤","status":1,"gp_url":"https://play.google.com/store/apps/details?id=com.exness.android.pa&amp;hl=en","icon_url":"http://pmp-cdn.modengs.net/plutus/client/mr.png"},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10005,"cpi_dsp_offer_id":"50683-4137853775"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":2161,"ecpm":33000000,"ad_serie_num":4},"ad_info":{"ad_id":770992,"click_url":"https://pms.ushareit.me/index.php?m=bug&amp;f=view&amp;bugID=14487","update_time":1675674758,"landing_page":"https://pms.ushareit.me/index.php?m=bug&amp;f=view&amp;bugID=14487","creative_info_map":{"6388817":{"creative_id":6388817,"creative_type":20,"format_id":61}},"dsp_type":1009,"landing_page_info":[{"id":56294,"template_id":15,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1672933404281_bOzHdYAC.html?__id__=56294&amp;__name__=biao&amp;isAuto=false&amp;isToNa=%7BisToNa%7D&amp;version=7ef8f9db7633fe98543da2f694684466"},{"id":56295,"template_id":17,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1672933404408_CceGVCss.html?__id__=56295&amp;__name__=biao&amp;isAuto=false&amp;isToNa=%7BisToNa%7D&amp;version=566dc3a15efd373042a20447e18db4ee"},{"id":56301,"template_id":5,"template_type":1},{"id":56302,"template_id":9,"template_type":1}],"ad_set":{"id":7290,"campaign":{"id":12223,"objective":2,"billing_type":4,"attribution_tracker_url":"https://app.adjust.com/cu01w2h?campaign={campaign}&amp;adgroup={adgroup}&amp;creative={creative}&amp;idfa={idfa}&amp;gps_adid={gaid}&amp;android_id={android_id}&amp;install_callback=https%3A%2F%2Fping.rqmob.com%2Fpostback%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26platform%3Dadjust&amp;reattribution_callback=https%3A%2F%2Fping.rqmob.com%2Fpostback%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26reinstalled_ts%3D%7Breinstalled_at%7D%26created_ts%3D%7Bcreated_at%7D%26is_retargeting%3D1%26platform%3Dadjust&amp;rejected_install_callback=https%3A%2F%2Fping.rqmob.com%2Frejected%3Fclickid%3D{clickid}%26rejection_reason%3D%7Brejection_reason%7D&amp;event_callback_wauti4=https%3A%2F%2Fping.rqmob.com%2Fevent%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26revenue_usd%3D%7Brevenue_usd%7D%26cost_type%3D%7Bcost_type%7D%26reporting_revenue%3D%7Breporting_revenue%7D%26event_token%3D%7Bevent%7D%26event_name%3D%7Bevent_name%7D%26platform%3Dadjust","bid_price":33000000,"ad_app_pkg_name":"com.didiglobal.passenger","start_time":1672933320,"end_time":4102416000,"ad_app_id":20695,"attribution_pi":1,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=29&amp;campid=2193132&amp;platform=adshonor&amp;pkg=com.didiglobal.passenger&amp;is_offline=1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","attribution_platform":"adjust","app_install_id":11707981,"is_online_ad":true,"ad_app":{"id":20695,"package_name":"com.didiglobal.passenger","app_name":"DiDi","status":1,"version_name":"7.3.4","version_code":"1107030404","package_size":101286802,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2022-07-21/com.didiglobal.passenger-1658389753.sapk","gp_url":"https://play.google.com/store/apps/details?id=com.didiglobal.passenger&amp;hl=en","silent_install_key":"42785bd53f98f3fc70074157431b8019729f60216c364e8b8720868019d7658757699678945bb69826c98bb6f01d7d5cbd39857a0d7401b2eac57b584ca65b2b","icon_url":"http://amp-cdn-test.rqmob.com/apps/2022-07-21/w3v08w9rz8nthtbv8mgp7farbh4r7gqr-1658389734.png","pi_status":1},"apk_select_info":{"origin_id":20695,"strategy_id":20695},"app_category":"MAPS_AND_NAVIGATION","campaign_id":"202301050009","so_line_id":13,"order_type":1,"campaign_name":"鑽ｈ��adx-yym02"},"start_time":1672933320,"end_time":4102416000,"placement_ids":[1600,2532,2756,2800]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":7160,"ecpm":8000000,"ad_serie_num":5},"ad_info":{"ad_id":100653950,"target_page":7,"creative_info_map":{"1000653950":{"creative_id":1000653950,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":100653950,"campaign":{"id":100653950,"objective":2,"billing_type":4,"bid_price":8000000,"ad_app_pkg_name":"com.yinshan.program.banda","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=1985311&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11000367,"is_online_ad":true,"ad_app":{"package_name":"com.yinshan.program.banda","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1352621"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":5796,"ecpm":3000000,"ad_serie_num":6},"ad_info":{"ad_id":100664283,"target_page":7,"creative_info_map":{"1000664283":{"creative_id":1000664283,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":100664283,"campaign":{"id":100664283,"objective":2,"billing_type":4,"bid_price":3000000,"ad_app_pkg_name":"id.moxa","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2034301&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":57999612,"is_online_ad":true,"ad_app":{"package_name":"id.moxa","app_name":"moxa","status":1,"version_name":"2.0.11","version_code":"3146044","package_size":54048839,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2021-09-16/moxa-1631771736.sapk","gp_url":"https://play.google.com/store/apps/details?id=id.moxa&amp;hl=en&amp;gl=ID","silent_install_key":"2b2805b0ce4cfde3bee6b5401d106401c923534864ddac95ec927b23b74809103e35c3d362c2ba941618b7317338979d94f7555e8d1c1e62e77278a1145be362","icon_url":"http://amp-cdn-test.rqmob.com/apps/2021-09-16/tmp_icon-1631771692.png","pi_status":2},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1400358"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":91,"ecpm":5500000,"ad_serie_num":7},"ad_info":{"ad_id":101721869,"target_page":7,"creative_info_map":{"1001721869":{"creative_id":1001721869,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101721869,"campaign":{"id":101721869,"objective":2,"billing_type":4,"bid_price":5500000,"ad_app_pkg_name":"id.com.uiux.mobile","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2222302&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":406422390,"is_online_ad":true,"ad_app":{"package_name":"id.com.uiux.mobile","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10016,"cpi_dsp_offer_id":"32430"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":2628,"ecpm":3000000,"ad_serie_num":8},"ad_info":{"ad_id":101811601,"target_page":7,"creative_info_map":{"1001811601":{"creative_id":1001811601,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101811601,"campaign":{"id":101811601,"objective":2,"billing_type":4,"bid_price":3000000,"ad_app_pkg_name":"com.octafx.copytrade","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2317489&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":406214573,"is_online_ad":true,"ad_app":{"package_name":"com.octafx.copytrade","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10024,"cpi_dsp_offer_id":"33539"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":3870,"ecpm":9000000,"ad_serie_num":9},"ad_info":{"ad_id":101702657,"target_page":7,"creative_info_map":{"1001702657":{"creative_id":1001702657,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101702657,"campaign":{"id":101702657,"objective":2,"billing_type":4,"bid_price":9000000,"ad_app_pkg_name":"id.co.linebank","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2202670&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11000601,"is_online_ad":true,"ad_app":{"package_name":"id.co.linebank","app_name":"LINE Bank","status":1,"version_name":"1.1.5","version_code":"10105","package_size":100350512,"cdn_url":"http://pmp-cdn.modengs.net/apps/2021-09-14/LINE Bank-1631657296.apk","gp_url":"https://play.google.com/store/apps/details?id=id.co.linebank&amp;hl=en&amp;gl=ID","silent_install_key":"2ce4488ab0b35365e5fc59387cff0369ffa49a65c8d6cbe010ab79a0e4036702f0cf57d53f2ae1d1a3771633d7db771f392f8630e1cf7b09ff4ec29223729480","icon_url":"http://pmp-cdn.modengs.net/apps/2021-09-14/tmp_icon-1631657256.png","pi_status":4},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1454572"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":6856,"ecpm":8400000,"ad_serie_num":10},"ad_info":{"ad_id":101701886,"target_page":7,"creative_info_map":{"1001701886":{"creative_id":1001701886,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101701886,"campaign":{"id":101701886,"objective":2,"billing_type":4,"bid_price":8400000,"ad_app_pkg_name":"com.octafx","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2201692&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":57930648,"is_online_ad":true,"ad_app":{"package_name":"com.octafx","status":1,"icon_url":"http://pmp-cdn.modengs.net/plutus/client/mr.png"},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1452122"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":3033,"ecpm":2280000,"ad_serie_num":11},"ad_info":{"ad_id":101699770,"target_page":7,"creative_info_map":{"1001699770":{"creative_id":1001699770,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101699770,"campaign":{"id":101699770,"objective":2,"billing_type":4,"bid_price":2280000,"ad_app_pkg_name":"app.ndtv.matahari","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2199127&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11697857,"is_online_ad":true,"ad_app":{"package_name":"app.ndtv.matahari","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1446162"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":5856,"ecpm":2000000,"ad_serie_num":12},"ad_info":{"ad_id":101675695,"target_page":7,"creative_info_map":{"1001675695":{"creative_id":1001675695,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101675695,"campaign":{"id":101675695,"objective":2,"billing_type":4,"bid_price":2000000,"ad_app_pkg_name":"com.boyaa.androidmarketid","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2133794&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11016903,"is_online_ad":true,"ad_app":{"package_name":"com.boyaa.androidmarketid","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10022,"cpi_dsp_offer_id":"24693949"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":4877,"ecpm":8000000,"ad_serie_num":13},"ad_info":{"ad_id":101680200,"target_page":7,"creative_info_map":{"1001680200":{"creative_id":1001680200,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101680200,"campaign":{"id":101680200,"objective":2,"billing_type":4,"bid_price":8000000,"ad_app_pkg_name":"com.pegipegi.android","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2163815&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11019145,"is_online_ad":true,"ad_app":{"package_name":"com.pegipegi.android","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1414921"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":596,"ecpm":6000000,"ad_serie_num":14},"ad_info":{"ad_id":101699802,"target_page":7,"creative_info_map":{"1001699802":{"creative_id":1001699802,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101699802,"campaign":{"id":101699802,"objective":2,"billing_type":4,"bid_price":6000000,"ad_app_pkg_name":"com.avatrade.mobile","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2199170&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":406424329,"is_online_ad":true,"ad_app":{"package_name":"com.avatrade.mobile","app_name":"AvaTradeGO","status":1,"version_name":"94.2.0","version_code":"1190940200","package_size":20393267,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2021-09-10/AvaTradeGO-1631269936.apk","gp_url":"https://play.google.com/store/apps/details?id=com.avatrade.mobile&amp;hl=en","silent_install_key":"8cf4657055bcc788e1173e17a329793781bd2b521aeeb7c972312a65f77edf2146e180944d06fc8d8742ca3bb2cb8ab4b804bea5ca9b59a415bfb6d6b156d6a4","icon_url":"http://amp-cdn-test.rqmob.com/apps/2021-09-10/tmp_icon-1631269918.png","pi_status":2},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1446885"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":5049,"ecpm":3000000,"ad_serie_num":15},"ad_info":{"ad_id":101754815,"target_page":7,"creative_info_map":{"1001754815":{"creative_id":1001754815,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101754815,"campaign":{"id":101754815,"objective":2,"billing_type":4,"bid_price":3000000,"ad_app_pkg_name":"com.grabtaxi.passenger","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2260052&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11028287,"is_online_ad":true,"ad_app":{"package_name":"com.grabtaxi.passenger","status":1},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10018,"cpi_dsp_offer_id":"1419943"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":5792,"ecpm":4000000,"ad_serie_num":16},"ad_info":{"ad_id":101702622,"target_page":7,"creative_info_map":{"1001702622":{"creative_id":1001702622,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101702622,"campaign":{"id":101702622,"objective":2,"billing_type":4,"bid_price":4000000,"ad_app_pkg_name":"com.binance.dev","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2202622&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11005846,"is_online_ad":true,"ad_app":{"package_name":"com.binance.dev","app_name":"甯佸畨","status":1,"version_name":"1.41.2","version_code":"10014102","package_size":98147267,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2021-06-07/nep34dti9mz9ffjh9a8hcuma981on7m1.apk","gp_url":"https://play.google.com/store/apps/details?id=com.binance.dev&amp;hl=en","silent_install_key":"4cafdac5b91d570d8b1976801bdad3c69c473594aab2c6f671b8d476b97b4d1ea61188e8c137d290314b5ca2ab0b0692b478a70578ff003553abbf1c6d1c98fb","icon_url":"http://amp-cdn-test.rqmob.com/apps/2021-06-07/79m8celndlpmhaqkia1pgiqnhc1wsppc.png","pi_status":2},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10010,"cpi_dsp_offer_id":"3040815"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":1548,"ecpm":8000000,"ad_serie_num":17},"ad_info":{"ad_id":101702745,"target_page":7,"creative_info_map":{"1001702745":{"creative_id":1001702745,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101702745,"campaign":{"id":101702745,"objective":2,"billing_type":4,"bid_price":8000000,"ad_app_pkg_name":"com.ixdigit.hanson","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2202786&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11738183,"is_online_ad":true,"ad_app":{"package_name":"com.ixdigit.hanson","app_name":"com.ixdigit.android.core.application.IXApplication","status":1,"version_name":"2.1.0.12.6.0","version_code":"21012060","package_size":44469420,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2023-01-11/com.ixdigit.hanson-1673431989.apk","gp_url":"https://play.google.com/store/apps/details?id=com.ixdigit.hanson&amp;hl=en","silent_install_key":"d332b72191b8f6ee8dc0701e7fc4233264d0b89f2deb0c45f259219ff4fb4c7fa5ea0ae8c8fa23577da2d6b4c8c66bd677f532e37ec6a3e5476264320d977bfb","icon_url":"http://amp-cdn-test.rqmob.com/icons/2023-01-11/j80s3tu94vt3cb2y7hp36ukbnl2taons.png","pi_status":2},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1455240"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":4018,"ecpm":9100000,"ad_serie_num":18},"ad_info":{"ad_id":101703132,"target_page":7,"creative_info_map":{"1001703132":{"creative_id":1001703132,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101703132,"campaign":{"id":101703132,"objective":2,"billing_type":4,"bid_price":9100000,"ad_app_pkg_name":"com.ocbcnisp.onemobileapp","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2203369&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11677977,"is_online_ad":true,"ad_app":{"package_name":"com.ocbcnisp.onemobileapp","app_name":"ONe Mobile","status":1,"version_name":"4.0.3","version_code":"300010","package_size":23989638,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2021-09-16/ONe Mobile-1631769416.sapk","gp_url":"https://play.google.com/store/apps/details?id=com.ocbcnisp.onemobileapp&amp;hl=en&amp;gl=ID","silent_install_key":"c70a47ece87d8e9afb5d31fd3b45bdb691e8ba883f76c53b5169aa95a538fe260b50e4edf375a3499e8091909364c8854a8881415ae41d99db32216cae067c65","icon_url":"http://amp-cdn-test.rqmob.com/apps/2021-09-16/tmp_icon-1631769409.png","pi_status":2},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10015,"cpi_dsp_offer_id":"1456343"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":6207,"ecpm":11000000,"ad_serie_num":19},"ad_info":{"ad_id":101701282,"target_page":7,"creative_info_map":{"1001701282":{"creative_id":1001701282,"creative_type":20,"format_id":61,"SubAction":3}},"dsp_type":1009,"ad_set":{"id":101701282,"campaign":{"id":101701282,"objective":2,"billing_type":4,"bid_price":11000000,"ad_app_pkg_name":"com.fintopia.idnEasycash.google","start_time":1283270400,"end_time":4102416000,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=30&amp;campid=2200912&amp;platform=shareit_attribution&amp;pkg={pkg}&amp;is_offline=-1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;publisher={publisher}&amp;oaid={oaid}","app_install_id":11698745,"is_online_ad":true,"ad_app":{"package_name":"com.fintopia.idnEasycash.google","app_name":"com.lingyue.easycash.EasyCashApplication","status":1,"version_name":"3.24.0","version_code":"32400","package_size":25386077,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2022-10-25/com.fintopia.idnEasycash.google-1666679910.apk","gp_url":"https://play.google.com/store/apps/details?id=com.fintopia.idnEasycash.google&amp;hl=en","silent_install_key":"101e4c33312499633493387d801563bda5e3cea3e6ac86f00fc1472b27835dc1c02a6e83df3ff3499a50e2a98c2b236278b9b70d6eaf8cf0d9865aad245b75c2","icon_url":"http://pmp-cdn.modengs.net/plutus/client/mr.png","pi_status":2},"apk_select_info":{},"order_type":1,"cpi_dsp_id":10004,"cpi_dsp_offer_id":"18173"},"start_time":1283270400,"end_time":4102416000,"placement_ids":[288,2756,2800,2532]},"auto_download":1,"silently_install":true}}]}],"call_ma_success":true,"log_data":{"user_app_info":{}}}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_oem_tran_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"retreival_ad_info":[{"pos_id":2963,"ad_info":[{"scoring":{"local_id":6473,"ecpm":99900000},"ad_info":{"ad_id":771281,"update_time":1677658347,"creative_info_map":{"6389709":{"creative_id":6389709,"creative_type":19,"format_id":54,"format_width":72,"format_height":72,"major_image_id":2733}},"dsp_type":1009,"landing_page_info":[{"id":59144,"template_id":15,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1677658351603_RYqXHOkH.html?__id__=59144&__name__=YallaLudo&isAuto=false&isToNa=%7BisToNa%7D&version=2dc701edccd4356818006876746344c2"},{"id":59145,"template_id":17,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1677658351713_lxjidUZf.html?__id__=59145&__name__=YallaLudo&isAuto=false&isToNa=%7BisToNa%7D&version=18fd5ef0012d6665b7d54520021f969a"},{"id":59156,"template_id":5,"template_type":1},{"id":59157,"template_id":9,"template_type":1}],"ad_set":{"id":7563,"campaign":{"id":12486,"objective":2,"billing_type":4,"attribution_tracker_url":"https://app.adjust.com/cu01w2h?campaign={campaign}&adgroup={adgroup}&creative={creative}&idfa={idfa}&gps_adid={gaid}&android_id={android_id}&install_callback=https%3A%2F%2Fping.rqmob.com%2Fpostback%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26platform%3Dadjust&reattribution_callback=https%3A%2F%2Fping.rqmob.com%2Fpostback%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26reinstalled_ts%3D%7Breinstalled_at%7D%26created_ts%3D%7Bcreated_at%7D%26is_retargeting%3D1%26platform%3Dadjust&rejected_install_callback=https%3A%2F%2Fping.rqmob.com%2Frejected%3Fclickid%3D{clickid}%26rejection_reason%3D%7Brejection_reason%7D&event_callback_wauti4=https%3A%2F%2Fping.rqmob.com%2Fevent%3Fclickid%3D{clickid}%26app_name%3D%7Bapp_name%7D%26android_id%3D%7Bandroid_id%7D%26gaid%3D%7Bgps_adid%7D%26oaid%3D%7Boaid%7D%26country_code%3D%7Bcountry%7D%26language%3D%7Blanguage%7D%26os_version%3D%7Bos_version%7D%26store%3D%7Bstore%7D%26campaign%3D%7Bcampaign_name%7D%26app_version_name%3D%7Bapp_version%7D%26revenue_usd%3D%7Brevenue_usd%7D%26cost_type%3D%7Bcost_type%7D%26reporting_revenue%3D%7Breporting_revenue%7D%26event_token%3D%7Bevent%7D%26event_name%3D%7Bevent_name%7D%26platform%3Dadjust","bid_price":99900000,"ad_app_pkg_name":"com.yalla.yallagames","start_time":1677600000,"end_time":4102416000,"ad_app_id":20625,"attribution_pi":1,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=29&campid=2199543&platform=adshonor&pkg=com.yalla.yallagames&is_offline=1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&amp_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}","attribution_platform":"adjust","app_install_id":388200542,"is_online_ad":true,"ad_app":{"id":20625,"package_name":"com.yalla.yallagames","app_name":"Yalla Ludo","status":1,"version_name":"1.2.9.0","version_code":"1020904","package_size":131142064,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2022-01-29/com.yalla.yallagames-1643466353.sapk","gp_url":"https://play.google.com/store/apps/details?id=com.yalla.yallagames&hl=en","silent_install_key":"27d92fbfaf0ddd2bb5a27361ff210d1acae4a5f19ba4125b87cb02c9fab403eaf7e9c86fc6ad131cc1271d6f13265411a53d22bc5f5f406ab68d57b06d2e404b","icon_url":"http://amp-cdn-test.rqmob.com/apps/2022-01-29/eqzbw0tazc6ftmjb690vwbhc1dl2m056-1643466335.png","pi_status":2},"apk_select_info":{"origin_id":20625,"strategy_id":20625},"app_category":"GAME_BOARD","campaign_id":"202303010014","so_line_id":1,"order_type":1,"campaign_name":"CPI_YallaLudo_20230301_鍕垮姩"},"start_time":1677600000,"end_time":4102416000,"placement_ids":[347,599,698,818,2052,2055,2061,2128,2132,2134,2136,2138,2140,2142,2160,2255,2295,2323,2364,2365,2389,2409,2419,2429,2430,2433,2434,2435,2442,2443,2455,2478,2483,2497,2498,2514,2522,2533,2563,2586,2589,2590,2592,2633,2634,2635,2636,2659,2663,2665,2666,2667,2668,2669,2670,2671,2675,2676,2682,2683,2684,2685,2686,2687,2688,2689,2696,2699,2704,2709,2718,2719,2747,2751,2755,2758,2759,2767,2769,2774,2775,2776,2783,2795,2803,2804,2811,2817,2819,2822,2824,2826,2827,2828,2829,2835,2836,2842,2844,2845,2858,2859,2860,2861,2862,2863,2871,2873,2874,2875,2876,2915,2917,2958,2960,2963,2972,2975,2976,2990,2997,2998,2999,3000]},"auto_download":1,"silently_install":true}},{"scoring":{"local_id":239,"ecpm":80000,"ad_serie_num":1},"ad_info":{"ad_id":770804,"update_time":1672300368,"creative_info_map":{"6388069":{"creative_id":6388069,"creative_type":20,"format_id":61}},"dsp_type":1009,"landing_page_info":[{"id":54931,"template_id":15,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1669628833432_iFgHteWN.html?__id__=54931&__name__=Title&isAuto=false&isToNa=%7BisToNa%7D&version=605bc62544f884296594542a86013407"},{"id":54932,"template_id":15,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1669628833560_IAuYgHzG.html?__id__=54932&__name__=Title&isAuto=false&isToNa=%7BisToNa%7D&version=137d398c27d38fd4d956ebc2bbe970cb"},{"id":54933,"template_id":17,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1669628833641_dujCTaCq.html?__id__=54933&__name__=Title&isAuto=false&isToNa=%7BisToNa%7D&version=c226c48b2ad811c25f41bb22b2009daa"},{"id":54934,"template_id":17,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1669628833802_clUfible.html?__id__=54934&__name__=Title&isAuto=false&isToNa=%7BisToNa%7D&version=dfcf00570240f83708ad74f20c6db1a4"},{"id":54935,"template_id":5,"template_type":1},{"id":54936,"template_id":5,"template_type":1},{"id":54937,"template_id":9,"template_type":1},{"id":54938,"template_id":9,"template_type":1}],"ad_set":{"id":7069,"campaign":{"id":7381,"objective":2,"billing_type":4,"attribution_tracker_url":"https://app.appsflyer.com/com.shopee.ph?pid=shareit_int&c=topappsicon&af_click_lookback=7d&clickid={clickid}&android_id={android_id}&advertising_id={gaid}&idfa={idfa}&device_id={device_id}&beyla_id={beyla_id}&device_model={device_model}&city={city}&state={state}&redirect=false&oaid={oaid}","bid_price":80000,"ad_app_pkg_name":"com.karimcastagnini.castlestorm.mg","start_time":1669628694,"end_time":4102416000,"ad_app_id":20731,"attribution_pi":1,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=29&campid=2181508&platform=adshonor&pkg=com.karimcastagnini.castlestorm.mg&is_offline=1&requestid={requestid}&gaid={gaid}&device_id={device_id}&android_id={android_id}&beyla_id={beyla_id}&imei={imei}&imsi={imsi}&c={c}&site_id={site_id}&c_id={c_id}&adset={adset}&adset_id={adset_id}&ad={ad}&ad_id={ad_id}&ad_type={ad_type}&cost_model={cost_model}&cost_currency={cost_currency}&cost_value={cost_value}&remote_ip={remote_ip}&country_code={country_code}&ip={ip}&os_version={os_version}&placement={placement}&other_category={other_category}&ext_info={ext_info}&uagent={uagent}&sid={sid}&channel_pkg={channel_pkg}&channel_pkg_ver={channel_pkg_ver}&app_type={app_type}&app_id={app_id}&real_attrplat={real_attrplat}&adpos_id={adpos_id}&midas_camp_id={midas_camp_id}&midas_traffic_type={midas_traffic_type}&amp_app_id={amp_app_id}&is_pre_install={is_pre_install}&cut_type={cut_type}&package_type={package_type}&publisher={publisher}&oaid={oaid}","attribution_platform":"appsflyer","is_online_ad":true,"ad_app":{"id":20731,"package_name":"com.karimcastagnini.castlestorm.mg","app_name":"com.stub.StubApp","status":1,"version_name":"1.6.1","version_code":"300161","package_size":40213500,"cdn_url":"http://amp-cdn-test.rqmob.com/apps/2022-04-08/com.karimcastagnini.castlestorm.mg-1649398744.apk","silent_install_key":"b6f819c3c8fe40155be8a43c479749e77d18a70a63f2dcadeb0b8c4fb7b76be02f20b676f3d93ae3dd30da229d72cfae3a69c2c2abb341f85285d093dc5ef0e5","icon_url":"http://amp-cdn-test.rqmob.com/icons/2022-04-08/a7g517tcuv82my975ejxuwnmci0df8fk.png","pi_status":4},"apk_select_info":{"origin_id":20731,"strategy_id":20731},"campaign_id":"202211280012","so_line_id":5,"order_type":1,"campaign_name":"RTB-adexchange-test2"},"start_time":1669628694,"end_time":4102416000,"placement_ids":[1596,1600,1637,1638,1639,1640,1641,1657,1662,2040,2050,2051,2052,2053,2054,2055,2056,2058,2059,2060,2061,2062,2064,2070,2085,2101,2128,2131,2132,2133,2134,2135,2136,2137,2138,2139,2140,2141,2142,2143,2180,2181,2183,2184,2188,2190,2282,2284,2315,2317,2318,2323,2324,2325,2329,2371,2378,2379,2382,2383,2386,2387,2389,2400,2401,2405,2408,2409,2410,2414,2415,2416,2418,2419,2424,2425,2426,2427,2428,2429,2430,2433,2434,2435,2436,2437,2440,2442,2443,2446,2447,2448,2449,2451,2452,2453,2454,2455,2494,2495,2496,2497,2498,2499,2500,2501,2510,2511,2514,2521,2522,2524,2532,2562,2563,2564,2565,2566,2568,2569,2570,2571,2572,2578,2579,2580,2583,2584,2585,2586,2587,2588,2589,2590,2592,2595,2596,2597,2601,2633,2635,2636,2661,2662,2663,2664,2665,2666,2667,2668,2669,2670,2671,2672,2673,2675,2676,2680,2682,2683,2684,2685,2686,2687,2688,2689,2690,2691,2693,2696,2698,2699,2700,2701,2703,2704,2708,2709,2712,2714,2715,2716,2718,2722,2723,2726,2727,2728,2742,2744,2747,2748,2751,2754,2755,2756,2759,2763,2774,2775,2776,2782,2783,2785,2786,2794,2795,2799,2800,2802,2804,2805,2808,2809,2810,2813,2815,2817,2818,2819,2820,2822,2823,2824,2825,2827,2833,2835,2836,2841,2856,2871,2873,2874,2875,2876,2878,2896,2897,2904,2917,2943,2963,2964,2965,2966,2991,2992,3000]},"auto_download":1,"silently_install":true}}]}],"call_ma_success":true,"log_data":{"user_app_info":{}}}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_ua_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"retreival_ad_info":[{"pos_id":2588,"ad_info":[{"scoring":{"local_id":6033,"ecpm":10000000},"ad_info":{"ad_id":294698,"update_time":1679644431,"creative_info_map":{"38743":{"creative_id":38743,"creative_type":5,"format_id":27,"format_width":600,"format_height":500,"packed_json":"{\\\"ad_animation_type_video\\\":0,\\\"animation_type\\\":0,\\\"countdown\\\":0,\\\"creative_id\\\":38743,\\\"creative_ver\\\":0,\\\"ext_info\\\":\\\"{\\\\\\\"need_cover\\\\\\\":false}\\\",\\\"height\\\":500,\\\"icon_url\\\":\\\"\\\",\\\"image_urls\\\":[\\\"http://static-dev.rqmob.com/test/sa/20210710/2491fa9c5d3f9a6b89294870ed78f3ff__FILE_CM____FILE_CM480_720____WEBP__.jpg\\\"],\\\"landing_page\\\":\\\"{\\\\\\\"items\\\\\\\":[{\\\\\\\"t\\\\\\\":\\\\\\\"image\\\\\\\",\\\\\\\"url\\\\\\\":\\\\\\\"http://static-dev.rqmob.com/test/sa/20220120/8943df731f9a2c022cc98a538b238765__FILE_CM____FILE_CM480_720____WEBP__.jpg\\\\\\\",\\\\\\\"h\\\\\\\":1518,\\\\\\\"w\\\\\\\":720,\\\\\\\"__tmpl_type\\\\\\\":\\\\\\\"Image\\\\\\\",\\\\\\\"img_material_id\\\\\\\":2409,\\\\\\\"video_material_id\\\\\\\":0},{\\\\\\\"t\\\\\\\":\\\\\\\"divider\\\\\\\",\\\\\\\"h\\\\\\\":109,\\\\\\\"w\\\\\\\":720,\\\\\\\"img_material_id\\\\\\\":0,\\\\\\\"video_material_id\\\\\\\":0},{\\\\\\\"t\\\\\\\":\\\\\\\"main_button\\\\\\\",\\\\\\\"h\\\\\\\":-1,\\\\\\\"w\\\\\\\":720,\\\\\\\"img_material_id\\\\\\\":0,\\\\\\\"video_material_id\\\\\\\":0,\\\\\\\"txt\\\\\\\":\\\\\\\"DOWNLOAD\\\\\\\",\\\\\\\"s\\\\\\\":1}],\\\\\\\"page_model\\\\\\\":5,\\\\\\\"title\\\\\\\":\\\\\\\"\\\\\\\"}\\\",\\\"layout_type\\\":-1,\\\"product_type\\\":0,\\\"scale_type\\\":1,\\\"support_jump\\\":false,\\\"title\\\":\\\"Recommended for you\\\",\\\"type\\\":5,\\\"width\\\":600}"}},"track_urls":"|||||||||||||||","dsp_type":1009,"landing_page_info":[{"id":32253,"template_id":15,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1646298168435_bOyQZYBH.html?__id__=32253&amp;__name__=&amp;isAuto=false&amp;isToNa=%7BisToNa%7D&amp;version=510a4c858444ead6ef05310792f3a863"},{"id":32254,"template_id":17,"template_type":2,"landing_page_url":"http://dbvh19g70uon7.cloudfront.net/page/1646298168756_scjcHYkj.html?__id__=32254&amp;__name__=&amp;isAuto=false&amp;isToNa=%7BisToNa%7D&amp;version=2b266578e61d9a7faee02554b1a02ff2"},{"id":32255,"template_id":5,"template_type":1},{"id":32256,"template_id":9,"template_type":1}],"ad_set":{"id":4669,"campaign":{"id":5020,"objective":2,"billing_type":4,"attribution_tracker_url":"https://app.appsflyer.com/com.shopee.ph?pid=shareit_int&amp;c=topappsicon&amp;af_click_lookback=7d","bid_price":10000000,"ad_app_pkg_name":"com.tdcm.trueidapp","start_time":1646236800,"end_time":4102416000,"ad_app_id":20500,"attribution_pi":1,"pack_attribution_tracker_url":"http://ping-test.rqmob.com/click?advid=29&amp;campid=1995523&amp;platform=adshonor&amp;pkg=com.tdcm.trueidapp&amp;is_offline=1&amp;requestid={requestid}&amp;gaid={gaid}&amp;device_id={device_id}&amp;android_id={android_id}&amp;beyla_id={beyla_id}&amp;imei={imei}&amp;imsi={imsi}&amp;c={c}&amp;site_id={site_id}&amp;c_id={c_id}&amp;adset={adset}&amp;adset_id={adset_id}&amp;ad={ad}&amp;ad_id={ad_id}&amp;ad_type={ad_type}&amp;cost_model={cost_model}&amp;cost_currency={cost_currency}&amp;cost_value={cost_value}&amp;remote_ip={remote_ip}&amp;country_code={country_code}&amp;ip={ip}&amp;os_version={os_version}&amp;placement={placement}&amp;other_category={other_category}&amp;ext_info={ext_info}&amp;uagent={uagent}&amp;sid={sid}&amp;channel_pkg={channel_pkg}&amp;channel_pkg_ver={channel_pkg_ver}&amp;app_type={app_type}&amp;app_id={app_id}&amp;real_attrplat={real_attrplat}&amp;adpos_id={adpos_id}&amp;midas_camp_id={midas_camp_id}&amp;midas_traffic_type={midas_traffic_type}&amp;_app_id={amp_app_id}&amp;is_pre_install={is_pre_install}&amp;cut_type={cut_type}&amp;package_type={package_type}&amp;t_l=0","attribution_platform":"appsflyer","app_install_id":11019416,"is_online_ad":true,"ad_app":{"id":20500,"package_name":"com.tdcm.trueidapp","app_name":"TrueID","status":1,"version_name":"3.6.0","version_code":"21030600","package_size":73505501,"gp_url":"https://play.google.com/store/apps/details?id=com.tdcm.trueidapp&amp;hl=en","silent_install_key":"08de8e24b9c401f92fe2cf00aa1eebfc81be00ebc14c409e57f1c34d5428a3a9afb67956598dbe6f22a492ffda51896e70e53dbeaa6eed58d8f01f4e5c75dfc4","icon_url":"http://pmp-cdn.modengs.net/apps/2022-02-25/mf89fcpbu6xk1gx6l99welog58wgf3ci-1645785550.webp","pi_status":2},"apk_select_info":{"origin_id":20500,"strategy_id":20500,"strategy_type":"6","traffic_key":"unknown","is_filter_sapk":true,"bucket_type":3},"app_category":"ENTERTAINMENT","ad_logo":1,"campaign_id":"202203030009","order_type":1,"campaign_name":"RTB-浼樺寲-yym"},"start_time":1646236800,"end_time":4102416000,"placement_ids":[1756,2652,2725,2906,2439,2444,2474,2485,2556,2640,2695,2865,2921,2954,125,176,437,691,696,697,822,823,824,2127,2180,2230,2231,2232,2233,2234,2235,2245,2251,2252,2272,2273,2274,2275,2377,2382,2387,2468,2469,2470,2476,2477,2490,2502,2503,2504,2508,2537,2550,2571,2588,2623,2625,2642,2730,2741,2745,2746]},"auto_download":1,"silently_install":true,"priority_weight":10000}}]}],"call_ma_success":true}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_oem_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82","bidid":"c54d2b02-b85d-4588-b693-82b116b7acbc","seatbid":[{"bid":[{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82","impid":"1","price":0.05,"burl":"https://test-track.lacunads.com/oem_win?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1&auction_price=${AUCTION_PRICE}","bundle":"com.remitly.androidapp","adid":"111111","crid":"aaaaaa","trackUrls":[{"eventType":2,"url":"http://trident.bigpinggo.com/click?advid=30&campid=1503722&platform=A_API_LacunaDSP&pkg=com.remitly.androidapp&clickid=6e0bbc2c-6974-4719-a5fd-91b280317b82&offer_name=com.remitly.androidapp_Global&creative_id={creative_id}&creative_name={creative_name}&affiliate_id={affiliate_id}&idfa={idfa}&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=Mozilla%2F5.0+%28Linux%3B+Android+8.0.0%3B+SM-J810G+Build%2FR16NW%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F89.0.4389.90+Mobile+Safari%2F537.36&os_platform=Android&os_version=8.0&device_model=MHA-L29&ip=130.0.36.18&country_code=US&ext={ext}&rta_token={rta_token}&san_campid=1000545&adset_id=15392&c_cto=0"},{"eventType":1,"url":"https://test-track.lacunads.com/oem_imp?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1&auction_price=${AUCTION_PRICE}"},{"eventType":2,"url":"https://test-track.lacunads.com/oem_clk?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1"}],"app_info":{"title":"Remitly: Send Money & Transfer","desc":"","icon":"http://pmp-cdn.modengs.net/plutus/client/mr.png","link_url":"https://play.google.com/store/apps/details?id=com.remitly.androidapp&hl=en","bundle":"com.remitly.androidapp"}}]},{"bid":[{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82","impid":"2","price":0.05,"burl":"https://test-track.lacunads.com/oem_win?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1&auction_price=${AUCTION_PRICE}","bundle":"com.abcdefg","adid":"222222","crid":"bbbbbb","trackUrls":[{"eventType":2,"url":"http://trident.bigpinggo.com/click?advid=30&campid=1503722&platform=A_API_LacunaDSP&pkg=com.remitly.androidapp&clickid=6e0bbc2c-6974-4719-a5fd-91b280317b82&offer_name=com.remitly.androidapp_Global&creative_id={creative_id}&creative_name={creative_name}&affiliate_id={affiliate_id}&idfa={idfa}&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=Mozilla%2F5.0+%28Linux%3B+Android+8.0.0%3B+SM-J810G+Build%2FR16NW%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F89.0.4389.90+Mobile+Safari%2F537.36&os_platform=Android&os_version=8.0&device_model=MHA-L29&ip=130.0.36.18&country_code=US&ext={ext}&rta_token={rta_token}&san_campid=1000545&adset_id=15392&c_cto=0"},{"eventType":1,"url":"https://test-track.lacunads.com/oem_imp?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1&auction_price=${AUCTION_PRICE}"},{"eventType":2,"url":"https://test-track.lacunads.com/oem_clk?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1"}],"app_info":{"title":"Remitly: Send Money & Transfer","desc":"","icon":"http://pmp-cdn.modengs.net/plutus/client/mr.png","link_url":"https://play.google.com/store/apps/details?id=com.remitly.androidapp&hl=en","bundle":"com.abcdefg"}}]},{"bid":[{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82","impid":"3","price":0.1,"burl":"https://test-track.lacunads.com/oem_win?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1&auction_price=${AUCTION_PRICE}","bundle":"com.higklmn","adid":"333333","crid":"ccccccc","trackUrls":[{"eventType":2,"url":"http://trident.bigpinggo.com/click?advid=30&campid=1503722&platform=A_API_LacunaDSP&pkg=com.remitly.androidapp&clickid=6e0bbc2c-6974-4719-a5fd-91b280317b82&offer_name=com.remitly.androidapp_Global&creative_id={creative_id}&creative_name={creative_name}&affiliate_id={affiliate_id}&idfa={idfa}&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=Mozilla%2F5.0+%28Linux%3B+Android+8.0.0%3B+SM-J810G+Build%2FR16NW%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F89.0.4389.90+Mobile+Safari%2F537.36&os_platform=Android&os_version=8.0&device_model=MHA-L29&ip=130.0.36.18&country_code=US&ext={ext}&rta_token={rta_token}&san_campid=1000545&adset_id=15392&c_cto=0"},{"eventType":1,"url":"https://test-track.lacunads.com/oem_imp?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1&auction_price=${AUCTION_PRICE}"},{"eventType":2,"url":"https://test-track.lacunads.com/oem_clk?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1"}],"app_info":{"title":"Remitly: Send Money & Transfer","desc":"","icon":"http://pmp-cdn.modengs.net/plutus/client/mr.png","link_url":"https://play.google.com/store/apps/details?id=com.remitly.androidapp&hl=en","bundle":"com.higklmn"}}]},{"bid":[{"id":"6e0bbc2c-6974-4719-a5fd-91b280317b82","impid":"4","price":0.2,"burl":"https://test-track.lacunads.com/oem_win?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1&auction_price=${AUCTION_PRICE}","bundle":"com.opqrst","adid":"444444","crid":"dddddd","trackUrls":[{"eventType":2,"url":"http://trident.bigpinggo.com/click?advid=30&campid=1503722&platform=A_API_LacunaDSP&pkg=com.remitly.androidapp&clickid=6e0bbc2c-6974-4719-a5fd-91b280317b82&offer_name=com.remitly.androidapp_Global&creative_id={creative_id}&creative_name={creative_name}&affiliate_id={affiliate_id}&idfa={idfa}&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=Mozilla%2F5.0+%28Linux%3B+Android+8.0.0%3B+SM-J810G+Build%2FR16NW%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F89.0.4389.90+Mobile+Safari%2F537.36&os_platform=Android&os_version=8.0&device_model=MHA-L29&ip=130.0.36.18&country_code=US&ext={ext}&rta_token={rta_token}&san_campid=1000545&adset_id=15392&c_cto=0"},{"eventType":1,"url":"https://test-track.lacunads.com/oem_imp?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1&auction_price=${AUCTION_PRICE}"},{"eventType":2,"url":"https://test-track.lacunads.com/oem_clk?viewid=5dVDtn7ZpVxJrfuUu3ixP_wwq08ZUzWYTR80_zOBOuMEmlpdWkoAbMhykF8eBZ0syf9w2AAmt5HzG8IYpiWnNkNAYt2GnTssGZlQz3YOI8FLQw_kvdE0K3reJ6iDkHLsdP8VMzGWcEVUEnfslW3erogHPOqeMwXMNe83PYicFrbi73pGugVblGqC2jAxUX6TReZds3W9RsdRxmnlHtq9GJmCb2Pi0V9A0wzVSg2bhGpMNaV_BTCfzJGA43zIoNErlf2msAuyBWcWce4Z52p3CtpxXvXLbVN5nV4yi2W7Z10h6j3GjfEhH5eCxx5o3JWOK-Yh7o_90wqWb6Vg1Pc_cEN4Ixf2WHXmOWe3fnitAVWfLVoftGA8hIxgUZEueSPi9ZqaxBnu&enc=1"}],"app_info":{"title":"Remitly: Send Money & Transfer","desc":"","icon":"http://pmp-cdn.modengs.net/plutus/client/mr.png","link_url":"https://play.google.com/store/apps/details?id=com.remitly.androidapp&hl=en","bundle":"com.opqrst"}}]}],"cur":"USD"}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	//topon
	r.HandleFunc("/mock/868/get_topon_banner_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"b7cda340-665b-175a-f49b-cb99d049b8df",
							"seatbid":[
								{
									"bid":[
										{
											"id":"649e371fd6899d00090e6353",
											"impid":"1",
											"price":0.1325,
											"burl":"https://example.com/adx/tracking?d=eyJUeXBlIjoiYWR4QnVybCIsImRhdGVfdGltZSI6IjIwMjMtMDYtMzAgMDk6NTk6NTkiLCJ0aW1lIjoxNjg4MDkwMzk5LCJwdWJsaXNoZXJfaWQiOjEzNDI1LCJmb3JtYXQiOjIsInBsX2lkX2ludCI6MzM0NzQ4LCJhcHBfaWRfaW50Ijo2MzE5MiwicGxfaWQiOiJiNjQ1OWU3MmUyOGZiOSIsImFwcF9pZCI6ImE2NDU5ZTcyMGU1MDY1IiwiYXBwX3ZuIjoiIiwic2RrX3ZlciI6IiIsInVwaWQiOiIiLCJ0X2dfaWQiOjAsImdyb19pZCI6MCwiZ2FpZCI6Ijk2NTRmYmUxLTE4NzgtNGI2YS05YTgyLTcwYWU0MWJiNGRlYiIsInJlcXVlc3RfaWQiOiJiN2NkYTM0MC02NjViLTE3NWEtZjQ5Yi1jYjk5ZDA0OWI4ZGYiLCJvZmZlcl9maXJtX2lkIjo2NiwiZHNwX2lkIjo3NCwiYmlkX2lkIjoiNjQ5ZTM3MWZkNjg5OWQwMDA5MGU2MzUzIiwiY2xpZW50X2lwIjoiMTcyLjU4LjEyMy4xODgiLCJnZW9fc2hvcnQiOiJVUyIsImRzcF9jdXJyZW5jeSI6IlVTRCIsImN1cnJlbmN5X3JhdGUiOjEsInJldmVudWVfY3V0X3JhdGUiOjAuMjUsInBrZyI6ImNvbS5uYXZpYXBwIiwicGxhdGZvcm0iOjEsImJ1bmRsZV9pbiI6ImNvbS52ZW5hbWF4Lm1heCIsImJ1bmRsZV9maW5hbCI6ImNvbS52ZW5hbWF4Lm1heCIsInRtYXgiOjI5NjIsImJpZGZsb29yIjowLjA3NjY2NjY2NjY2NjY2NjY3LCJydGJfYWRfdHlwZSI6IjIiLCJpbnN0bCI6MCwiZXhwaXJlIjozMCwiYWRvbWFpbiI6Im5hdmkuY29tIiwiY2F0IjoiIiwidWEiOiJNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgMTI7IG1vdG8gZyBwb3dlciAoMjAyMikgQnVpbGQvUzNSUVMzMi4yMC00Mi0xMC03OyB3dikgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi80LjAgQ2hyb21lLzExNC4wLjU3MzUuMTMwIE1vYmlsZSBTYWZhcmkvNTM3LjM2Iiwib3JpZ2luYWxfcHJpY2Vfc3RyIjoiQlUzTHN3M2hmVEVySy11d2RWTmFtblJBbWVSejFGUDY2RGhlSjAtejNIST0iLCJkaXNjb3VudF91c2RfcHJpY2Vfc3RyIjoiUXZhd3JTRDZmWS1HMThiN2hDS2V2dz09IiwiZHNwX2FwcCI6IjEyMTY1NDA4MiIsImRzcF9wbGFjZW1lbnQiOiJiYW5uZXIiLCJ2ZXIiOiJvdXRTZWFBcGkiLCJhc19idXJsIjoxfQ==",
											"adm":"<img src=\"https://example.com/adx/tracking?d=eyJUeXBlIjoiYWR4SW1wIiwiZGF0ZV90aW1lIjoiMjAyMy0wNi0zMCAwOTo1OTo1OSIsInRpbWUiOjE2ODgwOTAzOTksInB1Ymxpc2hlcl9pZCI6MTM0MjUsImZvcm1hdCI6MiwicGxfaWRfaW50IjozMzQ3NDgsImFwcF9pZF9pbnQiOjYzMTkyLCJwbF9pZCI6ImI2NDU5ZTcyZTI4ZmI5IiwiYXBwX2lkIjoiYTY0NTllNzIwZTUwNjUiLCJhcHBfdm4iOiIiLCJzZGtfdmVyIjoiIiwidXBpZCI6IiIsInRfZ19pZCI6MCwiZ3JvX2lkIjowLCJnYWlkIjoiOTY1NGZiZTEtMTg3OC00YjZhLTlhODItNzBhZTQxYmI0ZGViIiwicmVxdWVzdF9pZCI6ImI3Y2RhMzQwLTY2NWItMTc1YS1mNDliLWNiOTlkMDQ5YjhkZiIsIm9mZmVyX2Zpcm1faWQiOjY2LCJkc3BfaWQiOjc0LCJiaWRfaWQiOiI2NDllMzcxZmQ2ODk5ZDAwMDkwZTYzNTMiLCJjbGllbnRfaXAiOiIxNzIuNTguMTIzLjE4OCIsImdlb19zaG9ydCI6IlVTIiwiZHNwX2N1cnJlbmN5IjoiVVNEIiwiY3VycmVuY3lfcmF0ZSI6MSwicmV2ZW51ZV9jdXRfcmF0ZSI6MC4yNSwicGtnIjoiY29tLm5hdmlhcHAiLCJwbGF0Zm9ybSI6MSwiYnVuZGxlX2luIjoiY29tLnZlbmFtYXgubWF4IiwiYnVuZGxlX2ZpbmFsIjoiY29tLnZlbmFtYXgubWF4IiwidG1heCI6Mjk2MiwiYmlkZmxvb3IiOjAuMDc2NjY2NjY2NjY2NjY2NjcsInJ0Yl9hZF90eXBlIjoiMiIsImluc3RsIjowLCJleHBpcmUiOjMwLCJhZG9tYWluIjoibmF2aS5jb20iLCJjYXQiOiIiLCJ1YSI6Ik1vemlsbGEvNS4wIChMaW51eDsgQW5kcm9pZCAxMjsgbW90byBnIHBvd2VyICgyMDIyKSBCdWlsZC9TM1JRUzMyLjIwLTQyLTEwLTc7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvMTE0LjAuNTczNS4xMzAgTW9iaWxlIFNhZmFyaS81MzcuMzYiLCJvcmlnaW5hbF9wcmljZV9zdHIiOiJCVTNMc3czaGZURXJLLXV3ZFZOYW1uUkFtZVJ6MUZQNjZEaGVKMC16M0hJPSIsImRpc2NvdW50X3VzZF9wcmljZV9zdHIiOiJRdmF3clNENmZZLUcxOGI3aENLZXZ3PT0iLCJkc3BfYXBwIjoiMTIxNjU0MDgyIiwiZHNwX3BsYWNlbWVudCI6ImJhbm5lciIsInZlciI6Im91dFNlYUFwaSJ9\" width=\"1\" height=\"1\" hidden />\n<script>\nvar _tpClickUrl=\"https://tk.example.com/adx/tracking?d=eyJUeXBlIjoiYWR4Q2xpY2siLCJkYXRlX3RpbWUiOiIyMDIzLTA2LTMwIDA5OjU5OjU5IiwidGltZSI6MTY4ODA5MDM5OSwicHVibGlzaGVyX2lkIjoxMzQyNSwiZm9ybWF0IjoyLCJwbF9pZF9pbnQiOjMzNDc0OCwiYXBwX2lkX2ludCI6NjMxOTIsInBsX2lkIjoiYjY0NTllNzJlMjhmYjkiLCJhcHBfaWQiOiJhNjQ1OWU3MjBlNTA2NSIsImFwcF92biI6IiIsInNka192ZXIiOiIiLCJ1cGlkIjoiIiwidF9nX2lkIjowLCJncm9faWQiOjAsImdhaWQiOiI5NjU0ZmJlMS0xODc4LTRiNmEtOWE4Mi03MGFlNDFiYjRkZWIiLCJyZXF1ZXN0X2lkIjoiYjdjZGEzNDAtNjY1Yi0xNzVhLWY0OWItY2I5OWQwNDliOGRmIiwib2ZmZXJfZmlybV9pZCI6NjYsImRzcF9pZCI6NzQsImJpZF9pZCI6IjY0OWUzNzFmZDY4OTlkMDAwOTBlNjM1MyIsImNsaWVudF9pcCI6IjE3Mi41OC4xMjMuMTg4IiwiZ2VvX3Nob3J0IjoiVVMiLCJkc3BfY3VycmVuY3kiOiJVU0QiLCJjdXJyZW5jeV9yYXRlIjoxLCJyZXZlbnVlX2N1dF9yYXRlIjowLjI1LCJwa2ciOiJjb20ubmF2aWFwcCIsInBsYXRmb3JtIjoxLCJidW5kbGVfaW4iOiJjb20udmVuYW1heC5tYXgiLCJidW5kbGVfZmluYWwiOiJjb20udmVuYW1heC5tYXgiLCJ0bWF4IjoyOTYyLCJiaWRmbG9vciI6MC4wNzY2NjY2NjY2NjY2NjY2NywicnRiX2FkX3R5cGUiOiIyIiwiaW5zdGwiOjAsImV4cGlyZSI6MzAsImFkb21haW4iOiJuYXZpLmNvbSIsImNhdCI6IiIsInVhIjoiTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDEyOyBtb3RvIGcgcG93ZXIgKDIwMjIpIEJ1aWxkL1MzUlFTMzIuMjAtNDItMTAtNzsgd3YpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIFZlcnNpb24vNC4wIENocm9tZS8xMTQuMC41NzM1LjEzMCBNb2JpbGUgU2FmYXJpLzUzNy4zNiIsIm9yaWdpbmFsX3ByaWNlX3N0ciI6IkJVM0xzdzNoZlRFckstdXdkVk5hbW5SQW1lUnoxRlA2NkRoZUowLXozSEk9IiwiZGlzY291bnRfdXNkX3ByaWNlX3N0ciI6IlF2YXdyU0Q2ZlktRzE4YjdoQ0tldnc9PSIsImRzcF9hcHAiOiIxMjE2NTQwODIiLCJkc3BfcGxhY2VtZW50IjoiYmFubmVyIiwidmVyIjoib3V0U2VhQXBpIn0=\";\nfunction _tpHttpRequest(t){(new Image).src=t}var isSupportTouch=!1;document.body.addEventListener(\"touchstart\",function(t){isSupportTouch=!0,_tpHttpRequest(_tpClickUrl)}),document.body.addEventListener(\"click\",function(t){isSupportTouch||_tpHttpRequest(_tpClickUrl)});\n</script><div><a href=\"\" target=\"_blank\"><img src=\"https://cdn.example.com/Navi_creatives/Navi_300x250.jpg\" width=\"300\" height=\"250\" border=\"0\"></a></div> <img src=\"https://tracking.example.com/impression?apri=0.17666666&cpt=1&csid=topon-banner-android&wdid=1&exp=3600&bidid=&seatid=&dcamid=952805&dcreid=Navi_300x250&dadid=649e371fd6899d00090e6353&adm=navi.com&bundle=com.naviapp&rat=3&w=300&h=250&appid=121654082&appver=1.0.8&pm=1&devt=4&brand=Motorola&model=Moto+G+Power+%282022%29&netwk=6&os=android&osv=12&idfa=9654fbe1-1878-4b6a-9a82-70ae41bb4deb&aid=&imei=&ua=Mozilla%2F5.0+%28Linux%3B+Android+12%3B+moto+g+power+%282022%29+Build%2FS3RQS32.20-42-10-7%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F114.0.5735.130+Mobile+Safari%2F537.36&ip=172.58.123.188&cny=US&rts=1688090399&iswn=0&nexj=&mode=1&cbn=com.venamax.max&reqid=649e371fd6899d00090e6353&impid=imp_649e371fd6899d00090e6353&camid=1-952805&creid=1-Navi_300x250&spoid=645f10379951d&adxid=15&pritype=1&creqid=b7cda340-665b-175a-f49b-cb99d049b8df&dprir=0.01&dencap=NjQ5ZTM3MWZkNjg5OWQwMNiqirLFbJup9SW7cg%3D%3D&cpr=0.176667&resds=\" width=\"1\" height=\"1\" style=\"display:none;\">",
											"adid":"649e371fd6899d00090e6353",
											"adomain":[
												"navi.com"
											],
											"bundle":"com.naviapp",
											"cid":"952805",
											"crid":"Navi_300x250",
											"qagmediarating":3,
											"w":300,
											"h":250,
											"exp":3600
										}
									]
								}
							],
							"bidid":"649e371fd6899d00090e6353",
							"cur":"USD"
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_topon_native_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"cif1n63mgcdu3gjds0ag",
								"seatbid":[
									{
										"bid":[
											{
												"id":"ffde0a9c-a40c-4615-8e4b-4ed9b3f73f00",
												"impid":"1",
												"price":0.04549500000000001,
												"nurl":"https://track.example.com/win?viewid=5th4nW7N_Vceu_iStyezex0BVAHEvZmPjSCmWccTOzOHR3sfIv6zXS0A6kLESukQAWXbOaTIEozXRoPeIHe6s8MIAKWpxOR-eeBBcihn2fACGxQ1T0iRkigkkTbeXJtaGqeWec4_KgtJlWmnBRXOMCfrAW-TMvipnACCnW__Ovv_-PqLgl8CRit40Ksr1Zet4zrni3oNZBTDbZErb8GLLVWP_kG3rlBTFrrZGQwzphfaWmrPNiSUE5w7UMBiUQq72a4bPIVq8deLt8q24rCbwodwAi4QONXhrCOCI_NyzWhfXjzlwvQmz1em5EKN6UG4uF7j5gQ-UU71rNASz-C8QC7Hk4eIu30IYJoDsYSzFqdEX7E1Dtr9lZJP7r3ihoWjk3q0v7NwWYqzkiQGw0uKLY_5DXZj_hmVqQ7OnUL2Jcxoa3346aRb6nsOY_vrP3rWjzQLwly4t_UJX75LKYQ-5JTIS03MQana9LGYPO3GXGrL_rar1kPfAw9V5ZX2vT88uXT7yPg8wdWBUnmD64M56G5M_Zq5aCVlKSo77CdPBksG0EG5ab4d3UjNLj_YUSTHF6IICQlnjeE5BE_ni2Y8EpxVC9kHrD3LPZ5k-K6xwA-_pkN4Psj1rp4dW9cYkTUlY-bdo9pR1pcn0pezuCNdJ_N3Sbz9m9zjdC5qriC1zaQ5Tk8qKMje2xXuoSGT_kin-Nife0dy2CyoBlGI_xlFsuo1bg7hsc7W9aR5G-ncNbkTLJLHKkYzdK2nuBVgRdLG4dBNXoL2J84vjwyLKHBen7EfQwQBIP1rJcT8xb2mzf0K_vn9072CbJwW644UNJlSFl5sW-hwfE4LylzUCrb50suT0l0SYq_8zTMa-FdWCtWbdjZLR3HZr99BJteH62mDIFXrtXDy0h8hWN6fq_7F-e16A4chFQPu_Eyh62GnoN5N34vvV6EAmfr26ZSMDrhwO-SJpEZklUGqFymgoTwQJJ7dz1dK3D5pW5qNo03iw4DNCCF8Dq7I2akL4ghmbfc4FJcwSMwdFtT-DkVlz-6ObUai3tJ6G4ZTV9W3jGyT6d1Pexhj_-xARlw0duFmH-h9ixfPf7UWVBdgvFwpPPzpfmSTP5CgoeJd2fk64FQ_UP9bL5RMkKeL6hulBvmXbY8OLpSU8Tvw5akB2U38DHPe9RmaY_KxfS9VRGI5Y9oBcirjgpbN1TllLsVka0TVrwTGrEiQhn_OMvDCXFKdTNXb6rF32pEdPWuqnWltp9CzkoCQqWI6HP_IwbNFx2ua0RzS8MX_FpJPMqPqWtkojYxGWRdVfOQUUL4RE-iErcHfoU1TJ3EX3zzIE5mUgU2aQBRziYEOkXqP0z7h5jWOGM5LEHy3iA3cAaoP4OwQPNsUWURD0Plsm2u3HwpUfTy_kb2UxrgnAjbga68T7uEnj6F-URLVy_L1rp6KUd8JgSc1F2BvKlxoL51UiPdwQo65SfboFBsbl3fp1E8GtaGgaCkPvu-fFn_fyl7MMMiu_Fok47FheXeTmU80AmL9pdPk0_TOm_RGcnyusqOXP2XXkrhUUn-MiEz72XzpbFCht-YCE_F_nX9zKSoP5SHm1uyV6oKLMRSrDv0OIO8XfsswgrdyOyLcaO-sqMdrn7Qi7w2aJWbiVjTNx9wqQhefatdZO4mKkiRN0qRKI7I8O-7ZPFueWw1yYfA7F5ZaKl6F6heWWcoOwc0eueSCGYMz4CPdfdLeqQHdIFdwy77f9Pp26XXVtRC-nrhdZ2kdUFUzc-OLOnES7GKDK5G2o5e2UmrDfHEQrgS0KzzBoK9feT9BiECOXCDMgE0MW3b0Xy4mILHQjxlfCrmAC92effrAnqJakVwBNYKWb5supoiG9gybefzWzrl09hoFLEjizJlurZ2cDvirML13CtUIfBZl7w==&enc=1&auction_price=0.06066",
												"burl":"https://tk.example.com/adx/tracking?d=eyJUeXBlIjoiYWR4QnVybCIsImRhdGVfdGltZSI6IjIwMjMtMDYtMzAgMDg6MDI6MzMiLCJ0aW1lIjoxNjg4MDgzMzUzLCJwdWJsaXNoZXJfaWQiOjExOTU4LCJmb3JtYXQiOjAsInBsX2lkX2ludCI6Mjk0Mzk4LCJhcHBfaWRfaW50Ijo1NjA2MywicGxfaWQiOiJiNjNiNTRjMDZlZmFlOCIsImFwcF9pZCI6ImE2M2I1NGJlODRlMzFmIiwiYXBwX3ZuIjoiIiwic2RrX3ZlciI6IiIsInVwaWQiOiIiLCJ0X2dfaWQiOjAsImdyb19pZCI6MCwiZ2FpZCI6Ijc4YzRjODRlLTMxY2UtNGRkZi1hNjg5LTI1NjYxMmEyOTcwMCIsInJlcXVlc3RfaWQiOiJjaWYxbjYzbWdjZHUzZ2pkczBhZyIsIm9mZmVyX2Zpcm1faWQiOjY2LCJkc3BfaWQiOjU0LCJiaWRfaWQiOiJmZmRlMGE5Yy1hNDBjLTQ2MTUtOGU0Yi00ZWQ5YjNmNzNmMDAiLCJjbGllbnRfaXAiOiI5MS4xMjIuMjUwLjEyNCIsImdlb19zaG9ydCI6IlJVIiwiZHNwX2N1cnJlbmN5IjoiVVNEIiwiY3VycmVuY3lfcmF0ZSI6MSwicmV2ZW51ZV9jdXRfcmF0ZSI6MC4yNSwicGxhdGZvcm0iOjEsImJ1bmRsZV9pbiI6ImNvbS5iYW5kYWdhbWVzLm1wdXp6bGUuZ3AiLCJidW5kbGVfZmluYWwiOiJjb20uYmFuZGFnYW1lcy5tcHV6emxlLmdwIiwidG1heCI6NzUwLCJiaWRmbG9vciI6MC4wMTMzMzMzMzMzMzMzMzMzMzQsInJ0Yl9hZF90eXBlIjoiMCIsImluc3RsIjowLCJleHBpcmUiOjcyMCwiYWRvbWFpbiI6ImJ1bWFsdC5ydSIsImNhdCI6IiIsInVhIjoiTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDEzOyBTTS1YMjA1IEJ1aWxkL1RQMUEuMjIwNjI0LjAxNDsgd3YpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIFZlcnNpb24vNC4wIENocm9tZS8xMTQuMC41NzM1LjEzMSBTYWZhcmkvNTM3LjM2Iiwib3JpZ2luYWxfcHJpY2Vfc3RyIjoibnFjV0YyeFlXTGdzdHZhbXdWcU9OaERHUzVOakRnbUh4bnBzVmFoLUJ2OD0iLCJkaXNjb3VudF91c2RfcHJpY2Vfc3RyIjoiaUpENlkxcEZDYVRnaG5JN1FXRlM4ZVR5QWpKTG5ySDdfNkVVeDVubHZ6ND0iLCJkc3BfYXBwIjoiOTE4ODUiLCJkc3BfcGxhY2VtZW50IjoiU2hhcmVpdG5hdGl2ZSIsInZlciI6Im91dFNlYUFwaSIsImFzX2J1cmwiOjF9",
												"adm":"{\"native\":{\"ver\":\"1.1\",\"assets\":[{\"id\":3,\"title\":{\"text\":\"Упаковочная клейкая…\"}},{\"id\":4,\"data\":{\"value\":\"крафт-лента SafeLenta - антикражная! Окончательное высыхание 2 мин, после ленту нельзя отклеить без повреждения коробки!\"}},{\"id\":6,\"data\":{\"value\":\"Перейти\"}},{\"id\":7,\"img\":{\"url\":\"https://avatars.mds.example.net/get-direct/4380796/LWcxT57CUHebAR8xH-2SYg/wx1080\",\"w\":1080,\"h\":607}},{\"id\":1,\"img\":{\"url\":\"https://avatars.mds.example.net/get-direct/4380796/LWcxT57CUHebAR8xH-2SYg/wx1080\",\"w\":1080,\"h\":607}},{\"id\":2,\"video\":{\"vasttag\":\"\\u003c?xml version=\\\"1.0\\\" encoding=\\\"utf-8\\\"?\\u003e\\u003cVAST version=\\\"3.0\\\"\\u003e\\u003cAd id=\\\"15001\\\"\\u003e\\u003cInLine\\u003e\\u003cAdSystem\\u003eexample\\u003c/AdSystem\\u003e\\u003cAdTitle\\u003e\\u003c![CDATA[Упаковочная клейкая крафт-лента SafeLenta - антикражная!]]\\u003e\\u003c/AdTitle\\u003e\\u003cDescription\\u003e\\u003c![CDATA[Окончательное высыхание 2 мин, после ленту нельзя отклеить без повреждения коробки!]]\\u003e\\u003c/Description\\u003e\\u003cExtensions\\u003e\\u003cExtension type=\\\"example_ad_info\\\"\\u003e\\u003c![CDATA[{\\\"assets\\\":{},\\\"text\\\":{}}]]\\u003e\\u003c/Extension\\u003e\\u003c/Extensions\\u003e\\u003cSurvey\\u003ehttps://www.example.ru\\u003c/Survey\\u003e\\u003cCreatives\\u003e\\u003cCreative id=\\\"Linear\\\"\\u003e\\u003cLinear\\u003e\\u003cDuration\\u003e00:00:05\\u003c/Duration\\u003e\\u003cTrackingEvents\\u003e\\u003cTracking event=\\\"start\\\"\\u003e\\u003c![CDATA[https://an.example.ru/count/WPyejI_zOoVX2LaY0OqC0BDMLnUqThmUQeVhTx8wZbxdwtMuNcVtELaFvX86q76gI4DWOTNWWHOKE7bII5NCu62N7RuGTF3UtBam0dYpVTPEF4SYWMIjWMHD-aa2ejx8OFDderKgbKgbZ0yNDZk29EUgmXOXeuhOw9v8NXoNko1-JyuSp3G7CNC3LyniJAjXFr5tXeCIUJVuGu362zwOV67YH8aXL8sMIREjqcXfcko438rkTeHa2BpwnXcgB0F8mWIWJy1WzerR9kMoKAO4Nc7GFfKe2VDzELqPNk0wdTBfa_gqer7eivxcKMfzhLRnsXfatxKLpuFme5owC-COuwHlcQvqx0oPATEe_ES9ibjJPAMVO4X_QE84WqmPBPCTm93A4mfSJo7hAUXqZBY2LjTXl4-kA0et04mmXEbUG5J15HujHr1DFLzsrjpzRuoXAnipeHizdRrvKELAfL9PzUgmjUa2bFynW0_KBxue108E_B0ch-qw4O6Jx2F5v1aW9BrB_he-HQRRJ4p4dQY9ga0vLKFKOAG-3bfcwcIKFDq8~2=WLaejI_zOoVX2LdU05qA05DIHNwdBC8EEO2TCVJvNExisNxdpXlkV9N_F45sQugiz1CCwGzDd63GSPn1PEfgjt01G-bilTRBMTdiqKeFBGawEvnFNmhGEfya2z0wdo0Bq99WefBpXU67DVHQpELl5oGGGgwRFMTam6l8nVbymuIpzsEAZTcVRRw-x6o7pB0HpqlCR4ohOQUzUetkPaT8Xi2hRomT5yviwHyLoZD_UbinvEQJAuh23G0J366wlkoioRMWziNUof0VXW6YvOD56FA_1EnzxxBEwMkzquwVzM58PfbUUZi8iMgZTrz7YEVbpHduiWSkBYjFzcXWJtT-7A892bldlq23sRW_lwLrNJySpCRzUQttei84yifdbrlaibs44mPLI0KG_3SFd3HraTPeLaaKrSlRqbLCMLJ88Dq0mAPLJDbi3mGH8Q9sk6f80mfokk-EgK8XU0C0~2]]\\u003e\\u003c/Tracking\\u003e\\u003c/TrackingEvents\\u003e\\u003cMediaFiles\\u003e\\u003c/MediaFiles\\u003e\\u003cVideoClicks\\u003e\\u003cClickThrough id=\\\"AW\\\"\\u003e\\u003c![CDATA[https://an.example.ru/count/WUiejI_zOoVX2Ld70UqE01ERPj3dSxcpP_kTEs-uyxsWjUFrK3jSvJzFaDsQ8Ylz14FwGnCd67GSPn2PkjejNC1GEbklzN9MDdkqqaDBWiuEvzEN0hJE9ya2jCudo0Aq99Yev7nX-64DFuJz1wy7i_V0bOR0mgh10oqeSFAaaAgOmS5gCPU_NP102BbkzfoH0QyX5pXxFwRuiZERtGUwRG0pDt1to0mRyRp4pCOkU6FoXOaJ9OPGDLhQgzVeRdQ4P0Yy-iQDMoRbib2c15vXq3gTqkcJ-hIZKUYpdkPHQdsjLl7Q6cJVjHNFW_2WNBepunZZf6-PhdJi35awBgG7CAJov4zlAuRSV795KRW1O88XJFS2eWgky6WvWchg-RAp9jUNdJLt2sfVxAqC-M4O85e-M8GX_ql0tlSIiFUv1upzVLddzBNUQSVFUgSTJBCFGpBBCNHbOoPqOmSXLcruNjwSWIzxuEAoycIF1lPqvyUq1uTeWiBM-I_GO3RkpszfNTSF1xDnVrxhlQWmWRposUKMkMpNpOCgLogLogfNXwrD8r5-5LzQIxxueX085VmX89owmTLSltwUrnqRMzNQ9BmA2GKCtiOPEYJyICZGZzTBJV1u6o_xepmJkNC8uWauUK0QDY11VUrL_YkH5W00~2]]\\u003e\\u003c/ClickThrough\\u003e\\u003c/VideoClicks\\u003e\\u003cIcons\\u003e\\u003cIcon program=\\\"icon\\\" height=\\\"80\\\" width=\\\"80\\\" xPosition=\\\"left\\\" yPosition=\\\"bottom\\\"\\u003e\\u003cStaticResource creativeType=\\\"image/png\\\"\\u003e\\u003c![CDATA[https://yastatic.net/pcode/static/icons/icon-globe-192x192.png]]\\u003e\\u003c/StaticResource\\u003e\\u003c/Icon\\u003e\\u003cIcon program=\\\"AdChoices\\\" height=\\\"15\\\" width=\\\"15\\\" xPosition=\\\"right\\\" yPosition=\\\"bottom\\\"\\u003e\\u003cStaticResource creativeType=\\\"image/png\\\"\\u003e\\u003c![CDATA[https://yastatic.net/pcode/static/icons/icon-globe-64x64.png]]\\u003e\\u003c/StaticResource\\u003e\\u003cIconClicks\\u003e\\u003cIconClickThrough\\u003e\\u003c![CDATA[https://example.ru/tune/adv]]\\u003e\\u003c/IconClickThrough\\u003e\\u003c/IconClicks\\u003e\\u003c/Icon\\u003e\\u003c/Icons\\u003e\\u003c/Linear\\u003e\\u003c/Creative\\u003e\\u003c/Creatives\\u003e\\u003c/InLine\\u003e\\u003c/Ad\\u003e\\u003c/VAST\\u003e\"}}],\"link\":{\"url\":\"https://an.example.ru/count/WUiejI_zOoVX2Ld70UqE01ERPj3dSxcpP_kTEs-uyxsWjUFrK3jSvJzFaDsQ8Ylz14FwGnCd67GSPn2PkjejNC1GEbklzN9MDdkqqaDBWiuEvzEN0hJE9ya2jCudo0Aq99Yev7nX-64DFuJz1wy7i_V0bOR0mgh10oqeSFAaaAgOmS5gCPU_NP102BbkzfoH0QyX5pXxFwRuiZERtGUwRG0pDt1to0mRyRp4pCOkU6FoXOaJ9OPGDLhQgzVeRdQ4P0Yy-iQDMoRbib2c15vXq3gTqkcJ-hIZKUYpdkPHQdsjLl7Q6cJVjHNFW_2WNBepunZZf6-PhdJi35awBgG7CAJov4zlAuRSV795KRW1O88XJFS2eWgky6WvWchg-RAp9jUNdJLt2sfVxAqC-M4O85e-M8GX_ql0tlSIiFUv1upzVLddzBNUQSVFUgSTJBCFGpBBCNHbOoPqOmSXLcruNjwSWIzxuEAoycIF1lPqvyUq1uTeWiBM-I_GO3RkpszfNTSF1xDnVrxhlQWmWRposUKMkMpNpOCgLogLogfNXwrD8r5-5LzQIxxueX085VmX89owmTLSltwUrnqRMzNQ9BmA2GKCtiOPEYJyICZGZzTBJV1u6o_xepmJkNC8uWauUK0QDY11VUrL_YkH5W00~2\",\"clicktrackers\":[\"https://track.example.com/clk?viewid=5th4nW7N_Vceu_iStyezex0BVAHEvZmPjSCmWccTOzOHR3sfIv6zXS0A6kLESukQAWXbOaTIEozXRoPeIHe6s8MIAKWpxOR-eeBBcihn2fACGxQ1T0iRkigkkTbeXJtaGqeWec4_KgtJlWmnBRXOMCfrAW-TMvipnACCnW__Ovv_-PqLgl8CRit40Ksr1Zet4zrni3oNZBTDbZErb8GLLVWP_kG3rlBTFrrZGQwzphfaWmrPNiSUE5w7UMBiUQq72a4bPIVq8deLt8q24rCbwodwAi4QONXhrCOCI_NyzWhfXjzlwvQmz1em5EKN6UG4uF7j5gQ-UU71rNASz-C8QC7Hk4eIu30IYJoDsYSzFqdEX7E1Dtr9lZJP7r3ihoWjk3q0v7NwWYqzkiQGw0uKLY_5DXZj_hmVqQ7OnUL2Jcxoa3346aRb6nsOY_vrP3rWjzQLwly4t_UJX75LKYQ-5JTIS03MQana9LGYPO3GXGrL_rar1kPfAw9V5ZX2vT88uXT7yPg8wdWBUnmD64M56G5M_Zq5aCVlKSo77CdPBksG0EG5ab4d3UjNLj_YUSTHF6IICQlnjeE5BE_ni2Y8EpxVC9kHrD3LPZ5k-K6xwA-_pkN4Psj1rp4dW9cYkTUlY-bdo9pR1pcn0pezuCNdJ_N3Sbz9m9zjdC5qriC1zaQ5Tk8qKMje2xXuoSGT_kin-Nife0dy2CyoBlGI_xlFsuo1bg7hsc7W9aR5G-ncNbkTLJLHKkYzdK2nuBVgRdLG4dBNXoL2J84vjwyLKHBen7EfQwQBIP1rJcT8xb2mzf0K_vn9072CbJwW644UNJlSFl5sW-hwfE4LylzUCrb50suT0l0SYq_8zTMa-FdWCtWbdjZLR3HZr99BJteH62mDIFXrtXDy0h8hWN6fq_7F-e16A4chFQPu_Eyh62GnoN5N34vvV6EAmfr26ZSMDrhwO-SJpEZklUGqFymgoTwQJJ7dz1dK3D5pW5qNo03iw4DNCCF8Dq7I2akL4ghmbfc4FJcwSMwdFtT-DkVlz-6ObUai3tJ6G4ZTV9W3jGyT6d1Pexhj_-xARlw0duFmH-h9ixfPf7UWVBdgvFwpPPzpfmSTP5CgoeJd2fk64FQ_UP9bL5RMkKeL6hulBvmXbY8OLpSU8Tvw5akB2U38DHPe9RmaY_KxfS9VRGI5Y9oBcirjgpbN1TllLsVka0TVrwTGrEiQhn_OMvDCXFKdTNXb6rF32pEdPWuqnWltp9CzkoCQqWI6HP_IwbNFx2ua0RzS8MX_FpJPMqPqWtkojYxGWRdVfOQUUL4RE-iErcHfoU1TJ3EX3zzIE5mUgU2aQBRziYEOkXqP0z7h5jWOGM5LEHy3iA3cAaoP4OwQPNsUWURD0Plsm2u3HwpUfTy_kb2UxrgnAjbga68T7uEnj6F-URLVy_L1rp6KUd8JgSc1F2BvKlxoL51UiPdwQo65SfboFBsbl3fp1E8GtaGgaCkPvu-fFn_fyl7MMMiu_Fok47FheXeTmU80AmL9pdPk0_TOm_RGcnyusqOXP2XXkrhUUn-MiEz72XzpbFCht-YCE_F_nX9zKSoP5SHm1uyV6oKLMRSrDv0OIO8XfsswgrdyOyLcaO-sqMdrn7Qi7w2aJWbiVjTNx9wqQhefatdZO4mKkiRN0qRKI7I8O-7ZPFueWw1yYfA7F5ZaKl6F6heWWcoOwc0eueSCGYMz4CPdfdLeqQHdIFdwy77f9Pp26XXVtRC-nrhdZ2kdUFUzc-OLOnES7GKDK5G2o5e2UmrDfHEQrgS0KzzBoK9feT9BiECOXCDMgE0MW3b0Xy4mILHQjxlfCrmAC92effrAnqJakVwBNYKWb5supoiG9gybefzWzrl09hoFLEjizJlurZ2cDvirML13CtUIfBZl7w==\\u0026enc=1\",\"https://tk.example.com/adx/tracking?d=eyJUeXBlIjoiYWR4Q2xpY2siLCJkYXRlX3RpbWUiOiIyMDIzLTA2LTMwIDA4OjAyOjMzIiwidGltZSI6MTY4ODA4MzM1MywicHVibGlzaGVyX2lkIjoxMTk1OCwiZm9ybWF0IjowLCJwbF9pZF9pbnQiOjI5NDM5OCwiYXBwX2lkX2ludCI6NTYwNjMsInBsX2lkIjoiYjYzYjU0YzA2ZWZhZTgiLCJhcHBfaWQiOiJhNjNiNTRiZTg0ZTMxZiIsImFwcF92biI6IiIsInNka192ZXIiOiIiLCJ1cGlkIjoiIiwidF9nX2lkIjowLCJncm9faWQiOjAsImdhaWQiOiI3OGM0Yzg0ZS0zMWNlLTRkZGYtYTY4OS0yNTY2MTJhMjk3MDAiLCJyZXF1ZXN0X2lkIjoiY2lmMW42M21nY2R1M2dqZHMwYWciLCJvZmZlcl9maXJtX2lkIjo2NiwiZHNwX2lkIjo1NCwiYmlkX2lkIjoiZmZkZTBhOWMtYTQwYy00NjE1LThlNGItNGVkOWIzZjczZjAwIiwiY2xpZW50X2lwIjoiOTEuMTIyLjI1MC4xMjQiLCJnZW9fc2hvcnQiOiJSVSIsImRzcF9jdXJyZW5jeSI6IlVTRCIsImN1cnJlbmN5X3JhdGUiOjEsInJldmVudWVfY3V0X3JhdGUiOjAuMjUsInBsYXRmb3JtIjoxLCJidW5kbGVfaW4iOiJjb20uYmFuZGFnYW1lcy5tcHV6emxlLmdwIiwiYnVuZGxlX2ZpbmFsIjoiY29tLmJhbmRhZ2FtZXMubXB1enpsZS5ncCIsInRtYXgiOjc1MCwiYmlkZmxvb3IiOjAuMDEzMzMzMzMzMzMzMzMzMzM0LCJydGJfYWRfdHlwZSI6IjAiLCJpbnN0bCI6MCwiZXhwaXJlIjo3MjAsImFkb21haW4iOiJidW1hbHQucnUiLCJjYXQiOiIiLCJ1YSI6Ik1vemlsbGEvNS4wIChMaW51eDsgQW5kcm9pZCAxMzsgU00tWDIwNSBCdWlsZC9UUDFBLjIyMDYyNC4wMTQ7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvMTE0LjAuNTczNS4xMzEgU2FmYXJpLzUzNy4zNiIsIm9yaWdpbmFsX3ByaWNlX3N0ciI6Im5xY1dGMnhZV0xnc3R2YW13VnFPTmhER1M1TmpEZ21IeG5wc1ZhaC1Cdjg9IiwiZGlzY291bnRfdXNkX3ByaWNlX3N0ciI6ImlKRDZZMXBGQ2FUZ2huSTdRV0ZTOGVUeUFqSkxuckg3XzZFVXg1bmx2ejQ9IiwiZHNwX2FwcCI6IjkxODg1IiwiZHNwX3BsYWNlbWVudCI6IlNoYXJlaXRuYXRpdmUiLCJ2ZXIiOiJvdXRTZWFBcGkifQ==\"]},\"imptrackers\":[\"https://track.example.com/imp?viewid=5th4nW7N_Vceu_iStyezex0BVAHEvZmPjSCmWccTOzOHR3sfIv6zXS0A6kLESukQAWXbOaTIEozXRoPeIHe6s8MIAKWpxOR-eeBBcihn2fACGxQ1T0iRkigkkTbeXJtaGqeWec4_KgtJlWmnBRXOMCfrAW-TMvipnACCnW__Ovv_-PqLgl8CRit40Ksr1Zet4zrni3oNZBTDbZErb8GLLVWP_kG3rlBTFrrZGQwzphfaWmrPNiSUE5w7UMBiUQq72a4bPIVq8deLt8q24rCbwodwAi4QONXhrCOCI_NyzWhfXjzlwvQmz1em5EKN6UG4uF7j5gQ-UU71rNASz-C8QC7Hk4eIu30IYJoDsYSzFqdEX7E1Dtr9lZJP7r3ihoWjk3q0v7NwWYqzkiQGw0uKLY_5DXZj_hmVqQ7OnUL2Jcxoa3346aRb6nsOY_vrP3rWjzQLwly4t_UJX75LKYQ-5JTIS03MQana9LGYPO3GXGrL_rar1kPfAw9V5ZX2vT88uXT7yPg8wdWBUnmD64M56G5M_Zq5aCVlKSo77CdPBksG0EG5ab4d3UjNLj_YUSTHF6IICQlnjeE5BE_ni2Y8EpxVC9kHrD3LPZ5k-K6xwA-_pkN4Psj1rp4dW9cYkTUlY-bdo9pR1pcn0pezuCNdJ_N3Sbz9m9zjdC5qriC1zaQ5Tk8qKMje2xXuoSGT_kin-Nife0dy2CyoBlGI_xlFsuo1bg7hsc7W9aR5G-ncNbkTLJLHKkYzdK2nuBVgRdLG4dBNXoL2J84vjwyLKHBen7EfQwQBIP1rJcT8xb2mzf0K_vn9072CbJwW644UNJlSFl5sW-hwfE4LylzUCrb50suT0l0SYq_8zTMa-FdWCtWbdjZLR3HZr99BJteH62mDIFXrtXDy0h8hWN6fq_7F-e16A4chFQPu_Eyh62GnoN5N34vvV6EAmfr26ZSMDrhwO-SJpEZklUGqFymgoTwQJJ7dz1dK3D5pW5qNo03iw4DNCCF8Dq7I2akL4ghmbfc4FJcwSMwdFtT-DkVlz-6ObUai3tJ6G4ZTV9W3jGyT6d1Pexhj_-xARlw0duFmH-h9ixfPf7UWVBdgvFwpPPzpfmSTP5CgoeJd2fk64FQ_UP9bL5RMkKeL6hulBvmXbY8OLpSU8Tvw5akB2U38DHPe9RmaY_KxfS9VRGI5Y9oBcirjgpbN1TllLsVka0TVrwTGrEiQhn_OMvDCXFKdTNXb6rF32pEdPWuqnWltp9CzkoCQqWI6HP_IwbNFx2ua0RzS8MX_FpJPMqPqWtkojYxGWRdVfOQUUL4RE-iErcHfoU1TJ3EX3zzIE5mUgU2aQBRziYEOkXqP0z7h5jWOGM5LEHy3iA3cAaoP4OwQPNsUWURD0Plsm2u3HwpUfTy_kb2UxrgnAjbga68T7uEnj6F-URLVy_L1rp6KUd8JgSc1F2BvKlxoL51UiPdwQo65SfboFBsbl3fp1E8GtaGgaCkPvu-fFn_fyl7MMMiu_Fok47FheXeTmU80AmL9pdPk0_TOm_RGcnyusqOXP2XXkrhUUn-MiEz72XzpbFCht-YCE_F_nX9zKSoP5SHm1uyV6oKLMRSrDv0OIO8XfsswgrdyOyLcaO-sqMdrn7Qi7w2aJWbiVjTNx9wqQhefatdZO4mKkiRN0qRKI7I8O-7ZPFueWw1yYfA7F5ZaKl6F6heWWcoOwc0eueSCGYMz4CPdfdLeqQHdIFdwy77f9Pp26XXVtRC-nrhdZ2kdUFUzc-OLOnES7GKDK5G2o5e2UmrDfHEQrgS0KzzBoK9feT9BiECOXCDMgE0MW3b0Xy4mILHQjxlfCrmAC92effrAnqJakVwBNYKWb5supoiG9gybefzWzrl09hoFLEjizJlurZ2cDvirML13CtUIfBZl7w==\\u0026enc=1\\u0026auction_price=0.06066\",\"https://adrta.com/i?clid=srit\\u0026paid=srit\\u0026avid=10\\u0026caid=1_181682529\\u0026kv1=0x0\\u0026kv11=cif1n63mgcdu3gjds0ag\\u0026kv15=RUS\\u0026kv16=0.000000\\u0026kv17=0.000000\\u0026kv18=com.bandagames.mpuzzle.gp\\u0026kv19=78c4c84e-31ce-4ddf-a689-256612a29700\\u0026kv24=Mobile_InApp_Native\\u0026kv26=Android\\u0026kv4=91.122.250.124\\u0026plid=1_181682529_0x0\\u0026publisherId=6\\u0026siteId=0\",\"https://track.lacunads.com/win?viewid=5th4nW7N_Vceu_iStyezex0BVAHEvZmPjSCmWccTOzOHR3sfIv6zXS0A6kLESukQAWXbOaTIEozXRoPeIHe6s8MIAKWpxOR-eeBBcihn2fACGxQ1T0iRkigkkTbeXJtaGqeWec4_KgtJlWmnBRXOMCfrAW-TMvipnACCnW__Ovv_-PqLgl8CRit40Ksr1Zet4zrni3oNZBTDbZErb8GLLVWP_kG3rlBTFrrZGQwzphfaWmrPNiSUE5w7UMBiUQq72a4bPIVq8deLt8q24rCbwodwAi4QONXhrCOCI_NyzWhfXjzlwvQmz1em5EKN6UG4uF7j5gQ-UU71rNASz-C8QC7Hk4eIu30IYJoDsYSzFqdEX7E1Dtr9lZJP7r3ihoWjk3q0v7NwWYqzkiQGw0uKLY_5DXZj_hmVqQ7OnUL2Jcxoa3346aRb6nsOY_vrP3rWjzQLwly4t_UJX75LKYQ-5JTIS03MQana9LGYPO3GXGrL_rar1kPfAw9V5ZX2vT88uXT7yPg8wdWBUnmD64M56G5M_Zq5aCVlKSo77CdPBksG0EG5ab4d3UjNLj_YUSTHF6IICQlnjeE5BE_ni2Y8EpxVC9kHrD3LPZ5k-K6xwA-_pkN4Psj1rp4dW9cYkTUlY-bdo9pR1pcn0pezuCNdJ_N3Sbz9m9zjdC5qriC1zaQ5Tk8qKMje2xXuoSGT_kin-Nife0dy2CyoBlGI_xlFsuo1bg7hsc7W9aR5G-ncNbkTLJLHKkYzdK2nuBVgRdLG4dBNXoL2J84vjwyLKHBen7EfQwQBIP1rJcT8xb2mzf0K_vn9072CbJwW644UNJlSFl5sW-hwfE4LylzUCrb50suT0l0SYq_8zTMa-FdWCtWbdjZLR3HZr99BJteH62mDIFXrtXDy0h8hWN6fq_7F-e16A4chFQPu_Eyh62GnoN5N34vvV6EAmfr26ZSMDrhwO-SJpEZklUGqFymgoTwQJJ7dz1dK3D5pW5qNo03iw4DNCCF8Dq7I2akL4ghmbfc4FJcwSMwdFtT-DkVlz-6ObUai3tJ6G4ZTV9W3jGyT6d1Pexhj_-xARlw0duFmH-h9ixfPf7UWVBdgvFwpPPzpfmSTP5CgoeJd2fk64FQ_UP9bL5RMkKeL6hulBvmXbY8OLpSU8Tvw5akB2U38DHPe9RmaY_KxfS9VRGI5Y9oBcirjgpbN1TllLsVka0TVrwTGrEiQhn_OMvDCXFKdTNXb6rF32pEdPWuqnWltp9CzkoCQqWI6HP_IwbNFx2ua0RzS8MX_FpJPMqPqWtkojYxGWRdVfOQUUL4RE-iErcHfoU1TJ3EX3zzIE5mUgU2aQBRziYEOkXqP0z7h5jWOGM5LEHy3iA3cAaoP4OwQPNsUWURD0Plsm2u3HwpUfTy_kb2UxrgnAjbga68T7uEnj6F-URLVy_L1rp6KUd8JgSc1F2BvKlxoL51UiPdwQo65SfboFBsbl3fp1E8GtaGgaCkPvu-fFn_fyl7MMMiu_Fok47FheXeTmU80AmL9pdPk0_TOm_RGcnyusqOXP2XXkrhUUn-MiEz72XzpbFCht-YCE_F_nX9zKSoP5SHm1uyV6oKLMRSrDv0OIO8XfsswgrdyOyLcaO-sqMdrn7Qi7w2aJWbiVjTNx9wqQhefatdZO4mKkiRN0qRKI7I8O-7ZPFueWw1yYfA7F5ZaKl6F6heWWcoOwc0eueSCGYMz4CPdfdLeqQHdIFdwy77f9Pp26XXVtRC-nrhdZ2kdUFUzc-OLOnES7GKDK5G2o5e2UmrDfHEQrgS0KzzBoK9feT9BiECOXCDMgE0MW3b0Xy4mILHQjxlfCrmAC92effrAnqJakVwBNYKWb5supoiG9gybefzWzrl09hoFLEjizJlurZ2cDvirML13CtUIfBZl7w==\\u0026enc=1\\u0026auction_price=0.06066\",\"https://an.example.ru/count/WPyejI_zOoVX2LaY0OqC0BDMLnUqThmUQeVhTx8wZbxdwtMuNcVtELaFvX86q76gI4DWOTNWWHOKE7bII5NCu62N7RuGTF3UtBam0dYpVTPEF4SYWMIjWMHD-aa2ejx8OFDderKgbKgbZ0yNDZk29EUgmXOXeuhOw9v8NXoNko1-JyuSp3G7CNC3LyniJAjXFr5tXeCIUJVuGu362zwOV67YH8aXL8sMIREjqcXfcko438rkTeHa2BpwnXcgB0F8mWIWJy1WzerR9kMoKAO4Nc7GFfKe2VDzELqPNk0wdTBfa_gqer7eivxcKMfzhLRnsXfatxKLpuFme5owC-COuwHlcQvqx0oPATEe_ES9ibjJPAMVO4X_QE84WqmPBPCTm93A4mfSJo7hAUXqZBY2LjTXl4-kA0et04mmXEbUG5J15HujHr1DFLzsrjpzRuoXAnipeHizdRrvKELAfL9PzUgmjUa2bFynW0_KBxue108E_B0ch-qw4O6Jx2F5v1aW9BrB_he-HQRRJ4p4dQY9ga0vLKFKOAG-3bfcwcIKFDq8~2=WLaejI_zOoVX2LdU05qA05DIHNwdBC8EEO2TCVJvNExisNxdpXlkV9N_F45sQugiz1CCwGzDd63GSPn1PEfgjt01G-bilTRBMTdiqKeFBGawEvnFNmhGEfya2z0wdo0Bq99WefBpXU67DVHQpELl5oGGGgwRFMTam6l8nVbymuIpzsEAZTcVRRw-x6o7pB0HpqlCR4ohOQUzUetkPaT8Xi2hRomT5yviwHyLoZD_UbinvEQJAuh23G0J366wlkoioRMWziNUof0VXW6YvOD56FA_1EnzxxBEwMkzquwVzM58PfbUUZi8iMgZTrz7YEVbpHduiWSkBYjFzcXWJtT-7A892bldlq23sRW_lwLrNJySpCRzUQttei84yifdbrlaibs44mPLI0KG_3SFd3HraTPeLaaKrSlRqbLCMLJ88Dq0mAPLJDbi3mGH8Q9sk6f80mfokk-EgK8XU0C0~2\",\"https://tk.toponadss.com/adx/tracking?d=eyJUeXBlIjoiYWR4SW1wIiwiZGF0ZV90aW1lIjoiMjAyMy0wNi0zMCAwODowMjozMyIsInRpbWUiOjE2ODgwODMzNTMsInB1Ymxpc2hlcl9pZCI6MTE5NTgsImZvcm1hdCI6MCwicGxfaWRfaW50IjoyOTQzOTgsImFwcF9pZF9pbnQiOjU2MDYzLCJwbF9pZCI6ImI2M2I1NGMwNmVmYWU4IiwiYXBwX2lkIjoiYTYzYjU0YmU4NGUzMWYiLCJhcHBfdm4iOiIiLCJzZGtfdmVyIjoiIiwidXBpZCI6IiIsInRfZ19pZCI6MCwiZ3JvX2lkIjowLCJnYWlkIjoiNzhjNGM4NGUtMzFjZS00ZGRmLWE2ODktMjU2NjEyYTI5NzAwIiwicmVxdWVzdF9pZCI6ImNpZjFuNjNtZ2NkdTNnamRzMGFnIiwib2ZmZXJfZmlybV9pZCI6NjYsImRzcF9pZCI6NTQsImJpZF9pZCI6ImZmZGUwYTljLWE0MGMtNDYxNS04ZTRiLTRlZDliM2Y3M2YwMCIsImNsaWVudF9pcCI6IjkxLjEyMi4yNTAuMTI0IiwiZ2VvX3Nob3J0IjoiUlUiLCJkc3BfY3VycmVuY3kiOiJVU0QiLCJjdXJyZW5jeV9yYXRlIjoxLCJyZXZlbnVlX2N1dF9yYXRlIjowLjI1LCJwbGF0Zm9ybSI6MSwiYnVuZGxlX2luIjoiY29tLmJhbmRhZ2FtZXMubXB1enpsZS5ncCIsImJ1bmRsZV9maW5hbCI6ImNvbS5iYW5kYWdhbWVzLm1wdXp6bGUuZ3AiLCJ0bWF4Ijo3NTAsImJpZGZsb29yIjowLjAxMzMzMzMzMzMzMzMzMzMzNCwicnRiX2FkX3R5cGUiOiIwIiwiaW5zdGwiOjAsImV4cGlyZSI6NzIwLCJhZG9tYWluIjoiYnVtYWx0LnJ1IiwiY2F0IjoiIiwidWEiOiJNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgMTM7IFNNLVgyMDUgQnVpbGQvVFAxQS4yMjA2MjQuMDE0OyB3dikgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi80LjAgQ2hyb21lLzExNC4wLjU3MzUuMTMxIFNhZmFyaS81MzcuMzYiLCJvcmlnaW5hbF9wcmljZV9zdHIiOiJucWNXRjJ4WVdMZ3N0dmFtd1ZxT05oREdTNU5qRGdtSHhucHNWYWgtQnY4PSIsImRpc2NvdW50X3VzZF9wcmljZV9zdHIiOiJpSkQ2WTFwRkNhVGdobkk3UVdGUzhlVHlBakpMbnJIN182RVV4NW5sdno0PSIsImRzcF9hcHAiOiI5MTg4NSIsImRzcF9wbGFjZW1lbnQiOiJTaGFyZWl0bmF0aXZlIiwidmVyIjoib3V0U2VhQXBpIn0=\"]}}",
												"adid":"yabs.NzIwNTc2MDcyODc4MDM1NjE=",
												"adomain":[
													"bumalt.ru"
												],
												"cid":"1_181682529",
												"crid":"1_181682529_0x0"
											}
										]
									}
								],
								"bidid":"ffde0a9c-a40c-4615-8e4b-4ed9b3f73f00",
								"cur":"USD"
							}
						
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_topon_video_ad", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
							"id":"1688093999515857804_3360e976-16f2-11ee-9c84-3f5983d44daa_500352637_topon-rv",
							"seatbid":[
								{
									"bid":[
										{
											"id":"e2f050c3ad8b439d882bf0525ca60304",
											"impid":"1",
											"price":0.9573550000000001,
											"nurl":"https://bs-metadsp.yandex.ru/ssp_notifications/1NXRO-D30de200000000U9nJtF_oJEMaeCOdXk5lPpnTeB3Ko5BVTLuA0n1umaH2zQivSsEtRPPaI6K4YcU6lzrm3YHU2kBLNWKIhOmWSYQZe14BI2V3GaQhFcOmQoLZqpz6C6iZY7fm44ZiPHHOoiPmbB6lbO7XAYD8-2uZWuG0mr_6oGVrdUSoSG15bZ8ri3Cmh6Mo01fubka_45nb0gbyodJM-YxhAoQ12JSz8yFVCZ3SdsJ87F2VbU4l4ol88SmW1_GoCm79gSnSG78kiop86SnD3YeDd090ovbt9iOpd6JMv-N_r5LQQtAL6Um0PZcGuSokVyAvVp8xEF63dTp5ec9XP6bXQ6bYf9jiQ96gQv9YRg1hQP2SfckGOcbcPc1iQPwIPhc3XVa02zD9bWzBRBVCsPHrv2RBh9aWyD7-8SiPXiNgzm8UsBzb0Vds0yl-ovlFBRzqvS0Jnbyi2kQ9gPP9QGQffeOfAPQ9YSOfIIQMXgQ6gGQc5icMfcb9JQpyoVcvUWMP5VbvcSbwvUL52rXPJh86J-K6E-y2wuA6XA-ieC3vnSffQ6bWR6XgQc7Pm7Rqn_se3hO6bXbiLmu7Bs2xVsVTX2dFrxV6h6I3EUS0cyS9DkP7RBoLBivMOKQPkOw5FbdcBzd0vZ9yiFy9Pqraxpf-_VF34-_iTx0pdc2NUS3-PShprU_lhrh1z8FzflW2ZebF15Pmx9h3ju6ZwI_x3Wx6MnmOSVCK6ACDzhsSe22OoWcAfmMSh01LfzDq?ssp-notification-type=1&ssp-request-id=e2f050c3ad8b439d882bf0525ca60304&ssp-cur-price=1.5016&ssp-cur=USD",
											"burl":"https://tk.toponadss.com/adx/tracking?d=eyJUeXBlIjoiYWR4QnVybCIsImRhdGVfdGltZSI6IjIwMjMtMDYtMzAgMTA6NTk6NTkiLCJ0aW1lIjoxNjg4MDkzOTk5LCJwdWJsaXNoZXJfaWQiOjEzMjY2LCJmb3JtYXQiOjEsInBsX2lkX2ludCI6MzMyODIxLCJhcHBfaWRfaW50Ijo2MjM3OCwicGxfaWQiOiJiNjQ1NDcyMTQ4NDM5NiIsImFwcF9pZCI6ImE2NDQ3OGFjYjRjY2NkIiwiYXBwX3ZuIjoiIiwic2RrX3ZlciI6IiIsInVwaWQiOiIiLCJ0X2dfaWQiOjAsImdyb19pZCI6MCwiZ2FpZCI6IjgyOTgxOTQyLWMxOGUtNGQ3MC05MGNmLTBkOTY2MDM5N2RmOSIsInJlcXVlc3RfaWQiOiIxNjg4MDkzOTk5NTE1ODU3ODA0XzMzNjBlOTc2LTE2ZjItMTFlZS05Yzg0LTNmNTk4M2Q0NGRhYV81MDAzNTI2MzdfdG9wb24tcnYiLCJvZmZlcl9maXJtX2lkIjo2NiwiZHNwX2lkIjo3MCwiYmlkX2lkIjoiZTJmMDUwYzNhZDhiNDM5ZDg4MmJmMDUyNWNhNjAzMDQiLCJjbGllbnRfaXAiOiI1LjE4MS42Mi4xMjciLCJnZW9fc2hvcnQiOiJSVSIsImRzcF9jdXJyZW5jeSI6IlVTRCIsImN1cnJlbmN5X3JhdGUiOjEsInJldmVudWVfY3V0X3JhdGUiOjAuMTUsInBsYXRmb3JtIjoxLCJidW5kbGVfaW4iOiJjb20ucGl4b25pYy53d3IiLCJidW5kbGVfZmluYWwiOiJjb20ucGl4b25pYy53d3IiLCJ0bWF4IjoxNDY3LCJiaWRmbG9vciI6MC4wMTMwMDAwMDAwMDAwMDAwMDEsInJ0Yl9hZF90eXBlIjoiNSIsImluc3RsIjowLCJleHBpcmUiOjEyMCwiYWRvbWFpbiI6Im9zdHJvdm9rLnJ1IiwiY2F0IjoiSUFCMTksSUFCMjIsSUFCMjEiLCJ1YSI6Ik1vemlsbGEvNS4wIChMaW51eDsgQW5kcm9pZCAxMjsgQVNVU19JMDA1REEgQnVpbGQvU0tRMS4yMTA4MjEuMDAxOyB3dikgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi80LjAgQ2hyb21lLzExNC4wLjU3MzUuMTMwIE1vYmlsZSBTYWZhcmkvNTM3LjM2Iiwib3JpZ2luYWxfcHJpY2Vfc3RyIjoiMW1SWWttRXBEd1A3d1otWjA1RWRjQT09IiwiZGlzY291bnRfdXNkX3ByaWNlX3N0ciI6IjV6bE4zV1ZXcWV0c2lINlliY18wRmh5TkYxNDZMMjNpcVc4TFNpMGVMb2s9IiwiZHNwX2FwcCI6IjIwMDU3NSIsImRzcF9wbGFjZW1lbnQiOiJyZXdhcmRlZCB2aWRlbyIsInZlciI6Im91dFNlYUFwaSIsImFzX2J1cmwiOjF9",
											"adm":"<VAST version=\"3.0\"><Ad id=\"15001\"><InLine><AdSystem><![CDATA[Yandex]]></AdSystem><AdTitle><![CDATA[ostrovok.ru]]></AdTitle><Impression id=\"1c1355a7e7ae4566a160f0194835c61b\"><![CDATA[https://an.yandex.ru/rtbcount/1KBzau560e0200000000U9nJtF_oJEMaeCOdXk5lPpnTeB3Ko5BVTLuA0n1umaH2zQivSsEtRPPaI6K4YcU6lzrm3YHU2kBLNWKIhOmWSYQZe14BI2V3GaQhFcOmQoLZqpz6C6iZY7fm44ZiPHHOoiPmbB6lbO7XB2D8-2gOlCl88270y9Tnya5zvxbC742HfKmDx0mCQvaiWCQUPVeF15UPG2gVCfsrlilwYabWmatFoF2t30pt9nao1_ndPVZBn0fo3ZF86RtCJ41oAZCNa9pBB0jo37EJWqe3fm3G2Smx4-CPJZBhy_9_wYgjDRdA37O0inp8S6RNF-7SFvaTdFZ13kzYKR4mihGmj3Gnqaqsj4ZLDKanDz2rD4XEqpL8iJGpCp2sD4z9izp1mdo0XUaaouSbjblcRCewSfDbramG-6Z_aEKCmsBrUm4FxD-oWFpyWMN_vStdbb-wSk09uoyM1VF4LCkaD8DKKyEK5Cl4HEEK91DBGrF3L0DJYsHBqxIa9bR-vFnSlOBCYloyp6IzylAY1Qmi9rc3ftA3dNU1TK53mbTMKE3yOkKqj3GmDZGrjR1iO3lwutvK1ri3omosAuU35x3TlxDkmfJdwrjZLhB1d7E0pUC46_CZDjvA5sUhCAFCNCT2doppbsnWSna-sF-4ioQoTns__ldXYNVsEzWPJx3BF63_CkNvwlVtLwrW-i5-Ktm1nyGd0YiuTitXsy3Hz9Tz1mVZBGwCk7aAZD46-rvEK11CvGJ5qm9ELW3NLjH3?winprice=1.5016]]></Impression><Impression id=\"fd18bd1cf5d24f6e8d64783116b2d79f\"><![CDATA[https://an.yandex.ru/metacount/1NXRO-D30de200000000U9nJtF_oJEMaeCOdXk5lPpnTeB3Ko5BVTLuA0n1umaH2zQivSsEtRPPaI6K4YcU6lzrm3YHU2kBLNWKIhOmWSYQZe14BI2V3GaQhFcOmQoLZqpz6C6iZY7fm44ZiPHHOoiPmbB6lbO7XAYD8-2uZWuG0mr_6oGVrdUSoSG15bZ8ri3Cmh6Mo01fubka_45nb0gbyodJM-YxhAoQ12JSz8yFVCZ3SdsJ87F2VbU4l4ol88SmW1_GoCm79gSnSG78kiop86SnD3YeDd090ovbt9iOpd6JMv-N_r5LQQtAL6Um0PZcGuSokVyAvVp8xEF63dTp5ec9XP6bXQ6bYf9jiQ96gQv9YRg1hQP2SfckGOcbcPc1iQPwIPhc3XVa02zD9bWzBRBVCsPHrv2RBh9aWyD7-8SiPXiNgzm8UsBzb0Vds0yl-ovlFBRzqvS0Jnbyi2kQ9gPP9QGQffeOfAPQ9YSOfIIQMXgQ6gGQc5icMfcb9JQpyoVcvUWMP5VbvcSbwvUL52rXPJh86J-K6E-y2wuA6XA-ieC3vnSffQ6bWR6XgQc7Pm7Rqn_se3hO6bXbiLmu7Bs2xVsVTX2dFrxV6h6I3EUS0cyS9DkP7RBoLBivMOKQPkOw5FbdcBzd0vZ9yiFy9Pqraxpf-_VF34-_iTx0pdc2NUS3-PShprU_lhrh1z8FzflW2ZebF15Pmx9h3ju6ZwI_x3Wx6MnmOSVCK6ACDzhsSe22OoWcAfmMSh01LfzDq?&ssp-cur-price=1.5016&ssp-cur=USD]]></Impression><Impression id=\"0bc83ffbcbf9441db5f96c67a00c75a5\"><![CDATA[https://rus-trk.ascope.co/imp?cm=1315&sk=203705&af=3&ts=1688093999&id=e2f050c3ad8b439d882bf0525ca60304&pt=YTMwUjLx0zYwZyM2ITMuETPihnJ2EDM14SM9IHa&price=1.1263]]></Impression><Impression><![CDATA[https://tk.toponadss.com/adx/tracking?d=eyJUeXBlIjoiYWR4SW1wIiwiZGF0ZV90aW1lIjoiMjAyMy0wNi0zMCAxMDo1OTo1OSIsInRpbWUiOjE2ODgwOTM5OTksInB1Ymxpc2hlcl9pZCI6MTMyNjYsImZvcm1hdCI6MSwicGxfaWRfaW50IjozMzI4MjEsImFwcF9pZF9pbnQiOjYyMzc4LCJwbF9pZCI6ImI2NDU0NzIxNDg0Mzk2IiwiYXBwX2lkIjoiYTY0NDc4YWNiNGNjY2QiLCJhcHBfdm4iOiIiLCJzZGtfdmVyIjoiIiwidXBpZCI6IiIsInRfZ19pZCI6MCwiZ3JvX2lkIjowLCJnYWlkIjoiODI5ODE5NDItYzE4ZS00ZDcwLTkwY2YtMGQ5NjYwMzk3ZGY5IiwicmVxdWVzdF9pZCI6IjE2ODgwOTM5OTk1MTU4NTc4MDRfMzM2MGU5NzYtMTZmMi0xMWVlLTljODQtM2Y1OTgzZDQ0ZGFhXzUwMDM1MjYzN190b3Bvbi1ydiIsIm9mZmVyX2Zpcm1faWQiOjY2LCJkc3BfaWQiOjcwLCJiaWRfaWQiOiJlMmYwNTBjM2FkOGI0MzlkODgyYmYwNTI1Y2E2MDMwNCIsImNsaWVudF9pcCI6IjUuMTgxLjYyLjEyNyIsImdlb19zaG9ydCI6IlJVIiwiZHNwX2N1cnJlbmN5IjoiVVNEIiwiY3VycmVuY3lfcmF0ZSI6MSwicmV2ZW51ZV9jdXRfcmF0ZSI6MC4xNSwicGxhdGZvcm0iOjEsImJ1bmRsZV9pbiI6ImNvbS5waXhvbmljLnd3ciIsImJ1bmRsZV9maW5hbCI6ImNvbS5waXhvbmljLnd3ciIsInRtYXgiOjE0NjcsImJpZGZsb29yIjowLjAxMzAwMDAwMDAwMDAwMDAwMSwicnRiX2FkX3R5cGUiOiI1IiwiaW5zdGwiOjAsImV4cGlyZSI6MTIwLCJhZG9tYWluIjoib3N0cm92b2sucnUiLCJjYXQiOiJJQUIxOSxJQUIyMixJQUIyMSIsInVhIjoiTW96aWxsYS81LjAgKExpbnV4OyBBbmRyb2lkIDEyOyBBU1VTX0kwMDVEQSBCdWlsZC9TS1ExLjIxMDgyMS4wMDE7IHd2KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBWZXJzaW9uLzQuMCBDaHJvbWUvMTE0LjAuNTczNS4xMzAgTW9iaWxlIFNhZmFyaS81MzcuMzYiLCJvcmlnaW5hbF9wcmljZV9zdHIiOiIxbVJZa21FcER3UDd3Wi1aMDVFZGNBPT0iLCJkaXNjb3VudF91c2RfcHJpY2Vfc3RyIjoiNXpsTjNXVldxZXRzaUg2WWJjXzBGaHlORjE0NkwyM2lxVzhMU2kwZUxvaz0iLCJkc3BfYXBwIjoiMjAwNTc1IiwiZHNwX3BsYWNlbWVudCI6InJld2FyZGVkIHZpZGVvIiwidmVyIjoib3V0U2VhQXBpIn0=]]></Impression><Creatives><Creative id=\"Linear\"><Linear><Duration>00:00:30</Duration><Icons><Icon program=\"icon\" width=\"80\" height=\"80\" xPosition=\"left\" yPosition=\"bottom\" offset=\"0%\" duration=\"00:00:00\"><IconClicks><IconClickThrough></IconClickThrough></IconClicks><StaticResource creativeType=\"image/png\"><![CDATA[https://yastatic.net/pcode/static/icons/icon-globe-192x192.png]]></StaticResource><IFrameResource></IFrameResource></Icon><Icon program=\"AdChoices\" width=\"15\" height=\"15\" xPosition=\"right\" yPosition=\"bottom\" offset=\"0%\" duration=\"00:00:00\"><IconClicks><IconClickThrough><![CDATA[https://yandex.ru/tune/adv]]></IconClickThrough></IconClicks><StaticResource creativeType=\"image/png\"><![CDATA[https://yastatic.net/pcode/static/icons/icon-globe-64x64.png]]></StaticResource><IFrameResource></IFrameResource></Icon></Icons><TrackingEvents><Tracking event=\"start\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=0]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=1]]></Tracking><Tracking event=\"midpoint\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=2]]></Tracking><Tracking event=\"thirdQuartile\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=3]]></Tracking><Tracking event=\"complete\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=4]]></Tracking><Tracking event=\"mute\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=5]]></Tracking><Tracking event=\"unmute\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=6]]></Tracking><Tracking event=\"pause\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=7]]></Tracking><Tracking event=\"resume\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=8]]></Tracking><Tracking event=\"render\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=11]]></Tracking><Tracking event=\"impression\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=13]]></Tracking><Tracking event=\"complete\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=19]]></Tracking><Tracking event=\"progress\" offset=\"00:00:30\"><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=19]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://ad.adriver.ru/cgi-bin/rle.cgi?sid=1&bt=55&ad=760866&pid=3679617&bid=9126814&bn=9126814&exss=&rnd=2144079693]]></Tracking><Tracking event=\"firstQuartile\"><![CDATA[https://an.yandex.ru/count/WSGejI_zOoVX2Lad0RKD06DPOHUqThmUQeVhTx8wZbxdwtMuNcUtRPuSl-zSEe40Df-JzM0kOG0Tfueq11OYz109FM7md41nOmO7k-pmXQ23ruKHm9ljitRYEH8996iD96lIJn8GzKO6cpzlerKgbKgbZ0yNDZk29EUgmXOXeuhOw3vnQgnA4q2w7auSp3G7CNC3LyniJAjXju38s1UflSKhZG8qUaGmu1zPtn0cTZLQ9issIgFrDDaB6HhnGc3hQYiXrfYXYWyWUo40_G86iqyY1sNs1L18aJKCskTV68lgJ8gWamlSwRKgdTxEQ7HQTTRN8EUqTaghU6tDxQ3wYK_jVFhCTsOSOvpqBLCF9kmCUVcbPBojeDCovgLCe_AVBLZQLeLAvPKO6aJmH8j1feoMsT8osH6FN9HGu5PWef2bG46dOyGLjBeEutqAoTdxl5TYm7ILIcKHibAfR3v1P6fXTOElYbrJRrdMtFrlZD4MXacX6psTZM6RX1hpNg-aRp5NAfNAgbSA4ZPjWFAB000c3qnmWQ-43VoW9jVjkau4JaRUHKm_jEgpRP9Cuol4iHkrIdbmEvuOjRGydmCzz4SDhk7C0G00~2=WLaejI_zOoVX2Ldq05qA09DIHZHsfop23ldlZA7Fvt9dp_OxTzvmvwlyPmkMjbMXKdabXWOHFD5Yq74aRxchwk00IEVJbmAqpYV90hJE9yW2jBRDa-FtdKid4C16WebApdT64SOe3S92DZY5dSytJ6yRm7QVrqkaB6OsXXS2MrTR2R76cBiTh6sAVN5QWY4hmlisiwL_L2ZJNup0nP2RJ5emanwr1VDz2idP-xnNOi1qbsLDxjVxuWlGVfYsc_nW1AZyBmNpzr8mV-wCzkzUZCZNAj3qimeqCStqqlkzG3dIkwyZ95BsPcQUBtksChMj9titkonHVcm0TaVVbF_6i4r2ZNdlK9sMaiV5EYqbZrwxuSkS4f-DASnRzBfdl3B9XOl9QUxHFqxX4WAF3loTZSV7VNXq6YNh2tTyA1-CxrbCUxJi0e8DAuY2243XkXChRA10b8PFy0uHNGG0~2]]></Tracking><Tracking event=\"start\"><![CDATA[https://rus-trk.ascope.co/vdt?cm=1315&sk=203705&af=3&ts=1688093999&id=e2f050c3ad8b439d882bf0525ca60304&pt=YTMwUjLx0zYwZyM2ITMuETPihnJ2EDM14SM9IHa&ty=1]]></Tracking><Tracking event=\"complete\"><![CDATA[https://rus-trk.ascope.co/vdt?cm=1315&sk=203705&af=3&ts=1688093999&id=e2f050c3ad8b439d882bf0525ca60304&pt=YTMwUjLx0zYwZyM2ITMuETPihnJ2EDM14SM9IHa&ty=5]]></Tracking></TrackingEvents><VideoClicks><ClickThrough id=\"AW\"><![CDATA[https://an.yandex.ru/count/WVWejI_zOoVX2Lcr0LKF09CRQT3dSxcpP_kTEs-uyxsWjUFrK3jSvJzRiBIj2fNAAp4qYE295eEE8tdBNLK71q0wdxuKe7K-IHQWTJv15g2rRPuSl-zSEe40DX9CL7A-CuemHdq8BpzuFPYz1wuq11OYz109FM7md41nOmO7sy6J6XKiRTfWuOrNWx2G3KvXPwxTv_KIAOkPZMarx8XWm3-olY5Cx6gq8EXxZOuN0bjNMmcnnfYx7QnD8WTbzWLGI94r3Dgcs1Ud9HYoqsjLEhsTqUYqwgolGSvfxPHMyTgQsq7r4v_Q-VIPxymunZZfMwOUJDWPiWZlsvOUySX52GKk_SK8N0ovoQG5C-bHMp3Vfv2SlSzx9GPCLwbK5Y4hbShc4qIc5bOt-AhODVM-tXTVWFR3jDta1oD0vNyfcB-NWlbpPxH_zs94lbQ1ffzLe8bfZj4MXacX6xyNA4SYPMzcvlcoPojJstgoTxFB4cK0TaVVbF_6i4r2ZNdlK9sMaiV5EYqbZrwxuSkS4f-DASnRzBfdl3B9XOl9QUxHA-bRp9LAfRBgbOA4JRif_mNADWWckC1NGXJyO8wDJtpvIpyi-gmLMitjMc5pVA2Mq-pO9FMLDoEozyMGhBl3PNMtGMa9gRtx2LnGntp0Xmu9FbOiqNgwWtSO5p80~2]]></ClickThrough><ClickTracking><![CDATA[https://an.yandex.ru/tracking/WTWejI_zOoVX2Law0LKE02CQPJHsfop23ldlZ4sGjznLTN00fCwEIGB4FeSyemXZr8VmyeDNWUpRaG_x1nv3F8n5BC7piAK7u9fG377U0YG6bW6-IXX0nwcIxpXQQGWiHEWX4dh2u3c1uiOC3jRVV6tq8Ccvgh15oAWY7Cmq1_42J3ysZ9TGuWrYky_h5KjKYwmPmzI4lHYH6Yu8-bwZuuN0rjLMGgmnGnKEa7bP6OHOT0oWICpT3jOcaFTb18FkuSBt-HV6eZgJ8kZa0hU44zZ2Ni89R44RuwHlcNemOMV8Ivb6vJzRiBIj2fNAAp4qYE295eDC6IsJ7V78HGa5RfD92sRIexPYnCqO7URx59Apztcln83fAfNA8cIbKjbyWiZKmke6NnMxfbwDJxu0xOTfkyaFHe3A_r8mVoy5y-VEQF_knOXyhG9DFwj04zFylK0vqhjdArFRUh9tiyiIvSS6O7VqJVbl3Diar9Zt3jLfANdqMrhNzSKwBIMFNhlXovmIduqfp5lqEdC6XoaGl9A-nLofL2ghNYb8s3IFyXioLUsKIgKo3lpR80GetYEflw8wgf2wIrranrzlA5UorBSCjSnh89ozA_9GGgOWcCPTdon3k8f9Qly3~2?action-id=19]]></ClickTracking><ClickTracking><![CDATA[https://rus-trk.ascope.co/clk?cm=1315&sk=203705&af=3&ts=1688093999&id=e2f050c3ad8b439d882bf0525ca60304&pt=YTMwUjLx0zYwZyM2ITMuETPihnJ2EDM14SM9IHa]]></ClickTracking><ClickTracking><![CDATA[https://tk.toponadss.com/adx/tracking?d=eyJUeXBlIjoiYWR4Q2xpY2siLCJkYXRlX3RpbWUiOiIyMDIzLTA2LTMwIDEwOjU5OjU5IiwidGltZSI6MTY4ODA5Mzk5OSwicHVibGlzaGVyX2lkIjoxMzI2NiwiZm9ybWF0IjoxLCJwbF9pZF9pbnQiOjMzMjgyMSwiYXBwX2lkX2ludCI6NjIzNzgsInBsX2lkIjoiYjY0NTQ3MjE0ODQzOTYiLCJhcHBfaWQiOiJhNjQ0NzhhY2I0Y2NjZCIsImFwcF92biI6IiIsInNka192ZXIiOiIiLCJ1cGlkIjoiIiwidF9nX2lkIjowLCJncm9faWQiOjAsImdhaWQiOiI4Mjk4MTk0Mi1jMThlLTRkNzAtOTBjZi0wZDk2NjAzOTdkZjkiLCJyZXF1ZXN0X2lkIjoiMTY4ODA5Mzk5OTUxNTg1NzgwNF8zMzYwZTk3Ni0xNmYyLTExZWUtOWM4NC0zZjU5ODNkNDRkYWFfNTAwMzUyNjM3X3RvcG9uLXJ2Iiwib2ZmZXJfZmlybV9pZCI6NjYsImRzcF9pZCI6NzAsImJpZF9pZCI6ImUyZjA1MGMzYWQ4YjQzOWQ4ODJiZjA1MjVjYTYwMzA0IiwiY2xpZW50X2lwIjoiNS4xODEuNjIuMTI3IiwiZ2VvX3Nob3J0IjoiUlUiLCJkc3BfY3VycmVuY3kiOiJVU0QiLCJjdXJyZW5jeV9yYXRlIjoxLCJyZXZlbnVlX2N1dF9yYXRlIjowLjE1LCJwbGF0Zm9ybSI6MSwiYnVuZGxlX2luIjoiY29tLnBpeG9uaWMud3dyIiwiYnVuZGxlX2ZpbmFsIjoiY29tLnBpeG9uaWMud3dyIiwidG1heCI6MTQ2NywiYmlkZmxvb3IiOjAuMDEzMDAwMDAwMDAwMDAwMDAxLCJydGJfYWRfdHlwZSI6IjUiLCJpbnN0bCI6MCwiZXhwaXJlIjoxMjAsImFkb21haW4iOiJvc3Ryb3Zvay5ydSIsImNhdCI6IklBQjE5LElBQjIyLElBQjIxIiwidWEiOiJNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgMTI7IEFTVVNfSTAwNURBIEJ1aWxkL1NLUTEuMjEwODIxLjAwMTsgd3YpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIFZlcnNpb24vNC4wIENocm9tZS8xMTQuMC41NzM1LjEzMCBNb2JpbGUgU2FmYXJpLzUzNy4zNiIsIm9yaWdpbmFsX3ByaWNlX3N0ciI6IjFtUllrbUVwRHdQN3daLVowNUVkY0E9PSIsImRpc2NvdW50X3VzZF9wcmljZV9zdHIiOiI1emxOM1dWV3FldHNpSDZZYmNfMEZoeU5GMTQ2TDIzaXFXOExTaTBlTG9rPSIsImRzcF9hcHAiOiIyMDA1NzUiLCJkc3BfcGxhY2VtZW50IjoicmV3YXJkZWQgdmlkZW8iLCJ2ZXIiOiJvdXRTZWFBcGkifQ==]]></ClickTracking></VideoClicks><MediaFiles><MediaFile id=\"640p.mp4\" delivery=\"progressive\" type=\"video/mp4\" width=\"640\" height=\"360\" size=\"0\"><![CDATA[https://strm.yandex.ru/vh-canvas-converted/vod-content/1511467695061077301/ff61a008-8ba9-4148-8143-0a75bbf5521e/mp4/H264_640_360_1000.mp4?ssp-id=98384845&bid-id=e2f050c3ad8b439d882bf0525ca60304&uniq-id=16465532501705419199]]></MediaFile><MediaFile id=\"1280p.mp4\" delivery=\"progressive\" type=\"video/mp4\" width=\"1280\" height=\"720\" size=\"0\"><![CDATA[https://strm.yandex.ru/vh-canvas-converted/vod-content/1511467695061077301/ff61a008-8ba9-4148-8143-0a75bbf5521e/mp4/H264_1280_720_4480.mp4?ssp-id=98384845&bid-id=e2f050c3ad8b439d882bf0525ca60304&uniq-id=16465532501705419199]]></MediaFile><MediaFile id=\"854p.mp4\" delivery=\"progressive\" type=\"video/mp4\" width=\"854\" height=\"480\" size=\"0\"><![CDATA[https://strm.yandex.ru/vh-canvas-converted/vod-content/1511467695061077301/ff61a008-8ba9-4148-8143-0a75bbf5521e/mp4/H264_854_480_2420.mp4?ssp-id=98384845&bid-id=e2f050c3ad8b439d882bf0525ca60304&uniq-id=16465532501705419199]]></MediaFile><MediaFile id=\"426p.mp4\" delivery=\"progressive\" type=\"video/mp4\" width=\"426\" height=\"240\" size=\"0\"><![CDATA[https://strm.yandex.ru/vh-canvas-converted/vod-content/1511467695061077301/ff61a008-8ba9-4148-8143-0a75bbf5521e/mp4/H264_426_240_500.mp4?ssp-id=98384845&bid-id=e2f050c3ad8b439d882bf0525ca60304&uniq-id=16465532501705419199]]></MediaFile><MediaFile id=\"256p.mp4\" delivery=\"progressive\" type=\"video/mp4\" width=\"256\" height=\"144\" size=\"0\"><![CDATA[https://strm.yandex.ru/vh-canvas-converted/vod-content/1511467695061077301/ff61a008-8ba9-4148-8143-0a75bbf5521e/mp4/H264_256_144_300.mp4?ssp-id=98384845&bid-id=e2f050c3ad8b439d882bf0525ca60304&uniq-id=16465532501705419199]]></MediaFile><MediaFile id=\"1920p.mp4\" delivery=\"progressive\" type=\"video/mp4\" width=\"1920\" height=\"1080\" size=\"0\"><![CDATA[https://strm.yandex.ru/vh-canvas-converted/vod-content/1511467695061077301/ff61a008-8ba9-4148-8143-0a75bbf5521e/mp4/H264_1920_1080_7000.mp4?ssp-id=98384845&bid-id=e2f050c3ad8b439d882bf0525ca60304&uniq-id=16465532501705419199]]></MediaFile></MediaFiles></Linear></Creative></Creatives><Description></Description><Survey><![CDATA[https://www.yandex.ru]]></Survey><Extensions><Extension type=\"yandex_ad_info\"><![CDATA[{\"assets\":{\"first_frame\":[{\"url\":\"https://avatars.mds.yandex.net/get-vh/6213324/2a00000188a062d72c7d45380b50d247bacb/orig\",\"width\":\"1920\",\"height\":\"1080\"}]},\"text\":{}}]]></Extension></Extensions></InLine></Ad></VAST>",
											"adid":"yabs.NzIwNTc2MDg1OTg2NjQwNTg=",
											"adomain":[
												"ostrovok.ru"
											],
											"iurl":"https://bs-metadsp.yandex.ru/resource/spacer.gif",
											"cid":"1_190315589",
											"crid":"8_190315589_1920x1080_reach",
											"cat":[
												"IAB19",
												"IAB22",
												"IAB21"
											],
											"w":480,
											"h":320,
											"exp":1800
										}
									]
								}
							],
							"bidid":"e2f050c3ad8b439d882bf0525ca60304",
							"cur":"USD"
						}
					`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_midas_ua", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{"retreival_ad_info":[{"pos_id":2515,"ad_info":[{"scoring":{"local_id":4228,"ecpm":1000000},"ad_info":{"ad_id":1831836,"click_url":"https://ping.rqmob.com/click?advid=29\\u0026campid=1757672\\u0026platform=adshonor\\u0026pkg={pkg}\\u0026is_offline=1\\u0026requestid={requestid}\\u0026gaid={gaid}\\u0026device_id={device_id}\\u0026android_id={android_id}\\u0026beyla_id={beyla_id}\\u0026imei={imei}\\u0026imsi={imsi}\\u0026c={c}\\u0026site_id={site_id}\\u0026c_id={c_id}\\u0026adset={adset}\\u0026adset_id={adset_id}\\u0026ad={ad}\\u0026ad_id={ad_id}\\u0026ad_type={ad_type}\\u0026cost_model={cost_model}\\u0026cost_currency={cost_currency}\\u0026cost_value={cost_value}\\u0026remote_ip={remote_ip}\\u0026country_code={country_code}\\u0026ip={ip}\\u0026os_version={os_version}\\u0026placement={placement}\\u0026other_category={other_category}\\u0026ext_info={ext_info}\\u0026uagent={uagent}\\u0026sid={sid}\\u0026channel_pkg={channel_pkg}\\u0026channel_pkg_ver={channel_pkg_ver}\\u0026app_type={app_type}\\u0026app_id={app_id}\\u0026real_attrplat={real_attrplat}\\u0026adpos_id={adpos_id}\\u0026midas_camp_id={midas_camp_id}\\u0026midas_traffic_type={midas_traffic_type}\\u0026amp_app_id={amp_app_id}\\u0026is_pre_install={is_pre_install}\\u0026cut_type={cut_type}\\u0026package_type={package_type}\\u0026publisher={publisher}\\u0026oaid={oaid}","update_time":1688614726,"landing_page":"https://ping.rqmob.com/click?advid=29\\u0026campid=1757672\\u0026platform=adshonor\\u0026pkg={pkg}\\u0026is_offline=1\\u0026requestid={requestid}\\u0026gaid={gaid}\\u0026device_id={device_id}\\u0026android_id={android_id}\\u0026beyla_id={beyla_id}\\u0026imei={imei}\\u0026imsi={imsi}\\u0026c={c}\\u0026site_id={site_id}\\u0026c_id={c_id}\\u0026adset={adset}\\u0026adset_id={adset_id}\\u0026ad={ad}\\u0026ad_id={ad_id}\\u0026ad_type={ad_type}\\u0026cost_model={cost_model}\\u0026cost_currency={cost_currency}\\u0026cost_value={cost_value}\\u0026remote_ip={remote_ip}\\u0026country_code={country_code}\\u0026ip={ip}\\u0026os_version={os_version}\\u0026placement={placement}\\u0026other_category={other_category}\\u0026ext_info={ext_info}\\u0026uagent={uagent}\\u0026sid={sid}\\u0026channel_pkg={channel_pkg}\\u0026channel_pkg_ver={channel_pkg_ver}\\u0026app_type={app_type}\\u0026app_id={app_id}\\u0026real_attrplat={real_attrplat}\\u0026adpos_id={adpos_id}\\u0026midas_camp_id={midas_camp_id}\\u0026midas_traffic_type={midas_traffic_type}\\u0026amp_app_id={amp_app_id}\\u0026is_pre_install={is_pre_install}\\u0026cut_type={cut_type}\\u0026package_type={package_type}\\u0026publisher={publisher}\\u0026oaid={oaid}","target_page":3,"click_monitor_url":"https://ping.rqmob.com/click?advid=29\\u0026campid=1757673\\u0026platform=adshonor\\u0026pkg={pkg}\\u0026is_offline=1\\u0026requestid={requestid}\\u0026gaid={gaid}\\u0026device_id={device_id}\\u0026android_id={android_id}\\u0026beyla_id={beyla_id}\\u0026imei={imei}\\u0026imsi={imsi}\\u0026c={c}\\u0026site_id={site_id}\\u0026c_id={c_id}\\u0026adset={adset}\\u0026adset_id={adset_id}\\u0026ad={ad}\\u0026ad_id={ad_id}\\u0026ad_type={ad_type}\\u0026cost_model={cost_model}\\u0026cost_currency={cost_currency}\\u0026cost_value={cost_value}\\u0026remote_ip={remote_ip}\\u0026country_code={country_code}\\u0026ip={ip}\\u0026os_version={os_version}\\u0026placement={placement}\\u0026other_category={other_category}\\u0026ext_info={ext_info}\\u0026uagent={uagent}\\u0026sid={sid}\\u0026channel_pkg={channel_pkg}\\u0026channel_pkg_ver={channel_pkg_ver}\\u0026app_type={app_type}\\u0026app_id={app_id}\\u0026real_attrplat={real_attrplat}\\u0026adpos_id={adpos_id}\\u0026midas_camp_id={midas_camp_id}\\u0026midas_traffic_type={midas_traffic_type}\\u0026amp_app_id={amp_app_id}\\u0026is_pre_install={is_pre_install}\\u0026cut_type={cut_type}\\u0026package_type={package_type}\\u0026publisher={publisher}\\u0026oaid={oaid}","impr_monitor_url":"https://ad.adriver.ru/cgi-bin/rle.cgi?sid=1\\u0026bt=21\\u0026tuid=-1\\u0026ad=761042\\u0026pid=3671981\\u0026bid=9137128\\u0026bn=9137128\\u0026advid={advertising_id}\\u0026exss=\\u0026rnd=221660185","creative_info_map":{"23445875":{"creative_id":23445875,"creative_type":5,"format_id":26,"format_width":640,"format_height":100,"packed_json":"{\\\"ad_animation_type_video\\\":0,\\\"animation_type\\\":0,\\\"countdown\\\":0,\\\"creative_id\\\":23445875,\\\"creative_ver\\\":0,\\\"ext_info\\\":\\\"{\\\\\\\"need_cover\\\\\\\":false}\\\",\\\"height\\\":100,\\\"icon_url\\\":\\\"\\\",\\\"image_urls\\\":[\\\"http://static.rqmob.com/sa/20230706/d9bb3c090c1aae7c7af77ff43a39dc12__FILE_CM____FILE_CM480_720____WEBP__.jpg\\\"],\\\"landing_page\\\":\\\"\\\",\\\"layout_type\\\":-1,\\\"product_type\\\":0,\\\"scale_type\\\":1,\\\"support_jump\\\":false,\\\"type\\\":5,\\\"width\\\":640}","major_image_id":84232}},"track_urls":"|||||||||||||||","dsp_type":1008,"ad_set":{"id":17199,"campaign":{"id":12122,"objective":1,"billing_type":2,"bid_price":1000000,"start_time":1688515200,"end_time":1688774399,"attribution_pi":2,"is_online_ad":true,"ad_logo":1,"campaign_id":"202307050013","so_line_id":18872,"order_type":1,"campaign_name":"Ozon_CPM_RU_Masthead_7.5-7.7"},"start_time":1688515200,"end_time":1688774399,"placement_ids":[75,77,142,143,144,204,362,363,370,388,391,393,399,401,404,405,406,408,414,439,453,456,460,461,462,463,508,509,510,518,519,541,542,543,544,545,546,547,1128,1129,1130,1131,1132,1133,1134,1135,1226,1227,1228,1229,1230,1231,1243,1244,1245,1246,1253,1254,1255,1256,1258,1259,1260,1261,1262,1263,1264,1265,1272,1273,1274,1275,1276,1277,1278,1279,1280,1281,1288,1289,1290,1291,1292,1293,1294,1295,1296,1297,1323,1324,1325,1327,1328,1329,1330,1331,1337,1338,1339,1340,1341,1351,1352,1353,1354,1358,1359,1361,1362,1363,1373,1377,1378,1383,1387,1388,1389,1391,1392,1393,1394,1395,1396,1397,1398,1399,1400,1401,1402,1403,1404,1421,1422,1423,1424,1425,1426,1427,1428,1433,1434,1436,1437,1438,1439,1440,1450,1451,1452,1453,1460,1461,1462,1463,1464,1465,1466,1467,1468,1469,1470,1471,1472,1473,1474,1475,1477,1478,1479,1480,1481,1482,1483,1484,1485,1486,1487,1494,1495,1568,1569,1654,1751,1752,1753,1754,1755,1756,1757,1758,1759,1760,1761,1762,1763,1764,1765,1766,1767,1768,1769,1770,1771,1772,1774,1775,1776,1777,1781,1782,1783,1784,1785,1786,1787,1788,1789,1790,1791,1792,1793,1794,1795,1796,1797,1799,1803,1804,1805,1806,1807,1808,1809,1810,1811,1812,1813,1814,1815,1816,1817,1824,1825,1826,1827,1828,1829,1831,1832,1833,1834,1835,1836,1837,1838,1839,1840,1841,1842,1843,1844,1846,1849,1850,1851,1852,1853,1855,1856,1857,1858,1859,1860,1861,1862,1863,1864,1865,1866,1867,1868,1869,1870,1871,1872,1873,1876,1877,1878,1879,1880,1881,1887,1888,1889,1890,1891,1892,1893,1894,1895,1896,1897,1898,1899,1900,1901,1902,1903,1904,1905,1906,1907,1908,1909,1910,1911,1914,1915,1916,1917,1918,1919,1920,1921,1922,1923,1926,1928,1929,1933,1934,1935,1936,1937,1938,1939,1940,1947,1948,1949,1950,1951,1952,1953,1954,1958,1961,1962,1963,1966,1967,1968,1969,1970,1971,1972,1973,1974,1975,1976,1977,1978,1979,1980,1983,1984,1985,1986,1987,1988,1989,1991,1992,1993,1994,1995,1996,1997,1999,2000,2002,2003,2004,2005,2006,2007,2009,2010,2011,2012,2013,2014,2015,2018,2019,2020,2021,2022,2023,2024,2025,2026,2027,2028,2029,2030,2031,2032,2033,2034,2035,2036,2037,2038,2039,2040,2041,2044,2045,2047,2048,2049,2050,2051,2052,2053,2060,2061,2062,2063,2064,2065,2066,2067,2068,2069,2070,2073,2074,2075,2076,2077,2078,2079,2080,2081,2083,2084,2085,2086,2087,2088,2089,2090,2092,2093,2094,2095,2096,2097,2098,2099,2100,2101,2102,2103,2104,2105,2106,2107,2111,2112,2113,2114,2115,2116,2117,2119,2120,2121,2122,2123,2124,2125,2126,2128,2129,2130,2131,2132,2133,2134,2135,2136,2137,2138,2139,2140,2141,2142,2143,2144,2145,2146,2147,2148,2149,2150,2151,2152,2153,2154,2155,2156,2157,2158,2159,2160,2161,2162,2163,2164,2165,2166,2167,2168,2169,2170,2171,2172,2173,2174,2175,2176,2177,2178,2179,2180,2181,2184,2185,2186,2187,2188,2189,2190,2191,2192,2193,2194,2195,2196,2197,2198,2199,2200,2201,2202,2203,2204,2205,2206,2207,2208,2209,2210,2211,2212,2213,2214,2215,2216,2217,2218,2219,2220,2221,2222,2223,2224,2225,2226,2227,2228,2229,2230,2231,2232,2233,2234,2235,2236,2237,2238,2239,2240,2241,2242,2243,2244,2245,2246,2247,2248,2249,2250,2251,2252,2253,2254,2255,2256,2257,2258,2259,2260,2261,2262,2263,2264,2265,2266,2267,2268,2269,2270,2271,2272,2273,2274,2275,2276,2277,2278,2279,2280,2281,2282,2283,2284,2285,2286,2287,2288,2289,2290,2292,2293,2294,2295,2296,2297,2298,2299,2300,2301,2302,2303,2304,2305,2306,2307,2308,2309,2310,2311,2312,2313,2314,2315,2316,2317,2318,2319,2320,2321,2322,2323,2325,2326,2329,2330,2331,2332,2333,2334,2335,2336,2337,2338,2339,2340,2341,2342,2343,2344,2345,2346,2347,2348,2349,2350,2351,2352,2353,2354,2355,2356,2357,2358,2359,2360,2361,2363,2364,2365,2366,2367,2368,2369,2370,2372,2373,2374,2375,2376,2377,2378,2379,2380,2381,2382,2383,2384,2385,2386,2387,2388,2389,2390,2391,2392,2393,2394,2395,2396,2397,2398,2399,2400,2402,2403,2404,2405,2406,2407,2408,2409,2410,2411,2412,2413,2414,2415,2416,2417,2418,2419,2420,2421,2422,2423,2424,2425,2426,2427,2428,2429,2430,2431,2432,2433,2434,2435,2436,2437,2438,2439,2440,2441,2442,2443,2444,2445,2446,2447,2448,2449,2450,2451,2452,2453,2454,2455,2456,2457,2458,2459,2460,2461,2462,2463,2464,2465,2466,2467,2468,2469,2470,2471,2472,2473,2474,2475,2476,2477,2478,2479,2480,2482,2483,2484,2485,2486,2487,2488,2489,2490,2491,2492,2493,2494,2495,2496,2497,2498,2499,2500,2501,2502,2503,2504,2505,2506,2507,2508,2510,2511,2512,2513,2514,2515,2516,2517,2518,2519,2520,2521,2522,2523,2524,2525,2526,2527,2528,2529,2530,2531,2532,2533,2534,2535,2536,2537,2538,2540,2542,2543,2544,2545,2546,2547,2548,2549,2550,2551,2552,2553,2554,2555,2556,2557,2558,2559,2560,2561,2562,2563,2564,2565,2566,2567,2568,2569,2570,2571,2572,2573,2574,2575,2576,2577,2578,2579,2580,2581,2582,2583,2584,2586,2587,2588,2589,2590,2591,2592,2593,2594,2595,2596,2597,2598,2599,2600,2601,2602,2603,2604,2605,2606,2607,2616,2617,2618,2623,2624,2625,2626,2627,2628,2629,2630,2633,2634,2635,2636,2637,2638,2639,2640,2641,2642,2643,2644,2645,2646,2647,2648,2649,2650,2651,2652,2653,2654,2655,2656,2657,2658,2659,2660,2661,2662,2663,2664,2665,2666,2667,2668,2669,2670,2671,2672,2673,2674,2675,2676,2677,2678,2679,2680,2681,2682,2683,2684,2685,2686,2687,2688,2689,2690,2691,2692,2693,2694,2695,2696,2698,2699,2700,2701,2702,2703,2704,2705,2706,2707,2708,2709,2712,2713,2714,2722,2723,2724,2725,2726,2727,2732,2733,2737,2738,2740,2741,2742,2743,2744,2745,2746,2747,2748,2749,2750,2751,2752,2753,2754,2755,2756,2757,2758,2759,2761,2762,2764,2765,2766,2770,2771,2773,2777,2780,2781,2783,2784,2789,2790,2791,2792,2793,2794,2795,2796,2797,2798,2799,2800,2801,2805,2806,2807,2809,2810,2811,2812,2813,2814,2815,2816,2817,2818,2819,2822,2823,2824,2825,2826,2827,2828,2829,2830,2831,2832,2833,2834,2835,2836,2837,2838,2839,2840,2841,2842,2844,2846,2847,2848,2850,2851,2852,2853,2854,2855,2856,2857,2859,2860,2861,2862,2863,2864,2866,2868,2869,2871,2872,2874,2876,2877,2878,2881,2886,2887,2888,2889,2890,2891,2892,2895,2896,2897,2898,2899,2900,2902,2903,2904,2905,2906,2907,2908,2909,2910,2911,2912,2913,2915,2916,2917,2918,2919,2920,2921,2922,2923,2924,2925,2926,2927,2928,2929,2930,2931,2932,2933,2934,2935,2936,2937,2938,2939,2940,2941,2942,2943,2944,2945,2946,2947,2948,2949,2950,2951,2952,2953,2954,2955,2956,2957,2958,2959,2960,2961,2962,2963,2964,2965,2966,2967,2968,2969,2970,2971,2972,2973,2974,2975,2976,2977,2978,2979,2980,2981,2982,2983,2984,2985,2986,2987,2988,2989,2990,2991,2992,2993,2994,2995,2996,2997,2998,2999,3000,3001,3002,3003,3004,3005,3006,3007,3008,3009,3010,3011,3012,3013,3014,3015,3016,3017,3018,3019,3020,3021,3022,3023,3024,3025,3026,3027,3028,3029,3030,3031,3032,3033,3034,3035,3036,3037,3038,3039,3040,3041,3042,3043,3044,3045,3046,3047,3048,3049,3050,3051,3052,3053,3054,3055,3056,3057,3058,3059,3060,3061,3062,3063,3064,3065,3066,3067,3068,3069,3070,3071,3072,3073,3074,3075,3076,3077,3078,3079,3080,3081,3082,3083,3084,3085,3086,3087,3088,3089,3090,3091,3092,3093,3094,3095,3096,3097,3098,3099,3100,3101,3102,3103,3104,3105,3106,3107,3108,3109,3110,3111,3112,3113,3114,3115,3116,3117,3118,3119,3120,3121,3122,3123,3124,3125,3126,3127,3128,3130,3131,3132,3133]},"priority_weight":10000}}]}],"call_ma_success":true}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_oem_oppo", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"1160207280736_384164ab19eba19ceada6d22",
						"bidid":"56a8284b-2cb5-48bc-ab4f-a3c42d258d8c",
						"seatbid":[
							{
								"bid":[
									{
										"id":"1160207280736_384164ab19eba19ceada6d22",
										"impid":"1",
										"price":4.5,
										"nurl":"https://test-track.lacunads.com/oem_win?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWpVt5rSaVKYekN1yBD9f-zZYU_T7E_hACg6_mKDRs-9ZKqQqLrD3-Amy71hfMLMt-CT2xGL8kT5KvZyVv3zR6T94YvobB9OlEOD129d0dEI5aRcqiCoFgSecve-HpNkNryp24nOvX7IW-1_1y7YNxmMUCX4_bPJRITBLkYgSALULTRB_5kpj2awbO_Le0cd9x-XXzBQ58wEQf_r2eN6_OOMVcFNgd6QzMmRbvdxguHI4tS1NiZrephY3VHPbW90EseB4dmkQ4v3jsbeaz5RvFhYNC4RvFzVn89HJ8hZoVvMP6yDLH1ZUDT-vRq2dB5ctY0VDpDDM3j2Yf9qnVZSfy23xfjepdHYiasF8r_itBi2PB1MFMIwIeEqoGCOlHRxYI9ZO84c0SnkcRw&enc=1&auction_price=${AUCTION_PRICE}",
										"bundle":"com.adpdigital.mbs.ayande",
										"offer_id":"1996218",
										"trackUrls":[
											{
												"eventType":2,
												"url":"http://trident-test.bigpinggo.com/click?advid=30&campid=1996218&platform=felix&pkg=com.adpdigital.mbs.ayande&clickid=1160207280736_384164ab19eba19ceada6d22_1996218&offer_name=com.adpdigital.mbs.ayande_PH&creative_id={creative_id}&creative_name={creative_name}&affiliate_id=Vivo adx&idfa={idfa}&gaid=8c54269d-865f-4e7f-afee-32434231&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=&os_platform=Android&os_version=8.1.0&device_model=vivo 1811&ip=180.232.1.16&country_code=PH&ext={ext}&rta_token={rta_token}&san_campid=1000473&adset_id=4675&c_cto=0"
											},
											{
												"eventType":1,
												"url":"https://test-track.lacunads.com/oem_imp?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWpVt5rSaVKYekN1yBD9f-zZYU_T7E_hACg6_mKDRs-9ZKqQqLrD3-Amy71hfMLMt-CT2xGL8kT5KvZyVv3zR6T94YvobB9OlEOD129d0dEI5aRcqiCoFgSecve-HpNkNryp24nOvX7IW-1_1y7YNxmMUCX4_bPJRITBLkYgSALULTRB_5kpj2awbO_Le0cd9x-XXzBQ58wEQf_r2eN6_OOMVcFNgd6QzMmRbvdxguHI4tS1NiZrephY3VHPbW90EseB4dmkQ4v3jsbeaz5RvFhYNC4RvFzVn89HJ8hZoVvMP6yDLH1ZUDT-vRq2dB5ctY0VDpDDM3j2Yf9qnVZSfy23xfjepdHYiasF8r_itBi2PB1MFMIwIeEqoGCOlHRxYI9ZO84c0SnkcRw&enc=1&auction_price=${AUCTION_PRICE}"
											},
											{
												"eventType":2,
												"url":"https://test-track.lacunads.com/oem_clk?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWpVt5rSaVKYekN1yBD9f-zZYU_T7E_hACg6_mKDRs-9ZKqQqLrD3-Amy71hfMLMt-CT2xGL8kT5KvZyVv3zR6T94YvobB9OlEOD129d0dEI5aRcqiCoFgSecve-HpNkNryp24nOvX7IW-1_1y7YNxmMUCX4_bPJRITBLkYgSALULTRB_5kpj2awbO_Le0cd9x-XXzBQ58wEQf_r2eN6_OOMVcFNgd6QzMmRbvdxguHI4tS1NiZrephY3VHPbW90EseB4dmkQ4v3jsbeaz5RvFhYNC4RvFzVn89HJ8hZoVvMP6yDLH1ZUDT-vRq2dB5ctY0VDpDDM3j2Yf9qnVZSfy23xfjepdHYiasF8r_itBi2PB1MFMIwIeEqoGCOlHRxYI9ZO84c0SnkcRw&enc=1"
											}
										],
										"app_info":{
											"title":"com.adpdigital.mbs.ayande",
											"desc":"",
											"icon":"",
											"link_url":"https://play.google.com/store/apps/details?id=com.adpdigital.mbs.ayande&hl=en",
											"bundle":"com.adpdigital.mbs.ayande"
										}
									},
									{
										"id":"1160207280736_384164ab19eba19ceada6d22",
										"impid":"2",
										"price":4.5,
										"nurl":"https://test-track.lacunads.com/oem_win?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWoVt5rSaUyt2KGdIrZTV9DARZLXd3hnZ4wIlA84vV69h7HCXLK1WblrSxWNq-61ciGCWHhnNMrtZkJ5mpj_f63pAPA0-JxEptHPTlBKmG8dx6a3npzLS2toDJ9AlWV5-GhLIvclRO1oJNphOEU0eSauqpKGB19XC4CE7qcdUePZvYg-hDFmYFe233pQ1uF9Nte8lQr3qwiEaRmG5gfF4tzx24vFGiaAqF4yEBcUCP8fQ37WIk_qbN9EdpeNWSkVS5TUZB6zvZMvxA6g96OPFat8by6AwHnG7n6bBI8VgPqBh3N0zux-uB6hWkpErscvNsybjijOViRNpy9nDH0s-_vK8mpXCVdIyylyLSwxWeMHgKU8m_UtemNo7xCK32USyJ_XLXJ&enc=1&auction_price=${AUCTION_PRICE}",
										"bundle":"in.okcredit.merchant",
										"offer_id":"2146246",
										"trackUrls":[
											{
												"eventType":2,
												"url":"http://trident-test.bigpinggo.com/click?advid=30&campid=2146246&platform=felix&pkg=in.okcredit.merchant&clickid=1160207280736_384164ab19eba19ceada6d22_2146246&offer_name=in.okcredit.merchant_Global&creative_id={creative_id}&creative_name={creative_name}&affiliate_id=Vivo adx&idfa={idfa}&gaid=8c54269d-865f-4e7f-afee-32434231&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=&os_platform=Android&os_version=8.1.0&device_model=vivo 1811&ip=180.232.1.16&country_code=PH&ext={ext}&rta_token={rta_token}&san_campid=1000403&adset_id=6703&c_cto=0"
											},
											{
												"eventType":1,
												"url":"https://test-track.lacunads.com/oem_imp?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWoVt5rSaUyt2KGdIrZTV9DARZLXd3hnZ4wIlA84vV69h7HCXLK1WblrSxWNq-61ciGCWHhnNMrtZkJ5mpj_f63pAPA0-JxEptHPTlBKmG8dx6a3npzLS2toDJ9AlWV5-GhLIvclRO1oJNphOEU0eSauqpKGB19XC4CE7qcdUePZvYg-hDFmYFe233pQ1uF9Nte8lQr3qwiEaRmG5gfF4tzx24vFGiaAqF4yEBcUCP8fQ37WIk_qbN9EdpeNWSkVS5TUZB6zvZMvxA6g96OPFat8by6AwHnG7n6bBI8VgPqBh3N0zux-uB6hWkpErscvNsybjijOViRNpy9nDH0s-_vK8mpXCVdIyylyLSwxWeMHgKU8m_UtemNo7xCK32USyJ_XLXJ&enc=1&auction_price=${AUCTION_PRICE}"
											},
											{
												"eventType":2,
												"url":"https://test-track.lacunads.com/oem_clk?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWoVt5rSaUyt2KGdIrZTV9DARZLXd3hnZ4wIlA84vV69h7HCXLK1WblrSxWNq-61ciGCWHhnNMrtZkJ5mpj_f63pAPA0-JxEptHPTlBKmG8dx6a3npzLS2toDJ9AlWV5-GhLIvclRO1oJNphOEU0eSauqpKGB19XC4CE7qcdUePZvYg-hDFmYFe233pQ1uF9Nte8lQr3qwiEaRmG5gfF4tzx24vFGiaAqF4yEBcUCP8fQ37WIk_qbN9EdpeNWSkVS5TUZB6zvZMvxA6g96OPFat8by6AwHnG7n6bBI8VgPqBh3N0zux-uB6hWkpErscvNsybjijOViRNpy9nDH0s-_vK8mpXCVdIyylyLSwxWeMHgKU8m_UtemNo7xCK32USyJ_XLXJ&enc=1"
											}
										],
										"app_info":{
											"title":"OkCredit",
											"desc":"",
											"icon":"",
											"link_url":"https://play.google.com/store/apps/details?id=in.okcredit.merchant&hl=en",
											"bundle":"in.okcredit.merchant"
										}
									},
									{
										"id":"1160207280736_384164ab19eba19ceada6d22",
										"impid":"3",
										"price":4.5,
										"nurl":"https://test-track.lacunads.com/oem_win?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWrVttrSaXCDQFcj9mUr-954--wmmtEO37sMW5296FeXdOAYTxTtO-OaqLGINXSSsqKswhhzMylOn8hNpDmEBC0UYXajlSgrn037SveMII2Sm2qXjBTjQWVOHpGpaNjKD4db3ajLV4UfddZLx-0SoLlLq8NqeZWbvasQfJLPn3BePIzap1yoO3nEUUYPtTEwcL41GraweEXR5hYJfHntsjIIZ46bhWBZVVL7fLtq9-2BdBqbNKZhC_ACbGIYJ54yc-uQG-KCPBnRJDTT9Zty5jskhkplFgb-Y3tqPgtYWA677j_BAKbVedc_Q-tjQnP9Dj5JmLGrEMHPwqdNn-zuT4bzVKKrKP_zOQ9rh_ZiILd5RvIIzWegnndSSIN6fOiS7aNO37wOjaj-PEh&enc=1&auction_price=${AUCTION_PRICE}",
										"bundle":"com.jago.digitalBanking",
										"offer_id":"2146256",
										"trackUrls":[
											{
												"eventType":2,
												"url":"http://trident-test.bigpinggo.com/click?advid=30&campid=2146256&platform=felix&pkg=com.jago.digitalBanking&clickid=1160207280736_384164ab19eba19ceada6d22_2146256&offer_name=com.jago.digitalBanking_Global&creative_id={creative_id}&creative_name={creative_name}&affiliate_id=Vivo adx&idfa={idfa}&gaid=8c54269d-865f-4e7f-afee-32434231&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=&os_platform=Android&os_version=8.1.0&device_model=vivo 1811&ip=180.232.1.16&country_code=PH&ext={ext}&rta_token={rta_token}&san_campid=1000402&adset_id=6705&c_cto=0"
											},
											{
												"eventType":1,
												"url":"https://test-track.lacunads.com/oem_imp?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWrVttrSaXCDQFcj9mUr-954--wmmtEO37sMW5296FeXdOAYTxTtO-OaqLGINXSSsqKswhhzMylOn8hNpDmEBC0UYXajlSgrn037SveMII2Sm2qXjBTjQWVOHpGpaNjKD4db3ajLV4UfddZLx-0SoLlLq8NqeZWbvasQfJLPn3BePIzap1yoO3nEUUYPtTEwcL41GraweEXR5hYJfHntsjIIZ46bhWBZVVL7fLtq9-2BdBqbNKZhC_ACbGIYJ54yc-uQG-KCPBnRJDTT9Zty5jskhkplFgb-Y3tqPgtYWA677j_BAKbVedc_Q-tjQnP9Dj5JmLGrEMHPwqdNn-zuT4bzVKKrKP_zOQ9rh_ZiILd5RvIIzWegnndSSIN6fOiS7aNO37wOjaj-PEh&enc=1&auction_price=${AUCTION_PRICE}"
											},
											{
												"eventType":2,
												"url":"https://test-track.lacunads.com/oem_clk?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWrVttrSaXCDQFcj9mUr-954--wmmtEO37sMW5296FeXdOAYTxTtO-OaqLGINXSSsqKswhhzMylOn8hNpDmEBC0UYXajlSgrn037SveMII2Sm2qXjBTjQWVOHpGpaNjKD4db3ajLV4UfddZLx-0SoLlLq8NqeZWbvasQfJLPn3BePIzap1yoO3nEUUYPtTEwcL41GraweEXR5hYJfHntsjIIZ46bhWBZVVL7fLtq9-2BdBqbNKZhC_ACbGIYJ54yc-uQG-KCPBnRJDTT9Zty5jskhkplFgb-Y3tqPgtYWA677j_BAKbVedc_Q-tjQnP9Dj5JmLGrEMHPwqdNn-zuT4bzVKKrKP_zOQ9rh_ZiILd5RvIIzWegnndSSIN6fOiS7aNO37wOjaj-PEh&enc=1"
											}
										],
										"app_info":{
											"title":"Jago",
											"desc":"",
											"icon":"",
											"link_url":"https://play.google.com/store/apps/details?id=com.jago.digitalBanking&hl=en",
											"bundle":"com.jago.digitalBanking"
										}
									},
									{
										"id":"1160207280736_384164ab19eba19ceada6d22",
										"impid":"4",
										"price":4.5,
										"nurl":"https://test-track.lacunads.com/oem_win?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWsSMBrSaVaQGCIZCAkFE79-syjiUWIDRK6pAGplM2jBtNR3FTV3uiVugsxYyVGXzbokwulHoWSsi9zDdaJDAynIbiQrCwZn5mc84wMZz7WO7eWGP-YtVIcHBZhvgg5Qs4SdCpP6DyqbtPtP4Tnpj5-j-JzbMal0KK3i32kWE2BGihofRQR3U3QRZdXIt_fw9XMngpbGTPL7dfixJf0leE4OZXclpqrJb2B578ktnfoi2HFyJqYBjDn9TP1yFCYDESAaVpCxwKTZS-zHsdiX3TyPQkak4bobfi6w5gAx-cmfxio539hM7jOrx3Tla8448WM1LvRtPVphbOo66B9_mpJLEaSEd8wELJ3yAq1PrSsY4eYZQ==&enc=1&auction_price=${AUCTION_PRICE}",
										"bundle":"ajaib.co.id",
										"offer_id":"2009677",
										"trackUrls":[
											{
												"eventType":2,
												"url":"http://trident-test.bigpinggo.com/click?advid=30&campid=2009677&platform=felix&pkg=ajaib.co.id&clickid=1160207280736_384164ab19eba19ceada6d22_2009677&offer_name=ajaib.co.id_Global&creative_id={creative_id}&creative_name={creative_name}&affiliate_id=Vivo adx&idfa={idfa}&gaid=8c54269d-865f-4e7f-afee-32434231&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=&os_platform=Android&os_version=8.1.0&device_model=vivo 1811&ip=180.232.1.16&country_code=PH&ext={ext}&rta_token={rta_token}&san_campid=1000439&adset_id=3881&c_cto=0"
											},
											{
												"eventType":1,
												"url":"https://test-track.lacunads.com/oem_imp?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWsSMBrSaVaQGCIZCAkFE79-syjiUWIDRK6pAGplM2jBtNR3FTV3uiVugsxYyVGXzbokwulHoWSsi9zDdaJDAynIbiQrCwZn5mc84wMZz7WO7eWGP-YtVIcHBZhvgg5Qs4SdCpP6DyqbtPtP4Tnpj5-j-JzbMal0KK3i32kWE2BGihofRQR3U3QRZdXIt_fw9XMngpbGTPL7dfixJf0leE4OZXclpqrJb2B578ktnfoi2HFyJqYBjDn9TP1yFCYDESAaVpCxwKTZS-zHsdiX3TyPQkak4bobfi6w5gAx-cmfxio539hM7jOrx3Tla8448WM1LvRtPVphbOo66B9_mpJLEaSEd8wELJ3yAq1PrSsY4eYZQ==&enc=1&auction_price=${AUCTION_PRICE}"
											},
											{
												"eventType":2,
												"url":"https://test-track.lacunads.com/oem_clk?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlg1shbVALy7FstLO2_KVh9j4ejighVhSuw0tje4E8zP7s5AJMzXNk6zUGJRXFBfMoUVXLU-ycEWbaZfF20y02jLgf16du6VOkEeU7A8Fq1qPlnaE2xBjTX2WWsSMBrSaVaQGCIZCAkFE79-syjiUWIDRK6pAGplM2jBtNR3FTV3uiVugsxYyVGXzbokwulHoWSsi9zDdaJDAynIbiQrCwZn5mc84wMZz7WO7eWGP-YtVIcHBZhvgg5Qs4SdCpP6DyqbtPtP4Tnpj5-j-JzbMal0KK3i32kWE2BGihofRQR3U3QRZdXIt_fw9XMngpbGTPL7dfixJf0leE4OZXclpqrJb2B578ktnfoi2HFyJqYBjDn9TP1yFCYDESAaVpCxwKTZS-zHsdiX3TyPQkak4bobfi6w5gAx-cmfxio539hM7jOrx3Tla8448WM1LvRtPVphbOo66B9_mpJLEaSEd8wELJ3yAq1PrSsY4eYZQ==&enc=1"
											}
										],
										"app_info":{
											"title":"Ajaib",
											"desc":"",
											"icon":"",
											"link_url":"https://play.google.com/store/apps/details?id=ajaib.co.id&hl=en",
											"bundle":"ajaib.co.id"
										}
									}
								],
								"seat":"shareit"
							}
						],
						"cur":"USD"
					}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	r.HandleFunc("/mock/868/get_oem_dsp_vivo", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		sleepms := req.URL.Query().Get("sleepms")
		cpunum := req.URL.Query().Get("cpunum")
		print := req.URL.Query().Get("print")

		cpunumInt := 1
		if num, err := strconv.Atoi(cpunum); err == nil {
			cpunumInt = num
		}
		runtime.GOMAXPROCS(cpunumInt)

		sleepmsInt := 0
		if ms, err := strconv.Atoi(sleepms); err == nil {
			sleepmsInt = ms
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func(sleepmsInt int, w http.ResponseWriter, print string) {
			defer wg.Done()

			time.Sleep(time.Duration(int(sleepmsInt)) * time.Millisecond)
			nanoStr := strconv.FormatInt(time.Now().UnixNano(), 10)
			if print == "nano" {
				fmt.Fprintln(w, nanoStr)
			} else {
				s := `{
						"id":"1160207280736_384164ab19eba19ceada6d30",
						"bidid":"b2ae2291-cbd3-4ff1-a429-16aa7ec0132f",
						"seatbid":[
							{
								"bid":[
									{
										"id":"1160207280736_384164ab19eba19ceada6d30",
										"impid":"1",
										"price":6,
										"nurl":"https://track.lacunads.com/oem_win?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlq0chZSQDy7FN-c6-CNFVPnlSP-Z6sd2v7rNbUyjFs352FlL8eZxFx9vEeZWyy-LS5If50LGiVjtMq4HFwyH9djPbJyio4uZ0XLyiALzjvHNRm8r1ZMz8VUbXcl2uz7B21nNnpCC4tlTpsnXryGkAoXzHGTbA9Tfg-eMdzQ7qJJaXt1AP8I50IaG-LO8jpUL2P5memosJKeVK9TF42AX3sVB1tN7T92ZPloPxRb7XaD7ot8UOMVPty63yo6s4qsFAmKUczEckwxakCM7_kf2PFActcgp2AHIa9ZuSWxxBYJDPeSyno7EjtboZMr-fVAAHR5kA5AAffZXfyC0VubOn9xxmrLA4Ne35joNZFsH5PKScorvwdf4PRkr188Dl0ypbvJ5PqAwCmfv5h2pAvkrDAGZdj_ISDnLhKI0JS6QfaY80edeR5BRUcYfIo-PivLhkUfcufVkAYUs60NluXf2Mi4D8lS1reJralMxhZltGu&enc=1&auction_price=${AUCTION_PRICE}",
										"bundle":"com.unacademyapp",
										"offer_id":"1879393",
										"trackUrls":[
											{
												"eventType":2,
												"url":"http://trident.bigpinggo.com/click?advid=30&campid=1879393&platform=P_API_LacunaOEM&pkg=com.unacademyapp&clickid=1160207280736_384164ab19eba19ceada6d30_1879393&offer_name=com.unacademyapp_IN&creative_id={creative_id}&creative_name={creative_name}&affiliate_id=Vivo+adx&idfa={idfa}&gaid=8c54269d-865f-4e7f-afee-32434237&imei={imei}&android_id={android_id}&mac={mac}&oaid={oaid}&uagent=&os_platform=Android&os_version=8.1.0&device_model=vivo 1811&ip=127.0.0.1&country_code=IN&ext={ext}&rta_token={rta_token}&san_campid=1020528&adset_id=11904&c_cto=0&rtb_tag_id=2756"
											},
											{
												"eventType":1,
												"url":"https://track.lacunads.com/oem_imp?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlq0chZSQDy7FN-c6-CNFVPnlSP-Z6sd2v7rNbUyjFs352FlL8eZxFx9vEeZWyy-LS5If50LGiVjtMq4HFwyH9djPbJyio4uZ0XLyiALzjvHNRm8r1ZMz8VUbXcl2uz7B21nNnpCC4tlTpsnXryGkAoXzHGTbA9Tfg-eMdzQ7qJJaXt1AP8I50IaG-LO8jpUL2P5memosJKeVK9TF42AX3sVB1tN7T92ZPloPxRb7XaD7ot8UOMVPty63yo6s4qsFAmKUczEckwxakCM7_kf2PFActcgp2AHIa9ZuSWxxBYJDPeSyno7EjtboZMr-fVAAHR5kA5AAffZXfyC0VubOn9xxmrLA4Ne35joNZFsH5PKScorvwdf4PRkr188Dl0ypbvJ5PqAwCmfv5h2pAvkrDAGZdj_ISDnLhKI0JS6QfaY80edeR5BRUcYfIo-PivLhkUfcufVkAYUs60NluXf2Mi4D8lS1reJralMxhZltGu&enc=1&auction_price=${AUCTION_PRICE}"
											},
											{
												"eventType":2,
												"url":"https://track.lacunads.com/oem_clk?viewid=5thOjSqb-RpcwveXtS_vRKQeYnpruld6Tf6IWTwr2Dlq0chZSQDy7FN-c6-CNFVPnlSP-Z6sd2v7rNbUyjFs352FlL8eZxFx9vEeZWyy-LS5If50LGiVjtMq4HFwyH9djPbJyio4uZ0XLyiALzjvHNRm8r1ZMz8VUbXcl2uz7B21nNnpCC4tlTpsnXryGkAoXzHGTbA9Tfg-eMdzQ7qJJaXt1AP8I50IaG-LO8jpUL2P5memosJKeVK9TF42AX3sVB1tN7T92ZPloPxRb7XaD7ot8UOMVPty63yo6s4qsFAmKUczEckwxakCM7_kf2PFActcgp2AHIa9ZuSWxxBYJDPeSyno7EjtboZMr-fVAAHR5kA5AAffZXfyC0VubOn9xxmrLA4Ne35joNZFsH5PKScorvwdf4PRkr188Dl0ypbvJ5PqAwCmfv5h2pAvkrDAGZdj_ISDnLhKI0JS6QfaY80edeR5BRUcYfIo-PivLhkUfcufVkAYUs60NluXf2Mi4D8lS1reJralMxhZltGu&enc=1"
											}
										],
										"app_info":{
											"title":"Unacademy",
											"desc":"",
											"icon":"http://pmp-cdn.modengs.net/apps/2023-06-25/com.unacademyapp-1687701565.png",
											"link_url":"https://play.google.com/store/apps/details?id=com.unacademyapp&hl=en&gl=IN",
											"bundle":"com.unacademyapp"
										}
									}
								],
								"seat":"shareit"
							}
						],
						"cur":"USD"
					}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      r,
		WriteTimeout: time.Duration(100) * time.Second,
		ReadTimeout:  time.Duration(200) * time.Second,
		IdleTimeout:  time.Duration(200) * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	for {
		time.Sleep(time.Duration(360000) * time.Second)
	}

}
