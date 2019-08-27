# Gomodoro
A pomodoro CLI written in go.

I created this repository to practise Go and learn about flags, but I end up learning a little bit more.

To use this project you wont need any special setup, just run the executable and use the flags which are the follow ones:
1. -repeat lets you define how many pomodoros you will do, this flag is setup to be 1 by default. The cli will also add rests after every pomodoro.
2. -duration lets you define the time of your pomodoros, please use the following sintax: "ms", "s", "m", "h". Refer to this link for more insign on Go's duration type (https://golang.org/pkg/time/#Duration).
3. -rest lets you define the time for your rests, this will require the same sintax as the above flag.

