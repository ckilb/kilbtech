package route

import (
	"ckilb/kilbtech/dto"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"text/template"
)

type SprykerTemplateData struct {
	AssetCacheId    string
	Projects        []dto.Project
	MetaDescription string
	Headline        string
	SubHeadline     string
}

type Spryker struct {
	tpl *template.Template
}

func (r *Spryker) Path() string {
	return "/spryker-freelancer"
}

func (r *Spryker) Handler() http.Handler {
	assetCacheId := uuid.NewString()

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := r.tpl.Execute(w, &SprykerTemplateData{
			AssetCacheId:    assetCacheId,
			MetaDescription: "Christian Kilb has over 20 years of experience in web development and works with Spryker since 2017 as a freelancer, developer and consultant.",
			Headline:        "Spryker Freelancer &amp; Developer",
			SubHeadline:     "Spryker Developer & Consultant in Hamburg, Germany",
			Projects: []dto.Project{
				{
					LogoFileName: "marel.svg",
					Title:        "Marel",
					Position:     "Technical Lead, Customer Portal",
					Url:          "https://portal.marel.com",
					Description:  "As a <strong>Technical Lead</strong> I was asked to lead the implementation of a new customer portal platform in Spryker. Thanks to the new portal customers of one of the largest food processing companies worldwide are now able to place, view and manage their orders online. This is a big step towards digitalization for Marel.",
					LogoWidth:    1730,
					LogoHeight:   410,
				},
				{
					LogoFileName: "dept.svg",
					Title:        "DEPT",
					Position:     "Spryker Lead Developer, B2B e-commerce",
					Url:          "https://www.deptagency.com",
					Description:  "As a <strong>Spryker Lead Developer</strong>, I was asked to support DEPT with my expertise in several enterprise-level projects in an advisory and hands-on manner.\nIn close exchange with the customers, I collected, specified and implemented requirements.\n\nThe complexity of the projects often made it necessary to break down extensive business requirements and problems into smaller components - in order to finally communicate them to the team of experienced Spryker software developers in an understandable way.",
					LogoWidth:    730,
					LogoHeight:   200,
				},
				{
					LogoFileName: "tomtailor.svg",
					Title:        "Tom Tailor",
					Position:     "Spryker Senior Developer & Consultant, B2C e-commerce platform",
					Url:          "https://www.tom-tailor.de",
					Description:  "As a <strong>Spryker senior developer and consultant</strong>, I supported the expansion of the Tom Tailor online shop in close cooperation with my colleagues.\nIn addition to my work as a Spryker developer in the implementation, I also created various proofs of concepts, some of which I initiated proactively.\n\n(Junior) developers I was available in an advisory capacity to support them in the implementation of the requirements and further training of their knowledge in Spryker.",
					LogoWidth:    601,
					LogoHeight:   59,
				},
			},
		})

		if err != nil {
			log.Println(fmt.Errorf("executing spryker template: %w", err))
		}
	})
}

func NewSpryker(tpl *template.Template) Route {
	return &Spryker{tpl: tpl}
}
