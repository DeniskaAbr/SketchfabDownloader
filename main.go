package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

/*

https://sketchfab.com/i/search?
	count (max 24)
	q ( строка поиска )
 	categories ("animals-pets", "architecture", "art-abstract", "cars-vehicles", "characters-creatures",
				"cultural-heritage-history", "electronics-gadgets", "fashion-style", "food-drink", "furniture-home",
				"music", "nature-plants", "news-politics", "people", "places-travel", "science-technology",
				"sports-fitness", "weapons-military")
	in_store ( null 1 )
	processing_status ( succeeded ) //TODO: check what is
	type (models)
cursor ( cursor random)

// filters
//Features
	downloadable (null 1)
	animated ( null 1 )
	staffpicked ( null 1 )
	sound (null 1 )

//Sort by
	sort_by ( -relevance -likeCount -viewCount -publishedAt )

//Date
	date ( none(all) 1(day) 7(week) 31(month) )

	restricted ( 1(show restricted)

*/

/*

random UID

f4af9047bbc7452b9916220ba85611a7
c2952fe730eb4b9ab4f6af33c9f7a64b
dff2b08afc874d4997c5e5c47ef158a8
17ce9e6e8a904045aba90cebe113a852

TODO: Вывод коллекциями
previous: https://sketchfab.com/i/collections?count=24&cursor=cj0xJnA9Ng%3D%3D&date=7&restricted=1&sort_by=-subscriberCount

previous: https://sketchfab.com/i/collections/18d0504b692c4faab2f36a34335a30cc/models?cursor=cj0xJnA9MjAxNy0wNS0xNysxMiUzQTA4JTNBMTMuNjczNTQ2&restricted=1&sort_by=-collectedAt
*/

/*
TODO: Поиск через параметры в пунктах меню в магазине
https://sketchfab.com/store/categories All categories
https://sketchfab.com/store/3d-models?sort_by=-orderCount Best selling
https://sketchfab.com/store/3d-models?animated=1 Animatedscra
https://sketchfab.com/store/3d-models?pbr=1 PBR
https://sketchfab.com/store/3d-models?max_face_count=10000&min_face_count=0 Low poly
https://sketchfab.com/store/3d-models?min_face_count=100000&ref=header High poly
https://sketchfab.com/store/3d-models?file_formats=stl&ref=header 3D Printable
https://sketchfab.com/store/3d-models?q=scan&ref=header&sort_by=-orderCount 3DScan

*/

// пауза в секундах между реквестами
var ReqTime = 5 * time.Second // 60 * time.Second 60 секунд
var DownloadTime = 0 * time.Second
var MongoServer = "192.168.0.34:27017"

var ParserWG sync.WaitGroup
var ScraperWG sync.WaitGroup

var DEFAULT_USER_AGENT string
var FILESTORAGE_PATH string = "f:/FILESTORAGE"

const DATABASE = "test"
const COLLECTION = "sketchTest"

const BaseURL = "https://sketchfab.com/i/search?"

var Finder = finder

var Categories = [18]string{"animals-pets", "architecture", "art-abstract", "cars-vehicles", "characters-creatures", "cultural-heritage-history", "electronics-gadgets", "fashion-style", "food-drink", "furniture-home", "music", "nature-plants", "news-politics", "people", "places-travel", "science-technology", "sports-fitness", "weapons-military"}

// экземпляр соединения с монгой
var MongoClient *mongo.Client

var httpClient http.Client = http.Client{
	// ждем ответ указанное время если его не будет то будет ошибка
	Timeout: time.Second * 24000, // Timeout: time.Second * 2, // Maximum of 2 secs
}

func init() {

	// выбираем случайно юзерагент из списка, будет использован для запросов
	DEFAULT_USER_AGENT = USER_AGENTS[rand.Intn(len(USER_AGENTS))]
	// в качестве сида рандома берем текущее время
	rand.Seed(time.Now().UnixNano())

	// Нужно при работе через прокси
	// os.Setenv("HTTP_PROXY", "socks5://127.0.0.1:9050")

	// инициализирован ОРМ для монги
	MongoClient, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://" + MongoServer))

	// Создаем соединение
	err := MongoClient.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// Проверка соединения
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

}

