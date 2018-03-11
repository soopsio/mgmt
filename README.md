# mgmt

测试方式
curl -d '{"inventory":{"package":{"hosts":["10.11.6.160"]}},"playbook":"/package/package.yml","type":"playbook","extra_vars":{"command_path":"common/svn.sh","command_name":"svn代码下载","params":"'ins-test-internal_201801241024103760_300' 'ins-test-internal' 'https://10.10.10.248/svn/coding/dev9/ins-command-internal/trunk'","notice_url":"http://10.10.93.206:8099/TaskController/notice?taskId=ins-test-internal_201801241024103760_300"}}' "http://10.10.99.177:8081/api/v1/tasks?AccessKeyId=MIYiWYTKIlRZjxxhBbWe&SignatureMethod=HMAC-SHA1&SignatureNonce=ansible_201801241024121790_794&SignatureVersion=1.0&Timestamp=2018-01-24T10%3A24%3A12Z&Version=v2&Signature=9nHw%2FMDxQNy%2FRND9UkzftFpFXeU%3D" -H "Content-Type: application/json"

压力测试
hey -n 3000 -c 500 -d '{"inventory":{"package":{"hosts":["10.11.6.160"]}},"playbook":"/package/package.yml","type":"playbook","extra_vars":{"command_path":"common/svn.sh","command_name":"svn代码下载","params":"'ins-test-internal_201801241024103760_300' 'ins-test-internal' 'https://10.10.10.248/svn/coding/dev9/ins-command-internal/trunk'","notice_url":"http://10.10.93.206:8099/TaskController/notice?taskId=ins-test-internal_201801241024103760_300"}}' "http://10.10.99.177:8081/api/v1/tasks?AccessKeyId=MIYiWYTKIlRZjxxhBbWe&SignatureMethod=HMAC-SHA1&SignatureNonce=ansible_201801241024121790_794&SignatureVersion=1.0&Timestamp=2018-01-24T10%3A24%3A12Z&Version=v2&Signature=9nHw%2FMDxQNy%2FRND9UkzftFpFXeU%3D" -H "Content-Type: application/json"

验证
grep 'ins-test-internal_201801241024103760_300' /var/log/mgmt/logfile|jq .task_id > ids.txt

while read id;do redis-cli get ansible_task_id:`echo $id|tr -d '"'`:status ;done < ids.txt  > status.log
while read id;do redis-cli get `echo $id|tr -d '"'`;done < ids.txt  > result.log

