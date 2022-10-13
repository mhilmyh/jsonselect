## How to setup 
1. Clone this repository to your own local / server 
```
git clone git@github.com:mhilmyh/jsonselect.git
```

2. Build your binary from this repository
```
# this will ouput the jsonselect binary
go build
```
3. Make the jsonselect binary available to your need
```
# move to /usr/local/ folder
mv ./jsonselect /usr/local/jsonselect

# or you can register it to your PATH variable
# every os / shell has their way, figure it your self :)
```
4. Use the binary

## How to use
```
jsonselect [option] [path]
* option:
  -f string
        json file
  -o string
        output file 
  -p bool
        pretty print
* path:
    path of json you want to select.
    please read more in the gjson repository 
    here: https://github.com/tidwall/gjson
```

Here are some examples:

`jsonselect data.user_id < user.json`

`cat user.json | jsonselect data.user_id`

`jsonselect -f user.json -o user.out data.user_id`

`jsonselect -f user.json -o user.out -p=true data.user_id`
