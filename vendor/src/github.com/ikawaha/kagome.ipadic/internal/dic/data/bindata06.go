package data

import (
	"os"
	"time"
)

var _dicIpaUnkDic = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xfd\xdf\xcb\xc4\xc8\xf4\xbf\x8f\x81\xf1\x7f\x0f\x03\x83\xe9\xff\x6e\x66\x46\x46\x56\xdf\xfc\xa2\x82\x0c\x90\x00\x23\x33\x23\x9b\x4f\x6a\x5a\x89\xa7\x0b\x23\x0b\x03\x23\x7b\x50\x66\x7a\x06\x94\xcd\x16\x9e\x0a\xe2\x00\x99\x0c\x0c\xff\x18\x7f\x00\x0d\xd0\x60\xfc\xc7\x25\x05\xc1\x93\x54\x18\x80\x0c\x01\x08\xae\x2a\x02\x71\x84\x20\xb8\x4f\x02\xc4\x91\x80\xe0\x1c\x31\x10\x87\x0b\x82\x33\xd6\x00\x0d\x05\xc2\x7f\xe5\x6c\x48\x9a\x7d\x45\x90\x34\x27\xf5\x21\x69\x76\x8a\x60\x40\x58\x98\x54\x80\x64\x92\x9d\x0d\x03\x23\x17\x10\xfe\x53\x75\x42\x52\x92\x2c\x84\xac\xf9\x17\x92\xb1\xc9\x46\x48\x16\xfa\x45\x21\x99\x64\x3b\x03\xc4\x61\x83\xe0\xb9\xff\x90\x4c\xeb\xde\x82\x64\x5a\x71\x10\x92\x69\xbd\x12\x48\x06\xa4\x49\x41\x3d\xd5\xb2\x0b\xc9\x8e\xe2\x36\x24\x63\xdb\x91\x7d\xd8\x5e\x81\x64\x6c\xd2\x25\x24\x3d\xe5\x75\x48\xc6\x46\xba\x20\x39\xa5\x73\x1a\x88\x23\x07\xc1\xd7\x16\x21\x99\x96\x05\xb3\x3d\xaf\x04\xc9\xa4\x20\x23\x24\x3b\x42\x43\x90\x4c\xca\xda\x85\x64\x87\xe7\x2b\x24\x63\xaf\x7e\x61\x60\x14\x02\xc2\x7f\xae\x7d\x48\x4e\xef\x4c\x62\xe0\xfb\xdf\xce\x02\x4c\x3e\x1d\x0c\x8c\x2c\xa0\xa4\x20\x05\x64\x71\x8b\xf0\x30\x88\x09\x49\xb0\xf9\x71\x31\xf0\x28\xb1\x18\x08\xd8\xf0\xd9\x71\x78\x31\xf9\x08\x00\xa5\xd8\x78\x78\x58\xb8\xf8\xb8\xb8\xb8\x44\x38\x84\x38\x78\xff\xf7\x83\x52\xde\x04\x60\xca\x6b\x63\x60\xe0\xf9\xdf\x0a\xe2\x01\x43\x86\x07\x98\xa0\x58\xcd\x80\xe2\x1a\xec\x6c\x4f\x27\xf4\xbe\x58\x39\x8f\xe7\xe9\xec\x5d\xcf\xe6\x74\x42\x38\x6c\x4f\xe7\x6c\x78\x3a\x7f\x3e\xdb\x93\x1d\x0d\x2f\x3a\xd6\x30\x6a\x81\x20\x76\x85\xc8\x2a\x70\xaa\xd9\xb5\x0b\xc8\x22\xc2\xb0\xe7\x5b\x5b\x9e\xef\x9c\x82\x6e\x18\x9a\x1d\x40\x09\xce\x67\x2d\xf3\x9f\x76\x4f\x05\xca\x21\x04\x07\xd0\x81\x44\x07\x1d\xa6\x4f\xd8\x5e\xac\x98\xf1\xb4\x7f\x3b\x16\x09\x8a\x62\x85\x18\x77\x12\xed\x69\xec\x21\x88\xdb\xc9\x8f\x9b\xb6\x3e\x5d\xd2\xf9\xac\x6f\xe9\xf3\x6d\xb3\x06\x8b\x8f\x68\x92\x86\x48\xf1\x28\x55\x93\x18\xa9\x31\x42\x6c\x90\x33\x3f\x9b\xba\x81\x1c\x8f\x50\x3b\x3f\x52\x3f\x9b\x61\xf7\x1f\x34\xf3\x3d\x5f\xb9\xeb\xf9\xcc\xbd\x44\xa7\x64\x40\x00\x00\x00\xff\xff\x84\xbd\x61\xdb\xce\x07\x00\x00"

func dicIpaUnkDicBytes() ([]byte, error) {
	return bindataRead(
		_dicIpaUnkDic,
		"dic/ipa/unk.dic",
	)
}

func dicIpaUnkDic() (*asset, error) {
	bytes, err := dicIpaUnkDicBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/ipa/unk.dic", size: 1998, mode: os.FileMode(420), modTime: time.Unix(1496122634, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
