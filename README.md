# ClIliad

A CLI tools to check your stats from Iliad website

## Installation

`go get -u github.com/DLion/cliliad`

## Configuration

Put a `.cliliad.json` file in your `$HOME` dir

```json
{
    "user": "123456",
    "password": "yourpassword"
}
```

## Usage

```
cliliad -help

          oooo   o8o  oooo   o8o                  .o8
          '888   '"'  '888   '"'                 "888
 .ooooo.   888  oooo   888  oooo   .oooo.    .oooo888
d88' '"Y8  888  '888   888  '888  'P  )88b  d88' '888
888        888   888   888   888   .oP"888  888   888
888   .o8  888   888   888   888  d8(  888  888   888
'Y8bod8P' o888o o888o o888o o888o 'Y888""8o 'Y8bod88P"


Usage of cliliad:
  -calls
        Returns the time of the calls done
  -data
        Returns the number of the data sent
  -mms
        Returns the number of the mms sent
  -sms
        Returns the number of the sms sent
```

Retrieve all stats:
```
cliliad

          oooo   o8o  oooo   o8o                  .o8
          '888   '"'  '888   '"'                 "888
 .ooooo.   888  oooo   888  oooo   .oooo.    .oooo888
d88' '"Y8  888  '888   888  '888  'P  )88b  d88' '888
888        888   888   888   888   .oP"888  888   888
888   .o8  888   888   888   888  d8(  888  888   888
'Y8bod8P' o888o o888o o888o o888o 'Y888""8o 'Y8bod88P"


✉️  0 SMS 📤  0 MMS 📞  4m 37s 🌐  499,5MB %
```

## Author

* Domenico Luciani
* https://domenicoluciani.com
* [DLion92](https://twitter.com/DLion92)

# License
MIT
