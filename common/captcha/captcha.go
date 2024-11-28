package captcha

import (
	"log"

	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/slide"
)

var SlideCapt slide.Captcha

// 初始化验证码 滑动验证码
func InitCaptcha() {
	builder := slide.NewBuilder(
		//slide.WithGenGraphNumber(2),
		slide.WithEnableGraphVerticalRandom(true),
	)

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	graphs, err := tiles.GetTiles()
	if err != nil {
		log.Fatalln(err)
	}

	var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
	for i := 0; i < len(graphs); i++ {
		graph := graphs[i]
		newGraphs = append(newGraphs, &slide.GraphImage{
			OverlayImage: graph.OverlayImage,
			MaskImage:    graph.MaskImage,
			ShadowImage:  graph.ShadowImage,
		})
	}

	// set resources
	builder.SetResources(
		slide.WithGraphImages(newGraphs),
		slide.WithBackgrounds(imgs),
	)

	SlideCapt = builder.Make()
}

//
//func main() {
//	captData, err := slideCapt.Generate()
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	dotData := captData.GetData()
//	if dotData == nil {
//		log.Fatalln(">>>>> generate err")
//	}
//	x := dotData.X
//	Y := dotData.Y
//	fmt.Println(x)
//	fmt.Println(Y)
//	dots, _ := json.Marshal(dotData)
//	fmt.Println(">>>>> ", string(dots))
//	base64 := captData.GetMasterImage().ToBase64()
//
//	fmt.Println(base64)
//	base641 := captData.GetTileImage().ToBase64()
//
//	fmt.Println(base641)
//	err = captData.GetMasterImage().SaveToFile("./.caches/master.jpg", option.QualityNone)
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = captData.GetTileImage().SaveToFile("./.caches/thumb.png")
//	if err != nil {
//		fmt.Println(err)
//	}
//}
