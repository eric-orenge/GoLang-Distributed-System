brew install nsq

nsqlookupd

Open and maintain a connection to Twitter's streaming APIs looking for
any mention of the options

nsqlookupd



in one shell, start nsqlookupd:

$ nsqlookupd

in another shell, start nsqd:

$ nsqd --lookupd-tcp-address=127.0.0.1:4160

in another shell, start nsqadmin:

$ nsqadmin --lookupd-http-address=127.0.0.1:4161

publish an initial message (creates the topic in the cluster, too):

$ curl -d 'hello world 1' 'http://127.0.0.1:4151/pub?topic=test'

http://127.0.0.1:4171