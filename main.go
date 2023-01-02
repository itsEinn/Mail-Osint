// Mail Osint by Ein v0.1
// Author: theinn#0044
// Server: discord.gg/atv44

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"main/modules"
	"strconv"

	"os"

	"github.com/dimiro1/banner"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
)

func init() {
	templ := `{{ .Title "MAIL OSINT" "" 2 }}
   {{ .AnsiColor.BrightWhite }}v0.1{{ .AnsiColor.Default }}
   {{ .AnsiColor.BrightCyan }}discord - theinn#0044{{ .AnsiColor.Default }}
   Tarih: {{ .Now "Monday, 2 Jan 2006" }}`

	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	println()
}

func help_menu() {
	data := [][]string{
		{"-e", "Hedef mail belirleyin", "Gerekli"},
		{"-verify", "Hedef mail adresini doğrulayın", "Değil"},
		{"-social", "Hedef mail için sosyal medya taraması", "Değil"},
		{"-relateds", "Hedef mail adresi ile ilgili mailler ve etki alanlarını bulun", "Değil"},
		{"-leaks", "Hedef mail için şifre sızıntılarını bulun", "Değil"},
		{"-dumps", "Hedef mail için pastebin dökümlerini arayın", "Değil"},
		{"-domain", "Mail adresinin domaini hakkında daha fazla bilgi", "Değil"},
		{"-o", "txt dosyasının çıkacağı dizin", "Değil"},
		{"-v", "Versiyonu gösterir", "Değil"},
		{"-h", "Yardım", "Değil"},
		{"-all", "Hepsini kullan", "Değil"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Komutlar", "Açıklama", "Gereklimi?"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	color.Yellow("Örnek: go run main.go -e ornek@domain.com -all")
}

func verifyPrint(verifyData modules.VerifyStruct, emailRepData modules.EmailRepStruct, email string) {
	if verifyData.IsVerified {
		fmt.Println(email+" =>", color.GreenString("Verified \u2714"))
		outputText += email + " => Verified \n"
	} else {
		fmt.Println(email+" =>", color.RedString("Not Verified \u2718"))
		outputText += email + " => Not Verified \n"
	}
	if verifyData.IsDisposable {
		fmt.Println(email+" =>", color.RedString("Disposable \u2718"))
		outputText += email + " => Disposable \n"
	} else {
		fmt.Println(email+" =>", color.GreenString("Not Disposable \u2714"))
		outputText += email + " => Not Disposable \n"
	}

	if modules.GetAPIKey("EmailRep.io API Key") != "" {
		fmt.Println("\nEmailRep Data for", color.WhiteString(email))
		outputText += "\nEmailRep Data for " + email + "\n"
		fmt.Println("|- Reputation:", color.YellowString(emailRepData.Reputation))
		outputText += "|- Reputation: " + emailRepData.Reputation + "\n"
		fmt.Println("|- Kara Liste:", color.WhiteString(strconv.FormatBool(emailRepData.Details.Blacklisted)))
		outputText += "|- Kara Liste: " + strconv.FormatBool(emailRepData.Details.Blacklisted) + "\n"
		fmt.Println("|- Kötü amaçlı aktivite:", color.WhiteString(strconv.FormatBool(emailRepData.Details.MaliciousActivity)))
		outputText += "|- Kötü amaçlı aktivite: " + strconv.FormatBool(emailRepData.Details.MaliciousActivity) + "\n"
		fmt.Println("|- Kimlik Sızıntısı:", color.WhiteString(strconv.FormatBool(emailRepData.Details.CredentialsLeaked)))
		outputText += "|- Kimlik Sızıntısı: " + strconv.FormatBool(emailRepData.Details.CredentialsLeaked) + "\n"
		fmt.Println("|- İlk Görülme:", color.YellowString(emailRepData.Details.FirstSeen))
		outputText += "|- İlk Görülme: " + emailRepData.Details.FirstSeen + "\n"
		fmt.Println("|- Son Görülme:", color.YellowString(emailRepData.Details.LastSeen))
		outputText += "|- Son Görülme: " + emailRepData.Details.LastSeen + "\n"
		fmt.Println("|- Oluşturulma Tarihi:", color.WhiteString(strconv.Itoa(emailRepData.Details.DaysSinceDomainCreation)))
		outputText += "|- Oluşturulma Tarihi: " + strconv.Itoa(emailRepData.Details.DaysSinceDomainCreation) + "\n"
		fmt.Println("|- Spam:", color.WhiteString(strconv.FormatBool(emailRepData.Details.Spam)))
		outputText += "|- Spam: " + strconv.FormatBool(emailRepData.Details.Spam) + "\n"
		fmt.Println("|- Ücretsiz Sağlayıcı:", color.WhiteString(strconv.FormatBool(emailRepData.Details.FreeProvider)))
		outputText += "|- Ücretsiz Sağlayıcı: " + strconv.FormatBool(emailRepData.Details.FreeProvider) + "\n"
		fmt.Println("|- Gönderilebilir:", color.WhiteString(strconv.FormatBool(emailRepData.Details.Deliverable)))
		outputText += "|- Gönderilebilir: " + strconv.FormatBool(emailRepData.Details.Deliverable) + "\n"
		fmt.Println("|- Geçerli MX:", color.WhiteString(strconv.FormatBool(emailRepData.Details.ValidMx)))
		outputText += "|- Geçerli MX: " + strconv.FormatBool(emailRepData.Details.ValidMx) + "\n"
	} else {
		color.Red("EmailRep.io API key belirlenmemiş!")
		outputText += "EmailRep.io API key belirlenmemiş!\n"
	}
}

func socialPrint(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	outputText += "Sosyal medya arama sonuçları: \n"
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println("|- "+scanner.Text(), color.GreenString("\u2714"))
		outputText += "|- " + scanner.Text() + "\n"

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	e := os.Remove(filename)
	if e != nil {
		log.Fatal(e)
	}
}

func relatedPrint(relEmails []string, relDomains []string, fromGoogle []string, hunterData modules.HunterStruct) {

	fmt.Println("İlgili mail adresleri:")
	outputText += "İlgili mail adresleri: \n"
	for _, v := range relEmails {
		fmt.Println("|- "+v, color.GreenString("\u2714"))
		outputText += "|- " + v + "\n"
	}
	if modules.GetAPIKey("Hunter.io API Key") != "" {
		for _, v := range hunterData.Data.Emails {
			fmt.Println("|- "+v.Value, color.GreenString("\u2714"))
			outputText += "|- " + v.Value + "\n"
		}
	}
	println("")
	fmt.Println("İlgili Domainler:")
	outputText += "İlgili Domainler: \n"
	for _, v := range relDomains {
		fmt.Println("|- "+v, color.GreenString("\u2714"))
		outputText += "|- " + v + "\n"
	}
	for _, v := range fromGoogle {
		fmt.Println("|- "+v, color.GreenString("\u2714"))
		outputText += "|- " + v + "\n"
	}
}

func leakPrint(breachData modules.BreachDirectoryStruct, intelxData []string) {
	fmt.Println("Şifre sızıntıları:")
	outputText += "Şifre sızıntıları: \n"
	if breachData.Success {
		for _, v := range breachData.Result {
			for _, w := range v.Sources {
				fmt.Println("|- " + w)
				outputText += "|- " + w + "\n"
			}
			if v.HasPassword {
				fmt.Println("|-- "+v.Password, color.GreenString("\u2714"))
				outputText += "|-- " + v.Password + "\n"
			}
		}
	} else {
		fmt.Println("|- Şifre sızıntısı bulunamadı")
		outputText += "|- Şifre sızıntısı bulunamadı \n"
	}
	println("\nIntelx şifre sızıntıları:")
	if len(intelxData) > 0 {
		for _, v := range intelxData {
			fmt.Println("|- "+v, color.GreenString("\u2714"))
			outputText += "|- " + v + "\n"
		}
	} else {
		color.Red("intelx dosyası yok!")
		outputText += "intelx dosyası yok! \n"
	}
}

func dumpPrint(binData []string) {
	fmt.Println("Pastebin ve Throwbin arama sonuçları:")
	outputText += ("Pastebin ve Throwbin arama sonuçları:")  
	for _, v := range binData {
		fmt.Println("|- "+v, color.GreenString("\u2714"))
		outputText += "|- " + v + "\n"
	}
}

func domainPrint(table *tablewriter.Table, ipapi modules.IPAPIStruct) {
	println("\nDomain Bilgisi:")
	outputText += "Domain Bilgisi: \n"
	fmt.Println("|- IP: "+ipapi.IP, color.GreenString("\u2714"))
	outputText += "|- IP: " + ipapi.IP + "\n"
	fmt.Println("|- Şehir: "+ipapi.City, color.GreenString("\u2714"))
	outputText += "|- Şehir: " + ipapi.City + "\n"
	fmt.Println("|- Bölge: "+ipapi.Region, color.GreenString("\u2714"))
	outputText += "|- Bölge: " + ipapi.Region + "\n"
	fmt.Println("|- Bölge Kodu: "+ipapi.RegionCode, color.GreenString("\u2714"))
	outputText += "|- Bölge Kodu: " + ipapi.RegionCode + "\n"
	fmt.Println("|- Ülke: "+ipapi.Country, color.GreenString("\u2714"))
	outputText += "|- Ülke: " + ipapi.Country + "\n"
	fmt.Println("|- Ülke Kodu: "+ipapi.CountryCode, color.GreenString("\u2714"))
	outputText += "|- Ülke Kodu: " + ipapi.CountryCode + "\n"
	fmt.Println("|- Ülke Adı: "+ipapi.CountryName, color.GreenString("\u2714"))
	outputText += "|- Ülke Adı: " + ipapi.CountryName + "\n"
	fmt.Println("|- Posta Kodu: "+ipapi.Postal, color.GreenString("\u2714"))
	outputText += "|- Posta Kodu: " + ipapi.Postal + "\n"
	fmt.Println("|- Saat Dilimi: "+ipapi.Timezone, color.GreenString("\u2714"))
	outputText += "|- Saat Dilimi: " + ipapi.Timezone + "\n"
	fmt.Println("|- Ülke Telefon Kodu: "+ipapi.CountryCallingCode, color.GreenString("\u2714"))
	outputText += "|- Ülke Telefon Kodu: " + ipapi.CountryCallingCode + "\n"
	fmt.Println("|- Para Birimi: "+ipapi.Currency, color.GreenString("\u2714"))
	outputText += "|- Para Birimi: " + ipapi.Currency + "\n"
	fmt.Println("|- Organizasyon: "+ipapi.Org, color.GreenString("\u2714"))
	outputText += "|- Organizasyon: " + ipapi.Org + "\n"

	println("\nDNS Kayıtları:")
	table.Render()
}

var outputText string = ""

func main() {
	var email *string = flag.String("e", "", "Mail belirle")
	var verify *bool = flag.Bool("verify", false, "Hedef mail adresini doğrulayın")
	var social_accounts *bool = flag.Bool("social", false, "Maile kayıtlı olan hesapları bulma")
	var relateds *bool = flag.Bool("relateds", false, "Domainden ilgili mailleri ve domainleri bulma")
	var leaks *bool = flag.Bool("leaks", false, "Mailden şifre sızıntılarını bulma")
	var dumps *bool = flag.Bool("dumps", false, "Mailden pastebin dökümlerini bulma")
	var domain *bool = flag.Bool("domain", false, "Domain hakkında daha fazla bilgi")
	var output *bool = flag.Bool("o", false, "txt dosyasının çıkacağı dizim")
	var version *bool = flag.Bool("v", false, "Versiyon")
	var help *bool = flag.Bool("h", false, "Yardım")
	var all *bool = flag.Bool("all", false, "Tüm özellikler")
	flag.Parse()
	println("")
	if len(*email) == 0 {
		help_menu()
		os.Exit(0)
	} else if *help {
		help_menu()
		os.Exit(0)
	} else if *version {
		color.White("version: 0.1")
		os.Exit(0)
	} else if *all {

		var bar = progressbar.Default(100, "osinting")
		bar.Add(5)
		var verifyData = modules.VerifyEmail(*email)
		bar.Add(7)
		var emailRepData = modules.EmailRep(*email)
		bar.Add(8)
		modules.Runner(*email, "SocialScan")
		bar.Add(8)
		modules.Runner(*email, "Holehe")
		var relEmails = modules.RelatedEmails(*email)
		bar.Add(8)
		var relDomains = modules.RelatedDomains(*email)
		bar.Add(8)
		var fromGoogle = modules.Related_domains_from_google(*email)
		bar.Add(8)
		var hunterData = modules.Hunter(*email)
		bar.Add(8)
		var breachData = modules.BreachDirectory(*email)
		bar.Add(8)
		var binData = modules.BinSearch(*email)
		bar.Add(8)
		var ipapi = modules.IPAPI(*email)
		bar.Add(8)
		var table = modules.DNS_lookup(*email)
		bar.Add(8)
		var intelxData = modules.Intelx(*email)
		bar.Finish()
		verifyPrint(verifyData, emailRepData, *email)
		socialPrint("socialscantempresult.txt")
		socialPrint("holehetempresult.txt")
		relatedPrint(relEmails, relDomains, fromGoogle, hunterData)
		leakPrint(breachData, intelxData)
		dumpPrint(binData)
		domainPrint(table, ipapi)
		if *output {
			var filename = modules.FileWriter(*email, outputText)
			color.Green("\nÇıkış Dosyası: " + filename)
		}
		os.Exit(0)
	} else {
		if *verify {
			verifyPrint(modules.VerifyEmail(*email), modules.EmailRep(*email), *email)
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nÇıkış Dosyası: " + filename)
			}
			os.Exit(0)
		}
		if *social_accounts {
			fmt.Println("Social media accounts opened with", color.WhiteString(*email))
			modules.Runner(*email, "SocialScan")
			modules.Runner(*email, "Holehe")
			socialPrint("socialscantempresult.txt")
			socialPrint("holehetempresult.txt")
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nÇıkış Dosyası: " + filename)
			}
			os.Exit(0)
		}
		if *relateds {
			relatedPrint(modules.RelatedEmails(*email), modules.RelatedDomains(*email), modules.Related_domains_from_google(*email), modules.Hunter(*email))
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nÇıkış Dosyası: " + filename)
			}
			os.Exit(0)
		}
		if *leaks {
			leakPrint(modules.BreachDirectory(*email), modules.Intelx(*email))
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nÇıkış Dosyası: " + filename)
			}
			os.Exit(0)
		}
		if *dumps {
			dumpPrint(modules.BinSearch(*email))
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nÇıkış Dosyası: " + filename)
			}
		}
		if *domain {
			domainPrint(modules.DNS_lookup(*email), modules.IPAPI(*email))
			if *output {
				var filename = modules.FileWriter(*email, outputText)
				color.Green("\nÇıkış Dosyası: " + filename)
			}
			os.Exit(0)
		}
	}

}
