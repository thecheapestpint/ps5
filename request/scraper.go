package request

import (
	"fmt"
	"github.com/mxschmitt/playwright-go"
	"log"
)

type Scraper struct {
	Engine  *playwright.Playwright
	Browser playwright.Browser
	Page    playwright.Page
}

func New(url string) (*Scraper, error) {
	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	headless := true
	sandbox := false
	opts := playwright.BrowserTypeLaunchOptions{
		Headless:          &headless,
		ChromiumSandbox:   &sandbox,
	}
	browser, err := pw.Chromium.Launch(opts)

	if err != nil {
		return nil, fmt.Errorf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create page: %v", err)
	}

	if _, err = page.Goto(url); err != nil {
		fmt.Printf("could not goto: %v", err.Error())
	}

	return &Scraper{
		Engine:  pw,
		Browser: browser,
		Page:    page,
	}, nil
}

func (pw *Scraper) GetSelector(selector string) (playwright.ElementHandle, error) {
	entry, err := pw.Page.QuerySelector(selector)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (pw *Scraper) Stop() {
	pw.Browser.Close()
	pw.Engine.Stop()
}
