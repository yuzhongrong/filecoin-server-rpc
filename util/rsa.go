package util

import (
	"encoding/hex"
	"github.com/wumansgy/goEncrypt"
)

const (
	privateKey = `-----BEGIN  WUMAN RSA PRIVATE KEY -----
MIIEogIBAAKCAQEA1U74LWexD/qxGl69ZyepYN/yNMcvDlbBBCDqMGDonibci/cj
uT0RIZC1WplxYBzwgpTpUH7aePie2wahmRCel3eiDEsjY7UlLaI0pKFxZZDYdS1A
dmOInLGG3SPJQwzOsBM03syOvxXuDMQlZ7dmN/I/G7fldSjTGetkgaO8RWXr0JfQ
7byB9/9cpdg4rOkQs9PgDJ/K7+OiPdbawqPN+VvOg7V6jyDrfmllrCdOzV+17F2G
dvmGmzf1G3twh63+AhrL8upMbOZ0g26hlqYJ8mE0TnppmvcnNWUjoenlnIQDgggI
MNnZgZ65HshMzNDeENSgc2vmQarSxQ2gZIN3CQIDAQABAoIBAEIORW6SKNPg901K
P28daid00mWjtR/En9sucjdvGzo2oJ+7ddWcYpy5Wl/nGqP/8j8N7D6gOfmyTEdZ
g1uKOQKA6q7R6fCrnQrHq3O8BwjD1TRcQhUnd5vGohQDTAU2hx8ho4LHaAEmwmQz
rb2znrT+kKp9xVIFxXHGYoZ+9QOsHp/dlSS8v3mpmAdhKNSqC60MtLUEOCesDa7Z
zZXgYgNybNey9qf5ppIcx0t/8WW3vaY2r5Ga7DdIDdqwA34SZ2KTkVIQNagJue1f
OTcLcUlCw7/3X0vxOllZWc8yCkz4kWVTfSjVEYi6m4zsnwkFswxf+KenIFaWYhIg
PWSGiw0CgYEA5vlz0yfHDsVZSyFsa9PxobedwrrVpWzpDGXaqV/FtWTf9/SI9eC/
7lnq+LBLZGNslB9cG/dZWWLM6+ZWIeuKrztw/4QsjNAKVuaMTVO3q7bx31YUbnsJ
8rZsWxXE6mhEaQFoo2It0AFM7RwHDY4gEf9cxt6+vaZLgVvFau2lFm8CgYEA7GuE
C98jOSskfUDYh5yffB14M7mxDiOtiO2f4OnMoIkGVeh4b+giXcftpny+zdKEgMwH
sBk+pduJyXw24EMVtEFlxdpOZWhEeKV8Ityi6EUyi5EdJQyrzE9lUwqqCh0XRamR
pqxb9YxtV5186vaHkuncMYepXKfZ3eYhhye6xgcCgYA11nd1BJSVkNKbfJL1H1X3
SAFx3nLmOFiqFyO8zyIaggTimxFBnr2eJT9r8EvifnpUnGqv6hvdhfYWFn6FMY4G
Amj4ZiqsN+HxF5QkghsR33bJhBsHFY7gED15jb10lhE8GKP3UW80SNlRe3L5aeN3
znolsM3tDtISuP1vSy1r6QKBgFSLLjFAnkv3TZks80GrlKzBrRZyNQqlX40zzJSV
hwNxfL6D323FSWTX9fgva9wWiCO20pj6rhiJpYBT1xvjYYOQT2CtbJN+8d1i8D1X
QTpmZGjcf9ub6GOrkMRdb+kl9giHVvqPcGMi0IcgXmd3uYuj7YYYyUvFCnf7r8mx
P21JAoGACEof0m822Cx1nFjFEHq+fonN/59AQfpP83eFe2XAk2FwCn0Xpfz0vr6M
/D/62ooLeom8XsGV9Z9QkCSUO196Svj+uNYtwSO6gjr6Yw0y2u3mOlrnODKPdpT6
7QIH9wM97mekOcqcIdY0aUUWs3yPh+uha0k+kBgpLQ4u37grWK4=
-----END  WUMAN RSA PRIVATE KEY -----`
	publicKey = `-----BEGIN  WUMAN  RSA PUBLIC KEY -----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1U74LWexD/qxGl69Zyep
YN/yNMcvDlbBBCDqMGDonibci/cjuT0RIZC1WplxYBzwgpTpUH7aePie2wahmRCe
l3eiDEsjY7UlLaI0pKFxZZDYdS1AdmOInLGG3SPJQwzOsBM03syOvxXuDMQlZ7dm
N/I/G7fldSjTGetkgaO8RWXr0JfQ7byB9/9cpdg4rOkQs9PgDJ/K7+OiPdbawqPN
+VvOg7V6jyDrfmllrCdOzV+17F2GdvmGmzf1G3twh63+AhrL8upMbOZ0g26hlqYJ
8mE0TnppmvcnNWUjoenlnIQDgggIMNnZgZ65HshMzNDeENSgc2vmQarSxQ2gZIN3
CQIDAQAB
-----END  WUMAN  RSA PUBLIC KEY -----`
)

func GenerateKey()  {
	goEncrypt.GetRsaKey()//本地生成privatekey and  publickey file
}

func Encrypt(origin string) string{
	result,_:=goEncrypt.RsaEncrypt([]byte(origin),[]byte(publicKey))
	hexStr:=hex.EncodeToString(result)
	return hexStr
}


func Decrypt(encrypt string) string{
	decryptStr,_:=hex.DecodeString(encrypt)
	result,_:=goEncrypt.RsaDecrypt(decryptStr,[]byte(privateKey))
	return string(result)
}



