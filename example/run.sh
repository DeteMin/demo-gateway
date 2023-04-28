curl -X POST http://127.0.0.1:8110/v1/trans -d '{"slang":"cmn-Hans-CN","tlang":"en-US","text":""}'
echo "\n"
grpcurl -plaintext -d '{"slang":"cmn-Hans-CN","tlang":"en-US","text":""}'  127.0.0.1:8111  proto.TranslateService/Translate

ab -k -n 100000 -c 100 -p 'post.txt'  http://127.0.0.1:8110/v1/trans