package public_key_handler

import (
	"math/big"
	"testing"
)

func TestJWKHandler_decodeStringToUint64(t *testing.T) {
	encodedE := "AQAB"
	decodedE, err := decodeStringToUint64(encodedE)
	if err != nil {
		t.Error(err)
	}
	if want := uint64(65537); decodedE != want {
		t.Errorf("want %v, got %v", want, decodedE)
	}
}

func TestJWKHandler_decodeN(t *testing.T) {
	encodedN := "kaHdd1aU4kXXKYgQEvXiBuLFMJ7YXeLbQu7UJ8HCZNPnXv9nQkHevVa2Gwtbc6SWSenXYz4HArxDHHNlEpkhbF8Bxs92K-jAJgeoNDpq1QUFJYi2ey4-itCxMUU-IJGGspDRPnTZTOgmqHaOMaYRopzOaCSR7uHrocmMP0eFG5tyXo0cRLe30cX6U2rGOaVmExZCLrVWVtPi23_G4yGqU47RJ7xOTgnz7DtuIf82ImrUui6JVyKv1-wILsLnYmdjy7E0-WhoTa6t0CQKFzNBOYYaP8jiRIvcU9TJrS4fplkA-BgS8J_-Nky9d82NXgwDDV94vagTXsqI9GpKeKNGHfWWneiWVpqHhnuO4XxKW9ZeD1gaincaAz9mOll8oDarBtQV5q3vexeS1vRcXHlrzVbE9byGOofkZElLJoLrtv0WSNtYFUNjE5HusaZiyD1-aX85ZpiKeJFBlxo87ofFxEVOkrryDpEbEWj05UkZ4euNkR8zM8ZVcZg9maTNyRA-aC7O_ZU0zS5omVZxMdVVoQ5pZ02KvLehtBIVEIsQJbU6YmL7AQp_DTmFDbpnHMnEpLZr_R0ihTUfuzuQiyLReYHnwPPi8He3GQ4aVrDRM-xeXM6aE8ZDZP9WtQvmArNxAJxFsXrzJwTkyzcqI6WB5_Js_ZuJdb6-4maTVE4B2i0"
	decodedN, err := decodeStringToBytes(encodedN)
	if err != nil {
		t.Error(err)
	}

	var keyN big.Int
	keyN.SetBytes(decodedN)

	var want big.Int
	want.SetString(
		"594127889878685891478456104333332190432471737764406000539452636886042975021921315023399239382820240539920850652086172303118302121513683770856767635966396529616981953840094870260906436996498196630951175902954173995128651760367044675505077951658768145067373847946442892346553246467831551141033906088080735057525141872550063591176899928703029639815850642806627020247756125468273867620261412025850877657399965044701582445388910253579692694922453992481062579650621348536122957486849919111043632919795242042819819988370745469592287099228359891374596725447378354900663085724814267232703606903189475615570531423457763584586174249555849071942510438511792286965565180586084976121610054808256641095315044534619850229593459850862091041347353889328346963572629299831739227735207225724986044994209540507386212755001338554794796473093816972196952903942333297350404355824059062016192918730765588560815687594350946333089902516125664077030594976464025785403149852596239383936744720087615002046439505723959037452793407913628815843592475195493004967970255563734257982303248686998759420898370227027696510207040148231417097108875665670939543324321233459159766538944805746490433220380788572800126991841457345653756452768786913863080610807195302730724923949",
		0,
	)
	if keyN.Cmp(&want) != 0 {
		t.Errorf("want %v, got %v", want, keyN)
	}
}

func TestJWKHandler_getECDSAPublicKeyFromJWK(t *testing.T) {
	handler := NewPublicKeyHandler()

	jwk := ECDSAJSONWebKey{
		Kty: "EC",
		Crv: "P-256",
		Kid: "tn1AViVj7vhk4TrdghT8Mw==",
		X:   "sC1IpRQTKG3a_ULMXQZmP95vXUg3qWq1wUy_qIedfBU",
		Y:   "74SsB5DCdaML6rt99v0DVPRIoGh0WR3G8mxYUr4uUtM",
	}

	key, err := handler.getECDSAPublicKeyFromJWK(jwk)
	if err != nil {
		t.Error(err)
	}

	wantCrv := "P-256"
	if got := key.Key.Curve.Params().Name; got != wantCrv {
		t.Errorf("want %v, got %v", wantCrv, got)
	}

	var wantX big.Int
	wantX.SetString(
		"79687070844812207596098443007393002297593322043514267675772315167913943071765",
		0,
	)
	if got := key.Key.X; got.Cmp(&wantX) != 0 {
		t.Errorf("want %v, got %v", wantX, got)
	}

	var wantY big.Int
	wantY.SetString(
		"108337181928287654206268663555247825023969395620689137485556592744914024813267",
		0,
	)
	if got := key.Key.Y; got.Cmp(&wantY) != 0 {
		t.Errorf("want %v, got %v", wantY, got)
	}
}

func TestJWKHandler_getECDSAPublicKeyMapFromJWKs(t *testing.T) {
	handler := NewPublicKeyHandler()

	jwks := ECDSAJWKs{
		{
			Kty: "EC",
			Crv: "P-256",
			Kid: "tn1AViVj7vhk4TrdghT8Mw==",
			X:   "sC1IpRQTKG3a_ULMXQZmP95vXUg3qWq1wUy_qIedfBU",
			Y:   "74SsB5DCdaML6rt99v0DVPRIoGh0WR3G8mxYUr4uUtM",
		},
	}

	keyMap, err := handler.getECDSAPublicKeyMapFromJWKs(jwks)
	if err != nil {
		t.Error(err)
	}

	_, ok := keyMap["tn1AViVj7vhk4TrdghT8Mw=="]
	if !ok {
		t.Error("key not found in got keyMap")
	}
}
