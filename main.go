package hx_sdk

import (
	"fmt"
	"github.com/nbit99/hx_sdk/hx"
)

//Kxw17Y8T11kNrbaY8Y53aXkNvRo8tgYJGZaAYf9bUDBQKkfXXM3z HX5TDS4UrrUTAjmz5sQafYUrM37obZvCEyrVxJHd6teq5wiB7UDA HXNTyhBEVF312RfTyoQ878AhQwerayc7eazr <nil>
func main() {
	fmt.Println("just testing")
	wif, pubkey, addr, error := hx.GetNewPrivate()
	fmt.Println(wif, pubkey, addr, error)
	ref_info := hx.CalRefInfo("0021dd2d8f2ce56feb75c79614effdee4313bf22")
	fmt.Println(ref_info)                       //HXNXX9CLDUmh825kXzynvwx4azan2HKaL8aQ
	trx_data, err := hx.HxTransfer(ref_info, "Kxw17Y8T11kNrbaY8Y53aXkNvRo8tgYJGZaAYf9bUDBQKkfXXM3z", "08d1d10092bbdbb68c1613c93ded434805381fe73e845c59b5a97693fa1a778e", "HXNTyhBEVF312RfTyoQ878AhQwerayc7eazr", "HXNTyhBEVF312RfTyoQ878AhQwerayc7eazr", "HX", "0.11", "0.001", "aaaa", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("")
	fmt.Println("raw trx", string(trx_data))
	trx_data, err = hx.HxRegister(ref_info, "Kxw17Y8T11kNrbaY8Y53aXkNvRo8tgYJGZaAYf9bUDBQKkfXXM3z", "08d1d10092bbdbb68c1613c93ded434805381fe73e845c59b5a97693fa1a778e", "newtest", "HXNTyhBEVF312RfTyoQ878AhQwerayc7eazr", "5.001", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("")
	fmt.Println("register raw trx", string(trx_data))
	trx_data, err = hx.HxMining(ref_info, "Kxw17Y8T11kNrbaY8Y53aXkNvRo8tgYJGZaAYf9bUDBQKkfXXM3z", "08d1d10092bbdbb68c1613c93ded434805381fe73e845c59b5a97693fa1a778e", "HX", "1.2.105", "HXNTyhBEVF312RfTyoQ878AhQwerayc7eazr", "1", "0", "1.6.1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("")
	fmt.Println("HxMining raw trx", string(trx_data))
	trx_data, err = hx.HxForecloseBalance(ref_info, "Kxw17Y8T11kNrbaY8Y53aXkNvRo8tgYJGZaAYf9bUDBQKkfXXM3z", "08d1d10092bbdbb68c1613c93ded434805381fe73e845c59b5a97693fa1a778e", "HXNTyhBEVF312RfTyoQ878AhQwerayc7eazr", "1.2.105", "1.3.0", "1.6.1", "1", "0")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("")
	fmt.Println("HxForecloseBalance raw trx", string(trx_data))
	//hx.HxBind("")
}
