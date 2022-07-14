#/bin/bash
### http mock data ###
#LBS服务#
curl --location --request POST 'http://172.27.210.34:3771/add' \
--header 'Content-Type: application/json;charset=UTF-8' \
--data '{"path":"location_de?coordinate=44.000000,116,000000&device_id=m.10f1f287dcb5&ip=2.56.231.255&beyla_id=339e834d-d46c-44fc-bd5e-e61c22c3ce97&version=3","method":"GET","input":{"equals":{}},"output":{"data":{"version":3,"res":1,"country":{"name":"Russian Federation","code":"RU"},"province":{"name":"Moscow","code":"MOW"},"city":"","msg":"","country_de":{"name":"usa","code":"US"},"province_de":"va","city_de":"sterling","na_asn_info":{},"na_isp_info":{},"proxy_info":{"ip":"2.56.231.255","trans_id":"WcmJjYbst9"}}}}'

curl --location --request POST 'http://172.27.210.34:3771/add' \
--header 'Content-Type: application/json;charset=UTF-8' \
--data '{"path":"location_de?coordinate=40.000000,106.000000&device_id=a.89e0d99efb000557&ip=159.138.88.145&beyla_id=450740fe4f52470c9779f1c2e98e9257&version=3","method":"GET","input":{"equals":{}},"output":{"data":{"version":3,"res":1,"country":{"name":"China","code":"CN"},"country_de":{"name":"sgp","code":"SG"},"province":{"name":"Ningsia Hui Autonomous Region","code":""},"district":"","city":"Shitanjing","msg":""}}}'

curl http://172.27.210.34:3770/mock/location_de?coordinate='40.000000,106.000000'&device_id=a.89e0d99efb000557&ip=159.138.88.145&beyla_id=450740fe4f52470c9779f1c2e98e9257&version=3

#dmp服务#
curl --location --request POST 'http://172.27.210.34:3771/add' \
--header 'Content-Type: application/json;charset=UTF-8' \
--data '{"path":"dmp?device_id=a.7bc2c54021803a95&beyla_id=91788da8a6ac44b2a96d3d0d40d144fb&app_id=com.lenovo.anyshare.gps&biz=ads_midas","method":"GET","input":{"equals":{},"contains":null,"matches":null},"output":{"data":{"status":1,"msg":"未匹配到此device_id"}}}'

# curl --location --request POST 'http://172.27.210.34:3771/add' \
# --header 'Content-Type: application/json;charset=UTF-8' \
# --data '{"path":"dmp?device_id=a.89e0d99efb000557&beyla_id=450740fe4f52470c9779f1c2e98e9257&app_id=com.lenovo.anyshare.gps&biz=ads_midas","method":"GET","input":{"equals":{},"contains":null,"matches":null},"output":{"data":{"msg":"未匹配到此device_id","status":1,"user_info":{"age":"","app_list":null,"applist_all":null,"beyla_id":"","city":"","country":"","device_id":"","earning":"","gender":"","interest":null,"interest_active":"","language":"","phone":"","province":"","update_time":""}}}}'

curl --location --request POST 'http://172.27.210.34:3771/add' \
--header 'Content-Type: application/json;charset=UTF-8' \
--data '{"path":"dmp?device_id=a.89e0d99efb000557&beyla_id=450740fe4f52470c9779f1c2e98e9257","method":"GET","input":{"equals":{},"contains":null,"matches":null},"output":{"data":{"msg":"未匹配到此device_id","status":1,"user_info":{"age":"","app_list":null,"applist_all":null,"beyla_id":"","city":"","country":"","device_id":"","earning":"","gender":"","interest":null,"interest_active":"","language":"","phone":"","province":"","update_time":""}}}}'


