-- example HTTP POST script which demonstrates setting the
-- HTTP method, body, and adding a header
wrk.method = "POST"
wrk.body = [[
{
    "id": "32f94228-52d3-4bce-be88-5dd1d799ae46",
    "imp": [
        {
            "id": "6",
            "tagid": "2524",
            "bidfloor": 0.00,
            "bidfloorcur": "USD",
            "secure": 0,
            "native": {
                "request": "{\"native\":{\"assets\":[{\"id\":1,\"title\":{\"len\":30},\"required\":1},{\"id\":2,\"img\":{\"type\":1,\"wmin\":300,\"hmin\":300},\"required\":1},{\"id\":3,\"img\":{\"type\":3,\"wmin\":720,\"hmin\":1280},\"required\":1},{\"id\":4,\"data\":{\"type\":1,\"len\":25}},{\"id\":5,\"data\":{\"type\":2,\"len\":35}},{\"id\":6,\"data\":{\"type\":12,\"len\":15}}],\"ver\":\"1.2\",\"context\":1,\"plcmttype\":1}}"
            },
            "ext": {
                "ad_count": 40
            },
            "instl": 0
        }
    ],
    "app": {
        "id": "100461",
        "ver": "4050618",
        "bundle": "com.commsource.beautyplus"
    },
    "device": {
        "ip": "123.49.1.66",
        "geo": {
            "country": "BD",
            "type": 2
        },
        "ua": "Mozilla/5.0 (Linux; Android 11; CPH2269 Build/RP1A.200720.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/94.0.4606.85 Mobile Safari/537.36",
        "devicetype": 4,
        "make": "Xiaomi",
        "brand": "Xiaomi",
        "model": "Redmi 5",
        "os_bit": 64,
        "device_id": "m.F460E2DFE898",
        "os": "android",
        "osv": "8.1.1",
        "language": "en",
        "carrier": "310410",
        "connectiontype": 1,
        "ifa": "3e56c2e-149d-47bb-b978-44fbe63f0a003",
        "didsha1": "ecae9042f3e549a4fd1af27a1ed38b5f682fa3ca",
        "didmd5": "c98de3c3706ae57c89f97c0778c8121d",
        "dpidsha1": "a07713adea29cd23fd10a2e81a4fb07e23e00fb7",
        "dpidmd5": "af4935679d392191157f687f48907492"
    },
    "at": 1,
    "tmax": 2000000,
    "debug": 10101
}
]]
wrk.headers["Content-Type"] = "application/json"