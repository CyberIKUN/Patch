package src

import (
	"io/ioutil"
	"net/http"
)

func Engine(link string) string {
	request, err := http.NewRequest(http.MethodGet, link, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36")
	request.Header.Set("Cookie", "SECKEY_ABVK=8qVigN2MdEt3uMnKZgkYOBo/nrHUZ45meBbjCkRcoE0%3D; BMAP_SECKEY=KXWGiRf-Ac8b1siXsu7_ilYXz1-D50pieBo9gfN_rzqKhUwPdpagsJbOkLh2WIQMQ3riDm_MNQAdpmBwdhu8Pq1cmqzkmAL5Naq6JADftAImtMLmHH4OBalEpsPefhqjDvQaDGw-JxqG_tqSc_IcBtotn5LggqkQ_AIFpUpibO2YzpRtxpNgFF5KQ4vPxESn; sid=900ad74d-461e-4525-9be5-9c7881f8a1a4; ec=snpnh29i-1644465345762-ae1a9e7a293bb91582976; _exid=kJLcif%2BQTFpazdPHu4lnlcZNJcyaBd%2BwGNob9%2FRtqJGBO%2FJMKxUXAshILB%2BP26h7FGo%2FXgIxB3943amBe%2Bdq5A%3D%3D; _efmdata=902IHJux032ed04zRfl%2BLFD%2BW2fbT03hTXN9jE1JAnAIV3GSf%2F737xb15mgAMCYOHa3KrsmJ%2B6dkThf%2FzMziw3kvUlQM2QrvlaBclDVvjqI%3D; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1644465368; __channelId=905821%2C0; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1644465393")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		buffer, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		result := string(buffer)

		return result
	} else {
		return ""
	}
}