func main() {
	// Scrape("https://sketchfab.com/3d-models/lion-statue-optimized-62662d27c94d4994b2479b8de3a66ca7", "")
	// scrap("https://sketchfab.com/3d-models/car-police-car-concept-29c76ea294264212b0598f358fbce111", "29c76ea294264212b0598f358fbce111")
	// scrap("https://sketchfab.com/3d-models/future-car-hliAt7xc6YU2XjXh5G4lzmQwFmR", "hliAt7xc6YU2XjXh5G4lzmQwFmR")
	// scrap("https://sketchfab.com/3d-models/cartoon-prototype-car-88b89d3074cb4946a353ab990d1ff6a2", "88b89d3074cb4946a353ab990d1ff6a2")
	// scrap("https://sketchfab.com/3d-models/old-rusty-car-95baa20ebc5d4d2e869f0b549be838fe", "95baa20ebc5d4d2e869f0b549be838fe")
	// scrap("https://sketchfab.com/3d-models/car-bus-bagel-concept-f1d505c90f5a474783cc6c2c74fd7a6b", "f1d505c90f5a474783cc6c2c74fd7a6b")

	// parser()


	for i := 0; i < 3000; i++ {

		a, b, err := RandomDocument()

		fmt.Println("-----")
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(err)
		fmt.Println("-----")

		if err == nil {
			Scrape(a, b)
		}
	}


	// parser()

}

func parser() {
	url, err := makeURL()
	Check(err)
	fmt.Println(url)

	err = get(url, "", "1212", "", "")
	Check(err)
	ParserWG.Wait()
	fmt.Println("Main: Parser completed with no errors")
}

func scrap(url string, uid string) {
	Scrape(url, uid)
	fmt.Println("Main: Waiting for scrapper workers to finish")
	ScraperWG.Wait()
	fmt.Println("Main: Scrapper completed with no errors")

}

// получение номера не скачанного проекта
func GetUIDFromMongo() {

}

