package route

import (
	"ckilb/kilbtech/dto"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"text/template"
)

type HomeTemplateData struct {
	AssetCacheId    string
	Projects        []dto.Project
	MetaDescription string
	Headline        string
	SubHeadline     string
}

type Home struct {
	tpl *template.Template
}

func (r *Home) Path() string {
	return "/"
}

func (r *Home) Handler() http.Handler {
	assetCacheId := uuid.NewString()

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
		}

		err := r.tpl.Execute(w, &HomeTemplateData{
			AssetCacheId:    assetCacheId,
			MetaDescription: "With over 20 years of experience in software development, Christian Kilb has acquired a wealth of knowledge in technologies, strategies and leadership.",
			Headline:        "Christian Kilb",
			SubHeadline:     "Technical Lead | Software Architect | Developer",
			Projects: []dto.Project{
				{
					LogoFileName: "marel.svg",
					Title:        "Marel",
					Position:     "Technical Lead, Customer Portal",
					Url:          "https://portal.marel.com",
					Description:  "My responsibility as a <strong>Technical Lead</strong> was to create a new customer portal platform from scratch. Thanks to the new portal customers of one of the largest food processing companies worldwide are now able to place, view and manage their orders online. This is a big step towards digitalization for Marel.",
					LogoWidth:    1730,
					LogoHeight:   410,
				},
				{
					LogoFileName: "aboutyou.svg",
					Title:        "Legal You",
					Position:     "Product Lead, 3rd-party integration APIs",
					Url:          "https://aboutyou.tech",
					Description:  "As <strong>Product Lead</strong> for the ABOUT YOU Developer Center, my responsibility was to enable third-party apps to be able to connect to the shop ecosystem.\n\nHappily I had a team of six experienced software developers, with whom I coordinated the development of the platform, interfaces and software development kits.",
					LogoWidth:    1024,
					LogoHeight:   214,
				},
				{
					LogoFileName: "dept.svg",
					Title:        "DEPT",
					Position:     "Lead Developer, B2B e-commerce",
					Url:          "https://www.deptagency.com",
					Description:  "As a <strong>Lead Developer</strong>, I was able to support DEPT with my expertise in numerous enterprise-level projects in an advisory and hands-on manner.\nIn close exchange with the customers, I collected, specified and implemented requirements.\n\nThe complexity of the projects often made it necessary to break down extensive business requirements and problems into smaller components - in order to finally communicate them to the team of experienced software developers in an understandable way.",
					LogoWidth:    730,
					LogoHeight:   200,
				},
				{
					LogoFileName: "tomtailor.svg",
					Title:        "Tom Tailor",
					Position:     "Senior Developer & Consultant, B2C e-commerce platform",
					Url:          "https://www.tom-tailor.de",
					Description:  "As a <strong>senior developer and consultant</strong>, I supported the expansion of the Tom Tailor online shop in close cooperation with my colleagues.\nIn addition to my work as a software developer in the implementation, I also created various proofs of concepts, some of which I initiated proactively.\n\n(Junior) developers I was available in an advisory capacity to support them in the implementation of the requirements and further training of their knowledge, especially regarding S.O.L.I.D. coding principles.",
					LogoWidth:    601,
					LogoHeight:   59,
				},
				{
					LogoFileName: "kilb.svg",
					Title:        "Product Designer",
					Position:     "Owner",
					Url:          "https://store.shopware.com/en/kilb647502658205/kilb-product-designer-product-designer-for-shirts-cups-posters-cards...-plugin-version.html",
					Description:  "In addition to my work for my customers, I regularly expand my knowledge through my own projects. The Product Designer for Shopware and WooCommerce is particularly noteworthy. This is an extension that shop operators in medium-sized and enterprise companies can use to offer their customers an individual design of the products.",
					LogoWidth:    1544,
					LogoHeight:   553,
				},
			},
		})

		if err != nil {
			log.Println(fmt.Errorf("executing home template: %w", err))
		}

	})
}

func NewHome(tpl *template.Template) Route {
	return &Home{tpl: tpl}
}