### grpc mock data ###
#retrieval#
echo "{\"service\":\"RetrievalService\",\"method\":\"RetrieveV2\",\"input\":{\"equals\":{\"ability_info\":{\"support_adx\":true,\"support_auto_download\":true,\"support_ca\":true,\"support_cdn\":true,\"support_downloader\":true,\"support_mraidjs\":true,\"support_offline\":true,\"support_reserve\":true,\"support_silently_install\":true,\"support_video\":true},\"app_info\":{\"app_id\":\"com.lenovo.anyshare.gps\",\"app_key\":\"c55d4ae8-3265-4979-b90a-b8f70c6ab7c4\",\"app_name\":\"com.lenovo.anyshare.gps\",\"app_type\":1,\"app_ver\":4060808,\"app_version\":\"6.8.8_ww\",\"i_ms\":516208872,\"sdk_ver\":4060003,\"u_ms\":91024411},\"channel_id\":36,\"country_code\":\"SG\",\"cpi_abtest_info\":{\"top_n\":6},\"deviceInfo\":{\"android_id\":\"89e0d99efb000557\",\"connection_type\":2,\"cpu_abi\":\"armeabi-v7a_armeabi\",\"cpu_bit\":\"32\",\"device_model\":\"SM-A105G\",\"device_type\":4,\"dpi\":280,\"gaid\":\"c4d4b421-b567-4820-ade6-b07d919ab5ac1\",\"geo\":{\"city\":\"Shitanjing\",\"country\":\"China\",\"country_de\":\"sgp\",\"ioscode\":\"cn\",\"isocode_de\":\"sg\",\"lat\":40000000,\"lon\":106000000,\"province\":\"Ningsia Hui Autonomous Region\"},\"ip\":\"159.138.88.145\",\"lang\":\"zh\",\"manufacturer\":\"samsung\",\"os_type\":2,\"os_version\":\"30\",\"screen_height\":1382,\"screen_width\":720,\"web_user_agent\":\"PostmanRuntime/7.29.0\"},\"dpa_user_profile\":{},\"enable_debug\":true,\"pos_id\":332,\"pos_info\":{\"ad_count\":10,\"adpos_id\":332,\"app_type\":1,\"mediation_info\":{\"trace_info\":{}},\"phy_pos\":\"ad:layer_p_mppb1\",\"sdk_mediation_info\":{\"anchor_time_out\":3000,\"anchor_wait_time\":1000,\"border\":1,\"delay_time\":1000,\"load_bottom_ad_time\":0,\"network_config\":[{\"ad_type\":\"adshonor_p\",\"bid\":0,\"hb\":0,\"identity\":\"332\",\"load_type\":2,\"name\":\"adshonor\",\"npa\":0,\"priority\":2,\"view\":{\"item_id\":16112},\"view_id\":\"GPB9\"}],\"no_ad_view\":{\"item_id\":1},\"no_ad_view_id\":\"GAE=\",\"pos_id\":\"ad:layer_p_mppb1\",\"pos_view\":{\"block_ids\":[5186],\"config_id\":3626,\"pos_id\":\"ad:layer_p_mppb1\"},\"pos_view_id\":\"CKocEMIoGhBhZDpsYXllcl9wX21wcGIx\",\"sub_tab_name\":\"配置1.0\",\"wait_time\":0},\"shielding_strategy\":{\"pkg_names\":[\"br.com.cea.appb2c\"],\"strategy_id\":81},\"support_crt_size\":[2,28,1],\"supported_lp_template_ids\":[4,5,8,9,15,17],\"supported_lp_type\":[2,3,4,1,6,7],\"traffic_type\":[1,2]},\"sum_frequency_info\":{},\"topn\":200,\"trace_id\":\"478eb37a-0e1c-4e4b-89b3-2e40f3c3a555\",\"user_info\":{\"gaid\":\"c4d4b421-b567-4820-ade6-b07d919ab5ac1\",\"user_id\":\"450740fe4f52470c9779f1c2e98e9257\"},\"user_targeting_id\":[9005009,10000332,1010480000000,2010360000000,1060130000000,15000002,15000028,15000001,38542572099,38478172065,38559349718,1,20000030,17000002,17000003,17000004,17000001,17000006,17000007,26000004,26000005,26000008,26000009,26000015,26000017,18000000,100614,16036067,16036000,9005009,17183929992,30068831880,25773864584,27999999,30000001,31001000]}},\"output\":{\"data\":{\"ret_code\":1,\"pos_id\":332,\"ad_set_infos\":{\"5024\":{\"ad_info\":[{\"ad_id\":298602,\"scoring\":{\"local_id\":436,\"ecpm\":1000000},\"creatives\":[{\"creative_id\":1088982},{\"creative_id\":1088986}],\"is_online_ad\":true}]},\"5032\":{\"ad_info\":[{\"ad_id\":326808,\"scoring\":{\"local_id\":578,\"ecpm\":1000000},\"creatives\":[{\"creative_id\":1201815},{\"creative_id\":1201810}],\"is_online_ad\":true,\"is_offline_ad\":true}]}},\"whiteboard\":\"[{\\\"retrieval_recall\\\":[{\\\"index_recall_ad_set_id\\\":[5032,5024]}]},{\\\"retrieval_filter\\\":[{\\\"reservation_order_filter(AdSet)\\\":[]},{\\\"online_ad_recall_rate_filter(AdSet)\\\":[]},{\\\"filter_ads_overview\\\":null},{\\\"filter_crts_overview\\\":[{\\\"AdId\\\":326808,\\\"CrtId\\\":1201819,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1073368,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1201816,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1201812,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1073582,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":342197,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1201817,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1072033,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1072036,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1071617,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1073305,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1201818,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1201813,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":342056,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1201814,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":342195,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":326808,\\\"CrtId\\\":1201811,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":298602,\\\"CrtId\\\":1073582,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":298602,\\\"CrtId\\\":1073326,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":298602,\\\"CrtId\\\":1088983,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":298602,\\\"CrtId\\\":1088984,\\\"FilterReason\\\":\\\"FormatFilter\\\"},{\\\"AdId\\\":298602,\\\"CrtId\\\":1088985,\\\"FilterReason\\\":\\\"FormatFilter\\\"}]}]},{\\\"pre_filter\\\":[{\\\"over_budget_filter(AdSet)\\\":[]},{\\\"attribute_window_filter(AdSet)\\\":[]}]}]\"}}}" > /tmp/retrieval_add.json
curl -XPOST 'http://172.27.210.34:4771/add' --header 'Content-Type: text/plain' -d@/tmp/retrieval_add.json