func makeURL() (string, error) {

	var wordsPartsString = "a,ac,ad,af,ag,al,an,ap,as,at,a-,an-,ab,abs,-able,-ible,acer,acid,acri,act,ag,acu,-acy,-cy,-ade,aer,aero,ag,agi,ig,act,-age,agri,agro,-al,-al,-ial,-ical,alb,albo,ali,allo,alter,alt,am,ami,amor,ambi,ambul,-an,ana,ano,-ance,-ence,-ancy,-ency,andr,andro,ang,anim,ann,annu,enni,-ant,-ent,-ant,-ent,-ient,ante,anthrop,anti,ant,anti,antico,apo,ap,aph,aqu,-ar,-ary,arch,-ard,-art,aster,astr,-ate,-ate,-ate,-ation,auc,aug,aut,aud,audi,aur,aus,aug,auc,aut,auto,bar,be,belli,bene,bi,bine,bibl,bibli,biblio,bio,bi,brev,cad,cap,cas,ceiv,cept,capt,cid,cip,cad,cas,-cade,calor,capit,capt,carn,cat,cata,cath,caus,caut,cause,cuse,cus,ceas,ced,cede,ceed,cess,cent,centr,centri,chrom,chron,cide,cis,cise,circum,cit,civ,clam,claim,clin,clud,clusclaus,co,cog,col,coll,con,com,cor,cogn,gnos,com,con,contr,contra,counter,cord,cor,cardi,corp,cort,cosm,cour,cur,curr,curs,crat,cracy,cre,cresc,cret,crease,crea,cred,cresc,cret,crease,cru,crit,cur,curs,cura,cycl,cyclo,de-,dec,deca,dec,dign,dei,div,dem,demo,dent,dont,derm,di-,dy-,dia,dic,dict,dit,dis,dif,dit,doc,doct,domin,don,dorm,dox,-drome,duc,duct,dura,dynam,dys-,e-,ec-,eco-,ecto-,-ed,-ed,-en,-en,en-,em-,-ence,-ency,end-,epi-,equi-,-er,-ier,-er,-or,-er,-or,erg,-ery,-es,-ies,-es,-ies,-ess,-est,-iest,ev-,et-,ex-,exter-,extra-,extro-,fa,fess,fac,fact,fec,fect,fic,fas,fea,fall,fals,femto,fer,fic,feign,fain,fit,feat,fid,fid,fide,feder,fig,fila,fili,fin,fix,flex,flect,flict,flu,fluc,fluv,flux,-fold,for,fore,forc,fort,form,fract,frag,frai,fuge,-ful,-ful,fuse,-fy,gam,gastr,gastro,gen,gen,geo,germ,gest,giga,gin,gloss,glot,glu,glo,gor,grad,gress,gree,graph,gram,graf,grat,grav,greg,hale,heal,helio,hema,hemo,her,here,hes,hetero,hex,ses,sex,homo,hum,human,hydr,hydra,hydro,hyper,hypn,-ia,-ian,an,-iatry,-ic,-ic,ics,-ice,-ify,ignis,-ile,in,im,in,im,il,ir,infra,-ing,-ing,-ing,inter,intra,intro,-ion,-ish,-ism,-ist,-ite,-ity,ty,-ive,-ive,-ative,-itive,-ize,jac,ject,join,junct,judice,jug,junct,just,juven,labor,lau,lav,lot,lut,lect,leg,lig,leg,-less,levi,lex,leag,leg,liber,liver,lide,liter,loc,loco,log,logo,ology,loqu,locut,luc,lum,lun,lus,lust,lude,-ly,macr-,macer,magn,main,mal,man,manu,mand,mania,mar,mari,mer,matri,medi,mega,mem,ment,-ment,meso,meta,meter,metr,micro,migra,mill,kilo,milli,min,mis,mit,miss,mob,mov,mot,mon,mono,mor,mort,morph,multi,nano,nasc,nat,gnant,nai,nat,nasc,neo,-ness,neur,nom,nom,nym,nomen,nomin,non,non,nov,nox,noc,numer,numisma,ob,oc,of,op,oct,oligo,omni,onym,oper,-or,ortho,-ory,-ous,-eous,-ose,-ious,over,pac,pair,pare,paleo,pan,para,pat,pass,path,pater,patr,path,pathy,ped,pod,pedo,pel,puls,pend,pens,pond,per,peri,phage,phan,phas,phen,fan,phant,fant,phe,phil,phlegma,phobia,phobos,phon,phot,photo,pico,pict,plac,plais,pli,ply,plore,plu,plur,plus,pneuma,pneumon,pod,poli,poly,pon,pos,pound,pop,port,portion,post,pot,pre,pur,prehendere,prin,prim,prime,pro,proto,psych,punct,pute,quat,quad,quint,penta,quip,quir,quis,quest,quer,re,reg,recti,retro,ri,ridi,risi,rog,roga,rupt,sacr,sanc,secr,salv,salu,sanct,sat,satis,sci,scio,scientia,scope,scrib,script,se,sect,sec,sed,sess,sid,semi,sen,scen,sent,sens,sept,sequ,secu,sue,serv,-ship,sign,signi,simil,simul,sist,sta,stit,soci,sol,solus,solv,solu,solut,somn,soph,spec,spect,spi,spic,sper,sphere,spir,stand,stant,stab,stat,stan,sti,sta,st,stead,-ster,strain,strict,string,stige,stru,struct,stroy,stry,sub,suc,suf,sup,sur,sus,sume,sump,super,supra,syn,sym,tact,tang,tag,tig,ting,tain,ten,tent,tin,tect,teg,tele,tem,tempo,ten,tin,tain,tend,tent,tens,tera,term,terr,terra,test,the,theo,therm,thesis,thet,tire,tom,tor,tors,tort,tox,tract,tra,trai,treat,trans,tri,trib,tribute,turbo,typ,ultima,umber,umbraticum,un,uni,-ure,vac,vade,vale,vali,valu,veh,vect,ven,vent,ver,veri,verb,verv,vert,vers,vi,vic,vicis,vict,vinc,vid,vis,viv,vita,vivi,voc,voke,vol,volcan,volv,volt,vol,vor,-ward,-wise,with,-y,-y,zo"

	wordsPartsSlice := strings.Split(wordsPartsString, ",")

	for _, word := range wordsPartsSlice {
		b_word := []byte(word)

		if matched, _ := regexp.Match("-[a-z]", b_word); matched == true {
			fmt.Println("suffix")
		} else if matched, _ = regexp.Match("[a-z]-", b_word); matched == true {
			fmt.Println("prefix")
		}
		fmt.Println("root")
	}

	// var currentCursor string
	// var nextCursor string

	// constant params
	var url string = BaseURL
	Finder["count"] = "24"           // one page resultats
	Finder["cursor"] = ""            // link identity
	Finder["type"] = "models"        //scene type
	Finder["restricted"] = "1"       // sb_allows_age_restricted	"true" 18+
	Finder["sort_by"] = "-likeCount" // -relevance -likeCount -viewCount -publishedAt

	// nonstatic params
	Finder["q"] = ""       // find
	Finder["categories"] = Categories[1] // категория выборки" 17
	Finder["in_store"] = "1"  // в магазине

	params := getParams(Finder)

	if params != "" {
		url = url + params
	} else {
		return url, errors.New("нет параметров для поиска")
	}

	return url, nil
}

// формирует строку параметров для ссылки
func getParams(m map[string]string) string {

	params := ""

	for k, v := range m {
		if v != "" {
			params = params + "&" + k + "=" + v
		}
	}
	return params
}

// воркер парсера JSON
func worker(wg *sync.WaitGroup, url string, cursorNext string, cursorPrevious string, nextURL string, previousURL string) {
	defer wg.Done()

	fmt.Printf("Worker started: %s \n", cursorNext)
	var err = get(url, cursorNext, cursorPrevious, nextURL, previousURL)
	Check(err)
	fmt.Printf("Worker finished: %s \n", cursorNext)
}

