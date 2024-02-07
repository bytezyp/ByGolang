#!/usr/bin/env bash
export PATH=/usr/sbin:${PATH}
echo PATH

nowDate=$(date +'%Y-%m-%d')
echo $nowDate

/data/reportapi/bin/reportapi-bin --begin=$nowDate --end=$nowDate --media=ocean --report_type=getAccountDataRealtime --toml_file=/data/reportapi/conf/development.toml



