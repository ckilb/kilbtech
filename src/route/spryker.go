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
	TeaserImage     string
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
			TeaserImage:     "antelope",
			Projects: []dto.Project{
				{
					LogoFileName: "marel.svg",
					Title:        "Marel",
					Position:     "Technical Lead, Customer Portal",
					Url:          "https://portal.marel.com",
					Description:  "As a <strong>Technical Lead</strong>, I was tasked with leading the implementation of a new customer portal platform using Spryker for one of the world's <strong>largest food processing companies</strong>. Thanks to this new portal, customers can now easily place, view, and manage their orders online. This marks a significant milestone in Marel's digital transformation journey.",
					LogoWidth:    1730,
					LogoHeight:   410,
				},
				{
					LogoFileName: "dept.svg",
					Title:        "DEPT",
					Position:     "Spryker Lead Developer, B2B e-commerce",
					Url:          "https://www.deptagency.com",
					Description:  "As a <strong>Spryker Lead Developer</strong>, I provided advisory and hands-on support to DEPT on <strong>several enterprise-level projects</strong>. In close collaboration with customers, I gathered, specified, and implemented requirements.\n\nDue to the complexity of these projects, I often had to break down extensive business requirements and problems into smaller components to effectively communicate them to the experienced Spryker software development team.",
					LogoWidth:    730,
					LogoHeight:   200,
				},
				{
					LogoFileName: "tomtailor.svg",
					Title:        "Tom Tailor",
					Position:     "Spryker Senior Developer & Consultant, B2C e-commerce platform",
					Url:          "https://www.tom-tailor.de",
					Description:  "As a Spryker senior developer and consultant, I played a key role in the expansion of the Tom Tailor online shop in close collaboration with my colleagues. In addition to my implementation work as a Spryker developer, I proactively initiated and created various proofs of concepts.\n\nI also provided advisory support to (junior) developers to help them implement requirements and expand their knowledge of Spryker.",
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