#dpa#
# curl --location --request POST 'http://172.27.210.34:4771/add' \
# --header 'Content-Type: text/plain' \
# --data '{
#                         "service": "UserProfileService",
#                         "method":"GetUserProfile",
#                         "input":{
#                             "equals":{
#                         "beyla_id": "450740fe4f52470c9779f1c2e98e9257",
#                         "ctx": {
#                             "app_name": "com.lenovo.anyshare.gps",
#                             "app_version": "6.8.8_ww",
#                             "brand": "samsung",
#                             "city": "Shitanjing",
#                             "client_ip": "159.138.88.145",
#                             "country": "China",
#                             "device_type": 4,
#                             "language": "zh",
#                             "lat": 40000000,
#                             "lon": 106000000,
#                             "model": "SM-A105G",
#                             "network_type": 2,
#                             "os_version": "30",
#                             "pkg_name": "com.lenovo.anyshare.gps",
#                             "pos_ids": [
#                                 332
#                             ],
#                             "province": "Ningsia Hui Autonomous Region",
#                             "req_time": 1648890983,
#                             "screen_height": 1382,
#                             "screen_width": 720
#                         },
#                         "gaid": "c4d4b421-b567-4820-ade6-b07d919ab5ac1",
#                         "project": "online",
#                         "trace_id": "478eb37a-0e1c-4e4b-89b3-2e40f3c3a555",
#                         "traffic_source": 1
#                     }
#                         },
#                         "output":{
#                             "data":{
#                         "trace_id": "478eb37a-0e1c-4e4b-89b3-2e40f3c3a555",
#                         "user_profile": {}
#                     }
#                         }
#                     }'

