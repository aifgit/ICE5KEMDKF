/*
 * You have extremely bursty traffic on your network. To gain better insight
 * into the problem your engineering team has purchased a tool capable of
 * performing 1 millisecond monitoring.
 *
 * This tool takes Byte count samples of a link at 1ms intervals and sends out
 * updates every 1sec. In other words, the tool perform 1000 samples per second
 * but only emits information about the peak data once every second. It is
 * your task to decode this data and display some useful statistics about it.
 *
 * Further guidance:
 * In this project a microburst (scaled peak) is calculated by taking the 1ms
 * peak data and converting it to 1sec data. E.g. 20 Bytes in a ms would be a
 * 160 Kbps microburst.
 *
 * Watch your units - Bytes, bits, seconds, milliseconds...
 *
 */

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//To kepp this simple - pretend that the tool has a JSON API and this array
	//of JSON objects is the response to your HTTP request to that API.
	sampleData := []byte(`[{"epoch":1458151114,"offset":139,"bytes":659914},{"epoch":1458151115,"offset":256,"bytes":315929},{"epoch":1458151116,"offset":293,"bytes":1226342},{"epoch":1458151117,"offset":591,"bytes":44767},{"epoch":1458151118,"offset":437,"bytes":350451},{"epoch":1458151119,"offset":299,"bytes":1041392},{"epoch":1458151120,"offset":544,"bytes":740981},{"epoch":1458151121,"offset":958,"bytes":866369},{"epoch":1458151122,"offset":167,"bytes":14337},{"epoch":1458151123,"offset":864,"bytes":707923},{"epoch":1458151124,"offset":935,"bytes":446217},{"epoch":1458151125,"offset":404,"bytes":822836},{"epoch":1458151126,"offset":313,"bytes":82751},{"epoch":1458151127,"offset":627,"bytes":1014985},{"epoch":1458151128,"offset":261,"bytes":727751},{"epoch":1458151129,"offset":162,"bytes":816517},{"epoch":1458151130,"offset":113,"bytes":1202768},{"epoch":1458151131,"offset":478,"bytes":612443},{"epoch":1458151132,"offset":320,"bytes":74849},{"epoch":1458151133,"offset":177,"bytes":980455},{"epoch":1458151134,"offset":423,"bytes":454115},{"epoch":1458151135,"offset":696,"bytes":490218},{"epoch":1458151136,"offset":362,"bytes":1185284},{"epoch":1458151137,"offset":424,"bytes":59379},{"epoch":1458151138,"offset":705,"bytes":739385},{"epoch":1458151139,"offset":543,"bytes":591795},{"epoch":1458151140,"offset":920,"bytes":1164537},{"epoch":1458151141,"offset":713,"bytes":791824},{"epoch":1458151142,"offset":603,"bytes":171989},{"epoch":1458151143,"offset":629,"bytes":821781},{"epoch":1458151144,"offset":474,"bytes":338548},{"epoch":1458151145,"offset":328,"bytes":238712},{"epoch":1458151146,"offset":189,"bytes":815170},{"epoch":1458151147,"offset":363,"bytes":156409},{"epoch":1458151148,"offset":577,"bytes":791494},{"epoch":1458151149,"offset":362,"bytes":904194},{"epoch":1458151150,"offset":844,"bytes":1150977},{"epoch":1458151151,"offset":717,"bytes":143031},{"epoch":1458151152,"offset":195,"bytes":735485},{"epoch":1458151153,"offset":348,"bytes":762727},{"epoch":1458151154,"offset":342,"bytes":980912},{"epoch":1458151155,"offset":176,"bytes":173992},{"epoch":1458151156,"offset":141,"bytes":637096},{"epoch":1458151157,"offset":143,"bytes":584493},{"epoch":1458151158,"offset":712,"bytes":126895},{"epoch":1458151159,"offset":962,"bytes":990452},{"epoch":1458151160,"offset":348,"bytes":434155},{"epoch":1458151161,"offset":265,"bytes":595064},{"epoch":1458151162,"offset":565,"bytes":336235},{"epoch":1458151163,"offset":423,"bytes":599673},{"epoch":1458151164,"offset":588,"bytes":967172},{"epoch":1458151165,"offset":810,"bytes":179392},{"epoch":1458151166,"offset":427,"bytes":829249},{"epoch":1458151167,"offset":773,"bytes":270109},{"epoch":1458151168,"offset":503,"bytes":761710},{"epoch":1458151169,"offset":676,"bytes":391889},{"epoch":1458151170,"offset":774,"bytes":328144},{"epoch":1458151171,"offset":347,"bytes":340057},{"epoch":1458151172,"offset":485,"bytes":993455},{"epoch":1458151173,"offset":899,"bytes":398077},{"epoch":1458151174,"offset":640,"bytes":370301},{"epoch":1458151175,"offset":906,"bytes":241853},{"epoch":1458151176,"offset":207,"bytes":1138681},{"epoch":1458151177,"offset":153,"bytes":27703},{"epoch":1458151178,"offset":942,"bytes":98173},{"epoch":1458151179,"offset":220,"bytes":35397},{"epoch":1458151180,"offset":922,"bytes":240527},{"epoch":1458151181,"offset":279,"bytes":1131974},{"epoch":1458151182,"offset":985,"bytes":133825},{"epoch":1458151183,"offset":691,"bytes":577679},{"epoch":1458151184,"offset":928,"bytes":153748},{"epoch":1458151185,"offset":746,"bytes":803877},{"epoch":1458151186,"offset":508,"bytes":246946},{"epoch":1458151187,"offset":165,"bytes":1062693},{"epoch":1458151188,"offset":623,"bytes":882101},{"epoch":1458151189,"offset":628,"bytes":1035428},{"epoch":1458151190,"offset":774,"bytes":689566},{"epoch":1458151191,"offset":104,"bytes":809136},{"epoch":1458151192,"offset":209,"bytes":471921},{"epoch":1458151193,"offset":367,"bytes":1152837},{"epoch":1458151194,"offset":916,"bytes":769784},{"epoch":1458151195,"offset":578,"bytes":675706},{"epoch":1458151196,"offset":631,"bytes":1134207},{"epoch":1458151197,"offset":541,"bytes":128501},{"epoch":1458151198,"offset":538,"bytes":699665},{"epoch":1458151199,"offset":324,"bytes":1134616},{"epoch":1458151200,"offset":973,"bytes":1184776},{"epoch":1458151201,"offset":568,"bytes":391557},{"epoch":1458151202,"offset":965,"bytes":78877},{"epoch":1458151203,"offset":839,"bytes":771703},{"epoch":1458151204,"offset":288,"bytes":1102783},{"epoch":1458151205,"offset":449,"bytes":385522},{"epoch":1458151206,"offset":511,"bytes":852061},{"epoch":1458151207,"offset":508,"bytes":1216015},{"epoch":1458151208,"offset":190,"bytes":222670},{"epoch":1458151209,"offset":516,"bytes":61848},{"epoch":1458151210,"offset":950,"bytes":558370},{"epoch":1458151211,"offset":295,"bytes":752045},{"epoch":1458151212,"offset":973,"bytes":648167},{"epoch":1458151213,"offset":594,"bytes":101065},{"epoch":1458151214,"offset":615,"bytes":84316},{"epoch":1458151215,"offset":305,"bytes":1215296},{"epoch":1458151216,"offset":367,"bytes":1110095},{"epoch":1458151217,"offset":784,"bytes":1159181},{"epoch":1458151218,"offset":602,"bytes":616466},{"epoch":1458151219,"offset":335,"bytes":1103074},{"epoch":1458151220,"offset":587,"bytes":1140972},{"epoch":1458151221,"offset":570,"bytes":1199275},{"epoch":1458151222,"offset":931,"bytes":314213},{"epoch":1458151223,"offset":861,"bytes":962162},{"epoch":1458151224,"offset":574,"bytes":1030190},{"epoch":1458151225,"offset":592,"bytes":840628},{"epoch":1458151226,"offset":381,"bytes":132515},{"epoch":1458151227,"offset":490,"bytes":971173},{"epoch":1458151228,"offset":621,"bytes":901157},{"epoch":1458151229,"offset":422,"bytes":266042},{"epoch":1458151230,"offset":795,"bytes":975107},{"epoch":1458151231,"offset":533,"bytes":246653},{"epoch":1458151232,"offset":175,"bytes":549949},{"epoch":1458151233,"offset":860,"bytes":844015},{"epoch":1458151234,"offset":682,"bytes":87744},{"epoch":1458151235,"offset":291,"bytes":271554},{"epoch":1458151236,"offset":405,"bytes":73677},{"epoch":1458151237,"offset":970,"bytes":1008271},{"epoch":1458151238,"offset":496,"bytes":820261},{"epoch":1458151239,"offset":614,"bytes":593392},{"epoch":1458151240,"offset":343,"bytes":1103405},{"epoch":1458151241,"offset":961,"bytes":172078},{"epoch":1458151242,"offset":311,"bytes":589861},{"epoch":1458151243,"offset":921,"bytes":534633},{"epoch":1458151244,"offset":278,"bytes":1233207},{"epoch":1458151245,"offset":476,"bytes":490474},{"epoch":1458151246,"offset":509,"bytes":537385},{"epoch":1458151247,"offset":355,"bytes":215595},{"epoch":1458151248,"offset":344,"bytes":327410},{"epoch":1458151249,"offset":194,"bytes":1073025},{"epoch":1458151250,"offset":211,"bytes":187529},{"epoch":1458151251,"offset":781,"bytes":937953},{"epoch":1458151252,"offset":222,"bytes":779101},{"epoch":1458151253,"offset":236,"bytes":1015820},{"epoch":1458151254,"offset":738,"bytes":304211},{"epoch":1458151255,"offset":755,"bytes":47140},{"epoch":1458151256,"offset":653,"bytes":24809},{"epoch":1458151257,"offset":567,"bytes":709491},{"epoch":1458151258,"offset":136,"bytes":384692},{"epoch":1458151259,"offset":114,"bytes":578746},{"epoch":1458151260,"offset":163,"bytes":689400},{"epoch":1458151261,"offset":965,"bytes":194810},{"epoch":1458151262,"offset":875,"bytes":924027},{"epoch":1458151263,"offset":864,"bytes":724925},{"epoch":1458151264,"offset":122,"bytes":463802},{"epoch":1458151265,"offset":634,"bytes":244036},{"epoch":1458151266,"offset":515,"bytes":839179},{"epoch":1458151267,"offset":201,"bytes":874305},{"epoch":1458151268,"offset":191,"bytes":1247304},{"epoch":1458151269,"offset":380,"bytes":808422},{"epoch":1458151270,"offset":782,"bytes":222440},{"epoch":1458151271,"offset":788,"bytes":85329},{"epoch":1458151272,"offset":946,"bytes":481239},{"epoch":1458151273,"offset":424,"bytes":1025420},{"epoch":1458151274,"offset":754,"bytes":101581},{"epoch":1458151275,"offset":379,"bytes":864653},{"epoch":1458151276,"offset":896,"bytes":13085},{"epoch":1458151277,"offset":946,"bytes":983531},{"epoch":1458151278,"offset":574,"bytes":478153},{"epoch":1458151279,"offset":900,"bytes":24462},{"epoch":1458151280,"offset":562,"bytes":1184415},{"epoch":1458151281,"offset":715,"bytes":66987},{"epoch":1458151282,"offset":777,"bytes":1128507},{"epoch":1458151283,"offset":942,"bytes":402807},{"epoch":1458151284,"offset":987,"bytes":367426},{"epoch":1458151285,"offset":181,"bytes":137597},{"epoch":1458151286,"offset":331,"bytes":1242521},{"epoch":1458151287,"offset":713,"bytes":484821},{"epoch":1458151288,"offset":330,"bytes":991903},{"epoch":1458151289,"offset":123,"bytes":262814},{"epoch":1458151290,"offset":685,"bytes":367559},{"epoch":1458151291,"offset":310,"bytes":699541},{"epoch":1458151292,"offset":153,"bytes":1074583},{"epoch":1458151293,"offset":590,"bytes":710477},{"epoch":1458151294,"offset":932,"bytes":493525},{"epoch":1458151295,"offset":225,"bytes":213598},{"epoch":1458151296,"offset":822,"bytes":430639},{"epoch":1458151297,"offset":732,"bytes":1008299},{"epoch":1458151298,"offset":974,"bytes":60550},{"epoch":1458151299,"offset":792,"bytes":94437},{"epoch":1458151300,"offset":957,"bytes":539332},{"epoch":1458151301,"offset":843,"bytes":24841},{"epoch":1458151302,"offset":860,"bytes":302539},{"epoch":1458151303,"offset":466,"bytes":905876},{"epoch":1458151304,"offset":103,"bytes":960547},{"epoch":1458151305,"offset":630,"bytes":42136},{"epoch":1458151306,"offset":503,"bytes":251945},{"epoch":1458151307,"offset":833,"bytes":303917},{"epoch":1458151308,"offset":102,"bytes":529454},{"epoch":1458151309,"offset":748,"bytes":584023},{"epoch":1458151310,"offset":673,"bytes":1230935},{"epoch":1458151311,"offset":903,"bytes":202401},{"epoch":1458151312,"offset":669,"bytes":1182189},{"epoch":1458151313,"offset":434,"bytes":73768},{"epoch":1458151314,"offset":879,"bytes":1145014},{"epoch":1458151315,"offset":121,"bytes":143347},{"epoch":1458151316,"offset":287,"bytes":737473},{"epoch":1458151317,"offset":660,"bytes":900495},{"epoch":1458151318,"offset":878,"bytes":232443},{"epoch":1458151319,"offset":924,"bytes":322570},{"epoch":1458151320,"offset":877,"bytes":578053},{"epoch":1458151321,"offset":866,"bytes":1232465},{"epoch":1458151322,"offset":950,"bytes":163353},{"epoch":1458151323,"offset":601,"bytes":1118370},{"epoch":1458151324,"offset":422,"bytes":392037},{"epoch":1458151325,"offset":600,"bytes":1205156},{"epoch":1458151326,"offset":509,"bytes":665185},{"epoch":1458151327,"offset":646,"bytes":592873},{"epoch":1458151328,"offset":188,"bytes":290069},{"epoch":1458151329,"offset":142,"bytes":905885},{"epoch":1458151330,"offset":515,"bytes":412379},{"epoch":1458151331,"offset":905,"bytes":1145205},{"epoch":1458151332,"offset":755,"bytes":1058355},{"epoch":1458151333,"offset":129,"bytes":989118},{"epoch":1458151334,"offset":595,"bytes":1109938},{"epoch":1458151335,"offset":185,"bytes":728243},{"epoch":1458151336,"offset":376,"bytes":860981},{"epoch":1458151337,"offset":734,"bytes":393794},{"epoch":1458151338,"offset":580,"bytes":1033746},{"epoch":1458151339,"offset":765,"bytes":259321},{"epoch":1458151340,"offset":279,"bytes":589946},{"epoch":1458151341,"offset":965,"bytes":288319},{"epoch":1458151342,"offset":626,"bytes":421168},{"epoch":1458151343,"offset":924,"bytes":861236},{"epoch":1458151344,"offset":389,"bytes":104711},{"epoch":1458151345,"offset":531,"bytes":1187019},{"epoch":1458151346,"offset":968,"bytes":768653},{"epoch":1458151347,"offset":810,"bytes":488408},{"epoch":1458151348,"offset":812,"bytes":1142656},{"epoch":1458151349,"offset":261,"bytes":24188},{"epoch":1458151350,"offset":635,"bytes":806718},{"epoch":1458151351,"offset":112,"bytes":348537},{"epoch":1458151352,"offset":546,"bytes":224843},{"epoch":1458151353,"offset":268,"bytes":63707},{"epoch":1458151354,"offset":789,"bytes":356045},{"epoch":1458151355,"offset":439,"bytes":1021699},{"epoch":1458151356,"offset":737,"bytes":1229451},{"epoch":1458151357,"offset":382,"bytes":866579},{"epoch":1458151358,"offset":884,"bytes":107029},{"epoch":1458151359,"offset":582,"bytes":527120},{"epoch":1458151360,"offset":623,"bytes":488715},{"epoch":1458151361,"offset":745,"bytes":279765},{"epoch":1458151362,"offset":854,"bytes":1145299},{"epoch":1458151363,"offset":484,"bytes":879906},{"epoch":1458151364,"offset":665,"bytes":376477},{"epoch":1458151365,"offset":927,"bytes":1197858},{"epoch":1458151366,"offset":905,"bytes":865651},{"epoch":1458151367,"offset":106,"bytes":408828},{"epoch":1458151368,"offset":460,"bytes":727105},{"epoch":1458151369,"offset":197,"bytes":220356},{"epoch":1458151370,"offset":651,"bytes":278400},{"epoch":1458151371,"offset":812,"bytes":189771},{"epoch":1458151372,"offset":870,"bytes":200344},{"epoch":1458151373,"offset":771,"bytes":1930},{"epoch":1458151374,"offset":285,"bytes":184653},{"epoch":1458151375,"offset":427,"bytes":593598},{"epoch":1458151376,"offset":161,"bytes":1162313},{"epoch":1458151377,"offset":178,"bytes":37525},{"epoch":1458151378,"offset":424,"bytes":600923},{"epoch":1458151379,"offset":401,"bytes":798433},{"epoch":1458151380,"offset":353,"bytes":1153505},{"epoch":1458151381,"offset":795,"bytes":186339},{"epoch":1458151382,"offset":672,"bytes":583841},{"epoch":1458151383,"offset":458,"bytes":1184484},{"epoch":1458151384,"offset":847,"bytes":733062},{"epoch":1458151385,"offset":529,"bytes":542530},{"epoch":1458151386,"offset":218,"bytes":830785},{"epoch":1458151387,"offset":296,"bytes":70035},{"epoch":1458151388,"offset":518,"bytes":984608},{"epoch":1458151389,"offset":615,"bytes":180128},{"epoch":1458151390,"offset":937,"bytes":228825},{"epoch":1458151391,"offset":143,"bytes":927201},{"epoch":1458151392,"offset":784,"bytes":476579},{"epoch":1458151393,"offset":946,"bytes":924378},{"epoch":1458151394,"offset":135,"bytes":32536},{"epoch":1458151395,"offset":897,"bytes":940627},{"epoch":1458151396,"offset":739,"bytes":628493},{"epoch":1458151397,"offset":496,"bytes":1131501},{"epoch":1458151398,"offset":298,"bytes":413210},{"epoch":1458151399,"offset":699,"bytes":998888},{"epoch":1458151400,"offset":697,"bytes":497583},{"epoch":1458151401,"offset":258,"bytes":949681},{"epoch":1458151402,"offset":616,"bytes":761208},{"epoch":1458151403,"offset":585,"bytes":7996},{"epoch":1458151404,"offset":975,"bytes":928839},{"epoch":1458151405,"offset":625,"bytes":474243},{"epoch":1458151406,"offset":226,"bytes":83673},{"epoch":1458151407,"offset":852,"bytes":842853},{"epoch":1458151408,"offset":122,"bytes":102961},{"epoch":1458151409,"offset":119,"bytes":1156158},{"epoch":1458151410,"offset":457,"bytes":1042955},{"epoch":1458151411,"offset":302,"bytes":1070543},{"epoch":1458151412,"offset":911,"bytes":1154962},{"epoch":1458151413,"offset":421,"bytes":16334},{"epoch":1458151414,"offset":255,"bytes":912407},{"epoch":1458151415,"offset":973,"bytes":981361},{"epoch":1458151416,"offset":767,"bytes":793843},{"epoch":1458151417,"offset":332,"bytes":918426},{"epoch":1458151418,"offset":338,"bytes":279910},{"epoch":1458151419,"offset":967,"bytes":446198},{"epoch":1458151420,"offset":373,"bytes":38645},{"epoch":1458151421,"offset":539,"bytes":421111},{"epoch":1458151422,"offset":700,"bytes":1135358},{"epoch":1458151423,"offset":849,"bytes":1030534},{"epoch":1458151424,"offset":571,"bytes":592443},{"epoch":1458151425,"offset":332,"bytes":830211},{"epoch":1458151426,"offset":618,"bytes":612067},{"epoch":1458151427,"offset":324,"bytes":1156987},{"epoch":1458151428,"offset":194,"bytes":42624},{"epoch":1458151429,"offset":187,"bytes":800905},{"epoch":1458151430,"offset":400,"bytes":802581},{"epoch":1458151431,"offset":294,"bytes":83473},{"epoch":1458151432,"offset":539,"bytes":493931},{"epoch":1458151433,"offset":714,"bytes":319282},{"epoch":1458151434,"offset":771,"bytes":689151},{"epoch":1458151435,"offset":237,"bytes":1176771},{"epoch":1458151436,"offset":242,"bytes":591041},{"epoch":1458151437,"offset":164,"bytes":1068221},{"epoch":1458151438,"offset":488,"bytes":183426},{"epoch":1458151439,"offset":886,"bytes":417804},{"epoch":1458151440,"offset":611,"bytes":1126929},{"epoch":1458151441,"offset":168,"bytes":743169},{"epoch":1458151442,"offset":394,"bytes":134659},{"epoch":1458151443,"offset":775,"bytes":853521},{"epoch":1458151444,"offset":309,"bytes":675473},{"epoch":1458151445,"offset":578,"bytes":186323},{"epoch":1458151446,"offset":145,"bytes":296851},{"epoch":1458151447,"offset":742,"bytes":812950},{"epoch":1458151448,"offset":556,"bytes":701719},{"epoch":1458151449,"offset":469,"bytes":480992},{"epoch":1458151450,"offset":822,"bytes":10607},{"epoch":1458151451,"offset":564,"bytes":845111},{"epoch":1458151452,"offset":798,"bytes":409586},{"epoch":1458151453,"offset":896,"bytes":589434},{"epoch":1458151454,"offset":197,"bytes":1051654},{"epoch":1458151455,"offset":143,"bytes":953554},{"epoch":1458151456,"offset":376,"bytes":1114369},{"epoch":1458151457,"offset":194,"bytes":192498},{"epoch":1458151458,"offset":211,"bytes":278955},{"epoch":1458151459,"offset":664,"bytes":667770},{"epoch":1458151460,"offset":870,"bytes":92869},{"epoch":1458151461,"offset":680,"bytes":1174353},{"epoch":1458151462,"offset":521,"bytes":123516},{"epoch":1458151463,"offset":951,"bytes":22847},{"epoch":1458151464,"offset":933,"bytes":381256},{"epoch":1458151465,"offset":650,"bytes":953885},{"epoch":1458151466,"offset":396,"bytes":368153},{"epoch":1458151467,"offset":602,"bytes":544992},{"epoch":1458151468,"offset":543,"bytes":147620},{"epoch":1458151469,"offset":201,"bytes":1125031},{"epoch":1458151470,"offset":378,"bytes":631381},{"epoch":1458151471,"offset":876,"bytes":197272},{"epoch":1458151472,"offset":381,"bytes":865959},{"epoch":1458151473,"offset":147,"bytes":970016},{"epoch":1458151474,"offset":133,"bytes":740507},{"epoch":1458151475,"offset":854,"bytes":584625},{"epoch":1458151476,"offset":364,"bytes":95346}]`)

	type JData struct {
		Epoch  int32
		Offset int32
		Bytes  int32
	}

	//Decode JSON using json unmarshal

	//placeholder array variable for decoded data
	dataDecoded := []JData{}

	err := json.Unmarshal(sampleData, &dataDecoded)
	if err != nil {
		panic(err)
	}
	//fmt.Println(dataDecoded)

	//Perform calculations

	var peakBytes int32
	var peakTime int32
	var totalBytes float32
	var samplecount float32

	const BpmsTobps = 8000
	const BpmsToMbps = 0.008

	samplecount = float32(len(dataDecoded))

	for _, v := range dataDecoded {
		totalBytes += float32(v.Bytes)
		if v.Bytes > peakBytes {
			peakBytes = v.Bytes
			peakTime = v.Epoch
		}
	}

	//Calculations
	avgBytes := float32(totalBytes / samplecount)
	fmt.Println("Sample Count: ", samplecount)
	fmt.Println("Peak Max Bytes Observed: ", peakBytes)
	fmt.Println("Was Observed at this ms epoch: ", peakTime)
	fmt.Printf("Peak microburst in bps: %.3f \n", float32(float32(peakBytes)*BpmsTobps))
	fmt.Printf("Peak microburst in Mbps: %.3f \n", float32(float32(peakBytes)*BpmsToMbps))
	fmt.Printf("Mean Bytes Observed: %.3f \n", avgBytes)
	fmt.Println("Total Bytes Observed: ", totalBytes)

}