// запрашивает JSON с частью данных
func get(url string, cursorNext string, cursorPrevious string, nextURL string, previousURL string) error {

	// получам ссылку, пробуем запросить джисон
	// получаем джисон
	// проверяем последняя ли это часть данных
	// анмаршалим пишем в MongoDB

	// отладочный вывод в консоль

	// пауза добавлена для паузы между запросами
	time.Sleep(ReqTime) // time.Sleep(60 * time.Second) // 60 секунд

	if len(url) == 0 {
		return errors.New("ничего не передано, прерываем программу")
	}

	// формируем параметры запроса
	req, err := http.NewRequest(http.MethodGet, url, nil)
	// проверяем на ошибку
	Check(err)
	// устанавливаем хидер для запроса
	req.Header.Set("User-Agent", DEFAULT_USER_AGENT)
	req.Header.Set("Host", "")
	cookie := http.Cookie{Name: "sb_allows_age_restricted", Value: "true"}
	req.AddCookie(&cookie)

	//req.Header.Set( "Connection" , "keep-alive" )
	//req.Header.Set( "Upgrade-Insecure-Requests", "1" )
	//req.Header.Set("Accept-Encoding", "gzip, deflate")

	// fmt.Println("Запрашиваем данные по URL")
	// запрашиваем JSON с сайта
	res, getErr := httpClient.Do(req)
	// проверяем на ошибки
	Check(getErr)

	if res.StatusCode == 200 {

		// подготавливаем данные к анмаршаллингу
		var body bytes.Buffer
		_, err = io.Copy(&body, res.Body)

		if err != nil {
			// fmt.Println("nil")
		}

		// fmt.Println(string(body.Bytes()))

		// сопоставляем со структурой
		dat := DataPart{}
		if err := json.Unmarshal(body.Bytes(), &dat); err != nil {
			panic(err)
		}

		// fmt.Println(dat)
		// fmt.Println("отдаем на запись в БД")

		toMongo(dat)

		if (cursorPrevious != dat.Cursors.Previous || cursorNext != dat.Cursors.Next) && len(dat.Results) > 0 {

			if len(dat.Cursors.Next) != 0 && len(dat.Cursors.Next) != 0 {
				// fmt.Println("Main: Starting worker")
				ParserWG.Add(1)
				go worker(&ParserWG, dat.Next, dat.Cursors.Next, dat.Cursors.Previous, dat.Next, dat.Previous)
			}
		} else {
			return nil
		}

	} else if res.StatusCode == http.StatusTooManyRequests {
		return errors.New("нет параметров для поиска")

	} else {
		return errors.New("нет параметров для поиска")
	}

	return nil
}

// записывает данные о работах в базу данных
func toMongo(dataPart DataPart) {

	var result bson.M

	var collection = MongoClient.Database(DATABASE).Collection(COLLECTION)

	for index := range dataPart.Results {

		var uid = dataPart.Results[index].UID

		var data = dataPart.Results[index]

		err := collection.FindOne(context.TODO(), bson.D{{"uid", uid}}).Decode(&result)

		if err != nil { // ErrNoDocuments means that the filter did not match any documents in the collection
			if err == mongo.ErrNoDocuments {
				// если нет совпадений и документ не найден
				fmt.Println("документ с " + uid + " не найден, будем добавлять")

				// fmt.Println(dat.Data[index].User.Full_name)
				// fmt.Println(index)
				insertResult, err := collection.InsertOne(context.TODO(), data)

				Check(err)

				var filter = bson.D{{"uid", uid}}

				update := bson.D{
					{"$set", bson.D{{"checked", false}}}}

				_, err = collection.UpdateOne(context.TODO(), filter, update)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("Вставлен один документ с _id: ", insertResult.InsertedID)
				//return
			}
		}
	}
}

func RandomDocument() (string, string, error) {

	var res Result

	// ищем случайную запись у которой checked false, downloaded none
	var collection = MongoClient.Database(DATABASE).Collection(COLLECTION)
	var result bson.M
	// err := collection.FindOne(context.TODO(), bson.D{{"checked", false}, {"downloaded", bson.D{{"$exists", false}}}}).Decode(&result)

	err := collection.FindOne(context.TODO(),
		bson.D{
		{"checked", false},

		{"$or", []interface{}{
			bson.D{{"downloaded", bson.D{{"$exists", false}}}},
			bson.D{{"downloaded", false}}}},
		}).Decode(&result)


	if err != nil {
		return "", "", errors.New("error")
	}

	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &res)

	fmt.Println("выводим на экран ")
	/*	fmt.Println(s.Icons.Model3d)*/
	fmt.Println(res.UID)

	return res.Embedurl, res.UID, nil
}