# curl --location --request POST 'http://172.27.210.34:4771/add' \
# --header 'Content-Type: text/plain' \
# --data '{
#                         "service": "UserProfileService",
#                         "method":"GetUserProfile",
#                         "input":{
#                             "contains":{
#                         "beyla_id": "450740fe4f52470c9779f1c2e98e9257",
#                         "ctx": {
#                             "app_name": "com.lenovo.anyshare.gps",
#                             "app_version": "6.8.8_ww",
#                             "brand": "samsung",
#                             "city": "Shitanjing",
#                             "client_ip": "159.138.88.145",
#                             "country": "China",
#                             "device_type": 4,
#                             "language": "zh",
#                             "lat": 40000000,
#                             "lon": 106000000,
#                             "model": "SM-A105G",
#                             "network_type": 2,
#                             "os_version": "30",
#                             "pkg_name": "com.lenovo.anyshare.gps",
#                             "pos_ids": [
#                                 332
#                             ],
#                             "province": "Ningsia Hui Autonomous Region",
#                             "screen_height": 1382,
#                             "screen_width": 720
#                         },
#                         "gaid": "c4d4b421-b567-4820-ade6-b07d919ab5ac1",
#                         "project": "online",
#                         "trace_id": "478eb37a-0e1c-4e4b-89b3-2e40f3c3a555",
#                         "traffic_source": 1
#                     }
#                         },
#                         "output":{
#                             "data":{
#                         "trace_id": "478eb37a-0e1c-4e4b-89b3-2e40f3c3a555",
#                         "user_profile": {}
#                     }
#                         }
#                     }'


# curl --location --request POST 'http://172.27.210.34:4771/delete' \
# --header 'Content-Type: text/plain' \
# --data '{
#                         "service": "UserProfileService",
#                         "method":"GetUserProfile",
#                         "input":{
#                             "contains":{
#                         "beyla_id": "450740fe4f52470c9779f1c2e98e9257",
#                         "ctx": {
#                             "app_name": "com.lenovo.anyshare.gps",
#                             "app_version": "6.8.8_ww",
#                             "brand": "samsung",
#                             "city": "Shitanjing",
#                             "client_ip": "159.138.88.145",
#                             "country": "China",
#                             "device_type": 4,
#                             "language": "zh",
#                             "lat": 40000000,
#                             "lon": 106000000,
#                             "model": "SM-A105G",
#                             "network_type": 2,
#                             "os_version": "30",
#                             "pkg_name": "com.lenovo.anyshare.gps",
#                             "pos_ids": [
#                                 332
#                             ],
#                             "province": "Ningsia Hui Autonomous Region",
#                             "screen_height": 1382,
#                             "screen_width": 720
#                         },
#                         "gaid": "c4d4b421-b567-4820-ade6-b07d919ab5ac1",
#                         "project": "online",
#                         "trace_id": "478eb37a-0e1c-4e4b-89b3-2e40f3c3a555",
#                         "traffic_source": 1
#                     }
#                         },
#                         "output":{
#                             "data":{
#                         "trace_id": "478eb37a-0e1c-4e4b-89b3-2e40f3c3a555",
#                         "user_profile": {}
#                     }
#                         }
#                     }'


curl --location --request POST 'http://172.27.210.34:4771/add' \
--header 'Content-Type: text/plain' \
--data '{
                        "service": "UserProfileService",
                        "method":"GetUserProfile",
                        "input":{
                            "contains":{
                        "beyla_id": "450740fe4f52470c9779f1c2e98e9257",
                        "gaid": "c4d4b421-b567-4820-ade6-b07d919ab5ac1",
                        "project": "online",
                        "trace_id": "478eb37a-0e1c-4e4b-89b3-2e40f3c3a555",
                        "traffic_source": 1
                    }
                        },
                        "output":{
                            "data":{
                        "trace_id": "478eb37a-0e1c-4e4b-89b3-2e40f3c3a555",
                        "user_profile": {}
                    }
                        }
                    }'

#APPlist#
curl --location --request POST 'http://172.27.210.34:4771/add' \
--header 'Content-Type: text/plain' \
--data '{"service": "FilterService","method":"Filter","input":{"equals":{"beyla_id": "450740fe4f52470c9779f1c2e98e9257","cycle": 366,"gaid": "c4d4b421-b567-4820-ade6-b07d919ab5ac1","source": "midas"}},"output":{"data":{"Status":10000}}}'
