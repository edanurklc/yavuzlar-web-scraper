package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/mbndr/figlet4go"
)

func main() {

	ascii := figlet4go.NewAsciiRender()
	err := ascii.LoadFont("./fonts")
	if err != nil {
		fmt.Println("Font yüklenemedi, varsayılan yazı stili kullanılıyor.")
	}

	options := figlet4go.NewRenderOptions()
	options.FontName = "larry3d"
	morRenk, _ := figlet4go.NewTrueColorFromHexString("B026FF")
	options.FontColor = []figlet4go.Color{morRenk}

	renderStr, err := ascii.RenderOpts("EDDY", options)
	if err == nil {
		fmt.Print(renderStr)
	} else {
		fmt.Println("\n=== EDDY WEB SCRAPER ===")
	}

	hideDate := flag.Bool("date", false, "Tarih bilgisini gizleyin.")
	hideDescription := flag.Bool("description", false, "Açıklamayı gizleyin.")
	flag.Parse()

	for {
		fmt.Println("YAPMAK İSTEDİĞİNİZ İŞLEMİ SEÇİNİZ:")
		fmt.Println("1- thehackernews.com verilerini çekin")
		fmt.Println("2- bleepingcomputer.com verilerini çekin")
		fmt.Println("3- hackread.com verilerini çekin")
		fmt.Println("4- Programı sonlandır")

		var secim int = 0
		fmt.Scanln(&secim)

		c := colly.NewCollector()
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("HATA: Siteye erişilemiyor.", err)
			fmt.Println("Hata Kodu:", r.StatusCode)
		})

		switch secim {
		case 1:
			file, err := os.OpenFile("thehackernews.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("Dosya oluşturulurken bir hata oluştu:", err)
				continue
			}

			c.OnHTML("div.body-post", func(e *colly.HTMLElement) {
				baslik := e.ChildText("h2.home-title")
				tarih := e.ChildText("span.h-datetime")
				aciklama := e.ChildText("div.home-desc")

				file.WriteString("Başlık: " + baslik + "\n")
				if !*hideDate {
					file.WriteString("Tarih: " + tarih + "\n")
				}
				if !*hideDescription {
					file.WriteString("Açıklama: " + aciklama + "\n")
				}
				file.WriteString("---------------------------------------\n")
			})

			c.OnRequest(func(r *colly.Request) {
				fmt.Println("The Hacker News sitesine erişim sağlanıyor.")
			})

			c.Visit("https://thehackernews.com")
			file.Close()
			fmt.Println("Dosya oluşturuldu.")

		case 2:
			file, err := os.OpenFile("bleepingcomputer.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("Dosya oluşturulurken bir hata oluştu:", err)
				continue
			}

			c.OnHTML("div.bc_latest_news_text", func(e *colly.HTMLElement) {
				baslik := e.ChildText("h4 a")
				tarih := e.ChildText("li.bc_news_date")
				aciklama := e.ChildText("p")

				if baslik != "" {
					file.WriteString("Başlık: " + baslik + "\n")
					if !*hideDate {
						file.WriteString("Tarih: " + tarih + "\n")
					}
					if !*hideDescription {
						file.WriteString("Açıklama: " + aciklama + "\n")
					}
					file.WriteString("---------------------------------------\n")
				}
			})

			c.OnRequest(func(r *colly.Request) {
				fmt.Println("Bleeping Computer sitesine erişim sağlanıyor.")
			})

			c.Visit("https://www.bleepingcomputer.com/")
			file.Close()
			fmt.Println("Dosya oluşturuldu.")

		case 3:
			file, err := os.OpenFile("hackread.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("Dosya oluşturulurken bir hata oluştu:", err)
				continue
			}

			c.OnHTML("article", func(e *colly.HTMLElement) {
				baslik := e.ChildText("h2 a")
				tarih := e.ChildText("time")
				aciklama := e.ChildText("p")

				if baslik != "" {
					file.WriteString("Başlık: " + baslik + "\n")
					if !*hideDate {
						file.WriteString("Tarih: " + tarih + "\n")
					}
					if !*hideDescription {
						file.WriteString("Açıklama: " + aciklama + "\n")
					}
					file.WriteString("---------------------------------------\n")
				}
			})

			c.OnRequest(func(r *colly.Request) {
				fmt.Println("Hackread sitesine erişim sağlanıyor.")
			})

			c.Visit("https://hackread.com/")
			file.Close()
			fmt.Println("Dosya oluşturuldu.")

		case 4:
			fmt.Println("Program Sonlandırılıyor.")
			os.Exit(0)

		default:
			fmt.Println("Geçersiz seçim yaptınız. Tekrar deneyin!")
		}
	}
}
