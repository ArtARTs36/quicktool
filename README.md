# quicktool

quicktool - set of development mini tools.

## Commands

| Name           | Description           | Args           |
|----------------|-----------------------|----------------|
| uuid | Generate UUID |  |
| md5 | Encode value to md5 |  `value` - source value <br/> |
| jsonpretty | Pretty json |  `value` - source value <br/> |
| time | Get time |  |
| mkdirs | Make directories |  `range` - Range, example "0-20" <br/> `fileMask` - Mask to folder name, example "folder{number}" <br/> |
| jsonpath | Get value from json by path |  `path` - json path <br/> `source` - json path (string of content, file path) <br/> |
| user | Get user data |  |
| password | Generate password |  `length` - length of password <br/> |
| base64 | Base64 encode/decode |  `action` - action <br/> |
| range | generate range |  `length` - Range, example: 1-10 <br/> `delta` - delta value, example: +2 <br/> |
| gpg | gpg |  `action` -  <br/> `key-id` -  <br/> |
