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

	// mobupps banner ad
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
				s := `{"id":"xp90l0uqtll0c5c6kl2","cur":"USD","seatbid":[{"bid":[{"id":"000206e636279f4cfc41","impid":"1","price":0.073,"nurl":"https://t.adx.opera.com/win?a=a7314230524032&adomain=litres.ru&b=turkru.tv&cc=RU&cid=4028988167164987656&cm=1&crid=4028988167164987656&ct=&dvt=DESKTOP&ep=ep7226729523904&ext=JffUhbyX6FhEBira-RWv6HNh2Ei8QkcNRN5yyqJH1rtkdBp9YKcbFALlZBFdTs6VMAPuCCUKOMEiXBavLB7F-yg189l-sy7FGRsPk5h5kx1WlXlUxQp26AXhFodngyguN8t8AyP-RWcy2oW7jNeGYBYzVLg8K0dLdhCFXMmS86A%3D&impr_dl=1660827850378&m=m7314230524544&msz=300x250&ot=WINDOWS&pubId=pub7226695071424&radv=1495&rip=94.233.234.31&s=s7226771668160&said=fef65162602c7fc04d1f266b49d775e7&sc=-1&se=000206e636279f4cfc41&sh=-1&spid=31&srid=xp90l0uqtll0c5c6kl2&ss=3&u=630a321bb1d3c443&va=a7314230524032&vm=m7314230524544&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}","lurl":"https://t.adx.opera.com/loss?a=a7314230524032&adomain=litres.ru&b=turkru.tv&cc=RU&cid=4028988167164987656&cm=1&crid=4028988167164987656&ct=&dvt=DESKTOP&ep=ep7226729523904&ext=JffUhbyX6FhEBira-RWv6HNh2Ei8QkcNRN5yyqJH1rtkdBp9YKcbFALlZBFdTs6VMAPuCCUKOMEiXBavLB7F-yg189l-sy7FGRsPk5h5kx1WlXlUxQp26AXhFodngyguN8t8AyP-RWcy2oW7jNeGYBYzVLg8K0dLdhCFXMmS86A%3D&impr_dl=1660827850378&m=m7314230524544&msz=300x250&ot=WINDOWS&pubId=pub7226695071424&radv=1495&rip=94.233.234.31&s=s7226771668160&said=fef65162602c7fc04d1f266b49d775e7&sc=-1&se=000206e636279f4cfc41&sh=-1&spid=31&srid=xp90l0uqtll0c5c6kl2&ss=3&u=630a321bb1d3c443&va=a7314230524032&vm=m7314230524544&al=${AUCTION_LOSS}","adm":"<script type=\"text/javascript\">\n  function changeURLArg(url, arg, arg_val) {\n    var pattern = arg + '=([^&]*)';\n    var replaceText = arg + '=' + arg_val;\n    if (url.match(pattern)) {\n      var tmp = '/(' + arg + '=)([^&]*)/gi';\n      tmp = url.replace(eval(tmp), replaceText);\n      return tmp;\n    } else {\n      if (url.match('[\\?]')) {\n        return url + '&' + replaceText;\n      } else {\n        return url + '?' + replaceText;\n      }\n    }\n  }\n  function tracking(url) {\n    try {\n      navigator.sendBeacon(url);\n    } catch (error) {\n    }\n  }\n  function clickTracks(clkType) {\n    var urls = \"https://t.adx.opera.com/click?a=a7314230524032&adomain=litres.ru&b=turkru.tv&cc=RU&cid=4028988167164987656&cm=1&crid=4028988167164987656&ct=&dvt=DESKTOP&ep=ep7226729523904&ext=JffUhbyX6FhEBira-RWv6HNh2Ei8QkcNRN5yyqJH1rtkdBp9YKcbFALlZBFdTs6VMAPuCCUKOMEiXBavLB7F-yg189l-sy7FGRsPk5h5kx1WlXlUxQp26AXhFodngyguN8t8AyP-RWcy2oW7jNeGYBYzVLg8K0dLdhCFXMmS86A%3D&m=m7314230524544&msz=300x250&ot=WINDOWS&pubId=pub7226695071424&radv=1495&rip=94.233.234.31&s=s7226771668160&said=fef65162602c7fc04d1f266b49d775e7&sc=-1&se=000206e636279f4cfc41&sh=-1&spid=31&srid=xp90l0uqtll0c5c6kl2&ss=3&u=630a321bb1d3c443&va=a7314230524032&vm=m7314230524544&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}&ir=${IS_PREVIEW}\";\n    urls = urls.split(\",\");\n    for (var i = 0; i < urls.length; i++) {\n      var url = changeURLArg(urls[i], 'ct', clkType);\n      tracking(url);\n    }\n  }\n  var clickHandler = function() {\n    clickTracks(1)\n  }\n  var blurHandler = function() {\n    if (document.activeElement.tagName === \"IFRAME\") {\n      clickTracks(2)\n    }\n  }\n  var touchstartHandler = function() {\n    clickTracks(3)\n  }\n\n  document.addEventListener(\"DOMContentLoaded\", function () {\n    document.body.addEventListener('click', clickHandler);\n    document.body.addEventListener('touchstart', touchstartHandler);\n    window.addEventListener(\"blur\", blurHandler);\n  });\n\n  document.addEventListener(\"unload\", function () {\n    document.body.removeEventListener('click', clickHandler);\n    document.body.removeEventListener('touchstart', touchstartHandler);\n    window.removeEventListener(\"blur\", blurHandler);\n  });\n\n</script><head><style type='text/css'>body {margin:auto auto;text-align:center;} </style></head><div id=\"-oadx_000206e636279f4cfc41m73142305245441660826050379648726\"><div><html>\n<head>\n<meta charset=\"UTF-8\">\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/></head>\n<body style=\"margin: 0px; overflow:hidden;\"> \n<script type=\"text/javascript\">\nwindow.tbl_trc_domain = 'am-trc.taboola.com';\nwindow._taboola = window._taboola || [];\n_taboola.push({home:'auto'});\n!function (e, f, u, i) {\nif (!document.getElementById(i)){\ne.async = 1;\ne.src = u;\ne.id = i;\nf.parentNode.insertBefore(e, f);\n}\n}(document.createElement('script'),\ndocument.getElementsByTagName('script')[0],\n'//cdn.taboola.com/libtrc/example-placements/loader.js',\n'tb_loader_script');\nif(window.performance && typeof window.performance.mark == 'function')\n{window.performance.mark('tbl_ic');}\n</script>\n\n<div id=\"taboola-below-article-thumbnails\" style=\"height: 250px; width: 300px;\"></div>\n<script type=\"text/javascript\">\nwindow._taboola = window._taboola || [];\n_taboola.push({\nmode: 'SCOD_Test_300x250_1x1',\ncontainer: 'taboola-below-article-thumbnails',\nplacement: 'turkru.tv_Ppub7226695071424_Sfef65162602c7fc04d1f266b49d775e7_W300_H250_N1',\ntarget_type: 'mix',\n\"rtb-win\":{ \nbi:'000206e636279f4cfc41',\ncu:'USD',\nwp:'0.09109',\nwcb:'http://imprammp.taboola.com/st?cipid=66576435&ttype=0&cirid=08a8d08b-24c4-4e1d-b0e5-0b34d1e4a882&cicmp=67154105&cijs=1&dast=V7qJoBxgHm25KoM64CXQN1XcaZDemyzAR1XcaZDemyzAUAAAADBg0HJDBcGEeG4cQtGW1Ma9HKMXJLDCvXWjDxjEaOlWlhHE7WEIGm0-Fz3etFr8vX8rqLbncAAAAAeAAou5yC6NFtXI_o0W1clwAAAACgCKj4txC4AAAAAMD4____zwBIQLjVAIDlMHCG9fKw-wMA4OEBBABAAIMEAAAAoAQAAADg5P________8_ZuAefEaGAAA298agp-DaBwBACAAAIGtIAR6OB04ZpI4oQa-IEQAAAIBXuvrH0aROqLqq_v__-60ArgAAAgRh5b8wsnp1G9e3JmgwGS6Hw8Vst5iNlsPdbDUbVwI7TZeX5y55vQtBBpvNbLRczGuCBpPhcjhczHaL2Wg53M1Ws4UBAAAAjC3o3_L3XG53jd_tsv________9_s_-zfzQhF4WeNCEZMzO1Ht3G9bVfQAAANwCAtwC4mAOwAwAAALj7____95Ow0XK1GQ2Wo-FutRgOBqvl_gZiMBrgRAyWy8lkMdmtRqvRZrgbzQYLFIjBBCdkONpMVqPdajdZDiej0Wwz2SBFq1az0WYwXM0ms91uNRwMl6MRUrRmMZtMFrPRcrcZLCejwXAy3CGEbQYLz2QxsThGno1ptBlihm0GC89kMbE4Rp6NabTZE16Wh6fDN9ERAAAg4LEhR3PJZjOXbEZzzWKTIOay2VaL2WQ2mGx8M49hNHLMLLPZxLQc-XarlW-VP7dLxqfvnbKWLgXkdU29ZFB4YyWDQhsrGRTWWMmgUDZjJYNCGisZFMpYyaAwxkoGhTNWMiiUwVjJoDCWYyWDwhiOlQwKYzdWMiiM2VjJoDBWYyWDwhiNlQwKYzNWMiiMyVjJoDAWYyWDwhiMlQwKc6xkUIhjJYNCWWwTAAAAgBMCXhffZDKbLVeD3WI0GS05wzaDhWeymFgcI8_GNNrsCS_Lw9Phm9gsGgAAgIAc!&excid=93&tst=1&docw=0&s2s=true',\nrt:'1660826050326',\nrdc:'am.taboolasyndication.com',\nti:'5648',\nex:'OperaSCoD',\nbs:'fef65162602c7fc04d1f266b49d775e7',\nbp:'pub7226695071424',\nbd:'turkru.tv',\nsi:'17916'\n} \n,\nrec: {\"trc\":{\"si\":\"cb2411907026f3e933c0b5b1034f8eca\",\"sd\":\"v2_cb2411907026f3e933c0b5b1034f8eca_a8ae62c2-4760-42fd-818c-102827ee9605-tuct9f7b742_1660826050_1660826050_CIi3jgYQvt1UGIq6g8euqtuzsAEgASgBMLkBOMnrDEDymRBIs9naA1DAnydYAGAAaKjtzdPIpeaqL3AA\",\"ui\":\"a8ae62c2-4760-42fd-818c-102827ee9605-tuct9f7b742\",\"plc\":\"DESK\",\"wi\":\"-3325698876771888143\",\"cc\":\"RU\",\"route\":\"AM:AM:V\",\"el2r\":[\"bulk-metrics\",\"debug\",\"social\",\"abtests\",\"metrics\",\"perf\",\"supply-feature\"],\"uvpw\":\"1\",\"pi\":\"1388222\",\"cpb\":\"GNO629MGIJz__________wEqGWFtLnRhYm9vbGFzeW5kaWNhdGlvbi5jb20yDXRyYy1zY29kNDAxMjM4gNSn9wVAyesMSPKZEFCz2doDWMCfJ2MI_hYQiiAYE2RjCNcWENUfGCNkYwjSAxDgBhgIZGMIlhQQmRwYGGRjCIMvENc-GAlkYwj0FBCeHRgfZGMIpCcQijUYL2RqFDAwMDIwNmU2MzYyNzlmNGNmYzQxeAGAAQKIAaHXqM8BkAEc\",\"evh\":\"217325835\",\"evi\":{\"47\":\"5028|6794\"},\"vl\":[{\"ri\":\"9c405fad219c85984fef8ea32b187f9d\",\"uip\":\"turkru.tv_Ppub7226695071424_Sfef65162602c7fc04d1f266b49d775e7_W300_H250_N1\",\"ppb\":\"CIoF\",\"estimation_method\":\"EcpmEstimationMethodType_ESTIMATION\",\"original_ecpm\":\"0.11386413645642028\",\"baseline_variant\":\"false\",\"v\":[{\"thumbnail\":\"https://cdn.taboola.com/libtrc/static/thumbnails/0152b8e34e82f950278bd82e94be9878.gif\",\"all-thumbnails\":\"https://cdn.taboola.com/libtrc/static/thumbnails/195174768ca56fdc6f0c1ec91da69be0.jpg!-#@1000x600\",\"is-gift\":\"true\",\"origin\":\"default\",\"thumb-size\":\"1000x600\",\"description\":\"\",\"placement-label\":\"turkru.tv_Ppub7226695071424_Sfef65162602c7fc04d1f266b49d775e7_W300_H250_N1\",\"title\":\"ЛитРес: подписка за 399₽/месяц\",\"type\":\"text\",\"published-date\":\"1660217303\",\"branding-text\":\"ЛитРес\",\"url\":\"https://www.litres.ru/litres_subscription/?lfrom=1025278492&utm_source=taboola&utm_medium=video&utm_campaign=podpiska&tblci=GiCbMvUTXFYb4FcvwNWK0lmgVpboQvkCjCzJ1vGuNL25ZiC19VUouYKR6J7Bofct#tblciGiCbMvUTXFYb4FcvwNWK0lmgVpboQvkCjCzJ1vGuNL25ZiC19VUouYKR6J7Bofct\",\"duration\":\"0\",\"balancedPacingValue\":\"1.0\",\"sig\":\"f4ef82007d0a712c894d19c39a0e1a169f12da4e0f01\",\"constraintType\":\"SCALE\",\"item-id\":\"~~V1~~4028988167164987656~~DxPudjVzzvbSQ2dLcDfidnsDXnQuMLP-Dj1qSg5qobzTxvAnL2wqac4MyzR7uD46gj3kUkbS3FhelBtnsiJV6MhkDZRZzzIqDobN6rWmCPA3hYz5D3PLat6nhIftiT1lStQk6jaQcGGDUiGnTk93_4FS8qr-f53vDbiQQMMys25Xvi4P24MPI4rQxfLITK8T8ywME2ZwzKigjRsQeO32V8hbD21urIhLvAJGMPjxI-zqGDWsdg__ILQP_Zbbaqe-UWSjD2vpf1qi2yEN9nyKxQB8MQ5DwSrwujBax9ayWUY\",\"adomain\":\"litres.ru\",\"uploader\":\"\",\"cta-text\":\"Купить Сейчас\",\"is-syndicated\":\"true\",\"publisher\":\"aitagetlitres-sc\",\"id\":\"~~V1~~4028988167164987656~~DxPudjVzzvbSQ2dLcDfidnsDXnQuMLP-Dj1qSg5qobzTxvAnL2wqac4MyzR7uD46gj3kUkbS3FhelBtnsiJV6MhkDZRZzzIqDobN6rWmCPA3hYz5D3PLat6nhIftiT1lStQk6jaQcGGDUiGnTk93_4FS8qr-f53vDbiQQMMys25Xvi4P24MPI4rQxfLITK8T8ywME2ZwzKigjRsQeO32V8hbD21urIhLvAJGMPjxI-zqGDWsdg__ILQP_Zbbaqe-UWSjD2vpf1qi2yEN9nyKxQB8MQ5DwSrwujBax9ayWUY\",\"views\":\"0\"}]}],\"tslt\":{\"p-video-overlay\":{\"cancel\":\"Cancel\",\"goto\":\"Go To\"},\"read-more\":{\"DEFAULT_CAPTION\":\"Read%20More\"},\"next-up\":{\"BTN_TEXT\":\"Read Next Story\"},\"vignette\":{\"openBtn\":\"Learn More\",\"closeBtn\":\"Close\"},\"time-ago\":{\"yesterday\":\"Yesterday\",\"hours\":\"{0} hours ago\",\"hour\":\"1 hour ago\",\"minutes\":\"{0} minutes ago\",\"now\":\"Now\",\"today\":\"Today\",\"days\":\"{0} days ago\"},\"explore-more\":{\"POPUP_TEXT\":\"More stories to check out before you go\",\"TITLE_TEXT\":\"Keep on reading\"},\"adchoice\":{\"adChoiceBtn.title\":\"Why do I see this item?\"},\"userx\":{\"popover.content.questionnaire.options.uninteresting\":\"Uninteresting\",\"popover.content.questionnaire.options.racy\":\"Vulgar/Racy\",\"undoBtn.label\":\"Undo\",\"popover.title.scRemoved\":\"Sponsored link removed\",\"popover.content.questionnaire.options.repetitive\":\"Repetitive\",\"popover.title.thankYou\":\"Thank You!\",\"popover.title.removed\":\"Removed!\",\"popover.content.questionnaire.options.offensive\":\"Offensive\",\"popover.content.questionnaire.options.misleading\":\"Misleading\",\"removeBtn.title\":\"Remove this item\",\"popover.content.questionnaire.tellUsWhy\":\"Tell us why?\",\"popover.content.questionnaire.options.other\":\"Other\",\"popover.content.approval\":\"We will try not to show you this content anymore.\"}},\"cpcud\":{\"upc\":\"0.0\",\"upr\":\"0.0\"},\"dcga\":{\"pubConfigOverride\":{\"border-color\":\"black\",\"font-weight\":\"bold\",\"inherit-title-color\":\"true\",\"module-name\":\"cta-lazy-module\",\"enable-call-to-action-creative-component\":\"true\",\"disable-cta-on-custom-module\":\"true\"}}}}\n});\n</script>\n\n<script type=\"text/javascript\">\nwindow._taboola = window._taboola || [];\n_taboola.push({flush: true});\n</script>\n\n</body>\n</html></div><img id='adxImpressionTrackingPixel0' alt=\"\" src=\"https://t.adx.opera.com/impr?a=a7314230524032&adomain=litres.ru&b=turkru.tv&bl=1&cc=RU&cid=4028988167164987656&cm=1&crid=4028988167164987656&ct=&dvt=DESKTOP&ep=ep7226729523904&ext=JffUhbyX6FhEBira-RWv6HNh2Ei8QkcNRN5yyqJH1rtkdBp9YKcbFALlZBFdTs6VMAPuCCUKOMEiXBavLB7F-yg189l-sy7FGRsPk5h5kx1WlXlUxQp26AXhFodngyguN8t8AyP-RWcy2oW7jNeGYBYzVLg8K0dLdhCFXMmS86A%3D&impr_dl=1660827850378&m=m7314230524544&msz=300x250&ot=WINDOWS&pubId=pub7226695071424&radv=1495&rip=94.233.234.31&s=s7226771668160&said=fef65162602c7fc04d1f266b49d775e7&sc=-1&se=000206e636279f4cfc41&sh=-1&spid=31&srid=xp90l0uqtll0c5c6kl2&ss=3&u=630a321bb1d3c443&va=a7314230524032&vm=m7314230524544&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}\" style='width:0px;height:0px;border:0px;margin:0px;float:left;'/><img id='adxImpressionTrackingPixel1' alt=\"\" src=\"https://adrta.com/i?clid=opr&paid=opr&avid=adv7004556758144&kv7=pub7226695071424&publisherId=31&plid=m7314230524544-4028988167164987656&siteId=app7226716562816&caid=o7004556758592&lineItemId=a7314230524032&kv1=300x250&kv2=https%3A%2F%2Fturkru.tv&kv3=630a321bb1d3c443&kv4=94.233.234.31&kv10=&kv11=000206e636279f4cfc41&kv12=s7226771668160&kv15=RU&kv16=43.22860000&kv17=44.78190000&kv18=&kv19=&kv28=&kv23=&kv25=&kv26=windows&kv27=Mozilla%2F5.0+%28Windows+NT+10.0%3B+Win64%3B+x64%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F104.0.5112.81+Safari%2F537.36+Edg%2F104.0.1293.54&kv5=OpenRTB&kv24=Mobile_Web\" style='width:0px;height:0px;border:0px;margin:0px;float:left;'/><img id='adxImpressionTrackingPixel2' alt=\"\" src=\"https://sync.taboola.com/sg/OperaSCoD/1/cm\" style='width:0px;height:0px;border:0px;margin:0px;float:left;'/></div>","adomain":["litres.ru"],"cid":"4028988167164987656","crid":"4028988167164987656","h":250,"w":300,"exp":1800,"ext":{}}],"seat":"adv7004556758144"}],"bidid":"000206e636279f4cfc41"}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	// mobupps native ad
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
				s := `{"id":"e657d364-1164-49bf-b8cc-184e04737288","bidid":"c6ca008e9088c951f3ad52e5dc4960f0","seatbid":[{"bid":[{"id":"a6fe364e662e032e7d91e01eeb83fdbb","impid":"1","price":0.11284,"nurl":"http://b1.dcntr-ads.com/?win=nurl&sp=${AUCTION_PRICE}&t=native&uniq=511763b40a8546344164ad9895147f71","adm":"{\"native\":{\"ver\":\"1.1\",\"assets\":[{\"id\":1,\"required\":1,\"title\":{\"text\":\"Badoo\"}},{\"id\":2,\"img\":{\"url\":\"https://cdn.liftoff-creatives.io/customers/e12978da97/image/lambda_png/25d0fb7ccac186abdc31.png\",\"w\":1200,\"h\":627}},{\"id\":3,\"required\":1,\"img\":{\"url\":\"https://cdn.liftoff-creatives.io/customers/e12978da97/image/lambda_jpg_89/5b8003ee61c3a4c3114e.jpg\",\"w\":80,\"h\":80}},{\"id\":4,\"required\":1,\"data\":{\"value\":\"Otro estilo de citas. Añade tus intereses y habla antes de hacer match en Badoo.\",\"len\":80}},{\"id\":5,\"data\":{\"value\":\"INSTALAR\",\"len\":8}}],\"link\":{\"url\":\"https://play.google.com/store/apps/details?id=com.badoo.mobile&referrer=Liftoff_v.2_g.161950_a.0005076e65f4af012640_2832_c.123_t.ua_u.e5c4cb4f3d4a2167\",\"clicktrackers\":[\"https://click.liftoff.io/v1/campaign_click/Ch2McWAd1-6xzddPHmqZI6fgdFk_QUkhDlu2gVF0KFVthn--5jPDJL1uCORiMXvfSFTste9D-5snNiJl8USUrTkJfiXeLpLecjtEKa4cjpXe1Khf_-9p-jOOrSDtmGt025yj5ZuYDFcbgeFXyvz8D0-AoXlvRsdz6rV1fRafvr845eoYUu62nxVWWU58OIDug9etxabhuUsWRpC8QNu4LUu6YV3mQgFYycKCCQVkhUieyej2texwvMwBk2gx5wvIe2NFPgEX52XTaOyy-5AiQxVCTJdtLWtlxt8guUDIY7msrLkn462AqVTGVV7Y1tzVA6uAJK3BchrGLTfsqbp54ga4xaSYjGKtuOyOect_EZ11-NoaMkTPN6rdv87mqtnLDfYFBGbS9fwLauWflAas064gck6o_rIWMjlSip1ON6Z9wX1Q7TLL2aYdjgIhgdFGq9uT5slwnmt435mirs8u1zUwlJ4ju1SwTF7Ezb9shIum2WZq5t0WYJiRn8U4uamyvYKcT1SLmQw7omEP2Xx9kbTH3laHmLH7iUW2bl9exCtXRPbvtXuZytdAq-iD1qRA7e4hF1cABTfmwrgyfew7gyqffMkpuxq_P93xx89cxwjqILc\",\"https://t-odx.op-mobile.opera.com/click?a=a7005769032448&adomain=badoo.com&b=com.etermax.preguntados.lite&cc=ES&ch=third_party_sdk&cid=161950&cm=1&crid=304461&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep5738257056320&ext=awkKsd4F1X7kXA1Rd9wFUmHhnqyUMGIwf8CcUlqaL4Ioq0Bx0x13G-UL0R1l9-QejHDBC-pO0ERMNujZ9J_m6bWToxomGPKKkTzRPUqeqdIUpPXshi-eB5sTt6Be2h3C9vGeSF5arsi5HOtfG2d0Ag%3D%3D&gi=1B3E178C-313E-42D1-A607-8E3DFB6FDFE5&iabCat=IAB14-1&m=m7005769032832&msz=1200x627&ot=ANDROID&pc=1&pubId=pub4903319265728&rip=88.29.175.23&s=s4903325786432&said=af6df55abb31&sc=2&schp=1-2&se=0005076e65f4af012640&spid=194724177&srid=66e9fc2c806e0bf7b554d910a762&ss=2&stid=third_party_sdk&u=32871b0f481d9e97&va=a7005769032448&vm=m7005769032832&ac=USD&ap=0.18596&ir=${IS_PREVIEW}\"]},\"imptrackers\":[\"https://impression-europe.liftoff.io/opera/impression?ad_group_id=161950&channel_id=123&creative_id=304461&device_id_sha1=e5c4cb4f3d4a2167961759f3327f1846e2573093&source_app_store_id=com.etermax.preguntados.lite&bid_ts=1669965386480&auction_id=ocmCKsXcknnrLiQNXOWC&auction_price=0.245&origin=haggler-opera022\",\"https://opr.adx.opera.com/i?clid=opr&paid=opr&avid=adv7005106418624&caid=o7005106419136&plid=m7005769032832-304461&siteId=app4903323212480&kv7=pub4903319265728&publisherId=194724177&kv1=1200x627&kv4=88.29.175.23&kv3=32871b0f481d9e97&kv10=&kv12=s4903325786432&kv25=Trivia+Crack&kv15=ES&kv16=0.00000000&kv17=0.00000000&kv18=com.etermax.preguntados.lite&kv19=1B3E178C-313E-42D1-A607-8E3DFB6FDFE5&kv28=2201116SG&kv26=android&kv23=&kv27=Mozilla%2F5.0+%28Linux%3B+Android+11%3B+2201116SG+Build%2FRKQ1.211001.001%3B+wv%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Chrome%2F107.0.5304.141+Mobile+Safari%2F537.36&kv11=0005076e65f4af012640&lineItemId=a7005769032448&kv5=OpenRTB&kv24=Mobile_InApp_Native\",\"https://t-odx.op-mobile.opera.com/impr?a=a7005769032448&adomain=badoo.com&b=com.etermax.preguntados.lite&bl=1&cbu=H4sIAAAAAAAC_0yQu27rMBBE_4adZfEhyrqAcAuXCRKkijuBJpfSIhKXWFJGPj-wiiDlnHkUs9Say7_zGbfMUApSOsHOlKFZMVaKsUE6UwZ2fyL_XZhmpj1PGEZp5dC1wi8uJVgPorTwDK7iA55at8ZYKQI80D_BVBYnR-i88XcTdTBOSdsPVvbdELVWfZQXY0F1vW4HLQrt7GFyOU-lEh-TnrYGKvDmvpvMMO-pukClWbGCuGOYahmltcNgO32x5tIKt_uKlJ5l8tv1pdz8V0r8ih9vt_fP66-fGT2MbaNMJ4hxxjQubp5X4NNxQ6vUTwAAAP__IfgYdzQBAAA&cc=ES&ch=third_party_sdk&cid=161950&cm=1&crid=304461&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep5738257056320&ext=awkKsd4F1X7kXA1Rd9wFUmHhnqyUMGIwf8CcUlqaL4Ioq0Bx0x13G-UL0R1l9-QejHDBC-pO0ERMNujZ9J_m6bWToxomGPKKkTzRPUqeqdIUpPXshi-eB5sTt6Be2h3C9vGeSF5arsi5HOtfG2d0Ag%3D%3D&gi=1B3E178C-313E-42D1-A607-8E3DFB6FDFE5&iabCat=IAB14-1&impr_dl=1669967186547&m=m7005769032832&msz=1200x627&ot=ANDROID&pc=1&pubId=pub4903319265728&rip=88.29.175.23&s=s4903325786432&said=af6df55abb31&sc=2&schp=1-2&se=0005076e65f4af012640&spid=194724177&srid=66e9fc2c806e0bf7b554d910a762&ss=2&stid=third_party_sdk&u=32871b0f481d9e97&va=a7005769032448&vm=m7005769032832&ac=USD&ap=0.18596&ir=${IS_PREVIEW}\",\"https://ib1.bcc-ads.com/?win=impr&price=0.13715&prt=885_undefined&t=native&uniq=37e65a366d2df6ac4072918ea9367e85\",\"\",\"https://adrta.com/i?clid=deca&paid=deca&dvid=v&avid=opera_native_ron_11aug&caid=161950&publisherId=194724177&siteId=&plid=304461&kv3=a6dbe90ecdd44ec67064f771b914d44f3dbd9868&kv5=bidscube&kv16=&kv17=&kv23=Telefonica%20de%20Espana&kv1=undefinedxundefined&kv2=https%3A%2F%2Fplay.google.com%2Fstore%2Fapps%2Fdetails%3Fid%3Dcom.etermax.preguntados.lite&kv4=88.29.175.23&kv7=bidscube___dcntr&kv11=06a6e5af661c24f8f5527649ef1745f4&kv12=562488d7747b16478d2b0410a449&kv18=com.etermax.preguntados.lite&kv19=1B3E178C-313E-42D1-A607-8E3DFB6FDFE5&kv28=2201116SG&kv25=com.etermax.preguntados.lite&kv26=android&kv27=Mozilla%2F5.0%20(Linux%3B%20Android%2011%3B%202201116SG%20Build%2FRKQ1.211001.001%3B%20wv)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Version%2F4.0%20Chrome%2F107.0.5304.141%20Mobile%20Safari%2F537.36&kv24=Mobile_InApp_Native\",\"https://b1.dcntr-ads.com/?win=impr&price=${AUCTION_PRICE}&prt=49_undefined&t=native&uniq=511763b40a8546344164ad9895147f71\",\"\",\"https://adrta.com/i?clid=deca&paid=deca&dvid=v&avid=bdscb_native_mar31&caid=885_161950&publisherId=194724177&siteId=&plid=885_304461&kv3=a6dbe90ecdd44ec67064f771b914d44f3dbd9868&kv5=dcntrads&kv16=&kv17=&kv23=Telefonica%20de%20Espana&kv1=undefinedxundefined&kv2=https%3A%2F%2Fplay.google.com%2Fstore%2Fapps%2Fdetails%3Fid%3Dcom.etermax.preguntados.lite&kv4=88.29.175.23&kv7=dcntrads___mobupps&kv11=04136d991e7fb8a481411d903508b6ba&kv12=e657d364-1164-49bf-b8cc-184e04737288&kv18=com.etermax.preguntados.lite&kv19=1b3e178c-313e-42d1-a607-8e3dfb6fdfe5&kv28=2201116SG&kv25=com.etermax.preguntados.lite&kv26=android&kv27=Mozilla%2F5.0%20(Linux%3B%20Android%2011%3B%202201116SG%20Build%2FRKQ1.211001.001%3B%20wv)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Version%2F4.0%20Chrome%2F107.0.5304.141%20Mobile%20Safari%2F537.36&kv24=Mobile_InApp_Native\"]}}","adomain":["badoo.com"],"cat":["IAB14-1"],"adid":"55dbdb1c9686","iurl":"http://b1.dcntr-ads.com/?t=preview&k=55dbdb1c9686","cid":"49_885_161950","crid":"49_885_304461"}],"seat":"49"}],"cur":"USD"}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

	// mobupps video ad
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
				s := `{"id":"9a191d46-6054-4869-a547-9ef8d2426c55","seatbid":[{"bid":[{"id":"0002076e700e27e2ee83","impid":"1","price":2.384,"nurl":"https://t-odx.op-mobile.opera.com/win?a=a5522269118912\u0026adomain=wendys.com\u0026b=949295212\u0026cc=US\u0026cid=40021128\u0026cm=1\u0026crid=80859_55272534\u0026ct=\u0026dbbf=1.000\u0026dbmb=0.000\u0026dvt=PHONE\u0026ep=ep7270864899840\u0026ext=Cp85eZi9P-BMmPnnTXQkoyNWkG6zxQvYxK83vbIIOtLV8b2-5i1rN_9s1h8tAP9F15_Kezbi9-11_mALy0S-z7xcitQPmr18NGda6X6gOflvDwEB6zapWeuLZXXlfYqG_nFwff_rD0LbonlzBEY3Y3FZNwYObE6MejELWJEmg2g%3D\u0026gi=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026iabCat=IAB8-9\u0026impr_dl=1669969833967\u0026it=1\u0026m=m5522269119233\u0026ot=IOS\u0026pubId=pub7226695071424\u0026rip=172.58.103.70\u0026s=s7226745164736\u0026said=f5540250981342e753b259364d1da363\u0026sc=-1\u0026schp=-1\u0026se=0002076e700e27e2ee83\u0026spid=c1be64e41556140ba4007003a6a838e5\u0026srid=9a191d46-6054-4869-a547-9ef8d2426c55\u0026ss=3\u0026u=7cc3bfc211ced95f\u0026va=a5522269118912\u0026vd=120\u0026vm=m5522269119233\u0026ac=${AUCTION_CURRENCY}\u0026ap=${AUCTION_PRICE}","lurl":"https://t-odx.op-mobile.opera.com/loss?a=a5522269118912\u0026adomain=wendys.com\u0026b=949295212\u0026cc=US\u0026cid=40021128\u0026cm=1\u0026crid=80859_55272534\u0026ct=\u0026dbbf=1.000\u0026dbmb=0.000\u0026dvt=PHONE\u0026ep=ep7270864899840\u0026ext=Cp85eZi9P-BMmPnnTXQkoyNWkG6zxQvYxK83vbIIOtLV8b2-5i1rN_9s1h8tAP9F15_Kezbi9-11_mALy0S-z7xcitQPmr18NGda6X6gOflvDwEB6zapWeuLZXXlfYqG_nFwff_rD0LbonlzBEY3Y3FZNwYObE6MejELWJEmg2g%3D\u0026gi=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026iabCat=IAB8-9\u0026impr_dl=1669969833967\u0026it=1\u0026m=m5522269119233\u0026ot=IOS\u0026pubId=pub7226695071424\u0026rip=172.58.103.70\u0026s=s7226745164736\u0026said=f5540250981342e753b259364d1da363\u0026sc=-1\u0026schp=-1\u0026se=0002076e700e27e2ee83\u0026spid=c1be64e41556140ba4007003a6a838e5\u0026srid=9a191d46-6054-4869-a547-9ef8d2426c55\u0026ss=3\u0026u=7cc3bfc211ced95f\u0026va=a5522269118912\u0026vd=120\u0026vm=m5522269119233\u0026al=${AUCTION_LOSS}","adm":"\u003cVAST version=\"4.0\"\u003e\u003cAd id=\"55272534\"\u003e\u003cWrapper\u003e\u003cAdSystem\u003e\u003c![CDATA[Conversant]]\u003e\u003c/AdSystem\u003e\u003cVASTAdTagURI\u003e\u003c![CDATA[https://iad-usadmm.dotomi.com/fetch/html5/player/video/vast/2_0?cg=91\u0026dtmid=391591916636903898\u0026magic=42\u0026utype=0\u0026bidServerId=6238\u0026pnid=243\u0026pid=243\u0026ms=50\u0026trid=6238068087894529189\u0026dtm_server_id=1936\u0026comId=80859\u0026msgCampId=40021128\u0026tid=55272534\u0026parentMsgId=40021128\u0026ptid=50020278\u0026icb=0\u0026cgcb=-1\u0026fpc=0\u0026dvcid=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026mwp=2.980000\u0026btcurl=TriPeaks\u0026rt=4\u0026supplyType=3\u0026ctrl_ad_id=1\u0026iblob=h-ep3fq0CKWpnf2AvITJVhCi8eeOzTAaHVRyaVBlYWtzIFNvbGl0YWlyZTogQ2FyZCBHYW1lIgYyMTcyMDYqCTk0OTI5NTIxMjAAOiQ0ODhCQjU4RC1GN0IwLTQ5NUYtOTAyOC1EN0I1RTBFQzAyNTJSBUFwcGxlWgZJUEhPTkViBDE1LjVqBW9wZXJhcgUwLjAuMXgAggENMTcyLjU4LjEwMy43MJAB5ASaAQUwLjAuMaABA6oBBjEwMDI2ObIBBElBQjmyAQZJQUI5LTeyAQdJQUI5LTMwuAEDwAHL7MyY1eTc7QTIAf___________wHQAQDgAfitN-ABoPc24AHQ4TjoAdDju6O7jPKXxQHzAQoCVVMSAlVTGP4BKAA4AEAASABQAGAAbQAAAAB1AAAAAHoRVC1NT0JJTEUgVVNBIElOQy6KAQhULU1PQklMRZIBBk1PQklMRfQB-wEKAlVTEgJVUxj-ASICVFgoLDiXAlDqBFoFNzc1MjH8AYICCDEwMjcxODgxiAL___________8BmAIBoAIAqAKC7FCwAgHAAgLKAjwyMDQ5ODEwMjExfDg5ODI5Mjg0N3wxMDkwMTE0NDY2fDIxMTk2NjcyMzN8MTM2NDUxOTg3OXwwfC0xfDDgAgDoAgbzAgjprx8QmuniufUvGgExIeF6FK5H4co_KVVVVVVVVdU_9ALzAgjDsR8Q7v3jnMEwGgExIeF6FK5H4co_KVVVVVVVVdU_9ALzAgiXxxgQ_oO77q0wGgExIQAAAAAAAAAAKVVVVVVVVdU_9AL5AgAAAAAAAADAgQNczM8NTVnmP4kDzLIngc051z-RAwAAAAAAAPC_mQMAAAAAAADwv6EDAAAAAAAA8L-pAwAAAAAAAPA_sAMA8gMDVVNE-QMAAAAAAADwP4EEAAAAAAAAMUCJBNejcD0K1wdAkQQUrkfhehT-P5kETxHn1OHA-r6gBP-7xI_KMKgE5vzGAbAElBK5BLmDTAchFMhAwQQZdv3OFm_GP-kEAAAAAAAAAADxBAAAAAAAAAAA-AQAggUDaU9TiAUBkAUPmAUZqAUAsQUAAAAAAAAAALkFAAAAAAAAAADBBQAAAAAAAPC_yQUAAAAAAAAAANAFANgFAOkFAAAAAAAAAADxBQAAAAAAAAAA-QUAAAAAAAAAAIIGAklQmAb___________8BqAYAsAYB\u0026cturl=\u0026vpaid=1\u0026apis=6\u0026vcskippable_type=2\u0026vcskippable=false\u0026vrp=8\u0026min_duration=5\u0026max_duration=120\u0026dtm_user_ip=172.58.103.70]]\u003e\u003c/VASTAdTagURI\u003e\u003cImpression\u003e\u003c![CDATA[https://event.ad.cpe.dotomi.com/cvx/event/imp?enc=eAGNVmtzssgS_i9-8FNMhptAtqxTzS2ighLBxJzashBGRLkFUNRT57-fRpP47m7V7pkPMtP99NMzPe10_6dzqGgZh53nDiczgszITL_P9WXCSbLUeejUp6tOZPi-yEoyz_QliYiEMKjLaN3k5f6Vfh5oVZstx59xnIy4qqVgGZEl_YdOGB_jKs6zFs4gf6vj2IdOmq_jhGp-7SMN-U3mZVYWWIZF-3Uczso4oEjyKEsPneKwVvOq_ln6ZZ3R0qAIII_koRMX87qMswiJ0OmjID0yhHsUCVLF4abl5yVJUQRJ6xmiQnq8LBg9mbBSTxMVQSe6SlihdVwdiiI5u-cCmTk0zmoalX6Nu_-RpTSMr5LZbRcLWrana8_wSB7bKP0gbjZ9Bg9bJH5AU5rdgsYQVmQkqQVv_UrNswo1nee6PNCHTl7Qm0uknLPzleFNJggsbu7sQ5mgYlvXRfX89BT7Ye9Q-WGaPoZ5nafxY5CnT_SIfE9--JTEGxqcg4Q-BYeyROG_wjqNw8Ffb76b-lEcDHi2e6jx-APSReQKM-VIyxVaMDLX79KriiNE7OIVza9KMxz0WU7q4gXcZqQvEUnEzBFYmZHkbloNBNLFbSESNYLcxaPgVqwqQglPCMswrISwSPXT4ldRgbmCtoTFcKGDdiGwIitwfDcO1rjFIMJPj-mW9YDvhscAEf_PRXevNxRgRg3a_OoWGRqyPNctvr5pU1w1BEf3nhMDrhtXGBOMJR0w3Xid5OvBtkcLbvNJ1PFbkW1YOJruaLFVY4nS6cUFf7h4PfsLJVm-1RfTsI_rl4Qs35Lzh5tHDmucP1RluHxjEjNani03OFva8lN192TqmoLtmidrBzCNHTLVtqqz8_hXlXmxidlMXEewvWU9deE8VRkdZcyrqxjOBc62O5orntEEL6fkLfoYefp25u4XsaLpzGS3-FTe5ObjfbQNIq-Z7OBgvUcQRbrd-p-gj8lOb6wzz1kjUASY--B84XxQoJ8rSqvXTHa6NhVFTxRnl57B-Rg5nilMXIrz8GtuNQfQtQaGE9E6LxnqBqLjmrBZ3UczdMDRItjEtd0DJZ8FLA9Dh3d3OYTa7tCfirvZ-P3kDC_g5OpiYc0hWbgvM14ZA_CgA8xBceAFYO0ADoVpf2GYvy5UxnbJaOTqXrRY2IqpJ1Pn3B-Ds_UmHjNz9on1-oFn2H_NN47Sa_Rxy69Ho4V32vVgbqoLI8onWvwOifapGLlhXwLG2g0lWJqqqmEsdsFpqkWnGCb3c60kJQVTycGETxiroqE2EA0BoskYdg3esyNMr7b6aaNFODcFaxcRm2tOlrZvLFcntrZkNxrmgLtn7R3mxsWWLNdibc07Td2Im743zUYlaI_xi7QcovUFol1RniQnPWTxYeMdXyL9ZFKjb4yFIR_kq_Hia4TeSoZJi9eqV8kRj9wus_Tmn_Hx--kUOat8Koqf5Ia_xv0adfgTvwDRTQ6gQeTYwcWSbHeRpTN-r10mZhYFRGAuvVdovnAzdZU61tdCa459XftmkCa94o6DVQUWSJGl4d3qvV9sZryOWXEblqeOFJvuAo2MmSaEveN45X6zpVu3NxP2unsaZsx0CL2yHymznngyV2NrHOnC8fICa8DsHgvKJNVcCLaGtYXGcT7CIzc10tXLrLf_8YPnOynfTvHbAwf_U57me7Ibg6fswZul4H18ggeV492Rk71xX2gKZvJtYBzOv-JsMMCODJj-AX_6waNV71e8ab7APnFSWP8hJz9hCRUslev73rYCq7gY_KVqYlEq6Xwb0wSr9cZPKixMXxWofaX9OLrW8873830vUGpJsTwe6bU7uL73q-9X-w6CsFVnhyT5odXy1I-zqvP8705Ds_BctZWs8zvWZJoktPy158BXmSVin2JXQlmRspRK3EpmOQ4d3NBmW0jakvw3yJD6yX0X9FTTMmslt6p921vo1_R6IIps2CfJMhY3jpMF_qFT3rqgG7wzd8HW4FXDLdAsKM9FTcM2hncPeRlHcYaxyaLvytl5btsn7HfKvMGbaAu-v_HLGElCesQG6IvcurZKKPWL2Cj9lLZdWBsq5oH9_b__A8AqCTc\u0026__EXTRA__=__MACROS__]]\u003e\u003c/Impression\u003e\u003cImpression\u003e\u003c![CDATA[https://t-odx.op-mobile.opera.com/impr?a=a5522269118912\u0026adomain=wendys.com\u0026b=949295212\u0026bl=1\u0026cc=US\u0026cid=40021128\u0026cm=1\u0026crid=80859_55272534\u0026ct=\u0026dbbf=1.000\u0026dbmb=0.000\u0026dvt=PHONE\u0026ep=ep7270864899840\u0026ext=Cp85eZi9P-BMmPnnTXQkoyNWkG6zxQvYxK83vbIIOtLV8b2-5i1rN_9s1h8tAP9F15_Kezbi9-11_mALy0S-z7xcitQPmr18NGda6X6gOflvDwEB6zapWeuLZXXlfYqG_nFwff_rD0LbonlzBEY3Y3FZNwYObE6MejELWJEmg2g%3D\u0026gi=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026iabCat=IAB8-9\u0026impr_dl=1669969833967\u0026it=1\u0026m=m5522269119233\u0026ot=IOS\u0026pubId=pub7226695071424\u0026rip=172.58.103.70\u0026s=s7226745164736\u0026said=f5540250981342e753b259364d1da363\u0026sc=-1\u0026schp=-1\u0026se=0002076e700e27e2ee83\u0026spid=c1be64e41556140ba4007003a6a838e5\u0026srid=9a191d46-6054-4869-a547-9ef8d2426c55\u0026ss=3\u0026u=7cc3bfc211ced95f\u0026va=a5522269118912\u0026vd=120\u0026vm=m5522269119233]]\u003e\u003c/Impression\u003e\u003cImpression\u003e\u003c![CDATA[https://opr.adx.opera.com/i?clid=opr\u0026paid=opr\u0026dvid=v\u0026avid=adv5522269118272\u0026caid=o5522269118720\u0026plid=m5522269119233-80859_55272534\u0026siteId=app7226709609152\u0026kv7=pub7226695071424\u0026publisherId=c1be64e41556140ba4007003a6a838e5\u0026kv1=\u0026kv4=172.58.103.70\u0026kv3=7cc3bfc211ced95f\u0026kv10=\u0026kv12=s7226745164736\u0026kv25=TriPeaks+Solitaire%3A+Card+Game\u0026kv15=US\u0026kv16=29.64710000\u0026kv17=-95.48710000\u0026kv18=949295212\u0026kv19=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026kv28=IPHONE\u0026kv26=ios\u0026kv23=\u0026kv11=0002076e700e27e2ee83\u0026kv27=Mozilla%2F5.0+%28iPhone%3B+CPU+iPhone+OS+15_5+like+Mac+OS+X%29+AppleWebKit%2F605.1.15+%28KHTML%2C+like+Gecko%29+Mobile%2F15E148\u0026lineItemId=a5522269118912\u0026kv29=[ERRORCODE]\u0026kv30=[CONTENTPLAYHEAD]_[ADPLAYHEAD]\u0026kv33=[ASSETURI]\u0026kv34=[VASTVERSIONS]\u0026kv35=[IFATYPE]\u0026kv36=[IFA]\u0026kv37=[CLIENTUA]\u0026kv38=[SERVERUA]\u0026kv39=[DEVICEUA]\u0026kv40=[DEVICEIP]\u0026kv41=[LATLONG]\u0026kv42=[DOMAIN]\u0026kv43=[PAGEURL]\u0026kv44=\u0026kv45=[PLAYERSIZE]\u0026kv46=[REGULATIONS]\u0026kv47=[ADTYPE]\u0026kv48=[TRANSACTIONID]\u0026kv49=[BREAKPOSITION]\u0026kv50=[APPNAME]\u0026kv51=[PLACEMENTTYPE]\u0026kv54=[LAT]\u0026kv5=OpenRTB\u0026kv24=Mobile_InApp_Video]]\u003e\u003c/Impression\u003e\u003cError\u003e\u003c![CDATA[https://t-odx.op-mobile.opera.com/video?a=a5522269118912\u0026adomain=wendys.com\u0026b=949295212\u0026cc=US\u0026cid=40021128\u0026cm=1\u0026crid=80859_55272534\u0026ct=\u0026dbbf=1.000\u0026dbmb=0.000\u0026dvt=PHONE\u0026ep=ep7270864899840\u0026ext=Cp85eZi9P-BMmPnnTXQkoyNWkG6zxQvYxK83vbIIOtLV8b2-5i1rN_9s1h8tAP9F15_Kezbi9-11_mALy0S-z7xcitQPmr18NGda6X6gOflvDwEB6zapWeuLZXXlfYqG_nFwff_rD0LbonlzBEY3Y3FZNwYObE6MejELWJEmg2g%3D\u0026gi=488BB58D-F7B0-495F-9028-D7B5E0EC0252\u0026iabCat=IAB8-9\u0026impr_dl=1669969833967\u0026it=1\u0026m=m5522269119233\u0026ot=IOS\u0026pubId=pub7226695071424\u0026rip=172.58.103.70\u0026s=s7226745164736\u0026said=f5540250981342e753b259364d1da363\u0026sc=-1\u0026schp=-1\u0026se=0002076e700e27e2ee83\u0026spid=c1be64e41556140ba4007003a6a838e5\u0026srid=9a191d46-6054-4869-a547-9ef8d2426c55\u0026ss=3\u0026u=7cc3bfc211ced95f\u0026va=a5522269118912\u0026vd=120\u0026vel=1\u0026vm=m5522269119233\u0026evt=error]]\u003e\u003c/Error\u003e\u003cCreatives\u003e\u003cCreative id=\"80859_55272534\"\u003e\u003c/Creative\u003e\u003c/Creatives\u003e\u003cExtensions\u003e\u003c/Extensions\u003e\u003c/Wrapper\u003e\u003c/Ad\u003e\u003c/VAST\u003e","adomain":["wendys.com"],"cid":"40021128","crid":"80859_55272534","cat":["IAB8-9"],"exp":1800,"ext":{}}],"seat":"adv5522269118272"}],"bidid":"0002076e700e27e2ee83","cur":"USD"}`
				fmt.Fprintln(w, s)
			}
		}(sleepmsInt, w, print)

		wg.Wait()

		w.Header().Set("Connection", "close")
		w.Header().Set("Cache-Control", "no-cache")
	}).Methods("GET", "POST")

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
						"adm":"{\"native\":{\"assets\":[{\"id\":1,\"required\":1,\"img\":{\"type\":0,\"url\":\"http://dsp-adcreative.mobshark.net/adshark_dsp/1620899840630.jpg?x-oss-process=style/hq\",\"w\":1200,\"h\":627}}],\"link\":{\"url\":\"http://well.456halogames.com/subway-surfers/index.html?entry=sdk_ads\",\"clicktrackers\":[\"https://api.flat-ads.com/api/tracker/tracking/click?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424609&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424610&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=88876&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155356&dspid=100&srcid=3815373&cou=IDN&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&ag=6&ac=108&clickid=v1_2292_12606_88876_155356_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_108&et=click.47363.872080077674.927287201574.1&af_siteid=166_155356_2292\"],\"deeplink\":\"\"},\"imptrackers\":[\"https://api.flat-ads.com/api/tracker/tracking/impression?log_source=adx&bidid=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424609&bidid_id=6e0bbc2c-6974-4719-a5fd-91b280317b82.1667297549.21.57424610&impid=1&seat=498&price=0.1&secd_price=0.1&ver=4050878&cid=2292&adid=&crid=88876&ad_group_id=12606&abslot=&reqid=6e0bbc2c-6974-4719-a5fd-91b280317b82&ad_zone_id=155356&dspid=100&srcid=3815373&cou=IDN&gaid=a70afac7-c1c8-856e-b729-afc030b83a49&ag=6&ac=108&clickid=v1_2292_12606_88876_155356_IDN_1_6e0bbc2c-6974-4719-a5fd-91b280317b82_6_108&et=imp.03621.059429549382.104726872282.1\"],\"eventtrackers\":null}}",
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
						  "ext": {}
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
						  "price": 0.24,
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
				  "id": "d8b356ae-5c40-4d77-90b5-1401b2751254",
				  "seatbid": [
					{
					  "bid": [
						{
						  "id": "0004074a5d6c86804c80",
						  "impid": "0",
						  "price": 0.2454296,
						  "nurl": "https://t-odx.op-mobile.opera.com/win?a=a7005771013696&adomain=snap.com&b=com.mi.android.globalminusscreen&cc=IN&cid=160529&cm=1&crid=139087&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep5865235689408&ext=gpJ1cA2ilf-HT1DWiq0QvE6nd5yi_KW490CgkIYdYJN3WsS349HZQUturdfrOlM6euSgiZkNqMbYBOx0YNTkoaf9JPkSMV_2xEG5XS1uMsJuLw1QgljD3l2Sk4NC5WFy24ivHa0dyUT1H7TbOiKiPi4lkYlDTMlyB5OwCzpjg7A3D&gi=e3d4066d-c42b-4afd-bc8d-1b3b061269ba&iabCat=IAB1&impr_dl=1667554430843&m=m7005771014016&msz=1200x627&ot=ANDROID&pubId=pub5865193350528&rip=157.38.134.34&s=s5865212277120&said=337&sc=2&schp=1-2&se=0004074a5d6c86804c80&spid=637&srid=d8b356ae-5c40-4d77-90b5-1401b2751254&ss=2&u=1af1ac59ea3aff7b&va=a7005771013696&vm=m7005771014016&ac=${AUCTION_CURRENCY}&ap=${AUCTION_PRICE}",
						  "lurl": "https://t-odx.op-mobile.opera.com/loss?a=a7005771013696&adomain=snap.com&b=com.mi.android.globalminusscreen&cc=IN&cid=160529&cm=1&crid=139087&ct=&dbbf=1.000&dbmb=0.000&dvt=PHONE&ep=ep5865235689408&ext=gpJ1cA2ilf-HT1DWiq0QvE6nd5yi_KW490CgkIYdYJN3WsS349HZQUturdfrOlM6euSgiZkNqMbYBOx0YNTkoaf9JPkSMV_2xEG5XS1uMsJuLw1QgljD3l2Sk4NC5WFy24ivHa0dyUT1H7TbOiKiPi4lkYlDTMlyB5OwCzpjg7A3D&gi=e3d4066d-c42b-4afd-bc8d-1b3b061269ba&iabCat=IAB1&impr_dl=1667554430843&m=m7005771014016&msz=1200x627&ot=ANDROID&pubId=pub5865193350528&rip=157.38.134.34&s=s5865212277120&said=337&sc=2&schp=1-2&se=0004074a5d6c86804c80&spid=637&srid=d8b356ae-5c40-4d77-90b5-1401b2751254&ss=2&u=1af1ac59ea3aff7b&va=a7005771013696&vm=m7005771014016&al=${AUCTION_LOSS}",
						  "adm": "{\"native\":{\"ver\":\"1.2\",\"assets\":[{\"id\":1,\"required\":1,\"title\":{\"text\":\"Time Travel Today\"}},{\"id\":2,\"required\":1,\"img\":{\"url\":\"https://cdn.liftoff-creatives.io/customers/fd81ad044d/image/lambda_png/9fe85b85500f662a8154.png\",\"w\":300,\"h\":300}},{\"id\":3,\"required\":1,\"img\":{\"url\":\"https://cdn.liftoff-creatives.io/customers/fd81ad044d/image/lambda_jpg_85/0c36b820a6286001f10b.jpg\",\"w\":720,\"h\":1280}},{\"id\":4,\"data\":{\"value\":\"snap.com\",\"len\":8}},{\"id\":5,\"data\":{\"value\":\"See how you looked as baby, and how you'll look in 50 years.\",\"len\":60}},{\"id\":6,\"data\":{\"value\":\"INSTALL\",\"len\":7}}],\"link\":{\"url\":\"https://play.google.com/store/apps/details?id=com.snapchat.android\\u0026referrer=Liftoff_v.2_g.160529_a.0004074a5d6c86804c80_4016_c.123_t.ua_u.24ab72690a740ec3\",\"clicktrackers\":[\"https://click.liftoff.io/v1/campaign_click/9yONxExDsF8dj1FN4h2DVFx5R7Fu3VA17Wzi8PLu2SFXuPboLFzhXk4ZaR9SsGG6FZ7dqB5rcmqQD_PzgqjRDWdRCvRXnNfKaH4pp6anwMaWcBvl8SiOVP6yOiXnmQvybqSOumkNxD3ZoU9tARSKO9u_NaQZoW7SK7cBKwDzkFMLDXpytbC2gPnXG90z7_5frjjcsuqCKiC-a-J8HVq7XtIbid4_HLQcLHLu9uGp1Nq6SG1Wbt_a0bi1ZKTRNRKsv2yQKL3b0XP95cL_tXxqIryG8l6HbVdy__NSaqXgw8C8iZOdKqUO7EFvk0Mi7JDDErmcznNckxsiVZMzerrjZGY6rhI0Y3Q1i37ztTCdPZh5T7_OfRh-yqhcblY-GK7k8tEBnDChakYQRyaBFGepf6K63yDbleOE2MvERQPUsBMuVacVQAIwraBrVSK0GmqDu2I8IGuvtAtbOsiaQWomaw-3hyzApDHcnjXwonZgxoPk72xewv26sW-TP9Vmqw\",\"https://t-odx.op-mobile.opera.com/click?a=a7005771013696\\u0026adomain=snap.com\\u0026b=com.mi.android.globalminusscreen\\u0026cc=IN\\u0026cid=160529\\u0026cm=1\\u0026crid=139087\\u0026ct=\\u0026dbbf=1.000\\u0026dbmb=0.000\\u0026dvt=PHONE\\u0026ep=ep5865235689408\\u0026ext=gpJ1cA2ilf-HT1DWiq0QvE6nd5yi_KW490CgkIYdYJN3WsS349HZQUturdfrOlM6euSgiZkNqMbYBOx0YNTkoaf9JPkSMV_2xEG5XS1uMsJuLw1QgljD3l2Sk4NC5WFy24ivHa0dyUT1H7TbOiKiPi4lkYlDTMlyB5OwCzpjg7A3D\\u0026gi=e3d4066d-c42b-4afd-bc8d-1b3b061269ba\\u0026iabCat=IAB1\\u0026m=m7005771014016\\u0026msz=1200x627\\u0026ot=ANDROID\\u0026pubId=pub5865193350528\\u0026rip=157.38.134.34\\u0026s=s5865212277120\\u0026said=337\\u0026sc=2\\u0026schp=1-2\\u0026se=0004074a5d6c86804c80\\u0026spid=637\\u0026srid=d8b356ae-5c40-4d77-90b5-1401b2751254\\u0026ss=2\\u0026u=1af1ac59ea3aff7b\\u0026va=a7005771013696\\u0026vm=m7005771014016\\u0026ac=${AUCTION_CURRENCY}\\u0026ap=${AUCTION_PRICE}\\u0026ir=${IS_PREVIEW}\"]},\"imptrackers\":[\"https://impression-asia.liftoff.io/opera/impression?ad_group_id=160529\\u0026channel_id=123\\u0026creative_id=139087\\u0026device_id_sha1=24ab72690a740ec34c90dcee67f4eb2ac978842e\\u0026source_app_store_id=com.mi.android.globalminusscreen\\u0026bid_ts=1667547230799\\u0026auction_id=PnlaOPJeWKdjZDKGuNBv\\u0026auction_price=0.306787\\u0026origin=haggler-opera031\",\"https://opr.adx.opera.com/i?clid=opr\\u0026paid=opr\\u0026avid=adv7005106418624\\u0026caid=o7005106419136\\u0026plid=m7005771014016-139087\\u0026siteId=app5865202833792\\u0026kv7=pub5865193350528\\u0026publisherId=637\\u0026kv1=1200x627\\u0026kv4=157.38.134.34\\u0026kv3=1af1ac59ea3aff7b\\u0026kv10=\\u0026kv12=s5865212277120\\u0026kv25=MiuiApp\\u0026kv15=IN\\u0026kv16=0.00000000\\u0026kv17=0.00000000\\u0026kv18=com.mi.android.globalminusscreen\\u0026kv19=e3d4066d-c42b-4afd-bc8d-1b3b061269ba\\u0026kv28=redmi+note+5+pro\\u0026kv26=android\\u0026kv23=\\u0026kv27=Dalvik2F2.1.0+28Linux3B+U3B+Android+93B+Redmi+Note+5+Pro+MIUI2FV11.0.5.0.PEIMIXM29\\u0026kv11=0004074a5d6c86804c80\\u0026lineItemId=a7005771013696\\u0026kv5=OpenRTB\\u0026kv24=Mobile_InApp_Native\",\"https://t-odx.op-mobile.opera.com/impr?a=a7005771013696\\u0026adomain=snap.com\\u0026b=com.mi.android.globalminusscreen\\u0026bl=1\\u0026cbu=H4sIAAAAAAAC_0yQQWuEMBCF_01u68boGlOQQikUutDurdCLjMmoU2ImJHF_f1kPpcf3vcd3eGspMT-dz7TFhDkThxNkgsrTXHieK-IzR0zwb_AMblwS73EkN9SdvCgj7AohoD-IaoRNCIXueOTGyF4Lh3eyDzDmFepBtTBp1RkJupVom9Ya6Sxip-cWJwXW6L5vFYrMe7I4QoxjLpwOpeWt2qiC4BKTqxbPE_iNwp6zTYhBTOTGkoe66_Sl1aqR2hgBuy3E4SG4BQ-ft3f8urqf79fr2_7xcv_rYyKLg6wa2eleC060UBhWWBaP6XS8IZv6NwAA__-6lfdQOQEAAA\\u0026cc=IN\\u0026cid=160529\\u0026cm=1\\u0026crid=139087\\u0026ct=\\u0026dbbf=1.000\\u0026dbmb=0.000\\u0026dvt=PHONE\\u0026ep=ep5865235689408\\u0026ext=gpJ1cA2ilf-HT1DWiq0QvE6nd5yi_KW490CgkIYdYJN3WsS349HZQUturdfrOlM6euSgiZkNqMbYBOx0YNTkoaf9JPkSMV_2xEG5XS1uMsJuLw1QgljD3l2Sk4NC5WFy24ivHa0dyUT1H7TbOiKiPi4lkYlDTMlyB5OwCzpjg7A3D\\u0026gi=e3d4066d-c42b-4afd-bc8d-1b3b061269ba\\u0026iabCat=IAB1\\u0026impr_dl=1667554430843\\u0026m=m7005771014016\\u0026msz=1200x627\\u0026ot=ANDROID\\u0026pubId=pub5865193350528\\u0026rip=157.38.134.34\\u0026s=s5865212277120\\u0026said=337\\u0026sc=2\\u0026schp=1-2\\u0026se=0004074a5d6c86804c80\\u0026spid=637\\u0026srid=d8b356ae-5c40-4d77-90b5-1401b2751254\\u0026ss=2\\u0026u=1af1ac59ea3aff7b\\u0026va=a7005771013696\\u0026vm=m7005771014016\\u0026ac=${AUCTION_CURRENCY}\\u0026ap=${AUCTION_PRICE}\\u0026ir=${IS_PREVIEW}\"],\"privacy\":\"\"}}",
						  "adomain": [
							"snap.com"
						  ],
						  "bundle": "com.snapchat.android",
						  "iurl": "https://cdn.liftoff.io/customers/1378/previews/320d7bbbdb-ios-fad4aced99.png",
						  "cid": "160529",
						  "crid": "139087",
						  "cat": [
							"IAB1"
						  ],
						  "h": 627,
						  "w": 1200,
						  "exp": 7200,
						  "ext": {}
						}
					  ],
					  "seat": "adv7005106418624"
					}
				  ],
				  "bidid": "0004074a5d6c86804c80",
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
